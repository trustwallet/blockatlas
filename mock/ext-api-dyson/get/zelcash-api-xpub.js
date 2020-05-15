/// Mock for external Zelcash API
/// See:
/// curl "http://{Zelcash rpc}/v2/xpub/xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf?details=txs"
/// curl "http://localhost:3347/mock/zelcash-api/v2/xpub/xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf?details=txs"
/// curl "http://localhost:8437/v1/zelcash/xpub/xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf"

module.exports = {
    path: '/mock/zelcash-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fzelcash-api%2Fv2%2Fxpub%2Fxpub6C5soBeFd2uZLCcEvsqaoGXuh9UposMMfk2jSiBKMN8rJKs9NLqjPK51gWv9mYBpUY95GtHYsofwpPRdB6FJ56cEaTJGCba5GKv55wPNZNf%3Fdetails%3Dtxs";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
