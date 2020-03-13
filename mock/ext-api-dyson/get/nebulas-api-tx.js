/// Nebulas API Mock
/// See:
/// curl "http://localhost:3000/nebulas-api/tx?a=n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a&p=0"
/// curl "https://explorer-backend.nebulas.io/api/tx?a=n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a&p=0"
/// curl http://localhost:8420/v1/nebulas/n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a
module.exports = {
    path: "/nebulas-api/tx?",
    template: function(params, query, body) {
        //console.log(query)
        if (query.a === 'n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a') {
            return JSON.parse(`
                {
                    "code": 0,
                    "msg": "success",
                    "data": {
                        "txnList": [
                            {
                                "hash": "3cd98a767d8f5eaa87c8ffa95adbde3a4c08236b49ab0296f04057491e4bd1a3",
                                "block": {
                                    "hash": "9bb20444ae79a542af44ed64241a440759aa59363641fa1b7594266f6f16109d",
                                    "height": 2659950,
                                    "timestamp": null,
                                    "parentHash": null,
                                    "miner": null,
                                    "txnCnt": 0,
                                    "gasLimit": null,
                                    "avgGasPrice": null,
                                    "gasReward": null,
                                    "currentTimestamp": null,
                                    "timeDiff": null
                                },
                                "from": {
                                    "rank": null,
                                    "hash": "n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a",
                                    "alias": "",
                                    "balance": "3784780000000000",
                                    "percentage": null,
                                    "txCnt": null,
                                    "type": 0
                                },
                                "to": {
                                    "rank": null,
                                    "hash": "n1hiWG7Ce8HhTaJGzSJoAaJ9w1CJd7Do2rm",
                                    "alias": "",
                                    "balance": "200",
                                    "percentage": null,
                                    "txCnt": null,
                                    "type": 1
                                },
                                "status": 1,
                                "value": "0",
                                "nonce": 1,
                                "timestamp": 1562378850000,
                                "type": "call",
                                "gasPrice": "20000000000",
                                "gasLimit": "50000",
                                "gasUsed": "20711",
                                "createdAt": 1562379080000,
                                "currentTimestamp": 1585045274097,
                                "timeDiff": 22666424097,
                                "contractAddress": "",
                                "executeError": "",
                                "events": null,
                                "tokenName": null,
                                "decimal": null,
                                "txFee": "414220000000000"
                            },
                            {
                                "hash": "0d06074b64c37ed24a17162d1e0ef8a8e65122fc7758c90c969e61735bc4396a",
                                "block": {
                                    "hash": "faf4eb93b1a571c709bf1871e1c142570d3f5aeb787d325eba9a34a6ca56144b",
                                    "height": 2659949,
                                    "timestamp": null,
                                    "parentHash": null,
                                    "miner": null,
                                    "txnCnt": 0,
                                    "gasLimit": null,
                                    "avgGasPrice": null,
                                    "gasReward": null,
                                    "currentTimestamp": null,
                                    "timeDiff": null
                                },
                                "from": {
                                    "rank": null,
                                    "hash": "n1YDRkSbuzjRVYPbmGC9qULchXMeqS451Je",
                                    "alias": "",
                                    "balance": "37962400000000000000",
                                    "percentage": null,
                                    "txCnt": null,
                                    "type": 0
                                },
                                "to": {
                                    "rank": null,
                                    "hash": "n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a",
                                    "alias": "",
                                    "balance": "3784780000000000",
                                    "percentage": null,
                                    "txCnt": null,
                                    "type": 0
                                },
                                "status": 1,
                                "value": "10000000000000000",
                                "nonce": 210,
                                "timestamp": 1562378835000,
                                "type": "binary",
                                "gasPrice": "20000000000",
                                "gasLimit": "50000",
                                "gasUsed": "20000",
                                "createdAt": 1562379063000,
                                "currentTimestamp": 1585045274097,
                                "timeDiff": 22666439097,
                                "contractAddress": "",
                                "executeError": "",
                                "events": null,
                                "tokenName": null,
                                "decimal": null,
                                "txFee": "400000000000000"
                            }
                        ],
                        "totalPage": 1,
                        "maxDisplayCnt": 16,
                        "type": "address",
                        "currentPage": 0,
                        "txnCnt": 2
                    }
                }
            `);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
