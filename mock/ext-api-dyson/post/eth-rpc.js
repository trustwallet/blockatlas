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
            switch (body.params[0].data) {
                case '0x0178b8bfee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835':
                    // name lookup, vitalik.eth, part 2
                    return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000226159d592e2b063810a10ebf6dcbada94ed68b8"};
                case '0x02571be32337fcf9521666fd5114df2902fa9e6da5a6b004b5a192a0f55d2d9fab4f1047':
                    // name lookup, ourxyzwallet.xyz, part 1
                    return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000b30b09da480dd907907d98af5102c2034ed7df60"};
                case '0x02571be3a538cd174de170e8062c97c25c40a7ceb56aca81b12e6f5afc46f5a429999aa4':
                    // name lookup, vitalik.luxe, part 1
                    return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000d8a667312d5260f12a306ae7730c754d938da86c"};
                case '0x02571be3ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835':
                    // name lookup, vitalik.eth, part 1
                    return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000d8da6bf26964af9d7eed9e03e53415d37aa96045"};
                case '0x0178b8bf2337fcf9521666fd5114df2902fa9e6da5a6b004b5a192a0f55d2d9fab4f1047':
                    // name lookup, ourxyzwallet.xyz, part 2
                    return {jsonrpc: jsonrpc, id: id, result: "0x0000000000000000000000001da022710df5002339274aadee8d58218e9d6ab5"};
                case '0x0178b8bfa538cd174de170e8062c97c25c40a7ceb56aca81b12e6f5afc46f5a429999aa4':
                    // name lookup, vitalik.luxe, part 2
                    return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000bd5f5ec7ed5f19b53726344540296c02584a5237"};
                case '0x3b3b57de2337fcf9521666fd5114df2902fa9e6da5a6b004b5a192a0f55d2d9fab4f1047':
                    // name lookup, ourxyzwallet.xyz, part 5
                    return {jsonrpc: jsonrpc, id: id, result: "0x0000000000000000000000000c54eead78d555be3cbcd451424f9a27a7843935"};
                case '0x3b3b57dea538cd174de170e8062c97c25c40a7ceb56aca81b12e6f5afc46f5a429999aa4':
                    // name lookup, vitalik.luxe, part 4
                    return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000d8a667312d5260f12a306ae7730c754d938da86c"};
                
                case '0x3b3b57deeb4f647bea6caa36333c816d7b46fdcb05f9466ecacc140ea8c66faf15b3d9f1':
                    switch (body.params[0].to) {
                        case '0x1da022710df5002339274aadee8d58218e9d6ab5':
                            // name lookup, ourxyzwallet.xyz, part 3
                            return {jsonrpc: jsonrpc, id: id, result: "0x0000000000000000000000000000000000000000000000000000000000000000"};
                        case '0x226159d592e2b063810a10ebf6dcbada94ed68b8':
                            // name lookup, vitalik.eth, part 3
                            return {jsonrpc: jsonrpc, id: id, result: "0x000000000000000000000000eefb13c7d42efcc655e528da6d6f7bbcf9a2251d"};
                        case '0xbd5f5ec7ed5f19b53726344540296c02584a5237':
                            // name lookup, vitalik.luxe, part 3
                            return {jsonrpc: jsonrpc, id: id, result: "0x0000000000000000000000006109dd117aa5486605fc85e040ab00163a75c662"};
                    }

                case '0xf1cb7e062337fcf9521666fd5114df2902fa9e6da5a6b004b5a192a0f55d2d9fab4f1047000000000000000000000000000000000000000000000000000000000000003c':
                    switch (body.params[0].to) {
                        case '0x1da022710df5002339274aadee8d58218e9d6ab5':
                            // name lookup, ourxyzwallet.xyz, part 4
                            return {jsonrpc: jsonrpc, id: id, result: "0x"};
                    }    

                case '0xf1cb7e06ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835000000000000000000000000000000000000000000000000000000000000003c':
                    switch (body.params[0].to) {
                        case '0x226159d592e2b063810a10ebf6dcbada94ed68b8':
                            // name lookup, vitalik.eth, part 4
                            return {jsonrpc: jsonrpc, id: id, result: "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000014d8da6bf26964af9d7eed9e03e53415d37aa96045000000000000000000000000"};
                        case '0xbd5f5ec7ed5f19b53726344540296c02584a5237':
                            // name lookup, vitalik.luxe, part 4
                            return {jsonrpc: jsonrpc, id: id, result: "0x"};
                    }
            }
            // fallback
            return {error: "wrong data"};
        } else {
            // fallback
            return {error: "wrong method"};
        }
    }
};
