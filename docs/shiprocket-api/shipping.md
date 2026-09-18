# Shiprocket API — Couriers, Shipments, Tracking & NDR

> Generated from the published Postman collection behind [https://apidocs.shiprocket.in/](https://apidocs.shiprocket.in/) (collection "Shiprocket API", published id `SzYW1zB2`, version tag `latest`). Only content present in that collection is reproduced here; where the source omits something, the omission is stated.

## Couriers

Use these APIs to assign AWB to your order, check for courier serviceability, and request for the pickup of your order.

All available couriers, along with their codes, are mentioned below:

| **COURIER ID** | **COURIER NAME** |
|---|---|
| 1 | Blue Dart |
| 4 | Amazon Shipping 5Kg |
| 6 | DTDC Surface |
| 10 | Delhivery |
| 14 | Ecom Express Surface |
| 18 | DTDC 5kg |
| 19 | Ecom Express Surface 2kg |
| 23 | Xpressbees 1kg |
| 24 | Xpressbees 2kg |
| 25 | Xpressbees 5kg |
| 29 | Amazon Shipping 1Kg |
| 32 | Amazon Shipping 2Kg |
| 33 | Xpressbees |
| 35 | Aramex International |
| 39 | Delhivery Surface 5 Kgs |
| 43 | Delhivery Surface |
| 45 | Ecom Express Reverse |
| 46 | Shadowfax Reverse |
| 48 | Ekart Logistics |
| 51 | Xpressbees Surface |
| 54 | Ekart Logistics Surface |
| 55 | Blue Dart Surface |
| 58 | Shadowfax Surface |
| 60 | Ecom Premium and ROS |
| 61 | Delhivery Reverse |
| 64 | Delhivery 500G Savex (A) |
| 65 | Delhivery 500G Savex (S) |
| 66 | Ecom Exp Savex (A) |
| 67 | Ecom ROS Savex (A) |
| 69 | Kerry Indev Express Surface |
| 71 | Self Delivery |
| 72 | Bluedart Savex (A) |
| 75 | SRF Standard 500gm |
| 76 | Delhivery 5KG Savex (S) |
| 77 | Delhivery 10KG Savex (S) |
| 78 | Delhivery 20KG Savex (S) |
| 79 | SRF Standard 2kg |
| 81 | SRF Standard |
| 82 | DTDC 2kg |
| 83 | Delhivery 500G Savex RVP |
| 84 | Delhivery 500G Savex RVP-QC |
| 85 | Standard |
| 86 | Express |
| 87 | Delhivery 5KG Savex RVP (S) |
| 88 | Delhivery 10KG Savex RVP (S) |
| 89 | Delhivery 20KG Savex RVP (S) |
| 90 | Standard 500gm |
| 91 | Standard 2kg |
| 92 | Standard 5kg |
| 93 | Standard 10kg |
| 94 | Standard 20kg |
| 95 | Shadowfax Local |
| 97 | Dunzo Local |
| 98 | SRF Standard 5kg |
| 99 | Ecom Express ROS Reverse |
| 100 | Delhivery Surface 10 Kgs |
| 101 | Delhivery Surface 20 Kgs |
| 106 | Borzo |
| 107 | Borzo 5 Kg |
| 114 | SRF Standard 10kg |
| 115 | Shadowfax Savex |
| 116 | SRF Standard 20kg |
| 117 | SRF Express |
| 118 | Bluedart Savex Exchange |
| 119 | Identify Plus SDD Lite |
| 120 | Identify Plus SDD Standard |
| 125 | Xpressbees Reverse |
| 126 | Bluedart Savex TDD |
| 130 | Ekart 10kg |
| 138 | Delhivery Reverse 5kg |
| 140 | SRX Premium |
| 141 | SRF RUSH |
| 142 | Amazon Surface 500gm Prepaid |
| 144 | Xpressbees Reverse 2kg |
| 145 | Xpressbees Reverse 5 kg |
| 146 | Kerry Indev 2kg Surface |
| 150 | Xpressbees Reverse 1kg |
| 154 | Kerry Indev Express |
| 159 | Xpressbees 10kg |
| 170 | Ekart 2Kg |
| 171 | Ekart 5Kg |
| 181 | Amazon Shipping 10Kg |
| 182 | Amazon Shipping 20Kg |
| 195 | Amazon Surface 500gm COD |
| 196 | DTDC 500GMS |
| 235 | DTDC Express |
| 240 | SRX Priority |
| 313 | Wholemark |
| 501 | Bluedart CE Savex |
| 524 | Sm Shadowfax Surface 3Kg |
| 777 | Shadowfax DS |

### Generate AWB for Shipment

`POST https://apiv2.shiprocket.in/v1/external/courier/assign/awb`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API can be used to assign the AWB (Air Waybill Number) to your shipment. The AWB is a unique number that helps you track the shipment and get details about it.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `shipment_id` | YES | *integer* | The shipment id of the order you want to create the AWB for. | 16016920 |
| `courier_id` | NO | *integer* | The courier id of the courier service you want to select. The default courier is selected in case no id is specified. | 10 |
| `status` | NO | *string* | Use this to change the courier of a shipment. Value: reassign. Note that this can be done only once in 24 hours. | reassign |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "shipment_id": "",
  "courier_id": "",
  "status": ""
 
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
    "shipment_id": "16090281",
    "courier_id": "10"
}
```

```json
{
    "awb_assign_status": 1,
    "response": {
        "data": {
            "courier_company_id": 142,
            "awb_code": "321055706540",
            "cod": 0,
            "order_id": 281248157,
            "shipment_id": 16090281,
            "awb_code_status": 1,
            "assigned_date_time": {
                "date": "2022-11-25 11:17:52.878599",
                "timezone_type": 3,
                "timezone": "Asia/Kolkata"
            },
            "applied_weight": 0.5,
            "company_id": 25149,
            "courier_name": "Amazon Surface",
            "child_courier_name": null,
            "pickup_scheduled_date": "2022-11-25 14:00:00",
            "routing_code": "",
            "rto_routing_code": "",
            "invoice_no": "retail5769122647118",
            "transporter_id": "",
            "transporter_name": "",
            "shipped_by": {
                "shipper_company_name": "manoj",
                "shipper_address_1": "Aligarh",
                "shipper_address_2": "noida",
                "shipper_city": "Jammu",
                "shipper_state": "Jammu & Kashmir",
                "shipper_country": "India",
                "shipper_postcode": "110030",
                "shipper_first_mile_activated": 0,
                "shipper_phone": "8976967989",
                "lat": "32.731899",
                "long": "74.860376",
                "shipper_email": "hdhd@gshd.com",
                "rto_company_name": "test",
                "rto_address_1": "Unnamed Road, Bengaluru, Karnataka 560060, India",
                "rto_address_2": "Katrabrahmpur",
                "rto_city": "Bangalore",
                "rto_state": "Karnataka",
                "rto_country": "India",
                "rto_postcode": "560060",
                "rto_phone": "9999999999",
                "rto_email": "test@test.com"
            }
        }
    }
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

Example request body:

```json
{
    "shipment_id": "",
    "courier_id": ""
}
```

```json
{
    "message": "Required field missing",
    "errors": {
        "shipment_id": [
            "The shipment id field is required."
        ]
    },
    "status_code": 422
}
```

#### Invalid Data — HTTP 400 Bad Request

Example request body:

```json
{
    "shipment_id": "11111111"
}
```

```json
{
    "message": "Oops! Cannot reassign courier for this shipment.",
    "status_code": 400
}
```

### List of Couriers

`GET https://apiv2.shiprocket.in/v1/external/courier/courierListWithCounts`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to check the list of couriers and the related information with Shiprocket based on the search criteria.<br>**Note:**

- This API will work on a company level.
- You can use filters to sort your data. By default, all the couriers are shown if no filter is used.
- total_courier_count will change based on the filter used.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `type` | NO | *string* | Use this as a parameter. Possible values are : active/inactive/all | [https://apiv2.shiprocket.in/v1/external/courier/courierListWithCounts?type=active](https://apiv2.shiprocket.in/v1/external/courier/courierListWithCounts?type=active') |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
"total_courier_count": 72,
"serviceable_pincodes_count": 26416,
"pickup_pincodes_count": 25501,
"total_rto_count": 3380,
"total_oda_count": 2032,
"courier_count": 12,
"courier_data": [
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 53,
"min_weight": 1,
"base_courier_id": 53,
"name": "Gati Surface 1 Kg",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "Gati Surface 1 Kg",
"service_type": 1,
"mode": 0,
"image": {
"logo": "post_order/img/courier/GATI-SURFACE.png",
"small_logo": "post_order/img/courier/thumb/gati.jpg",
"email_logo_s3_path": "gati.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "On Request",
"call_before_delivery": "Available",
"activated_date": "2021-10-14",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 40,
"min_weight": 5,
"base_courier_id": 40,
"name": "Gati Surface 5 Kg",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "Gati Surface 5 Kg",
"service_type": 1,
"mode": 0,
"image": {
"logo": "post_order/img/courier/GATI-SURFACE.png",
"small_logo": "post_order/img/courier/thumb/gati.jpg",
"email_logo_s3_path": "gati.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "On Request",
"call_before_delivery": "Available",
"activated_date": "2021-10-14",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 50,
"min_weight": 0.5,
"base_courier_id": 50,
"name": "Wow Express",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "Wow Express",
"service_type": 1,
"mode": 1,
"image": {
"logo": "post_order/img/courier/wow_express.png",
"small_logo": "post_order/img/courier/thumb/wow_express.jpg",
"email_logo_s3_path": "easyco.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "On Request",
"call_before_delivery": "Available",
"activated_date": "2021-10-14",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 16,
"min_weight": 0.5,
"base_courier_id": 16,
"name": "Dotzot",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "Dotzot",
"service_type": 1,
"mode": 1,
"image": {
"logo": "post_order/img/courier/DOTZOT.png",
"small_logo": "post_order/img/courier/thumb/dotzot.jpg",
"email_logo_s3_path": "dozot.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "On Request",
"call_before_delivery": "Available",
"activated_date": "2021-10-14",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 3,
"min_weight": 0,
"base_courier_id": 3,
"name": "ARAMEX",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "ARAMEX",
"service_type": 1,
"mode": 0,
"image": {
"logo": "post_order/img/courier/ARAMEX.png",
"small_logo": "post_order/img/courier/thumb/aramex.jpg",
"email_logo_s3_path": "aramax.png"
},
"realtime_tracking": "MIS",
"delivery_boy_contact": "Not Available",
"pod_available": "On Request",
"call_before_delivery": "Not Available",
"activated_date": "2021-10-14",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 2,
"min_weight": 0.5,
"base_courier_id": 2,
"name": "FedEx",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "FedEx",
"service_type": 1,
"mode": 1,
"image": {
"logo": "post_order/img/courier/FEDEX.png",
"small_logo": "post_order/img/courier/thumb/fedex.jpg",
"email_logo_s3_path": "fedex.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "Instant",
"call_before_delivery": "Available",
"activated_date": "2021-06-21",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 41,
"min_weight": 0.5,
"base_courier_id": 41,
"name": "FedEx Flat Rate",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "FedEx Flat Rate",
"service_type": 1,
"mode": 1,
"image": {
"logo": "post_order/img/courier/FEDEX.png",
"small_logo": "post_order/img/courier/thumb/fedex.jpg",
"email_logo_s3_path": "fedex.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "Instant",
"call_before_delivery": "Available",
"activated_date": "2021-06-21",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 12,
"min_weight": 10,
"base_courier_id": 12,
"name": "FedEx Surface 10 Kg",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "FedEx Surface 10 Kg",
"service_type": 1,
"mode": 0,
"image": {
"logo": "post_order/img/courier/FEDEX-SURFACE.png",
"small_logo": "post_order/img/courier/thumb/fedex.jpg",
"email_logo_s3_path": "fedex.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "Instant",
"call_before_delivery": "Available",
"activated_date": "2021-06-21",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 42,
"min_weight": 5,
"base_courier_id": 42,
"name": "FedEx Surface 5 Kg",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "FedEx Surface 5 Kg",
"service_type": 1,
"mode": 0,
"image": {
"logo": "post_order/img/courier/FEDEX.png",
"small_logo": "post_order/img/courier/thumb/fedex.jpg",
"email_logo_s3_path": "fedex.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "Instant",
"call_before_delivery": "Available",
"activated_date": "2021-06-21",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 62,
"min_weight": 1,
"base_courier_id": 62,
"name": "FedEx Surface 1 Kg",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "FedEx Surface 1 Kg",
"service_type": 1,
"mode": 0,
"image": {
"logo": "post_order/img/courier/FEDEX-SURFACE.png",
"small_logo": "post_order/img/courier/thumb/fedex.jpg",
"email_logo_s3_path": "fedex.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Not Available",
"pod_available": "Instant",
"call_before_delivery": "Available",
"activated_date": "2021-06-21",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 1,
"min_weight": 0.5,
"base_courier_id": 1,
"name": "Blue Dart",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "Blue Dart",
"service_type": 1,
"mode": 1,
"image": {
"logo": "post_order/img/courier/BLUEDART.png",
"small_logo": "post_order/img/courier/thumb/bluedart.jpg",
"email_logo_s3_path": "Bluedart.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Available",
"pod_available": "Instant",
"call_before_delivery": "Available",
"activated_date": "2021-03-12",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
},
{
"is_own_key_courier": 0,
"ownkey_courier_id": 0,
"id": 55,
"min_weight": 0.5,
"base_courier_id": 55,
"name": "Blue Dart Surface",
"use_sr_postcodes": 1,
"type": 1,
"status": 1,
"courier_type": 0,
"master_company": "Blue Dart Surface",
"service_type": 1,
"mode": 0,
"image": {
"logo": "post_order/img/courier/DARTPLUS.png",
"small_logo": "post_order/img/courier/thumb/bluedart.jpg",
"email_logo_s3_path": "Bluedart.png"
},
"realtime_tracking": "Real Time",
"delivery_boy_contact": "Available",
"pod_available": "Instant",
"call_before_delivery": "Available",
"activated_date": "2021-03-12",
"newest_date": null,
"shipment_count": "",
"is_hyperlocal": 0
}
]
}
```

### Check Courier Serviceability

`GET https://apiv2.shiprocket.in/v1/external/courier/serviceability/`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to check the availability of couriers between the pickup and delivery postal codes.<br>Further details like the estimated time of delivery, the rates along with the ids are also shown.

**Note:**

- One of either the 'order_id' or 'cod' and 'weight' is required. If you specify the order id, the cod and weight fields are not required and vice versa.
- You can add further fields to add the shipment details and filter the search.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `pickup_postcode` | YES | *integer* | Postcode from where the order will be picked. | 110030 |
| `delivery_postcode` | YES | *integer* | Postcode where the order will be delivered | 122002 |
| `order_id` | NO | *integer* | If order id is already created in Shiprocket panel then you can use this shiprocket order id in servicibility | 123456 |
| `cod` | CONDITIONAL YES | *boolean* | 1 for Cash on Delivery and 0 for Prepaid orders. | 1 |
| `weight` | CONDITIONAL YES | *string* | The weight of shipment in kgs. | 2 |
| `length` | NO | *integer* | The length of the shipment in cms. | 15 |
| `breadth` | NO | *integer* | The breadth of the shipment in cms. | 10 |
| `height` | NO | *integer* | The height of the shipment in cms. | 5 |
| `declared_value` | NO | *integer* | The price of the order shipment in rupees. | 50 |
| `mode` | NO | *string* | The mode of travel. Either: Surface or Air | Air |
| `is_return` | NO | *integer* | Whether the order is a return order or not. 1 in case of Yes and 0 for No. (declared_value field is required in case you use this parameter) | 0 |
| `couriers_type` | NO | *integer* | Use this to filter out and show only "documents" couriers like XB documents, etc. The only accepted value is 1. | 1 |
| `only_local` | NO | *integer* | Use this to filter out and show only Hyperlocal couriers. The only accepted value is 1. | 1 |
| `qc_check` | CONDITIONAL YES | *integer* | Use this filter to show only the QC-enabled couriers. is_return has to be set to 1 | 1 |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "company_auto_shipment_insurance_setting": true,
    "covid_zones": {
        "delivery_zone": null,
        "pickup_zone": null
    },
    "currency": "INR",
    "data": {
        "available_courier_companies": [
            {
                "air_max_weight": "0.00",
                "assured_amount": 0,
                "base_courier_id": null,
                "base_weight": "",
                "blocked": 0,
                "call_before_delivery": "Available",
                "charge_weight": 0.5,
                "city": "Mandi",
                "cod": 1,
                "cod_charges": 0,
                "cod_multiplier": 0.01,
                "cost": "",
                "courier_company_id": 43,
                "courier_name": "Delhivery Surface",
                "courier_type": "0",
                "coverage_charges": 0,
                "cutoff_time": "11:00",
                "delivery_boy_contact": "Not Available",
                "delivery_performance": 5,
                "description": "",
                "edd": "",
                "entry_tax": 0,
                "estimated_delivery_days": "4",
                "etd": "Jul 01, 2024",
                "etd_hours": 91,
                "freight_charge": 54,
                "id": 459245934,
                "is_custom_rate": 1,
                "is_hyperlocal": false,
                "is_international": 0,
                "is_rto_address_available": true,
                "is_surface": true,
                "local_region": 0,
                "metro": 0,
                "min_weight": 0.5,
                "mode": 0,
                "new_edd": 0,
                "odablock": false,
                "other_charges": 0,
                "others": "{\"allow_postcode_auto_sync\":1,\"cancel_real_time\":true}",
                "pickup_availability": "0",
                "pickup_performance": 4.7,
                "pickup_priority": "",
                "pickup_supress_hours": 0,
                "pod_available": "Instant",
                "postcode": "175019",
                "qc_courier": 0,
                "rank": "",
                "rate": 54,
                "rating": 4.9,
                "realtime_tracking": "Real Time",
                "region": 1,
                "rto_charges": 54,
                "rto_performance": 5,
                "seconds_left_for_pickup": 0,
                "secure_shipment_disabled": false,
                "ship_type": 1,
                "state": "Himachal Pradesh",
                "suppress_date": "",
                "suppress_text": "",
                "suppression_dates": null,
                "surface_max_weight": "4.00",
                "tracking_performance": 5,
                "volumetric_max_weight": null,
                "weight_cases": 4.6,
                "zone": "z_e"
            },
            {
                "air_max_weight": "0.00",
                "assured_amount": 0,
                "base_courier_id": null,
                "base_weight": "",
                "blocked": 0,
                "call_before_delivery": "Available",
                "charge_weight": 2,
                "city": "MANDI",
                "cod": 1,
                "cod_charges": 0,
                "cod_multiplier": 0,
                "cost": "",
                "courier_company_id": 225,
                "courier_name": "India Post-Business Parcel Surface Prepaid",
                "courier_type": "0",
                "coverage_charges": 8.71,
                "cutoff_time": "10:00",
                "delivery_boy_contact": "Not Available",
                "delivery_performance": 4.5,
                "description": "",
                "edd": "",
                "entry_tax": 0,
                "estimated_delivery_days": "8",
                "etd": "Jul 05, 2024",
                "etd_hours": 184,
                "freight_charge": 123.9,
                "id": 396397972,
                "is_custom_rate": 0,
                "is_hyperlocal": false,
                "is_international": 0,
                "is_rto_address_available": true,
                "is_surface": true,
                "local_region": 0,
                "metro": 0,
                "min_weight": 2,
                "mode": 0,
                "new_edd": 0,
                "odablock": false,
                "other_charges": 0,
                "others": "{\"cancel_real_time\":true}",
                "pickup_availability": "0",
                "pickup_performance": 4.4,
                "pickup_priority": "",
                "pickup_supress_hours": 0,
                "pod_available": "On Request",
                "postcode": "175019",
                "qc_courier": 0,
                "rank": "",
                "rate": 123.9,
                "rating": 4.4,
                "realtime_tracking": "Real Time",
                "region": 5,
                "rto_charges": 0,
                "rto_performance": 4.4,
                "seconds_left_for_pickup": 0,
                "secure_shipment_disabled": false,
                "ship_type": 1,
                "state": "HIMACHAL PRADESH",
                "suppress_date": "",
                "suppress_text": "",
                "suppression_dates": {
                    "action_on": "2023-11-11 20:29:04",
                    "delay_remark": "Festival",
                    "delivery_delay_by": 2065762,
                    "delivery_delay_days": "1",
                    "delivery_delay_from": "2023-11-13",
                    "delivery_delay_to": "2023-11-13",
                    "pickup_delay_by": 2065762,
                    "pickup_delay_days": "1",
                    "pickup_delay_from": "2023-11-13",
                    "pickup_delay_to": "2023-11-13"
                },
                "surface_max_weight": "35.00",
                "tracking_performance": 4.5,
                "volumetric_max_weight": 35,
                "weight_cases": 4.2,
                "zone": "z_e"
            },
            {
                "air_max_weight": "0.00",
                "assured_amount": 0,
                "base_courier_id": null,
                "base_weight": "",
                "blocked": 0,
                "call_before_delivery": "Available",
                "charge_weight": 2,
                "city": "SUNDERNAGAR",
                "cod": 1,
                "cod_charges": 0,
                "cod_multiplier": 0,
                "cost": "",
                "courier_company_id": 19,
                "courier_name": "Ecom Express Surface 2kg",
                "courier_type": "0",
                "coverage_charges": 0,
                "cutoff_time": "11:00",
                "delivery_boy_contact": "Not Available",
                "delivery_performance": 5,
                "description": "",
                "edd": "",
                "entry_tax": 0,
                "estimated_delivery_days": "6",
                "etd": "Jul 03, 2024",
                "etd_hours": 144,
                "freight_charge": 150.12,
                "id": 456774028,
                "is_custom_rate": 0,
                "is_hyperlocal": false,
                "is_international": 0,
                "is_rto_address_available": true,
                "is_surface": true,
                "local_region": 0,
                "metro": 0,
                "min_weight": 2,
                "mode": 0,
                "new_edd": 0,
                "odablock": false,
                "other_charges": 0,
                "others": "{\"allow_postcode_auto_sync\":1,\"cancel_real_time\":true,\"wec\":1}",
                "pickup_availability": "0",
                "pickup_performance": 4.5,
                "pickup_priority": "",
                "pickup_supress_hours": 0,
                "pod_available": "Instant",
                "postcode": "175019",
                "qc_courier": 0,
                "rank": "",
                "rate": 150.12,
                "rating": 4.62,
                "realtime_tracking": "Real Time",
                "region": 1,
                "rto_charges": 142.6,
                "rto_performance": 5,
                "seconds_left_for_pickup": 0,
                "secure_shipment_disabled": false,
                "ship_type": 1,
                "state": "HIMACHAL PRADESH",
                "suppress_date": "",
                "suppress_text": "",
                "suppression_dates": {
                    "action_on": "2024-06-15 23:27:37",
                    "delay_remark": "Festival",
                    "delivery_delay_by": 2065762,
                    "delivery_delay_days": "1",
                    "delivery_delay_from": "2024-06-17",
                    "delivery_delay_to": "2024-06-17",
                    "pickup_delay_by": 2065762,
                    "pickup_delay_days": "1",
                    "pickup_delay_from": "2024-06-17",
                    "pickup_delay_to": "2024-06-17"
                },
                "surface_max_weight": "30.00",
                "tracking_performance": 4,
                "volumetric_max_weight": null,
                "weight_cases": 4.6,
                "zone": "z_e"
            }
        ],
        "blocked_courier_companies": [
            {
                "block_reason": "Operational Issues",
                "courier_company_id": 33,
                "courier_name": "Xpressbees Air",
                "postcode": "175019"
            }
        ],
        "child_courier_id": null,
        "is_recommendation_enabled": 1,
        "recommendation_advance_rule": 0,
        "recommended_by": {
            "id": 6,
            "title": "Recommendation By Shiprocket"
        },
        "recommended_courier_company_id": 43,
        "shiprocket_recommended_courier_id": 43
    },
    "dg_courier": 0,
    "eligible_for_insurance": false,
    "insurace_opted_at_order_creation": false,
    "is_allow_templatized_pricing": true,
    "is_latlong": 0,
    "is_old_zone_opted": false,
    "is_zone_from_mongo": true,
    "label_generate_type": 2,
    "on_new_zone": 2,
    "seller_address": [],
    "status": 200,
    "user_insurance_manadatory": false
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

```json
{
    "message": "Required field missing",
    "errors": {
        "weight": [
            "The weight field is required when order id is not present."
        ],
        "cod": [
            "The cod field is required when order id is not present."
        ],
        "order_id": [
            "The order id field is required when pickup postcode / delivery postcode / cod / weight is not present."
        ]
    },
    "status_code": 422
}
```

#### Invalid Data — HTTP 200 OK

```json
{
    "status": 404,
    "message": "Order does not exist!"
}
```

### Request for Shipment Pickup

`POST https://apiv2.shiprocket.in/v1/external/courier/generate/pickup`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to create a pickup request for your order shipment. The API returns the pickup status along with the estimated pickup time.<br>You will have to call the 'Generate Manifest' API after the successful response of this API.

**Note:**

- The AWB must be already generated for the shipment id to generate the pickup request.
- Only one shipment_id can be passed at a time.
- In case the pickup_date falls on a holiday or Sunday, the shipment will be booked for the next available date.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `shipment_id` | YES | integer | The shipment id of the shipment which is requested for pickup. | [16090109] |
| `status` | NO | *string* | Use this field to retry if the pickup request fails. Value: retry | retry |
| `pickup_date` | NO | *Array of Dates* | Use this field to schedule a pickup for a future date. The date format would be YYYY-MM-DD. You can reschedule the date till that date has not passed | ["2022-06-04"] |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	"shipment_id": [16090109]
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
	"shipment_id": [16091084]
	
}
```

```json
{
    "pickup_status": 1,
    "response": {
        "pickup_scheduled_date": "2021-12-10 12:39:54",
        "pickup_token_number": "Reference No: 194_BIGFOOT 1966840_11122021",
        "status": 3,
        "others": "{\"tier_id\":5,\"etd_zone\":\"z_e\",\"etd_hours\":\"{\\\"assign_to_pick\\\":6.9000000000000004,\\\"pick_to_ship\\\":22.600000000000001,\\\"ship_to_deliver\\\":151.40000000000001,\\\"etd_zone\\\":\\\"z_e\\\",\\\"pick_to_ship_table\\\":\\\"dev_etd_pickup_to_ship\\\",\\\"ship_to_deliver_table\\\":\\\"dev_etd_ship_to_deliver\\\"}\",\"actual_etd\":\"2021-12-18 00:36:03\",\"routing_code\":\"S2\\/S-69\\/1B\\/016\",\"addition_in_etd\":[\"deduction_of_6_and_half_hours\"],\"shipment_metadata\":{\"type\":\"ship\",\"device\":\"WebKit\",\"platform\":\"desktop\",\"client_ip\":\"94.237.77.195\",\"created_at\":\"2021-12-10 12:36:03\",\"request_type\":\"web\"},\"templatized_pricing\":0,\"selected_courier_type\":\"Best in price\",\"recommended_courier_data\":{\"etd\":\"Dec 19, 2021\",\"price\":153,\"rating\":3.6,\"courier_id\":54},\"recommendation_advance_rule\":null,\"dynamic_weight\":\"1.00\"}",
        "pickup_generated_date": {
            "date": "2021-12-10 12:39:54.034695",
            "timezone_type": 3,
            "timezone": "Asia/Kolkata"
        },
        "data": "Pickup is confirmed by Xpressbees 1kg For AWB :- 143254213727423"
    }
}
```

#### Missing Fields — HTTP 400 Bad Request

Example request body:

```json
{
	"shipment_id": [""]
	
}
```

```json
{
    "message": "Invalid shipment id",
    "status_code": 400
}
```

#### Invalid Data — HTTP 400 Bad Request

Example request body:

```json
{
	"shipment_id": [121121212]
	
}
```

```json
{
    "message": "Invalid shipment id",
    "status_code": 400
}
```

### Upload Blocked Pincodes

`POST https://serviceability.shiprocket.in/v1/external/blocked-pincodes/upload`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value `Bearer {{token}}`.

**Request headers**

| Header | Value |
|---|---|
| `Authorization` | `Bearer {{token}}` |
| `Content-Type` | `application/json` |

**Description**

Adds pincodes to the seller's delivery‑block list, or removes them.

The `block` action runs **per‑pincode classification** — valid entries are saved, invalid or duplicate entries are recorded in **Activity Logs** as a downloadable CSV instead of failing the whole request. The `unblock` action is **all‑or‑nothing**: the first invalid pincode aborts the call.

#### Authentication

Pass the bearer token via header or query parameter:

| Method | Where | Example |
|---|---|---|
| Header (standard) | `Authorization` | `Authorization: {{vault:bearer-token}}` |
| Query parameter | `?token=` |  |

#### Request Body Parameters

| Field | Type | Required | Description |
|---|---|---|---|
| `postcode` | object | Yes | Container for the `delivery_blocked` array. |
| `postcode.delivery_blocked` | string[] | Yes | Pincodes to add to / remove from the delivery block list. |
| `action` | string | Yes | `"block"` or `"unblock"`. Trimmed and lowercased server‑side. |

#### Notes

- **Idempotency** — `block` is naturally idempotent. `unblock` is similarly idempotent for already‑removed pincodes.
- **Request size cap** — 1 MiB.
- **Unblock does not write to Activity Logs** — only `block` produces a per‑pincode result CSV; `unblock` surfaces the first malformed entry inline in the response.
- **Timestamps** — stored in IST (`Asia/Kolkata`).

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{"postcode": {"delivery_blocked": ["110001", "560034"]}, "action": "block"}
```

**Example responses:** none published in the source collection for this request.

### Get Blocked Pincodes

`GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value `Bearer {{token}}`.

**Request headers**

| Header | Value |
|---|---|
| `Authorization` | `Bearer {{token}}` |

**Description**

Reads the seller's delivery‑block list. **Multi‑mode** — the mode is selected by query parameters:

| Query parameter | Mode |
|---|---|
| `?is_download=1` | CSV download (highest priority) |
| `?search=` (non‑empty after trim) | Prefix search |
| (default — neither of the above) | Paginated list |

> Only `delivery_blocked` data is exposed.

#### Authentication

Pass the bearer token via header or query parameter:

| Method | Where | Example |
|---|---|---|
| Header (standard) | `Authorization` | `Authorization: Bearer` |
| Query parameter | `?token=` | `?token=` |

#### Query Parameters

| Param | Type | Required | Default | Description |
|---|---|---|---|---|
| `is_download` | int | No | — | Set to `1` for a presigned S3 link to a CSV of all pincodes. |
| `search` | string | No | — | Prefix‑match filter. Frontend should enforce a 4‑char minimum; server does not validate. |
| `per_page` | int | No | `15` | Page size (paginated mode only). Range `1–15` (silently capped at 15). |
| `current_page` | int | No | `1` | Page number (paginated mode only). Must be `>= 1`. |

#### Mode Precedence

1. `is_download=1` — wins over everything
2. `search` (non‑empty after trim) — wins over pagination
3. Default — paginated list

#### Notes

- **Search isn't validated server‑side** — pass any string; the server does prefix matching as‑is.
- **CSV download URL is presigned for 24 hours** — generate a fresh link by calling the endpoint again after expiry.
- `per_page` is silently capped at 15; values above 15 are treated as 15.

**Example responses**

#### Paginated List — Default Page — HTTP 200 OK

```json
{"data":{"delivery_blocked":["110001","110002","110003","110004","110005","110006","110007","110008","110009","110010","110011","110012","110013","110014","110015"],"total":42,"per_page":15,"current_page":1,"last_page":3}}
```

#### Paginated List — Custom Page & Size — HTTP 200 OK

Example request: `GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get?per_page=5&current_page=2`

```json
{"data":{"delivery_blocked":["110006","110007","110008","110009","110010"],"total":42,"per_page":5,"current_page":2,"last_page":9}}
```

#### Paginated List — Page Beyond Data — HTTP 200 OK

Example request: `GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get?per_page=15&current_page=99`

```json
{"data":{"delivery_blocked":[],"total":42,"per_page":15,"current_page":99,"last_page":3}}
```

#### Paginated List — No Blocked Pincodes — HTTP 200 OK

```json
{"data":{"delivery_blocked":[],"total":0,"per_page":15,"current_page":1,"last_page":1}}
```

#### Prefix Search — Matching Results — HTTP 200 OK

Example request: `GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get?search=1100`

```json
{"data":{"delivery_blocked":["110001","110002","110003","110004","110005"]}}
```

#### Prefix Search — No Matches — HTTP 200 OK

Example request: `GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get?search=9999`

```json
{"data":{"delivery_blocked":[]}}
```

#### CSV Download — With Data — HTTP 200 OK

Example request: `GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get?is_download=1`

```json
{"data":{"url":"https://s3.amazonaws.com/shiprocket-blocked-pincodes/seller_123_blocked.csv?X-Amz-Expires=86400&X-Amz-Signature=abc123"}}
```

#### CSV Download — No Data — HTTP 200 OK

Example request: `GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get?is_download=1`

```json
{"data":{"url":null}}
```

#### Error — Invalid Pagination Params — HTTP 400 Bad Request

Example request: `GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get?per_page=0&current_page=0`

```json
{"message":"The per page must be at least 1. The current page must be at least 1.","errors":{"per_page":["The per page must be at least 1."],"current_page":["The current page must be at least 1."]}}
```

#### Error — Unauthorized — HTTP 401 Unauthorized

```json
{"message":"Unauthenticated."}
```

#### Error — Internal Server Error — HTTP 500 Internal Server Error

```json
{"message":"Internal server error."}
```

#### Error — CSV Generation Failed — HTTP 500 Internal Server Error

Example request: `GET https://serviceability.shiprocket.in/v1/external/block-pincodes/get?is_download=1`

```json
{"message":"Failed to generate CSV. Please try again later."}
```

## Shipments

These APIs can be used to get shipment details. You can either get all shipment details at once or the details of a particular shipment.

Below are some associated status codes related to the shipments:

##### Status Codes:

| **STATUS CODE** | **DESCRIPTION** |
|---|---|
| 1 | AWB Assigned |
| 2 | Label Generated |
| 3 | Pickup Scheduled/Generated |
| 4 | Pickup Queued |
| 5 | Manifest Generated |
| 6 | Shipped |
| 7 | Delivered |
| 8 | Cancelled |
| 9 | RTO Initiated |
| 10 | RTO Delivered |
| 11 | Pending |
| 12 | Lost |
| 13 | Pickup Error |
| 14 | RTO Acknowledged |
| 15 | Pickup Rescheduled |
| 16 | Cancellation Requested |
| 17 | Out For Delivery |
| 18 | In Transit |
| 19 | Out For Pickup |
| 20 | Pickup Exception |
| 21 | Undelivered |
| 22 | Delayed |
| 23 | Partial_Delivered |
| 24 | Destroyed |
| 25 | Damaged |
| 26 | Fulfilled |
| 38 | Reached at Destination |
| 39 | Misrouted |
| 40 | RTO NDR |
| 41 | RTO OFD |
| 42 | Picked Up |
| 43 | Self Fulfilled |
| 44 | DISPOSED_OFF |
| 45 | CANCELLED_BEFORE_DISPATCHED |
| 46 | RTO_IN_TRANSIT |
| 47 | QC Failed |
| 48 | Reached Warehouse |
| 49 | Custom Cleared |
| 50 | In Flight |
| 51 | Handover to Courier |
| 52 | Shipment Booked |
| 54 | In Transit Overseas |
| 55 | Connection Aligned |
| 56 | Reached Overseas Warehouse |
| 57 | Custom Cleared Overseas |
| 59 | Box Packing |

### Get All Shipment Details

`GET https://apiv2.shiprocket.in/v1/external/shipments`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to get the shipment details of all the shipments in your Shiprocket account.

The data is displayed in the default format if no filter or sort condition is passed. You can use the filter and sort conditions to streamline your data.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| ` sort` | NO | *string* | The sort format: ASC or DESC | ASC |
| `sort_by` | NO | *string* | The field to sort by. | id |
| `filter` | NO | *string* | The filter value. | 141121 |
| `filter_by` | NO | *string* | The filed to filter by. | id |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": [
        {
            "number": "",
            "code": "",
            "id": 3322791,
            "order_id": 3324748,
            "products": [
                {
                    "name": "gsdgw",
                    "sku": "123",
                    "quantity": 2
                }
            ],
            "awb": "109123535421",
            "status": "CANCELED",
            "created_at": "28th Aug 2018 07:11 PM",
            "channel_id": 76893,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "payment_method": "cod"
        },
        {
            "number": "",
            "code": "",
            "id": 3322802,
            "order_id": 3324759,
            "products": [
                {
                    "name": "Black tshirt XL",
                    "sku": "BlackTshirt",
                    "quantity": 2
                }
            ],
            "awb": "109123535594",
            "status": "CANCELED",
            "created_at": "28th Aug 2018 07:14 PM",
            "channel_id": 76893,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "payment_method": "cod"
        }
        
    ],
    "meta": {
        "pagination": {
            "total": 19631,
            "count": 15,
            "per_page": 2,
            "current_page": 1,
            "total_pages": 1309,
            "links": {
                "next": "https://apiv2.shiprocket.in/v1/external/shipments?page=2"
            }
        }
    }
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/shipment`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Get Details of Specific Shipment

`GET https://apiv2.shiprocket.in/v1/external/shipments`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Get the details of a specific Shipment by passing the value of shipment_id in the endpoint URL. No other body parameters are required.

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/shipments/16016920](https://apiv2.shiprocket.in/v1/external/shipments/16016920) |

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/shipments/16016920`

```json
{
    "data": {
        "id": 16016920,
        "order_id": 16167171,
        "channel_id": 76893,
        "company_id": 67216,
        "invoice_no": null,
        "invoice_date": null,
        "courier": null,
        "sr_courier_id": null,
        "awb": null,
        "awb_assign_date": null,
        "pickup_generated_date": null,
        "pickup_token_number": null,
        "method": "Standard",
        "weight": "0.000",
        "dimensions": "0.00x0.00x0.00",
        "quantity": 1,
        "cost": "0.00",
        "tax": "0.00",
        "cod_charges": "0.00",
        "total": "9000.00",
        "shipping_address": {
            "city": "New Delhi",
            "state": "DELHI",
            "address": "House 221B, Leaf Village",
            "country": "India",
            "pincode": "110002",
            "address_2": "Near Hokage House",
            "company_name": null
        },
        "customer_details": null,
        "status": 8,
        "shipped_date": null,
        "delivered_date": null,
        "returned_date": null,
        "label_url": null,
        "manifest_url": null,
        "created_at": {
            "date": "2019-07-31 12:37:41.000000",
            "timezone_type": 3,
            "timezone": "Asia/Kolkata"
        },
        "updated_at": {
            "date": "2019-07-31 15:57:11.000000",
            "timezone_type": 3,
            "timezone": "Asia/Kolkata"
        }
    }
}
```

#### Invalid or Wrong Data — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/shipments/100100`

```json
{
	
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/shipments/{shipment_id}`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Cancel a Shipment

`POST https://apiv2.shiprocket.in/v1/external/orders/cancel/shipment/awbs`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to cancel a created shipment before the "Out for Pickup" state. Multiple AWBs can be passed together as an array to cancel them simultaneously.

**Limit: 2000 AWBs**

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `awbs` | YES | *string* | The AWB/List of AWBs that need to be canceled. | 19041211125783 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
"awbs": ["19041211125783"]
}
```

**Example responses**

#### Successful Call — HTTP 204 No Content

```json
{
"message": "Bulk Shipment cancellation is in progress. Please wait for some time."
}
```

## Labels | Manifests | Invoice

These APIs are designed for the generation of manifests, invoices and labels for your shipment. You can also get manifest details and print a manifest using the respective API's.

### Generate Manifest

`POST https://apiv2.shiprocket.in/v1/external/manifests/generate`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Using this API, you can generate the manifest for your order. This API generates the manifest and displays the download URL of the same.

**Note:**

- Multiple ids can also be passed as an array for bulk generation of manifests.
- AWB must be assigned and pickup requested on the shipment id to generate manifest.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `shipment_id` | YES | *integer* | The shipment id of the order. Multiple ids can be passed as an array, separated by commas. | [16090109] |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	"shipment_id": []
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
	"shipment_id": [16090109]
}
```

```json
{
    "status": 1,
    "manifest_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/25149/manifest/MANIFEST-3051.pdf"
}
```

#### Bad Request — HTTP 400 Bad Request

```json
{
    "message": "Manifest already generated for some shipment_ids, Please try again after removin them",
    "status_code": 400,
    "already_manifested_shipment_ids": [
        16090109
    ]
}
```

#### Missing Field or Invalid Data — HTTP 200 OK

Example request body:

```json
{
	"shipment_id": ""
}
```

```json
{
    "status": 1,
    "manifest_url": ""
}
```

### Print Manifest

`POST https://apiv2.shiprocket.in/v1/external/manifests/print`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to print the generated manifest of orders at an individual level.

**Note**

- Manifest needs to be generated first for this API to print it. Use the 'Generate Manifest' API to do the same.
- Multiple order ids can be passed together.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_ids` | YES | *integer* | The Shiprocket order id of whose manifest is to be generated. Multiple ids can be passed together as an array. | [16090109] |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	"order_ids": []
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
	"order_ids": [16240904]
}
```

```json
{
    "manifest_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/manifest/c_25149/print_manifests/115261_fedex-surface_D52FE1564654197.pdf"
}
```

#### Missing or Invalid Data — HTTP 200 OK

Example request body:

```json
{
	"order_ids": [16161616]
}
```

```json
{
    "manifest_url": ""
}
```

#### Wrong Format — HTTP 500 Internal Server Error

Example request body:

```json
{
	"order_ids": ""
}
```

```json
{
    "message": "Invalid argument supplied for foreach()",
    "status_code": 500
}
```

### Generate Label

`POST https://apiv2.shiprocket.in/v1/external/courier/generate/label`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Generate the label of order by passing the shipment id in the form of an array. This API displays the URL of the generated label.

**Note:**

- The AWB must be assigned to the shipment to generate the label.
- 'shipment_id' must be passed as an array.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `shipment_id` | YES | *integer* | The shipment id of the order whose label is to be generated. Multiple ids can be passed together as an array. | [16104408,16104409] |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	"shipment_id": []
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
	"shipment_id": ["16104408"]
}
```

```json
{
    "label_created": 1,
    "label_url": "https://kr-shipmultichannel.s3.ap-southeast-1.amazonaws.com/25149/labels/shipping-label-16104408-788830567028.pdf",
    "response": "Label has been created and uploaded successfully!",
    "not_created": []
}
```

#### Wrong Format — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

Example request body:

```json
{
	"shipment_id": "16161616"
}
```

```json
{
    "message": "Oops! Invalid Data.",
    "errors": {
        "shipment_id": [
            "The shipment id must be an array."
        ]
    },
    "status_code": 422
}
```

#### Missing or Invalid Data — HTTP 200 OK

Example request body:

```json
{
	"shipment_id": [""]
}
```

```json
{
    "label_created": 0,
    "not_created": [],
    "response": "No valid shipment ids to print label"
}
```

### Generate Invoice

`POST https://apiv2.shiprocket.in/v1/external/orders/print/invoice`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to generate the invoice for your order by passing the respective Shiprocket order ids.

The generated invoice URL is displayed as a response. Multiple ids can be passed together as an array.

**Note**

- Order ids must be passed as an array.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `ids` | YES | *integer* | The Shiprocket order id of the orders whose invoices are to be created. Multiple ids can be passed together as an array. | [16255275,16255276] |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	"ids": []

}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
	"ids": ["16255275"]

}
```

```json
{
    "is_invoice_created": true,
    "invoice_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/25149/invoices/KD101019281564656872.pdf",
    "not_created": []
}
```

#### Missing Fields or Invalid Data — HTTP 400 Bad Request

Example request body:

```json
{
    "ids": [""]
}
```

```json
{
    "is_invoice_created": false,
    "message": "Invoice could not be created",
    "not_created": []
}
```

#### Wrong Format — HTTP 500 Internal Server Error

Example request body:

```json
{
    "ids": "16255275"
}
```

```json
{
    "message": "Invalid argument supplied for foreach()",
    "status_code": 500
}
```

### Generate Label + Invoice (Combined)

`POST https://apiv2.shiprocket.in/v1/external/courier/generate/label-invoice`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Generates a single combined PDF containing the shipping label followed by the invoice for each shipment. Output preserves the input order of `shipment_ids`.

**Note:**

- `shipment_ids` is mandatory for external requests.
- Not supported for reseller-enabled accounts.
- `completed: true` in the response means the request ran synchronously (always true for external callers) — it does not guarantee every shipment succeeded. Check `success_count` / `error_count`; on partial/full failure `error_file_url` points to the error details.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `shipment_ids` | YES | *integer* | Shipment IDs to generate the combined label+invoice for. Multiple ids can be passed as an array. Max 200 per request. | [123456789, 123456790] |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{"shipment_ids": []}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
	"shipment_ids": [123456789, 123456790]
}
```

```json
{
  "completed": true,
  "file_url": "https://shiprocket.s3.amazonaws.com/combined-label-invoice.pdf",
  "error_file_url": "",
  "success_count": 2,
  "error_count": 0
}
```

#### Missing or Invalid Data — HTTP 400 Bad Request

Example request body:

```json
{
	"shipment_ids": []
}
```

```text
{\n    \"message\": \"shipment_ids is required for external requests\"\n}
```

#### Max Limit Exceeded — HTTP 400 Bad Request

Example request body:

```json
{
	"shipment_ids": [100000001, 100000002]
}
```

```text
{\n    \"message\": \"Max 200 shipment_ids allowed per external request\"\n}
```

#### Reseller Not Supported — HTTP 400 Bad Request

Example request body:

```json
{
	"shipment_ids": [123456789]
}
```

```text
{\n    \"message\": \"Combined label+invoice not supported for reseller setup\"\n}
```

## NDR

These APIs can be used to get details about your shipments that are in NDR, as well as manage them, which includes fetching NDR shipments from your account (all or specific) and taking actions such as Reattempt and RTO based on the reason.

### Get All NDR Shipments

`GET https://apiv2.shiprocket.in/v1/external/ndr/all`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value ``. Note: the Postman request itself is flagged "No Auth"; the collection-wide guideline still states Bearer authorization.

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |
| `Authorization` | `` |

**Description**

This API call will display a list of all the shipments that are in NDR in your Shiprocket account.

You can also sort and filter the data according to your needs by passing the optional parameters. Not passing anything will display the data in the default format.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `page` | NO | *integer* | The page number to display. | 5 |
| `per_page` | NO | *integer* | The number of entries per page. | 5 |
| `to` | NO | *string* | The end date. | 2021-08-02 |
| `from` | NO | *string* | The start date. | 2021-08-02 |
| `search` | NO | *string* | Search for AWB. | 224477 |

**Example responses**

#### Successful Call — HTTP 200 OK

```text
{
        "data": [
                {
                        "id": 94711332,
                        "shipment_id": 94323606,
                        "customer_name": "John Doe",
                        "customer_email": "john.doe@gmail.com",
                        "customer_phone": "9999999998",
                        "customer_address": "#400, Ground floor, Valley",
                        "customer_address_2": "View estate, Near Indira Canteen",
                        "customer_city": "Gurgaon",
                        "customer_state": "Haryana",
                        "customer_pincode": "122003",
                        "payment_status": "2",
                        "status": "UNDELIVERED",
                        "status_code": 36,
                        "payment_method": "prepaid",
                        "created_at": "12 Mar 2021, 04:16 PM",
                        "reason": "Customer Asked For Future Delivery",
                        "attempts": 1,
                        "ndr_raised_at": "2021-03-17 22:51:17",
                        "courier": "FedEx",
                        "awb_code": "784698160933",
                        "escalation_status": "N/A",
                        "product_name": "Cricket kit",
                        "product_price": "1484.01",
                        "shipment_channel_id": 152865,
                        "history": [
                                {
                                        "id": 50347646,
                                        "ndr_id": 14933497,
                                        "ndr_reason": "Customer Asked For Future Delivery",
                                        "action_by": 3,
                                        "ndr_attempt": 1,
                                        "medium": null,
                                        "ndr_push_status": 0,
                                        "comment": "",
                                        "call_center_call_recording": "",
                                        "call_center_recording_date": "",
                                        "proof_recording": null,
                                        "proof_image": null,
                                        "sms_response": "No Response",
                                        "ndr_raised_at": "2021-03-17 22:51:17"
                                }
                        ],
                        "delivered_date": ""
                },
                {
                        "id": 94279207,
                        "shipment_id": 93891718,
                        "customer_name": " Jane Doe",
                        "customer_email": "jane.doe@gmail.com",
                        "customer_phone": "9999999998",
                        "customer_address": "416 udyog vihar 3",
                        "customer_address_2": "near UC",
                        "customer_city": "Gurgaon",
                        "customer_state": "Haryana",
                        "customer_pincode": "122003",
                        "payment_status": "2",
                        "status": "UNDELIVERED",
                        "status_code": 36,
                        "payment_method": "prepaid",
                        "created_at": "10 Mar 2021, 03:14 PM",
                        "reason": "Customer Asked For Future Delivery",
                        "attempts": 2,
                        "ndr_raised_at": "2021-03-17 22:51:13",
                        "courier": "FedEx",
                        "awb_code": "784613220808",
                        "escalation_status": "N/A",
                        "product_name": "Maa ka Aachaar",
                        "product_price": "1019.00",
                        "shipment_channel_id": 152865,
                        "history": [
                                {
                                        "id": 50227533,
                                        "ndr_id": 14908011,
                                        "ndr_reason": "Customer Asked For Future Delivery",
                                        "action_by": 3,
                                        "ndr_attempt": 1,
                                        "medium": null,
                                        "ndr_push_status": 0,
                                        "comment": "",
                                        "call_center_call_recording": "",
                                        "call_center_recording_date": "",
                                        "proof_recording": null,
                                        "proof_image": null,
                                        "sms_response": "No Response",
                                        "ndr_raised_at": "2021-03-16 21:54:32"
                                },
                                {
                                        "id": 50244772,
                                        "ndr_id": 14908011,
                                        "ndr_reason": "Customer Asked For Future Delivery",
                                        "action_by": 2,
                                        "ndr_attempt": 1,
                                        "medium": null,
                                        "ndr_push_status": 2,
                                        "comment": null,
                                        "call_center_call_recording": "",
                                        "call_center_recording_date": "",
                                        "proof_recording": null,
                                        "proof_image": null,
                                        "sms_response": "No Response",
                                        "ndr_raised_at": "2021-03-16 21:54:32"
                                },
                                {
                                        "id": 50347643,
                                        "ndr_id": 14908011,
                                        "ndr_reason": "Customer Asked For Future Delivery",
                                        "action_by": 3,
                                        "ndr_attempt": 2,
                                        "medium": null,
                                        "ndr_push_status": 0,
                                        "comment": "",
                                        "call_center_call_recording": "",
                                        "call_center_recording_date": "",
                                        "proof_recording": null,
                                        "proof_image": null,
                                        "sms_response": "No Response",
                                        "ndr_raised_at": "2021-03-17 22:51:13"
                                }
                        ],
                        "delivered_date": ""
                }
                
        ],
        "meta": {
                "pagination": {
                        "total": 3,
                        "count": 3,
                        "per_page": 15,
                        "current_page": 1,
                        "total_pages": 1,
                        "links": {
                                "next": "https://apiv2.shiprocket.in/v1/external/ndr/all?page=2"
                        }
                }
        }
}
```

### Get Specific NDR Shipment Details

`GET https://apiv2.shiprocket.in/v1/external/ndr/{AWB}`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value `Bearer {token}`.

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |
| `Authorization` | `Bearer {token}` |

**Path placeholders:** `AWB` (see the parameter table in the description for meaning where the source provides one).

**Description**

Get the shipment details of a particular order through this API by passing the AWB number in the endpoint URL itself. You can get the details like AWB, NDR Attempt, NDR Reason, Customer Details, Product Details, Courier.

Type in your AWB code in place of {AWB}. No other body parameters are required.

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/ndr/94711332](https://apiv2.shiprocket.in/v1/external/ndr/94711332) |

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/ndr/94711332`

```json
{
    "data": [
        {
            "id": 94711332,
            "shipment_id": 94323606,
            "customer_name": "John Doe",
            "customer_email": "john@gmail.com",
            "customer_phone": "999999998",
            "customer_address": "#123, Ground floor, Valley apartment",
            "customer_address_2": "near view estate",
            "customer_city": "Gurgaon",
            "customer_state": "Haryana",
            "customer_pincode": "122003",
            "payment_status": "2",
            "status": "UNDELIVERED",
            "status_code": 36,
            "payment_method": "prepaid",
            "created_at": "12 Mar 2021, 04:16 PM",
            "reason": "Customer Asked For Future Delivery",
            "attempts": 1,
            "ndr_raised_at": "2021-03-17 22:51:17",
            "courier": "FedEx",
            "awb_code": "8373927474982",
            "escalation_status": "N/A",
            "product_name": "Gaming Chair",
            "product_price": "14484.01",
            "shipment_channel_id": 152865,
            "history": [
                {
                    "id": 50347646,
                    "ndr_id": 14933497,
                    "ndr_reason": "Customer Asked For Future Delivery",
                    "action_by": 3,
                    "ndr_attempt": 1,
                    "medium": null,
                    "ndr_push_status": 0,
                    "comment": "",
                    "call_center_call_recording": "",
                    "call_center_recording_date": "",
                    "proof_recording": null,
                    "proof_image": null,
                    "sms_response": "No Response",
                    "ndr_raised_at": "2021-03-17 22:51:17"
                }
            ],
            "delivered_date": ""
        }
    ],
    "meta": {
        "pagination": {
            "total": 1,
            "count": 1,
            "per_page": 15,
            "current_page": 1,
            "total_pages": 1,
            "links": {}
        }
    }
}
```

### Action NDR

`POST https://apiv2.shiprocket.in/v1/external/ndr/{awb}/action`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Path placeholders:** `awb` (see the parameter table in the description for meaning where the source provides one).

**Description**

This API will let you take actions like Reattempt and RTO on the shipments that are in NDR.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `action` | YES | *string* | The action needs to be specified | ‘fake-attempt’ or ‘re-attempt’ or 'return' |
| `comments` | YES | *string* | Any comment can be mentioned | The Byer does not want the product |
| `phone` | NO | *string* | The phone number will be updated at the time of re-attempt and fake-attempt | 9999988888 |
| `proof_audio` | CONDITIONAL YES | *string* | URL of the audio which will be updated at the time of the fake attempt | [https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/imports/ndr/1655100133_file_example_MP3_700KB.mp3](https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/imports/ndr/1655100133_file_example_MP3_700KB.mp3) |
| `proof_image` | CONDITIONAL YES | *string* | URL of the image which will be updated at the time of the fake attempt | [https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/imports/ndr/img_7687678678.jpg](https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/imports/ndr/img_7687678678.jpg) |
| `remarks` | CONDITIONAL YES | *string* | Remarks will be updated at the time of a fake attempt | Delivery Requested |
| `address1` | NO | *string* | address1 will be updated at the time of re-attempt and fake-attempt | U-56, sector-23, Noida, India |
| `address2` | NO | *string* | addres2 will be updated at the time of re-attempt and fake-attempt | U-56, sector-23, Noida, India |
| `deferred_date` | NO | *date(string)* | Deferred date will be updated as preferred_date at the time of fake-attempt and re-attempt | 2022-08-10 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
   "action": "",
   "comments": ""
}
```

**Example responses**

#### Successful Call — HTTP 202 Accepted

Example request: `POST https://apiv2.shiprocket.in/v1/external/ndr/8805225468/action`

Example request body:

```json
{
   "action": "return",
   "comments": "The Buyer does not want the product"
 
}
```

```json
{
    "status": "Data Updated Sucessfully"
}
```

## Tracking

Use these APIs to get the tracking details of your shipments through the AWB code or the Shipment ID.

#### Shipment Status Codes:

| **STATUS CODE** |  | **DESCRIPTION** |
|---|---|---|
| 6 |  | Shipped |
| 7 |  | Delivered |
| 8 |  | Canceled |
| 9 |  | RTO Initiated |
| 10 |  | RTO Delivered |
| 12 |  | Lost |
| 13 |  | Pickup Error |
| 14 |  | RTO Acknowledged |
| 15 |  | Pickup Rescheduled |
| 16 |  | Cancellation Requested |
| 17 |  | Out For Delivery |
| 18 |  | In Transit |
| 19 |  | Out For Pickup |
| 20 |  | Pickup Exception |
| 21 |  | Undelivered |
| 22 |  | Delayed |
| 23 |  | Partial_Delivered |
| 24 |  | DESTROYED |
| 25 |  | DAMAGED |
| 26 |  | FULFILLED |
| 27 |  | Pickup Booked |
| 38 |  | REACHED AT DESTINATION HUB |
| 39 |  | MISROUTED |
| 40 |  | RTO_NDR |
| 41 |  | RTO_OFD |
| 42 |  | PICKED UP |
| 43 |  | SELF FULFILLED |
| 44 |  | DISPOSED OFF |
| 45 |  | CANCELLED_BEFORE_DISPATCHED |
| 46 |  | RTO IN INTRANSIT |
| 47 |  | QC FAILED |
| 48 |  | Reached Warehouse |
| 49 |  | Custom Cleared |
| 50 |  | In Flight |
| 51 |  | Handover to Courier |
| 52 |  | Shipment Booked |
| 54 |  | In Transit Overseas |
| 55 |  | Connection Aligned |
| 56 |  | Reached Overseas Warehouse |
| 57 |  | Custom Cleared Overseas |
| 59 |  | Box Packing |
| 68 |  | PROCESSED AT WAREHOUSE |
| 60 |  | FC Allocated |
| 61 |  | Picklist Generated |
| 62 |  | Ready To Pack |
| 63 |  | Packed |
| 67 |  | FC MANIFEST GENERATED |
| 71 |  | HANDOVER EXCEPTION |
| 72 |  | PACKED EXCEPTION |
| 75 |  | RTO_LOCK |
| 76 |  | UNTRACEABLE |
| 77 |  | ISSUE_RELATED_TO_THE_RECIPIENT |
| 78 |  | REACHED_BACK_AT_SELLER_CITY |

### Get Tracking through AWB

`GET https://apiv2.shiprocket.in/v1/external/courier/track/awb/{awb_code}`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Path placeholders:** `awb_code` (see the parameter table in the description for meaning where the source provides one).

**Description**

Get the tracking details of your shipment by entering the AWB code of the same in the endpoint URL itself. No other body parameters are required to access this API.

The response is displayed in JSON format.

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/courier/track/awb/788830567028](https://apiv2.shiprocket.in/v1/external/courier/track/awb/788830567028) |

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/courier/track/awb/788830567028`

```json
{
    "tracking_data": {
        "track_status": 1,
        "shipment_status": 7,
        "shipment_track": [
            {
                "id": 236612717,
                "awb_code": "141123221084922",
                "courier_company_id": 51,
                "shipment_id": 236612717,
                "order_id": 237157589,
                "pickup_date": "2022-07-18 20:28:00",
                "delivered_date": "2022-07-19 11:37:00",
                "weight": "0.30",
                "packages": 1,
                "current_status": "Delivered",
                "delivered_to": "Chittoor",
                "destination": "Chittoor",
                "consignee_name": "",
                "origin": "Banglore",
                "courier_agent_details": null,
                "courier_name": "Xpressbees Surface",
                "edd": null,
                "pod": "Available",
                "pod_status": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/courier/51/pod/141123221084922.png"
            }
        ],
        "shipment_track_activities": [
            {
                "date": "2022-07-19 11:37:00",
                "status": "DLVD",
                "activity": "Delivered",
                "location": "MADANPALLI, Madanapalli, ANDHRA PRADESH",
                "sr-status": "7",
                "sr-status-label": "DELIVERED"
            },
            {
                "date": "2022-07-19 08:57:00",
                "status": "OFD",
                "activity": "Out for Delivery Out for delivery: 383439-Nandinayani Reddy Bhaskara Sitics Logistics  (356231) (383439)-PDS22200085719383439-FromMob , MobileNo:- 9963133564",
                "location": "MADANPALLI, Madanapalli, ANDHRA PRADESH",
                "sr-status": "17",
                "sr-status-label": "OUT FOR DELIVERY"
            },
            {
                "date": "2022-07-19 07:33:00",
                "status": "RAD",
                "activity": "Reached at Destination Shipment BagOut From Bag : nxbg03894488",
                "location": "MADANPALLI, Madanapalli, ANDHRA PRADESH",
                "sr-status": "38",
                "sr-status-label": "REACHED AT DESTINATION HUB"
            },
            {
                "date": "2022-07-18 21:02:00",
                "status": "IT",
                "activity": "InTransit Shipment added in Bag nxbg03894488",
                "location": "BLR/FC1, BANGALORE, KARNATAKA",
                "sr-status": "18",
                "sr-status-label": "IN TRANSIT"
            },
            {
                "date": "2022-07-18 20:28:00",
                "status": "PKD",
                "activity": "Picked Shipment InScan from Manifest",
                "location": "BLR/FC1, BANGALORE, KARNATAKA",
                "sr-status": "6",
                "sr-status-label": "SHIPPED"
            },
            {
                "date": "2022-07-18 13:50:00",
                "status": "PUD",
                "activity": "PickDone ",
                "location": "RTO/CHD, BANGALORE, KARNATAKA",
                "sr-status": "42",
                "sr-status-label": "PICKED UP"
            },
            {
                "date": "2022-07-18 10:04:00",
                "status": "OFP",
                "activity": "Out for Pickup ",
                "location": "RTO/CHD, BANGALORE, KARNATAKA",
                "sr-status": "19",
                "sr-status-label": "OUT FOR PICKUP"
            },
            {
                "date": "2022-07-18 09:51:00",
                "status": "DRC",
                "activity": "Pending Manifest Data Received",
                "location": "RTO/CHD, BANGALORE, KARNATAKA",
                "sr-status": "NA",
                "sr-status-label": "NA"
            }
        ],
        "track_url": "https://shiprocket.co//tracking/141123221084922",
        "etd": "2022-07-20 19:28:00",
        "qc_response": {
            "qc_image": "",
            "qc_failed_reason": ""
        }
    }
}
```

#### Missing Fields — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/courier/track/awb/`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

#### Invalid Data — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/courier/track/awb/111111111111`

```json
{
    "tracking_data": {
        "track_status": 0,
        "error": "Aahh! There is no activities found in our DB. Please have some patience it will be updated soon."
    }
}
```

### Get Tracking Data for Multiple AWBS

`POST https://apiv2.shiprocket.in/v1/external/courier/track/awbs`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Get the tracking details of multiple shipments by entering their AWB codes together as an array.

The response is displayed in JSON format.

**Notes:**

- Data must be passed as a string array.
- Maximum of 50 AWB codes is supported at a time.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `awbs` | YES | *string* | The AWB codes of the shipments. Must be passed as an array. | ["788830567028","788829354408"] |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "awbs": ["",""]
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
    "awbs": ["788830567028","788829354408"]
}
```

```json
{
    "788829354408": {
        "tracking_data": {
            "track_status": 1,
            "shipment_status": 1,
            "shipment_track": [
                {
                    "id": 8067757,
                    "awb_code": "788829354408",
                    "courier_company_id": 2,
                    "shipment_id": null,
                    "order_id": 16240551,
                    "pickup_date": null,
                    "delivered_date": null,
                    "weight": "2.5",
                    "packages": 1,
                    "current_status": "AWB Assigned",
                    "delivered_to": "New Delhi",
                    "destination": "New Delhi",
                    "consignee_name": "Naruto",
                    "origin": "Jammu",
                    "courier_agent_details": null
                }
            ],
            "shipment_track_activities": [
                {
                    "date": "2019-08-01 02:05:05",
                    "activity": "Shipment information sent to FedEx - OC",
                    "location": "NA"
                }
            ],
            "track_url": "https://app.shiprocket.in/tracking/awb/788829354408"
        }
    },
    "788830567028": {
        "tracking_data": {
            "track_status": 1,
            "shipment_status": 3,
            "shipment_track": [
                {
                    "id": 8087109,
                    "awb_code": "788830567028",
                    "courier_company_id": 2,
                    "shipment_id": null,
                    "order_id": 16255275,
                    "pickup_date": null,
                    "delivered_date": null,
                    "weight": "2.5",
                    "packages": 1,
                    "current_status": "Pickup Generated",
                    "delivered_to": "New Delhi",
                    "destination": "New Delhi",
                    "consignee_name": "Naruto",
                    "origin": "Jammu",
                    "courier_agent_details": null
                }
            ],
            "shipment_track_activities": [
                {
                    "date": "2019-08-01 05:20:55",
                    "activity": "Shipment information sent to FedEx - OC",
                    "location": "NA"
                }
            ],
            "track_url": "https://app.shiprocket.in/tracking/awb/788830567028"
        }
    }
}
```

#### Invalid or Missing Data — HTTP 200 OK

Example request body:

```json
{
    "awbs": ["134232541121","123456712345"]
}
```

```json
{
	
}
```

#### Wrong Format — HTTP 500 Internal Server Error

Example request body:

```text
{
    "awbs": "788830567028","788829354408"
}
```

```json
{
    "message": "Invalid argument supplied for foreach()",
    "status_code": 500
}
```

### Get Tracking through Shipment ID

`GET https://apiv2.shiprocket.in/v1/external/courier/track/shipment/{shipment_id}`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Path placeholders:** `shipment_id` (see the parameter table in the description for meaning where the source provides one).

**Description**

Get the tracking details of your shipment by entering the shipment_id of the same in the endpoint URL. No other body parameters are required to access this API.

The response is displayed in JSON format.

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/courier/track/shipment/16104408](https://apiv2.shiprocket.in/v1/external/courier/track/shipment/16104408) |

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/courier/track/shipment/16104408`

```json
{
    "tracking_data": {
        "track_status": 1,
        "shipment_status": 42,
        "shipment_track": [
            {
                "id": 185584215,
                "awb_code": "1091188857722",
                "courier_company_id": 10,
                "shipment_id": 168347943,
                "order_id": 168807908,
                "pickup_date": null,
                "delivered_date": null,
                "weight": "0.10",
                "packages": 1,
                "current_status": "PICKED UP",
                "delivered_to": "Mumbai",
                "destination": "Mumbai",
                "consignee_name": "Musarrat",
                "origin": "PALWAL",
                "courier_agent_details": null,
                "edd": "2021-12-27 23:23:18"
            }
        ],
        "shipment_track_activities": [
            {
                "date": "2021-12-23 14:23:18",
                "status": "X-PPOM",
                "activity": "In Transit - Shipment picked up",
                "location": "Palwal_NewColony_D (Haryana)",
                "sr-status": "42"
            },
            {
                "date": "2021-12-23 14:19:37",
                "status": "FMPUR-101",
                "activity": "Manifested - Pickup scheduled",
                "location": "Palwal_NewColony_D (Haryana)",
                "sr-status": "NA"
            },
            {
                "date": "2021-12-23 14:19:34",
                "status": "X-UCI",
                "activity": "Manifested - Consignment Manifested",
                "location": "Palwal_NewColony_D (Haryana)",
                "sr-status": "5"
            }
        ],
        "track_url": "https://shiprocket.co//tracking/1091188857722",
        "etd": "2021-12-28 10:19:35"
    }
}
```

#### Invalid Data — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/courier/track/shipment/00000000`

```json
{
    "tracking_data": {
        "track_status": 0,
        "error": "Aahh! There is no activities found in our DB. Please have some patience it will be updated soon."
    }
}
```

#### Missing Fields — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/courier/track/shipment/`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Get Tracking Data through Order ID

`GET https://apiv2.shiprocket.in/v1/external/courier/track?order_id=123&channel_id=12345`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value `Bearer {token}`.

**Request headers**

| Header | Value |
|---|---|
| `Authorization` | `Bearer {token}` |

**Query parameters in the request template**

| Name | Example value |
|---|---|
| `order_id` | `123` |
| `channel_id` | `12345` |

**Description**

Get the tracking details of your shipment by entering the Order ID of the same in the endpoint URL itself. If you have the same order ID in more than one channel, then use the param channel_id.

The response is displayed in JSON format.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *string* | The order ID/number of your store. | NO-123 |
| `channel_id` | NO | *integer* | Channel ID corresponding to the store | 12345 |

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/courier/track?order_id=123&channel_id=12345](https://apiv2.shiprocket.in/v1/external/courier/track?order_id=123&channel_id=12345) |

**Example responses**

#### Get Tracking Data through Order iD — HTTP 200 OK

```json
[
    {
        "tracking_data": {
            "track_status": 1,
            "shipment_status": 42,
            "shipment_track": [
                {
                    "id": 185584215,
                    "awb_code": "1091188857722",
                    "courier_company_id": 10,
                    "shipment_id": 168347943,
                    "order_id": 168807908,
                    "pickup_date": null,
                    "delivered_date": null,
                    "weight": "0.10",
                    "packages": 1,
                    "current_status": "PICKED UP",
                    "delivered_to": "Mumbai",
                    "destination": "Mumbai",
                    "consignee_name": "Musarrat",
                    "origin": "PALWAL",
                    "courier_agent_details": null,
                    "edd": "2021-12-27 23:23:18"
                }
            ],
            "shipment_track_activities": [
                {
                    "date": "2021-12-23 14:23:18",
                    "status": "X-PPOM",
                    "activity": "In Transit - Shipment picked up",
                    "location": "Palwal_NewColony_D (Haryana)",
                    "sr-status": "42"
                },
                {
                    "date": "2021-12-23 14:19:37",
                    "status": "FMPUR-101",
                    "activity": "Manifested - Pickup scheduled",
                    "location": "Palwal_NewColony_D (Haryana)",
                    "sr-status": "NA"
                },
                {
                    "date": "2021-12-23 14:19:34",
                    "status": "X-UCI",
                    "activity": "Manifested - Consignment Manifested",
                    "location": "Palwal_NewColony_D (Haryana)",
                    "sr-status": "5"
                }
            ],
            "track_url": "https://shiprocket.co//tracking/1091188857722",
            "etd": "2021-12-28 10:19:35"
        }
    }
]
```

## Pickup Addresses

Use these APIs to get a list of all available pickup locations in your account or add a new pickup location.

### Get All Pickup Locations

`GET https://apiv2.shiprocket.in/v1/external/settings/company/pickup`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Description**

Get a list of all pickup locations that have been added to your Shiprocket account through this API.

No parameters are required to use this API.

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": {
        "shipping_address": [
            {
                "id": 4984500,
                "pickup_location": "Casa Moderna",
                "address_type": null,
                "address": "1900 GF, Sector 45",
                "address_2": "near Park 1",
                "updated_address": false,
                "old_address": "",
                "old_address2": "",
                "tag": "",
                "tag_value": "",
                "instruction": "",
                "city": "Gurgaon",
                "state": "Haryana",
                "country": "India",
                "pin_code": "122003",
                "email": "abcxyz@gmail.com",
                "is_first_mile_pickup": 0,
                "phone": "9667667496",
                "name": "CASA MODERNA",
                "company_id": 1138022,
                "gstin": null,
                "vendor_name": null,
                "status": 2,
                "phone_verified": 1,
                "lat": "28.44786453607",
                "long": "77.074250297017",
                "open_time": null,
                "close_time": null,
                "warehouse_code": null,
                "alternate_phone": "",
                "rto_address_id": 4984500,
                "lat_long_status": 1,
                "new": 1,
                "associated_rto_address": null,
                "is_primary_location": 1
            }
        ],
        "allow_more": "true",
        "is_blackbox_seller": false,
        "company_name": "CASA MODERNA",
        "recent_addresses": []
    }
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/settings/company/picku`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Add a New Pickup Location

`POST https://apiv2.shiprocket.in/v1/external/settings/company/addpickup`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API can be used to add a new pickup location to your account.<br>Pass the minimum required parameters to add the location.<br>Further details to the address can be added if required.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `pickup_location` | YES | *string* | The nickname of the new pickup location. Max 36 characters. | Home |
| `name` | YES | *string* | The shipper's name. | Deadpool |
| `email` | YES | *string* | The shipper's email address. | [deadpool@chimichanga.com](mailto:deadpool@chimichanga.com) |
| `phone` | YES | *integer* | Shipper's phone number. | 9777777779 |
| `address` | YES | *string* | Shipper's primary address. Max 80 characters. | Mutant Facility, Sector 3 |
| `address_2` | NO | *string* | Additional address details. | House number 34 |
| `city` | YES | *string* | Pickup location city name. | Pune |
| `state` | YES | *string* | Pickup location state name. | Maharashtra |
| `country` | YES | *string* | Pickup location country. | India |
| `pin_code` | YES | *integer* | Pickup location pincode. | 110022 |
| `lat` | NO | *float* | Pickup location Latitude. | 22.4064 |
| `long` | NO | *float* | Pickup location Longitude. | 69.0747 |
| `address_type` | NO | *string* | To be given if address of different vendor is to be provided with pickup address | vendor |
| `vendor_name` | NO | *string* | Name of vendor if address_type is vendor | John |
| `gstin` | NO | *string* | gstin of vendor | 09XXXCH7409R1XXX |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	"pickup_location": "",
	"name": "",
	"email": "",
	"phone": "",
	"address": "",
	"address_2": "",
	"city": "",
	"state":"",
	"country": "",
	"pin_code": ""
	
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
	"pickup_location": "Home",
	"name": "Deadpool",
	"email": "deadpool@chimichanga.com",
	"phone": "9777777779",
	"address": "Mutant Facility, Sector 3 ",
	"address_2": "",
	"city": "Pune",
	"state":"Maharshtra",
	"country": "India",
	"pin_code": "110022"
	
}
```

```json
{
    "success": true,
    "address": {
        "company_id": 25149,
        "pickup_code": "TESTADI",
        "address": "Mutant Facility, Sector 3",
        "address_2": "",
        "address_type": null,
        "city": "South West Delhi",
        "state": "Maharshtra",
        "country": "India",
        "gstin": null,
        "pin_code": "110022",
        "phone": "9777777779",
        "email": "deadpool@chimichanga.com",
        "name": "Deadpool",
        "alternate_phone": null,
        "lat": null,
        "long": null,
        "status": 1,
        "phone_verified": 0,
        "rto_address_id": 1468067,
        "extra_info": "{\"source\":3}",
        "updated_at": "2021-10-12 11:51:48",
        "created_at": "2021-10-12 11:51:48",
        "id": 1856901
    },
    "pickup_id": 1856901,
    "company_name": "ShiprocketTest",
    "full_name": "API"
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

Example request body:

```json
{
	"pickup_location": "Home",
	"name": "Deadpool",
	"email": "deadpool@chimichanga.com",
	"phone": "",
	"address": "Mutant Facility, Sector 3 ",
	"address_2": "",
	"city": "Pune",
	"state":"Maharshtra",
	"country": "India",
	"pin_code": "110022"
	
}
```

```json
{
    "message": "Not a valid mobile number",
    "errors": {
        "phone": [
            "The phone field is required."
        ]
    },
    "status_code": 422
}
```
