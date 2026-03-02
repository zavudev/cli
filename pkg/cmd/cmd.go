// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/stainless-sdks/zavudev-cli/internal/autocomplete"
	"github.com/stainless-sdks/zavudev-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "zavudev",
		Usage:   "CLI for the zavudev API",
		Suggest: true,
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("ZAVUDEV_API_KEY"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagesRetrieve,
					&messagesList,
					&messagesReact,
					&messagesSend,
				},
			},
			{
				Name:     "templates",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&templatesCreate,
					&templatesRetrieve,
					&templatesList,
					&templatesDelete,
					&templatesSubmit,
				},
			},
			{
				Name:     "senders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersCreate,
					&sendersRetrieve,
					&sendersUpdate,
					&sendersList,
					&sendersDelete,
					&sendersGetProfile,
					&sendersRegenerateWebhookSecret,
					&sendersUpdateProfile,
					&sendersUploadProfilePicture,
				},
			},
			{
				Name:     "senders:agent",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersAgentCreate,
					&sendersAgentRetrieve,
					&sendersAgentUpdate,
					&sendersAgentDelete,
					&sendersAgentStats,
				},
			},
			{
				Name:     "senders:agent:executions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersAgentExecutionsList,
				},
			},
			{
				Name:     "senders:agent:flows",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersAgentFlowsCreate,
					&sendersAgentFlowsRetrieve,
					&sendersAgentFlowsUpdate,
					&sendersAgentFlowsList,
					&sendersAgentFlowsDelete,
					&sendersAgentFlowsDuplicate,
				},
			},
			{
				Name:     "senders:agent:tools",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersAgentToolsCreate,
					&sendersAgentToolsRetrieve,
					&sendersAgentToolsUpdate,
					&sendersAgentToolsList,
					&sendersAgentToolsDelete,
					&sendersAgentToolsTest,
				},
			},
			{
				Name:     "senders:agent:knowledge-bases",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersAgentKnowledgeBasesCreate,
					&sendersAgentKnowledgeBasesRetrieve,
					&sendersAgentKnowledgeBasesUpdate,
					&sendersAgentKnowledgeBasesList,
					&sendersAgentKnowledgeBasesDelete,
				},
			},
			{
				Name:     "senders:agent:knowledge-bases:documents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersAgentKnowledgeBasesDocumentsCreate,
					&sendersAgentKnowledgeBasesDocumentsList,
					&sendersAgentKnowledgeBasesDocumentsDelete,
				},
			},
			{
				Name:     "contacts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&contactsRetrieve,
					&contactsUpdate,
					&contactsList,
					&contactsRetrieveByPhone,
				},
			},
			{
				Name:     "broadcasts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&broadcastsCreate,
					&broadcastsRetrieve,
					&broadcastsUpdate,
					&broadcastsList,
					&broadcastsDelete,
					&broadcastsCancel,
					&broadcastsProgress,
					&broadcastsReschedule,
					&broadcastsSend,
				},
			},
			{
				Name:     "broadcasts:contacts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&broadcastsContactsList,
					&broadcastsContactsAdd,
					&broadcastsContactsRemove,
				},
			},
			{
				Name:     "introspect",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&introspectValidatePhone,
				},
			},
			{
				Name:     "phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&phoneNumbersRetrieve,
					&phoneNumbersUpdate,
					&phoneNumbersList,
					&phoneNumbersPurchase,
					&phoneNumbersRelease,
					&phoneNumbersRequirements,
					&phoneNumbersSearchAvailable,
				},
			},
			{
				Name:     "addresses",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&addressesCreate,
					&addressesRetrieve,
					&addressesList,
					&addressesDelete,
				},
			},
			{
				Name:     "regulatory-documents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&regulatoryDocumentsCreate,
					&regulatoryDocumentsRetrieve,
					&regulatoryDocumentsList,
					&regulatoryDocumentsDelete,
					&regulatoryDocumentsUploadURL,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "zavudev @manpages [-o zavudev.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "zavudev.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "zavudev.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
