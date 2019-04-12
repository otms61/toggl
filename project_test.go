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

func getTestProject() Project {
	return Project{
		ID:            111358164,
		Wid:           2108335,
		Name:          "test project",
		Billable:      false,
		IsPrivate:     true,
		Active:        true,
		Template:      false,
		At:            time.Date(2018, 4, 12, 7, 49, 15, 0, time.UTC),
		CreatedAt:     time.Date(2018, 4, 12, 7, 49, 15, 0, time.UTC),
		Color:         "3",
		AutoEstimates: false,
		ActualHours:   314,
		HexColor:      "#fb8b14",
	}
}

func getTestUser() User {
	return User{
		ID:      56651129,
		Pid:     111358164,
		UID:     2941647,
		Wid:     2108335,
		Manager: true,
		Rate:    0,
	}
}

func TestGetProject(t *testing.T) {
	expectedURL := "/api/v8/projects/1"
	expected := getTestProject()

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data Project `json:"data"`
		}{
			Data: getTestProject(),
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

	project, err := api.GetProject(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *project) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestCreateProject(t *testing.T) {
	expectedURL := "/api/v8/projects"
	expected := getTestProject()

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data Project `json:"data"`
		}{
			Data: getTestProject(),
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

	project, err := api.CreateProject(context.Background(), "test project", 2108335, 9999, false, 8888)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *project) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestUpdateProject(t *testing.T) {
	expectedURL := "/api/v8/projects/111358164"
	expected := getTestProject()

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data Project `json:"data"`
		}{
			Data: getTestProject(),
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

	project, err := api.UpdateProject(context.Background(), 111358164, "test project", 2108335, 9999, false, 8888)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *project) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestDeleteProject(t *testing.T) {
	expectedURL := "/api/v8/projects/111358164"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal([]int{111358164})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil

	})
	api := New("test", OptionHTTPClient(client))

	err := api.DeleteProject(context.Background(), 111358164)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestGetProjectUsers(t *testing.T) {
	expectedURL := "/api/v8/projects/1/project_users"
	expected := []User{getTestUser()}

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal([]User{getTestUser()})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil

	})
	api := New("test", OptionHTTPClient(client))

	user, err := api.GetProjectUsers(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *user) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}
