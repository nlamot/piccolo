package population

//go:generate go run github.com/vektra/mockery/cmd/mockery -name PopulationService -output ./mock/ -outpkg mock
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
