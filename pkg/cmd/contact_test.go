// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
	"github.com/zavudev/cli/internal/requestflag"
)

func TestContactsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts", "create",
			"--channel", "{channel: sms, identifier: '+14155551234', countryCode: US, isPrimary: true, label: work}",
			"--display-name", "John Doe",
			"--metadata", "{foo: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(contactsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts", "create",
			"--channel.channel", "sms",
			"--channel.identifier", "+14155551234",
			"--channel.country-code", "US",
			"--channel.is-primary=true",
			"--channel.label", "work",
			"--display-name", "John Doe",
			"--metadata", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"channels:\n" +
			"  - channel: sms\n" +
			"    identifier: '+14155551234'\n" +
			"    countryCode: US\n" +
			"    isPrimary: true\n" +
			"    label: work\n" +
			"displayName: John Doe\n" +
			"metadata:\n" +
			"  foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"contacts", "create",
		)
	})
}

func TestContactsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts", "retrieve",
			"--contact-id", "contactId",
		)
	})
}

func TestContactsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts", "update",
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
			t, pipeData,
			"--api-key", "string",
			"contacts", "update",
			"--contact-id", "contactId",
		)
	})
}

func TestContactsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
			"--phone-number", "phoneNumber",
		)
	})
}

func TestContactsDismissMergeSuggestion(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts", "dismiss-merge-suggestion",
			"--contact-id", "contactId",
		)
	})
}

func TestContactsMerge(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts", "merge",
			"--contact-id", "contactId",
			"--source-contact-id", "jx7xyz789",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("sourceContactId: jx7xyz789")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"contacts", "merge",
			"--contact-id", "contactId",
		)
	})
}

func TestContactsRetrieveByPhone(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts", "retrieve-by-phone",
			"--phone-number", "phoneNumber",
		)
	})
}
