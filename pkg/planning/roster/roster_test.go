package roster_test

import "piccolo.com/planner/pkg/planning/roster"

var aRoster = roster.Roster{
	Lessons: []roster.Lesson{
		{
			ID: 1381,
			ClassID: "1AMa1",
			ProfessorID: "JN",
			SubjectID: "AAR",
			ClassRoomID: "0115",
			Moment: roster.LessonMoment{
				DayOfWeek: 2,
				LessonOfDay: 3,
			},
		},
	},
}