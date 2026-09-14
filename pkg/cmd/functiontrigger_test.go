// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestFunctionsTriggersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:triggers", "create",
			"--function-id", "functionId",
			"--event-type", "message.inbound",
			"--sender-id", "null",
			"--cron", "0 9 * * 1-5",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"eventTypes:\n" +
			"  - message.inbound\n" +
			"senderIds:\n" +
			"  - null\n" +
			"cron: 0 9 * * 1-5\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions:triggers", "create",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsTriggersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:triggers", "update",
			"--trigger-id", "triggerId",
			"--active=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("active: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions:triggers", "update",
			"--trigger-id", "triggerId",
		)
	})
}

func TestFunctionsTriggersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:triggers", "list",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsTriggersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:triggers", "delete",
			"--trigger-id", "triggerId",
		)
	})
}
