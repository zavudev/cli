// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestIntrospectValidatePhone(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "introspect", "validate-phone",
			"--api-key", "string",
			"--phone-number", "+56912345678",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("phoneNumber: '+56912345678'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "introspect", "validate-phone",
			"--api-key", "string",
		)
	})
}
