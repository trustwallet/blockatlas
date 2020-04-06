/// Tron API Mock
/// See:
/// curl http://localhost:3000/tron-api/wallet/listwitnesses
/// curl http://localhost:8420/v2/tron/staking/validators

module.exports = {
    path: "/tron-api/wallet/:operation",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.operation === 'getaccount') {
            return {balance: 1, assetV2: [], votes: [], frozen: []}
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
