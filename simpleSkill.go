package skillgo

// SimpleTextType : SimpleText Struct Type
type SimpleTextType struct {
	simpleText struct {
		text string
	}
}

// SimpleText : SimpleText SkillResponse
func SimpleText(text string) SimpleTextType {
	response := SimpleTextType{}
	response.simpleText.text = text

	return response
}

// SimpleImageType : SimpleImage Struct Type
type SimpleImageType struct {
	simpleImage struct {
		imageURL string
		altText  string
	}
}

// SimpleImage : SimpleImage SkillResponse
func SimpleImage(imageURL string, altText string) SimpleImageType {
	response := SimpleImageType{}
	response.simpleImage.imageURL = imageURL
	response.simpleImage.altText = altText

	return response
}
