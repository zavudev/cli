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

var sendersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create sender",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "email-address",
			Usage:    "From-address for the email channel (e.g. noreply@yourdomain.com). The address's domain must be a verified email domain in your project. Setting this attaches the email channel to the sender.",
			BodyPath: "emailAddress",
		},
		&requestflag.Flag[string]{
			Name:     "email-domain-id",
			Usage:    "ID of the verified email domain to attach. Optional — resolved from `emailAddress`'s domain when omitted.",
			BodyPath: "emailDomainId",
		},
		&requestflag.Flag[string]{
			Name:     "email-from-name",
			Usage:    "Display name shown in the recipient's inbox for the email channel.",
			BodyPath: "emailFromName",
		},
		&requestflag.Flag[bool]{
			Name:     "email-receiving-enabled",
			Usage:    "Enable inbound email receiving on this sender. Requires a verified MX record on the domain; ignored otherwise.",
			BodyPath: "emailReceivingEnabled",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-sms-oneway",
			Usage:    "Enable the one-way SMS channel (`sms_oneway`). Needs nothing else — no phone number, no credential — so it is the fastest way to get a sender that can send. Recipients cannot reply. Confirm with `sms_oneway` in the `channels` array on the response.",
			Default:  false,
			BodyPath: "enableSmsOneway",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-voice",
			Usage:    "Let this sender place and answer phone calls. Requires `phoneNumber`; enabling it without one returns 400. Check the `channels` array on the response to confirm `voice` is on.",
			Default:  false,
			BodyPath: "enableVoice",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "Phone number in E.164 format, and it must be a number your project already owns (see `GET /v1/phone-numbers`). The number is routed to the sender as part of this call, which is what turns the SMS channel on. Passing a number the project does not own, or one already attached to another sender, returns 400 rather than creating a sender that cannot send. Omit for an email-only sender.",
			BodyPath: "phoneNumber",
		},
		&requestflag.Flag[bool]{
			Name:     "set-as-default",
			Default:  false,
			BodyPath: "setAsDefault",
		},
		&requestflag.Flag[[]string]{
			Name:     "webhook-event",
			Usage:    "Events to subscribe to.",
			BodyPath: "webhookEvents",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-signature-version",
			Usage:    "Which `X-Zavu-Signature` scheme this receiver is sent.\n\n- `v1`: `v1=HMAC_SHA256(secret, body)`. The scheme used before this was configurable. Existing webhooks stay on it until you move them.\n- `v2`: `v2=HMAC_SHA256(secret, \"{t}.{body}\")`. The current scheme, and the default for new senders. It signs the timestamp together with the body.\n- `v1+v2`: both signatures, sharing one `t`. The migration setting: a receiver reading either one works, so you can deploy and confirm your new verifier before switching over.\n\nMoving from `v1` straight to `v2` returns `400`. Set `v1+v2` first. See https://docs.zavu.dev/guides/receiving-messages/signature-migration",
			BodyPath: "webhookSignatureVersion",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "HTTPS URL for webhook events.",
			BodyPath: "webhookUrl",
		},
	},
	Action:          handleSendersCreate,
	HideHelpCommand: true,
}

var sendersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get sender",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersRetrieve,
	HideHelpCommand: true,
}

var sendersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update sender",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
		&requestflag.Flag[string]{
			Name:     "email-address",
			Usage:    "Attach or change the sender's email from-address (e.g. noreply@yourdomain.com). The domain must be a verified email domain in your project.",
			BodyPath: "emailAddress",
		},
		&requestflag.Flag[bool]{
			Name:     "email-catch-all-enabled",
			Usage:    "Enable or disable domain catch-all. When enabled (with emailReceivingEnabled true), this sender receives email for any address at its domain. Ignored (treated as false) if receiving is not enabled.",
			BodyPath: "emailCatchAllEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "email-domain-id",
			Usage:    "ID of the verified email domain to attach. Optional — resolved from `emailAddress`'s domain when omitted.",
			BodyPath: "emailDomainId",
		},
		&requestflag.Flag[string]{
			Name:     "email-from-name",
			Usage:    "Display name shown in the recipient's inbox for the email channel.",
			BodyPath: "emailFromName",
		},
		&requestflag.Flag[bool]{
			Name:     "email-receiving-enabled",
			Usage:    "Enable or disable inbound email receiving for this sender.",
			BodyPath: "emailReceivingEnabled",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-sms-oneway",
			Usage:    "Turn the one-way SMS channel on or off. Enabling needs nothing else and takes effect immediately; disabling removes the channel from the sender. Confirm with the `channels` array on the response.",
			BodyPath: "enableSmsOneway",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-voice",
			Usage:    "Turn the voice channel on or off. The sender must already have a phone number provisioned for calls; enabling it otherwise returns 400 instead of storing a flag that changes nothing. Confirm with the `channels` array on the response.",
			BodyPath: "enableVoice",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "set-as-default",
			BodyPath: "setAsDefault",
		},
		&requestflag.Flag[bool]{
			Name:     "webhook-active",
			Usage:    "Whether the webhook is active.",
			BodyPath: "webhookActive",
		},
		&requestflag.Flag[[]string]{
			Name:     "webhook-event",
			Usage:    "Events to subscribe to.",
			BodyPath: "webhookEvents",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-signature-version",
			Usage:    "Which `X-Zavu-Signature` scheme this receiver is sent.\n\n- `v1`: `v1=HMAC_SHA256(secret, body)`. The scheme used before this was configurable. Existing webhooks stay on it until you move them.\n- `v2`: `v2=HMAC_SHA256(secret, \"{t}.{body}\")`. The current scheme, and the default for new senders. It signs the timestamp together with the body.\n- `v1+v2`: both signatures, sharing one `t`. The migration setting: a receiver reading either one works, so you can deploy and confirm your new verifier before switching over.\n\nMoving from `v1` straight to `v2` returns `400`. Set `v1+v2` first. See https://docs.zavu.dev/guides/receiving-messages/signature-migration",
			BodyPath: "webhookSignatureVersion",
		},
		&requestflag.Flag[*string]{
			Name:     "webhook-url",
			Usage:    "HTTPS URL for webhook events. Set to null to remove webhook.",
			BodyPath: "webhookUrl",
		},
	},
	Action:          handleSendersUpdate,
	HideHelpCommand: true,
}

