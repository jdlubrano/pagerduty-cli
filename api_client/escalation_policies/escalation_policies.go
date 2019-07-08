package escalation_policies

import (
  "github.com/jdlubrano/pagerduty-cli/api_client"
)

type EscalationPoliciesData struct {
  EscalationPolicies []EscalationPolicy `json:"escalation_policies"`
}

type EscalationPolicy struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

func ListForQuery(client *api_client.ApiClient, query string) ([]EscalationPolicy, error) {
  params := &map[string]string{"query": query}

  resp, err := client.Get("/escalation_policies", params)

  if err != nil {
    return []EscalationPolicy{}, err
  }

  var data EscalationPoliciesData
  resp.ParseInto(&data)
  return data.EscalationPolicies, nil
}
