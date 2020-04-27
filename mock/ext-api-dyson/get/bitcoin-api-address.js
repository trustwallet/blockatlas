/// Mock for external Bitcoin API
/// See:
/// curl "http://localhost:3347/bitcoin-api/v2/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj?details=txs"
/// curl "https://btc1.trezor.io/api/v2/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj?details=txs"
/// curl "http://localhost:8437/v1/bitcoin/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"

module.exports = {
    path: '/bitcoin-api/v2/address/:address?',
    template: function (params, query, body) {
        switch (params.address) {
            case 'bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj':
                return JSON.parse(`
                {
                  "page": 1,
                  "totalPages": 1,
                  "itemsOnPage": 1000,
                  "address": "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj",
                  "balance": "15245",
                  "totalReceived": "91089",
                  "totalSent": "75844",
                  "unconfirmedBalance": "0",
                  "unconfirmedTxs": 0,
                  "txs": 12,
                  "transactions": [{
                          "txid": "36b1e721a25ea3ac2fcc09a92d4ff1e2ae4ed70d593e276806d9a9fd2a901132",
                          "version": 1,
                          "vin": [{
                              "txid": "c6a4c82d5c7a342796e7d81237ab399918d3205f791ebc40e63501cac28c32be",
                              "vout": 1,
                              "n": 0,
                              "addresses": [
                                  "bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"
                              ],
                              "isAddress": true,
                              "value": "4829"
                          }],
                          "vout": [{
                                  "value": "1000",
                                  "n": 0,
                                  "hex": "a9141ba5187259ae22520f99ec3522d320279066aafd87",
                                  "addresses": [
                                      "34DBwkzPz6yeqU1LoZLVzdCm91oeyxq37T"
                                  ],
                                  "isAddress": true
                              },
                              {
                                  "value": "2699",
                                  "n": 1,
                                  "hex": "00148f347b3ea1ce113b09e096125a3ad3310dd0a947",
                                  "addresses": [
                                      "bc1q3u68k04pecgnkz0qjcf95wknxyxap2287gyzrg"
                                  ],
                                  "isAddress": true
                              }
                          ],
                          "blockHash": "0000000000000000000fccc2c78e235d6ea3da4013ea98c8376c7c41f7f71938",
                          "blockHeight": 624894,
                          "confirmations": 2494,
                          "blockTime": 1586299624,
                          "value": "3699",
                          "valueIn": "4829",
                          "fees": "1130",
                          "hex": "01000000000101be328cc2ca0135e640bc1e795f20d3189939ab3712d8e79627347a5c2dc8a4c601000000000000000002e80300000000000017a9141ba5187259ae22520f99ec3522d320279066aafd878b0a0000000000001600148f347b3ea1ce113b09e096125a3ad3310dd0a94702483045022100c251d7e2497b42d57ff00c92eeb8a6faa86a3d5d77a9d257b7e464400306fa5502204f532ab553ba711779cf1f93cab47c4f5e0f11dbc6017781c79b76caa6ea322b012102624729d04d58fa33eb3f7be9fe9c307c60aa1bad52f4ffbacc78ba3808faf62500000000"
                      },
                      {
                          "txid": "8180545030fcfb7b14ee90acb01b606c5a68fe1d520bcf140bd0097108c2b7f4",
                          "version": 1,
                          "vin": [{
                              "txid": "270a6490cc806d99aa84bb8079b543d9d7a255fc59364890e7a4cb57970daef8",
                              "vout": 1,
                              "sequence": 4294967294,
                              "n": 0,
                              "addresses": [
                                  "bc1q3230a7cqt2drewuza8qff4c4gpt4muy9qyqknw"
                              ],
                              "isAddress": true,
                              "value": "1699"
                          }],
                          "vout": [{
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
                          "confirmations": 5371,
                          "blockTime": 1584485709,
                          "value": "1473",
                          "valueIn": "1699",
                          "fees": "226",
                          "hex": "01000000000101f8ae0d9757cba4e790483659fc55a2d7d943b57980bb84aa996d80cc90640a270100000000feffffff0277010000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a4a040000000000001600141a475acd52ae04da60ab33bf373c9255cea3169a024730440220683d9c2b40958fd706623783017300ebbc9b73ce1db571963d9bb40679c3079c02207ee68976e225e1e98d20e66d4d8a498cddf3268d0d7b66827a2ed3ac4330f137012102b7f0cb54c2a6a4da3372fe17d7a07475f575777364633e89dacb39eb1fd5a09c00000000"
                      }
                  ]
                }
                `);
        }
        return { error: "Not implemented" };
    }
}
