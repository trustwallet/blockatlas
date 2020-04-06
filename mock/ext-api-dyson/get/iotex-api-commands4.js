/// Iotex API Mock
/// See:
/// curl "http://localhost:3000/iotex-api/staking/delegations/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
/// curl "https://{iotex_rpc}/v1/staking/delegations/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
/// curl "http://localhost:8420/v2/iotex/staking/delegations/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m?Authorization=Bearer"

module.exports = {
    path: "/iotex-api/:command1/:command2/:arg3/:arg4?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch(params.command1) {
            case 'staking':
                switch(params.command2) {
                    case 'delegations':
                        switch (params.arg3) {
                            case 'io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m':
                                return JSON.parse(`
                                    [{
                                        "delegator": {
                                            "id": "iotxplorerio",
                                            "status": true,
                                            "info": {
                                                "name": "iotxplorer",
                                                "description": "",
                                                "image": "https://imgc.iotex.io/dokc3pa1x/image/upload/v1551121446/delegates/iotxplorer/Group_2.png",
                                                "website": "https://twitter.com/iotxplorer"
                                            },
                                            "details": {
                                                "reward": {"annual": 0},
                                                "locktime": 259200,
                                                "minimum_amount": "100000000000000000000"
                                            }
                                        },
                                        "value": "100000000000000000000",
                                        "status": "active"
                                    }]
                                `);
                        }
                }
        }

        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};

