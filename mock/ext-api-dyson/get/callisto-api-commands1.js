/// Callisto API Mock
/// See:
/// curl "http://localhost:3347/mock/callisto-api/transactions?address=0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7"
/// curl "http://localhost:3347/mock/callisto-api/tokens?address=0xc3d5b69f65027ddf48f894e6e90121293a2f6615"
/// curl "https://{callisto rpc}/transactions?address=0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7"
/// curl "https://{callisto rpc}/tokens?address=0xc3d5b69f65027ddf48f894e6e90121293a2f6615"
/// curl "http://localhost:8437/v1/callisto/0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7"
/// curl "http://localhost:8437/v2/callisto/tokens/0xc3d5b69f65027ddf48f894e6e90121293a2f6615?Authorization=Bearer"

module.exports = {
    path: '/mock/callisto-api/:command1?',
    template: function(params, query, body) {
        //console.log(query);
        switch (params.command1) {
            case 'transactions':
                if (query.address === '0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fcallisto-api%2Ftransactions%3Faddress%3D0x3083a7ec44ca2b038d4be4b0798152f948f0f3d7.json";
                    var json = require(fn);
                    return json;
                }
                break;

            case 'tokens':
                if (query.address === '0xc3d5b69f65027ddf48f894e6e90121293a2f6615') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fcallisto-api%2Ftokens%3Faddress%3D0xc3d5b69f65027ddf48f894e6e90121293a2f6615.json";
                    var json = require(fn);
                    return json;
                }
                break;        
        }
        
        return {error: "Not implemented"};
    }
};
