package main

import (
	"encoding/json"
	"fmt"
	apns "github.com/anachronistic/apns"
	"github.com/gin-gonic/gin"
	"github.com/muto_gcm"
	"github.com/parnurzeal/gorequest"
)

func main() {

	r := gin.Default()

	r.POST("/send", func(c *gin.Context) {

		var combined muto_gcm.Combined
		var notification muto_gcm.Notification

		c.BindJSON(&combined)

		// Mapping the Notification content and other options to the corresponding server
		// i.e. GCM for Android and APNs for iOS
		for _, val := range combined.DevicesList {

			if val.Platform == "Android" {

				notification.TargetDeviceIDs = notification.TargetDeviceIDs[:0] // Empty the slice for previous entries
				notification.TargetDeviceIDs = append(notification.TargetDeviceIDs, val.DeviceID)

				notification.CollapseKey = combined.CollapseKey
				notification.Priority = combined.Priority
				notification.ContentAvailable = combined.ContentAvailableGcm
				notification.DelayWhileIdle = combined.DelayWhileIdle
				notification.TimeToLive = combined.TimeToLive
				notification.RestrictedPackageName = combined.RestrictedPackageName
				notification.DryRun = combined.DryRun

				notification.Payload.Title = combined.Contents.Title
				notification.Payload.Body = combined.Contents.Body
				notification.Payload.Icon = combined.Payload.Icon
				notification.Payload.Sound = combined.Payload.Sound
				notification.Payload.Tag = combined.Payload.Tag
				notification.Payload.Color = combined.Payload.Color
				notification.Payload.ClickAction = combined.Payload.ClickAction
				notification.Payload.BodyLocKey = combined.Payload.BodyLocKey
				notification.Payload.BodyLocArgs = combined.Payload.BodyLocArgs
				notification.Payload.TitleLocKey = combined.Payload.TitleLocKey
				notification.Payload.TitleLocArgs = combined.Payload.TitleLocArgs

				b, _ := json.Marshal(notification) // Convert the Notification object into a JSON object

				request := gorequest.New()
				resp, body, errs := request.Post(muto_gcm.GcmAPI).
					Set("Content-Type", c.ContentType()).
					Set("Authorization", muto_gcm.Authorization).
					Send(string(b)).
					End()
				fmt.Print("Response : %#v %#v %#v", resp, "\n", "Body : ", body, "\n", "Error : ", errs)

			} else {

				dict := apns.NewAlertDictionary()
				dict.Body = combined.Contents.Body

				// Including them gave unexpected results initially, but supported features by APNs.
				// dict.Title = combined.AlertDict.Title
				// dict.TitleLocKey = combined.AlertDict.TitleLocKey
				// dict.TitleLocArgs = combined.AlertDict.TitleLocArgs

				dict.ActionLocKey = combined.AlertDict.ActionLocKey
				dict.LocKey = combined.AlertDict.LocKey
				dict.LocArgs = combined.AlertDict.LocArgs
				dict.LaunchImage = combined.AlertDict.LaunchImage

				payload := apns.NewPayload()
				payload.Alert = dict
				payload.Badge = combined.Badge
				payload.Sound = combined.Sound
				payload.ContentAvailable = combined.ContentAvailable
				payload.Category = combined.Category

				pn := apns.NewPushNotification()
				pn.DeviceToken = val.DeviceID
				pn.AddPayload(payload)

				client := apns.NewClient(muto_gcm.ApnsAPI, "/Users/shivam_mac/Downloads/cert.pem", "/Users/shivam_mac/Downloads/key.pem")
				resp := client.Send(pn)
				fmt.Println("Response from the APNS Server : ", resp)

				alert, _ := pn.PayloadString()
				fmt.Println(alert)

			}

		}
	})
	fmt.Println("Going live on :9011")
	r.Run(":9011")
}
