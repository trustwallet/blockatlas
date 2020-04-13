/// Ethereum API Mock
/// See:
/// curl "http://localhost:3000/eth-api/transactions?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:3000/eth-api/tokens?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "https://{eth rpc}/transactions?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "https://{eth rpc}/tokens?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:8420/v1/ethereum/0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:8420/v2/ethereum/tokens/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?Authorization=Bearer"

module.exports = {
    path: '/eth-api/:command1?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.command1) {
            case 'transactions':
                if (query.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
                    return JSON.parse(`{
                        "docs": [
                            {
                                "operations": [],
                                "contract": null,
                                "_id": "0x2a9fd94735e273526a2cde57c6a19b9d488e0c9a960565c3e19be5e12d4b4b47",
                                "blockNumber": 9551915,
                                "time": 1582624428,
                                "nonce": 227,
                                "from": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                                "to": "0x1717f94202c126ef71d6C562de253Fe95eEbDD5f",
                                "value": "17635730000000000",
                                "gas": "21000",
                                "gasPrice": "4320000000",
                                "gasUsed": "21000",
                                "input": "0x",
                                "error": "",
                                "id": "0x2a9fd94735e273526a2cde57c6a19b9d488e0c9a960565c3e19be5e12d4b4b47",
                                "timeStamp": "1582624428"
                            },
                            {
                                "operations": [
                                    {
                                        "transactionId": "0xb669b69afee75c6ef073a603600041d3708d54da8d43cab7b35ee66baa7510d3-0",
                                        "contract": {
                                            "address": "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
                                            "symbol": "BAT",
                                            "decimals": 18,
                                            "totalSupply": "1500000000000000000000000000",
                                            "name": "Basic Attention Token",
                                            "updatedAt": "2020-03-23T06:01:25.975Z"
                                        },
                                        "from": "0xeCe114137b2e9Dbf29712BDC39639EB0B72B41b8",
                                        "to": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                                        "type": "token_transfer",
                                        "value": "400000000000000000",
                                        "id": null
                                    }
                                ],
                                "contract": null,
                                "_id": "0xb669b69afee75c6ef073a603600041d3708d54da8d43cab7b35ee66baa7510d3",
                                "blockNumber": 9519169,
                                "time": 1582189159,
                                "nonce": 16,
                                "from": "0xeCe114137b2e9Dbf29712BDC39639EB0B72B41b8",
                                "to": "0x0D8775F648430679A709E98d2b0Cb6250d2887EF",
                                "value": "0",
                                "gas": "51839",
                                "gasPrice": "11500000000",
                                "gasUsed": "37028",
                                "input": "0xa9059cbb0000000000000000000000000875bcab22de3d02402bc38aee4104e1239374a7000000000000000000000000000000000000000000000000058d15e176280000",
                                "error": "",
                                "id": "0xb669b69afee75c6ef073a603600041d3708d54da8d43cab7b35ee66baa7510d3",
                                "timeStamp": "1582189159"
                            }
                        ],
                        "total": 2
                    }`);
                }

            case 'tokens':
                if (query.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
                    return JSON.parse(`
                        {
                            "total": 17,
                            "docs": [
                                {
                                    "address": "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
                                    "name": "Kyber Network Crystal",
                                    "decimals": 18,
                                    "symbol": "KNC"
                                },
                                {
                                    "address": "0xB8c77482e45F1F44dE1745F52C74426C631bDD52",
                                    "name": "BNB",
                                    "decimals": 18,
                                    "symbol": "BNB"
                                },
                                {
                                    "address": "0x0D8775F648430679A709E98d2b0Cb6250d2887EF",
                                    "name": "Basic Attention Token",
                                    "decimals": 18,
                                    "symbol": "BAT"
                                },
                                {
                                    "address": "0x226bb599a12C826476e3A771454697EA52E9E220",
                                    "name": "Propy",
                                    "decimals": 8,
                                    "symbol": "PRO"
                                },
                                {
                                    "address": "0xf3Db5Fa2C66B7aF3Eb0C0b782510816cbe4813b8",
                                    "name": "Everex",
                                    "decimals": 4,
                                    "symbol": "EVX"
                                },
                                {
                                    "address": "0x85e076361cc813A908Ff672F9BAd1541474402b2",
                                    "name": "Telcoin",
                                    "decimals": 2,
                                    "symbol": "TEL"
                                },
                                {
                                    "address": "0xD73bE539d6B2076BaB83CA6Ba62DfE189aBC6Bbe",
                                    "name": "BlockchainCuties",
                                    "decimals": 0,
                                    "symbol": "BC"
                                },
                                {
                                    "address": "0x0000000000085d4780B73119b644AE5ecd22b376",
                                    "name": "TrueUSD",
                                    "decimals": 18,
                                    "symbol": "TUSD"
                                },
                                {
                                    "address": "0xFBeef911Dc5821886e1dda71586d90eD28174B7d",
                                    "name": "KnownOriginDigitalAsset",
                                    "decimals": 0,
                                    "symbol": "KODA"
                                },
                                {
                                    "address": "0xc3761EB917CD790B30dAD99f6Cc5b4Ff93C4F9eA",
                                    "name": "ERC20",
                                    "decimals": 18,
                                    "symbol": "ERC20"
                                },
                                {
                                    "address": "0x77FE30b2cf39245267C0a5084B66a560f1cF9E1f",
                                    "name": "Azbit",
                                    "decimals": 18,
                                    "symbol": "AZ"
                                },
                                {
                                    "address": "0x7f3EaB3491Ed282197038F1B89CA33D7e5ADffBa",
                                    "name": "Coin-coin coinslot.com",
                                    "decimals": 8,
                                    "symbol": "CC coinslot.com"
                                },
                                {
                                    "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                                    "name": "Tether USD",
                                    "decimals": 6,
                                    "symbol": "USDT"
                                },
                                {
                                    "address": "0xE1Ac9Eb7cDDAbfd9e5CA49c23bd521aFcDF8BE49",
                                    "name": "Mycion",
                                    "decimals": 18,
                                    "symbol": "MYC"
                                },
                                {
                                    "address": "0xF629cBd94d3791C9250152BD8dfBDF380E2a3B9c",
                                    "name": "Enjin Coin",
                                    "decimals": 18,
                                    "symbol": "ENJ"
                                },
                                {
                                    "address": "0x467Bccd9d29f223BcE8043b84E8C8B282827790F",
                                    "name": "Telcoin",
                                    "decimals": 2,
                                    "symbol": "TEL"
                                },
                                {
                                    "address": "0xC12D1c73eE7DC3615BA4e37E4ABFdbDDFA38907E",
                                    "name": "KickToken",
                                    "decimals": 8,
                                    "symbol": "KICK"
                                }
                            ]
                        }
                    `);
                }
        }

        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
