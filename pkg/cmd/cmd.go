// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
	"github.com/zavudev/cli/internal/autocomplete"
	"github.com/zavudev/cli/internal/requestflag"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "zavudev",
		Usage:     "CLI for the zavudev API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
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
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
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
					&messagesListAttachments,
					&messagesReact,
					&messagesSend,
					&messagesShowTyping,
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
					&templatesSync,
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
					&sendersAgentExecutionsRetrieve,
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
					&sendersAgentToolsListTestRuns,
					&sendersAgentToolsTest,
				},
			},
			{
				Name:     "senders:agent:tools:webhook",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersAgentToolsWebhookRotateSecret,
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
					&sendersAgentKnowledgeBasesDocumentsRetrieveDocument,
					&sendersAgentKnowledgeBasesDocumentsUpdateDocument,
				},
			},
			{
				Name:     "senders:whatsapp-sync",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersWhatsappSyncRetrieve,
					&sendersWhatsappSyncStartContactsSync,
					&sendersWhatsappSyncStartHistorySync,
				},
			},
			{
				Name:     "senders:telegram",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendersTelegramConnect,
					&sendersTelegramDisconnect,
				},
			},
			{
				Name:     "contacts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&contactsCreate,
					&contactsRetrieve,
					&contactsUpdate,
					&contactsList,
					&contactsDelete,
					&contactsDismissMergeSuggestion,
					&contactsMerge,
					&contactsRetrieveByPhone,
				},
			},
			{
				Name:     "contacts:channels",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&contactsChannelsUpdate,
					&contactsChannelsAdd,
					&contactsChannelsRemove,
					&contactsChannelsSetPrimary,
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
					&broadcastsEscalateReview,
					&broadcastsProgress,
					&broadcastsReschedule,
					&broadcastsRetryReview,
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
					&introspectValidateEmail,
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
				Name:     "invitations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&invitationsCreate,
					&invitationsRetrieve,
					&invitationsList,
					&invitationsCancel,
				},
			},
			{
				Name:     "urls",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&urlsEscalate,
					&urlsListVerified,
					&urlsRetrieveDetails,
					&urlsSubmitForVerification,
				},
			},
			{
				Name:     "balance",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&balanceRetrieve,
				},
			},
			{
				Name:     "sub-accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&subAccountsCreate,
					&subAccountsRetrieve,
					&subAccountsUpdate,
					&subAccountsList,
					&subAccountsDeactivate,
					&subAccountsGetBalance,
				},
			},
			{
				Name:     "sub-accounts:api-keys",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&subAccountsAPIKeysCreate,
					&subAccountsAPIKeysList,
					&subAccountsAPIKeysRevoke,
				},
			},
			{
				Name:     "number-10dlc:brands",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&number10dlcBrandsCreate,
					&number10dlcBrandsRetrieve,
					&number10dlcBrandsUpdate,
					&number10dlcBrandsList,
					&number10dlcBrandsDelete,
					&number10dlcBrandsListUseCases,
					&number10dlcBrandsSubmit,
					&number10dlcBrandsSyncStatus,
				},
			},
			{
				Name:     "number-10dlc:campaigns",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&number10dlcCampaignsCreate,
					&number10dlcCampaignsRetrieve,
					&number10dlcCampaignsUpdate,
					&number10dlcCampaignsList,
					&number10dlcCampaignsDelete,
					&number10dlcCampaignsSubmit,
					&number10dlcCampaignsSyncStatus,
				},
			},
			{
				Name:     "number-10dlc:campaigns:phone-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&number10dlcCampaignsPhoneNumbersList,
					&number10dlcCampaignsPhoneNumbersAssign,
					&number10dlcCampaignsPhoneNumbersUnassign,
				},
			},
			{
				Name:     "me",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&meRetrieve,
				},
			},
			{
				Name:     "functions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&functionsCreate,
					&functionsRetrieve,
					&functionsUpdate,
					&functionsDelete,
					&functionsDeploy,
					&functionsGetDeployment,
					&functionsListDeployments,
					&functionsListEventTypes,
					&functionsRollbackDeployment,
					&functionsTailLogs,
				},
			},
			{
				Name:     "functions:secrets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&functionsSecretsList,
					&functionsSecretsSet,
					&functionsSecretsUnset,
				},
			},
			{
				Name:     "functions:triggers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&functionsTriggersCreate,
					&functionsTriggersUpdate,
					&functionsTriggersList,
					&functionsTriggersDelete,
				},
			},
			{
				Name:     "functions:git-link",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&functionsGitLinkRetrieve,
					&functionsGitLinkUpdate,
					&functionsGitLinkDeployNow,
					&functionsGitLinkLink,
					&functionsGitLinkUnlink,
				},
			},
			{
				Name:     "conversations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&conversationsRetrieve,
					&conversationsList,
					&conversationsListMessages,
					&conversationsMarkAsRead,
				},
			},
			{
				Name:     "calls",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&callsCreate,
					&callsRetrieve,
					&callsList,
					&callsHangup,
				},
			},
			{
				Name:     "agent-templates",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentTemplatesRetrieve,
					&agentTemplatesList,
				},
			},
			{
				Name:     "email-domains",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&emailDomainsCreate,
					&emailDomainsRetrieve,
					&emailDomainsList,
					&emailDomainsDelete,
					&emailDomainsVerify,
				},
			},
			{
				Name:     "agents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentsCreate,
					&agentsRetrieve,
					&agentsUpdate,
					&agentsList,
					&agentsDelete,
					&agentsListVoices,
					&agentsTest,
				},
			},
			{
				Name:     "agents:senders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentsSendersConnect,
					&agentsSendersDisconnect,
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
