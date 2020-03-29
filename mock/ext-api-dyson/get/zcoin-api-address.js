/// Mock for external Zcoin API
/// See:
/// curl "http://{Zcoin rpc}/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn?details=txs"
/// curl "http://localhost:3000/zcoin-api/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn?details=txs"
/// curl "http://localhost:8420/v1/zcoin/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"

module.exports = {
    path: '/zcoin-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "addrStr": "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn",
                        "balance": "0.6310911",
                        "totalReceived": "15.65104974",
                        "totalSent": "15.01995864",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 18,
                        "txs": [
                            {
                                "txid": "11a051ebde2108ba2b517f6513dffd34aa099c90a67a4e3e988bd5ddde9ddb45",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "86fb43e30f9bc2a8a6492e67e2b893680526987ae8b5ec56e2587381813927ef",
                                        "vout": 1,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "483045022100dde950aa0f6d73c2e78afdb66a138ff3ca481983977baad260c0632dfb2b4ee202202c840111ac82603f436bb698a8d9a50f97d0155b2d4c4576ed66ad4ad6b901020121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15"
                                        },
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "value": "0.63028838"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.001",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                            "addresses": [
                                                "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "0.6292477",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                            "addresses": [
                                                "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                            ]
                                        },
                                        "spent": true,
                                        "spentTxId": "50cb7a384e031c2823147c440884dbbf6ecef313925f113768406ab7776229fd",
                                        "spentHeight": 251156
                                    }
                                ],
                                "blockhash": "5ba0f8c0e58986d1eb5036e766e28ed541d0993cb717fc1aea46d714a89945e7",
                                "blockheight": 250864,
                                "confirmations": 1845,
                                "time": 1584571576,
                                "blocktime": 1584571576,
                                "valueOut": "0.6302477",
                                "valueIn": "0.63028838",
                                "fees": "0.00004068",
                                "hex": "0100000001ef273981817358e256ecb5e87a9826056893b8e2672e49a6a8c29b0fe343fb86010000006b483045022100dde950aa0f6d73c2e78afdb66a138ff3ca481983977baad260c0632dfb2b4ee202202c840111ac82603f436bb698a8d9a50f97d0155b2d4c4576ed66ad4ad6b901020121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b150000000002a0860100000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ace227c003000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
                            {
                                "txid": "86fb43e30f9bc2a8a6492e67e2b893680526987ae8b5ec56e2587381813927ef",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "5cf7ebaed000a76d67652ed9a91f6e1e57a6c0d74b4445a533bdda1f36be601e",
                                        "vout": 1,
                                        "sequence": 4294967294,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "483045022100b3638e753ae9136ee38906754c8d7e44e10e3f1b08dc37948ba317a5e28eb8e60220391cca6fb14eb0392c03df3bfb9639ae76a523ad7bc3f669483d5afa7ed629570121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15"
                                        },
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "value": "0.64032906"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.01",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914cffef031eead7d332c245df1372dcf4980a0127c88ac",
                                            "addresses": [
                                                "aKgF22yWfjBqekrbkhkFYqenzsw9zfRch8"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "0.63028838",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                            "addresses": [
                                                "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                            ]
                                        },
                                        "spent": true,
                                        "spentTxId": "11a051ebde2108ba2b517f6513dffd34aa099c90a67a4e3e988bd5ddde9ddb45",
                                        "spentHeight": 250864
                                    }
                                ],
                                "blockhash": "fcfea4d17cde4cfc8d113d923e7e2978513e902448f1087fd448ae4162696f4f",
                                "blockheight": 244409,
                                "confirmations": 8300,
                                "time": 1582620606,
                                "blocktime": 1582620606,
                                "valueOut": "0.64028838",
                                "valueIn": "0.64032906",
                                "fees": "0.00004068",
                                "hex": "01000000011e60be361fdabd33a545444bd7c0a6571e6e1fa9d92e65676da700d0aeebf75c010000006b483045022100b3638e753ae9136ee38906754c8d7e44e10e3f1b08dc37948ba317a5e28eb8e60220391cca6fb14eb0392c03df3bfb9639ae76a523ad7bc3f669483d5afa7ed629570121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15feffffff0240420f00000000001976a914cffef031eead7d332c245df1372dcf4980a0127c88ac66bec103000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
                            {
                                "txid": "4178eb3a84a34a9b508dbb61abbe725de115de2daea619e1cfa68eadf001fe51",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "6de60b862ef752e24f1417c255f750e8e2e2b64ce4160eccbcdd82663a36daeb",
                                        "vout": 0,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "483045022100e51595fc45741b1d43f5cef9b662605a98634644fc4281ddb471d5feee389bc402207b8b6d103f94ff1278a5d364b92b395021496c39b3f06241ca584036567fb54d012102ed835d6c0f1f4c0df53bbd728376362180010101f0ce9d79bf4ce61363d14b4b"
                                        },
                                        "addresses": [
                                            "aDVvWiM5PTv3QrUbrWQW3hPVLiFTMmzAHr"
                                        ],
                                        "value": "0.001"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.00096544",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                            "addresses": [
                                                "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "e3d64bc33182f0b74b6247c7b37a9f8ff3c6852fe1c08116d460d067919345d1",
                                "blockheight": 241112,
                                "confirmations": 11597,
                                "time": 1581616858,
                                "blocktime": 1581616858,
                                "valueOut": "0.00096544",
                                "valueIn": "0.001",
                                "fees": "0.00003456",
                                "hex": "0100000001ebda363a6682ddbccc0e16e44cb6e2e2e850f755c217144fe252f72e860be66d000000006b483045022100e51595fc45741b1d43f5cef9b662605a98634644fc4281ddb471d5feee389bc402207b8b6d103f94ff1278a5d364b92b395021496c39b3f06241ca584036567fb54d012102ed835d6c0f1f4c0df53bbd728376362180010101f0ce9d79bf4ce61363d14b4b000000000120790100000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
                            {
                                "txid": "5cf7ebaed000a76d67652ed9a91f6e1e57a6c0d74b4445a533bdda1f36be601e",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "5482cea818e70b136de121fb7a3443f4fb78b3e3f107712d7aa1dc671d0e3241",
                                        "vout": 1,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "483045022100c2a0003ab30fe8f3ffbb20e9cc38232256ce27b287c24293eff40fe4224f933502202882f50cb71e93fb432c8f244ae7dbe7f52f492452a1da68003cc0d43e9fea08012103cf5fc72d3c15e517443f8fc4295c816e5a791d7e8971943ce7c45a7bfd1ca8e8"
                                        },
                                        "addresses": [
                                            "aL96Xr2E1pMW3vcsZ5QV6BYLFE9jdCVBJL"
                                        ],
                                        "value": "3.64036974"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "3",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a9147157fe38b8108bcc0f8b6ea5cab365e2233ed0e888ac",
                                            "addresses": [
                                                "aB3mWSqhFzszMJ3h8sxTedFYDhZCDBkeQT"
                                            ]
                                        },
                                        "spent": true,
                                        "spentTxId": "736d2ed9a2398511cc98f30df178db721716c1996fd8ab67970a5ac41a795e72",
                                        "spentIndex": 2,
                                        "spentHeight": 239559
                                    },
                                    {
                                        "value": "0.64032906",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                            "addresses": [
                                                "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                            ]
                                        },
                                        "spent": true,
                                        "spentTxId": "86fb43e30f9bc2a8a6492e67e2b893680526987ae8b5ec56e2587381813927ef",
                                        "spentHeight": 244409
                                    }
                                ],
                                "blockhash": "07d7c24529ee402be1a94b6b5a70d000c6002f0fa7201c1f5ef4ef7dd87f0117",
                                "blockheight": 239551,
                                "confirmations": 13158,
                                "time": 1581138788,
                                "blocktime": 1581138788,
                                "valueOut": "3.64032906",
                                "valueIn": "3.64036974",
                                "fees": "0.00004068",
                                "hex": "010000000141320e1d67dca17a2d7107f1e3b378fbf443347afb21e16d130be718a8ce8254010000006b483045022100c2a0003ab30fe8f3ffbb20e9cc38232256ce27b287c24293eff40fe4224f933502202882f50cb71e93fb432c8f244ae7dbe7f52f492452a1da68003cc0d43e9fea08012103cf5fc72d3c15e517443f8fc4295c816e5a791d7e8971943ce7c45a7bfd1ca8e8000000000200a3e111000000001976a9147157fe38b8108bcc0f8b6ea5cab365e2233ed0e888ac8a10d103000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
                            {
                                "txid": "c4f897a1d64e6c870184ff26a3ad8f5f4a94a6667f6d3486f21e52530d8ecea2",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "3054b05d5f9624264bf7bfbb140749472c58670cc935efa5aab4a96c6430ed19",
                                        "vout": 0,
                                        "sequence": 4294967294,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "473044022062705d4b0f266a25646c410e3a735b7665dafb85655bf4b9f3ce89ecde4306f20220702f5fd1f3d6e593c3e3e9a37adf6431b1920b55151872dc9c8f4dc166df1a8f0121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15"
                                        },
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "value": "3.6504511"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "3.65041042",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914e5712d04a22c551a8261aea3f4a8d98ddb057f5988ac",
                                            "addresses": [
                                                "aMde4RuTWiNzoZN53SWZ7i7tKNbhSwt57U"
                                            ]
                                        },
                                        "spent": true,
                                        "spentTxId": "5482cea818e70b136de121fb7a3443f4fb78b3e3f107712d7aa1dc671d0e3241",
                                        "spentHeight": 238210
                                    }
                                ],
                                "blockhash": "1afa27baa3146c5ae36cb93612f928b794b306d4d8b552636fb3f35891a7adb8",
                                "blockheight": 238176,
                                "confirmations": 14533,
                                "time": 1580709451,
                                "blocktime": 1580709451,
                                "valueOut": "3.65041042",
                                "valueIn": "3.6504511",
                                "fees": "0.00004068",
                                "hex": "010000000119ed30646ca9b4aaa5ef35c90c67582c47490714bbbff74b2624965f5db05430000000006a473044022062705d4b0f266a25646c410e3a735b7665dafb85655bf4b9f3ce89ecde4306f20220702f5fd1f3d6e593c3e3e9a37adf6431b1920b55151872dc9c8f4dc166df1a8f0121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15feffffff019215c215000000001976a914e5712d04a22c551a8261aea3f4a8d98ddb057f5988ac00000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
