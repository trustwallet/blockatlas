/// Mock for external Bitcoin API
/// See:
/// curl "https://btc1.trezor.io/api/v2/xpub/zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC?details=txs"
/// curl "http://localhost:3000/bitcoin-api/v2/xpub/zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC?details=txs"
/// curl "http://localhost:8420/v1/bitcoin/xpub/zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC"

module.exports = {
    path: '/bitcoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC",
                        "balance": "37093",
                        "totalReceived": "4218808",
                        "totalSent": "4181715",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 92,
                        "transactions": [
                            {
                                "txid": "c6a4c82d5c7a342796e7d81237ab399918d3205f791ebc40e63501cac28c32be",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "d39d78838d5a3c870aa5acdc9e518945bd6ad86af093b38f85c714f790f28c76",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                                        ],
                                        "isAddress": true,
                                        "value": "6055"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1000",
                                        "n": 0,
                                        "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                                        "addresses": [
                                            "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "4829",
                                        "n": 1,
                                        "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                                        "addresses": [
                                            "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "00000000000000000006ff4c09a5e05c4a28dffe54beb7ef0c61f8dd940a951e",
                                "blockHeight": 622017,
                                "confirmations": 780,
                                "blockTime": 1584485709,
                                "value": "5829",
                                "valueIn": "6055",
                                "fees": "226",
                                "hex": "01000000000101768cf290f714c7858fb393f06ad86abd4589519edcaca50a873c5a8d83789dd301000000000000000002e8030000000000001600141a475acd52ae04da60ab33bf373c9255cea3169add120000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a02483045022100cb25ac0d9e69ddaaedd49e3a3f9efb218b38bf49c2abf44ff2266bfb941bd3b202206b32cc1337f9a6a176fabadd6aefdf5a95d479e5db856d4e3e8eee7fefc26773012102624729d04d58fa33eb3f7be9fe9c307c60aa1bad52f4ffbacc78ba3808faf62500000000"
                            },
                            {
                                "txid": "8180545030fcfb7b14ee90acb01b606c5a68fe1d520bcf140bd0097108c2b7f4",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "270a6490cc806d99aa84bb8079b543d9d7a255fc59364890e7a4cb57970daef8",
                                        "vout": 1,
                                        "sequence": 4294967294,
                                        "n": 0,
                                        "addresses": [
                                            "bc1q3230a7cqt2drewuza8qff4c4gpt4muy9qyqknw"
                                        ],
                                        "isAddress": true,
                                        "value": "1699"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "375",
                                        "n": 0,
                                        "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                                        "addresses": [
                                            "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "1098",
                                        "n": 1,
                                        "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                                        "addresses": [
                                            "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "00000000000000000006ff4c09a5e05c4a28dffe54beb7ef0c61f8dd940a951e",
                                "blockHeight": 622017,
                                "confirmations": 780,
                                "blockTime": 1584485709,
                                "value": "1473",
                                "valueIn": "1699",
                                "fees": "226",
                                "hex": "01000000000101f8ae0d9757cba4e790483659fc55a2d7d943b57980bb84aa996d80cc90640a270100000000feffffff0277010000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a4a040000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a024730440220683d9c2b40958fd706623783017300ebbc9b73ce1db571963d9bb40679c3079c02207ee68976e225e1e98d20e66d4d8a498cddf3268d0d7b66827a2ed3ac4330f137012102b7f0cb54c2a6a4da3372fe17d7a07475f575777364633e89dacb39eb1fd5a09c00000000"
                            },
                            {
                                "txid": "d39d78838d5a3c870aa5acdc9e518945bd6ad86af093b38f85c714f790f28c76",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "c215e9c6f9533d584effc4d31d08736a0ef4f717b7d42563ebba053f5f3bc1d5",
                                        "vout": 1,
                                        "sequence": 4294967287,
                                        "n": 0,
                                        "addresses": [
                                            "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                                        ],
                                        "isAddress": true,
                                        "value": "8185"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1000",
                                        "n": 0,
                                        "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                                        "addresses": [
                                            "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "6055",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                                        "addresses": [
                                            "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "0000000000000000000ab8b2230f5b63d4d4be1a0aa6b9296cfc09e48c8e42e0",
                                "blockHeight": 619385,
                                "confirmations": 3412,
                                "blockTime": 1582902175,
                                "value": "7055",
                                "valueIn": "8185",
                                "fees": "1130",
                                "hex": "01000000000101d5c13b5f3f05baeb6325d4b717f7f40e6a73081dd3c4ff4e583d53f9c6e915c20100000000f7ffffff02e8030000000000001600141a475acd52ae04da60ab33bf373c9255cea3169aa7170000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a024730440220158b0d8c5fb6cb9f60d0af05ce8b63abe934dca8eca380013bc7e8463420810f02207d7e6594586ec8196546e2520b6775c2e8659c8f266acb3b861e99f13cf1b380012102624729d04d58fa33eb3f7be9fe9c307c60aa1bad52f4ffbacc78ba3808faf62500000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
