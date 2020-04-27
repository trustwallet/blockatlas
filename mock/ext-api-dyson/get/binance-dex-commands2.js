/// Binance chain block explorer API Mock, Dex
/// See:
/// curl "http://localhost:3347/binance-dex/v1/account/bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m"
/// curl "http://localhost:3347/binance-dex/v1/tokens?limit=1000&offset=0"
/// curl "https://dex.binance.org/api/v1/account/bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m"
/// curl "https://dex.binance.org/api/v1/tokens?limit=1000&offset=0"
/// curl "http://localhost:8437/v2/binance/tokens/bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m?Authorization=Bearer"

module.exports = {
    path: '/binance-dex/:version/:command1/:command2?',
    template: function(params, query, body) {
        //console.log(params);
        //console.log(query);
        switch (params.version) {
            case 'v1':
                switch (params.command1) {
                    case 'account':
                        switch (params.command2) {
                            case 'bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m':
                                return JSON.parse(`
                                    {
                                        "account_number": 51,
                                        "address": "bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m",
                                        "balances": [
                                            {
                                                "free": "1732627268.91580163",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "CHZ-ECD"
                                            },
                                            {
                                                "free": "927279.98843502",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "GTO-908"
                                            },
                                            {
                                                "free": "22990392.97970594",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "LBA-340"
                                            },
                                            {
                                                "free": "13504.20603578",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "LTC-F07"
                                            },
                                            {
                                                "free": "1.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "SLV-986"
                                            },
                                            {
                                                "free": "3400.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "TWT-8C2"
                                            },
                                            {
                                                "free": "417595410.39908617",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "COS-2E4"
                                            },
                                            {
                                                "free": "8859448.20434497",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "XRP-BF2"
                                            },
                                            {
                                                "free": "66199647.50538896",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "VRAB-B56"
                                            },
                                            {
                                                "free": "565.32000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "AERGO-46B"
                                            },
                                            {
                                                "free": "145.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ART-3C9"
                                            },
                                            {
                                                "free": "483312.47553546",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BOLT-4C6"
                                            },
                                            {
                                                "free": "85668472.69714541",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ERD-D06"
                                            },
                                            {
                                                "free": "372741.55000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "MDAB-D42"
                                            },
                                            {
                                                "free": "56343499.10367935",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ONE-5F9"
                                            },
                                            {
                                                "free": "2000000.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BCPT-95A"
                                            },
                                            {
                                                "free": "5.00110000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BEAR-14C"
                                            },
                                            {
                                                "free": "2000.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "GIV-94E"
                                            },
                                            {
                                                "free": "9.44356120",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "NEXO-A84"
                                            },
                                            {
                                                "free": "193182441.20916349",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "UND-EBC"
                                            },
                                            {
                                                "free": "569871.98092216",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "TUSDB-888"
                                            },
                                            {
                                                "free": "316344310.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BKRW-AB7"
                                            },
                                            {
                                                "free": "0.00091449",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "EOSBEAR-721"
                                            },
                                            {
                                                "free": "10918.75299816",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "EOSBULL-F0D"
                                            },
                                            {
                                                "free": "296751584.10699693",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "TROY-9B8"
                                            },
                                            {
                                                "free": "450587749.03659246",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "WINB-41F"
                                            },
                                            {
                                                "free": "1438139.59497068",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ARN-71B"
                                            },
                                            {
                                                "free": "19597996.48343213",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "COTI-CBB"
                                            },
                                            {
                                                "free": "3572612.16675809",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "LTO-BDF"
                                            },
                                            {
                                                "free": "1.22227000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "XRPBEAR-00B"
                                            },
                                            {
                                                "free": "8619272.37637807",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "CAN-677"
                                            },
                                            {
                                                "free": "338507.95870700",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "HNST-3C9"
                                            },
                                            {
                                                "free": "173741504.28850373",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BTTB-D31"
                                            },
                                            {
                                                "free": "65500.72562560",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BUSD-BD1"
                                            },
                                            {
                                                "free": "8549.17210965",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ETH-1C9"
                                            },
                                            {
                                                "free": "1654100.82380293",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "TOMOB-4BC"
                                            },
                                            {
                                                "free": "5.00455774",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "XRPBULL-E7C"
                                            },
                                            {
                                                "free": "35.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "AWC-986"
                                            },
                                            {
                                                "free": "1587996.65265528",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BNB"
                                            },
                                            {
                                                "free": "45605.00637205",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "QARK-FCE"
                                            },
                                            {
                                                "free": "1406863.84448749",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ARPA-575"
                                            },
                                            {
                                                "free": "1619733.62576473",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ENTRP-C8D"
                                            },
                                            {
                                                "free": "0.01392000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "MTXLT-286"
                                            },
                                            {
                                                "free": "2340.58605759",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BCH-1FD"
                                            },
                                            {
                                                "free": "3.49158551",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BULL-BE4"
                                            },
                                            {
                                                "free": "1.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ECO-083"
                                            },
                                            {
                                                "free": "0.50000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "WICC-01D"
                                            },
                                            {
                                                "free": "4235575.51194174",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "KAVA-10C"
                                            },
                                            {
                                                "free": "10.00000000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "PVT-554"
                                            },
                                            {
                                                "free": "9772.00009772",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "UPX-F3E"
                                            },
                                            {
                                                "free": "221000.60154058",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "USDSB-1AC"
                                            },
                                            {
                                                "free": "4372695.27174854",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "WRX-ED1"
                                            },
                                            {
                                                "free": "19906076.03857030",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "PHB-2DF"
                                            },
                                            {
                                                "free": "12696933.52213685",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "TRXB-2E6"
                                            },
                                            {
                                                "free": "267793524.10892121",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BLINK-9C6"
                                            },
                                            {
                                                "free": "3370.02600563",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "BTCB-1DE"
                                            },
                                            {
                                                "free": "103467039.59858401",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "CBM-4B2"
                                            },
                                            {
                                                "free": "4095309.32370000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "DUSK-45E"
                                            },
                                            {
                                                "free": "82964636.96725382",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "FTM-A64"
                                            },
                                            {
                                                "free": "182686435.86912742",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ANKR-E97"
                                            },
                                            {
                                                "free": "5.00100000",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ETHBEAR-B2B"
                                            },
                                            {
                                                "free": "0.00872511",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "ETHBULL-D33"
                                            },
                                            {
                                                "free": "39286407.28666093",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "MATIC-84A"
                                            },
                                            {
                                                "free": "20813555.86251786",
                                                "frozen": "0.00000000",
                                                "locked": "0.00000000",
                                                "symbol": "MITH-C76"
                                            }
                                        ],
                                        "flags": 0,
                                        "public_key": [
                                            3,
                                            86,
                                            224,
                                            165,
                                            128,
                                            56,
                                            154,
                                            111,
                                            210,
                                            204,
                                            145,
                                            205,
                                            82,
                                            92,
                                            109,
                                            90,
                                            77,
                                            128,
                                            84,
                                            175,
                                            112,
                                            223,
                                            23,
                                            72,
                                            78,
                                            88,
                                            103,
                                            143,
                                            159,
                                            87,
                                            74,
                                            11,
                                            77
                                        ],
                                        "sequence": 547072
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
