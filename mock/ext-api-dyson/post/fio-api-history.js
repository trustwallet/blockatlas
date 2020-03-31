/// FIO RPC API Mock, history API
/// curl -H "Content-Type: application/json" -d '{"account_name": "ezsmbcy2opod"}' https://fio.eosphere.io/v1/history/get_actions
/// curl -H "Content-Type: application/json" -d '{"account_name": "ezsmbcy2opod"}' http://localhost:3000/fio-api/v1/history/get_actions
/// curl "http://localhost:8420/v1/fio/FIO7Q3XfQ2ocGP1zYst6Sfx5qrsiZ865Cu8So2atrub9JN94so7gt"

module.exports = {
    path: '/fio-api/v1/history/:action',
    template: function(params, query, body) {
        console.log(params);
        console.log(body);
        if (params.action === 'get_actions') {
            if (body.account_name === 'ezsmbcy2opod') {
                return JSON.parse(`
                    {
                        "actions": [
                            {
                                "global_action_seq": 507874,
                                "account_action_seq": 0,
                                "block_num": 358689,
                                "block_time": "2020-03-27T01:54:24.500",
                                "action_trace": {
                                    "receipt": {
                                        "receiver": "ezsmbcy2opod",
                                        "response": "{'status': 'OK','fee_collected':2000000000}",
                                        "act_digest": "6ef1d997fd449b6fefb8df19401c71425d366a24aab4613ca2059478fef7915f",
                                        "global_sequence": 507874,
                                        "recv_sequence": 1,
                                        "auth_sequence": [
                                            [
                                                "f5axfpgffiqz",
                                                125
                                            ]
                                        ],
                                        "code_sequence": 1,
                                        "abi_sequence": 1
                                    },
                                    "receiver": "ezsmbcy2opod",
                                    "act": {
                                        "account": "fio.token",
                                        "name": "trnsfiopubky",
                                        "authorization": [
                                            {
                                                "actor": "f5axfpgffiqz",
                                                "permission": "active"
                                            }
                                        ],
                                        "data": {
                                            "payee_public_key": "FIO7Q3XfQ2ocGP1zYst6Sfx5qrsiZ865Cu8So2atrub9JN94so7gt",
                                            "amount": 700000000000,
                                            "max_fee": 2000000000,
                                            "actor": "f5axfpgffiqz",
                                            "tpid": ""
                                        },
                                        "hex_data": "3546494f375133586651326f634750317a5973743653667835717273695a383635437538536f326174727562394a4e3934736f376774005840fba20000000094357700000000f0ad5b8bd5d54d5900"
                                    },
                                    "context_free": false,
                                    "elapsed": 6,
                                    "console": "",
                                    "trx_id": "2e0a7dc3640768e1d644cee871734dd2efa23e65a54c438c1ba03801d7386fb7",
                                    "block_num": 358689,
                                    "block_time": "2020-03-27T01:54:24.500",
                                    "producer_block_id": "00057921b2a5f55fba29f7f08cd49dd095a18f0de2e4e06ba0833ca771ad699b",
                                    "account_ram_deltas": [],
                                    "except": null,
                                    "error_code": null,
                                    "action_ordinal": 6,
                                    "creator_action_ordinal": 1,
                                    "closest_unnotified_ancestor_action_ordinal": 1
                                }
                            },
                            {
                                "global_action_seq": 508968,
                                "account_action_seq": 1,
                                "block_num": 359762,
                                "block_time": "2020-03-27T02:03:21.000",
                                "action_trace": {
                                    "receipt": {
                                        "receiver": "fio.address",
                                        "response": "{'status': 'OK','expiration':'2021-03-27T02:03:21','fee_collected':40000000000}",
                                        "act_digest": "e054494384148e58122a10c3398d9e01a09a79cd6c22da0203653906afd79b3b",
                                        "global_sequence": 508968,
                                        "recv_sequence": 34025,
                                        "auth_sequence": [
                                            [
                                                "ezsmbcy2opod",
                                                1
                                            ]
                                        ],
                                        "code_sequence": 1,
                                        "abi_sequence": 1
                                    },
                                    "receiver": "fio.address",
                                    "act": {
                                        "account": "fio.address",
                                        "name": "regaddress",
                                        "authorization": [
                                            {
                                                "actor": "ezsmbcy2opod",
                                                "permission": "active"
                                            }
                                        ],
                                        "data": {
                                            "fio_address": "bp@eosph",
                                            "owner_fio_public_key": "FIO7Q3XfQ2ocGP1zYst6Sfx5qrsiZ865Cu8So2atrub9JN94so7gt",
                                            "max_fee": 40000000000,
                                            "actor": "ezsmbcy2opod",
                                            "tpid": ""
                                        },
                                        "hex_data": "08627040656f7370683546494f375133586651326f634750317a5973743653667835717273695a383635437538536f326174727562394a4e3934736f37677400902f50090000009068a5c2a323f15700"
                                    },
                                    "context_free": false,
                                    "elapsed": 2834,
                                    "console": "",
                                    "trx_id": "b7dd60839ccce11cd175fce6816da04bbce9b70825661005d2ea5d3572408c04",
                                    "block_num": 359762,
                                    "block_time": "2020-03-27T02:03:21.000",
                                    "producer_block_id": "00057d5209ad0141445b3ee2eb1ba7265ab4c7093f7bc4c2142e2eac4c8e5d10",
                                    "account_ram_deltas": [
                                        {
                                            "account": "ezsmbcy2opod",
                                            "delta": 798
                                        }
                                    ],
                                    "except": null,
                                    "error_code": null,
                                    "action_ordinal": 1,
                                    "creator_action_ordinal": 0,
                                    "closest_unnotified_ancestor_action_ordinal": 0
                                }
                            },
                            {
                                "global_action_seq": 508970,
                                "account_action_seq": 2,
                                "block_num": 359762,
                                "block_time": "2020-03-27T02:03:21.000",
                                "action_trace": {
                                    "receipt": {
                                        "receiver": "ezsmbcy2opod",
                                        "response": "",
                                        "act_digest": "e6774ab921eabe540049667756efb5567b83b291b8a21c5bec81c8b5dc98c366",
                                        "global_sequence": 508970,
                                        "recv_sequence": 2,
                                        "auth_sequence": [
                                            [
                                                "eosio",
                                                430085
                                            ]
                                        ],
                                        "code_sequence": 1,
                                        "abi_sequence": 1
                                    },
                                    "receiver": "ezsmbcy2opod",
                                    "act": {
                                        "account": "fio.token",
                                        "name": "transfer",
                                        "authorization": [
                                            {
                                                "actor": "eosio",
                                                "permission": "active"
                                            }
                                        ],
                                        "data": {
                                            "from": "ezsmbcy2opod",
                                            "to": "fio.treasury",
                                            "quantity": "40.000000000 FIO",
                                            "memo": "FIO API fees. Thank you."
                                        },
                                        "hex_data": "9068a5c2a323f157e0afc646dd0ca85b00902f50090000000946494f000000001846494f2041504920666565732e205468616e6b20796f752e"
                                    },
                                    "context_free": false,
                                    "elapsed": 6,
                                    "console": "",
                                    "trx_id": "b7dd60839ccce11cd175fce6816da04bbce9b70825661005d2ea5d3572408c04",
                                    "block_num": 359762,
                                    "block_time": "2020-03-27T02:03:21.000",
                                    "producer_block_id": "00057d5209ad0141445b3ee2eb1ba7265ab4c7093f7bc4c2142e2eac4c8e5d10",
                                    "account_ram_deltas": [],
                                    "except": null,
                                    "error_code": null,
                                    "action_ordinal": 6,
                                    "creator_action_ordinal": 2,
                                    "closest_unnotified_ancestor_action_ordinal": 2
                                }
                            },
                            {
                                "global_action_seq": 510503,
                                "account_action_seq": 3,
                                "block_num": 361281,
                                "block_time": "2020-03-27T02:16:00.500",
                                "action_trace": {
                                    "receipt": {
                                        "receiver": "eosio",
                                        "response": "{'status': 'OK','fee_collected':200000000000}",
                                        "act_digest": "52b1631c84896baaebd4b81a80c2cfab4f9741e0323d2845d1b3df5683ad7c3c",
                                        "global_sequence": 510503,
                                        "recv_sequence": 404854,
                                        "auth_sequence": [
                                            [
                                                "ezsmbcy2opod",
                                                2
                                            ]
                                        ],
                                        "code_sequence": 2,
                                        "abi_sequence": 2
                                    },
                                    "receiver": "eosio",
                                    "act": {
                                        "account": "eosio",
                                        "name": "regproducer",
                                        "authorization": [
                                            {
                                                "actor": "ezsmbcy2opod",
                                                "permission": "active"
                                            }
                                        ],
                                        "data": {
                                            "fio_address": "bp@eosph",
                                            "fio_pub_key": "FIO5hZB8REVEirigba4N7TKa67MYm4HFwtiKz6GZJ2eRi5Paxwixz",
                                            "url": "https://www.eosph.io",
                                            "location": 10,
                                            "actor": "ezsmbcy2opod",
                                            "max_fee": 400000000000
                                        },
                                        "hex_data": "08627040656f7370683546494f35685a423852455645697269676261344e37544b6136374d596d3448467774694b7a36475a4a32655269355061787769787a1468747470733a2f2f7777772e656f7370682e696f0a009068a5c2a323f15700a0db215d000000"
                                    },
                                    "context_free": false,
                                    "elapsed": 2207,
                                    "console": "",
                                    "trx_id": "5cd8d902a675d608d4c82eca892c7973deaa6aefcc58ed9bcbd382ebdc7271c8",
                                    "block_num": 361281,
                                    "block_time": "2020-03-27T02:16:00.500",
                                    "producer_block_id": "00058341700331f78adea1eec39e7150b743f171c9fd9dfff8ea7310f5e6b515",
                                    "account_ram_deltas": [
                                        {
                                            "account": "ezsmbcy2opod",
                                            "delta": 635
                                        }
                                    ],
                                    "except": null,
                                    "error_code": null,
                                    "action_ordinal": 1,
                                    "creator_action_ordinal": 0,
                                    "closest_unnotified_ancestor_action_ordinal": 0
                                }
                            },
                            {
                                "global_action_seq": 510505,
                                "account_action_seq": 4,
                                "block_num": 361281,
                                "block_time": "2020-03-27T02:16:00.500",
                                "action_trace": {
                                    "receipt": {
                                        "receiver": "ezsmbcy2opod",
                                        "response": "",
                                        "act_digest": "a8dd1a5cfe8db5fb3ca38eed9e348acd73150d88211f93e0d37e332289fe76a3",
                                        "global_sequence": 510505,
                                        "recv_sequence": 3,
                                        "auth_sequence": [
                                            [
                                                "eosio",
                                                431614
                                            ]
                                        ],
                                        "code_sequence": 1,
                                        "abi_sequence": 1
                                    },
                                    "receiver": "ezsmbcy2opod",
                                    "act": {
                                        "account": "fio.token",
                                        "name": "transfer",
                                        "authorization": [
                                            {
                                                "actor": "eosio",
                                                "permission": "active"
                                            }
                                        ],
                                        "data": {
                                            "from": "ezsmbcy2opod",
                                            "to": "fio.treasury",
                                            "quantity": "200.000000000 FIO",
                                            "memo": "FIO API fees. Thank you."
                                        },
                                        "hex_data": "9068a5c2a323f157e0afc646dd0ca85b00d0ed902e0000000946494f000000001846494f2041504920666565732e205468616e6b20796f752e"
                                    },
                                    "context_free": false,
                                    "elapsed": 6,
                                    "console": "",
                                    "trx_id": "5cd8d902a675d608d4c82eca892c7973deaa6aefcc58ed9bcbd382ebdc7271c8",
                                    "block_num": 361281,
                                    "block_time": "2020-03-27T02:16:00.500",
                                    "producer_block_id": "00058341700331f78adea1eec39e7150b743f171c9fd9dfff8ea7310f5e6b515",
                                    "account_ram_deltas": [],
                                    "except": null,
                                    "error_code": null,
                                    "action_ordinal": 6,
                                    "creator_action_ordinal": 2,
                                    "closest_unnotified_ancestor_action_ordinal": 2
                                }
                            }
                        ],
                        "last_irreversible_block": 928886
                    }
                `);
            }
        }
        return {error: 'Not implemented'};
    }
};
