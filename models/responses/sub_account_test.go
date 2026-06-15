package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubAccount_JSONRoundTrip(t *testing.T) {
	subAccountID := uuid.MustParse("2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a")
	parentAccountID := uuid.MustParse("9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1")
	createdAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	lastActivityAt := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)

	t.Run("openapi example", func(t *testing.T) {
		data := []byte(`{
			"object": "sub_account",
			"id": "2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a",
			"parent_account_id": "9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1",
			"name": "Acme Subsidiary",
			"website": "acme.example.com",
			"status": "active",
			"monthly_credit": 0,
			"created_at": "2024-01-01T00:00:00Z",
			"domain_count": 2,
			"member_count": 3,
			"last_activity_at": "2024-01-15T12:00:00Z"
		}`)

		var subAccount SubAccount
		err := json.Unmarshal(data, &subAccount)
		require.NoError(t, err)

		assert.Equal(t, "sub_account", subAccount.Object)
		assert.Equal(t, subAccountID, subAccount.ID)
		assert.Equal(t, parentAccountID, subAccount.ParentAccountID)
		assert.Equal(t, "Acme Subsidiary", subAccount.Name)
		assert.Equal(t, "acme.example.com", subAccount.Website)
		assert.Equal(t, "active", subAccount.Status)
		assert.Equal(t, int64(0), subAccount.MonthlyCredit)
		assert.Equal(t, int64(2), subAccount.DomainCount)
		assert.Equal(t, int64(3), subAccount.MemberCount)
		assert.True(t, subAccount.CreatedAt.Equal(createdAt))
		require.NotNil(t, subAccount.LastActivityAt)
		assert.True(t, subAccount.LastActivityAt.Equal(lastActivityAt))

		jsonData, err := json.Marshal(subAccount)
		require.NoError(t, err)

		var decoded SubAccount
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, subAccount, decoded)
	})

	t.Run("nullable last activity", func(t *testing.T) {
		subAccount := SubAccount{
			Object:          "sub_account",
			ID:              subAccountID,
			ParentAccountID: parentAccountID,
			CreatedAt:       createdAt,
			Name:            "Acme Subsidiary",
			Website:         "acme.example.com",
			Status:          "active",
			MonthlyCredit:   50000,
			DomainCount:     0,
			MemberCount:     1,
			LastActivityAt:  nil,
		}

		jsonData, err := json.Marshal(subAccount)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(jsonData, &result)
		require.NoError(t, err)
		assert.Contains(t, result, "last_activity_at")
		assert.Nil(t, result["last_activity_at"])

		var decoded SubAccount
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)
		assert.Nil(t, decoded.LastActivityAt)
		assert.Equal(t, subAccount.ID, decoded.ID)
	})
}

func TestPaginatedSubAccountsResponse_JSONRoundTrip(t *testing.T) {
	data := []byte(`{
		"object": "list",
		"data": [
			{
				"object": "sub_account",
				"id": "2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a",
				"parent_account_id": "9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1",
				"name": "Acme Subsidiary",
				"website": "acme.example.com",
				"status": "active",
				"monthly_credit": 0,
				"created_at": "2024-01-01T00:00:00Z",
				"domain_count": 2,
				"member_count": 3,
				"last_activity_at": null
			}
		],
		"pagination": {
			"has_more": true,
			"next_cursor": "next-page",
			"previous_cursor": "previous-page"
		}
	}`)

	var response PaginatedSubAccountsResponse
	err := json.Unmarshal(data, &response)
	require.NoError(t, err)

	assert.Equal(t, "list", response.Object)
	require.Len(t, response.Data, 1)
	assert.Equal(t, "Acme Subsidiary", response.Data[0].Name)
	assert.Nil(t, response.Data[0].LastActivityAt)
	assert.True(t, response.Pagination.HasMore)
	require.NotNil(t, response.Pagination.NextCursor)
	assert.Equal(t, "next-page", *response.Pagination.NextCursor)
	require.NotNil(t, response.Pagination.PreviousCursor)
	assert.Equal(t, "previous-page", *response.Pagination.PreviousCursor)

	jsonData, err := json.Marshal(response)
	require.NoError(t, err)

	var decoded PaginatedSubAccountsResponse
	err = json.Unmarshal(jsonData, &decoded)
	require.NoError(t, err)
	assert.Equal(t, response, decoded)
}

func TestSubAccountUsageResponse_JSONRoundTrip(t *testing.T) {
	parentAccountID := uuid.MustParse("9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1")
	subAccountID := uuid.MustParse("2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a")
	periodStart := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	periodEnd := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

	data := []byte(`{
		"billing_period": {
			"start": "2024-01-01T00:00:00Z",
			"end": "2024-02-01T00:00:00Z"
		},
		"currency": "usd",
		"allocation_method": "proportional",
		"allocation_note": "allocated_cost is a proportional share of the parent's pooled invoice for the period, not what the sub would pay on its own plan.",
		"parent": {
			"account_id": "9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1",
			"reception_count": 1000000,
			"allocated_cost": 20.0
		},
		"sub_accounts": [
			{
				"account_id": "2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a",
				"name": "Acme Subsidiary",
				"reception_count": 3000000,
				"allocated_cost": 60.0
			}
		],
		"removed_sub_accounts": {
			"reception_count": 0,
			"allocated_cost": 0.0
		},
		"total": {
			"reception_count": 4000000,
			"allocated_cost": 80.0
		}
	}`)

	var response SubAccountUsageResponse
	err := json.Unmarshal(data, &response)
	require.NoError(t, err)

	assert.True(t, response.BillingPeriod.Start.Equal(periodStart))
	assert.True(t, response.BillingPeriod.End.Equal(periodEnd))
	assert.Equal(t, "usd", response.Currency)
	assert.Equal(t, "proportional", response.AllocationMethod)

	require.NotNil(t, response.Parent.AccountID)
	assert.Equal(t, parentAccountID, *response.Parent.AccountID)
	assert.Nil(t, response.Parent.Name)
	assert.Equal(t, int64(1000000), response.Parent.ReceptionCount)
	assert.Equal(t, 20.0, response.Parent.AllocatedCost)

	require.Len(t, response.SubAccounts, 1)
	require.NotNil(t, response.SubAccounts[0].AccountID)
	assert.Equal(t, subAccountID, *response.SubAccounts[0].AccountID)
	require.NotNil(t, response.SubAccounts[0].Name)
	assert.Equal(t, "Acme Subsidiary", *response.SubAccounts[0].Name)
	assert.Equal(t, int64(3000000), response.SubAccounts[0].ReceptionCount)
	assert.Equal(t, 60.0, response.SubAccounts[0].AllocatedCost)

	assert.Nil(t, response.RemovedSubAccounts.AccountID)
	assert.Nil(t, response.RemovedSubAccounts.Name)
	assert.Equal(t, int64(0), response.RemovedSubAccounts.ReceptionCount)
	assert.Equal(t, 0.0, response.RemovedSubAccounts.AllocatedCost)

	assert.Nil(t, response.Total.AccountID)
	assert.Nil(t, response.Total.Name)
	assert.Equal(t, int64(4000000), response.Total.ReceptionCount)
	assert.Equal(t, 80.0, response.Total.AllocatedCost)

	jsonData, err := json.Marshal(response)
	require.NoError(t, err)

	var decoded SubAccountUsageResponse
	err = json.Unmarshal(jsonData, &decoded)
	require.NoError(t, err)
	assert.Equal(t, response, decoded)
}
