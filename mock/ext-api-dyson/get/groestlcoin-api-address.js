/// Mock for external Groestlcoin API
/// See:
/// curl "http://{groestlcoin rpc}/api/v2/address/33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj?details=txs"
/// curl "http://localhost:3347/mock/groestlcoin-api/v2/address/33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj?details=txs"
/// curl "http://localhost:8437/v1/groestlcoin/address/33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj"

module.exports = {
    path: '/mock/groestlcoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case '33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fgroestlcoin-api%2Fv2%2Faddress%2F33Ym3fecmWaHD19jymYt6fGd9TqSDQFfQj%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
