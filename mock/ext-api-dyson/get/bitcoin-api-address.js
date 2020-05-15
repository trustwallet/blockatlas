/// Mock for external Bitcoin API
/// See:
/// curl "http://localhost:3347/mock/bitcoin-api/v2/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj?details=txs"
/// curl "https://btc1.trezor.io/api/v2/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj?details=txs"
/// curl "http://localhost:8437/v1/bitcoin/address/bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj"

module.exports = {
    path: '/mock/bitcoin-api/v2/address/:address?',
    template: function (params, query, body) {
        switch (params.address) {
            case 'bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fbitcoin-api%2Fv2%2Faddress%2Fbc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return { error: "Not implemented" };
    }
}
