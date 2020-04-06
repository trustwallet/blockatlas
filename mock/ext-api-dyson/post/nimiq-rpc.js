/// Nimiq RPC API Mock
/// See:
/// curl -H "Content-Type: application/json" -d '{"jsonrpc": "2.0", "method": "getTransactionsByAddress", "params": [ "NQ94HC9AK9D83FSJM6PT8XGNNMXLR0E53Y07", "25" ], "id": "getTransactionsByAddress"}' https://{nimiq_rpc}
/// curl -H "Content-Type: application/json" -d '{"jsonrpc": "2.0", "method": "getTransactionsByAddress", "params": [ "NQ94HC9AK9D83FSJM6PT8XGNNMXLR0E53Y07", "25" ], "id": "getTransactionsByAddress"}' http://localhost:3000/nimiq-rpc
/// curl "http://localhost:8420/v1/nimiq/NQ94HC9AK9D83FSJM6PT8XGNNMXLR0E53Y07"

module.exports = {
    path: '/nimiq-rpc',
    template: function(params, query, body) {
        //console.log(body);
        //console.log('body.method', body.method);
        if (body.method === 'getTransactionsByAddress') {
            //console.log('body.params', body.params);
            if (body.params.length == 0) {
                return {error: "missing parameters"};
            }
            var address = body.params[0];
            switch (address) {
                case 'NQ94HC9AK9D83FSJM6PT8XGNNMXLR0E53Y07':
                    return JSON.parse(`
                        {
                            "jsonrpc": "2.0",
                            "result": [
                                {
                                    "hash": "8674d985ac1ea2a75fe9b35ed11d629cef8adfed50db6a5eb125351e622bb405",
                                    "blockHash": "11b0a3aaf64119cfbc66b7fdb8c38d5734a7f50d5bbb0116ee579eb1de513c3f",
                                    "blockNumber": 543545,
                                    "timestamp": 1556458272,
                                    "confirmations": 475868,
                                    "from": "21c7fc976349d39188b3d239b3b86ab66daeafc2",
                                    "fromAddress": "NQ32 473Y R5T3 979R 325K S8UT 7E3A NRNS VBX2",
                                    "to": "8b12a9a5a81bf52a9afb47a16b57d4c81c51fc07",
                                    "toAddress": "NQ94 HC9A K9D8 3FSJ M6PT 8XGN NMXL R0E5 3Y07",
                                    "value": 42839,
                                    "fee": 138,
                                    "data": null,
                                    "flags": 0
                                },
                                {
                                    "hash": "c3c6e87cba620095a2aeb9898f07077d184e20a68f53798032857cc842dfb0c0",
                                    "blockHash": "9b51f58b8ca881e2b182e1f544d09e4f1fe265b0f478fd41bf0128099013abd8",
                                    "blockNumber": 354736,
                                    "timestamp": 1545080725,
                                    "confirmations": 664677,
                                    "from": "e9a799c4fce6198cfb2eaacdeda5231078227d62",
                                    "fromAddress": "NQ61 V6KR KH7U UQCQ RXRE MB6X T993 21U2 4YB2",
                                    "to": "8b12a9a5a81bf52a9afb47a16b57d4c81c51fc07",
                                    "toAddress": "NQ94 HC9A K9D8 3FSJ M6PT 8XGN NMXL R0E5 3Y07",
                                    "value": 99808,
                                    "fee": 192,
                                    "data": "3c7363726970743e64656275676765723b3c2f7363726970743e",
                                    "flags": 0
                                }
                            ],
                            "id": "getTransactionsByAddress"
                        }
                    `);
            }
            // fallback
            return {error: "wrong data"};
        } else {
            // fallback
            return {error: "wrong method"};
        }
    }
};
