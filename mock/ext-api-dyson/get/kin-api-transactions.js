/// Kin API Mock, transactions
/// See:
/// curl "http://localhost:3000/kin-api/transactions/b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a?"
/// curl "https://horizon-block-explorer.kininfrastructure.com/transactions/b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a?"
/// curl http://localhost:8420/v1/kin/GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX
module.exports = {
    path: "/kin-api/transactions/:txid?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.txid === 'b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a') {
            return JSON.parse(`
            {
                "memo": "Namilak8",
                "_links": {
                  "self": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions/b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a"
                  },
                  "account": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH"
                  },
                  "ledger": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/ledgers/7485062"
                  },
                  "operations": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions/b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a/operations{?cursor,limit,order}",
                    "templated": true
                  },
                  "effects": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions/b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a/effects{?cursor,limit,order}",
                    "templated": true
                  },
                  "precedes": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions?order=asc\u0026cursor=32148096498577408"
                  },
                  "succeeds": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions?order=desc\u0026cursor=32148096498577408"
                  }
                },
                "id": "b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a",
                "paging_token": "32148096498577408",
                "hash": "b2131beb8e0c57dfd3309914fce08ac4e094c51dc918087171150dd5b996076a",
                "ledger": 7485062,
                "created_at": "2020-03-24T01:43:00Z",
                "source_account": "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH",
                "source_account_sequence": "3873257342118155",
                "fee_paid": 100,
                "operation_count": 1,
                "envelope_xdr": "AAAAAE6qZ+LUs9bu+pfd6c/LbS0iA0G7JR6ac5R9XIIyqC6IAAAAZAANwrUAACkLAAAAAAAAAAEAAAAITmFtaWxhazgAAAABAAAAAQAAAABOqmfi1LPW7vqX3enPy20tIgNBuyUemnOUfVyCMqguiAAAAAEAAAAAbDIyvkAp46X97VW4/zmT18NY7BAFyxQxDNv5ve/FylMAAAAAAAAAA4aZXmAAAAAAAAAAATKoLogAAABAiivAdbjfQvWTMheamksBZp9yBc9oRN2E2OpPhTCeT/bICu6eSr3FJy+WqYlieBlXtdVbBaoCWOlR07Mi4QqjDQ==",
                "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAA=",
                "result_meta_xdr": "AAAAAAAAAAEAAAAEAAAAAwByNoYAAAAAAAAAAE6qZ+LUs9bu+pfd6c/LbS0iA0G7JR6ac5R9XIIyqC6IAABHRPE+rKwADcK1AAApCwAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQByNoYAAAAAAAAAAE6qZ+LUs9bu+pfd6c/LbS0iA0G7JR6ac5R9XIIyqC6IAABHQWqlTkwADcK1AAApCwAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAwBwLBkAAAAAAAAAAGwyMr5AKeOl/e1VuP85k9fDWOwQBcsUMQzb+b3vxcpTAAAAAzCVZ9QAcAeUAAAAAwAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQByNoYAAAAAAAAAAGwyMr5AKeOl/e1VuP85k9fDWOwQBcsUMQzb+b3vxcpTAAAABrcuxjQAcAeUAAAAAwAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA",
                "fee_meta_xdr": "AAAAAgAAAAMAcindAAAAAAAAAABOqmfi1LPW7vqX3enPy20tIgNBuyUemnOUfVyCMqguiAAAR0TxPq0QAA3CtQAAKQoAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAcjaGAAAAAAAAAABOqmfi1LPW7vqX3enPy20tIgNBuyUemnOUfVyCMqguiAAAR0TxPqysAA3CtQAAKQsAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
                "memo_type": "text",
                "signatures": [
                  "iivAdbjfQvWTMheamksBZp9yBc9oRN2E2OpPhTCeT/bICu6eSr3FJy+WqYlieBlXtdVbBaoCWOlR07Mi4QqjDQ=="
                ]
              }
            `);
        }
        if (params.txid === 'eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed') {
            return JSON.parse(`
            {
                "memo": "",
                "_links": {
                  "self": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions/eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed"
                  },
                  "account": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/accounts/GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH"
                  },
                  "ledger": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/ledgers/7481821"
                  },
                  "operations": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions/eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed/operations{?cursor,limit,order}",
                    "templated": true
                  },
                  "effects": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions/eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed/effects{?cursor,limit,order}",
                    "templated": true
                  },
                  "precedes": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions?order=asc\u0026cursor=32134176509775872"
                  },
                  "succeeds": {
                    "href": "https://horizon-block-explorer.kininfrastructure.com/transactions?order=desc\u0026cursor=32134176509775872"
                  }
                },
                "id": "eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed",
                "paging_token": "32134176509775872",
                "hash": "eb070218bfde8c48af6071432bc9f19b7690b9dc50535cc135f67ea046e70bed",
                "ledger": 7481821,
                "created_at": "2020-03-23T21:12:01Z",
                "source_account": "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH",
                "source_account_sequence": "3873257342118154",
                "fee_paid": 100,
                "operation_count": 1,
                "envelope_xdr": "AAAAAE6qZ+LUs9bu+pfd6c/LbS0iA0G7JR6ac5R9XIIyqC6IAAAAZAANwrUAACkKAAAAAAAAAAEAAAAAAAAAAQAAAAEAAAAATqpn4tSz1u76l93pz8ttLSIDQbslHppzlH1cgjKoLogAAAABAAAAADqMsEcsu8a3aP7yB8M2R8im3p4D+wRWZFJY8tvI0bNRAAAAAAAAAF1caQX1AAAAAAAAAAEyqC6IAAAAQFsrhOcyOXVWLIFf8aSJh37W1bx4ZP3qKXPejdB2M0z+FG09ggG0uDiV7FAsmRvzPUGoEgI3reEw+eAFByO6dAs=",
                "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAA=",
                "result_meta_xdr": "AAAAAAAAAAEAAAAEAAAAAwByKaQAAAAAAAAAADqMsEcsu8a3aP7yB8M2R8im3p4D+wRWZFJY8tvI0bNRAAAAdGpUDqAAcik+AAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQByKd0AAAAAAAAAADqMsEcsu8a3aP7yB8M2R8im3p4D+wRWZFJY8tvI0bNRAAAA0ca9FJUAcik+AAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAwByKd0AAAAAAAAAAE6qZ+LUs9bu+pfd6c/LbS0iA0G7JR6ac5R9XIIyqC6IAABHok2nswUADcK1AAApCgAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQByKd0AAAAAAAAAAE6qZ+LUs9bu+pfd6c/LbS0iA0G7JR6ac5R9XIIyqC6IAABHRPE+rRAADcK1AAApCgAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA",
                "fee_meta_xdr": "AAAAAgAAAAMAcimkAAAAAAAAAABOqmfi1LPW7vqX3enPy20tIgNBuyUemnOUfVyCMqguiAAAR6JNp7NpAA3CtQAAKQkAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAcindAAAAAAAAAABOqmfi1LPW7vqX3enPy20tIgNBuyUemnOUfVyCMqguiAAAR6JNp7MFAA3CtQAAKQoAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
                "memo_type": "text",
                "signatures": [
                  "WyuE5zI5dVYsgV/xpImHftbVvHhk/eopc96N0HYzTP4UbT2CAbS4OJXsUCyZG/M9QagSAjet4TD54AUHI7p0Cw=="
                ]
              }
            `);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
