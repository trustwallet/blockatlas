/// Mock for external Zcash API
/// See:
/// curl "http://{Zcash rpc}/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs"
/// curl "http://localhost:3000/zcash-api/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs"
/// curl "http://localhost:8420/v1/zcash/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX"

module.exports = {
    path: '/zcash-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 't1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 2,
                        "addrStr": "t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX",
                        "balance": "0.12344656",
                        "totalReceived": "1096.63825939",
                        "totalSent": "1096.51481283",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 1636,
                        "txs": [
                            {
                                "txid": "d504755667e0fb61543e16d09aff5f84733f500190a07acb4a6784b4dc4c99e7",
                                "version": 4,
                                "vin": [
                                    {
                                        "txid": "",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": null,
                                        "value": ""
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914219d1bea95fc9c401d4073a20f840c71fc1a920288ac",
                                            "addresses": [
                                                "t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "1.25",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "a914e445cfa944b6f2bdacefbda904a81d5fdd26d77f87",
                                            "addresses": [
                                                "t3fNcdBUbycvbCtsD2n9q3LuxG7jVPvFB8L"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "0.0625",
                                        "n": 2,
                                        "scriptPubKey": {
                                            "hex": "76a914f82c1c9f44630fe1a9ab4cb1b5023c91ee4ddae688ac",
                                            "addresses": [
                                                "t1gVpRrLdr8R2M2RaGisWZRspEF8EZvLZHv"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "4.9375",
                                        "n": 3,
                                        "scriptPubKey": {
                                            "hex": "76a914521d713a3660b86d03ab6c68358433c30389ea1f88ac",
                                            "addresses": [
                                                "t1RMngkaWf2F8EDrpzQ1fhLBhfi89L14tN2"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "000000000259ba5a4b53e1ce8e6828fc780f8858283620aa27418ce80681c1b5",
                                "blockheight": 768336,
                                "confirmations": 7,
                                "time": 1584747959,
                                "blocktime": 1584747959,
                                "valueOut": "6.25",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff500350b90b0044656661756c7420732d6e6f6d7020706f6f6c2068747470733a2f2f6769746875622e636f6d2f732d6e6f6d702f732d6e6f6d702f77696b692f496e73696768742d706f6f6c2d6c696e6bffffffff0400000000000000001976a914219d1bea95fc9c401d4073a20f840c71fc1a920288ac405973070000000017a914e445cfa944b6f2bdacefbda904a81d5fdd26d77f87105e5f00000000001976a914f82c1c9f44630fe1a9ab4cb1b5023c91ee4ddae688acf0066e1d000000001976a914521d713a3660b86d03ab6c68358433c30389ea1f88ac00000000000000000000000000000000000000"
                            },
                            {
                                "txid": "95e12ff1a8ca7a19d5af92c6a480d1c1da4cf998dcba08329d151acb45d78bbb",
                                "version": 4,
                                "vin": [
                                    {
                                        "txid": "",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": null,
                                        "value": ""
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914219d1bea95fc9c401d4073a20f840c71fc1a920288ac",
                                            "addresses": [
                                                "t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "1.25",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "a914e445cfa944b6f2bdacefbda904a81d5fdd26d77f87",
                                            "addresses": [
                                                "t3fNcdBUbycvbCtsD2n9q3LuxG7jVPvFB8L"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "0.0625",
                                        "n": 2,
                                        "scriptPubKey": {
                                            "hex": "76a914f82c1c9f44630fe1a9ab4cb1b5023c91ee4ddae688ac",
                                            "addresses": [
                                                "t1gVpRrLdr8R2M2RaGisWZRspEF8EZvLZHv"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "4.9375",
                                        "n": 3,
                                        "scriptPubKey": {
                                            "hex": "76a914521d713a3660b86d03ab6c68358433c30389ea1f88ac",
                                            "addresses": [
                                                "t1RMngkaWf2F8EDrpzQ1fhLBhfi89L14tN2"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "0000000000f747c6311d3658bad9bab649e21dd365075a1df8b98a88c3871458",
                                "blockheight": 768281,
                                "confirmations": 62,
                                "time": 1584744347,
                                "blocktime": 1584744347,
                                "valueOut": "6.25",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff500319b90b0044656661756c7420732d6e6f6d7020706f6f6c2068747470733a2f2f6769746875622e636f6d2f732d6e6f6d702f732d6e6f6d702f77696b692f496e73696768742d706f6f6c2d6c696e6bffffffff0400000000000000001976a914219d1bea95fc9c401d4073a20f840c71fc1a920288ac405973070000000017a914e445cfa944b6f2bdacefbda904a81d5fdd26d77f87105e5f00000000001976a914f82c1c9f44630fe1a9ab4cb1b5023c91ee4ddae688acf0066e1d000000001976a914521d713a3660b86d03ab6c68358433c30389ea1f88ac00000000000000000000000000000000000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
