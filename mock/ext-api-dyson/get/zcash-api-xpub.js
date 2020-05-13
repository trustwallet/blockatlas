/// Mock for external Zcash API
/// See:
/// curl "http://{Zcash rpc}/v2/xpub/xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS?details=txs"
/// curl "http://localhost:3347/mock/zcash-api/v2/xpub/xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS?details=txs"
/// curl "http://localhost:8437/v1/zcash/xpub/xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS"

module.exports = {
    path: '/mock/zcash-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fzcash-api%2Fv2%2Fxpub%2Fxpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
