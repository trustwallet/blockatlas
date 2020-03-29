/// Mock for external Doge API
/// See:
/// curl "http://{doge rpc}/v2/xpub/dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN?details=txs"
/// curl "http://localhost:3000/doge-api/v2/xpub/dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN?details=txs"
/// curl "http://localhost:8420/v1/doge/xpub/dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN"

module.exports = {
    path: '/doge-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN",
                        "balance": "53537405232",
                        "totalReceived": "2701368049752",
                        "totalSent": "2647830644520",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 2,
                        "transactions": [
                            {
                                "txid": "8d5fc1e686ff042b4811cb6b2df406b98b2484e973cd663b442c296f2d2f78b9",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "f60da53d80e50e4aef26e58ca89cbbe90b554baefcf67c54b0a2648420907498",
                                        "vout": 1,
                                        "sequence": 4294967288,
                                        "n": 0,
                                        "addresses": [
                                            "DF5TedCytjJRQdvTbMPJuMwbMWoveWZYp3"
                                        ],
                                        "isAddress": true,
                                        "value": "533600000",
                                        "hex": "483045022100f357d4cb4fbef56bc76de9894e330ce485085c3cbe58dc9680db2a35dcf99dbf02204d7ef785782d75025c117c5e70ae009104b7cdd7cd1554d5cf4cb884958c06be01210391c22d3c06172377ff770f0832a4decf91ee38b975eb65451c38faa6bef861b5"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "100000000",
                                        "n": 0,
                                        "hex": "76a914054eed59ece5b6aa1ef15f3713765e708cf3d5f688ac",
                                        "addresses": [
                                            "D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "275400000",
                                        "n": 1,
                                        "hex": "76a9149ff8db746d4f6b11acab2bda9288dde3b8428d8e88ac",
                                        "addresses": [
                                            "DKix6fTygojRpBFbADfkYKDX9nZ2Y7Huqq"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "ece2c433edfaf7143ba79f381ea40ac8b7e38ce6c4ff0771ae1e3dda9246824a",
                                "blockHeight": 3099770,
                                "confirmations": 59977,
                                "blockTime": 1581351126,
                                "value": "375400000",
                                "valueIn": "533600000",
                                "fees": "158200000",
                                "hex": "0100000001987490208464a2b0547cf6fcae4b550be9bb9ca88ce526ef4a0ee5803da50df6010000006b483045022100f357d4cb4fbef56bc76de9894e330ce485085c3cbe58dc9680db2a35dcf99dbf02204d7ef785782d75025c117c5e70ae009104b7cdd7cd1554d5cf4cb884958c06be01210391c22d3c06172377ff770f0832a4decf91ee38b975eb65451c38faa6bef861b5f8ffffff0200e1f505000000001976a914054eed59ece5b6aa1ef15f3713765e708cf3d5f688ac40456a10000000001976a9149ff8db746d4f6b11acab2bda9288dde3b8428d8e88ac00000000"
                            },
                            {
                                "txid": "f60da53d80e50e4aef26e58ca89cbbe90b554baefcf67c54b0a2648420907498",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "c6a51ed2d4bb163c324f58659d98142b3fbe44503289eb6c7749414e63111dfb",
                                        "vout": 1,
                                        "sequence": 4294967291,
                                        "n": 0,
                                        "addresses": [
                                            "D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh"
                                        ],
                                        "isAddress": true,
                                        "value": "841800000",
                                        "hex": "483045022100f80b337309fcefc7886364f9a87ff9afcadf876542a5a36b543547f7e9d43ef10220397d6aaa649814b68dc98a1ed1027de34c8aa06cbe480e34bbe5a7ba5f282402012102cc82e0824da5cc96d5967cfc06d4b82e29c44ce10cb86028d15b5e67ceb1b6ca"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "150000000",
                                        "n": 0,
                                        "hex": "76a914e82178c73b5744342916f6a7a944af4fa37aebff88ac",
                                        "addresses": [
                                            "DSJVQRY3wyVPrXbEDMSeoFufWNSsMeKEuj"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "533600000",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "76a9146d01372d8c138cddcc142ade8a7ebc84948b87df88ac",
                                        "addresses": [
                                            "DF5TedCytjJRQdvTbMPJuMwbMWoveWZYp3"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "d4a57feed1683117d412944fc6581eace1bc3a8be5ebab93cf2dd99f5918374d",
                                "blockHeight": 3088989,
                                "confirmations": 70758,
                                "blockTime": 1580673745,
                                "value": "683600000",
                                "valueIn": "841800000",
                                "fees": "158200000",
                                "hex": "0100000001fb1d11634e4149776ceb89325044be3f2b14989d65584f323c16bbd4d21ea5c6010000006b483045022100f80b337309fcefc7886364f9a87ff9afcadf876542a5a36b543547f7e9d43ef10220397d6aaa649814b68dc98a1ed1027de34c8aa06cbe480e34bbe5a7ba5f282402012102cc82e0824da5cc96d5967cfc06d4b82e29c44ce10cb86028d15b5e67ceb1b6cafbffffff0280d1f008000000001976a914e82178c73b5744342916f6a7a944af4fa37aebff88ac0017ce1f000000001976a9146d01372d8c138cddcc142ade8a7ebc84948b87df88ac00000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
