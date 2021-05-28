package roster

type Roster struct {
	UUID    string
	Lessons []Lesson
}

type Lesson struct {
	ID          int
	ClassID     string
	ProfessorID string
	SubjectID   string
	Moment      LessonMoment
}

type LessonMoment struct {
	DayOfWeek   int
	LessonOfDay int
}