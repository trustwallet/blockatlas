/// Mock for external Qtum API
/// See:
/// curl "http://{qtum rpc}/api/v2/address/QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ?details=txs"
/// curl "http://localhost:3347/mock/qtum-api/v2/address/QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ?details=txs"
/// curl "http://localhost:8437/v1/qtum/address/QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ"

module.exports = {
    path: '/mock/qtum-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'QZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fqtum-api%2Fv2%2Faddress%2FQZJbNrGT3cZ1J1AEHtgH3JWM7uLBNAejLZ%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
