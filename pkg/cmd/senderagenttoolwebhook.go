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

var sendersAgentToolsWebhookRotateSecret = cli.Command{
	Name:    "rotate-secret",
	Usage:   "Generate a new signing secret for this tool. The previous one stops working on\nthe next call, with no overlap, so update your endpoint first. The tool keeps\nits id, so flows that reference it by name are unaffected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
		&requestflag.Flag[string]{
			Name:      "tool-id",
			Required:  true,
			PathParam: "toolId",
		},
	},
	Action:          handleSendersAgentToolsWebhookRotateSecret,
	HideHelpCommand: true,
}

func handleSendersAgentToolsWebhookRotateSecret(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
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

	params := zavudev.SenderAgentToolWebhookRotateSecretParams{
		SenderID: cmd.Value("sender-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Agent.Tools.Webhook.RotateSecret(
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "senders:agent:tools:webhook rotate-secret",
		Transform:      transform,
	})
}
