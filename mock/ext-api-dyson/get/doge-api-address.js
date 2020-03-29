/// Mock for external Doge API
/// See:
/// curl "http://{doge rpc}/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh?details=txs"
/// curl "http://localhost:3000/doge-api/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh?details=txs"
/// curl "http://localhost:8420/v1/doge/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh"

module.exports = {
    path: '/doge-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "addrStr": "D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh",
                        "balance": "513.94795942",
                        "totalReceived": "1602.02987826",
                        "totalSent": "1088.08191884",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 11,
                        "txs": [
                            {
                                "txid": "8d5fc1e686ff042b4811cb6b2df406b98b2484e973cd663b442c296f2d2f78b9",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "f60da53d80e50e4aef26e58ca89cbbe90b554baefcf67c54b0a2648420907498",
                                        "vout": 1,
                                        "sequence": 4294967288,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "483045022100f357d4cb4fbef56bc76de9894e330ce485085c3cbe58dc9680db2a35dcf99dbf02204d7ef785782d75025c117c5e70ae009104b7cdd7cd1554d5cf4cb884958c06be01210391c22d3c06172377ff770f0832a4decf91ee38b975eb65451c38faa6bef861b5"
                                        },
                                        "addresses": [
                                            "DF5TedCytjJRQdvTbMPJuMwbMWoveWZYp3"
                                        ],
                                        "value": "5.336"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914054eed59ece5b6aa1ef15f3713765e708cf3d5f688ac",
                                            "addresses": [
                                                "D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "2.754",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a9149ff8db746d4f6b11acab2bda9288dde3b8428d8e88ac",
                                            "addresses": [
                                                "DKix6fTygojRpBFbADfkYKDX9nZ2Y7Huqq"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "ece2c433edfaf7143ba79f381ea40ac8b7e38ce6c4ff0771ae1e3dda9246824a",
                                "blockheight": 3099770,
                                "confirmations": 59971,
                                "time": 1581351126,
                                "blocktime": 1581351126,
                                "valueOut": "3.754",
                                "valueIn": "5.336",
                                "fees": "1.582",
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
                                        "scriptSig": {
                                            "hex": "483045022100f80b337309fcefc7886364f9a87ff9afcadf876542a5a36b543547f7e9d43ef10220397d6aaa649814b68dc98a1ed1027de34c8aa06cbe480e34bbe5a7ba5f282402012102cc82e0824da5cc96d5967cfc06d4b82e29c44ce10cb86028d15b5e67ceb1b6ca"
                                        },
                                        "addresses": [
                                            "D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh"
                                        ],
                                        "value": "8.418"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1.5",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914e82178c73b5744342916f6a7a944af4fa37aebff88ac",
                                            "addresses": [
                                                "DSJVQRY3wyVPrXbEDMSeoFufWNSsMeKEuj"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "5.336",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a9146d01372d8c138cddcc142ade8a7ebc84948b87df88ac",
                                            "addresses": [
                                                "DF5TedCytjJRQdvTbMPJuMwbMWoveWZYp3"
                                            ]
                                        },
                                        "spent": true
                                    }
                                ],
                                "blockhash": "d4a57feed1683117d412944fc6581eace1bc3a8be5ebab93cf2dd99f5918374d",
                                "blockheight": 3088989,
                                "confirmations": 70752,
                                "time": 1580673745,
                                "blocktime": 1580673745,
                                "valueOut": "6.836",
                                "valueIn": "8.418",
                                "fees": "1.582",
                                "hex": "0100000001fb1d11634e4149776ceb89325044be3f2b14989d65584f323c16bbd4d21ea5c6010000006b483045022100f80b337309fcefc7886364f9a87ff9afcadf876542a5a36b543547f7e9d43ef10220397d6aaa649814b68dc98a1ed1027de34c8aa06cbe480e34bbe5a7ba5f282402012102cc82e0824da5cc96d5967cfc06d4b82e29c44ce10cb86028d15b5e67ceb1b6cafbffffff0280d1f008000000001976a914e82178c73b5744342916f6a7a944af4fa37aebff88ac0017ce1f000000001976a9146d01372d8c138cddcc142ade8a7ebc84948b87df88ac00000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
