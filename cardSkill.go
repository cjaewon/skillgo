package skillgo

// BasicCardType : BasicCard Struct Type
type BasicCardType struct {
	BasicCard struct {
		Title       string `json:"title, omitempty"`
		Description string `json:"description, omitempty"`
		Thumbnail   struct {
			ImageURL string `json:"imageUrl"`
		} `json:"thumbnail"`
		// Profile 미지원
		Profile struct {
			ImageURL string `json:"imageUrl"`
			Nickname string `json:"nickname"`
		} `json:"profile, omitempty"`
		// Social 미지원
		Social struct {
			Like    int `json:"like"`
			Comment int `json:"comment"`
			Share   int `json:"share"`
		} `json:"social, omitempty"`
		Buttons []struct {
			Action      string `json:"action"`
			Label       string `json:"label"`
			MessageText string `json:"messageText,omitempty"`
			WebLinkURL  string `json:"webLinkUrl,omitempty"`
		} `json:"buttons"`
	} `json:"basicCard"`
}

// BasicCard : BasicCard SkillResponse
func BasicCard() BasicCardType {
	resposne := BasicCardType{}

	return resposne
}