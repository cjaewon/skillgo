package skillgo

// CarouselType : Carousel Struct Type
type CarouselType struct {
	Type string `json:"type"`
	// BasicCardType, CommerceCard
	Items  []interface{}      `json:"items"`
	Header CarouselHeaderType `json:"header,omitempty"`
}

// Carousel : Carousel SkillResponse
func Carousel(Type string, items []interface{}, header CarouselHeaderType) CarouselType {
	response := CarouselType{}

	response.Type = Type
	response.Header = header

	for _, item := range items {
		switch v := item.(type) {
		case BasicCardType:
			response.Items = append(response.Items, v.BasicCard)
		case CommerceCardType:
			response.Items = append(response.Items, v.CommerceCard)
		}
	}

	return response
}
