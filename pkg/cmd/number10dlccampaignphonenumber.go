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

var number10dlcCampaignsPhoneNumbersList = cli.Command{
	Name:    "list",
	Usage:   "List phone numbers assigned to a 10DLC campaign.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "campaign-id",
			Required:  true,
			PathParam: "campaignId",
		},
	},
	Action:          handleNumber10dlcCampaignsPhoneNumbersList,
	HideHelpCommand: true,
}

var number10dlcCampaignsPhoneNumbersAssign = cli.Command{
	Name:    "assign",
	Usage:   "Assign a US phone number to an approved 10DLC campaign. The campaign must be in\napproved status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "campaign-id",
			Required:  true,
			PathParam: "campaignId",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number-id",
			Usage:    "ID of the phone number to assign.",
			Required: true,
			BodyPath: "phoneNumberId",
		},
	},
	Action:          handleNumber10dlcCampaignsPhoneNumbersAssign,
	HideHelpCommand: true,
}

var number10dlcCampaignsPhoneNumbersUnassign = cli.Command{
	Name:    "unassign",
	Usage:   "Remove a phone number assignment from a 10DLC campaign.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "campaign-id",
			Required:  true,
			PathParam: "campaignId",
		},
		&requestflag.Flag[string]{
			Name:      "assignment-id",
			Required:  true,
			PathParam: "assignmentId",
		},
	},
	Action:          handleNumber10dlcCampaignsPhoneNumbersUnassign,
	HideHelpCommand: true,
}

func handleNumber10dlcCampaignsPhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Number10dlc.Campaigns.PhoneNumbers.List(ctx, cmd.Value("campaign-id").(string), options...)
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
		Title:          "number-10dlc:campaigns:phone-numbers list",
		Transform:      transform,
	})
}

func handleNumber10dlcCampaignsPhoneNumbersAssign(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.Number10dlcCampaignPhoneNumberAssignParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Number10dlc.Campaigns.PhoneNumbers.Assign(
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
		Title:          "number-10dlc:campaigns:phone-numbers assign",
		Transform:      transform,
	})
}

func handleNumber10dlcCampaignsPhoneNumbersUnassign(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assignment-id") && len(unusedArgs) > 0 {
		cmd.Set("assignment-id", unusedArgs[0])
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

	params := zavudev.Number10dlcCampaignPhoneNumberUnassignParams{
		CampaignID: cmd.Value("campaign-id").(string),
	}

	return client.Number10dlc.Campaigns.PhoneNumbers.Unassign(
		ctx,
		cmd.Value("assignment-id").(string),
		params,
		options...,
	)
}
