/// Mock for external Zelcash API
/// See:
/// curl "http://{Zelcash rpc}/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa?details=txs&pageSize=25"
/// curl "http://localhost:3000/zelcash-api/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa?details=txs"
/// curl "http://localhost:8420/v1/zelcash/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa"

module.exports = {
    path: '/zelcash-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 't1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 4175,
                        "itemsOnPage": 25,
                        "addrStr": "t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa",
                        "balance": "6300.00809211",
                        "totalReceived": "10963952.85564162",
                        "totalSent": "10957652.84754951",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 104369,
                        "txs": [
                            {
                                "txid": "ca8b3fa115c5eae0f056d9b629680c521f4d9d271abbe86d4c482c54117e15a1",
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
                                        "value": "112.5001",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a91404e2699cec5f44280540fb752c7660aa3ba857cc88ac",
                                            "addresses": [
                                                "t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "5.625",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a91402664e0d583d97b55b71320eea199c20ecbab3a588ac",
                                            "addresses": [
                                                "t1J6HuzEqw8Aicm1WLL8YX7jo7Syy7zkTjs"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "9.375",
                                        "n": 2,
                                        "scriptPubKey": {
                                            "hex": "76a9146096374c8e00616fd777c9497d9fce18eae3948d88ac",
                                            "addresses": [
                                                "t1SgJpwdmPSgw4j6eAdVVY1duWkbixygptP"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "22.5",
                                        "n": 3,
                                        "scriptPubKey": {
                                            "hex": "76a9145a5f63ff65c2b1f4124ed2f5aa6fd9db53ec2f4488ac",
                                            "addresses": [
                                                "t1S7T6TRuYuodWyAkiRVxRj5XgfFs1G3PWY"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "000000351df0a7b2a629657aecc89ce10d07f7c5a9ad7f03fefb3442b1a7a521",
                                "blockheight": 562387,
                                "confirmations": 1,
                                "time": 1585124312,
                                "blocktime": 1585124312,
                                "valueOut": "150.0001",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff2003d3940800324d696e6572732068747470733a2f2f326d696e6572732e636f6dffffffff0490878d9e020000001976a91404e2699cec5f44280540fb752c7660aa3ba857cc88aca0118721000000001976a91402664e0d583d97b55b71320eea199c20ecbab3a588ac601de137000000001976a9146096374c8e00616fd777c9497d9fce18eae3948d88ac80461c86000000001976a9145a5f63ff65c2b1f4124ed2f5aa6fd9db53ec2f4488ac00000000000000000000000000000000000000"
                            },
                            {
                                "txid": "96990e262110a89d8e372df3126cf72023b07643325c9b8b38f4c3651c1931b3",
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
                                        "value": "112.5",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a91404e2699cec5f44280540fb752c7660aa3ba857cc88ac",
                                            "addresses": [
                                                "t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "5.625",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a914361dfb3217992096d4fc8cfac8a076cdc4e9aaf588ac",
                                            "addresses": [
                                                "t1NokQpvLQo1Ti4MzvqPM24tn5psXEbJAWu"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "9.375",
                                        "n": 2,
                                        "scriptPubKey": {
                                            "hex": "76a9142ccfb039f8a3057684832334b2b96f38299d9a9988ac",
                                            "addresses": [
                                                "t1MxYYDbod3vJNDpmCXPF99SpoFZ8j4k8dZ"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "22.5",
                                        "n": 3,
                                        "scriptPubKey": {
                                            "hex": "76a9143d94fc18dd332c386014f9e5b05cce17710a2ce688ac",
                                            "addresses": [
                                                "t1PVDhgMueDh8vY3JfkZoXcGrwATCfkwaL9"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "00000027379c0e0420ab3ebfc88a0f1bf59f10ced3317ce70f64ffa53f6a7775",
                                "blockheight": 562386,
                                "confirmations": 2,
                                "time": 1585124285,
                                "blocktime": 1585124285,
                                "valueOut": "150",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff2003d2940800324d696e6572732068747470733a2f2f326d696e6572732e636f6dffffffff0480608d9e020000001976a91404e2699cec5f44280540fb752c7660aa3ba857cc88aca0118721000000001976a914361dfb3217992096d4fc8cfac8a076cdc4e9aaf588ac601de137000000001976a9142ccfb039f8a3057684832334b2b96f38299d9a9988ac80461c86000000001976a9143d94fc18dd332c386014f9e5b05cce17710a2ce688ac00000000000000000000000000000000000000"
                            },
                            {
                                "txid": "aa1012b3cfcc36b0d6e10c25d0550d8ff482462cfc39e2ae93b03e9f7021bc00",
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
                                        "value": "112.5",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a91404e2699cec5f44280540fb752c7660aa3ba857cc88ac",
                                            "addresses": [
                                                "t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "5.625",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a914490683ffa4c8c06eca7f0ede150cf267df5e920988ac",
                                            "addresses": [
                                                "t1QXj93zU8HVsfTpqJTHCnoZ9QvR2ho1rp7"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "9.375",
                                        "n": 2,
                                        "scriptPubKey": {
                                            "hex": "76a91415f6ecf50ff229de0462f7003dcd5c5656f1e56888ac",
                                            "addresses": [
                                                "t1Ksk14EFS9GqwGUF5kqnyR9zpYnKDAfxA5"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "22.5",
                                        "n": 3,
                                        "scriptPubKey": {
                                            "hex": "76a914ee494e77ec55c57b2f5b42b1487a15fb3264227188ac",
                                            "addresses": [
                                                "t1fbYeRABbZfzuxViLQbJ4d1dBTET48FgnV"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "000000938163fab733308a221b50373250c7e8316e1420992abb833208152916",
                                "blockheight": 562385,
                                "confirmations": 3,
                                "time": 1585124162,
                                "blocktime": 1585124162,
                                "valueOut": "150",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff2003d1940800324d696e6572732068747470733a2f2f326d696e6572732e636f6dffffffff0480608d9e020000001976a91404e2699cec5f44280540fb752c7660aa3ba857cc88aca0118721000000001976a914490683ffa4c8c06eca7f0ede150cf267df5e920988ac601de137000000001976a91415f6ecf50ff229de0462f7003dcd5c5656f1e56888ac80461c86000000001976a914ee494e77ec55c57b2f5b42b1487a15fb3264227188ac00000000000000000000000000000000000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