var sendersList = cli.Command{
	Name:    "list",
	Usage:   "List senders",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSendersList,
	HideHelpCommand: true,
}

var sendersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete sender",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersDelete,
	HideHelpCommand: true,
}

var sendersGetProfile = cli.Command{
	Name:    "get-profile",
	Usage:   "Get the WhatsApp Business profile for a sender. The sender must have a WhatsApp\nBusiness Account connected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersGetProfile,
	HideHelpCommand: true,
}

var sendersRegenerateWebhookSecret = cli.Command{
	Name:    "regenerate-webhook-secret",
	Usage:   "Regenerate the webhook secret for a sender. The old secret will be invalidated\nimmediately.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersRegenerateWebhookSecret,
	HideHelpCommand: true,
}

var sendersUpdateProfile = cli.Command{
	Name:    "update-profile",
	Usage:   "Update the WhatsApp Business profile for a sender. The sender must have a\nWhatsApp Business Account connected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
		&requestflag.Flag[string]{
			Name:     "about",
			Usage:    "Short description of the business (max 139 characters).",
			BodyPath: "about",
		},
		&requestflag.Flag[string]{
			Name:     "address",
			Usage:    "Physical address of the business (max 256 characters).",
			BodyPath: "address",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Extended description of the business (max 512 characters).",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Business email address.",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "vertical",
			Usage:    "Business category for WhatsApp Business profile.",
			BodyPath: "vertical",
		},
		&requestflag.Flag[[]string]{
			Name:     "website",
			Usage:    "Business website URLs (maximum 2).",
			BodyPath: "websites",
		},
	},
	Action:          handleSendersUpdateProfile,
	HideHelpCommand: true,
}

var sendersUploadProfilePicture = cli.Command{
	Name:    "upload-profile-picture",
	Usage:   "Upload a new profile picture for the WhatsApp Business profile. The image will\nbe uploaded to Meta and set as the profile picture.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
		&requestflag.Flag[string]{
			Name:     "image-url",
			Usage:    "URL of the image to upload.",
			Required: true,
			BodyPath: "imageUrl",
		},
		&requestflag.Flag[string]{
			Name:     "mime-type",
			Usage:    "MIME type of the image.",
			Required: true,
			BodyPath: "mimeType",
		},
	},
	Action:          handleSendersUploadProfilePicture,
	HideHelpCommand: true,
}

func handleSendersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.SenderNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.New(ctx, params, options...)
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
		Title:          "senders create",
		Transform:      transform,
	})
}

func handleSendersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Senders.Get(ctx, cmd.Value("sender-id").(string), options...)
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
		Title:          "senders retrieve",
		Transform:      transform,
	})
}

func handleSendersUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.SenderUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Update(
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
		Title:          "senders update",
		Transform:      transform,
	})
}

func handleSendersList(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.SenderListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Senders.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "senders list",
			Transform:      transform,
		})
	} else {
		iter := client.Senders.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "senders list",
			Transform:      transform,
		})
	}
}

func handleSendersDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Senders.Delete(ctx, cmd.Value("sender-id").(string), options...)
}

func handleSendersGetProfile(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Senders.GetProfile(ctx, cmd.Value("sender-id").(string), options...)
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
		Title:          "senders get-profile",
		Transform:      transform,
	})
}

func handleSendersRegenerateWebhookSecret(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Senders.RegenerateWebhookSecret(ctx, cmd.Value("sender-id").(string), options...)
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
		Title:          "senders regenerate-webhook-secret",
		Transform:      transform,
	})
}

func handleSendersUpdateProfile(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.SenderUpdateProfileParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.UpdateProfile(
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
		Title:          "senders update-profile",
		Transform:      transform,
	})
}

func handleSendersUploadProfilePicture(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.SenderUploadProfilePictureParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.UploadProfilePicture(
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
		Title:          "senders upload-profile-picture",
		Transform:      transform,
	})
}
