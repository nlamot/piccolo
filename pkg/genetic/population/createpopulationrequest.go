package population

type CreatePopulationRequest struct {
	Size int `json:"size"`
	RosterID string `json:"rosterID"`
	InternIDs []string `json:"internIDs"`
}
