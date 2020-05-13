/// Mock for external Viacoin API
/// See:
/// curl "http://{Viacoin rpc}/v2/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A?details=txs"
/// curl "http://localhost:3347/mock/viacoin-api/v2/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A?details=txs"
/// curl "http://localhost:8437/v1/viacoin/address/VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A"

module.exports = {
    path: '/mock/viacoin-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fviacoin-api%2Fv2%2Faddress%2FVdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
