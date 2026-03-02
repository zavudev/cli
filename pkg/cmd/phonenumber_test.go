// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestPhoneNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers", "retrieve",
		"--api-key", "string",
		"--phone-number-id", "phoneNumberId",
	)
}

func TestPhoneNumbersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers", "update",
		"--api-key", "string",
		"--phone-number-id", "phoneNumberId",
		"--name", "Support Line",
		"--sender-id", "senderId",
	)
}

func TestPhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--limit", "100",
		"--status", "active",
	)
}

func TestPhoneNumbersPurchase(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers", "purchase",
		"--api-key", "string",
		"--phone-number", "+15551234567",
		"--name", "Primary Line",
	)
}

func TestPhoneNumbersRelease(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers", "release",
		"--api-key", "string",
		"--phone-number-id", "phoneNumberId",
	)
}

func TestPhoneNumbersRequirements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers", "requirements",
		"--api-key", "string",
		"--country-code", "xx",
		"--type", "local",
	)
}

func TestPhoneNumbersSearchAvailable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers", "search-available",
		"--api-key", "string",
		"--country-code", "xx",
		"--contains", "contains",
		"--limit", "50",
		"--type", "local",
	)
}
