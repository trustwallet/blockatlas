/// Mock for external Dash API
/// See:
/// curl "http://{dash rpc}/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG?details=txs"
/// curl "http://localhost:3000/dash-api/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG?details=txs"
/// curl "http://localhost:8420/v1/dash/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG"

module.exports = {
    path: '/dash-api/address/:address?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 1000,
                        "addrStr": "XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG",
                        "balance": "167.97080211",
                        "totalReceived": "1302.03239791",
                        "totalSent": "1134.0615958",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxApperances": 0,
                        "txApperances": 618,
                        "txs": [
                            {
                                "txid": "35541e151818606a512bb88acddb449449f46531e97f1ed0a5fcc0d7a8091e41",
                                "version": 3,
                                "vin": [
                                    {
                                        "txid": "",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": null,
                                        "value": ""
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1.55334749",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac",
                                            "addresses": [
                                                "XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "1.55334763",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a91404dfe99f6dda028e4e6a90595096e63787a4a01c88ac",
                                            "addresses": [
                                                "Xb8cmjtK67y9T2haqrcDicoAMByDesLcAe"
                                            ]
                                        },
                                        "spent": true
                                    }
                                ],
                                "blockhash": "000000000000000f432c35cbd89c013333cf591c72d15959a51100625d263243",
                                "blockheight": 1239931,
                                "confirmations": 3276,
                                "time": 1584614323,
                                "blocktime": 1584614323,
                                "valueOut": "3.10669512",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "03000500010000000000000000000000000000000000000000000000000000000000000000ffffffff27037beb121a4d696e656420627920416e74506f6f6c347c00330120ad45211b6b76000001000000ffffffff025d384209000000001976a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac6b384209000000001976a91404dfe99f6dda028e4e6a90595096e63787a4a01c88ac000000004602007beb12009d9eb909c943fd6b4aad565dc083dc2b53736359cc678af072352bbf409a942f3b0192d52c43692520241eb04505406f34adf6805fd5ab29113959133e8e1d3d"
                            },
                            {
                                "txid": "f47a4929bcd8767819b6969c83893a76d5b26a145c74ace416ce4c454166186e",
                                "version": 3,
                                "vin": [
                                    {
                                        "txid": "",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": null,
                                        "value": ""
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1.56344682",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac",
                                            "addresses": [
                                                "XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG"
                                            ]
                                        },
                                        "spent": false
                                    },
                                    {
                                        "value": "1.56344689",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a91404dfe99f6dda028e4e6a90595096e63787a4a01c88ac",
                                            "addresses": [
                                                "Xb8cmjtK67y9T2haqrcDicoAMByDesLcAe"
                                            ]
                                        },
                                        "spent": true
                                    }
                                ],
                                "blockhash": "000000000000001709a13d55d31782a8d26ea3691c4d3433794b9ff2ee0b0ac7",
                                "blockheight": 1237951,
                                "confirmations": 5256,
                                "time": 1584303664,
                                "blocktime": 1584303664,
                                "valueOut": "3.12689371",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "03000500010000000000000000000000000000000000000000000000000000000000000000ffffffff2703bfe3121a4d696e656420627920416e74506f6f6c357c006502207193bf88753900000d030000ffffffff026aa15109000000001976a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac71a15109000000001976a91404dfe99f6dda028e4e6a90595096e63787a4a01c88ac00000000460200bfe312006d5f02b8c49b966902195a3b2b581b4a5847901d5c66bf9820a35884c6a91b51b20d38a728281d7ba3a9b3c157593c9121850977cababcba9ab71729684d9723"
                            },
                            {
                                "txid": "aaceebd8a59a777c8d27ede7d086f605a04e3463bd3465fd3fa02518fcd24abd",
                                "version": 3,
                                "vin": [
                                    {
                                        "txid": "",
                                        "vout": 0,
                                        "sequence": 4294967295,
                                        "n": 0,
                                        "scriptSig": {},
                                        "addresses": null,
                                        "value": ""
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "1.55873001",
                                        "n": 0,
                                        "scriptPubKey": {
                                            "hex": "76a9142832b5c571b5686c4a08dae4091d856c4f9b190a88ac",
                                            "addresses": [
                                                "XeMPcKeVDN9bkECGDC7ggtf9QsX5thgKAx"
                                            ]
                                        },
                                        "spent": true
                                    },
                                    {
                                        "value": "1.55872988",
                                        "n": 1,
                                        "scriptPubKey": {
                                            "hex": "76a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac",
                                            "addresses": [
                                                "XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG"
                                            ]
                                        },
                                        "spent": false
                                    }
                                ],
                                "blockhash": "00000000000000021075e42769f6a1bd6ac477574266102b324197c158bf3a14",
                                "blockheight": 1237949,
                                "confirmations": 5258,
                                "time": 1584302904,
                                "blocktime": 1584302904,
                                "valueOut": "3.11745989",
                                "valueIn": "0",
                                "fees": "0",
                                "hex": "03000500010000000000000000000000000000000000000000000000000000000000000000ffffffff1703bde31212563c785f430fb8fa7417000000002f4e614effffffff02e96e4a09000000001976a9142832b5c571b5686c4a08dae4091d856c4f9b190a88acdc6e4a09000000001976a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac00000000460200bde312006d5f02b8c49b966902195a3b2b581b4a5847901d5c66bf9820a35884c6a91b51b20d38a728281d7ba3a9b3c157593c9121850977cababcba9ab71729684d9723"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
