/// Mock for external Doge API
/// See:
/// curl "http://{doge rpc}/api/v2/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh?details=txs"
/// curl "http://localhost:3347/mock/doge-api/v2/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh?details=txs"
/// curl "http://localhost:8437/v1/doge/address/D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh"

module.exports = {
    path: '/mock/doge-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'D5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fdoge-api%2Fv2%2Faddress%2FD5dAUAx3Ezg1q4dRgzKTBsxp4VJietWkDh%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
