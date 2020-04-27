/// Mock for external Digibyte API
/// See:
/// curl "http://{digibyte rpc}/api/v2/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi?details=txs"
/// curl "http://localhost:3347/digibyte-api/v2/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi?details=txs"
/// curl "http://localhost:8437/v1/digibyte/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi"

module.exports = {
    path: '/digibyte-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 1821,
                    "itemsOnPage": 2,
                    "address": "DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi",
                    "balance": "0",
                    "totalReceived": "238524073234",
                    "totalSent": "238524073234",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 3642,
                    "transactions": [
                        {
                            "txid": "2b7ba9b43d615fe03bc14bc18201d1bd73c7a0f8aad428accf51725743f9073a",
                            "version": 1,
                            "vin": [
                                {
                                    "txid": "cba15c1782f15d1f3a5dfd6f06bad6ee03183d4d2bb3699bf72b9a423f4a97a3",
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "addresses": [
                                        "DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi"
                                    ],
                                    "isAddress": true,
                                    "value": "47451000",
                                    "hex": "473044022015642218d196b2bf010e10eb0eff952e0a6bd5bcb8aecd09751599bef8ca5df802205d277e3f6ac47ad360785df4462ffd2092618a58bcb0c8ae94751f26e5918fa30141041dd9c41a6c1a3ea6c9fd53eecc0314c79386eb56f380552d46644d106bdbc4dfd389027e1498313480441bbf5130461af785656deaea87348dd7ac523eaeff48"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "47450000",
                                    "n": 0,
                                    "spent": true,
                                    "hex": "76a914562e330d8879eeaf17c72c8ad7f3acb4ebb6262c88ac",
                                    "addresses": [
                                        "DCzmzkMBqEz2tLn47W9YuNAV9cFzuWCydW"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "0",
                                    "n": 1,
                                    "hex": "6a45524d555453422e41432e5448000000000000280ca398bee3241347368c2d4d758575c7fe0079c3ea7e1cb73919815c2534d932c4b305dd86f2aafcfc843c264594069454de",
                                    "addresses": [
                                        "OP_RETURN 524d555453422e41432e5448000000000000280ca398bee3241347368c2d4d758575c7fe0079c3ea7e1cb73919815c2534d932c4b305dd86f2aafcfc843c264594069454de"
                                    ],
                                    "isAddress": false
                                }
                            ],
                            "blockHash": "00000000000000018af4a98ade5bb1d29d65871fd99b3cbaa8ea6266ea2c16a5",
                            "blockHeight": 10712917,
                            "confirmations": 1047,
                            "blockTime": 1587685950,
                            "value": "47450000",
                            "valueIn": "47451000",
                            "fees": "1000",
                            "hex": "0100000001a3974a3f429a2bf79b69b32b4d3d1803eed6ba066ffd5d3a1f5df182175ca1cb000000008a473044022015642218d196b2bf010e10eb0eff952e0a6bd5bcb8aecd09751599bef8ca5df802205d277e3f6ac47ad360785df4462ffd2092618a58bcb0c8ae94751f26e5918fa30141041dd9c41a6c1a3ea6c9fd53eecc0314c79386eb56f380552d46644d106bdbc4dfd389027e1498313480441bbf5130461af785656deaea87348dd7ac523eaeff48ffffffff029007d402000000001976a914562e330d8879eeaf17c72c8ad7f3acb4ebb6262c88ac0000000000000000476a45524d555453422e41432e5448000000000000280ca398bee3241347368c2d4d758575c7fe0079c3ea7e1cb73919815c2534d932c4b305dd86f2aafcfc843c264594069454de00000000"
                        },
                        {
                            "txid": "cba15c1782f15d1f3a5dfd6f06bad6ee03183d4d2bb3699bf72b9a423f4a97a3",
                            "version": 1,
                            "vin": [
                                {
                                    "txid": "fae67ad84a3f6a7e7648f12122d649bb711178f0fd129286e2566df1f31fa08f",
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "addresses": [
                                        "DCzmzkMBqEz2tLn47W9YuNAV9cFzuWCydW"
                                    ],
                                    "isAddress": true,
                                    "value": "47452000",
                                    "hex": "46304302205da10139f9ff5f0c0579237108634edce8dafa82feb7f63f0104b4596c5d0c98021f75ca71439e45a274b8953fa383d51938403043d961b59e9b078ef50a6fba01012102aa3f5da884ab5654137cf660576e590830f6ec177ca84629d5a0264a1e494057"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "47451000",
                                    "n": 0,
                                    "spent": true,
                                    "hex": "76a9146aa65418d9de46d678eb71db1d2616f5a96acdc588ac",
                                    "addresses": [
                                        "DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "0",
                                    "n": 1,
                                    "hex": "6a45524d555453422e41432e5448000000000000280ca398bee3241347368c2d4d758575c7fe0079c3ea7e1cb73919815c2534d932c4b305dd86f2aafcfc843c264594069454de",
                                    "addresses": [
                                        "OP_RETURN 524d555453422e41432e5448000000000000280ca398bee3241347368c2d4d758575c7fe0079c3ea7e1cb73919815c2534d932c4b305dd86f2aafcfc843c264594069454de"
                                    ],
                                    "isAddress": false
                                }
                            ],
                            "blockHash": "00000000000000018af4a98ade5bb1d29d65871fd99b3cbaa8ea6266ea2c16a5",
                            "blockHeight": 10712917,
                            "confirmations": 1047,
                            "blockTime": 1587685950,
                            "value": "47451000",
                            "valueIn": "47452000",
                            "fees": "1000",
                            "hex": "01000000018fa01ff3f16d56e2869212fdf0781171bb49d62221f148767e6a3f4ad87ae6fa000000006946304302205da10139f9ff5f0c0579237108634edce8dafa82feb7f63f0104b4596c5d0c98021f75ca71439e45a274b8953fa383d51938403043d961b59e9b078ef50a6fba01012102aa3f5da884ab5654137cf660576e590830f6ec177ca84629d5a0264a1e494057ffffffff02780bd402000000001976a9146aa65418d9de46d678eb71db1d2616f5a96acdc588ac0000000000000000476a45524d555453422e41432e5448000000000000280ca398bee3241347368c2d4d758575c7fe0079c3ea7e1cb73919815c2534d932c4b305dd86f2aafcfc843c264594069454de00000000"
                        }
                    ]
                }                
                `);
        }
        return {error: "Not implemented"};
    }
}
