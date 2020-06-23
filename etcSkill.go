package skillgo

// ThumbnailType : Thumbnail Struct Type
type ThumbnailType struct {
	ImageURL   string   `json:"imageUrl"`
	Link       LinkType `json:"link,omitempty"`
	FixedRatio bool     `json:"fixedRatio,omitempty"`
	Width      int      `json:"width,omitempty"`
	Height     int      `json:"height,omitempty"`
}

// LinkType : Link Struct Type
type LinkType struct {
	PC     string `json:"pc,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Web    string `json:"web,omitempty"`
}

// ButtonType : Button Struct Type
type ButtonType struct {
	Label       string                 `json:"label"`
	Action      string                 `json:"action"`
	WebLinkURL  string                 `json:"webLinkUrl,omitempty"`
	MessageText string                 `json:"messageText,omitempty"`
	PhoneNumber string                 `json:"phoneNumber,omitempty"`
	BlockID     map[string]interface{} `json:"blockId,omitempty"`
}

// CarouselHeaderType : CarouselHeader Struct Type
type CarouselHeaderType struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Thumbnail   ThumbnailType `json:"thumbnail"`
}

// ProfileType : Profile Struct Type
type ProfileType struct {
	Nickname string `json:"nickname"`
	ImageURL string `json:"imageUrl"`
}
