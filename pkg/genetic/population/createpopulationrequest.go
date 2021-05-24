package population

type GeneratePopulationRequest struct {
	Size      int      `json:"size"`
	RosterID  string   `json:"rosterID"`
	InternIDs []string `json:"internIDs"`
}
