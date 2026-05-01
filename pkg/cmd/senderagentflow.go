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

var sendersAgentFlowsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new flow for an agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			Required: true,
			BodyPath: "steps",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "trigger",
			Required: true,
			BodyPath: "trigger",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Default:  false,
			BodyPath: "enabled",
		},
		&requestflag.Flag[int64]{
			Name:     "priority",
			Default:  0,
			BodyPath: "priority",
		},
	},
	Action:          handleSendersAgentFlowsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"step": {
		&requestflag.InnerFlag[string]{
			Name:       "step.id",
			Usage:      "Unique step identifier.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "step.config",
			Usage:      "Step configuration (varies by type).",
			InnerField: "config",
		},
		&requestflag.InnerFlag[string]{
			Name:       "step.type",
			Usage:      "Type of flow step.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "step.next-step-id",
			Usage:      "ID of the next step to execute.",
			InnerField: "nextStepId",
		},
	},
	"trigger": {
		&requestflag.InnerFlag[string]{
			Name:       "trigger.type",
			Usage:      "Type of trigger for a flow.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trigger.intent",
			Usage:      "Intent that triggers the flow (for intent type).",
			InnerField: "intent",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "trigger.keywords",
			Usage:      "Keywords that trigger the flow (for keyword type).",
			InnerField: "keywords",
		},
	},
})

var sendersAgentFlowsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a specific flow.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "flow-id",
			Required: true,
		},
	},
	Action:          handleSendersAgentFlowsRetrieve,
	HideHelpCommand: true,
}

var sendersAgentFlowsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a flow.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "flow-id",
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
		&requestflag.Flag[int64]{
			Name:     "priority",
			BodyPath: "priority",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			BodyPath: "steps",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "trigger",
			BodyPath: "trigger",
		},
	},
	Action:          handleSendersAgentFlowsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"step": {
		&requestflag.InnerFlag[string]{
			Name:       "step.id",
			Usage:      "Unique step identifier.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "step.config",
			Usage:      "Step configuration (varies by type).",
			InnerField: "config",
		},
		&requestflag.InnerFlag[string]{
			Name:       "step.type",
			Usage:      "Type of flow step.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "step.next-step-id",
			Usage:      "ID of the next step to execute.",
			InnerField: "nextStepId",
		},
	},
	"trigger": {
		&requestflag.InnerFlag[string]{
			Name:       "trigger.type",
			Usage:      "Type of trigger for a flow.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trigger.intent",
			Usage:      "Intent that triggers the flow (for intent type).",
			InnerField: "intent",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "trigger.keywords",
			Usage:      "Keywords that trigger the flow (for keyword type).",
			InnerField: "keywords",
		},
	},
})

var sendersAgentFlowsList = cli.Command{
	Name:    "list",
	Usage:   "List flows for an agent.",
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
	Action:          handleSendersAgentFlowsList,
	HideHelpCommand: true,
}

var sendersAgentFlowsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a flow. Cannot delete flows with active sessions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "flow-id",
			Required: true,
		},
	},
	Action:          handleSendersAgentFlowsDelete,
	HideHelpCommand: true,
}

var sendersAgentFlowsDuplicate = cli.Command{
	Name:    "duplicate",
	Usage:   "Create a copy of an existing flow with a new name.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "flow-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "new-name",
			Required: true,
			BodyPath: "newName",
		},
	},
	Action:          handleSendersAgentFlowsDuplicate,
	HideHelpCommand: true,
}

func handleSendersAgentFlowsCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentFlowNewParams{}

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
	_, err = client.Senders.Agent.Flows.New(
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "senders:agent:flows create",
		Transform:      transform,
	})
}

func handleSendersAgentFlowsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("flow-id") && len(unusedArgs) > 0 {
		cmd.Set("flow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentFlowGetParams{
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
	_, err = client.Senders.Agent.Flows.Get(
		ctx,
		cmd.Value("flow-id").(string),
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
		Title:          "senders:agent:flows retrieve",
		Transform:      transform,
	})
}

func handleSendersAgentFlowsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("flow-id") && len(unusedArgs) > 0 {
		cmd.Set("flow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentFlowUpdateParams{
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
	_, err = client.Senders.Agent.Flows.Update(
		ctx,
		cmd.Value("flow-id").(string),
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
		Title:          "senders:agent:flows update",
		Transform:      transform,
	})
}

func handleSendersAgentFlowsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentFlowListParams{}

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
		_, err = client.Senders.Agent.Flows.List(
			ctx,
			cmd.Value("sender-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "senders:agent:flows list",
			Transform:      transform,
		})
	} else {
		iter := client.Senders.Agent.Flows.ListAutoPaging(
			ctx,
			cmd.Value("sender-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "senders:agent:flows list",
			Transform:      transform,
		})
	}
}

func handleSendersAgentFlowsDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("flow-id") && len(unusedArgs) > 0 {
		cmd.Set("flow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentFlowDeleteParams{
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

	return client.Senders.Agent.Flows.Delete(
		ctx,
		cmd.Value("flow-id").(string),
		params,
		options...,
	)
}

func handleSendersAgentFlowsDuplicate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("flow-id") && len(unusedArgs) > 0 {
		cmd.Set("flow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentFlowDuplicateParams{
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
	_, err = client.Senders.Agent.Flows.Duplicate(
		ctx,
		cmd.Value("flow-id").(string),
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
		Title:          "senders:agent:flows duplicate",
		Transform:      transform,
	})
}
