// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/zavudev/cli/internal/apiquery"
	"github.com/zavudev/cli/internal/requestflag"
	"github.com/zavudev/sdk-go"
	"github.com/zavudev/sdk-go/option"
)

var sendersAgentToolsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new tool for an agent. Tools allow the agent to call external webhooks.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parameters",
			Required: true,
			BodyPath: "parameters",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Must be HTTPS.",
			Required: true,
			BodyPath: "webhookUrl",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Default:  true,
			BodyPath: "enabled",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-secret",
			Usage:    "Optional secret for webhook signature verification.",
			BodyPath: "webhookSecret",
		},
	},
	Action:          handleSendersAgentToolsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"parameters": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "parameters.properties",
			InnerField: "properties",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "parameters.required",
			InnerField: "required",
		},
		&requestflag.InnerFlag[string]{
			Name:       "parameters.type",
			Usage:      `Allowed values: "object".`,
			InnerField: "type",
		},
	},
})

var sendersAgentToolsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a specific tool.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
	},
	Action:          handleSendersAgentToolsRetrieve,
	HideHelpCommand: true,
}

var sendersAgentToolsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a tool.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parameters",
			BodyPath: "parameters",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-secret",
			BodyPath: "webhookSecret",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			BodyPath: "webhookUrl",
		},
	},
	Action:          handleSendersAgentToolsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"parameters": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "parameters.properties",
			InnerField: "properties",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "parameters.required",
			InnerField: "required",
		},
		&requestflag.InnerFlag[string]{
			Name:       "parameters.type",
			Usage:      `Allowed values: "object".`,
			InnerField: "type",
		},
	},
})

var sendersAgentToolsList = cli.Command{
	Name:    "list",
	Usage:   "List tools for an agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[bool]{
			Name:      "enabled",
			QueryPath: "enabled",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSendersAgentToolsList,
	HideHelpCommand: true,
}

var sendersAgentToolsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a tool.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
	},
	Action:          handleSendersAgentToolsDelete,
	HideHelpCommand: true,
}

var sendersAgentToolsTest = cli.Command{
	Name:    "test",
	Usage:   "Test a tool by triggering its webhook with test parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "test-params",
			Usage:    "Parameters to pass to the tool for testing.",
			Required: true,
			BodyPath: "testParams",
		},
	},
	Action:          handleSendersAgentToolsTest,
	HideHelpCommand: true,
}

func handleSendersAgentToolsCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentToolNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Agent.Tools.New(
		ctx,
		cmd.Value("sender-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "senders:agent:tools create", obj, format, explicitFormat, transform)
}

func handleSendersAgentToolsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentToolGetParams{
		SenderID: cmd.Value("sender-id").(string),
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
	_, err = client.Senders.Agent.Tools.Get(
		ctx,
		cmd.Value("tool-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "senders:agent:tools retrieve", obj, format, explicitFormat, transform)
}

func handleSendersAgentToolsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentToolUpdateParams{
		SenderID: cmd.Value("sender-id").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Agent.Tools.Update(
		ctx,
		cmd.Value("tool-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "senders:agent:tools update", obj, format, explicitFormat, transform)
}

func handleSendersAgentToolsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentToolListParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Senders.Agent.Tools.List(
			ctx,
			cmd.Value("sender-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "senders:agent:tools list", obj, format, explicitFormat, transform)
	} else {
		iter := client.Senders.Agent.Tools.ListAutoPaging(
			ctx,
			cmd.Value("sender-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "senders:agent:tools list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleSendersAgentToolsDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentToolDeleteParams{
		SenderID: cmd.Value("sender-id").(string),
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

	return client.Senders.Agent.Tools.Delete(
		ctx,
		cmd.Value("tool-id").(string),
		params,
		options...,
	)
}

func handleSendersAgentToolsTest(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentToolTestParams{
		SenderID: cmd.Value("sender-id").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Agent.Tools.Test(
		ctx,
		cmd.Value("tool-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "senders:agent:tools test", obj, format, explicitFormat, transform)
}
