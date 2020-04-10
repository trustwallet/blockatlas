/// Callisto API Mock
/// See:
/// curl "http://localhost:3000/callisto-api/transactions?address=0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7"
/// curl "http://{callisto rpc}/transactions?address=0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7"
/// curl http://localhost:8420/v1/callisto/0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7

module.exports = {
    path: '/callisto-api/transactions',
    template: function(params, query, body) {
        //console.log(query)
        if (query.address === '0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7') {
            return JSON.parse(`
                {
                    "docs": [
                        {
                            "operations": [],
                            "contract": null,
                            "_id": "0x4052f55ee615a783abaa22a62100a45d2eee27ecab3a599e9d6691371336463e",
                            "blockNumber": 4699868,
                            "time": 1584952936,
                            "nonce": 7973,
                            "from": "0x28D2c7db63a9fC5f0b1eE3Ee2B7c0c350Eb9256A",
                            "to": "0x3083a7Ec44ca2b038D4BE4B0798152F948f0f3d7",
                            "value": "120071054450000000000",
                            "gas": "21000",
                            "gasPrice": "50000000000",
                            "gasUsed": "21000",
                            "input": "0x",
                            "error": "",
                            "id": "0x4052f55ee615a783abaa22a62100a45d2eee27ecab3a599e9d6691371336463e",
                            "timeStamp": "1584952936"
                        },
                        {
                            "operations": [],
                            "contract": null,
                            "_id": "0x8d50f2e48d3616e55c4adf2b579aa4696459714c2777baf01d43333389acf9c5",
                            "blockNumber": 4699693,
                            "time": 1584950541,
                            "nonce": 1466,
                            "from": "0x1212c7aE2d01E3d174C759EF616c46C9C50bd94f",
                            "to": "0x3083a7Ec44ca2b038D4BE4B0798152F948f0f3d7",
                            "value": "179512617388000000000",
                            "gas": "21000",
                            "gasPrice": "50000000000",
                            "gasUsed": "21000",
                            "input": "0x",
                            "error": "",
                            "id": "0x8d50f2e48d3616e55c4adf2b579aa4696459714c2777baf01d43333389acf9c5",
                            "timeStamp": "1584950541"
                        }
                    ],
                    "total": 2
                }
            `);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
