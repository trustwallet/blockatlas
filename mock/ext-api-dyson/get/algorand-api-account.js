/// Mock for external Algorand API
/// curl "http://localhost:3000/algorand-api/v1/account/4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U/transactions?"
/// curl "http://{algorand rpc}/v1/account/4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U/transactions?"
/// curl http://localhost:8420/v1/algorand/4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U

module.exports = {
    path: '/algorand-api/v1/account/:account/transactions?',
    template: function(params, query, body) {
        //console.log(params)
        switch (params.account) {
            case '4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U':
                return JSON.parse(`
                {
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
                }
                `);
        }
        // fallback
        return {error: "Not implemented"}
    }
};
