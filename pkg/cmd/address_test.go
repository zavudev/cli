// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestAddressesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "create",
			"--api-key", "string",
			"--country-code", "DE",
			"--locality", "Berlin",
			"--postal-code", "10115",
			"--street-address", "123 Main St",
			"--administrative-area", "administrativeArea",
			"--business-name", "businessName",
			"--extended-address", "extendedAddress",
			"--first-name", "John",
			"--last-name", "Doe",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"countryCode: DE\n" +
			"locality: Berlin\n" +
			"postalCode: '10115'\n" +
			"streetAddress: 123 Main St\n" +
			"administrativeArea: administrativeArea\n" +
			"businessName: businessName\n" +
			"extendedAddress: extendedAddress\n" +
			"firstName: John\n" +
			"lastName: Doe\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "addresses", "create",
			"--api-key", "string",
		)
	})
}

func TestAddressesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "retrieve",
			"--api-key", "string",
			"--address-id", "addressId",
		)
	})
}

func TestAddressesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestAddressesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "delete",
			"--api-key", "string",
			"--address-id", "addressId",
		)
	})
}
