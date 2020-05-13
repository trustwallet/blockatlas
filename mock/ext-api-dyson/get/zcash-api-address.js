/// Mock for external Zcash API
/// See:
/// curl "http://{Zcash rpc}/api/v2/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs"
/// curl "http://localhost:3347/mock/zcash-api/v2/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs"
/// curl "http://localhost:8437/v1/zcash/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX"

module.exports = {
    path: '/mock/zcash-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 't1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fzcash-api%2Fv2%2Faddress%2Ft1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
