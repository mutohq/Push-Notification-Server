# What are Push Notifications?
Push notifications are a medium by which your application notifies a user of new messages or events even when the user is not actively using your application. Various platforms such as Android, iOS, Windows, etc. support this feature. Google Cloud Messaging(GCM) and Apple Push Notifications Service(APNS) are the APIs developed by the corresponding companies to allow your application to send Push Notifications to Android and iOS devices respectively.
Infact, even browsers such as Safari, Chrome and Mozilla amongst others, have support for Push Notifications.

# Why Push-Notification-Server?
Considering the existence of so many differing platforms, you'd have to write code for each of the platforms you intend to support. Wouldn't it be wonderful if you could just plug in the target device, the contents of the notification and the rest figures out by itself? This service does exactly that!
The Push-Notification-Server forwards your Push Notification to the API service of the intended target devices. You only need to give the Target Devices and the Notification Payload and the server takes care of the rest.

# Installation
go get github.com/mutohq/Push-Notification-Server

#Usage
 **Device ID filled in manually**

```
package main

import (
      "fmt"
      "github.com/mutohq/Push-Notification-Server"
      "encoding/json"
      "github.com/parnurzeal/gorequest"
)
  
func main(){
  var c muto_gcm.Combined 
  c.Contents.Title = "Notification Title"
  c.Contents.Body = "NOtification Body"
  c.TargetDeviceIDs = []string{"device_id_1", "Some_more_device_ids"}
  b, _ := json.Marshal(request)
	req := gorequest.New()
	resp, body, errs := req.Post("http://localhost:9010/send").
		Set("Content-Type", "application/json").
		Send(string(b)).
		End()
	fmt.Println("\n", resp, body, errs)
}
```
# The Combined Struct
Probably the axle of the server, this struct combines the payload structure fields used by GCM and APNS **locally**. The original format supported by the APIs is mapped later on when the Platform is identified.
```
// A Local struct to integrate all the supported fields of both APIs before bifurcating the payload based on Platform.

type Combined struct {
	
	// GCM fields
	
	TargetDeviceIDs       []string    `json:"registration_ids,omitempty"`
	CollapseKey           string      `json:"collapse_key,omitempty"`
	Priority              string      `json:"priority,omitempty"`
	ContentAvailableGcm   bool        `json:"content_available,omitempty"`
	DelayWhileIdle        bool        `json:"delay_while_idle,omitempty"`
	TimeToLive            int         `json:"time_to_live,omitempty"`
	RestrictedPackageName string      `json:"restricted_package_name,omitempty"`
	DryRun                bool        `json:"dry_run,omitempty"`
	Payload               PayloadBody `json:"notification,omitempty"` ////  notification
	Contents              Content     `json:"contents,omitempty"`
	
	 // APNS fields

	Alert            interface{} `json:"alert,omitempty"` ////    alert
	Badge            int         `json:"badge,omitempty"`
	Sound            string      `json:"sound,omitempty"`
	ContentAvailable int         `json:"content-available,omitempty"`
	Category         string      `json:"category,omitempty"`
	
	// An array to store the DeviceIDs and their Platforms

	DevicesList []Device `json"deviceslist,omitempty"`
	
	// Not a standard field. Included to incorporate AlertDictionary type in Combined type.

	AlertDict AlertDictionary `json:"alertdic,omitempty"`
}

```
