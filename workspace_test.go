package toggl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func getTestWorkspace() Workspace {
	return Workspace{
		ID:                          3278506,
		Name:                        "Home",
		Premium:                     false,
		Admin:                       true,
		DefaultHourlyRate:           0,
		DefaultCurrency:             "USD",
		OnlyAdminsMayCreateProjects: false,
		OnlyAdminsSeeBillableRates:  false,
		Rounding:                    1,
		RoundingMinutes:             0,
		At:                          time.Date(2018, 4, 12, 7, 49, 15, 0, time.Local),
		LogoURL:                     "",
	}
}

func getTestWorkspaceUser() WorkspaceUser {
	return WorkspaceUser{
		ID:             4808871,
		UID:            2941647,
		Wid:            3278506,
		Admin:          true,
		Owner:          true,
		Active:         true,
		Email:          "test@a.a",
		Timezone:       "Asia/Tokyo",
		Inactive:       false,
		At:             time.Date(2018, 4, 12, 7, 49, 0, 0, time.Local),
		Name:           "saso",
		GroupIds:       []int{},
		Rate:           0,
		LabourCost:     0,
		InviteURL:      "",
		InvitationCode: "",
		AvatarFileName: "",
	}
}

func TestGetWrokspaces(t *testing.T) {
	expected := []Workspace{getTestWorkspace(), getTestWorkspace()}
	expectedURL := "/api/v8/workspaces"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal([]Workspace{
			getTestWorkspace(), getTestWorkspace(),
		})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil

	})
	api := New("test", OptionHTTPClient(client))

	workspaces, err := api.GetWrokspaces(context.Background())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *workspaces) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestGetWorkspaceProjects(t *testing.T) {
	expected := []Project{getTestProject(), getTestProject()}
	expectedURL := "/api/v8/workspaces/1/projects"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal([]Project{
			getTestProject(), getTestProject(),
		})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil

	})
	api := New("test", OptionHTTPClient(client))

	projects, err := api.GetWorkspaceProjects(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *projects) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestGetWorkspaceTags(t *testing.T) {
	expected := []Tag{getTestTag(), getTestTag()}
	expectedURL := "/api/v8/workspaces/1/tags"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal([]Tag{
			getTestTag(), getTestTag(),
		})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil

	})
	api := New("test", OptionHTTPClient(client))

	tags, err := api.GetWorkspaceTags(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *tags) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestGetWorkspaceUsers(t *testing.T) {
	expected := []WorkspaceUser{getTestWorkspaceUser()}
	expectedURL := "/api/v8/workspaces/1/workspace_users"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal([]WorkspaceUser{getTestWorkspaceUser()})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil

	})
	api := New("test", OptionHTTPClient(client))

	users, err := api.GetWorkspaceUsers(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *users) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}
