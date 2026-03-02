// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
)

func TestSendersAgentToolsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:tools", "create",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--description", "Get the status of a customer order",
		"--name", "get_order_status",
		"--parameters", "{properties: {order_id: {description: The order ID to look up, type: string}}, required: [order_id], type: object}",
		"--webhook-url", "https://api.example.com/webhooks/order-status",
		"--enabled=true",
		"--webhook-secret", "whsec_...",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sendersAgentToolsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:tools", "create",
		"--sender-id", "senderId",
		"--description", "Get the status of a customer order",
		"--name", "get_order_status",
		"--parameters.properties", "{order_id: {description: The order ID to look up, type: string}}",
		"--parameters.required", "[order_id]",
		"--parameters.type", "object",
		"--webhook-url", "https://api.example.com/webhooks/order-status",
		"--enabled=true",
		"--webhook-secret", "whsec_...",
	)
}

func TestSendersAgentToolsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:tools", "retrieve",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--tool-id", "toolId",
	)
}

func TestSendersAgentToolsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:tools", "update",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--tool-id", "toolId",
		"--description", "description",
		"--enabled=true",
		"--name", "name",
		"--parameters", "{properties: {foo: {description: description, type: type}}, required: [string], type: object}",
		"--webhook-secret", "webhookSecret",
		"--webhook-url", "https://example.com",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sendersAgentToolsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:tools", "update",
		"--sender-id", "senderId",
		"--tool-id", "toolId",
		"--description", "description",
		"--enabled=true",
		"--name", "name",
		"--parameters.properties", "{foo: {description: description, type: type}}",
		"--parameters.required", "[string]",
		"--parameters.type", "object",
		"--webhook-secret", "webhookSecret",
		"--webhook-url", "https://example.com",
	)
}

func TestSendersAgentToolsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:tools", "list",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--cursor", "cursor",
		"--enabled=true",
		"--limit", "100",
	)
}

func TestSendersAgentToolsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:tools", "delete",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--tool-id", "toolId",
	)
}

func TestSendersAgentToolsTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"senders:agent:tools", "test",
		"--api-key", "string",
		"--sender-id", "senderId",
		"--tool-id", "toolId",
		"--test-params", "{order_id: bar}",
	)
}
