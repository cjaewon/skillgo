package skillgo

// SimpleTextType : SimpleText Struct Type
type SimpleTextType struct {
	SimpleText struct {
		Text string `json:"text"`
	} `json:"simpleText"`
}

// SimpleText : SimpleText SkillResponse
func SimpleText(text string) SimpleTextType {
	response := SimpleTextType{}
	response.SimpleText.Text = text

	return response
}

// SimpleImageType : SimpleImage Struct Type
type SimpleImageType struct {
	SimpleImage struct {
		ImageURL string `json:"imageUrl"`
		AltText  string `json:"altText"`
	} `json:"simpleImage"`
}

// SimpleImage : SimpleImage SkillResponse
func SimpleImage(imageURL string, altText string) SimpleImageType {
	response := SimpleImageType{}
	response.SimpleImage.ImageURL = imageURL
	response.SimpleImage.AltText = altText

	return response
}
