/// Kin API Mock, accounts
/// See:
/// curl "http://localhost:3347/mock/kin-api/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH/payments?order=desc&limit=25"
/// curl "https://horizon-block-explorer.kininfrastructure.com/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH/payments?order=desc&limit=25"
/// curl http://localhost:8437/v1/kin/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH
module.exports = {
    path: "/mock/kin-api/accounts/:address/:operation?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.operation === 'payments') {
            if (params.address === 'GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH') {
              var fn = "../../ext-api-data/get/" +
                "mock%2Fkin-api%2Faccounts%2FGBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH%2Fpayments%3Flimit%3D25%26order%3Ddesc.json";
              var json = require(fn);
              return json;
            }
        }
        
        return {error: "Not implemented"};
    }
};
