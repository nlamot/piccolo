package population

import "piccolo.com/planner/pkg/common/gcp"

type PopulationRepository interface {
}

type populationRepository struct {
	firestoreClient gcp.FirestoreClient
}

func ProvidePopulationRepository(firestoreClient gcp.FirestoreClient) PopulationRepository{
	return &populationRepository{
		firestoreClient: firestoreClient,
	}
}