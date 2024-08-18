package clydetools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const leageTwoLeagueId = "184"
const clydeTeamId = "6762"

var CurrentTime = func() time.Time {
	return time.Now()
}

type Root struct {
	Get        string        `json:"get"`
	Parameters Parameters    `json:"parameters"`
	Errors     []interface{} `json:"errors"` // Since the errors array is empty, it can hold any type
	Results    int           `json:"results"`
	Paging     Paging        `json:"paging"`
	Response   []Response    `json:"response"`
}

type Parameters struct {
	League string `json:"league"`
	Season string `json:"season"`
	Team   string `json:"team"`
	Next   string `json:"next"`
}

type Paging struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type Response struct {
	Fixture Fixture `json:"fixture"`
	League  League  `json:"league"`
	Teams   Teams   `json:"teams"`
	Goals   Goals   `json:"goals"`
	Score   Score   `json:"score"`
}

type Fixture struct {
	ID        int     `json:"id"`
	Referee   *string `json:"referee"` // Using *string to handle potential null values
	Timezone  string  `json:"timezone"`
	Date      string  `json:"date"`
	Timestamp int64   `json:"timestamp"`
	Periods   Periods `json:"periods"`
	Venue     Venue   `json:"venue"`
	Status    Status  `json:"status"`
}

type Periods struct {
	First  *int `json:"first"`
	Second *int `json:"second"`
}

type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Status struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed *int   `json:"elapsed"`
}

type League struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
	Round   string `json:"round"`
}

type Teams struct {
	Home Team `json:"home"`
	Away Team `json:"away"`
}

type Team struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner *bool  `json:"winner"`
}

type Goals struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

type Score struct {
	Halftime  Halftime  `json:"halftime"`
	Fulltime  Fulltime  `json:"fulltime"`
	Extratime Extratime `json:"extratime"`
	Penalty   Penalty   `json:"penalty"`
}

type Halftime struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

type Fulltime struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

type Extratime struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

type Penalty struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

func GetCurrentSeasonYear() string {
	year, month, _ := CurrentTime().Date()

	if month < 8 {
		year = year - 1
	}

	return strconv.Itoa(year)
}

func parseFixtures(body []byte) [10]string {
	var fixtures [10]string
	var fixturesResponse Root
	json.Unmarshal(body, &fixturesResponse)
	for i, value := range fixturesResponse.Response {
		home := value.Teams.Home.Name
		away := value.Teams.Away.Name
		fixture := fmt.Sprintf("%s vs %s", home, away)
		fixtures[i] = fixture
	}
	return fixtures
}

func getFixtures() [10]string {
	url := "https://v3.football.api-sports.io/fixtures?league=" + leageTwoLeagueId + "&season=" + GetCurrentSeasonYear() + "&team=" + clydeTeamId + "&next=10"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return [10]string{}
	}
	req.Header.Add("x-rapidapi-key", "d932a7c27dd5f4e3eace3d48c0378b40")
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return [10]string{}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return [10]string{}
	}

	return parseFixtures(body)
}

func GetFixtures() [10]string {
	var fixtures [10]string = getFixtures()
	return fixtures
}
