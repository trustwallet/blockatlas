/// Mock for external Zelcash API
/// See:
/// curl "http://{Zelcash rpc}/api/v2/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa?details=txs&pageSize=25"
/// curl "http://localhost:3347/zelcash-api/v2/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa?details=txs"
/// curl "http://localhost:8437/v1/zelcash/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa"

module.exports = {
    path: '/zelcash-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 't1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 58223,
                    "itemsOnPage": 2,
                    "address": "t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa",
                    "balance": "787501194898",
                    "totalReceived": "1224870418723238",
                    "totalSent": "1224082917528340",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 116445,
                    "transactions": [
                        {
                            "txid": "58b5ce7d497d95e21ad9ebe2f2a3fb480483f77ca4de3ce1bc9a95eb9c0d18e1",
                            "version": 4,
                            "vin": [
                                {
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "coinbase": "0322e80800324d696e6572732068747470733a2f2f326d696e6572732e636f6d"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "11250000000",
                                    "n": 0,
                                    "hex": "76a91404e2699cec5f44280540fb752c7660aa3ba857cc88ac",
                                    "addresses": [
                                        "t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa"
                                    ]
                                },
                                {
                                    "value": "562500000",
                                    "n": 1,
                                    "hex": "76a914cc069c2af44790f67f6d1a26fa4668cd18ae3cee88ac",
                                    "addresses": [
                                        "t1cUPmjzAv1UHmVDbdhwJTXQWER1Y9RXrab"
                                    ]
                                },
                                {
                                    "value": "937500000",
                                    "n": 2,
                                    "hex": "76a9147e89c1c8a911588d4dd2df5ee0906de9b367631c88ac",
                                    "addresses": [
                                        "t1VQgB2axkuPhYQfsbvUG7Wfhfm7WBEYiEb"
                                    ]
                                },
                                {
                                    "value": "2250000000",
                                    "n": 3,
                                    "hex": "76a914c33faabb9d5162c37c0b29dcdb357a39738637b888ac",
                                    "addresses": [
                                        "t1bfz3Q3E1adjroiPrGXaWsWNJ7u3WJHBgS"
                                    ]
                                }
                            ],
                            "blockHash": "0000009f069aa63aef11ed375763532f3863bb2f03135827e1d6dfc5af3c4746",
                            "blockHeight": 583714,
                            "confirmations": 2,
                            "blockTime": 1587702257,
                            "value": "15000000000",
                            "valueIn": "0",
                            "fees": "0",
                            "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff200322e80800324d696e6572732068747470733a2f2f326d696e6572732e636f6dffffffff0480608d9e020000001976a91404e2699cec5f44280540fb752c7660aa3ba857cc88aca0118721000000001976a914cc069c2af44790f67f6d1a26fa4668cd18ae3cee88ac601de137000000001976a9147e89c1c8a911588d4dd2df5ee0906de9b367631c88ac80461c86000000001976a914c33faabb9d5162c37c0b29dcdb357a39738637b888ac00000000000000000000000000000000000000"
                        },
                        {
                            "txid": "eaddfc88f931b67c3046b8046fa96831fdeb590db9be4b1e59e706479b0f3012",
                            "version": 4,
                            "vin": [
                                {
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "coinbase": "031fe80800324d696e6572732068747470733a2f2f326d696e6572732e636f6d"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "11250011000",
                                    "n": 0,
                                    "hex": "76a91404e2699cec5f44280540fb752c7660aa3ba857cc88ac",
                                    "addresses": [
                                        "t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa"
                                    ]
                                },
                                {
                                    "value": "562500000",
                                    "n": 1,
                                    "hex": "76a9148a8b9570ece58cfce36487fe68da103964f2ff4288ac",
                                    "addresses": [
                                        "t1WWAUFtKyytFzuXpn3Rzc6fGFndzD6xWLK"
                                    ]
                                },
                                {
                                    "value": "937500000",
                                    "n": 2,
                                    "hex": "76a914c6e488b6eaa3becba6a1f439d3a5c05a41b6b77f88ac",
                                    "addresses": [
                                        "t1c1Fa9ARp7KvHgEi3SKefpkLzy2yBTZvZc"
                                    ]
                                },
                                {
                                    "value": "2250000000",
                                    "n": 3,
                                    "hex": "76a91446f4a65996b3689d296445f2a1170f9f87c2f82e88ac",
                                    "addresses": [
                                        "t1QLnPL6ozQjArqX4QSDesMNkXTFWUUhbeF"
                                    ]
                                }
                            ],
                            "blockHash": "0000006926830de74fa8629389bb21802ecbfe254ab52198678de133ed6c3b22",
                            "blockHeight": 583711,
                            "confirmations": 5,
                            "blockTime": 1587701924,
                            "value": "15000011000",
                            "valueIn": "0",
                            "fees": "0",
                            "hex": "0400008085202f89010000000000000000000000000000000000000000000000000000000000000000ffffffff20031fe80800324d696e6572732068747470733a2f2f326d696e6572732e636f6dffffffff04788b8d9e020000001976a91404e2699cec5f44280540fb752c7660aa3ba857cc88aca0118721000000001976a9148a8b9570ece58cfce36487fe68da103964f2ff4288ac601de137000000001976a914c6e488b6eaa3becba6a1f439d3a5c05a41b6b77f88ac80461c86000000001976a91446f4a65996b3689d296445f2a1170f9f87c2f82e88ac00000000000000000000000000000000000000"
                        }
                    ]
                }
                `);
        }
        return {error: "Not implemented"};
    }
}
