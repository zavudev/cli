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

var templatesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a WhatsApp message template. Note: Templates must be approved by Meta\nbefore use.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "body",
			Usage:    "Default template body. Used when no channel-specific body is set.",
			Required: true,
			BodyPath: "body",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Default:  "en",
			Required: true,
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "add-security-recommendation",
			Usage:    "Add 'Do not share this code' disclaimer. Only for AUTHENTICATION templates.",
			BodyPath: "addSecurityRecommendation",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "button",
			Usage:    "Template buttons (max 3).",
			BodyPath: "buttons",
		},
		&requestflag.Flag[int64]{
			Name:     "code-expiration-minutes",
			Usage:    "Code expiration time in minutes. Only for AUTHENTICATION templates.",
			BodyPath: "codeExpirationMinutes",
		},
		&requestflag.Flag[string]{
			Name:     "footer",
			Usage:    "Footer text for the template.",
			BodyPath: "footer",
		},
		&requestflag.Flag[string]{
			Name:     "header-content",
			Usage:    "Header content (text string or media URL).",
			BodyPath: "headerContent",
		},
		&requestflag.Flag[string]{
			Name:     "header-type",
			Usage:    "Type of header for the template.",
			BodyPath: "headerType",
		},
		&requestflag.Flag[string]{
			Name:     "instagram-body",
			Usage:    "Channel-specific body for Instagram. Falls back to `body` if not set.",
			BodyPath: "instagramBody",
		},
		&requestflag.Flag[string]{
			Name:     "sms-body",
			Usage:    "Channel-specific body for SMS. Falls back to `body` if not set.",
			BodyPath: "smsBody",
		},
		&requestflag.Flag[string]{
			Name:     "telegram-body",
			Usage:    "Channel-specific body for Telegram. Falls back to `body` if not set.",
			BodyPath: "telegramBody",
		},
		&requestflag.Flag[[]string]{
			Name:     "variable",
			BodyPath: "variables",
		},
		&requestflag.Flag[string]{
			Name:     "whatsapp-category",
			Usage:    "WhatsApp template category.",
			BodyPath: "whatsappCategory",
		},
	},
	Action:          handleTemplatesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"button": {
		&requestflag.InnerFlag[string]{
			Name:       "button.text",
			InnerField: "text",
		},
		&requestflag.InnerFlag[string]{
			Name:       "button.type",
			Usage:      `Allowed values: "quick_reply", "url", "phone", "otp".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "button.otp-type",
			Usage:      "Required when type is 'otp'. COPY_CODE shows copy button, ONE_TAP enables Android autofill.",
			InnerField: "otpType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "button.package-name",
			Usage:      "Android package name. Required for ONE_TAP buttons.",
			InnerField: "packageName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "button.phone-number",
			InnerField: "phoneNumber",
		},
		&requestflag.InnerFlag[string]{
			Name:       "button.signature-hash",
			Usage:      "Android app signature hash. Required for ONE_TAP buttons.",
			InnerField: "signatureHash",
		},
		&requestflag.InnerFlag[string]{
			Name:       "button.url",
			InnerField: "url",
		},
	},
})

var templatesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get template",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "template-id",
			Required: true,
		},
	},
	Action:          handleTemplatesRetrieve,
	HideHelpCommand: true,
}

var templatesList = cli.Command{
	Name:    "list",
	Usage:   "List WhatsApp message templates for this project.",
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
	Action:          handleTemplatesList,
	HideHelpCommand: true,
}

var templatesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete template",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "template-id",
			Required: true,
		},
	},
	Action:          handleTemplatesDelete,
	HideHelpCommand: true,
}

var templatesSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Submit a WhatsApp template to Meta for approval. The template must be in draft\nstatus and associated with a sender that has a WhatsApp Business Account\nconfigured.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "template-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Usage:    "The sender ID with the WhatsApp Business Account to submit the template to.",
			Required: true,
			BodyPath: "senderId",
		},
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "WhatsApp template category.",
			BodyPath: "category",
		},
	},
	Action:          handleTemplatesSubmit,
	HideHelpCommand: true,
}

func handleTemplatesCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.TemplateNewParams{}

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
	_, err = client.Templates.New(ctx, params, options...)
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
		Title:          "templates create",
		Transform:      transform,
	})
}

func handleTemplatesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
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
	_, err = client.Templates.Get(ctx, cmd.Value("template-id").(string), options...)
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
		Title:          "templates retrieve",
		Transform:      transform,
	})
}

func handleTemplatesList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.TemplateListParams{}

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
		_, err = client.Templates.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "templates list",
			Transform:      transform,
		})
	} else {
		iter := client.Templates.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "templates list",
			Transform:      transform,
		})
	}
}

func handleTemplatesDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
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

	return client.Templates.Delete(ctx, cmd.Value("template-id").(string), options...)
}

func handleTemplatesSubmit(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.TemplateSubmitParams{}

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
	_, err = client.Templates.Submit(
		ctx,
		cmd.Value("template-id").(string),
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
		Title:          "templates submit",
		Transform:      transform,
	})
}
