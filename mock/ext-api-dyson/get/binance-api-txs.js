/// Binance chain block explorer API Mock, tx
/// Returns:
/// - Multi-transaction transaction for a specific address
///   see http://localhost:3000/binance-rpc/v1/transactions?address=bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q&limit=25&startTime=1578585526000&txAsset=BNB&txType=TRANSFER
///   see https://dex.binance.org/api/v1/transactions?address=bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q&limit=25&startTime=1578585526000&txAsset=BNB&txType=TRANSFER
///   see http://localhost:8420/v1/binance/bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q
/// - empty response for other txHash'es

module.exports = {
    path: '/binance-rpc/v1/transactions',
    template: function (params, query, body) {
        return JSON.parse(`{
            "tx": [{
                "txHash": "91AFC91FCEFC1C556A79A25F640622F379F75F2CD23FD2A806B27BE049EFA828",
                "blockHeight": 79462376,
                "txType": "TRANSFER",
                "timeStamp": "2020-04-06T14:35:03.436Z",
                "fromAddr": "bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q",
                "toAddr": "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23",
                "value": "500.00000000",
                "txAsset": "BNB",
                "txFee": "0.00037500",
                "proposalId": null,
                "txAge": 266311,
                "orderId": null,
                "code": 0,
                "data": null,
                "confirmBlocks": 0,
                "memo": "109830695",
                "source": 2,
                "sequence": 68
            }], "total": 1
        }`)
    }
};
