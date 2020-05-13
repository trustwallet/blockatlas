/// POA API Mock
/// See:
/// curl "http://localhost:3347/mock/poa-api/transactions?address=0x55798eCbF17ce1241d543c22dCE46134c13b4bc0"
/// curl "http://localhost:3347/mock/poa-api/tokens?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "https://{poa rpc}/transactions?address=0x55798eCbF17ce1241d543c22dCE46134c13b4bc0"
/// curl "https://{poa rpc}/tokens?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "http://localhost:8437/v1/poa/0x55798eCbF17ce1241d543c22dCE46134c13b4bc0"
/// curl "http://localhost:8437/v2/poa/tokens/0x0875BCab22dE3d02402bc38aEe4104e1239374a7?Authorization=Bearer"

module.exports = {
    path: '/mock/poa-api/:command1?',
    template: function(params, query, body) {
        //console.log(params);
        //console.log(query);
        switch (params.command1) {
            case 'tokens':
                if (query.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fpoa-api%2Ftokens%3Faddress%3D0x0875BCab22dE3d02402bc38aEe4104e1239374a7.json";
                    var json = require(fn);
                    return json;
                }
                break;
                
            case 'transactions':
                if (query.address === '0x55798eCbF17ce1241d543c22dCE46134c13b4bc0') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fpoa-api%2Ftransactions%3Faddress%3D0x55798eCbF17ce1241d543c22dCE46134c13b4bc0.json";
                    var json = require(fn);
                    return json;
                }
                break;
        }

        return {error: "Not implemented"};
    }
};
