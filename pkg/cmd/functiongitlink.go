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

var functionsGitLinkRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "The link and its last deploy. Never returns the webhook secret.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
	},
	Action:          handleFunctionsGitLinkRetrieve,
	HideHelpCommand: true,
}

var functionsGitLinkUpdate = cli.Command{
	Name:    "update",
	Usage:   "Change the branch, the root directory, or whether pushes deploy. Pass at least\none field. `rootDir: null` clears the subdirectory.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-deploy",
			BodyPath: "autoDeploy",
		},
		&requestflag.Flag[string]{
			Name:     "branch",
			BodyPath: "branch",
		},
		&requestflag.Flag[*string]{
			Name:     "root-dir",
			BodyPath: "rootDir",
		},
	},
	Action:          handleFunctionsGitLinkUpdate,
	HideHelpCommand: true,
}

var functionsGitLinkDeployNow = cli.Command{
	Name:    "deploy-now",
	Usage:   "Fetch the linked branch and deploy it without waiting for a push. Returns\nimmediately; follow the outcome with `GET /v1/functions/{functionId}/git-link`,\nwhose `lastStatus` and `lastError` describe the run.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
	},
	Action:          handleFunctionsGitLinkDeployNow,
	HideHelpCommand: true,
}

var functionsGitLinkLink = cli.Command{
	Name:    "link",
	Usage:   "Bind a repository to this function so every push to `branch` deploys it. A\nfunction holds at most one link; linking again returns 400.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Required: true,
			BodyPath: "owner",
		},
		&requestflag.Flag[string]{
			Name:     "repo",
			Required: true,
			BodyPath: "repo",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-deploy",
			Default:  true,
			BodyPath: "autoDeploy",
		},
		&requestflag.Flag[string]{
			Name:     "branch",
			Default:  "main",
			BodyPath: "branch",
		},
		&requestflag.Flag[string]{
			Name:     "root-dir",
			Usage:    "Subdirectory holding the project, for monorepos.",
			BodyPath: "rootDir",
		},
	},
	Action:          handleFunctionsGitLinkLink,
	HideHelpCommand: true,
}

var functionsGitLinkUnlink = cli.Command{
	Name:    "unlink",
	Usage:   "Remove the link. The function and its deployments stay. A manual webhook left in\nthe repository stops being accepted, so remove it there too.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
	},
	Action:          handleFunctionsGitLinkUnlink,
	HideHelpCommand: true,
}

func handleFunctionsGitLinkRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Functions.GitLink.Get(ctx, cmd.Value("function-id").(string), options...)
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
		Title:          "functions:git-link retrieve",
		Transform:      transform,
	})
}

func handleFunctionsGitLinkUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.FunctionGitLinkUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.GitLink.Update(
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
		Title:          "functions:git-link update",
		Transform:      transform,
	})
}

func handleFunctionsGitLinkDeployNow(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Functions.GitLink.DeployNow(ctx, cmd.Value("function-id").(string), options...)
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
		Title:          "functions:git-link deploy-now",
		Transform:      transform,
	})
}

func handleFunctionsGitLinkLink(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.FunctionGitLinkLinkParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.GitLink.Link(
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
		Title:          "functions:git-link link",
		Transform:      transform,
	})
}

func handleFunctionsGitLinkUnlink(ctx context.Context, cmd *cli.Command) error {
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

	return client.Functions.GitLink.Unlink(ctx, cmd.Value("function-id").(string), options...)
}
