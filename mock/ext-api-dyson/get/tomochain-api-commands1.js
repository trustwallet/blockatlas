/// Tomochain API Mock
/// See:
/// curl "http://localhost:3347/mock/tomochain-api/transactions?address=0x17e4c16605e32adead5fa371bf6117df34ca0200"
/// curl "http://localhost:3347/mock/tomochain-api/tokens?address=0x8b353021189375591723e7384262f45709a3c3dc"
/// curl "https://{Tomochain rpc}/transactions?address=0x17e4c16605e32adead5fa371bf6117df34ca0200"
/// curl "https://{Tomochain rpc}/tokens?address=0x8b353021189375591723e7384262f45709a3c3dc"
/// curl "http://localhost:8437/v1/tomochain/0x17e4c16605e32adead5fa371bf6117df34ca0200"
/// curl "http://localhost:8437/v2/tomochain/tokens/0x8b353021189375591723e7384262f45709a3c3dc?Authorization=Bearer"

module.exports = {
    path: '/mock/tomochain-api/:command1?',
    template: function(params, query) {
        switch (params.command1) {
            case 'transactions':
                if (query.address === '0x17e4c16605e32adead5fa371bf6117df34ca0200') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Ftomochain-api%2Ftransactions%3Faddress%3D0x17e4c16605e32adead5fa371bf6117df34ca0200.json";
                    var json = require(fn);
                    return json;
                }
                break;

            case 'tokens':
                if (query.address === '0x8b353021189375591723e7384262f45709a3c3dc') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Ftomochain-api%2Ftokens%3Faddress%3D0x8b353021189375591723e7384262f45709a3c3dc.json";
                    var json = require(fn);
                    return json;
                }
                break;        
        }

        return {error: "Not implemented"};
    }
};
