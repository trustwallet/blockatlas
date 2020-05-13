/// Mock for external Zcoin API
/// See:
/// curl "http://{Zcoin rpc}/v2/xpub/xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK?details=txs"
/// curl "http://localhost:3347/mock/zcoin-api/v2/xpub/xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK?details=txs"
/// curl "http://localhost:8437/v1/zcoin/xpub/xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK"

module.exports = {
    path: '/mock/zcoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fzcoin-api%2Fv2%2Fxpub%2Fxpub6Cgu6WtTyo99pRtTabwscog2ncj4BUbTWzk7bt7habdLYwgnXLEWH3TuR1789QSTPVsPjLMa2KQzHffyZHTkLQQyRxeEBmWHaETS2btF5fK%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
