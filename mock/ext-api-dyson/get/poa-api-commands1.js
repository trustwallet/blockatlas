/// POA API Mock
/// See:
/// curl "http://localhost:3000/poa-api/transactions?address=0x55798eCbF17ce1241d543c22dCE46134c13b4bc0"
/// curl "http://{poa rpc}/transactions?address=0x55798eCbF17ce1241d543c22dCE46134c13b4bc0"
/// curl http://localhost:8420/v1/poa/0x55798eCbF17ce1241d543c22dCE46134c13b4bc0

module.exports = {
    path: '/poa-api/transactions',
    template: function(params, query, body) {
        //console.log(query)
        if (query.address === '0x55798eCbF17ce1241d543c22dCE46134c13b4bc0') {
            return JSON.parse(`
                {
                    "docs": [
                        {
                            "operations": [
                                {
                                    "transactionId": "0x544f5c54ad3c0b048ae81ef6f14507f934c4e53031bb2d39e074871c22f5ff9d-0",
                                    "contract": {
                                        "address": "0xab2f2dd3120de530d38936ee09a74a6d17e3da44",
                                        "decimals": 18,
                                        "name": "GeonCoin",
                                        "symbol": "GC",
                                        "totalSupply": "100000000000000000000000000",
                                        "updatedAt": "2020-02-26T23:41:29.763Z"
                                    },
                                    "from": "0x55798eCbF17ce1241d543c22dCE46134c13b4bc0",
                                    "to": "0xE87E46032847e097F5bB51404d19A0449c704EE8",
                                    "type": "token_transfer",
                                    "value": "1000000000000000000",
                                    "id": null
                                }
                            ],
                            "contract": null,
                            "_id": "0x544f5c54ad3c0b048ae81ef6f14507f934c4e53031bb2d39e074871c22f5ff9d",
                            "blockNumber": 14115901,
                            "time": 1584448420,
                            "nonce": 1,
                            "from": "0x55798eCbF17ce1241d543c22dCE46134c13b4bc0",
                            "to": "0xAb2f2Dd3120dE530d38936EE09A74a6d17e3Da44",
                            "value": "0",
                            "gas": "120000",
                            "gasPrice": "1000000000",
                            "gasUsed": "78484",
                            "input": "0xb88a3f725e70c38d1d03b5568e2f4d1600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a7640000",
                            "error": "",
                            "id": "0x544f5c54ad3c0b048ae81ef6f14507f934c4e53031bb2d39e074871c22f5ff9d",
                            "timeStamp": "1584448420"
                        },
                        {
                            "operations": [],
                            "contract": null,
                            "_id": "0xcde1f2c3a262012f624065f167616db3160e763f4938ccfa0fcc742d5725eae2",
                            "blockNumber": 14115900,
                            "time": 1584448415,
                            "nonce": 109291,
                            "from": "0x628cF7150f8242b20B2ADB064D27183D2a026130",
                            "to": "0x55798eCbF17ce1241d543c22dCE46134c13b4bc0",
                            "value": "600000000000000",
                            "gas": "120000",
                            "gasPrice": "1000000000",
                            "gasUsed": "21000",
                            "input": "0x",
                            "error": "",
                            "id": "0xcde1f2c3a262012f624065f167616db3160e763f4938ccfa0fcc742d5725eae2",
                            "timeStamp": "1584448415"
                        },
                        {
                            "operations": [],
                            "contract": null,
                            "_id": "0x90ef177d3448d379560218f99657027661a7238fb62c793dc2a7091a0d0d88e8",
                            "blockNumber": 14115899,
                            "time": 1584448410,
                            "nonce": 0,
                            "from": "0x55798eCbF17ce1241d543c22dCE46134c13b4bc0",
                            "to": "0xE87E46032847e097F5bB51404d19A0449c704EE8",
                            "value": "0",
                            "gas": "120000",
                            "gasPrice": "1000000000",
                            "gasUsed": "100890",
                            "input": "0x1e76c9445e70c38d1d03b5568e2f4d160000000000000000000000000000000000000000",
                            "error": "",
                            "id": "0x90ef177d3448d379560218f99657027661a7238fb62c793dc2a7091a0d0d88e8",
                            "timeStamp": "1584448410"
                        },
                        {
                            "operations": [],
                            "contract": null,
                            "_id": "0x6fbe91787b01664f99ea928fb1a314aed4fcd777a9fae063ed74b838315120c3",
                            "blockNumber": 14115898,
                            "time": 1584448405,
                            "nonce": 109290,
                            "from": "0x628cF7150f8242b20B2ADB064D27183D2a026130",
                            "to": "0x55798eCbF17ce1241d543c22dCE46134c13b4bc0",
                            "value": "600000000000000",
                            "gas": "120000",
                            "gasPrice": "1000000000",
                            "gasUsed": "21000",
                            "input": "0x",
                            "error": "",
                            "id": "0x6fbe91787b01664f99ea928fb1a314aed4fcd777a9fae063ed74b838315120c3",
                            "timeStamp": "1584448405"
                        }
                    ],
                    "total": 4
                }
            `);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
