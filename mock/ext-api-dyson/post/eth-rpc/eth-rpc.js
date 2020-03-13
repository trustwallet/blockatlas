/// ETH RPC API Mock
/// See:
/// curl -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","id":2,"method":"eth_call","params":[{"data":"0x02571be3ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835","from":"0x0000000000000000000000000000000000000000","to":"0x00000000000c2e074ec69a0dfb2997ba6c7d2e1e"},"latest"]}' http://localhost:3000/eth-rpc
/// curl "http://localhost:8420/v2/ns/lookup?name=vitalik.eth&coins=60"

module.exports = {
    path: '/eth-rpc',
    template: function(params, query, body) {
        //console.log(body);
        var jsonrpc = body.jsonrpc;
        var id = body.id;
        //console.log('body.method', body.method);
        if (body.method === 'net_version') {
            return {jsonrpc: jsonrpc, id: id, result: "1"}
        } else if (body.method === 'eth_call') {
            //console.log('body.params', body.params);
            //console.log('body.params[0].data', body.params[0].data);
            if (body.params[0].data === '0x02571be3ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835') {
                // name lookup, part 1
                return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000d8da6bf26964af9d7eed9e03e53415d37aa96045"}
            } else if (body.params[0].data === '0x0178b8bfee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835') {
                // name lookup, part 2
                return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000226159d592e2b063810a10ebf6dcbada94ed68b8"}
            } else if (body.params[0].data === '0x3b3b57deeb4f647bea6caa36333c816d7b46fdcb05f9466ecacc140ea8c66faf15b3d9f1') {
                // name lookup, part 3
                return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000eefb13c7d42efcc655e528da6d6f7bbcf9a2251d"}
            } else if (body.params[0].data === '0xf1cb7e06ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835000000000000000000000000000000000000000000000000000000000000003c') {
                // name lookup, part 4
                return {jsonrpc: jsonrpc, id: id, result: "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000014d8da6bf26964af9d7eed9e03e53415d37aa96045000000000000000000000000"}
            }
        }
    }
};
