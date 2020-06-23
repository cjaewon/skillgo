package skillgo

// ThumbnailType : Thumbnail Struct Type
type ThumbnailType struct {
	ImageURL   string   `json:"imageUrl"`
	Link       LinkType `json:"link,omitempty"`
	FixedRatio bool     `json:"fixedRatio,omitempty"`
	Width      int      `json:"width,omitempty"`
	Height     int      `json:"height,omitempty"`
}

// Thumbnail : Thumbnail SkillResponse
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

// LinkType : Link Struct Type
type LinkType struct {
	PC     string `json:"pc,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Web    string `json:"web,omitempty"`
}

// Link : Link SkillResponse
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

// ActionConfig : Config in Action
type ActionConfig struct {
	WebLinkURL  string                 `json:"webLinkUrl,omitempty"`
	MessageText string                 `json:"messageText,omitempty"`
	PhoneNumber string                 `json:"phoneNumber,omitempty"`
	BlockID     map[string]interface{} `json:"blockId,omitempty"`
}

// ButtonType : Button Struct Type
type ButtonType struct {
	Label  string `json:"label"`
	Action string `json:"action"`
	ActionConfig
}

// Button : Button SkillResponse
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

// ProfileType : Profile Struct Type
type ProfileType struct {
	Nickname string `json:"nickname"`
	ImageURL string `json:"imageUrl"`
}

// Profile : Profile SkillResponse
func Profile(nickname string, imageURL string) ProfileType {
	response := ProfileType{}

	response.Nickname = nickname
	if imageURL != "" {
		response.ImageURL = imageURL
	}

	return response
}

// CarouselHeaderType : CarouselHeader Struct Type
type CarouselHeaderType struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Thumbnail   ThumbnailType `json:"thumbnail"`
}

// CarouselHeader : CarouselHeader SkillResponse
func CarouselHeader(title string, description string, thumbnail ThumbnailType) CarouselHeaderType {
	response := CarouselHeaderType{}

	response.Title = title
	response.Description = description
	response.Thumbnail = thumbnail

	return response
}
