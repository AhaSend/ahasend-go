package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newSubAccountsTestClient(t *testing.T, handler http.HandlerFunc) (*APIClient, func()) {
	t.Helper()

	server := httptest.NewServer(handler)
	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	cfg := NewConfiguration()
	cfg.Host = serverURL.Host
	cfg.Scheme = serverURL.Scheme
	cfg.APIKey = "test-key"
	cfg.RetryConfig.Enabled = false

	return NewAPIClientWithConfig(cfg), server.Close
}

func subAccountResponseJSON(accountID, subAccountID uuid.UUID, status string) string {
	return fmt.Sprintf(`{
		"object": "sub_account",
		"id": "%s",
		"parent_account_id": "%s",
		"name": "Acme Subsidiary",
		"website": "acme.example.com",
		"status": "%s",
		"monthly_credit": 50000,
		"created_at": "2024-01-01T00:00:00Z",
		"domain_count": 2,
		"member_count": 3,
		"last_activity_at": null
	}`, subAccountID.String(), accountID.String(), status)
}

func paginatedSubAccountsResponseJSON(accountID, subAccountID uuid.UUID) string {
	return fmt.Sprintf(`{
		"object": "list",
		"data": [%s],
		"pagination": {
			"has_more": false
		}
	}`, subAccountResponseJSON(accountID, subAccountID, "active"))
}

func subAccountsUsageResponseJSON(accountID, subAccountID uuid.UUID) string {
	return fmt.Sprintf(`{
		"billing_period": {
			"start": "2024-01-01T00:00:00Z",
			"end": "2024-02-01T00:00:00Z"
		},
		"currency": "usd",
		"allocation_method": "proportional",
		"allocation_note": "allocated_cost is a proportional share of the parent's pooled invoice for the period.",
		"parent": {
			"account_id": "%s",
			"reception_count": 1000000,
			"allocated_cost": 20.0
		},
		"sub_accounts": [
			{
				"account_id": "%s",
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
	}`, accountID.String(), subAccountID.String())
}

func subAccountAPIKeyResponseJSON(accountID, keyID, scopeID uuid.UUID, secretKey *string) string {
	secretKeyJSON := ""
	if secretKey != nil {
		secretKeyJSON = fmt.Sprintf(`,
		"secret_key": %q`, *secretKey)
	}

	return fmt.Sprintf(`{
		"object": "api_key",
		"id": "%s",
		"created_at": "2024-01-01T00:05:00Z",
		"updated_at": "2024-01-01T00:05:00Z",
		"account_id": "%s",
		"label": "Bootstrap key",
		"public_key": "aha-pk-child-public-key",
		"last_used_at": null%s,
		"scopes": [
			{
				"id": "%s",
				"created_at": "2024-01-01T00:05:00Z",
				"updated_at": "2024-01-01T00:05:00Z",
				"api_key_id": "%s",
				"scope": "messages:send:all",
				"domain_id": null
			}
		]
	}`, keyID.String(), accountID.String(), secretKeyJSON, scopeID.String(), keyID.String())
}

func paginatedSubAccountAPIKeysResponseJSON(accountID, keyID, scopeID uuid.UUID) string {
	return fmt.Sprintf(`{
		"object": "list",
		"data": [%s],
		"pagination": {
			"has_more": false
		}
	}`, subAccountAPIKeyResponseJSON(accountID, keyID, scopeID, nil))
}

