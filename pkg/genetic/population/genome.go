package population

import (
	"strconv"
	"strings"
)

type InternCalendar struct {
	Fitness float32
	InternMoments []InternMoment
}

type InternMoment struct {
	LessonID int
	WeekID int
}

func (c InternCalendar) GetIdentifier() string {
	var genome []string
	for _, v := range c.InternMoments {
		genome = append(genome, strconv.Itoa(v.LessonID) + "-" + strconv.Itoa(v.WeekID))
	}
	return strings.Join(genome, ";")
}