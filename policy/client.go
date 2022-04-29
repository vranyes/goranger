package policy

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/vranyes/goranger/client"
)

type PolicyClient struct {
	Client      *client.RangerClient
	policy_path string
}

func NewPolicyClient(base_url, username, password string, ssl_disabled bool) PolicyClient {
	c := client.NewClient(base_url, username, password, ssl_disabled)

	return PolicyClient{
		Client:      &c,
		policy_path: "/policy",
	}
}

func (p PolicyClient) CreatePolicy(policy Policy) (Policy, error) {
	resource_path := "/"
	response, err := p.post(resource_path, policy)
	if err != nil {
		return Policy{}, err
	}

	err = json.Unmarshal(response, &policy)
	return policy, err
}

func (p PolicyClient) GetPolicyById(policyId int) (Policy, error) {
	resource_path := "/" + strconv.Itoa(policyId)
	response, err := p.get(resource_path)
	if err != nil {
		return Policy{}, err
	}

	var policy Policy
	err = json.Unmarshal(response, &policy)
	return policy, err
}

func (p PolicyClient) GetPolicyByResources(resource map[string]PolicyResource) (int, error) {
	policies, err := p.ReadAll()
	if err != nil {
		return -1, err
	}

	for _, p := range policies {
		if reflect.DeepEqual(resource, p.Resources) {
			return p.Id, nil
		}
	}
	return -1, errors.New("Couldn't find matching")
}

func (p PolicyClient) ReadAll() ([]Policy, error) {
	response, err := p.get("/")
	if err != nil {
		return []Policy{}, err
	}

	var policies []Policy
	err = json.Unmarshal(response, &policies)
	return policies, nil
}

func (p PolicyClient) UpdatePolicy(policy Policy) (Policy, error) {
	resource_path := "/" + strconv.Itoa(policy.Id)
	response, err := p.put(resource_path, policy)
	if err != nil {
		return Policy{}, err
	}

	err = json.Unmarshal(response, &policy)
	return policy, err
}

func (p PolicyClient) DeletePolicy(policy Policy) (bool, error) {
	resource_path := "/" + strconv.Itoa(policy.Id)
	_, err := p.put(resource_path, policy)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p PolicyClient) DeletePolicyById(policyId int) (bool, error) {
	resource_path := "/" + strconv.Itoa(policyId)
	_, err := p.delete(resource_path)
	if err != nil {
		return false, err
	}
	return true, nil
}

// HTTP verbs with the policy path applied
func (p PolicyClient) get(path string) ([]byte, error) {
	request_path := p.policy_path + path

	resp, err := p.Client.RequestHandler("GET", request_path, http.NoBody)
	return resp, err
}

func (p PolicyClient) post(path string, policy Policy) ([]byte, error) {
	request_path := p.policy_path + path
	body, err := json.Marshal(policy)

	resp, err := p.Client.RequestHandler("POST", request_path, bytes.NewBuffer(body))
	return resp, err
}

func (p PolicyClient) put(path string, policy Policy) ([]byte, error) {
	request_path := p.policy_path + path
	body, err := json.Marshal(policy)

	resp, err := p.Client.RequestHandler("PUT", request_path, bytes.NewBuffer(body))
	return resp, err
}

func (p PolicyClient) delete(path string) ([]byte, error) {
	request_path := p.policy_path + path

	resp, err := p.Client.RequestHandler("DELETE", request_path, http.NoBody)
	return resp, err
}
