package swing

// ABI holds the ABI for the NewSale event.
var ABI = `
[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "bytes32",
        "name": "id",
        "type": "bytes32"
      },
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "bridgeTransferId",
        "type": "bytes32"
      },
      {
        "indexed": true,
        "internalType": "bytes32",
        "name": "bridge",
        "type": "bytes32"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "from",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "fromToken",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "bridgeToken",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "destToken",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "fromAmount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "bridgeAmount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "dstAmount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "enum DataTypes.SwapStatus",
        "name": "status",
        "type": "uint8"
      }
    ],
    "name": "CrosschainSwapRequest",
    "type": "event"
  }
]
`
