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

var broadcastsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new broadcast campaign. Add contacts after creation, then send.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "channel",
			Usage:    "Broadcast delivery channel. Use 'smart' for per-contact intelligent routing.",
			Required: true,
			BodyPath: "channel",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the broadcast campaign.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "content",
			Usage:    "Content for non-text broadcast message types.",
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "email-html-body",
			Usage:    "HTML body for email broadcasts.",
			BodyPath: "emailHtmlBody",
		},
		&requestflag.Flag[string]{
			Name:     "email-subject",
			Usage:    "Email subject line. Required for email broadcasts.",
			BodyPath: "emailSubject",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Idempotency key to prevent duplicate broadcasts.",
			BodyPath: "idempotencyKey",
		},
		&requestflag.Flag[string]{
			Name:     "message-type",
			Usage:    "Type of message for broadcast.",
			BodyPath: "messageType",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "scheduled-at",
			Usage:    "Schedule the broadcast for future delivery.",
			BodyPath: "scheduledAt",
		},
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Usage:    "Sender profile ID. Uses default sender if omitted.",
			BodyPath: "senderId",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Text content or caption. Supports template variables: {{name}}, {{1}}, etc.",
			BodyPath: "text",
		},
	},
	Action:          handleBroadcastsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"content": {
		&requestflag.InnerFlag[string]{
			Name:       "content.filename",
			Usage:      "Filename for documents.",
			InnerField: "filename",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.media-id",
			Usage:      "Media ID if already uploaded.",
			InnerField: "mediaId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.media-url",
			Usage:      "URL of the media file.",
			InnerField: "mediaUrl",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.mime-type",
			Usage:      "MIME type of the media.",
			InnerField: "mimeType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.template-id",
			Usage:      "Template ID for template messages.",
			InnerField: "templateId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "content.template-variables",
			Usage:      "Default template variables (can be overridden per contact).",
			InnerField: "templateVariables",
		},
	},
})

var broadcastsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get broadcast",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
	},
	Action:          handleBroadcastsRetrieve,
	HideHelpCommand: true,
}

var broadcastsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a broadcast in draft status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "content",
			Usage:    "Content for non-text broadcast message types.",
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "email-html-body",
			BodyPath: "emailHtmlBody",
		},
		&requestflag.Flag[string]{
			Name:     "email-subject",
			BodyPath: "emailSubject",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			BodyPath: "text",
		},
	},
	Action:          handleBroadcastsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"content": {
		&requestflag.InnerFlag[string]{
			Name:       "content.filename",
			Usage:      "Filename for documents.",
			InnerField: "filename",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.media-id",
			Usage:      "Media ID if already uploaded.",
			InnerField: "mediaId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.media-url",
			Usage:      "URL of the media file.",
			InnerField: "mediaUrl",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.mime-type",
			Usage:      "MIME type of the media.",
			InnerField: "mimeType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.template-id",
			Usage:      "Template ID for template messages.",
			InnerField: "templateId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "content.template-variables",
			Usage:      "Default template variables (can be overridden per contact).",
			InnerField: "templateVariables",
		},
	},
})

var broadcastsList = cli.Command{
	Name:    "list",
	Usage:   "List broadcasts for this project.",
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
			Usage:     "Current status of the broadcast.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBroadcastsList,
	HideHelpCommand: true,
}

var broadcastsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a broadcast in draft status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
	},
	Action:          handleBroadcastsDelete,
	HideHelpCommand: true,
}

var broadcastsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a broadcast. Pending contacts will be skipped, but already queued\nmessages may still be delivered.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
	},
	Action:          handleBroadcastsCancel,
	HideHelpCommand: true,
}

var broadcastsEscalateReview = cli.Command{
	Name:    "escalate-review",
	Usage:   "Request manual review by the Zavu team for a rejected broadcast. Use this after\nautomated review rejection if you believe the content is legitimate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
	},
	Action:          handleBroadcastsEscalateReview,
	HideHelpCommand: true,
}

var broadcastsProgress = cli.Command{
	Name:    "progress",
	Usage:   "Get real-time progress of a broadcast including delivery counts and estimated\ncompletion time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
	},
	Action:          handleBroadcastsProgress,
	HideHelpCommand: true,
}

var broadcastsReschedule = cli.Command{
	Name:    "reschedule",
	Usage:   "Update the scheduled time for a broadcast. The broadcast must be in scheduled\nstatus.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "scheduled-at",
			Usage:    "New scheduled time for the broadcast.",
			Required: true,
			BodyPath: "scheduledAt",
		},
	},
	Action:          handleBroadcastsReschedule,
	HideHelpCommand: true,
}

var broadcastsRetryReview = cli.Command{
	Name:    "retry-review",
	Usage:   "Resubmit a rejected broadcast for AI review after editing content. Maximum 3\nreview attempts allowed per broadcast.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
	},
	Action:          handleBroadcastsRetryReview,
	HideHelpCommand: true,
}

var broadcastsSend = cli.Command{
	Name:    "send",
	Usage:   "Start sending the broadcast immediately or schedule for later. Broadcasts go\nthrough automated AI content review before sending. If the review passes, the\nbroadcast proceeds. If rejected, use PATCH to edit content, then call POST\n/retry-review. Reserves the estimated cost from your balance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "broadcast-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "scheduled-at",
			Usage:    "Schedule for future delivery. Omit to send immediately.",
			BodyPath: "scheduledAt",
		},
	},
	Action:          handleBroadcastsSend,
	HideHelpCommand: true,
}

func handleBroadcastsCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.BroadcastNewParams{}

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
	_, err = client.Broadcasts.New(ctx, params, options...)
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
		Title:          "broadcasts create",
		Transform:      transform,
	})
}

func handleBroadcastsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.Get(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts retrieve",
		Transform:      transform,
	})
}

func handleBroadcastsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.BroadcastUpdateParams{}

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
	_, err = client.Broadcasts.Update(
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
		Title:          "broadcasts update",
		Transform:      transform,
	})
}

func handleBroadcastsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.BroadcastListParams{}

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
		_, err = client.Broadcasts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "broadcasts list",
			Transform:      transform,
		})
	} else {
		iter := client.Broadcasts.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "broadcasts list",
			Transform:      transform,
		})
	}
}

func handleBroadcastsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Broadcasts.Delete(ctx, cmd.Value("broadcast-id").(string), options...)
}

func handleBroadcastsCancel(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.Cancel(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts cancel",
		Transform:      transform,
	})
}

func handleBroadcastsEscalateReview(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.EscalateReview(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts escalate-review",
		Transform:      transform,
	})
}

func handleBroadcastsProgress(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.Progress(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts progress",
		Transform:      transform,
	})
}

func handleBroadcastsReschedule(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.BroadcastRescheduleParams{}

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
	_, err = client.Broadcasts.Reschedule(
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
		Title:          "broadcasts reschedule",
		Transform:      transform,
	})
}

func handleBroadcastsRetryReview(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.RetryReview(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts retry-review",
		Transform:      transform,
	})
}

func handleBroadcastsSend(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.BroadcastSendParams{}

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
	_, err = client.Broadcasts.Send(
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
		Title:          "broadcasts send",
		Transform:      transform,
	})
}
