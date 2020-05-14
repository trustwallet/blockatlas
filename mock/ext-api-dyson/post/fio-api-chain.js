/// FIO RPC API Mock, chain API
/// Returns:
/// - public address for certain fio name and coin combinations
/// - public address not found message for other input
/// See:
/// curl -H "Content-Type: application/json" -d '{"fio_address":"trust@trust","token_code":"BTC","chain_code":"BTC"}' "http://localhost:3347/mock/fio-api/v1/chain/get_pub_address"
/// curl -H "Content-Type: application/json" -d '{"fio_address":"trust@trust","token_code":"BTC","chain_code":"BTC"}' "http://testnet.fioprotocol.io/v1/chain/get_pub_address"
/// curl "http://localhost:8437/v2/ns/lookup?name=trust@trust&coins=60"

module.exports = {
    path: '/mock/fio-api/v1/chain/:action',
    template: function(params, query, body) {
        if (params.action === 'get_pub_address') {
            switch (body.fio_address) {
                case 'trust@trust':
                    switch (body.token_code) {
                        case 'BTC':
                            return {public_address: 'bc1qvy4074rggkdr2pzw5vpnn62eg0smzlxwp70d7v'};
                        case 'ETH':
                            return {public_address: '0xce5cB6c92Da37bbBa91Bd40D4C9D4D724A3a8F51'};
                        case 'BNB':
                            return {public_address: 'bnb1ts3dg54apwlvr9hupv2n0j6e46q54znnusjk9s'};
                    }
                    break;
        
                case 'trust@trustwallet':
                    return {public_address: '0xce5cB6c92Da37bbBa91Bd40D4C9D4D724A3a0001'};
        
                case 'name@somefiodomain':
                    return {public_address: '0xce5cB6c92Da37bbBa91Bd40D4C9D4D724A3a0002'};
            }
            return {message: 'Public address not found'};
        }
        return {error: "Not implemented"};
    }
};
