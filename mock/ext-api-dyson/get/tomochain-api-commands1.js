/// Tomochain API Mock
/// See:
/// curl "http://localhost:3000/tomochain-api/transactions?address=0x17e4c16605e32adead5fa371bf6117df34ca0200"
/// curl "http://{Tomochain rpc}/transactions?address=0x17e4c16605e32adead5fa371bf6117df34ca0200"
/// curl http://localhost:8420/v1/tomochain/0x17e4c16605e32adead5fa371bf6117df34ca0200

module.exports = {
    path: '/tomochain-api/transactions',
    template: function(params, query, body) {
        //console.log(query)
        if (query.address === '0x17e4c16605e32adead5fa371bf6117df34ca0200') {
            return JSON.parse(`{
                "docs": [
                    {
                        "operations": [],
                        "contract": null,
                        "_id": "0xb2af5ad5eebaba1da3452e0281f9b6ec9e00a5bb832793a912eb3959b6c2fdd2",
                        "blockNumber": 18455252,
                        "time": 1584963130,
                        "nonce": 892939,
                        "from": "0x17e4C16605E32ADEaD5FA371BF6117Df34Ca0200",
                        "to": "0x0000000000000000000000000000000000000089",
                        "value": "0",
                        "gas": "200000",
                        "gasPrice": "0",
                        "gasUsed": "0",
                        "input": "0xe341eaa40000000000000000000000000000000000000000000000000000000001199ad29a43ad0df08d1ff128de7642208ce68720dc77bbe7e02acb2685951f73efea56",
                        "error": "",
                        "id": "0xb2af5ad5eebaba1da3452e0281f9b6ec9e00a5bb832793a912eb3959b6c2fdd2",
                        "timeStamp": "1584963130"
                    },
                    {
                        "operations": [],
                        "contract": null,
                        "_id": "0x9181d170fb77eb55e1dda42c55080230397702204247e8b242c8f6588f61f41d",
                        "blockNumber": 18455237,
                        "time": 1584963100,
                        "nonce": 892938,
                        "from": "0x17e4C16605E32ADEaD5FA371BF6117Df34Ca0200",
                        "to": "0x0000000000000000000000000000000000000089",
                        "value": "0",
                        "gas": "200000",
                        "gasPrice": "0",
                        "gasUsed": "0",
                        "input": "0xe341eaa40000000000000000000000000000000000000000000000000000000001199ac3f1234069ee58f791cc422978d794e7f3c1e82d580013038badd5493ed854b6e4",
                        "error": "",
                        "id": "0x9181d170fb77eb55e1dda42c55080230397702204247e8b242c8f6588f61f41d",
                        "timeStamp": "1584963100"
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
