/// Tron API Mock
/// See:
/// curl http://localhost:3347/mock/tron-api/wallet/listwitnesses
/// curl http://localhost:8437/v2/tron/staking/validators

module.exports = {
    path: "/mock/tron-api/wallet/:operation?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.operation === 'listwitnesses') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Ftron-api%2Fwallet%2Flistwitnesses.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"};
    }
};
