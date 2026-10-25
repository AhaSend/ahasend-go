package common

import (
	"bytes"
	"encoding/json"
	"errors"
)

// SegmentSchemaVersion is the version of the SegmentDefinition format. The
// API refuses any other version.
const SegmentSchemaVersion = 1

// Segment match modes: whether a contact must match all of a definition's or
// group's conditions, or any of them.
const (
	SegmentMatchAll = "all"
	SegmentMatchAny = "any"
)

// Segment states.
//
//   - SegmentStateBuilding: the members are being found for the current
//     definition. A campaign that uses the segment cannot be sent until it is
//     ready.
//   - SegmentStateReady: the members follow the current definition.
//   - SegmentStateFailed: the members could not be found, and the failure
//     reason says why. Any update of the segment, an empty one included,
//     builds it again.
const (
	SegmentStateBuilding = "building"
	SegmentStateReady    = "ready"
	SegmentStateFailed   = "failed"
)

// Segment condition fields that are not contact fields of the account. A
// contact field of the account is named "attributes.<key>".
//
//   - SegmentFieldEmail, SegmentFieldFirstName and SegmentFieldLastName are
//     text fields, and SegmentFieldCreatedAt, when the contact was created,
//     is a timestamp field. None of them takes the null operators.
//   - SegmentFieldList takes SegmentOpIn or SegmentOpNotIn and a list ID.
//     SegmentOpIn means a confirmed membership of that list.
//   - SegmentFieldCampaign takes a received, opened or clicked operator, or
//     its negation, and a campaign ID.
//   - SegmentFieldAnyCampaign takes a received, opened or clicked operator,
//     SegmentOpNotOpened or SegmentOpNotClicked, and a relative time such as
//     "14 days ago".
//   - SegmentFieldCampaignCount takes an at-least or fewer-than operator and
//     a SegmentCampaignCount value.
const (
	SegmentFieldEmail         = "email"
	SegmentFieldFirstName     = "first_name"
	SegmentFieldLastName      = "last_name"
	SegmentFieldCreatedAt     = "created_at"
	SegmentFieldList          = "list"
	SegmentFieldCampaign      = "campaign"
	SegmentFieldAnyCampaign   = "any_campaign"
	SegmentFieldCampaignCount = "campaign_count"
)

// Segment condition operators. Each applies to some fields only, and the API
// refuses an operator that does not apply to its field. A negated operator on
// a contact field also matches a contact that has no value for the field.
const (
	// Text fields take these.
	SegmentOpEqual         = "equal"
	SegmentOpNotEqual      = "not_equal"
	SegmentOpContains      = "contains"
	SegmentOpNotContains   = "not_contains"
	SegmentOpBeginsWith    = "begins_with"
	SegmentOpNotBeginsWith = "not_begins_with"
	SegmentOpEndsWith      = "ends_with"
	SegmentOpNotEndsWith   = "not_ends_with"
	// SegmentOpIsEmpty and SegmentOpIsNotEmpty take no value.
	SegmentOpIsEmpty    = "is_empty"
	SegmentOpIsNotEmpty = "is_not_empty"

	// Number fields take these, and SegmentOpEqual and SegmentOpNotEqual.
	SegmentOpLess           = "less"
	SegmentOpLessOrEqual    = "less_or_equal"
	SegmentOpGreater        = "greater"
	SegmentOpGreaterOrEqual = "greater_or_equal"

	// Date and timestamp fields take these, and SegmentOpEqual and
	// SegmentOpNotEqual, with a date such as "2024-06-15" or a relative time
	// such as "today" or "30 days ago". Timestamp fields also take a time
	// such as "2 hours from now".
	SegmentOpBefore     = "before"
	SegmentOpBeforeOrOn = "before_or_on"
	SegmentOpAfter      = "after"
	SegmentOpAfterOrOn  = "after_or_on"

	// Custom date, number, timestamp and boolean fields take these, with no
	// value.
	SegmentOpIsNull    = "is_null"
	SegmentOpIsNotNull = "is_not_null"

	// Boolean fields take these, with no value.
	SegmentOpIsTrue  = "is_true"
	SegmentOpIsFalse = "is_false"

	// Custom date fields take these, which read only the month and day.
	// SegmentOpMonthIs takes a month from 1 (January) to 12 (December).
	// SegmentOpAnniversaryWithin takes a number of days N from 0 to 365 and
	// matches when the month and day fall on today or one of the next N days
	// in the segment's time zone.
	SegmentOpMonthIs           = "month_is"
	SegmentOpAnniversaryWithin = "anniversary_within"

	// SegmentFieldList takes these, with a list ID.
	SegmentOpIn    = "in"
	SegmentOpNotIn = "not_in"

	// SegmentFieldCampaign takes these, with a campaign ID.
	// SegmentFieldAnyCampaign takes all but SegmentOpNotReceived, with a
	// relative time. SegmentOpNotOpened and SegmentOpNotClicked match only
	// contacts who received a campaign.
	SegmentOpReceived    = "received"
	SegmentOpNotReceived = "not_received"
	SegmentOpOpened      = "opened"
	SegmentOpNotOpened   = "not_opened"
	SegmentOpClicked     = "clicked"
	SegmentOpNotClicked  = "not_clicked"

	// SegmentFieldCampaignCount takes these, with a SegmentCampaignCount
	// value. The fewer-than operators match only a contact who received at
	// least one campaign sent in the window.
	SegmentOpReceivedAtLeast   = "received_at_least"
	SegmentOpReceivedFewerThan = "received_fewer_than"
	SegmentOpOpenedAtLeast     = "opened_at_least"
	SegmentOpOpenedFewerThan   = "opened_fewer_than"
	SegmentOpClickedAtLeast    = "clicked_at_least"
	SegmentOpClickedFewerThan  = "clicked_fewer_than"
)

