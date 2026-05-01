// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestURLsListVerified(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"urls", "list-verified",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
			"--status", "pending",
		)
	})
}

func TestURLsRetrieveDetails(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"urls", "retrieve-details",
			"--url-id", "urlId",
		)
	})
}

func TestURLsSubmitForVerification(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"urls", "submit-for-verification",
			"--url", "https://example.com/page",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("url: https://example.com/page")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"urls", "submit-for-verification",
		)
	})
}
