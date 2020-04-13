/// Iotex API Mock
/// See:
/// curl "http://localhost:3000/iotex-api/staking/validators?status=bonded"
/// curl "http://localhost:3000/iotex-api/accounts/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
/// curl "https://{iotex_rpc}/v1/staking/validators?status=bonded"
/// curl "https://{iotex_rpc}/v1/accounts/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
/// curl "http://localhost:8420/v2/iotex/staking/validators"
/// curl "http://localhost:8420/v2/iotex/staking/delegations/io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m?Authorization=Bearer"

module.exports = {
    path: "/iotex-api/:command1/:command2?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch(params.command1) {
            case 'accounts':
                switch(params.command2) {
                    case 'io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m':
                        // status=bonded
                        return JSON.parse(`{"accountMeta":{"address":"io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m","balance":"15578806681028832609262","nonce":"583","pendingNonce":"584","numActions":"647"}}`);

                    case 'io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5':
                        return JSON.parse(`{"accountMeta":{"address":"io1vg808avg2ydye8djl2axmkc9j0xhzu6vdaw6g5","balance":"33829692159741946535","nonce":"4","pendingNonce":"5","numActions":"57"}}`);
                }

            case 'staking':
                switch(params.command2) {
                    case 'validators':
                        // status=bonded
                        return JSON.parse(`
                            [
                                {
                                    "id": "huobiwallet",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexcore",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "pubxpayments",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "droute",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "coredev",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "enlightiv",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iosg",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "royalland",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "gamefantasy#",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "laomao",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "cpc",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "hashbuy",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "hotbit",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotxplorerio",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexteam",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "airfoil",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "capitmu",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "ducapital",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "pnp",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "longz",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexlab",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "hofancrypto",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "draperdragon",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "yvalidator",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "satoshi",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "hashquark",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "mrtrump",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "elitex",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "ratels",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "rockx",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "blockfolio",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "preangel",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "consensusnet",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "blockboost",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "metanyx",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "coingecko",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "thebottoken#",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "cobo",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "zhcapital",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "infstones",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "tgb",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexgeeks",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotask",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "cryptolionsx",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "bittaker",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "keys",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "snzholding",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexunion",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "raketat8",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "wannodes",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "citex2018",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "wetez",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "eon",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "whales",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "eatliverun",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexicu",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "everstake",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "nodeasy",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "slowmist",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "alphacoin",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexhub",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "blackpool",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "superiotex",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "link",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexmainnet",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "lanhu",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "piexgo",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "iotexbgogo",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "meter",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "bitwires",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
                                    }
                                },
                                {
                                    "id": "elink",
                                    "status": true,
                                    "details": {
                                        "reward": {
                                            "annual": 0
                                        },
                                        "locktime": 259200,
                                        "minimum_amount": "100000000000000000000"
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

