package population_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"piccolo.com/planner/pkg/genetic/population"
)

func TestInternCalendarToString(t *testing.T) {
	var calendar = population.InternCalendar{
		Fitness: 0.3232323,
		InternMoments: []population.InternMoment{	
			{
				LessonID: 1,
				WeekID: 12,
			},	
			{
				LessonID: 3,
				WeekID: 13,
			},	
			{
				LessonID: 9,
				WeekID: 35,
			},	
			{
				LessonID: 2,
				WeekID: 10,
			},
		},
	}
	assert.Equal(t, "1-12;3-13;9-35;2-10", calendar.GetIdentifier())
}