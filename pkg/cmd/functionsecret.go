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

var functionsSecretsList = cli.Command{
	Name:    "list",
	Usage:   "Lists every secret key set on the function. Plaintext is NEVER returned — only\nthe last 4 characters of each value, for visual confirmation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
	},
	Action:          handleFunctionsSecretsList,
	HideHelpCommand: true,
}

var functionsSecretsSet = cli.Command{
	Name:    "set",
	Usage:   "Create or update a secret on a function. Marks the function out-of-sync; the\nnext `POST /deploy` re-publishes the Lambda with the new env. Keys must match\n`[A-Z_][A-Z0-9_]*` (uppercase env-var style) and cannot start with reserved\nprefixes (AWS*, LAMBDA*, etc).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
		&requestflag.Flag[string]{
			Name:      "key",
			Required:  true,
			PathParam: "key",
		},
		&requestflag.Flag[string]{
			Name:     "value",
			Required: true,
			BodyPath: "value",
		},
	},
	Action:          handleFunctionsSecretsSet,
	HideHelpCommand: true,
}

var functionsSecretsUnset = cli.Command{
	Name:    "unset",
	Usage:   "Remove a secret from a function. Doesn't take effect on the running Lambda until\nthe next deploy.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
		&requestflag.Flag[string]{
			Name:      "key",
			Required:  true,
			PathParam: "key",
		},
	},
	Action:          handleFunctionsSecretsUnset,
	HideHelpCommand: true,
}

func handleFunctionsSecretsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Functions.Secrets.List(ctx, cmd.Value("function-id").(string), options...)
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
		Title:          "functions:secrets list",
		Transform:      transform,
	})
}

func handleFunctionsSecretsSet(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key") && len(unusedArgs) > 0 {
		cmd.Set("key", unusedArgs[0])
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

	params := zavudev.FunctionSecretSetParams{
		FunctionID: cmd.Value("function-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Secrets.Set(
		ctx,
		cmd.Value("key").(string),
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
		Title:          "functions:secrets set",
		Transform:      transform,
	})
}

func handleFunctionsSecretsUnset(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key") && len(unusedArgs) > 0 {
		cmd.Set("key", unusedArgs[0])
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

	params := zavudev.FunctionSecretUnsetParams{
		FunctionID: cmd.Value("function-id").(string),
	}

	return client.Functions.Secrets.Unset(
		ctx,
		cmd.Value("key").(string),
		params,
		options...,
	)
}
