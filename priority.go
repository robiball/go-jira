package jira

import "fmt"

// PriorityService handles projects for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project
type PriorityService struct {
	client *Client
}

// Priority represents a priority of a JIRA issue.
// Typical types are "Normal", "Moderate", "Urgent", ...
type Priority struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	StatusColor string `json:"statusColor,omitempty" structs:"statusColor,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
}

// PriorityScheme represents a priority scheme
type PriorityScheme struct {
	Expand          string    `json:"expand" structs:"expand,omitempty"`
	Self            string    `json:"self" structs:"self,omitempty"`
	ID              int       `json:"id" structs:"id,omitempty"`
	Name            string    `json:"name" structs:"name,omitempty"`
	Description     string    `json:"Description" structs:"description,omitempty"`
	DefaultOptionId string    `json:"defaultOptionId" structs:"defaultOptionId,omitempty"`
	OptionIds       []string  `json:"optionIds" structs:"optionIds,omitempty"`
	DefaultScheme   bool      `json:"defaultScheme" structs:"defaultScheme,omitempty"`
	ProjectKeys     []string  `json:"projectKeys" structs:"projectKeys,omitempty"`
}

// Get returns a full representation of a priority
// JIRA will attempt to identify the priority by the priorityID parameter.
//
// JIRA API docs: https://docs.atlassian.com/software/jira/docs/api/REST/7.7.0/#api/2/priority
func (s *PriorityService) Get(priorityID string) (*Priority, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/priority/%s", priorityID)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	p := new(Priority)
	resp, err := s.client.Do(req, p)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return p, resp, nil
}


// GetPriorityScheme returns a full representation of a priority scheme
// JIRA will attempt to identify the scheme by the prioritySchemeID parameter.
//
// JIRA API docs: https://docs.atlassian.com/software/jira/docs/api/REST/7.7.0/#api/2/priorityschemes
func (s *PriorityService) GetScheme(prioritySchemeID string) (*PriorityScheme, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/priorityschemes/%s", prioritySchemeID)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	ps := new(PriorityScheme)
	resp, err := s.client.Do(req, ps)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return ps, resp, nil
}
