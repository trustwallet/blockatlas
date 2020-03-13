/// Mock for external Digibyte API
/// See:
/// curl "http://{digibyte rpc}/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi?details=txs"
/// curl "http://localhost:3000/digibyte-api/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi?details=txs"
/// curl "http://localhost:8420/v1/digibyte/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi"

module.exports = {
    path: '/digibyte-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 4,
                        "itemsOnPage": 1000,
                        "addrStr": "DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi",
                        "balance": "0",
                        "totalReceived": "2317.03950234",
                        "totalSent": "2317.03950234",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 3428,
                        "txs": [
                            {
                                "txid": "3ad1035d37f47061599a5055284a8b2acc0facc6039353b1822f0a3c0f1d6906",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "14f845047925fb36e5a0c569f2c1d2b7c5be4543e46920b49eb732860bf5b02a",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "47304402202fc0e7eebec384bcff141800990f80742d2ed97db02784fb383cd4cd42aca71e02202be0f292a8209bff5682e05424170632bb7023a7aaca0dca1faf54af1e8eec16012102aa3f5da884ab5654137cf660576e590830f6ec177ca84629d5a0264a1e494057"
                                        },
                                        "addresses": [
                                            "DCzmzkMBqEz2tLn47W9YuNAV9cFzuWCydW"
                                        ],
                                        "value": "0.7988"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.79879",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a9146aa65418d9de46d678eb71db1d2616f5a96acdc588ac",
                                            "addresses": [
                                                "DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi"
                                            ]
                                        },
                                        "spent": true
                                    },
                                    {
                                        "value": "0",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "6a45524d555453422e41432e544800000000000025aba398bee3241347368c2d4d758575c7fe00138abe258456087c4f640ae3274c452f87d110bc12ee1f77ed9b365a49d6c547",
                                            "addresses": [
                                                "OP_RETURN 524d555453422e41432e544800000000000025aba398bee3241347368c2d4d758575c7fe00138abe258456087c4f640ae3274c452f87d110bc12ee1f77ed9b365a49d6c547"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "0000000000000007fa73216ba848e3d029f91a088f628cc1780f333d7e13311d",
                                "blockheight": 10510734,
                                "confirmations": 30461,
                                "time": 1584668841,
                                "blocktime": 1584668841,
                                "valueOut": "0.79879",
                                "valueIn": "0.7988",
                                "fees": "0.00001",
                                "hex": "01000000012ab0f50b8632b79eb42069e44345bec5b7d2c1f269c5a0e536fb25790445f814000000006a47304402202fc0e7eebec384bcff141800990f80742d2ed97db02784fb383cd4cd42aca71e02202be0f292a8209bff5682e05424170632bb7023a7aaca0dca1faf54af1e8eec16012102aa3f5da884ab5654137cf660576e590830f6ec177ca84629d5a0264a1e494057ffffffff0258dbc204000000001976a9146aa65418d9de46d678eb71db1d2616f5a96acdc588ac0000000000000000476a45524d555453422e41432e544800000000000025aba398bee3241347368c2d4d758575c7fe00138abe258456087c4f640ae3274c452f87d110bc12ee1f77ed9b365a49d6c54700000000"
                            },
                            {
                                "txid": "d7efaffeded550697d7385d3eb07e539237615792b29e3639df427ce919c8d6e",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "a7613f7242b3a24ca6f1de2c15be605411d01f6d86ed1cf0037da9d0e908bf0e",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "47304402201d09be67618dc89a882c389560e73713302cb3d120adbae814e06ed15b2a17e002204595a11648173fe085a820b034394c1264ddd30b0e2dfdb07d3ebb46067145c10141041dd9c41a6c1a3ea6c9fd53eecc0314c79386eb56f380552d46644d106bdbc4dfd389027e1498313480441bbf5130461af785656deaea87348dd7ac523eaeff48"
                                        },
                                        "addresses": [
                                            "DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi"
                                        ],
                                        "value": "0.71187"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.71186",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914562e330d8879eeaf17c72c8ad7f3acb4ebb6262c88ac",
                                            "addresses": [
                                                "DCzmzkMBqEz2tLn47W9YuNAV9cFzuWCydW"
                                            ]
                                        },
                                        "spent": true
                                    },
                                    {
                                        "value": "0",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "6a45524d555453422e41432e544800000000000025a8a398bee3241347368c2d4d758575c7fe00f810517db039aa4a3055e73c5eca51dccdad9a31ff0deccfa3ab44611f3f1891",
                                            "addresses": [
                                                "OP_RETURN 524d555453422e41432e544800000000000025a8a398bee3241347368c2d4d758575c7fe00f810517db039aa4a3055e73c5eca51dccdad9a31ff0deccfa3ab44611f3f1891"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "6518a4fe30d601f87370aca047d1e2bcc3eaf72a194fc05f475bdbf24ca6d36e",
                                "blockheight": 10493876,
                                "confirmations": 47319,
                                "time": 1584416875,
                                "blocktime": 1584416875,
                                "valueOut": "0.71186",
                                "valueIn": "0.71187",
                                "fees": "0.00001",
                                "hex": "01000000010ebf08e9d0a97d03f01ced866d1fd0115460be152cdef1a64ca2b342723f61a7000000008a47304402201d09be67618dc89a882c389560e73713302cb3d120adbae814e06ed15b2a17e002204595a11648173fe085a820b034394c1264ddd30b0e2dfdb07d3ebb46067145c10141041dd9c41a6c1a3ea6c9fd53eecc0314c79386eb56f380552d46644d106bdbc4dfd389027e1498313480441bbf5130461af785656deaea87348dd7ac523eaeff48ffffffff0250363e04000000001976a914562e330d8879eeaf17c72c8ad7f3acb4ebb6262c88ac0000000000000000476a45524d555453422e41432e544800000000000025a8a398bee3241347368c2d4d758575c7fe00f810517db039aa4a3055e73c5eca51dccdad9a31ff0deccfa3ab44611f3f189100000000"
                            },
                            {
                                "txid": "a7613f7242b3a24ca6f1de2c15be605411d01f6d86ed1cf0037da9d0e908bf0e",
                                "version": 1,
                                "vin": [
                                    {
                                        "txid": "3e4e8f95cf76c3a1bb91283e8c55ba5a7c103cc31b1c25c4587916666fbb8263",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {
                                            "hex": "47304402202560cf31d48b023fb44383c79f1685420af0365d87377589f9fa493f66c6d97f02201724f2c0c41ff384fc53d128fec351093d8f2d54332cc7191153008080c4f9a7012102aa3f5da884ab5654137cf660576e590830f6ec177ca84629d5a0264a1e494057"
                                        },
                                        "addresses": [
                                            "DCzmzkMBqEz2tLn47W9YuNAV9cFzuWCydW"
                                        ],
                                        "value": "0.71188"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "0.71187",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a9146aa65418d9de46d678eb71db1d2616f5a96acdc588ac",
                                            "addresses": [
                                                "DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi"
                                            ]
                                        },
                                        "spent": true
                                    },
                                    {
                                        "value": "0",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "6a45524d555453422e41432e544800000000000025a8a398bee3241347368c2d4d758575c7fe00f810517db039aa4a3055e73c5eca51dccdad9a31ff0deccfa3ab44611f3f1891",
                                            "addresses": [
                                                "OP_RETURN 524d555453422e41432e544800000000000025a8a398bee3241347368c2d4d758575c7fe00f810517db039aa4a3055e73c5eca51dccdad9a31ff0deccfa3ab44611f3f1891"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "6518a4fe30d601f87370aca047d1e2bcc3eaf72a194fc05f475bdbf24ca6d36e",
                                "blockheight": 10493876,
                                "confirmations": 47319,
                                "time": 1584416875,
                                "blocktime": 1584416875,
                                "valueOut": "0.71187",
                                "valueIn": "0.71188",
                                "fees": "0.00001",
                                "hex": "01000000016382bb6f66167958c4251c1bc33c107c5aba558c3e2891bba1c376cf958f4e3e000000006a47304402202560cf31d48b023fb44383c79f1685420af0365d87377589f9fa493f66c6d97f02201724f2c0c41ff384fc53d128fec351093d8f2d54332cc7191153008080c4f9a7012102aa3f5da884ab5654137cf660576e590830f6ec177ca84629d5a0264a1e494057ffffffff02383a3e04000000001976a9146aa65418d9de46d678eb71db1d2616f5a96acdc588ac0000000000000000476a45524d555453422e41432e544800000000000025a8a398bee3241347368c2d4d758575c7fe00f810517db039aa4a3055e73c5eca51dccdad9a31ff0deccfa3ab44611f3f189100000000"
                            }
                        ]
                    }                
                `);
        }
        return {error: "Not implemented"};
    }
}
