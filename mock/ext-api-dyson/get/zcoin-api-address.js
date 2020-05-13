/// Mock for external Zcoin API
/// See:
/// curl "http://{Zcoin rpc}/api/v2/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn?details=txs"
/// curl "http://localhost:3347/zcoin-api/v2/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn?details=txs"
/// curl "http://localhost:8437/v1/zcoin/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"

module.exports = {
    path: '/zcoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 9,
                    "itemsOnPage": 2,
                    "address": "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn",
                    "balance": "63109110",
                    "totalReceived": "1565104974",
                    "totalSent": "1501995864",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 18,
                    "transactions": [
                        {
                            "txid": "f1db892b13a2cb34a9a1ed0890b85050c43b95e06dc085d7f063e9207e984609",
                            "version": 1,
                            "vin": [
                                {
                                    "txid": "cd33ac3382c4a5b2eb70f6d7625289446430ff9484326ffc234bc498c274b564",
                                    "vout": 1,
                                    "n": 0,
                                    "addresses": [
                                        "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                    ],
                                    "isAddress": true,
                                    "value": "51916634",
                                    "hex": "483045022100dc5a93fa3bd1f61115062700756163da79860b38fe348aa9c3bc5d9c578c8d55022048f2f300dce9858fa3a24e05084bb892287a623b40271289cb5fce9783e77b730121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "13113178",
                                    "n": 0,
                                    "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                    "addresses": [
                                        "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "38799388",
                                    "n": 1,
                                    "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                    "addresses": [
                                        "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "f85889dc565371828f3c0821f2507ca2fa2f83eae557b7d3ba7270eab6ac7328",
                            "blockHeight": 251191,
                            "confirmations": 10102,
                            "blockTime": 1584663676,
                            "value": "51912566",
                            "valueIn": "51916634",
                            "fees": "4068",
                            "hex": "010000000164b574c298c44b23fc6f328494ff306444895262d7f670ebb2a5c48233ac33cd010000006b483045022100dc5a93fa3bd1f61115062700756163da79860b38fe348aa9c3bc5d9c578c8d55022048f2f300dce9858fa3a24e05084bb892287a623b40271289cb5fce9783e77b730121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b1500000000025a17c800000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac1c085002000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                        },
                        {
                            "txid": "cd33ac3382c4a5b2eb70f6d7625289446430ff9484326ffc234bc498c274b564",
                            "version": 1,
                            "vin": [
                                {
                                    "txid": "50cb7a384e031c2823147c440884dbbf6ecef313925f113768406ab7776229fd",
                                    "vout": 1,
                                    "n": 0,
                                    "addresses": [
                                        "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                    ],
                                    "isAddress": true,
                                    "value": "61920702",
                                    "hex": "483045022100aff10fef2849988470ac3411bd463ae03a17a88336283290095f8247dc0531a4022064f2c0e60662effb5b46c2a70318331b007df7441a209e2bf6489924c73952f50121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "10000000",
                                    "n": 0,
                                    "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                    "addresses": [
                                        "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "51916634",
                                    "n": 1,
                                    "spent": true,
                                    "spentTxId": "f1db892b13a2cb34a9a1ed0890b85050c43b95e06dc085d7f063e9207e984609",
                                    "spentHeight": 251191,
                                    "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                    "addresses": [
                                        "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "f85889dc565371828f3c0821f2507ca2fa2f83eae557b7d3ba7270eab6ac7328",
                            "blockHeight": 251191,
                            "confirmations": 10102,
                            "blockTime": 1584663676,
                            "value": "61916634",
                            "valueIn": "61920702",
                            "fees": "4068",
                            "hex": "0100000001fd296277b76a406837115f9213f3ce6ebfdb8408447c1423281c034e387acb50010000006b483045022100aff10fef2849988470ac3411bd463ae03a17a88336283290095f8247dc0531a4022064f2c0e60662effb5b46c2a70318331b007df7441a209e2bf6489924c73952f50121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15000000000280969800000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac5a2f1803000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                        }
                    ]
                }
                `);
        }
        return {error: "Not implemented"};
    }
}
