package roster

//go:generate mockery --name RosterService --output ./mock/ --outpkg mock
type RosterService interface {
	// Create creates the new roster and returns the UUID
	Create(string, Roster) (string, error)
}

type rosterService struct {
	repository RosterRepository
}

func (s *rosterService) Create(organisationUUID string, roster Roster) (string, error) {
	return s.repository.Create(organisationUUID, roster)
}

func ProvideRosterService(repository RosterRepository) RosterService {
	return &rosterService{
		repository: repository,
	}
}
