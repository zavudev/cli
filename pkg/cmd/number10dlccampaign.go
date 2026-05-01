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

var number10dlcCampaignsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a 10DLC campaign under an existing brand. The campaign starts in draft\nstatus. Submit it for carrier review using the submit endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "affiliate-marketing",
			Required: true,
			BodyPath: "affiliateMarketing",
		},
		&requestflag.Flag[bool]{
			Name:     "age-gated",
			Required: true,
			BodyPath: "ageGated",
		},
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Usage:    "ID of the brand to create this campaign under.",
			Required: true,
			BodyPath: "brandId",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "direct-lending",
			Required: true,
			BodyPath: "directLending",
		},
		&requestflag.Flag[bool]{
			Name:     "embedded-link",
			Required: true,
			BodyPath: "embeddedLink",
		},
		&requestflag.Flag[bool]{
			Name:     "embedded-phone",
			Required: true,
			BodyPath: "embeddedPhone",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "number-pooling",
			Required: true,
			BodyPath: "numberPooling",
		},
		&requestflag.Flag[[]string]{
			Name:     "sample-message",
			Required: true,
			BodyPath: "sampleMessages",
		},
		&requestflag.Flag[bool]{
			Name:     "subscriber-help",
			Required: true,
			BodyPath: "subscriberHelp",
		},
		&requestflag.Flag[bool]{
			Name:     "subscriber-opt-in",
			Required: true,
			BodyPath: "subscriberOptIn",
		},
		&requestflag.Flag[bool]{
			Name:     "subscriber-opt-out",
			Required: true,
			BodyPath: "subscriberOptOut",
		},
		&requestflag.Flag[string]{
			Name:     "use-case",
			Usage:    "Campaign use case (e.g., ACCOUNT_NOTIFICATION, MARKETING, 2FA).",
			Required: true,
			BodyPath: "useCase",
		},
		&requestflag.Flag[string]{
			Name:     "help-message",
			BodyPath: "helpMessage",
		},
		&requestflag.Flag[string]{
			Name:     "message-flow",
			BodyPath: "messageFlow",
		},
		&requestflag.Flag[[]string]{
			Name:     "opt-in-keyword",
			BodyPath: "optInKeywords",
		},
		&requestflag.Flag[[]string]{
			Name:     "opt-out-keyword",
			BodyPath: "optOutKeywords",
		},
		&requestflag.Flag[[]string]{
			Name:     "sub-use-case",
			BodyPath: "subUseCases",
		},
	},
	Action:          handleNumber10dlcCampaignsCreate,
	HideHelpCommand: true,
}

var number10dlcCampaignsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get 10DLC campaign",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "campaign-id",
			Required:  true,
			PathParam: "campaignId",
		},
	},
	Action:          handleNumber10dlcCampaignsRetrieve,
	HideHelpCommand: true,
}

var number10dlcCampaignsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a 10DLC campaign in draft status. Cannot update after submission.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "campaign-id",
			Required:  true,
			PathParam: "campaignId",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "help-message",
			BodyPath: "helpMessage",
		},
		&requestflag.Flag[string]{
			Name:     "message-flow",
			BodyPath: "messageFlow",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "opt-in-keyword",
			BodyPath: "optInKeywords",
		},
		&requestflag.Flag[[]string]{
			Name:     "opt-out-keyword",
			BodyPath: "optOutKeywords",
		},
		&requestflag.Flag[[]string]{
			Name:     "sample-message",
			BodyPath: "sampleMessages",
		},
	},
	Action:          handleNumber10dlcCampaignsUpdate,
	HideHelpCommand: true,
}

var number10dlcCampaignsList = cli.Command{
	Name:    "list",
	Usage:   "List 10DLC campaign registrations for this project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "brand-id",
			Usage:     "Filter campaigns by brand ID.",
			QueryPath: "brandId",
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
	Action:          handleNumber10dlcCampaignsList,
	HideHelpCommand: true,
}

var number10dlcCampaignsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete 10DLC campaign",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "campaign-id",
			Required:  true,
			PathParam: "campaignId",
		},
	},
	Action:          handleNumber10dlcCampaignsDelete,
	HideHelpCommand: true,
}

var number10dlcCampaignsSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Submit a draft campaign for carrier review. The campaign must be in draft status\nand its brand must be verified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "campaign-id",
			Required:  true,
			PathParam: "campaignId",
		},
	},
	Action:          handleNumber10dlcCampaignsSubmit,
	HideHelpCommand: true,
}

var number10dlcCampaignsSyncStatus = cli.Command{
	Name:    "sync-status",
	Usage:   "Sync the campaign status with the registration provider. Use this to check for\napproval updates after submission.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "campaign-id",
			Required:  true,
			PathParam: "campaignId",
		},
	},
	Action:          handleNumber10dlcCampaignsSyncStatus,
	HideHelpCommand: true,
}

func handleNumber10dlcCampaignsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.Number10dlcCampaignNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Number10dlc.Campaigns.New(ctx, params, options...)
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
		Title:          "number-10dlc:campaigns create",
		Transform:      transform,
	})
}

func handleNumber10dlcCampaignsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Number10dlc.Campaigns.Get(ctx, cmd.Value("campaign-id").(string), options...)
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
		Title:          "number-10dlc:campaigns retrieve",
		Transform:      transform,
	})
}

func handleNumber10dlcCampaignsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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

	params := zavudev.Number10dlcCampaignUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Number10dlc.Campaigns.Update(
		ctx,
		cmd.Value("campaign-id").(string),
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
		Title:          "number-10dlc:campaigns update",
		Transform:      transform,
	})
}

func handleNumber10dlcCampaignsList(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.Number10dlcCampaignListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Number10dlc.Campaigns.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "number-10dlc:campaigns list",
			Transform:      transform,
		})
	} else {
		iter := client.Number10dlc.Campaigns.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "number-10dlc:campaigns list",
			Transform:      transform,
		})
	}
}

func handleNumber10dlcCampaignsDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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

	return client.Number10dlc.Campaigns.Delete(ctx, cmd.Value("campaign-id").(string), options...)
}

func handleNumber10dlcCampaignsSubmit(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Number10dlc.Campaigns.Submit(ctx, cmd.Value("campaign-id").(string), options...)
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
		Title:          "number-10dlc:campaigns submit",
		Transform:      transform,
	})
}

func handleNumber10dlcCampaignsSyncStatus(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Number10dlc.Campaigns.SyncStatus(ctx, cmd.Value("campaign-id").(string), options...)
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
		Title:          "number-10dlc:campaigns sync-status",
		Transform:      transform,
	})
}
