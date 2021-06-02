package roster

import (
	"context"

	"piccolo.com/planner/pkg/common/gcp"
)

type RosterRepository interface {
	Create(string, Roster) (string, error)
}

type rosterRepository struct {
	firestoreClient gcp.FirestoreClient
}

func (r *rosterRepository) Create(organisationUUID string, roster Roster) (string, error) {
	rosters := r.firestoreClient.Collection("organisation-data/" + organisationUUID + "/rosters")
	doc, _, err := rosters.Add(context.Background(), roster)
	return doc.ID(), err
}

func ProvideRosterRepository(firestoreClient gcp.FirestoreClient) RosterRepository {
	return &rosterRepository{
		firestoreClient: firestoreClient,
	}
}
