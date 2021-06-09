package population

import "piccolo.com/planner/pkg/common/gcp"

//go:generate go run github.com/vektra/mockery/cmd/mockery -name PopulationRepository -output ./mock/ -outpkg mock
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