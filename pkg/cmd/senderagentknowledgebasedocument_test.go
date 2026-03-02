// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestSendersAgentKnowledgeBasesDocumentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:knowledge-bases:documents", "create",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--kb-id", "kbId",
		"--content", "Our return policy allows returns within 30 days of purchase...",
		"--title", "Return Policy",
	)
}

func TestSendersAgentKnowledgeBasesDocumentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:knowledge-bases:documents", "list",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--kb-id", "kbId",
		"--cursor", "cursor",
		"--limit", "100",
	)
}

func TestSendersAgentKnowledgeBasesDocumentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:knowledge-bases:documents", "delete",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--kb-id", "kbId",
		"--doc-id", "docId",
	)
}
