/// Binance chain block explorer API Mock, txs
/// Return
/// - two txs, one of them multi, for address 'bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp'
/// - empty list for other addresses
/// See:
/// curl "http://localhost:3000/binance-explorer-api/v1/txs?address=bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp&page=1&rows=20&txAsset=BNB&txType=TRANSFER"
/// curl "https://explorer.binance.org/api/v1/txs?address=bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp&page=1&rows=20&txAsset=BNB&txType=TRANSFER"
/// curl "http://localhost:8420/v1/binance/bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp"

module.exports = {
    path: '/binance-explorer-api/v1/txs',
    template: function(params, query, body) {
        if (query['address'] == 'bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp') {
            return {
                txNums: 2,
                txArray: [
                    {
                        txHash: "4A5E688755E3588B89C5F0325274430AA52786E527313739C542BC69436E18F4",
                        blockHeight: Number.parseFrom('63282272'),
                        txType: "TRANSFER",
                        timeStamp: Number.parseFrom('1579689046296'),
                        fromAddr: "bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp",
                        toAddr: "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23",
                        value: 0.5541546,
                        txAsset: "BNB",
                        txFee: Number.parseFrom('0.000375'),
                        txAge: 2351199,
                        code: 0,
                        log: "Msg 0: ",
                        confirmBlocks: 0,
                        memo: "102101832",
                        source: 0,
                        hasChildren: 0
                    },
                    {
                        txHash: "F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC",
                        blockHeight: Number.parseFrom('63280715'),
                        txType: "TRANSFER",
                        timeStamp: Number.parseFrom('1579688431580'),
                        txFee: Number.parseFrom('0.00060'),
                        txAge: 2351814,
                        code: 0,
                        log: "Msg 0: ",
                        confirmBlocks: 0,
                        memo: "Trust Wallet Redeem",
                        source: 0,
                        hasChildren: 1
                    }
                ]
            }
        }

        // not found, address
        return {
            txNums: 0,
            txArray: []
        }
    }
};
