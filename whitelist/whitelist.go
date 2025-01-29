package whitelist

import (
	_ "embed"
	"encoding/json"
)

//go:embed protocols.json
var protocols []byte

const (
	Gaming              string = "Gaming"
	Bridge              string = "Bridge"
	DEX                 string = "DEX"
	Lending             string = "Lending"
	DeFi                string = "DeFi"
	ENS                 string = "ENS"
	NFT                 string = "NFT"
	CrossChain          string = "Cross-chain"
	Derivatives         string = "Derivatives"
	BridgeAggregator    string = "Bridge Aggregator"
	SocialFi            string = "SocialFi"
	Analytics           string = "Analytics"
	AI                  string = "AI"
	PredictionMarket    string = "Prediction Market"
	LiquidityManagement string = "Liquidity Management"
	Staking             string = "Staking"
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
	Category *string `json:"category,omitempty"`
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
