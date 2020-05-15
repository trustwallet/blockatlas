/// Mock for external Ravencoin API
/// See:
/// curl "http://{Ravencoin rpc}/api/v2/address/RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo?details=txs"
/// curl "http://localhost:3347/mock/ravencoin-api/v2/address/RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo?details=txs"
/// curl "http://localhost:8437/v1/ravencoin/address/RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo"

module.exports = {
    path: '/mock/ravencoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fravencoin-api%2Fv2%2Faddress%2FRGkwvrUors8DtmhKy5bddFwRCTZaunjpvo%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
