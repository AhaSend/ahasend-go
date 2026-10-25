package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// segmentDefinitionExample is the SegmentDefinition example of the API
// specification.
const segmentDefinitionExample = `{
	"schema_version":1,
	"match":"all",
	"conditions":[
		{"field":"attributes.country","op":"equal","value":"DE"},
		{"field":"created_at","op":"after","value":"30 days ago"},
		{"field":"list","op":"in","value":"6f1c2a52-8f2e-4d8e-9a3b-1c2d3e4f5a6b"},
		{"match":"any","conditions":[
			{"field":"campaign","op":"opened","value":"9a2e7c41-0b3d-4f5e-8a6b-7c8d9e0f1a2b"},
			{"field":"any_campaign","op":"clicked","value":"14 days ago"},
			{"field":"campaign_count","op":"opened_at_least","value":{"count":3,"since":"90 days ago"}}
		]}
	]
}`

func TestSegmentDefinitionRoundTripsTheSpecExample(t *testing.T) {
	var definition SegmentDefinition
	require.NoError(t, json.Unmarshal([]byte(segmentDefinitionExample), &definition))

	assert.Equal(t, SegmentSchemaVersion, definition.SchemaVersion)
	assert.Equal(t, SegmentMatchAll, definition.Match)
	require.Len(t, definition.Conditions, 4)
	for _, rule := range definition.Conditions[:3] {
		require.NotNil(t, rule.Condition)
		assert.Nil(t, rule.Group)
	}
	group := definition.Conditions[3].Group
	require.NotNil(t, group)
	assert.Nil(t, definition.Conditions[3].Condition)
	assert.Equal(t, SegmentMatchAny, group.Match)
	require.Len(t, group.Conditions, 3)
	count := group.Conditions[2].Value
	require.NotNil(t, count)
	assert.Equal(t, &SegmentCampaignCount{Count: 3, Since: "90 days ago"}, count.Count)

	encoded, err := json.Marshal(definition)
	require.NoError(t, err)
	assert.JSONEq(t, segmentDefinitionExample, string(encoded))
}

func TestSegmentRuleRefusesBothOrNeitherPart(t *testing.T) {
	condition := &SegmentCondition{Field: SegmentFieldEmail, Op: SegmentOpIsNotEmpty}
	group := &SegmentGroup{Match: SegmentMatchAll, Conditions: []SegmentCondition{*condition}}

	for name, rule := range map[string]SegmentRule{
		"both":    {Condition: condition, Group: group},
		"neither": {},
	} {
		t.Run(name, func(t *testing.T) {
			_, err := json.Marshal(rule)
			assert.Error(t, err)
		})
	}
}

func TestSegmentValueRefusesNoFieldOrTwoFields(t *testing.T) {
	text := "DE"
	yes := true

	for name, value := range map[string]SegmentValue{
		"no field":   {},
		"two fields": {String: &text, Bool: &yes},
	} {
		t.Run(name, func(t *testing.T) {
			_, err := json.Marshal(value)
			assert.Error(t, err)
		})
	}
}

func TestSegmentValueRoundTrips(t *testing.T) {
	text := "30 days ago"
	whole := float64(3)
	decimal := 2.5
	yes := true
	no := false

	tests := []struct {
		name  string
		json  string
		value SegmentValue
	}{
		{name: "string", json: `"30 days ago"`, value: SegmentValue{String: &text}},
		{name: "whole number", json: `3`, value: SegmentValue{Number: &whole}},
		{name: "decimal", json: `2.5`, value: SegmentValue{Number: &decimal}},
		{name: "true", json: `true`, value: SegmentValue{Bool: &yes}},
		{name: "false", json: `false`, value: SegmentValue{Bool: &no}},
		{name: "count", json: `{"count":3,"since":"90 days ago"}`, value: SegmentValue{Count: &SegmentCampaignCount{Count: 3, Since: "90 days ago"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var decoded SegmentValue
			require.NoError(t, json.Unmarshal([]byte(tt.json), &decoded))
			assert.Equal(t, tt.value, decoded)

			encoded, err := json.Marshal(decoded)
			require.NoError(t, err)
			assert.JSONEq(t, tt.json, string(encoded))
		})
	}
}

func TestSegmentConditionWithoutAValueOmitsIt(t *testing.T) {
	encoded, err := json.Marshal(SegmentCondition{Field: SegmentFieldEmail, Op: SegmentOpIsEmpty})

	require.NoError(t, err)
	assert.JSONEq(t, `{"field":"email","op":"is_empty"}`, string(encoded))
}

func TestSegmentConditionDecodesANullValueAsNil(t *testing.T) {
	var condition SegmentCondition
	require.NoError(t, json.Unmarshal([]byte(`{"field":"email","op":"is_empty","value":null}`), &condition))

	assert.Nil(t, condition.Value)
}

func TestSegmentRuleRefusesNullEntry(t *testing.T) {
	var definition SegmentDefinition
	err := json.Unmarshal([]byte(`{"schema_version":1,"match":"all","conditions":[null]}`), &definition)
	require.Error(t, err)
}
