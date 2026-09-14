// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestContactsChannelsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts:channels", "update",
			"--contact-id", "contactId",
			"--channel-id", "channelId",
			"--label", "label",
			"--metadata", "{foo: string}",
			"--verified=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"label: label\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"verified: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"contacts:channels", "update",
			"--contact-id", "contactId",
			"--channel-id", "channelId",
		)
	})
}

func TestContactsChannelsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts:channels", "add",
			"--contact-id", "contactId",
			"--channel", "email",
			"--identifier", "john.work@company.com",
			"--country-code", "US",
			"--is-primary=true",
			"--label", "work",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"channel: email\n" +
			"identifier: john.work@company.com\n" +
			"countryCode: US\n" +
			"isPrimary: true\n" +
			"label: work\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"contacts:channels", "add",
			"--contact-id", "contactId",
		)
	})
}

func TestContactsChannelsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts:channels", "remove",
			"--contact-id", "contactId",
			"--channel-id", "channelId",
		)
	})
}

func TestContactsChannelsSetPrimary(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"contacts:channels", "set-primary",
			"--contact-id", "contactId",
			"--channel-id", "channelId",
		)
	})
}
