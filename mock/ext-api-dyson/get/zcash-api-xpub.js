/// Mock for external Zcash API
/// See:
/// curl "http://{Zcash rpc}/v2/xpub/xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS?details=txs"
/// curl "http://localhost:3000/zcash-api/v2/xpub/xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS?details=txs"
/// curl "http://localhost:8420/v1/zcash/xpub/xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS"

module.exports = {
    path: '/zcash-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS':
                return JSON.parse(`
                    {
                        "page": 1,
                        "totalPages": 1,
                        "itemsOnPage": 2,
                        "address": "xpub6CCXGBJ13akWuKSn4iU7CqQeXzLyDC2Y1Z83Mmg2xz11PX2EeZJJKRECz29iN4eHewRh8yfb7FpnCcjYbkqn6ynHnXW3jczPcJcenThfFeS",
                        "balance": "846466",
                        "totalReceived": "14304018",
                        "totalSent": "13457552",
                        "unconfirmedBalance": "0",
                        "unconfirmedTxs": 0,
                        "txs": 38,
                        "transactions": [
                            {
                                "txid": "f2438a93039faf08d39bd3df1f7b5f19a2c29ffe8753127e2956ab4461adab35",
                                "version": 4,
                                "vin": [
                                    {
                                        "txid": "ee120b714991e6166b8ff4b6419cbfee657eca87675245ce8585787be3b07d70",
                                        "vout": 1,
                                        "n": 0,
                                        "addresses": [
                                            "t1LJWoRDU14zUG4TGumHiobd9WBUjqLE5FU"
                                        ],
                                        "isAddress": true,
                                        "value": "956466",
                                        "hex": "483045022100d2ac7b2b17218572f3dc163063463be5143a8d3c2126b52eb73d49c9c4959da00220057417baaad53c168e82dd44a7f964759c7d584e0770b6284a8e1b0261f1c7e0012102a71eeea42aa1b4e30a409f9186bd88dc84dec5c404964f386c323f557ad268d9"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "100000",
                                        "n": 0,
                                        "hex": "76a914a2511fc37f67f3b5627affce89fb3f7bb186412888ac",
                                        "addresses": [
                                            "t1Yfrf1dssDLmaMBsq2LFKWPbS5vH3nGpa2"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "846466",
                                        "n": 1,
                                        "hex": "76a9146fd73e7c147d8ccc15fda31d8429e70f302b843988ac",
                                        "addresses": [
                                            "t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "00000000005a238fecd893e55fba7072af65545fd43eaad63402c211cc172a61",
                                "blockHeight": 750255,
                                "confirmations": 18102,
                                "blockTime": 1583385097,
                                "value": "946466",
                                "valueIn": "956466",
                                "fees": "10000",
                                "hex": "0400008085202f8901707db0e37b788585ce45526787ca7e65eebf9c41b6f48f6b16e69149710b12ee010000006b483045022100d2ac7b2b17218572f3dc163063463be5143a8d3c2126b52eb73d49c9c4959da00220057417baaad53c168e82dd44a7f964759c7d584e0770b6284a8e1b0261f1c7e0012102a71eeea42aa1b4e30a409f9186bd88dc84dec5c404964f386c323f557ad268d90000000002a0860100000000001976a914a2511fc37f67f3b5627affce89fb3f7bb186412888ac82ea0c00000000001976a9146fd73e7c147d8ccc15fda31d8429e70f302b843988ac00000000000000000000000000000000000000"
                            },
                            {
                                "txid": "ee120b714991e6166b8ff4b6419cbfee657eca87675245ce8585787be3b07d70",
                                "version": 4,
                                "vin": [
                                    {
                                        "txid": "b49791a277ff137e37f9e163c8a5bc25492f741c275f9d36d5c0c713497c176c",
                                        "vout": 1,
                                        "sequence": 2147483646,
                                        "n": 0,
                                        "addresses": [
                                            "t1fuw3P3r5xbJXFqpaHUWb3TMz6xo9yPXEL"
                                        ],
                                        "isAddress": true,
                                        "value": "996466",
                                        "hex": "483045022100e7ddc39a2b976ee277d71f6130334084f1f21ef4f17cdd81514cf8af2181df1b02201c9a76a45b948c87940dc47525f1660d6c9c1e44ab22e4bcc16af0ce99c4c1890121031a03d4f4982aaa59d4087434907c9578fe77eadc73491b32f7373aaf9c17902e"
                                    }
                                ],
                                "vout": [
                                    {
                                        "value": "30000",
                                        "n": 0,
                                        "spent": true,
                                        "hex": "76a914db49ca3524ee1e99aa097c6c9860ffcb5e5ed94a88ac",
                                        "addresses": [
                                            "t1ds6PMeHUvAYtASPqCuEXzT1DsteXWwtTv"
                                        ],
                                        "isAddress": true
                                    },
                                    {
                                        "value": "956466",
                                        "n": 1,
                                        "spent": true,
                                        "hex": "76a9141aa64e287339ad3fa041bc7a429bbcd57ac8a80088ac",
                                        "addresses": [
                                            "t1LJWoRDU14zUG4TGumHiobd9WBUjqLE5FU"
                                        ],
                                        "isAddress": true
                                    }
                                ],
                                "blockHash": "00000000008dcc62d679d7eaf05ac18403198f5f39b3e4e7f36869673855779b",
                                "blockHeight": 685806,
                                "confirmations": 82551,
                                "blockTime": 1578529212,
                                "value": "986466",
                                "valueIn": "996466",
                                "fees": "10000",
                                "hex": "0400008085202f89016c177c4913c7c0d5369d5f271c742f4925bca5c863e1f9377e13ff77a29197b4010000006b483045022100e7ddc39a2b976ee277d71f6130334084f1f21ef4f17cdd81514cf8af2181df1b02201c9a76a45b948c87940dc47525f1660d6c9c1e44ab22e4bcc16af0ce99c4c1890121031a03d4f4982aaa59d4087434907c9578fe77eadc73491b32f7373aaf9c17902efeffff7f0230750000000000001976a914db49ca3524ee1e99aa097c6c9860ffcb5e5ed94a88ac32980e00000000001976a9141aa64e287339ad3fa041bc7a429bbcd57ac8a80088ac00000000000000000000000000000000000000"
                            }
                        ]
                    }
                `);
        }
        return {error: "Not implemented"};
    }
}
