/// Kava API Mock
/// See:
/// curl "http://localhost:3347/mock/kava-api/txs?transfer.recipient=kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m&page=1&limit=25"
/// curl "http://localhost:3347/mock/kava-api/txs?message.sender=kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m&page=1&limit=25"
/// curl "http://{kava rpc}/txs?transfer.recipient=kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m&page=1&limit=25"
/// curl "http://{kava rpc}/txs?message.sender=kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m&page=1&limit=25"
/// curl http://localhost:8437/v1/kava/kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m
module.exports = {
    path: "/mock/kava-api/txs?",
    template: function(params, query, body) {
        //console.log(query)
        if (query["transfer.recipient"] === 'kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fkava-api%2Ftxs%3Flimit%3D25%26page%3D1%26transfer.recipient%3Dkava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m.json";
            var json = require(fn);
            return json;
        }

        if (query["message.sender"] === 'kava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fkava-api%2Ftxs%3Flimit%3D25%26message.sender%3Dkava1l8va9zyl50cpzv447c694k3jndelc9ygtfll2m%26page%3D1.json";
            var json = require(fn);
            return json;
        }

        return {error: "Not implemented"};
    }
};

