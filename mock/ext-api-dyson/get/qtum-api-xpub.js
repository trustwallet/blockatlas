/// Mock for external Qtum API
/// See:
/// curl "http://{qtum rpc}/v2/xpub/xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd?details=txs"
/// curl "http://localhost:3347/mock/qtum-api/v2/xpub/xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd?details=txs"
/// curl "http://localhost:8437/v1/qtum/xpub/xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd"

module.exports = {
    path: '/mock/qtum-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fqtum-api%2Fv2%2Fxpub%2Fxpub6CvFuU1yPwHjMekXqgEZjcQy22ZWiKgRUY6yAneNNyk1trZhV6ZBFSY8Vt2wygTXTVHBkfi4n823vm79yiw42w6xTL2UjKyh2W9V88sXoNd%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
