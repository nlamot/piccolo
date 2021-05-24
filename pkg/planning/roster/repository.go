package roster

type RosterRepository interface {
	Create(Roster) string
}

type rosterRepository struct {
}

func (r *rosterRepository) Create(roster Roster) string {
	return ""
}

func ProvideRosterRepository() RosterRepository {
	return &rosterRepository{}
}