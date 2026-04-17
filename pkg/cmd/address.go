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

var addressesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a regulatory address for phone number purchases. Some countries require a\nverified address before phone numbers can be activated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "country-code",
			Required: true,
			BodyPath: "countryCode",
		},
		&requestflag.Flag[string]{
			Name:     "locality",
			Required: true,
			BodyPath: "locality",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Required: true,
			BodyPath: "postalCode",
		},
		&requestflag.Flag[string]{
			Name:     "street-address",
			Required: true,
			BodyPath: "streetAddress",
		},
		&requestflag.Flag[string]{
			Name:     "administrative-area",
			BodyPath: "administrativeArea",
		},
		&requestflag.Flag[string]{
			Name:     "business-name",
			BodyPath: "businessName",
		},
		&requestflag.Flag[string]{
			Name:     "extended-address",
			BodyPath: "extendedAddress",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			BodyPath: "lastName",
		},
	},
	Action:          handleAddressesCreate,
	HideHelpCommand: true,
}

var addressesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a specific regulatory address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "address-id",
			Required: true,
		},
	},
	Action:          handleAddressesRetrieve,
	HideHelpCommand: true,
}

var addressesList = cli.Command{
	Name:    "list",
	Usage:   "List regulatory addresses for this project.",
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
	Action:          handleAddressesList,
	HideHelpCommand: true,
}

var addressesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a regulatory address. Cannot delete addresses that are in use.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "address-id",
			Required: true,
		},
	},
	Action:          handleAddressesDelete,
	HideHelpCommand: true,
}

func handleAddressesCreate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.AddressNewParams{}

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
	_, err = client.Addresses.New(ctx, params, options...)
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
		Title:          "addresses create",
		Transform:      transform,
	})
}

func handleAddressesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("address-id") && len(unusedArgs) > 0 {
		cmd.Set("address-id", unusedArgs[0])
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
	_, err = client.Addresses.Get(ctx, cmd.Value("address-id").(string), options...)
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
		Title:          "addresses retrieve",
		Transform:      transform,
	})
}

func handleAddressesList(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := zavudev.AddressListParams{}

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
		_, err = client.Addresses.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "addresses list",
			Transform:      transform,
		})
	} else {
		iter := client.Addresses.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "addresses list",
			Transform:      transform,
		})
	}
}

func handleAddressesDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("address-id") && len(unusedArgs) > 0 {
		cmd.Set("address-id", unusedArgs[0])
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

	return client.Addresses.Delete(ctx, cmd.Value("address-id").(string), options...)
}
