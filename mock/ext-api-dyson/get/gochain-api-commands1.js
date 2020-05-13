/// Gochain API Mock
/// See:
/// curl "http://localhost:3347/mock/gochain-api/transactions?address=0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896"
/// curl "http://localhost:3347/mock/gochain-api/tokens?address=0x0Fd98FB42C439E5F6484f7E71Caa6661d81d0628"
/// curl "https://{go rpc}/transactions?address=0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896"
/// curl "https://{go rpc}/ 
/// curl "http://localhost:8437/v1/gochain/0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896"
/// curl "http://localhost:8437/v2/gochain/tokens/0x0Fd98FB42C439E5F6484f7E71Caa6661d81d0628?Authorization=Bearer"

module.exports = {
    path: '/mock/gochain-api/:command1?',
    template: function(params, query) {
        switch (params.command1) {
            case 'transactions':
                if (query.address === '0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fgochain-api%2Ftransactions%3Faddress%3D0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896.json";
                    var json = require(fn);
                    return json;
                }
                break;

            case 'tokens':
                if (query.address === '0x0Fd98FB42C439E5F6484f7E71Caa6661d81d0628') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fgochain-api%2Ftokens%3Faddress%3D0x0Fd98FB42C439E5F6484f7E71Caa6661d81d0628.json";
                    var json = require(fn);
                    return json;
                }
                break;
        }

        return {error: "Not implemented"};
    }
};
