// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
	"github.com/zavudev/cli/internal/requestflag"
)

func TestSendersAgentToolsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:tools", "create",
			"--sender-id", "senderId",
			"--description", "Get the status of a customer order",
			"--name", "get_order_status",
			"--parameters", "{properties: {order_id: {description: The order ID to look up, type: string}}, required: [order_id], type: object}",
			"--webhook-url", "https://api.example.com/webhooks/order-status",
			"--enabled=true",
			"--webhook-secret", "whsec_...",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sendersAgentToolsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: Get the status of a customer order\n" +
			"name: get_order_status\n" +
			"parameters:\n" +
			"  properties:\n" +
			"    order_id:\n" +
			"      description: The order ID to look up\n" +
			"      type: string\n" +
			"  required:\n" +
			"    - order_id\n" +
			"  type: object\n" +
			"webhookUrl: https://api.example.com/webhooks/order-status\n" +
			"enabled: true\n" +
			"webhookSecret: whsec_...\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"senders:agent:tools", "create",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersAgentToolsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:tools", "retrieve",
			"--sender-id", "senderId",
			"--tool-id", "toolId",
		)
	})
}

func TestSendersAgentToolsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:tools", "update",
			"--sender-id", "senderId",
			"--tool-id", "toolId",
			"--description", "description",
			"--enabled=true",
			"--name", "name",
			"--parameters", "{properties: {foo: {description: description, type: type}}, required: [string], type: object}",
			"--webhook-secret", "webhookSecret",
			"--webhook-url", "https://example.com",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sendersAgentToolsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"enabled: true\n" +
			"name: name\n" +
			"parameters:\n" +
			"  properties:\n" +
			"    foo:\n" +
			"      description: description\n" +
			"      type: type\n" +
			"  required:\n" +
			"    - string\n" +
			"  type: object\n" +
			"webhookSecret: webhookSecret\n" +
			"webhookUrl: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"senders:agent:tools", "update",
			"--sender-id", "senderId",
			"--tool-id", "toolId",
		)
	})
}

func TestSendersAgentToolsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:tools", "list",
			"--max-items", "10",
			"--sender-id", "senderId",
			"--cursor", "cursor",
			"--enabled=true",
			"--limit", "100",
		)
	})
}

func TestSendersAgentToolsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:tools", "delete",
			"--sender-id", "senderId",
			"--tool-id", "toolId",
		)
	})
}

func TestSendersAgentToolsTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"senders:agent:tools", "test",
			"--sender-id", "senderId",
			"--tool-id", "toolId",
			"--test-params", "{order_id: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"testParams:\n" +
			"  order_id: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"senders:agent:tools", "test",
			"--sender-id", "senderId",
			"--tool-id", "toolId",
		)
	})
}
