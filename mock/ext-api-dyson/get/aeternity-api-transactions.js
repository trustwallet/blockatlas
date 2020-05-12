/// Aeternity API Mock
/// See:
/// curl "http://localhost:3347/mock/aeternity-api/middleware/transactions/account/ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb"
/// curl "https://mdw.aepps.com/middleware/transactions/account/ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb"
/// curl http://localhost:8437/v1/aeternity/ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb
module.exports = {
    path: "/mock/aeternity-api/middleware/transactions/:type/:address?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.type === 'account') {
            if (params.address === 'ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb') {
                var fn = "../../data/get/" +
                    "mock%2Faeternity-api%2Fmiddleware%2Ftransactions%2Faccount%2Fak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb.json";
                let json = require(fn);
                return json;
            }
        }

        return {error: "Not implemented"};
    }
};
