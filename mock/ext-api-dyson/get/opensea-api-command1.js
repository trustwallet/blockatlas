/// OpenSea API Mock
/// See:

/// curl "http://localhost:3347/opensea-api/api/v1/assets/?collection=unstoppable-domains&limit=300&owner=0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d"
/// curl "https://api.opensea.io/api/v1/assets/?collection=unstoppable-domains&limit=300&owner=0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d"
/// curl "http://localhost:8437/v4/ethereum/collections/0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d/collection/unstoppable-domains"

/// curl "http://localhost:3347/opensea-api/api/v1/collections/?asset_owner=0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d&limit=1000"
/// curl "https://api.opensea.io/api/v1/collections?asset_owner=0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d&limit=1000"
/// curl -d '{"60":["0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d"]}' http://localhost:8437/v4/collectibles/categories

module.exports = {
    path: "/opensea-api/api/v1/:command1/?",
    template: function(params, query) {
        switch (params.command1) {
            case 'assets':
                if (query.owner == '0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d' &&
                    query.collection == 'unstoppable-domains') {
                    return JSON.parse(`
                        {
                            "assets": [
                                {
                                    "token_id": "107221826727469155718680588460379721284635325845036369130032935095484960793482",
                                    "num_sales": 0,
                                    "background_color": "4C47F7",
                                    "image_url": "https://storage.opensea.io/files/fec912a0664aecaa908ee888177ebfc6.svg",
                                    "image_preview_url": "https://lh3.googleusercontent.com/cqCHmaBgbAtaE5i-jRKLGbJy2C-AUnDXN8zpkw8JPKnqRNGXdMJme_ThTFL30mDN9u9piJ4ACcfHmmDAyaXfIocN=s250",
                                    "image_thumbnail_url": "https://lh3.googleusercontent.com/cqCHmaBgbAtaE5i-jRKLGbJy2C-AUnDXN8zpkw8JPKnqRNGXdMJme_ThTFL30mDN9u9piJ4ACcfHmmDAyaXfIocN=s128",
                                    "image_original_url": null,
                                    "animation_url": null,
                                    "animation_original_url": null,
                                    "name": "sloth",
                                    "description": "A .crypto blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.",
                                    "external_link": "https://unstoppabledomains.com/search?searchTerm=sloth.crypto",
                                    "asset_contract": {
                                        "address": "0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe",
                                        "asset_contract_type": "non-fungible",
                                        "created_date": "2019-12-10T08:55:01.176657",
                                        "name": ".crypto",
                                        "nft_version": "3.0",
                                        "opensea_version": null,
                                        "owner": null,
                                        "schema_name": "ERC721",
                                        "symbol": "",
                                        "total_supply": null,
                                        "description": "Simplify your crypto currency payments with human readable names and build censorship resistant websites. Purchase your blockchain domains today!",
                                        "external_link": "https://unstoppabledomains.com/",
                                        "image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB=s60",
                                        "default_to_fiat": false,
                                        "dev_buyer_fee_basis_points": 0,
                                        "dev_seller_fee_basis_points": 0,
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": 0,
                                        "opensea_seller_fee_basis_points": 250,
                                        "buyer_fee_basis_points": 0,
                                        "seller_fee_basis_points": 250,
                                        "payout_address": null
                                    },
                                    "owner": {
                                        "user": null,
                                        "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/6.png",
                                        "address": "0x84e79d544b4b13bc3560069cfd56a9d5bbe7521d",
                                        "config": "",
                                        "discord_id": ""
                                    },
                                    "permalink": "https://opensea.io/assets/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/107221826727469155718680588460379721284635325845036369130032935095484960793482",
                                    "collection": {
                                        "banner_image_url": null,
                                        "chat_url": null,
                                        "created_date": "2019-12-10T08:55:01.591598",
                                        "default_to_fiat": false,
                                        "description": "Simplify your crypto currency payments with human readable names and build censorship resistant websites. Purchase your blockchain domains today!",
                                        "dev_buyer_fee_basis_points": "0",
                                        "dev_seller_fee_basis_points": "0",
                                        "display_data": {
                                            "card_display_style": "padded",
                                            "images": [
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/100260309576151364545964131889528085814542346820268672379352985861671742542725-1576285529.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/4421676099993643179128103015054658901819404648065294942165491120042640407039-1576285527.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/73695583292838981607767254059063047756835875089503078348607895437407313924749-1576285524.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/27205116107504397094907367426483105387345010197719784382188136305164707242210-1576285522.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/101523982898827340730362103529507400599025467135859971366386174116683457290786-1576285522.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/113871464558189091408051847721365479982421488926705310425753631895495013528420-1576285520.png"
                                            ]
                                        },
                                        "external_url": "https://unstoppabledomains.com/",
                                        "featured": false,
                                        "featured_image_url": "https://storage.googleapis.com/opensea-static/featured-images/unstoppable-domains-featured.png",
                                        "hidden": false,
                                        "safelist_request_status": "approved",
                                        "image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB=s60",
                                        "is_subject_to_whitelist": false,
                                        "large_image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB",
                                        "name": "Unstoppable Domains",
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": "0",
                                        "opensea_seller_fee_basis_points": "250",
                                        "payout_address": null,
                                        "require_email": false,
                                        "short_description": null,
                                        "slug": "unstoppable-domains",
                                        "wiki_url": null
                                    },
                                    "decimals": 0,
                                    "auctions": null,
                                    "sell_orders": [],
                                    "traits": [
                                        {
                                            "trait_type": "level",
                                            "value": 2,
                                            "display_type": null,
                                            "max_value": null,
                                            "trait_count": 0,
                                            "order": null
                                        },
                                        {
                                            "trait_type": "domain",
                                            "value": "sloth.crypto",
                                            "display_type": null,
                                            "max_value": null,
                                            "trait_count": 1,
                                            "order": null
                                        },
                                        {
                                            "trait_type": "type",
                                            "value": "standard",
                                            "display_type": null,
                                            "max_value": null,
                                            "trait_count": 52377,
                                            "order": null
                                        }
                                    ],
                                    "last_sale": null,
                                    "top_bid": null,
                                    "current_price": null,
                                    "current_escrow_price": null,
                                    "listing_date": null,
                                    "is_presale": false,
                                    "transfer_fee_payment_token": null,
                                    "transfer_fee": null
                                },
                                {
                                    "token_id": "39602885785430968515488172908991286516392379555774938600072517130912304722435",
                                    "num_sales": 0,
                                    "background_color": "4C47F7",
                                    "image_url": "https://storage.opensea.io/files/a9d70320f34fb5fb5183d1a9913d2267.svg",
                                    "image_preview_url": "https://lh3.googleusercontent.com/vbjIXUNL8ghOTEn7W6j_ALfwbFl4EOaYdqZjIwmtgdPQ92kzTd3jJjk7fsjJuW2IKkCT80JPuJ6PX_dChnOKI0yK=s250",
                                    "image_thumbnail_url": "https://lh3.googleusercontent.com/vbjIXUNL8ghOTEn7W6j_ALfwbFl4EOaYdqZjIwmtgdPQ92kzTd3jJjk7fsjJuW2IKkCT80JPuJ6PX_dChnOKI0yK=s128",
                                    "image_original_url": null,
                                    "animation_url": null,
                                    "animation_original_url": null,
                                    "name": "vikmeup",
                                    "description": "A .crypto blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.",
                                    "external_link": "https://unstoppabledomains.com/search?searchTerm=vikmeup.crypto",
                                    "asset_contract": {
                                        "address": "0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe",
                                        "asset_contract_type": "non-fungible",
                                        "created_date": "2019-12-10T08:55:01.176657",
                                        "name": ".crypto",
                                        "nft_version": "3.0",
                                        "opensea_version": null,
                                        "owner": null,
                                        "schema_name": "ERC721",
                                        "symbol": "",
                                        "total_supply": null,
                                        "description": "Simplify your crypto currency payments with human readable names and build censorship resistant websites. Purchase your blockchain domains today!",
                                        "external_link": "https://unstoppabledomains.com/",
                                        "image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB=s60",
                                        "default_to_fiat": false,
                                        "dev_buyer_fee_basis_points": 0,
                                        "dev_seller_fee_basis_points": 0,
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": 0,
                                        "opensea_seller_fee_basis_points": 250,
                                        "buyer_fee_basis_points": 0,
                                        "seller_fee_basis_points": 250,
                                        "payout_address": null
                                    },
                                    "owner": {
                                        "user": null,
                                        "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/6.png",
                                        "address": "0x84e79d544b4b13bc3560069cfd56a9d5bbe7521d",
                                        "config": "",
                                        "discord_id": ""
                                    },
                                    "permalink": "https://opensea.io/assets/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/39602885785430968515488172908991286516392379555774938600072517130912304722435",
                                    "collection": {
                                        "banner_image_url": null,
                                        "chat_url": null,
                                        "created_date": "2019-12-10T08:55:01.591598",
                                        "default_to_fiat": false,
                                        "description": "Simplify your crypto currency payments with human readable names and build censorship resistant websites. Purchase your blockchain domains today!",
                                        "dev_buyer_fee_basis_points": "0",
                                        "dev_seller_fee_basis_points": "0",
                                        "display_data": {
                                            "card_display_style": "padded",
                                            "images": [
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/100260309576151364545964131889528085814542346820268672379352985861671742542725-1576285529.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/4421676099993643179128103015054658901819404648065294942165491120042640407039-1576285527.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/73695583292838981607767254059063047756835875089503078348607895437407313924749-1576285524.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/27205116107504397094907367426483105387345010197719784382188136305164707242210-1576285522.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/101523982898827340730362103529507400599025467135859971366386174116683457290786-1576285522.png",
                                                "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/113871464558189091408051847721365479982421488926705310425753631895495013528420-1576285520.png"
                                            ]
                                        },
                                        "external_url": "https://unstoppabledomains.com/",
                                        "featured": false,
                                        "featured_image_url": "https://storage.googleapis.com/opensea-static/featured-images/unstoppable-domains-featured.png",
                                        "hidden": false,
                                        "safelist_request_status": "approved",
                                        "image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB=s60",
                                        "is_subject_to_whitelist": false,
                                        "large_image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB",
                                        "name": "Unstoppable Domains",
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": "0",
                                        "opensea_seller_fee_basis_points": "250",
                                        "payout_address": null,
                                        "require_email": false,
                                        "short_description": null,
                                        "slug": "unstoppable-domains",
                                        "wiki_url": null
                                    },
                                    "decimals": 0,
                                    "auctions": null,
                                    "sell_orders": [],
                                    "traits": [
                                        {
                                            "trait_type": "level",
                                            "value": 2,
                                            "display_type": null,
                                            "max_value": null,
                                            "trait_count": 0,
                                            "order": null
                                        },
                                        {
                                            "trait_type": "domain",
                                            "value": "vikmeup.crypto",
                                            "display_type": null,
                                            "max_value": null,
                                            "trait_count": 1,
                                            "order": null
                                        },
                                        {
                                            "trait_type": "type",
                                            "value": "standard",
                                            "display_type": null,
                                            "max_value": null,
                                            "trait_count": 52377,
                                            "order": null
                                        }
                                    ],
                                    "last_sale": null,
                                    "top_bid": null,
                                    "current_price": null,
                                    "current_escrow_price": null,
                                    "listing_date": null,
                                    "is_presale": false,
                                    "transfer_fee_payment_token": null,
                                    "transfer_fee": null
                                }
                            ]
                        }
                    `);
                }
                break;

            case 'collections':
                if (query.asset_owner == '0x84E79D544B4b13bC3560069cfD56A9D5bbE7521d') {
                    return JSON.parse(`
                        [
                            {
                                "primary_asset_contracts": [
                                    {
                                        "address": "0x1d963688fe2209a98db35c67a041524822cf04ff",
                                        "asset_contract_type": "non-fungible",
                                        "created_date": "2019-01-23T02:17:28.643618",
                                        "name": "MarbleCards",
                                        "nft_version": "3.0",
                                        "opensea_version": null,
                                        "owner": 204604,
                                        "schema_name": "ERC721",
                                        "symbol": "MRBLNFT",
                                        "total_supply": null,
                                        "description": "Claim the most amazing web pages. Remember that every web page can only be marbled once and by one person only. Once a card is created, that URL is claimed forever. Now go create some classics!",
                                        "external_link": "https://marble.cards",
                                        "image_url": "https://lh3.googleusercontent.com/JHs53JRA6f3VcBqqORnoL4_q4kLDeZkDgZkmbY3iziyQQ14IRtP3mQglePCmHpXE_fit88FH8cAFMUA3j54mivAA=s60",
                                        "default_to_fiat": false,
                                        "dev_buyer_fee_basis_points": 0,
                                        "dev_seller_fee_basis_points": 250,
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": 0,
                                        "opensea_seller_fee_basis_points": 250,
                                        "buyer_fee_basis_points": 0,
                                        "seller_fee_basis_points": 500,
                                        "payout_address": "0xee68e4c594b96efc19a9d7d2a33901651ce967a2"
                                    }
                                ],
                                "traits": {
                                    "Collection ID": {
                                        "min": 1,
                                        "max": 5865
                                    },
                                    "level": {
                                        "min": 1,
                                        "max": 100
                                    },
                                    "collection_id": {
                                        "min": 1,
                                        "max": 4033
                                    }
                                },
                                "stats": {
                                    "seven_day_volume": 1.01145277777778,
                                    "seven_day_change": 4.163107594577744,
                                    "total_volume": 350.909262241233,
                                    "count": 57728,
                                    "num_owners": 1091,
                                    "market_cap": 3029.695577899588,
                                    "average_price": 0.164129682994029,
                                    "items_sold": 2026
                                },
                                "banner_image_url": null,
                                "chat_url": null,
                                "created_date": "2019-04-26T22:13:21.421079",
                                "default_to_fiat": false,
                                "description": "Claim the most amazing web pages. Remember that every web page can only be marbled once and by one person only. Once a card is created, that URL is claimed forever. Now go create some classics!",
                                "dev_buyer_fee_basis_points": "0",
                                "dev_seller_fee_basis_points": "250",
                                "display_data": {
                                    "card_display_style": "contain",
                                    "images": [
                                        "https://storage.opensea.io/0x1d963688fe2209a98db35c67a041524822cf04ff/18612-1552667321.png",
                                        "https://storage.opensea.io/0x1d963688fe2209a98db35c67a041524822cf04ff/18905-1552770831.png",
                                        "https://storage.opensea.io/0x1d963688fe2209a98db35c67a041524822cf04ff/18704-1552750420.png",
                                        "https://storage.opensea.io/0x1d963688fe2209a98db35c67a041524822cf04ff/8829-1550886013.png",
                                        "https://storage.opensea.io/0x1d963688fe2209a98db35c67a041524822cf04ff/18951-1552772148.png",
                                        "https://storage.opensea.io/0x1d963688fe2209a98db35c67a041524822cf04ff/18950-1552772146.png"
                                    ]
                                },
                                "external_url": "https://marble.cards",
                                "featured": false,
                                "featured_image_url": "https://storage.opensea.io/0x1d963688fe2209a98db35c67a041524822cf04ff-featured-1556589463.png",
                                "hidden": false,
                                "safelist_request_status": "approved",
                                "image_url": "https://lh3.googleusercontent.com/JHs53JRA6f3VcBqqORnoL4_q4kLDeZkDgZkmbY3iziyQQ14IRtP3mQglePCmHpXE_fit88FH8cAFMUA3j54mivAA=s60",
                                "is_subject_to_whitelist": false,
                                "large_image_url": "https://lh3.googleusercontent.com/JHs53JRA6f3VcBqqORnoL4_q4kLDeZkDgZkmbY3iziyQQ14IRtP3mQglePCmHpXE_fit88FH8cAFMUA3j54mivAA",
                                "name": "MarbleCards",
                                "only_proxied_transfers": false,
                                "opensea_buyer_fee_basis_points": "0",
                                "opensea_seller_fee_basis_points": "250",
                                "payout_address": "0xee68e4c594b96efc19a9d7d2a33901651ce967a2",
                                "require_email": false,
                                "short_description": "Claim websites as unique crypto collectibles",
                                "slug": "marblecards",
                                "wiki_url": null,
                                "owned_asset_count": 1
                            },
                            {
                                "primary_asset_contracts": [
                                    {
                                        "address": "0x57f1887a8bf19b14fc0df6fd9b2acc9af147ea85",
                                        "asset_contract_type": "non-fungible",
                                        "created_date": "2019-05-08T21:59:29.327544",
                                        "name": "ENS",
                                        "nft_version": "3.0",
                                        "opensea_version": null,
                                        "owner": 279872,
                                        "schema_name": "ERC721",
                                        "symbol": "ENS",
                                        "total_supply": null,
                                        "description": "Ethereum Name Service (ENS) domains are secure domain names for the decentralized world. ENS domains provide a way for users to map human readable names to blockchain and non-blockchain resources, like Ethereum addresses, IPFS hashes, or website URLs. ENS domains can be bought and sold on secondary markets.",
                                        "external_link": "https://ens.domains",
                                        "image_url": "https://lh3.googleusercontent.com/0cOqWoYA7xL9CkUjGlxsjreSYBdrUBE0c6EO1COG4XE8UeP-Z30ckqUNiL872zHQHQU5MUNMNhfDpyXIP17hRSC5HQ=s60",
                                        "default_to_fiat": false,
                                        "dev_buyer_fee_basis_points": 0,
                                        "dev_seller_fee_basis_points": 0,
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": 0,
                                        "opensea_seller_fee_basis_points": 250,
                                        "buyer_fee_basis_points": 0,
                                        "seller_fee_basis_points": 250,
                                        "payout_address": ""
                                    }
                                ],
                                "traits": {
                                    "Length": {
                                        "min": 3,
                                        "max": 70
                                    }
                                },
                                "stats": {
                                    "seven_day_volume": 61.5400674228928,
                                    "seven_day_change": 24.471538567710834,
                                    "total_volume": 6330.17513707834,
                                    "count": 352132,
                                    "num_owners": 30272,
                                    "market_cap": 349849.79099535465,
                                    "average_price": 0.713098472127782,
                                    "items_sold": 8868
                                },
                                "banner_image_url": null,
                                "chat_url": null,
                                "created_date": "2019-05-08T21:59:36.282454",
                                "default_to_fiat": false,
                                "description": "Ethereum Name Service (ENS) domains are secure domain names for the decentralized world. ENS domains provide a way for users to map human readable names to blockchain and non-blockchain resources, like Ethereum addresses, IPFS hashes, or website URLs. ENS domains can be bought and sold on secondary markets.",
                                "dev_buyer_fee_basis_points": "0",
                                "dev_seller_fee_basis_points": "0",
                                "display_data": {
                                    "card_display_style": "cover",
                                    "images": [
                                        "https://storage.opensea.io/0xfac7bea255a6990f749363002136af6556b31e04/26693679858283927161893791132395207263838168762475401800426021835398462679008-1565293871.png",
                                        "https://storage.opensea.io/0xfac7bea255a6990f749363002136af6556b31e04/14216961695379335495094368768338390872453914418325715259759390099828279953313-1565293758.png",
                                        "https://storage.opensea.io/0xfac7bea255a6990f749363002136af6556b31e04/28032727276679833339346855127845420552876278365164540739421161968857583632995-1565293757.png",
                                        "https://storage.opensea.io/0xfac7bea255a6990f749363002136af6556b31e04/5939065732706496924433039736405736608461463576171087944642507927497598723058-1565293758.png",
                                        "https://storage.opensea.io/0xfac7bea255a6990f749363002136af6556b31e04/8268680357553423178495517182551079885250260646681574205692425339615070583014-1565293756.png",
                                        "https://storage.opensea.io/0xfac7bea255a6990f749363002136af6556b31e04/96156505551778134260224136134970011352406026760333505247707872912354314952431-1565293756.png"
                                    ]
                                },
                                "external_url": "https://ens.domains",
                                "featured": false,
                                "featured_image_url": "https://storage.googleapis.com/opensea-static/official-ens-logo.png",
                                "hidden": false,
                                "safelist_request_status": "approved",
                                "image_url": "https://lh3.googleusercontent.com/0cOqWoYA7xL9CkUjGlxsjreSYBdrUBE0c6EO1COG4XE8UeP-Z30ckqUNiL872zHQHQU5MUNMNhfDpyXIP17hRSC5HQ=s60",
                                "is_subject_to_whitelist": false,
                                "large_image_url": "https://lh3.googleusercontent.com/0cOqWoYA7xL9CkUjGlxsjreSYBdrUBE0c6EO1COG4XE8UeP-Z30ckqUNiL872zHQHQU5MUNMNhfDpyXIP17hRSC5HQ",
                                "name": "Ethereum Name Service (ENS)",
                                "only_proxied_transfers": false,
                                "opensea_buyer_fee_basis_points": "0",
                                "opensea_seller_fee_basis_points": "250",
                                "payout_address": "",
                                "require_email": false,
                                "short_description": null,
                                "slug": "ens",
                                "wiki_url": null,
                                "owned_asset_count": 2
                            },
                            {
                                "primary_asset_contracts": [
                                    {
                                        "address": "0x2fb5d7dda4f1f20f974a0fdd547c38674e8d940c",
                                        "asset_contract_type": "non-fungible",
                                        "created_date": "2019-10-24T11:06:57.511707",
                                        "name": "KnightStory Item",
                                        "nft_version": "3.0",
                                        "opensea_version": null,
                                        "owner": 1672886,
                                        "schema_name": "ERC721",
                                        "symbol": "",
                                        "total_supply": "1",
                                        "description": "Knight Story is an innovative mobile RPG powered by blockchain. The game is the second title of Biscuit; developed EOS Knights, the legendary blockchain game.",
                                        "external_link": "https://knightstory.io",
                                        "image_url": "https://lh3.googleusercontent.com/Hwr0JNz9lHdTeu3mZTawVun-BdKRf-zSpi5ZUDxirBbPs_-hW92qHfh25QcTzeGCPy0FRULooZyTJ6MlRh8qaq4=s60",
                                        "default_to_fiat": false,
                                        "dev_buyer_fee_basis_points": 0,
                                        "dev_seller_fee_basis_points": 0,
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": 0,
                                        "opensea_seller_fee_basis_points": 250,
                                        "buyer_fee_basis_points": 0,
                                        "seller_fee_basis_points": 250,
                                        "payout_address": "0x9702a479115788294232c384a5c1f42c881789fe"
                                    }
                                ],
                                "traits": {
                                    "hp": {
                                        "min": 2,
                                        "max": 5502
                                    },
                                    "defense (set)": {
                                        "min": 70,
                                        "max": 288
                                    },
                                    "attack (set)": {
                                        "min": 60,
                                        "max": 714
                                    },
                                    "score": {
                                        "min": 28,
                                        "max": 100
                                    },
                                    "luck": {
                                        "min": 2,
                                        "max": 1079
                                    },
                                    "magic_bean_purchase_bonus": {
                                        "min": 2,
                                        "max": 128
                                    },
                                    "luck (set)": {
                                        "min": 40,
                                        "max": 294
                                    },
                                    "defense": {
                                        "min": 2,
                                        "max": 2242
                                    },
                                    "hp (set)": {
                                        "min": 60,
                                        "max": 756
                                    },
                                    "None": {
                                        "min": 0,
                                        "max": 0
                                    },
                                    "attack": {
                                        "min": 7,
                                        "max": 5502
                                    }
                                },
                                "stats": {
                                    "seven_day_volume": 16.6826531170265,
                                    "seven_day_change": -0.42001600555972113,
                                    "total_volume": 555.213393323626,
                                    "count": 39491,
                                    "num_owners": 5650,
                                    "market_cap": 1592.9963396829778,
                                    "average_price": 0.0318228574152362,
                                    "items_sold": 17443
                                },
                                "banner_image_url": null,
                                "chat_url": null,
                                "created_date": "2019-10-24T11:06:58.609791",
                                "default_to_fiat": false,
                                "description": "Knight Story is an innovative mobile RPG powered by blockchain. The game is the second title of Biscuit; developed EOS Knights, the legendary blockchain game.",
                                "dev_buyer_fee_basis_points": "0",
                                "dev_seller_fee_basis_points": "0",
                                "display_data": {
                                    "card_display_style": "padded",
                                    "images": [
                                        "https://storage.opensea.io/0x2fb5d7dda4f1f20f974a0fdd547c38674e8d940c/1-1571915237.svg",
                                        "https://storage.opensea.io/0x2fb5d7dda4f1f20f974a0fdd547c38674e8d940c/2-1571982585.svg",
                                        "https://storage.opensea.io/0x2fb5d7dda4f1f20f974a0fdd547c38674e8d940c/6-1571984918.svg",
                                        "https://storage.opensea.io/0x2fb5d7dda4f1f20f974a0fdd547c38674e8d940c/5-1571984898.svg",
                                        "https://storage.opensea.io/0x2fb5d7dda4f1f20f974a0fdd547c38674e8d940c/4-1571984887.svg",
                                        "https://storage.opensea.io/0x2fb5d7dda4f1f20f974a0fdd547c38674e8d940c/3-1571984879.svg"
                                    ]
                                },
                                "external_url": "https://knightstory.io",
                                "featured": false,
                                "featured_image_url": null,
                                "hidden": false,
                                "safelist_request_status": "approved",
                                "image_url": "https://lh3.googleusercontent.com/Hwr0JNz9lHdTeu3mZTawVun-BdKRf-zSpi5ZUDxirBbPs_-hW92qHfh25QcTzeGCPy0FRULooZyTJ6MlRh8qaq4=s60",
                                "is_subject_to_whitelist": false,
                                "large_image_url": "https://lh3.googleusercontent.com/Hwr0JNz9lHdTeu3mZTawVun-BdKRf-zSpi5ZUDxirBbPs_-hW92qHfh25QcTzeGCPy0FRULooZyTJ6MlRh8qaq4",
                                "name": "KnightStory",
                                "only_proxied_transfers": false,
                                "opensea_buyer_fee_basis_points": "0",
                                "opensea_seller_fee_basis_points": "250",
                                "payout_address": "0x9702a479115788294232c384a5c1f42c881789fe",
                                "require_email": false,
                                "short_description": null,
                                "slug": "knightstory",
                                "wiki_url": null,
                                "owned_asset_count": 1
                            },
                            {
                                "primary_asset_contracts": [
                                    {
                                        "address": "0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe",
                                        "asset_contract_type": "non-fungible",
                                        "created_date": "2019-12-10T08:55:01.176657",
                                        "name": ".crypto",
                                        "nft_version": "3.0",
                                        "opensea_version": null,
                                        "owner": null,
                                        "schema_name": "ERC721",
                                        "symbol": "",
                                        "total_supply": null,
                                        "description": "Simplify your crypto currency payments with human readable names and build censorship resistant websites. Purchase your blockchain domains today!",
                                        "external_link": "https://unstoppabledomains.com/",
                                        "image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB=s60",
                                        "default_to_fiat": false,
                                        "dev_buyer_fee_basis_points": 0,
                                        "dev_seller_fee_basis_points": 0,
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": 0,
                                        "opensea_seller_fee_basis_points": 250,
                                        "buyer_fee_basis_points": 0,
                                        "seller_fee_basis_points": 250,
                                        "payout_address": null
                                    }
                                ],
                                "traits": {
                                    "level": {
                                        "min": 1,
                                        "max": 3
                                    }
                                },
                                "stats": {
                                    "seven_day_volume": 8.232,
                                    "seven_day_change": -0.43929053087580205,
                                    "total_volume": 46.0238785864631,
                                    "count": 56375,
                                    "num_owners": 4213,
                                    "market_cap": 8820.628142903694,
                                    "average_price": 0.104362536477241,
                                    "items_sold": 441
                                },
                                "banner_image_url": null,
                                "chat_url": null,
                                "created_date": "2019-12-10T08:55:01.591598",
                                "default_to_fiat": false,
                                "description": "Simplify your crypto currency payments with human readable names and build censorship resistant websites. Purchase your blockchain domains today!",
                                "dev_buyer_fee_basis_points": "0",
                                "dev_seller_fee_basis_points": "0",
                                "display_data": {
                                    "card_display_style": "padded",
                                    "images": [
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/100260309576151364545964131889528085814542346820268672379352985861671742542725-1576285529.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/4421676099993643179128103015054658901819404648065294942165491120042640407039-1576285527.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/73695583292838981607767254059063047756835875089503078348607895437407313924749-1576285524.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/27205116107504397094907367426483105387345010197719784382188136305164707242210-1576285522.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/101523982898827340730362103529507400599025467135859971366386174116683457290786-1576285522.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/113871464558189091408051847721365479982421488926705310425753631895495013528420-1576285520.png"
                                    ]
                                },
                                "external_url": "https://unstoppabledomains.com/",
                                "featured": false,
                                "featured_image_url": "https://storage.googleapis.com/opensea-static/featured-images/unstoppable-domains-featured.png",
                                "hidden": false,
                                "safelist_request_status": "approved",
                                "image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB=s60",
                                "is_subject_to_whitelist": false,
                                "large_image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB",
                                "name": "Unstoppable Domains",
                                "only_proxied_transfers": false,
                                "opensea_buyer_fee_basis_points": "0",
                                "opensea_seller_fee_basis_points": "250",
                                "payout_address": null,
                                "require_email": false,
                                "short_description": null,
                                "slug": "unstoppable-domains",
                                "wiki_url": null,
                                "owned_asset_count": 2
                            },
                            {
                                "primary_asset_contracts": [
                                    {
                                        "address": "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
                                        "asset_contract_type": "semi-fungible",
                                        "created_date": "2019-08-02T23:43:14.666153",
                                        "name": "Enjin",
                                        "nft_version": null,
                                        "opensea_version": null,
                                        "owner": null,
                                        "schema_name": "ERC1155",
                                        "symbol": "",
                                        "total_supply": null,
                                        "description": "Enjin assets are unique digital ERC1155 assets used in a variety of games in the Enjin multiverse.",
                                        "external_link": "https://enjinx.io/",
                                        "image_url": "https://lh3.googleusercontent.com/pz9RPxNoHxFTJNNySYV5bXjsWlajAiDiI1A5m5OvUaS1fd8N64yViclbRQqM8HViBTIUPrYgQ-w49h36NHL0D1Y=s60",
                                        "default_to_fiat": false,
                                        "dev_buyer_fee_basis_points": 0,
                                        "dev_seller_fee_basis_points": 0,
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": 0,
                                        "opensea_seller_fee_basis_points": 250,
                                        "buyer_fee_basis_points": 0,
                                        "seller_fee_basis_points": 250,
                                        "payout_address": null
                                    }
                                ],
                                "traits": {},
                                "stats": {
                                    "seven_day_volume": 0.6148,
                                    "seven_day_change": 3.6330067822155234,
                                    "total_volume": 284.785297517294,
                                    "count": 34146,
                                    "num_owners": 29430,
                                    "market_cap": 907.2470250000015,
                                    "average_price": 0.21673158106339,
                                    "items_sold": 1306
                                },
                                "banner_image_url": null,
                                "chat_url": null,
                                "created_date": "2019-12-15T03:51:34.864843",
                                "default_to_fiat": false,
                                "description": "The season of giving is upon us, and we come bearing gifts! Join us in the spirit of giving to receive one of our first-ever Binance Collectibles! “HO HO HODL: Binance Collectibles Series 1.”",
                                "dev_buyer_fee_basis_points": "0",
                                "dev_seller_fee_basis_points": "0",
                                "display_data": {
                                    "card_display_style": "contain",
                                    "images": [
                                        "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c/50885195465617477053098479556454830685103047942629492242239710623321420727762-1576342335.jpg",
                                        "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c/50885195465617477053098479556454830685103047942629492242239710623321420726407-1576341423.jpg",
                                        "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c/50885195465617477053098479556454830685103047942629492242239710623321420727252-1576342120.jpg",
                                        "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c/50885195465617477053098479556454830685103047942629492242239710623321420727146-1576342000.jpg",
                                        "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c/50885195465617477053098479556454830685103047942629492242239710623321420727063-1576341986.jpg",
                                        "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c/50885195465617477053098479556454830685103047942629492242239710623321420727778-1576342346.jpg"
                                    ]
                                },
                                "external_url": "https://www.binance.com/en/blog/411120468693135360/Earn-a-Guaranteed-Binance-NFT--HO-HO-HODL-Binance-Collectibles-Series-1",
                                "featured": false,
                                "featured_image_url": null,
                                "hidden": false,
                                "safelist_request_status": "approved",
                                "image_url": "https://lh3.googleusercontent.com/vbRXgbAGZVvBQw5q-qmV0tF3HHKCJeomBz5oHFTehsv2q6xuY7UyndXSWgCWqj2GGJM77DLFP-vLNVnaKYnVoD8=s60",
                                "is_subject_to_whitelist": false,
                                "large_image_url": "https://lh3.googleusercontent.com/vbRXgbAGZVvBQw5q-qmV0tF3HHKCJeomBz5oHFTehsv2q6xuY7UyndXSWgCWqj2GGJM77DLFP-vLNVnaKYnVoD8",
                                "name": "Binance",
                                "only_proxied_transfers": false,
                                "opensea_buyer_fee_basis_points": "0",
                                "opensea_seller_fee_basis_points": "250",
                                "payout_address": null,
                                "require_email": false,
                                "short_description": null,
                                "slug": "binance",
                                "wiki_url": null,
                                "owned_asset_count": 4
                            },
                            {
                                "primary_asset_contracts": [
                                    {
                                        "address": "0x3eea5bf894236f4b7a6f1451bca89a9c91f49719",
                                        "asset_contract_type": "non-fungible",
                                        "created_date": "2019-12-24T15:08:28.581742",
                                        "name": "Trust Collectible",
                                        "nft_version": "3.0",
                                        "opensea_version": null,
                                        "owner": 1982280,
                                        "schema_name": "ERC721",
                                        "symbol": "",
                                        "total_supply": "25",
                                        "description": "Your friendly talkative crypto app.",
                                        "external_link": "https://trustwallet.com",
                                        "image_url": "https://storage.opensea.io/trust-collectible-1577200601.png",
                                        "default_to_fiat": false,
                                        "dev_buyer_fee_basis_points": 0,
                                        "dev_seller_fee_basis_points": 0,
                                        "only_proxied_transfers": false,
                                        "opensea_buyer_fee_basis_points": 0,
                                        "opensea_seller_fee_basis_points": 250,
                                        "buyer_fee_basis_points": 0,
                                        "seller_fee_basis_points": 250,
                                        "payout_address": null
                                    }
                                ],
                                "traits": {
                                    "level": {
                                        "min": 2,
                                        "max": 2
                                    },
                                    "generation": {
                                        "min": 1,
                                        "max": 1
                                    }
                                },
                                "stats": {
                                    "seven_day_volume": 0,
                                    "seven_day_change": 0,
                                    "total_volume": 1.12550814873532,
                                    "count": 50,
                                    "num_owners": 22,
                                    "market_cap": 0,
                                    "average_price": 0.0229695540558228,
                                    "items_sold": 47
                                },
                                "banner_image_url": null,
                                "chat_url": null,
                                "created_date": "2019-12-24T15:08:28.833176",
                                "default_to_fiat": false,
                                "description": "Your friendly talkative crypto app.",
                                "dev_buyer_fee_basis_points": "0",
                                "dev_seller_fee_basis_points": "0",
                                "display_data": {
                                    "card_display_style": "cover",
                                    "images": [
                                        "https://storage.opensea.io/0x3eea5bf894236f4b7a6f1451bca89a9c91f49719/1-1577200113.png",
                                        "https://storage.opensea.io/0x3eea5bf894236f4b7a6f1451bca89a9c91f49719/25-1577200136.png",
                                        "https://storage.opensea.io/0x3eea5bf894236f4b7a6f1451bca89a9c91f49719/24-1577200136.png",
                                        "https://storage.opensea.io/0x3eea5bf894236f4b7a6f1451bca89a9c91f49719/21-1577200135.png",
                                        "https://storage.opensea.io/0x3eea5bf894236f4b7a6f1451bca89a9c91f49719/23-1577200135.png",
                                        "https://storage.opensea.io/0x3eea5bf894236f4b7a6f1451bca89a9c91f49719/22-1577200133.png"
                                    ]
                                },
                                "external_url": "https://trustwallet.com",
                                "featured": false,
                                "featured_image_url": null,
                                "hidden": false,
                                "safelist_request_status": "approved",
                                "image_url": "https://storage.opensea.io/trust-collectible-1577200601.png",
                                "is_subject_to_whitelist": false,
                                "large_image_url": "https://storage.opensea.io/trust-collectible-large-1577200602.png",
                                "name": "Trust Collectibles",
                                "only_proxied_transfers": false,
                                "opensea_buyer_fee_basis_points": "0",
                                "opensea_seller_fee_basis_points": "250",
                                "payout_address": null,
                                "require_email": false,
                                "short_description": null,
                                "slug": "trust-collectible",
                                "wiki_url": null,
                                "owned_asset_count": 1
                            },
                            {
                                "primary_asset_contracts": [],
                                "traits": {
                                    "level": {
                                        "min": 2,
                                        "max": 2
                                    }
                                },
                                "stats": {
                                    "seven_day_volume": 2.11109527819312,
                                    "seven_day_change": 2.279116617261758,
                                    "total_volume": 2.78059527819312,
                                    "count": 14085,
                                    "num_owners": 8511,
                                    "market_cap": 126.33769191403255,
                                    "average_price": 0.00896966218771974,
                                    "items_sold": 310
                                },
                                "banner_image_url": null,
                                "chat_url": null,
                                "created_date": "2020-04-13T20:06:38.168795",
                                "default_to_fiat": false,
                                "description": "Simplify your crypto currency payments with human readable names and build censorship resistant websites. Purchase your blockchain domains today!",
                                "dev_buyer_fee_basis_points": "0",
                                "dev_seller_fee_basis_points": "0",
                                "display_data": {
                                    "card_display_style": "padded",
                                    "images": [
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/100260309576151364545964131889528085814542346820268672379352985861671742542725-1576285529.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/4421676099993643179128103015054658901819404648065294942165491120042640407039-1576285527.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/73695583292838981607767254059063047756835875089503078348607895437407313924749-1576285524.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/27205116107504397094907367426483105387345010197719784382188136305164707242210-1576285522.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/101523982898827340730362103529507400599025467135859971366386174116683457290786-1576285522.png",
                                        "https://storage.opensea.io/0xd1e5b0ff1287aa9f9a268759062e4ab08b9dacbe/113871464558189091408051847721365479982421488926705310425753631895495013528420-1576285520.png"
                                    ]
                                },
                                "external_url": "https://unstoppabledomains.com/",
                                "featured": false,
                                "featured_image_url": "https://storage.googleapis.com/opensea-static/featured-images/unstoppable-domains-featured.png",
                                "hidden": false,
                                "safelist_request_status": "approved",
                                "image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB=s60",
                                "is_subject_to_whitelist": false,
                                "large_image_url": "https://lh3.googleusercontent.com/Ak1PwcaxSjJmX2ZR4XN1GOYw1ZqQPlJo48FJaD_RHJykq9p_-lrmLDDv2x5cjgDncxJphoSkWL4hEPA_693aXHJB",
                                "name": "Unstoppable Domains Animals",
                                "only_proxied_transfers": false,
                                "opensea_buyer_fee_basis_points": "0",
                                "opensea_seller_fee_basis_points": "250",
                                "payout_address": null,
                                "require_email": false,
                                "short_description": null,
                                "slug": "unstoppable-domains-animals",
                                "wiki_url": null,
                                "owned_asset_count": 1
                            }
                        ]
                    `);
                }
                break;
        }

        return {error: "Not implemented"};
    }
};
