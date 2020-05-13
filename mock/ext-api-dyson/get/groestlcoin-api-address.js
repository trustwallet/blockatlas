/// Mock for external Groestlcoin API
/// See:
/// curl "http://{groestlcoin rpc}/api/v2/address/33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj?details=txs"
/// curl "http://localhost:3347/groestlcoin-api/v2/address/33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj?details=txs"
/// curl "http://localhost:8437/v1/groestlcoin/address/33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj"

module.exports = {
    path: '/groestlcoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case '33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 1,
                    "itemsOnPage": 2,
                    "address": "33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj",
                    "balance": "0",
                    "totalReceived": "5951153060",
                    "totalSent": "5951153060",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 2,
                    "transactions": [
                        {
                            "txid": "2640aa5de0c9603da1c0d9c16b2fd3fa0a17b1472c3aa02559d3ef5e1defceb5",
                            "version": 2,
                            "lockTime": 2959295,
                            "vin": [
                                {
                                    "txid": "d8e0d23dedc2c89e9de317e7a54bdc3d26f514918a9360aa04e271c4d8074c28",
                                    "sequence": 4294967294,
                                    "n": 0,
                                    "addresses": [
                                        "33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj"
                                    ],
                                    "isAddress": true,
                                    "value": "5951153060",
                                    "hex": "160014d6c589125f084df1e3286fcd55446b64dc7de130"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "1151149700",
                                    "n": 0,
                                    "spent": true,
                                    "hex": "a91436d64490426cc347a50bdd3f8db2ef20d62949f587",
                                    "addresses": [
                                        "36gy6VVstfso35mS89pBg1PiUcYY3Gesar"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "4800000000",
                                    "n": 1,
                                    "spent": true,
                                    "hex": "76a914fcad3abf614562d224c6cc8b0e00d2fa9016404388ac",
                                    "addresses": [
                                        "FtCkFSrwrgiJzjQzGRZvjHzrmHp4HJeGYm"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "0000000000000a79428395294255704ed877847d93c6d36108dc8184b71c1f0a",
                            "blockHeight": 2959365,
                            "confirmations": 100119,
                            "blockTime": 1581386699,
                            "value": "5951149700",
                            "valueIn": "5951153060",
                            "fees": "3360",
                            "hex": "02000000000101284c07d8c471e204aa60938a9114f5263ddc4ba5e717e39d9ec8c2ed3dd2e0d80000000017160014d6c589125f084df1e3286fcd55446b64dc7de130feffffff0284269d440000000017a91436d64490426cc347a50bdd3f8db2ef20d62949f58700301a1e010000001976a914fcad3abf614562d224c6cc8b0e00d2fa9016404388ac02473044022034f3f2ab2d021a27ba999aebb40016f921433c39149d6908fe1e96d914c5c96402203d5d12127f64a01429775090abb445b5af2ec90803372c92499a35e12e229adb0121033ca60a0478fee5583e52c3b85c4dacb81faa9c4a10ad8b4f574c1b050f814463bf272d00"
                        },
                        {
                            "txid": "d8e0d23dedc2c89e9de317e7a54bdc3d26f514918a9360aa04e271c4d8074c28",
                            "version": 2,
                            "lockTime": 2959360,
                            "vin": [
                                {
                                    "txid": "2ed852f7881270ec108c86482d609f818ee21ae07033796fb77cb8e52fa86ccd",
                                    "sequence": 4294967294,
                                    "n": 0,
                                    "addresses": [
                                        "Fg4WGddhNYayAF3mTPDNCFCEqrXydAd6Vu"
                                    ],
                                    "isAddress": true,
                                    "value": "29751157520",
                                    "hex": "47304402201fe0aff8dd5c35be49f824216adb51d0749878dc2d759d0dd6d4ed6612ca09cd02203ad7d2b2eaa52f341f325934cca2c2758ac4c4b7b8bc279e7748d48ab98b484901210246e7e23df8acf0a305b94009d9b60eca7230ce747493201a9574f7ccb03775e9"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "5951153060",
                                    "n": 0,
                                    "spent": true,
                                    "hex": "a914146081496e97dbb864af7df601184f8ec3624aa787",
                                    "addresses": [
                                        "33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "23800000000",
                                    "n": 1,
                                    "spent": true,
                                    "hex": "76a914fcad3abf614562d224c6cc8b0e00d2fa9016404388ac",
                                    "addresses": [
                                        "FtCkFSrwrgiJzjQzGRZvjHzrmHp4HJeGYm"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "0000000000000a79428395294255704ed877847d93c6d36108dc8184b71c1f0a",
                            "blockHeight": 2959365,
                            "confirmations": 100119,
                            "blockTime": 1581386699,
                            "value": "29751153060",
                            "valueIn": "29751157520",
                            "fees": "4460",
                            "hex": "0200000001cd6ca82fe5b87cb76f793370e01ae28e819f602d48868c10ec701288f752d82e000000006a47304402201fe0aff8dd5c35be49f824216adb51d0749878dc2d759d0dd6d4ed6612ca09cd02203ad7d2b2eaa52f341f325934cca2c2758ac4c4b7b8bc279e7748d48ab98b484901210246e7e23df8acf0a305b94009d9b60eca7230ce747493201a9574f7ccb03775e9feffffff02a463b7620100000017a914146081496e97dbb864af7df601184f8ec3624aa787002e978a050000001976a914fcad3abf614562d224c6cc8b0e00d2fa9016404388ac00282d00"
                        }
                    ]
                }
                `);
        }
        return {error: "Not implemented"};
    }
}
