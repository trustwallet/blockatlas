/// Mock for external Zilliqa API
/// See:
/// curl -H "X-APIKEY: YOUR_API_KEY" "https://api.viewblock.io/v1/zilliqa/addresses/zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z/txs"
/// curl "http://localhost:3000/zilliqa-api/addresses/zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z/txs"
/// curl "http://localhost:8420/v1/zilliqa/zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z"

module.exports = {
    path: '/zilliqa-api/addresses/:address/txs',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z':
                return JSON.parse(`
                    [
                        {
                            "hash": "0xa5d0e0cefd6f114e9659f15016f9f4a303a8367eb373beb6909ee44f7f39834d",
                            "blockHeight": 459052,
                            "from": "zil10lx2eurx5hexaca0lshdr75czr025cevqu83uz",
                            "to": "zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z",
                            "value": "1000000000000",
                            "fee": "1000000000",
                            "timestamp": 1583384813191,
                            "signature": "0x00AF9733DF8FAEFA0FCC9BBAD21DB30A4815749C19CBFFA5A3BCE75A199E4C84FB4728CB946A2F043B8FCA338C8EEC26855CD7B88D8925E69F87CADF5D072343",
                            "direction": "in",
                            "nonce": 25,
                            "receiptSuccess": true,
                            "events": []
                        },
                        {
                            "hash": "0xe29a7e17402c0c067af4c285dedc79114fe62f23d39843c15756db4641f1a00d",
                            "blockHeight": 414452,
                            "from": "zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z",
                            "to": "zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z",
                            "value": "10000000000",
                            "fee": "1000000000",
                            "timestamp": 1580891803587,
                            "signature": "0xA5F2C86098A37149CDDD0884009FF1032714DDAC07B2831C87E1FEEB16B8D9B5EB5D310D042003B3AA27A31BE16FD62CD41E5B1F5108475FEB3EB5B8524DA3F4",
                            "direction": "self",
                            "nonce": 47,
                            "receiptSuccess": true,
                            "events": []
                        }
                    ]
                `);
        }
        return {error: "Not implemented"};
    }
}
