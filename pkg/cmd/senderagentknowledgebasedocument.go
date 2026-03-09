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

var sendersAgentKnowledgeBasesDocumentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Add a document to a knowledge base. The document will be automatically processed\nfor RAG.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "kb-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "content",
			Required: true,
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
	},
	Action:          handleSendersAgentKnowledgeBasesDocumentsCreate,
	HideHelpCommand: true,
}

var sendersAgentKnowledgeBasesDocumentsList = cli.Command{
	Name:    "list",
	Usage:   "List documents in a knowledge base.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "kb-id",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSendersAgentKnowledgeBasesDocumentsList,
	HideHelpCommand: true,
}

var sendersAgentKnowledgeBasesDocumentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a document from a knowledge base.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "kb-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "doc-id",
			Required: true,
		},
	},
	Action:          handleSendersAgentKnowledgeBasesDocumentsDelete,
	HideHelpCommand: true,
}

func handleSendersAgentKnowledgeBasesDocumentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("kb-id") && len(unusedArgs) > 0 {
		cmd.Set("kb-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentKnowledgeBaseDocumentNewParams{
		SenderID: cmd.Value("sender-id").(string),
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
	_, err = client.Senders.Agent.KnowledgeBases.Documents.New(
		ctx,
		cmd.Value("kb-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "senders:agent:knowledge-bases:documents create", obj, format, transform)
}

func handleSendersAgentKnowledgeBasesDocumentsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("kb-id") && len(unusedArgs) > 0 {
		cmd.Set("kb-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentKnowledgeBaseDocumentListParams{
		SenderID: cmd.Value("sender-id").(string),
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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Senders.Agent.KnowledgeBases.Documents.List(
			ctx,
			cmd.Value("kb-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "senders:agent:knowledge-bases:documents list", obj, format, transform)
	} else {
		iter := client.Senders.Agent.KnowledgeBases.Documents.ListAutoPaging(
			ctx,
			cmd.Value("kb-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "senders:agent:knowledge-bases:documents list", iter, format, transform, maxItems)
	}
}

func handleSendersAgentKnowledgeBasesDocumentsDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("doc-id") && len(unusedArgs) > 0 {
		cmd.Set("doc-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.SenderAgentKnowledgeBaseDocumentDeleteParams{
		SenderID: cmd.Value("sender-id").(string),
		KBID:     cmd.Value("kb-id").(string),
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

	return client.Senders.Agent.KnowledgeBases.Documents.Delete(
		ctx,
		cmd.Value("doc-id").(string),
		params,
		options...,
	)
}
