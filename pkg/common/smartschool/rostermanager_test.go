package smartschool_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"piccolo.com/planner/pkg/common/smartschool"
	"piccolo.com/planner/pkg/planning/roster"
)

var rosterManager = smartschool.ProvideRosterManager()
var csvFile *csv.Reader
var expected roster.Roster

func TestImportRosterCSV(t *testing.T) {
	givenAValidRosterCSV(t)
	var result, err = rosterManager.ImportCSV(csvFile)
	
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestImportRosterCSVReturnsErrorForInvalidFileNil(t *testing.T) {
	csvFile = nil
	var result, err = rosterManager.ImportCSV(csvFile)
	assert.NotNil(t, err)
	assert.Equal(t, roster.Roster{}, result)
} 

func TestImportRosterCSVReturnsErrorForMissingFile(t *testing.T) {
	givenMissingRosterCSV(t)
	var result, err = rosterManager.ImportCSV(csvFile)
	assert.NotNil(t, err)
	assert.Equal(t, roster.Roster{}, result)
}

func TestImportRosterCSVReturnsErrorForInvalidFile(t *testing.T) {
	givenInvalidRosterCSV(t)
	var result, err = rosterManager.ImportCSV(csvFile)
	assert.NotNil(t, err)
	assert.Equal(t, roster.Roster{}, result)
} 


func givenAValidRosterCSV(t *testing.T) {
	file, err := os.Open("resources/test-roster-valid.csv")
	if err != nil {
		fmt.Println(err)
		assert.Fail(t,"Test resources test-roster-valid.csv not found.")
	}
	csvFile = csv.NewReader(file)
	csvFile.Comma = ';'
	expected = roster.Roster{
		UUID: "", // Needs to be empty for a new roster
		Lessons: []roster.Lesson{
			{
				ID: 1381,
				ClassID: "1AMa1",
				ProfessorID: "JN",
				SubjectID: "AAR",
				Moment: roster.LessonMoment{
					DayOfWeek: 2,
					LessonOfDay: 3,
				},
			},
			{
				ID: 1387,
				ClassID: "1ALa",
				ProfessorID: "VDBM",
				SubjectID: "KLS",
				Moment: roster.LessonMoment{
					DayOfWeek: 3,
					LessonOfDay: 2,
				},
			},
		},
	}
}

func givenMissingRosterCSV(t *testing.T) {
	file, _ := os.Open("resources/does-not-exist.csv")
	csvFile = csv.NewReader(file)
}

func givenInvalidRosterCSV(t *testing.T) {
	file, err := os.Open("resources/test-roster-invalid.csv")
	if err != nil {
		fmt.Println(err)
		assert.Fail(t,"Test resources test-roster-invalid.csv not found.")
	}
	csvFile = csv.NewReader(file)
	csvFile.Comma = ';'
}