package population

type PopulationService interface {
	Generate(GeneratePopulationRequest)
}

type populationService struct {
	repository PopulationRepository
}

func (p *populationService) Generate(request GeneratePopulationRequest) {

}

func ProvidePopulationService(repository PopulationRepository) PopulationService {
	return &populationService{
		repository: repository,
	}
}
