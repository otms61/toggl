package toggl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// TimeEntry contain all the information of a time entries
type TimeEntry struct {
	ID          int       `json:"id"`
	GUID        string    `json:"guid"`
	Wid         int       `json:"wid"`
	Pid         int       `json:"pid"`
	Description string    `json:"description"`
	Billable    bool      `json:"billable"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
	Duration    int       `json:"duration"`
	Duronly     bool      `json:"duronly"`
	At          time.Time `json:"at"`
	UID         int       `json:"uid"`
	Tags        []string  `json:"tags"`
	Project     Project
}

// TimeEntryRequest is used to create time entry
type TimeEntryRequest struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Start       string   `json:"start,omitempty"`
	Duration    int      `json:"duration,omitempty"`
	Pid         int      `json:"pid"`
	CreatedWith string   `json:"created_with"`
}

// TimeEntryResponse is the wrapper of the TimeEntry for API server specification.
type TimeEntryResponse struct {
	Data TimeEntry `json:"data"`
}

// BulkUpdateTagsRequest is the wrapper of the TimeEntry for API server specification.
type BulkUpdateTagsRequest struct {
	Tags      []string `json:"tags"`
	TagAction string   `json:"tag_action"`
}

// BulkUpdateTagsResponse is the wrapper of the TimeEntry for API server specification.
type BulkUpdateTagsResponse struct {
	Data []TimeEntry `json:"data"`
}

// GetRunningTimeEntry will retrive running time entry.
func (c *Client) GetRunningTimeEntry(ctx context.Context) (*TimeEntry, error) {
	spath := "v8/time_entries/current"
	response := &TimeEntryResponse{}

	err := c.get(ctx, spath, nil, response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetTimeEntries will retrive time entries in specific time range.
func (c *Client) GetTimeEntries(ctx context.Context, start, end time.Time) (*[]TimeEntry, error) {
	spath := fmt.Sprintf("v8/time_entries")
	response := &[]TimeEntry{}

	params := url.Values{}
	params.Add("start_date", start.Format("2006-01-02T15:04:05+09:00"))
	params.Add("end_date", end.Format("2006-01-02T15:04:05+09:00"))

	err := c.get(ctx, spath, params, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetTimeEntry will retrive time entries in specific time range.
func (c *Client) GetTimeEntry(ctx context.Context, id int) (*TimeEntry, error) {
	spath := fmt.Sprintf("v8/time_entries/%d", id)
	response := &TimeEntryResponse{}

	err := c.get(ctx, spath, nil, response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// StartTimeEntry creates a new running time entry based in the given configuration.
func (c *Client) StartTimeEntry(ctx context.Context, projectID int, description string, tags []string, createdWith string) (*TimeEntry, error) {
	spath := "v8/time_entries/start"
	response := &TimeEntryResponse{}

	params := struct {
		TimeEntry TimeEntryRequest `json:"time_entry"`
	}{
		TimeEntry: TimeEntryRequest{
			Description: description,
			Tags:        tags,
			Pid:         projectID,
			CreatedWith: createdWith,
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

// StopTimeEntry will stop the running time entry.
func (c *Client) StopTimeEntry(ctx context.Context, id int) error {
	spath := fmt.Sprintf("v8/time_entries/%d/stop", id)
	response := &TimeEntryResponse{}

	err := c.put(ctx, spath, nil, response)
	if err != nil {
		return err
	}

	return nil
}

// CreateTimeEntry creates a new time entry based in the given configuration.
func (c *Client) CreateTimeEntry(ctx context.Context, projectID int, description string, start time.Time, duration int, tags []string, createdWith string) (*TimeEntry, error) {
	spath := "v8/time_entries"
	response := &TimeEntryResponse{}

	params := struct {
		TimeEntry TimeEntryRequest `json:"time_entry"`
	}{
		TimeEntry: TimeEntryRequest{
			Description: description,
			Tags:        tags,
			Start:       start.Format("2006-01-02T15:04:05+09:00"),
			Duration:    duration,
			Pid:         projectID,
			CreatedWith: createdWith,
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

// UpdateTimeEntry creates a new project based in the given configuration.
func (c *Client) UpdateTimeEntry(ctx context.Context, timeEntryID int, projectID int, description string, start time.Time, duration int, tags []string, createdWith string) (*TimeEntry, error) {
	spath := fmt.Sprintf("v8/time_entries/%d", timeEntryID)
	response := &TimeEntryResponse{}

	params := struct {
		TimeEntry TimeEntryRequest `json:"time_entry"`
	}{
		TimeEntry: TimeEntryRequest{
			Description: description,
			Tags:        tags,
			Start:       start.Format("2006-01-02T15:04:05+09:00"),
			Duration:    duration,
			Pid:         projectID,
			CreatedWith: createdWith,
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

// DeleteTimeEntry will delete the time entry.
func (c *Client) DeleteTimeEntry(ctx context.Context, id int) error {
	spath := fmt.Sprintf("v8/time_entries/%d", id)
	response := &[]int{}

	err := c.delete(ctx, spath, nil, response)
	if err != nil {
		return err
	}

	return nil
}

// BulkUpdateTimeEntriesTags creates a new project based in the given configuration.
func (c *Client) BulkUpdateTimeEntriesTags(ctx context.Context, timeEntryIDs []int, tags []string, action string) (*[]TimeEntry, error) {
	// BulkUpdate endpoint needs ID separated by ",".
	var ids bytes.Buffer
	ids.WriteString(fmt.Sprint(timeEntryIDs[0]))
	timeEntryIDs = timeEntryIDs[1:]
	for _, ID := range timeEntryIDs {
		ids.WriteString(",")
		s := fmt.Sprint(ID)
		ids.WriteString(s)
	}

	spath := fmt.Sprintf("v8/time_entries/%s", ids.String())
	response := &BulkUpdateTagsResponse{}

	params := struct {
		TimeEntry BulkUpdateTagsRequest `json:"time_entry"`
	}{
		TimeEntry: BulkUpdateTagsRequest{
			Tags:      tags,
			TagAction: action,
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
