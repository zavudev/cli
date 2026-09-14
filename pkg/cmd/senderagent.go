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

var sendersAgentCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an AI agent for a sender. Each sender can have at most one agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
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
		&requestflag.Flag[string]{
			Name:     "api-key",
			Usage:    "API key for the LLM provider. Required unless provider is 'zavu'.",
			BodyPath: "apiKey",
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
			Usage:    "Voice Agent configuration. Enable this to let the agent answer and place phone calls with Zavu's managed voice pipeline. Requires the Voice Agents feature to be enabled for your team.",
			BodyPath: "voice",
		},
	},
	Action:          handleSendersAgentCreate,
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

var sendersAgentRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get the AI agent configuration for a sender.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersAgentRetrieve,
	HideHelpCommand: true,
}

var sendersAgentUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an AI agent's configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
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
	Action:          handleSendersAgentUpdate,
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

var sendersAgentDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an AI agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersAgentDelete,
	HideHelpCommand: true,
}

var sendersAgentStats = cli.Command{
	Name:    "stats",
	Usage:   "Get statistics for an AI agent including invocations, tokens, and costs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sender-id",
			Required:  true,
			PathParam: "senderId",
		},
	},
	Action:          handleSendersAgentStats,
	HideHelpCommand: true,
}

func handleSendersAgentCreate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := zavudev.SenderAgentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Agent.New(
		ctx,
		cmd.Value("sender-id").(string),
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
		Title:          "senders:agent create",
		Transform:      transform,
	})
}

func handleSendersAgentRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Senders.Agent.Get(ctx, cmd.Value("sender-id").(string), options...)
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
		Title:          "senders:agent retrieve",
		Transform:      transform,
	})
}

func handleSendersAgentUpdate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := zavudev.SenderAgentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Senders.Agent.Update(
		ctx,
		cmd.Value("sender-id").(string),
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
		Title:          "senders:agent update",
		Transform:      transform,
	})
}

func handleSendersAgentDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Senders.Agent.Delete(ctx, cmd.Value("sender-id").(string), options...)
}

func handleSendersAgentStats(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Senders.Agent.Stats(ctx, cmd.Value("sender-id").(string), options...)
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
		Title:          "senders:agent stats",
		Transform:      transform,
	})
}
