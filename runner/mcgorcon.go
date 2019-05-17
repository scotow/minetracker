package runner

import (
	"sync"

	"github.com/bearbin/mcgorcon"
)

func NewDirect(cred Credentials) (*Direct, error) {
	d := new(Direct)
	d.cred = cred

	client, err := mcgorcon.Dial(cred.Hostname, cred.Port, cred.Password)
	if err != nil {
		return nil, err
	}
	d.client = client

	return d, nil
}

type Direct struct {
	cred   Credentials
	client mcgorcon.Client
	lock   sync.Mutex
}

func (d *Direct) Run(cmd string) (string, error) {
	d.lock.Lock()
	defer d.lock.Unlock()

	return d.client.SendCommand(cmd)
}
