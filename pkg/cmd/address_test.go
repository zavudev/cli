// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestAddressesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "create",
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
}

func TestAddressesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "retrieve",
		"--api-key", "string",
		"--address-id", "addressId",
	)
}

func TestAddressesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--limit", "100",
	)
}

func TestAddressesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "delete",
		"--api-key", "string",
		"--address-id", "addressId",
	)
}