func TestSubAccountsAPISuccessPaths(t *testing.T) {
	accountID := uuid.MustParse("9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1")
	subAccountID := uuid.MustParse("2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a")
	subAccountPath := "/v2/accounts/" + accountID.String() + "/sub-accounts/" + subAccountID.String()

	t.Run("ListSubAccounts", func(t *testing.T) {
		client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/v2/accounts/"+accountID.String()+"/sub-accounts", r.URL.Path)
			assert.Equal(t, "100", r.URL.Query().Get("limit"))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(paginatedSubAccountsResponseJSON(accountID, subAccountID)))
		})
		defer cleanup()

		resp, httpResp, err := client.SubAccountsAPI.ListSubAccounts(context.Background(), accountID, nil)

		require.NoError(t, err)
		require.NotNil(t, httpResp)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		require.NotNil(t, resp)
		require.Len(t, resp.Data, 1)
		assert.Equal(t, subAccountID, resp.Data[0].ID)
		assert.False(t, resp.Pagination.HasMore)
	})

	t.Run("CreateSubAccount", func(t *testing.T) {
		var body map[string]interface{}
		client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/v2/accounts/"+accountID.String()+"/sub-accounts", r.URL.Path)
			assert.Equal(t, "create-sub-account-test", r.Header.Get("Idempotency-Key"))

			err := json.NewDecoder(r.Body).Decode(&body)
			require.NoError(t, err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(subAccountResponseJSON(accountID, subAccountID, "active")))
		})
		defer cleanup()

		resp, httpResp, err := client.SubAccountsAPI.CreateSubAccount(context.Background(), accountID, requests.CreateSubAccountRequest{
			Name:          "Acme Subsidiary",
			Website:       "acme.example.com",
			MonthlyCredit: ahasend.Int64(50000),
		}, WithIdempotencyKey("create-sub-account-test"))

		require.NoError(t, err)
		require.NotNil(t, httpResp)
		assert.Equal(t, http.StatusCreated, httpResp.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, subAccountID, resp.ID)
		assert.Equal(t, "Acme Subsidiary", body["name"])
		assert.Equal(t, "acme.example.com", body["website"])
		assert.Equal(t, float64(50000), body["monthly_credit"])
	})

	t.Run("GetSubAccountsUsage", func(t *testing.T) {
		client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/v2/accounts/"+accountID.String()+"/sub-accounts/usage", r.URL.Path)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(subAccountsUsageResponseJSON(accountID, subAccountID)))
		})
		defer cleanup()

		resp, httpResp, err := client.SubAccountsAPI.GetSubAccountsUsage(context.Background(), accountID)

		require.NoError(t, err)
		require.NotNil(t, httpResp)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Parent.AccountID)
		assert.Equal(t, accountID, *resp.Parent.AccountID)
		require.Len(t, resp.SubAccounts, 1)
		require.NotNil(t, resp.SubAccounts[0].AccountID)
		assert.Equal(t, subAccountID, *resp.SubAccounts[0].AccountID)
	})

	t.Run("GetSubAccount", func(t *testing.T) {
		client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, subAccountPath, r.URL.Path)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(subAccountResponseJSON(accountID, subAccountID, "active")))
		})
		defer cleanup()

		resp, httpResp, err := client.SubAccountsAPI.GetSubAccount(context.Background(), accountID, subAccountID)

		require.NoError(t, err)
		require.NotNil(t, httpResp)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, subAccountID, resp.ID)
		assert.Equal(t, accountID, resp.ParentAccountID)
	})

	t.Run("UpdateSubAccount", func(t *testing.T) {
		var body map[string]interface{}
		client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Equal(t, subAccountPath, r.URL.Path)

			err := json.NewDecoder(r.Body).Decode(&body)
			require.NoError(t, err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(subAccountResponseJSON(accountID, subAccountID, "active")))
		})
		defer cleanup()

		resp, httpResp, err := client.SubAccountsAPI.UpdateSubAccount(context.Background(), accountID, subAccountID, requests.UpdateSubAccountRequest{
			Name: ahasend.String("Updated Subsidiary"),
		})

		require.NoError(t, err)
		require.NotNil(t, httpResp)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, subAccountID, resp.ID)
		assert.Equal(t, "Updated Subsidiary", body["name"])
		assert.NotContains(t, body, "website")
	})

	t.Run("DeleteSubAccount", func(t *testing.T) {
		client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Equal(t, subAccountPath, r.URL.Path)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"message":"sub account deleted"}`))
		})
		defer cleanup()

		resp, httpResp, err := client.SubAccountsAPI.DeleteSubAccount(context.Background(), accountID, subAccountID)

		require.NoError(t, err)
		require.NotNil(t, httpResp)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "sub account deleted", resp.Message)
	})

	t.Run("SuspendSubAccount", func(t *testing.T) {
		var body map[string]interface{}
		client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, subAccountPath+"/suspend", r.URL.Path)

			err := json.NewDecoder(r.Body).Decode(&body)
			require.NoError(t, err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(subAccountResponseJSON(accountID, subAccountID, "suspended")))
		})
		defer cleanup()

		resp, httpResp, err := client.SubAccountsAPI.SuspendSubAccount(context.Background(), accountID, subAccountID, requests.SuspendSubAccountRequest{
			Reason: "Customer requested temporary pause",
		})

		require.NoError(t, err)
		require.NotNil(t, httpResp)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "suspended", resp.Status)
		assert.Equal(t, "Customer requested temporary pause", body["reason"])
	})

	t.Run("UnsuspendSubAccount", func(t *testing.T) {
		client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, subAccountPath+"/unsuspend", r.URL.Path)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(subAccountResponseJSON(accountID, subAccountID, "active")))
		})
		defer cleanup()

		resp, httpResp, err := client.SubAccountsAPI.UnsuspendSubAccount(context.Background(), accountID, subAccountID)

		require.NoError(t, err)
		require.NotNil(t, httpResp)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "active", resp.Status)
	})
}

func TestSubAccountAPIKeyMethods(t *testing.T) {
	accountID := uuid.MustParse("9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1")
	subAccountID := uuid.MustParse("2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a")
	keyID := uuid.MustParse("13b3aa8e-78d3-48a1-92d2-4b8b1228c2dd")
	scopeID := uuid.MustParse("c574470d-76ef-4f74-9b24-70a583a17e03")
	apiKeysPath := "/v2/accounts/" + accountID.String() + "/sub-accounts/" + subAccountID.String() + "/api-keys"
	apiKeyPath := apiKeysPath + "/" + keyID.String()
	secretKey := "aha-sk-child-secret-key"

	tests := []struct {
		name            string
		wantMethod      string
		wantPath        string
		wantStatus      int
		wantIdempotency string
		wantQuery       map[string]string
		writeResponse   func(http.ResponseWriter)
		call            func(*testing.T, *APIClient) (*http.Response, error)
		assertBody      func(*testing.T, map[string]interface{})
	}{
		{
			name:       "ListSubAccountAPIKeys",
			wantMethod: http.MethodGet,
			wantPath:   apiKeysPath,
			wantStatus: http.StatusOK,
			wantQuery: map[string]string{
				"limit": "100",
			},
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(paginatedSubAccountAPIKeysResponseJSON(subAccountID, keyID, scopeID)))
			},
			call: func(t *testing.T, client *APIClient) (*http.Response, error) {
				resp, httpResp, err := client.SubAccountsAPI.ListSubAccountAPIKeys(context.Background(), accountID, subAccountID, nil)
				if err == nil {
					require.NotNil(t, resp)
					require.Len(t, resp.Data, 1)
					assert.Equal(t, keyID, resp.Data[0].ID)
					assert.Equal(t, subAccountID, resp.Data[0].AccountID)
					assert.Nil(t, resp.Data[0].SecretKey)
					assert.False(t, resp.Pagination.HasMore)
				}
				return httpResp, err
			},
		},
		{
			name:            "CreateSubAccountAPIKey",
			wantMethod:      http.MethodPost,
			wantPath:        apiKeysPath,
			wantStatus:      http.StatusCreated,
			wantIdempotency: "create-sub-account-api-key-test",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(subAccountAPIKeyResponseJSON(subAccountID, keyID, scopeID, &secretKey)))
			},
			call: func(t *testing.T, client *APIClient) (*http.Response, error) {
				resp, httpResp, err := client.SubAccountsAPI.CreateSubAccountAPIKey(context.Background(), accountID, subAccountID, requests.CreateAPIKeyRequest{
					Label:  "Bootstrap key",
					Scopes: []string{"messages:send:all", "domains:read"},
				}, WithIdempotencyKey("create-sub-account-api-key-test"))
				if err == nil {
					require.NotNil(t, resp)
					assert.Equal(t, keyID, resp.ID)
					assert.Equal(t, subAccountID, resp.AccountID)
					require.NotNil(t, resp.SecretKey)
					assert.Equal(t, secretKey, *resp.SecretKey)
					require.Len(t, resp.Scopes, 1)
					assert.Equal(t, "messages:send:all", resp.Scopes[0].Scope)
				}
				return httpResp, err
			},
			assertBody: func(t *testing.T, body map[string]interface{}) {
				assert.Equal(t, "Bootstrap key", body["label"])
				assert.Equal(t, []interface{}{"messages:send:all", "domains:read"}, body["scopes"])
			},
		},
		{
			name:       "GetSubAccountAPIKey",
			wantMethod: http.MethodGet,
			wantPath:   apiKeyPath,
			wantStatus: http.StatusOK,
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(subAccountAPIKeyResponseJSON(subAccountID, keyID, scopeID, nil)))
			},
			call: func(t *testing.T, client *APIClient) (*http.Response, error) {
				resp, httpResp, err := client.SubAccountsAPI.GetSubAccountAPIKey(context.Background(), accountID, subAccountID, keyID)
				if err == nil {
					require.NotNil(t, resp)
					assert.Equal(t, keyID, resp.ID)
					assert.Equal(t, subAccountID, resp.AccountID)
					assert.Nil(t, resp.SecretKey)
				}
				return httpResp, err
			},
		},
		{
			name:       "UpdateSubAccountAPIKey",
			wantMethod: http.MethodPut,
			wantPath:   apiKeyPath,
			wantStatus: http.StatusOK,
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(subAccountAPIKeyResponseJSON(subAccountID, keyID, scopeID, nil)))
			},
			call: func(t *testing.T, client *APIClient) (*http.Response, error) {
				scopes := []string{"messages:send:all"}
				resp, httpResp, err := client.SubAccountsAPI.UpdateSubAccountAPIKey(context.Background(), accountID, subAccountID, keyID, requests.UpdateAPIKeyRequest{
					Label:  ahasend.String("Updated bootstrap key"),
					Scopes: &scopes,
				})
				if err == nil {
					require.NotNil(t, resp)
					assert.Equal(t, keyID, resp.ID)
					assert.Equal(t, subAccountID, resp.AccountID)
					assert.Nil(t, resp.SecretKey)
				}
				return httpResp, err
			},
			assertBody: func(t *testing.T, body map[string]interface{}) {
				assert.Equal(t, "Updated bootstrap key", body["label"])
				assert.Equal(t, []interface{}{"messages:send:all"}, body["scopes"])
			},
		},
		{
			name:       "DeleteSubAccountAPIKey",
			wantMethod: http.MethodDelete,
			wantPath:   apiKeyPath,
			wantStatus: http.StatusOK,
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"message":"sub account api key deleted"}`))
			},
			call: func(t *testing.T, client *APIClient) (*http.Response, error) {
				resp, httpResp, err := client.SubAccountsAPI.DeleteSubAccountAPIKey(context.Background(), accountID, subAccountID, keyID)
				if err == nil {
					require.NotNil(t, resp)
					assert.Equal(t, "sub account api key deleted", resp.Message)
				}
				return httpResp, err
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var body map[string]interface{}
			client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.wantMethod, r.Method)
				assert.Equal(t, tt.wantPath, r.URL.Path)
				assert.Equal(t, tt.wantIdempotency, r.Header.Get("Idempotency-Key"))
				for _, key := range []string{"limit", "after", "before", "cursor"} {
					assert.Equal(t, tt.wantQuery[key], r.URL.Query().Get(key))
				}

				if tt.assertBody != nil {
					err := json.NewDecoder(r.Body).Decode(&body)
					require.NoError(t, err)
				}

				w.Header().Set("Content-Type", "application/json")
				tt.writeResponse(w)
			})
			defer cleanup()

			httpResp, err := tt.call(t, client)

			require.NoError(t, err)
			require.NotNil(t, httpResp)
			assert.Equal(t, tt.wantStatus, httpResp.StatusCode)
			if tt.assertBody != nil {
				require.NotNil(t, body)
				tt.assertBody(t, body)
			}
		})
	}
}

