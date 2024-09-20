package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mazdazoomzoom/finalturnracing/pkg/models"
)

type APIData struct {
	Series1 []Race `json:"series_1"`
	Series2 []Race `json:"series_2"`
	Series3 []Race `json:"series_3"`
}

type Race struct {
	RaceID        int    `json:"race_id"`
	SeriesID      int    `json:"series_id"`
	RaceSeason    int    `json:"race_season"`
	RaceName      string `json:"race_name"`
	RaceTypeID    int    `json:"race_type_id"`
	TrackName     string `json:"track_name"`
	ScheduledLaps int    `json:"scheduled_laps"`
	Stage1Laps    int    `json:"stage_1_laps"`
	Stage2Laps    int    `json:"stage_2_laps"`
	Stage3Laps    int    `json:"stage_3_laps"`
	Stage4Laps    int    `json:"stage_4_laps"`
	PlayoffRound  int    `json:"playoff_round"`
}

func getSchedule() ([]models.Schedule, error) {
	year := time.Now().Year()
	api_url_prefix := os.Getenv("NASCAR_API")
	api_url_suffix := os.Getenv("NASCAR_SCHEDULE_API_URL")
	url := fmt.Sprintf("%s/%d/%s", api_url_prefix, year, api_url_suffix)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching schedule: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	var data APIData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v\n", err)
		return nil, err
	}

	var schedule []models.Schedule

	// Series 1
	for _, item := range data.Series1 {
		schedule = append(schedule, models.Schedule{
			RaceID:        item.RaceID,
			SeriesID:      item.SeriesID,
			RaceSeason:    item.RaceSeason,
			RaceName:      item.RaceName,
			RaceTypeID:    item.RaceTypeID,
			TrackName:     item.TrackName,
			ScheduledLaps: item.ScheduledLaps,
			Stage1Laps:    item.Stage1Laps,
			Stage2Laps:    item.Stage2Laps,
			Stage3Laps:    item.Stage3Laps,
			Stage4Laps:    item.Stage4Laps,
			PlayoffRound:  item.PlayoffRound,
		})
	}

	// Series 2
	for _, item := range data.Series2 {
		schedule = append(schedule, models.Schedule{
			RaceID:        item.RaceID,
			SeriesID:      item.SeriesID,
			RaceSeason:    item.RaceSeason,
			RaceName:      item.RaceName,
			RaceTypeID:    item.RaceTypeID,
			TrackName:     item.TrackName,
			ScheduledLaps: item.ScheduledLaps,
			Stage1Laps:    item.Stage1Laps,
			Stage2Laps:    item.Stage2Laps,
			Stage3Laps:    item.Stage3Laps,
			Stage4Laps:    item.Stage4Laps,
			PlayoffRound:  item.PlayoffRound,
		})
	}

	// Series 3
	for _, item := range data.Series3 {
		schedule = append(schedule, models.Schedule{
			RaceID:        item.RaceID,
			SeriesID:      item.SeriesID,
			RaceSeason:    item.RaceSeason,
			RaceName:      item.RaceName,
			RaceTypeID:    item.RaceTypeID,
			TrackName:     item.TrackName,
			ScheduledLaps: item.ScheduledLaps,
			Stage1Laps:    item.Stage1Laps,
			Stage2Laps:    item.Stage2Laps,
			Stage3Laps:    item.Stage3Laps,
			Stage4Laps:    item.Stage4Laps,
			PlayoffRound:  item.PlayoffRound,
		})
	}

	return schedule, nil
}

func GetSchedule() {
	getSchedule()
}
