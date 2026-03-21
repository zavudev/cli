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

var messagesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get message by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
	},
	Action:          handleMessagesRetrieve,
	HideHelpCommand: true,
}

var messagesList = cli.Command{
	Name:    "list",
	Usage:   "List messages previously sent by this project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "channel",
			Usage:     "Delivery channel. Use 'auto' for intelligent routing.",
			QueryPath: "channel",
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
			Usage:     `Allowed values: "queued", "sending", "sent", "delivered", "failed", "received", "pending_url_verification".`,
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			QueryPath: "to",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleMessagesList,
	HideHelpCommand: true,
}

var messagesReact = cli.Command{
	Name:    "react",
	Usage:   "Send an emoji reaction to an existing WhatsApp message. Reactions are only\nsupported for WhatsApp messages.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "emoji",
			Usage:    "Single emoji character to react with.",
			Required: true,
			BodyPath: "emoji",
		},
		&requestflag.Flag[string]{
			Name:       "zavu-sender",
			HeaderPath: "Zavu-Sender",
		},
	},
	Action:          handleMessagesReact,
	HideHelpCommand: true,
}

var messagesSend = requestflag.WithInnerFlags(cli.Command{
	Name:    "send",
	Usage:   "Send a message to a recipient via SMS or WhatsApp.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Recipient phone number in E.164 format, email address, or numeric chat ID (for Telegram/Instagram).",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[string]{
			Name:     "channel",
			Usage:    "Delivery channel. Use 'auto' for intelligent routing.",
			BodyPath: "channel",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "content",
			Usage:    "Content for non-text message types (WhatsApp and Telegram).",
			BodyPath: "content",
		},
		&requestflag.Flag[bool]{
			Name:     "fallback-enabled",
			Usage:    "Whether to enable automatic fallback to SMS if WhatsApp fails. Defaults to true.",
			Default:  true,
			BodyPath: "fallbackEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "html-body",
			Usage:    "HTML body for email messages. If provided, email will be sent as multipart with both text and HTML.",
			BodyPath: "htmlBody",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Optional idempotency key to avoid duplicate sends.",
			BodyPath: "idempotencyKey",
		},
		&requestflag.Flag[string]{
			Name:     "message-type",
			Usage:    "Type of message. Non-text types are supported by WhatsApp and Telegram (varies by type).",
			BodyPath: "messageType",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Arbitrary metadata to associate with the message.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "reply-to",
			Usage:    "Reply-To email address for email messages.",
			BodyPath: "replyTo",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Email subject line. Required when channel is 'email' or recipient is an email address.",
			BodyPath: "subject",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Text body for text messages or caption for media messages.",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "voice-language",
			Usage:    "Language code for voice text-to-speech (e.g., 'en-US', 'es-ES', 'pt-BR'). If omitted, language is auto-detected from recipient's country code.",
			BodyPath: "voiceLanguage",
		},
		&requestflag.Flag[string]{
			Name:       "zavu-sender",
			HeaderPath: "Zavu-Sender",
		},
	},
	Action:          handleMessagesSend,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"content": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "content.buttons",
			Usage:      "Interactive buttons (max 3).",
			InnerField: "buttons",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "content.contacts",
			Usage:      "Contact cards for contact messages.",
			InnerField: "contacts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.emoji",
			Usage:      "Emoji for reaction messages.",
			InnerField: "emoji",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.filename",
			Usage:      "Filename for documents.",
			InnerField: "filename",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "content.latitude",
			Usage:      "Latitude for location messages.",
			InnerField: "latitude",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.list-button",
			Usage:      "Button text for list messages.",
			InnerField: "listButton",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.location-address",
			Usage:      "Address of the location.",
			InnerField: "locationAddress",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.location-name",
			Usage:      "Name of the location.",
			InnerField: "locationName",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "content.longitude",
			Usage:      "Longitude for location messages.",
			InnerField: "longitude",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.media-id",
			Usage:      "WhatsApp media ID if already uploaded.",
			InnerField: "mediaId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.media-url",
			Usage:      "URL of the media file (for image, video, audio, document, sticker).",
			InnerField: "mediaUrl",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.mime-type",
			Usage:      "MIME type of the media.",
			InnerField: "mimeType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.react-to-message-id",
			Usage:      "Message ID to react to.",
			InnerField: "reactToMessageId",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "content.sections",
			Usage:      "Sections for list messages.",
			InnerField: "sections",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.template-id",
			Usage:      "Template ID for template messages.",
			InnerField: "templateId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "content.template-variables",
			Usage:      "Variables for template rendering. Keys are variable positions (1, 2, 3...).",
			InnerField: "templateVariables",
		},
	},
})

func handleMessagesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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
	_, err = client.Messages.Get(ctx, cmd.Value("message-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages retrieve", obj, format, transform)
}

func handleMessagesList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.MessageListParams{}

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
		_, err = client.Messages.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messages list", obj, format, transform)
	} else {
		iter := client.Messages.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "messages list", iter, format, transform, maxItems)
	}
}

func handleMessagesReact(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.MessageReactParams{}

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
	_, err = client.Messages.React(
		ctx,
		cmd.Value("message-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages react", obj, format, transform)
}

func handleMessagesSend(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.MessageSendParams{}

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
	_, err = client.Messages.Send(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages send", obj, format, transform)
}
