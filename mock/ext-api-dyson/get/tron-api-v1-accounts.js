/// Tron API Mock
/// See:
/// curl "http://localhost:3000/tron-api/v1/accounts/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB/transactions?token_id=&limit=25&order_by=block_timestamp,desc"
/// curl "http://localhost:3000/tron-api/v1/accounts/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB"
/// curl "http://{Tron rpc}/v1/accounts/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB/transactions?token_id=&limit=25&order_by=block_timestamp,desc"
/// curl "http://{Tron rpc}/v1/accounts/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB"
/// curl "http://localhost:8420/v1/tron/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB"
/// curl "http://localhost:8420/v2/tron/tokens/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB?Authorization=Bearer"

module.exports = {
    path: "/tron-api/v1/accounts/:address/:operation?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.operation === 'transactions') {
            if (params.address === 'TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB') {
                return JSON.parse(`
                    {
                        "success": true,
                        "meta": {
                            "at": 1585038748264,
                            "page_size": 2
                        },
                        "data": [
                            {
                                "block_timestamp": 1553870898000,
                                "raw_data": {
                                    "contract": [
                                        {
                                            "parameter": {
                                                "type_url": "type.googleapis.com/protocol.TransferAssetContract",
                                                "value": {
                                                    "amount": 10149740000,
                                                    "asset_name": "1002000",
                                                    "owner_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261",
                                                    "to_address": "4139fec4d95bb59f45a727f9234020adaf2cec9e20"
                                                }
                                            },
                                            "type": "TransferAssetContract"
                                        }
                                    ],
                                    "expiration": 1553870955000,
                                    "fee_limit": 0,
                                    "ref_block_bytes": "b560",
                                    "ref_block_hash": "2d9026ee979db6b6",
                                    "timestamp": 1553870897024
                                },
                                "raw_data_hex": "0a02b56022082d9026ee979db6b640f8c3b4cf9c2d5a77080212730a32747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e736665724173736574436f6e7472616374123d0a07313030323030301215410583a68a3bcd86c25ab1bee482bac04a216b02611a154139fec4d95bb59f45a727f9234020adaf2cec9e2020e0fbe2e7257080ffb0cf9c2d",
                                "ret": [
                                    {
                                        "code": "SUCESS",
                                        "contractRet": "SUCCESS",
                                        "fee": 0
                                    }
                                ],
                                "signature": [
                                    "10c57bda2346f4cc7e287a2b696fcd8dcbec273d84df8b8d41fcf49d13e02fdf1ca539d1b631feabd58e3f0a8c501c8d74595774fd149cfba16770eb8a4bd18100"
                                ],
                                "txID": "d9206168ee601935bb19de30a613f735972f1ab59178ba05d078ff1ae19c0602"
                            },
                            {
                                "block_timestamp": 1553864040000,
                                "raw_data": {
                                    "contract": [
                                        {
                                            "parameter": {
                                                "type_url": "type.googleapis.com/protocol.TransferContract",
                                                "value": {
                                                    "amount": 278720000,
                                                    "owner_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261",
                                                    "to_address": "4139fec4d95bb59f45a727f9234020adaf2cec9e20"
                                                }
                                            },
                                            "type": "TransferContract"
                                        }
                                    ],
                                    "expiration": 1553864097000,
                                    "fee_limit": 0,
                                    "ref_block_bytes": "ac7f",
                                    "ref_block_hash": "648c4cb65453b4de",
                                    "timestamp": 1553864038710
                                },
                                "raw_data_hex": "0a02ac7f2208648c4cb65453b4de40e8f991cc9c2d5a69080112650a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412340a15410583a68a3bcd86c25ab1bee482bac04a216b026112154139fec4d95bb59f45a727f9234020adaf2cec9e201880dcf3840170b6b28ecc9c2d",
                                "ret": [
                                    {
                                        "code": "SUCESS",
                                        "contractRet": "SUCCESS",
                                        "fee": 0
                                    }
                                ],
                                "signature": [
                                    "25052c3c653de8c06dc2294f9078a621179102282799f14f4a6950b805e66ea17c1a894a34922d92b7b14aad27cc07aeb129cb31570d55e9f21a0e3f5dd11ce701"
                                ],
                                "txID": "3ef5e225ce5bdd01333286e4ab4413ae3da8b80c24b26c5811ae78962940a8ca"
                            }
                        ]
                    }
                `);
            }
        }

        if (typeof params.operation === 'undefined') {
            if (params.address === 'TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB') {
                return JSON.parse(`
                    {
                        "success": true,
                        "meta": {
                            "at": 1586264949424,
                            "page_size": 1
                        },
                        "data": [
                            {
                                "account_resource": {},
                                "active_permission": [
                                    {
                                        "id": 2,
                                        "keys": [
                                            {
                                                "address": "4139fec4d95bb59f45a727f9234020adaf2cec9e20",
                                                "weight": 1
                                            }
                                        ],
                                        "operations": "7fff1fc0033e0000000000000000000000000000000000000000000000000000",
                                        "permission_name": "active",
                                        "threshold": 1,
                                        "type": "Active"
                                    }
                                ],
                                "address": "4139fec4d95bb59f45a727f9234020adaf2cec9e20",
                                "allowance": 848829,
                                "assetV2": [
                                    {
                                        "key": "1002000",
                                        "value": 10183240058
                                    },
                                    {
                                        "key": "1002798",
                                        "value": 10000000
                                    },
                                    {
                                        "key": "1002814",
                                        "value": 10000000
                                    }
                                ],
                                "balance": 278720000,
                                "create_time": 1553864037000,
                                "free_asset_net_usageV2": [
                                    {
                                        "key": "1002000",
                                        "value": 0
                                    },
                                    {
                                        "key": "1002798",
                                        "value": 0
                                    },
                                    {
                                        "key": "1002814",
                                        "value": 0
                                    }
                                ],
                                "latest_consume_free_time": 1575142326000,
                                "latest_consume_time": 1576767612000,
                                "latest_opration_time": 1576767612000,
                                "owner_permission": {
                                    "keys": [
                                        {
                                            "address": "4139fec4d95bb59f45a727f9234020adaf2cec9e20",
                                            "weight": 1
                                        }
                                    ],
                                    "permission_name": "owner",
                                    "threshold": 1
                                },
                                "trc20": [
                                    {
                                        "TLa2f6VPqDgRE67v1736s7bJ8Ray5wYjU7": "50356946"
                                    }
                                ]
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
