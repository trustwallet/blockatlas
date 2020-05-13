/// Mock for external Zcoin API
/// See:
/// curl "http://{Zcoin rpc}/api/v2/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn?details=txs"
/// curl "http://localhost:3347/mock/zcoin-api/v2/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn?details=txs"
/// curl "http://localhost:8437/v1/zcoin/address/a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn"

module.exports = {
    path: '/mock/zcoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'a8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fzcoin-api%2Fv2%2Faddress%2Fa8EF4cpenEgEn9hm2NL5KfFK1UmSZZaQVn%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
