/// Mock for external Ravencoin API
/// See:
/// curl "http://{Ravencoin rpc}/api/v2/address/RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo?details=txs"
/// curl "http://localhost:3347/ravencoin-api/v2/address/RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo?details=txs"
/// curl "http://localhost:8437/v1/ravencoin/address/RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo"

module.exports = {
    path: '/ravencoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 1,
                    "itemsOnPage": 1000,
                    "address": "RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo",
                    "balance": "8525568094",
                    "totalReceived": "8525568094",
                    "totalSent": "0",
                    "unconfirmedBalance": "-8525568094",
                    "unconfirmedTxs": 1,
                    "txs": 1,
                    "transactions": [
                        {
                            "txid": "fc226aad6fc28e1204b747042e8c8b25f5d28424b9669bd162ba6a0149df2f71",
                            "version": 2,
                            "lockTime": 1201754,
                            "vin": [
                                {
                                    "txid": "1222cb57d31bfb439e94080266c33ebb0134cdb90999a9fad99455d09a15159b",
                                    "sequence": 4294967294,
                                    "n": 0,
                                    "addresses": [
                                        "RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo"
                                    ],
                                    "isAddress": true,
                                    "value": "8525568094",
                                    "hex": "483045022100dc56832c815e7294dd51c7bb1ba342af80f5d25c3bc44c9689a55ccd94d18d01022063d3420930736fb1c1ab9a851ec0d1ff01dc906c66685b35972e3ab26b2d5d040121032dd4ba9d193e1912a6b48bd5642cb3579415ebabfc61e9fec9690fcd466dea15"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "78460623",
                                    "n": 0,
                                    "hex": "76a9141657456724a83ca9e4c4cc0b65c98b6ffdba5bae88ac",
                                    "addresses": [
                                        "RBKKVSR79YjBSGBE5pymUCaa871qogE5aY"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "8446853390",
                                    "n": 1,
                                    "hex": "76a9142e664ae2c04d6292d50bcd65b644553eea7d7f1388ac",
                                    "addresses": [
                                        "RDWXhkQyUkgmumGbzdR6JetLChVzUbStRq"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHeight": -1,
                            "confirmations": 0,
                            "blockTime": 1587702074,
                            "value": "8525314013",
                            "valueIn": "8525568094",
                            "fees": "254081",
                            "hex": "02000000019b15159ad05594d9faa99909b9cd3401bb3ec3660208949e43fb1bd357cb2212000000006b483045022100dc56832c815e7294dd51c7bb1ba342af80f5d25c3bc44c9689a55ccd94d18d01022063d3420930736fb1c1ab9a851ec0d1ff01dc906c66685b35972e3ab26b2d5d040121032dd4ba9d193e1912a6b48bd5642cb3579415ebabfc61e9fec9690fcd466dea15feffffff02cf36ad04000000001976a9141657456724a83ca9e4c4cc0b65c98b6ffdba5bae88ac0ec178f7010000001976a9142e664ae2c04d6292d50bcd65b644553eea7d7f1388ac5a561200"
                        },
                        {
                            "txid": "1222cb57d31bfb439e94080266c33ebb0134cdb90999a9fad99455d09a15159b",
                            "version": 2,
                            "lockTime": 1201750,
                            "vin": [
                                {
                                    "txid": "003d748c2c3541535093e840b4c5e54b966915eb522bac06028bc8d1376e1e63",
                                    "sequence": 4294967294,
                                    "n": 0,
                                    "addresses": [
                                        "R9kyoRBmiF89o5ACFXF4EY7GawGjURP1Z5"
                                    ],
                                    "isAddress": true,
                                    "value": "8530842427",
                                    "hex": "47304402207ea927cfb1c7d8e14b067aa5b892cfde8b7b96b8bd95866b906c13e168581fd20220423bee1c4e25f66cc5ac27d94a33540307edb2a9a8874e800a1e1f875194dabf012103ded5613a24b22a8dc0c416f3ac4b436372a5b581b42e2d45f00fbd042cb066f1"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "8525568094",
                                    "n": 0,
                                    "hex": "76a9145208b754b23169797a19e0ffaa9041ab5e29b9ef88ac",
                                    "addresses": [
                                        "RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "5020252",
                                    "n": 1,
                                    "spent": true,
                                    "hex": "76a914bb7590eef377d63adcda0d080073035ad9a86bfc88ac",
                                    "addresses": [
                                        "RSNPH2YPN69PNRj6BYq4V3orCyC7nahG9R"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "00000000000003efc9fb9dd4d5deb4e6ca393266191771d5d2e2426a191e69b3",
                            "blockHeight": 1201752,
                            "confirmations": 3,
                            "blockTime": 1587701825,
                            "value": "8530588346",
                            "valueIn": "8530842427",
                            "fees": "254081",
                            "hex": "0200000001631e6e37d1c88b0206ac2b52eb1569964be5c5b440e893505341352c8c743d00000000006a47304402207ea927cfb1c7d8e14b067aa5b892cfde8b7b96b8bd95866b906c13e168581fd20220423bee1c4e25f66cc5ac27d94a33540307edb2a9a8874e800a1e1f875194dabf012103ded5613a24b22a8dc0c416f3ac4b436372a5b581b42e2d45f00fbd042cb066f1feffffff025ed829fc010000001976a9145208b754b23169797a19e0ffaa9041ab5e29b9ef88ac5c9a4c00000000001976a914bb7590eef377d63adcda0d080073035ad9a86bfc88ac56561200"
                        }
                    ]
                }
                `);
        }
        return {error: "Not implemented"};
    }
}
