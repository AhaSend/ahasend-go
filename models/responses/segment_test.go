package responses

import (
	"encoding/json"
	"testing"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// segmentFixture is a ready segment with one condition.
const segmentFixture = `{
	"object":"segment",
	"id":"33333333-3333-4333-8333-333333333333",
	"created_at":"2026-10-20T10:00:00Z",
	"updated_at":"2026-10-20T11:00:00Z",
	"name":"Germany",
	"description":"",
	"definition":{"schema_version":1,"match":"all","conditions":[{"field":"attributes.country","op":"equal","value":"DE"}]},
	"timezone":"Europe/Berlin",
	"state":"ready",
	"failure_reason":null,
	"member_count":12,
	"last_evaluated_at":"2026-10-20T11:05:00Z"`

func TestSegmentDetailDecodesASetMarketableCount(t *testing.T) {
	payload := segmentFixture + `,"marketable_count":9}`

	var segment SegmentDetail
	require.NoError(t, json.Unmarshal([]byte(payload), &segment))

	assert.Equal(t, uuid.MustParse("33333333-3333-4333-8333-333333333333"), segment.ID)
	assert.Equal(t, common.SegmentStateReady, segment.State)
	assert.Equal(t, 12, segment.MemberCount)
	require.NotNil(t, segment.LastEvaluatedAt)
	require.NotNil(t, segment.MarketableCount)
	assert.Equal(t, 9, *segment.MarketableCount)
	require.Len(t, segment.Definition.Conditions, 1)
	require.NotNil(t, segment.Definition.Conditions[0].Condition)
	require.NotNil(t, segment.Definition.Conditions[0].Condition.Value)
	require.NotNil(t, segment.Definition.Conditions[0].Condition.Value.String)
	assert.Equal(t, "DE", *segment.Definition.Conditions[0].Condition.Value.String)

	encoded, err := json.Marshal(segment)
	require.NoError(t, err)
	assert.JSONEq(t, payload, string(encoded))
}

func TestSegmentDetailDecodesANullMarketableCount(t *testing.T) {
	var segment SegmentDetail
	require.NoError(t, json.Unmarshal([]byte(segmentFixture+`,"marketable_count":null}`), &segment))

	assert.Nil(t, segment.MarketableCount)
}

func TestSegmentDecodesAFailedState(t *testing.T) {
	payload := []byte(`{
		"object":"segment",
		"id":"33333333-3333-4333-8333-333333333333",
		"created_at":"2026-10-20T10:00:00Z",
		"updated_at":"2026-10-20T11:00:00Z",
		"name":"Germany",
		"description":"Customers in Germany",
		"definition":{"schema_version":1,"match":"any","conditions":[{"field":"email","op":"is_not_empty"}]},
		"timezone":"UTC",
		"state":"failed",
		"failure_reason":"The members could not be found in time.",
		"member_count":0,
		"last_evaluated_at":null
	}`)

	var segment Segment
	require.NoError(t, json.Unmarshal(payload, &segment))

	assert.Equal(t, common.SegmentStateFailed, segment.State)
	require.NotNil(t, segment.FailureReason)
	assert.Equal(t, "The members could not be found in time.", *segment.FailureReason)
	assert.Nil(t, segment.LastEvaluatedAt)
	assert.Equal(t, common.SegmentMatchAny, segment.Definition.Match)
	require.Len(t, segment.Definition.Conditions, 1)
	require.NotNil(t, segment.Definition.Conditions[0].Condition)
	assert.Nil(t, segment.Definition.Conditions[0].Condition.Value)

	encoded, err := json.Marshal(segment)
	require.NoError(t, err)
	assert.JSONEq(t, string(payload), string(encoded))
}

func TestSegmentPreviewDecodesANullCount(t *testing.T) {
	var preview SegmentPreview
	require.NoError(t, json.Unmarshal([]byte(`{"object":"segment_preview","count":null,"contacts":[{"contact_id":"11111111-1111-4111-8111-111111111111","email":"ann@example.com"}]}`), &preview))

	assert.Equal(t, "segment_preview", preview.Object)
	assert.Nil(t, preview.Count)
	require.Len(t, preview.Contacts, 1)
	assert.Equal(t, "ann@example.com", preview.Contacts[0].Email)
}
