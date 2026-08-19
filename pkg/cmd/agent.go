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

var agentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an agent without a sender. It is created disabled; connect a sender and\nenable it when you are ready for it to answer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "model",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "LLM provider for the AI agent.",
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "system-prompt",
			Required: true,
			BodyPath: "systemPrompt",
		},
		&requestflag.Flag[int64]{
			Name:     "context-window-messages",
			Default:  10,
			BodyPath: "contextWindowMessages",
		},
		&requestflag.Flag[bool]{
			Name:     "include-contact-metadata",
			Default:  true,
			BodyPath: "includeContactMetadata",
		},
		&requestflag.Flag[int64]{
			Name:     "max-tokens",
			BodyPath: "maxTokens",
		},
		&requestflag.Flag[float64]{
			Name:     "temperature",
			BodyPath: "temperature",
		},
		&requestflag.Flag[[]string]{
			Name:     "trigger-on-channel",
			Default:  []string{"*"},
			BodyPath: "triggerOnChannels",
		},
		&requestflag.Flag[[]string]{
			Name:     "trigger-on-message-type",
			Default:  []string{"text"},
			BodyPath: "triggerOnMessageTypes",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "voice",
			Usage:    "Voice Agent configuration on a sender's AI agent. Controls how the agent behaves on inbound and outbound phone calls through Zavu's managed voice pipeline (speech recognition, the agent's LLM, and speech synthesis, with real-time interruption handling). Requires the Voice Agents feature to be enabled for your team.",
			BodyPath: "voice",
		},
	},
	Action:          handleAgentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"voice": {
		&requestflag.InnerFlag[bool]{
			Name:       "voice.enabled",
			Usage:      "Whether the agent handles voice calls. When false, the sender's number is not answered by the voice agent and outbound calls are rejected.",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.greeting",
			Usage:      "Opening line the agent speaks when the call connects. If omitted, the agent waits for the caller to speak first.",
			InnerField: "greeting",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "voice.greetings",
			Usage:      "Greeting per language, keyed by language code. Used when the caller's language differs from the one `greeting` is written in.",
			InnerField: "greetings",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "voice.interruptible",
			Usage:      "Whether the caller can interrupt the agent while it is speaking (barge-in). When true, the agent stops talking as soon as the caller starts.",
			InnerField: "interruptible",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.language",
			Usage:      "BCP-47 language code used for both speech recognition and speech synthesis (e.g. `en`, `es`, `pt-BR`). Auto-detected from the recipient when omitted.",
			InnerField: "language",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "voice.max-call-duration-minutes",
			Usage:      "Hard limit on call length in minutes. The call ends automatically when reached.",
			InnerField: "maxCallDurationMinutes",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "voice.max-idle-seconds",
			Usage:      "How long the agent waits during silence before ending the call.",
			InnerField: "maxIdleSeconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.model",
			Usage:      "Model that runs the conversation, co-located in the voice network for lowest latency. Independent of the model used for text messaging. Derived from the agent's text model when omitted.",
			InnerField: "model",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "voice.record-calls",
			Usage:      "Whether the call audio is recorded.",
			InnerField: "recordCalls",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.stt-model",
			Usage:      "Speech-recognition model. Uses the default when omitted.",
			InnerField: "sttModel",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.stt-provider",
			Usage:      "Speech-recognition provider. Uses the default when omitted.",
			InnerField: "sttProvider",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.transfer-phone-number",
			Usage:      "E.164 phone number the agent can transfer the call to. When set, the agent is given a transfer tool it can use to hand the call to a human.",
			InnerField: "transferPhoneNumber",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.tts-provider",
			Usage:      "Speech-synthesis provider. Uses the default when omitted.",
			InnerField: "ttsProvider",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.tts-voice-id",
			Usage:      "Identifier of the synthesized voice that speaks. Choose from the voices available in the dashboard. Uses a neutral default when omitted.",
			InnerField: "ttsVoiceId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.voicemail-action",
			Usage:      "What the agent does when an answering machine or voicemail is detected on an outbound call.",
			InnerField: "voicemailAction",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.voicemail-message",
			Usage:      "Message spoken when `voicemailAction` is `leave_message`. Falls back to `greeting` when omitted.",
			InnerField: "voicemailMessage",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "voice.voice-speed",
			Usage:      "Speech rate. 1.0 is natural. Only honoured by voices that support rate control; ignored by the others.",
			InnerField: "voiceSpeed",
		},
	},
})

var agentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
	},
	Action:          handleAgentsRetrieve,
	HideHelpCommand: true,
}

var agentsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:     "api-key",
			BodyPath: "apiKey",
		},
		&requestflag.Flag[int64]{
			Name:     "context-window-messages",
			BodyPath: "contextWindowMessages",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "include-contact-metadata",
			BodyPath: "includeContactMetadata",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-tokens",
			BodyPath: "maxTokens",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "LLM provider for the AI agent.",
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "system-prompt",
			BodyPath: "systemPrompt",
		},
		&requestflag.Flag[*float64]{
			Name:     "temperature",
			BodyPath: "temperature",
		},
		&requestflag.Flag[[]string]{
			Name:     "trigger-on-channel",
			BodyPath: "triggerOnChannels",
		},
		&requestflag.Flag[[]string]{
			Name:     "trigger-on-message-type",
			BodyPath: "triggerOnMessageTypes",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "voice",
			Usage:    "Voice Agent configuration. Patch this object to enable voice, change the greeting, or adjust call limits. Requires the Voice Agents feature to be enabled for your team.",
			BodyPath: "voice",
		},
	},
	Action:          handleAgentsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"voice": {
		&requestflag.InnerFlag[bool]{
			Name:       "voice.enabled",
			Usage:      "Whether the agent handles voice calls. When false, the sender's number is not answered by the voice agent and outbound calls are rejected.",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.greeting",
			Usage:      "Opening line the agent speaks when the call connects. If omitted, the agent waits for the caller to speak first.",
			InnerField: "greeting",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "voice.greetings",
			Usage:      "Greeting per language, keyed by language code. Used when the caller's language differs from the one `greeting` is written in.",
			InnerField: "greetings",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "voice.interruptible",
			Usage:      "Whether the caller can interrupt the agent while it is speaking (barge-in). When true, the agent stops talking as soon as the caller starts.",
			InnerField: "interruptible",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.language",
			Usage:      "BCP-47 language code used for both speech recognition and speech synthesis (e.g. `en`, `es`, `pt-BR`). Auto-detected from the recipient when omitted.",
			InnerField: "language",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "voice.max-call-duration-minutes",
			Usage:      "Hard limit on call length in minutes. The call ends automatically when reached.",
			InnerField: "maxCallDurationMinutes",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "voice.max-idle-seconds",
			Usage:      "How long the agent waits during silence before ending the call.",
			InnerField: "maxIdleSeconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.model",
			Usage:      "Model that runs the conversation, co-located in the voice network for lowest latency. Independent of the model used for text messaging. Derived from the agent's text model when omitted.",
			InnerField: "model",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "voice.record-calls",
			Usage:      "Whether the call audio is recorded.",
			InnerField: "recordCalls",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.stt-model",
			Usage:      "Speech-recognition model. Uses the default when omitted.",
			InnerField: "sttModel",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.stt-provider",
			Usage:      "Speech-recognition provider. Uses the default when omitted.",
			InnerField: "sttProvider",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.transfer-phone-number",
			Usage:      "E.164 phone number the agent can transfer the call to. When set, the agent is given a transfer tool it can use to hand the call to a human.",
			InnerField: "transferPhoneNumber",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.tts-provider",
			Usage:      "Speech-synthesis provider. Uses the default when omitted.",
			InnerField: "ttsProvider",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.tts-voice-id",
			Usage:      "Identifier of the synthesized voice that speaks. Choose from the voices available in the dashboard. Uses a neutral default when omitted.",
			InnerField: "ttsVoiceId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.voicemail-action",
			Usage:      "What the agent does when an answering machine or voicemail is detected on an outbound call.",
			InnerField: "voicemailAction",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.voicemail-message",
			Usage:      "Message spoken when `voicemailAction` is `leave_message`. Falls back to `greeting` when omitted.",
			InnerField: "voicemailMessage",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "voice.voice-speed",
			Usage:      "Speech rate. 1.0 is natural. Only honoured by voices that support rate control; ignored by the others.",
			InnerField: "voiceSpeed",
		},
	},
})

