/// Mock for external Groestlcoin API
/// See:
/// curl "http://{groestlcoin rpc}/v2/xpub/zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf?details=txs"
/// curl "http://localhost:3347/mock/groestlcoin-api/v2/xpub/zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf?details=txs"
/// curl "http://localhost:8437/v1/groestlcoin/xpub/zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf"

module.exports = {
    path: '/mock/groestlcoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'zpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fgroestlcoin-api%2Fv2%2Fxpub%2Fzpub6rWUMiiVPxjWVHffT8x3AfcbyDu8SZJAiuKUTBmhxT7Bvqk1WitxndDStG1qHN6XzRM7JgsaRaVccRFW3AprWk4Fpaev1N6QSp1aNnP5JPf%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
