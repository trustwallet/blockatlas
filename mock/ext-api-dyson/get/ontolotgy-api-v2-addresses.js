/// Ontology API Mock
/// See:
/// curl "http://localhost:3000/ontology-api/v2/addresses/AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7/transactions?page_size=20&page_number=1"
/// curl "https://explorer.ont.io/v2/addresses/AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7/transactions?page_size=20&page_number=1"
/// curl http://localhost:8420/v1/ontology/AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7
module.exports = {
    path: "/ontology-api/v2/addresses/:address/:operation?",
    template: function(params, query, body) {
        //console.log(params)
        if (params.address === 'AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7') {
            return JSON.parse(`
                {
                    "code": 0,
                    "msg": "SUCCESS",
                    "result": [
                        {
                            "tx_hash": "d3be042ae21a313bc70fbfea62d11fb32731126101859a66f523e081c3fd429c",
                            "tx_type": 209,
                            "tx_time": 1582189609,
                            "block_height": 7836941,
                            "fee": "0.01",
                            "block_index": 1,
                            "confirm_flag": 1,
                            "transfers": [
                                {
                                    "amount": "0.4",
                                    "from_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "to_address": "Abm1FqnU4qur9bviXmrD5YnNixXGvMsi9R",
                                    "asset_name": "ong",
                                    "contract_hash": "0200000000000000000000000000000000000000"
                                },
                                {
                                    "amount": "0.01",
                                    "from_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "to_address": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
                                    "asset_name": "ong",
                                    "contract_hash": "0200000000000000000000000000000000000000"
                                }
                            ]
                        },
                        {
                            "tx_hash": "5c9a5f708953ee6cd13623e9d3fd610530159fafe71a8d207d890ed7b97f3ced",
                            "tx_type": 209,
                            "tx_time": 1582188446,
                            "block_height": 7836832,
                            "fee": "0.01",
                            "block_index": 1,
                            "confirm_flag": 1,
                            "transfers": [
                                {
                                    "amount": "0.02",
                                    "from_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "to_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "asset_name": "ong",
                                    "contract_hash": "0200000000000000000000000000000000000000"
                                },
                                {
                                    "amount": "0.01",
                                    "from_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "to_address": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
                                    "asset_name": "ong",
                                    "contract_hash": "0200000000000000000000000000000000000000"
                                }
                            ]
                        },
                        {
                            "tx_hash": "7e35568117aaed4e041d281087b2f47da91a4ac183c370d40040aa2b496d5878",
                            "tx_type": 209,
                            "tx_time": 1582122942,
                            "block_height": 7832679,
                            "fee": "0.01",
                            "block_index": 1,
                            "confirm_flag": 1,
                            "transfers": [
                                {
                                    "amount": "0.011074432",
                                    "from_address": "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV",
                                    "to_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "asset_name": "ong",
                                    "contract_hash": "0200000000000000000000000000000000000000"
                                },
                                {
                                    "amount": "0.01",
                                    "from_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "to_address": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
                                    "asset_name": "ong",
                                    "contract_hash": "0200000000000000000000000000000000000000"
                                }
                            ]
                        },
                        {
                            "tx_hash": "0ae15301f25a5067c075a25e722f0624a813a1352363197de2dd5f61ebd2998d",
                            "tx_type": 209,
                            "tx_time": 1581924056,
                            "block_height": 7819845,
                            "fee": "0.01",
                            "block_index": 1,
                            "confirm_flag": 1,
                            "transfers": [
                                {
                                    "amount": "1",
                                    "from_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "to_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "asset_name": "ont",
                                    "contract_hash": "0100000000000000000000000000000000000000"
                                },
                                {
                                    "amount": "0.01",
                                    "from_address": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
                                    "to_address": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
                                    "asset_name": "ong",
                                    "contract_hash": "0200000000000000000000000000000000000000"
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
