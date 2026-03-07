// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
)

func TestBroadcastsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "create",
		"--api-key", "string",
		"--channel", "sms",
		"--name", "Black Friday Sale",
		"--content", "{filename: filename, mediaId: mediaId, mediaUrl: mediaUrl, mimeType: mimeType, templateId: templateId, templateVariables: {foo: string}}",
		"--email-html-body", "emailHtmlBody",
		"--email-subject", "emailSubject",
		"--idempotency-key", "idempotencyKey",
		"--message-type", "text",
		"--metadata", "{foo: string}",
		"--scheduled-at", "'2019-12-27T18:11:19.117Z'",
		"--sender-id", "senderId",
		"--text", "Hi {{name}}, check out our Black Friday deals! Use code FRIDAY20 for 20% off.",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(broadcastsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "create",
		"--api-key", "string",
		"--channel", "sms",
		"--name", "Black Friday Sale",
		"--content.filename", "filename",
		"--content.media-id", "mediaId",
		"--content.media-url", "mediaUrl",
		"--content.mime-type", "mimeType",
		"--content.template-id", "templateId",
		"--content.template-variables", "{foo: string}",
		"--email-html-body", "emailHtmlBody",
		"--email-subject", "emailSubject",
		"--idempotency-key", "idempotencyKey",
		"--message-type", "text",
		"--metadata", "{foo: string}",
		"--scheduled-at", "'2019-12-27T18:11:19.117Z'",
		"--sender-id", "senderId",
		"--text", "Hi {{name}}, check out our Black Friday deals! Use code FRIDAY20 for 20% off.",
	)
}

func TestBroadcastsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "retrieve",
		"--api-key", "string",
		"--broadcast-id", "broadcastId",
	)
}

func TestBroadcastsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "update",
		"--api-key", "string",
		"--broadcast-id", "broadcastId",
		"--content", "{filename: filename, mediaId: mediaId, mediaUrl: mediaUrl, mimeType: mimeType, templateId: templateId, templateVariables: {foo: string}}",
		"--email-html-body", "emailHtmlBody",
		"--email-subject", "emailSubject",
		"--metadata", "{foo: string}",
		"--name", "name",
		"--text", "text",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(broadcastsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "update",
		"--api-key", "string",
		"--broadcast-id", "broadcastId",
		"--content.filename", "filename",
		"--content.media-id", "mediaId",
		"--content.media-url", "mediaUrl",
		"--content.mime-type", "mimeType",
		"--content.template-id", "templateId",
		"--content.template-variables", "{foo: string}",
		"--email-html-body", "emailHtmlBody",
		"--email-subject", "emailSubject",
		"--metadata", "{foo: string}",
		"--name", "name",
		"--text", "text",
	)
}

func TestBroadcastsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--limit", "100",
		"--status", "draft",
	)
}

func TestBroadcastsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "delete",
		"--api-key", "string",
		"--broadcast-id", "broadcastId",
	)
}

func TestBroadcastsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "cancel",
		"--api-key", "string",
		"--broadcast-id", "broadcastId",
	)
}

func TestBroadcastsProgress(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "progress",
		"--api-key", "string",
		"--broadcast-id", "broadcastId",
	)
}

func TestBroadcastsReschedule(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "reschedule",
		"--api-key", "string",
		"--broadcast-id", "broadcastId",
		"--scheduled-at", "'2024-01-15T14:00:00Z'",
	)
}

func TestBroadcastsSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"broadcasts", "send",
		"--api-key", "string",
		"--broadcast-id", "broadcastId",
		"--scheduled-at", "'2019-12-27T18:11:19.117Z'",
	)
}
