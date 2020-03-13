/// Binance chain block explorer API Mock, txs
/// Return
/// - for address 'bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp' 2 txs, one of them multi
/// - for address 'bnb1mnf0g6zhy2rpy63w8ryc6yp77s0fmqfyzjtkvd' 1 tx
/// - empty list for other addresses
/// See:
/// curl "http://localhost:3000/explorer-api/v1/txs?address=bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp&page=1&rows=20&txAsset=BNB&txType=TRANSFER"
/// curl "https://explorer.binance.org/api/v1/txs?address=bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp&page=1&rows=20&txAsset=BNB&txType=TRANSFER"
/// curl "http://localhost:8420/v1/binance/bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp"

module.exports = {
    path: '/binance-api/v1/txs',
    template: function(params, query, body) {
        //console.log(query);
        var address = query['address'];
        //console.log('address', address);
        switch (address) {
            case 'bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp':
                return JSON.parse(`{
                    "txNums": 2,
                    "txArray": [
                        {
                            "txHash": "4A5E688755E3588B89C5F0325274430AA52786E527313739C542BC69436E18F4",
                            "blockHeight": 63282272,
                            "txType": "TRANSFER",
                            "timeStamp": 1579689046296,
                            "fromAddr": "bnb1563k58pc3keeuwkhlrxwz7sdsetyn9l7gdnznp",
                            "toAddr": "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23",
                            "value": 0.55415460,
                            "txAsset": "BNB",
                            "txFee": 0.00037500,
                            "txAge": 5273327,
                            "code": 0,
                            "log": "Msg 0: ",
                            "confirmBlocks": 0,
                            "memo": "102101832",
                            "source": 0,
                            "hasChildren": 0
                        },
                        {
                            "txHash": "F53BB470A3B6B83977CFFE5D5F9937FB1CBB8785FBE818D9B38AD43F3ECD82BC",
                            "blockHeight": 63280715,
                            "txType": "TRANSFER",
                            "timeStamp": 1579688431580,
                            "txFee": 0.00060000,
                            "txAge": 5273941,
                            "code": 0,
                            "log": "Msg 0: ",
                            "confirmBlocks": 0,
                            "memo": "Trust Wallet Redeem",
                            "source": 0,
                            "hasChildren": 1
                        }
                    ]
                }`);

            case 'bnb1mnf0g6zhy2rpy63w8ryc6yp77s0fmqfyzjtkvd':
                return JSON.parse(`{
                    "txNums": 1,
                    "txArray": [
                        {
                            "txHash": "DDA939C475755B2A00123CDA37A5EC442F502EA2983014D04F647C3CB9FB57A0",
                            "blockHeight": 75110907,
                            "txType": "TRANSFER",
                            "timeStamp": 1584450870939,
                            "fromAddr": "bnb166fvwtfra2atta5mm8mygt2fyltktqctgpedcg",
                            "toAddr": "bnb1mnf0g6zhy2rpy63w8ryc6yp77s0fmqfyzjtkvd",
                            "value": 0.00037500,
                            "txAsset": "BNB",
                            "txFee": 0.00037500,
                            "txAge": 10170,
                            "code": 0,
                            "log": "Msg 0: ",
                            "confirmBlocks": 0,
                            "memo": "",
                            "source": 0,
                            "hasChildren": 0
                        }
                    ]
                }`);
        }

        // not found, address
        return {txNums: 0, txArray: []}
    }
};
