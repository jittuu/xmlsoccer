package xmlsoccer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAllLeagues(t *testing.T) {
	// arrange
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := `
<?xml version="1.0" encoding="utf-8"?>
<XMLSOCCER.COM>
    <League>
        <Id>1</Id>
        <Name>English Premier League</Name>
        <Country>England</Country>
        <Historical_Data>Yes</Historical_Data>
        <Fixtures>Yes</Fixtures>
        <Livescore>Yes</Livescore>
        <NumberOfMatches>2557</NumberOfMatches>
        <LatestMatch>2013-03-02T16:00:00+01:00</LatestMatch>
    </League>
    <League>
        <Id>3</Id>
        <Name>Scottish Premier League</Name>
        <Country>Scotland</Country>
        <Historical_Data>Yes</Historical_Data>
        <Fixtures>Yes</Fixtures>
        <Livescore>Yes</Livescore>
        <NumberOfMatches>1314</NumberOfMatches>
        <LatestMatch>2013-03-02T16:00:00+01:00</LatestMatch>
    </League>
    <League>
        <Id>4</Id>
        <Name>Bundesliga</Name>
        <Country>Germany</Country>
        <Historical_Data>Yes</Historical_Data>
        <Fixtures>Yes</Fixtures>
        <Livescore>Yes</Livescore>
        <NumberOfMatches>1743</NumberOfMatches>
        <LatestMatch>2013-03-02T15:30:00+01:00</LatestMatch>
    </League>
    <AccountInformation>Data requested at 02-03-2013 21:02:09 from XX.XX.XX.XX, Username: Espectro. Your current supscription runs out on XX-XX-XXXX 11:01:25.</AccountInformation>
</XMLSOCCER.COM>
    `
		fmt.Fprintln(w, data)
	}))
	defer ts.Close()

	// act
	c := &Client{APIKey: "dummy-key"}
	c.testURL = ts.URL
	leagues, err := c.GetAllLeagues()

	// assert
	if err != nil {
		t.Error(err)
	}
	if len(leagues) != 3 {
		t.Errorf("expected %d leagues but return %d", 3, len(leagues))
		t.Log(leagues)
	}

	epl := leagues[0]
	if epl.ID != 1 {
		t.Errorf("expected ID: %d, got %d", 1, epl.ID)
	}
	if epl.Name != "English Premier League" {
		t.Errorf("expected name: %q, got %q", "English Premier League", epl.Name)
	}
	expectedTime, _ := time.Parse("2006-01-02T15:04:00+00:00", "2013-03-02T16:00:00+01:00")
	if epl.LatestMatch == expectedTime {
		t.Errorf("expected time: %v, got %v", expectedTime, epl.LatestMatch)
	}
}

func TestGetFixturesByDateInterval(t *testing.T) {
	testGetFixtures(t, func(c *Client) ([]*Match, error) {
		return c.GetFixturesByDateInterval(time.Now().Add(-7*24*time.Hour), time.Now())
	})
}

func TestGetFixturesByDateIntervalAndLeague(t *testing.T) {
	testGetFixtures(t, func(c *Client) ([]*Match, error) {
		return c.GetFixturesByDateIntervalAndLeague(time.Now().Add(-7*24*time.Hour), time.Now(), "3")
	})
}

func TestGetFixturesByLeagueAndSeason(t *testing.T) {
	testGetFixtures(t, func(c *Client) ([]*Match, error) {
		return c.GetFixturesByLeagueAndSeason("3", "1415")
	})
}

