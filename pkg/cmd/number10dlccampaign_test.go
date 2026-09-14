// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestNumber10dlcCampaignsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns", "create",
			"--affiliate-marketing=false",
			"--age-gated=false",
			"--brand-id", "brand_abc123",
			"--description", "Send order status updates and shipping notifications to customers who opted in.",
			"--direct-lending=false",
			"--embedded-link=true",
			"--embedded-phone=false",
			"--name", "Order Notifications",
			"--number-pooling=false",
			"--sample-message", "Hi {{name}}, your order #{{order_id}} has shipped! Track it at {{url}}",
			"--sample-message", "Your order #{{order_id}} has been delivered. Thank you for your purchase!",
			"--subscriber-help=true",
			"--subscriber-opt-in=true",
			"--subscriber-opt-out=true",
			"--use-case", "ACCOUNT_NOTIFICATION",
			"--help-message", "helpMessage",
			"--message-flow", "messageFlow",
			"--opt-in-keyword", "string",
			"--opt-out-keyword", "string",
			"--sub-use-case", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"affiliateMarketing: false\n" +
			"ageGated: false\n" +
			"brandId: brand_abc123\n" +
			"description: >-\n" +
			"  Send order status updates and shipping notifications to customers who opted\n" +
			"  in.\n" +
			"directLending: false\n" +
			"embeddedLink: true\n" +
			"embeddedPhone: false\n" +
			"name: Order Notifications\n" +
			"numberPooling: false\n" +
			"sampleMessages:\n" +
			"  - 'Hi {{name}}, your order #{{order_id}} has shipped! Track it at {{url}}'\n" +
			"  - 'Your order #{{order_id}} has been delivered. Thank you for your purchase!'\n" +
			"subscriberHelp: true\n" +
			"subscriberOptIn: true\n" +
			"subscriberOptOut: true\n" +
			"useCase: ACCOUNT_NOTIFICATION\n" +
			"helpMessage: helpMessage\n" +
			"messageFlow: messageFlow\n" +
			"optInKeywords:\n" +
			"  - string\n" +
			"optOutKeywords:\n" +
			"  - string\n" +
			"subUseCases:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"number-10dlc:campaigns", "create",
		)
	})
}

func TestNumber10dlcCampaignsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns", "retrieve",
			"--campaign-id", "campaignId",
		)
	})
}

func TestNumber10dlcCampaignsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns", "update",
			"--campaign-id", "campaignId",
			"--description", "description",
			"--help-message", "helpMessage",
			"--message-flow", "messageFlow",
			"--name", "name",
			"--opt-in-keyword", "string",
			"--opt-out-keyword", "string",
			"--sample-message", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"helpMessage: helpMessage\n" +
			"messageFlow: messageFlow\n" +
			"name: name\n" +
			"optInKeywords:\n" +
			"  - string\n" +
			"optOutKeywords:\n" +
			"  - string\n" +
			"sampleMessages:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"number-10dlc:campaigns", "update",
			"--campaign-id", "campaignId",
		)
	})
}

func TestNumber10dlcCampaignsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns", "list",
			"--max-items", "10",
			"--brand-id", "brandId",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestNumber10dlcCampaignsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns", "delete",
			"--campaign-id", "campaignId",
		)
	})
}

func TestNumber10dlcCampaignsSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns", "submit",
			"--campaign-id", "campaignId",
		)
	})
}

func TestNumber10dlcCampaignsSyncStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:campaigns", "sync-status",
			"--campaign-id", "campaignId",
		)
	})
}
