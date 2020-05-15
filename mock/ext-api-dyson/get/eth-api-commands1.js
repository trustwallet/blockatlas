/// Ethereum API Mock
/// See:
/// curl "http://localhost:3347/mock/eth-api/transactions?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:3347/mock/eth-api/tokens?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "https://{eth rpc}/transactions?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "https://{eth rpc}/tokens?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:8437/v1/ethereum/0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:8437/v2/ethereum/tokens/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?Authorization=Bearer"

module.exports = {
    path: '/mock/eth-api/:command1?',
    template: function(params, query) {
        switch (params.command1) {
            case 'transactions':
                if (query.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Feth-api%2Ftransactions%3Faddress%3D0x0875BCab22dE3d02402bc38aEe4104e1239374a7.json";
                    var json = require(fn);
                    return json;
                }

            case 'tokens':
                if (query.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Feth-api%2Ftokens%3Faddress%3D0x0875BCab22dE3d02402bc38aEe4104e1239374a7.json";
                    var json = require(fn);
                    return json;
                }
        }

        return {error: "Not implemented"};
    }
};
