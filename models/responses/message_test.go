package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessage_JSONMarshaling(t *testing.T) {
	now := time.Now().Truncate(time.Second) // Truncate for JSON precision
	domainID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	accountID := uuid.MustParse("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
	apiID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	deliveredAt := now.Add(5 * time.Minute)
	bounceClass := "hard"
	refMsgID := int64(789)
	msg := Message{
		Object:               "message",
		CreatedAt:            now,
		UpdatedAt:            now,
		SentAt:               &now,
		DeliveredAt:          &deliveredAt,
		RetainUntil:          now.Add(30 * 24 * time.Hour),
		MessageID:            "msg123@example.com",
		ID:                   apiID,
		Subject:              "Test Message",
		Tags:                 []string{"test", "automated"},
		Sender:               "sender@example.com",
		Recipient:            "recipient@example.com",
		Direction:            "outbound",
		Status:               "delivered",
		NumAttempts:          1,
		DeliveryAttempts:     []DeliveryEvent{{Time: now, Log: "Delivered", Status: "delivered"}},
		IsBounceNotification: false,
		BounceClassification: &bounceClass,
		ClickCount:           2,
		OpenCount:            1,
		ReferenceMessageID:   &refMsgID,
		DomainID:             domainID,
		AccountID:            accountID,
	}

	// Marshal to JSON
	data, err := json.Marshal(msg)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled Message
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, msg.Object, unmarshaled.Object)
	assert.Equal(t, msg.Subject, unmarshaled.Subject)
	assert.Equal(t, msg.Status, unmarshaled.Status)
	assert.Equal(t, msg.Sender, unmarshaled.Sender)
	assert.Equal(t, msg.Recipient, unmarshaled.Recipient)
	assert.Equal(t, msg.Tags, unmarshaled.Tags)

	// Verify pointer fields
	assert.NotNil(t, unmarshaled.SentAt)
	assert.True(t, unmarshaled.SentAt.Equal(now))

	assert.NotNil(t, unmarshaled.BounceClassification)
	assert.Equal(t, "hard", *unmarshaled.BounceClassification)

	assert.NotNil(t, unmarshaled.ReferenceMessageID)
	assert.Equal(t, int64(789), *unmarshaled.ReferenceMessageID)
}

func TestCreateMessageResponse_JSONMarshaling(t *testing.T) {
	response := CreateMessageResponse{
		Object: "list",
		Data: []CreateSingleMessageResponse{
			{
				Object:    "message",
				ID:        func() *string { s := "msg_123"; return &s }(),
				Recipient: common.Recipient{Email: "test@example.com"},
				Status:    "queued",
			},
			{
				Object:    "message",
				Recipient: common.Recipient{Email: "failed@example.com"},
				Status:    "failed",
				Error:     func() *string { s := "Invalid email"; return &s }(),
			},
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(response)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled CreateMessageResponse
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	// Verify
	assert.Equal(t, response.Object, unmarshaled.Object)
	assert.Len(t, unmarshaled.Data, 2)

	// Check first message
	first := unmarshaled.Data[0]
	assert.Equal(t, "message", first.Object)
	assert.NotNil(t, first.ID)
	assert.Equal(t, "msg_123", *first.ID)

	// Check second message
	second := unmarshaled.Data[1]
	assert.Equal(t, "failed", second.Status)
	assert.NotNil(t, second.Error)
	assert.Equal(t, "Invalid email", *second.Error)
}

func TestMessageSchedule_JSONMarshaling(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	expires := now.Add(24 * time.Hour)
	schedule := MessageSchedule{
		FirstAttempt: &now,
		Expires:      &expires,
	}

	// Marshal to JSON
	data, err := json.Marshal(schedule)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled MessageSchedule
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	// Verify
	assert.NotNil(t, unmarshaled.FirstAttempt)
	assert.True(t, unmarshaled.FirstAttempt.Equal(now))

	assert.NotNil(t, unmarshaled.Expires)
	assert.True(t, unmarshaled.Expires.Equal(now.Add(24*time.Hour)))
}

func TestMessage_OptionalFields_JSONMarshaling(t *testing.T) {
	t.Run("minimal message without optional fields", func(t *testing.T) {
		msg := Message{
			Object:               "message",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
			RetainUntil:          time.Now().Add(24 * time.Hour),
			MessageID:            "test@example.com",
			ID:                   uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Subject:              "Test",
			Tags:                 []string{},
			Sender:               "sender@example.com",
			Recipient:            "recipient@example.com",
			Direction:            "outbound",
			Status:               "queued",
			NumAttempts:          0,
			DeliveryAttempts:     []DeliveryEvent{},
			IsBounceNotification: false,
			ClickCount:           0,
			OpenCount:            0,
			DomainID:             uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd"),
			AccountID:            uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc"),
		}

		data, err := json.Marshal(msg)
		require.NoError(t, err)

		// Parse as generic map to check omitempty behavior
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// These should be omitted
		assert.NotContains(t, result, "sent_at")
		assert.NotContains(t, result, "delivered_at")
		assert.NotContains(t, result, "bounce_classification")
		assert.NotContains(t, result, "reference_message_id")

		// These should be present
		assert.Contains(t, result, "object")
		assert.Contains(t, result, "subject")
		assert.Contains(t, result, "status")
	})

	t.Run("complete message with all optional fields", func(t *testing.T) {
		now := time.Now().Truncate(time.Second)
		bounceClass := "soft"
		refID := int64(12345)

		msg := Message{
			Object:               "message",
			CreatedAt:            now,
			UpdatedAt:            now,
			SentAt:               &now,
			DeliveredAt:          &now,
			RetainUntil:          now.Add(24 * time.Hour),
			MessageID:            "test@example.com",
			ID:                   uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Subject:              "Test",
			Tags:                 []string{"test"},
			Sender:               "sender@example.com",
			Recipient:            "recipient@example.com",
			Direction:            "outbound",
			Status:               "delivered",
			NumAttempts:          2,
			DeliveryAttempts:     []DeliveryEvent{{Time: now, Log: "Delivered", Status: "delivered"}},
			IsBounceNotification: true,
			BounceClassification: &bounceClass,
			ClickCount:           3,
			OpenCount:            1,
			ReferenceMessageID:   &refID,
			DomainID:             uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd"),
			AccountID:            uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc"),
		}

		// Marshal to JSON
		data, err := json.Marshal(msg)
		require.NoError(t, err)

		// Unmarshal back
		var unmarshaled Message
		err = json.Unmarshal(data, &unmarshaled)
		require.NoError(t, err)

		// Verify all fields
		assert.Equal(t, msg.Object, unmarshaled.Object)
		assert.Equal(t, msg.MessageID, unmarshaled.MessageID)
		assert.Equal(t, msg.Subject, unmarshaled.Subject)
		assert.Equal(t, msg.Status, unmarshaled.Status)
		assert.Equal(t, msg.IsBounceNotification, unmarshaled.IsBounceNotification)
		assert.NotNil(t, unmarshaled.BounceClassification)
		assert.Equal(t, "soft", *unmarshaled.BounceClassification)
		assert.NotNil(t, unmarshaled.ReferenceMessageID)
		assert.Equal(t, int64(12345), *unmarshaled.ReferenceMessageID)
	})
}
