/// Mock for external Dash API
/// See:
/// curl "http://{dash rpc}/v2/xpub/xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD?details=txs"
/// curl "http://localhost:3347/mock/dash-api/v2/xpub/xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD?details=txs"
/// curl "http://localhost:8437/v1/dash/xpub/xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD"

module.exports = {
    path: '/mock/dash-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fdash-api%2Fv2%2Fxpub%2Fxpub6CKAjCUKKPW7bzYEG5mzRsmzyTRp7XzauqFWNmpGVNqMqsSQpLMCN3ygEmD6ZEGVocNDrDhE7SeGot78noEWpwPDbJjfxREHC848sxNrUkD%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
