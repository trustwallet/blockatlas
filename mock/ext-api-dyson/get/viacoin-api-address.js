/// Mock for external Viacoin API
/// See:
/// curl "http://{Viacoin rpc}/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A?details=txs"
/// curl "http://localhost:3000/viacoin-api/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A?details=txs"
/// curl "http://localhost:8420/v1/viacoin/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A"

module.exports = {
    path: '/viacoin-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 2,
                        "addrStr": "VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A",
                        "balance": "737.5107688",
                        "totalReceived": "183234.99094131",
                        "totalSent": "182497.48017251",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 1523480,
                        "txs": [
                            {
                                "txid": "dadbe1f13dc5ecd1cfb0d8ae139a4e55493eb9ab2c937cb78b67a5d9c1ded7e8",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": null,
                                        "value": ""
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.00976562",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a91424cc424c1e5e977175d2b20012554d39024bd68f88ac",
                                            "addresses": [
                                                "VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "496ff7e4cf450449ade88fc8bf214fd39a6d4203da8f8256492c589bf28de1a7",
                                "blockheight": 7423549,
                                "confirmations": 6,
                                "time": 1584742971,
                                "blocktime": 1584742971,
                                "valueOut": "0.00976562",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff09033d4671045e75423bffffffff01b2e60e00000000001976a91424cc424c1e5e977175d2b20012554d39024bd68f88ac00000000"
                            },
                            {
                                "txid": "be2a7b573af425f8f309561b3db4631f905626057da94399d7e44f97595c7dbe",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": null,
                                        "value": ""
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.00976562",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a91424cc424c1e5e977175d2b20012554d39024bd68f88ac",
                                            "addresses": [
                                                "VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "36c57ff80410c9c32790b75e566fdf855c6302a6a09f6dc75b362f9c6b3c884c",
                                "blockheight": 7423545,
                                "confirmations": 10,
                                "time": 1584742872,
                                "blocktime": 1584742872,
                                "valueOut": "0.00976562",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0903394671045e7541d8ffffffff01b2e60e00000000001976a91424cc424c1e5e977175d2b20012554d39024bd68f88ac00000000"
                            }
                        ]
                    }                
                `);
        }
        return {error: "Not implemented"};
    }
}
