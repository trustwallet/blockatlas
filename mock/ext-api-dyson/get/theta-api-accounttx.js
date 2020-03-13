/// Theta API Mock
/// See:
/// curl "http://localhost:3000/theta-api/accounttx/0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f?type=2&pageNumber=1&limitNumber=100&isEqualType=true"
/// curl "https://explorer.thetatoken.org:9000/api/accounttx/0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f?type=2&pageNumber=1&limitNumber=100&isEqualType=true"
/// curl http://localhost:8420/v1/theta/0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f
module.exports = {
    path: "/theta-api/accounttx/:address?",
    template: function(params, query, body) {
        //console.log(params)
        if (params.address === '0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f') {
            return JSON.parse(`
                {
                    "type": "account_tx_list",
                    "body": [
                        {
                            "_id": "0x85aaf834470dd858dc2f60c7bd535ddf08c3fcfef9456f43676462950a355b78",
                            "block_height": "4160784",
                            "data": {
                                "fee": {
                                    "thetawei": "0",
                                    "tfuelwei": "2000000000000"
                                },
                                "inputs": [
                                    {
                                        "address": "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
                                        "coins": {
                                            "thetawei": "1000000000000000",
                                            "tfuelwei": "2000000000000"
                                        },
                                        "sequence": "52",
                                        "signature": "0x60ae320a7a0f7f6e72ed5080830d2156f145a191197092df9eb88e7b8c7637ad01894b128b81d33b67a6516e3a6c7dc064576eb40b9516b4dc127e9845ede8c300"
                                    }
                                ],
                                "outputs": [
                                    {
                                        "address": "0x082a2aef39b6473c55ba7e28c122ed3aea0de381",
                                        "coins": {
                                            "thetawei": "1000000000000000",
                                            "tfuelwei": "0"
                                        }
                                    }
                                ]
                            },
                            "hash": "0x85aaf834470dd858dc2f60c7bd535ddf08c3fcfef9456f43676462950a355b78",
                            "number": 31836666,
                            "status": "finalized",
                            "timestamp": "1579009947",
                            "type": 2
                        },
                        {
                            "_id": "0x43804bbc7ff254f5eb3ff0a9bc1feed5788cbbc3eb5b535c52de4c47c8963dfd",
                            "block_height": "4160736",
                            "data": {
                                "fee": {
                                    "thetawei": "0",
                                    "tfuelwei": "2000000000000"
                                },
                                "inputs": [
                                    {
                                        "address": "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
                                        "coins": {
                                            "thetawei": "1000000000000000",
                                            "tfuelwei": "2000000000000"
                                        },
                                        "sequence": "51",
                                        "signature": "0x23cd4f14859ef6a9cd0bd820591101cd53c74d3421e01b42ca0671501afdd6f73dcf2ec8e35ceba773319ba552c0b170daa3460bf6125992e846f9d0cf839a9700"
                                    }
                                ],
                                "outputs": [
                                    {
                                        "address": "0x082a2aef39b6473c55ba7e28c122ed3aea0de381",
                                        "coins": {
                                            "thetawei": "1000000000000000",
                                            "tfuelwei": "0"
                                        }
                                    }
                                ]
                            },
                            "hash": "0x43804bbc7ff254f5eb3ff0a9bc1feed5788cbbc3eb5b535c52de4c47c8963dfd",
                            "number": 31836430,
                            "status": "finalized",
                            "timestamp": "1579009642",
                            "type": 2
                        }
                    ],
                    "totalPageNumber": 1,
                    "currentPageNumber": 1
                }
            `);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
