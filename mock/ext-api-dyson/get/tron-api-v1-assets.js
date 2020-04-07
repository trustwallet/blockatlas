/// Tron API Mock
/// See:
/// curl "http://localhost:3000/tron-api/v1/assets/1002798"
/// curl "https://{tron_rpc}/v1/assets/1002798"
/// curl "http://localhost:8420/v2/tron/tokens/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB?Authorization=Bearer"

module.exports = {
    path: "/tron-api/v1/assets/:arg2?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.arg2) {
            case '1002000':
                    return JSON.parse(`
                        {
                            "success": true,
                            "meta": {
                                "at": 1586265512148,
                                "page_size": 1
                            },
                            "data": [
                                {
                                    "id": "1002000",
                                    "abbr": "BTT",
                                    "description": "Official Token of BitTorrent Protocol",
                                    "name": "BitTorrent",
                                    "num": 1,
                                    "precision": 6,
                                    "total_supply": "990000000000000000",
                                    "trx_num": 1,
                                    "url": "www.bittorrent.com",
                                    "owner_address": "4137fa1a56eb8c503624701d776d95f6dae1d9f0d6",
                                    "start_time": 1548000000000,
                                    "end_time": 1548000001000
                                }
                            ]
                        }
                    `);
                    
            case '1002798':
                return JSON.parse(`
                    {
                        "success": true,
                        "meta": {
                            "at": 1586265317378,
                            "page_size": 1
                        },
                        "data": [
                            {
                                "id": "1002798",
                                "abbr": "EPICAL",
                                "description": "The token of the game Builder III",
                                "name": "EPICAL",
                                "num": 1000000000,
                                "precision": 6,
                                "total_supply": "10000000000000000",
                                "trx_num": 1000000000,
                                "url": "https://builder3.fun",
                                "vote_score": 0,
                                "owner_address": "41133b084f5225a9112f4e8527db26b381a32babd0",
                                "start_time": 1574966426578,
                                "end_time": 1574966486578
                            }
                        ]
                    }
                `);



        case '1002814':
                return JSON.parse(`
                    {
                        "success": true,
                        "meta": {
                            "at": 1586265542289,
                            "page_size": 1
                        },
                        "data": [
                            {
                                "id": "1002814",
                                "abbr": "AX",
                                "description": "The token of the game TrainX",
                                "name": "AX",
                                "num": 1000000000,
                                "precision": 6,
                                "total_supply": "10000000000000000",
                                "trx_num": 1000000000,
                                "url": "https://trainx.fun",
                                "vote_score": 0,
                                "owner_address": "418e267ead411aaaf671be100a7afe587d4eab0d71",
                                "start_time": 1576483805985,
                                "end_time": 1576483865985
                            }
                        ]
                    }
                `);
        }

        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
