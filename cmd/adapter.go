package cmd

type adapter string

func (a adapter) String() string {
	return string(a)
}

const (
	IzumiLP                  adapter = "IzumiLP"
	RitsuLP                  adapter = "RitsuLP"
	TransactionSender        adapter = "TransactionSender"
	NftDeployed              adapter = "NftDeployed"
	GamingWhitelist          adapter = "GamingWhitelist"
	DotTaikoDomains          adapter = "DotTaikoDomains"
	OkxOrderFulfilled        adapter = "OkxOrderFulfilled"
	LoopexNewSale            adapter = "LoopexNewSale"
	OmnihubContractDeployed  adapter = "OmnihubContractDeployed"
	Nfts2meCollectionCreated adapter = "Nfts2meCollectionCreated"
	ConftTokenSold           adapter = "ConftTokenSold"
	DripsLock                adapter = "DripsLock"
)

func adapterz() []adapter {
	return []adapter{
		IzumiLP,
		RitsuLP,
		TransactionSender,
		NftDeployed,
		GamingWhitelist,
		DotTaikoDomains,
		OkxOrderFulfilled,
		LoopexNewSale,
		OmnihubContractDeployed,
		Nfts2meCollectionCreated,
		ConftTokenSold,
		DripsLock,
	}
}
