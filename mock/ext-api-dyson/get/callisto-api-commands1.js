/// Callisto API Mock
/// See:
/// curl "http://localhost:3347/callisto-api/transactions?address=0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7"
/// curl "http://localhost:3347/callisto-api/tokens?address=0xc3d5b69f65027ddf48f894e6e90121293a2f6615"
/// curl "https://{callisto rpc}/transactions?address=0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7"
/// curl "https://{callisto rpc}/tokens?address=0xc3d5b69f65027ddf48f894e6e90121293a2f6615"
/// curl "http://localhost:8437/v1/callisto/0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7"
/// curl "http://localhost:8437/v2/callisto/tokens/0xc3d5b69f65027ddf48f894e6e90121293a2f6615?Authorization=Bearer"

module.exports = {
    path: '/callisto-api/:command1?',
    template: function(params, query, body) {
        //console.log(query);
        switch (params.command1) {
            case 'transactions':
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
                break;

            case 'tokens':
                if (query.address === '0xc3d5b69f65027ddf48f894e6e90121293a2f6615') {
                    return JSON.parse(`
                        {
                            "total": 1,
                            "docs": [
                                {
                                    "address": "0xc3d5B69F65027dDF48f894E6e90121293a2F6615",
                                    "name": "The Hitchhikers Guide to the Galaxy",
                                    "decimals": 18,
                                    "symbol": "H2G2"
                                }
                            ]
                        }
                    `);
                }
                break;        
        }
        
        return {error: "Not implemented"};
    }
};
