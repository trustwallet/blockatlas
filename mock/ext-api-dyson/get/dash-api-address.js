/// Mock for external Dash API
/// See:
/// curl "http://{dash rpc}/api/v2/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG?details=txs"
/// curl "http://localhost:3347/mock/dash-api/v2/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG?details=txs"
/// curl "http://localhost:8437/v1/dash/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG"

module.exports = {
    path: '/mock/dash-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fdash-api%2Fv2%2Faddress%2FXrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
