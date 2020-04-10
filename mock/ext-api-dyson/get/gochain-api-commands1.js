/// Gochain API Mock
/// See:
/// curl "http://localhost:3000/gochain-api/transactions?address=0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896"
/// curl "http://{go rpc}/transactions?address=0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896"
/// curl http://localhost:8420/v1/gochain/0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896

module.exports = {
    path: '/gochain-api/transactions',
    template: function(params, query, body) {
        //console.log(query)
        if (query.address === '0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896') {
            return JSON.parse(`{
                "docs": [
                    {
                        "operations": [],
                        "contract": null,
                        "_id": "0x2637d7b0e7f66134e983e6af921700718b8a49885a9215ee8033ac384d5d1be7",
                        "blockNumber": 10070535,
                        "time": 1576793173,
                        "nonce": 6,
                        "from": "0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896",
                        "to": "0x65a0e277B2d0eEF62b8C30b280B4C49f83a104d3",
                        "value": "860000000000000000000000",
                        "gas": "90000",
                        "gasPrice": "2000000000",
                        "gasUsed": "21000",
                        "input": "0x",
                        "error": "",
                        "id": "0x2637d7b0e7f66134e983e6af921700718b8a49885a9215ee8033ac384d5d1be7",
                        "timeStamp": "1576793173"
                    },
                    {
                        "operations": [],
                        "contract": null,
                        "_id": "0x472c94385a46072713bb2bf3503fc1610135a70e415046bf61a937789c205f4c",
                        "blockNumber": 10070520,
                        "time": 1576793098,
                        "nonce": 5,
                        "from": "0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896",
                        "to": "0x65a0e277B2d0eEF62b8C30b280B4C49f83a104d3",
                        "value": "800000000000000000000",
                        "gas": "90000",
                        "gasPrice": "2000000000",
                        "gasUsed": "21000",
                        "input": "0x",
                        "error": "",
                        "id": "0x472c94385a46072713bb2bf3503fc1610135a70e415046bf61a937789c205f4c",
                        "timeStamp": "1576793098"
                    }
                ],
                "total": 2
            }`);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
