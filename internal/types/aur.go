package types

type AURResponse struct {
	Results []AurPackage `json:"results"`
}

type AurPackage struct {
	Name        string `json:"Name"`
	Version     string `json:"Version"`
	Description string `json:"Description"`
	Maintainer  string `json:"Maintainer"`
	NumVotes    int    `json:"NumVotes"`
	UpstreamURL string `json:"URL"`
}
