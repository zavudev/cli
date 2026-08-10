// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/zavudev/cli/internal/apiquery"
	"github.com/zavudev/cli/internal/requestflag"
	"github.com/zavudev/sdk-go"
	"github.com/zavudev/sdk-go/option"
)

var functionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new Zavu Function. The function starts in `draft` status. A dedicated\nAPI key is auto-provisioned and injected as the `ZAVU_API_KEY` secret so the\nfunction can call back into the Zavu API without manual setup.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			Usage:    "URL-safe identifier (lowercase, digits, hyphens). Must be unique per project.",
			Required: true,
			BodyPath: "slug",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dependencies",
			Usage:    "npm dependencies. Keys are package names, values are semver ranges.",
			BodyPath: "dependencies",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "http-enabled",
			Usage:    "Whether to expose a public HTTPS URL for this function.",
			Default:  false,
			BodyPath: "httpEnabled",
		},
		&requestflag.Flag[int64]{
			Name:     "memory-mb",
			Usage:    "Allowed values: 128, 256, 512, 1024.",
			Default:  256,
			BodyPath: "memoryMb",
		},
		&requestflag.Flag[string]{
			Name:     "runtime",
			Usage:    "Runtime the function is deployed on.",
			BodyPath: "runtime",
		},
		&requestflag.Flag[string]{
			Name:     "source-code",
			Usage:    "TypeScript source code for the function entry point (max ~900KB).",
			BodyPath: "sourceCode",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-sec",
			Usage:    "Per-invocation timeout in seconds. Event and cron invocations are asynchronous, so a long timeout only bounds cost; a tool called during a live conversation holds up the reply, and a function exposed over HTTP is additionally bounded by the platform's HTTP response limit.",
			Default:  30,
			BodyPath: "timeoutSec",
		},
	},
	Action:          handleFunctionsCreate,
	HideHelpCommand: true,
}

var functionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get function",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
	},
	Action:          handleFunctionsRetrieve,
	HideHelpCommand: true,
}

var functionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an existing function. `sourceCode` / `dependencies` edit the draft\nwithout triggering a build — they go live on the next\n`POST /v1/functions/{functionId}/deploy`. `httpEnabled` is applied to the\ndeployed function immediately, so turning the public endpoint on or off does not\nrequire a redeploy.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dependencies",
			Usage:    "New dependency map (replaces existing dependencies).",
			BodyPath: "dependencies",
		},
		&requestflag.Flag[bool]{
			Name:     "http-enabled",
			Usage:    "Expose the function on its public HTTPS URL, or take it down. Applies to the already-deployed function without redeploying; the URL is returned as `publicUrl`.",
			BodyPath: "httpEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "source-code",
			Usage:    "New source code for the draft (replaces it).",
			BodyPath: "sourceCode",
		},
	},
	Action:          handleFunctionsUpdate,
	HideHelpCommand: true,
}

var functionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently delete a function and cascade: triggers, secrets, deployment\nhistory, managed agents+tools, and revoke the auto-provisioned API key. The AWS\nLambda + log group are torn down asynchronously.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
	},
	Action:          handleFunctionsDelete,
	HideHelpCommand: true,
}

var functionsDeploy = cli.Command{
	Name:    "deploy",
	Usage:   "Publish the function. If `sourceCode` or `dependencies` are provided in the\nbody, they replace the current draft before deployment. Returns immediately with\na deployment ID — poll `GET /v1/functions/deployments/{deploymentId}` until\nstatus is `active` or `failed`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dependencies",
			Usage:    "New dependency map (replaces existing dependencies).",
			BodyPath: "dependencies",
		},
		&requestflag.Flag[string]{
			Name:     "source-code",
			Usage:    "New source code to publish (replaces the draft).",
			BodyPath: "sourceCode",
		},
	},
	Action:          handleFunctionsDeploy,
	HideHelpCommand: true,
}

var functionsGetDeployment = cli.Command{
	Name:    "get-deployment",
	Usage:   "Fetch a deployment to poll its status during a deploy.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "deployment-id",
			Required:  true,
			PathParam: "deploymentId",
		},
	},
	Action:          handleFunctionsGetDeployment,
	HideHelpCommand: true,
}

var functionsTailLogs = cli.Command{
	Name:    "tail-logs",
	Usage:   "Fetch invocation logs for a function. Logs are paginated via `nextToken`. Pass\n`startTime` / `endTime` (Unix epoch milliseconds) to bound the window, or\n`filterPattern` to filter messages.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
		&requestflag.Flag[int64]{
			Name:      "end-time",
			Usage:     "End of the log window in Unix epoch milliseconds.",
			QueryPath: "endTime",
		},
		&requestflag.Flag[string]{
			Name:      "filter-pattern",
			QueryPath: "filterPattern",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "next-token",
			QueryPath: "nextToken",
		},
		&requestflag.Flag[int64]{
			Name:      "start-time",
			Usage:     "Start of the log window in Unix epoch milliseconds.",
			QueryPath: "startTime",
		},
	},
	Action:          handleFunctionsTailLogs,
	HideHelpCommand: true,
}

func handleFunctionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := zavudev.FunctionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "functions create",
		Transform:      transform,
	})
}

func handleFunctionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-id") && len(unusedArgs) > 0 {
		cmd.Set("function-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Get(ctx, cmd.Value("function-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "functions retrieve",
		Transform:      transform,
	})
}

func handleFunctionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-id") && len(unusedArgs) > 0 {
		cmd.Set("function-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := zavudev.FunctionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Update(
		ctx,
		cmd.Value("function-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "functions update",
		Transform:      transform,
	})
}

func handleFunctionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-id") && len(unusedArgs) > 0 {
		cmd.Set("function-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Delete(ctx, cmd.Value("function-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "functions delete",
		Transform:      transform,
	})
}

func handleFunctionsDeploy(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-id") && len(unusedArgs) > 0 {
		cmd.Set("function-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := zavudev.FunctionDeployParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Deploy(
		ctx,
		cmd.Value("function-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "functions deploy",
		Transform:      transform,
	})
}

func handleFunctionsGetDeployment(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("deployment-id") && len(unusedArgs) > 0 {
		cmd.Set("deployment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.GetDeployment(ctx, cmd.Value("deployment-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "functions get-deployment",
		Transform:      transform,
	})
}

func handleFunctionsTailLogs(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-id") && len(unusedArgs) > 0 {
		cmd.Set("function-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := zavudev.FunctionTailLogsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.TailLogs(
		ctx,
		cmd.Value("function-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "functions tail-logs",
		Transform:      transform,
	})
}
