package main

import (
	"encoding/json"
	"fmt"
	"github.com/muto_gcm"
	"github.com/parnurzeal/gorequest"
)

func main() {

	var request muto_gcm.Request
	var combined muto_gcm.Combined

	request.Email = "Bhavyantkolli@gmail.com"
	request.Username = "bhavya"
	combined.Priority = "high"
	combined.Contents.Title = "Combined Working or not?"
	combined.Contents.Body = "Yes, it is working man."

	args := make([]string, 1)
	args[0] = "localized args"

	combined.AlertDict.ActionLocKey = "Play a Game!"
	combined.AlertDict.LocKey = "localized key"
	combined.AlertDict.LocArgs = args
	combined.AlertDict.LaunchImage = "image.jpg"
	combined.Badge = 42
	combined.Sound = "bingbong.aiff"

	b, _ := json.Marshal(request)
	req := gorequest.New()
	resp, body, errs := req.Post("http://localhost:9010/token").
		Set("Content-Type", "application/json").
		Send(string(b)).
		End()
	fmt.Println("\n", resp, body, errs)
}
