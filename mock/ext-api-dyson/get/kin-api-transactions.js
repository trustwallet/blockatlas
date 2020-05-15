/// Kin API Mock, transactions
/// See:
/// curl "http://localhost:3347/mock/kin-api/transactions/b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a"
/// curl "https://horizon-block-explorer.kininfrastructure.com/transactions/b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a"
/// curl http://localhost:8437/v1/kin/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX
module.exports = {
    path: "/mock/kin-api/transactions/:txid?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.txid === 'b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fkin-api%2Ftransactions%2Fb2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a%3F.json";
            var json = require(fn);
            return json;
        }
        if (params.txid === 'eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fkin-api%2Ftransactions%2Feb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"};
    }
};
