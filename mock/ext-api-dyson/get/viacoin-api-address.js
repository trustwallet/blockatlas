/// Mock for external Viacoin API
/// See:
/// curl "http://{Viacoin rpc}/v2/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A?details=txs"
/// curl "http://localhost:3347/viacoin-api/v2/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A?details=txs"
/// curl "http://localhost:8437/v1/viacoin/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A"

module.exports = {
    path: '/viacoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 785695,
                    "itemsOnPage": 2,
                    "address": "VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A",
                    "balance": "120312392935",
                    "totalReceived": "18370512600313",
                    "totalSent": "18250200207378",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 1571390,
                    "transactions": [
                        {
                            "txid": "1b311427cd4749cbebd5c1fc70896b9e92a58d6935ba75bb40ed4a9e22e4a6cb",
                            "version": 1,
                            "vin": [
                                {
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "coinbase": "03a02873045ea2915d"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "976562",
                                    "n": 0,
                                    "hex": "76a91424cc424c1e5e977175d2b20012554d39024bd68f88ac",
                                    "addresses": [
                                        "VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A"
                                    ]
                                }
                            ],
                            "blockHash": "c590a4a7357d308917c6d3d8a8bd75b27968db346b965f2cd6e2be639e792f24",
                            "blockHeight": 7547040,
                            "confirmations": 3,
                            "blockTime": 1587712349,
                            "value": "976562",
                            "valueIn": "0",
                            "fees": "0",
                            "hex": "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0903a02873045ea2915dffffffff01b2e60e00000000001976a91424cc424c1e5e977175d2b20012554d39024bd68f88ac00000000"
                        },
                        {
                            "txid": "e7393042e4b7755ba320a80ee734b0f0bd496a46fa3c9e9982c1a8176b6815b4",
                            "version": 1,
                            "vin": [
                                {
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "coinbase": "039e2873045ea29101"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "976562",
                                    "n": 0,
                                    "hex": "76a91424cc424c1e5e977175d2b20012554d39024bd68f88ac",
                                    "addresses": [
                                        "VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A"
                                    ]
                                }
                            ],
                            "blockHash": "87f245b6882d8d1a0a2122c5dd5e00a8a6ea0e811d8bc53cde5ba00a1a7b63ce",
                            "blockHeight": 7547038,
                            "confirmations": 5,
                            "blockTime": 1587712257,
                            "value": "976562",
                            "valueIn": "0",
                            "fees": "0",
                            "hex": "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff09039e2873045ea29101ffffffff01b2e60e00000000001976a91424cc424c1e5e977175d2b20012554d39024bd68f88ac00000000"
                        }
                    ]
                }                
                `);
        }
        return {error: "Not implemented"};
    }
}
