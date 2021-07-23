package notification

type Notifier interface {
	Notify(content []string) error
}
