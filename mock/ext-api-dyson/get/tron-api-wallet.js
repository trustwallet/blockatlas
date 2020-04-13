/// Tron API Mock
/// See:
/// curl http://localhost:3000/tron-api/wallet/listwitnesses
/// curl http://localhost:8420/v2/tron/staking/validators

module.exports = {
    path: "/tron-api/wallet/:operation?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.operation === 'listwitnesses') {
            return JSON.parse(`
                {
                    "witnesses": [

                        {
                            "address": "4178c842ee63b253f8f0d2955bbc582c661a078c9d",
                            "voteCount": 13546358416,
                            "url": "https://www.binance.com/en/staking",
                            "totalProduced": 197868,
                            "totalMissed": 68,
                            "latestBlockNum": 18523838,
                            "latestSlotNum": 528611454,
                            "isJobs": true
                        },
                        {
                            "address": "41d25855804e4e65de904faf3ac74b0bdfc53fac76",
                            "voteCount": 698985941,
                            "url": "https://www.bitguild.com",
                            "totalProduced": 594077,
                            "totalMissed": 2605,
                            "latestBlockNum": 18523841,
                            "latestSlotNum": 528611457,
                            "isJobs": true
                        },
                        {
                            "address": "41d376d829440505ea13c9d1c455317d51b62e4ab6",
                            "voteCount": 317223159,
                            "url": "http://blockchain.org",
                            "totalProduced": 439090,
                            "totalMissed": 2650,
                            "latestBlockNum": 18523376,
                            "latestSlotNum": 528610983,
                            "isJobs": true
                        },
                        {
                            "address": "4192c5d96c3b847268f4cb3e33b87ecfc67b5ce3de",
                            "voteCount": 331299971,
                            "url": "https://infstones.io/",
                            "totalProduced": 528971,
                            "totalMissed": 2232,
                            "latestBlockNum": 18523824,
                            "latestSlotNum": 528611440,
                            "isJobs": true
                        },
                        {
                            "address": "417bdd2efb4401c50b6ad255e6428ba688e0b83f81",
                            "voteCount": 292566620,
                            "url": "https://minergate.com",
                            "totalProduced": 365859,
                            "totalMissed": 798,
                            "latestBlockNum": 18523360,
                            "latestSlotNum": 528610967,
                            "isJobs": true
                        },
                        {
                            "address": "414d1ef8673f916debb7e2515a8f3ecaf2611034aa",
                            "voteCount": 392387821,
                            "url": "https://www.sesameseed.org",
                            "totalProduced": 667595,
                            "totalMissed": 5738,
                            "latestBlockNum": 18523844,
                            "latestSlotNum": 528611460,
                            "isJobs": true
                        },
                        {
                            "address": "41de9c3c2276abe2da70a7cdb34a205ecf7750d063",
                            "voteCount": 4575746,
                            "url": "https://www.tron-family.de"
                        }
                    ]
                }
            `)
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
