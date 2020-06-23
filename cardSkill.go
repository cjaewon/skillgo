package skillgo

// BasicCardType : BasicCard Struct Type
type BasicCardType struct {
	BasicCard struct {
		Title       string        `json:"title,omitempty"`
		Description string        `json:"description,omitempty"`
		Thumbnail   ThumbnailType `json:"thumbnail"`
		// Profile 미지원
		Profile struct {
			ImageURL string `json:"imageUrl"`
			Nickname string `json:"nickname"`
		} `json:"profile,omitempty"`
		// Social 미지원
		Social struct {
			Like    int `json:"like"`
			Comment int `json:"comment"`
			Share   int `json:"share"`
		} `json:"social,omitempty"`
		Buttons []ButtonType `json:"buttons,omitempty"`
	} `json:"basicCard"`
}

// BasicCard : BasicCard SkillResponse
func BasicCard(title string, description string, thumbnail ThumbnailType, buttons []ButtonType) BasicCardType {
	resposne := BasicCardType{}
	resposne.BasicCard.Thumbnail = thumbnail

	if title != "" {
		resposne.BasicCard.Title = title
	}
	if description != "" {
		resposne.BasicCard.Description = description
	}

	if len(buttons) != 0 {
		resposne.BasicCard.Buttons = buttons
	}

	return resposne
}
