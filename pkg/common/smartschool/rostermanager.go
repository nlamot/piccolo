package smartschool

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"

	"piccolo.com/planner/pkg/planning/roster"
)

type RosterManager interface {
	ImportCSV(*csv.Reader) (roster.Roster, error)
}

type rosterManager struct {
}

func (r *rosterManager) ImportCSV(csv *csv.Reader) (roster.Roster, error) {
	var importedRoster = roster.Roster{}  
	if (csv == nil) {
		return importedRoster, errors.New("failed to import roster csv, reader must not be nil")
	}
	csvLines, err := csv.ReadAll()
    if err != nil {
		fmt.Println("Failed to import roster csv.")
		return importedRoster, err
    }  
    for _, line := range csvLines {
		importedRoster.Lessons = append(importedRoster.Lessons, parseLesson(line))
    }
	return importedRoster, nil
}

func ProvideRosterManager() RosterManager {
	return &rosterManager{}
}

func parseLesson(csvLine []string) roster.Lesson {
	var id, _ = strconv.Atoi(csvLine[0])
	var dayOfWeek, _ = strconv.Atoi(csvLine[5])
	var lessonOfDay, _ = strconv.Atoi(csvLine[6])
	return roster.Lesson{
		ID: id,
		ClassID: csvLine[1],
		ProfessorID: csvLine[2],
		SubjectID: csvLine[3],
		Moment: roster.LessonMoment{
			DayOfWeek: dayOfWeek,
			LessonOfDay: lessonOfDay,
		},
	}
}