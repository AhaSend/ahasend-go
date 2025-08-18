package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount_JSONMarshaling(t *testing.T) {
	accountID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	createdAt := time.Now().UTC().Truncate(time.Second)
	updatedAt := createdAt.Add(time.Hour)

	t.Run("minimal account", func(t *testing.T) {
		account := Account{
			Object:    "account",
			ID:        accountID,
			Name:      "Test Account",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(account)
		require.NoError(t, err)

		// Should not contain optional fields
		assert.NotContains(t, string(jsonData), "website")
		assert.NotContains(t, string(jsonData), "about")
		assert.NotContains(t, string(jsonData), "track_opens")

		// Unmarshal and verify
		var decoded Account
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, account.Object, decoded.Object)
		assert.Equal(t, account.ID, decoded.ID)
		assert.Equal(t, account.Name, decoded.Name)
		assert.True(t, account.CreatedAt.Equal(decoded.CreatedAt))
		assert.True(t, account.UpdatedAt.Equal(decoded.UpdatedAt))

		// Optional fields should be nil
		assert.Nil(t, decoded.Website)
		assert.Nil(t, decoded.About)
		assert.Nil(t, decoded.TrackOpens)
	})

	t.Run("complete account with all optional fields", func(t *testing.T) {
		website := "https://example.com"
		about := "Test company description"
		trackOpens := true
		trackClicks := false
		rejectBad := true
		rejectMistyped := false
		metadataRetention := int32(30)
		dataRetention := int32(90)

		account := Account{
			Object:                   "account",
			ID:                       accountID,
			Name:                     "Complete Account",
			CreatedAt:                createdAt,
			UpdatedAt:                updatedAt,
			Website:                  &website,
			About:                    &about,
			TrackOpens:               &trackOpens,
			TrackClicks:              &trackClicks,
			RejectBadRecipients:      &rejectBad,
			RejectMistypedRecipients: &rejectMistyped,
			MessageMetadataRetention: &metadataRetention,
			MessageDataRetention:     &dataRetention,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(account)
		require.NoError(t, err)

		// Should contain all optional fields
		assert.Contains(t, string(jsonData), "website")
		assert.Contains(t, string(jsonData), "about")
		assert.Contains(t, string(jsonData), "track_opens")
		assert.Contains(t, string(jsonData), "track_clicks")

		// Unmarshal and verify
		var decoded Account
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, account.Name, decoded.Name)
		assert.Equal(t, website, *decoded.Website)
		assert.Equal(t, about, *decoded.About)
		assert.Equal(t, trackOpens, *decoded.TrackOpens)
		assert.Equal(t, trackClicks, *decoded.TrackClicks)
		assert.Equal(t, rejectBad, *decoded.RejectBadRecipients)
		assert.Equal(t, rejectMistyped, *decoded.RejectMistypedRecipients)
		assert.Equal(t, metadataRetention, *decoded.MessageMetadataRetention)
		assert.Equal(t, dataRetention, *decoded.MessageDataRetention)
	})
}

func TestUserAccount_JSONMarshaling(t *testing.T) {
	userAccountID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	userID := uuid.MustParse("11111111-2222-3333-4444-555555555555")
	accountID := uuid.MustParse("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
	createdAt := time.Now().UTC().Truncate(time.Second)
	updatedAt := createdAt.Add(time.Hour)

	userAccount := UserAccount{
		ID:        userAccountID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		UserID:    userID,
		AccountID: accountID,
		Role:      "Administrator",
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(userAccount)
	require.NoError(t, err)

	// Unmarshal and verify
	var decoded UserAccount
	err = json.Unmarshal(jsonData, &decoded)
	require.NoError(t, err)

	assert.Equal(t, userAccount.ID, decoded.ID)
	assert.True(t, userAccount.CreatedAt.Equal(decoded.CreatedAt))
	assert.True(t, userAccount.UpdatedAt.Equal(decoded.UpdatedAt))
	assert.Equal(t, userAccount.UserID, decoded.UserID)
	assert.Equal(t, userAccount.AccountID, decoded.AccountID)
	assert.Equal(t, userAccount.Role, decoded.Role)
}

func TestAccountMembersResponse_JSONMarshaling(t *testing.T) {
	response := AccountMembersResponse{
		Object: "list",
		Data: []UserAccount{
			{
				ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				UserID:    uuid.MustParse("22222222-2222-2222-2222-222222222222"),
				AccountID: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
				Role:      "Administrator",
			},
			{
				ID:        uuid.MustParse("44444444-4444-4444-4444-444444444444"),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				UserID:    uuid.MustParse("55555555-5555-5555-5555-555555555555"),
				AccountID: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
				Role:      "Developer",
			},
		},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(response)
	require.NoError(t, err)

	// Unmarshal and verify
	var decoded AccountMembersResponse
	err = json.Unmarshal(jsonData, &decoded)
	require.NoError(t, err)

	assert.Equal(t, response.Object, decoded.Object)
	assert.Len(t, decoded.Data, 2)
	assert.Equal(t, response.Data[0].Role, decoded.Data[0].Role)
	assert.Equal(t, response.Data[1].Role, decoded.Data[1].Role)
}
