/// Mock for external Bitcoin API
/// See:
/// curl "http://localhost:3000/bitcoin-api/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj?details=txs"
/// curl "https://btc1.trezor.io/api/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj?details=txs"
/// curl "http://localhost:8420/v1/bitcoin/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"

module.exports = {
    path: '/bitcoin-api/address/:address?',
    template: function (params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 1,
                    "itemsOnPage": 1000,
                    "addrStr": "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj",
                    "balance": "0.00020074",
                    "totalReceived": "0.00091089",
                    "totalSent": "0.00071015",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxApperances": 0,
                    "txApperances": 2,
                    "txs": [
                      {
                        "txid": "8180545030fcfb7b14ee90acb01b606c5a68fe1d520bcf140bd0097108c2b7f4",
                        "version": 1,
                        "vin": [
                          {
                            "txid": "270a6490cc806d99aa84bb8079b543d9d7a255fc59364890e7a4cb57970daef8",
                            "vout": 1,
                            "sequence": 4294967294,
                            "n": 0,
                            "scriptSig": {},
                            "addresses": [
                              "bc1q3230a7cqt2drewuza8qff4c4gpt4muy9qyqknw"
                            ],
                            "value": "0.00001699"
                          }
                        ],
                        "vout": [
                          {
                            "value": "0.00000375",
                            "n": 0,
                            "scriptPubKey": {
                              "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                              "addresses": [
                                "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                              ]
                            },
                            "spent": false
                          },
                          {
                            "value": "0.00001098",
                            "n": 1,
                            "scriptPubKey": {
                              "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                              "addresses": [
                                "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                              ]
                            },
                            "spent": false
                          }
                        ],
                        "blockhash": "00000000000000000006ff4c09a5e05c4a28dffe54beb7ef0c61f8dd940a951e",
                        "blockheight": 622017,
                        "confirmations": 780,
                        "time": 1584485709,
                        "blocktime": 1584485709,
                        "valueOut": "0.00001473",
                        "valueIn": "0.00001699",
                        "fees": "0.00000226",
                        "hex": "01000000000101f8ae0d9757cba4e790483659fc55a2d7d943b57980bb84aa996d80cc90640a270100000000feffffff0277010000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a4a040000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a024730440220683d9c2b40958fd706623783017300ebbc9b73ce1db571963d9bb40679c3079c02207ee68976e225e1e98d20e66d4d8a498cddf3268d0d7b66827a2ed3ac4330f137012102b7f0cb54c2a6a4da3372fe17d7a07475f575777364633e89dacb39eb1fd5a09c00000000"
                      },
                      {
                        "txid": "c6a4c82d5c7a342796e7d81237ab399918d3205f791ebc40e63501cac28c32be",
                        "version": 1,
                        "vin": [
                          {
                            "txid": "d39d78838d5a3c870aa5acdc9e518945bd6ad86af093b38f85c714f790f28c76",
                            "vout": 1,
                            "n": 0,
                            "scriptSig": {},
                            "addresses": [
                              "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                            ],
                            "value": "0.00006055"
                          }
                        ],
                        "vout": [
                          {
                            "value": "0.00001",
                            "n": 0,
                            "scriptPubKey": {
                              "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                              "addresses": [
                                "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                              ]
                            },
                            "spent": false
                          },
                          {
                            "value": "0.00004829",
                            "n": 1,
                            "scriptPubKey": {
                              "hex": "00141a475acd52ae04da60ab33bf373c9255cea3169a",
                              "addresses": [
                                "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                              ]
                            },
                            "spent": false
                          }
                        ],
                        "blockhash": "00000000000000000006ff4c09a5e05c4a28dffe54beb7ef0c61f8dd940a951e",
                        "blockheight": 622017,
                        "confirmations": 780,
                        "time": 1584485709,
                        "blocktime": 1584485709,
                        "valueOut": "0.00005829",
                        "valueIn": "0.00006055",
                        "fees": "0.00000226",
                        "hex": "01000000000101768cf290f714c7858fb393f06ad86abd4589519edcaca50a873c5a8d83789dd301000000000000000002e8030000000000001600141a475acd52ae04da60ab33bf373c9255cea3169add120000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a02483045022100cb25ac0d9e69ddaaedd49e3a3f9efb218b38bf49c2abf44ff2266bfb941bd3b202206b32cc1337f9a6a176fabadd6aefdf5a95d479e5db856d4e3e8eee7fefc26773012102624729d04d58fa33eb3f7be9fe9c307c60aa1bad52f4ffbacc78ba3808faf62500000000"
                      }
                    ]
                  }
                `);
        }
        return { error: "Not implemented" };
    }
}
