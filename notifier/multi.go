package notifier

import "sync"

func NewMultiNotifier(notifiers ...Notifier) *MultiNotifier {
	mn := new(MultiNotifier)
	mn.notifiers = notifiers
	return mn
}

type MultiNotifier struct {
	notifiers []Notifier
}

func (mn *MultiNotifier) Notify(data string) error {
	for _, n := range mn.notifiers {
		err := n.Notify(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDynamicMultiNotifier(notifiers ...Notifier) *DynamicMultiNotifier {
	dmn := new(DynamicMultiNotifier)
	dmn.notifiers = notifiers
	return dmn
}

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

func (dmn *DynamicMultiNotifier) Add(notifier Notifier) {
	dmn.lock.Lock()
	dmn.notifiers = append(dmn.notifiers, notifier)
	dmn.lock.Unlock()
}
