package xmlsoccer

import (
	"encoding/xml"
	"time"
)

const (
	// TimeNotStarted is the value "Not Started" for Match's Time field
	TimeNotStarted = "Not Started"
	// TimeHalftime is the value "Halftime" for Match's Time field
	TimeHalftime = "Halftime"
	// TimeFinished is the value "Finished" for Match's Time field
	TimeFinished = "Finished"
	// TimeFinishedAET is the value "Finished AET" for Match's Time field
	TimeFinishedAET = "Finished AET"
	// TimeFinishedAP is the value "Finished AP" for Match's Time field
	TimeFinishedAP = "Finished AP"
	// TimeWaitingForPenalty is the value "Waiting for Penalty" for Match's Time field
	TimeWaitingForPenalty = "Waiting for Penalty"
	// TimeCancelled is the value "Cancelled" for Match's Time field
	TimeCancelled = "Cancelled"
	// TimePostponed is the value "Postponed" for Match's Time field
	TimePostponed = "Postponed"
	// TimeAbandoned is the value "Abandoned" for Match's Time field
	TimeAbandoned = "Abandoned"

	iso8601Layout   = "2006-01-02T15:04:05-07:00"
	dateparamLayout = "2006-01-02 15:04"
)

type xmlroot struct {
	XMLName xml.Name  `xml:"XMLSOCCER.COM"`
	Leagues []*League `xml:"League"`
	Matches []*Match  `xml:"Match"`
	Teams   []*Team   `xml:"Team"`
}

// League represents a soccer League
type League struct {
	// Id is unique identifier of a league
	ID int `xml:"Id"`

	// Name is the name of a league
	Name string

	// LatestMatch is the date of the last match for the league
	LatestMatch time.Time
}

// Match represents a soccer Match
type Match struct {
	// ID is unique identifier of a Match
	ID int `xml:"Id"`

	// StartDate is match start date time
	StartDate time.Time `xml:"Date"`

	// Round is the n-th play match
	Round int

	// HomeTeamName is home team's name
	HomeTeamName string `xml:"HomeTeam"`

	// HomeTeamID is home team's unique identifier
	HomeTeamID int `xml:"HomeTeam_Id"`

	// HomeGoals is home team's goals
	HomeGoals int

	// AwayTeamName is away team's name
	AwayTeamName string `xml:"AwayTeam"`

	// AwayTeamID is away team's unique identifier
	AwayTeamID int `xml:"AwayTeam_Id"`

	// AwayGoals is away team's goals
	AwayGoals int

	// Time can contain the following values
	// Not Started
	// X (where X is the minute of the match)
	// Halftime
	// Finished
	// Finished AET (Added Extra Time)
	// Finished AP (Added Penalty)
	// Waiting for Penalty
	// Cancelled
	// Postponed
	// Abandoned
	Time string
}

// Team represents a soccer team
type Team struct {
	// ID is unique identifier of a Team
	ID int `xml:"Team_Id"`

	// Name is team's name
	Name string

	// Country is team's home country
	Country string

	// WikiLink is a http link to wikipedia if there is any
	WikiLink string `xml:"WIKILink"`
}
