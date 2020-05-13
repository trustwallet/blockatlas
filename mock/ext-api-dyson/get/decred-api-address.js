/// Mock for external Decred API
/// See:
/// curl "http://{decred rpc}/api/v2/address/DsTxPUVFxXeNgu5fzozr4mTR4tqqMaKcvpY?details=txs"
/// curl "http://localhost:3347/mock/decred-api/v2/address/DsTxPUVFxXeNgu5fzozr4mTR4tqqMaKcvpY?details=txs"
/// curl "http://localhost:8437/v1/decred/address/DsTxPUVFxXeNgu5fzozr4mTR4tqqMaKcvpY"

module.exports = {
    path: '/mock/decred-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'DsTxPUVFxXeNgu5fzozr4mTR4tqqMaKcvpY':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fdecred-api%2Fv2%2Faddress%2FDsTxPUVFxXeNgu5fzozr4mTR4tqqMaKcvpY%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
