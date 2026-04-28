// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestNumber10dlcBrandsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:brands", "create",
			"--city", "San Francisco",
			"--country", "US",
			"--display-name", "Acme Corp",
			"--email", "compliance@acme.com",
			"--entity-type", "PRIVATE_PROFIT",
			"--phone", "+14155551234",
			"--postal-code", "94102",
			"--state", "CA",
			"--street", "123 Main St",
			"--vertical", "Technology",
			"--company-name", "Acme Corporation",
			"--ein", "12-3456789",
			"--first-name", "firstName",
			"--last-name", "lastName",
			"--stock-exchange", "stockExchange",
			"--stock-symbol", "stockSymbol",
			"--website", "https://acme.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"city: San Francisco\n" +
			"country: US\n" +
			"displayName: Acme Corp\n" +
			"email: compliance@acme.com\n" +
			"entityType: PRIVATE_PROFIT\n" +
			"phone: '+14155551234'\n" +
			"postalCode: '94102'\n" +
			"state: CA\n" +
			"street: 123 Main St\n" +
			"vertical: Technology\n" +
			"companyName: Acme Corporation\n" +
			"ein: 12-3456789\n" +
			"firstName: firstName\n" +
			"lastName: lastName\n" +
			"stockExchange: stockExchange\n" +
			"stockSymbol: stockSymbol\n" +
			"website: https://acme.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"number-10dlc:brands", "create",
		)
	})
}

func TestNumber10dlcBrandsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:brands", "retrieve",
			"--brand-id", "brandId",
		)
	})
}

func TestNumber10dlcBrandsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:brands", "update",
			"--brand-id", "brandId",
			"--city", "city",
			"--company-name", "companyName",
			"--country", "xx",
			"--display-name", "displayName",
			"--ein", "ein",
			"--email", "dev@stainless.com",
			"--entity-type", "PRIVATE_PROFIT",
			"--first-name", "firstName",
			"--last-name", "lastName",
			"--phone", "phone",
			"--postal-code", "postalCode",
			"--state", "state",
			"--stock-exchange", "stockExchange",
			"--stock-symbol", "stockSymbol",
			"--street", "street",
			"--vertical", "vertical",
			"--website", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"city: city\n" +
			"companyName: companyName\n" +
			"country: xx\n" +
			"displayName: displayName\n" +
			"ein: ein\n" +
			"email: dev@stainless.com\n" +
			"entityType: PRIVATE_PROFIT\n" +
			"firstName: firstName\n" +
			"lastName: lastName\n" +
			"phone: phone\n" +
			"postalCode: postalCode\n" +
			"state: state\n" +
			"stockExchange: stockExchange\n" +
			"stockSymbol: stockSymbol\n" +
			"street: street\n" +
			"vertical: vertical\n" +
			"website: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"number-10dlc:brands", "update",
			"--brand-id", "brandId",
		)
	})
}

func TestNumber10dlcBrandsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:brands", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestNumber10dlcBrandsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:brands", "delete",
			"--brand-id", "brandId",
		)
	})
}

func TestNumber10dlcBrandsListUseCases(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:brands", "list-use-cases",
		)
	})
}

func TestNumber10dlcBrandsSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:brands", "submit",
			"--brand-id", "brandId",
		)
	})
}

func TestNumber10dlcBrandsSyncStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-10dlc:brands", "sync-status",
			"--brand-id", "brandId",
		)
	})
}
