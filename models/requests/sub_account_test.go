package requests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubAccountRequests_JSON(t *testing.T) {
	t.Run("create omits nil optional monthly credit", func(t *testing.T) {
		request := CreateSubAccountRequest{
			Name:    "Acme Subsidiary",
			Website: "acme.example.com",
		}

		data, err := json.Marshal(request)
		require.NoError(t, err)
		assert.JSONEq(t, `{"name":"Acme Subsidiary","website":"acme.example.com"}`, string(data))

		var decoded CreateSubAccountRequest
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)
		assert.Equal(t, request, decoded)
	})

	t.Run("create includes provided monthly credit", func(t *testing.T) {
		monthlyCredit := int64(50000)
		request := CreateSubAccountRequest{
			Name:          "Acme Subsidiary",
			Website:       "acme.example.com",
			MonthlyCredit: &monthlyCredit,
		}

		data, err := json.Marshal(request)
		require.NoError(t, err)
		assert.JSONEq(t, `{"name":"Acme Subsidiary","website":"acme.example.com","monthly_credit":50000}`, string(data))

		var decoded CreateSubAccountRequest
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)
		require.NotNil(t, decoded.MonthlyCredit)
		assert.Equal(t, monthlyCredit, *decoded.MonthlyCredit)
	})

	t.Run("update omits nil fields", func(t *testing.T) {
		name := "Acme Subsidiary"
		request := UpdateSubAccountRequest{Name: &name}

		data, err := json.Marshal(request)
		require.NoError(t, err)
		assert.JSONEq(t, `{"name":"Acme Subsidiary"}`, string(data))
	})

	t.Run("suspend uses reason", func(t *testing.T) {
		request := SuspendSubAccountRequest{Reason: "Customer requested temporary pause"}

		data, err := json.Marshal(request)
		require.NoError(t, err)
		assert.JSONEq(t, `{"reason":"Customer requested temporary pause"}`, string(data))
	})
}

func TestCreateSubAccountRequest_Validate(t *testing.T) {
	validCredit := int64(0)
	negativeCredit := int64(-1)
	tooLargeCredit := maxMonthlyCredit + 1

	tests := []struct {
		name    string
		request CreateSubAccountRequest
		wantErr bool
	}{
		{
			name: "valid without monthly credit",
			request: CreateSubAccountRequest{
				Name:    "Acme Subsidiary",
				Website: "acme.example.com",
			},
		},
		{
			name: "valid with zero monthly credit",
			request: CreateSubAccountRequest{
				Name:          "Acme Subsidiary",
				Website:       "acme.example.com",
				MonthlyCredit: &validCredit,
			},
		},
		{
			name: "blank name",
			request: CreateSubAccountRequest{
				Name:    "   ",
				Website: "acme.example.com",
			},
			wantErr: true,
		},
		{
			name: "overlong name",
			request: CreateSubAccountRequest{
				Name:    strings.Repeat("a", maxRequestNameLength+1),
				Website: "acme.example.com",
			},
			wantErr: true,
		},
		{
			name: "blank website",
			request: CreateSubAccountRequest{
				Name:    "Acme Subsidiary",
				Website: "   ",
			},
			wantErr: true,
		},
		{
			name: "negative monthly credit",
			request: CreateSubAccountRequest{
				Name:          "Acme Subsidiary",
				Website:       "acme.example.com",
				MonthlyCredit: &negativeCredit,
			},
			wantErr: true,
		},
		{
			name: "too large monthly credit",
			request: CreateSubAccountRequest{
				Name:          "Acme Subsidiary",
				Website:       "acme.example.com",
				MonthlyCredit: &tooLargeCredit,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestUpdateSubAccountRequest_Validate(t *testing.T) {
	name := "Acme Subsidiary"
	website := "acme.example.com"
	monthlyCredit := int64(50000)
	blankName := " "
	overlongName := strings.Repeat("a", maxRequestNameLength+1)
	blankWebsite := " "
	negativeCredit := int64(-1)
	tooLargeCredit := maxMonthlyCredit + 1

	tests := []struct {
		name    string
		request UpdateSubAccountRequest
		wantErr bool
	}{
		{
			name:    "all nil",
			request: UpdateSubAccountRequest{},
			wantErr: true,
		},
		{
			name:    "valid name only",
			request: UpdateSubAccountRequest{Name: &name},
		},
		{
			name:    "valid website only",
			request: UpdateSubAccountRequest{Website: &website},
		},
		{
			name:    "valid monthly credit only",
			request: UpdateSubAccountRequest{MonthlyCredit: &monthlyCredit},
		},
		{
			name:    "blank name",
			request: UpdateSubAccountRequest{Name: &blankName},
			wantErr: true,
		},
		{
			name:    "overlong name",
			request: UpdateSubAccountRequest{Name: &overlongName},
			wantErr: true,
		},
		{
			name:    "blank website",
			request: UpdateSubAccountRequest{Website: &blankWebsite},
			wantErr: true,
		},
		{
			name:    "negative monthly credit",
			request: UpdateSubAccountRequest{MonthlyCredit: &negativeCredit},
			wantErr: true,
		},
		{
			name:    "too large monthly credit",
			request: UpdateSubAccountRequest{MonthlyCredit: &tooLargeCredit},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestSuspendSubAccountRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request SuspendSubAccountRequest
		wantErr bool
	}{
		{
			name:    "valid reason",
			request: SuspendSubAccountRequest{Reason: "Customer requested temporary pause"},
		},
		{
			name:    "blank reason",
			request: SuspendSubAccountRequest{Reason: "   "},
			wantErr: true,
		},
		{
			name:    "overlong reason",
			request: SuspendSubAccountRequest{Reason: strings.Repeat("a", maxSuspensionReasonLength+1)},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
