
/// Unstoppabledomains API Mock, lookup
/// Returns:
/// - public address for certain name and coin combinations
/// - public address not found message for other input
/// See:
/// curl "http://localhost:3000/unstoppabledomains/api/v1//dpantani.zil"
/// curl "https://unstoppabledomains.com/api/v1/dpantani.zil"
/// curl "http://localhost:8420/v2/ns/lookup?name=dpantani.zil&coins=313"

module.exports = {
    path: '/unstoppabledomains/api/v1//:name?',
    template: function(params, query, body) { return lookup(params, query, body); }
};

function lookup(params, query, body) {
    var addr = getAddresses(params.name);
    if (addr == '') {
        return {addresses: {}, meta: {owner: null, type: "ZNS", ttl: 0}, claimed: false};
    }
    return addr;
}

function getAddresses(address) {
    if (address !== 'dpantani.zil' && address !== 'dpantani.crypto') { return ''; }
    return {
        addresses: {
            BTC: "bc1qd7eystu9xl53hkyxm4kyg7h5yk4p436sqx6f27",
            ETH: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
            ZIL: "zil1vdntvlk47j9kh9a85klqcd9rvgze06ruhmna64"
        },
        meta: {
            owner: "0xe4da405976f315ab91ff9a43a51972cc94739aa8",
            type: "ZNS",
            ttl: 0
        }
    }
};
