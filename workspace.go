package toggl

import (
	"context"
	"fmt"
	"time"
)

// Workspace contain all the information of a workspace
type Workspace struct {
	ID                          int       `json:"id"`
	Name                        string    `json:"name"`
	Premium                     bool      `json:"premium"`
	Admin                       bool      `json:"admin"`
	DefaultHourlyRate           int       `json:"default_hourly_rate"`
	DefaultCurrency             string    `json:"default_currency"`
	OnlyAdminsMayCreateProjects bool      `json:"only_admins_may_create_projects"`
	OnlyAdminsSeeBillableRates  bool      `json:"only_admins_see_billable_rates"`
	Rounding                    int       `json:"rounding"`
	RoundingMinutes             int       `json:"rounding_minutes"`
	At                          time.Time `json:"at"`
	LogoURL                     string    `json:"logo_url"`
}

// WorkspaceUser is not the User struct. WorkspaceUser connect between user and workspaces.
type WorkspaceUser struct {
	ID             int       `json:"id"`
	UID            int       `json:"uid"`
	Wid            int       `json:"wid"`
	Admin          bool      `json:"admin"`
	Owner          bool      `json:"owner"`
	Active         bool      `json:"active"`
	Email          string    `json:"email"`
	Timezone       string    `json:"timezone"`
	Inactive       bool      `json:"inactive"`
	At             time.Time `json:"at"`
	Name           string    `json:"name"`
	GroupIds       []int     `json:"group_ids"`
	Rate           int       `json:"rate"`
	LabourCost     int       `json:"labour_cost"`
	InviteURL      string    `json:"invite_url"`
	InvitationCode string    `json:"invitation_code"`
	AvatarFileName string    `json:"avatar_file_name"`
}

// GetWrokspaces will retrive the all workspace of the token onwner.
func (c *Client) GetWrokspaces(ctx context.Context) (*[]Workspace, error) {
	spath := "v8/workspaces"
	response := &[]Workspace{}

	err := c.get(ctx, spath, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetWorkspaceProjects will retrive the all active workspace projects.
func (c *Client) GetWorkspaceProjects(ctx context.Context, id int) (*[]Project, error) {
	spath := fmt.Sprintf("v8/workspaces/%d/projects", id)
	response := &[]Project{}

	err := c.get(ctx, spath, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetWorkspaceTags will retrive the all workspace tags.
func (c *Client) GetWorkspaceTags(ctx context.Context, id int) (*[]Tag, error) {
	spath := fmt.Sprintf("v8/workspaces/%d/tags", id)
	response := &[]Tag{}

	err := c.get(ctx, spath, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetWorkspaceUsers will retrive the all workspace tags.
func (c *Client) GetWorkspaceUsers(ctx context.Context, id int) (*[]WorkspaceUser, error) {
	spath := fmt.Sprintf("v8/workspaces/%d/workspace_users", id)
	response := &[]WorkspaceUser{}

	err := c.get(ctx, spath, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
