package avalon_claim

var ABI = `
[
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "user",
                "type": "address",
                "indexed": true
            },
            {
                "internalType": "uint256",
                "name": "avlAmount",
                "type": "uint256",
                "indexed": false
            },
            {
                "internalType": "uint256",
                "name": "usdaAmount",
                "type": "uint256",
                "indexed": false
            }
        ],
        "type": "event",
        "name": "Claimed",
        "anonymous": false
    }
]
`
