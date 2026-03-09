// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestSendersAgentExecutionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders:agent:executions", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--sender-id", "senderId",
			"--cursor", "cursor",
			"--limit", "100",
			"--status", "success",
		)
	})
}
