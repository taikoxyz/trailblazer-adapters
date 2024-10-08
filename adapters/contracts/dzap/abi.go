package dzap

var ABI = `[
	{
	"anonymous": false,
	"name": "MultiSwapped",
	"inputs": [
		{ "type": "bytes32", "name": "transactionId" },
		{ "type": "address", "name": "integrator", "indexed": true },
		{ "type": "address", "name": "sender", "indexed": true },
		{ "type": "address", "name": "recipient" },
		{
		"type": "tuple[]",
		"name": "swapInfo",
		"components": [
			{ "type": "address", "name": "dex", "indexed": false },
			{ "type": "address", "name": "fromToken", "indexed": false },
			{ "type": "address", "name": "toToken", "indexed": false },
			{ "type": "uint256", "name": "fromAmount", "indexed": false },
			{ "type": "uint256", "name": "leftOverFromAmount", "indexed": false },
			{ "type": "uint256", "name": "returnToAmount", "indexed": false }
		]
		}
	]
	},
	{
	"anonymous": false,
	"name": "Swapped",
	"inputs": [
		{ "type": "bytes32", "name": "transactionId" },
		{ "type": "address", "name": "integrator", "indexed": true },
		{ "type": "address", "name": "sender", "indexed": true },
		{ "type": "address", "name": "recipient" },
		{
		"type": "tuple",
		"name": "swapInfo",
		"components": [
			{ "type": "address", "name": "dex", "indexed": false },
			{ "type": "address", "name": "fromToken", "indexed": false },
			{ "type": "address", "name": "toToken", "indexed": false },
			{ "type": "uint256", "name": "fromAmount", "indexed": false },
			{ "type": "uint256", "name": "leftOverFromAmount", "indexed": false },
			{ "type": "uint256", "name": "returnToAmount", "indexed": false }
		]
		}
	]
	},
	{
	"anonymous": false,
	"name": "SwappedSingleToken",
	"inputs": [
		{ "type": "bytes32", "name": "transactionId" },
		{ "type": "address", "name": "sender", "indexed": true },
		{ "type": "address", "name": "recipient" },
		{
		"type": "tuple",
		"name": "swapInfo",
		"components": [
			{ "type": "address", "name": "dex", "indexed": false },
			{ "type": "address", "name": "fromToken", "indexed": false },
			{ "type": "address", "name": "toToken", "indexed": false },
			{ "type": "uint256", "name": "fromAmount", "indexed": false },
			{ "type": "uint256", "name": "leftOverFromAmount", "indexed": false },
			{ "type": "uint256", "name": "returnToAmount", "indexed": false }
		]
		}
	]
	},
]
`