var agentsList = cli.Command{
	Name:    "list",
	Usage:   "Every agent in the project, newest first — including agents that are not\nconnected to any sender yet, which the sender-scoped routes cannot reach. Each\nitem carries `senderIds`, the senders the agent answers on.",
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
	Action:          handleAgentsList,
	HideHelpCommand: true,
}

var agentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
	},
	Action:          handleAgentsDelete,
	HideHelpCommand: true,
}

var agentsListVoices = cli.Command{
	Name:    "list-voices",
	Usage:   "The voices an agent can speak with, for `voice.ttsVoiceId`. Filter by `language`\nto get the ones that speak it; a voice can still be used with `language: auto`,\nwhere the agent follows the caller and keeps the chosen voice.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "language",
			Usage:     "BCP-47 tag (`en`, `es`, `pt-BR`). Omit, or pass `auto`, for every voice.",
			QueryPath: "language",
		},
	},
	Action:          handleAgentsListVoices,
	HideHelpCommand: true,
}

var agentsTest = requestflag.WithInnerFlags(cli.Command{
	Name:    "test",
	Usage:   "Run the agent's prompt, model and knowledge base against a message and return\nthe reply instead of delivering it. Writes nothing and charges nothing, so it is\nsafe to call repeatedly while iterating on a prompt.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:     "message",
			Usage:    "What to say to the agent.",
			Required: true,
			BodyPath: "message",
		},
		&requestflag.Flag[bool]{
			Name:     "execute-tools",
			Usage:    "Run the tools the agent calls instead of reporting the choice and stopping.\n\nOff by default because a tool handler talks to the outside world: a rehearsal that charges a card is not a rehearsal. Turn it on to exercise the loop that actually matters — the model picks a tool, the handler answers, the model replies with the result — without sending a message to anyone. What ran comes back in `executedToolCalls`.",
			Default:  false,
			BodyPath: "executeTools",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "history",
			Usage:    "Prior turns, oldest first, to exercise multi-turn behaviour without persisting a thread. Trimmed to the agent's context window.",
			BodyPath: "history",
		},
		&requestflag.Flag[bool]{
			Name:     "use-knowledge-base",
			Usage:    "Set false to skip retrieval and isolate prompt behaviour from the knowledge base.",
			Default:  true,
			BodyPath: "useKnowledgeBase",
		},
	},
	Action:          handleAgentsTest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"history": {
		&requestflag.InnerFlag[string]{
			Name:       "history.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "history.role",
			Usage:      `Allowed values: "user", "assistant".`,
			InnerField: "role",
		},
	},
})

func handleAgentsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.AgentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.New(ctx, params, options...)
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
		Title:          "agents create",
		Transform:      transform,
	})
}

func handleAgentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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
	_, err = client.Agents.Get(ctx, cmd.Value("agent-id").(string), options...)
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
		Title:          "agents retrieve",
		Transform:      transform,
	})
}

func handleAgentsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := zavudev.AgentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Update(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "agents update",
		Transform:      transform,
	})
}

func handleAgentsList(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.AgentListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Agents.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "agents list",
			Transform:      transform,
		})
	} else {
		iter := client.Agents.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "agents list",
			Transform:      transform,
		})
	}
}

func handleAgentsDelete(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	return client.Agents.Delete(ctx, cmd.Value("agent-id").(string), options...)
}

func handleAgentsListVoices(ctx context.Context, cmd *cli.Command) error {
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

	params := zavudev.AgentListVoicesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.ListVoices(ctx, params, options...)
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
		Title:          "agents list-voices",
		Transform:      transform,
	})
}

func handleAgentsTest(ctx context.Context, cmd *cli.Command) error {
	client := zavudev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := zavudev.AgentTestParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Test(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "agents test",
		Transform:      transform,
	})
}
