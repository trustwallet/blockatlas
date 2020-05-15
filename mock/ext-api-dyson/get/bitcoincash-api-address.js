/// Mock for external Bitcoincash API
/// See:
/// curl "https://{bch rpc}/api/v2/address/bitcoincash:qq07l6rr5lsdm3m80qxw80ku2ex0tj76vvsxpvmgme?details=txs"
/// curl "http://localhost:3347/mock/bitcoincash-api/v2/address/bitcoincash:qq07l6rr5lsdm3m80qxw80ku2ex0tj76vvsxpvmgme?details=txs"
/// curl "http://localhost:8437/v1/bitcoincash/address/bitcoincash:qq07l6rr5lsdm3m80qxw80ku2ex0tj76vvsxpvmgme"

module.exports = {
    path: '/mock/bitcoincash-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'bitcoincash:qq07l6rr5lsdm3m80qxw80ku2ex0tj76vvsxpvmgme':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fbitcoincash-api%2Fv2%2Faddress%2Fbitcoincash%3Aqq07l6rr5lsdm3m80qxw80ku2ex0tj76vvsxpvmgme%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
