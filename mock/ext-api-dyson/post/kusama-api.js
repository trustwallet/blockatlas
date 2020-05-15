/// Kusama RPC API Mock
/// See:
/// curl -H "Content-Type: application/json" -d '{"address": "HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK", "row": 25}' https://kusama.subscan.io/api/scan/transfers
/// curl -H "Content-Type: application/json" -d '{"address": "HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK", "row": 25}' http://localhost:3347/mock/kusama-rpc/scan/transfers
/// curl "http://localhost:8437/v1/kusama/HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK"

module.exports = {
    path: '/mock/kusama-rpc/scan/transfers',
    template: function(params, query, body) {
        switch (body.address) {
            case 'HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK':
                var fn = "../../ext-api-data/post/" +
                    "mock%2Fkusama-api%2Fscan%2Ftransfers.json";
                var json = require(fn);
                return json;

        }
        return {error: "Not implemented"};
    }
};
