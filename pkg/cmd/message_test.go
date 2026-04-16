// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
	"github.com/zavudev/cli/internal/requestflag"
)

func TestMessagesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "retrieve",
			"--message-id", "messageId",
		)
	})
}

func TestMessagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "list",
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
			t,
			"--api-key", "string",
			"messages", "react",
			"--message-id", "messageId",
			"--emoji", "👍",
			"--zavu-sender", "sender_12345",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("emoji: 👍")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages", "react",
			"--message-id", "messageId",
			"--zavu-sender", "sender_12345",
		)
	})
}

func TestMessagesSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "send",
			"--to", "+56912345678",
			"--attachment", "{filename: invoice.pdf, content: content, content_id: logo, content_type: application/pdf, path: https://example.com}",
			"--channel", "auto",
			"--content", "{buttons: [{id: id, title: title}], contacts: [{name: name, phones: [string]}], ctaDisplayText: See Dates, ctaHeaderMediaUrl: https://example.com, ctaHeaderText: ctaHeaderText, ctaHeaderType: text, ctaUrl: https://example.com/schedule, emoji: emoji, filename: invoice.pdf, footerText: Dates subject to change., latitude: 0, listButton: listButton, locationAddress: locationAddress, locationName: locationName, longitude: 0, mediaId: mediaId, mediaUrl: https://example.com/image.jpg, mimeType: image/jpeg, reactToMessageId: reactToMessageId, sections: [{rows: [{id: id, title: title, description: description}], title: title}], templateId: templateId, templateVariables: {'1': John, '2': ORD-12345}}",
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
			t,
			"--api-key", "string",
			"messages", "send",
			"--to", "+56912345678",
			"--attachment.filename", "invoice.pdf",
			"--attachment.content", "content",
			"--attachment.content-id", "logo",
			"--attachment.content-type", "application/pdf",
			"--attachment.path", "https://example.com",
			"--channel", "auto",
			"--content.buttons", "[{id: id, title: title}]",
			"--content.contacts", "[{name: name, phones: [string]}]",
			"--content.cta-display-text", "See Dates",
			"--content.cta-header-media-url", "https://example.com",
			"--content.cta-header-text", "ctaHeaderText",
			"--content.cta-header-type", "text",
			"--content.cta-url", "https://example.com/schedule",
			"--content.emoji", "emoji",
			"--content.filename", "invoice.pdf",
			"--content.footer-text", "Dates subject to change.",
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
			"attachments:\n" +
			"  - filename: invoice.pdf\n" +
			"    content: content\n" +
			"    content_id: logo\n" +
			"    content_type: application/pdf\n" +
			"    path: https://example.com\n" +
			"channel: auto\n" +
			"content:\n" +
			"  buttons:\n" +
			"    - id: id\n" +
			"      title: title\n" +
			"  contacts:\n" +
			"    - name: name\n" +
			"      phones:\n" +
			"        - string\n" +
			"  ctaDisplayText: See Dates\n" +
			"  ctaHeaderMediaUrl: https://example.com\n" +
			"  ctaHeaderText: ctaHeaderText\n" +
			"  ctaHeaderType: text\n" +
			"  ctaUrl: https://example.com/schedule\n" +
			"  emoji: emoji\n" +
			"  filename: invoice.pdf\n" +
			"  footerText: Dates subject to change.\n" +
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
			t, pipeData,
			"--api-key", "string",
			"messages", "send",
			"--zavu-sender", "sender_12345",
		)
	})
}
