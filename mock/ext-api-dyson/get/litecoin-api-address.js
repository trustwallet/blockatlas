/// Mock for external Litecoin API
/// See:
/// curl "http://{ltc rpc}/address/ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept?details=txs"
/// curl "http://localhost:3000/litecoin-api/address/ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept?details=txs"
/// curl "http://localhost:8420/v1/litecoin/address/ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept"

module.exports = {
    path: '/litecoin-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "addrStr": "ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept",
                        "balance": "0.01688078",
                        "totalReceived": "0.25679292",
                        "totalSent": "0.23991214",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 19,
                        "txs": [
                            {
                                "txid": "bfce02fdd69ec01c08486b75fa5e4a93d82a0e6f346c8eade6437affb0f817b1",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "674092ed2b166f0f21ba6287b1df23abb9612e140e77bd52f2885312c69bf867",
                                        "vout": 1,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": [
                                            "ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept"
                                        ],
                                        "value": "0.0078853"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.001",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "0014ccf05f3bc453f7f66fb48b124bd79690bc309e22",
                                            "addresses": [
                                                "ltc1qenc97w7y20mlvma53vfyh4ukjz7rp83zf892wx"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "0.00688078",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "00140ee85acd7206ba404a684e9655b54d32f242b7c5",
                                            "addresses": [
                                                "ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "ae8df4fb9cbbd07375eef217dd25e8b918a1b2746907c23cf4f2e15cd18fccdd",
                                "blockheight": 1804076,
                                "confirmations": 7531,
                                "time": 1583910871,
                                "blocktime": 1583910871,
                                "valueOut": "0.00788078",
                                "valueIn": "0.0078853",
                                "fees": "0.00000452",
                                "hex": "0100000000010167f89bc6125388f252bd770e142e61b9ab23dfb18762ba210f6f162bed92406701000000000000000002a086010000000000160014ccf05f3bc453f7f66fb48b124bd79690bc309e22ce7f0a00000000001600140ee85acd7206ba404a684e9655b54d32f242b7c502483045022100a40df50e2b419a35001f40d508e911f36d01497ebe5225dc93c81f53a30452fb022034c086517326be293019c2cbf215d6b9a53aa0ab4a0c2b890a9f1b05e8052aac012102872f9e841a8150ab716574ff897d915f575ed9abe9b4c9426617f6a1d8b3bbd100000000"
                            },
                            {
                                "txid": "674092ed2b166f0f21ba6287b1df23abb9612e140e77bd52f2885312c69bf867",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "f8ddc3eca55aa4bffd27b030fc3b35ceb8d503b31a03dafd32df56eec552fe99",
                                        "vout": 1,
                                        "sequence": 4294967290,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": [
                                            "ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept"
                                        ],
                                        "value": "0.01788982"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.01",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "00140ee85acd7206ba404a684e9655b54d32f242b7c5",
                                            "addresses": [
                                                "ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "0.0078853",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "00140ee85acd7206ba404a684e9655b54d32f242b7c5",
                                            "addresses": [
                                                "ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept"
                                            ]
                                        },
                                        "spent": true
                                    }
                                ],
                                "blockhash": "dd92763dcb8477d8f012e5ec0598ebf3c135a15780d7491754fff40d3d24dd9b",
                                "blockheight": 1803562,
                                "confirmations": 8045,
                                "time": 1583831802,
                                "blocktime": 1583831802,
                                "valueOut": "0.0178853",
                                "valueIn": "0.01788982",
                                "fees": "0.00000452",
                                "hex": "0100000000010199fe52c5ee56df32fdda031ab303d5b8ce353bfc30b027fdbfa45aa5ecc3ddf80100000000faffffff0240420f00000000001600140ee85acd7206ba404a684e9655b54d32f242b7c532080c00000000001600140ee85acd7206ba404a684e9655b54d32f242b7c502483045022100bbde918ff538468c440424323bbe8ba78eef14d9dde357c8e6496905d540098c02205700d984884edff918f827289898c16ec9ad30344a637ea85dc55f7f836143bd012102872f9e841a8150ab716574ff897d915f575ed9abe9b4c9426617f6a1d8b3bbd100000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
