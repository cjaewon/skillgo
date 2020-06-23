package skillgo

import "reflect"

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

// CommerceCardType : CommerceCard Struct Type
type CommerceCardType struct {
	CommerceCard struct {
		Description     string          `json:"description"`
		Price           int             `json:"price"`
		Currency        string          `json:"currency"`
		Discount        int             `json:"discount,omitempty"`
		DiscountRate    int             `json:"discountRate,omitempty"`
		DiscountedPrice int             `json:"dicountedPrice,omitempty"`
		Thumbnails      []ThumbnailType `json:"thumbnails"`
		Profile         ProfileType     `json:"profile,omitempty"`
		Buttons         []ButtonType    `json:"buttons"`
	} `json:"commerceCard"`
}

// CommerceCard : CommerceCard SkillResponse
func CommerceCard(
	description string,
	price int,
	currency string,
	discount int,
	discountRate int,
	discountedPrice int,
	thumbnails []ThumbnailType,
	profile ProfileType,
	buttons []ButtonType,
) CommerceCardType {
	response := CommerceCardType{}

	response.CommerceCard.Description = description
	response.CommerceCard.Price = price
	response.CommerceCard.Currency = currency

	if discount != 0 {
		response.CommerceCard.Discount = discount
	}
	if discountRate != 0 {
		response.CommerceCard.DiscountRate = discountRate
	}
	if discountedPrice != 0 {
		response.CommerceCard.DiscountedPrice = discountedPrice
	}

	response.CommerceCard.Buttons = buttons

	return response
}

// ListCardType : ListCard Struct Type
type ListCardType struct {
	ListCard struct {
		Header  ListItemType   `json:"header"`
		Items   []ListItemType `json:"items"`
		Buttons []ButtonType   `json:"buttons,omitempty"`
	} `json:"listCard"`
}

// ListItemType : ListItem Struct Type
type ListItemType struct {
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	ImageURL    string   `json:"imageUrl,omitempty"`
	Link        LinkType `json:"link,omitempty"`
}

// ListCard : ListCard SkillResponse
func ListCard(header ListItemType, items []ListItemType, buttons []ButtonType) ListCardType {
	response := ListCardType{}

	response.ListCard.Header = header
	response.ListCard.Items = items

	if len(buttons) != 0 {
		response.ListCard.Buttons = buttons
	}

	return response
}

// ListItem : ListItem for ListCard
func ListItem(title string, description string, imageURL string, link LinkType) ListItemType {
	response := ListItemType{}

	response.Title = title
	if description != "" {
		response.Description = description
	}
	if imageURL != "" {
		response.ImageURL = imageURL
	}
	if !reflect.DeepEqual(link, LinkType{}) {
		response.Link = link
	}

	return response
}
