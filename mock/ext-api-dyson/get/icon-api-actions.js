/// Icon API Mock
/// See:
/// curl "http://localhost:3000/icon-api/address/txList?address=hxee691e7bccc4eb11fee922896e9f51490e62b12e&count=25"
/// curl "https://tracker.icon.foundation/v3/address/txList?address=hxee691e7bccc4eb11fee922896e9f51490e62b12e&count=25"
/// curl http://localhost:8420/v1/icon/hxee691e7bccc4eb11fee922896e9f51490e62b12e
module.exports = {
    path: "/icon-api/address/:action?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.action === 'txList') {
            if (query.address === 'hxee691e7bccc4eb11fee922896e9f51490e62b12e') {
                return JSON.parse(`
                    {
                        "data": [
                            {
                                "txHash": "0x3b1a382884091e683350d6285d908406c149a030a7d52eb347f4571c1c904382",
                                "height": 7044559,
                                "createDate": "2019-08-13T06:53:56.000+0000",
                                "fromAddr": "hxee691e7bccc4eb11fee922896e9f51490e62b12e",
                                "toAddr": "hx06d5b88bb7089033a2d8a2b61b8a305ecff8774f",
                                "txType": "0",
                                "dataType": "icx",
                                "amount": "0.498",
                                "fee": "0.001",
                                "state": 1,
                                "errorMsg": null,
                                "targetContractAddr": null,
                                "id": null
                            },
                            {
                                "txHash": "0x990b7f289aa465369e638b4f719c99f0d050f9756c19081ba86b2bae5b8e2f0d",
                                "height": 361987,
                                "createDate": "2019-04-17T01:16:02.000+0000",
                                "fromAddr": "hxee691e7bccc4eb11fee922896e9f51490e62b12e",
                                "toAddr": "hxee691e7bccc4eb11fee922896e9f51490e62b12e",
                                "txType": "0",
                                "dataType": "icx",
                                "amount": "0.00347",
                                "fee": "0.001",
                                "state": 1,
                                "errorMsg": null,
                                "targetContractAddr": null,
                                "id": null
                            },
                            {
                                "txHash": "0x3a455daca1f5e4588adbc6ac5cdfd84126d0617d3ebcf5610eb44b15dfc71402",
                                "height": 112100,
                                "createDate": "2018-11-23T21:52:15.000+0000",
                                "fromAddr": "hx3709f4e615f072158247ca4973218e1d40e0ea35",
                                "toAddr": "hxee691e7bccc4eb11fee922896e9f51490e62b12e",
                                "txType": "0",
                                "dataType": "icx",
                                "amount": "0.5",
                                "fee": "0.001",
                                "state": 1,
                                "errorMsg": null,
                                "targetContractAddr": null,
                                "id": null
                            }
                        ],
                        "listSize": 3,
                        "totalSize": 3,
                        "result": "200",
                        "description": "success"
                    }
                `);
            }
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
