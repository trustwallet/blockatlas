/// Aion API Mock
/// See:
/// curl "http://localhost:3347/mock/aion-api/getTransactionsByAddress?accountAddress=0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed&size=25"
/// curl "https://mainnet-api.theoan.com/aion/dashboard/getTransactionsByAddress?accountAddress=0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed&size=25"
/// curl http://localhost:8437/v1/aion/0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed
module.exports = {
    path: "/mock/aion-api/getTransactionsByAddress?",
    template: function(params, query, body) {
        //console.log(query)
        if (query.accountAddress === '0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Faion-api%2FgetTransactionsByAddress%3FaccountAddress%3D0xa04f0117864ccf5013861a89f08c6fc790284d72356c8a362025d31b855ed6ed%26size%3D25.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"};
    }
};
