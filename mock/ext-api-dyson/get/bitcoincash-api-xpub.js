/// Mock for external Bitcoincash API
/// See:
/// curl "http://{bch rpc}/v2/xpub/xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX?details=txs"
/// curl "http://localhost:3000/bitcoincash-api/v2/xpub/xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX?details=txs"
/// curl "http://localhost:8420/v1/bitcoincash/xpub/xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX"

module.exports = {
    path: '/bitcoincash-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX",
                        "balance": "177672",
                        "totalReceived": "13221795",
                        "totalSent": "13044123",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 72,
                        "transactions": [
                            {
                                "txid": "269d428f01fbe49cd6d2c2ca5e6e2f0ff68aece905313932156078d4341d347a",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "5536327a38fddb32a92a98582aca4e85b4e50f40b9d3c03db360eaca7b512f97",
                                        "n": 0,
                                        "addresses": [
                                            "bitcoincash:qzvqpsyaupwcu24dear47va7m96gr3rf4g9gcq4h25"
                                        ],
                                        "isAddress": true,
                                        "value": "188124",
                                        "hex": "473044022057e811d95dfce16c7f1d1da574ab4ca1bf883359da34a1e7a380f5045fcffc960220149c357343f0172301f06b914a2844d3755a3e990a236ae33f71afed52f0849d41210259dc89eeb4c7b349faab20e932be226a1ed7d582f25b64e9abf84f73e1d25fd5"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "10000",
                                        "n": 0,
                                        "hex": "76a9147222fa3e0a256cc45823a641aa4060e66718276288ac",
                                        "addresses": [
                                            "bitcoincash:qpez9737pgjke3zcywnyr2jqvrnxwxp8vgu3nnxf6x"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "177672",
                                        "n": 1,
                                        "hex": "76a9149800c09de05d8e2aadcf475f33bed97481c469aa88ac",
                                        "addresses": [
                                            "bitcoincash:qzvqpsyaupwcu24dear47va7m96gr3rf4g9gcq4h25"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "000000000000000000e0846c675bda2ea268eaa290287aef83f8877c40f07743",
                                "blockHeight": 620315,
                                "confirmations": 7640,
                                "blockTime": 1580517507,
                                "value": "187672",
                                "valueIn": "188124",
                                "fees": "452",
                                "hex": "0100000001972f517bcaea60b33dc0d3b9400fe5b4854eca2a58982aa932dbfd387a323655000000006a473044022057e811d95dfce16c7f1d1da574ab4ca1bf883359da34a1e7a380f5045fcffc960220149c357343f0172301f06b914a2844d3755a3e990a236ae33f71afed52f0849d41210259dc89eeb4c7b349faab20e932be226a1ed7d582f25b64e9abf84f73e1d25fd5000000000210270000000000001976a9147222fa3e0a256cc45823a641aa4060e66718276288ac08b60200000000001976a9149800c09de05d8e2aadcf475f33bed97481c469aa88ac00000000"
                            },
                            {
                                "txid": "5536327a38fddb32a92a98582aca4e85b4e50f40b9d3c03db360eaca7b512f97",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "feac34105f3f0bd442f75acaff6b2e840b9e2c0e32e897393d9d0fe47fe76ba1",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "bitcoincash:qzvqpsyaupwcu24dear47va7m96gr3rf4g9gcq4h25"
                                        ],
                                        "isAddress": true,
                                        "value": "38804",
                                        "hex": "483045022100dbbf18561baf7fdb46c21d27801594209f2067f1d7735e245c56001c4e462d3c0220016e052ebbcf350bd64a257f7cc52ef23730fe50b24355d098f87c528fcd839641210259dc89eeb4c7b349faab20e932be226a1ed7d582f25b64e9abf84f73e1d25fd5"
                                    },
                                    {
                                        "txid": "01b8958ff2a939534269d05af5ac747238eb3a8c91b2a564107af20e4de55680",
                                        "n": 1,
                                        "addresses": [
                                            "bitcoincash:qzvqpsyaupwcu24dear47va7m96gr3rf4g9gcq4h25"
                                        ],
                                        "isAddress": true,
                                        "value": "150000",
                                        "hex": "47304402202100334789951706405faa9ddc191702b88d9438d5a461d1f008a9bd097bd9b0022062cd93003ad63bfbb16386b90e619c2b32d9caef5229d8943c11309f26b6c19f41210259dc89eeb4c7b349faab20e932be226a1ed7d582f25b64e9abf84f73e1d25fd5"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "188124",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a9149800c09de05d8e2aadcf475f33bed97481c469aa88ac",
                                        "addresses": [
                                            "bitcoincash:qzvqpsyaupwcu24dear47va7m96gr3rf4g9gcq4h25"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "00000000000000000028351f97cd867acbba39b5a81137ce35a1ac18444aaa76",
                                "blockHeight": 619773,
                                "confirmations": 8182,
                                "blockTime": 1580201823,
                                "value": "188124",
                                "valueIn": "188804",
                                "fees": "680",
                                "hex": "0100000002a16be77fe40f9d3d3997e8320e2c9e0b842e6bffca5af742d40b3f5f1034acfe010000006b483045022100dbbf18561baf7fdb46c21d27801594209f2067f1d7735e245c56001c4e462d3c0220016e052ebbcf350bd64a257f7cc52ef23730fe50b24355d098f87c528fcd839641210259dc89eeb4c7b349faab20e932be226a1ed7d582f25b64e9abf84f73e1d25fd5000000008056e54d0ef27a1064a5b2918c3aeb387274acf55ad069425339a9f28f95b801000000006a47304402202100334789951706405faa9ddc191702b88d9438d5a461d1f008a9bd097bd9b0022062cd93003ad63bfbb16386b90e619c2b32d9caef5229d8943c11309f26b6c19f41210259dc89eeb4c7b349faab20e932be226a1ed7d582f25b64e9abf84f73e1d25fd50000000001dcde0200000000001976a9149800c09de05d8e2aadcf475f33bed97481c469aa88ac00000000"
                            },
                            {
                                "txid": "feac34105f3f0bd442f75acaff6b2e840b9e2c0e32e897393d9d0fe47fe76ba1",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "6f2d6fb46039f192f82cfd9bc85dcf0575fdf5eda9b726643314be883c8c0559",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "bitcoincash:qzvqpsyaupwcu24dear47va7m96gr3rf4g9gcq4h25"
                                        ],
                                        "isAddress": true,
                                        "value": "49256",
                                        "hex": "47304402207df4bd44f74e9a077b84f547b29c595e808bfb3ebc23e070e975905e2012044c02206fa66e30f9afe063c5a5c30d0bb20c6d0eece782a9b6211676180c25d745e98c41210259dc89eeb4c7b349faab20e932be226a1ed7d582f25b64e9abf84f73e1d25fd5"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "10000",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a914045891bcbb214544c47a8dbc75350584d8042bc288ac",
                                        "addresses": [
                                            "bitcoincash:qqz93yduhvs523xy02xmcaf4qkzdspptcgq4dscr8c"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "38804",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "76a9149800c09de05d8e2aadcf475f33bed97481c469aa88ac",
                                        "addresses": [
                                            "bitcoincash:qzvqpsyaupwcu24dear47va7m96gr3rf4g9gcq4h25"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "0000000000000000015997ed69ac8b1e1d19c006bca7fb08faf3442890b2fed5",
                                "blockHeight": 619764,
                                "confirmations": 8191,
                                "blockTime": 1580198433,
                                "value": "48804",
                                "valueIn": "49256",
                                "fees": "452",
                                "hex": "010000000159058c3c88be14336426b7a9edf5fd7505cf5dc89bfd2cf892f13960b46f2d6f010000006a47304402207df4bd44f74e9a077b84f547b29c595e808bfb3ebc23e070e975905e2012044c02206fa66e30f9afe063c5a5c30d0bb20c6d0eece782a9b6211676180c25d745e98c41210259dc89eeb4c7b349faab20e932be226a1ed7d582f25b64e9abf84f73e1d25fd5000000000210270000000000001976a914045891bcbb214544c47a8dbc75350584d8042bc288ac94970000000000001976a9149800c09de05d8e2aadcf475f33bed97481c469aa88ac00000000"
                            }
                        ],
                        "usedTokens": 71,
                        "tokens": [
                            {
                                "type": "XPUBAddress",
                                "name": "bitcoincash:qzvqpsyaupwcu24dear47va7m96gr3rf4g9gcq4h25",
                                "path": "m/44'/145'/0'/0/0",
                                "transfers": 11,
                                "decimals": 8,
                                "balance": "177672",
                                "totalReceived": "850002",
                                "totalSent": "672330"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
