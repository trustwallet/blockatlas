/// FIO RPC API Mock, chain API
/// Returns:
/// - public address for certain fio name and coin combinations
/// - public address not found message for other input
/// See:
/// curl -H "Content-Type: application/json" -d '{"fio_address":"trust@trust","token_code":"BTC","chain_code":"BTC"}' "http://localhost:3000/fio-api/v1/chain/get_pub_address"
/// curl -H "Content-Type: application/json" -d '{"fio_address":"trust@trust","token_code":"BTC","chain_code":"BTC"}' "http://testnet.fioprotocol.io/v1/chain/get_pub_address"
/// curl "http://localhost:8420/v2/ns/lookup?name=trust@trust&coins=60"

module.exports = {
    path: '/fio-api/v1/chain/:action',
    template: function(params, query, body) {
        if (params.action === 'get_pub_address') {
            var addr = getAddress(body.fio_address, body.token_code);
            if (addr == '') {
                return {message: 'Public address not found'};
            }
            return {public_address: addr};
        }
        return {error: 'Not implemented'};
    }
};

function getAddress(fio_address, token_code) {
    switch (fio_address) {
        case 'trust@trust':
            switch (token_code) {
                case 'BTC': return 'bc1qvy4074rggkdr2pzw5vpnn62eg0smzlxwp70d7v';
                case 'ETH': return '0xce5cB6c92Da37bbBa91Bd40D4C9D4D724A3a8F51';
                case 'BNB': return 'bnb1ts3dg54apwlvr9hupv2n0j6e46q54znnusjk9s';
                default: return '';
            }

        case 'trust@trustwallet':
            return '0xce5cB6c92Da37bbBa91Bd40D4C9D4D724A3a0001';

        case 'name@somefiodomain':
            return '0xce5cB6c92Da37bbBa91Bd40D4C9D4D724A3a0002';
    }
    return '';
};
