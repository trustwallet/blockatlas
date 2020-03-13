/// Ripple API Mock
/// See:
/// curl "http://localhost:3000/ripple-api/accounts/rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1/transactions?type=Payment&descending=false&limit=25"
/// curl "https://data.ripple.com/v2/accounts/rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1/transactions?type=Payment&descending=false&limit=25"
/// curl http://localhost:8420/v1/ripple/rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1
module.exports = {
    path: "/ripple-api/accounts/:address/transactions?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        if (params.address === 'rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1') {
            return JSON.parse(`
                {
                    "result": "success",
                    "count": 2,
                    "transactions": [
                        {
                            "hash": "20F2F95E4612296B1A82CC72F0DC53C9EAA8DA1557A0D0AD6C1E3BCA7E67E7CE",
                            "ledger_index": 44734671,
                            "date": "2019-01-28T19:22:50+00:00",
                            "tx": {
                                "TransactionType": "Payment",
                                "Flags": 2147483648,
                                "Sequence": 13,
                                "LastLedgerSequence": 44913253,
                                "Amount": "392834642660000",
                                "Fee": "500",
                                "SigningPubKey": "",
                                "Account": "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1",
                                "Destination": "rUEchj8k8jvVXKGhfVYvrs5ntVEtc5F2hK",
                                "Signers": [
                                    {
                                        "Signer": {
                                            "SigningPubKey": "03807050F9E271B2E49B0FF658362EF37DBFDD31435E610B6E11C52879DF8A9907",
                                            "TxnSignature": "30440220247A6624588612C73EEC2CD5F8D0A74E7FF7EE36E1ECFFF413703E928654824702205E9F5DB179A9146331935C5B5762AAAE52AE09EFDDEFDA6C596A53AE870464D2",
                                            "Account": "rGSxFjoqmWz54PycrgQBQ5dB6e7TUpMxzq"
                                        }
                                    },
                                    {
                                        "Signer": {
                                            "SigningPubKey": "029A7A9E7A6175690C6E3D46EBBA33DA9EBF096EB5942E21C4A8C98F5606E6E6B8",
                                            "TxnSignature": "304402201DFB1E6C459753DAAD5AA6D0D44501D5D1B067F18733E2DE23F3DA9E6F4A03E602201BD890DCDDA34E8304C7C29209EAAB3C34DE07C9FD4CF2537848EA4CA02FA1BB",
                                            "Account": "rHH2gS6XikKoEt9i1xAxEBvXLHKFr7oLn8"
                                        }
                                    },
                                    {
                                        "Signer": {
                                            "SigningPubKey": "03797EFEC20C8DBF46C20FCC75AB22168AAC34D9E18BFD81C3CC007DC80C389686",
                                            "TxnSignature": "304402205E1D7DE31B9E88848E370AD7D5D77FF16CED0807A3514918BA41718040A9721F02207A832DA5575C0B66B9FF0AC6C1C0BFAF15359727221A97D2027A69CFDA953297",
                                            "Account": "rMY6Wm2RWQLN4d3Jjz15MKP74GJWVQE2pb"
                                        }
                                    },
                                    {
                                        "Signer": {
                                            "SigningPubKey": "0368CFF85FEF68857A445A3F51FFC6119F2A113C9590E898351CE49F61BA71E6E8",
                                            "TxnSignature": "304402204BB844E6102A40079C9616AB888C7332F000353665B4FA972525B261DB3DDED6022018586112E834D6BC3A2A2151A4F39304B1357EB11617EE84ED55D862D7978E7C",
                                            "Account": "rP5xpZ5KzPih69fLhG3NYvZEDfLmSEViUk"
                                        }
                                    }
                                ]
                            },
                            "meta": {
                                "TransactionIndex": 2,
                                "AffectedNodes": [
                                    {
                                        "ModifiedNode": {
                                            "LedgerEntryType": "AccountRoot",
                                            "PreviousTxnLgrSeq": 44734265,
                                            "PreviousTxnID": "AB83FFE7B46FF4E44BC084DFBEBBCFAA09913770766CCBFC0B1858E8195CE94E",
                                            "LedgerIndex": "459E688FB21500B75D59CBFA49EB4C17C54E81D32A7A69CF1910691AB6AA4DCC",
                                            "PreviousFields": {
                                                "Balance": "392834747660000"
                                            },
                                            "FinalFields": {
                                                "Flags": 0,
                                                "Sequence": 1,
                                                "OwnerCount": 0,
                                                "Balance": "785669390320000",
                                                "Account": "rUEchj8k8jvVXKGhfVYvrs5ntVEtc5F2hK"
                                            }
                                        }
                                    },
                                    {
                                        "ModifiedNode": {
                                            "LedgerEntryType": "AccountRoot",
                                            "PreviousTxnLgrSeq": 44734265,
                                            "PreviousTxnID": "AB83FFE7B46FF4E44BC084DFBEBBCFAA09913770766CCBFC0B1858E8195CE94E",
                                            "LedgerIndex": "564241023DCB6F74760910F17F78B179AEC159C701BBACD99A1D3259D77D3CFF",
                                            "PreviousFields": {
                                                "Sequence": 13,
                                                "Balance": "4123164791269591"
                                            },
                                            "FinalFields": {
                                                "Flags": 1048576,
                                                "Sequence": 14,
                                                "OwnerCount": 10,
                                                "Balance": "3730330148609091",
                                                "Account": "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1"
                                            }
                                        }
                                    }
                                ],
                                "TransactionResult": "tesSUCCESS",
                                "delivered_amount": "392834642660000"
                            }
                        },
                        {
                            "hash": "000BC2AFC047BE45C43886109720F44B6F3F2434D59B1A16A9B2B4FC9B9C5A13",
                            "ledger_index": 51969362,
                            "date": "2019-12-11T00:00:01+00:00",
                            "tx": {
                                "TransactionType": "Payment",
                                "Flags": 2147483648,
                                "Sequence": 14,
                                "LastLedgerSequence": 52151515,
                                "Amount": "220303137120000",
                                "Fee": "5000",
                                "SigningPubKey": "",
                                "Account": "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1",
                                "Destination": "r3a8tn1ubcP13np3giaLYzkmVKfnhLP2BL",
                                "Signers": [
                                    {
                                        "Signer": {
                                            "SigningPubKey": "03807050F9E271B2E49B0FF658362EF37DBFDD31435E610B6E11C52879DF8A9907",
                                            "TxnSignature": "3045022100BC1FBB9457B5A8684AACA1E949DB2640CE1D0C9907AA060D334655F4162D27D0022028DD17C6E2AAD9863D2FA7BE7F0D80D4181CB9B9349371959E5DB2D69BF662D3",
                                            "Account": "rGSxFjoqmWz54PycrgQBQ5dB6e7TUpMxzq"
                                        }
                                    },
                                    {
                                        "Signer": {
                                            "SigningPubKey": "029A7A9E7A6175690C6E3D46EBBA33DA9EBF096EB5942E21C4A8C98F5606E6E6B8",
                                            "TxnSignature": "3044022014FA8252577C836D14D796B59C74F377A23EE1D335AEB68ED21A4BF5D4D05BBE022054361EEC50A8F3C4A4AF4EFE61C83B399F7B5240C42C2FD4BBBCEAB1CFF138B4",
                                            "Account": "rHH2gS6XikKoEt9i1xAxEBvXLHKFr7oLn8"
                                        }
                                    },
                                    {
                                        "Signer": {
                                            "SigningPubKey": "03797EFEC20C8DBF46C20FCC75AB22168AAC34D9E18BFD81C3CC007DC80C389686",
                                            "TxnSignature": "3045022100B407D26B1A4ABAA6C0BC9C99EDB0DCB280852EF21974D82C8C8E5866FB08187A02207255B8CD26CCEFD8ACA1DD63A767CCC127AE6EA2115CC2F89D1B798EBFA3316C",
                                            "Account": "rMY6Wm2RWQLN4d3Jjz15MKP74GJWVQE2pb"
                                        }
                                    },
                                    {
                                        "Signer": {
                                            "SigningPubKey": "0368CFF85FEF68857A445A3F51FFC6119F2A113C9590E898351CE49F61BA71E6E8",
                                            "TxnSignature": "3045022100A0A5FD62D9EC81D1F0C5474A7EEB31C45423D647764B352F947617C383CD9E6902201896720DB3B9876FE56DEA05E4F75B3F90A20B3A5EF95E39703F80F5166BD1F6",
                                            "Account": "rP5xpZ5KzPih69fLhG3NYvZEDfLmSEViUk"
                                        }
                                    }
                                ]
                            },
                            "meta": {
                                "TransactionIndex": 0,
                                "AffectedNodes": [
                                    {
                                        "ModifiedNode": {
                                            "LedgerEntryType": "AccountRoot",
                                            "PreviousTxnLgrSeq": 51901464,
                                            "PreviousTxnID": "607A53B712B938B1475D194EC4F3318E6EC4BACD430D6C5437AA4018F1754A2A",
                                            "LedgerIndex": "45806910346E5E79AE202994014742DEDA552A6197C19E03C4DC5C5466A00DB2",
                                            "PreviousFields": {
                                                "Balance": "120000000"
                                            },
                                            "FinalFields": {
                                                "Flags": 0,
                                                "Sequence": 2,
                                                "OwnerCount": 0,
                                                "Balance": "220303257120000",
                                                "Account": "r3a8tn1ubcP13np3giaLYzkmVKfnhLP2BL"
                                            }
                                        }
                                    },
                                    {
                                        "ModifiedNode": {
                                            "LedgerEntryType": "AccountRoot",
                                            "PreviousTxnLgrSeq": 44734671,
                                            "PreviousTxnID": "20F2F95E4612296B1A82CC72F0DC53C9EAA8DA1557A0D0AD6C1E3BCA7E67E7CE",
                                            "LedgerIndex": "564241023DCB6F74760910F17F78B179AEC159C701BBACD99A1D3259D77D3CFF",
                                            "PreviousFields": {
                                                "Sequence": 14,
                                                "Balance": "3730330148609091"
                                            },
                                            "FinalFields": {
                                                "Flags": 1048576,
                                                "Sequence": 15,
                                                "OwnerCount": 10,
                                                "Balance": "3510027011484091",
                                                "Account": "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1"
                                            }
                                        }
                                    }
                                ],
                                "TransactionResult": "tesSUCCESS",
                                "delivered_amount": "220303137120000"
                            }
                        }
                    ]
                }
            `);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
