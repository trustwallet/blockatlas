/// OpenSea API Mock
/// See:

/// curl "http://localhost:3347/mock/opensea-api/api/v1/assets/?collection=unstoppable-domains&limit=300&owner=0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d"
/// curl "https://api.opensea.io/api/v1/assets/?collection=unstoppable-domains&limit=300&owner=0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d"
/// curl "http://localhost:8437/v4/ethereum/collections/0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d/collection/unstoppable-domains"

/// curl "http://localhost:3347/mock/opensea-api/api/v1/collections/?asset_owner=0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d&limit=1000"
/// curl "https://api.opensea.io/api/v1/collections?asset_owner=0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d&limit=1000"
/// curl -d '{"60":["0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d"]}' http://localhost:8437/v4/collectibles/categories

module.exports = {
    path: "/mock/opensea-api/api/v1/:command1/?",
    template: function(params, query) {
        switch (params.command1) {
            case 'assets':
                if (query.owner == '0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d' &&
                    query.collection == 'unstoppable-domains') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fopensea-api%2Fapi%2Fv1%2Fassets%2F%3Fcollection%3Dunstoppable-domains%26limit%3D300%26owner%3D0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d.json";
                    var json = require(fn);
                    return json;
                }
                break;

            case 'collections':
                if (query.asset_owner == '0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fopensea-api%2Fapi%2Fv1%2Fcollections%2F%3Fasset_owner%3D0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d%26limit%3D1000.json";
                    var json = require(fn);
                    return json;
                }
                break;
        }

        return {error: "Not implemented"};
    }
};
