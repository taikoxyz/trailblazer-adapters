package whitelist

import (
	_ "embed"
	"encoding/json"
)

//go:embed protocols.json
var protocols []byte

type Protocol struct {
	// Name of the protocol
	Name string `json:"name"`
	// Slug of the protocol
	Slug string `json:"slug"`
	// Contracts holds the protocols' contract addresses
	Contracts []string `json:"contracts"`
	// Twitter reference of the protocol
	Twitter *string `json:"twitter,omitempty"`
	// Logo of the protocol
	Logo *string `json:"logo,omitempty"`
}

// Protocols that are whitelisted
func Protocols() ([]Protocol, error) {
	var ps []Protocol

	err := json.Unmarshal(protocols, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
