/// Aion API Mock
/// See:
/// curl "http://localhost:3000/aion-api/getTransactionsByAddress?accountAddress=0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed&size=25"
/// curl "https://mainnet-api.theoan.com/aion/dashboard/getTransactionsByAddress?accountAddress=0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed&size=25"
/// curl http://localhost:8420/v1/aion/0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed
module.exports = {
    path: "/aion-api/getTransactionsByAddress?",
    template: function(params, query, body) {
        //console.log(query)
        if (query.accountAddress === '0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed') {
            return JSON.parse(`
                {
                    "page": {
                        "number": 0,
                        "size": 25,
                        "totalPages": 1,
                        "start": 1582532306,
                        "end": 1585037906,
                        "totalElements": 2
                    },
                    "content": [
                        {
                            "blockHash": "f9aab3a58f9211d9d37458d8a87e49f2e00d1d242390d4650cd4092897650797",
                            "nrgPrice": 10000000000,
                            "toAddr": "a0c00cfc4c53bb6d49f385b91f58a23dd7b1dc024976b391bb8b81eb9e8801ab",
                            "contractAddr": "",
                            "data": "",
                            "year": 2020,
                            "internalTransactionCount": 0,
                            "transactionIndex": 0,
                            "type": "DEFAULT",
                            "nonce": "1e2f",
                            "transactionHash": "511d3a4aafbdd26825a1655b5c7525df5bea8a8ef0f887d5490b490055b53df7",
                            "transactionTimestamp": 1584349754,
                            "nrgConsumed": 21000,
                            "month": 3,
                            "blockNumber": 5619712,
                            "blockTimestamp": 1584349754,
                            "transactionLog": "[]",
                            "fromAddr": "a04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed",
                            "day": 16,
                            "value": "1.002217611212000000000000000000000000",
                            "txError": ""
                        },
                        {
                            "blockHash": "ed8d3d92da7228c35c363f55383c602b6f5041b37ff165231245cfa83b863ae7",
                            "nrgPrice": 10000000000,
                            "toAddr": "a04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed",
                            "contractAddr": "",
                            "data": "",
                            "year": 2020,
                            "internalTransactionCount": 0,
                            "transactionIndex": 2,
                            "type": "DEFAULT",
                            "nonce": "2b2e55",
                            "transactionHash": "84789693b108bc3c437f8dccbd517682f611036d54621ec1f4db26f661fde468",
                            "transactionTimestamp": 1584349550,
                            "nrgConsumed": 21000,
                            "month": 3,
                            "blockNumber": 5619690,
                            "blockTimestamp": 1584349550,
                            "transactionLog": "[]",
                            "fromAddr": "a00983f07c11ee9160a64dd3ba3dc3d1f88332a2869f25725f56cbd0be32ef7a",
                            "day": 16,
                            "value": "1.002427611212000000000000000000000000",
                            "txError": ""
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
