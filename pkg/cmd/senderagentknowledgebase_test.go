// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestSendersAgentKnowledgeBasesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:knowledge-bases", "create",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--name", "Product FAQ",
		"--description", "Frequently asked questions about our products",
	)
}

func TestSendersAgentKnowledgeBasesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:knowledge-bases", "retrieve",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--kb-id", "kbId",
	)
}

func TestSendersAgentKnowledgeBasesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:knowledge-bases", "update",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--kb-id", "kbId",
		"--description", "description",
		"--name", "name",
	)
}

func TestSendersAgentKnowledgeBasesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:knowledge-bases", "list",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--cursor", "cursor",
		"--limit", "100",
	)
}

func TestSendersAgentKnowledgeBasesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:knowledge-bases", "delete",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--kb-id", "kbId",
	)
}