func TestListSubAccountsSerializesPagination(t *testing.T) {
	accountID := uuid.MustParse("9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1")
	subAccountID := uuid.MustParse("2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a")

	tests := []struct {
		name       string
		pagination *common.PaginationParams
		wantQuery  map[string]string
	}{
		{
			name:       "nil pagination uses default limit",
			pagination: nil,
			wantQuery: map[string]string{
				"limit": "100",
			},
		},
		{
			name: "explicit limit",
			pagination: &common.PaginationParams{
				Limit: ahasend.Int32(25),
			},
			wantQuery: map[string]string{
				"limit": "25",
			},
		},
		{
			name: "after takes priority over before and cursor",
			pagination: &common.PaginationParams{
				Limit:  ahasend.Int32(10),
				After:  ahasend.String("next-page"),
				Before: ahasend.String("previous-page"),
				Cursor: ahasend.String("legacy-page"),
			},
			wantQuery: map[string]string{
				"limit": "10",
				"after": "next-page",
			},
		},
		{
			name: "before takes priority over cursor",
			pagination: &common.PaginationParams{
				Before: ahasend.String("previous-page"),
				Cursor: ahasend.String("legacy-page"),
			},
			wantQuery: map[string]string{
				"limit":  "100",
				"before": "previous-page",
			},
		},
		{
			name: "legacy cursor",
			pagination: &common.PaginationParams{
				Cursor: ahasend.String("legacy-page"),
			},
			wantQuery: map[string]string{
				"limit":  "100",
				"cursor": "legacy-page",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotQuery url.Values
			client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				gotQuery = r.URL.Query()
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(paginatedSubAccountsResponseJSON(accountID, subAccountID)))
			})
			defer cleanup()

			_, _, err := client.SubAccountsAPI.ListSubAccounts(context.Background(), accountID, tt.pagination)
			require.NoError(t, err)
			require.NotNil(t, gotQuery)

			assert.Equal(t, tt.wantQuery["limit"], gotQuery.Get("limit"))
			for _, key := range []string{"after", "before", "cursor"} {
				assert.Equal(t, tt.wantQuery[key], gotQuery.Get(key))
			}
		})
	}
}

