/// Mock for external Doge API
/// See:
/// curl "http://{doge rpc}/api/v2/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh?details=txs"
/// curl "http://localhost:3347/doge-api/v2/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh?details=txs"
/// curl "http://localhost:8437/v1/doge/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh"

module.exports = {
    path: '/doge-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 6,
                    "itemsOnPage": 2,
                    "address": "D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh",
                    "balance": "50394795942",
                    "totalReceived": "160202987826",
                    "totalSent": "109808191884",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 12,
                    "transactions": [
                        {
                            "txid": "fee8381fed7406a48a8ea9e62328a3134064210626e81489ed8906e18c433bdf",
                            "version": 1,
                            "vin": [
                                {
                                    "txid": "c6a51ed2d4bb163c324f58659d98142b3fbe44503289eb6c7749414e63111dfb",
                                    "n": 0,
                                    "addresses": [
                                        "D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh"
                                    ],
                                    "isAddress": true,
                                    "value": "1000000000",
                                    "hex": "47304402200833c3184e768e5b6da836d9824f98b65704a24229a02e572ac760bb304ab9fb022068d102da9400afaeee3a6b42930a966ff107f5c68affde4a0239dab9fa21123d012102cc82e0824da5cc96d5967cfc06d4b82e29c44ce10cb86028d15b5e67ceb1b6ca"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "134700000",
                                    "n": 0,
                                    "hex": "76a9142ba977acdb30bc18d350df32f38a777313ae86b088ac",
                                    "addresses": [
                                        "D97xcKB9EAPmgNExbB5nAXDAQ9s5ZJ91J4"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "707100000",
                                    "n": 1,
                                    "hex": "76a914b358390833fd8371733c1d11477bdb2d185e61e988ac",
                                    "addresses": [
                                        "DMVPBRTD6bgqr3fk2yTdz97hS4AA2gzj8g"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "b94fd58d99710000ff0991ebf82c85330ed706ea18bcd4265e1f22747b91ddfe",
                            "blockHeight": 3170421,
                            "confirmations": 30293,
                            "blockTime": 1585794154,
                            "value": "841800000",
                            "valueIn": "1000000000",
                            "fees": "158200000",
                            "hex": "0100000001fb1d11634e4149776ceb89325044be3f2b14989d65584f323c16bbd4d21ea5c6000000006a47304402200833c3184e768e5b6da836d9824f98b65704a24229a02e572ac760bb304ab9fb022068d102da9400afaeee3a6b42930a966ff107f5c68affde4a0239dab9fa21123d012102cc82e0824da5cc96d5967cfc06d4b82e29c44ce10cb86028d15b5e67ceb1b6ca0000000002e05b0708000000001976a9142ba977acdb30bc18d350df32f38a777313ae86b088ac607d252a000000001976a914b358390833fd8371733c1d11477bdb2d185e61e988ac00000000"
                        },
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
                            "confirmations": 100944,
                            "blockTime": 1581351126,
                            "value": "375400000",
                            "valueIn": "533600000",
                            "fees": "158200000",
                            "hex": "0100000001987490208464a2b0547cf6fcae4b550be9bb9ca88ce526ef4a0ee5803da50df6010000006b483045022100f357d4cb4fbef56bc76de9894e330ce485085c3cbe58dc9680db2a35dcf99dbf02204d7ef785782d75025c117c5e70ae009104b7cdd7cd1554d5cf4cb884958c06be01210391c22d3c06172377ff770f0832a4decf91ee38b975eb65451c38faa6bef861b5f8ffffff0200e1f505000000001976a914054eed59ece5b6aa1ef15f3713765e708cf3d5f688ac40456a10000000001976a9149ff8db746d4f6b11acab2bda9288dde3b8428d8e88ac00000000"
                        }
                    ]
                }
                `);
        }
        return {error: "Not implemented"};
    }
}
