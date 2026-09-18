# Shiprocket API — Account, Wallet & Billing

> Generated from the published Postman collection behind [https://apidocs.shiprocket.in/](https://apidocs.shiprocket.in/) (collection "Shiprocket API", published id `SzYW1zB2`, version tag `latest`). Only content present in that collection is reproduced here; where the source omits something, the omission is stated.

## Account

_No folder-level description in the source collection._

### Get Wallet Balance

`GET https://apiv2.shiprocket.in/v1/external/account/details/wallet-balance`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to get the Wallet balance details of your Shiprocket account. No parameters are required to access this API.

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": {
        "balance_amount": "-291539.83"
    }
}
```

## Statement Details

Get your Shiprocket account's statement details using this API.

### Get Statement Details

`GET https://apiv2.shiprocket.in/v1/external/account/details/statement`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to get the account statement details of your Shiprocket account. No parameters are required to access this API. However, you sort and filter the data displayed using the optional parameters.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| ` page` | NO | *integer* | The page number you want to display. | 5 |
| `per_page` | NO | *integer* | The number of orders to get per request. | 2 |
| `from` | NO | *string* | From a specific date. | 2017-08-12 |
| `to` | NO | *string* | To a specific date. | 2017-09-12 |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": [
        {
            "transaction_id": "",
            "order_id": "",
            "channel_order_id": "",
            "awb_code": "",
            "return_awb_code": null,
            "applied_weight": "",
            "charged_weight": "",
            "billed_weight": "",
            "action": "",
            "charge": "",
            "description": "Wallet Balance",
            "debit_amount": "",
            "credit_amount": "",
            "balance_amount": "0",
            "balance_weight": 0,
            "volumetric_weight": "",
            "entered_weight": "",
            "created_at": "",
            "can_ship": true
        }
    ]
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/account/details/statemen`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

## Discrepancy Details

Use this API to get the discrepancy details of your account.

### Get Dicrepancy Data

`GET https://apiv2.shiprocket.in/v1/external/billing/discrepancy`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Get the discrepancy data associated with your account, if any. No parameters are required to use this API.

The data is displayed in JSON format.

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "status": 200,
    "data": [],
    "upper_fold_text": "The entered weight for the below given shipment was incorrect. Please allow us to make deduction on the basis of correct charged weight as shared by the courier company.",
    "lower_fild_text": "To raise a ticket for your queries, please call us on 011-39595108 or email us at support@shiprocket.in"
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/billing/discrepanc`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```
