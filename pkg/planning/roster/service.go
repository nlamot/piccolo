package roster

type RosterService interface {
	// Create creates the new roster and returns the UUID
	Create(Roster) string
}

type rosterService struct {	
	repository RosterRepository
}

func (s *rosterService) Create(r Roster) string {
	return s.repository.Create(r)
}

func ProvideRosterService(repository RosterRepository) RosterService {
	return &rosterService{
		repository: repository,
	}
}