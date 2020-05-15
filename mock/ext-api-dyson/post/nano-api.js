/// Nano RPC Mock
/// curl -H 'Content-Type: application/json' -d '{"action":"account_history","account":"nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z","count":"25"}' http://localhost:3347/mock/nano-api
/// curl -H 'Content-Type: application/json' -d '{"action":"account_history","account":"nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z","count":"25"}' https://{nano_rpc}
/// curl "http://localhost:8437/v1/nano/nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z"

module.exports = {
    path: '/mock/nano-api',
    template: function(params, query, body) {
        //console.log("curl -H 'Content-Type: application/json' -d '", JSON.stringify(body), "' https://{nano_rpc}");
        if (body.action === 'account_history') {
            //console.log('body.account', body.account);
            if (body.account === 'nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z') {
                var fn = "../../ext-api-data/post/" +
                    "mock%2Fnano-api.json";
                var json = require(fn);
                return json;
            }
        }

        return {error: "Not implemented"};
    }
};
