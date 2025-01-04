package whitelist

import (
	_ "embed"
	"encoding/json"
)

//go:embed protocols.json
var protocols []byte

type Category string

const (
	Gaming           Category = "Gaming"
	Bridge           Category = "Bridge"
	DEX              Category = "DEX"
	Lending          Category = "Lending"
	DeFi             Category = "DeFi"
	ENS              Category = "ENS"
	NFT              Category = "NFT"
	CrossChain       Category = "Cross-chain"
	Derivatives      Category = "Derivatives"
	BridgeAggregator Category = "Bridge Aggregator"
)

type Protocol struct {
	// Mandatory fields
	// Name of the protocol
	Name string `json:"name"`
	// Slug of the protocol
	Slug string `json:"slug"`

	// Mandatory fields for Trailblazer competitions
	// Contracts holds the protocols' contract addresses
	Contracts []string `json:"contracts"`

	// Optional fields
	// Short description of the protocol
	Description *string `json:"description,omitempty"`
	// Category of the protocol
	Category *Category `json:"category,omitempty"`
	// Twitter reference of the protocol
	Twitter *string `json:"twitter,omitempty"`
	// Website of the protocol
	Website *string `json:"website,omitempty"`
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
