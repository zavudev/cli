// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
)

func TestTemplatesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "templates", "create",
			"--api-key", "string",
			"--body", "Hi {{1}}, your order {{2}} has been confirmed and will ship within 24 hours.",
			"--language", "en",
			"--name", "order_confirmation",
			"--add-security-recommendation=true",
			"--button", "{text: text, type: quick_reply, otpType: COPY_CODE, packageName: packageName, phoneNumber: phoneNumber, signatureHash: signatureHash, url: https://example.com}",
			"--code-expiration-minutes", "1",
			"--variable", "customer_name",
			"--variable", "order_id",
			"--whatsapp-category", "UTILITY",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(templatesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "templates", "create",
			"--api-key", "string",
			"--body", "Hi {{1}}, your order {{2}} has been confirmed and will ship within 24 hours.",
			"--language", "en",
			"--name", "order_confirmation",
			"--add-security-recommendation=true",
			"--button.text", "text",
			"--button.type", "quick_reply",
			"--button.otp-type", "COPY_CODE",
			"--button.package-name", "packageName",
			"--button.phone-number", "phoneNumber",
			"--button.signature-hash", "signatureHash",
			"--button.url", "https://example.com",
			"--code-expiration-minutes", "1",
			"--variable", "customer_name",
			"--variable", "order_id",
			"--whatsapp-category", "UTILITY",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"body: Hi {{1}}, your order {{2}} has been confirmed and will ship within 24 hours.\n" +
			"language: en\n" +
			"name: order_confirmation\n" +
			"addSecurityRecommendation: true\n" +
			"buttons:\n" +
			"  - text: text\n" +
			"    type: quick_reply\n" +
			"    otpType: COPY_CODE\n" +
			"    packageName: packageName\n" +
			"    phoneNumber: phoneNumber\n" +
			"    signatureHash: signatureHash\n" +
			"    url: https://example.com\n" +
			"codeExpirationMinutes: 1\n" +
			"variables:\n" +
			"  - customer_name\n" +
			"  - order_id\n" +
			"whatsappCategory: UTILITY\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "templates", "create",
			"--api-key", "string",
		)
	})
}

func TestTemplatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "templates", "retrieve",
			"--api-key", "string",
			"--template-id", "templateId",
		)
	})
}

func TestTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "templates", "list",
			"--api-key", "string",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestTemplatesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "templates", "delete",
			"--api-key", "string",
			"--template-id", "templateId",
		)
	})
}

func TestTemplatesSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "templates", "submit",
			"--api-key", "string",
			"--template-id", "templateId",
			"--sender-id", "sender_abc123",
			"--category", "UTILITY",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"senderId: sender_abc123\n" +
			"category: UTILITY\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "templates", "submit",
			"--api-key", "string",
			"--template-id", "templateId",
		)
	})
}
