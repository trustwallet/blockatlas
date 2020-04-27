/// Tron API Mock
/// See:
/// curl http://localhost:3347/tron-api/wallet/listwitnesses
/// curl http://localhost:8437/v2/tron/staking/validators

module.exports = {
    path: "/tron-api/wallet/:operation",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.operation === 'getaccount') {
            return {balance: 1, assetV2: [], votes: [], frozen: []}
        }

        return {error: "Not implemented"};
    }
};
