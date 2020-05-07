/// Binance chain block explorer mock transaction
/// Returns:
/// - Multi-transaction transaction for a specific address
///   1. newman calls: http://localhost:8437/v1/binance/bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q
///   2. Block Atlas calls internally : https://{binance_explorer}/api/v1/txs?address=bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q&page=1&rows=25&txType=TRANSFER
///   3. Dyson response mock data: http://localhost:3347/binance-explorer/api/v1/txs?address=bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q&page=1&rows=25&txType=TRANSFER
/// - empty response for other txHash'es

// Example contains each type, BNB transfer, BEP2 transfer, BEP2 transfer as multisend transaction
module.exports = {
    path: '/binance-explorer/api/v1/txs',
    template: function (params, query, body) {
        return JSON.parse(`{
    "txNums": 3,
    "txArray": [
        {
            "txHash": "963D8C627DE1E739845C0AC0C0EAC3387B53806C77289ACB8C1653CD6D62C9CC",
            "blockHeight": 84191249,
            "txType": "TRANSFER",
            "timeStamp": 1588086370574,
            "fromAddr": "bnb14cjy0yl23xkf0hnw3ql295v8nghqstvlzkvqpl",
            "toAddr": "bnb1mr5f97rx5wnkfcakx9fcpvljmx2s6kwqc08yur",
            "value": 2800.00000000,
            "txAsset": "TWT-8C2",
            "txFee": 0.00037500,
            "txAge": 446701,
            "code": 0,
            "log": "Msg 0: ",
            "confirmBlocks": 0,
            "memo": "",
            "source": 0,
            "hasChildren": 0
        },
        {
            "txHash": "4577CB3B5B202696E9E0B093A6DA973C7DD9CBC6808DA1326872745C35F3C089",
            "blockHeight": 84191216,
            "txType": "TRANSFER",
            "timeStamp": 1588086357686,
            "fromAddr": "bnb1mr5f97rx5wnkfcakx9fcpvljmx2s6kwqc08yur",
            "toAddr": "bnb14cjy0yl23xkf0hnw3ql295v8nghqstvlzkvqpl",
            "value": 0.00040000,
            "txAsset": "BNB",
            "txFee": 0.00037500,
            "txAge": 446714,
            "code": 0,
            "log": "Msg 0: ",
            "confirmBlocks": 0,
            "memo": "",
            "source": 0,
            "hasChildren": 0
        },
        {
            "txHash": "FAD8C1C5E450BE5E0913B12007AAEACC307F8CFFAFFB0844A9F83155E1235C25",
            "blockHeight": 80167666,
            "txType": "TRANSFER",
            "timeStamp": 1586464452922,
            "txFee": 0.29970000,
            "txAge": 2068619,
            "code": 0,
            "log": "Msg 0: ",
            "confirmBlocks": 0,
            "memo": "",
            "source": 0,
            "hasChildren": 1
        }
    ]
}`)
    }
};
