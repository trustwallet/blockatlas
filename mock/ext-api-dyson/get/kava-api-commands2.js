/// Kava API Mock
/// See:
/// curl "http://localhost:3000/kava-api/staking/validators?status=bonded"
/// curl "http://localhost:3000/kava-api/staking/pool"
/// curl "http://localhost:3000/kava-api/minting/inflation"
/// curl "https://{kava_rpc}/staking/validators?status=bonded"
/// curl "https://{kava_rpc}/staking/pool"
/// curl "https://{kava_rpc}/minting/inflation"
/// curl "http://localhost:8420/v2/kava/staking/validators"
/// curl "http://localhost:8420/v2/kava/staking/delegations/kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m?Authorization=Bearer"

module.exports = {
    path: "/kava-api/:command1/:command2?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch(params.command1) {
            case 'minting':
                switch(params.command2) {
                    case 'inflation':
                        // status=bonded
                        return JSON.parse(`{"height":"1793129","result":"0.060343307052455856"}`);
                }

            case 'txs': {
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
            }

            case 'staking':
                switch(params.command2) {
                    case 'pool':
                        return JSON.parse(`
                            {
                                "height":"1793093",
                                "result": {
                                    "not_bonded_tokens": "1664170404961",
                                    "bonded_tokens": "86469758994132"
                                }
                            }
                        `);

                    case 'validators':
                        // status=bonded
                        return JSON.parse(`
                            {"height":"1793067","result":[
                                {
                                "operator_address": "kavavaloper1qyc2cfl0nw8r95dsdw534x99wq0xcj9rmxpl7z",
                                "consensus_pubkey": "kavavalconspub1zcjduepqyj4j29k7hn58g7n6ert7mm4m7d0kllrx6h5rzzgpvjdt69r80zsq3az2xq",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23006899100",
                                "delegator_shares": "23009200000.000000000000000000",
                                "description": {
                                    "moniker": "Stake Capital",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": "'Trustless Digital Asset Management', Twitter: @StakeCapital, operated by @bneiluj @leopoldjoy"
                                },
                                "unbonding_height": "1415519",
                                "unbonding_time": "2020-03-28T07:43:47.102756457Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.500000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.500000000000000000"
                                    },
                                    "update_time": "2020-02-05T09:55:51.495267155Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1qfy0e2w62g6j4jg5djcqd4py3zsaeqexjplj2d",
                                "consensus_pubkey": "kavavalconspub1zcjduepqftv90yrm4g4w4mq47rdx9f384yxegfx7qfpkft89x2nlzafv36pq47wgls",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20069063571",
                                "delegator_shares": "20069063571.000000000000000000",
                                "description": {
                                    "moniker": "DragonStake",
                                    "identity": "EA61A46F31742B22",
                                    "website": "https://dragonstake.io",
                                    "security_contact": "dragonstake@protonmail.com",
                                    "details": "Forking the Banks"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "1.000000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "1.000000000000000000"
                                    },
                                    "update_time": "2019-11-20T18:38:18.607985797Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1q3y9qga5hf360dmzta67vp54qz25tmv4hhkk4t",
                                "consensus_pubkey": "kavavalconspub1zcjduepquu3m094f3j9jnklursgngn667tt3rz0ahpt4f7406qzqclc42mnqxlrn7e",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20004700000",
                                "delegator_shares": "20004700000.000000000000000000",
                                "description": {
                                    "moniker": "funky",
                                    "identity": "1EBAA06E87B6DD60",
                                    "website": "https://kava-funkyvalidator.nl",
                                    "security_contact": "",
                                    "details": "Validating with love in Holland for the world :)"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.110000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-21T00:42:13.651988876Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvl7z9xh",
                                "consensus_pubkey": "kavavalconspub1zcjduepqfkq6kfq6v3avvrc6qvh0tr6h97qhqxcxw6yhmsgwrclpm3qdqwwq4zl4kc",
                                "jailed": false,
                                "status": 2,
                                "tokens": "106190067334",
                                "delegator_shares": "106190067334.000000000000000000",
                                "description": {
                                    "moniker": "WeStaking",
                                    "identity": "DA9C5AD3E308E426",
                                    "website": "https://westaking.io",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2020-03-17T04:24:29.070264845Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1pn5k9c5pxmg5f0rycpl9rrx6k6mk85scxf06zx",
                                "consensus_pubkey": "kavavalconspub1zcjduepq5gqsp02excuz03zequ5xkygsj8t0su4g46p00q9p0se5kkhr2fmsusw7w7",
                                "jailed": false,
                                "status": 2,
                                "tokens": "22999900000",
                                "delegator_shares": "22999900000.000000000000000000",
                                "description": {
                                    "moniker": "Stir",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-15T14:45:39.203670522Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1phd8jz25lumudc7ac7rhmupvcqcv7lg3c8dprc",
                                "consensus_pubkey": "kavavalconspub1zcjduepquugppwhq0s4t5q8vh6n3j0t4t3susfasdfapa0zp8a8c36tpgj6qpqvs3s",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2675001000000",
                                "delegator_shares": "2675001000000.000000000000000000",
                                "description": {
                                    "moniker": "the_valiator",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "1.000000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-12-13T08:47:32.032003423Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1pceqe8we7drpfqrutchwy3f99800hhzuw6cc84",
                                "consensus_pubkey": "kavavalconspub1zcjduepqypw0nlaweu77hnlmy37gy82d0cn57g07yy9qrkm36tx5zhp6rjzqzfqta0",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23459640449",
                                "delegator_shares": "23459640449.000000000000000000",
                                "description": {
                                    "moniker": "Galaxy",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.500000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-25T09:53:43.24669958Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1rgcgqmnkeffks7enrv6hk5u4wg3nzfkmqlzjqd",
                                "consensus_pubkey": "kavavalconspub1zcjduepqv3f0m888swgzamf2mh3vh2gaehredz9uktv9rjcyfdjzvzvsthks4e5f4t",
                                "jailed": false,
                                "status": 2,
                                "tokens": "21000000000",
                                "delegator_shares": "21000000000.000000000000000000",
                                "description": {
                                    "moniker": "syncnode",
                                    "identity": "F422F328C14AFBFA",
                                    "website": "wallet.syncnode.ro",
                                    "security_contact": "",
                                    "details": "https://www.linkedin.com/in/gbunea/"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.190000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-03-26T09:09:30.282258822Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1yrm63pqvld8uyzkavz55p2cktpm2gm8jd8xxlu",
                                "consensus_pubkey": "kavavalconspub1zcjduepq6nje57k4x0n0uua5dl26ufpgj9jw2n945hkrv066fv54924vnmeq60d92r",
                                "jailed": false,
                                "status": 2,
                                "tokens": "200010000000",
                                "delegator_shares": "200010000000.000000000000000000",
                                "description": {
                                    "moniker": "Nodeasy.com",
                                    "identity": "F7BABF2C95B02A9F",
                                    "website": "https://www.nodeasy.com",
                                    "security_contact": "",
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
                                    "update_time": "2019-12-04T03:53:57.870109232Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1yjj2wfers947l6n5pynpgsqlz7svc5n8ssl6ye",
                                "consensus_pubkey": "kavavalconspub1zcjduepqftnrqaely46s7tzfkczuqfecuk3664wks5www53x6cj7rtl90zhqq6pacf",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20410000001",
                                "delegator_shares": "20410000001.000000000000000000",
                                "description": {
                                    "moniker": "StakingHub",
                                    "identity": "25D7A4013B5C2A8F",
                                    "website": "http://www.stakinghub.net/",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.200000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-01-14T23:30:14.565194467Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1yna6lete8nwwwctsalzdg04ldqaz73gtn33ydq",
                                "consensus_pubkey": "kavavalconspub1zcjduepqzwu26st99545k4qp79vuuslynk5qaxne9z4ucjfzs5g6646klwrqg5dxfa",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23163160000",
                                "delegator_shares": "23163160000.000000000000000000",
                                "description": {
                                    "moniker": "ChainodeTech",
                                    "identity": "E34BF744FF5BA8A9",
                                    "website": "https://chainode.tech/",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-15T14:35:31.864391537Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper196anr2ycsalg806dz29afklnecuaupvkh5qz6c",
                                "consensus_pubkey": "kavavalconspub1zcjduepqh525w7mx9apyra8jqeclcs90e6u9p22jrqaw8ymuf7mug2m5cg6sw7v7yu",
                                "jailed": false,
                                "status": 2,
                                "tokens": "47900816757",
                                "delegator_shares": "47900816757.000000000000000000",
                                "description": {
                                    "moniker": "Phorest 🌲",
                                    "identity": "A5281CA10EBCDCB5",
                                    "website": "http://phorest.xyz/kava",
                                    "security_contact": "",
                                    "details": "Secure and stable PoS validator"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.000000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2020-01-20T14:02:20.813362703Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1x9hq3rjc48t5upcsr3c209ycgekfasne3l5nkc",
                                "consensus_pubkey": "kavavalconspub1zcjduepqvrf8r43ymaj39x73luu4qsp9sn4sw4t2lh77799mau2739xtrlnsqxyqju",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23114805000",
                                "delegator_shares": "23114805000.000000000000000000",
                                "description": {
                                    "moniker": "Masternode24.de",
                                    "identity": "7EBA7730A85C865C",
                                    "website": "https://masternode24.de/",
                                    "security_contact": "",
                                    "details": "Investieren Sie mit uns zusammen in die besten Blockchain Projekte"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2020-02-20T16:21:56.085890332Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1xftqdxvq0xkv2mu8c5y0jrsc578tak4m9u0s44",
                                "consensus_pubkey": "kavavalconspub1zcjduepqngw3s4hpzj8xscq6dfg4m8a6qeh5xwau0sjrknmlvg2lqdlgp0zszj7ewc",
                                "jailed": false,
                                "status": 2,
                                "tokens": "28160800000",
                                "delegator_shares": "28160800000.000000000000000000",
                                "description": {
                                    "moniker": "Ping.pub",
                                    "identity": "6EA723DA332200B2",
                                    "website": "",
                                    "security_contact": "",
                                    "details": "We are one of the most secure and stable validator, welcome to delegate to us. 我们是最安全，最稳定，性价比最高的验证人节点，欢迎委托给我们！"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.020000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-21T09:12:17.124698721Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1xka27j0jvmq97yunj5wp8fv242lycmax8ejlaf",
                                "consensus_pubkey": "kavavalconspub1zcjduepqx4zjl7vsh8pg4l3ndceu88h3027q0vvmpf7anv5c4xxq0htng5kqf24q07",
                                "jailed": false,
                                "status": 2,
                                "tokens": "3047139316200",
                                "delegator_shares": "3047139316200.000000000000000000",
                                "description": {
                                    "moniker": "page2",
                                    "identity": "",
                                    "website": "https://coinmarketcap.com/2/",
                                    "security_contact": "",
                                    "details": "Lorem Ipsum"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.050000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1xhxzmj8fvkqn76knay9x2chfra826369dhdu2c",
                                "consensus_pubkey": "kavavalconspub1zcjduepqwsmxfu5vdulvnrwej06v4jj5hvdxxqqk82flvznhxkh5whrdnxms2z2kjx",
                                "jailed": false,
                                "status": 2,
                                "tokens": "1050206000000",
                                "delegator_shares": "1050206000000.000000000000000000",
                                "description": {
                                    "moniker": "Figment Networks",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.250000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper18zksjhrefqew0zahmts894p8asscufxvdfq702",
                                "consensus_pubkey": "kavavalconspub1zcjduepq7qm0l9hy8m8a4my03x7rfvcy3vru0y0kv29n2fwcpmrvjda8n23skzza6m",
                                "jailed": false,
                                "status": 2,
                                "tokens": "14040219",
                                "delegator_shares": "14040219.000000000000000000",
                                "description": {
                                    "moniker": "Consulnode",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2020-03-04T19:27:58.259245834Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper18s9m5d5cjf0humjv7mkq8pm47kchwm0r0369cx",
                                "consensus_pubkey": "kavavalconspub1zcjduepqphfe0vnahkpurcc6jn3kesfch6c0afnph8qv9ryx29gpec7rtf3q968vu3",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20010000000",
                                "delegator_shares": "20010000000.000000000000000000",
                                "description": {
                                    "moniker": "Stake5 Labs",
                                    "identity": "C7595F18CC5D4FA6",
                                    "website": "https://www.stake5labs.com",
                                    "security_contact": "",
                                    "details": "Professional PoS node contributor"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:01:13.609429004Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper18cf35l7req0k6ulqapeyv830mrrucn9xj87plr",
                                "consensus_pubkey": "kavavalconspub1zcjduepq7rr7drhvr6d67nydvwxtcqv3k0uv6krfn9ktm23l77eqrra9af9s828zrk",
                                "jailed": false,
                                "status": 2,
                                "tokens": "1171728959900",
                                "delegator_shares": "1171846144514.451445144514451445",
                                "description": {
                                    "moniker": "mxcpospool",
                                    "identity": "",
                                    "website": "https://www.mxc.co",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "1433843",
                                "unbonding_time": "2020-03-29T19:11:06.277346205Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-12-13T10:05:51.300009009Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1g20mhvcpjxp6gzlwhtfcphjehwcl2njqydgu7q",
                                "consensus_pubkey": "kavavalconspub1zcjduepqcg7v7vpgw7tj2f4z6yle5vpwutm3n5zjmlx2dpypphxd0cdvyaeqyeyteg",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23142100000",
                                "delegator_shares": "23142100000.000000000000000000",
                                "description": {
                                    "moniker": "pi-londoner",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.200000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-02-10T23:44:11.226734593Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1gtd040dmljyaty9tkhq0mlqz7saer48v4d608x",
                                "consensus_pubkey": "kavavalconspub1zcjduepqqy2kzt636d5ak6l6gxsxexgtadcm8d59qh3wumgtswpa8hau7apsk4ecma",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2391603214",
                                "delegator_shares": "2391603214.000000000000000000",
                                "description": {
                                    "moniker": "kava.bi23",
                                    "identity": "EB3470949B3E89E2",
                                    "website": "https://kava.bi23.com",
                                    "security_contact": "",
                                    "details": "Bi23 focuses on the Crypto-Assets, providing customers with Staking and DeFi services."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.000000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-12-04T23:42:01.494010426Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1g4qpetrj59a29e4wxpe74x93q4df2czjh8r9ak",
                                "consensus_pubkey": "kavavalconspub1zcjduepqe3evdke2eyzv7ngkwj0w4qarpqad47s6eqsqavesmzhz664ewmzsre8fkj",
                                "jailed": false,
                                "status": 2,
                                "tokens": "978579000000",
                                "delegator_shares": "978579000000.000000000000000000",
                                "description": {
                                    "moniker": "IOSG",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1ffcujj05v6220ccxa6qdnpz3j48ng024ykh2df",
                                "consensus_pubkey": "kavavalconspub1zcjduepqxzqld0trat9nu9j53rzxwvpnpjekxf02eeg8mw3xewjky2qgy69sgh0tcw",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2207656905814",
                                "delegator_shares": "2207656905814.000000000000000000",
                                "description": {
                                    "moniker": "🐠stake.fish",
                                    "identity": "90B597A673FC950E",
                                    "website": "http://stake.fish",
                                    "security_contact": "",
                                    "details": "stake.fish is a reliable validator for PoS blockchains. Stake with us. We know staking."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "1.000000000000000000"
                                    },
                                    "update_time": "2019-12-03T18:59:42.821933888Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1fw7vjc3fphahqxpdjypddlulnltxws8g0mrds7",
                                "consensus_pubkey": "kavavalconspub1zcjduepq0hlydw7ztc5fp7e6k3snlafkz4h45tyaqzxxw5j87g4trdy826yskkmmgc",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20090000000",
                                "delegator_shares": "20090000000.000000000000000000",
                                "description": {
                                    "moniker": "BasBlock",
                                    "identity": "E0A6A3980E464A66",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "1318309",
                                "unbonding_time": "2020-03-20T12:28:42.433079809Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.130000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-02-01T11:28:28.795491395Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1fmas0qlsucg4qwf8mqyrylcg3uluz2ffg8952q",
                                "consensus_pubkey": "kavavalconspub1zcjduepqkv4fkh6u8l068xusy5dhz2jpz969nzszntuavsc0uu6sxxr9wafqewxkac",
                                "jailed": false,
                                "status": 2,
                                "tokens": "22612704810",
                                "delegator_shares": "22612704810.000000000000000000",
                                "description": {
                                    "moniker": "mintonium",
                                    "identity": "0FF7C542EF9422AB",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:00:18.218278987Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper12qn7y04wzr5s3h4dmdtre4q9f4nvc03a9a9qsz",
                                "consensus_pubkey": "kavavalconspub1zcjduepqsnl9ahysm5hypmpfuwyqkha280aa9ushdlvvqza4u7rfmf0pvdxqlla0wj",
                                "jailed": false,
                                "status": 2,
                                "tokens": "22996800100",
                                "delegator_shares": "22999100000.000000000000000000",
                                "description": {
                                    "moniker": "dexhybrid",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "1321471",
                                "unbonding_time": "2020-03-20T18:38:12.071818929Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.250000000000000000",
                                    "max_change_rate": "0.020000000000000000"
                                    },
                                    "update_time": "2019-11-16T14:33:31.092253347Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper12r77hhj6ylvvl4etjm0fmpzh07jum8u7qqd695",
                                "consensus_pubkey": "kavavalconspub1zcjduepq3ndp4jxqsvxlp5crspvu6tcra7zg7cnnxzm3hx6sxhcd409ytzyskcadfv",
                                "jailed": false,
                                "status": 2,
                                "tokens": "10000000",
                                "delegator_shares": "10000000.000000000000000000",
                                "description": {
                                    "moniker": "stefansky",
                                    "identity": "878AB3284CD8CBED",
                                    "website": "http://stefancondurachi.com",
                                    "security_contact": "",
                                    "details": "Software Architect"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-02-03T05:32:50.725944775Z"
                                },
                                "min_self_delegation": "1000000"
                                },
                                {
                                "operator_address": "kavavaloper129kkennsm7na34lu6sn4kwxp7ewes58y4fx6y9",
                                "consensus_pubkey": "kavavalconspub1zcjduepqz354kpmk0970fpvhw6va6ufhasrnh4sgzh82cfd5egm9ggl5gy6qlgwqjt",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23011000000",
                                "delegator_shares": "23011000000.000000000000000000",
                                "description": {
                                    "moniker": "stake.zone",
                                    "identity": "0A888728046018EC",
                                    "website": "http://stake.zone",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.250000000000000000",
                                    "max_change_rate": "0.040000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:01:07.22222101Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper12g40q2parn5z9ewh5xpltmayv6y0q3zs6ddmdg",
                                "consensus_pubkey": "kavavalconspub1zcjduepq0plta99ns8ec5h5sj8x4ufzkdpsz6xu7ckp62hkagfxey242qzgstm0ldw",
                                "jailed": false,
                                "status": 2,
                                "tokens": "5085834241836",
                                "delegator_shares": "5085834241836.000000000000000000",
                                "description": {
                                    "moniker": "P2P.ORG - P2P Validator",
                                    "identity": "E12F4695036D8072",
                                    "website": "https://p2p.org",
                                    "security_contact": "vladimir.m@p2p.org",
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
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper12et238paeqxyhvk2pfs20ygfj6ct0dx2ccsdz5",
                                "consensus_pubkey": "kavavalconspub1zcjduepqzmls593jxfkz8unn306c8nh757n24ztxhnzwzvpuadalk0skuvfsaamxwx",
                                "jailed": false,
                                "status": 2,
                                "tokens": "3000000",
                                "delegator_shares": "3000000.000000000000000000",
                                "description": {
                                    "moniker": "hippo",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "1.000000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2020-03-10T14:45:34.603416711Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1t5l8ht0wxpd4lpe4cweftrpg5kyn3qp437yvnr",
                                "consensus_pubkey": "kavavalconspub1zcjduepqe58zj7f5zjqu023l3sd96qwurtzy7ayv2549z6mca8zmalvtlxfs80qsja",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20517050000",
                                "delegator_shares": "20517050000.000000000000000000",
                                "description": {
                                    "moniker": "DelegaNetworks",
                                    "identity": "1BED7C08416A619F",
                                    "website": "https://delega.io",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-24T20:37:25.840600387Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1vyx7wt8s8dwcspdt7dy49sq4l3jwyhxyndakmm",
                                "consensus_pubkey": "kavavalconspub1zcjduepq2v03gwaywadc7ar423rxmy7k7ev3y3tkvsetq7w30j0whymjkacsxjg9lt",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2999557167252",
                                "delegator_shares": "2999557167252.000000000000000000",
                                "description": {
                                    "moniker": "InfStones",
                                    "identity": "39A41C2FDE0AD040",
                                    "website": "https://infstones.io",
                                    "security_contact": "",
                                    "details": "Fueling Blockchain Beyond Infinity"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.010000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1vfawvtzvmcjkpqzhezvhk9q5tv5t7x8smz95uu",
                                "consensus_pubkey": "kavavalconspub1zcjduepq2yhf7ntk9dkkqpg07ead9nzjmldf2ar6tx6uz8jeaqjedrj9keqsq0hn02",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23079000000",
                                "delegator_shares": "23079000000.000000000000000000",
                                "description": {
                                    "moniker": "Newroad Network",
                                    "identity": "91FAC10CA3718A17",
                                    "website": "https://newroad.network/kava/",
                                    "security_contact": "",
                                    "details": "We provide a professional delegation service for multiple Proof of Stake networks. We use a secure and redundant setup. Visit our website for more information."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.350000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2020-02-03T04:21:51.013448233Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1vw35vclatlrcmzuaxf2fleyuk38xa7xf4vdaq2",
                                "consensus_pubkey": "kavavalconspub1zcjduepqnrpursauds6wjnqht0979kjr0pvk5jp4c0cv46feazqv4qappqsqn5rd22",
                                "jailed": false,
                                "status": 2,
                                "tokens": "875050100000",
                                "delegator_shares": "875050100000.000000000000000000",
                                "description": {
                                    "moniker": "hashquark",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "1.000000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1vw4t8ge2ephu0wuhcmclcw04ag2vzj9rpdme2x",
                                "consensus_pubkey": "kavavalconspub1zcjduepq0fjxw9wq8swcrk26j2vk8qm90d0vn94c2c7cs0ru5rgy2zlyv5eq8q5kvc",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20153000000",
                                "delegator_shares": "20153000000.000000000000000000",
                                "description": {
                                    "moniker": "UbikCapital",
                                    "identity": "8265DEAF50B61DF7",
                                    "website": "https://ubik.capital",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.500000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2019-12-25T21:46:08.338409716Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1vuylvflgy75d8zr07ta8x0gcqynvkvw70et853",
                                "consensus_pubkey": "kavavalconspub1zcjduepqvk6tye9w7399g7xys6j5ycw5rly8wz92mp0pevgmrv30qc4nva5sc2u25u",
                                "jailed": false,
                                "status": 2,
                                "tokens": "1874980100000",
                                "delegator_shares": "1874980100000.000000000000000000",
                                "description": {
                                    "moniker": "Lemniscap",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.250000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1dwae0ny4uuvacucakm9v8r8mxhw82zack4cn7y",
                                "consensus_pubkey": "kavavalconspub1zcjduepq3rfkl40umrfy0wq9u9m5mc5zv30x4lmj08vup2u5d90t4fu3qe7qvq5scl",
                                "jailed": false,
                                "status": 2,
                                "tokens": "3317768400000",
                                "delegator_shares": "3317768400000.000000000000000000",
                                "description": {
                                    "moniker": "OKExPool",
                                    "identity": "",
                                    "website": "www.okex.com/pool",
                                    "security_contact": "",
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
                                    "update_time": "2019-12-20T02:25:09.919481267Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1dj63l3z0smqa8au5yek37nmj0xd60zs9dwmdh0",
                                "consensus_pubkey": "kavavalconspub1zcjduepqq0hj2jx5x7p65t56c7s750e85yp0u6xvydk6eh5m45xq755rjqlsthzygr",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20100000782",
                                "delegator_shares": "20100000782.000000000000000000",
                                "description": {
                                    "moniker": "Yn",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-15T14:04:52.626128649Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1dntlhrw3jrej6ssdp64yfmkkz08ykyx4n7hphh",
                                "consensus_pubkey": "kavavalconspub1zcjduepqxd9wm06lw0p6nrysw2ecmusgrmsq6438ht4wh7gx64a0m6vnu4qs7jhyjk",
                                "jailed": false,
                                "status": 2,
                                "tokens": "4001049500000",
                                "delegator_shares": "4001049500000.000000000000000000",
                                "description": {
                                    "moniker": "Staked",
                                    "identity": "E7BFA6515FB02B3B",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-12-03T16:23:23.639110694Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1dede4flaq24j2g9u8f83vkqrqxe6cwzrxt5zsu",
                                "consensus_pubkey": "kavavalconspub1zcjduepqstxrft39yrg5hgalcnjqurms5tj2yqpw9y85axv9gsxyld62cneqg6k7nk",
                                "jailed": false,
                                "status": 2,
                                "tokens": "104328140934",
                                "delegator_shares": "104328140934.000000000000000000",
                                "description": {
                                    "moniker": "POS Bakerz",
                                    "identity": "3AFAE7268F4DFD10",
                                    "website": "https://posbakerz.com/",
                                    "security_contact": "",
                                    "details": "Secure, Reliable and Efficient Staking-as-a-Service"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.098500000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-03-20T12:29:19.763049675Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1davtd9w5yadvlg5x7aw0kpkyhckk5m5ecrmnyl",
                                "consensus_pubkey": "kavavalconspub1zcjduepq9x39e7g3m2hv8v6nxcx2lmy6fzaeqmvdz4uqtpm8qaxtr6527yrsh3fkgl",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23000000000",
                                "delegator_shares": "23000000000.000000000000000000",
                                "description": {
                                    "moniker": "Cyberili",
                                    "identity": "3554F49719B1BF6F",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-15T14:55:31.701406451Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1wtcn3ylrsp4qp4urlhkf78kkt8f2ch7p0f0p98",
                                "consensus_pubkey": "kavavalconspub1zcjduepqjffmrcy39vaecs9rjaju3hzq97nggr2l0zg96k2u3c54mfx20hks43wvdy",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23124628440",
                                "delegator_shares": "23124628440.000000000000000000",
                                "description": {
                                    "moniker": "007",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.200000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2020-03-20T19:41:53.352622226Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1wu8m65vqazssv2rh8rthv532hzggfr3h9azwz9",
                                "consensus_pubkey": "kavavalconspub1zcjduepqqg29usle4d3hgt4eadn3epppa2h6dsga9mj6sxvyj5prchufgr5qmrneka",
                                "jailed": false,
                                "status": 2,
                                "tokens": "5005157202698",
                                "delegator_shares": "5005157202698.000000000000000000",
                                "description": {
                                    "moniker": "Binance Staking",
                                    "identity": "",
                                    "website": "https://binance.com",
                                    "security_contact": "",
                                    "details": "Exchange the world!"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "1.000000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1spkjjtwks5zt0dqexj8rv8ljwmwxe8ufkraukq",
                                "consensus_pubkey": "kavavalconspub1zcjduepq9gppak7hsjrwvyxz2y9hy830gzsx2r4nax43vh2lprj0wj3rvghsyu6xev",
                                "jailed": false,
                                "status": 2,
                                "tokens": "22950000000",
                                "delegator_shares": "22950000000.000000000000000000",
                                "description": {
                                    "moniker": "Jupiter.IAM",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-17T11:35:54.989770823Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1srhded4xw0krcwlvddtcyycuh70fp5ry9yvp86",
                                "consensus_pubkey": "kavavalconspub1zcjduepq64ugztwcunjrpema5pqwherkfznms8ykc4jx7syc6tm5ztv58keqxlc6fv",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2402020264361",
                                "delegator_shares": "2402020264361.000000000000000000",
                                "description": {
                                    "moniker": "TRGCapital",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-12-12T17:38:19.894601876Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1s8akp0nq7z7vuf5g9agdswcwq7vm9uwudk0han",
                                "consensus_pubkey": "kavavalconspub1zcjduepqpgz29c6efzyj96y0kkjrxns67nsm5t3rae0ra2nmrx4fm5j33nkqn066rz",
                                "jailed": false,
                                "status": 2,
                                "tokens": "22950000000",
                                "delegator_shares": "22950000000.000000000000000000",
                                "description": {
                                    "moniker": "IvoryTower",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-17T02:12:20.831453541Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper13fxkk4730cqglgdv7w0mdelyx07myyq76uf9f3",
                                "consensus_pubkey": "kavavalconspub1zcjduepquwqhhvst3t6pg7rzutmvr3dma46ux6nggfxrju4f6u2dza9n3l7s54jgu7",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23213610000",
                                "delegator_shares": "23215931593.159315931593159315",
                                "description": {
                                    "moniker": "Finantra.com",
                                    "identity": "",
                                    "website": "http://finantra.com",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "1122467",
                                "unbonding_time": "2020-03-04T18:59:03.214161979Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.500000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2020-03-14T19:38:24.407313925Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper13vstf6ecmfe4p0gaufumk569sdawhtrf8gu56h",
                                "consensus_pubkey": "kavavalconspub1zcjduepq2ye3234dt2jes9g49mf5sdy7f7xfshekg5xyjm7vq7m07s9uq82qger9yk",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23081092090",
                                "delegator_shares": "23081092090.000000000000000000",
                                "description": {
                                    "moniker": "QuantaFrontier.com",
                                    "identity": "",
                                    "website": "https://QuantaFrontier.com",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.120000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2020-03-10T19:08:38.751434231Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper13dfu6et0m8zm4hachudvn62w0c2zvcmrqn8cs3",
                                "consensus_pubkey": "kavavalconspub1zcjduepqcctrta82mcm79j0e6ttgguekmljvud0y4rlggrm9uw9t49qemujsd69h59",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23144338847",
                                "delegator_shares": "23144338847.000000000000000000",
                                "description": {
                                    "moniker": "Sunshine",
                                    "identity": "24DB71E8B5076192",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.300000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2020-04-05T16:35:57.63239062Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1jyuv7z9at27elvmnmzh2v39dc06r9kjcy59xkr",
                                "consensus_pubkey": "kavavalconspub1zcjduepqcznteuyudn9zmfzyu6cksgh9qv3e9sc9fkyu32djy9d29ertyjrqfwcnvj",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2799950000000",
                                "delegator_shares": "2799950000000.000000000000000000",
                                "description": {
                                    "moniker": "DAF",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.050000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1j26c4k2jj9tv95whdhva3e8v2fcm4s3dsgstd2",
                                "consensus_pubkey": "kavavalconspub1zcjduepq3ra0zhkgeyqzksspztcx475dtlra5lckdpzwht29mk9fuepnkjqsw2rly7",
                                "jailed": false,
                                "status": 2,
                                "tokens": "3269187409856",
                                "delegator_shares": "3269187409856.000000000000000000",
                                "description": {
                                    "moniker": "DokiaCapital",
                                    "identity": "",
                                    "website": "https://staking.dokia.cloud",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.120000000000000000",
                                    "max_rate": "0.120000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1jhraz4ftxl2pd37knmeua7wjxghmlskwf7x8pf",
                                "consensus_pubkey": "kavavalconspub1zcjduepq6rxswfd9puq5t7jlp52cetv5s7cqm4l4n5tx4gqsclg9ngswe45qwhz2dm",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23443000000",
                                "delegator_shares": "23443000000.000000000000000000",
                                "description": {
                                    "moniker": "mariposa2",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-12-26T13:44:37.328086675Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1jlzx4js09d9zzyuhuz8sfdweklrapzacuhsxq5",
                                "consensus_pubkey": "kavavalconspub1zcjduepqg94y638yaj7es28zdktmar4jn9wu0v2l4h8mr40gerr2rek8q3rq9d3gh6",
                                "jailed": false,
                                "status": 2,
                                "tokens": "3099479960136",
                                "delegator_shares": "3099479960136.000000000000000000",
                                "description": {
                                    "moniker": "shiprekt",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.000000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1jl42l225565y3hm9dm4my33hjgdzleucqryhlx",
                                "consensus_pubkey": "kavavalconspub1zcjduepqtlmej0vprfmun69nsnmuv56j3dctzw90tmk0f805lqxcqn9sr77s4nzkx0",
                                "jailed": false,
                                "status": 2,
                                "tokens": "34047476004",
                                "delegator_shares": "34047476004.000000000000000000",
                                "description": {
                                    "moniker": "melea-◮👁◭",
                                    "identity": "4BE49EABAA41B8BF",
                                    "website": "https://meleatrust.com/kava/",
                                    "security_contact": "meleacrypto@gmail.com",
                                    "details": "Validator service secure and trusted, awarded in Game Of Steaks"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.000000000000000000",
                                    "max_rate": "0.250000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2019-12-27T14:23:19.174760869Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1nxgg4grsc0fwh893mks62d3x3r6uazgpj3m3cr",
                                "consensus_pubkey": "kavavalconspub1zcjduepqkk7jeymtt9rk3jlhxtxu95akmxfusj44mqq2wuq8c2hckulpga0q4yrauv",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23085863116",
                                "delegator_shares": "23085863116.000000000000000000",
                                "description": {
                                    "moniker": "Easy 2 Stake",
                                    "identity": "2C877AC873132C91",
                                    "website": "www.easy2stake.com",
                                    "security_contact": "",
                                    "details": "Easy.Stake.Trust. As easy and as simple as you would click next. Complete transparency and trust with a secure and stable validator."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:00:10.631092707Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1njvaku4qg9pxmx9jgjks36xrxfd6fyqs3tgs4d",
                                "consensus_pubkey": "kavavalconspub1zcjduepqsxtm3s99l58x9e3sqfxx4z88v6t4pyhku6fnu60drwmq8zm4jwjqx257u9",
                                "jailed": false,
                                "status": 2,
                                "tokens": "1000002000000",
                                "delegator_shares": "1000002000000.000000000000000000",
                                "description": {
                                    "moniker": "BitMax Staking",
                                    "identity": "",
                                    "website": "https://bitmax.io/",
                                    "security_contact": "",
                                    "details": "Trading Platform Industry Leader driven by Product Innovation"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-01-15T09:09:05.545733701Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1nnwwu4km0alut2q8vhg7zjt45wyehddpwlfrmj",
                                "consensus_pubkey": "kavavalconspub1zcjduepqey9flxqkdw5ewlskac3wltu4yl93asdamt66prm393zc8g93dcuszezutf",
                                "jailed": false,
                                "status": 2,
                                "tokens": "5681499001",
                                "delegator_shares": "5682067150.013579192799429688",
                                "description": {
                                    "moniker": "Wetez",
                                    "identity": "26FA2B24F46A98EF",
                                    "website": "https://www.wetez.io",
                                    "security_contact": "",
                                    "details": "Wetez is the most professional team in the POS ( Proof of Stake) field.Wetez是POS领域最专业的团队，为POS带来的权益做更多赋能。"
                                },
                                "unbonding_height": "448243",
                                "unbonding_time": "2020-01-11T04:03:40.112446434Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.010000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2019-11-17T01:33:27.008856884Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper15kwwzz908wl0qv4w66a5ee70kytpmfc9khvah6",
                                "consensus_pubkey": "kavavalconspub1zcjduepq8ylkkt83ktmljrmmk7evkj600j3p9xjnm4kf44z33hyr3cum448sqw2a6j",
                                "jailed": false,
                                "status": 2,
                                "tokens": "6118996081",
                                "delegator_shares": "6118996081.000000000000000000",
                                "description": {
                                    "moniker": "Staking4All",
                                    "identity": "",
                                    "website": "https://www.staking4all.org",
                                    "security_contact": "",
                                    "details": "Validator for Proof of Stake blockchains. Delegate to us for a easy staking experience"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.300000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-03-25T15:53:47.321864774Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper15urq2dtp9qce4fyc85m6upwm9xul3049dcs7da",
                                "consensus_pubkey": "kavavalconspub1zcjduepqsglrgfudcqw4la7e54lfdu2va9z5gd4r8z6wrwe8m5rtz3d2vtsq55njyr",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2475995760455",
                                "delegator_shares": "2475995760455.000000000000000000",
                                "description": {
                                    "moniker": "Chorus One",
                                    "identity": "00B79D689B7DC1CE",
                                    "website": "https://chorus.one",
                                    "security_contact": "security@chorus.one",
                                    "details": "Secure Kava and shape its future by delegating to Chorus One, a highly secure and stable validator. By delegating, you agree to the terms of service."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.075000000000000000",
                                    "max_rate": "0.300000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper14fkp35j5nkvtztmxmsxh88jks6p3w8u7p76zs9",
                                "consensus_pubkey": "kavavalconspub1zcjduepqlgrd20pqzw4dmkdkg3nd65mra077pjne8l0d43wsvw93y9dw450smyh79n",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23654500000",
                                "delegator_shares": "23654500000.000000000000000000",
                                "description": {
                                    "moniker": "Bit Cat🐱",
                                    "identity": "FAB46CEEAEAB9FA1",
                                    "website": "https://www.bitcat365.com",
                                    "security_contact": "",
                                    "details": "Secure and stable KAVA validator service from China team"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:01:37.718892109Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper140g8fnnl46mlvfhygj3zvjqlku6x0fwu6lgey7",
                                "consensus_pubkey": "kavavalconspub1zcjduepqvtvkhh22hgfvp865tj4uwltv0hu7fs3vwmxwrl0n2mdpfuzj0p0qes2k9e",
                                "jailed": false,
                                "status": 2,
                                "tokens": "4051502234537",
                                "delegator_shares": "4051502234537.000000000000000000",
                                "description": {
                                    "moniker": "Cosmostation",
                                    "identity": "AE4C403A6E7AA1AC",
                                    "website": "https://www.cosmostation.io",
                                    "security_contact": "",
                                    "details": "Delegate your Kava and start earning your staking rewards."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.150000000000000000",
                                    "max_rate": "0.300000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-18T01:27:53.436530462Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper14kn0kk33szpwus9nh8n87fjel8djx0y02c7me3",
                                "consensus_pubkey": "kavavalconspub1zcjduepqk8fq7ssjps5r40j32lmhxrdzfg72k8dknc3xndljjtlghc9u9nrshrlf3r",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2883221903903",
                                "delegator_shares": "2883221903903.000000000000000000",
                                "description": {
                                    "moniker": "Forbole",
                                    "identity": "2861F5EE06627224",
                                    "website": "https://www.forbole.com",
                                    "security_contact": "info@forbole.com",
                                    "details": "Professional PoS validator based in Hong Kong"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.095000000000000000",
                                    "max_rate": "0.300000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1kgddca7qj96z0qcxr2c45z73cfl0c75p27tsg6",
                                "consensus_pubkey": "kavavalconspub1zcjduepqzah30dnsz855gnufkqlqnq50ksyc4jetprhz4ugg2rmuzhh3gvwsfxan38",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2156193000000",
                                "delegator_shares": "2156193000000.000000000000000000",
                                "description": {
                                    "moniker": "ChainLayer",
                                    "identity": "AD3CDBC91802F94A",
                                    "website": "https://www.chainlayer.io",
                                    "security_contact": "",
                                    "details": "Secure and reliable validator. TG: https://t.me/chainlayer"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.100000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1kwj4l5putuymgxw9kx8emh3e5dpaca0hnf3zdy",
                                "consensus_pubkey": "kavavalconspub1zcjduepqmqwfzwnmjdnkmt9k75k3q7rez3lr08m9fqev06rh6m2sdke8kmusjhywm9",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23173000000",
                                "delegator_shares": "23173000000.000000000000000000",
                                "description": {
                                    "moniker": "InChainWorks",
                                    "identity": "BE448F9ECAB40ABE",
                                    "website": "https://inchain.works",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.200000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-12-06T10:37:03.352901001Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1k4kxxfkhhwvyzxxxgksefkzpxahp9wfc572esl",
                                "consensus_pubkey": "kavavalconspub1zcjduepqt0k39t3zmz460are7rf9n97u20qhfmukp8ujeqhe0jg2q7ye99cqprr8nr",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23215000000",
                                "delegator_shares": "23215000000.000000000000000000",
                                "description": {
                                    "moniker": "Mr.K",
                                    "identity": "74D3AF53635231D9",
                                    "website": "",
                                    "security_contact": "",
                                    "details": "Validator on Tendermint base project"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.500000000000000000",
                                    "max_rate": "0.500000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-29T03:49:52.013497324Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1kc6vzheht92jwf0gtzhjk6jjht67rxhal9z04v",
                                "consensus_pubkey": "kavavalconspub1zcjduepqt8065r783whzn3w47uuhm0t8rw2ke0zda6awa4j5lemlhw029mtqs79pcs",
                                "jailed": false,
                                "status": 2,
                                "tokens": "20028000000",
                                "delegator_shares": "20028000000.000000000000000000",
                                "description": {
                                    "moniker": "moonli.me",
                                    "identity": "662AEC27BD90D036",
                                    "website": "https://moonli.me",
                                    "security_contact": "",
                                    "details": "How to Delegete: https://go.moonli.me/kava"
                                },
                                "unbonding_height": "972523",
                                "unbonding_time": "2020-02-21T16:44:26.247476886Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.120000000000000000",
                                    "max_rate": "0.300000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:03:28.153247409Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1k760ypy9tzhp6l2rmg06sq4n74z0d3rejwwaa0",
                                "consensus_pubkey": "kavavalconspub1zcjduepqqnxh5k4mhggtdx2u9j9xgq0etpkxpevg333sm954efqema6mn5yqqyf0gw",
                                "jailed": false,
                                "status": 2,
                                "tokens": "25190876386",
                                "delegator_shares": "25190876386.000000000000000000",
                                "description": {
                                    "moniker": "novy",
                                    "identity": "5422BE13389B2B4D",
                                    "website": "https://novy.pw",
                                    "security_contact": "",
                                    "details": "Let's validate this network!"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-12-29T20:44:00.222304348Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1h9ulmhqv5e2373khk6s9n0wtrfc5qavre09fxl",
                                "consensus_pubkey": "kavavalconspub1zcjduepqjqsd0jkhftta7fuc9e45f3yat25ja9xknzrdqye6zk9kyx40wkaq50hcdv",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23182088934",
                                "delegator_shares": "23182088934.000000000000000000",
                                "description": {
                                    "moniker": "alexDcrypto",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.190000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-01-02T22:54:47.953329027Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1hezl6xwva28xt0hk204dllalagenmfsnuu50j6",
                                "consensus_pubkey": "kavavalconspub1zcjduepqwf05tj86hm83n55lha6zgeyf9jyyz053cwq957nhhv933v9z9sgqdcmt5s",
                                "jailed": false,
                                "status": 2,
                                "tokens": "169401100000",
                                "delegator_shares": "169401100000.000000000000000000",
                                "description": {
                                    "moniker": "01node.com",
                                    "identity": "22823CD59617B8E3",
                                    "website": "https://01node.com",
                                    "security_contact": "",
                                    "details": "01node Staking Services"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.150000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:00:18.218278987Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1c9ye54e3pzwm3e0zpdlel6pnavrj9qqvh0atdq",
                                "consensus_pubkey": "kavavalconspub1zcjduepqc8585309s34yq92calayzy4xc4hu8rrx4pkrymewdy5qyjkx5tjqjgp7tq",
                                "jailed": false,
                                "status": 2,
                                "tokens": "4116869613316",
                                "delegator_shares": "4116869613316.000000000000000000",
                                "description": {
                                    "moniker": "StakeWith.Us",
                                    "identity": "609F83752053AD57",
                                    "website": "https://stakewith.us",
                                    "security_contact": "node@stakewith.us",
                                    "details": "Secured Staking Made Easy. Put Your Crypto to Work - Hassle Free. Disclaimer: Delegators should understand that delegation comes with slashing risk. By delegating to StakeWithUs Pte Ltd, you acknowledge that StakeWithUs Pte Ltd is not liable for any losses on your investment."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.250000000000000000",
                                    "max_change_rate": "0.050000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1cj9cdx9mg95lhvpquym08ncgzpjvhmnwdvm5kc",
                                "consensus_pubkey": "kavavalconspub1zcjduepq3w00ypwl5652c8unvwhn4urpkpcpztkusgqsdztdsq8fhycw224qwlcd33",
                                "jailed": false,
                                "status": 2,
                                "tokens": "868075391",
                                "delegator_shares": "868075391.000000000000000000",
                                "description": {
                                    "moniker": "Sifu Ventures",
                                    "identity": "",
                                    "website": "http://sifu.ventures",
                                    "security_contact": "",
                                    "details": "dPos Tokens Investment Fund"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.050000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-03-24T08:49:40.805033852Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1ceun2qqw65qce5la33j8zv8ltyyaqqfctl35n4",
                                "consensus_pubkey": "kavavalconspub1zcjduepqsqtuf26ph4szcxutxscz5eq0g9tnqvedhdf203cha77m8xlnuy0sy3nky9",
                                "jailed": false,
                                "status": 2,
                                "tokens": "3266752395198",
                                "delegator_shares": "3266752395198.000000000000000000",
                                "description": {
                                    "moniker": "sikka",
                                    "identity": "",
                                    "website": "https://sikka.tech",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.010000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "1.000000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper168g9nn9vnamhsnjkm7uqqee9f3v07flgwwddf9",
                                "consensus_pubkey": "kavavalconspub1zcjduepq29lrp34q3dveyt8r2mnc3kd82y5a5dsq39laj6yhkwans0ay2u9smmru3r",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23164473513",
                                "delegator_shares": "23164473513.000000000000000000",
                                "description": {
                                    "moniker": "sebytza05",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.140000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:00:18.218278987Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1645czvg787l0jr6mawhs9dm3mljnggj003yed9",
                                "consensus_pubkey": "kavavalconspub1zcjduepqws4aku5x2z0806kpyd750yzdpwrergpj82akyj0zz4ft6t38c4esdztwm8",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23138575200",
                                "delegator_shares": "23143203500.319736666961944921",
                                "description": {
                                    "moniker": "Mike Pocket",
                                    "identity": "C57B29418AE33CC0",
                                    "website": "",
                                    "security_contact": "",
                                    "details": "GO KAVA GO"
                                },
                                "unbonding_height": "1426431",
                                "unbonding_time": "2020-03-29T04:47:39.676931865Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.200000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-03-08T11:13:06.439174757Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper16lnfpgn6llvn4fstg5nfrljj6aaxyee9z59jqd",
                                "consensus_pubkey": "kavavalconspub1zcjduepquwh8pq39pwddp3etxk82d3u57kdfnvftgxamx4kcqduk8q02c3xqsyfaaa",
                                "jailed": false,
                                "status": 2,
                                "tokens": "2833230807310",
                                "delegator_shares": "2833230807310.000000000000000000",
                                "description": {
                                    "moniker": "pylonvalidator",
                                    "identity": "0979483D4F669CFF",
                                    "website": "https://pylonvalidator.com",
                                    "security_contact": "security@pylonvalidator.com",
                                    "details": "'It doesn't matter whether the cat is black or white, as long as it catches mice.' --Deng Xiaoping"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "1.000000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1m9y0c7j7wxyu3nqtmeevyfzzpga8jhdyqw42wx",
                                "consensus_pubkey": "kavavalconspub1zcjduepqqpq85gztxdvheh48pugwp624mm6ey6u3464pj80c6s2epa8m9uwq4xr44m",
                                "jailed": false,
                                "status": 2,
                                "tokens": "1255000300000",
                                "delegator_shares": "1255000300000.000000000000000000",
                                "description": {
                                    "moniker": "B-Harvest",
                                    "identity": "8957C5091FBF4192",
                                    "website": "https://bharvest.io",
                                    "security_contact": "contact@bharvest.io",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.300000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-20T04:34:50.0221158Z"
                                },
                                "min_self_delegation": "1000"
                                },
                                {
                                "operator_address": "kavavaloper1mu78xhlr705mzgwqykcafp4xy3kgatvwzrww8z",
                                "consensus_pubkey": "kavavalconspub1zcjduepqhux83t9q55l0em60vq5a0p22zpf6kcqk3qs7cyzcc5huhycp6qnsutwsz3",
                                "jailed": false,
                                "status": 2,
                                "tokens": "21000000000",
                                "delegator_shares": "21000000000.000000000000000000",
                                "description": {
                                    "moniker": "joesatri",
                                    "identity": "649DD29EE41766EB",
                                    "website": "",
                                    "security_contact": "",
                                    "details": "Game of Stakes Never Been Jailed Crew | Kava Founder Badge"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.110000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-15T15:01:47.310662882Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1uvz0vus7ktxt47cermscwe3k9gs7h9ag05sh6g",
                                "consensus_pubkey": "kavavalconspub1zcjduepq06xpflh9qkvgax8avl4wen49z48svsv4vplxk2jp2p4qx2htckzsfe2mwq",
                                "jailed": false,
                                "status": 2,
                                "tokens": "32995000000",
                                "delegator_shares": "32995000000.000000000000000000",
                                "description": {
                                    "moniker": "Genesis Lab",
                                    "identity": "C1A123F2723041F0",
                                    "website": "https://genesislab.net",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-15T14:15:12.174174763Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1udg4pal8c9gffv4l7zhvza027z0y345gd6esa6",
                                "consensus_pubkey": "kavavalconspub1zcjduepqpdjehvuwljjjp5se8t9xselu0hlneh2np3vrxwd0f4v6csdahclqktmkdp",
                                "jailed": false,
                                "status": 2,
                                "tokens": "1739934640784",
                                "delegator_shares": "1739934640784.000000000000000000",
                                "description": {
                                    "moniker": "ATEAM",
                                    "identity": "0CB9A4E7643FF992",
                                    "website": "",
                                    "security_contact": "",
                                    "details": "Kava_Founding member | Cosmos_GOS: Never Jailed \u0026 Top-Tier | Terra_Validator Drill: Top 3 | IOV_Validator candidates score for mainnet: Top 5 | Lino_Founding Validator Prize Winners: Top 7 | IRISnet_Nyancat: Perfect score"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.050000000000000000",
                                    "max_rate": "0.300000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-19T01:30:06.727558554Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1u0kfndes0pf8dstaunsgumv7scsnmy09p3ln9r",
                                "consensus_pubkey": "kavavalconspub1zcjduepqxzxd363ej2v0kljvfvy5ff4dka90gnkkfemej9p6rkqv3zac9qhsy8p567",
                                "jailed": false,
                                "status": 2,
                                "tokens": "4800705821800",
                                "delegator_shares": "4800705821800.000000000000000000",
                                "description": {
                                    "moniker": "SNZPool",
                                    "identity": "FF2019D4CF1F3185",
                                    "website": "https://snzholding.com",
                                    "security_contact": "",
                                    "details": "SNZ Pool is a professional \u0026 reliable POS validator for a dozen of projects like Cosmos, IRISnet, EOS, ONT, Loom, etc."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.100000000000000000"
                                    },
                                    "update_time": "2019-11-15T14:18:22.510513847Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1u3jf6c2f85kmjldxsncnhsdp44nac7v5j7vzpc",
                                "consensus_pubkey": "kavavalconspub1zcjduepqvd3xz6346n447pvl3s3s73ee3zwwn54f4ejwjfue3ppvpjgdfxwqu95mqu",
                                "jailed": false,
                                "status": 2,
                                "tokens": "3314184935364",
                                "delegator_shares": "3314184935364.000000000000000000",
                                "description": {
                                    "moniker": "HuobiPool",
                                    "identity": "23536C5BDE3EB949",
                                    "website": "https://www.huobipool.com/",
                                    "security_contact": "",
                                    "details": "Huobi Pool is a sub-brand of Huobi Group, which is an important part of the global ecological strategy of Huobi.Huobi Pool has become one of the largest POS communities in the Asia- Pacific region, the leading POW pool and nodes of a number of public chains."
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-12-12T15:50:40.940881395Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1u3hqe2m7vm59l30tyaqd3zurz864dlsg7nq83f",
                                "consensus_pubkey": "kavavalconspub1zcjduepqrtnnezp4g67f7hh8fs7gtx23pmdfjwu37uqguq4hsrglqkfmyensl4xvdt",
                                "jailed": false,
                                "status": 2,
                                "tokens": "4039881397431",
                                "delegator_shares": "4040285425961.491814603996529363",
                                "description": {
                                    "moniker": "Delchain",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "388052",
                                "unbonding_time": "2020-01-06T09:05:10.544078799Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.050000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-05T14:00:00Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1uu9zguynwjyyyq57jnm3j2q47d5c362wafgsmf",
                                "consensus_pubkey": "kavavalconspub1zcjduepqn5uazvr3ew6ktstgshs9pfmh3vh9f9m36e76era3hxffxk8pr5asux5v9a",
                                "jailed": false,
                                "status": 2,
                                "tokens": "1000000",
                                "delegator_shares": "1000000.000000000000000000",
                                "description": {
                                    "moniker": "sentry_lab",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "1.000000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-12-13T08:30:36.49152285Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1u7vsaanwt4e5mdzmxuqurmccxjka3h0ns2n3f5",
                                "consensus_pubkey": "kavavalconspub1zcjduepq7uy7ln7afmgrfutfpzhyy0llcndntx9q55nl5puqxcq5dqdzrnqswahrha",
                                "jailed": false,
                                "status": 2,
                                "tokens": "1000009950010",
                                "delegator_shares": "1000009950010.000000000000000000",
                                "description": {
                                    "moniker": "BiKi Pool",
                                    "identity": "",
                                    "website": "https://www.biki.com",
                                    "security_contact": "",
                                    "details": "Popular Token's Choice"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.100000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-03-04T12:26:32.258743177Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1atelyappg8x4ndgtj6h2jng532a5hvuke2cajr",
                                "consensus_pubkey": "kavavalconspub1zcjduepqrwah2t8gezgv79stgnzju9ank4rpe85juf6xf4zzfhwgfawc5jpse287y0",
                                "jailed": false,
                                "status": 2,
                                "tokens": "567000000",
                                "delegator_shares": "567000000.000000000000000000",
                                "description": {
                                    "moniker": "Nosce",
                                    "identity": "",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.007000000000000000",
                                    "max_rate": "0.007700000000000000",
                                    "max_change_rate": "0.000100000000000000"
                                    },
                                    "update_time": "2020-02-19T04:09:38.903736876Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper1aj9r9ll8m72xr5pmxr9fmum88k864tqg39nkfu",
                                "consensus_pubkey": "kavavalconspub1zcjduepq9nq9gdnvtl5as9sgl92exde7hju797hl20c9htje2chgwnlkwluq8hmtu0",
                                "jailed": false,
                                "status": 2,
                                "tokens": "21000000000",
                                "delegator_shares": "21000000000.000000000000000000",
                                "description": {
                                    "moniker": "kytzu",
                                    "identity": "909A480D5643CCC5",
                                    "website": "https://www.linkedin.com/in/calinchitu",
                                    "security_contact": "",
                                    "details": "Blockchain consultant and developer"
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.190000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2019-11-15T16:04:53.477387115Z"
                                },
                                "min_self_delegation": "1000000"
                                },
                                {
                                "operator_address": "kavavaloper1ajwfalplxnhkwhsycfax36yyyxpxz3450s800x",
                                "consensus_pubkey": "kavavalconspub1zcjduepq33ekkklqez4tdd37mrvptrkawejnj0p5caa0a6c7efqs8n7eg37qa2vega",
                                "jailed": false,
                                "status": 2,
                                "tokens": "42770134655",
                                "delegator_shares": "42770134655.000000000000000000",
                                "description": {
                                    "moniker": "Ryabina.io - Staking provider in awesome networks",
                                    "identity": "6B1C8FDD84CF4ACD",
                                    "website": "https://ryabina.io",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.010000000000000000",
                                    "max_rate": "0.200000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-02-24T18:08:47.865265697Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper17wcggpjx007uc09s8y4hwrj8f228mlwez945ey",
                                "consensus_pubkey": "kavavalconspub1zcjduepq5jtcd7kxpft40l65gm5p2ecg8rcxp59nml424u88lxkpxu6yl9cq6hxvzd",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23092086042",
                                "delegator_shares": "23092086042.000000000000000000",
                                "description": {
                                    "moniker": "Inotel",
                                    "identity": "975D494265B1AC25",
                                    "website": "",
                                    "security_contact": "",
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
                                    "update_time": "2019-11-15T14:04:03.245058425Z"
                                },
                                "min_self_delegation": "1"
                                },
                                {
                                "operator_address": "kavavaloper17498ffqdj49zca4jm7mdf3eevq7uhcsgjvm0uk",
                                "consensus_pubkey": "kavavalconspub1zcjduepq0mk39ccslman3ame4s3dn5exml8upqearp089sa5tae6pchdyrcsmvwmhn",
                                "jailed": false,
                                "status": 2,
                                "tokens": "23699000000",
                                "delegator_shares": "23699000000.000000000000000000",
                                "description": {
                                    "moniker": "kaval",
                                    "identity": "5281C9C7BE5C2646",
                                    "website": "",
                                    "security_contact": "",
                                    "details": ""
                                },
                                "unbonding_height": "0",
                                "unbonding_time": "1970-01-01T00:00:00Z",
                                "commission": {
                                    "commission_rates": {
                                    "rate": "0.550000000000000000",
                                    "max_rate": "1.000000000000000000",
                                    "max_change_rate": "0.010000000000000000"
                                    },
                                    "update_time": "2020-03-07T10:24:27.851934685Z"
                                },
                                "min_self_delegation": "1"
                                }
                            ]}
                        `);
                }
        }

        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};

