/// Mock for external Qtum API
/// See:
/// curl "http://{qtum rpc}/v2/xpub/xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd?details=txs"
/// curl "http://localhost:3000/qtum-api/v2/xpub/xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd?details=txs"
/// curl "http://localhost:8420/v1/qtum/xpub/xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd"

module.exports = {
    path: '/qtum-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd",
                        "balance": "235357610",
                        "totalReceived": "78913733296",
                        "totalSent": "78678375686",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 37,
                        "transactions": [
                            {
                                "txid": "62438bb658856c3a08b89ca80e199e7031f98956ea86a135f5d6306660230f67",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "782e18a7b1bf612972e773fd6ccdcbe0c4514c661c31a7942f5dcd4807c5ab68",
                                        "sequence": 4294967293,
                                        "n": 0,
                                        "addresses": [
                                            "QWLQSMPF5WKAqhXxePJvjUJNjxDHrPCbjB"
                                        ],
                                        "isAddress": true,
                                        "value": "100000",
                                        "hex": "483045022100b660d73ff3bd7cc52d1052b933828cc2bb23b7a94500dff4bf00d5640ff4951702206ac6ecdaf3818dc57d54f8e860bbf212846200ccb2d1a323bf5b59868dcba872012102d7062f4af80f3a31f67c928e141627d807f0641b0bb0d466f4a62ee31230aa1e"
                                    },
                                    {
                                        "txid": "80b7c6bcba625ad4abb69f7e92521a782681e984982408d40ad35a24b3f78297",
                                        "sequence": 4294967294,
                                        "n": 1,
                                        "addresses": [
                                            "QgqVXuCCN1m4Hfb9HCkrzbgUSkoodKsrVY"
                                        ],
                                        "isAddress": true,
                                        "value": "100000",
                                        "hex": "48304502210090ea4a4aac51d8734ad9ffecf0e5c14be50a8941b8cac6a492c002a0be99e5e102203550b52b7ad546adf8b1e5ed00b7aa755cc296d4bd61004b2ace43bfe0007d17012103692f9ee47370b3a70e8d7b6ac36c705b2d3296bb37c0f02e15158f3a6001ce5b"
                                    },
                                    {
                                        "txid": "9c3a6eb9fa07868a2e59530fe2aa3ea29a21d3818a70343097adbe571ee2f0be",
                                        "sequence": 4294967292,
                                        "n": 2,
                                        "addresses": [
                                            "QSEXLG5qct47UtMHiXNg2Y27KSDSAEHJp9"
                                        ],
                                        "isAddress": true,
                                        "value": "10000000",
                                        "hex": "473044022055806e0939aa734d296ac364d4edf430bcd4fc15a377d35205362d09f7367d8e02202cc4e691c31ece66abb9adcd20bf8e3f500a2254edaaf2c241529cee670972440121025249520b3d48efd435e6cc30c9af238d054461d51f8beeaf75c4de50a647acbc"
                                    },
                                    {
                                        "txid": "9c3a6eb9fa07868a2e59530fe2aa3ea29a21d3818a70343097adbe571ee2f0be",
                                        "vout": 1,
                                        "sequence": 4294967291,
                                        "n": 3,
                                        "addresses": [
                                            "Qfw9Poi2c4gGEkAAuVyG21PsBmNGsmZKFC"
                                        ],
                                        "isAddress": true,
                                        "value": "225494620",
                                        "hex": "483045022100e9d2126507af25eed231d27db52794bf6dc67c33800ebafd99b44b13394e1a2d02206b5abdce728114fa75624cd50d70f3c7da77c5b00433fa306b203e4bfb7d96320121020d708ba820b87e195b5790fa2723033cfec259ce5af4270a431cd7297d7704d3"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "235357610",
                                        "n": 0,
                                        "hex": "76a9148b4f969602a4579f9c982e7d30c3a06e2f0cacc888ac",
                                        "addresses": [
                                            "QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "6e71fa96ca29a6e4cbd10defaa4fc69ff09adec45fc666d82afade3f75206cc1",
                                "blockHeight": 563426,
                                "confirmations": 11012,
                                "blockTime": 1583720608,
                                "value": "235357610",
                                "valueIn": "235694620",
                                "fees": "337010",
                                "hex": "010000000468abc50748cd5d2f94a7311c664c51c4e0cbcd6cfd73e7722961bfb1a7182e78000000006b483045022100b660d73ff3bd7cc52d1052b933828cc2bb23b7a94500dff4bf00d5640ff4951702206ac6ecdaf3818dc57d54f8e860bbf212846200ccb2d1a323bf5b59868dcba872012102d7062f4af80f3a31f67c928e141627d807f0641b0bb0d466f4a62ee31230aa1efdffffff9782f7b3245ad30ad408249884e98126781a52927e9fb6abd45a62babcc6b780000000006b48304502210090ea4a4aac51d8734ad9ffecf0e5c14be50a8941b8cac6a492c002a0be99e5e102203550b52b7ad546adf8b1e5ed00b7aa755cc296d4bd61004b2ace43bfe0007d17012103692f9ee47370b3a70e8d7b6ac36c705b2d3296bb37c0f02e15158f3a6001ce5bfeffffffbef0e21e57bead973034708a81d3219aa23eaae20f53592e8a8607fab96e3a9c000000006a473044022055806e0939aa734d296ac364d4edf430bcd4fc15a377d35205362d09f7367d8e02202cc4e691c31ece66abb9adcd20bf8e3f500a2254edaaf2c241529cee670972440121025249520b3d48efd435e6cc30c9af238d054461d51f8beeaf75c4de50a647acbcfcffffffbef0e21e57bead973034708a81d3219aa23eaae20f53592e8a8607fab96e3a9c010000006b483045022100e9d2126507af25eed231d27db52794bf6dc67c33800ebafd99b44b13394e1a2d02206b5abdce728114fa75624cd50d70f3c7da77c5b00433fa306b203e4bfb7d96320121020d708ba820b87e195b5790fa2723033cfec259ce5af4270a431cd7297d7704d3fbffffff01aa45070e000000001976a9148b4f969602a4579f9c982e7d30c3a06e2f0cacc888ac00000000"
                            },
                            {
                                "txid": "9c3a6eb9fa07868a2e59530fe2aa3ea29a21d3818a70343097adbe571ee2f0be",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "1d597a91e1810fdbfcb712c11c4776f5384ab4e7bd94a2fdc6538b9afdc1e6d0",
                                        "vout": 1,
                                        "sequence": 4294967292,
                                        "n": 0,
                                        "addresses": [
                                            "QZsdyvpJnczcDPJoySkNTaSh8VP8XAXeMz"
                                        ],
                                        "isAddress": true,
                                        "value": "235607620",
                                        "hex": "4730440220679ad69958629be2516a0d93457ba24f154fd42e23a4f3f2896272fe97940eed0220395098fa364232f74440e50043d6b2868829cb11faf56057d12f68f4bac82bdc0121036ea9140776d2e6a3ec7dd32b9795b47cf592e90e4159f7db793019b069e78b46"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "10000000",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a9143dc1aae7701ab1f7be54fc69808c092ddb5a574088ac",
                                        "addresses": [
                                            "QSEXLG5qct47UtMHiXNg2Y27KSDSAEHJp9"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "225494620",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "76a914d40a0b271f69b1d2f43f60c540325c9039b4a05588ac",
                                        "addresses": [
                                            "Qfw9Poi2c4gGEkAAuVyG21PsBmNGsmZKFC"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "a7a05f41f5f07faa5d3ca2cf5fb7de6a0cf46bc4ee465028ceda72ea5e6b5669",
                                "blockHeight": 479076,
                                "confirmations": 95362,
                                "blockTime": 1572922480,
                                "value": "235494620",
                                "valueIn": "235607620",
                                "fees": "113000",
                                "hex": "0100000001d0e6c1fd9a8b53c6fda294bde7b44a38f576471cc112b7fcdb0f81e1917a591d010000006a4730440220679ad69958629be2516a0d93457ba24f154fd42e23a4f3f2896272fe97940eed0220395098fa364232f74440e50043d6b2868829cb11faf56057d12f68f4bac82bdc0121036ea9140776d2e6a3ec7dd32b9795b47cf592e90e4159f7db793019b069e78b46fcffffff0280969800000000001976a9143dc1aae7701ab1f7be54fc69808c092ddb5a574088ac5cc6700d000000001976a914d40a0b271f69b1d2f43f60c540325c9039b4a05588ac00000000"
                            }
                        ],
                        "usedTokens": 33,
                        "tokens": [
                            {
                                "type": "XPUBAddress",
                                "name": "QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ",
                                "path": "m/44'/2301'/0'/0/0",
                                "transfers": 5,
                                "decimals": 8,
                                "balance": "235357610",
                                "totalReceived": "43335357610",
                                "totalSent": "43100000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
