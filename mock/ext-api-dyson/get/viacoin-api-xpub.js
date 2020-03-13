/// Mock for external Viacoin API
/// See:
/// curl "http://{Viacoin rpc}/v2/xpub/zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK?details=txs"
/// curl "http://localhost:3000/viacoin-api/v2/xpub/zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK?details=txs"
/// curl "http://localhost:8420/v1/viacoin/xpub/zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK"

module.exports = {
    path: '/viacoin-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 1,
                    "itemsOnPage": 1,
                    "address": "zpub6qVn6ubhK9tfepuABqy8wBXXn3qUZTbpqyNBqLyqakqTrZZD9rXZ3L5MZ945g8Mu7vmMSbC7vfLtLatTgxAnVJ8ECCtwmKqCo6TJm2ZsFJK",
                    "balance": "741749040",
                    "totalReceived": "57216749588",
                    "totalSent": "56475000548",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 39,
                    "transactions": [
                        {
                            "txid": "bdff404da14940abd84f7cb741fe8f77ad5de5aefbb74254caee683fe2a9b540",
                            "version": 1,
                            "vin": [
                                {
                                    "txid": "32790f42de714507b6a1e6001a4a334600ac316937a98489c19821bba0d5c74f",
                                    "vout": 1,
                                    "n": 0,
                                    "addresses": [
                                        "via1qxdul6szs3mppq8s3cdge5vtpthnw0dnn06m5qp"
                                    ],
                                    "value": "741862718"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "100000000",
                                    "n": 0,
                                    "hex": "00143379fd40508ec2101e11c3519a31615de6e7b673",
                                    "addresses": [
                                        "via1qxdul6szs3mppq8s3cdge5vtpthnw0dnn06m5qp"
                                    ]
                                },
                                {
                                    "value": "641749040",
                                    "n": 1,
                                    "hex": "00143379fd40508ec2101e11c3519a31615de6e7b673",
                                    "addresses": [
                                        "via1qxdul6szs3mppq8s3cdge5vtpthnw0dnn06m5qp"
                                    ]
                                }
                            ],
                            "blockHash": "0853179f0baef807426b977e4a9ae260ba1a3a3726dbb7513ff4150991f5827a",
                            "blockHeight": 7305935,
                            "confirmations": 117673,
                            "blockTime": 1581914336,
                            "value": "741749040",
                            "valueIn": "741862718",
                            "fees": "113678",
                            "hex": "010000000001014fc7d5a0bb2198c18984a9376931ac0046334a1a00e6a1b6074571de420f79320100000000000000000200e1f505000000001600143379fd40508ec2101e11c3519a31615de6e7b67330504026000000001600143379fd40508ec2101e11c3519a31615de6e7b67302473044022049b1d030bd9ec88ec343a52971d354b470d1357413675e985161bf2d5f8ca8ba02205d393b4d5232ae983faba4237e0c8feace5483493bf4da35313ce62459c87c66012102b5e09e1dbb76a0eeebea67bf80069ce41c26c81d0f06d1dd42e712ebf55de4bc00000000"
                        }
                    ],
                    "usedTokens": 40,
                    "tokens": [
                        {
                            "type": "XPUBAddress",
                            "name": "via1qxdul6szs3mppq8s3cdge5vtpthnw0dnn06m5qp",
                            "path": "m/84'/14'/0'/0/0",
                            "transfers": 7,
                            "decimals": 8,
                            "balance": "741749040",
                            "totalReceived": "9445588154",
                            "totalSent": "8703839114"
                        }
                    ]
                }
                `);
        }
        return {error: "Not implemented"};
    }
}
