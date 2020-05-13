/// Mock for external Zilliqa API
/// See:
/// curl -H "X-APIKEY: YOUR_API_KEY" "https://api.viewblock.io/v1/zilliqa/addresses/zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z/txs"
/// curl "http://localhost:3347/mock/zilliqa-api/addresses/zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z/txs"
/// curl "http://localhost:8437/v1/zilliqa/zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z"

module.exports = {
    path: '/mock/zilliqa-api/addresses/:address/txs',
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.address) {
            case 'zil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fzilliqa-api%2Faddresses%2Fzil1l8ddxvejeam70qang54wnqkgtmlu5mwlgzy64z%2Ftxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
