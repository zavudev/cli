// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
	"github.com/zavudev/cli/internal/requestflag"
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
			"--regulatory-requirement", "{fieldValue: jd7x2k3m4n5p6q7r8s9t0abc, requirementType: 8c5b1a2e-0f3d-4f5b-9a61-2c7e4d9b1f10}",
			"--type", "local",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(phoneNumbersPurchase)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "purchase",
			"--phone-number", "+15551234567",
			"--name", "Primary Line",
			"--regulatory-requirement.field-value", "jd7x2k3m4n5p6q7r8s9t0abc",
			"--regulatory-requirement.requirement-type", "8c5b1a2e-0f3d-4f5b-9a61-2c7e4d9b1f10",
			"--type", "local",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phoneNumber: '+15551234567'\n" +
			"name: Primary Line\n" +
			"regulatoryRequirements:\n" +
			"  - fieldValue: jd7x2k3m4n5p6q7r8s9t0abc\n" +
			"    requirementType: 8c5b1a2e-0f3d-4f5b-9a61-2c7e4d9b1f10\n" +
			"type: local\n")
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
			"--phone-number", "phoneNumber",
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
			"--capabilities", "voice,sms",
			"--contains", "contains",
			"--limit", "50",
			"--type", "local",
		)
	})
}
