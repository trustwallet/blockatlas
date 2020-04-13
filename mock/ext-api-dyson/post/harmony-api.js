/// Harmony RPC Mock
/// curl -H 'Content-Type: application/json' -d ' {"jsonrpc":"2.0","method":"hmy_getTransactionsHistory","params":[{"address":"one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv","fullTx":true}],"id":"hmy_getTransactionsHistory"} ' http://localhost:3000/harmony-api
/// curl -H 'Content-Type: application/json' -d ' {"jsonrpc":"2.0","method":"hmy_getTransactionsHistory","params":[{"address":"one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv","fullTx":true}],"id":"hmy_getTransactionsHistory"} ' https://{harmony_rpc}
/// curl "http://localhost:8420/v2/harmony/one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv"

module.exports = {
    path: '/harmony-api',
    template: function(params, query, body) {
        //console.log("curl -H 'Content-Type: application/json' -d '", JSON.stringify(body), "' https://{harmony_rpc}");
        if (body.method === 'hmy_getTransactionsHistory') {
            //console.log('body.params[0].address', body.params[0].address);
            if (body.params[0].address === 'one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv') {
                return JSON.parse(`{
                    "jsonrpc": "2.0",
                    "id": "hmy_getTransactionsHistory",
                    "result": {
                        "transactions": [
                            {
                                "blockHash": "0x1b45da220f94f41bd42dc0eb89a5e83ce77b37cc78dbc9032bd10a5bbab13674",
                                "blockNumber": "0x1b2cdb",
                                "from": "one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv",
                                "timestamp": "0x5df7cb45",
                                "gas": "0x33450",
                                "gasPrice": "0x3b9aca00",
                                "hash": "0x6bfd003239bf137855342f6266c456ae1b342e886bd1965a47edd4c74dd1a687",
                                "input": "0x",
                                "nonce": "0x0",
                                "to": "one1syjs6cnfwd9fgrhng03dyzs07suwtywwreczmk",
                                "transactionIndex": "0x0",
                                "value": "0x1bbfef6a6ff9c000",
                                "shardID": 0,
                                "toShardID": 0,
                                "v": "0x26",
                                "r": "0x44c2a7b8c38612a3ce8b51967947c0c9756dda8d82d0f12137451ce33bbef390",
                                "s": "0x34234c6c26fe36bda246db845692fd248cb4060cda0d1c6056298b39486ed4f2"
                            },
                            {
                                "blockHash": "0x58cc90d1b9302d0121b426f5fd35014acbc009defe36c5bbff6ac920d6005526",
                                "blockNumber": "0x1b2cf2",
                                "from": "one1syjs6cnfwd9fgrhng03dyzs07suwtywwreczmk",
                                "timestamp": "0x5df7cc02",
                                "gas": "0x33450",
                                "gasPrice": "0x3b9aca00",
                                "hash": "0xf54f6f86df1624f38fcedf56eb248854a13a733e0b0f92d4b31b9cc7196d90b0",
                                "input": "0x",
                                "nonce": "0x20",
                                "to": "one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv",
                                "transactionIndex": "0x0",
                                "value": "0x1bc0ae68df60e000",
                                "shardID": 0,
                                "toShardID": 0,
                                "v": "0x25",
                                "r": "0x46d14df06948f055e519ac244bcaaa3c3446ed20671d41db01546ddaad32503b",
                                "s": "0x50d13aeb94f2e8c760723afa4d892f3ca17690451a57465ea7e39c5d347efe8"
                            },
                            {
                                "blockHash": "0xa84376f0f4d082b76355ccc6ec491660c0062d9ef3f39985ca00986e8dcbacc1",
                                "blockNumber": "0x1b2d1b",
                                "from": "one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv",
                                "timestamp": "0x5df7cd50",
                                "gas": "0x33450",
                                "gasPrice": "0x3b9aca00",
                                "hash": "0x73355f625d0f7ef31512657c0f669e710c8c91a67180cd28bc919cfca0a2af82",
                                "input": "0x",
                                "nonce": "0x1",
                                "to": "one1syjs6cnfwd9fgrhng03dyzs07suwtywwreczmk",
                                "transactionIndex": "0x0",
                                "value": "0x1bc09b4f6dd69000",
                                "shardID": 0,
                                "toShardID": 0,
                                "v": "0x26",
                                "r": "0xf31271b93e446bce8cc83eb997183146b4da2378a5f2cb1fc4ae64fbd65a93f5",
                                "s": "0x77e8e5693a87394204f96693a15f2157459b0329a2ddab97da0bf1cfb6fbce8c"
                            }
                        ]
                    }
                }`);
            }
            return {jsonrpc:"2.0",id:"hmy_getTransactionsHistory",result:{"transactions":[]}};
        }
        return {error: 'Invalid request'};
    }
};
