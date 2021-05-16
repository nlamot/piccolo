package population

type PopulationService interface {
}

type populationService struct {
	repository PopulationRepository
}

func ProvidePopulationService(repository PopulationRepository) PopulationService{
	return &populationService{
		repository: repository,
	}
}