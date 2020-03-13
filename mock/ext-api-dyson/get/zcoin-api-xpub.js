/// Mock for external Zcoin API
/// See:
/// curl "http://{Zcoin rpc}/v2/xpub/xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK?details=txs"
/// curl "http://localhost:3000/zcoin-api/v2/xpub/xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK?details=txs"
/// curl "http://localhost:8420/v1/zcoin/xpub/xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK"

module.exports = {
    path: '/zcoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK",
                        "balance": "63109110",
                        "totalReceived": "19765346316",
                        "totalSent": "19702237206",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 54,
                        "transactions": [
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
                                "confirmations": 1518,
                                "blockTime": 1584663676,
                                "value": "61916634",
                                "valueIn": "61920702",
                                "fees": "4068",
                                "hex": "0100000001fd296277b76a406837115f9213f3ce6ebfdb8408447c1423281c034e387acb50010000006b483045022100aff10fef2849988470ac3411bd463ae03a17a88336283290095f8247dc0531a4022064f2c0e60662effb5b46c2a70318331b007df7441a209e2bf6489924c73952f50121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15000000000280969800000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac5a2f1803000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
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
                                "confirmations": 1518,
                                "blockTime": 1584663676,
                                "value": "51912566",
                                "valueIn": "51916634",
                                "fees": "4068",
                                "hex": "010000000164b574c298c44b23fc6f328494ff306444895262d7f670ebb2a5c48233ac33cd010000006b483045022100dc5a93fa3bd1f61115062700756163da79860b38fe348aa9c3bc5d9c578c8d55022048f2f300dce9858fa3a24e05084bb892287a623b40271289cb5fce9783e77b730121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b1500000000025a17c800000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac1c085002000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
                            {
                                "txid": "50cb7a384e031c2823147c440884dbbf6ecef313925f113768406ab7776229fd",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "11a051ebde2108ba2b517f6513dffd34aa099c90a67a4e3e988bd5ddde9ddb45",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true,
                                        "value": "62924770",
                                        "hex": "4830450221009e0fe641502a598f5f470ff3a5cb4bd46e2906e0e09c7a8c93d80bfd61cfb26e0220199247204747faf99e8605e7c6af4a1a8489c2a133b92f1f547990cef62cc9f60121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1000000",
                                        "n": 0,
                                        "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "61920702",
                                        "n": 1,
                                        "spent": true,
                                        "spentTxId": "cd33ac3382c4a5b2eb70f6d7625289446430ff9484326ffc234bc498c274b564",
                                        "spentHeight": 251191,
                                        "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "1cb40b2f2d4618cd9606b0f476d2ccfeee4740bb76cd46e2c75df6e783017cb1",
                                "blockHeight": 251156,
                                "confirmations": 1553,
                                "blockTime": 1584656561,
                                "value": "62920702",
                                "valueIn": "62924770",
                                "fees": "4068",
                                "hex": "010000000145db9ddeddd58b983e4e7aa6909c09aa34fddf13657f512bba0821deeb51a011010000006b4830450221009e0fe641502a598f5f470ff3a5cb4bd46e2906e0e09c7a8c93d80bfd61cfb26e0220199247204747faf99e8605e7c6af4a1a8489c2a133b92f1f547990cef62cc9f60121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15000000000240420f00000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788acbed5b003000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
                            {
                                "txid": "11a051ebde2108ba2b517f6513dffd34aa099c90a67a4e3e988bd5ddde9ddb45",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "86fb43e30f9bc2a8a6492e67e2b893680526987ae8b5ec56e2587381813927ef",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true,
                                        "value": "63028838",
                                        "hex": "483045022100dde950aa0f6d73c2e78afdb66a138ff3ca481983977baad260c0632dfb2b4ee202202c840111ac82603f436bb698a8d9a50f97d0155b2d4c4576ed66ad4ad6b901020121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "100000",
                                        "n": 0,
                                        "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "62924770",
                                        "n": 1,
                                        "spent": true,
                                        "spentTxId": "50cb7a384e031c2823147c440884dbbf6ecef313925f113768406ab7776229fd",
                                        "spentHeight": 251156,
                                        "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "5ba0f8c0e58986d1eb5036e766e28ed541d0993cb717fc1aea46d714a89945e7",
                                "blockHeight": 250864,
                                "confirmations": 1845,
                                "blockTime": 1584571576,
                                "value": "63024770",
                                "valueIn": "63028838",
                                "fees": "4068",
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
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true,
                                        "value": "64032906",
                                        "hex": "483045022100b3638e753ae9136ee38906754c8d7e44e10e3f1b08dc37948ba317a5e28eb8e60220391cca6fb14eb0392c03df3bfb9639ae76a523ad7bc3f669483d5afa7ed629570121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1000000",
                                        "n": 0,
                                        "hex": "76a914cffef031eead7d332c245df1372dcf4980a0127c88ac",
                                        "addresses": [
                                            "aKgF22yWfjBqekrbkhkFYqenzsw9zfRch8"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "63028838",
                                        "n": 1,
                                        "spent": true,
                                        "spentTxId": "11a051ebde2108ba2b517f6513dffd34aa099c90a67a4e3e988bd5ddde9ddb45",
                                        "spentHeight": 250864,
                                        "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "fcfea4d17cde4cfc8d113d923e7e2978513e902448f1087fd448ae4162696f4f",
                                "blockHeight": 244409,
                                "confirmations": 8300,
                                "blockTime": 1582620606,
                                "value": "64028838",
                                "valueIn": "64032906",
                                "fees": "4068",
                                "hex": "01000000011e60be361fdabd33a545444bd7c0a6571e6e1fa9d92e65676da700d0aeebf75c010000006b483045022100b3638e753ae9136ee38906754c8d7e44e10e3f1b08dc37948ba317a5e28eb8e60220391cca6fb14eb0392c03df3bfb9639ae76a523ad7bc3f669483d5afa7ed629570121030987bda3c78bd6e2dd765f3767cc3611cd05c9e3c0d49873f2cd68d3fd305b15feffffff0240420f00000000001976a914cffef031eead7d332c245df1372dcf4980a0127c88ac66bec103000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
                            {
                                "txid": "4178eb3a84a34a9b508dbb61abbe725de115de2daea619e1cfa68eadf001fe51",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "6de60b862ef752e24f1417c255f750e8e2e2b64ce4160eccbcdd82663a36daeb",
                                        "n": 0,
                                        "addresses": [
                                            "aDVvWiM5PTv3QrUbrWQW3hPVLiFTMmzAHr"
                                        ],
                                        "isAddress": true,
                                        "value": "100000",
                                        "hex": "483045022100e51595fc45741b1d43f5cef9b662605a98634644fc4281ddb471d5feee389bc402207b8b6d103f94ff1278a5d364b92b395021496c39b3f06241ca584036567fb54d012102ed835d6c0f1f4c0df53bbd728376362180010101f0ce9d79bf4ce61363d14b4b"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "96544",
                                        "n": 0,
                                        "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "e3d64bc33182f0b74b6247c7b37a9f8ff3c6852fe1c08116d460d067919345d1",
                                "blockHeight": 241112,
                                "confirmations": 11597,
                                "blockTime": 1581616858,
                                "value": "96544",
                                "valueIn": "100000",
                                "fees": "3456",
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
                                        "addresses": [
                                            "aL96Xr2E1pMW3vcsZ5QV6BYLFE9jdCVBJL"
                                        ],
                                        "isAddress": true,
                                        "value": "364036974",
                                        "hex": "483045022100c2a0003ab30fe8f3ffbb20e9cc38232256ce27b287c24293eff40fe4224f933502202882f50cb71e93fb432c8f244ae7dbe7f52f492452a1da68003cc0d43e9fea08012103cf5fc72d3c15e517443f8fc4295c816e5a791d7e8971943ce7c45a7bfd1ca8e8"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "300000000",
                                        "n": 0,
                                        "spent": true,
                                        "spentTxId": "736d2ed9a2398511cc98f30df178db721716c1996fd8ab67970a5ac41a795e72",
                                        "spentIndex": 2,
                                        "spentHeight": 239559,
                                        "hex": "76a9147157fe38b8108bcc0f8b6ea5cab365e2233ed0e888ac",
                                        "addresses": [
                                            "aB3mWSqhFzszMJ3h8sxTedFYDhZCDBkeQT"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "64032906",
                                        "n": 1,
                                        "spent": true,
                                        "spentTxId": "86fb43e30f9bc2a8a6492e67e2b893680526987ae8b5ec56e2587381813927ef",
                                        "spentHeight": 244409,
                                        "hex": "76a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac",
                                        "addresses": [
                                            "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "07d7c24529ee402be1a94b6b5a70d000c6002f0fa7201c1f5ef4ef7dd87f0117",
                                "blockHeight": 239551,
                                "confirmations": 13158,
                                "blockTime": 1581138788,
                                "value": "364032906",
                                "valueIn": "364036974",
                                "fees": "4068",
                                "hex": "010000000141320e1d67dca17a2d7107f1e3b378fbf443347afb21e16d130be718a8ce8254010000006b483045022100c2a0003ab30fe8f3ffbb20e9cc38232256ce27b287c24293eff40fe4224f933502202882f50cb71e93fb432c8f244ae7dbe7f52f492452a1da68003cc0d43e9fea08012103cf5fc72d3c15e517443f8fc4295c816e5a791d7e8971943ce7c45a7bfd1ca8e8000000000200a3e111000000001976a9147157fe38b8108bcc0f8b6ea5cab365e2233ed0e888ac8a10d103000000001976a914526ac6dd84927e5e26dc6c89976d58ea4c1a4d4788ac00000000"
                            },
                            {
                                "txid": "5482cea818e70b136de121fb7a3443f4fb78b3e3f107712d7aa1dc671d0e3241",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "c4f897a1d64e6c870184ff26a3ad8f5f4a94a6667f6d3486f21e52530d8ecea2",
                                        "sequence": 4294967294,
                                        "n": 0,
                                        "addresses": [
                                            "aMde4RuTWiNzoZN53SWZ7i7tKNbhSwt57U"
                                        ],
                                        "isAddress": true,
                                        "value": "365041042",
                                        "hex": "4730440220260b40074ffa84cf4f747433efdc524a347bf9f2b995a212a7c7fafeeb9105180220222edad73b70511698a9dc68ae246b78a7ba9daf3c1b44fb5d7fdd4d8d008d890121035209cb9a9c86a775cb9593b4c6fdde4f23e4bb4193cfeae0706080702189efa3"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1000000",
                                        "n": 0,
                                        "spent": true,
                                        "spentTxId": "c70c9d3bf599a9e3553b6235f6468c69c1fa54aa0505e49d81e37aa1a749379a",
                                        "spentHeight": 238297,
                                        "hex": "76a91480be9e18824eae9cc7771291b7a017a458fb48b388ac",
                                        "addresses": [
                                            "aCTCaoLEy6ZbKQGKjBe1kprQ4foEa2cBVv"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "364036974",
                                        "n": 1,
                                        "spent": true,
                                        "spentTxId": "5cf7ebaed000a76d67652ed9a91f6e1e57a6c0d74b4445a533bdda1f36be601e",
                                        "spentHeight": 239551,
                                        "hex": "76a914d51315e0d5624c84e9da869bd34735926022ae5888ac",
                                        "addresses": [
                                            "aL96Xr2E1pMW3vcsZ5QV6BYLFE9jdCVBJL"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "fc464897e8c6bcbd5bede4109a24075894a61d3a2d7aa06cc0bc1a37b07d1903",
                                "blockHeight": 238210,
                                "confirmations": 14499,
                                "blockTime": 1580718580,
                                "value": "365036974",
                                "valueIn": "365041042",
                                "fees": "4068",
                                "hex": "0100000001a2ce8e0d53521ef286346d7f66a6944a5f8fada326ff8401876c4ed6a197f8c4000000006a4730440220260b40074ffa84cf4f747433efdc524a347bf9f2b995a212a7c7fafeeb9105180220222edad73b70511698a9dc68ae246b78a7ba9daf3c1b44fb5d7fdd4d8d008d890121035209cb9a9c86a775cb9593b4c6fdde4f23e4bb4193cfeae0706080702189efa3feffffff0240420f00000000001976a91480be9e18824eae9cc7771291b7a017a458fb48b388ac6ec3b215000000001976a914d51315e0d5624c84e9da869bd34735926022ae5888ac00000000"
                            }
                        ],
                        "usedTokens": 52,
                        "tokens": [
                            {
                                "type": "XPUBAddress",
                                "name": "a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn",
                                "path": "m/44'/136'/0'/0/0",
                                "transfers": 18,
                                "decimals": 8,
                                "balance": "63109110",
                                "totalReceived": "1565104974",
                                "totalSent": "1501995864"
                            }
                        ]
                    }                
                `);
        }
        return {error: "Not implemented"};
    }
}
