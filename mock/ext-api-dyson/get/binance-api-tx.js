/// Binance chain block explorer API Mock, tx
/// Returns:
/// - Multi-transaction transaction for a specific address
///   see http://localhost:3000/binance-api/v1/tx?txHash=F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC
///   see https://explorer.binance.org/api/v1/tx?txHash=F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC
///   see http://localhost:8420/v1/binance/bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp
///   see http://localhost:8420/v1/binance/bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp?token=BUSD-BD1
/// - empty response for other txHash'es

module.exports = {
    path: '/binance-api/v1/tx',
    template: function(params, query, body) {
        if (query['txHash'] == 'F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC') {
            return {
                txHash: "F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC",
                blockHeight: Number.parseFrom('63280715'),
                txType: "TRANSFER",
                timeStamp: Number.parseFrom('1579688431580'),
                txFee: Number.parseFrom('0.00060'),
                txAge: Number.parseFrom('2350509'),
                code: 0,
                log: "Msg 0: ",
                confirmBlocks: 5818526,
                memo: "Trust Wallet Redeem",
                source: 0,
                sequence: 175,
                hasChildren: 1,
                subTxsDto: {
                    totalNum: 2,
                    pageSize: 15,
                    subTxDtoList: [
                        {
                            hash: "F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC",
                            height: Number.parseFrom('63280715'),
                            type: "TRANSFER",
                            value: 0.00375,
                            asset: "BNB",
                            fromAddr: "bnb1rhv98jcx2yu26shxedskttjzpkvsrz4nd226yv",
                            toAddr: "bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp",
                            fee: Number.parseFrom('0.00060')
                        },
                        {
                            hash: "F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC",
                            height: Number.parseFrom('63280715'),
                            type: "TRANSFER",
                            value: 10.0,
                            asset: "BUSD-BD1",
                            fromAddr: "bnb1rhv98jcx2yu26shxedskttjzpkvsrz4nd226yv",
                            toAddr: "bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp",
                            fee: null
                        }
                    ]
                }
            };
        }

        // not found txHash, return empty response
        return {}
    }
};