func TestListSubAccountAPIKeysSerializesPagination(t *testing.T) {
	accountID := uuid.MustParse("9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1")
	subAccountID := uuid.MustParse("2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a")
	keyID := uuid.MustParse("13b3aa8e-78d3-48a1-92d2-4b8b1228c2dd")
	scopeID := uuid.MustParse("c574470d-76ef-4f74-9b24-70a583a17e03")

	tests := []struct {
		name       string
		pagination *common.PaginationParams
		wantQuery  map[string]string
	}{
		{
			name:       "nil pagination uses default limit",
			pagination: nil,
			wantQuery: map[string]string{
				"limit": "100",
			},
		},
		{
			name: "explicit limit",
			pagination: &common.PaginationParams{
				Limit: ahasend.Int32(25),
			},
			wantQuery: map[string]string{
				"limit": "25",
			},
		},
		{
			name: "after takes priority over before and cursor",
			pagination: &common.PaginationParams{
				Limit:  ahasend.Int32(10),
				After:  ahasend.String("next-page"),
				Before: ahasend.String("previous-page"),
				Cursor: ahasend.String("legacy-page"),
			},
			wantQuery: map[string]string{
				"limit": "10",
				"after": "next-page",
			},
		},
		{
			name: "before takes priority over cursor",
			pagination: &common.PaginationParams{
				Before: ahasend.String("previous-page"),
				Cursor: ahasend.String("legacy-page"),
			},
			wantQuery: map[string]string{
				"limit":  "100",
				"before": "previous-page",
			},
		},
		{
			name: "legacy cursor",
			pagination: &common.PaginationParams{
				Cursor: ahasend.String("legacy-page"),
			},
			wantQuery: map[string]string{
				"limit":  "100",
				"cursor": "legacy-page",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var gotQuery url.Values
			client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				gotQuery = r.URL.Query()
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(paginatedSubAccountAPIKeysResponseJSON(subAccountID, keyID, scopeID)))
			})
			defer cleanup()

			_, _, err := client.SubAccountsAPI.ListSubAccountAPIKeys(context.Background(), accountID, subAccountID, tt.pagination)
			require.NoError(t, err)
			require.NotNil(t, gotQuery)

			assert.Equal(t, tt.wantQuery["limit"], gotQuery.Get("limit"))
			for _, key := range []string{"after", "before", "cursor"} {
				assert.Equal(t, tt.wantQuery[key], gotQuery.Get(key))
			}
		})
	}
}

