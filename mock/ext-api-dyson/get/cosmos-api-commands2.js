/// Cosmos API Mock
/// See:
/// curl "http://localhost:3000/cosmos-api/staking/validators?status=bonded"
/// curl "http://localhost:3000/cosmos-api/staking/pool"
/// curl "http://localhost:3000/cosmos-api/minting/inflation"
/// curl "https://{cosmos_rpc}/staking/validators?status=bonded"
/// curl "https://{cosmos_rpc}/staking/pool"
/// curl "https://{cosmos_rpc}/minting/inflation"
/// curl "http://localhost:8420/v2/cosmos/staking/validators"
/// curl "http://localhost:8420/v2/cosmos/staking/delegations/cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq?Authorization=Bearer"

module.exports = {
    path: "/cosmos-api/:command1/:command2?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch(params.command1) {
            case 'minting':
                switch(params.command2) {
                    case 'inflation':
                        // status=bonded
                        return JSON.parse(`{"height":"1419036","result":"0.070000000000000000"}`);
                }

            case 'txs': {
                if (query["transfer.recipient"] === 'cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq') {
                    return JSON.parse(`
                        {
                            "total_count": "1",
                            "count": "1",
                            "page_number": "1",
                            "page_total": "1",
                            "limit": "25",
                            "txs": [
                                {
                                    "height": "26616",
                                    "txhash": "3FEE8F10FBA5505DE2A8D3EF220D5CE900CA76B96208234BBF89B8075743A230",
                                    "data": "0C0886C1BEF00510E7FFF6F801",
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
                                                            "key": "sender",
                                                            "value": "cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl"
                                                        },
                                                        {
                                                            "key": "sender",
                                                            "value": "cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh"
                                                        },
                                                        {
                                                            "key": "module",
                                                            "value": "staking"
                                                        },
                                                        {
                                                            "key": "sender",
                                                            "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                        },
                                                        {
                                                            "key": "action",
                                                            "value": "begin_unbonding"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "type": "transfer",
                                                    "attributes": [
                                                        {
                                                            "key": "recipient",
                                                            "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                        },
                                                        {
                                                            "key": "amount",
                                                            "value": "1147uatom"
                                                        },
                                                        {
                                                            "key": "recipient",
                                                            "value": "cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r"
                                                        },
                                                        {
                                                            "key": "amount",
                                                            "value": "2203000uatom"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "type": "unbond",
                                                    "attributes": [
                                                        {
                                                            "key": "validator",
                                                            "value": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5"
                                                        },
                                                        {
                                                            "key": "amount",
                                                            "value": "2203000"
                                                        },
                                                        {
                                                            "key": "completion_time",
                                                            "value": "2020-01-03T20:13:58Z"
                                                        }
                                                    ]
                                                }
                                            ]
                                        }
                                    ],
                                    "gas_wanted": "200000",
                                    "gas_used": "127111",
                                    "tx": {
                                        "type": "cosmos-sdk/StdTx",
                                        "value": {
                                            "msg": [
                                                {
                                                    "type": "cosmos-sdk/MsgUndelegate",
                                                    "value": {
                                                        "delegator_address": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq",
                                                        "validator_address": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5",
                                                        "amount": {
                                                            "denom": "uatom",
                                                            "amount": "2203000"
                                                        }
                                                    }
                                                }
                                            ],
                                            "fee": {
                                                "amount": [
                                                    {
                                                        "denom": "uatom",
                                                        "amount": "1000"
                                                    }
                                                ],
                                                "gas": "200000"
                                            },
                                            "signatures": [
                                                {
                                                    "pub_key": {
                                                        "type": "tendermint/PubKeySecp256k1",
                                                        "value": "A782zo6TI2H3DfHJ7X1WHOJz6p4fUYVRYhb/XqMTcVQt"
                                                    },
                                                    "signature": "fJkOgQX9FZz6UJJo48AA3qxc2v/vFrjxlMgq3iOhDmVUWjA2P6X+pPbPkSB1SHlBFlVbsNF/DDPB0pvP9LRPTg=="
                                                }
                                            ],
                                            "memo": ""
                                        }
                                    },
                                    "timestamp": "2019-12-13T20:13:58Z",
                                    "events": [
                                        {
                                            "type": "message",
                                            "attributes": [
                                                {
                                                    "key": "sender",
                                                    "value": "cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl"
                                                },
                                                {
                                                    "key": "sender",
                                                    "value": "cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh"
                                                },
                                                {
                                                    "key": "module",
                                                    "value": "staking"
                                                },
                                                {
                                                    "key": "sender",
                                                    "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                },
                                                {
                                                    "key": "action",
                                                    "value": "begin_unbonding"
                                                }
                                            ]
                                        },
                                        {
                                            "type": "transfer",
                                            "attributes": [
                                                {
                                                    "key": "recipient",
                                                    "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "1147uatom"
                                                },
                                                {
                                                    "key": "recipient",
                                                    "value": "cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "2203000uatom"
                                                }
                                            ]
                                        },
                                        {
                                            "type": "unbond",
                                            "attributes": [
                                                {
                                                    "key": "validator",
                                                    "value": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "2203000"
                                                },
                                                {
                                                    "key": "completion_time",
                                                    "value": "2020-01-03T20:13:58Z"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    `);
                }
        
                if (query["message.sender"] === 'cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq') {
                    return JSON.parse(`
                        {
                            "total_count": "2",
                            "count": "2",
                            "page_number": "1",
                            "page_total": "1",
                            "limit": "25",
                            "txs": [
                                {
                                    "height": "26616",
                                    "txhash": "3FEE8F10FBA5505DE2A8D3EF220D5CE900CA76B96208234BBF89B8075743A230",
                                    "data": "0C0886C1BEF00510E7FFF6F801",
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
                                                            "key": "sender",
                                                            "value": "cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl"
                                                        },
                                                        {
                                                            "key": "sender",
                                                            "value": "cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh"
                                                        },
                                                        {
                                                            "key": "module",
                                                            "value": "staking"
                                                        },
                                                        {
                                                            "key": "sender",
                                                            "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                        },
                                                        {
                                                            "key": "action",
                                                            "value": "begin_unbonding"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "type": "transfer",
                                                    "attributes": [
                                                        {
                                                            "key": "recipient",
                                                            "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                        },
                                                        {
                                                            "key": "amount",
                                                            "value": "1147uatom"
                                                        },
                                                        {
                                                            "key": "recipient",
                                                            "value": "cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r"
                                                        },
                                                        {
                                                            "key": "amount",
                                                            "value": "2203000uatom"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "type": "unbond",
                                                    "attributes": [
                                                        {
                                                            "key": "validator",
                                                            "value": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5"
                                                        },
                                                        {
                                                            "key": "amount",
                                                            "value": "2203000"
                                                        },
                                                        {
                                                            "key": "completion_time",
                                                            "value": "2020-01-03T20:13:58Z"
                                                        }
                                                    ]
                                                }
                                            ]
                                        }
                                    ],
                                    "gas_wanted": "200000",
                                    "gas_used": "127111",
                                    "tx": {
                                        "type": "cosmos-sdk/StdTx",
                                        "value": {
                                            "msg": [
                                                {
                                                    "type": "cosmos-sdk/MsgUndelegate",
                                                    "value": {
                                                        "delegator_address": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq",
                                                        "validator_address": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5",
                                                        "amount": {
                                                            "denom": "uatom",
                                                            "amount": "2203000"
                                                        }
                                                    }
                                                }
                                            ],
                                            "fee": {
                                                "amount": [
                                                    {
                                                        "denom": "uatom",
                                                        "amount": "1000"
                                                    }
                                                ],
                                                "gas": "200000"
                                            },
                                            "signatures": [
                                                {
                                                    "pub_key": {
                                                        "type": "tendermint/PubKeySecp256k1",
                                                        "value": "A782zo6TI2H3DfHJ7X1WHOJz6p4fUYVRYhb/XqMTcVQt"
                                                    },
                                                    "signature": "fJkOgQX9FZz6UJJo48AA3qxc2v/vFrjxlMgq3iOhDmVUWjA2P6X+pPbPkSB1SHlBFlVbsNF/DDPB0pvP9LRPTg=="
                                                }
                                            ],
                                            "memo": ""
                                        }
                                    },
                                    "timestamp": "2019-12-13T20:13:58Z",
                                    "events": [
                                        {
                                            "type": "message",
                                            "attributes": [
                                                {
                                                    "key": "sender",
                                                    "value": "cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl"
                                                },
                                                {
                                                    "key": "sender",
                                                    "value": "cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh"
                                                },
                                                {
                                                    "key": "module",
                                                    "value": "staking"
                                                },
                                                {
                                                    "key": "sender",
                                                    "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                },
                                                {
                                                    "key": "action",
                                                    "value": "begin_unbonding"
                                                }
                                            ]
                                        },
                                        {
                                            "type": "transfer",
                                            "attributes": [
                                                {
                                                    "key": "recipient",
                                                    "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "1147uatom"
                                                },
                                                {
                                                    "key": "recipient",
                                                    "value": "cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "2203000uatom"
                                                }
                                            ]
                                        },
                                        {
                                            "type": "unbond",
                                            "attributes": [
                                                {
                                                    "key": "validator",
                                                    "value": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "2203000"
                                                },
                                                {
                                                    "key": "completion_time",
                                                    "value": "2020-01-03T20:13:58Z"
                                                }
                                            ]
                                        }
                                    ]
                                },
                                {
                                    "height": "404179",
                                    "txhash": "93E43518BAE4BC137605BBB7FD5D31FDAE6427ECE57EC299C43CE786FDAEBC63",
                                    "logs": [
                                        {
                                            "msg_index": 0,
                                            "success": true,
                                            "log": "",
                                            "events": [
                                                {
                                                    "type": "delegate",
                                                    "attributes": [
                                                        {
                                                            "key": "validator",
                                                            "value": "cosmosvaloper17h2x3j7u44qkrq0sk8ul0r2qr440rwgjkfg0gh"
                                                        },
                                                        {
                                                            "key": "amount",
                                                            "value": "2211271"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "type": "message",
                                                    "attributes": [
                                                        {
                                                            "key": "module",
                                                            "value": "staking"
                                                        },
                                                        {
                                                            "key": "sender",
                                                            "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                        },
                                                        {
                                                            "key": "action",
                                                            "value": "delegate"
                                                        }
                                                    ]
                                                }
                                            ]
                                        }
                                    ],
                                    "gas_wanted": "200000",
                                    "gas_used": "93720",
                                    "tx": {
                                        "type": "cosmos-sdk/StdTx",
                                        "value": {
                                            "msg": [
                                                {
                                                    "type": "cosmos-sdk/MsgDelegate",
                                                    "value": {
                                                        "delegator_address": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq",
                                                        "validator_address": "cosmosvaloper17h2x3j7u44qkrq0sk8ul0r2qr440rwgjkfg0gh",
                                                        "amount": {
                                                            "denom": "uatom",
                                                            "amount": "2211271"
                                                        }
                                                    }
                                                }
                                            ],
                                            "fee": {
                                                "amount": [
                                                    {
                                                        "denom": "uatom",
                                                        "amount": "1000"
                                                    }
                                                ],
                                                "gas": "200000"
                                            },
                                            "signatures": [
                                                {
                                                    "pub_key": {
                                                        "type": "tendermint/PubKeySecp256k1",
                                                        "value": "A782zo6TI2H3DfHJ7X1WHOJz6p4fUYVRYhb/XqMTcVQt"
                                                    },
                                                    "signature": "X1/NzRdb+HUxE7N9gMk39XI8TyHFobLRQtsX4QxZZbYeKmEYOOZ7FyNSGqgmipCkpuysBeh6fNnbXmo3IFzFEQ=="
                                                }
                                            ],
                                            "memo": ""
                                        }
                                    },
                                    "timestamp": "2020-01-13T15:23:12Z",
                                    "events": [
                                        {
                                            "type": "delegate",
                                            "attributes": [
                                                {
                                                    "key": "validator",
                                                    "value": "cosmosvaloper17h2x3j7u44qkrq0sk8ul0r2qr440rwgjkfg0gh"
                                                },
                                                {
                                                    "key": "amount",
                                                    "value": "2211271"
                                                }
                                            ]
                                        },
                                        {
                                            "type": "message",
                                            "attributes": [
                                                {
                                                    "key": "module",
                                                    "value": "staking"
                                                },
                                                {
                                                    "key": "sender",
                                                    "value": "cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq"
                                                },
                                                {
                                                    "key": "action",
                                                    "value": "delegate"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    `);
                }        
            }
            
            case 'staking':
                switch(params.command2) {
                    case 'pool':
                        return JSON.parse(`
                            {
                                "height":"1419019",
                                "result": {
                                    "not_bonded_tokens": "2422309183727",
                                    "bonded_tokens": "182732778407119"
                                }
                            }
                        `);

                    case 'validators':
                        return JSON.parse(`
                            {
                                "height": "1418977",
                                "result": [
                                    {
                                        "operator_address": "cosmosvaloper1qdxmyqkvt8jsxpn5pp45a38ngs36mn2604cqk9",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqmq2l5ehgl3rxwjgzgr6sgzp69qwjl5ufvtyvacc9ms8p3phazl2qh3ulfw",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "94117413314",
                                        "delegator_shares": "94117413314.000000000000000000",
                                        "description": {
                                            "moniker": "真本聪\u0026IOSG",
                                            "identity": "8A79F44CC25D26DF",
                                            "website": "realsatoshi.net",
                                            "details": "To The Moon Then Cosmos. We are a crypto community and venture capital combined staking service provider"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-06-24T19:55:42.604341211Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-08-05T07:24:32.572614545Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqwrjpn0slu86e32zfu5xxg8l42uk40guuw6er44vw2yl6s7wc38est6l0ux",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "5289222663449",
                                        "delegator_shares": "5289222663449.000000000000000000",
                                        "description": {
                                            "moniker": "Certus One",
                                            "identity": "ABD51DF68C0D1ECF",
                                            "website": "https://certus.one",
                                            "details": "Stake and earn rewards with the most secure and stable validator. Winner of the Game of Stakes. Operated by Certus One Inc. By delegating, you confirm that you are aware of the risk of slashing and that Certus One Inc is not liable for any potential damages to your investment."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.125000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1qs8tnw2t8l6amtzvdemnnsq9dzk0ag0z52uzay",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqsszd2gzte82dzt0xpa3w0ky8lxhjs6zpd5ft8akmkscwujpftymsnt83qc",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1605450345336",
                                        "delegator_shares": "1605450345336.000000000000000000",
                                        "description": {
                                            "moniker": "Castlenode",
                                            "identity": "F685CC35D748424C",
                                            "website": "https://www.castlenode.com/cosmos",
                                            "details": "Castlenode is a validator operator focused on security and run by experienced professionals. Please read our Terms and Conditions on our website before delegating"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.090000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1pz0lfq40sa63n0wany3v95x3yvznc5gyf8u28w",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqv03r9c26w884sqw0zdqg9cdx4xxwgn3dd2s6wa3x346tm83nxudqwhm3jl",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "296744671239",
                                        "delegator_shares": "296744671239.000000000000000000",
                                        "description": {
                                            "moniker": "Cobo",
                                            "identity": "3B7C85200D5B57A9",
                                            "website": "https://cobo.com/",
                                            "details": "Cobo is the first leading company in the world to offer Proof-of-Stake (PoS) and masternode rewards on user holdings, making it easy for users to grow their digital assets effortlessly."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-05-18T11:34:21.668987388Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-06-15T05:45:23.950986198Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1pgsjyvkg3y2m7qas534zzdhsqsxqyph2jh3uck",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq5tycuhznpjuwlr93ypflreaddk0r66hdryzgmfvmlj2ekdd5sqrqnxclgg",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "228520445065",
                                        "delegator_shares": "228520445065.000000000000000000",
                                        "description": {
                                            "moniker": "OneSixtyTwo",
                                            "identity": "1C136F82A18BB2E2",
                                            "website": "",
                                            "details": "OneSixtyTwo"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-03-20T16:43:30.502949835Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqtcsm8lp7n6ph98vd59qa9esgyuysuntww9juz5wynxrhzpspmuuq6g5pzg",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1125811000585",
                                        "delegator_shares": "1125811000585.000000000000000000",
                                        "description": {
                                            "moniker": "WeStaking",
                                            "identity": "DA9C5AD3E308E426",
                                            "website": "https://www.westaking.io",
                                            "details": "Delegate your atom to us for the staking rewards. We will do our best as secure and stable validator."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-09-07T12:24:58.270714195Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.030000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-03-29T00:37:22.163935678Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1pjmngrwcsatsuyy8m3qrunaun67sr9x7z5r2qs",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqcdav5ylt2zst90qmuh8e5w07xmxv9y6wufp5k9ngzmx7v9qewqtqkcq4z8",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "654800077136",
                                        "delegator_shares": "654800077136.000000000000000000",
                                        "description": {
                                            "moniker": "Cypher Core",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1p650epkdwj0jte6sjc3ep0n3wz6jc9ehh8jutg",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqmy5ja8dee3sprtmxkekcuvwz20y2x9d6llcsjl6p6553rpqev0eql6qccf",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "111582433714",
                                        "delegator_shares": "111615914997.731210796675788592",
                                        "description": {
                                            "moniker": "Cosmos Suisse",
                                            "identity": "2D875495A911A943",
                                            "website": "https://cosmos-suisse.com",
                                            "details": "An awesome Validator based in Crypto Valley, Switzerland."
                                        },
                                        "unbonding_height": "970101",
                                        "unbonding_time": "2020-03-21T06:33:36.130547135Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.295000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-04-05T06:40:17.027131003Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1zqgheeawp7cmqk27dgyctd80rd8ryhqs6la9wc",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq5tm478lhn4du7l46yp6fgu9e8fcfks0j4kf87pk4z8vc3clckxzqvh2q72",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "685457015651",
                                        "delegator_shares": "685457015651.000000000000000000",
                                        "description": {
                                            "moniker": "melea-◮👁◭",
                                            "identity": "4BE49EABAA41B8BF",
                                            "website": "https://meleatrust.com/",
                                            "details": "Validator service secure and trusted, awarded in Game Of Steaks"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-06-26T23:50:09.38840326Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-11-27T05:18:46.405341672Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1rpgtz9pskr5geavkjz02caqmeep7cwwpv73axj",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq04y0dtylyed2f8drc9t78dmptfuta7l6xyujwmsgrqefs0sxpgjsnzpsj6",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "3995820273418",
                                        "delegator_shares": "3996619557261.806706040418750517",
                                        "description": {
                                            "moniker": "Blockpower",
                                            "identity": "C374865C7ADE710B",
                                            "website": "http://blockpower.capital/cosmos",
                                            "details": "We self-bond 1m atoms and will  keep our fee low  till transfer enabled. We are top-ten validator in Tezos and operate professional staking services for multiple PoS platforms. Twitter: https://twitter.com/blockpowercap and TG: https://t.me/blockpowercosmos"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-11-04T04:14:24.222468286Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.030000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-31T19:58:53.172697095Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1r8kyvg4me2upnvlk26n2ay0zd5t4jktna8hhxp",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqsl29jhapyxf4fa5a947dmnlvrt266fm959hfg2c657ju80hq0ljs3qejr7",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "250002000000",
                                        "delegator_shares": "250002000000.000000000000000000",
                                        "description": {
                                            "moniker": "noma",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-03-14T03:50:13.123302789Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1rwh0cxa72d3yle3r4l8gd7vyphrmjy2kpe4x72",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq5kg8xls9l35ftulkm2rt70hexeeyr5cqqkcv4h7936z5uasvvazqla8eck",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "5699344397878",
                                        "delegator_shares": "5699344397878.000000000000000000",
                                        "description": {
                                            "moniker": "SparkPool",
                                            "identity": "DE8E37240061B04E",
                                            "website": "https://cosmos.sparkpool.com/",
                                            "details": "The biggest Ethereum mining pool, we can be a reliable validator with our 3 years"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.040000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-09-26T06:05:19.676181518Z"
                                        },
                                        "min_self_delegation": "100000000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1rcp29q3hpd246n6qak7jluqep4v006cdsc2kkl",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq7mft6gfls57a0a42d7uhx656cckhfvtrlmw744jv4q0mvlv0dypskehfk8",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "243735228951",
                                        "delegator_shares": "243735228951.000000000000000000",
                                        "description": {
                                            "moniker": "2nd only to Certus One in GoS: in3s.com",
                                            "identity": "0CE19EE3E4BA48E5",
                                            "website": "https://in3s.com/Services#CosmosValidator",
                                            "details": "https://in3s.com/Delegate#Delegate"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "1.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-03-18T20:31:50.23335594Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1y0us8xvsvfvqkk9c6nt5cfyu5au5tww2ztve7q",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqqhkzzgfc876q287ny2v9lqwmjpenuzkzlsnstrmg7krwrrf0pfqs9fxvs0",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "12973804826",
                                        "delegator_shares": "12973804826.000000000000000000",
                                        "description": {
                                            "moniker": "Swiss Staking",
                                            "identity": "165F85FC0194320D",
                                            "website": "https://swiss-staking.ch",
                                            "details": "Experienced PoS Validator. We refund downtime slashing to 100%."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-12-11T18:25:42.711862818Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper19yy989ka5usws6gsd8vl94y7l6ssgdwsrnscjc",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqf22yaz9nsxlh043qm0tmupw8pnpver2n8lm3mwz6jzmsql76fkmqa482y8",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "500001550260",
                                        "delegator_shares": "500001550260.000000000000000000",
                                        "description": {
                                            "moniker": "OKEx Pool",
                                            "identity": "406C257E090E70AA",
                                            "website": "http://www.okpool.top",
                                            "details": "An innovative mining pool with higher yield"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-09-26T10:31:26.722694911Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1998928nfs697ep5d825y5jah0nq9zrtd00yyj7",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqlecm0rrfrr0vfgl624s7su9xvd3ycsaetndeuw2c7us0v8vfyfsq7cqz80",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "105880988262",
                                        "delegator_shares": "105902167539.392512073112670643",
                                        "description": {
                                            "moniker": "HLT",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-12-22T20:58:53.786852369Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-06-17T11:29:36.731435956Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper199mlc7fr6ll5t54w7tts7f4s0cvnqgc59nmuxf",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqajqpmv4j70a08ahs8lyjt8qk28ffa77zjegd7yajghchy8au575qmmxuyt",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "571392869594",
                                        "delegator_shares": "571392869594.000000000000000000",
                                        "description": {
                                            "moniker": "Velocity V1",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-05-30T20:41:10.689796224Z"
                                        },
                                        "min_self_delegation": "1000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper19v9ej55ataqrfl39v83pf4e0dm69u89rngf928",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqa32xgy8y5s69j6dynjmr3rlpv5dv35whz2tyaedwtjeqckm5gg4s2hj8ss",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "14356837431",
                                        "delegator_shares": "14356837431.000000000000000000",
                                        "description": {
                                            "moniker": "blockscape",
                                            "identity": "C46C8329BB5F48D8",
                                            "website": "https://blockscape.network/",
                                            "details": "By delegating, you confirm that you are aware of the risk of slashing and that M-Way Solutions GmbH is not liable for any potential damages to your investment."
                                        },
                                        "unbonding_height": "483",
                                        "unbonding_time": "2020-01-01T17:32:11.54021567Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.099900000000000000",
                                                "max_rate": "0.399900000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-12-02T09:18:02.304119479Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper19j2hd230c3hw6ds843yu8akc0xgvdvyuz9v02v",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq9ge7uqrfp9qkdapzd29tjtwrqpt2mm9meptx395ygxgm40tdc8ysrzj40a",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "441435636408",
                                        "delegator_shares": "441435636408.000000000000000000",
                                        "description": {
                                            "moniker": "syncnode",
                                            "identity": "F422F328C14AFBFA",
                                            "website": "wallet.syncnode.ro",
                                            "details": "email: g@ysncnode.ro || Operator's LinkedIn: https://www.linkedin.com/in/gbunea/ || Telegram Channel: https://t.me/syncnodeValidator || Blog: https://medium.com/syncnode-validator"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-01-15T11:41:35.747062104Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper19kwwdw0j64xhrmgkz49l0lmu5uyujjayxakwsn",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq668g4epaumjtx35rk3ucz2nlm7l7zuewkt0kzutg9hha859zjxmsvl2v67",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "630448698935",
                                        "delegator_shares": "630448698935.000000000000000000",
                                        "description": {
                                            "moniker": "Firmamint",
                                            "identity": "2FE4BC7A59E09FD0",
                                            "website": "https://www.firmamint.io/",
                                            "details": "The FUTURE is at STAKE - Proudly Canadian -  Tier 1 WINNER of Game of Stakes Adversarial Testnet"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.030000000000000000"
                                            },
                                            "update_time": "2019-03-18T22:02:43.333950761Z"
                                        },
                                        "min_self_delegation": "1500000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1xym2qygmr9vanpa0m7ndk3n0qxgey3ffzcyd5c",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq5evf9e4qhkxym6lx0ur2mwuz5e09u2v7u54yz3wcfanqwvhkc7rqcgpmlw",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "100726900000",
                                        "delegator_shares": "100726900000.000000000000000000",
                                        "description": {
                                            "moniker": "🐡grant.fish",
                                            "identity": "BE328F9A089F50C9",
                                            "website": "http://grant.fish",
                                            "details": "Providing grants to projects contributing to the Cosmos ecosystem."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "1.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-12-11T17:14:13.747872899Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1x8rr4hcf54nz6hfckyy2n05sxss54h8wz9puzg",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq8xpk5nh78lmgg0s0qqdyh4xtcw66xemt8anjsr6hrvlhauq252kstq7zr7",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "876463600211",
                                        "delegator_shares": "876463600211.000000000000000000",
                                        "description": {
                                            "moniker": "cosmosgbt",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-10-18T04:22:16.555340871Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1x88j7vp2xnw3zec8ur3g4waxycyz7m0mahdv3p",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqhm6gjjkwecqyfrgey96s5up7drnspnl4t3rdr79grklkg9ff6zaqnfl2dg",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "2025689252407",
                                        "delegator_shares": "2025689252407.000000000000000000",
                                        "description": {
                                            "moniker": "Staking Facilities",
                                            "identity": "6B0DF6793DE1FB1F",
                                            "website": "stakingfacilities.com",
                                            "details": "Earn rewards with one of the most experienced and secure validators. More than 150k USD in customer rewards paid out. We exclude liability for any financial damage resulting from delegating."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-23T22:11:43.172644866Z"
                                        },
                                        "min_self_delegation": "100000000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1x065cjlgejk2p2la0029akfvdy52gtq9mm58ta",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqny59kv2elgh89tq9qr4jje2n0my4gyvh2hlnydzljtt542a5plwswtas48",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "150298393037",
                                        "delegator_shares": "150298393037.000000000000000000",
                                        "description": {
                                            "moniker": "MathWallet麦子钱包",
                                            "identity": "58320327FF6C928C",
                                            "website": "https://mathwallet.org",
                                            "details": "Math Wallet is a powerful and secure universal crypto wallet that enables storage of all BTC, ETH/ERC20, NEO/NEP5, EOS/ENU/Telos, TRON, ONT, BinanceChain, Cosmos/IRISnet tokens, supports cross-chain token exchange and a multi-chain dApp store."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-07-21T13:52:13.251185336Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-09-09T03:44:27.230973438Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1grgelyng2v6v3t8z87wu3sxgt9m5s03xfytvz7",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqdgvppnyr5c9pulsrmzr9e9rp7qpgm9jwp5yu8g3aumekgjugxacq8a9p2c",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "5364854020103",
                                        "delegator_shares": "5364854020103.000000000000000000",
                                        "description": {
                                            "moniker": "iqlusion",
                                            "identity": "DCB176E79AE7D51F",
                                            "website": "iqlusion.io",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "100000000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1gdg6qqe5a3u483unqlqsnullja23g0xvqkxtk0",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqlsvqw4vacnv2qtwxmm8tq32lrcdhau3szqxevrrzy8u0jvwz69pqphnkzk",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "76705740036",
                                        "delegator_shares": "76705740036.000000000000000000",
                                        "description": {
                                            "moniker": "zugerselfdelegation",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "1.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2020-02-21T16:46:55.500729256Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1fqzqejwkk898fcslw4z4eeqjzesynvrdfr5hte",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqd4hvh0rwfkhtwrj4ly3ptyxs8pyfaser57wx2tcnzzp0rlref90sxm5kwr",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "412308417647",
                                        "delegator_shares": "412308417647.000000000000000000",
                                        "description": {
                                            "moniker": "commercio.network",
                                            "identity": "ADBDB0178E4441BE",
                                            "website": "https://commercio.network",
                                            "details": "The Documents Blockchain"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-06-23T11:25:09.272949514Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.090000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-05-25T10:53:14.835657659Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ff0dw8kawsnxkrgj7p65kvw7jxxakyf8n583gx",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqu08f7tuce8k88tgewhwer69kfvk5az3cn5lz3v8phl8gvu9nxu8qhrjxfj",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "685381971996",
                                        "delegator_shares": "685450516977.642009537883274736",
                                        "description": {
                                            "moniker": "Compass",
                                            "identity": "72CB5AAAAFB1CE69",
                                            "website": "http://val.network/",
                                            "details": "EasyZone is a decentralized light client, which means users can access account, stake and earn rewords with local key store.  For the time being we focus on Tendermint ecosystem, including Cosmos, QOS and Irisnet etc. Winner of the Game of Stakes."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-10-14T04:10:07.163539573Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1fhr7e04ct0zslmkzqt9smakg3sxrdve6ulclj2",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqzu34dgs2p6ysz52hpdycls4jcfwgnf2pvxv0eh539ypmadkjfmes6mwaa3",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "717001927145",
                                        "delegator_shares": "717001927145.000000000000000000",
                                        "description": {
                                            "moniker": "POS Bakerz",
                                            "identity": "3AFAE7268F4DFD10",
                                            "website": "https://posbakerz.com/",
                                            "details": "Secure, Reliable and Efficient Staking-as-a-Service"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-03-02T09:39:28.188566867Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepquuc5k7egx8ymejamr27c3sgw6tyhmt0eq0ak4qvflvhxx56nvjzsx9etmd",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "3596427902503",
                                        "delegator_shares": "3596427902503.000000000000000000",
                                        "description": {
                                            "moniker": "Huobi Wallet",
                                            "identity": "CF01514DBF6583FE",
                                            "website": "https://www.huobiwallet.com",
                                            "details": "Huobi Wallet is a leading multi-chain wallet from Huobi Group and also offering competitive Proof-of-Stake (PoS) rewards for users to grow their digital assets. As a Cosmos validator, Huobi Wallet offers the one and only hard slashing insurance fund to our delegators"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.020000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-01-06T02:06:36.387455308Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper125umsz3fws7gepn5ccsh0sv4gre9r6a3tccz4r",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqkzserh5aetg6m33vrfcsuz60qpkgkpkpzkh8gp9lsd7cxffgph4qjzkxft",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "6281466956",
                                        "delegator_shares": "6281466956.000000000000000000",
                                        "description": {
                                            "moniker": "Moonstake",
                                            "identity": "742F7B64C32DF7A6",
                                            "website": "https://www.moonstake.io/",
                                            "details": "Shoot For The Moon"
                                        },
                                        "unbonding_height": "950642",
                                        "unbonding_time": "2020-03-19T15:59:10.375998429Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.030000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2020-03-03T10:05:22.79594206Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper124maqmcqv8tquy764ktz7cu0gxnzfw54n3vww8",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq279zs0zgd53ujngs6v2hus2le9rk2a2rs66j4yvv6ewvvxn29yqqk95h8x",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "636354488555",
                                        "delegator_shares": "636354488555.000000000000000000",
                                        "description": {
                                            "moniker": "Simply Staking",
                                            "identity": "F74595D6D5D568A2",
                                            "website": "https://www.simply-vc.com.mt",
                                            "details": "Simply VC runs highly reliable and secure infrastructure in our own datacentre in Malta, built with the aim of supporting the growth of the blockchain ecosystem."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.070000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-05-11T17:29:59.537654017Z"
                                        },
                                        "min_self_delegation": "1000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1tflk30mq5vgqjdly92kkhhq3raev2hnz6eete3",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqwcvy8hyw2phdp080ggj7prxv972rvqc9gwyjnl0uwf7uxn63s8vqdctdcw",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "164930806828",
                                        "delegator_shares": "164930806828.000000000000000000",
                                        "description": {
                                            "moniker": "Everstake",
                                            "identity": "",
                                            "website": "https://everstake.one",
                                            "details": "Reliable and experienced staking service provider from Ukraine. Visit our website for more details."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-12-07T19:10:59.878559804Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.030000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-07-14T19:56:46.186760169Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ttfytaf43nkytzp8hkfjfgjc693ky4t3y2n2ku",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq2nfs6lcwu6ksq54yf0ptgrmjjrnm5p5ywng3x0t0767m777hvctq30rwcs",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "68002415747",
                                        "delegator_shares": "68009216666.737140135826986143",
                                        "description": {
                                            "moniker": "StarCluster",
                                            "identity": "F97B6EF4FD82202F",
                                            "website": "https://starcluster.tech",
                                            "details": "With decades of passion invested in tech, we provide a team of top security and infrastructure engineers as a service. With extensive knowledge in the blockchain space and having run a successful ICO, we are confident that providing our experience \u0026 skills could benefit many."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-12-24T15:27:58.293852371Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.150000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-10-26T21:50:11.42045872Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1te8nxpc2myjfrhaty0dnzdhs5ahdh5agzuym9v",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqjg26g27dtvjqstyqktmp4jsn98473vfz0mek2eyklfp0yqapav5szdrvpd",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "3519949961612",
                                        "delegator_shares": "3519949961612.000000000000000000",
                                        "description": {
                                            "moniker": "CoinoneNode",
                                            "identity": "F4E86EE9BD73A11F",
                                            "website": "https://node.coinone.co.kr",
                                            "details": "The more, the easier. Coinone Node manages your assets securely."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-14T09:22:28.276013718Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1vrg6ruw00lhszl4sjgwt5ldvl8z0f7pfp5va85",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq2hfnf0rvk6nksvtzkqly4vy362sencfkt7tgsrvx30krj5vxw0asa7hmjh",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "342783139733",
                                        "delegator_shares": "342783139733.000000000000000000",
                                        "description": {
                                            "moniker": "SSSnodes",
                                            "identity": "C5B68615F8828EC0",
                                            "website": "http://sssnodes.com/cosmos",
                                            "details": "Stake and earn rewards with the most secure and stable validator. Winner of the Game of Stakes. Operated by SSSnodes, a Corp. focused on delegation service for multiple Proof of Stake networks, such as COSMOS, ChainX, IOST, CyberMiles, Lambda, ONT, etc."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.018000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-11-28T08:09:59.47667909Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1vf44d85es37hwl9f4h9gv0e064m0lla60j9luj",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqtj2urav4g9wex3hku588au0x4sucrc9lpky46zp5u8w4mvd584sqmcxxhs",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "7246081767965",
                                        "delegator_shares": "7246081767965.000000000000000000",
                                        "description": {
                                            "moniker": "MultiChain Ventures",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.020000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-08-21T08:01:01.670548948Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1v5y0tg0jllvxf5c3afml8s3awue0ymju89frut",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq9kun5ty55rl3lnmf46tfxhj06as8h7zpxcdhujm6d708ffn6kgss43q6u9",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "5856792267830",
                                        "delegator_shares": "5856792267830.000000000000000000",
                                        "description": {
                                            "moniker": "Zero Knowledge Validator (ZKV)",
                                            "identity": "3E38E52A12F94561",
                                            "website": "https://zkvalidator.com/",
                                            "details": "Zero Knowledge Validator: Stake \u0026 Support ZKP Research \u0026 Privacy Tech"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-11-11T15:48:18.737920871Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1vk706z2tfnqhdg6jrkngyx7f463jq58nj0x7p7",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq0qdudwaluzpy5ptluu5umck5fp2ns2qf2kjp367qlr7yf65agx5s0n8h7f",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "26192977888",
                                        "delegator_shares": "26192977888.000000000000000000",
                                        "description": {
                                            "moniker": "Public Payments",
                                            "identity": "1850627141D54797",
                                            "website": "https://www.publicpayments.io",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-01-02T12:25:53.791107872Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1d0aup392g3enru7eash83sedqclaxvp7fzh6gk",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepql9j7qpwvfl0pspymhesj48t3t0aazjx0m2jwjuyxd7zw53hqnkss4hmasl",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "194657773848",
                                        "delegator_shares": "194677241406.157421240224999535",
                                        "description": {
                                            "moniker": "Stir",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-06-21T20:37:09.115641759Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-14T13:03:10.437070061Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1dse76yk5jmj85jsd77ewsczc4k3u4s7a870wtj",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqjdc9grwq5mxdtg26hv9t75y3ltsau3rtmg6p72p0dh8343nj4s6qr6xymd",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "399992036873",
                                        "delegator_shares": "400032040000.820061751426683457",
                                        "description": {
                                            "moniker": "gf.network",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "10001",
                                        "unbonding_time": "2020-01-02T11:58:18.459331455Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.080000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-05-27T12:23:08.242833465Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1de7qx00pz2j6gn9k88ntxxylelkazfk3g8fgh9",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqwr5p8j076mfydn7wckqz748lr0j50zwgsftnfpvgz6rz0rkvvqwqg5fyaf",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "222020931706",
                                        "delegator_shares": "222020931706.000000000000000000",
                                        "description": {
                                            "moniker": "Cosmic Validator",
                                            "identity": "FF4B91B50B71CEDA",
                                            "website": "",
                                            "details": "A reliable, passionate and service oriented Cosmos, Polkadot and Sentinel validator. We are long term trusted community members and have received delegation from the Interchain Foundation (ICF) as a reward for our continuous support and effort."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-10-07T05:35:40.25581859Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1d7ufwp2rgfj7s7pfw2q7vm2lc9txmr8vh77ztr",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq2x456pyef8jdnjgk8j62fuug24xfg9cnnjl66ewtgttwr00phz8sdzatkj",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "117924360855",
                                        "delegator_shares": "117924360855.000000000000000000",
                                        "description": {
                                            "moniker": "Cybernetic Destiny",
                                            "identity": "",
                                            "website": "",
                                            "details": "The future you’ve always wanted."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2020-02-07T23:19:03.834895518Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1wp9jne5t3e4au7u8gfep90g59j0qdhpeqvlg7n",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq6740f8r23xr74w94l5ew9fh6n8wquutgm22pw6yyrydq507mgdkqghsjtd",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "92787621513",
                                        "delegator_shares": "92787621513.000000000000000000",
                                        "description": {
                                            "moniker": "Newroad Network",
                                            "identity": "F898ACE263EC1C4E",
                                            "website": "https://newroad.network/cosmos/",
                                            "details": "We provide a professional delegation service for multiple Proof of Stake networks. We use a secure and redundant setup. Visit our website for more information."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2020-02-26T08:44:32.038638755Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1wtv0kp6ydt03edd8kyr5arr4f3yc52vp3u2x3u",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq0zyaquh6c8vmjfzft43uqylf2ejjpjcvup2zrtsk40uyz8xsq29s0k4eaw",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "353202000001",
                                        "delegator_shares": "353202000001.000000000000000000",
                                        "description": {
                                            "moniker": "kytzu",
                                            "identity": "909A480D5643CCC5",
                                            "website": "https://www.linkedin.com/in/calinchitu",
                                            "details": "Blockchain consultant, running on IPSX infrastructure (calin@ip.sx)"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-07-13T04:46:58.395203387Z"
                                        },
                                        "min_self_delegation": "1000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1wdrypwex63geqswmcy5qynv4w3z3dyef2qmyna",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqs0et7kpf82glsw5j9jnppekrpa7kl6gr6xk67ztqg9ynmhgj82ks9edcrw",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "457217580134",
                                        "delegator_shares": "457217580134.000000000000000000",
                                        "description": {
                                            "moniker": "Genesis Lab",
                                            "identity": "C1A123F2723041F0",
                                            "website": "https://genesislab.net",
                                            "details": "Genesis Lab is a blockchain-focused development company and validation nodes operator in PoS networks"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.070000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-15T17:32:00.387570437Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1wd8vquxza6svsvvnh489tdvg9vdjjfpepcfvf9",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqf4qptg2va4frg7g35c4plnlq0ngny64qv34a9vzhy39vuv8jwjfsexun5p",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "6427445224",
                                        "delegator_shares": "6427445224.000000000000000000",
                                        "description": {
                                            "moniker": "DeBank",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-12-25T08:04:37.22829156Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1wwspfe7whh3zu4ql5rvpg044lyk6cuu7fpnd9e",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqz7ylkz6wapcwa5k3gn7vh6efd0qk7cqwp2rxh5lf0l7lhvsh4m8qhgp2zf",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "5699129041",
                                        "delegator_shares": "5699129041.000000000000000000",
                                        "description": {
                                            "moniker": "Bit Cat🐱",
                                            "identity": "FAB46CEEAEAB9FA1",
                                            "website": "https://www.bitcat365.com",
                                            "details": "Secure and stable validator service from China team"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.020000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-07-30T09:12:32.806551501Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1w0494h0l4mneaq7ajkrcjvn73m2n04l87j2nst",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepquxeuhp0gj88u5mrukuazzrxc4rnjjuakls4fr2gzxlwj4f9p8lfs965r7z",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "96449451062",
                                        "delegator_shares": "96478391551.647848054450574736",
                                        "description": {
                                            "moniker": "Angel",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "804087",
                                        "unbonding_time": "2020-03-07T13:08:49.834368524Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.500000000000000000"
                                            },
                                            "update_time": "2019-03-18T15:40:26.600089741Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1w42lm7zv55jrh5ggpecg0v643qeatfkd9aqf3f",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqz679nxu2dkfd6y9hytqwvf2z4yuevraqykkm2464ag4e6z278h3qdq92xu",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "761888512720",
                                        "delegator_shares": "761888512720.000000000000000000",
                                        "description": {
                                            "moniker": "Mythos",
                                            "identity": "2E9FDF34351A5112",
                                            "website": "https://mythos.services",
                                            "details": "Staking and validator services for crypto networks"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1we6knm8qartmmh2r0qfpsz6pq0s7emv3e0meuw",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq6adydsk7nw3d63qtn30t5rexhfg56pq44sw4l9ld0tcj6jvnx30s5xw9ar",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1817891090486",
                                        "delegator_shares": "1817891090486.000000000000000000",
                                        "description": {
                                            "moniker": "Staked",
                                            "identity": "E7BFA6515FB02B3B",
                                            "website": "https://staked.us/",
                                            "details": "Staked operates highly available and highly secure, institutional grade staking infrastructure for leading proof-of-stake (PoS) protocols."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.020000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1wauge4px27c257nfn4k3329wteddqw7gs3n66u",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqqaffdhuhdtr0d6nl8twpraxps74q3mxn68qknrex465yd9cc9l0qeh6lkk",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "166849428224",
                                        "delegator_shares": "166849428224.000000000000000000",
                                        "description": {
                                            "moniker": "DappPub",
                                            "identity": "BB2D113EFC6DFDC4",
                                            "website": "https://dapp.pub",
                                            "details": "DappPub, Unleashing the Power of DApps"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-05-15T10:37:35.020580984Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper102ruvpv2srmunfffxavttxnhezln6fnc54at8c",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq9weu2v0za8fdcvx0w3ps972k5v7sm6h5as9qaznc437vwpfxu37q0f3lyg",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "419722382279",
                                        "delegator_shares": "419722382279.000000000000000000",
                                        "description": {
                                            "moniker": "Ztake.org",
                                            "identity": "09A303A2C724C591",
                                            "website": "https://ztake.org/",
                                            "details": "Support reliable independent validator"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.070000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-08-14T05:12:52.848294105Z"
                                        },
                                        "min_self_delegation": "10"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1000ya26q2cmh399q4c5aaacd9lmmdqp90kw2jn",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqe93asg05nlnj30ej2pe3r8rkeryyuflhtfw3clqjphxn4j3u27msrr63nk",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "457608104422",
                                        "delegator_shares": "457608104422.000000000000000000",
                                        "description": {
                                            "moniker": "Staking Fund",
                                            "identity": "805F39B20E881861",
                                            "website": "https://www.staking.fund",
                                            "details": "Staking Fund has been participating in the validating role since early 2018 and is a proud member of the Never Jailed Crew of Game of Stakes."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.120000000000000000",
                                                "max_rate": "0.334000000000000000",
                                                "max_change_rate": "0.012019031323000000"
                                            },
                                            "update_time": "2020-02-02T00:47:46.138000758Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper10nzaaeh2kq28t3nqsh5m8kmyv90vx7ym5mpakx",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqmuspsp8739l0lgn2qz0arargk6ccfy2p82mwflsrsqzwpvhuh5usuwykf6",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "73145988986",
                                        "delegator_shares": "73145988986.000000000000000000",
                                        "description": {
                                            "moniker": "Blockdaemon",
                                            "identity": "8F898657DE26D645",
                                            "website": "https://blockdaemon.com/node-marketplace/#staking",
                                            "details": "Blockdaemon provides maximum uptime for the Cosmos network so that you can be confident your node will be there, ready and secure, for optimal reward generation. Contact us to stake on Cosmos today."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-12-19T03:19:44.114438899Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-10-29T19:55:36.41178089Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper10e4vsut6suau8tk9m6dnrm0slgd6npe3jx5xpv",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqteacnywz7urnac46wtrcy34myyj82j250ny7866yffypdgavae5s0lf4a0",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "3607499246627",
                                        "delegator_shares": "3607499246627.000000000000000000",
                                        "description": {
                                            "moniker": "B-Harvest",
                                            "identity": "8957C5091FBF4192",
                                            "website": "https://bharvest.io",
                                            "details": "B-Harvest focus on the value of high standard security \u0026 stability, active community participation on Cosmos Network, and real world practical use-case of blockchain technology."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-24T11:21:51.694254562Z"
                                        },
                                        "min_self_delegation": "9000000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1sxx9mszve0gaedz5ld7qdkjkfv8z992ax69k08",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqjnnwe2jsywv0kfc97pz04zkm7tc9k2437cde2my3y5js9t7cw9mstfg3sa",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "2545294437814",
                                        "delegator_shares": "2545294437814.000000000000000000",
                                        "description": {
                                            "moniker": "validator.network | Security first. Highly available.",
                                            "identity": "357F80896B3311B4",
                                            "website": "https://validator.network",
                                            "details": "Highly resilient and secure validator operating out of Northern Europe. See website for terms of service."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-06-30T08:32:22.562118158Z"
                                        },
                                        "min_self_delegation": "100000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1sd4tl9aljmmezzudugs7zlaya7pg2895ws8tfs",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq8y846wm58fmmuctxp7csqmaz3594xnykcean0lp722ntf6u5ycaqss4prd",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1080650483982",
                                        "delegator_shares": "1080650483982.000000000000000000",
                                        "description": {
                                            "moniker": "InfStones (Infinity Stones)",
                                            "identity": "39A41C2FDE0AD040",
                                            "website": "https://infstones.io",
                                            "details": "Fueling Blockchain Beyond Infinity!"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-05-10T06:56:55.530159974Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1s05va5d09xlq3et8mapsesqh6r5lqy7mkhwshm",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqgx5xdrx0xktl5r8e3w7vj329fgh3fnep8ahgx8027nd5nkjxzuqs5us5en",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "557166703758",
                                        "delegator_shares": "557166703758.000000000000000000",
                                        "description": {
                                            "moniker": "Wetez",
                                            "identity": "26FA2B24F46A98EF",
                                            "website": "https://www.wetez.io",
                                            "details": "Wetez is the most professional team in the POS ( Proof of Stake) field.Wetez是POS领域最专业的团队，为POS带来的权益做更多赋能。"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-03-25T02:35:12.908955321Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ssm0d433seakyak8kcf93yefhknjleeds4y3em",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqrgyyjxpe0ujefxwnkpmqz9m0hj03y09tdz9lwc0s7mvy469hulfq69f8sd",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1548438266517",
                                        "delegator_shares": "1548438266517.000000000000000000",
                                        "description": {
                                            "moniker": "IRISnet-Bianjie",
                                            "identity": "DB667A6F239969F5",
                                            "website": "https://irisnet.org/irisnet-bianjie",
                                            "details": "Interchain Service Hub for NextGen Distributed Applications."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.020000000000000000"
                                            },
                                            "update_time": "2019-06-21T04:33:57.044358454Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq6fpkt3qn9xd7u44478ypkhrvtx45uhfj3uhdny420hzgsssrvh3qnzwdpe",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "12598748851227",
                                        "delegator_shares": "12598748851227.000000000000000000",
                                        "description": {
                                            "moniker": "🐠stake.fish",
                                            "identity": "90B597A673FC950E",
                                            "website": "stake.fish",
                                            "details": "We are the leading staking service provider for blockchain projects. Join our community to help secure networks and earn rewards. We know staking."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.040000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1s6x9fy4wc49wj9jx4jv6czredqsmp46h7vnk2q",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqfkkyuexns2l7rw2mx2ms988heah0rjv42e9q88scc3ms5hzg45psycrvr4",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "735675714417",
                                        "delegator_shares": "735675714417.000000000000000000",
                                        "description": {
                                            "moniker": "SNZPool",
                                            "identity": "FF2019D4CF1F3185",
                                            "website": "https://snzholding.com",
                                            "details": "SNZ is a crypto assets capital, consulting agency, community builder and professional \u0026 reliable POS validator for a dozen of projects like Cosmos, IRISnet, EOS, ONT, Loom, etc."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-05-31T14:02:51.669843605Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1s6t3wzx6mcv3pjg5fp2ddzplm3gj4pg6d330wg",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqncdd6lvm4r42eke822e5eg0alentpvlxjwzat7nvpynlp0vcu55sl5z96g",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "245002260000",
                                        "delegator_shares": "245002260000.000000000000000000",
                                        "description": {
                                            "moniker": "omega3",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-03-22T01:22:56.42711479Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1s65zmn32zugl57ysj47s7vmfcek0rtd7he7wde",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqrc6g9m2eyy4zs7kyeph8vk5ldpgnceveelc39zf7lc32j8k3shqssevdlg",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "194202250000",
                                        "delegator_shares": "194202250000.000000000000000000",
                                        "description": {
                                            "moniker": "firstblock",
                                            "identity": "23D9B8528FC93D58",
                                            "website": "https://firstblock.io",
                                            "details": "You Delegate. We Validate."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-01-09T03:45:47.112144406Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1s7jnk7t6yqzensdgpvkvkag022udk842qdjdtd",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqnnh28nlj55sc329ppnhcr0xx7kuc9vnsp3dpwc28wdhhxtjc7jfs9k57f7",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "336505666793",
                                        "delegator_shares": "336505666793.000000000000000000",
                                        "description": {
                                            "moniker": "Blockscale",
                                            "identity": "F38EDEA063FD446C",
                                            "website": "https://blockscale.net",
                                            "details": "Planet-scale blockchain infrastructure."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-06-14T07:33:07.032714186Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.250000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-09-21T01:23:56.63727699Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper132juzk0gdmwuxvx4phug7m3ymyatxlh9734g4w",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq9xu9z6ky3nz3k544ar4zhupjehkxdlpmt2l90kekxkrvuu7hxfgslcdqwy",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "2451022343872",
                                        "delegator_shares": "2451022343872.000000000000000000",
                                        "description": {
                                            "moniker": "P2P.ORG - P2P Validator",
                                            "identity": "E12F4695036D8072",
                                            "website": "https://p2p.org",
                                            "details": "One of the winners of Cosmos Game of Stakes. We provide a simple, secure and intelligent staking service to help you generate rewards on your blockchain assets across 9+ networks within a single interface. Let’s stake together - p2p.org."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.010000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-12-11T17:13:40.302520375Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper130mdu9a0etmeuw52qfxk73pn0ga6gawkxsrlwf",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqfahazsjeru5wqulfuzklmkh272ggss2ru6fk00zq2fmlfzcq773sqlqe42",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1127978319378",
                                        "delegator_shares": "1127978319378.000000000000000000",
                                        "description": {
                                            "moniker": "jackzampolin",
                                            "identity": "0979483D4F669CFF",
                                            "website": "https://pylonvalidator.com",
                                            "details": "'You must construct additional pylons' -StarCraft"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.150000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-03-25T15:28:01.171914203Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper13sduv92y3xdhy3rpmhakrc3v7t37e7ps9l0kpv",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqqddwwkhkfrsd66u49kg3h6q36t4kv557vlszqaed4c3y936ncq9s0r0tm2",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1647613576303",
                                        "delegator_shares": "1647613576303.000000000000000000",
                                        "description": {
                                            "moniker": "nylira.net",
                                            "identity": "6A0D65E29A4CBC8E",
                                            "website": "https://nylira.net",
                                            "details": "Stake and earn with security and peace of mind. Operated by Peng Zhong."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-05-14T21:33:23.86382871Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper13ce7hhqa0z3tpc2l7jm0lcvwe073hdkkpp2nst",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqmx0dcpmd9uq7ueck8d880lrt8tp9kvkfmaz0mtv0arye2cda2zrsrlla3n",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "61552309043",
                                        "delegator_shares": "61552309043.000000000000000000",
                                        "description": {
                                            "moniker": "RockX",
                                            "identity": "A15B586AB203F14E",
                                            "website": "www.rockx.com",
                                            "details": "Unlocking the full value of digital assets and decentralized governance"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-12-24T17:24:12.507160243Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-09-05T08:00:55.144718017Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1jxv0u20scum4trha72c7ltfgfqef6nsch7q6cu",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqnru7aa6ayyuwddd5qsa6tvutzs7xl9jk6pfx4ka5dr4y9d3q6eesgz9rz7",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "302293155457",
                                        "delegator_shares": "302293155457.000000000000000000",
                                        "description": {
                                            "moniker": "Ping.pub",
                                            "identity": "6EA723DA332200B2",
                                            "website": "https://ping.pub",
                                            "details": "We are one of the most secure and stable validator, welcome to delegate to us. 我们是最安全，最稳定，性价比最高的验证人节点，欢迎委托给我们！"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.020000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.200000000000000000"
                                            },
                                            "update_time": "2019-11-18T12:14:35.328556Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1j0vaeh27t4rll7zhmarwcuq8xtrmvqhudrgcky",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqvn4a4skwj88c8e0jvns3qjrhyy0whvnuwmth3k8kexvqk5vupw4qsdje47",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1071888285988",
                                        "delegator_shares": "1071888285988.000000000000000000",
                                        "description": {
                                            "moniker": "chainflow-cosmos-prodval-01",
                                            "identity": "81D443FA08A4A926",
                                            "website": "https://chainflow.io/cosmos",
                                            "details": "Operated by Chris Remus (Twitter @cjremus) / Validating since the Validator Working Group formed in October 2017"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.120000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-04-01T15:19:20.666364183Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1jlr62guqwrwkdt4m3y00zh2rrsamhjf9num5xr",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq5e8w7t7k9pwfewgrwy8vn6cghk0x49chx64vt0054yl4wwsmjgrqfackxm",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "723622196447",
                                        "delegator_shares": "723622196447.000000000000000000",
                                        "description": {
                                            "moniker": "StakeWith.Us",
                                            "identity": "609F83752053AD57",
                                            "website": "https://stakewith.us",
                                            "details": "Secured Staking Made Easy. Put Your Crypto to Work - Hassle Free. Disclaimer: Delegators should understand that delegation comes with slashing risk. By delegating to StakeWithUs Pte Ltd, you acknowledge that StakeWithUs Pte Ltd is not liable for any losses on your investment."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-04-22T12:01:25.715205208Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1n5pu2rtz4e2skaeatcmlexza7kheedzh8a2680",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqnd9kzfhhvuv5k2cq62yu0e5v73ymsgxa0wlen9c7999ucwg7hg6qdm34pm",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "576060849627",
                                        "delegator_shares": "576060849627.000000000000000000",
                                        "description": {
                                            "moniker": "BlockMatrix 🚀",
                                            "identity": "DA33F58EC17769B4",
                                            "website": "https://blockmatrix.network",
                                            "details": "Experienced validator across multiple PoS and DPoS networks. Winners in the Game of Stakes. Cosmos FTW!"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "10"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1nm0rrq86ucezaf8uj35pq9fpwr5r82clzyvtd8",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqsnngsdda53d9aqwezvpsx4uh2nkwkn4nra5lw4tyl9n3m02q4kvsrqq0pw",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "145001000000",
                                        "delegator_shares": "145001000000.000000000000000000",
                                        "description": {
                                            "moniker": "Cthulhu",
                                            "identity": "Cthulhu",
                                            "website": "Cthul.hu",
                                            "details": "Cthulhu"
                                        },
                                        "unbonding_height": "644758",
                                        "unbonding_time": "2020-02-23T10:33:51.593569686Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "1.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2020-01-22T08:27:20.214359212Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper15r4tc0m6hc7z8drq3dzlrtcs6rq2q9l2nvwher",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqjcp9ez3dzmvsdfcw2h5kllmqvjgqnhtlvhad4q9wzcqf34gf6ewq6zl5mm",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "730533169795",
                                        "delegator_shares": "730533169795.000000000000000000",
                                        "description": {
                                            "moniker": "DragonStake",
                                            "identity": "EA61A46F31742B22",
                                            "website": "https://dragonstake.io",
                                            "details": "Forking the Banks"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-03-24T18:52:31.67294751Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper159eexl76jlygrxnfreehl3j9tt70d8wfnn39fw",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq0jhujmf2ur4uk7al8pht6rpxwf7a24gqmhggeche0w09fglj9tmss4a0ql",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "8594550000",
                                        "delegator_shares": "8594550000.000000000000000000",
                                        "description": {
                                            "moniker": "fishegg.net",
                                            "identity": "",
                                            "website": "http://www.fishegg.net",
                                            "details": "welcome to join staking"
                                        },
                                        "unbonding_height": "7949",
                                        "unbonding_time": "2020-01-02T07:58:37.125916904Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-12-13T14:17:35.977287639Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper156gqf9837u7d4c4678yt3rl4ls9c5vuursrrzf",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqtw8862dhw8uty58d6t2szfd6kqram2t234zjteaaeem6l45wclaq8l60gn",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "9408964703954",
                                        "delegator_shares": "9408964703954.000000000000000000",
                                        "description": {
                                            "moniker": "Binance Staking",
                                            "identity": "",
                                            "website": "https://binance.com",
                                            "details": "Exchange the world"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.025000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-03-06T07:47:39.033746293Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper15urq2dtp9qce4fyc85m6upwm9xul3049e02707",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqjc07nu2ya8tyzl8m385rnc382pkulwt2gh8yary73f3a96jak7pqsf63xf",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "4800643669275",
                                        "delegator_shares": "4800643669275.000000000000000000",
                                        "description": {
                                            "moniker": "Chorus One",
                                            "identity": "00B79D689B7DC1CE",
                                            "website": "https://chorus.one/",
                                            "details": "Secure Cosmos and shape its future by delegating to Chorus One, a highly secure and stable validator. By delegating, you agree to the terms of service at: https://chorus.one/cosmos/tos"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.075000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.020000000000000000"
                                            },
                                            "update_time": "2019-08-13T17:43:26.871706216Z"
                                        },
                                        "min_self_delegation": "10"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1402ggxz5u6vm29sqztwqq8vxs3ke6dmwl2z5dk",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqd4qh9dqgyce948kn48r2aqk7qlgtuwdmfeewj0z9aj4dr30v33cq6pmlql",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "9831948191",
                                        "delegator_shares": "9832931400.858711612156228525",
                                        "description": {
                                            "moniker": "Cosmoon",
                                            "identity": "8935A6F323FA0881",
                                            "website": "https://cosmoon.org/",
                                            "details": " Professional Stake Rewards Service "
                                        },
                                        "unbonding_height": "1126591",
                                        "unbonding_time": "2020-04-03T05:56:12.40115085Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2020-01-27T22:18:16.158833459Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper14kn0kk33szpwus9nh8n87fjel8djx0y070ymmj",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqmfxl36td7rcdzszzrk6c7kzp5l3jlw4lnxz8zms3py7qcsa9xlns7zxfd6",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "2279800547152",
                                        "delegator_shares": "2279800547152.000000000000000000",
                                        "description": {
                                            "moniker": "Forbole",
                                            "identity": "2861F5EE06627224",
                                            "website": "https://www.forbole.com/cosmos-hub-validator/",
                                            "details": "As a prominent validator and contributor in Cosmos since 2017, Forbole is devoted to build a stronger Cosmos ecosystem. We are award winners in Game of Stakes and HackAtom. Please join our [community](https://t.me/forbole) or visit [our website](https://www.forbole.com/)."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.095000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-11-10T16:52:29.417872059Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper14k4pzckkre6uxxyd2lnhnpp8sngys9m6hl6ml7",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepquhlqdhjw4qp2c2t6qh5z7tfk52qc72623f0etc8f3n7hy8uuh25ql34fvu",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "11370687848139",
                                        "delegator_shares": "11370687848139.000000000000000000",
                                        "description": {
                                            "moniker": "Polychain Labs",
                                            "identity": "A51CE3B9CD649C3F",
                                            "website": "https://cosmos.polychainlabs.com",
                                            "details": "Secure staking with Polychain Labs, the most experienced institutional grade staking team."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-04-22T04:57:29.717811274Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper146kwpzhmleafmhtaxulfptyhnvwxzlvm87hwnm",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqfc7vnpgls3an0w2pv60pu4vr30p2dxqlmhmlrdv0m38y3tg689vs5qg4u5",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "85928911865",
                                        "delegator_shares": "85928911865.000000000000000000",
                                        "description": {
                                            "moniker": "🌐 KysenPool.io",
                                            "identity": "2474A8FCC4426BC5",
                                            "website": "https://www.kysenpool.io",
                                            "details": "Based in Silicon Valley. Help secure Cosmos by delegating to Kysen. Validators are backed by HSMs in Tier 3 enterprise-grade data centers."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-09-06T03:47:00.801583378Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.079000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-03-23T06:10:09.194879379Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper14az9dmutwtz4vuycvae8csm4wwwtm0aumtlppe",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq59t2nm3ph5k6uc804w0n7ey69ul8ntee2dy47d7u53q248ud822sunv93j",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1745806120297",
                                        "delegator_shares": "1745980718357.935064116083909306",
                                        "description": {
                                            "moniker": "F4RM",
                                            "identity": "181FAE6C0E4FA498",
                                            "website": "http://www.f4rm.io",
                                            "details": "F4RM - secure network validation \u0026 pooled digital asset staking"
                                        },
                                        "unbonding_height": "546055",
                                        "unbonding_time": "2020-02-15T08:28:57.815140271Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-15T10:36:44.571237954Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper14l0fp639yudfl46zauvv8rkzjgd4u0zk2aseys",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq7jsrkl9fgqk0wj3ahmfr8pgxj6vakj2wzn656s8pehh0zhv2w5as5gd80a",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1970826215820",
                                        "delegator_shares": "1970826215820.000000000000000000",
                                        "description": {
                                            "moniker": "ATEAM",
                                            "identity": "0CB9A4E7643FF992",
                                            "website": "nodeateam.com",
                                            "details": "[GOS Never Jailed crew \u0026 Top-Tier] Node A-Team promises to provide validator node operation services at the highest quality"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.099000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-03-15T00:18:37.289241198Z"
                                        },
                                        "min_self_delegation": "5000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper14lultfckehtszvzw4ehu0apvsr77afvyju5zzy",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqp0j4vum7ryt6nl6zsgq9ar347afmq2c5z6jmzeavv2p2ns6m0dgs5zmg4z",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "11215950232520",
                                        "delegator_shares": "11215950232520.000000000000000000",
                                        "description": {
                                            "moniker": "DokiaCapital",
                                            "identity": "25422F4ADF3F6765",
                                            "website": "https://staking.dokia.cloud",
                                            "details": "Downtime is not an option for Dokia Capital. We operate an enterprise-grade infrastructure that is robust and secure."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.150000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1k9a0cs97vul8w2vwknlfmpez6prv8klv03lv3d",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqfgpyq4xk4s96ksmkfrr7juea9kmdxkl5ht94xgpxe240743u9cvsht489p",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "900772962204",
                                        "delegator_shares": "900772962204.000000000000000000",
                                        "description": {
                                            "moniker": "Stake Capital",
                                            "identity": "1DD9A932591FA928",
                                            "website": "https://stake.capital",
                                            "details": "'Trustless Digital Asset Management', Twitter: @StakeCapital, operated by @bneiluj @leopoldjoy"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.080000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.030000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1kgddca7qj96z0qcxr2c45z73cfl0c75p7f3s2e",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqay5ldqdmyzy9qfr93enxmm7cwsd5aafz6huqvczytqahpw4twa8qvtrwhv",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "599883677526",
                                        "delegator_shares": "599883677526.000000000000000000",
                                        "description": {
                                            "moniker": "ChainLayer",
                                            "identity": "AD3CDBC91802F94A",
                                            "website": "https://www.chainlayer.io",
                                            "details": "Secure and reliable validator. TG: https://t.me/chainlayer"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-04-01T18:02:13.514368678Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ktecz4dr56j9tsfh7nwg8s9suvhfu70qykhu5s",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq4euv7ertqhgvxrla583fg9g6z2v2dzrkl9spche4j4r23vukmx2q8gqvev",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "5766932920",
                                        "delegator_shares": "5766932920.000000000000000000",
                                        "description": {
                                            "moniker": "Dawns.World",
                                            "identity": "AA70E5B206F952A3",
                                            "website": "https://dawns.world",
                                            "details": "To discover token's intrinsic real value and enhance its liquidity"
                                        },
                                        "unbonding_height": "1282180",
                                        "unbonding_time": "2020-04-16T04:36:32.319738894Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-12-17T15:57:42.95929171Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1kj0h4kn4z5xvedu2nd9c4a9a559wvpuvu0h6qn",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqvc5xdrpvduse3fc084s56n4a6dhzudyzjmywjx25fkgw2fhsj70searwgy",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1564634211770",
                                        "delegator_shares": "1564634211770.000000000000000000",
                                        "description": {
                                            "moniker": "Cryptium Labs",
                                            "identity": "5A309B5CA189D8B3",
                                            "website": "https://cryptium.ch",
                                            "details": "Secure and available validation from the Swiss Alps"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.110000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-27T11:16:35.284063151Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1kn3wugetjuy4zetlq6wadchfhvu3x740ae6z6x",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqc8slfqdszcd85wzzweuanv0em4h4gdc5wkh3et6e7t8z93z24u0s0rdlx2",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "3445647543790",
                                        "delegator_shares": "3445647543790.000000000000000000",
                                        "description": {
                                            "moniker": "HuobiPool",
                                            "identity": "23536C5BDE3EB949",
                                            "website": "https://www.huobipool.com/",
                                            "details": "Huobi Pool is a sub-brand of Huobi Group, which is an important part of the global ecological strategy of Huobi.Huobi Pool has become one of the largest POS communities in the Asia- Pacific region, the leading POW pool and nodes of a number of public chains."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.010000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-04-05T10:43:26.93220487Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1hvsdf03tl6w5pnfvfv5g8uphjd4wfw2h4gvnl7",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq5l63vgd8m9chc3c32wn5lthzsax6xxdylpzmhqmjwrgfhd3m2swsj2wc2d",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "100005087848",
                                        "delegator_shares": "100015089026.194888154307376423",
                                        "description": {
                                            "moniker": "Atom.Bi23",
                                            "identity": "EB3470949B3E89E2",
                                            "website": "https://atom.bi23.com",
                                            "details": "Bi23 focuses on the Crypto-Assets, providing customers with Staking and DeFi services."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-08-16T15:34:04.702756371Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.500000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-03-11T04:30:32.134491966Z"
                                        },
                                        "min_self_delegation": "2"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1hjct6q7npsspsg3dgvzk3sdf89spmlpfdn6m9d",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqnltddase4lqjcfhup8ymg0qex3srakg54ppv06pstvwdjxkm6tmq08znvs",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "4653603791655",
                                        "delegator_shares": "4653603791655.000000000000000000",
                                        "description": {
                                            "moniker": "Figment Networks",
                                            "identity": "E5F274B870BDA01D",
                                            "website": "https://figment.network",
                                            "details": "Makers of Hubble and Canada’s largest Cosmos validator, Figment is the easiest and most secure way to stake your Atoms."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.090000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-12-06T12:17:54.693866931Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1crqm3598z6qmyn2kkcl9dz7uqs4qdqnr6s8jdn",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqt0fpxxufuuhavfqh8zg3pjnnwdvvzw9huemzxe59kpjt5e3xprhs7d8khn",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "447018398440",
                                        "delegator_shares": "447018398440.000000000000000000",
                                        "description": {
                                            "moniker": "Bison Trails",
                                            "identity": "A296556FF603197C",
                                            "website": "bisontrails.co",
                                            "details": "Bison Trails is the easiest way to run infrastructure on multiple blockchains."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.080000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-10-30T20:21:03.030621767Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1cgh5ksjwy2sd407lyre4l3uj2fdrqhpkzp06e6",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq3f6wnsk6k6qu6g8n5vly4z7ajw7q930wh3qx6zkxhktnh49l40kszf5lry",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "932859245807",
                                        "delegator_shares": "932859245807.000000000000000000",
                                        "description": {
                                            "moniker": "HashQuark",
                                            "identity": "31AFBBE0A52FA1ED",
                                            "website": "https://www.hashquark.io",
                                            "details": "Staking made easier!"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-10-23T02:48:39.22540691Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1clpqr4nrk4khgkxj78fcwwh6dl3uw4epsluffn",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq0dc9apn3pz2x2qyujcnl2heqq4aceput2uaucuvhrjts75q0rv5smjjn7v",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "6103920214139",
                                        "delegator_shares": "6103920214139.000000000000000000",
                                        "description": {
                                            "moniker": "Cosmostation",
                                            "identity": "AE4C403A6E7AA1AC",
                                            "website": "https://www.cosmostation.io",
                                            "details": "CØSMOSTATION Validator. Delegate your atoms and Start Earning Staking Rewards"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.120000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "10"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ey69r37gfxvxg62sh4r0ktpuc46pzjrm873ae8",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqg6y8magedjwr9p6s2c28zp28jdjtecxhn97ew6tnuzqklg63zgfspp9y3n",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "10361463918660",
                                        "delegator_shares": "10361463917414.465324179280152398",
                                        "description": {
                                            "moniker": "Sikka",
                                            "identity": "https://keybase.io/team/sikka",
                                            "website": "sikka.tech",
                                            "details": "Sunny Aggarwal (@sunnya97) and Dev Ojha (@ValarDragon)"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.030000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-11-01T04:08:08.548659287Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1e859xaue4k2jzqw20cv6l7p3tmc378pc3k8g2u",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepql9gmstvlcyxpa7ndyl08y694xxc8r03e8nn0zfvl3x255ev6q7rq9hevuz",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "2164966230",
                                        "delegator_shares": "2164966230.000000000000000000",
                                        "description": {
                                            "moniker": "#fuckgoogle",
                                            "identity": "",
                                            "website": "https://github.com/cybercongress",
                                            "details": ""
                                        },
                                        "unbonding_height": "1357109",
                                        "unbonding_time": "2020-04-22T09:09:26.8630982Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-12-12T12:34:44.598269429Z"
                                        },
                                        "min_self_delegation": "10"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1et77usu8q2hargvyusl4qzryev8x8t9wwqkxfs",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqfx0p8s3gmxmaftkkazw5wag4sfau3vgcn20ut4dn5rv2nr8ddq2s59rnvq",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "6938017012",
                                        "delegator_shares": "6938017012.000000000000000000",
                                        "description": {
                                            "moniker": "👾replicator.network",
                                            "identity": "9203983F91296B66",
                                            "website": "https://replicator.network",
                                            "details": ""
                                        },
                                        "unbonding_height": "8201",
                                        "unbonding_time": "2020-01-02T08:27:54.43078638Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-12-12T04:01:41.501864578Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1e0plfg475phrsvrlzw8gwppeva0zk5yg9fgg8c",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqz83dmnt4g6w0w6syrf433mwpk86zejxnq6e336xtxd8pg9jtxkgq732tpu",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "331202100073",
                                        "delegator_shares": "331202100073.000000000000000000",
                                        "description": {
                                            "moniker": "Easy 2 Stake",
                                            "identity": "2C877AC873132C91",
                                            "website": "www.easy2stake.com",
                                            "details": "Easy.Stake.Trust. as easy and as simple as you would click next. Complete transparency and trust with a secure and stable validator. GoS winner, Never Jailed Crew"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.300000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2019-09-27T08:35:53.771679331Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1e0jnq2sun3dzjh8p2xq95kk0expwmd7sj6x59m",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqfndze9l7th79g0m4fguf5cueksmywl6sw2xhuz7gp6jatr39ffuqn3exk9",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "116267146692",
                                        "delegator_shares": "116267146692.000000000000000000",
                                        "description": {
                                            "moniker": "Fission Labs",
                                            "identity": "7DAC30FBD99879B0",
                                            "website": "https://fissionlabs.io/",
                                            "details": "Fission Labs - Blockchain infrastructure and development services"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-10-27T15:17:00.635446897Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1eh5mwu044gd5ntkkc2xgfg8247mgc56fz4sdg3",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq8hu49qdl5594rzxmdsww3hleu8phxrajjfsseqjere9mjrrrv9tq35mll4",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "2135806970012",
                                        "delegator_shares": "2135806970012.000000000000000000",
                                        "description": {
                                            "moniker": "BouBouNode",
                                            "identity": "",
                                            "website": "https://boubounode.com",
                                            "details": "AI-based Validator. #1 AI Validator on Game of Stakes. Fairly priced. Don't trust (humans), verify. Made with BouBou love."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.061000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ehkfl7palwrh6w2hhr2yfrgrq8jetgucudztfe",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqvmmhug9hcmm26ce7we0n3esavn4c6tfcfd6zgnuj732ls7khjq4srpg0ft",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1138068194456",
                                        "delegator_shares": "1138068194456.000000000000000000",
                                        "description": {
                                            "moniker": "KalpaTech",
                                            "identity": "B4AD06F0EB355573",
                                            "website": "http://kalpatech.co",
                                            "details": "KalpaTech | Genesis Validator | Game of Stakes winner | Services dedicated exclusively for Cosmos  Hub | All resources put in one network"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-28T22:00:22.763748595Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ec3p6a75mqwkv33zt543n6cnxqwun37rr5xlqv",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqd85nu5nelvcyyzcsrr0yaglh8rfvn6cv9pp3p0hgmwtk8hf3cazqc7vz5c",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1049353149259",
                                        "delegator_shares": "1049353149259.000000000000000000",
                                        "description": {
                                            "moniker": "lunamint",
                                            "identity": "4F26823468DD7518",
                                            "website": "https://lunamint.com",
                                            "details": "Always adding value to Cosmos. Check out Lunagram, the Cosmos wallet built into Telegram."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.150000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1emaa7mwgpnpmc7yptm728ytp9quamsvu837nc0",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqfuxvufupnsm7v5anpwd7z8ec70z2k209j7xclnm25zz7vauhyc5qjgxx3h",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "525602853385",
                                        "delegator_shares": "525602853385.000000000000000000",
                                        "description": {
                                            "moniker": "kochacolaj",
                                            "identity": "1E9CE94FD0BA5CFEB901F90BC658D64D85B134D2",
                                            "website": "https://blog.cosmos.network/game-of-stakes-closing-ceremonies-eddb71d3b114#147d",
                                            "details": "Top 5 Game Of Stakes winner"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "1.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2020-02-29T17:40:11.439163513Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1eup5t8pp8jq354heck53qtama7vss9l354kh6r",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqxh4s2zj52uhfssu7u2xyhmnk5f7g9ty368twxkkcfllsq3fqaw9sdl6rj9",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "8452580358",
                                        "delegator_shares": "8452580358.000000000000000000",
                                        "description": {
                                            "moniker": "IZ0",
                                            "identity": "BF964D76855711CC",
                                            "website": "www.izo.ro",
                                            "details": "Izo Data Network ! Commission is 0% . Please join our [community](https://t.me/IzoData) or visit [website](https://www.izo.ro/)! "
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.100000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-12-26T08:56:10.399933806Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper16qme5yxucnaj6snx35nmwze0wyxr8wfgqxsqfw",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqwnhw3azrlhnx9kaujvn0es9u26e4a3af6hye6e9j0pl2tlpx9k3s59zwh0",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "260865453794",
                                        "delegator_shares": "260865453794.000000000000000000",
                                        "description": {
                                            "moniker": "KIRA Staking",
                                            "identity": "C86C8FF08A5269DC",
                                            "website": "https://kiraex.com",
                                            "details": "Kira Core Staking Services - Sentry, KMS, HSM, High Availability \u0026 Double Sign Protection"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.010000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2020-04-01T21:07:55.105865438Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper16v3f95amtvpewuajjcdsvaekuuy4yyzups85ec",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqmgwzcm3aqmc8nln9u4q5ydsjwx6rzqrch6p243x2gtzetnx5l3ls432euv",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "147885388389",
                                        "delegator_shares": "147885388389.000000000000000000",
                                        "description": {
                                            "moniker": "BlockPool",
                                            "identity": "",
                                            "website": "www.blockpool.com",
                                            "details": "Power the staking economy"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.020000000000000000",
                                                "max_rate": "0.030000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-06-22T05:19:17.897735561Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1648ynlpdw7fqa2axt0w2yp3fk542junl7rsvq6",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqf8llkc4p43lksktsqzr5nmgmw4ln9pzym2vp4kqfrny8xrgnqrsq76djjc",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "797707710682",
                                        "delegator_shares": "797787489383.461009934458905148",
                                        "description": {
                                            "moniker": "Any Labs",
                                            "identity": "B2D07CA3CCC907CE",
                                            "website": "https://anylabs.io",
                                            "details": "Blockchain staking and consultancy based in Japan."
                                        },
                                        "unbonding_height": "25880",
                                        "unbonding_time": "2020-01-03T18:48:44.276425288Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-05-08T11:30:02.004384504Z"
                                        },
                                        "min_self_delegation": "100"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper16k579jk6yt2cwmqx9dz5xvq9fug2tekvlu9qdv",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq55mjplg9gy979ua9r5qmk2wr5nysmputt28j0zsgadn933lyh32sh20cmm",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "928458345391",
                                        "delegator_shares": "928551200492.966284028096640587",
                                        "description": {
                                            "moniker": "Cephalopod Equipment",
                                            "identity": "6408AA029ADBE364",
                                            "website": "https://cephalopod.equipment",
                                            "details": "Cephalopod Equipment - infrastructure for decentralized intelligence"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "2019-12-01T09:34:39.548038382Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.081100000000000000",
                                                "max_rate": "0.420000000000000000",
                                                "max_change_rate": "0.011800000000000000"
                                            },
                                            "update_time": "2019-09-19T12:26:43.48042061Z"
                                        },
                                        "min_self_delegation": "100000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1mykn77lkynl8fkwvl9tqg369u0zajzzcdhkptq",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqft6uxfmfjjce0p7ke4h0zc38x4d9d38wlmrgcc47flru92qq3ydq76mrsf",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "89943764382",
                                        "delegator_shares": "89943764382.000000000000000000",
                                        "description": {
                                            "moniker": "Nodeasy.com",
                                            "identity": "AB006A79DBD8FC57",
                                            "website": "https://www.nodeasy.com",
                                            "details": "Nodeasy.com，助你进入Staking时代！"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2019-10-28T13:02:55.182998044Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1m83cwjucw9nt8xm66u8xavvy6v9m7xfspcszc5",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq7qnf5r40z7esjc2utrjrvzxg9sfd683hw0805ek85ddchdcptthqjnzxxu",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "159214462425",
                                        "delegator_shares": "159214462425.000000000000000000",
                                        "description": {
                                            "moniker": "Fenbushi US - Staked",
                                            "identity": "CC4B238C8F9FB2BE",
                                            "website": "https://fenbushi.vc",
                                            "details": "Fenbushi Capital is the first and most active blockchain-focused venture capital firm in Asia. Staked is the leading provider of validation technology and services. We're bringing our combined skills to Cosmos."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-11-15T18:07:08.897219336Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ma02nlc7lchu7caufyrrqt4r6v2mpsj90y9wzd",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqxtu8am2qmf0qnglqtvkar9gaclhccfn29tsp7n82vasrtnc8m2fsulp4h2",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "3464569661466",
                                        "delegator_shares": "3464569661466.000000000000000000",
                                        "description": {
                                            "moniker": "hashtower",
                                            "identity": "0BBBAE1FD11AEBAF",
                                            "website": "http://hashtower.com",
                                            "details": "Hashtower Actwo COSMOS Validator"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.030000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.200000000000000000"
                                            },
                                            "update_time": "2019-07-16T08:46:06.560077744Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1uxh465053nq3at4dn0jywgwq3s9sme3la3drx6",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqc5y2du793cjut0cn6v7thp3xlvphggk6rt2dhw9ekjla5wtkm7nstmv5vy",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "520672159921",
                                        "delegator_shares": "520672159921.000000000000000000",
                                        "description": {
                                            "moniker": "Bison Trails",
                                            "identity": "A296556FF603197C",
                                            "website": "https://bisontrails.co",
                                            "details": "Bison Trails is the easiest way to run infrastructure on multiple blockchains."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.500000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1udpsgkgyutgsglauk9vk9rs03a3skc62gup9ny",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq38tmpw8wujah8nhvkkd26tskx4p7l0qx2kqwfkp3hj8644e9kevqxe5zl2",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "5000000000",
                                        "delegator_shares": "5000000000.000000000000000000",
                                        "description": {
                                            "moniker": "TEST_NODE",
                                            "identity": "",
                                            "website": "",
                                            "details": "Test Node."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.075000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-12-30T11:04:05.791970925Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1uhnsxv6m83jj3328mhrql7yax3nge5svrv6t6c",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepql42t7mstnewp5rgweteuw95hawzystll7mq8dl24n5yh0th7q2jqetcy07",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "720793765456",
                                        "delegator_shares": "720865851960.598464555143928151",
                                        "description": {
                                            "moniker": "Skystar Capital",
                                            "identity": "",
                                            "website": "",
                                            "details": ""
                                        },
                                        "unbonding_height": "1168707",
                                        "unbonding_time": "2020-04-06T17:50:47.447745873Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1u6ddcsjueax884l3tfrs66497c7g86skn7pa0u",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq87zcnf8sm4ewacjafqujfevt8rhwj5qk9uwtx4ef89ctuqmndkeq446ahw",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "809489002763",
                                        "delegator_shares": "809489002763.000000000000000000",
                                        "description": {
                                            "moniker": "Sentinel",
                                            "identity": "D54C8032CF19C407",
                                            "website": "https://sentinel.co",
                                            "details": "We are team Sentinel, developer of infrastructure tools on Cosmos \u0026 other networks.Winner in the Uptime category during GOS.Developed the first working version of the dVPN which runs on Ethereum \u0026 Sentinel's own Tendermint TestNet"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1uutuwrwt3z2a5z8z3uasml3rftlpmu25aga5c6",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqarrl0ppddzyczwvcqwf3jwd9qwkhxfy6lcv8ep4msk293mlxg39qgf77y3",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "894288539493",
                                        "delegator_shares": "894288539493.000000000000000000",
                                        "description": {
                                            "moniker": "Delega Networks♾ ",
                                            "identity": "1BED7C08416A619F",
                                            "website": "https://delega.io",
                                            "details": "Nodes managed by wimel"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.180000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-04-05T00:45:01.073794845Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1ul2me6vukg2vac2p6ltxmqlaa7jywdgt8q76ag",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq0cet8ez89wj4yz8uencych7aldc5wyyrpx6jvh6n6kxxslumln5sxkq922",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1208359862997",
                                        "delegator_shares": "1208359862997.000000000000000000",
                                        "description": {
                                            "moniker": "HyperblocksPro",
                                            "identity": "B073FA5BAD230585",
                                            "website": "https://hyperblocks.pro/",
                                            "details": "Secure the network and earn rewards with Hyperblocks.pro, one of the first companies in the world fully focused on Proof Of Stake protocols"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.250000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-04-05T23:57:19.320271237Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1a3yjj7d3qnx4spgvjcwjq9cw9snrrrhu5h6jll",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqxjkll4nxla8gtekx2ueq3tc8vv4e6rn79jmdead30jeqlm3kc7eqqx7hs8",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "10041185315",
                                        "delegator_shares": "10041185315.000000000000000000",
                                        "description": {
                                            "moniker": "Coinbase Custody",
                                            "identity": "Coinbase Custody",
                                            "website": "https://custody.coinbase.com",
                                            "details": "Coinbase Custody Cosmos Validator"
                                        },
                                        "unbonding_height": "676534",
                                        "unbonding_time": "2020-02-26T01:41:46.759697123Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.200000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "1.000000000000000000"
                                            },
                                            "update_time": "2020-02-05T00:29:51.081896503Z"
                                        },
                                        "min_self_delegation": "1000000"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper17h2x3j7u44qkrq0sk8ul0r2qr440rwgjkfg0gh",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqc9ppxzktam9v39d9q07h6n98cdm7cgg4l65vq5yvtgruxp5h0yhs8tup68",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "202801407477",
                                        "delegator_shares": "202801407477.000000000000000000",
                                        "description": {
                                            "moniker": "FRESHATOMS",
                                            "identity": "63575EE3F0F9FAFC",
                                            "website": "https://freshatoms.com",
                                            "details": "FreshAtoms runs on bare metal in a SSAE16 SOC2 certified Tier 3 datacenter with geographically distributed private sentry nodes, YubiHSM2 hardware protected keys, with 24/7 monitoring, alerting, and analytics."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-09-23T20:29:05.781322875Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper17mggn4znyeyg25wd7498qxl7r2jhgue8u4qjcq",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqlzmd0spn9m0m3eq9zp93d4w6e5tugamv44yqjzyacelnvra634fqnfec0r",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1079608204356",
                                        "delegator_shares": "1079608204356.000000000000000000",
                                        "description": {
                                            "moniker": "01node",
                                            "identity": "22823CD59617B8E3",
                                            "website": "https://01node.com",
                                            "details": "01node Professional Staking Services for Cosmos, Iris, Terra, Solana, Kava, Polkadot, Skale"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.100000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2019-03-13T23:00:00Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1l9fwl9c77zx850htsr20pq3ltc379xt86ndelm",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqfjgjrj4heptaw6h9nhtkng8hw2zsq5c6e9xzwvnjjx2n6pc7x6yq4ny2qc",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "5057317038",
                                        "delegator_shares": "5057317038.000000000000000000",
                                        "description": {
                                            "moniker": "CosmosLink",
                                            "identity": "3F7807C66CE770B0",
                                            "website": "cosmoslink.network",
                                            "details": "Based on Cosmos network digital asset security value-added service provider"
                                        },
                                        "unbonding_height": "1276519",
                                        "unbonding_time": "2020-04-15T17:26:54.325724353Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-03-18T16:46:59.079404223Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1lktjhnzkpkz3ehrg8psvmwhafg56kfss3q3t8m",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqelcwpat987h9yq0ck6g9fsc8t0mththk547gwvk0w4wnkpl0stnspr3hdc",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "1720474496626",
                                        "delegator_shares": "1720474496626.000000000000000000",
                                        "description": {
                                            "moniker": "Umbrella ☔",
                                            "identity": "A530AC4D75991FE2",
                                            "website": "https://umbrellavalidator.com",
                                            "details": "One of the winners of Cosmos Game of Stakes, and HackAtom3."
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.070400000000000000",
                                                "max_rate": "1.000000000000000000",
                                                "max_change_rate": "0.100000000000000000"
                                            },
                                            "update_time": "2019-08-05T07:10:23.689753607Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1lcwxu50rvvgf9v6jy6q5mrzyhlszwtjxhtscmp",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepqh3jg5ld5xg5q5mcxrzn6fcuq696qqa3ut3azskphpdrty37ervjqcn8mfj",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "9713509339",
                                        "delegator_shares": "9713509339.000000000000000000",
                                        "description": {
                                            "moniker": "stake.zone",
                                            "identity": "0A888728046018EC",
                                            "website": "http://stake.zone",
                                            "details": "operated by nuevax"
                                        },
                                        "unbonding_height": "0",
                                        "unbonding_time": "1970-01-01T00:00:00Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.000000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.050000000000000000"
                                            },
                                            "update_time": "2020-02-13T02:20:41.320594366Z"
                                        },
                                        "min_self_delegation": "1"
                                    },
                                    {
                                        "operator_address": "cosmosvaloper1l6udzyaz8xaxv4hpagwauacm95jlcec3xlht2u",
                                        "consensus_pubkey": "cosmosvalconspub1zcjduepq9t4r8jgr09rsscgacnaklxuf55pg0wq5zfwzw8ycawms5x0hrfhqks3dpe",
                                        "jailed": false,
                                        "status": 2,
                                        "tokens": "10802365900",
                                        "delegator_shares": "10802365900.000000000000000000",
                                        "description": {
                                            "moniker": "StakeHouse",
                                            "identity": "A1AAB1D6D0E8F976",
                                            "website": "stakehouse.org",
                                            "details": "Low fees. No hassle. Enjoy your meal."
                                        },
                                        "unbonding_height": "701066",
                                        "unbonding_time": "2020-02-28T02:04:00.114690756Z",
                                        "commission": {
                                            "commission_rates": {
                                                "rate": "0.010000000000000000",
                                                "max_rate": "0.200000000000000000",
                                                "max_change_rate": "0.010000000000000000"
                                            },
                                            "update_time": "2020-02-11T07:24:35.723449554Z"
                                        },
                                        "min_self_delegation": "1"
                                    }
                                ]
                            }
                        `);
                }
        }

        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};

