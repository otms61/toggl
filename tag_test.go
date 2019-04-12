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
)

func getTestTag() Tag {
	return Tag{
		ID:   5740596,
		Wid:  3278506,
		Name: "fun",
	}
}

func TestCreateTag(t *testing.T) {
	expected := getTestTag()
	expectedURL := "/api/v8/tags"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(tagResponse{
			Data: getTestTag(),
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

	tag, err := api.CreateTag(context.Background(), "fun", 3278506)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *tag) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestUpdateTag(t *testing.T) {
	expected := getTestTag()
	expectedURL := "/api/v8/tags/5740596"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		b, err := json.Marshal(tagResponse{
			Data: getTestTag(),
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

	tag, err := api.UpdateTag(context.Background(), 5740596, "fun")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, *tag) {
		t.Fatal(errors.New("Response is incorrect"))
	}
}

func TestDeleteTag(t *testing.T) {
	expectedURL := "/api/v8/tags/5740596"

	client := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("Expected URL '%s', got %s", expectedURL, req.URL.Path)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("null"))),
		}, nil

	})
	api := New("test", OptionHTTPClient(client))

	err := api.DeleteTag(context.Background(), 5740596)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}
