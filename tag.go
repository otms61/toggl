package toggl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// Tag contain all the information of a tag
type Tag struct {
	ID   int    `json:"id,omitempty"`
	Wid  int    `json:"wid,omitempty"`
	Name string `json:"name"`
}

type tagResponse struct {
	Data Tag `json:"data"`
}

// CreateTag creates a new tag based in the given cofiguration.
func (c *Client) CreateTag(ctx context.Context, name string, workspaceID int) (*Tag, error) {
	spath := "v8/tags"
	response := &tagResponse{}

	params := struct {
		Tag Tag `json:"tag"`
	}{
		Tag: Tag{
			Wid:  workspaceID,
			Name: name,
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

// UpdateTag creates a new tag based in the given cofiguration.
func (c *Client) UpdateTag(ctx context.Context, id int, name string) (*Tag, error) {
	spath := fmt.Sprintf("v8/tags/%d", id)
	response := &tagResponse{}

	params := struct {
		Tag Tag `json:"tag"`
	}{
		Tag: Tag{
			Name: name,
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

// DeleteTag will delete the tag.
func (c *Client) DeleteTag(ctx context.Context, id int) error {
	spath := fmt.Sprintf("v8/tags/%d", id)
	response := &[]int{}

	err := c.delete(ctx, spath, nil, response)
	if err != nil {
		return err
	}

	return nil
}
