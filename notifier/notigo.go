package notifier

import "github.com/scotow/notigo"

func NewNotigoNotifier(key, title string) *NotigoNotifier {
	nn := new(NotigoNotifier)
	nn.key = notigo.Key(key)
	nn.title = title
	return nn
}

type NotigoNotifier struct {
	key   notigo.Key
	title string
}

func (nn *NotigoNotifier) Notify(data string) error {
	return nn.key.Send(notigo.NewNotification(nn.title, data))
}
