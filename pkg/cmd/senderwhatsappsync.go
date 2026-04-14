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

var sendersWhatsappSyncRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get the current sync status for a sender's WhatsApp coexistence account. Only\navailable for senders connected in coexistence mode (WhatsApp Business App +\nCloud API).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
	},
	Action:          handleSendersWhatsappSyncRetrieve,
	HideHelpCommand: true,
}

var sendersWhatsappSyncStartContactsSync = cli.Command{
	Name:    "start-contacts-sync",
	Usage:   "Initiate contact names sync from the WhatsApp Business App. This imports contact\nnames stored in the app to Zavu. Only available for coexistence accounts with\nactive status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
	},
	Action:          handleSendersWhatsappSyncStartContactsSync,
	HideHelpCommand: true,
}

var sendersWhatsappSyncStartHistorySync = cli.Command{
	Name:    "start-history-sync",
	Usage:   "Initiate message history sync from the WhatsApp Business App. This sends a\nrequest to the account owner to approve sharing their conversation history. Only\navailable for coexistence accounts with active status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sender-id",
			Required: true,
		},
	},
	Action:          handleSendersWhatsappSyncStartHistorySync,
	HideHelpCommand: true,
}

func handleSendersWhatsappSyncRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
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
	_, err = client.Senders.WhatsappSync.Get(ctx, cmd.Value("sender-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "senders:whatsapp-sync retrieve", obj, format, transform)
}

func handleSendersWhatsappSyncStartContactsSync(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
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
	_, err = client.Senders.WhatsappSync.StartContactsSync(ctx, cmd.Value("sender-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "senders:whatsapp-sync start-contacts-sync", obj, format, transform)
}

func handleSendersWhatsappSyncStartHistorySync(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sender-id") && len(unusedArgs) > 0 {
		cmd.Set("sender-id", unusedArgs[0])
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
	_, err = client.Senders.WhatsappSync.StartHistorySync(ctx, cmd.Value("sender-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "senders:whatsapp-sync start-history-sync", obj, format, transform)
}
