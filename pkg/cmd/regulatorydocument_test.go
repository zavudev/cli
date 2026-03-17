// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestRegulatoryDocumentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"regulatory-documents", "create",
			"--document-type", "passport",
			"--file-size", "102400",
			"--mime-type", "image/jpeg",
			"--name", "Passport Scan",
			"--storage-id", "kg2abc123...",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"documentType: passport\n" +
			"fileSize: 102400\n" +
			"mimeType: image/jpeg\n" +
			"name: Passport Scan\n" +
			"storageId: kg2abc123...\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"regulatory-documents", "create",
		)
	})
}

func TestRegulatoryDocumentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"regulatory-documents", "retrieve",
			"--document-id", "documentId",
		)
	})
}

func TestRegulatoryDocumentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"regulatory-documents", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestRegulatoryDocumentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"regulatory-documents", "delete",
			"--document-id", "documentId",
		)
	})
}

func TestRegulatoryDocumentsUploadURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"regulatory-documents", "upload-url",
		)
	})
}
