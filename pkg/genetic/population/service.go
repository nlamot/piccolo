package population

type PopulationService interface {
	Create(CreatePopulationRequest)
}

type populationService struct {
	repository PopulationRepository
}

func (p *populationService) Create(request CreatePopulationRequest) {
	
}

func ProvidePopulationService(repository PopulationRepository) PopulationService{
	return &populationService{
		repository: repository,
	}
}