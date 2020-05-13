/// Mock for external Litecoin API
/// See:
/// curl "http://{ltc rpc}/api/v2/address/ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept?details=txs"
/// curl "http://localhost:3347/mock/litecoin-api/v2/address/ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept?details=txs"
/// curl "http://localhost:8437/v1/litecoin/address/ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept"

module.exports = {
    path: '/mock/litecoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'ltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Flitecoin-api%2Fv2%2Faddress%2Fltc1qpm594ntjq6ayqjngf6t9td2dxtey9d7985eept%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
