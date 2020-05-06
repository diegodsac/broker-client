package v1

import (
	"encoding/json"
)

// GetInstance GetInstance
func (I *Instance) GetInstance(h Broker) {
	h.URLinstance = h.Server + "/data/instance" + h.UID
	body, err := getBroker(h)
	err = json.Unmarshal(body, &I)
	if err != nil {
		panic(err)
	}
}

//GetInstanceList GetInstanceList
func (I *Instances) GetInstanceList(b Broker) {
	b.URLinstance = b.Server + "/data/instance" + b.Instancia + b.Hostname
	body, err := getBroker(b)
	err = json.Unmarshal(body, &I.Items)
	if err != nil {
		panic(err)
	}
}
