package roster

import "piccolo.com/planner/pkg/common/gcp"

type RosterRepository interface {
	Create(Roster) string
}

type rosterRepository struct {
	firestoreClient gcp.FirestoreClient
}

func (r *rosterRepository) Create(roster Roster) string {
	return ""
}

func ProvideRosterRepository(firestoreClient gcp.FirestoreClient) RosterRepository {
	return &rosterRepository{
		firestoreClient: firestoreClient,
	}
}