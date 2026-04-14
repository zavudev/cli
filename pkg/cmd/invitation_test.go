// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestInvitationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invitations", "create",
			"--allowed-phone-country", "US",
			"--allowed-phone-country", "MX",
			"--client-email", "contact@acme.com",
			"--client-name", "Acme Corp",
			"--client-phone", "+14155551234",
			"--expires-in-days", "1",
			"--phone-number-id", "pn_abc123",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowedPhoneCountries:\n" +
			"  - US\n" +
			"  - MX\n" +
			"clientEmail: contact@acme.com\n" +
			"clientName: Acme Corp\n" +
			"clientPhone: '+14155551234'\n" +
			"expiresInDays: 1\n" +
			"phoneNumberId: pn_abc123\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"invitations", "create",
		)
	})
}

func TestInvitationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invitations", "retrieve",
			"--invitation-id", "invitationId",
		)
	})
}

func TestInvitationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invitations", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
			"--status", "pending",
		)
	})
}

func TestInvitationsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invitations", "cancel",
			"--invitation-id", "invitationId",
		)
	})
}
