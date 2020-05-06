package v1

import (
	"fmt"
	"net/url"
)

// SetURLDelete deleta
func (H *Broker) SetURLDelete() {
	tmpURL, _ := url.Parse(H.Server)
	tmpURL.Path = H.Provision
	tmpURL.Path = H.UID
	params := url.Values{}
	params.Add("accepts_incomplete", "true")
	params.Add("plan_id", H.PlanID)
	params.Add("service_id", H.ServiceID)
	tmpURL.RawQuery = params.Encode()
	H.URL = tmpURL.String()
}

// SetURLCreate teste
func (H *Broker) SetURLCreate() {
	tmpURL, _ := url.Parse(H.Server)
	tmpURL.Path = H.Provision + H.UID
	params := url.Values{}
	params.Add("accepts_incomplete", "true")
	tmpURL.RawQuery = params.Encode()
	H.URLProvision = tmpURL.String()
	fmt.Printf("%s\n", H.URL)
}
