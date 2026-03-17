// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestSendersAgentKnowledgeBasesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:knowledge-bases", "create",
			"--sender-id", "senderId",
			"--name", "Product FAQ",
			"--description", "Frequently asked questions about our products",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Product FAQ\n" +
			"description: Frequently asked questions about our products\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"senders:agent:knowledge-bases", "create",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersAgentKnowledgeBasesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:knowledge-bases", "retrieve",
			"--sender-id", "senderId",
			"--kb-id", "kbId",
		)
	})
}

func TestSendersAgentKnowledgeBasesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:knowledge-bases", "update",
			"--sender-id", "senderId",
			"--kb-id", "kbId",
			"--description", "description",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"senders:agent:knowledge-bases", "update",
			"--sender-id", "senderId",
			"--kb-id", "kbId",
		)
	})
}

func TestSendersAgentKnowledgeBasesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:knowledge-bases", "list",
			"--max-items", "10",
			"--sender-id", "senderId",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestSendersAgentKnowledgeBasesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:knowledge-bases", "delete",
			"--sender-id", "senderId",
			"--kb-id", "kbId",
		)
	})
}
