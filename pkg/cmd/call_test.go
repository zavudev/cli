// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestCallsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "create",
			"--to", "+56912345678",
			"--greeting", "greeting",
			"--language", "es-ES",
			"--max-duration-minutes", "1",
			"--metadata", "{foo: string}",
			"--sender-id", "sender_12345",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"to: '+56912345678'\n" +
			"greeting: greeting\n" +
			"language: es-ES\n" +
			"maxDurationMinutes: 1\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"senderId: sender_12345\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls", "create",
		)
	})
}

func TestCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "retrieve",
			"--call-id", "callId",
		)
	})
}

func TestCallsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--direction", "inbound",
			"--limit", "100",
			"--status", "queued",
		)
	})
}

func TestCallsHangup(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "hangup",
			"--call-id", "callId",
		)
	})
}
