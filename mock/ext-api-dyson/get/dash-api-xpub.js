/// Mock for external Dash API
/// See:
/// curl "http://{dash rpc}/v2/xpub/xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD?details=txs"
/// curl "http://localhost:3000/dash-api/v2/xpub/xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD?details=txs"
/// curl "http://localhost:8420/v1/dash/xpub/xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD"

module.exports = {
    path: '/dash-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD",
                        "balance": "2597761",
                        "totalReceived": "27175423",
                        "totalSent": "24577662",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 52,
                        "transactions": [
                            {
                                "txid": "2b67e2fbe6a212286243bc539cca3c1d877e85ffec7c925e34f2bfb7b6cc498c",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "227f1995f5d0b0adcd4f014159710892fa4b66a02ff7ad1e0fc4c1b84ca27cb2",
                                        "n": 0,
                                        "addresses": [
                                            "XyPpEePUKruNEVgdp5jWSakfvoQTnkZxhL"
                                        ],
                                        "isAddress": true,
                                        "value": "34508",
                                        "hex": "483045022100bebce18899214200115fbb2e3b16cd7a6b61cca2cf3e1ffc3b26872f66b714f702203316ce9f1cfb7cef751dbba46015ea76d3c409d6b913b6e434553ebe66680713012103f556051304c43ef031a17615fd8b056c1ac7b0a50926137bf918e72be002f606"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "10000",
                                        "n": 0,
                                        "hex": "76a9149b80787d880fd9f6a822fad78037c9563790a6fa88ac",
                                        "addresses": [
                                            "Xps4UHWgH1Ge9HAJijuJMA1i6Ybhq8JNkq"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "18858",
                                        "n": 1,
                                        "hex": "76a9149b80787d880fd9f6a822fad78037c9563790a6fa88ac",
                                        "addresses": [
                                            "Xps4UHWgH1Ge9HAJijuJMA1i6Ybhq8JNkq"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "000000000000000b3e9287157ae29c1c19bbd1e65660a5bbe723261d0fc57737",
                                "blockHeight": 1226578,
                                "confirmations": 16630,
                                "blockTime": 1582509671,
                                "value": "28858",
                                "valueIn": "34508",
                                "fees": "5650",
                                "hex": "0100000001b27ca24cb8c1c40f1eadf72fa0664bfa9208715941014fcdadb0d0f595197f22000000006b483045022100bebce18899214200115fbb2e3b16cd7a6b61cca2cf3e1ffc3b26872f66b714f702203316ce9f1cfb7cef751dbba46015ea76d3c409d6b913b6e434553ebe66680713012103f556051304c43ef031a17615fd8b056c1ac7b0a50926137bf918e72be002f606000000000210270000000000001976a9149b80787d880fd9f6a822fad78037c9563790a6fa88acaa490000000000001976a9149b80787d880fd9f6a822fad78037c9563790a6fa88ac00000000"
                            },
                            {
                                "txid": "dc8e1219da91a24c2eca773a6853d7a0602eed509a4b71afd553f8b4c218b309",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "501c3c7ecaae272de760e81d284e44644a0475c42e3ee15ae16a125650211c8e",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "Xps4UHWgH1Ge9HAJijuJMA1i6Ybhq8JNkq"
                                        ],
                                        "isAddress": true,
                                        "value": "768334",
                                        "hex": "483045022100805ae010fb81b3659f687d1f9da9efa27fa220d34c0f9c54dde4f8e8358e7c2e022003c2c89eb0372616ad50b74f627097e3ccf2ecd940bad616b0bf7caaab58615c01210359d6639a8c603a3ecf0384373935c90a93477668c148929be3b5752fcdf64600"
                                    },
                                    {
                                        "txid": "501c3c7ecaae272de760e81d284e44644a0475c42e3ee15ae16a125650211c8e",
                                        "n": 1,
                                        "addresses": [
                                            "Xps4UHWgH1Ge9HAJijuJMA1i6Ybhq8JNkq"
                                        ],
                                        "isAddress": true,
                                        "value": "809686",
                                        "hex": "483045022100ff644abb86e4272384e64de00d7ac2b1ed01e6848f59707f1bd93913b0f2ceeb02205ba0be94a02a2af3d7178977d828dda4268da4f83c00d2af1331e4274a5454a201210359d6639a8c603a3ecf0384373935c90a93477668c148929be3b5752fcdf64600"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1000000",
                                        "n": 0,
                                        "hex": "76a9149b80787d880fd9f6a822fad78037c9563790a6fa88ac",
                                        "addresses": [
                                            "Xps4UHWgH1Ge9HAJijuJMA1i6Ybhq8JNkq"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "576150",
                                        "n": 1,
                                        "hex": "76a9149b80787d880fd9f6a822fad78037c9563790a6fa88ac",
                                        "addresses": [
                                            "Xps4UHWgH1Ge9HAJijuJMA1i6Ybhq8JNkq"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "000000000000000f727266a1a157af489fec0a009938dadeca89bc3edd1bc017",
                                "blockHeight": 1215128,
                                "confirmations": 28080,
                                "blockTime": 1580704472,
                                "value": "1576150",
                                "valueIn": "1578020",
                                "fees": "1870",
                                "hex": "01000000028e1c215056126ae15ae13e2ec475044a64444e281de860e72d27aeca7e3c1c50010000006b483045022100805ae010fb81b3659f687d1f9da9efa27fa220d34c0f9c54dde4f8e8358e7c2e022003c2c89eb0372616ad50b74f627097e3ccf2ecd940bad616b0bf7caaab58615c01210359d6639a8c603a3ecf0384373935c90a93477668c148929be3b5752fcdf64600000000008e1c215056126ae15ae13e2ec475044a64444e281de860e72d27aeca7e3c1c50000000006b483045022100ff644abb86e4272384e64de00d7ac2b1ed01e6848f59707f1bd93913b0f2ceeb02205ba0be94a02a2af3d7178977d828dda4268da4f83c00d2af1331e4274a5454a201210359d6639a8c603a3ecf0384373935c90a93477668c148929be3b5752fcdf64600000000000240420f00000000001976a9149b80787d880fd9f6a822fad78037c9563790a6fa88ac96ca0800000000001976a9149b80787d880fd9f6a822fad78037c9563790a6fa88ac00000000"
                            }
                        ],
                        "usedTokens": 47,
                        "tokens": [
                            {
                                "type": "XPUBAddress",
                                "name": "Xps4UHWgH1Ge9HAJijuJMA1i6Ybhq8JNkq",
                                "path": "m/44'/5'/0'/0/0",
                                "transfers": 13,
                                "decimals": 8,
                                "balance": "2543095",
                                "totalReceived": "4275043",
                                "totalSent": "1731948"
                            },
                            {
                                "type": "XPUBAddress",
                                "name": "XezMSYZ3ucD2tNGVME1j3A8tM5q3NvigKu",
                                "path": "m/44'/5'/0'/0/15",
                                "transfers": 1,
                                "decimals": 8,
                                "balance": "10000",
                                "totalReceived": "10000",
                                "totalSent": "0"
                            },
                            {
                                "type": "XPUBAddress",
                                "name": "XmM5KX1Zpepji5V6RF7iypCkXBi4hSronS",
                                "path": "m/44'/5'/0'/0/18",
                                "transfers": 1,
                                "decimals": 8,
                                "balance": "10000",
                                "totalReceived": "10000",
                                "totalSent": "0"
                            },
                            {
                                "type": "XPUBAddress",
                                "name": "XcLBw1EUz5HPDJ5pZrSbeyXYBWcVCabZ56",
                                "path": "m/44'/5'/0'/1/24",
                                "transfers": 1,
                                "decimals": 8,
                                "balance": "34666",
                                "totalReceived": "34666",
                                "totalSent": "0"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
