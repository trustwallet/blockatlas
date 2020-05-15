/// Ethereum Blockbook API Mock
/// See:
/// curl "http://localhost:3347/mock/eth-blockbook-api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?details=txs"
/// curl "http://localhost:3347/mock/eth-blockbook-api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?details=tokenBalances"
/// curl "https://{eth blockbook api}/api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?details=txs"
/// curl "https://{eth blockbook api}/api/v2/address/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?details=tokenBalances"
/// curl "http://localhost:8437/v1/ethereum/0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:8437/v2/ethereum/tokens/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?Authorization=Bearer"

module.exports = {
    path: '/mock/eth-blockbook-api/v2/address/:address?',
    template: function(params, query) {
        if (params.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
            if (query.details === 'tokenBalances') {
                var fn = "../../ext-api-data/get/" +
                    "mock%2Feth-blockbook-api%2Fv2%2Faddress%2F0x0875BCab22dE3d02402bc38aEe4104e1239374a7%3Fdetails%3DtokenBalances.json";
                var json = require(fn);
                return json;
            }
            var fn = "../../ext-api-data/get/" +
                "mock%2Feth-blockbook-api%2Fv2%2Faddress%2F0x0875BCab22dE3d02402bc38aEe4104e1239374a7%3Fdetails%3Dtxs.json";
            var json = require(fn);
            return json;
        }

        return {error: "Not implemented"};
    }
};
