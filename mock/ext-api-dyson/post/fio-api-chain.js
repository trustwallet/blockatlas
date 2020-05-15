/// FIO RPC API Mock, chain API
/// Returns:
/// - public address for certain fio name and coin combinations
/// - public address not found message for other input
/// See:
/// curl -H "Content-Type: application/json" -d '{"fio_address":"trust@trust","token_code":"BTC","chain_code":"BTC"}' "http://localhost:3347/mock/fio-api/v1/chain/get_pub_address"
/// curl -H "Content-Type: application/json" -d '{"fio_address":"trust@trust","token_code":"BTC","chain_code":"BTC"}' "https://fio.greymass.com/v1/chain/get_pub_address"
/// curl "http://localhost:8437/v2/ns/lookup?name=trust@trust&coins=60"

module.exports = {
    path: '/mock/fio-api/v1/chain/:action',
    template: function(params, query, body) {
        if (params.action === 'get_pub_address') {
            switch (body.fio_address) {
                case 'trust@trust':
                    switch (body.token_code) {
                        case 'ETH':
                            var fn = "../../ext-api-data/post/" +
                                "mock%2Ffio-api%2Fv1%2Fchain%2Fget_pub_address.json";
                            var json = require(fn);
                            return json;
                    }
                    break;
        
                case 'trust@trustwallet':
                    var fn = "../../ext-api-data/post/" +
                        "mock%2Ffio-api%2Fv1%2Fchain%2Fget_pub_address.0001.json";
                    var json = require(fn);
                    return json;
        
                case 'name@somefiodomain':
                    var fn = "../../ext-api-data/post/" +
                        "mock%2Ffio-api%2Fv1%2Fchain%2Fget_pub_address.0002.json";
                    var json = require(fn);
                    return json;
            }
        }
        return {error: "Not implemented"};
    }
};
