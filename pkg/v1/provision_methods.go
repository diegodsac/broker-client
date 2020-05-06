package v1

import (
	"encoding/json"
)

// Update Update
func (P *Provision) Update() {

}

// String conversao
func (P *Provision) String() (string, error) {
	b, err := json.Marshal(P)
	return string(b), err
}

//SetValues SetValues
func (P *Provision) SetValues(h Broker) {
	P.id.ID = ID{
		OrganizationGUID: h.OrganizationGUID, SpaceGUID: h.SpaceGUID,
		ServiceID: h.ServiceID,
		PlanID:    h.PlanID,
	}
	P.setParameters(h)
	P.Context = Context{
		ProjectID: h.Hostname,
		ClusterID: h.ClusterID,
		Namespace: h.Namespace,
		Platform:  h.Platform,
	}
	P.broker = h
}

//Create teste
func (P *Provision) Create() {
	var tmp Instances
	var tmpBroker Broker
	tmpBroker = P.broker

	tmpBroker.URLBroker = URLBroker{
		Instancia: "?context/user_id=k8s-bbcloud&context.project_id=",
		Server:    "http://oaas-broker.nuvem.bb.com.br",
	}
	tmp.GetInstanceList(tmpBroker)
	if len(tmp.Items) > 0 {
		if tmp.Items[0].State == "PROVISIONED" {
			return
		}
	}
	P.broker.SetURLCreate()
	body, err := create(*P, P.broker)
	err = json.Unmarshal(body, &P.Return)
	if err != nil {
		panic(err)
	}

}
