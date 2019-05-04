// Package rocket implements the Rocket.Chat integration for nions commandline client
//
// Rocket.Chat example with curl
// curl -X POST -H 'Content-Type: application/json' --data '{"text":"Example message","attachments":[{"title":"Rocket.Chat","title_link":"https://rocket.chat","text":"Rocket.Chat, the best open source chat","image_url":"/images/integration-attachment-example.png","color":"#764FA5"}]}' https://chat.example.com/hooks/$ROCKET_TOKEN
package rocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// Default message values
var (
	DefaultText                = "Hello from Nions CLI!"
	DefaultAttachmentTitle     = "Nions CLI"
	DefaultAttachmentTitleLink = "https://github.com/iwittkau/nions"
	DefaultAttachmentText      = "This is a message from the Nions CLI client."
	DefaultAttachmentImageURL  = ""
	DefaultAttachmentColor     = "#00B2AD"
)

// Exec executes the Rocket.Chat integration
func Exec() error {

	var (
		text                string
		attachmentTitle     string
		attachmentTitleLink string
		attachmentText      string
		attachmentImageURL  string
		attachmentColor     string
		short               bool
		debug               bool
	)

	flag.StringVar(&text, "message", DefaultText, "set rocket mesage")
	flag.StringVar(&attachmentTitle, "title", DefaultAttachmentTitle, "set rocket attachment title")
	flag.StringVar(&attachmentTitleLink, "link", DefaultAttachmentTitleLink, "set rocket attachment title link")
	flag.StringVar(&attachmentText, "text", DefaultAttachmentText, "set rocket attachment text")
	flag.StringVar(&attachmentImageURL, "image", DefaultAttachmentImageURL, "set rocket attachment image")
	flag.StringVar(&attachmentColor, "color", DefaultAttachmentColor, "set rocket attachment color")
	flag.BoolVar(&short, "short", false, "disable attachment")
	flag.BoolVar(&debug, "d", false, "enable debugging")
	flag.Parse()

	token := os.Getenv("NIONS_ROCKET_TOKEN")
	instance := os.Getenv("NIONS_ROCKET_INSTANCE")
	if token == "" {
		return errors.New("rocket chat: empty token")
	}
	if instance == "" {
		return errors.New("rocket chat: empty instance")
	}

	message := Message{
		Text: text,
	}

	if !short {
		message.Attachments = []Attachment{{
			Title:     attachmentTitle,
			TitleLink: attachmentTitleLink,
			Text:      attachmentText,
			ImageURL:  attachmentImageURL,
			Color:     attachmentColor,
		}}
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	u := url.URL{}
	u.Scheme = "https"
	u.Host = instance
	u.Path = "hooks/" + token

	resp, err := http.DefaultClient.Post(u.String(), "application/json", bytes.NewBuffer(data))
	if debug {
		fmt.Println(u.String())
		fmt.Println(string(data))
	}
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("rocket chat: status_code=%d status=%s", resp.StatusCode, resp.Status)
}
