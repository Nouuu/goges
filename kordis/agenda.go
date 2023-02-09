package kordis

import (
	"encoding/json"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

type Agenda struct {
	ResponseCode int           `json:"response_code"`
	Version      string        `json:"version"`
	Result       []AgendaEvent `json:"result"`
	Links        []string      `json:"links"`
}

type AgendaEvent struct {
	ReservationId         int        `json:"reservation_id"`
	Rooms                 []Room     `json:"rooms"`
	Type                  string     `json:"type"`
	Modality              string     `json:"modality"`
	Author                int        `json:"author"`
	CreateDate            int64      `json:"create_date"` // Unix timestamp milliseconds
	StartDate             int64      `json:"start_date"`  // Unix timestamp milliseconds
	EndDate               int64      `json:"end_date"`    // Unix timestamp milliseconds
	Comment               string     `json:"comment"`
	Classes               []string   `json:"classes"`
	Name                  string     `json:"name"`
	Discipline            Discipline `json:"discipline"`
	Teacher               string     `json:"teacher"`
	Promotion             string     `json:"promotion"`
	PrestationType        int        `json:"prestation_type"`
	IsElectronicSignature bool       `json:"is_electronic_signature"`
	Links                 []string   `json:"links"`
}

type Room struct {
	Links     []string `json:"links"`
	RoomId    int      `json:"room_id"`
	Name      string   `json:"name"`
	Floor     string   `json:"floor"`
	Campus    string   `json:"campus"`
	Color     string   `json:"color"`
	Latitude  string   `json:"latitude"`
	Longitude string   `json:"longitude"`
}

type Discipline struct {
	Coef             any      `json:"coef"`
	Ects             any      `json:"ects"`
	Name             string   `json:"name"`
	Teacher          string   `json:"teacher"`
	Trimester        string   `json:"trimester"`
	Year             int      `json:"year"`
	Links            []string `json:"links"`
	HasDocuments     any      `json:"has_documents"`
	HasGrades        any      `json:"has_grades"`
	NbStudents       int      `json:"nb_students"`
	RcId             int      `json:"rc_id"`
	SchoolId         int      `json:"school_id"`
	StudentGroupId   int      `json:"student_group_id"`
	StudentGroupName string   `json:"student_group_name"`
	SyllabusId       any      `json:"syllabus_id"`
	TeacherId        int      `json:"teacher_id"`
	TrimesterId      int      `json:"trimester_id"`
}

func (mygesApi *KordisApi) GetAgenda(start time.Time, end time.Time) (Agenda, error) {
	queryParams := map[string]string{
		"start": strconv.FormatInt(start.UnixMilli(), 10),
		"end":   strconv.FormatInt(end.UnixMilli(), 10),
	}
	response, err := mygesApi.Get(kordisAgendaUrl, queryParams)
	if err != nil {
		panic(err)
	}
	var agenda Agenda
	err = json.Unmarshal(response.Body(), &agenda)
	return agenda, err
}

func (mygesApi *KordisApi) GetAgendaFromNow(days int) (Agenda, error) {
	start := carbon.Now().StartOfDay().Carbon2Time()
	end := carbon.Now().AddDays(int(math.Max(float64(days), 0))).EndOfDay().Carbon2Time()
	return mygesApi.GetAgenda(start, end)
}

func PrintAgenda(agenda Agenda) {
	lang := carbon.NewLanguage()
	lang.SetLocale("fr")
	c := carbon.SetLanguage(lang).SetTimezone(os.Getenv("TZ"))

	sort.Slice(agenda.Result, func(i, j int) bool {
		return agenda.Result[i].StartDate < agenda.Result[j].StartDate
	})

	for _, event := range agenda.Result {
		startDate := c.CreateFromTimestamp(event.StartDate / 1000)
		endDate := c.CreateFromTimestamp(event.EndDate / 1000)
		fmt.Printf("%-*s %s - %-*s %-*s %-*s",
			25,
			startDate.Format("l d F"),
			startDate.Format("H:i"),
			10,
			endDate.Format("H:i"),
			50,
			event.Name,
			20,
			event.Teacher,
		)
		if event.Rooms != nil && len(event.Rooms) > 0 {
			fmt.Printf("%s (%s)", event.Rooms[0].Name, event.Rooms[0].Campus)
		}
		fmt.Println()
	}
}
