package requests

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	maxRequestNameLength      = 255
	maxRequestWebsiteLength   = 255
	maxAPIKeyLabelLength      = 255
	maxSuspensionReasonLength = 500
	minMonthlyCredit          = int64(0)
	maxMonthlyCredit          = int64(1000000000)
)

func validateRequiredString(field, value string, maxLength int) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s must not be blank", field)
	}

	if utf8.RuneCountInString(value) > maxLength {
		return fmt.Errorf("%s must be at most %d characters", field, maxLength)
	}

	return nil
}

func validateOptionalString(field string, value *string, maxLength int) error {
	if value == nil {
		return nil
	}

	return validateRequiredString(field, *value, maxLength)
}

func validateOptionalInt64Range(field string, value *int64, minValue, maxValue int64) error {
	if value == nil {
		return nil
	}

	if *value < minValue || *value > maxValue {
		return fmt.Errorf("%s must be between %d and %d", field, minValue, maxValue)
	}

	return nil
}

func validateAtLeastOneField(provided bool) error {
	if !provided {
		return fmt.Errorf("at least one field must be provided")
	}

	return nil
}

func validateOptionalNonEmptyStringSlice(field string, value *[]string) error {
	if value == nil {
		return nil
	}

	if len(*value) == 0 {
		return fmt.Errorf("%s must contain at least one item", field)
	}

	return nil
}
