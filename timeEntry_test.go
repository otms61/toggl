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

func getTestTimeEntry() TimeEntry {
	return TimeEntry{
		ID:          1111111111,
		GUID:        "31424791f57b72c9cbbd80ed0f85790a",
		Wid:         3278506,
		Pid:         123456789,
		Description: "toggl test",
		Billable:    false,
		Start:       time.Date(2018, 4, 12, 7, 49, 0, 0, time.Local),
		Duration:    30,
		Duronly:     false,
		At:          time.Date(2018, 4, 12, 7, 49, 30, 0, time.Local),
		UID:         2941647,
		Tags: []string{
			"fun",
		},
	}
}

func TestRunningTimeEntry(t *testing.T) {
	expected := getTestTimeEntry()
	expectedURL := "/api/v8/time_entries/current"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data TimeEntry `json:"data"`
		}{
			Data: getTestTimeEntry(),
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

	timeEntry, err := api.GetRunningTimeEntry(context.Background())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *timeEntry) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestGetTimeEntry(t *testing.T) {
	expected := getTestTimeEntry()
	expectedURL := "/api/v8/time_entries/1111111111"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data TimeEntry `json:"data"`
		}{
			Data: getTestTimeEntry(),
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

	timeEntry, err := api.GetTimeEntry(context.Background(), 1111111111)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *timeEntry) {
		fmt.Println(expected)
		fmt.Printf("val:%#v type:%T\n", expected.Wid, expected.Wid)
		fmt.Println(*timeEntry)
		fmt.Printf("val:%#v type:%T\n", (*timeEntry).Wid, (*timeEntry).Wid)
		t.Fatal(errors.New("Response is incorrect"))
	}

}

func TestCreateTimeEntry(t *testing.T) {
	expected := getTestTimeEntry()
	expectedURL := "/api/v8/time_entries"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data TimeEntry `json:"data"`
		}{
			Data: getTestTimeEntry(),
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

	start := time.Date(2018, 4, 12, 7, 49, 0, 0, time.Local)
	timeEntry, err := api.CreateTimeEntry(context.Background(), 123456789, "toggl test", start, 30, []string{"fun"}, "golang")

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *timeEntry) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestUpdateTimeEntry(t *testing.T) {
	expected := getTestTimeEntry()
	expectedURL := "/api/v8/time_entries/1111111111"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data TimeEntry `json:"data"`
		}{
			Data: getTestTimeEntry(),
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

	start := time.Date(2018, 4, 12, 7, 49, 0, 0, time.Local)
	timeEntry, err := api.UpdateTimeEntry(context.Background(), 1111111111, 123456789, "toggl test", start, 30, []string{"fun"}, "golang")

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *timeEntry) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestDeleteTimeEntry(t *testing.T) {
	expectedURL := "/api/v8/time_entries/1111111111"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal([]int{1111111111})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil

	})
	api := New("test", OptionHTTPClient(client))

	err := api.DeleteTimeEntry(context.Background(), 1111111111)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestStartTimeEntry(t *testing.T) {
	expected := getTestTimeEntry()
	expectedURL := "/api/v8/time_entries/start"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data TimeEntry `json:"data"`
		}{
			Data: getTestTimeEntry(),
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

	timeEntry, err := api.StartTimeEntry(context.Background(), 123456789, "toggl test", []string{"fun"}, "golang")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *timeEntry) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestStopTimeEntry(t *testing.T) {
	expectedURL := "/api/v8/time_entries/1111111111/stop"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(struct {
			Data TimeEntry `json:"data"`
		}{
			Data: getTestTimeEntry(),
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

	err := api.StopTimeEntry(context.Background(), 1111111111)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}
