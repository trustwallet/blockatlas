/// Waves API Mock
/// See:
/// curl "http://localhost:3347/mock/waves-api/transactions/address/3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD/limit/25"
/// curl "https://nodes.wavesnodes.com/transactions/address/3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD/limit/25"
/// curl http://localhost:8437/v1/waves/3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD
module.exports = {
    path: "/mock/waves-api/transactions/address/:address/limit/:limit",
    template: function (params, query) {
        //console.log(params)
        if (params.address === '3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fwaves-api%2Ftransactions%2Faddress%2F3PJ4q4sqriJs2y7Z45wmbLrbmV9MDecbPxD%2Flimit%2F25.json";
            var json = require(fn);
            return json;
        }

        return {error: "Not implemented"};
    }
}
;
