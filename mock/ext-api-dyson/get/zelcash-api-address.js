/// Mock for external Zelcash API
/// See:
/// curl "http://{Zelcash rpc}/api/v2/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa?details=txs&pageSize=25"
/// curl "http://localhost:3347/mock/zelcash-api/v2/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa?details=txs"
/// curl "http://localhost:8437/v1/zelcash/address/t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa"

module.exports = {
    path: '/mock/zelcash-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 't1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fzelcash-api%2Fv2%2Faddress%2Ft1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
