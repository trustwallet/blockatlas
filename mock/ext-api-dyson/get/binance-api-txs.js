/// Binance chain block explorer API Mock, txs
/// Return
/// - for address 'bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp' 2 txs, one of them multi
/// - for address 'bnb1mnf0g6zhy2rpy63w8ryc6yp77s0fmqfyzjtkvd' 1 tx
/// - empty list for other addresses
/// See:
/// curl "http://localhost:3347/mock/explorer-api/v1/txs?address=bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp&page=1&rows=20&txAsset=BNB&txType=TRANSFER"
/// curl "https://explorer.binance.org/api/v1/txs?address=bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp&page=1&rows=20&txAsset=BNB&txType=TRANSFER"
/// curl "http://localhost:8437/v1/binance/bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp"

module.exports = {
    path: '/mock/binance-api/v1/txs',
    template: function(params, query, body) {
        //console.log(query);
        var address = query['address'];
        //console.log('address', address);
        switch (address) {
            case 'bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fbinance-api%2Fv1%2Ftxs%3Faddress%3Dbnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp%26page%3D1%26rows%3D20%26txAsset%3DBNB%26txType%3DTRANSFER.json";
                var json = require(fn);
                return json;

            case 'bnb1mnf0g6zhy2rpy63w8ryc6yp77s0fmqfyzjtkvd':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fbinance-api%2Fv1%2Ftxs%3Faddress%3Dbnb1mnf0g6zhy2rpy63w8ryc6yp77s0fmqfyzjtkvd%26page%3D1%26rows%3D20%26txAsset%3DBNB%26txType%3DTRANSFER.json";
                var json = require(fn);
                return json;
        }

        // not found, address
        return {txNums: 0, txArray: []}
    }
};