func TestSubAccountRequestValidationShortCircuitsBeforeHTTP(t *testing.T) {
	var requestCount int
	client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusOK)
	})
	defer cleanup()

	accountID := uuid.New()
	subAccountID := uuid.New()
	keyID := uuid.New()

	tests := []struct {
		name string
		call func() (*http.Response, error)
	}{
		{
			name: "create",
			call: func() (*http.Response, error) {
				_, resp, err := client.SubAccountsAPI.CreateSubAccount(context.Background(), accountID, requests.CreateSubAccountRequest{
					Name:    "",
					Website: "https://example.com",
				})
				return resp, err
			},
		},
		{
			name: "update",
			call: func() (*http.Response, error) {
				_, resp, err := client.SubAccountsAPI.UpdateSubAccount(context.Background(), accountID, subAccountID, requests.UpdateSubAccountRequest{})
				return resp, err
			},
		},
		{
			name: "suspend",
			call: func() (*http.Response, error) {
				_, resp, err := client.SubAccountsAPI.SuspendSubAccount(context.Background(), accountID, subAccountID, requests.SuspendSubAccountRequest{
					Reason: "   ",
				})
				return resp, err
			},
		},
		{
			name: "update api key empty scopes",
			call: func() (*http.Response, error) {
				scopes := []string{}
				_, resp, err := client.SubAccountsAPI.UpdateSubAccountAPIKey(context.Background(), accountID, subAccountID, keyID, requests.UpdateAPIKeyRequest{
					Scopes: &scopes,
				})
				return resp, err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.call()
			require.Error(t, err)
			assert.Nil(t, resp)

			var apiErr *APIError
			require.True(t, errors.As(err, &apiErr))
			assert.Equal(t, ErrorTypeValidation, apiErr.Type)
		})
	}

	assert.Equal(t, 0, requestCount)
}

