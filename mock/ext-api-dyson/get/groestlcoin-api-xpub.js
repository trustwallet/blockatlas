/// Mock for external Groestlcoin API
/// See:
/// curl "http://{groestlcoin rpc}/v2/xpub/zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf?details=txs"
/// curl "http://localhost:3000/groestlcoin-api/v2/xpub/zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf?details=txs"
/// curl "http://localhost:8420/v1/groestlcoin/xpub/zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf"

module.exports = {
    path: '/groestlcoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf",
                        "balance": "412844353",
                        "totalReceived": "289739972697",
                        "totalSent": "289327128344",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 98,
                        "transactions": [
                            {
                                "txid": "686c651223b937b1223560a60631ad79ad17351c88bba67cad8ea0c95fccbb83",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "1e3362ef26063cae601362721fcde1c0e856b7d81abc9e541ba38a0b15330689",
                                        "sequence": 2147483644,
                                        "n": 0,
                                        "addresses": [
                                            "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04"
                                        ],
                                        "isAddress": true,
                                        "value": "100000000"
                                    },
                                    {
                                        "txid": "1e3362ef26063cae601362721fcde1c0e856b7d81abc9e541ba38a0b15330689",
                                        "vout": 1,
                                        "sequence": 2147483645,
                                        "n": 1,
                                        "addresses": [
                                            "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04"
                                        ],
                                        "isAddress": true,
                                        "value": "212864353"
                                    },
                                    {
                                        "txid": "dd7780ee2529f7030153737c8ee2a16ef32817cbd63b6aba8553c4ccbeac368d",
                                        "sequence": 2147483646,
                                        "n": 2,
                                        "addresses": [
                                            "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04"
                                        ],
                                        "isAddress": true,
                                        "value": "100000000"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "10000000",
                                        "n": 0,
                                        "hex": "0014e9c2ceb5078d184f14f31c6cb1b3f633fdb42395",
                                        "addresses": [
                                            "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "402844353",
                                        "n": 1,
                                        "hex": "0014e9c2ceb5078d184f14f31c6cb1b3f633fdb42395",
                                        "addresses": [
                                            "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "00000000000003a48c30ddfaa9c875b902e8372c448413ee65dba418601a5b7e",
                                "blockHeight": 2998112,
                                "confirmations": 20605,
                                "blockTime": 1583831414,
                                "value": "412844353",
                                "valueIn": "412864353",
                                "fees": "20000",
                                "hex": "01000000000103890633150b8aa31b549ebc1ad8b756e8c0e1cd1f72621360ae3c0626ef62331e0000000000fcffff7f890633150b8aa31b549ebc1ad8b756e8c0e1cd1f72621360ae3c0626ef62331e0100000000fdffff7f8d36acbeccc45385ba6a3bd6cb1728f36ea1e28e7c73530103f72925ee8077dd0000000000feffff7f028096980000000000160014e9c2ceb5078d184f14f31c6cb1b3f633fdb42395c1ea021800000000160014e9c2ceb5078d184f14f31c6cb1b3f633fdb423950247304402201b4acd75cb1186fba1405bd604a2e24a960d94a39e14316f4e6ef5289808e4c90220569d652a91ecc149d543d7aab0653a646ce60d8e3bed053d421907432bd7c9bd012102ab2d3112dfb83d7e08949ec32128e24a4be5c588640eee5028cdba56aae7b285024730440220426733a0a12dd1c34142d9814e74987fdbb41438189fb0bbed74ab2cd420a9b4022046dc588c6690b1ecba351b8a981ed3db63df9de95df49d5a8333b087b349b6c3012102ab2d3112dfb83d7e08949ec32128e24a4be5c588640eee5028cdba56aae7b285024730440220291f4932101719b70630233d53503854f456a0a944453a6105f4e2b95fde526b0220617d4dd0b4dd12e8917d15f80dadf617ec584fe650a31ca8f07f4d6869256549012102ab2d3112dfb83d7e08949ec32128e24a4be5c588640eee5028cdba56aae7b28500000000"
                            },
                            {
                                "txid": "1e3362ef26063cae601362721fcde1c0e856b7d81abc9e541ba38a0b15330689",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "dd7780ee2529f7030153737c8ee2a16ef32817cbd63b6aba8553c4ccbeac368d",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04"
                                        ],
                                        "isAddress": true,
                                        "value": "312884353"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "100000000",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "0014e9c2ceb5078d184f14f31c6cb1b3f633fdb42395",
                                        "addresses": [
                                            "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "212864353",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "0014e9c2ceb5078d184f14f31c6cb1b3f633fdb42395",
                                        "addresses": [
                                            "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "000000000000034560a3a07b6498c9b937da914a4f683d1575197a02c4b69a83",
                                "blockHeight": 2989928,
                                "confirmations": 28789,
                                "blockTime": 1583315379,
                                "value": "312864353",
                                "valueIn": "312884353",
                                "fees": "20000",
                                "hex": "010000000001018d36acbeccc45385ba6a3bd6cb1728f36ea1e28e7c73530103f72925ee8077dd0100000000000000000200e1f50500000000160014e9c2ceb5078d184f14f31c6cb1b3f633fdb42395610db00c00000000160014e9c2ceb5078d184f14f31c6cb1b3f633fdb4239502483045022100e0e7e5484f9d06120f0e3a63471d353d90b65ba63b0fbf640c714608414ffa5d02204556b0c2554bc33ebfb338d806fa1589df1e25afe251c5a7869cfad6e6ef0a61012102ab2d3112dfb83d7e08949ec32128e24a4be5c588640eee5028cdba56aae7b28500000000"
                            }
                        ],
                        "usedTokens": 79,
                        "tokens": [
                            {
                                "type": "XPUBAddress",
                                "name": "grs1qa8pvadg835vy798nr3ktrvlkx07mggu4hqvv04",
                                "path": "m/84'/17'/0'/0/0",
                                "transfers": 53,
                                "decimals": 8,
                                "balance": "412844353",
                                "totalReceived": "34760610350",
                                "totalSent": "34347765997"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
