package types

type AURPackage struct {
	Name        string `json:"Name"`
	Version     string `json:"Version"`
	Description string `json:"Description"`
	Maintainer  string `json:"Maintainer"`
	NumVotes    int    `json:"NumVotes"`
	UpstreamURL string `json:"URL"`
}

type AURResponse struct {
	Results []AURPackage `json:"results"`
}
