// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/zavudev/cli/internal/mocktest"
)

func TestFunctionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "create",
			"--name", "Order Bot",
			"--slug", "order-bot",
			"--dependencies", "{openai: ^4.20.0}",
			"--description", "Replies to order status questions on WhatsApp.",
			"--http-enabled=true",
			"--memory-mb", "128",
			"--runtime", "nodejs24",
			"--source-code", "import { defineFunction } from '@zavudev/functions';\n\nexport default defineFunction(async (event, ctx) => {\n  ctx.log('received', event.type);\n});\n",
			"--timeout-sec", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Order Bot\n" +
			"slug: order-bot\n" +
			"dependencies:\n" +
			"  openai: ^4.20.0\n" +
			"description: Replies to order status questions on WhatsApp.\n" +
			"httpEnabled: true\n" +
			"memoryMb: 128\n" +
			"runtime: nodejs24\n" +
			"sourceCode: |\n" +
			"  import { defineFunction } from '@zavudev/functions';\n" +
			"\n" +
			"  export default defineFunction(async (event, ctx) => {\n" +
			"    ctx.log('received', event.type);\n" +
			"  });\n" +
			"timeoutSec: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions", "create",
		)
	})
}

func TestFunctionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "retrieve",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "update",
			"--function-id", "functionId",
			"--dependencies", "{foo: string}",
			"--http-enabled=true",
			"--source-code", "sourceCode",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"dependencies:\n" +
			"  foo: string\n" +
			"httpEnabled: true\n" +
			"sourceCode: sourceCode\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions", "update",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "delete",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsDeploy(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "deploy",
			"--function-id", "functionId",
			"--dependencies", "{foo: string}",
			"--source-code", "sourceCode",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"dependencies:\n" +
			"  foo: string\n" +
			"sourceCode: sourceCode\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions", "deploy",
			"--function-id", "functionId",
		)
	})
}

func TestFunctionsGetDeployment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "get-deployment",
			"--deployment-id", "deploymentId",
		)
	})
}

func TestFunctionsTailLogs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "tail-logs",
			"--function-id", "functionId",
			"--end-time", "0",
			"--filter-pattern", "filterPattern",
			"--limit", "1",
			"--next-token", "nextToken",
			"--start-time", "0",
		)
	})
}
