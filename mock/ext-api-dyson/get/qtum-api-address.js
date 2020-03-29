/// Mock for external Qtum API
/// See:
/// curl "http://{qtum rpc}/address/QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ?details=txs"
/// curl "http://localhost:3000/qtum-api/address/QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ?details=txs"
/// curl "http://localhost:8420/v1/qtum/address/QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ"

module.exports = {
    path: '/qtum-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "addrStr": "QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ",
                        "balance": "2.3535761",
                        "totalReceived": "433.3535761",
                        "totalSent": "431",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 5,
                        "txs": [
                            {
                                "txid": "62438bb658856c3a08b89ca80e199e7031f98956ea86a135f5d6306660230f67",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "782e18a7b1bf612972e773fd6ccdcbe0c4514c661c31a7942f5dcd4807c5ab68",
                                        "vout": 0,
                                        "sequence": 4294967293,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "483045022100b660d73ff3bd7cc52d1052b933828cc2bb23b7a94500dff4bf00d5640ff4951702206ac6ecdaf3818dc57d54f8e860bbf212846200ccb2d1a323bf5b59868dcba872012102d7062f4af80f3a31f67c928e141627d807f0641b0bb0d466f4a62ee31230aa1e"
                                        },
                                        "addresses": [
                                            "QWLQSMPF5WKAqhXxePJvjUJNjxDHrPCbjB"
                                        ],
                                        "value": "0.001"
                                    },
                                    {
                                        "txid": "80b7c6bcba625ad4abb69f7e92521a782681e984982408d40ad35a24b3f78297",
                                        "vout": 0,
                                        "sequence": 4294967294,
                                        "n": 1,
                                        "scriptSig": {
                                            "hex": "48304502210090ea4a4aac51d8734ad9ffecf0e5c14be50a8941b8cac6a492c002a0be99e5e102203550b52b7ad546adf8b1e5ed00b7aa755cc296d4bd61004b2ace43bfe0007d17012103692f9ee47370b3a70e8d7b6ac36c705b2d3296bb37c0f02e15158f3a6001ce5b"
                                        },
                                        "addresses": [
                                            "QgqVXuCCN1m4Hfb9HCkrzbgUSkoodKsrVY"
                                        ],
                                        "value": "0.001"
                                    },
                                    {
                                        "txid": "9c3a6eb9fa07868a2e59530fe2aa3ea29a21d3818a70343097adbe571ee2f0be",
                                        "vout": 0,
                                        "sequence": 4294967292,
                                        "n": 2,
                                        "scriptSig": {
                                            "hex": "473044022055806e0939aa734d296ac364d4edf430bcd4fc15a377d35205362d09f7367d8e02202cc4e691c31ece66abb9adcd20bf8e3f500a2254edaaf2c241529cee670972440121025249520b3d48efd435e6cc30c9af238d054461d51f8beeaf75c4de50a647acbc"
                                        },
                                        "addresses": [
                                            "QSEXLG5qct47UtMHiXNg2Y27KSDSAEHJp9"
                                        ],
                                        "value": "0.1"
                                    },
                                    {
                                        "txid": "9c3a6eb9fa07868a2e59530fe2aa3ea29a21d3818a70343097adbe571ee2f0be",
                                        "vout": 1,
                                        "sequence": 4294967291,
                                        "n": 3,
                                        "scriptSig": {
                                            "hex": "483045022100e9d2126507af25eed231d27db52794bf6dc67c33800ebafd99b44b13394e1a2d02206b5abdce728114fa75624cd50d70f3c7da77c5b00433fa306b203e4bfb7d96320121020d708ba820b87e195b5790fa2723033cfec259ce5af4270a431cd7297d7704d3"
                                        },
                                        "addresses": [
                                            "Qfw9Poi2c4gGEkAAuVyG21PsBmNGsmZKFC"
                                        ],
                                        "value": "2.2549462"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "2.3535761",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a9148b4f969602a4579f9c982e7d30c3a06e2f0cacc888ac",
                                            "addresses": [
                                                "QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "6e71fa96ca29a6e4cbd10defaa4fc69ff09adec45fc666d82afade3f75206cc1",
                                "blockheight": 563426,
                                "confirmations": 11011,
                                "time": 1583720608,
                                "blocktime": 1583720608,
                                "valueOut": "2.3535761",
                                "valueIn": "2.3569462",
                                "fees": "0.0033701",
                                "hex": "010000000468abc50748cd5d2f94a7311c664c51c4e0cbcd6cfd73e7722961bfb1a7182e78000000006b483045022100b660d73ff3bd7cc52d1052b933828cc2bb23b7a94500dff4bf00d5640ff4951702206ac6ecdaf3818dc57d54f8e860bbf212846200ccb2d1a323bf5b59868dcba872012102d7062f4af80f3a31f67c928e141627d807f0641b0bb0d466f4a62ee31230aa1efdffffff9782f7b3245ad30ad408249884e98126781a52927e9fb6abd45a62babcc6b780000000006b48304502210090ea4a4aac51d8734ad9ffecf0e5c14be50a8941b8cac6a492c002a0be99e5e102203550b52b7ad546adf8b1e5ed00b7aa755cc296d4bd61004b2ace43bfe0007d17012103692f9ee47370b3a70e8d7b6ac36c705b2d3296bb37c0f02e15158f3a6001ce5bfeffffffbef0e21e57bead973034708a81d3219aa23eaae20f53592e8a8607fab96e3a9c000000006a473044022055806e0939aa734d296ac364d4edf430bcd4fc15a377d35205362d09f7367d8e02202cc4e691c31ece66abb9adcd20bf8e3f500a2254edaaf2c241529cee670972440121025249520b3d48efd435e6cc30c9af238d054461d51f8beeaf75c4de50a647acbcfcffffffbef0e21e57bead973034708a81d3219aa23eaae20f53592e8a8607fab96e3a9c010000006b483045022100e9d2126507af25eed231d27db52794bf6dc67c33800ebafd99b44b13394e1a2d02206b5abdce728114fa75624cd50d70f3c7da77c5b00433fa306b203e4bfb7d96320121020d708ba820b87e195b5790fa2723033cfec259ce5af4270a431cd7297d7704d3fbffffff01aa45070e000000001976a9148b4f969602a4579f9c982e7d30c3a06e2f0cacc888ac00000000"
                            },
                            {
                                "txid": "ef335bef783551cb054acd220555f1588cc4b03c0f47cc0839220c81d3ceb88d",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "506d07570d75914d5bd319d591374fb9cd1633f5c8b3cc800387e104298e8a62",
                                        "vout": 1,
                                        "sequence": 4294967294,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "4830450221008676b34896a922ac913b0922319f98ab4bb5223f479c19d62b92f0f187b2759d02204e6c2f57337f9b3ef262e2262b3e04a971635c450fbe4278db8b2b25e5a816540121027f67fba269482ce4eecfbcc31e7acefaf17a8cb3b4f7ccfb5447c71688e1081e"
                                        },
                                        "addresses": [
                                            "QfLtauKWxzF6pDf2ceXJe5oqizNHVMBneH"
                                        ],
                                        "value": "1.98984808"
                                    },
                                    {
                                        "txid": "e22c3e2ea4d285b71a73c41cdbb51f7c75d80b6f473e7f4aade6df5bf8771942",
                                        "vout": 0,
                                        "sequence": 4294967292,
                                        "n": 1,
                                        "scriptSig": {
                                            "hex": "483045022100e33513c3a4251a92af231c1d3e3bec9d13be2289c1fb203cb0096c6c073a8b840220206f696ee4e38da215f0b1e2c9d2005f085b8bb967d030caa70439046df64d4601210317cfe4e8b0e35f4c2c561921fe019910e3ddbe78d500a8ff546cf1843f723764"
                                        },
                                        "addresses": [
                                            "QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ"
                                        ],
                                        "value": "430"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "430",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a91455c89563888363b725aa0ebd8191a2e9ea54474d88ac",
                                            "addresses": [
                                                "QURZqPoBfuXXoHDkCHCvhGJa1DP2Ahj1KX"
                                            ]
                                        },
                                        "spent": true
                                    },
                                    {
                                        "value": "1.98831468",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a914bb29ba1df86d8507e0de32f187fe401cc3817fc388ac",
                                            "addresses": [
                                                "QdfcTR3TQHjGHcKLiHmiPLRfVPFRWazJPg"
                                            ]
                                        },
                                        "spent": true
                                    }
                                ],
                                "blockhash": "918141389623f4f880b1ba279bb70ee91de6a005e1a9d523d21a66e8c55f0502",
                                "blockheight": 393505,
                                "confirmations": 180932,
                                "time": 1560879520,
                                "blocktime": 1560879520,
                                "valueOut": "431.98831468",
                                "valueIn": "431.98984808",
                                "fees": "0.0015334",
                                "hex": "0100000002628a8e2904e1870380ccb3c8f53316cdb94f3791d519d35b4d91750d57076d50010000006b4830450221008676b34896a922ac913b0922319f98ab4bb5223f479c19d62b92f0f187b2759d02204e6c2f57337f9b3ef262e2262b3e04a971635c450fbe4278db8b2b25e5a816540121027f67fba269482ce4eecfbcc31e7acefaf17a8cb3b4f7ccfb5447c71688e1081efeffffff421977f85bdfe6ad4a7f3e476f0bd8757c1fb5db1cc4731ab785d2a42e3e2ce2000000006b483045022100e33513c3a4251a92af231c1d3e3bec9d13be2289c1fb203cb0096c6c073a8b840220206f696ee4e38da215f0b1e2c9d2005f085b8bb967d030caa70439046df64d4601210317cfe4e8b0e35f4c2c561921fe019910e3ddbe78d500a8ff546cf1843f723764fcffffff0200eeff020a0000001976a91455c89563888363b725aa0ebd8191a2e9ea54474d88ac6cedd90b000000001976a914bb29ba1df86d8507e0de32f187fe401cc3817fc388ac00000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
