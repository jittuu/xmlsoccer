package xmlsoccer

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Client is to call webservice
type Client struct {
	*http.Client

	// it will be zero value ("") while not in testing
	testURL string

	// API key to access the service
	APIKey string

	// the base url for webservice
	BaseURL string
}

// ErrMissingAPIKey represents error when client makes request without APIKey
var (
	ErrMissingAPIKey = errors.New("APIKey is requried")
	DemoURL          = "http://www.xmlsoccer.com/FootballDataDemo.asmx"
	FullURL          = "http://www.xmlsoccer.com/FootballData.asmx"
)

// GetAllLeagues returns all published leagues
func (c *Client) GetAllLeagues() ([]*League, error) {
	result := xmlroot{}
	err := c.invokeService("GetAllLeagues", url.Values{}, &result)
	if err != nil {
		return nil, err
	}

	return result.Leagues, nil
}

// GetFixturesByDateInterval returns all match fixtures between the given interval
func (c *Client) GetFixturesByDateInterval(startDate, endDate time.Time) ([]*Match, error) {
	result := xmlroot{}
	s, e := convertToCET(startDate, endDate)
	err := c.invokeService("GetFixturesByDateInterval",
		url.Values{"startDateString": {s.Format(dateparamLayout)}, "endDateString": {e.Format(dateparamLayout)}},
		&result)
	if err != nil {
		return nil, err
	}

	return result.Matches, nil
}

// GetFixturesByDateIntervalAndLeague returns all match fixtures for the given league between the given interval
func (c *Client) GetFixturesByDateIntervalAndLeague(startDate, endDate time.Time, league string) ([]*Match, error) {
	result := xmlroot{}
	s, e := convertToCET(startDate, endDate)
	err := c.invokeService("GetFixturesByDateIntervalAndLeague",
		url.Values{
			"startDateString": {s.Format(dateparamLayout)},
			"endDateString":   {e.Format(dateparamLayout)},
			"league":          {league},
		},
		&result)
	if err != nil {
		return nil, err
	}

	return result.Matches, nil
}

// GetFixturesByLeagueAndSeason returns all match fixtures for the given league and season
func (c *Client) GetFixturesByLeagueAndSeason(league, season string) ([]*Match, error) {
	result := xmlroot{}
	err := c.invokeService("GetFixturesByLeagueAndSeason",
		url.Values{
			"league":           {league},
			"seasonDateString": {season},
		},
		&result)
	if err != nil {
		return nil, err
	}

	return result.Matches, nil
}

// GetAllTeamsByLeagueAndSeason returns all teams for the given league and season
func (c *Client) GetAllTeamsByLeagueAndSeason(league, season string) ([]*Team, error) {
	result := xmlroot{}
	err := c.invokeService("GetAllTeamsByLeagueAndSeason",
		url.Values{
			"league":           {league},
			"seasonDateString": {season},
		},
		&result)
	if err != nil {
		return nil, err
	}

	return result.Teams, nil
}

func (c *Client) invokeService(serviceName string, data url.Values, v interface{}) error {
	if c.APIKey == "" {
		return ErrMissingAPIKey
	}
	data.Add("ApiKey", c.APIKey)

	if c.Client == nil {
		c.Client = http.DefaultClient
	}

	resp, err := c.PostForm(c.postURL(serviceName), data)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(content, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) postURL(serviceName string) string {
	if c.testURL != "" {
		return c.testURL
	}

	return c.BaseURL + "/" + serviceName
}

func convertToCET(start, end time.Time) (time.Time, time.Time) {
	cet, _ := time.LoadLocation("CET")
	return start.In(cet), end.In(cet)
}