// SegmentDefinition represents the rules of a segment. The root holds
// conditions and groups, and a group holds conditions only. A definition has
// at most 50 conditions in all, and at most 5 of them read campaigns. The API
// refuses a definition with 400 and a message that names each problem by its
// path, such as conditions[3].value.
type SegmentDefinition struct {
	// SchemaVersion is SegmentSchemaVersion.
	SchemaVersion int `json:"schema_version"`
	// Match is SegmentMatchAll or SegmentMatchAny.
	Match string `json:"match"`
	// Conditions holds at least one rule.
	Conditions []SegmentRule `json:"conditions"`
}

// SegmentGroup represents a nested group of conditions. A group cannot hold
// another group.
type SegmentGroup struct {
	// Match is SegmentMatchAll or SegmentMatchAny.
	Match string `json:"match"`
	// Conditions holds at least one condition.
	Conditions []SegmentCondition `json:"conditions"`
}

// SegmentCondition represents one rule: an operator that compares a field
// with a value.
type SegmentCondition struct {
	// Field is one of the SegmentField constants, or "attributes.<key>" for
	// a contact field of the account.
	Field string `json:"field"`
	// Op is one of the SegmentOp constants that applies to Field.
	Op string `json:"op"`
	// Value is the value Op compares with, nil for an operator that takes
	// none, such as SegmentOpIsEmpty. SegmentOpMonthIs and
	// SegmentOpAnniversaryWithin take a whole number, and
	// SegmentFieldCampaignCount a SegmentCampaignCount.
	Value *SegmentValue `json:"value,omitempty"`
}

// SegmentCampaignCount represents the value of a SegmentFieldCampaignCount
// condition: a number of campaigns sent since a time.
type SegmentCampaignCount struct {
	// Count is the number of campaigns, from 1 to 1000, or from 2 for
	// SegmentOpReceivedFewerThan.
	Count int `json:"count"`
	// Since is when the window starts: "M days ago", "M weeks ago",
	// "M months ago" or "M years ago", with M at most 365 days, 52 weeks,
	// 12 months or 1 year. It starts at the beginning of its day in the
	// segment's time zone.
	Since string `json:"since"`
}

// SegmentRule represents one entry of a SegmentDefinition: a condition or a
// group. Set exactly one of Condition and Group.
type SegmentRule struct {
	Condition *SegmentCondition
	Group     *SegmentGroup
}

// MarshalJSON writes the condition or the group that is set. It fails when
// both or neither are set.
func (r SegmentRule) MarshalJSON() ([]byte, error) {
	switch {
	case r.Condition != nil && r.Group != nil:
		return nil, errors.New("segment rule: set Condition or Group, not both")
	case r.Condition != nil:
		return json.Marshal(r.Condition)
	case r.Group != nil:
		return json.Marshal(r.Group)
	default:
		return nil, errors.New("segment rule: set Condition or Group")
	}
}

// UnmarshalJSON reads an object with a "match" key as a group and any other
// object as a condition.
func (r *SegmentRule) UnmarshalJSON(data []byte) error {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return errors.New("segment rule: null is not a condition or a group")
	}

	var keys map[string]json.RawMessage
	if err := json.Unmarshal(data, &keys); err != nil {
		return err
	}

	if _, ok := keys["match"]; ok {
		var group SegmentGroup
		if err := json.Unmarshal(data, &group); err != nil {
			return err
		}
		*r = SegmentRule{Group: &group}
		return nil
	}

	var condition SegmentCondition
	if err := json.Unmarshal(data, &condition); err != nil {
		return err
	}
	*r = SegmentRule{Condition: &condition}
	return nil
}

// SegmentValue represents the value of a SegmentCondition: a JSON string,
// number, boolean or SegmentCampaignCount object. Set exactly one field. A
// condition with no value has a nil *SegmentValue.
type SegmentValue struct {
	String *string
	Number *float64
	Bool   *bool
	Count  *SegmentCampaignCount
}

// MarshalJSON writes the field that is set. It fails when no field or more
// than one is set.
func (v SegmentValue) MarshalJSON() ([]byte, error) {
	set := 0
	for _, isSet := range []bool{v.String != nil, v.Number != nil, v.Bool != nil, v.Count != nil} {
		if isSet {
			set++
		}
	}
	if set != 1 {
		return nil, errors.New("segment value: set exactly one of String, Number, Bool and Count")
	}

	switch {
	case v.String != nil:
		return json.Marshal(*v.String)
	case v.Number != nil:
		return json.Marshal(*v.Number)
	case v.Bool != nil:
		return json.Marshal(*v.Bool)
	default:
		return json.Marshal(v.Count)
	}
}

// UnmarshalJSON reads a JSON string, number, boolean or object into the
// matching field.
func (v *SegmentValue) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return errors.New("segment value: empty JSON")
	}

	var value SegmentValue
	var err error
	switch trimmed[0] {
	case 'n':
		if !bytes.Equal(trimmed, []byte("null")) {
			return errors.New("segment value: invalid JSON")
		}
		return nil
	case '"':
		var s string
		err = json.Unmarshal(trimmed, &s)
		value.String = &s
	case 't', 'f':
		var b bool
		err = json.Unmarshal(trimmed, &b)
		value.Bool = &b
	case '{':
		var count SegmentCampaignCount
		err = json.Unmarshal(trimmed, &count)
		value.Count = &count
	default:
		var n float64
		err = json.Unmarshal(trimmed, &n)
		value.Number = &n
	}
	if err != nil {
		return err
	}

	*v = value
	return nil
}
