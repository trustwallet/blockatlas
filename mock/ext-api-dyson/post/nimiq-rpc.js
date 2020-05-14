/// Nimiq RPC API Mock
/// See:
/// curl -H "Content-Type: application/json" -d '{"jsonrpc": "2.0", "method": "getTransactionsByAddress", "params": [ "NQ94HC9AK9D83FSJM6PT8XGNNMXLR0E53Y07", "25" ], "id": "getTransactionsByAddress"}' https://{nimiq_rpc}
/// curl -H "Content-Type: application/json" -d '{"jsonrpc": "2.0", "method": "getTransactionsByAddress", "params": [ "NQ94HC9AK9D83FSJM6PT8XGNNMXLR0E53Y07", "25" ], "id": "getTransactionsByAddress"}' http://localhost:3347/mock/nimiq-rpc
/// curl "http://localhost:8437/v1/nimiq/NQ94HC9AK9D83FSJM6PT8XGNNMXLR0E53Y07"

module.exports = {
    path: '/mock/nimiq-rpc',
    template: function(params, query, body) {
        if (body.method === 'getTransactionsByAddress') {
            //console.log('body.params', body.params);
            if (body.params.length == 0) {
                return {error: "missing parameters"};
            }
            var address = body.params[0];
            switch (address) {
                case 'NQ94HC9AK9D83FSJM6PT8XGNNMXLR0E53Y07':
                    var fn = "../../ext-api-data/post/" +
                        "mock%2Fnimiq-rpc.json";
                    var json = require(fn);
                    return json;    
            }
        }
        return {error: "Not implemented"};
    }
};
