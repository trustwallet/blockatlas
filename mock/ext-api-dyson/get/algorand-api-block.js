/// Mock for external Algorand API
/// curl "http://localhost:3000/algorand-api/v1/block/5478346?"
/// curl "http://{algorand rpc}/v1/block/5478346?"
/// curl http://localhost:8420/v1/algorand/4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U

module.exports = {
    path: '/algorand-api/v1/block/:block?',
    template: function(params, query, body) {
        //console.log(params)
        switch (params.block) {
            case '5478346':
                return JSON.parse(`
                {
                    "hash": "SU4PUD4YK5DPIXUXDLEZOEHVVIIP4SWZORFPJBW2J22QO444NSAA",
                    "previousBlockHash": "CC4CQRE2U5AC7S7JLJHDVVGIDCHGWTDFIZM3P34TVTC7AP4BA2AQ",
                    "seed": "NLAQIQXPCZI5QMKMRFFKMIKANNBC4PTFW5QFVAK7U6FE477T2TRA",
                    "proposer": "WCFKBT4NEPDUCQP7MGLIWV4QGXIBOCXJFVYNTXF3HZHFHGQGHGR5WWA7XA",
                    "round": 5478346,
                    "period": 0,
                    "txnRoot": "KFADZYI64LDUX76VWVVV3MROORV5RXAVNZA7MVTWOP2QTGW6DZ2A",
                    "reward": 111421,
                    "rate": 25999967,
                    "frac": 2639736378,
                    "txns": {
                        "transactions": [
                            {
                                "type": "pay",
                                "tx": "JZTKP6RRFGUDOZ7DQIR26DQWRDYZVYG2G3WDE4WLBJ77AMAEBFBA",
                                "from": "5TSQNIL54GB545B3WLC6OVH653SHAELMHU6MSVNGTUNMOEHAMWG7EC3AA4",
                                "fee": 1000,
                                "first-round": 5478300,
                                "last-round": 5478749,
                                "noteb64": "sHLxsLBrP3o=",
                                "round": 5478346,
                                "payment": {
                                    "to": "4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U",
                                    "amount": 1,
                                    "torewards": 2052177,
                                    "closerewards": 0
                                },
                                "fromrewards": 0,
                                "genesisID": "mainnet-v1.0",
                                "genesishashb64": "wGHE2Pwdvd7S12BL5FaOP20EGYesN73ktiC1qzkkit8="
                            }
                        ]
                    },
                    "timestamp": 1584332782,
                    "currentProtocol": "https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa",
                    "nextProtocol": "",
                    "nextProtocolApprovals": 0,
                    "nextProtocolVoteBefore": 0,
                    "nextProtocolSwitchOn": 0,
                    "upgradePropose": "",
                    "upgradeApprove": false
                }
                `);
        }
        // fallback
        return {error: "Not implemented"}
    }
};
