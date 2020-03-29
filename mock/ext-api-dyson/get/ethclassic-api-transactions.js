/// Ethereum Classic API Mock
/// See:
/// curl "http://localhost:3000/ethclassic-api/transactions?address=0x7d2d0e153026fb428b885d86de50768d4cfeac37"
/// curl "http://{etc rpc}/transactions?address=0x7d2d0e153026fb428b885d86de50768d4cfeac37"
/// curl http://localhost:8420/v1/ethereumclassic/0x7d2d0e153026fb428b885d86de50768d4cfeac37

module.exports = {
    path: '/ethclassic-api/transactions',
    template: function(params, query, body) {
        //console.log(query)
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
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
