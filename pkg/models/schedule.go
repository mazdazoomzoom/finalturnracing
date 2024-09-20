package models

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	RaceID        int
	SeriesID      int
	RaceSeason    int
	RaceName      string
	RaceTypeID    int
	TrackName     string
	ScheduledLaps int
	Stage1Laps    int
	Stage2Laps    int
	Stage3Laps    int
	Stage4Laps    int
	PlayoffRound  int
}
