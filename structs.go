package muto_gcm

const GcmAPI = "https://android.googleapis.com/gcm/send"
const Authorization = "key=AIzaSyBfI5t4-GW5VovfzQ6BpvhTd2dkUB7L9R0"

const ApnsAPI = "gateway.sandbox.push.apple.com:2195"


// Stores DevviceID/Token and Platform for any given device
type Device struct {
	DeviceID string `json:"deviceid,omitempty"`
	Platform string `json:"platform,omitempty"`
}

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

// Original GCM struct format
type Notification struct {
	TargetDeviceIDs       []string    `json:"registration_ids,omitempty"`
	Payload               PayloadBody `json:"notification,omitempty"`
	CollapseKey           string      `json:"collapse_key,omitempty"`
	Priority              string      `json:"priority,omitempty"`
	ContentAvailable      bool        `json:"content_available,omitempty"`
	DelayWhileIdle        bool        `json:"delay_while_idle,omitempty"`
	TimeToLive            int         `json:"time_to_live,omitempty"`
	RestrictedPackageName string      `json:"restricted_package_name,omitempty"`
	DryRun                bool        `json:"dry_run,omitempty"`
}

// GCM Payload fields
type PayloadBody struct {
	Title        string   `json:"title,omitempty"`
	Body         string   `json:"body,omitempty"`
	Icon         string   `json:"icon,omitempty"`
	Sound        string   `json:"sound,omitempty"`
	Tag          string   `json:"tag,omitempty"`
	Color        string   `json:"color,omitempty"`
	ClickAction  string   `json:"click_action,omitempty"`
	BodyLocKey   string   `json:"body_loc_key,omitempty"`
	BodyLocArgs  []string `json:"body_loc_args,omitempty"`
	TitleLocKey  string   `json:"title_loc_key,omitempty"`
	TitleLocArgs []string `json:"title_loc_args,omitempty"`
}

// APNS original struct format
type AlertDictionary struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`

	TitleLocKey  string   `json:"title-loc-key,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
}

// Captures the common features of GCM and APNS payload
type Content struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

// Reuired to query the Database and find out DeviceID and Platform.
type Request struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
