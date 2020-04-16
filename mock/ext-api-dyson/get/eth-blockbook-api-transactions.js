/// Ethereum Blockbook API Mock
/// See:
/// curl "http://localhost:3000/eth-blockbook-api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?details=txs"
/// curl "http://localhost:3000/eth-blockbook-api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?details=tokens"
/// curl "https://{eth blockbook api}/api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?details=txs"
/// curl "https://{eth blockbook api}/api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?details=tokens"
/// curl "http://localhost:8420/v1/ethereum/0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:8420/v2/ethereum/tokens/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?Authorization=Bearer"

module.exports = {
    path: '/eth-blockbook-api/v2/address/:address?',
    template: function(params, query, body) {
        //console.log(params);
        //console.log(query);
        if (params.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
            if (query.details === 'tokens') {
                return JSON.parse(`
                    {
                        "address": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                        "balance": "182976771756327797",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 263,
                        "nonTokenTxs": 243,
                        "nonce": "231",
                        "tokens": [
                            {
                                "type": "ERC20",
                                "name": "Kyber Network Crystal",
                                "contract": "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
                                "transfers": 13,
                                "symbol": "KNC",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "BNB",
                                "contract": "0xB8c77482e45F1F44dE1745F52C74426C631bDD52",
                                "transfers": 4,
                                "symbol": "BNB",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "Basic Attention Token",
                                "contract": "0x0D8775F648430679A709E98d2b0Cb6250d2887EF",
                                "transfers": 40,
                                "symbol": "BAT",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "Propy",
                                "contract": "0x226bb599a12C826476e3A771454697EA52E9E220",
                                "transfers": 1,
                                "symbol": "PRO",
                                "decimals": 8
                            },
                            {
                                "type": "ERC20",
                                "name": "Everex",
                                "contract": "0xf3Db5Fa2C66B7aF3Eb0C0b782510816cbe4813b8",
                                "transfers": 2,
                                "symbol": "EVX",
                                "decimals": 4
                            },
                            {
                                "type": "ERC20",
                                "name": "Telcoin",
                                "contract": "0x85e076361cc813A908Ff672F9BAd1541474402b2",
                                "transfers": 61,
                                "symbol": "TEL",
                                "decimals": 2
                            },
                            {
                                "type": "ERC20",
                                "name": "blockwell.ai KYC Casper Token",
                                "contract": "0x212D95FcCdF0366343350f486bda1ceAfC0C2d63",
                                "transfers": 1,
                                "symbol": "blockwell.ai KYC Casper Token",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "VeChain Token",
                                "contract": "0xD850942eF8811f2A866692A623011bDE52a462C1",
                                "transfers": 1,
                                "symbol": "VEN",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "BlockchainCuties",
                                "contract": "0xD73bE539d6B2076BaB83CA6Ba62DfE189aBC6Bbe",
                                "transfers": 2,
                                "symbol": "BC",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "TrueUSD",
                                "contract": "0x0000000000085d4780B73119b644AE5ecd22b376",
                                "transfers": 5,
                                "symbol": "TUSD",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "ERC20",
                                "contract": "0xc3761EB917CD790B30dAD99f6Cc5b4Ff93C4F9eA",
                                "transfers": 1,
                                "symbol": "ERC20",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "Azbit",
                                "contract": "0x77FE30b2cf39245267C0a5084B66a560f1cF9E1f",
                                "transfers": 1,
                                "symbol": "AZ",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "Coin-coin coinslot.com",
                                "contract": "0x7f3EaB3491Ed282197038F1B89CA33D7e5ADffBa",
                                "transfers": 1,
                                "symbol": "CC coinslot.com",
                                "decimals": 8
                            },
                            {
                                "type": "ERC20",
                                "name": "Tether USD",
                                "contract": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                                "transfers": 2,
                                "symbol": "USDT",
                                "decimals": 6
                            },
                            {
                                "type": "ERC20",
                                "name": "Mycion",
                                "contract": "0xE1Ac9Eb7cDDAbfd9e5CA49c23bd521aFcDF8BE49",
                                "transfers": 1,
                                "symbol": "MYC",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "Enjin Coin",
                                "contract": "0xF629cBd94d3791C9250152BD8dfBDF380E2a3B9c",
                                "transfers": 1,
                                "symbol": "ENJ",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "Telcoin",
                                "contract": "0x467Bccd9d29f223BcE8043b84E8C8B282827790F",
                                "transfers": 1,
                                "symbol": "TEL",
                                "decimals": 2
                            },
                            {
                                "type": "ERC20",
                                "name": "KickToken",
                                "contract": "0xC12D1c73eE7DC3615BA4e37E4ABFdbDDFA38907E",
                                "transfers": 1,
                                "symbol": "KICK",
                                "decimals": 8
                            },
                            {
                                "type": "ERC20",
                                "name": "OMGToken",
                                "contract": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
                                "transfers": 2,
                                "symbol": "OMG",
                                "decimals": 18
                            },
                            {
                                "type": "ERC20",
                                "name": "Minereum",
                                "contract": "0xc92e74b131D7b1D46E60e07F3FaE5d8877Dd03F0",
                                "transfers": 1,
                                "symbol": "MNE",
                                "decimals": 8
                            }
                        ]
                    }
                `);
            }
            return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 11,
                    "itemsOnPage": 25,
                    "address": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                    "balance": "185659745674589722",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 259,
                    "nonTokenTxs": 240,
                    "transactions": [
                        {
                            "txid": "0x2a9fd94735e273526a2cde57c6a19b9d488e0c9a960565c3e19be5e12d4b4b47",
                            "vin": [
                                {
                                    "n": 0,
                                    "addresses": [
                                        "0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "vout": [
                                {
                                    "value": "17635730000000000",
                                    "n": 0,
                                    "addresses": [
                                        "0x1717f94202c126ef71d6C562de253Fe95eEbDD5f"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "0x5b8b39099d025a8feeed7575ebff6d0b2508e0542d6279c1b3b4f8b893e74296",
                            "blockHeight": 9551915,
                            "confirmations": 231227,
                            "blockTime": 1582624428,
                            "value": "17635730000000000",
                            "fees": "90720000000000",
                            "ethereumSpecific": {
                                "status": 1,
                                "nonce": 227,
                                "gasLimit": 21000,
                                "gasUsed": 21000,
                                "gasPrice": "4320000000"
                            }
                        },
                        {
                            "txid": "0xb669b69afee75c6ef073a603600041d3708d54da8d43cab7b35ee66baa7510d3",
                            "vin": [
                                {
                                    "n": 0,
                                    "addresses": [
                                        "0xeCe114137b2e9Dbf29712BDC39639EB0B72B41b8"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "vout": [
                                {
                                    "value": "0",
                                    "n": 0,
                                    "addresses": [
                                        "0x0D8775F648430679A709E98d2b0Cb6250d2887EF"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "0x770332c9b4ea42e518f8c5a0178ca5732fd5edfe5e8e011e1362e6598de262d5",
                            "blockHeight": 9519169,
                            "confirmations": 263973,
                            "blockTime": 1582189159,
                            "value": "0",
                            "fees": "425822000000000",
                            "tokenTransfers": [
                                {
                                    "type": "ERC20",
                                    "from": "0xeCe114137b2e9Dbf29712BDC39639EB0B72B41b8",
                                    "to": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                                    "token": "0x0D8775F648430679A709E98d2b0Cb6250d2887EF",
                                    "name": "Basic Attention Token",
                                    "symbol": "BAT",
                                    "decimals": 18,
                                    "value": "400000000000000000"
                                }
                            ],
                            "ethereumSpecific": {
                                "status": 1,
                                "nonce": 16,
                                "gasLimit": 51839,
                                "gasUsed": 37028,
                                "gasPrice": "11500000000"
                            }
                        }
                    ],
                    "nonce": "228",
                    "tokens": []
                }
            `);
        }

        return {error: "Not implemented"};
    }
};
