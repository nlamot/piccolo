package population

import "piccolo.com/planner/pkg/common/gcp/firestore"

//go:generate go run github.com/vektra/mockery/cmd/mockery -name PopulationRepository -output ./mock/ -outpkg mock
type PopulationRepository interface {
}

type populationRepository struct {
	firestoreClient firestore.FirestoreClient
}

func ProvidePopulationRepository(firestoreClient firestore.FirestoreClient) PopulationRepository{
	return &populationRepository{
		firestoreClient: firestoreClient,
	}
}