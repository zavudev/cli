// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestFunctionsGitLinkRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:git-link", "retrieve",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsGitLinkUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:git-link", "update",
			"--function-id", "functionId",
			"--auto-deploy=false",
			"--branch", "branch",
			"--root-dir", "rootDir",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"autoDeploy: false\n" +
			"branch: branch\n" +
			"rootDir: rootDir\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions:git-link", "update",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsGitLinkDeployNow(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:git-link", "deploy-now",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsGitLinkLink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:git-link", "link",
			"--function-id", "functionId",
			"--owner", "acme",
			"--repo", "order-bot",
			"--auto-deploy=true",
			"--branch", "main",
			"--root-dir", "apps/bot",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"owner: acme\n" +
			"repo: order-bot\n" +
			"autoDeploy: true\n" +
			"branch: main\n" +
			"rootDir: apps/bot\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions:git-link", "link",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsGitLinkUnlink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:git-link", "unlink",
			"--function-id", "functionId",
		)
	})
}
