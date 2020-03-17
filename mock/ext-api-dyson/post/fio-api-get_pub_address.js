/// FIO RPC API Mock, get_pub_address
/// Returns:
/// - public address for certain fio name and coin combinations
/// - public address not found message for other input
/// See:
/// curl -H "Content-Type: application/json" -d '{"fio_address":"adam@fiotestnet","token_code":"BTC","chain_code":"BTC"}' "http://localhost:3000/fio-api/v1/chain/get_pub_address"
/// curl -H "Content-Type: application/json" -d '{"fio_address":"adam@fiotestnet","token_code":"BTC","chain_code":"BTC"}' "http://testnet.fioprotocol.io/v1/chain/get_pub_address"
/// curl "http://localhost:8420/v2/ns/lookup?name=adam@fiotestnet&coins=60"

module.exports = {
    path: '/fio-api/v1/chain/get_pub_address',
    template: function(params, query, body) {
        var addr = getAddress(body.fio_address, body.token_code);
        if (addr == '') {
            return {message: 'Public address not found'};
        }
        return {public_address: addr};
    }
};

function getAddress(fio_address, token_code) {
    if (fio_address !== 'adam@fiotestnet') { return ''; }
    switch (token_code) {
        case 'BTC': return 'bc1qvy4074rggkdr2pzw5vpnn62eg0smzlxwp70d7v';
        case 'ETH': return '0xce5cB6c92Da37bbBa91Bd40D4C9D4D724A3a8F51';
        case 'BNB': return 'bnb1ts3dg54apwlvr9hupv2n0j6e46q54znnusjk9s';
        default: return '';
    }
};
