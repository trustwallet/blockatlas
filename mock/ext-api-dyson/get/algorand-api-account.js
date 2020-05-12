/// Mock for external Algorand API
/// curl "http://localhost:3347/mock/algorand-api/v1/account/4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U/transactions"
/// curl "https://{algorand rpc}/v1/account/4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U/transactions"
/// curl http://localhost:8437/v1/algorand/4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U

module.exports = {
    path: '/mock/algorand-api/v1/account/:account/transactions?',
    template: function(params, query, body) {
        //console.log(params)
        switch (params.account) {
            case '4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Falgorand-api%2Fv1%2Faccount%2F4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U%2Ftransactions.json";
                let json = require(fn);
                return json;
        }

        return {error: "Not implemented"}
    }
};
