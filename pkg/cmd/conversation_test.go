// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestConversationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations", "retrieve",
			"--conversation-id", "conversationId",
		)
	})
}

func TestConversationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations", "list",
			"--max-items", "10",
			"--channel", "sms",
			"--cursor", "cursor",
			"--limit", "100",
			"--search", "+56912345678",
			"--sender-id", "senderId",
		)
	})
}

func TestConversationsListMessages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations", "list-messages",
			"--max-items", "10",
			"--conversation-id", "conversationId",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestConversationsMarkAsRead(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations", "mark-as-read",
			"--conversation-id", "conversationId",
		)
	})
}
