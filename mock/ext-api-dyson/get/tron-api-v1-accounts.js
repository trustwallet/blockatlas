/// Tron API Mock
/// See:
/// curl "http://localhost:3347/mock/tron-api/v1/accounts/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB/transactions?token_id=&limit=25&order_by=block_timestamp,desc"
/// curl "http://localhost:3347/mock/tron-api/v1/accounts/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB"
/// curl "http://{Tron rpc}/v1/accounts/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB/transactions?token_id=&limit=25&order_by=block_timestamp,desc"
/// curl "http://{Tron rpc}/v1/accounts/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB"
/// curl "http://localhost:8437/v1/tron/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB"
/// curl "http://localhost:8437/v2/tron/tokens/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB?Authorization=Bearer"

module.exports = {
    path: "/mock/tron-api/v1/accounts/:address/:operation?",
    template: function(params, query) {
        if (params.operation === 'transactions') {
            if (params.address === 'TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB') {
                var fn = "../../ext-api-data/get/" +
                    "mock%2Ftron-api%2Fv1%2Faccounts%2FTFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB%2Ftransactions%3Flimit%3D25%26order_by%3Dblock_timestamp%2Cdesc%26token_id%3D.json";
                var json = require(fn);
                return json;
            }
        }

        if (typeof params.operation === 'undefined') {
            if (params.address === 'TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB') {
                var fn = "../../ext-api-data/get/" +
                    "mock%2Ftron-api%2Fv1%2Faccounts%2FTFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB.json";
                var json = require(fn);
                return json;
            }
        }

        return {error: "Not implemented"};
    }
};
