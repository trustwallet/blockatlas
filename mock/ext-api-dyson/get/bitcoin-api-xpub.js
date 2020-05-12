/// Mock for external Bitcoin API
/// See:
/// curl "https://btc1.trezor.io/api/v2/xpub/zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC?details=txs"
/// curl "http://localhost:3347/mock/bitcoin-api/v2/xpub/zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC?details=txs"
/// curl "http://localhost:8437/v1/bitcoin/xpub/zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC"

module.exports = {
    path: '/mock/bitcoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        switch (params.xpubkey) {
            case 'zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fbitcoin-api%2Fv2%2Fxpub%2Fzpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
