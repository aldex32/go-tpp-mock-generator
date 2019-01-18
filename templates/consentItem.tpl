{
  "Item": {
    "access": {
      "M": {
        "balances": {
          "L": [
            {
              "M": {
                "currency": {
                  "S": "EUR"
                },
                "iban": {
                  "S": "{{.Iban}}"
                }
              }
            }
          ]
        },
        "transactions": {
          "L": [
            {
              "M": {
                "currency": {
                  "S": "EUR"
                },
                "iban": {
                  "S": "{{.Iban}}"
                }
              }
            }
          ]
        }
      }
    },
    "combinedServiceIndicator": {
      "BOOL": false
    },
    "consentId": {
      "S": "{{.ConsentId}}"
    },
    "consentStatus": {
      "S": "{{.Status}}"
    },
    "created": {
      "S": "2019-01-15T15:06:52.920Z"
    },
    "frequencyPerDay": {
      "N": "{{.FrequencyPerDay}}"
    },
    "lastUpdated": {
      "S": "{{.LastUpdated}}"
    },
    "recurringIndicator": {
      "BOOL": {{.RecurringIndicator}}
    },
    "requestId": {
      "S": "{{.RequestId}}"
    },
    "tppAuthorizationNumber": {
      "S": "{{.TppAuthorizationNumber}}"
    },
    "tppNokRedirectUri": {
      "S": "https://www.tpp.com"
    },
    "tppRedirectUri": {
      "S": "https://www.tpp.com"
    },
    "transactionStatusDetailed": {
      "S": "{{.TransactionStatusDetailed}}"
    },
    "validUntil": {
      "S": "{{.ValidUntil}}"
    }
  }
}