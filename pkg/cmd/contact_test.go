// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
)

func TestContactsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contacts", "retrieve",
			"--api-key", "string",
			"--contact-id", "contactId",
		)
	})
}

func TestContactsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contacts", "update",
			"--api-key", "string",
			"--contact-id", "contactId",
			"--default-channel", "sms",
			"--metadata", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"defaultChannel: sms\n" +
			"metadata:\n" +
			"  foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "contacts", "update",
			"--api-key", "string",
			"--contact-id", "contactId",
		)
	})
}

func TestContactsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contacts", "list",
			"--api-key", "string",
			"--cursor", "cursor",
			"--limit", "100",
			"--phone-number", "phoneNumber",
		)
	})
}

func TestContactsRetrieveByPhone(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contacts", "retrieve-by-phone",
			"--api-key", "string",
			"--phone-number", "phoneNumber",
		)
	})
}
