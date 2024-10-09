package cmd

type adapter string

func (a adapter) String() string {
	return string(a)
}

const (
	Izumi                 adapter = "Izumi"
	RitsuLP               adapter = "RitsuLP"
	NewTransactionSender  adapter = "NewTransactionSender"
	NftDeployed           adapter = "NftDeployed"
	GamingWhitelist       adapter = "GamingWhitelist"
	DotTaikoIndexer       adapter = "DotTaikoIndexer"
	OrderFulfilledIndexer adapter = "OrderFulfilledIndexer"
	NewSaleIndexer        adapter = "NewSaleIndexer"
	ContractDeployed      adapter = "ContractDeployed"
	CollectionCreated     adapter = "CollectionCreated"
	TokenSold             adapter = "TokenSold"
	Drips                 adapter = "Drips"
)

func adapterz() []adapter {
	return []adapter{
		Izumi,
		RitsuLP,
		NewTransactionSender,
		NftDeployed,
		GamingWhitelist,
		DotTaikoIndexer,
		OrderFulfilledIndexer,
		NewSaleIndexer,
		ContractDeployed,
		CollectionCreated,
		TokenSold,
		Drips,
	}
}
