/// Mock for external Bitcoincash API
/// See:
/// curl "https://{bch rpc}/v2/xpub/xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX?details=txs"
/// curl "http://localhost:3347/mock/bitcoincash-api/v2/xpub/xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX?details=txs"
/// curl "http://localhost:8437/v1/bitcoincash/xpub/xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX"

module.exports = {
    path: '/mock/bitcoincash-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX':
                var fn = "../../data/get/" +
                    "mock%2Fbitcoincash-api%2Fv2%2Fxpub%2Fxpub6Bq3UUphocwroXkhA9sn8ACnZpJNuwaBehgo7WbDi2DULYnvT72Uzgsv9cE5EiP8ThDYdMyZREfbpkUY4KZ88ZaUQxXciBcZ1soSi1d8xtX%3Fdetails%3Dtxs.json";
                let json = require(fn);
                return json;
        }

        return {error: "Not implemented"};
    }
}
