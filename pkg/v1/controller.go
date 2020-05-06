package v1

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

func getBroker(B Broker) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", B.URLinstance, nil)
	if err != nil {
		return nil, nil
	}
	req.Header.Set("X-Instance-API-Version", "2.13")
	req.Header.Set("X-Instance-Api-Originating-Identity", B.BrokerID)
	req.Header.Set("Authorization", B.AuthKEY)
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	return body, nil
}

func create(p Provision, b Broker) ([]byte, error) {
	var payload *strings.Reader
	provString, _ := p.String()
	payload = strings.NewReader(provString)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("PUT", b.URLProvision, payload)
	req.Header.Set("X-Broker-API-Version", b.BrokerVER)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Broker-Api-Originating-Identity", b.BrokerID)
	req.Header.Set("Authorization", b.AuthKEY)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("cache-control", "no-cache")
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	return body, nil
}

func delete(p Provision, b Broker) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("DELETE", b.URLDeprovision, nil)
	if err != nil {
		return nil, nil
	}
	req.Header.Set("X-Broker-API-Version", b.BrokerVER)
	req.Header.Set("X-Broker-Api-Originating-Identity", b.BrokerID)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", b.AuthKEY)
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	return body, nil
}
