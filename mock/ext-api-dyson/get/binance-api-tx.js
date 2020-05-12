/// Binance chain block explorer API Mock, tx
/// Returns:
/// - Multi-transaction transaction for a specific address
///   see http://localhost:3347/mock/binance-api/v1/tx?txHash=F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC
///   see https://explorer.binance.org/api/v1/tx?txHash=F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC
///   see http://localhost:8437/v1/binance/bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp
///   see http://localhost:8437/v1/binance/bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp?token=BUSD-BD1
/// - empty response for other txHash'es

module.exports = {
    path: '/mock/binance-api/v1/tx',
    template: function(params, query, body) {
        if (query['txHash'] == 'F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC') {
            var fn = "../../ext-api-data/get/" +
                "get/mock%2Fbinance-api%2Fv1%2Ftx%3FtxHash%3DF53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC.json";
            let json = require(fn);
            return json;
        }

        // not found txHash, return empty response
        return {}
    }
};
