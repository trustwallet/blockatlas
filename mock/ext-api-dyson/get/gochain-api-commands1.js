/// Gochain API Mock
/// See:
/// curl "http://localhost:3347/gochain-api/transactions?address=0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896"
/// curl "http://localhost:3347/gochain-api/tokens?address=0x0Fd98FB42C439E5F6484f7E71Caa6661d81d0628"
/// curl "https://{go rpc}/transactions?address=0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896"
/// curl "https://{go rpc}/ 
/// curl "http://localhost:8437/v1/gochain/0xEd7F2e81B0264177e0df8f275f97Fd74Fa51A896"
/// curl "http://localhost:8437/v2/gochain/tokens/0x0Fd98FB42C439E5F6484f7E71Caa6661d81d0628?Authorization=Bearer"

module.exports = {
    path: '/gochain-api/:command1?',
    template: function(params, query, body) {
        //console.log(params);
        //console.log(query);
        switch (params.command1) {
            case 'transactions':
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
                break;

            case 'tokens':
                if (query.address === '0x0Fd98FB42C439E5F6484f7E71Caa6661d81d0628') {
                    return JSON.parse(`
                        {
                            "total": 1,
                            "docs": [
                                {
                                    "address": "0x5f16Fa0B5c9d779a3C8d46859a27973Ff3511188",
                                    "name": "pukkamex",
                                    "decimals": 18,
                                    "symbol": "PUX"
                                }
                            ]
                        }
                    `);
                }
                break;
        }

        return {error: "Not implemented"};
    }
};
