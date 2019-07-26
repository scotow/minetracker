package notifier

import "sync"

// Create a MultiNotifier using the specified notifiers.
func NewMultiNotifier(notifiers ...Notifier) *MultiNotifier {
	mn := new(MultiNotifier)
	mn.notifiers = notifiers
	return mn
}

// A simple utility Notifier that iterate over a list of Notifier and call their Notify function.
type MultiNotifier struct {
	notifiers []Notifier
}

// Iterate over a list of Notifier and call their Notify function.
func (mn *MultiNotifier) Notify(data string) error {
	for _, n := range mn.notifiers {
		err := n.Notify(data)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create a DynamicMultiNotifier with the specified notifiers.
func NewDynamicMultiNotifier(notifiers ...Notifier) *DynamicMultiNotifier {
	dmn := new(DynamicMultiNotifier)
	dmn.notifiers = notifiers
	return dmn
}

// A DynamicMultiNotifier is like a MultiNotifier but that allows the insertion of new Notifier on the fly.
type DynamicMultiNotifier struct {
	MultiNotifier
	lock sync.RWMutex
}

func (dmn *DynamicMultiNotifier) Notify(data string) error {
	dmn.lock.RLock()
	for _, n := range dmn.notifiers {
		err := n.Notify(data)
		if err != nil {
			dmn.lock.RUnlock()
			return err
		}
	}

	dmn.lock.RUnlock()
	return nil
}

// Add a notifier to the list.
// This method is thread-safe.
func (dmn *DynamicMultiNotifier) Add(notifier Notifier) {
	dmn.lock.Lock()
	dmn.notifiers = append(dmn.notifiers, notifier)
	dmn.lock.Unlock()
}
