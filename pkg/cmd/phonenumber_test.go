// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestPhoneNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "retrieve",
			"--phone-number-id", "phoneNumberId",
		)
	})
}

func TestPhoneNumbersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "update",
			"--phone-number-id", "phoneNumberId",
			"--name", "Support Line",
			"--sender-id", "senderId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Support Line\n" +
			"senderId: senderId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"phone-numbers", "update",
			"--phone-number-id", "phoneNumberId",
		)
	})
}

func TestPhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
			"--status", "active",
		)
	})
}

func TestPhoneNumbersPurchase(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "purchase",
			"--phone-number", "+15551234567",
			"--name", "Primary Line",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phoneNumber: '+15551234567'\n" +
			"name: Primary Line\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"phone-numbers", "purchase",
		)
	})
}

func TestPhoneNumbersRelease(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "release",
			"--phone-number-id", "phoneNumberId",
		)
	})
}

func TestPhoneNumbersRequirements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "requirements",
			"--country-code", "xx",
			"--type", "local",
		)
	})
}

func TestPhoneNumbersSearchAvailable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "search-available",
			"--country-code", "xx",
			"--contains", "contains",
			"--limit", "50",
			"--type", "local",
		)
	})
}