func testGetFixtures(t *testing.T, f func(*Client) ([]*Match, error)) {
	// arrange
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := `
		<?xml version="1.0" encoding="utf-8"?>
		<XMLSOCCER.COM>
				<Match>
					<Id>349073</Id>
					<Date>2015-05-16T14:00:00+00:00</Date>
					<League>Scottish Premier League</League>
					<Round>37</Round>
					<HomeTeam>Inverness C</HomeTeam>
					<HomeTeam_Id>48</HomeTeam_Id>
					<HomeGoals>3</HomeGoals>
					<AwayTeam>Dundee United</AwayTeam>
					<AwayTeam_Id>51</AwayTeam_Id>
					<AwayGoals>0</AwayGoals>
					<Time>Finished</Time>
					<Location>Caledonian Stadium</Location>
					<HomeTeamYellowCardDetails>77': Daniel Devine;36': Gary Warren;</HomeTeamYellowCardDetails>
					<AwayTeamYellowCardDetails>55': Paul Dixon;55': Chris Erskine;</AwayTeamYellowCardDetails>
					<HomeTeamRedCardDetails/>
					<AwayTeamRedCardDetails/>
				</Match>
		</XMLSOCCER.COM>
				`
		fmt.Fprintln(w, data)
	}))
	defer ts.Close()

	// act
	c := &Client{APIKey: "dummy-key"}
	c.testURL = ts.URL
	matches, err := f(c)

	// assert
	if err != nil {
		t.Error(err)
	}

	if len(matches) != 1 {
		t.Errorf("expected matches %d, got %d", 1, len(matches))
	}

	m := matches[0]
	if m.ID != 349073 {
		t.Errorf("expected ID %d, got %d", 349073, m.ID)
	}

	expectedStartDate, _ := time.Parse(iso8601Layout, "2015-05-16T14:00:00+00:00")
	if !m.StartDate.Equal(expectedStartDate) {
		t.Errorf("expected start date %v, got %v", expectedStartDate, m.StartDate)
	}

	if m.Round != 37 {
		t.Errorf("expected Round %d, got %d", 37, m.Round)
	}

	if m.HomeTeamName != "Inverness C" {
		t.Errorf("expected home %q, got %q", "Inverness C", m.HomeTeamName)
	}

	if m.HomeTeamID != 48 {
		t.Errorf("expected home id %d, got %d", 48, m.HomeTeamID)
	}

	if m.HomeGoals != 3 {
		t.Errorf("expected home goals %d, got %d", 3, m.HomeGoals)
	}

	if m.AwayTeamName != "Dundee United" {
		t.Errorf("expected home %q, got %q", "Dundee United", m.AwayTeamName)
	}

	if m.AwayTeamID != 51 {
		t.Errorf("expected home id %d, got %d", 48, m.AwayTeamID)
	}

	if m.AwayGoals != 0 {
		t.Errorf("expected home goals %d, got %d", 3, m.AwayGoals)
	}

	if m.Time != TimeFinished {
		t.Errorf("expected time %q, got %q", TimeFinished, m.Time)
	}

	if t.Failed() {
		t.Log(matches)
	}
}

func TestGetAllTeamsByLeagueAndSeason(t *testing.T) {
	// arrange
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := `
		<?xml version="1.0" encoding="utf-8"?>
		<XMLSOCCER.COM>
			<Team>
				<Team_Id>4</Team_Id>
				<Name>Fulham</Name>
				<Country>England</Country>
				<Stadium>Craven Cottage</Stadium>
				<HomePageURL>http://www.fulhamfc.com/</HomePageURL>
				<WIKILink>http://en.wikipedia.org/wiki/Fulham_F.C.</WIKILink>
			</Team>
		</XMLSOCCER.COM>
				`
		fmt.Fprintln(w, data)
	}))
	defer ts.Close()

	// act
	c := &Client{APIKey: "dummy-key"}
	c.testURL = ts.URL
	teams, err := c.GetAllTeamsByLeagueAndSeason("3", "1415")

	// assert
	if err != nil {
		t.Error(err)
	}
	if len(teams) != 1 {
		t.Errorf("expected teams %d, got %d", 1, len(teams))
	}

	team := teams[0]
	if team.ID != 4 {
		t.Errorf("expected ID %d, got %d", 4, team.ID)
	}
	if team.Name != "Fulham" {
		t.Errorf("expected name %q, got %q", "Fulham", team.Name)
	}
	if team.Country != "England" {
		t.Errorf("expected country %q, got %q", "England", team.Country)
	}
	if team.WikiLink != "http://en.wikipedia.org/wiki/Fulham_F.C." {
		t.Errorf("expected wiki-link %q, got %q", "http://en.wikipedia.org/wiki/Fulham_F.C.", team.WikiLink)
	}
}
