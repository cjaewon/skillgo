package skillgo

// SimpleTextType is SimpleText Template which located in SkillTemplate.Outputs
type SimpleTextType struct {
	SimpleText struct {
		Text string `json:"text"`
	} `json:"simpleText"`
}

// SimpleText creates SimpleTextType SkillResponse
func SimpleText(text string) SimpleTextType {
	response := SimpleTextType{}
	response.SimpleText.Text = text

	return response
}

// SimpleImageType is SimpleImage Template which located in SkillTemplate.Outputs
type SimpleImageType struct {
	SimpleImage struct {
		ImageURL string `json:"imageUrl"`
		AltText  string `json:"altText"`
	} `json:"simpleImage"`
}

// SimpleImage creates SimpleImage SkillResponse
func SimpleImage(imageURL string, altText string) SimpleImageType {
	response := SimpleImageType{}
	response.SimpleImage.ImageURL = imageURL
	response.SimpleImage.AltText = altText

	return response
}
