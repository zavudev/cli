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

var functionsTriggersCreate = cli.Command{
	Name:    "create",
	Usage:   "Subscribe a function to one or more event types, optionally scoped to specific\nsenders. Provide eventTypes and senderIds (use null in senderIds for all\nsenders); a trigger is created for each event type and sender combination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
		&requestflag.Flag[[]string]{
			Name:     "event-type",
			Usage:    "Event types to subscribe to.",
			Required: true,
			BodyPath: "eventTypes",
		},
		&requestflag.Flag[[]string]{
			Name:     "sender-id",
			Usage:    "Senders to scope the triggers to. Use null for all senders.",
			Required: true,
			BodyPath: "senderIds",
		},
		&requestflag.Flag[string]{
			Name:     "cron",
			Usage:    "Required when eventTypes includes `cron`: a 5-field cron expression (minute hour day-of-month month day-of-week), evaluated in UTC.",
			BodyPath: "cron",
		},
	},
	Action:          handleFunctionsTriggersCreate,
	HideHelpCommand: true,
}

var functionsTriggersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Enable or disable a trigger",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "trigger-id",
			Required:  true,
			PathParam: "triggerId",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Required: true,
			BodyPath: "active",
		},
	},
	Action:          handleFunctionsTriggersUpdate,
	HideHelpCommand: true,
}

var functionsTriggersList = cli.Command{
	Name:    "list",
	Usage:   "List function triggers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-id",
			Required:  true,
			PathParam: "functionId",
		},
	},
	Action:          handleFunctionsTriggersList,
	HideHelpCommand: true,
}

var functionsTriggersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a trigger",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "trigger-id",
			Required:  true,
			PathParam: "triggerId",
		},
	},
	Action:          handleFunctionsTriggersDelete,
	HideHelpCommand: true,
}

func handleFunctionsTriggersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.FunctionTriggerNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Triggers.New(
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
		Title:          "functions:triggers create",
		Transform:      transform,
	})
}

func handleFunctionsTriggersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("trigger-id") && len(unusedArgs) > 0 {
		cmd.Set("trigger-id", unusedArgs[0])
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

	params := zavudev.FunctionTriggerUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Triggers.Update(
		ctx,
		cmd.Value("trigger-id").(string),
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
		Title:          "functions:triggers update",
		Transform:      transform,
	})
}

func handleFunctionsTriggersList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Functions.Triggers.List(ctx, cmd.Value("function-id").(string), options...)
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
		Title:          "functions:triggers list",
		Transform:      transform,
	})
}

func handleFunctionsTriggersDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("trigger-id") && len(unusedArgs) > 0 {
		cmd.Set("trigger-id", unusedArgs[0])
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

	return client.Functions.Triggers.Delete(ctx, cmd.Value("trigger-id").(string), options...)
}
