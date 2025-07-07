package singleagent

import (
	"github.com/cloudwego/eino/schema"
	"gorm.io/gorm"

	"github.com/coze-dev/coze-studio/backend/api/model/crossdomain/agentrun"
	"github.com/coze-dev/coze-studio/backend/api/model/crossdomain/plugin"
	"github.com/coze-dev/coze-studio/backend/api/model/ocean/cloud/bot_common"
	"github.com/coze-dev/coze-studio/backend/crossdomain/contract/crossworkflow"
)

type AgentRuntime struct {
	AgentVersion     string
	IsDraft          bool
	SpaceID          int64
	ConnectorID      int64
	PreRetrieveTools []*agentrun.Tool
}

type EventType string

const (
	EventTypeOfFinalAnswer  EventType = "final_answer"
	EventTypeOfToolsMessage EventType = "tools_message"
	EventTypeOfFuncCall     EventType = "func_call"
	EventTypeOfSuggest      EventType = "suggest"
	EventTypeOfKnowledge    EventType = "knowledge"
	EventTypeOfInterrupt    EventType = "interrupt"
)

type AgentEvent struct {
	EventType EventType

	FinalAnswer  *schema.StreamReader[*schema.Message]
	ToolsMessage []*schema.Message
	FuncCall     *schema.Message
	Suggest      *schema.Message
	Knowledge    []*schema.Document
	Interrupt    *InterruptInfo
}

type SingleAgent struct {
	AgentID   int64
	CreatorID int64
	SpaceID   int64
	Name      string
	Desc      string
	IconURI   string
	CreatedAt int64
	UpdatedAt int64
	Version   string
	DeletedAt gorm.DeletedAt

	VariablesMetaID         *int64
	OnboardingInfo          *bot_common.OnboardingInfo
	ModelInfo               *bot_common.ModelInfo
	Prompt                  *bot_common.PromptInfo
	Plugin                  []*bot_common.PluginInfo
	Knowledge               *bot_common.Knowledge
	Workflow                []*bot_common.WorkflowInfo
	SuggestReply            *bot_common.SuggestReplyInfo
	JumpConfig              *bot_common.JumpConfig
	BackgroundImageInfoList []*bot_common.BackgroundImageInfo
	Database                []*bot_common.Database
	ShortcutCommand         []string
}

type InterruptEventType int64

const (
	InterruptEventType_LocalPlugin         InterruptEventType = 1
	InterruptEventType_Question            InterruptEventType = 2
	InterruptEventType_RequireInfos        InterruptEventType = 3
	InterruptEventType_SceneChat           InterruptEventType = 4
	InterruptEventType_InputNode           InterruptEventType = 5
	InterruptEventType_WorkflowLocalPlugin InterruptEventType = 6
	InterruptEventType_OauthPlugin         InterruptEventType = 7
	InterruptEventType_WorkflowLLM         InterruptEventType = 100
)

type InterruptInfo struct {
	AllToolInterruptData map[string]*plugin.ToolInterruptEvent
	AllWfInterruptData   map[string]*crossworkflow.ToolInterruptEvent
	ToolCallID           string
	InterruptType        InterruptEventType
	InterruptID          string
}

type ExecuteRequest struct {
	Identity *AgentIdentity
	UserID   string

	Input        *schema.Message
	History      []*schema.Message
	ResumeInfo   *InterruptInfo
	PreCallTools []*agentrun.ToolsRetriever
}

type AgentIdentity struct {
	AgentID int64
	// State   AgentState
	Version     string
	IsDraft     bool
	ConnectorID int64
}
