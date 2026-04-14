// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestExportsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"exports", "create",
			"--data-type", "messages",
			"--data-type", "conversations",
			"--date-from", "'2024-01-01T00:00:00Z'",
			"--date-to", "'2024-12-31T23:59:59Z'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"dataTypes:\n" +
			"  - messages\n" +
			"  - conversations\n" +
			"dateFrom: '2024-01-01T00:00:00Z'\n" +
			"dateTo: '2024-12-31T23:59:59Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"exports", "create",
		)
	})
}

func TestExportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"exports", "retrieve",
			"--export-id", "exportId",
		)
	})
}

func TestExportsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"exports", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
			"--status", "pending",
		)
	})
}
