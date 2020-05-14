/// FIO RPC API Mock, history API
/// curl -H "Content-Type: application/json" -d '{"account_name": "ezsmbcy2opod"}' https://fio.greymass.com/v1/history/get_actions
/// curl -H "Content-Type: application/json" -d '{"account_name": "ezsmbcy2opod"}' http://localhost:3347/mock/fio-api/v1/history/get_actions
/// curl "http://localhost:8437/v1/fio/FIO7Q3XfQ2ocGP1zYst6Sfx5qrsiZ865Cu8So2atrub9JN94so7gt"

module.exports = {
    path: '/mock/fio-api/v1/history/:action',
    template: function(params, query, body) {
        if (params.action === 'get_actions') {
            switch (body.account_name) {
                case 'ezsmbcy2opod': // FIO7Q3XfQ2ocGP1zYst6Sfx5qrsiZ865Cu8So2atrub9JN94so7gt
                    var fn = "../../ext-api-data/post/" +
                        "mock%2Ffio-api%2Fv1%2Fhistory%2Fget_actions.json";
                    var json = require(fn);
                    return json;

                case 'gmdncuvoqxfn': // FIO6gZthsHigy7wXeev4MKS4MuoygkxQ1yirmmUqpoubDWLJTASa2
                    var fn = "../../ext-api-data/post/" +
                        "mock%2Ffio-api%2Fv1%2Fhistory%2Fget_actions.0001.json";
                    var json = require(fn);
                    return json;

                case 'fio.treasury':
                    var fn = "../../ext-api-data/post/" +
                        "mock%2Ffio-api%2Fv1%2Fhistory%2Fget_actions.0002.json";
                    var json = require(fn);
                    return json;
            }
        }
        return {error: 'Not implemented'};
    }
};
