// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/zavudev-cli/internal/mocktest"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
)

func TestBroadcastsContactsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "broadcasts:contacts", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--broadcast-id", "broadcastId",
			"--cursor", "cursor",
			"--limit", "100",
			"--status", "pending",
		)
	})
}

func TestBroadcastsContactsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "broadcasts:contacts", "add",
			"--api-key", "string",
			"--broadcast-id", "broadcastId",
			"--contact", "{recipient: '+14155551234', templateVariables: {name: John, order_id: ORD-001}}",
			"--contact", "{recipient: '+14155555678', templateVariables: {name: Jane, order_id: ORD-002}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(broadcastsContactsAdd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "broadcasts:contacts", "add",
			"--api-key", "string",
			"--broadcast-id", "broadcastId",
			"--contact.recipient", "+14155551234",
			"--contact.template-variables", "{name: John, order_id: ORD-001}",
			"--contact.recipient", "+14155555678",
			"--contact.template-variables", "{name: Jane, order_id: ORD-002}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contacts:\n" +
			"  - recipient: '+14155551234'\n" +
			"    templateVariables:\n" +
			"      name: John\n" +
			"      order_id: ORD-001\n" +
			"  - recipient: '+14155555678'\n" +
			"    templateVariables:\n" +
			"      name: Jane\n" +
			"      order_id: ORD-002\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "broadcasts:contacts", "add",
			"--api-key", "string",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsContactsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "broadcasts:contacts", "remove",
			"--api-key", "string",
			"--broadcast-id", "broadcastId",
			"--contact-id", "contactId",
		)
	})
}
