package roster

type Roster struct {
	UUID    string 		`firestore:"uuid"`
	Lessons []Lesson	`firestore:"lessons"`
}

type Lesson struct {
	ID          int				`firestore:"id"`
	ClassID     string			`firestore:"classId"`
	ProfessorID string			`firestore:"professorId"`
	SubjectID   string			`firestore:"subjectId"`
	Moment      LessonMoment	`firestore:"moment"`
}

type LessonMoment struct {
	DayOfWeek   int		`firestore:"dayOfWeek"`
	LessonOfDay int		`firestore:"lessonOfDay"`
}