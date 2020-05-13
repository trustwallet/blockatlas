/// Mock for external Tezos API, transactions
/// See
/// curl "https://api.tzstats.com/explorer/account/tz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8/op?limit=25&order=desc&type=transaction%2Cdelegation"
/// curl "http://localhost:3347/mock/tezos-api/account/tz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8/op?limit=25&order=desc&type=transaction%2Cdelegation"
/// curl "http://localhost:8437/v1/tezos/tz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8"

module.exports = {
    path: '/mock/tezos-api/account/:account/:op?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.account === 'tz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Ftezos-api%2Faccount%2Ftz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8%2Fop%3Flimit%3D25%26order%3Ddesc%26type%3Dtransaction%2Cdelegation.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"}
    }
};
