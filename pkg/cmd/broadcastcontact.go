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

var broadcastsContactsList = cli.Command{
	Name:    "list",
	Usage:   "List contacts in a broadcast with optional status filter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
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
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
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
			Name:       "contact.template-button-variables",
			Usage:      "Per-contact button variables for dynamic URL/OTP buttons. Keys are the button index (0, 1, 2).",
			InnerField: "templateButtonVariables",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "contact.template-variables",
			Usage:      "Per-contact body variables. Keys are positions (1, 2, ...) matching the order placeholders appear in the template body.",
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
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
		&requestflag.Flag[string]{
			Name:      "contact-id",
			Required:  true,
			PathParam: "contactId",
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

	params := zavudev.BroadcastContactListParams{}

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
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "broadcasts:contacts list",
			Transform:      transform,
		})
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
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "broadcasts:contacts list",
			Transform:      transform,
		})
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

	params := zavudev.BroadcastContactAddParams{}

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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "broadcasts:contacts add",
		Transform:      transform,
	})
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

	params := zavudev.BroadcastContactRemoveParams{
		BroadcastID: cmd.Value("broadcast-id").(string),
	}

	return client.Broadcasts.Contacts.Remove(
		ctx,
		cmd.Value("contact-id").(string),
		params,
		options...,
	)
}
