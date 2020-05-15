/// Stellar API Mock, transactions
/// See:
/// curl "http://localhost:3347/mock/stellar-api/transactions/2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029?"
/// curl "https://horizon.stellar.org/transactions/2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029?"
/// curl http://localhost:8437/v1/stellar/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX
module.exports = {
    path: "/mock/stellar-api/transactions/:txid?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.txid === '2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fstellar-api%2Ftransactions%2F2912d519b2c2174b0147a9e02208f3ed14820228913142f8c6a5cd360783c029%3F.json";
            var json = require(fn);
            return json;
        }
        if (params.txid === '23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fstellar-api%2Ftransactions%2F23fa4d3c787f22a1755afa4275761ffba516ec4f2e039440ec02a5522b388f2c.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"};
    }
};
