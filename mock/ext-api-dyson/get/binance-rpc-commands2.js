/// Binance chain RPC Mock
/// See:
/// curl "http://localhost:3347/binance-rpc/v1/account/bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q"
/// curl "http://localhost:3347/binance-rpc/v1/tokens?limit=1000&offset=0"
/// curl "https://{binance_rpc}/v1/account/bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q"
/// curl "https://{binance_rpc}/v1/tokens?limit=1000&offset=0"
/// curl "http://localhost:8437/v2/binance/tokens/bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q?Authorization=Bearer"

module.exports = {
    path: '/binance-rpc/:version/:command1/:command2?',
    template: function(params, query, body) {
        switch (params.version) {
            case 'v1':
                switch (params.command1) {
                    case 'account':
                        switch (params.command2) {
                            case 'bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q':
                                return JSON.parse(`
                                    {
                                        "account_number": 273171,
                                        "address": "bnb1jeu6gscugy6l2wyatxthkh2hmer4hzevgcmf0q",
                                        "balances": [
                                            {
                                                "free": "226.53110295",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BNB"
                                            },
                                            {
                                                "free": "2623.96917801",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BUSD-BD1"
                                            },
                                            {
                                                "free": "0.05000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "TWT-8C2"
                                            }
                                        ],
                                        "flags": 0,
                                        "public_key": [
                                            2, 142, 117
                                        ],
                                        "sequence": 75
                                    }
                                `);
                        }
                        break;

                    case 'tokens':
                        return JSON.parse(`
                            [
                                {
                                    "mintable": true,
                                    "name": "Africa Stable-Coin",
                                    "original_symbol": "ABCD",
                                    "owner": "bnb1ujvzeuft0ezf9fu4u0mk52t8mc7t8geyfkevms",
                                    "symbol": "ABCD-5D8",
                                    "total_supply": "3347000.00000000"
                                },
                                {
                                    "mintable": false,
                                    "name": "Aditus",
                                    "original_symbol": "ADI",
                                    "owner": "bnb1djdymfgzknmcsu9dzm9s0uavdszn0cl82z4hps",
                                    "symbol": "ADI-6BB",
                                    "total_supply": "750000000.00000000"
                                },
                                {
                                    "mintable": false,
                                    "name": "Aergo",
                                    "original_symbol": "AERGO",
                                    "owner": "bnb1llqhwwwmh878844tm3g8v47k0t7xtnhl4hggjl",
                                    "symbol": "AERGO-46B",
                                    "total_supply": "500000000.00000000"
                                },
                                {
                                    "mintable": false,
                                    "name": "Alaris",
                                    "original_symbol": "ALA",
                                    "owner": "bnb1pmdkvw6cquwylr46wcrl82xzmul0y2jpj5cwx7",
                                    "symbol": "ALA-DCD",
                                    "total_supply": "60000000.00000000"
                                },
                                {
                                    "mintable": false,
                                    "name": "ANKR",
                                    "original_symbol": "ANKR",
                                    "owner": "bnb1hvg059mkwleum35j6y2qjn4fvmgl7zxtlah4tn",
                                    "symbol": "ANKR-E97",
                                    "total_supply": "10000000000.00000000"
                                },
                                {
                                    "mintable": false,
                                    "name": "Aeron",
                                    "original_symbol": "ARN",
                                    "owner": "bnb1dq8ae0ayztqp99peggq5sygzf3n7u2ze4t0jne",
                                    "symbol": "ARN-71B",
                                    "total_supply": "20000000.00000000"
                                },
                                {
                                    "mintable": true,
                                    "name": "ARPA",
                                    "original_symbol": "ARPA",
                                    "owner": "bnb1mecnt25u3j9ne7th5av7hqvnmzvyrr7ny8hg8c",
                                    "symbol": "ARPA-575",
                                    "total_supply": "12000000.00000000"
                                },
                                {
                                    "mintable": false,
                                    "name": "Maecenas ART Token",
                                    "original_symbol": "ART",
                                    "owner": "bnb13plj9kycvcew5v0achpatnd5l5pacys9h0gu8l",
                                    "symbol": "ART-3C9",
                                    "total_supply": "100000000.00000000"
                                },
                                {
                                    "mintable": true,
                                    "name": "Atlas Protocol",
                                    "original_symbol": "ATP",
                                    "owner": "bnb1msw3avv894nlpeu0vn4qlkl0r65a3rp7gtz5hf",
                                    "symbol": "ATP-38C",
                                    "total_supply": "40000000.00000000"
                                },
                                {
                                    "mintable": false,
                                    "name": "Travala.com Token",
                                    "original_symbol": "AVA",
                                    "owner": "bnb1dm9c7gccgd07td5r69m50u8fg8danfgqvlhj6c",
                                    "symbol": "AVA-645",
                                    "total_supply": "61383832.00000000"
                                }
                            ]
                        `);
                }
        }

        // not found, address
        return {txNums: 0, txArray: []}
    }
};
