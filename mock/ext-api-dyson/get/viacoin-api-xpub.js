/// Mock for external Viacoin API
/// See:
/// curl "http://{Viacoin rpc}/v2/xpub/zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK?details=txs"
/// curl "http://localhost:3347/mock/viacoin-api/v2/xpub/zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK?details=txs"
/// curl "http://localhost:8437/v1/viacoin/xpub/zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK"

module.exports = {
    path: '/mock/viacoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fviacoin-api%2Fv2%2Fxpub%2Fzpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
