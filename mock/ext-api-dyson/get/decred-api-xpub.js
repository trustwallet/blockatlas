/// Mock for external Decred API
/// See:
/// curl "http://{decred rpc}v2/xpub/dpubZFf1tYMxcku9nvDRxnYdE4yrEESrkuQFRq5RwA4KoYQKpDSRszN2emePTwLgfQpd4mZHGrHbQkKPZdjH1BcopomXRnr5Gt43rjpNEfeuJLN?details=txs"
/// curl "http://localhost:3347/mock/decred-api/v2/xpub/dpubZFf1tYMxcku9nvDRxnYdE4yrEESrkuQFRq5RwA4KoYQKpDSRszN2emePTwLgfQpd4mZHGrHbQkKPZdjH1BcopomXRnr5Gt43rjpNEfeuJLN?details=txs"
/// curl "http://localhost:8437/v1/decred/xpub/dpubZFf1tYMxcku9nvDRxnYdE4yrEESrkuQFRq5RwA4KoYQKpDSRszN2emePTwLgfQpd4mZHGrHbQkKPZdjH1BcopomXRnr5Gt43rjpNEfeuJLN"

module.exports = {
    path: '/mock/decred-api/v2/xpub/:xpubkey?',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.xpubkey) {
            case 'dpubZFf1tYMxcku9nvDRxnYdE4yrEESrkuQFRq5RwA4KoYQKpDSRszN2emePTwLgfQpd4mZHGrHbQkKPZdjH1BcopomXRnr5Gt43rjpNEfeuJLN':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fdecred-api%2Fv2%2Fxpub%2FdpubZFf1tYMxcku9nvDRxnYdE4yrEESrkuQFRq5RwA4KoYQKpDSRszN2emePTwLgfQpd4mZHGrHbQkKPZdjH1BcopomXRnr5Gt43rjpNEfeuJLN%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
