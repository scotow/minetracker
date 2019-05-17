package runner

import (
	"sync"

	"github.com/bearbin/mcgorcon"
)

func NewDirectRunner(cred Credentials) (*DirectRunner, error) {
	dr := new(DirectRunner)
	dr.cred = cred

	client, err := mcgorcon.Dial(cred.Hostname, cred.Port, cred.Password)
	if err != nil {
		return nil, err
	}
	dr.client = client

	return dr, nil
}

type DirectRunner struct {
	cred   Credentials
	client mcgorcon.Client
	lock   sync.Mutex
}

func (dr *DirectRunner) Run(cmd string) (string, error) {
	dr.lock.Lock()
	defer dr.lock.Unlock()

	return dr.client.SendCommand(cmd)
}
