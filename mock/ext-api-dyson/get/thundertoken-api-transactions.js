/// Thundertoken API Mock
/// See:
/// curl "http://localhost:3000/thundertoken-api/transactions?address=0x0b230def08139f18a86536d9cfa150f04435414c"
/// curl "http://{Thundertoken rpc}/transactions?address=0x0b230def08139f18a86536d9cfa150f04435414c"
/// curl http://localhost:8420/v1/thundertoken/0x0b230def08139f18a86536d9cfa150f04435414c

module.exports = {
    path: '/thundertoken-api/transactions',
    template: function(params, query, body) {
        //console.log(query)
        if (query.address === '0x0b230def08139f18a86536d9cfa150f04435414c') {
            return JSON.parse(`{
                "docs": [
                    {
                        "operations": [],
                        "contract": null,
                        "_id": "0x80929a2170969364e06d253d2b853fe0b2ca3a307e61d629946b9b76e2f983cf",
                        "blockNumber": 33514713,
                        "time": 1584945288,
                        "nonce": 312,
                        "from": "0x0B230dEf08139F18a86536d9CFa150f04435414c",
                        "to": "0x83Ed8e6135e079a49085CA2f2F5E502691ccC0B1",
                        "value": "0",
                        "gas": "4700000",
                        "gasPrice": "4000000000",
                        "gasUsed": "421569",
                        "input": "0xc78b6dea0000000000000000000000000000000000000000000000000000000000000800",
                        "error": "",
                        "id": "0x80929a2170969364e06d253d2b853fe0b2ca3a307e61d629946b9b76e2f983cf",
                        "timeStamp": "1584945288"
                    },
                    {
                        "operations": [
                            {
                                "transactionId": "0xa8f78b26a4bf9e89bd2662cb2dcd2e9031a8d38df2ed85d6d52526589bbfed00-0",
                                "contract": {
                                    "address": "0x51bca9300d034a7e6ab379399a317bdfa0d4835b",
                                    "decimals": 18,
                                    "name": "The Third Identity",
                                    "symbol": "TTI",
                                    "totalSupply": "10000000000000000000000000000",
                                    "updatedAt": "2020-02-26T21:42:19.178Z"
                                },
                                "from": "0x3270b4e0916B4556BF236B6E5D4377282b3255e5",
                                "to": "0x0B230dEf08139F18a86536d9CFa150f04435414c",
                                "type": "token_transfer",
                                "value": "292820000000000000000",
                                "id": null
                            }
                        ],
                        "contract": null,
                        "_id": "0xa8f78b26a4bf9e89bd2662cb2dcd2e9031a8d38df2ed85d6d52526589bbfed00",
                        "blockNumber": 33514370,
                        "time": 1584944945,
                        "nonce": 311,
                        "from": "0x0B230dEf08139F18a86536d9CFa150f04435414c",
                        "to": "0x3270b4e0916B4556BF236B6E5D4377282b3255e5",
                        "value": "0",
                        "gas": "4700000",
                        "gasPrice": "2880000000",
                        "gasUsed": "1511414",
                        "input": "0x4008d2f3",
                        "error": "",
                        "id": "0xa8f78b26a4bf9e89bd2662cb2dcd2e9031a8d38df2ed85d6d52526589bbfed00",
                        "timeStamp": "1584944945"
                    }
                ],
                "total": 2
            }`);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
