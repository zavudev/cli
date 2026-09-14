// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestSendersWhatsappSyncRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:whatsapp-sync", "retrieve",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersWhatsappSyncStartContactsSync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:whatsapp-sync", "start-contacts-sync",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersWhatsappSyncStartHistorySync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:whatsapp-sync", "start-history-sync",
			"--sender-id", "senderId",
		)
	})
}
