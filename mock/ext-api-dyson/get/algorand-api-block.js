/// Mock for external Algorand API
/// curl "http://localhost:3347/mock/algorand-api/v1/block/5478346"
/// curl "https://{algorand rpc}/v1/block/5478346"
/// curl http://localhost:8437/v1/algorand/4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U

module.exports = {
    path: '/mock/algorand-api/v1/block/:block?',
    template: function(params, query, body) {
        //console.log(params)
        switch (params.block) {
            case '5478346':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Falgorand-api%2Fv1%2Fblock%2F5478346.json";
                let json = require(fn);
                return json;
        }

        return {error: "Not implemented"}
    }
};
