/// Binance chain block explorer API Mock, Dex
/// See:
/// curl "http://localhost:3347/mock/binance-dex/v1/account/bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m"
/// curl "http://localhost:3347/mock/binance-dex/v1/tokens?limit=1000&offset=0"
/// curl "https://dex.binance.org/api/v1/account/bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m"
/// curl "https://dex.binance.org/api/v1/tokens?limit=1000&offset=0"
/// curl "http://localhost:8437/v2/binance/tokens/bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m?Authorization=Bearer"

module.exports = {
    path: '/mock/binance-dex/:version/:command1/:command2?',
    template: function(params, query, body) {
        switch (params.version) {
            case 'v1':
                switch (params.command1) {
                    case 'account':
                        switch (params.command2) {
                            case 'bnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m':
                                var fn = "../../ext-api-data/get/" +
                                    "mock%2Fbinance-dex%2Fv1%2Faccount%2Fbnb1jxfh2g85q3v0tdq56fnevx6xcxtcnhtsmcu64m.json";
                                var json = require(fn);
                                return json;
                        }
                        break;

                    case 'tokens':
                            var fn = "../../ext-api-data/get/" +
                            "mock%2Fbinance-dex%2Fv1%2Ftokens%3Flimit%3D1000%26offset%3D0.json";
                        var json = require(fn);
                        return json;
                }
        }

        // not found, address
        return {txNums: 0, txArray: []}
    }
};
