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

var number10dlcBrandsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a 10DLC brand registration. The brand starts in draft status. Submit it\nfor review using the submit endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "city",
			Required: true,
			BodyPath: "city",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Two-letter ISO country code.",
			Required: true,
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name of the brand.",
			Required: true,
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "entity-type",
			Usage:    "Business entity type for 10DLC brand registration.",
			Required: true,
			BodyPath: "entityType",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			Usage:    "Contact phone in E.164 format.",
			Required: true,
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Required: true,
			BodyPath: "postalCode",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Required: true,
			BodyPath: "state",
		},
		&requestflag.Flag[string]{
			Name:     "street",
			Required: true,
			BodyPath: "street",
		},
		&requestflag.Flag[string]{
			Name:     "vertical",
			Usage:    "Industry vertical.",
			Required: true,
			BodyPath: "vertical",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "Legal company name.",
			BodyPath: "companyName",
		},
		&requestflag.Flag[string]{
			Name:     "ein",
			Usage:    "Employer Identification Number (format: XX-XXXXXXX).",
			BodyPath: "ein",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			BodyPath: "lastName",
		},
		&requestflag.Flag[string]{
			Name:     "stock-exchange",
			BodyPath: "stockExchange",
		},
		&requestflag.Flag[string]{
			Name:     "stock-symbol",
			BodyPath: "stockSymbol",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			BodyPath: "website",
		},
	},
	Action:          handleNumber10dlcBrandsCreate,
	HideHelpCommand: true,
}

var number10dlcBrandsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get 10DLC brand",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleNumber10dlcBrandsRetrieve,
	HideHelpCommand: true,
}

var number10dlcBrandsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a 10DLC brand in draft status. Cannot update after submission.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "city",
			BodyPath: "city",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			BodyPath: "companyName",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "ein",
			BodyPath: "ein",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "entity-type",
			Usage:    "Business entity type for 10DLC brand registration.",
			BodyPath: "entityType",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			BodyPath: "lastName",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			BodyPath: "postalCode",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			BodyPath: "state",
		},
		&requestflag.Flag[string]{
			Name:     "stock-exchange",
			BodyPath: "stockExchange",
		},
		&requestflag.Flag[string]{
			Name:     "stock-symbol",
			BodyPath: "stockSymbol",
		},
		&requestflag.Flag[string]{
			Name:     "street",
			BodyPath: "street",
		},
		&requestflag.Flag[string]{
			Name:     "vertical",
			BodyPath: "vertical",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			BodyPath: "website",
		},
	},
	Action:          handleNumber10dlcBrandsUpdate,
	HideHelpCommand: true,
}

var number10dlcBrandsList = cli.Command{
	Name:    "list",
	Usage:   "List 10DLC brand registrations for this project.",
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
	Action:          handleNumber10dlcBrandsList,
	HideHelpCommand: true,
}

var number10dlcBrandsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete 10DLC brand",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleNumber10dlcBrandsDelete,
	HideHelpCommand: true,
}

var number10dlcBrandsListUseCases = cli.Command{
	Name:            "list-use-cases",
	Usage:           "List available use cases for 10DLC campaign registration.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleNumber10dlcBrandsListUseCases,
	HideHelpCommand: true,
}

var number10dlcBrandsSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Submit a draft brand to The Campaign Registry (TCR) for vetting. The brand must\nbe in draft status. A $35 registration fee is charged from your balance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleNumber10dlcBrandsSubmit,
	HideHelpCommand: true,
}

var number10dlcBrandsSyncStatus = cli.Command{
	Name:    "sync-status",
	Usage:   "Sync the brand status with the registration provider. Use this to check for\napproval updates after submission.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleNumber10dlcBrandsSyncStatus,
	HideHelpCommand: true,
}

func handleNumber10dlcBrandsCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.Number10dlcBrandNewParams{}

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
	_, err = client.Number10dlc.Brands.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "number-10dlc:brands create", obj, format, explicitFormat, transform)
}

func handleNumber10dlcBrandsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Number10dlc.Brands.Get(ctx, cmd.Value("brand-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "number-10dlc:brands retrieve", obj, format, explicitFormat, transform)
}

func handleNumber10dlcBrandsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.Number10dlcBrandUpdateParams{}

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
	_, err = client.Number10dlc.Brands.Update(
		ctx,
		cmd.Value("brand-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "number-10dlc:brands update", obj, format, explicitFormat, transform)
}

func handleNumber10dlcBrandsList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.Number10dlcBrandListParams{}

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
		_, err = client.Number10dlc.Brands.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "number-10dlc:brands list", obj, format, explicitFormat, transform)
	} else {
		iter := client.Number10dlc.Brands.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "number-10dlc:brands list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleNumber10dlcBrandsDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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

	return client.Number10dlc.Brands.Delete(ctx, cmd.Value("brand-id").(string), options...)
}

func handleNumber10dlcBrandsListUseCases(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Number10dlc.Brands.ListUseCases(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "number-10dlc:brands list-use-cases", obj, format, explicitFormat, transform)
}

func handleNumber10dlcBrandsSubmit(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Number10dlc.Brands.Submit(ctx, cmd.Value("brand-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "number-10dlc:brands submit", obj, format, explicitFormat, transform)
}

func handleNumber10dlcBrandsSyncStatus(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Number10dlc.Brands.SyncStatus(ctx, cmd.Value("brand-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "number-10dlc:brands sync-status", obj, format, explicitFormat, transform)
}
