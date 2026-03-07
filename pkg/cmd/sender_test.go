// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestSendersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "create",
			"--api-key", "string",
			"--name", "name",
			"--phone-number", "phoneNumber",
			"--set-as-default=true",
			"--webhook-event", "message.queued",
			"--webhook-url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"phoneNumber: phoneNumber\n" +
			"setAsDefault: true\n" +
			"webhookEvents:\n" +
			"  - message.queued\n" +
			"webhookUrl: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "senders", "create",
			"--api-key", "string",
		)
	})
}

func TestSendersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "retrieve",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "update",
			"--api-key", "string",
			"--sender-id", "senderId",
			"--email-receiving-enabled=true",
			"--name", "name",
			"--set-as-default=true",
			"--webhook-active=true",
			"--webhook-event", "message.queued",
			"--webhook-url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"emailReceivingEnabled: true\n" +
			"name: name\n" +
			"setAsDefault: true\n" +
			"webhookActive: true\n" +
			"webhookEvents:\n" +
			"  - message.queued\n" +
			"webhookUrl: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "senders", "update",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestSendersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "delete",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersGetProfile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "get-profile",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersRegenerateWebhookSecret(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "regenerate-webhook-secret",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersUpdateProfile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "update-profile",
			"--api-key", "string",
			"--sender-id", "senderId",
			"--about", "Succulent specialists!",
			"--address", "address",
			"--description", "We specialize in providing high-quality succulents.",
			"--email", "contact@example.com",
			"--vertical", "RETAIL",
			"--website", "https://www.example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"about: Succulent specialists!\n" +
			"address: address\n" +
			"description: We specialize in providing high-quality succulents.\n" +
			"email: contact@example.com\n" +
			"vertical: RETAIL\n" +
			"websites:\n" +
			"  - https://www.example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "senders", "update-profile",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersUploadProfilePicture(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders", "upload-profile-picture",
			"--api-key", "string",
			"--sender-id", "senderId",
			"--image-url", "https://example.com/profile.jpg",
			"--mime-type", "image/jpeg",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"imageUrl: https://example.com/profile.jpg\n" +
			"mimeType: image/jpeg\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "senders", "upload-profile-picture",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}
