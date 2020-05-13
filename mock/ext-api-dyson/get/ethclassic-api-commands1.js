/// Ethereum Classic API Mock
/// See:
/// curl "http://localhost:3347/mock/ethclassic-api/transactions?address=0x7d2d0e153026fb428b885d86de50768d4cfeac37"
/// curl "http://localhost:3347/mock/ethclassic-api/tokens?address=0xa12105efa0663147bddee178f6a741ac15676b79"
/// curl "https://{etc rpc}/transactions?address=0x7d2d0e153026fb428b885d86de50768d4cfeac37"
/// curl "https://{etc rpc}/tokens?address=0xa12105efa0663147bddee178f6a741ac15676b79"
/// curl "http://localhost:8437/v1/ethereumclassic/0x7d2d0e153026fb428b885d86de50768d4cfeac37"
/// curl "http://localhost:8437/v2/classic/tokens/0xa12105efa0663147bddee178f6a741ac15676b79?Authorization=Bearer"

module.exports = {
    path: '/mock/ethclassic-api/:command1?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.command1) {
            case 'transactions':
                if (query.address === '0x7d2d0e153026fb428b885d86de50768d4cfeac37') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fethclassic-api%2Ftransactions%3Faddress%3D0x7d2d0e153026fb428b885d86de50768d4cfeac37.json";
                    var json = require(fn);
                    return json;
                }
                break;

            case 'tokens':
                if (query.address === '0xa12105efa0663147bddee178f6a741ac15676b79') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fethclassic-api%2Ftokens%3Faddress%3D0xa12105efa0663147bddee178f6a741ac15676b79.json";
                    var json = require(fn);
                    return json;
                }
                break;
        }

        return {error: "Not implemented"};
    }
};
