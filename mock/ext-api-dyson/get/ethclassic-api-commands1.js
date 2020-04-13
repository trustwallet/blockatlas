/// Ethereum Classic API Mock
/// See:
/// curl "http://localhost:3000/ethclassic-api/transactions?address=0x7d2d0e153026fb428b885d86de50768d4cfeac37"
/// curl "http://localhost:3000/ethclassic-api/tokens?address=0xa12105efa0663147bddee178f6a741ac15676b79"
/// curl "https://{etc rpc}/transactions?address=0x7d2d0e153026fb428b885d86de50768d4cfeac37"
/// curl "https://{etc rpc}/tokens?address=0xa12105efa0663147bddee178f6a741ac15676b79"
/// curl "http://localhost:8420/v1/ethereumclassic/0x7d2d0e153026fb428b885d86de50768d4cfeac37"
/// curl "http://localhost:8420/v2/classic/tokens/0xa12105efa0663147bddee178f6a741ac15676b79?Authorization=Bearer"

module.exports = {
    path: '/ethclassic-api/:command1?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.command1) {
            case 'transactions':
                if (query.address === '0x7d2d0e153026fb428b885d86de50768d4cfeac37') {
                    return JSON.parse(`
                        {
                            "docs": [
                                {
                                    "operations": [],
                                    "contract": null,
                                    "_id": "0x8439b360732a43b4bbebe4f00b9f8b2dbe14c6bd52d70598ab7fe6c8503383fe",
                                    "blockNumber": 9833329,
                                    "time": 1582237362,
                                    "nonce": 7071,
                                    "from": "0x2d4BeCA01fDE8963b3A2DBf8a838e0F2EBaB0FF9",
                                    "to": "0x7D2D0E153026fB428B885D86De50768D4cFeaC37",
                                    "value": "18529160000000000",
                                    "gas": "21000",
                                    "gasPrice": "1000000000",
                                    "gasUsed": "21000",
                                    "input": "0x",
                                    "error": "",
                                    "id": "0x8439b360732a43b4bbebe4f00b9f8b2dbe14c6bd52d70598ab7fe6c8503383fe",
                                    "timeStamp": "1582237362"
                                },
                                {
                                    "operations": [],
                                    "contract": null,
                                    "_id": "0x524fb5fedace7bee35c165ac7f641d689122cb21390db7f7384507473a5ba058",
                                    "blockNumber": 9804290,
                                    "time": 1581854371,
                                    "nonce": 6918,
                                    "from": "0x2d4BeCA01fDE8963b3A2DBf8a838e0F2EBaB0FF9",
                                    "to": "0x7D2D0E153026fB428B885D86De50768D4cFeaC37",
                                    "value": "16208400000000000",
                                    "gas": "21000",
                                    "gasPrice": "1000000000",
                                    "gasUsed": "21000",
                                    "input": "0x",
                                    "error": "",
                                    "id": "0x524fb5fedace7bee35c165ac7f641d689122cb21390db7f7384507473a5ba058",
                                    "timeStamp": "1581854371"
                                }
                            ],
                            "total": 2
                        }
                    `);
                }

            case 'tokens':
                if (query.address === '0xa12105efa0663147bddee178f6a741ac15676b79') {
                    return JSON.parse(`
                        {
                            "total": 2,
                            "docs": [
                                {
                                    "address": "0xCA68fE57A0E9987F940Ebcc65fe5F96E7bC30128",
                                    "name": "Litecoin Classic Token",
                                    "decimals": 8,
                                    "symbol": "LCT"
                                },
                                {
                                    "address": "0x2B682bd9d5c31E67a95cbdF0292017C02E51923C",
                                    "name": "Ether Klown",
                                    "decimals": 6,
                                    "symbol": "KLOWN2"
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
