package v1

func (P *Provision) setParameters(h Broker) {
	P.Parameters = Parameters{
		ExVarZone:  h.ExVarZone,
		Hostname:   h.URL,
		Tipo:       h.Tipo,
		TTL:        h.TTL,
		StrRecords: h.StrRecords,
		Descricao:  h.Descricao,
		Reverso:    true,
	}
}
