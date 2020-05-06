package v1

type URLS struct {
	URL            string `json:"url,omitempty"`
	URLProvision   string
	URLDeprovision string
	URLUpdate      string
	URLinstance    string
}

type Broker struct {
	Head
	ID
	URLBroker
	BodyBroker
}

type Head struct {
	AuthKEY   string `json:"authkey,omitempty"`
	BrokerID  string `json:"brokerIdentity,omitempty"`
	BrokerVER string `json:"brokerversion,omitempty"`
	UID       string
}

type BodyBroker struct {
	Parameters
	Context
}

type URLBroker struct {
	Server            string
	Provision         string
	Instancia         string
	AcceptsIncomplete string `json:"accepts_incomplete,omitempty"`
	URLS
}
