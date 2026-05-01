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

var sendersAgentCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an AI agent for a sender. Each sender can have at most one agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "LLM provider for the AI agent.",
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "system-prompt",
			Required: true,
			BodyPath: "systemPrompt",
		},
		&requestflag.Flag[string]{
			Name:     "api-key",
			Usage:    "API key for the LLM provider. Required unless provider is 'zavu'.",
			BodyPath: "apiKey",
		},
		&requestflag.Flag[int64]{
			Name:     "context-window-messages",
			Default:  10,
			BodyPath: "contextWindowMessages",
		},
		&requestflag.Flag[bool]{
			Name:     "include-contact-metadata",
			Default:  true,
			BodyPath: "includeContactMetadata",
		},
		&requestflag.Flag[int64]{
			Name:     "max-tokens",
			BodyPath: "maxTokens",
		},
		&requestflag.Flag[float64]{
			Name:     "temperature",
			BodyPath: "temperature",
		},
		&requestflag.Flag[[]string]{
			Name:     "trigger-on-channel",
			Default:  []string{"*"},
			BodyPath: "triggerOnChannels",
		},
		&requestflag.Flag[[]string]{
			Name:     "trigger-on-message-type",
			Default:  []string{"text"},
			BodyPath: "triggerOnMessageTypes",
		},
	},
	Action:          handleSendersAgentCreate,
	HideHelpCommand: true,
}

var sendersAgentRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get the AI agent configuration for a sender.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersAgentRetrieve,
	HideHelpCommand: true,
}

var sendersAgentUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an AI agent's configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
		&requestflag.Flag[string]{
			Name:     "api-key",
			BodyPath: "apiKey",
		},
		&requestflag.Flag[int64]{
			Name:     "context-window-messages",
			BodyPath: "contextWindowMessages",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "include-contact-metadata",
			BodyPath: "includeContactMetadata",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-tokens",
			BodyPath: "maxTokens",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "LLM provider for the AI agent.",
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "system-prompt",
			BodyPath: "systemPrompt",
		},
		&requestflag.Flag[*float64]{
			Name:     "temperature",
			BodyPath: "temperature",
		},
		&requestflag.Flag[[]string]{
			Name:     "trigger-on-channel",
			BodyPath: "triggerOnChannels",
		},
		&requestflag.Flag[[]string]{
			Name:     "trigger-on-message-type",
			BodyPath: "triggerOnMessageTypes",
		},
	},
	Action:          handleSendersAgentUpdate,
	HideHelpCommand: true,
}

var sendersAgentDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an AI agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersAgentDelete,
	HideHelpCommand: true,
}

var sendersAgentStats = cli.Command{
	Name:    "stats",
	Usage:   "Get statistics for an AI agent including invocations, tokens, and costs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersAgentStats,
	HideHelpCommand: true,
}

func handleSendersAgentCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
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

	params := zavudev.SenderAgentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Agent.New(
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
		Title:          "senders:agent create",
		Transform:      transform,
	})
}

func handleSendersAgentRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
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
	_, err = client.Senders.Agent.Get(ctx, cmd.Value("sender-id").(string), options...)
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
		Title:          "senders:agent retrieve",
		Transform:      transform,
	})
}

func handleSendersAgentUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
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

	params := zavudev.SenderAgentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Agent.Update(
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
		Title:          "senders:agent update",
		Transform:      transform,
	})
}

func handleSendersAgentDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
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

	return client.Senders.Agent.Delete(ctx, cmd.Value("sender-id").(string), options...)
}

func handleSendersAgentStats(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
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
	_, err = client.Senders.Agent.Stats(ctx, cmd.Value("sender-id").(string), options...)
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
		Title:          "senders:agent stats",
		Transform:      transform,
	})
}
