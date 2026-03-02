// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/zavudev-cli/internal/apiquery"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
	"github.com/stainless-sdks/zavudev-go"
	"github.com/stainless-sdks/zavudev-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var sendersAgentExecutionsList = cli.Command{
	Name:    "list",
	Usage:   "List recent agent executions with pagination.",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Status of an agent execution.",
			QueryPath: "status",
		},
	},
	Action:          handleSendersAgentExecutionsList,
	HideHelpCommand: true,
}

func handleSendersAgentExecutionsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentExecutionListParams{}

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
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Senders.Agent.Executions.List(
			ctx,
			cmd.Value("sender-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "senders:agent:executions list", obj, format, transform)
	} else {
		iter := client.Senders.Agent.Executions.ListAutoPaging(
			ctx,
			cmd.Value("sender-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "senders:agent:executions list", iter, format, transform)
	}
}
