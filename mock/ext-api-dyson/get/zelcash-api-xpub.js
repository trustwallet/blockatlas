/// Mock for external Zelcash API
/// See:
/// curl "http://{Zelcash rpc}/v2/xpub/xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf?details=txs"
/// curl "http://localhost:3000/zelcash-api/v2/xpub/xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf?details=txs"
/// curl "http://localhost:8420/v1/zelcash/xpub/xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf"

module.exports = {
    path: '/zelcash-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "address": "xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf",
                        "balance": "9209317640",
                        "totalReceived": "135114824440",
                        "totalSent": "125905506800",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 29,
                        "transactions": [
                            {
                                "txid": "640ea080f060fbfcc6722f861c93483551a831020f6731a37ba5383da71b4972",
                                "version": 4,
                                "vin": [
                                    {
                                        "txid": "a3db9e10aa146805477b9012130e6f68d473975095b82c7e6e6ab5f409da9eef",
                                        "n": 0,
                                        "addresses": [
                                            "t1PHmJzMdeErPQGL9WEavrVakYiCzxKXi9z"
                                        ],
                                        "value": "1000000",
                                        "hex": "47304402205f7e4f5c17f6bf9880bd0afecec581253419a7b94a440fc9d4a4f4c35ee5b69102203a5c1ba8fc40098e9de6d4df71bb364e9028ddf9eb6ec31e42a1293c19170d8d012103aa6d9069ced0b9b556ac6eb7468d5f15422b16903507e8651824c1b7e6e36644"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "100000",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a9143b6a6100245510218dc29b313b988df0628ab36388ac",
                                        "addresses": [
                                            "t1PHmJzMdeErPQGL9WEavrVakYiCzxKXi9z"
                                        ]
                                    },
                                    {
                                        "value": "890000",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "76a9143b6a6100245510218dc29b313b988df0628ab36388ac",
                                        "addresses": [
                                            "t1PHmJzMdeErPQGL9WEavrVakYiCzxKXi9z"
                                        ]
                                    }
                                ],
                                "blockHash": "0000004957cd52c4654e44b0a02cd33be907ae3ae47949e44984cccaf0e34e96",
                                "blockHeight": 560762,
                                "confirmations": 1627,
                                "blockTime": 1584928030,
                                "value": "990000",
                                "valueIn": "1000000",
                                "fees": "10000",
                                "hex": "0400008085202f8901ef9eda09f4b56a6e7e2cb895509773d4686f0e1312907b47056814aa109edba3000000006a47304402205f7e4f5c17f6bf9880bd0afecec581253419a7b94a440fc9d4a4f4c35ee5b69102203a5c1ba8fc40098e9de6d4df71bb364e9028ddf9eb6ec31e42a1293c19170d8d012103aa6d9069ced0b9b556ac6eb7468d5f15422b16903507e8651824c1b7e6e366440000000002a0860100000000001976a9143b6a6100245510218dc29b313b988df0628ab36388ac90940d00000000001976a9143b6a6100245510218dc29b313b988df0628ab36388ac00000000000000000000000000000000000000"
                            },
                            {
                                "txid": "a3db9e10aa146805477b9012130e6f68d473975095b82c7e6e6ab5f409da9eef",
                                "version": 4,
                                "vin": [
                                    {
                                        "txid": "2de19629bbf78af114ff9cf3332229dc336be4758481b72828f1bd786e8cc41e",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "t1PHmJzMdeErPQGL9WEavrVakYiCzxKXi9z"
                                        ],
                                        "value": "899960000",
                                        "hex": "4830450221008f14795cde861210cbd9a1c3d89a29f0efe5c3cdceb2049795d5a9a84cc2571f02203c4e360575b7e5a45b067382b7d4e3fadbc80382eaa32e08a0bba46ae9bc7cb9012103aa6d9069ced0b9b556ac6eb7468d5f15422b16903507e8651824c1b7e6e36644"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1000000",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a9143b6a6100245510218dc29b313b988df0628ab36388ac",
                                        "addresses": [
                                            "t1PHmJzMdeErPQGL9WEavrVakYiCzxKXi9z"
                                        ]
                                    },
                                    {
                                        "value": "898950000",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "76a9143b6a6100245510218dc29b313b988df0628ab36388ac",
                                        "addresses": [
                                            "t1PHmJzMdeErPQGL9WEavrVakYiCzxKXi9z"
                                        ]
                                    }
                                ],
                                "blockHash": "0000008979ea8aac4d45ae6b5db721845d2d909e723b8fbb87d81f528a7ceefe",
                                "blockHeight": 558864,
                                "confirmations": 3525,
                                "blockTime": 1584697832,
                                "value": "899950000",
                                "valueIn": "899960000",
                                "fees": "10000",
                                "hex": "0400008085202f89011ec48c6e78bdf12828b7818475e46b33dc292233f39cff14f18af7bb2996e12d010000006b4830450221008f14795cde861210cbd9a1c3d89a29f0efe5c3cdceb2049795d5a9a84cc2571f02203c4e360575b7e5a45b067382b7d4e3fadbc80382eaa32e08a0bba46ae9bc7cb9012103aa6d9069ced0b9b556ac6eb7468d5f15422b16903507e8651824c1b7e6e36644000000000240420f00000000001976a9143b6a6100245510218dc29b313b988df0628ab36388ac70e39435000000001976a9143b6a6100245510218dc29b313b988df0628ab36388ac00000000000000000000000000000000000000"
                            },
                            {
                                "txid": "2de19629bbf78af114ff9cf3332229dc336be4758481b72828f1bd786e8cc41e",
                                "version": 4,
                                "vin": [
                                    {
                                        "txid": "7db9135b1dabcecccff9e4d403802d1bf0d7850438233b3091b876331bab551b",
                                        "n": 0,
                                        "addresses": [
                                            "t1MpHLd2vryyTufEMF3WBHy7VP63YKigi53"
                                        ],
                                        "value": "999970000",
                                        "hex": "473044022040534cefd8cfbe8f2705032a3636067fb8612ec3449f75699fcd894e5f2dc9960220363c26ddbbdb840fc76fbb950aecbfd6fa7d03e64da1c5890c302f0bb6f007a8012103d6bf5faa2c95ae8c4572f0755f45dba20dcd806ef8344a5dc2a54996874661c4"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "100000000",
                                        "n": 0,
                                        "hex": "76a914a394f62bbe0a47af97171e99bd8ff4fed9cae33188ac",
                                        "addresses": [
                                            "t1YnYbdwU1ReMUbN46byAQNtnDPqMRYxFPJ"
                                        ]
                                    },
                                    {
                                        "value": "899960000",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "76a9143b6a6100245510218dc29b313b988df0628ab36388ac",
                                        "addresses": [
                                            "t1PHmJzMdeErPQGL9WEavrVakYiCzxKXi9z"
                                        ]
                                    }
                                ],
                                "blockHash": "0000007f0db0ed2ac5f5747e7746dc5e016e1f2a84df54115e85eee4d7aecb3b",
                                "blockHeight": 547987,
                                "confirmations": 14402,
                                "blockTime": 1583385317,
                                "value": "999960000",
                                "valueIn": "999970000",
                                "fees": "10000",
                                "hex": "0400008085202f89011b55ab1b3376b891303b23380485d7f01b2d8003d4e4f9cfccceab1d5b13b97d000000006a473044022040534cefd8cfbe8f2705032a3636067fb8612ec3449f75699fcd894e5f2dc9960220363c26ddbbdb840fc76fbb950aecbfd6fa7d03e64da1c5890c302f0bb6f007a8012103d6bf5faa2c95ae8c4572f0755f45dba20dcd806ef8344a5dc2a54996874661c4000000000200e1f505000000001976a914a394f62bbe0a47af97171e99bd8ff4fed9cae33188acc04ca435000000001976a9143b6a6100245510218dc29b313b988df0628ab36388ac00000000000000000000000000000000000000"
                            },
                            {
                                "txid": "7db9135b1dabcecccff9e4d403802d1bf0d7850438233b3091b876331bab551b",
                                "version": 4,
                                "vin": [
                                    {
                                        "txid": "3ce3913910d9aab7ddb308b16abd6d5d66c258737b5f86e829cc741fa317820b",
                                        "n": 0,
                                        "addresses": [
                                            "t1YnYbdwU1ReMUbN46byAQNtnDPqMRYxFPJ"
                                        ],
                                        "value": "999980000",
                                        "hex": "483045022100fdf9542e8c44a218601d427ac9d8262866083c3ac1312c1448c345ed54782901022077a4c5e1d22657e000a2c76dc6f873fcb70b7298bd627dd3dc7bb24fa6175ab10121033985b538c7c62cd419a5b24e0dd8d322174303f376e025fddd1f19da2cebcc50"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "999970000",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a9142b3fac242eb3d7e20130082190605a5f7750037788ac",
                                        "addresses": [
                                            "t1MpHLd2vryyTufEMF3WBHy7VP63YKigi53"
                                        ]
                                    }
                                ],
                                "blockHash": "00000021b6af345a3d3010cdf9811fa9ca1f1163841519f2906ababa766b21ec",
                                "blockHeight": 494906,
                                "confirmations": 67483,
                                "blockTime": 1576977235,
                                "value": "999970000",
                                "valueIn": "999980000",
                                "fees": "10000",
                                "hex": "0400008085202f89010b8217a31f74cc29e8865f7b7358c2665d6dbd6ab108b3ddb7aad9103991e33c000000006b483045022100fdf9542e8c44a218601d427ac9d8262866083c3ac1312c1448c345ed54782901022077a4c5e1d22657e000a2c76dc6f873fcb70b7298bd627dd3dc7bb24fa6175ab10121033985b538c7c62cd419a5b24e0dd8d322174303f376e025fddd1f19da2cebcc500000000001d0549a3b000000001976a9142b3fac242eb3d7e20130082190605a5f7750037788ac00000000000000000000000000000000000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
