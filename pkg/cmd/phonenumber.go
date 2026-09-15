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

var phoneNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get details of a specific phone number.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number-id",
			Required:  true,
			PathParam: "phoneNumberId",
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
			Name:      "phone-number-id",
			Required:  true,
			PathParam: "phoneNumberId",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Custom name for the phone number. Set to null to clear.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "sender-id",
			Usage:    "Sender ID to assign the phone number to. Set to null to unassign. A number under regulatory review is recorded now and connected to the sender when approved; a rejected number is refused.",
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
			Usage:     "Billing state of an owned number, separate from `regulatoryStatus`. `pending` is legacy and is not written to numbers today. The SDKs carry `active`, `suspended` and `pending` only; `releasing` and `released` are returned by the REST API until their next release.",
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

var phoneNumbersPurchase = requestflag.WithInnerFlags(cli.Command{
	Name:    "purchase",
	Usage:   "Purchase an available phone number. Requires a paid plan: the Free plan cannot\npurchase phone numbers and receives `402` with code `paid_plan_required`.",
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
		&requestflag.Flag[[]map[string]any]{
			Name:     "regulatory-requirement",
			Usage:    "Regulatory information, for numbers whose requirements list is not empty. Get the list with `GET /v1/phone-numbers/requirements?phoneNumber=...` and send one entry per requirement id, except `action` requirements, which take no value. Every required id must be present, once, and no unknown id may be sent; otherwise the purchase is refused with `400 invalid_request` before anything is charged.\n\nThe information is kept for your project under the number's country and `type`. A later purchase there may omit this field if what is kept still covers that number's requirements. Omit it for numbers without requirements.",
			BodyPath: "regulatoryRequirements",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of phone number. `mobile` is stocked in countries where no geographic (`local`) or non-geographic (`national`) inventory exists, and in several markets it is the only type that can receive SMS.",
			BodyPath: "type",
		},
	},
	Action:          handlePhoneNumbersPurchase,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"regulatory-requirement": {
		&requestflag.InnerFlag[string]{
			Name:       "regulatory-requirement.field-value",
			Usage:      "Depends on the requirement's `type`: the text itself for `textual`; for `address`, the `id` of an address created in this project with `POST /v1/addresses`; for `document`, the `id` of a document created with `POST /v1/documents`. An address or document from another project, or one rejected in review, is refused.",
			InnerField: "fieldValue",
		},
		&requestflag.InnerFlag[string]{
			Name:       "regulatory-requirement.requirement-type",
			Usage:      "A `requirementTypes[].id` from `GET /v1/phone-numbers/requirements`. Each id may appear only once.",
			InnerField: "requirementType",
		},
	},
})

var phoneNumbersRelease = cli.Command{
	Name:    "release",
	Usage:   "Release a phone number. The phone number must not be assigned to a sender.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number-id",
			Required:  true,
			PathParam: "phoneNumberId",
		},
	},
	Action:          handlePhoneNumbersRelease,
	HideHelpCommand: true,
}

var phoneNumbersRequirements = cli.Command{
	Name:    "requirements",
	Usage:   "Get the regulatory information needed to buy a phone number, for one specific\nnumber or for a country and number type. Prefer `phoneNumber`: the response is\nthen exactly the list the purchase of that number validates against. Pass each\n`requirementTypes[].id` back as `requirementType` in `regulatoryRequirements` on\n`POST /v1/phone-numbers`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Usage:     "Two-letter ISO country code. Required unless `phoneNumber` is given.",
			QueryPath: "countryCode",
		},
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Usage:     "E.164 number from `GET /v1/phone-numbers/available`, with `+` encoded as `%2B`. Returns the requirements the purchase of that number checks. Takes precedence over `countryCode`.",
			QueryPath: "phoneNumber",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Type of phone number. `mobile` is stocked in countries where no geographic (`local`) or non-geographic (`national`) inventory exists, and in several markets it is the only type that can receive SMS.",
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
			Name:      "capabilities",
			Usage:     "Comma-separated capabilities the number must have: `sms`, `voice`, `mms`. Numbers missing any of them are dropped.",
			QueryPath: "capabilities",
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
			Usage:     "Type of phone number. `mobile` is stocked in countries where no geographic (`local`) or non-geographic (`national`) inventory exists, and in several markets it is the only type that can receive SMS.",
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "phone-numbers retrieve",
		Transform:      transform,
	})
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

	params := zavudev.PhoneNumberUpdateParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "phone-numbers update",
		Transform:      transform,
	})
}

func handlePhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.PhoneNumberListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.PhoneNumbers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "phone-numbers list",
			Transform:      transform,
		})
	} else {
		iter := client.PhoneNumbers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "phone-numbers list",
			Transform:      transform,
		})
	}
}

func handlePhoneNumbersPurchase(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.PhoneNumberPurchaseParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PhoneNumbers.Purchase(ctx, params, options...)
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
		Title:          "phone-numbers purchase",
		Transform:      transform,
	})
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

	params := zavudev.PhoneNumberRequirementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PhoneNumbers.Requirements(ctx, params, options...)
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
		Title:          "phone-numbers requirements",
		Transform:      transform,
	})
}

func handlePhoneNumbersSearchAvailable(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.PhoneNumberSearchAvailableParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PhoneNumbers.SearchAvailable(ctx, params, options...)
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
		Title:          "phone-numbers search-available",
		Transform:      transform,
	})
}
