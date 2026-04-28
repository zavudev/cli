// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
	"github.com/zavudev/cli/internal/requestflag"
)

func TestBroadcastsContactsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts:contacts", "list",
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
			t,
			"--api-key", "string",
			"broadcasts:contacts", "add",
			"--broadcast-id", "broadcastId",
			"--contact", "{recipient: '+14155551234', templateButtonVariables: {'0': abc-report-token}, templateVariables: {name: John, order_id: ORD-001}}",
			"--contact", "{recipient: '+14155555678', templateButtonVariables: {'0': abc-report-token}, templateVariables: {name: Jane, order_id: ORD-002}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(broadcastsContactsAdd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts:contacts", "add",
			"--broadcast-id", "broadcastId",
			"--contact.recipient", "+14155551234",
			"--contact.template-button-variables", "{'0': abc-report-token}",
			"--contact.template-variables", "{name: John, order_id: ORD-001}",
			"--contact.recipient", "+14155555678",
			"--contact.template-button-variables", "{'0': abc-report-token}",
			"--contact.template-variables", "{name: Jane, order_id: ORD-002}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contacts:\n" +
			"  - recipient: '+14155551234'\n" +
			"    templateButtonVariables:\n" +
			"      '0': abc-report-token\n" +
			"    templateVariables:\n" +
			"      name: John\n" +
			"      order_id: ORD-001\n" +
			"  - recipient: '+14155555678'\n" +
			"    templateButtonVariables:\n" +
			"      '0': abc-report-token\n" +
			"    templateVariables:\n" +
			"      name: Jane\n" +
			"      order_id: ORD-002\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"broadcasts:contacts", "add",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsContactsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts:contacts", "remove",
			"--broadcast-id", "broadcastId",
			"--contact-id", "contactId",
		)
	})
}