func TestSubAccountAPIErrorTypes(t *testing.T) {
	accountID := uuid.MustParse("9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1")
	subAccountID := uuid.MustParse("2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a")

	tests := []struct {
		statusCode int
		wantType   ErrorType
	}{
		{statusCode: http.StatusUnauthorized, wantType: ErrorTypeAuthentication},
		{statusCode: http.StatusForbidden, wantType: ErrorTypePermission},
		{statusCode: http.StatusNotFound, wantType: ErrorTypeNotFound},
		{statusCode: http.StatusConflict, wantType: ErrorTypeConflict},
		{statusCode: http.StatusUnprocessableEntity, wantType: ErrorTypeIdempotency},
		{statusCode: http.StatusInternalServerError, wantType: ErrorTypeServer},
	}

	for _, tt := range tests {
		t.Run(http.StatusText(tt.statusCode), func(t *testing.T) {
			client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(`{"message":"sub account error"}`))
			})
			defer cleanup()

			resp, httpResp, err := client.SubAccountsAPI.GetSubAccount(context.Background(), accountID, subAccountID)

			require.Error(t, err)
			assert.NotNil(t, resp)
			require.NotNil(t, httpResp)
			assert.Equal(t, tt.statusCode, httpResp.StatusCode)

			var apiErr *APIError
			require.True(t, errors.As(err, &apiErr))
			assert.Equal(t, tt.wantType, apiErr.Type)
			assert.Equal(t, tt.statusCode, apiErr.StatusCode)
			assert.Equal(t, "sub account error", apiErr.Message)
		})
	}
}

