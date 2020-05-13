/// Mock for external Ravencoin API
/// See:
/// curl "http://{Ravencoin rpc}/v2/xpub/xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B?details=txs"
/// curl "http://localhost:3347/mock/ravencoin-api/v2/xpub/xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B?details=txs"
/// curl "http://localhost:8437/v1/ravencoin/xpub/xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B"

module.exports = {
    path: '/mock/ravencoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fravencoin-api%2Fv2%2Fxpub%2Fxpub6BrkWQHMnuGcvKowEn2hpvnZ41SiCsu38mgFThKU3nMzPUN9r9C26puf18rfVdHH3nDwSkeMgsjVniNDKUk5arxekekGpNyVLsWihYAfC5B%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
