/// Mock for external Zcash API
/// See:
/// curl "http://{Zcash rpc}/api/v2/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs"
/// curl "http://localhost:3347/zcash-api/v2/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs"
/// curl "http://localhost:8437/v1/zcash/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX"

module.exports = {
    path: '/zcash-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 't1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 917,
                    "itemsOnPage": 2,
                    "address": "t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX",
                    "balance": "12344656",
                    "totalReceived": "109663825939",
                    "totalSent": "109651481283",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 1833,
                    "transactions": [
                        {
                            "txid": "ca640e5215e141a4f9c235dd91e4b99fd1e0defdd5da27d78ec2cee02d3493dc",
                            "version": 4,
                            "vin": [
                                {
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "isAddress": false,
                                    "coinbase": "0323520c0044656661756c7420732d6e6f6d7020706f6f6c2068747470733a2f2f6769746875622e636f6d2f732d6e6f6d702f732d6e6f6d702f77696b692f496e73696768742d706f6f6c2d6c696e6b"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "0",
                                    "n": 0,
                                    "hex": "76a914219d1bea95fc9c401d4073a20f840c71fc1a920288ac",
                                    "addresses": [
                                        "t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "125000000",
                                    "n": 1,
                                    "hex": "a914cdcd95aa6892db5fc849ac15804d2b9dd035a3b787",
                                    "addresses": [
                                        "t3dKojUU2EMjs28nHV84TvkVEUDu1M1FaEx"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "6250000",
                                    "n": 2,
                                    "hex": "76a914f82c1c9f44630fe1a9ab4cb1b5023c91ee4ddae688ac",
                                    "addresses": [
                                        "t1gVpRrLdr8R2M2RaGisWZRspEF8EZvLZHv"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "493750000",
                                    "n": 3,
                                    "hex": "76a914521d713a3660b86d03ab6c68358433c30389ea1f88ac",
                                    "addresses": [
                                        "t1RMngkaWf2F8EDrpzQ1fhLBhfi89L14tN2"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "000000000031b763658a48c735fa02d9e5181b12683cc66a93d79caf5f66e9b0",
                            "blockHeight": 807459,
                            "confirmations": 89,
                            "blockTime": 1587696399,
                            "value": "625000000",
                            "valueIn": "0",
                            "fees": "0",
                            "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff500323520c0044656661756c7420732d6e6f6d7020706f6f6c2068747470733a2f2f6769746875622e636f6d2f732d6e6f6d702f732d6e6f6d702f77696b692f496e73696768742d706f6f6c2d6c696e6bffffffff0400000000000000001976a914219d1bea95fc9c401d4073a20f840c71fc1a920288ac405973070000000017a914cdcd95aa6892db5fc849ac15804d2b9dd035a3b787105e5f00000000001976a914f82c1c9f44630fe1a9ab4cb1b5023c91ee4ddae688acf0066e1d000000001976a914521d713a3660b86d03ab6c68358433c30389ea1f88ac00000000000000000000000000000000000000"
                        },
                        {
                            "txid": "7bbf19b071907f23bd3d124d8f28a2aa8445a38d9f8a915487e9a36686c9094f",
                            "version": 4,
                            "vin": [
                                {
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "isAddress": false,
                                    "coinbase": "03a2510c0044656661756c7420732d6e6f6d7020706f6f6c2068747470733a2f2f6769746875622e636f6d2f732d6e6f6d702f732d6e6f6d702f77696b692f496e73696768742d706f6f6c2d6c696e6b"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "0",
                                    "n": 0,
                                    "hex": "76a914219d1bea95fc9c401d4073a20f840c71fc1a920288ac",
                                    "addresses": [
                                        "t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "125000000",
                                    "n": 1,
                                    "hex": "a914cdcd95aa6892db5fc849ac15804d2b9dd035a3b787",
                                    "addresses": [
                                        "t3dKojUU2EMjs28nHV84TvkVEUDu1M1FaEx"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "6250000",
                                    "n": 2,
                                    "spent": true,
                                    "hex": "76a914f82c1c9f44630fe1a9ab4cb1b5023c91ee4ddae688ac",
                                    "addresses": [
                                        "t1gVpRrLdr8R2M2RaGisWZRspEF8EZvLZHv"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "493750000",
                                    "n": 3,
                                    "spent": true,
                                    "hex": "76a914521d713a3660b86d03ab6c68358433c30389ea1f88ac",
                                    "addresses": [
                                        "t1RMngkaWf2F8EDrpzQ1fhLBhfi89L14tN2"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "00000000030219218c61edcf72dcd6c3512bbf9d3e5ffc34b607625d476dc558",
                            "blockHeight": 807330,
                            "confirmations": 218,
                            "blockTime": 1587686790,
                            "value": "625000000",
                            "valueIn": "0",
                            "fees": "0",
                            "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff5003a2510c0044656661756c7420732d6e6f6d7020706f6f6c2068747470733a2f2f6769746875622e636f6d2f732d6e6f6d702f732d6e6f6d702f77696b692f496e73696768742d706f6f6c2d6c696e6bffffffff0400000000000000001976a914219d1bea95fc9c401d4073a20f840c71fc1a920288ac405973070000000017a914cdcd95aa6892db5fc849ac15804d2b9dd035a3b787105e5f00000000001976a914f82c1c9f44630fe1a9ab4cb1b5023c91ee4ddae688acf0066e1d000000001976a914521d713a3660b86d03ab6c68358433c30389ea1f88ac00000000000000000000000000000000000000"
                        }
                    ]
                }
                `);
        }
        return {error: "Not implemented"};
    }
}