func TestSubAccountAPIKeyErrorTypes(t *testing.T) {
	accountID := uuid.MustParse("9d0cf9d0-4f5e-4674-bcf1-8ec39968b6e1")
	subAccountID := uuid.MustParse("2f3c5d2a-9ef8-4c91-a5f4-79990c8c1d3a")
	keyID := uuid.MustParse("13b3aa8e-78d3-48a1-92d2-4b8b1228c2dd")

	statuses := []struct {
		statusCode int
		wantType   ErrorType
	}{
		{statusCode: http.StatusUnauthorized, wantType: ErrorTypeAuthentication},
		{statusCode: http.StatusForbidden, wantType: ErrorTypePermission},
		{statusCode: http.StatusNotFound, wantType: ErrorTypeNotFound},
		{statusCode: http.StatusConflict, wantType: ErrorTypeConflict},
		{statusCode: http.StatusUnprocessableEntity, wantType: ErrorTypeIdempotency},
		{statusCode: http.StatusInternalServerError, wantType: ErrorTypeServer},
	}

	surfaces := []struct {
		name string
		call func(*APIClient) (*http.Response, error)
	}{
		{
			name: "create",
			call: func(client *APIClient) (*http.Response, error) {
				_, resp, err := client.SubAccountsAPI.CreateSubAccountAPIKey(context.Background(), accountID, subAccountID, requests.CreateAPIKeyRequest{
					Label:  "Bootstrap key",
					Scopes: []string{"messages:send:all"},
				})
				return resp, err
			},
		},
		{
			name: "update",
			call: func(client *APIClient) (*http.Response, error) {
				_, resp, err := client.SubAccountsAPI.UpdateSubAccountAPIKey(context.Background(), accountID, subAccountID, keyID, requests.UpdateAPIKeyRequest{
					Label: ahasend.String("Updated bootstrap key"),
				})
				return resp, err
			},
		},
		{
			name: "delete",
			call: func(client *APIClient) (*http.Response, error) {
				_, resp, err := client.SubAccountsAPI.DeleteSubAccountAPIKey(context.Background(), accountID, subAccountID, keyID)
				return resp, err
			},
		},
	}

	for _, surface := range surfaces {
		surface := surface
		for _, tt := range statuses {
			tt := tt
			t.Run(surface.name+"/"+http.StatusText(tt.statusCode), func(t *testing.T) {
				client, cleanup := newSubAccountsTestClient(t, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(tt.statusCode)
					_, _ = w.Write([]byte(`{"message":"sub account api key error"}`))
				})
				defer cleanup()

				httpResp, err := surface.call(client)

				require.Error(t, err)
				require.NotNil(t, httpResp)
				assert.Equal(t, tt.statusCode, httpResp.StatusCode)

				var apiErr *APIError
				require.True(t, errors.As(err, &apiErr))
				assert.Equal(t, tt.wantType, apiErr.Type)
				assert.Equal(t, tt.statusCode, apiErr.StatusCode)
				assert.Equal(t, "sub account api key error", apiErr.Message)
			})
		}
	}
}
