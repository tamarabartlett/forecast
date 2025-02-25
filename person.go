package forecast

import (
	"errors"
	"fmt"
	"time"
)

type peopleContainer struct {
	People People `json:"people"`
}

type personContainer struct {
	Person Person `json:"person"`
}

// People is a list of people
type People []Person

// Person is a person who is being scheduled in Forecast
type Person struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Login          string    `json:"login"`
	Admin          bool      `json:"admin"`
	Archived       bool      `json:"archived"`
	Subscribed     bool      `json:"subscribed"`
	AvatarURL      string    `json:"avatar_url"`
	Teams          []string  `json:"teams"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedByID    int       `json:"updated_by_id"`
	HarvestUserID  int       `json:"harvest_user_id"`
	WeeklyCapacity int       `json:"weekly_capacity"`
	WorkingDays    struct {
		Monday    bool `json:"monday"`
		Tuesday   bool `json:"tuesday"`
		Wednesday bool `json:"wednesday"`
		Thursday  bool `json:"thursday"`
		Friday    bool `json:"friday"`
	} `json:"working_days"`
	ColorBlind bool `json:"color_blind"`
}

// People returns all people being scheduled in Forecast
func (api *API) People() (People, error) {
	var container peopleContainer
	err := api.do("people", &container)
	if err != nil {
		return nil, err
	}
	return container.People, nil
}

// Person returns the person with the requested ID
func (api *API) Person(id int) (*Person, error) {
	if id == 0 {
		return nil, errors.New("cannot retrieve a person with an id of 0")
	}
	var container personContainer
	err := api.do(fmt.Sprintf("people/%v", id), &container)
	if err != nil {
		return nil, err
	}
	return &container.Person, nil
}
