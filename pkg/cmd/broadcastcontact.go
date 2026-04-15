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

var broadcastsContactsList = cli.Command{
	Name:    "list",
	Usage:   "List contacts in a broadcast with optional status filter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
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
			Usage:     "Status of a contact within a broadcast.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBroadcastsContactsList,
	HideHelpCommand: true,
}

var broadcastsContactsAdd = requestflag.WithInnerFlags(cli.Command{
	Name:    "add",
	Usage:   "Add contacts to a broadcast in batch. Maximum 1000 contacts per request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "contact",
			Usage:    "List of contacts to add (max 1000 per request).",
			Required: true,
			BodyPath: "contacts",
		},
	},
	Action:          handleBroadcastsContactsAdd,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"contact": {
		&requestflag.InnerFlag[string]{
			Name:       "contact.recipient",
			Usage:      "Phone number (E.164) or email address.",
			InnerField: "recipient",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "contact.template-variables",
			Usage:      "Per-contact template variables to personalize the message.",
			InnerField: "templateVariables",
		},
	},
})

var broadcastsContactsRemove = cli.Command{
	Name:    "remove",
	Usage:   "Remove a contact from a broadcast in draft status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
	},
	Action:          handleBroadcastsContactsRemove,
	HideHelpCommand: true,
}

func handleBroadcastsContactsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.BroadcastContactListParams{}

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
		_, err = client.Broadcasts.Contacts.List(
			ctx,
			cmd.Value("broadcast-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "broadcasts:contacts list", obj, format, explicitFormat, transform)
	} else {
		iter := client.Broadcasts.Contacts.ListAutoPaging(
			ctx,
			cmd.Value("broadcast-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "broadcasts:contacts list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleBroadcastsContactsAdd(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.BroadcastContactAddParams{}

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
	_, err = client.Broadcasts.Contacts.Add(
		ctx,
		cmd.Value("broadcast-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "broadcasts:contacts add", obj, format, explicitFormat, transform)
}

func handleBroadcastsContactsRemove(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("contact-id") && len(unusedArgs) > 0 {
		cmd.Set("contact-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.BroadcastContactRemoveParams{
		BroadcastID: cmd.Value("broadcast-id").(string),
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

	return client.Broadcasts.Contacts.Remove(
		ctx,
		cmd.Value("contact-id").(string),
		params,
		options...,
	)
}
