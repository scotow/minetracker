package notigo

import "github.com/scotow/notigo"

// Create Notigo Notifier.
// key is the secret key provided by IFTTT. title is the title of the Push Notification.
func NewNotigoNotifier(key, title string) *NotigoNotifier {
	nn := new(NotigoNotifier)
	nn.key = notigo.Key(key)
	nn.title = title
	return nn
}

// The Notigo Notifier use the scotow/notigo library to send IFTTT Push Notifications.
type NotigoNotifier struct {
	key   notigo.Key
	title string
}

func (nn *NotigoNotifier) Notify(data string) error {
	return nn.key.Send(notigo.NewNotification(nn.title, data))
}
