package magpie

// ABI holds the ABI for the OrderFulfilled event.
var ABI = `
[
    {
        "name": "ApprovalFailed",
        "type": "error",
        "inputs": []
    },
    {
        "name": "ExpiredTransaction",
        "type": "error",
        "inputs": []
    },
    {
        "name": "InsufficientAmountOut",
        "type": "error",
        "inputs": []
    },
    {
        "name": "InvalidCall",
        "type": "error",
        "inputs": []
    },
    {
        "name": "InvalidCommand",
        "type": "error",
        "inputs": []
    },
    {
        "name": "InvalidSelector",
        "type": "error",
        "inputs": []
    },
    {
        "name": "InvalidSequenceType",
        "type": "error",
        "inputs": []
    },
    {
        "name": "InvalidSequencesLength",
        "type": "error",
        "inputs": []
    },
    {
        "name": "InvalidTransferFrom",
        "type": "error",
        "inputs": []
    },
    {
        "name": "TransferFailed",
        "type": "error",
        "inputs": []
    },
    {
        "name": "TransferFromFailed",
        "type": "error",
        "inputs": []
    },
    {
        "name": "UniswapV3InvalidAmount",
        "type": "error",
        "inputs": []
    },
    {
        "name": "OwnershipTransferStarted",
        "type": "event",
        "inputs": [
            {
                "name": "previousOwner",
                "type": "address",
                "indexed": true,
                "internalType": "address"
            },
            {
                "name": "newOwner",
                "type": "address",
                "indexed": true,
                "internalType": "address"
            }
        ],
        "anonymous": false
    },
    {
        "name": "OwnershipTransferred",
        "type": "event",
        "inputs": [
            {
                "name": "previousOwner",
                "type": "address",
                "indexed": true,
                "internalType": "address"
            },
            {
                "name": "newOwner",
                "type": "address",
                "indexed": true,
                "internalType": "address"
            }
        ],
        "anonymous": false
    },
    {
        "name": "Swap",
        "type": "event",
        "inputs": [
            {
                "name": "fromAddress",
                "type": "address",
                "indexed": true,
                "internalType": "address"
            },
            {
                "name": "toAddress",
                "type": "address",
                "indexed": true,
                "internalType": "address"
            },
            {
                "name": "fromAssetAddress",
                "type": "address",
                "indexed": false,
                "internalType": "address"
            },
            {
                "name": "toAssetAddress",
                "type": "address",
                "indexed": false,
                "internalType": "address"
            },
            {
                "name": "amountIn",
                "type": "uint256",
                "indexed": false,
                "internalType": "uint256"
            },
            {
                "name": "amountOut",
                "type": "uint256",
                "indexed": false,
                "internalType": "uint256"
            }
        ],
        "anonymous": false
    },
    {
        "type": "fallback",
        "stateMutability": "nonpayable"
    },
    {
        "name": "acceptOwnership",
        "type": "function",
        "inputs": [],
        "outputs": [],
        "stateMutability": "nonpayable"
    },
    {
        "name": "estimateSwapGas",
        "type": "function",
        "inputs": [
            {
                "name": "",
                "type": "bytes",
                "internalType": "bytes"
            }
        ],
        "outputs": [
            {
                "name": "amountOut",
                "type": "uint256",
                "internalType": "uint256"
            },
            {
                "name": "gasUsed",
                "type": "uint256",
                "internalType": "uint256"
            }
        ],
        "stateMutability": "payable"
    },
    {
        "name": "multicall",
        "type": "function",
        "inputs": [
            {
                "name": "data",
                "type": "bytes[]",
                "internalType": "bytes[]"
            }
        ],
        "outputs": [
            {
                "name": "results",
                "type": "bytes[]",
                "internalType": "bytes[]"
            }
        ],
        "stateMutability": "nonpayable"
    },
    {
        "name": "owner",
        "type": "function",
        "inputs": [],
        "outputs": [
            {
                "name": "",
                "type": "address",
                "internalType": "address"
            }
        ],
        "stateMutability": "view"
    },
    {
        "name": "pendingOwner",
        "type": "function",
        "inputs": [],
        "outputs": [
            {
                "name": "",
                "type": "address",
                "internalType": "address"
            }
        ],
        "stateMutability": "view"
    },
    {
        "name": "renounceOwnership",
        "type": "function",
        "inputs": [],
        "outputs": [],
        "stateMutability": "nonpayable"
    },
    {
        "name": "silentSwap",
        "type": "function",
        "inputs": [
            {
                "name": "",
                "type": "bytes",
                "internalType": "bytes"
            }
        ],
        "outputs": [
            {
                "name": "amountOut",
                "type": "uint256",
                "internalType": "uint256"
            }
        ],
        "stateMutability": "payable"
    },
    {
        "name": "swap",
        "type": "function",
        "inputs": [
            {
                "name": "",
                "type": "bytes",
                "internalType": "bytes"
            }
        ],
        "outputs": [
            {
                "name": "amountOut",
                "type": "uint256",
                "internalType": "uint256"
            }
        ],
        "stateMutability": "payable"
    },
    {
        "name": "transferOwnership",
        "type": "function",
        "inputs": [
            {
                "name": "newOwner",
                "type": "address",
                "internalType": "address"
            }
        ],
        "outputs": [],
        "stateMutability": "nonpayable"
    },
    {
        "name": "updateSelector",
        "type": "function",
        "inputs": [
            {
                "name": "commandType",
                "type": "uint16",
                "internalType": "uint16"
            },
            {
                "name": "selector",
                "type": "bytes4",
                "internalType": "bytes4"
            }
        ],
        "outputs": [],
        "stateMutability": "nonpayable"
    },
    {
        "type": "receive",
        "stateMutability": "payable"
    }
]`