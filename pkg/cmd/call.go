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

var callsCreate = cli.Command{
	Name:    "create",
	Usage:   "Place an outbound voice call answered by the voice agent configured on the\nsender. Zavu dials the recipient and runs the conversation through its managed\nvoice pipeline (speech recognition, the agent's LLM, and speech synthesis, with\nreal-time interruption handling).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Recipient phone number in E.164 format.",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[string]{
			Name:     "greeting",
			Usage:    "Overrides the agent's configured greeting for this call only.",
			BodyPath: "greeting",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "Language the agent speaks on this call only, as a BCP-47 tag (`en`, `es`, `es-ES`, `pt-BR`), or `auto` to detect the caller's language and follow it. Overrides the agent's configured language for speech recognition, the agent's replies, and the synthesized voice. If the agent uses a custom voice you supplied, that voice is kept and only the language changes. When omitted, the agent's configured language is used.",
			BodyPath: "language",
		},
		&requestflag.Flag[int64]{
			Name:     "max-duration-minutes",
			Usage:    "Overrides the agent's maximum call duration for this call only.",
			BodyPath: "maxDurationMinutes",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Arbitrary metadata to associate with the call. Returned on the call object and included in voice webhooks.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Usage:    "Sender profile that places the call. Uses the project's default sender if omitted. The sender's agent must have voice enabled.",
			BodyPath: "senderId",
		},
	},
	Action:          handleCallsCreate,
	HideHelpCommand: true,
}

var callsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single voice call, including its full transcript once the\nconversation has produced turns.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "call-id",
			Required:  true,
			PathParam: "callId",
		},
	},
	Action:          handleCallsRetrieve,
	HideHelpCommand: true,
}

var callsList = cli.Command{
	Name:    "list",
	Usage:   "List voice calls for this project, most recent first. Transcripts are omitted\nfrom the list; fetch a single call to get its transcript.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "direction",
			Usage:     "Whether the call was placed by Zavu (outbound) or received from a caller (inbound).",
			QueryPath: "direction",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Lifecycle status of a voice call.\n- `queued`: outbound call created, not yet dialing.\n- `ringing`: dialing (outbound) or received and ringing (inbound).\n- `in_progress`: answered, the agent is connected.\n- `completed`: ended after a conversation.\n- `failed`: could not be completed.\n- `busy`: the line was busy.\n- `no_answer`: rang but was not answered.\n- `canceled`: canceled before it was answered.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleCallsList,
	HideHelpCommand: true,
}

var callsHangup = cli.Command{
	Name:    "hangup",
	Usage:   "End an active voice call. The call must still be ringing or in progress. Not\navailable with test-mode API keys.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "call-id",
			Required:  true,
			PathParam: "callId",
		},
	},
	Action:          handleCallsHangup,
	HideHelpCommand: true,
}

func handleCallsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.CallNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Calls.New(ctx, params, options...)
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
		Title:          "calls create",
		Transform:      transform,
	})
}

func handleCallsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-id") && len(unusedArgs) > 0 {
		cmd.Set("call-id", unusedArgs[0])
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
	_, err = client.Calls.Get(ctx, cmd.Value("call-id").(string), options...)
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
		Title:          "calls retrieve",
		Transform:      transform,
	})
}

func handleCallsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := zavudev.CallListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Calls.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "calls list",
			Transform:      transform,
		})
	} else {
		iter := client.Calls.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "calls list",
			Transform:      transform,
		})
	}
}

func handleCallsHangup(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-id") && len(unusedArgs) > 0 {
		cmd.Set("call-id", unusedArgs[0])
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
	_, err = client.Calls.Hangup(ctx, cmd.Value("call-id").(string), options...)
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
		Title:          "calls hangup",
		Transform:      transform,
	})
}
