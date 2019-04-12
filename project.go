package toggl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Project contain all the information of a project.
type Project struct {
	ID            int       `json:"id"`
	Wid           int       `json:"wid"`
	Name          string    `json:"name"`
	Billable      bool      `json:"billable"`
	IsPrivate     bool      `json:"is_private"`
	Active        bool      `json:"active"`
	Template      bool      `json:"template"`
	At            time.Time `json:"at"`
	CreatedAt     time.Time `json:"created_at"`
	Color         string    `json:"color"`
	AutoEstimates bool      `json:"auto_estimates"`
	ActualHours   int       `json:"actual_hours"`
	HexColor      string    `json:"hex_color"`
}

type projectResponse struct {
	Data Project `json:"data"`
}

type projectRequest struct {
	Name       string `json:"name"`
	Wid        int    `json:"wid"`
	TemplateID int    `json:"template_id"`
	IsPrivate  bool   `json:"is_private"`
	Cid        int    `json:"cid"`
}

// GetProject will retrive the complete project information.
func (c *Client) GetProject(ctx context.Context, id int) (*Project, error) {
	spath := fmt.Sprintf("v8/projects/%d", id)
	response := &projectResponse{}

	err := c.get(ctx, spath, nil, response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// CreateProject creates a new project based in the given configuration with a custom context.
func (c *Client) CreateProject(ctx context.Context, name string, workspaceID int, templateID int, isPrivate bool, clientID int) (*Project, error) {
	spath := "v8/projects"
	response := &projectResponse{}

	params := struct {
		Project projectRequest `json:"project"`
	}{
		Project: projectRequest{
			Name:       name,
			Wid:        workspaceID,
			TemplateID: templateID,
			IsPrivate:  isPrivate,
			Cid:        clientID,
		},
	}
	j, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	err = c.post(ctx, spath, bytes.NewBuffer(j), response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// UpdateProject updates the project based in the given configuration.
func (c *Client) UpdateProject(ctx context.Context, projectID int, name string, workspaceID int, templateID int, isPrivate bool, clientID int) (*Project, error) {
	spath := fmt.Sprintf("v8/projects/%d", projectID)
	response := &projectResponse{}

	params := struct {
		Project projectRequest `json:"project"`
	}{
		Project: projectRequest{
			Name:       name,
			Wid:        workspaceID,
			TemplateID: templateID,
			IsPrivate:  isPrivate,
			Cid:        clientID,
		},
	}
	j, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	err = c.put(ctx, spath, bytes.NewBuffer(j), response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// DeleteProject will delete the project.
func (c *Client) DeleteProject(ctx context.Context, id int) error {
	spath := fmt.Sprintf("v8/projects/%d", id)
	response := &[]int{}

	err := c.delete(ctx, spath, nil, response)
	if err != nil {
		return err
	}

	return nil
}

// GetProjectUsers will retrive the complete project information.
func (c *Client) GetProjectUsers(ctx context.Context, id int) (*[]User, error) {
	spath := fmt.Sprintf("v8/projects/%d/project_users", id)
	response := &[]User{}

	err := c.get(ctx, spath, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
