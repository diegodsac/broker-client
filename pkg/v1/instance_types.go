package v1

type parameters struct {
	Parameters
}

//Context informacoes do deploy
type Context struct {
	//ProjectID é a parte do Hostname da url
	ProjectID string `json:"project_id,omitempty"`
	//ClusterID uid de quem criou
	ClusterID string `json:"clusterid,omitempty"`
	//Namespace registro do servidor DNS do broker
	Namespace string `json:"namespace,omitempty"`
	//Platform a plataforma que requisitou a criação
	Platform string `json:"platform,omitempty"`
}

//Context informacoes do deploy
type context struct {
	Context
}

//Instance json do broker
type Instance struct {
	//InstanceID UID fornecido para criaçao
	InstanceID string `json:"instance_id,omitempty"`
	//Inputs valores de criação do dns
	Inputs parameters `json:"inputs,omitempty"`
	//Contexts conjunto de informações fornecido pelo requisitante
	Contexts context `json:"context,omitempty"`
	//DeprovisionOperation uid da operaçao
	DeprovisionOperation string `json:"deprovision_operation,omitempty"`
	//ProvisionOperation uid da operaçao
	ProvisionOperation string `json:"provision_operation,omitempty"`
	//State informa o estado atual de provisionamento
	State string `json:"state,omitempty"`
	// Outputs              string `json:"outputs,omitempty"`
	// Children string `json:"children,omitempty"`
}

// Instances contains a list of Instance
type Instances struct {
	//Items lista de Instance
	Items []Instance `json:"items"`
}

//Provision valores que são o payload
type Provision struct {
	id
	Parameters Parameters `json:"parameters,omitempty"`
	Context    Context    `json:"context,omitempty"`
	Return     Return
	// head       Head
	broker Broker
}

//Return retorno do broker
type Return struct {
	Operation   string `json:"operation,omitempty"`
	Description string `json:"description,omitempty"`
}

type id struct {
	ID
}

//ID chaves de identificação
type ID struct {
	ServiceID        string `json:"service_id,omitempty"`
	OrganizationGUID string `json:"organization_guid,omitempty"`
	SpaceGUID        string `json:"space_guid,omitempty"`
	PlanID           string `json:"plan_id,omitempty"`
}
