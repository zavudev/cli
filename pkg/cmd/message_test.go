// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
)

func TestMessagesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messages", "retrieve",
			"--api-key", "string",
			"--message-id", "messageId",
		)
	})
}

func TestMessagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messages", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--channel", "auto",
			"--cursor", "cursor",
			"--limit", "100",
			"--status", "queued",
			"--to", "to",
		)
	})
}

func TestMessagesReact(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messages", "react",
			"--api-key", "string",
			"--message-id", "messageId",
			"--emoji", "👍",
			"--zavu-sender", "sender_12345",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("emoji: 👍")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messages", "react",
			"--api-key", "string",
			"--message-id", "messageId",
			"--zavu-sender", "sender_12345",
		)
	})
}

func TestMessagesSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messages", "send",
			"--api-key", "string",
			"--to", "+56912345678",
			"--channel", "auto",
			"--content", "{buttons: [{id: id, title: title}], contacts: [{name: name, phones: [string]}], emoji: emoji, filename: invoice.pdf, latitude: 0, listButton: listButton, locationAddress: locationAddress, locationName: locationName, longitude: 0, mediaId: mediaId, mediaUrl: https://example.com/image.jpg, mimeType: image/jpeg, reactToMessageId: reactToMessageId, sections: [{rows: [{id: id, title: title, description: description}], title: title}], templateId: templateId, templateVariables: {'1': John, '2': ORD-12345}}",
			"--fallback-enabled=true",
			"--html-body", "htmlBody",
			"--idempotency-key", "msg_01HZY4ZP7VQY2J3BRW7Z6G0QGE",
			"--message-type", "text",
			"--metadata", "{foo: string}",
			"--reply-to", "support@example.com",
			"--subject", "Your order confirmation",
			"--text", "Your verification code is 123456",
			"--voice-language", "es-ES",
			"--zavu-sender", "sender_12345",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messagesSend)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "messages", "send",
			"--api-key", "string",
			"--to", "+56912345678",
			"--channel", "auto",
			"--content.buttons", "[{id: id, title: title}]",
			"--content.contacts", "[{name: name, phones: [string]}]",
			"--content.emoji", "emoji",
			"--content.filename", "invoice.pdf",
			"--content.latitude", "0",
			"--content.list-button", "listButton",
			"--content.location-address", "locationAddress",
			"--content.location-name", "locationName",
			"--content.longitude", "0",
			"--content.media-id", "mediaId",
			"--content.media-url", "https://example.com/image.jpg",
			"--content.mime-type", "image/jpeg",
			"--content.react-to-message-id", "reactToMessageId",
			"--content.sections", "[{rows: [{id: id, title: title, description: description}], title: title}]",
			"--content.template-id", "templateId",
			"--content.template-variables", "{'1': John, '2': ORD-12345}",
			"--fallback-enabled=true",
			"--html-body", "htmlBody",
			"--idempotency-key", "msg_01HZY4ZP7VQY2J3BRW7Z6G0QGE",
			"--message-type", "text",
			"--metadata", "{foo: string}",
			"--reply-to", "support@example.com",
			"--subject", "Your order confirmation",
			"--text", "Your verification code is 123456",
			"--voice-language", "es-ES",
			"--zavu-sender", "sender_12345",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"to: '+56912345678'\n" +
			"channel: auto\n" +
			"content:\n" +
			"  buttons:\n" +
			"    - id: id\n" +
			"      title: title\n" +
			"  contacts:\n" +
			"    - name: name\n" +
			"      phones:\n" +
			"        - string\n" +
			"  emoji: emoji\n" +
			"  filename: invoice.pdf\n" +
			"  latitude: 0\n" +
			"  listButton: listButton\n" +
			"  locationAddress: locationAddress\n" +
			"  locationName: locationName\n" +
			"  longitude: 0\n" +
			"  mediaId: mediaId\n" +
			"  mediaUrl: https://example.com/image.jpg\n" +
			"  mimeType: image/jpeg\n" +
			"  reactToMessageId: reactToMessageId\n" +
			"  sections:\n" +
			"    - rows:\n" +
			"        - id: id\n" +
			"          title: title\n" +
			"          description: description\n" +
			"      title: title\n" +
			"  templateId: templateId\n" +
			"  templateVariables:\n" +
			"    '1': John\n" +
			"    '2': ORD-12345\n" +
			"fallbackEnabled: true\n" +
			"htmlBody: htmlBody\n" +
			"idempotencyKey: msg_01HZY4ZP7VQY2J3BRW7Z6G0QGE\n" +
			"messageType: text\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"replyTo: support@example.com\n" +
			"subject: Your order confirmation\n" +
			"text: Your verification code is 123456\n" +
			"voiceLanguage: es-ES\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messages", "send",
			"--api-key", "string",
			"--zavu-sender", "sender_12345",
		)
	})
}
