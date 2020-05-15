/// Stellar API Mock, accounts
/// See:
/// curl "http://localhost:3347/mock/stellar-api/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX/payments?order=desc&limit=25"
/// curl "https://horizon.stellar.org/accounts/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX/payments?order=desc&limit=25"
/// curl http://localhost:8437/v1/stellar/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX
module.exports = {
    path: "/mock/stellar-api/accounts/:address/:operation?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.operation === 'payments') {
            if (params.address === 'GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX') {
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fstellar-api%2Faccounts%2FGDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX%2Fpayments%3Flimit%3D25%26order%3Ddesc.json";
                var json = require(fn);
                return json;
            }
        }
        
        return {error: "Not implemented"};
    }
};
