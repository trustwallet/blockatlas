/// Mock for external Doge API
/// See:
/// curl "http://{doge rpc}/v2/xpub/dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN?details=txs"
/// curl "http://localhost:3347/mock/doge-api/v2/xpub/dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN?details=txs"
/// curl "http://localhost:8437/v1/doge/xpub/dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN"

module.exports = {
    path: '/mock/doge-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'dgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fdoge-api%2Fv2%2Fxpub%2Fdgub8rceyfsEvGDexmvJcBqiKBrmuxWGgYJxHjtbouHTwTfQrCQcMjxyNf6vUPY4dUp23QtReFy6WGedutBk9XUaYNupUqVAZcweqGhfsudUELN%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
