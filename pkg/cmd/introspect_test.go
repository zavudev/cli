// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestIntrospectValidatePhone(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"introspect", "validate-phone",
		"--api-key", "string",
		"--phone-number", "+56912345678",
	)
}
