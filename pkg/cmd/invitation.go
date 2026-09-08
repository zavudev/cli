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

var invitationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a partner invitation link for a client to connect a Meta channel. The\nclient opens the returned `url` and authorizes with Meta; the resulting sender\nis created in your project when they finish, and the invitation transitions to\n`completed`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "allowed-phone-country",
			Usage:    "ISO country codes for allowed phone numbers. Only valid when `connectionType` is `whatsapp_waba` — sending it with `messenger` returns 400.",
			BodyPath: "allowedPhoneCountries",
		},
		&requestflag.Flag[string]{
			Name:     "client-email",
			Usage:    "Email of the client being invited.",
			BodyPath: "clientEmail",
		},
		&requestflag.Flag[string]{
			Name:     "client-name",
			Usage:    "Name of the client being invited.",
			BodyPath: "clientName",
		},
		&requestflag.Flag[string]{
			Name:     "client-phone",
			Usage:    "Phone number of the client in E.164 format.",
			BodyPath: "clientPhone",
		},
		&requestflag.Flag[string]{
			Name:     "connection-type",
			Usage:    "Which Meta channel the client connects, and how.\n- `whatsapp_waba` (default): Meta's embedded signup links an official WhatsApp Business Account. Accepts `phoneNumberId` and `allowedPhoneCountries`.\n- `messenger`: the client authorizes with Facebook and picks a Facebook Page they administer. The Page's Messenger inbox — including Marketplace chats — is routed to Zavu. They must be an admin of at least one Page. A Page can only be connected to one Zavu project at a time: if the client picks a Page that another project already connected, the newer connection wins and the older one is disconnected.\n\nOne invitation connects one channel. To onboard a client on several channels, create one invitation per channel; each completes into its own sender.",
			Default:  "whatsapp_waba",
			BodyPath: "connectionType",
		},
		&requestflag.Flag[int64]{
			Name:     "expires-in-days",
			Usage:    "Number of days until the invitation expires.",
			Default:  7,
			BodyPath: "expiresInDays",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number-id",
			Usage:    "ID of a Zavu phone number to pre-assign for WhatsApp registration. If provided, the client will use this number instead of their own. Only valid when `connectionType` is `whatsapp_waba` — sending it with `messenger` returns 400, since a Facebook Page has no phone number.",
			BodyPath: "phoneNumberId",
		},
	},
	Action:          handleInvitationsCreate,
	HideHelpCommand: true,
}

var invitationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get invitation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "invitation-id",
			Required:  true,
			PathParam: "invitationId",
		},
	},
	Action:          handleInvitationsRetrieve,
	HideHelpCommand: true,
}

var invitationsList = cli.Command{
	Name:    "list",
	Usage:   "List partner invitations for this project.",
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
			Name:      "status",
			Usage:     "Current status of the partner invitation.\n\n`failed` means the client started the connection and it did not finish (they cancelled Meta's dialog, denied a permission, or abandoned the tab). A failed invitation is still usable: the same link can be retried, and it moves back to `in_progress` when the client tries again.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleInvitationsList,
	HideHelpCommand: true,
}

var invitationsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel an active invitation. The client will no longer be able to use the\ninvitation link.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "invitation-id",
			Required:  true,
			PathParam: "invitationId",
		},
	},
	Action:          handleInvitationsCancel,
	HideHelpCommand: true,
}

func handleInvitationsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.InvitationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Invitations.New(ctx, params, options...)
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
		Title:          "invitations create",
		Transform:      transform,
	})
}

func handleInvitationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invitation-id") && len(unusedArgs) > 0 {
		cmd.Set("invitation-id", unusedArgs[0])
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
	_, err = client.Invitations.Get(ctx, cmd.Value("invitation-id").(string), options...)
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
		Title:          "invitations retrieve",
		Transform:      transform,
	})
}

func handleInvitationsList(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.InvitationListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Invitations.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "invitations list",
			Transform:      transform,
		})
	} else {
		iter := client.Invitations.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "invitations list",
			Transform:      transform,
		})
	}
}

func handleInvitationsCancel(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invitation-id") && len(unusedArgs) > 0 {
		cmd.Set("invitation-id", unusedArgs[0])
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
	_, err = client.Invitations.Cancel(ctx, cmd.Value("invitation-id").(string), options...)
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
		Title:          "invitations cancel",
		Transform:      transform,
	})
}
