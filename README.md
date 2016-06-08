# What are Push Notifications?
Push notifications are a medium by which your application notifies a user of new messages or events even when the user is not actively using your application. Various platforms such as Android, iOS, Windows, etc. support this feature. Google Cloud Messaging(GCM) and Apple Push Notifications Service(APNS) are the APIs developed by the corresponding companies to allow your application to send Push Notifications to Android and iOS devices respectively.
Infact, even browsers such as Safari, Chrome and Mozilla amongst others, have support for Push Notifications.

# Why Push-Notification-Server?
Considering the existence of so many differing platforms, you'd have to write code for each of the platforms you intend to support. Wouldn't it be wonderful if you could just plug in the target device, the contents of the notification and the rest figures out by itself? This service does exactly that!
The Push-Notification-Server forwards your Push Notification to the API service of the intended target devices. You only need to give the Target Devices and the Notification Payload and the server takes care of the rest.
