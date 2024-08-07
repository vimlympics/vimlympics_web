package model

type Record struct {
	Name    string
	Country string
	Gold    int
	Silver  int
	Bronze  int
}

type LeaderboardRecords []Record
