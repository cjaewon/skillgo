package skillgo

// SkillPayload : Information that the bot system delivers to the skill server
type SkillPayload struct {
	UserRequest struct {
		Timezone string                 `json:"timezone"`
		Params   map[string]interface{} `json:"params"`
		Block    struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"block"`
		Utterance string `json:"utterance"`
		Lang      string `json:"lang"`
		User      struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			Properties struct {
				PlusfriendUserKey string `json:"plusfriendUserKey"`
			} `json:"properties"`
		} `json:"user"`
	} `json:"userRequest"`
	Contexts []interface{} `json:"contexts"`
	Bot      struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"bot"`
	Action struct {
		Name        string            `json:"name"`
		ClientExtra map[string]string `json:"clientExtra"`
		Params      struct {
		} `json:"params"`
		ID           string `json:"id"`
		DetailParams struct {
		} `json:"detailParams"`
	} `json:"action"`
}

// SkillResponse : SkillResponse
type SkillResponse struct {
	Version  string        `json:"version"`
	Template SkillTemplate `json:"template,omitempty"`

	Context interface{} `json:"context,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SkillTemplate : SkillTemplate
type SkillTemplate struct {
	Outputs      []interface{} `json:"outputs,omitempty"`
	QuickReplies []interface{} `json:"quickReplies,omitempty"`
}

// ContextControl : ContextControl
type ContextControl struct {
	Values []struct {
		Name     string            `json:"name"`
		LifeSpan int               `json:"lifeSpan"`
		Params   map[string]string `json:"params,omitempty"`
	}
}
