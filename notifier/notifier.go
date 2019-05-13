package notifier

type Notifier interface {
	Notify(string) error
}
