package v1

// Parameters valores que criam o dns
type Parameters struct {
	//ExVarZone Zona onde foi criada
	ExVarZone string `json:"ex_var_zone,omitempty"`
	//Hostname registro básico do nome DNS
	Hostname string `json:"hostname,omitempty"`
	//Tipo trata da forma de acesso ao domínio.(A,CNAME...)
	Tipo string `json:"tipo,omitempty"`
	//TTL o tempo de cache do record DNS
	TTL string `json:"ttl,omitempty"`
	//StrRecords identifica os servidores DNS responsáveis ​​por um nome de domínio
	StrRecords string `json:"str_records,omitempty"`
	//Descricao breve descrição
	Descricao string `json:"descricao,omitempty"`
	//Reverso verifica a autenticidade de endereços
	Reverso bool `json:"reverso,omitempty"`
}

type DnsServiceSpec struct {
	Server      string `json:"server,omitempty"`
	Zone        string `json:"zone,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Type        string `json:"type,omitempty"`
	Ttl         int32  `json:"ttl,omitempty"`
	Records     string `json:"records,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

type DnsZoneSpec struct {
	AuthKey          string `json:"authkey,omitempty"`
	Server           string `json:"server,omitempty"`
	Zone             string `json:"zone,omitempty"`
	Records          string `json:"records,omitempty"`
	Type             string `json:"type,omitempty"`
	TTL              string `json:"ttl,omitempty"`
	Description      string `json:"description,omitempty"`
	BrokerVersion    string `json:"BrokerVersion,omitempty"`
	BrokerIdentity   string `json:"BrokerIdentity,omitempty"`
	Clusterid        string `json:"clusterid,omitempty"`
	Namespace        string `json:"namespace,omitempty"`
	Platform         string `json:"platform,omitempty"`
	Serviceid        string `json:"Serviceid,omitempty"`
	Planid           string `json:"Planid,omitempty"`
	Spaceguid        string `json:"spaceguid,omitempty"`
	Organizationguid string `json:"organizationguid,omitempty"`
	Url1             string `json:"Url1,omitempty"`
	Url2             string `json:"Url2,omitempty"`
}
