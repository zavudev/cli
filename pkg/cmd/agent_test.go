// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
	"github.com/zavudev/cli/internal/requestflag"
)

func TestAgentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "create",
			"--model", "model",
			"--name", "name",
			"--provider", "openai",
			"--system-prompt", "systemPrompt",
			"--context-window-messages", "1",
			"--include-contact-metadata=true",
			"--max-tokens", "1",
			"--temperature", "0",
			"--trigger-on-channel", "string",
			"--trigger-on-message-type", "string",
			"--voice", "{enabled: true, greeting: 'Hi, thanks for calling Acme. How can I help you today?', greetings: {es: 'Hola, soy Atlas. Preguntame lo que quieras.'}, interruptible: true, language: en, maxCallDurationMinutes: 1, maxIdleSeconds: 5, model: openai/gpt-4o, recordCalls: true, sttModel: sttModel, sttProvider: sttProvider, transferPhoneNumber: '+14155551234', ttsProvider: ttsProvider, ttsVoiceId: aria, voicemailAction: hangup, voicemailMessage: voicemailMessage, voiceSpeed: 0.5}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "create",
			"--model", "model",
			"--name", "name",
			"--provider", "openai",
			"--system-prompt", "systemPrompt",
			"--context-window-messages", "1",
			"--include-contact-metadata=true",
			"--max-tokens", "1",
			"--temperature", "0",
			"--trigger-on-channel", "string",
			"--trigger-on-message-type", "string",
			"--voice.enabled=true",
			"--voice.greeting", "Hi, thanks for calling Acme. How can I help you today?",
			"--voice.greetings", "{es: 'Hola, soy Atlas. Preguntame lo que quieras.'}",
			"--voice.interruptible=true",
			"--voice.language", "en",
			"--voice.max-call-duration-minutes", "1",
			"--voice.max-idle-seconds", "5",
			"--voice.model", "openai/gpt-4o",
			"--voice.record-calls=true",
			"--voice.stt-model", "sttModel",
			"--voice.stt-provider", "sttProvider",
			"--voice.transfer-phone-number", "+14155551234",
			"--voice.tts-provider", "ttsProvider",
			"--voice.tts-voice-id", "aria",
			"--voice.voicemail-action", "hangup",
			"--voice.voicemail-message", "voicemailMessage",
			"--voice.voice-speed", "0.5",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"model: model\n" +
			"name: name\n" +
			"provider: openai\n" +
			"systemPrompt: systemPrompt\n" +
			"contextWindowMessages: 1\n" +
			"includeContactMetadata: true\n" +
			"maxTokens: 1\n" +
			"temperature: 0\n" +
			"triggerOnChannels:\n" +
			"  - string\n" +
			"triggerOnMessageTypes:\n" +
			"  - string\n" +
			"voice:\n" +
			"  enabled: true\n" +
			"  greeting: Hi, thanks for calling Acme. How can I help you today?\n" +
			"  greetings:\n" +
			"    es: Hola, soy Atlas. Preguntame lo que quieras.\n" +
			"  interruptible: true\n" +
			"  language: en\n" +
			"  maxCallDurationMinutes: 1\n" +
			"  maxIdleSeconds: 5\n" +
			"  model: openai/gpt-4o\n" +
			"  recordCalls: true\n" +
			"  sttModel: sttModel\n" +
			"  sttProvider: sttProvider\n" +
			"  transferPhoneNumber: '+14155551234'\n" +
			"  ttsProvider: ttsProvider\n" +
			"  ttsVoiceId: aria\n" +
			"  voicemailAction: hangup\n" +
			"  voicemailMessage: voicemailMessage\n" +
			"  voiceSpeed: 0.5\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "create",
		)
	})
}

func TestAgentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "retrieve",
			"--agent-id", "agentId",
		)
	})
}

func TestAgentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "agentId",
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
			"--voice", "{enabled: true, greeting: 'Hi, thanks for calling Acme. How can I help you today?', greetings: {es: 'Hola, soy Atlas. Preguntame lo que quieras.'}, interruptible: true, language: en, maxCallDurationMinutes: 1, maxIdleSeconds: 5, model: openai/gpt-4o, recordCalls: true, sttModel: sttModel, sttProvider: sttProvider, transferPhoneNumber: '+14155551234', ttsProvider: ttsProvider, ttsVoiceId: aria, voicemailAction: hangup, voicemailMessage: voicemailMessage, voiceSpeed: 0.5}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "agentId",
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
			"--voice.enabled=true",
			"--voice.greeting", "Hi, thanks for calling Acme. How can I help you today?",
			"--voice.greetings", "{es: 'Hola, soy Atlas. Preguntame lo que quieras.'}",
			"--voice.interruptible=true",
			"--voice.language", "en",
			"--voice.max-call-duration-minutes", "1",
			"--voice.max-idle-seconds", "5",
			"--voice.model", "openai/gpt-4o",
			"--voice.record-calls=true",
			"--voice.stt-model", "sttModel",
			"--voice.stt-provider", "sttProvider",
			"--voice.transfer-phone-number", "+14155551234",
			"--voice.tts-provider", "ttsProvider",
			"--voice.tts-voice-id", "aria",
			"--voice.voicemail-action", "hangup",
			"--voice.voicemail-message", "voicemailMessage",
			"--voice.voice-speed", "0.5",
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
			"  - string\n" +
			"voice:\n" +
			"  enabled: true\n" +
			"  greeting: Hi, thanks for calling Acme. How can I help you today?\n" +
			"  greetings:\n" +
			"    es: Hola, soy Atlas. Preguntame lo que quieras.\n" +
			"  interruptible: true\n" +
			"  language: en\n" +
			"  maxCallDurationMinutes: 1\n" +
			"  maxIdleSeconds: 5\n" +
			"  model: openai/gpt-4o\n" +
			"  recordCalls: true\n" +
			"  sttModel: sttModel\n" +
			"  sttProvider: sttProvider\n" +
			"  transferPhoneNumber: '+14155551234'\n" +
			"  ttsProvider: ttsProvider\n" +
			"  ttsVoiceId: aria\n" +
			"  voicemailAction: hangup\n" +
			"  voicemailMessage: voicemailMessage\n" +
			"  voiceSpeed: 0.5\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "agentId",
		)
	})
}

func TestAgentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestAgentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "delete",
			"--agent-id", "agentId",
		)
	})
}

func TestAgentsListVoices(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "list-voices",
			"--language", "es",
		)
	})
}

func TestAgentsTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "test",
			"--agent-id", "agentId",
			"--message", "Where is order ORD-12345?",
			"--execute-tools=true",
			"--history", "{content: content, role: user}",
			"--use-knowledge-base=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsTest)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "test",
			"--agent-id", "agentId",
			"--message", "Where is order ORD-12345?",
			"--execute-tools=true",
			"--history.content", "content",
			"--history.role", "user",
			"--use-knowledge-base=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"message: Where is order ORD-12345?\n" +
			"executeTools: true\n" +
			"history:\n" +
			"  - content: content\n" +
			"    role: user\n" +
			"useKnowledgeBase: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "test",
			"--agent-id", "agentId",
		)
	})
}
