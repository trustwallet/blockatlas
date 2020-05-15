/// Harmony RPC Mock
/// curl -H 'Content-Type: application/json' -d '{"jsonrpc":"2.0","method":"hmy_getTransactionsHistory","params":[{"address":"one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv","fullTx":true}],"id":1}' http://localhost:3347/mock/harmony-api
/// curl -H 'Content-Type: application/json' -d '{"jsonrpc":"2.0","method":"hmy_getTransactionsHistory","params":[{"address":"one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv","fullTx":true}],"id":1}' https://{harmony_rpc}
/// curl "http://localhost:8437/v2/harmony/one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv"

module.exports = {
    path: '/mock/harmony-api',
    template: function(params, query, body) {
        //console.log("curl -H 'Content-Type: application/json' -d '", JSON.stringify(body), "' https://{harmony_rpc}");
        if (body.method === 'hmy_getTransactionsHistory') {
            //console.log('body.params[0].address', body.params[0].address);
            if (body.params[0].address === 'one1e4mr7tp0a76wnhv9xd0wzentdjnjnsh3fwzgfv') {
                var fn = "../../ext-api-data/post/" +
                    "mock%2Fharmony-api.json";
                var json = require(fn);
                return json;
            }
        }
        return {error: "Not implemented"};
    }
};
