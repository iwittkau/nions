package rocket

// Message is the standard rocket chat message
type Message struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachments are the rocket.Message attachments
type Attachment struct {
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
	ImageURL  string `json:"image_url"`
	Color     string `json:"color"`
}

// message , attachment: title, image url, color
