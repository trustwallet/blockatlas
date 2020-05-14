/// Tron API Mock
/// See:
/// curl http://localhost:3347/mock/tron-api/wallet/listwitnesses
/// curl http://localhost:8437/v2/tron/staking/delegations/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB?Authorization=Bearer

module.exports = {
    path: "/mock/tron-api/wallet/:operation",
    template: function(params, query, body) {
        if (params.operation === 'getaccount') {
            var fn = "../../ext-api-data/post/" +
                "mock%2Ftron-api%2Fwallet%2Fgetaccount.json";
            var json = require(fn);
            return json;
        }
        return {error: "Not implemented"};
    }
};
