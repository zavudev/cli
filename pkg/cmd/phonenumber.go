// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/zavudev-cli/internal/apiquery"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
	"github.com/stainless-sdks/zavudev-go"
	"github.com/stainless-sdks/zavudev-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var phoneNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get details of a specific phone number.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number-id",
			Required: true,
		},
	},
	Action:          handlePhoneNumbersRetrieve,
	HideHelpCommand: true,
}

var phoneNumbersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a phone number's name or sender assignment.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Custom name for the phone number. Set to null to clear.",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "sender-id",
			Usage:    "Sender ID to assign the phone number to. Set to null to unassign.",
			BodyPath: "senderId",
		},
	},
	Action:          handlePhoneNumbersUpdate,
	HideHelpCommand: true,
}

var phoneNumbersList = cli.Command{
	Name:    "list",
	Usage:   "List all phone numbers owned by this project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handlePhoneNumbersList,
	HideHelpCommand: true,
}

var phoneNumbersPurchase = cli.Command{
	Name:    "purchase",
	Usage:   "Purchase an available phone number. The first US phone number is free for each\nteam.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "Phone number in E.164 format.",
			Required: true,
			BodyPath: "phoneNumber",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Optional custom name for the phone number.",
			BodyPath: "name",
		},
	},
	Action:          handlePhoneNumbersPurchase,
	HideHelpCommand: true,
}

var phoneNumbersRelease = cli.Command{
	Name:    "release",
	Usage:   "Release a phone number. The phone number must not be assigned to a sender.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number-id",
			Required: true,
		},
	},
	Action:          handlePhoneNumbersRelease,
	HideHelpCommand: true,
}

var phoneNumbersRequirements = cli.Command{
	Name:    "requirements",
	Usage:   "Get regulatory requirements for purchasing phone numbers in a specific country.\nSome countries require additional documentation (addresses, identity documents)\nbefore phone numbers can be activated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Usage:     "Two-letter ISO country code.",
			Required:  true,
			QueryPath: "countryCode",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			QueryPath: "type",
		},
	},
	Action:          handlePhoneNumbersRequirements,
	HideHelpCommand: true,
}

var phoneNumbersSearchAvailable = cli.Command{
	Name:    "search-available",
	Usage:   "Search for available phone numbers to purchase by country and type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Usage:     "Two-letter ISO country code.",
			Required:  true,
			QueryPath: "countryCode",
		},
		&requestflag.Flag[string]{
			Name:      "contains",
			Usage:     "Search for numbers containing this string.",
			QueryPath: "contains",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return.",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			QueryPath: "type",
		},
	},
	Action:          handlePhoneNumbersSearchAvailable,
	HideHelpCommand: true,
}

func handlePhoneNumbersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
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
	_, err = client.PhoneNumbers.Get(ctx, cmd.Value("phone-number-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers retrieve", obj, format, transform)
}

func handlePhoneNumbersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.PhoneNumberUpdateParams{}

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
	_, err = client.PhoneNumbers.Update(
		ctx,
		cmd.Value("phone-number-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers update", obj, format, transform)
}

func handlePhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.PhoneNumberListParams{}

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
		_, err = client.PhoneNumbers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "phone-numbers list", obj, format, transform)
	} else {
		iter := client.PhoneNumbers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "phone-numbers list", iter, format, transform, maxItems)
	}
}

func handlePhoneNumbersPurchase(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.PhoneNumberPurchaseParams{}

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
	_, err = client.PhoneNumbers.Purchase(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers purchase", obj, format, transform)
}

func handlePhoneNumbersRelease(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
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

	return client.PhoneNumbers.Release(ctx, cmd.Value("phone-number-id").(string), options...)
}

func handlePhoneNumbersRequirements(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.PhoneNumberRequirementsParams{}

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
	_, err = client.PhoneNumbers.Requirements(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers requirements", obj, format, transform)
}

func handlePhoneNumbersSearchAvailable(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.PhoneNumberSearchAvailableParams{}

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
	_, err = client.PhoneNumbers.SearchAvailable(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers search-available", obj, format, transform)
}
