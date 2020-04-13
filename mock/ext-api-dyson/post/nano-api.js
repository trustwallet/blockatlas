/// Nano RPC Mock
/// curl -H 'Content-Type: application/json' -d ' {"action":"account_history","account":"nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z","count":"25"} ' http://localhost:3000/nano-api
/// curl -H 'Content-Type: application/json' -d ' {"action":"account_history","account":"nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z","count":"25"} ' https://{nano_rpc}
/// curl "http://localhost:8420/v1/nano/nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z"

module.exports = {
    path: '/nano-api',
    template: function(params, query, body) {
        //console.log("curl -H 'Content-Type: application/json' -d '", JSON.stringify(body), "' https://{nano_rpc}");
        if (body.action === 'account_history') {
            //console.log('body.account', body.account);
            if (body.account === 'nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z') {
                return JSON.parse(`{
                    "account": "nano_36e7qfxrpixge3xxujtpc87c77mn9ubu3bhywfjkr1trnubtd4qswwydhn9z",
                    "history": [
                        {
                            "type": "send",
                            "account": "nano_1trqphog5noig7z888asnjejcie8z1iopxyepcjdo1atps8whxiuwd51ehbw",
                            "amount": "1083000821328155744662798729216",
                            "local_timestamp": "1576911470",
                            "height": "8",
                            "hash": "05A23A254272028F349BB7E74115A5101B2048786A5437D1EBBE8807CEFF1E82"
                        },
                        {
                            "type": "receive",
                            "account": "nano_1trqphog5noig7z888asnjejcie8z1iopxyepcjdo1atps8whxiuwd51ehbw",
                            "amount": "1047300821328155744662798729216",
                            "local_timestamp": "1576911457",
                            "height": "7",
                            "hash": "13B8146F059C4D5695DCBCEEC750EAEDDDD8F436A6FAD569A107F2FAC0F899D5"
                        },
                        {
                            "type": "receive",
                            "account": "nano_1trqphog5noig7z888asnjejcie8z1iopxyepcjdo1atps8whxiuwd51ehbw",
                            "amount": "34700000000000000000000000000",
                            "local_timestamp": "1576911446",
                            "height": "6",
                            "hash": "8A2A5840C9286B35D998F5AD535851750C1AFE7B0C7D3AA61E06A1628EDD0E94"
                        }
                    ]
                }`);
            }
            return {error: 'Bad account number'};
        }
        var return4Codacy = {error: 'Invalid request'};
        return return4Codacy;
    }
};
