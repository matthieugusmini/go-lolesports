// Package lolesports provides a client for interacting with the unofficial LoL Esports API.
//
// Note: This client relies on an unofficial API that could change or become unavailable at any time
// if Riot Games modifies or discontinues their API or for any other reason.
package lolesports

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	baseURL = "https://esports-api.lolesports.com/persisted/gw/"

	headerAPIKey = "X-Api-Key" //nolint:gosec

	//nolint:gosec
	// apiKey is the key used by the official https://lolesports.com website.
	// This is not a private API key.
	apiKey = "0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z"
)

// ClientOption represents a functional option to customize a [Client].
type ClientOption func(*Client)

// WithHTTPClient sets the [Client] with  a custom [http.Client].
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// Client fetches data from the unofficial LoL Esports API.
type Client struct {
	httpClient *http.Client
}

// NewClient returns a newly instanciated [Client].
//
// By default it uses the [http.DefaultClient] internally but
// you can customize it using [ClientOption].
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// GetScheduleOptions represents the options for fetching a schedule.
type GetScheduleOptions struct {
	// LeagueIDs is a list of league IDs used to filter the fetched events.
	// When set, only events related to the specified leagues will be fetched.
	LeagueIDs []string

	// PageToken is a base64 encoded token that identifies a specific schedule page.
	// This can be used to fetch a particular page of the schedule.
	PageToken *string
}

// GetSchedule retrieves the LoL Esports schedule. Without any options set,
// it retrieves the most recent events by default. You can use [GetScheduleOptions]
// to customize your request.
func (c *Client) GetSchedule(ctx context.Context, opts *GetScheduleOptions) (Schedule, error) {
	params := map[string]string{}
	if opts != nil {
		for _, leagueID := range opts.LeagueIDs {
			params["leagueId"] = strings.Join([]string{params["leagueID"], leagueID}, ",")
		}
		if opts.PageToken != nil {
			params["pageToken"] = *opts.PageToken
		}
	}

	req, err := newRequest(ctx, "getSchedule", params)
	if err != nil {
		return Schedule{}, fmt.Errorf("could not create request: %w", err)
	}

	var responseBody struct {
		Data struct {
			Schedule Schedule `json:"schedule"`
		} `json:"data"`
	}
	err = c.doRequest(req, &responseBody)
	if err != nil {
		return Schedule{}, err
	}
	return responseBody.Data.Schedule, nil
}

// GetSeasonsOptions represents the options for fetching seasons.
type GetSeasonsOptions struct {
	// ID can be set to only retrieve a specific season.
	ID *string
}

// GetSeasons fetches all LoL Esports seasons from the beginning of the records.
// Optionally, you can use [GetSeasonsOptions] to request a specific season.
func (c *Client) GetSeasons(ctx context.Context, opts *GetSeasonsOptions) ([]Season, error) {
	params := map[string]string{}
	if opts != nil {
		if opts.ID != nil {
			params["id"] = *opts.ID
		}
	}
	req, err := newRequest(ctx, "getSeasons", params)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	var responseBody struct {
		Data struct {
			Seasons []Season `json:"seasons"`
		} `json:"data"`
	}
	err = c.doRequest(req, &responseBody)
	if err != nil {
		return nil, err
	}
	return responseBody.Data.Seasons, nil
}

// GetStandings retrieves the standings for each tournament specified by the given tournament IDs.
func (c *Client) GetStandings(ctx context.Context, tournamentIDs []string) ([]Standings, error) {
	params := map[string]string{
		"tournamentId": strings.Join(tournamentIDs, ","),
	}
	req, err := newRequest(ctx, "getStandings", params)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	var responseBody struct {
		Data struct {
			Standings []Standings `json:"standings"`
		} `json:"data"`
	}
	err = c.doRequest(req, &responseBody)
	if err != nil {
		return nil, err
	}
	return responseBody.Data.Standings, nil
}

func (c *Client) doRequest(req *http.Request, response any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return fmt.Errorf("could not decode the response body: %w", err)
	}
	return nil
}

func newRequest(
	ctx context.Context,
	endpoint string,
	params map[string]string,
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set(headerAPIKey, apiKey)
	q := req.URL.Query()
	q.Add("hl", "en-US")
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}
