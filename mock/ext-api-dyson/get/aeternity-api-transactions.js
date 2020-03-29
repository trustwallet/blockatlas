/// Aeternity API Mock
/// See:
/// curl "http://localhost:3000/aeternity-api/middleware/transactions/account/ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb?"
/// curl "https://mdw.aepps.com/middleware/transactions/account/ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb?"
/// curl http://localhost:8420/v1/aeternity/ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb
module.exports = {
    path: "/aeternity-api/middleware/transactions/:type/:address?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.type === 'account') {
            if (params.address === 'ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb') {
                return JSON.parse(`
                    [
                        {
                            "block_hash": "mh_2V2PJ732cTfDqD4CaWr6W1s1fN531PqiVydMiy7gooccQFoGvM",
                            "block_height": 222994,
                            "hash": "th_WfeoRYXd13MMmDBzMXjBBdaNSSaRXkwUwJ7tFxJ7ZKPQVNXC4",
                            "signatures": [
                                "sg_MLBxuqnHUD21i747SKnQuyNh1LzLuH7RDHqY1mngZyJMDWL9LFcuMbKdo2XMbwkMmskFsgfsZZBwgyA5w4Xs3ghynF2AF"
                            ],
                            "time": 1583633007081,
                            "tx": {
                                "amount": 1.99979979992e21,
                                "fee": 16840000000000,
                                "nonce": 3,
                                "payload": "ba_Xfbg4g==",
                                "recipient_id": "ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb",
                                "sender_id": "ak_26dGJ7S5ba7LhUXPo8vMPe4QuNaQbskWnM5jkopPuim3PGny9Y",
                                "type": "SpendTx",
                                "version": 1
                            }
                        },
                        {
                            "block_hash": "mh_Z1a6HSs9vE5XDf8JtSiQMN5UpX49Spgnf9Sn32TiXxKwSbsP6",
                            "block_height": 186035,
                            "hash": "th_2wGEGmqwRMPYdv3vz1EE7ueJEveMTzEWzosgUH9ytERjGMv3rX",
                            "signatures": [
                                "sg_66c2LXv4rvYnn4r2fZ4H1wxieDUuVEUF6fLCB3M13MjB9QNfEH4q9YThUWJCo5WzWKxzqjb3RbumhXdUZh8YDq3v8KZ2M"
                            ],
                            "time": 1576946253731,
                            "tx": {
                                "amount": 1.00079975e21,
                                "fee": 50000000000000,
                                "nonce": 3,
                                "payload": "ba_Xfbg4g==",
                                "recipient_id": "ak_26dGJ7S5ba7LhUXPo8vMPe4QuNaQbskWnM5jkopPuim3PGny9Y",
                                "sender_id": "ak_2WGWYMgWy1opZxgA8AVzGCTavCQyUBtbKx5SrCX6E4kmDZMtJb",
                                "ttl": 186135,
                                "type": "SpendTx",
                                "version": 1
                            }
                        }
                    ]
                `);
            }
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
