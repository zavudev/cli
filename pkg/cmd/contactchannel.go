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

var contactsChannelsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a contact's channel properties.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "channel-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "label",
			Usage:    "Optional label for the channel. Set to null to clear.",
			BodyPath: "label",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[bool]{
			Name:     "verified",
			Usage:    "Whether the channel is verified.",
			BodyPath: "verified",
		},
	},
	Action:          handleContactsChannelsUpdate,
	HideHelpCommand: true,
}

var contactsChannelsAdd = cli.Command{
	Name:    "add",
	Usage:   "Add a new communication channel to an existing contact.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "channel",
			Usage:    "Channel type.",
			Required: true,
			BodyPath: "channel",
		},
		&requestflag.Flag[string]{
			Name:     "identifier",
			Usage:    "Channel identifier (phone number in E.164 format or email address).",
			Required: true,
			BodyPath: "identifier",
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Usage:    "ISO country code for phone numbers.",
			BodyPath: "countryCode",
		},
		&requestflag.Flag[bool]{
			Name:     "is-primary",
			Usage:    "Whether this should be the primary channel for its type.",
			Default:  false,
			BodyPath: "isPrimary",
		},
		&requestflag.Flag[string]{
			Name:     "label",
			Usage:    "Optional label for the channel.",
			BodyPath: "label",
		},
	},
	Action:          handleContactsChannelsAdd,
	HideHelpCommand: true,
}

var contactsChannelsRemove = cli.Command{
	Name:    "remove",
	Usage:   "Remove a communication channel from a contact. Cannot remove the last channel.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "channel-id",
			Required: true,
		},
	},
	Action:          handleContactsChannelsRemove,
	HideHelpCommand: true,
}

var contactsChannelsSetPrimary = cli.Command{
	Name:    "set-primary",
	Usage:   "Set a channel as the primary channel for its type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "channel-id",
			Required: true,
		},
	},
	Action:          handleContactsChannelsSetPrimary,
	HideHelpCommand: true,
}

func handleContactsChannelsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("channel-id") && len(unusedArgs) > 0 {
		cmd.Set("channel-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.ContactChannelUpdateParams{
		ContactID: cmd.Value("contact-id").(string),
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
	_, err = client.Contacts.Channels.Update(
		ctx,
		cmd.Value("channel-id").(string),
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
		Title:          "contacts:channels update",
		Transform:      transform,
	})
}

func handleContactsChannelsAdd(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("contact-id") && len(unusedArgs) > 0 {
		cmd.Set("contact-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.ContactChannelAddParams{}

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
	_, err = client.Contacts.Channels.Add(
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
		Title:          "contacts:channels add",
		Transform:      transform,
	})
}

func handleContactsChannelsRemove(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("channel-id") && len(unusedArgs) > 0 {
		cmd.Set("channel-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.ContactChannelRemoveParams{
		ContactID: cmd.Value("contact-id").(string),
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

	return client.Contacts.Channels.Remove(
		ctx,
		cmd.Value("channel-id").(string),
		params,
		options...,
	)
}

func handleContactsChannelsSetPrimary(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("channel-id") && len(unusedArgs) > 0 {
		cmd.Set("channel-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.ContactChannelSetPrimaryParams{
		ContactID: cmd.Value("contact-id").(string),
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
	_, err = client.Contacts.Channels.SetPrimary(
		ctx,
		cmd.Value("channel-id").(string),
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
		Title:          "contacts:channels set-primary",
		Transform:      transform,
	})
}
