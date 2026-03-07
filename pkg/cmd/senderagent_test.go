// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestSendersAgentCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders:agent", "create",
			"--api-key", "string",
			"--sender-id", "senderId",
			"--model", "gpt-4o-mini",
			"--name", "Customer Support",
			"--provider", "openai",
			"--system-prompt", "You are a helpful customer support agent. Be friendly and concise.",
			"--api-key", "sk-...",
			"--context-window-messages", "1",
			"--include-contact-metadata=true",
			"--max-tokens", "1",
			"--temperature", "0",
			"--trigger-on-channel", "string",
			"--trigger-on-message-type", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"model: gpt-4o-mini\n" +
			"name: Customer Support\n" +
			"provider: openai\n" +
			"systemPrompt: You are a helpful customer support agent. Be friendly and concise.\n" +
			"apiKey: sk-...\n" +
			"contextWindowMessages: 1\n" +
			"includeContactMetadata: true\n" +
			"maxTokens: 1\n" +
			"temperature: 0\n" +
			"triggerOnChannels:\n" +
			"  - string\n" +
			"triggerOnMessageTypes:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "senders:agent", "create",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersAgentRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders:agent", "retrieve",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersAgentUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders:agent", "update",
			"--api-key", "string",
			"--sender-id", "senderId",
			"--api-key", "apiKey",
			"--context-window-messages", "1",
			"--enabled=true",
			"--include-contact-metadata=true",
			"--max-tokens", "1",
			"--model", "model",
			"--name", "name",
			"--provider", "openai",
			"--system-prompt", "systemPrompt",
			"--temperature", "0",
			"--trigger-on-channel", "string",
			"--trigger-on-message-type", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"apiKey: apiKey\n" +
			"contextWindowMessages: 1\n" +
			"enabled: true\n" +
			"includeContactMetadata: true\n" +
			"maxTokens: 1\n" +
			"model: model\n" +
			"name: name\n" +
			"provider: openai\n" +
			"systemPrompt: systemPrompt\n" +
			"temperature: 0\n" +
			"triggerOnChannels:\n" +
			"  - string\n" +
			"triggerOnMessageTypes:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "senders:agent", "update",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersAgentDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders:agent", "delete",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}

func TestSendersAgentStats(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "senders:agent", "stats",
			"--api-key", "string",
			"--sender-id", "senderId",
		)
	})
}
