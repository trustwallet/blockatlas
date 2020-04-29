/// Binance chain block explorer API Mock, tx
/// Returns:
/// - Multi-transaction transaction for a specific address
///   see http://localhost:3000/binance-api/v1/transactions?address=bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q&limit=25&startTime=1578585526000&txAsset=BNB&txType=TRANSFER
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
        // tx: [{
        //     txHash: "DDA939C475755B2A00123CDA37A5EC442F502EA2983014D04F647C3CB9FB57A0",
        //     blockHeight: parseFloat('75110907'),
        //     txType: "TRANSFER",
        //     timeStamp: "2020-03-17T13:14:30.939Z",
        //     fromAddr: "bnb166fvwtfra2atta5mm8mygt2fyltktqctgpedcg",
        //     toAddr: "bnb1mnf0g6zhy2rpy63w8ryc6yp77s0fmqfyzjtkvd",
        //     value: parseFloat('0.00037500'),
        //     txAsset: "BNB",
        //     txFee: parseFloat('0.00060'),
        //     proposalId: null,
        //     txAge: parseFloat('1997110'),
        //     orderId: null,
        //     code: parseFloat('0'),
        //     data: null,
        //     confirmBlocks: parseFloat('0'),
        //     memo: "",
        //     source: parseFloat('2'),
        //     sequence: parseFloat('946')
        // }
        // ],
        // total: 1
    }
};
