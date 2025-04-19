package lolesports

import "time"

type Schedule struct {
	Updated time.Time `json:"updated"`
	Pages   Pages     `json:"pages"`
	Events  []Event   `json:"events"`
}

type Pages struct {
	Older string `json:"older"`
	Newer string `json:"newer"`
}

type EventType string

const (
	EventTypeMatch = "match"
	EventTypeShow  = "show"
)

type EventState string

const (
	EventStateUnstarted  = "unstarted"
	EventStateInProgress = "inProgress"
	EventStateCompleted  = "completed"
)

type Event struct {
	StartTime time.Time  `json:"startTime"`
	BlockName string     `json:"blockName"`
	Match     Match      `json:"match"`
	State     EventState `json:"state"`
	Type      string     `json:"type"`
	League    League     `json:"league"`
}

type Match struct {
	ID               string   `json:"id"`
	PreviousMatchIDs []string `json:"previousMatchIds"`
	Flags            []string `json:"flags"`
	Teams            []Team   `json:"teams"`
	Strategy         Strategy `json:"strategy"`
}

type Team struct {
	ID     string  `json:"id"`
	Slug   string  `json:"slug"`
	Name   string  `json:"name"`
	Code   string  `json:"code"`
	Image  string  `json:"image"`
	Result *Result `json:"result"`
	Record *Record `json:"record"`
}

type Result struct {
	Outcome  *string `json:"outcome"`
	GameWins int     `json:"gameWins"`
}

type Record struct {
	Losses int `json:"losses"`
	Wins   int `json:"wins"`
}

type Strategy struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type MatchStrategyType string

const (
	MatchStrategyTypeBestOf = "bestOf"
)

type Standings struct {
	Stages []Stage `json:"stages"`
}

type Stage struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Slug     string    `json:"slug"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Name     string    `json:"name"`
	Matches  []Match   `json:"matches"`
	Rankings []Ranking `json:"rankings"`
}

type Ranking struct {
	Ordinal int    `json:"ordinal"`
	Teams   []Team `json:"teams"`
}

type Season struct {
	ID          string    `json:"id"`
	Description *string   `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Splits      []Split   `json:"splits"`
}

type Split struct {
	ID          string       `json:"id"`
	Description *string      `json:"description"`
	Name        string       `json:"name"`
	Slug        string       `json:"slug"`
	StartTime   time.Time    `json:"startTime"`
	EndTime     time.Time    `json:"endTime"`
	Region      string       `json:"region"`
	Tournaments []Tournament `json:"tournaments"`
}

type Tournament struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	League League `json:"league"`
}

type League struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Image           string          `json:"image"`
	DisplayPriority DisplayPriority `json:"displayPriority"`
}

type DisplayPriority struct {
	Position int    `json:"position"`
	Status   string `json:"status"`
}
