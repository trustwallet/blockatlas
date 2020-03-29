/// Kin API Mock, accounts
/// See:
/// curl "http://localhost:3000/kin-api/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH/payments?order=desc&limit=25"
/// curl "https://horizon-block-explorer.kininfrastructure.com/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH/payments?order=desc&limit=25"
/// curl http://localhost:8420/v1/kin/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH
module.exports = {
    path: "/kin-api/accounts/:address/:operation?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.operation === 'payments') {
            if (params.address === 'GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH') {
                return JSON.parse(`
                {
                    "_links": {
                      "self": {
                        "href": "https://horizon-block-explorer.kininfrastructure.com/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH/payments?cursor=\u0026limit=25\u0026order=desc"
                      },
                      "next": {
                        "href": "https://horizon-block-explorer.kininfrastructure.com/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH/payments?cursor=31839477328764929\u0026limit=25\u0026order=desc"
                      },
                      "prev": {
                        "href": "https://horizon-block-explorer.kininfrastructure.com/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH/payments?cursor=32148096498577409\u0026limit=25\u0026order=asc"
                      }
                    },
                    "_embedded": {
                      "records": [
                        {
                          "_links": {
                            "self": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/operations/32148096498577409"
                            },
                            "transaction": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/transactions/b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a"
                            },
                            "effects": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/operations/32148096498577409/effects"
                            },
                            "succeeds": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/effects?order=desc\u0026cursor=32148096498577409"
                            },
                            "precedes": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/effects?order=asc\u0026cursor=32148096498577409"
                            }
                          },
                          "id": "32148096498577409",
                          "paging_token": "32148096498577409",
                          "source_account": "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH",
                          "type": "payment",
                          "type_i": 1,
                          "created_at": "2020-03-24T01:43:00Z",
                          "transaction_hash": "b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a",
                          "asset_type": "native",
                          "from": "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH",
                          "to": "GBWDEMV6IAU6HJP55VK3R7ZZSPL4GWHMCAC4WFBRBTN7TPPPYXFFH4AR",
                          "amount": "151431.00000"
                        },
                        {
                          "_links": {
                            "self": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/operations/32134176509775873"
                            },
                            "transaction": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/transactions/eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed"
                            },
                            "effects": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/operations/32134176509775873/effects"
                            },
                            "succeeds": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/effects?order=desc\u0026cursor=32134176509775873"
                           },
                            "precedes": {
                              "href": "https://horizon-block-explorer.kininfrastructure.com/effects?order=asc\u0026cursor=32134176509775873"
                            }
                          },
                          "id": "32134176509775873",
                          "paging_token": "32134176509775873",
                          "source_account": "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH",
                          "type": "payment",
                          "type_i": 1,
                          "created_at": "2020-03-23T21:12:01Z",
                          "transaction_hash": "eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed",
                          "asset_type": "native",
                          "from": "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH",
                          "to": "GA5IZMCHFS54NN3I73ZAPQZWI7EKNXU6AP5QIVTEKJMPFW6I2GZVD3JO",
                          "amount": "4009823.45205"
                        }
                      ]
                    }
                  }
                `);
            }
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
