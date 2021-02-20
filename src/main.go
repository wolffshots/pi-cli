package main

import (
	"log"
	"os"
	"time"
)

// constant values required by Pi-CLI
const (
	// port that the Pi-Hole API is defaulted to
	defaultPort = 80
	// the default refresh rate of the data in seconds
	defaultRefreshS = 1
	// the name of the configuration file
	configFileName = "config.json"
	// the starting setting for the number of queries that are included in the live log
	defaultAmountOfQueries = 10
)

// stores the data needed by Pi-CLI during runtime
type PiCLIData struct {
	Settings            *Settings
	FormattedAPIAddress string
	APIKey              string
	LastUpdated         time.Time
	ShowKeybindsScreen  bool
	Keybinds            []string
}

var piCLIData = PiCLIData{
	Keybinds: []string{
		"",
		"---------- Query Log ----------",
		"",
		"         [E]  Increase number of queries in query log",
		"         [D]  Decrease number of queries in query log",
		"  [UP ARROW]  Scroll up query log",
		"[DOWN ARROW]  Scroll down query log",
		"",
		"---------- Misc. ----------",
		"",
		"[Q]  Quit Pi-CLI",
	},
}

var summary = Summary{
	QueriesToday:        "",
	BlockedToday:        "",
	PercentBlockedToday: "",
	DomainsOnBlocklist:  "",
	Status:              "",
	PrivacyLevel:        "",
	PrivacyLevelNumberMapping: map[string]string{
		"0": "Show Everything",
		"1": "Hide Domains",
		"2": "Hide Domains and Clients",
		"3": "Anonymous",
	},
	TotalClientsSeen: "",
}

var topItems = TopItems{
	TopQueries:       map[string]int{},
	TopAds:           map[string]int{},
	PrettyTopQueries: []string{},
	PrettyTopAds:     []string{},
}

var allQueries = AllQueries{
	Queries:              make([]Query, defaultAmountOfQueries),
	AmountOfQueriesInLog: defaultAmountOfQueries,
	Table:                []string{},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
