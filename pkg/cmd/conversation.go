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

var conversationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get conversation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "conversation-id",
			Required:  true,
			PathParam: "conversationId",
		},
	},
	Action:          handleConversationsRetrieve,
	HideHelpCommand: true,
}

var conversationsList = cli.Command{
	Name:    "list",
	Usage:   "List inbox threads, most recently active first. A conversation groups every\nmessage with one contact across channels, which is what you need to build an\ninbox: `GET /v1/messages` returns a flat log with no thread to hang it on.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "channel",
			Usage:     "Keep only threads that have carried this channel.",
			QueryPath: "channel",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque cursor from a previous response's `nextCursor`. Do not construct it.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Search threads by identity: phone number (any format — `+1 (555) 123-4567` and `15551234567` both match), email address (full or local part), WhatsApp group subject, WhatsApp username, or BSUID. Matching is by whole word, with prefix matching on the last term, so `mar` finds `maria@example.com` and `+1555` finds `+15551234567`; a fragment from the middle or end of a number (`4567`) does not match.\n\nIt does **not** search message bodies — only who the thread is with.\n\nResults come back ranked by relevance rather than by recency, so the usual \"most recently active first\" ordering does not apply while `q` is set. `senderId` and `channel` still narrow the results, and `cursor` paginates them as usual. An empty or whitespace-only `q` returns no items rather than the full list.",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Usage:     "Keep only threads last handled by this sender.",
			QueryPath: "senderId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleConversationsList,
	HideHelpCommand: true,
}

var conversationsListMessages = cli.Command{
	Name:    "list-messages",
	Usage:   "Messages in this thread, newest first, across every channel it has carried.\nReply with `POST /v1/messages`, passing the conversation's `senderId` as the\n`Zavu-Sender` header so the answer leaves from the number the contact already\nknows.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "conversation-id",
			Required:  true,
			PathParam: "conversationId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque cursor from a previous response's `nextCursor`.",
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
	Action:          handleConversationsListMessages,
	HideHelpCommand: true,
}

var conversationsMarkAsRead = cli.Command{
	Name:    "mark-as-read",
	Usage:   "Reset the thread's `unreadCount` to zero. Marks the thread read in your own\ninbox only: it does not send a read receipt to the contact.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "conversation-id",
			Required:  true,
			PathParam: "conversationId",
		},
	},
	Action:          handleConversationsMarkAsRead,
	HideHelpCommand: true,
}

func handleConversationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
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
	_, err = client.Conversations.Get(ctx, cmd.Value("conversation-id").(string), options...)
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
		Title:          "conversations retrieve",
		Transform:      transform,
	})
}

func handleConversationsList(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.ConversationListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Conversations.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "conversations list",
			Transform:      transform,
		})
	} else {
		iter := client.Conversations.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "conversations list",
			Transform:      transform,
		})
	}
}

func handleConversationsListMessages(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
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

	params := zavudev.ConversationListMessagesParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Conversations.ListMessages(
			ctx,
			cmd.Value("conversation-id").(string),
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
			Title:          "conversations list-messages",
			Transform:      transform,
		})
	} else {
		iter := client.Conversations.ListMessagesAutoPaging(
			ctx,
			cmd.Value("conversation-id").(string),
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
			Title:          "conversations list-messages",
			Transform:      transform,
		})
	}
}

func handleConversationsMarkAsRead(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
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
	_, err = client.Conversations.MarkAsRead(ctx, cmd.Value("conversation-id").(string), options...)
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
		Title:          "conversations mark-as-read",
		Transform:      transform,
	})
}
