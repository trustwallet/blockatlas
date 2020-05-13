/// Mock for external Digibyte API
/// See:
/// curl "http://{digibyte rpc}/v2/xpub/zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow?details=txs"
/// curl "http://localhost:3347/mock/digibyte-api/v2/xpub/zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow?details=txs"
/// curl "http://localhost:8437/v1/digibyte/xpub/zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow"

module.exports = {
    path: '/mock/digibyte-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'zpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fdigibyte-api%2Fv2%2Fxpub%2Fzpub6ricE56nzsDeTAVo4w68vQQ3tRvR6C18JjKVsgbiRFjEawGV9SuS2gfkpm5qFxjbTNPPuvAA3cqRsxNHxFwVnpYD2Lawjtb3wowbFdwmjow%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
