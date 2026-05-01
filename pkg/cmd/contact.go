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

var contactsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new contact with one or more communication channels.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "channel",
			Usage:    "Communication channels for the contact.",
			Required: true,
			BodyPath: "channels",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name for the contact.",
			BodyPath: "displayName",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Arbitrary metadata to associate with the contact.",
			BodyPath: "metadata",
		},
	},
	Action:          handleContactsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"channel": {
		&requestflag.InnerFlag[string]{
			Name:       "channel.channel",
			Usage:      "Channel type.",
			InnerField: "channel",
		},
		&requestflag.InnerFlag[string]{
			Name:       "channel.identifier",
			Usage:      "Channel identifier (phone number in E.164 format or email address).",
			InnerField: "identifier",
		},
		&requestflag.InnerFlag[string]{
			Name:       "channel.country-code",
			Usage:      "ISO country code for phone numbers.",
			InnerField: "countryCode",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "channel.is-primary",
			Usage:      "Whether this should be the primary channel for its type.",
			InnerField: "isPrimary",
		},
		&requestflag.InnerFlag[string]{
			Name:       "channel.label",
			Usage:      "Optional label for the channel.",
			InnerField: "label",
		},
	},
})

var contactsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get contact",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
	},
	Action:          handleContactsRetrieve,
	HideHelpCommand: true,
}

var contactsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update contact",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
		&requestflag.Flag[*string]{
			Name:     "default-channel",
			Usage:    "Preferred channel for this contact. Set to null to clear.",
			BodyPath: "defaultChannel",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
	},
	Action:          handleContactsUpdate,
	HideHelpCommand: true,
}

var contactsList = cli.Command{
	Name:    "list",
	Usage:   "List contacts with their communication channels.",
	Suggest: true,
	Flags: []cli.Flag{
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
			Name:      "phone-number",
			QueryPath: "phoneNumber",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleContactsList,
	HideHelpCommand: true,
}

var contactsDismissMergeSuggestion = cli.Command{
	Name:    "dismiss-merge-suggestion",
	Usage:   "Dismiss the merge suggestion for a contact.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
	},
	Action:          handleContactsDismissMergeSuggestion,
	HideHelpCommand: true,
}

var contactsMerge = cli.Command{
	Name:    "merge",
	Usage:   "Merge a source contact into this contact. All channels from the source contact\nwill be moved to the target contact, and the source contact will be marked as\nmerged.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "source-contact-id",
			Usage:    "ID of the contact to merge into the target contact. The source contact will be marked as merged.",
			Required: true,
			BodyPath: "sourceContactId",
		},
	},
	Action:          handleContactsMerge,
	HideHelpCommand: true,
}

var contactsRetrieveByPhone = cli.Command{
	Name:    "retrieve-by-phone",
	Usage:   "Get contact by phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
		},
	},
	Action:          handleContactsRetrieveByPhone,
	HideHelpCommand: true,
}

func handleContactsCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.ContactNewParams{}

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
	_, err = client.Contacts.New(ctx, params, options...)
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
		Title:          "contacts create",
		Transform:      transform,
	})
}

func handleContactsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Contacts.Get(ctx, cmd.Value("contact-id").(string), options...)
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
		Title:          "contacts retrieve",
		Transform:      transform,
	})
}

func handleContactsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("contact-id") && len(unusedArgs) > 0 {
		cmd.Set("contact-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.ContactUpdateParams{}

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
	_, err = client.Contacts.Update(
		ctx,
		cmd.Value("contact-id").(string),
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
		Title:          "contacts update",
		Transform:      transform,
	})
}

func handleContactsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.ContactListParams{}

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
		_, err = client.Contacts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "contacts list",
			Transform:      transform,
		})
	} else {
		iter := client.Contacts.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "contacts list",
			Transform:      transform,
		})
	}
}

func handleContactsDismissMergeSuggestion(ctx context.Context, cmd *cli.Command) error {
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

	return client.Contacts.DismissMergeSuggestion(ctx, cmd.Value("contact-id").(string), options...)
}

func handleContactsMerge(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("contact-id") && len(unusedArgs) > 0 {
		cmd.Set("contact-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.ContactMergeParams{}

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
	_, err = client.Contacts.Merge(
		ctx,
		cmd.Value("contact-id").(string),
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
		Title:          "contacts merge",
		Transform:      transform,
	})
}

func handleContactsRetrieveByPhone(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
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
	_, err = client.Contacts.GetByPhone(ctx, cmd.Value("phone-number").(string), options...)
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
		Title:          "contacts retrieve-by-phone",
		Transform:      transform,
	})
}
