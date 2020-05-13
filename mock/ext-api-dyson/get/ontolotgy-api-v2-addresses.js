/// Ontology API Mock
/// See:
/// curl "http://localhost:3347/mock/ontology-api/v2/addresses/AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7/transactions?page_size=20&page_number=1"
/// curl "https://explorer.ont.io/v2/addresses/AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7/transactions?page_size=20&page_number=1"
/// curl http://localhost:8437/v1/ontology/AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7
module.exports = {
    path: "/mock/ontology-api/v2/addresses/:address/:operation?",
    template: function(params, query, body) {
        //console.log(params)
        if (params.address === 'AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fontology-api%2Fv2%2Faddresses%2FAUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7%2Ftransactions%3Fpage_number%3D1%26page_size%3D20.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"};
    }
};
