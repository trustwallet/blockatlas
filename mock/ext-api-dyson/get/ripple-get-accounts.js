/// Ripple API Mock
/// See:
/// curl "http://localhost:3347/mock/ripple-api/accounts/rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1/transactions?type=Payment&descending=false&limit=25"
/// curl "https://data.ripple.com/v2/accounts/rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1/transactions?type=Payment&descending=false&limit=25"
/// curl http://localhost:8437/v1/ripple/rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1
module.exports = {
    path: "/mock/ripple-api/accounts/:address/transactions?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.address === 'rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fripple-api%2Faccounts%2FrMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1%2Ftransactions%3Fdescending%3Dfalse%26limit%3D25%26type%3DPayment.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"};
    }
};
