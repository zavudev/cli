// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestNumber10dlcCampaignsPhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns:phone-numbers", "list",
			"--campaign-id", "campaignId",
		)
	})
}

func TestNumber10dlcCampaignsPhoneNumbersAssign(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns:phone-numbers", "assign",
			"--campaign-id", "campaignId",
			"--phone-number-id", "pn_abc123",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("phoneNumberId: pn_abc123")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"number-10dlc:campaigns:phone-numbers", "assign",
			"--campaign-id", "campaignId",
		)
	})
}

func TestNumber10dlcCampaignsPhoneNumbersUnassign(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns:phone-numbers", "unassign",
			"--campaign-id", "campaignId",
			"--assignment-id", "assignmentId",
		)
	})
}
