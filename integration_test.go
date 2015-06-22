// +build integration

package xmlsoccer

import (
	"flag"
	"testing"
	"time"
)

var (
	apiKeyPtr = flag.String("api", "", "provide api key to run integration tests")
)

func TestIntegrationGetAllLeagues(t *testing.T) {
	checkAPIKey(t)

	c := DemoClient(*apiKeyPtr)
	leagues, err := c.GetAllLeagues()

	if err != nil {
		t.Error(err)
	}

	if len(leagues) < 1 {
		t.Errorf("expected more than one leagues, got %d", len(leagues))
	}
}

func TestIntegrationGetFixturesByDateInterval(t *testing.T) {
	testIntegrationGetFixtures(t, func(c *Client) ([]Match, error) {
		return c.GetFixturesByDateInterval(time.Now().Add(-10*24*time.Hour), time.Now())
	})
}

func TestIntegrationGetFixturesByDateIntervalAndLeague(t *testing.T) {
	testIntegrationGetFixtures(t, func(c *Client) ([]Match, error) {
		return c.GetFixturesByDateIntervalAndLeague(time.Now().Add(-10*24*time.Hour), time.Now(), "3")
	})
}

func TestIntegrationGetFixturesByLeagueAndSeason(t *testing.T) {
	testIntegrationGetFixtures(t, func(c *Client) ([]Match, error) {
		return c.GetFixturesByLeagueAndSeason("3", "1415")
	})
}

func testIntegrationGetFixtures(t *testing.T, f func(*Client) ([]Match, error)) {
	checkAPIKey(t)

	c := DemoClient(*apiKeyPtr)
	matches, err := f(c)

	if err != nil {
		t.Error(err)
	}

	if len(matches) < 1 {
		t.Errorf("expected more than one match, got %d", len(matches))
	}
}

func TestIntegrationGetAllTeamsByLeagueAndSeason(t *testing.T) {
	checkAPIKey(t)

	c := DemoClient(*apiKeyPtr)
	teams, err := c.GetAllTeamsByLeagueAndSeason("3", "1415")

	if err != nil {
		t.Error(err)
	}

	if len(teams) < 1 {
		t.Errorf("expected more than one team, got %d", len(teams))
	}
}

func checkAPIKey(t *testing.T) {
	apiKey := *apiKeyPtr
	if apiKey == "" {
		t.Log("API key is required to run integration test. Please provide via command line flag")
		t.Fail()
	}
}
