// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
)

func TestSendersAgentFlowsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:flows", "create",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--name", "Lead Capture",
		"--step", "{id: welcome, config: {text: bar}, type: message, nextStepId: ask_name}",
		"--step", "{id: ask_name, config: {variable: bar, prompt: bar}, type: collect, nextStepId: nextStepId}",
		"--trigger", "{type: keyword, intent: intent, keywords: [info, pricing, demo]}",
		"--description", "Capture lead information",
		"--enabled=true",
		"--priority", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sendersAgentFlowsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:flows", "create",
		"--sender-id", "senderId",
		"--name", "Lead Capture",
		"--step.id", "welcome",
		"--step.config", "{text: bar}",
		"--step.type", "message",
		"--step.next-step-id", "ask_name",
		"--step.id", "ask_name",
		"--step.config", "{variable: bar, prompt: bar}",
		"--step.type", "collect",
		"--step.next-step-id", "nextStepId",
		"--trigger.type", "keyword",
		"--trigger.intent", "intent",
		"--trigger.keywords", "[info, pricing, demo]",
		"--description", "Capture lead information",
		"--enabled=true",
		"--priority", "0",
	)
}

func TestSendersAgentFlowsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:flows", "retrieve",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--flow-id", "flowId",
	)
}

func TestSendersAgentFlowsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:flows", "update",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--flow-id", "flowId",
		"--description", "description",
		"--enabled=true",
		"--name", "name",
		"--priority", "0",
		"--step", "{id: id, config: {foo: bar}, type: message, nextStepId: nextStepId}",
		"--trigger", "{type: keyword, intent: intent, keywords: [string]}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sendersAgentFlowsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:flows", "update",
		"--sender-id", "senderId",
		"--flow-id", "flowId",
		"--description", "description",
		"--enabled=true",
		"--name", "name",
		"--priority", "0",
		"--step.id", "id",
		"--step.config", "{foo: bar}",
		"--step.type", "message",
		"--step.next-step-id", "nextStepId",
		"--trigger.type", "keyword",
		"--trigger.intent", "intent",
		"--trigger.keywords", "[string]",
	)
}

func TestSendersAgentFlowsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:flows", "list",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--cursor", "cursor",
		"--enabled=true",
		"--limit", "100",
	)
}

func TestSendersAgentFlowsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:flows", "delete",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--flow-id", "flowId",
	)
}

func TestSendersAgentFlowsDuplicate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:flows", "duplicate",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--flow-id", "flowId",
		"--new-name", "Lead Capture (Copy)",
	)
}
