package skillgo

// ThumbnailType is Thumbnail template which located in CarouselHeaderType
type ThumbnailType struct {
	ImageURL   string   `json:"imageUrl"`
	Link       LinkType `json:"link,omitempty"`
	FixedRatio bool     `json:"fixedRatio,omitempty"`
	Width      int      `json:"width,omitempty"`
	Height     int      `json:"height,omitempty"`
}

// Thumbnail creates Thumbnail SkillResponse
func Thumbnail(imageURL string, link LinkType, fixedRatio bool, width int, height int) ThumbnailType {
	response := ThumbnailType{}
	response.ImageURL = imageURL
	response.Link = link
	response.FixedRatio = fixedRatio

	if fixedRatio {
		response.Width = width
		response.Height = height
	}

	return response
}

// LinkType is Link template
type LinkType struct {
	PC     string `json:"pc,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Web    string `json:"web,omitempty"`
}

// Link creates Link SkillResponse
func Link(pc string, mobile string, web string) LinkType {
	response := LinkType{}

	if pc != "" {
		response.PC = pc
	}
	if mobile != "" {
		response.Mobile = mobile
	}
	if web != "" {
		response.Web = web
	}

	return response
}

// ActionConfig represents actions
type ActionConfig struct {
	WebLinkURL  string                 `json:"webLinkUrl,omitempty"`
	MessageText string                 `json:"messageText,omitempty"`
	PhoneNumber string                 `json:"phoneNumber,omitempty"`
	BlockID     map[string]interface{} `json:"blockId,omitempty"`
}

// ButtonType is Button Template
type ButtonType struct {
	Label  string `json:"label"`
	Action string `json:"action"`
	ActionConfig
}

// Button creates Button SkillResponse
func Button(label string, action string, config ActionConfig) ButtonType {
	response := ButtonType{}

	response.Label = label
	response.Action = action

	if action == "webLink" {
		response.WebLinkURL = config.WebLinkURL
	} else if action == "message" || action == "block" {
		response.MessageText = config.MessageText
	}

	if action == "phone" {
		response.PhoneNumber = config.PhoneNumber
	} else if action == "block" {
		response.BlockID = config.BlockID
	}

	return response
}

// ProfileType is Profile template
type ProfileType struct {
	Nickname string `json:"nickname"`
	ImageURL string `json:"imageUrl"`
}

// Profile creates Profile SkillResponse
func Profile(nickname string, imageURL string) ProfileType {
	response := ProfileType{}

	response.Nickname = nickname
	if imageURL != "" {
		response.ImageURL = imageURL
	}

	return response
}

// CarouselHeaderType is CarouselHeader Template which located in Carousel
type CarouselHeaderType struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Thumbnail   ThumbnailType `json:"thumbnail"`
}

// CarouselHeader creates CarouselHeader SkillResponse
func CarouselHeader(title string, description string, thumbnail ThumbnailType) CarouselHeaderType {
	response := CarouselHeaderType{}

	response.Title = title
	response.Description = description
	response.Thumbnail = thumbnail

	return response
}

// QuickRepliesType is QucikReplies Template
type QuickRepliesType struct {
	Label       string      `json:"label"`
	Action      string      `json:"action"`
	MessageText string      `json:"messageText,omitempty"`
	BlockID     string      `json:"blockId,omitempty"`
	Extra       interface{} `json:"extra,omitempty"`
}

// QuickReplies creates QuickReplies SkillResponse
func QuickReplies(label string, action string, messageText string, blockID string, extra interface{}) QuickRepliesType {
	response := QuickRepliesType{}

	response.Label = label
	response.Action = action
	response.MessageText = messageText
	response.BlockID = blockID
	response.Extra = extra

	return response
}
