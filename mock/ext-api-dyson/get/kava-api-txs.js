/// Kava API Mock
/// See:
/// curl "http://localhost:3000/kava-api/txs?transfer.recipient=kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m&page=1&limit=25"
/// curl "http://localhost:3000/kava-api/txs?message.sender=kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m&page=1&limit=25"
/// curl "http://{kava rpc}/txs?transfer.recipient=kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m&page=1&limit=25"
/// curl "http://{kava rpc}/txs?message.sender=kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m&page=1&limit=25"
/// curl http://localhost:8420/v1/kava/kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m
module.exports = {
    path: "/kava-api/txs?",
    template: function(params, query, body) {
        //console.log(query)
        if (query["transfer.recipient"] === 'kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m') {
            return JSON.parse(`
                {
                    "total_count": "1",
                    "count": "1",
                    "page_number": "1",
                    "page_total": "1",
                    "limit": "25",
                    "txs": [
                        {
                            "height": "793039",
                            "txhash": "84FFA6CC8428B7A5E82003A7BCA6D1FAB2DF296AD18BE73EE61CBCBF67623880",
                            "logs": [
                                {
                                    "msg_index": 0,
                                    "success": true,
                                    "log": "",
                                    "events": [
                                        {
                                            "type": "message",
                                            "attributes": [
                                                {
                                                    "key": "action",
                                                    "value": "send"
                                                },
                                                {
                                                    "key": "sender",
                                                    "value": "kava1g2rqynrdw22ca7y8a4cj3sncylspsg7r04fcr7"
                                                },
                                                {
                                                    "key": "module",
                                                    "value": "bank"
                                                }
                                            ]
                                        },
                                        {
                                            "type": "transfer",
                                            "attributes": [
                                                {
                                                    "key": "recipient",
                                                    "value": "kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "9999000ukava"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ],
                            "gas_wanted": "200000",
                            "gas_used": "70019",
                            "tx": {
                                "type": "cosmos-sdk/StdTx",
                                "value": {
                                    "msg": [
                                        {
                                            "type": "cosmos-sdk/MsgSend",
                                            "value": {
                                                "from_address": "kava1g2rqynrdw22ca7y8a4cj3sncylspsg7r04fcr7",
                                                "to_address": "kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m",
                                                "amount": [
                                                    {
                                                        "denom": "ukava",
                                                        "amount": "9999000"
                                                    }
                                                ]
                                            }
                                        }
                                    ],
                                    "fee": {
                                        "amount": [
                                            {
                                                "denom": "ukava",
                                                "amount": "1000"
                                            }
                                        ],
                                        "gas": "200000"
                                    },
                                    "signatures": [
                                        {
                                            "pub_key": {
                                                "type": "tendermint/PubKeySecp256k1",
                                                "value": "As3eDIXFkLdMkBep+aV+6q9ccVRkA94snsXq2iL5bCTt"
                                            },
                                            "signature": "QFx3iARLhLeuCCHRUbFAZTePzq+vGAYd4lB/F4JuCGYw1TrfIbAzJhmUHbXHcB/GC+f7Myv5xNrnlCPfCR53uA=="
                                        }
                                    ],
                                    "memo": ""
                                }
                            },
                            "timestamp": "2020-01-17T11:17:36Z",
                            "events": [
                                {
                                    "type": "message",
                                    "attributes": [
                                        {
                                            "key": "action",
                                            "value": "send"
                                        },
                                        {
                                            "key": "sender",
                                            "value": "kava1g2rqynrdw22ca7y8a4cj3sncylspsg7r04fcr7"
                                        },
                                        {
                                            "key": "module",
                                            "value": "bank"
                                        }
                                    ]
                                },
                                {
                                    "type": "transfer",
                                    "attributes": [
                                        {
                                            "key": "recipient",
                                            "value": "kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m"
                                        },
                                        {
                                            "key": "amount",
                                            "value": "9999000ukava"
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            `);
        }

        if (query["message.sender"] === 'kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m') {
            return JSON.parse(`
                {
                    "total_count": "1",
                    "count": "1",
                    "page_number": "1",
                    "page_total": "1",
                    "limit": "25",
                    "txs": [
                        {
                            "height": "1101012",
                            "txhash": "30EFE9E830D317F84629F9DC35177577DF6713D78D354C2C469DE633900303BC",
                            "logs": [
                                {
                                    "msg_index": 0,
                                    "success": true,
                                    "log": "",
                                    "events": [
                                        {
                                            "type": "message",
                                            "attributes": [
                                                {
                                                    "key": "action",
                                                    "value": "send"
                                                },
                                                {
                                                    "key": "sender",
                                                    "value": "kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m"
                                                },
                                                {
                                                    "key": "module",
                                                    "value": "bank"
                                                }
                                            ]
                                        },
                                        {
                                            "type": "transfer",
                                            "attributes": [
                                                {
                                                    "key": "recipient",
                                                    "value": "kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "9998000ukava"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ],
                            "gas_wanted": "200000",
                            "gas_used": "68317",
                            "tx": {
                                "type": "cosmos-sdk/StdTx",
                                "value": {
                                    "msg": [
                                        {
                                            "type": "cosmos-sdk/MsgSend",
                                            "value": {
                                                "from_address": "kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m",
                                                "to_address": "kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7",
                                                "amount": [
                                                    {
                                                        "denom": "ukava",
                                                        "amount": "9998000"
                                                    }
                                                ]
                                            }
                                        }
                                    ],
                                    "fee": {
                                        "amount": [
                                            {
                                                "denom": "ukava",
                                                "amount": "1000"
                                            }
                                        ],
                                        "gas": "200000"
                                    },
                                    "signatures": [
                                        {
                                            "pub_key": {
                                                "type": "tendermint/PubKeySecp256k1",
                                                "value": "AvgGasyPvRpG68UQ6IGPVKN+HrJix24tgwcYXIcIfsna"
                                            },
                                            "signature": "iST3iM8TtkR5VHTFFhJE1zi+Nh6Invrg4hKgLEY2HCQ+R6Qj9L3jET4yRtAg6+QxitJab27+etoPjJbK20V/Cw=="
                                        }
                                    ],
                                    "memo": "100009547"
                                }
                            },
                            "timestamp": "2020-02-11T01:33:46Z",
                            "events": [
                                {
                                    "type": "message",
                                    "attributes": [
                                        {
                                            "key": "action",
                                            "value": "send"
                                        },
                                        {
                                            "key": "sender",
                                            "value": "kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m"
                                        },
                                        {
                                            "key": "module",
                                            "value": "bank"
                                        }
                                    ]
                                },
                                {
                                    "type": "transfer",
                                    "attributes": [
                                        {
                                            "key": "recipient",
                                            "value": "kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7"
                                        },
                                        {
                                            "key": "amount",
                                            "value": "9998000ukava"
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            `);
        }

        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};

