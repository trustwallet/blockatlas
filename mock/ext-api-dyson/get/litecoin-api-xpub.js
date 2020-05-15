/// Mock for external Litecoin API
/// See:
/// curl "http://{ltc rpc}/v2/xpub/zpub6rpF5Uxuz4KKWePLorSz2QrHMmk1iiZvGUGgtSHpor8yiGekyRuWf5ZNmf6GUKB4v3ibQDuZp5v8RnjEGq58kR3WPtGPn8Lrg677MQ8YeKu?details=txs"
/// curl "http://localhost:3347/mock/litecoin-api/v2/xpub/zpub6rpF5Uxuz4KKWePLorSz2QrHMmk1iiZvGUGgtSHpor8yiGekyRuWf5ZNmf6GUKB4v3ibQDuZp5v8RnjEGq58kR3WPtGPn8Lrg677MQ8YeKu?details=txs"
/// curl "http://localhost:8437/v1/litecoin/xpub/zpub6rpF5Uxuz4KKWePLorSz2QrHMmk1iiZvGUGgtSHpor8yiGekyRuWf5ZNmf6GUKB4v3ibQDuZp5v8RnjEGq58kR3WPtGPn8Lrg677MQ8YeKu"

module.exports = {
    path: '/mock/litecoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'zpub6rpF5Uxuz4KKWePLorSz2QrHMmk1iiZvGUGgtSHpor8yiGekyRuWf5ZNmf6GUKB4v3ibQDuZp5v8RnjEGq58kR3WPtGPn8Lrg677MQ8YeKu':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Flitecoin-api%2Fv2%2Fxpub%2Fzpub6rpF5Uxuz4KKWePLorSz2QrHMmk1iiZvGUGgtSHpor8yiGekyRuWf5ZNmf6GUKB4v3ibQDuZp5v8RnjEGq58kR3WPtGPn8Lrg677MQ8YeKu%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
