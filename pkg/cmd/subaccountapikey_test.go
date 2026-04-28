// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestSubAccountsAPIKeysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sub-accounts:api-keys", "create",
			"--id", "id",
			"--name", "Production Key",
			"--environment", "live",
			"--permission", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Production Key\n" +
			"environment: live\n" +
			"permissions:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sub-accounts:api-keys", "create",
			"--id", "id",
		)
	})
}

func TestSubAccountsAPIKeysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sub-accounts:api-keys", "list",
			"--id", "id",
		)
	})
}

func TestSubAccountsAPIKeysRevoke(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sub-accounts:api-keys", "revoke",
			"--id", "id",
			"--key-id", "keyId",
		)
	})
}
