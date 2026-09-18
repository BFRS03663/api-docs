# Shiprocket API — Hyperlocal & International

> Generated from the published Postman collection behind [https://apidocs.shiprocket.in/](https://apidocs.shiprocket.in/) (collection "Shiprocket API", published id `SzYW1zB2`, version tag `latest`). Only content present in that collection is reproduced here; where the source omits something, the omission is stated.

## Hyperlocal

_No folder-level description in the source collection._

## Hyperlocal / Orders

These APIs can be used to get details about your created orders, as well as manage them, which include fetching orders from your channel, syncing their status, and exporting them to a CSV file.

#### Filters:

The 'filter' values are mentioned below. Use these to filter your data:

| **filter_by** | **filter** | **description** |
|---|---|---|
| payment_method |  |  |
| -> | cod | cash on delivery orders |
| -> | Prepaid | for prepaid orders |
| channel_order_id |  |  |
| -> | 123 | Your order id |
| status |  |  |
| -> | 1 | New |
| -> | 2 | Invoiced |
| -> | 3 | Ready To Ship |
| -> | 4 | Pickup Scheduled |
| -> | 5 | Canceled |
| -> | 6 | Shipped |
| -> | 7 | Delivered |
| -> | 8 | ePayment Failed |
| -> | 9 | Returned |
| -> | 10 | Unmapped |
| -> | 11 | Unfulfillable |
| -> | 12 | Pickup Queue |
| -> | 13 | Pickup Rescheduled |
| -> | 14 | Pickup Error// Created when there is an error on the pickup schedule |
| -> | 15 | RTO Initiated |
| -> | 16 | RTO Delivered |
| -> | 17 | RTO Acknowledged |
| -> | 18 | Cancellation Requested |
| -> | 19 | Out for Delivery |
| -> | 20 | In Transit |
| -> | 21 | Return Pending |
| -> | 22 | Return Initiated |
| -> | 23 | Return Pickup Queued |
| -> | 24 | Return Pickup Error |
| -> | 25 | Return In Transit |
| -> | 26 | Return Delivered |
| -> | 27 | Return Cancelled |
| -> | 28 | Return Pickup Generated |
| -> | 29 | Return Cancellation Requested |
| -> | 30 | Return Pickup Cancelled |
| -> | 31 | Return Pickup Rescheduled |
| -> | 32 | Return Picked Up |
| -> | 33 | Lost |
| -> | 34 | Out For Pickup |
| -> | 35 | Pickup Exception |
| -> | 36 | Undelivered |
| -> | 37 | Delivery Delayed |
| -> | 38 | Partial Delivered |
| -> | 39 | Destroyed |
| -> | 40 | Damaged |
| -> | 41 | Fulfilled |
| -> | 42 | Archived |
| -> | 43 | Reached Destination Hub |
| -> | 44 | Misrouted |
| -> | 45 | RTO_OFD |
| -> | 46 | RTO_NDR |
| -> | 47 | Return Out For Pickup |
| -> | 48 | Return Out For Delivery |
| -> | 49 | Return Pickup Exception |
| -> | 50 | Return Undelivered |
| -> | 51 | Picked Up |
| -> | 52 | Self Fulfilled |
| -> | 53 | Disposed Off |
| -> | 54 | Canceled before Dispatched |
| -> | 55 | RTO In-Transit |
| -> | 57 | QC Failed |
| -> | 58 | Reached Warehouse |
| -> | 59 | Custom Cleared |
| -> | 60 | In Flight |
| -> | 61 | Handover to Courier |
| -> | 62 | Booked |
| -> | 64 | In Transit Overseas |
| -> | 65 | Connection Aligned |
| -> | 66 | Reached Overseas Warehouse |
| -> | 67 | Custom Cleared Overseas |
| -> | 68 | RETURN ACKNOWLEGED |
| -> | 69 | Box Packing |
| -> | 70 | Pickup Booked |
| -> | 71 | DARKSTORE SCHEDULED |
| -> | 72 | Allocation in Progress |
| -> | 81 | PROCESSED AT WAREHOUSE |
| -> | 73 | FC Allocated |
| -> | 74 | Picklist Generated |
| -> | 75 | Ready to Pack |
| -> | 76 | Packed |
| -> | 80 | FC MANIFEST GENERATED |
| -> | 82 | PACKED EXCEPTION |
| -> | 83 | HANDOVER EXCEPTION |
| -> | 87 | RTO_LOCK |
| -> | 88 | UNTRACEABLE |
| -> | 89 | ISSUE_RELATED_TO_THE_RECIPIENT |
| -> | 90 | REACHED_BACK_AT_SELLER_CITY |

### Create Custom Order

`POST https://apiv2.shiprocket.in/v1/external/orders/create/adhoc`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to create a quick custom order. Quick orders are the ones where we do not store the product details in the master catalogue.

You have to pass all the required params at the minimum to create a quick custom order. You can add additional parameters as per your preference.

**Note:**

- In case the 'shipping_is_billing' field is false, further shipping detail fields are required.

If no channel id is passed, the order will be assigned to the default custom channel. If the channel id is not known, use the 'Get All Channels' API to get the list of all integrated channels in your Shiprocket account.

- order_id field cannot be equal to an already existing id. Doing so does not change or affect the existing order.
- New orders cannot be created with order id's same as that of cancelled orders. If error 422 shows up despite filling in the correct details, consider changing the order_id.
- Be sure to input the correct calculated sub_total amount. The total is not calculated automatically through the API.
- The 'order_id' returned in the response is the Shiprocket order_id. Please save this order ID as we will use this in future API calls.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *string* | The order id you want to specify to the order. Max char: 50. (Avoid passing character values as this contradicts some other API calls). | 224477 or 224-477 |
| `order_date` | YES | *string* | The date of order creation in yyyy-mm-dd format. Time is additional. | 2019-07-24 11:11 |
| `pickup_location` | YES | *string* | The name of the pickup location added in your Shiprocket account. This cannot be a new location. | Jammu |
| `channel_id` | NO | *integer* | Mention this in case you need to assign the order to a particular channel. Deafult is 'Custom'. | 27022 |
| `comment` | NO | *string* | Option to add 'From' field to the shipment. To do this, enter the name in the following format: 'Reseller: [name]'. | Reseller: Divine |
| `reseller_name` | NO | *string* | The 'from' name if you want to print. Use 'Reseller: [name]' | Reseller: Divine |
| `company_name` | NO | *string* | Name of the company. | Amazon |
| `billing_customer_name` | YES | *string* | First name of the billed customer. | John |
| `billing_last_name` | NO | *string* | Last name of the billed customer. | Doe |
| `billing_address` | YES | *string* | address details of the billed customer. | Civil line, House 20 |
| `billing_address_2` | NO | *string* | Further address details of the billed customer. | Near Hokage House |
| `billing_city` | YES | *string* | Billing address city. Max char: 30. | New Delhi |
| `billing_pincode` | YES | *integer* | Pincode of the billing address. | 110002 |
| `billing_state` | YES | *string* | Billing address state. | Delhi |
| `billing_country` | YES | *string* | Billing address country. | India |
| `billing_email` | YES | *string* | Email address of the billed customer. | [John@doe.com](https://mailto:John@doe.com) |
| `billing_phone` | YES | *integer* | The phone number of the billing customer. | 9856321472 |
| `billing_alternate_phone` | NO | *integer* | Alternate phone number of the billing customer. | 8604690454 |
| `shipping_is_billing` | YES | *boolean* | Whether the shipping address is the same as billing address. 1 or 'true' for yes and 0 or 'false' for no. | true |
| `shipping_customer_name` | CONDITIONAL YES | *string* | Name of the customer the order is shipped to. Required in case billing is not same as shipping. | Jane |
| `shipping_last_name` | NO | *string* | Last name of the shipping customer. | Doe |
| `shipping_address` | CONDITIONAL YES | *string* | Address of the Shipping customer. Required in case billing is not same as shipping. | Lane number 69 |
| `shipping_address_2` | NO | *string* | Further address details of shipping customer. | Andheri |
| `billing_isd_code` | NO | *string* | ISD code of the billing address. | +91 |
| `shipping_city` | CONDITIONAL YES | *string* | Shipping address city. | Mumbai |
| `shipping_pincode` | CONDITIONAL YES | *integer* | Shipping address pincode. | 200912 |
| `shipping_country` | CONDITIONAL YES | *string* | Shipping address country. | India |
| `shipping_state` | CONDITIONAL YES | *string* | Shipping address state. | Maharashtra |
| `shipping_email` | NO | *string* | Email of the shipping customer. | [Jane@doe.com](https://mailto:Jane@doe.com) |
| `shipping_phone` | CONDITIONAL YES | *integer* | Phone no. of the shipping customer. |  |
| `longitude` | YES | *float* | Delivery Longitude. Mandatory in case of Hyper local shipments | 77.06745147705078 |
| `latitude` | YES | *float* | Delivery Latitude. Mandatory in case of Hyper local shipments | 28.50724220275879 |
| `order_items` | YES | / | List of items and their relevant fields in the form of Array. | / |
| `name` | YES | *string* | Name of the product. | Jeans |
| `sku` | YES | *string* | The sku id of the product. | cbs123 |
| `units` | YES | *integer* | No of units that are to be shipped. | 10 |
| `selling_price` | YES | *integer* | The selling price per unit in Rupee. Inclusive of GST. | 900 |
| `discount` | NO | *integer* | The discount amount in Rupee. Inclusive of tax. | 10 |
| `tax` | NO | *integer* | The tax percentage on the item. | 5 |
| `hsn` | NO | *integer* | Harmonised System Nomenclature code. Used to determine the category of taxation the goods fall under. | 44122 |
| `payment_method` | YES | *string* | The method of payment. Can be either COD (Cash on delivery) Or Prepaid. | COD |
| `shipping_charges` | NO | *integer* | Shipping charges if any in Rupee. | 5 |
| `giftwrap_charges` | NO | *integer* | Giftwrap charges if any in Rupee. | 5 |
| `transaction_charges` | NO | *integer* | Transaction charges if any in Rupee. | 5 |
| `total_discount` | NO | *integer* | The total discount amount in Rupee. | 15 |
| `sub_total` | YES | *integer* | Calculated sub total amount in Rupee after deductions. | 9010 |
| `length` | YES | *float* | The length of the item in cms. Must be more than 0.5. | 10 |
| `breadth` | YES | *float* | The breadth of the item in cms. Must be more than 0.5. | 10 |
| `height` | YES | *float* | The height of the item in cms. Must be more than 0.5. | 10 |
| `weight` | YES | *float* | The weight of the item in kgs. Must be more than 0. | 2.5 |
| `ewaybill_no` | NO | *string* | Details relating to the shipment of goods. . | K92373490 |
| `customer_gstin` | NO | *string* | Goods and Services Tax Identification Number. | 29ABCDE1234F2Z5 |
| `invoice_number` | NO | *string* |  |  |
| `order_type` | NO | *string* | Key to differentiate between Essentials or Non Essentials Shipments. Order type can only be ESSENTIALS or NON ESSENTIALS. Please note it is case sensitive and blank values are allowed. | ESSENTIALS |
| `checkout_shipping_method` | NO | *string* | Only for SRF users. | a. SR_RUSH: SDD, NDD b. SR_STANDARD: Surface Delivery c. SR_EXPRESS: Air Delivery d. SR_QUICK: 3 hrs delivery |
| `what3words_address` | NO | *string* | What3words is a proprietary geocode system designed to identify any location on the surface of Earth with a resolution of about 3 meters. The system encodes geographic coordinates into three permanently fixed dictionary words. | toddler.geologist.animated |
| `is_insurance_opt` | NO | *boolean* | To secure shipments above the order value of Rs 2500 | true |
| `is_document` | NO | *integer* | To create a document order | 1 or 0 |
| `shipping_method` | YES | *string* | Shipping method use HL in case of Hyper local shipments | HL |

**Request body template** (as published; empty strings are placeholders to fill in)

```text
{
    "order_id": "",
    "order_date": "",
    "pickup_location": "",
    "channel_id": "",
    "comment": "",
    "reseller_name": "",
    "company_name": "",
    "billing_customer_name": "",
    "billing_last_name": "",
    "billing_address": "",
    "billing_address_2": "",
    "billing_isd_code": "",
    "billing_city": "",
    "billing_pincode": "",
    "billing_state": "",
    "billing_country": "",
    "billing_email": "",
    "billing_phone": "",
    "billing_alternate_phone":"",
    "shipping_is_billing": "",
    "shipping_customer_name": "",
    "shipping_last_name": "",
    "shipping_address": "",
    "shipping_address_2": "",
    "shipping_city": "",
    "shipping_pincode": "",
    "shipping_country": "",
    "shipping_state": "",
    "shipping_email": "",
    "shipping_phone": "",
    "order_items": [
        {
            "name": "",
            "sku": "",
            "units": "",
            "selling_price": "",
            "discount": "",
            "tax": "",
            "hsn": ""
        }
    ],
    "payment_method": "",
    "shipping_charges": "",
    "giftwrap_charges": "",
    "transaction_charges": "",
    "total_discount": "",
    "sub_total": "",
    "length": "",
    "breadth": "",
    "height": "",
    "weight": "",
    "ewaybill_no": "",
    "customer_gstin": "",
    "invoice_number":"",
    "order_type":"",
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
  "order_id": "224-447",
  "order_date": "2019-07-24 11:11",
  "pickup_location": "Jammu",
  "channel_id": "",
  "comment": "Reseller: M/s Goku",
  "billing_customer_name": "Naruto",
  "billing_last_name": "Uzumaki",
  "billing_address": "House 221B, Leaf Village",
  "billing_address_2": "Near Hokage House",
  "billing_city": "New Delhi",
  "billing_pincode": "110002",
  "billing_state": "Delhi",
  "billing_country": "India",
  "billing_email": "naruto@uzumaki.com",
  "billing_phone": "9876543210",
  "shipping_is_billing": true,
  "shipping_customer_name": "",
  "shipping_last_name": "",
  "shipping_address": "",
  "shipping_address_2": "",
  "shipping_city": "",
  "shipping_pincode": "",
  "shipping_country": "",
  "shipping_state": "",
  "shipping_email": "",
  "shipping_phone": "",
  "order_items": [
    {
      "name": "Kunai",
      "sku": "chakra123",
      "units": 10,
      "selling_price": "900",
      "discount": "",
      "tax": "",
      "hsn": 441122
    }
  ],
  "payment_method": "Prepaid",
  "shipping_charges": 0,
  "giftwrap_charges": 0,
  "transaction_charges": 0,
  "total_discount": 0,
  "sub_total": 9000,
  "length": 10,
  "breadth": 15,
  "height": 20,
  "weight": 2.5
}
```

```json
{
    "order_id": 16161616,
    "shipment_id": 15151515,
    "status": "NEW",
    "status_code": 1,
    "onboarding_completed_now": 0,
    "awb_code": null,
    "courier_company_id": null,
    "courier_name": null
}
```

#### Invalid Data — HTTP 400 Bad Request

Example request body:

```json
{
  "order_id": "224-477",
  "order_date": "2019-07-24 11:11",
  "pickup_location": "Jammu",
  "channel_id": "12345",
  "comment": "Reseller: M/s Goku",
  "billing_customer_name": "Naruto",
  "billing_last_name": "Uzumaki",
  "billing_address": "House 221B, Leaf Village",
  "billing_address_2": "Near Hokage House",
  "billing_city": "New Delhi",
  "billing_pincode": "110002",
  "billing_state": "Delhi",
  "billing_country": "India",
  "billing_email": "naruto@uzumaki.com",
  "billing_phone": "9876543210",
  "shipping_is_billing": true,
  "shipping_customer_name": "",
  "shipping_last_name": "",
  "shipping_address": "",
  "shipping_address_2": "",
  "shipping_city": "",
  "shipping_pincode": "",
  "shipping_country": "",
  "shipping_state": "",
  "shipping_email": "",
  "shipping_phone": "",
  "order_items": [
    {
      "name": "Kunai",
      "sku": "chakra123",
      "units": 10,
      "selling_price": "900",
      "discount": "",
      "tax": "",
      "hsn": 441122
    }
  ],
  "payment_method": "Prepaid",
  "shipping_charges": 0,
  "giftwrap_charges": 0,
  "transaction_charges": 0,
  "total_discount": 0,
  "sub_total": 9000,
  "length": 10,
  "breadth": 15,
  "height": 20,
  "weight": 2.5
}
```

```json
{
    "message": "Given channel id does not exist",
    "status_code": 400
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

Example request body:

```json
{
  "order_id": "",
  "order_date": "2019-07-24 11:11",
  "pickup_location": "Jammu",
  "channel_id": "",
  "comment": "Reseller: M/s Goku",
  "billing_customer_name": "Naruto",
  "billing_last_name": "Uzumaki",
  "billing_address": "House 221B, Leaf Village",
  "billing_address_2": "Near Hokage House",
  "billing_city": "New Delhi",
  "billing_pincode": "110002",
  "billing_state": "Delhi",
  "billing_country": "India",
  "billing_email": "naruto@uzumaki.com",
  "billing_phone": "9876543210",
  "shipping_is_billing": true,
  "shipping_customer_name": "",
  "shipping_last_name": "",
  "shipping_address": "",
  "shipping_address_2": "",
  "shipping_city": "",
  "shipping_pincode": "",
  "shipping_country": "",
  "shipping_state": "",
  "shipping_email": "",
  "shipping_phone": "",
  "order_items": [
    {
      "name": "Kunai",
      "sku": "chakra123",
      "units": 10,
      "selling_price": "900",
      "discount": "",
      "tax": "",
      "hsn": 441122
    }
  ],
  "payment_method": "Prepaid",
  "shipping_charges": 0,
  "giftwrap_charges": 0,
  "transaction_charges": 0,
  "total_discount": 0,
  "sub_total": 9000,
  "length": 10,
  "breadth": 15,
  "height": 20,
  "weight": 2.5
}
```

```json
{
    "message": "Oops! Invalid Data.",
    "errors": {
        "order_id": [
            "The order id field is required."
        ]
    },
    "status_code": 422
}
```

#### Create Custom Order — (no HTTP status recorded in source)

_(empty response body in source)_

### Get all Orders

`GET https://apiv2.shiprocket.in/v1/external/orders`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value ``. Note: the Postman request itself is flagged "No Auth"; the collection-wide guideline still states Bearer authorization.

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |
| `Authorization` | `` |

**Description**

This API call will display a list of all created and available orders in your Shiprocket account. The product and shipment details are displayed as sub-arrays within each order detail.

You can also sort and filter the data according to your needs by passing the optional parameters. Not passing anything will display the data in the default format.

You can also fetch the data based on the order update date. PFB the validations:

```
1. if updated_to => passed and updated_from not passed =>, error will come => Please send update_from date along with update_to date
2. if updated_from => passed and updated_to => paased => and from-to > 30 then error => Difference between updated_from and update_to date should not be greater than 30 days.
3. if updated_From < 30 then error => Updated_from date should not be less than 30 days from current date
```

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `page` | NO | *integer* | The page number to display. | 5 |
| `per_page` | NO | *integer* | The number of entries per page. | 5 |
| `sort` | NO | *string* | Sort conditions: ASC or DESC | ASC |
| `sort_by` | NO | *string* | The Field to sort by: id or status | id |
| `to` | NO | *string* | The end date. | 2018-07-24 |
| `from` | NO | *string* | The start date. | 2019-07-24 |
| `filter_by` | NO | *string* | Field to filter by. | status, payment_method, delivery_country, channel_order_id |
| `filter` | NO | *string* | Value of the field |  |
| `search` | NO | *string* | Search for AWB or by Channel order_id (order id specified by you). | 224477 |
| `pickup_location` | NO | *string* | Search Orders on the basis of pickup location. | xyz |
| `channel_id` | NO | *integer* | Channel ID Returned in Get Integrated Channels API | 123 |
| `fbs` | NO | *integer* | Use this filter if you want to view and filter the SRF orders | 0 or 1 |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": [
        {
            "id": 16178831,
            "channel_id": 76893,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_order_id": "224-4779888",
            "customer_name": "Majin Bu",
            "customer_email": "naruto@uzumaki.com",
            "customer_phone": "9988998899",
            "pickup_location": "hell",
            "payment_status": "",
            "total": "9000.00",
            "tax": "0.00",
            "sla": "2 days",
            "shipping_method": "SR",
            "expedited": 0,
            "status": "CANCELED",
            "status_code": 5,
            "payment_method": "prepaid",
            "is_international": 0,
            "purpose_of_shipment": 0,
            "channel_created_at": "24 Jul 2019, 11:11 AM",
            "created_at": "31 Jul 2019, 03:03 PM",
            "products": [
                {
                    "id": 18769728,
                    "channel_order_product_id": "18769728",
                    "name": "Kunai",
                    "channel_sku": "chakra123",
                    "quantity": 10,
                    "product_id": 17484610,
                    "available": 50,
                    "status": "CANCELED",
                    "hsn": "441122"
                }
            ],
            "shipments": [
                {
                    "id": 16028538,
                    "isd_code": "",
                    "courier": "",
                    "weight": 0,
                    "dimensions": "0.00x0.00x0.00",
                    "pickup_scheduled_date": null,
                    "pickup_token_number": null,
                    "awb": "",
                    "return_awb": "",
                    "volumetric_weight": 0,
                    "pod": null,
                    "etd": "NA",
                    "rto_delivered_date": "0000-00-00 00:00:00",
                    "delivered_date": null,
                    "etd_escalation_btn": false
                }
            ],
            "activities": [
                "ORDER_CREATED",
                "ADDRESS_MODIFIED",
                "ORDER_CANCELLED"
            ],
            "allow_return": 0,
            "is_incomplete": 0,
            "errors": [],
            "show_escalation_btn": 0,
            "escalation_status": "",
            "escalation_history": []
        }
    ],
    "meta": {
        "pagination": {
            "total": 1,
            "count": 1,
            "per_page": 1,
            "current_page": 1,
            "total_pages": 1,
            "links": {}
        }
    }
}
```

#### Successful Call With No Parameters — HTTP 200 OK

```json
{
    "data": [
        {
            "id": 16178831,
            "channel_id": 76893,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_order_id": "224-4779888",
            "customer_name": "Majin Bu",
            "customer_email": "naruto@uzumaki.com",
            "customer_phone": "9988998899",
            "pickup_location": "hell",
            "payment_status": "",
            "total": "9000.00",
            "tax": "0.00",
            "sla": "2 days",
            "shipping_method": "SR",
            "expedited": 0,
            "status": "CANCELED",
            "status_code": 5,
            "payment_method": "prepaid",
            "is_international": 0,
            "purpose_of_shipment": 0,
            "channel_created_at": "24 Jul 2019, 11:11 AM",
            "created_at": "31 Jul 2019, 03:03 PM",
            "products": [
                {
                    "id": 18769728,
                    "channel_order_product_id": "18769728",
                    "name": "Kunai",
                    "channel_sku": "chakra123",
                    "quantity": 10,
                    "product_id": 17484610,
                    "available": 50,
                    "status": "CANCELED",
                    "hsn": "441122"
                }
            ],
            "shipments": [
                {
                    "id": 16028538,
                    "isd_code": "",
                    "courier": "",
                    "weight": 0,
                    "dimensions": "0.00x0.00x0.00",
                    "pickup_scheduled_date": null,
                    "pickup_token_number": null,
                    "awb": "",
                    "return_awb": "",
                    "volumetric_weight": 0,
                    "pod": null,
                    "etd": "NA",
                    "rto_delivered_date": "0000-00-00 00:00:00",
                    "delivered_date": null,
                    "etd_escalation_btn": false
                }
            ],
            "activities": [
                "ORDER_CREATED",
                "ADDRESS_MODIFIED",
                "ORDER_CANCELLED"
            ],
            "allow_return": 0,
            "is_incomplete": 0,
            "errors": [],
            "show_escalation_btn": 0,
            "escalation_status": "",
            "escalation_history": []
        },
        {
            "id": 16177223,
            "channel_id": 76893,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_order_id": "224-47798",
            "customer_name": "Naruto Uzumaki",
            "customer_email": "naruto@uzumaki.com",
            "customer_phone": "9876543210",
            "pickup_location": "okooi",
            "payment_status": "",
            "total": "9000.00",
            "tax": "0.00",
            "sla": "2 days",
            "shipping_method": "SR",
            "expedited": 0,
            "status": "NEW",
            "status_code": 1,
            "payment_method": "prepaid",
            "is_international": 0,
            "purpose_of_shipment": 0,
            "channel_created_at": "24 Jul 2019, 11:11 AM",
            "created_at": "31 Jul 2019, 02:39 PM",
            "products": [
                {
                    "id": 18766013,
                    "channel_order_product_id": "18766013",
                    "name": "Kunai",
                    "channel_sku": "chakra123",
                    "quantity": 10,
                    "product_id": 17484610,
                    "available": 50,
                    "status": "UNDEFINED",
                    "hsn": "441122"
                }
            ],
            "shipments": [
                {
                    "id": 16026933,
                    "isd_code": "+91",
                    "courier": "",
                    "weight": 0,
                    "dimensions": "0.00x0.00x0.00",
                    "pickup_scheduled_date": null,
                    "pickup_token_number": null,
                    "awb": "",
                    "return_awb": "",
                    "volumetric_weight": 0,
                    "pod": null,
                    "etd": "NA",
                    "rto_delivered_date": "0000-00-00 00:00:00",
                    "delivered_date": null,
                    "etd_escalation_btn": false
                }
            ],
            "activities": [
                "ORDER_CREATED"
            ],
            "allow_return": 0,
            "is_incomplete": 0,
            "errors": [],
            "show_escalation_btn": 0,
            "escalation_status": "",
            "escalation_history": []
        }
    ],
    "meta": {
        "pagination": {
            "total": 19631,
            "count": 15,
            "per_page": 15,
            "current_page": 1,
            "total_pages": 1309,
            "links": {
                "next": "https://apiv2.shiprocket.in/v1/external/orders?page=2"
            }
        }
    }
}
```

### Get Specific Order Details

`GET https://apiv2.shiprocket.in/v1/external/orders/show`

This entry in the source collection is byte-for-byte identical to **Orders / Get Specific Order Details**. See [Get Specific Order Details](orders.md#get-specific-order-details) for the full details (description, parameters, examples).

### Export your orders

`POST https://apiv2.shiprocket.in/v1/external/orders/export`

This entry in the source collection is byte-for-byte identical to **Orders / Export your orders**. See [Export your orders](orders.md#export-your-orders) for the full details (description, parameters, examples).

## Hyperlocal / Couriers

Use these APIs to assign AWB to your order, check for courier serviceability, and request for the pickup of your order.

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

```text
{"success" => true, "message" => "We are processing your request"}
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
| `is_new_hyperlocal` | YES | *boolean* | 1 in case of hyper-local shipments | 1 |
| `lat_from` | YES | *float* | Pickup Latitude. Mandatory in case of Hyper local shipments | 28.509223937988 |
| `long_from` | YES | *float* | Pickup Longitude. Mandatory in case of Hyper local shipments | 77.067848205566 |
| `lat_to` | YES | *float* | Delivery Latitude. Mandatory in case of Hyper local shipments | 28.50724220275879 |
| `long_to` | YES | *float* | Delivery Longitude. Mandatory in case of Hyper local shipments | 77.06745147705078 |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{

   "status": true,

   "data": [

       {

           "courier_name": "Shiprocket Quick",

           "rates": "345.3"

       }

   ]

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

## Hyperlocal / Tracking

Use these APIs to get the tracking details of your shipments through the AWB code or the Shipment ID.

#### Shipment Status Codes:

| **STATUS CODE** | **DESCRIPTION** |
|---|---|
| 6 | Shipped |
| 7 | Delivered |
| 8 | Canceled |
| 9 | RTO Initiated |
| 10 | RTO Delivered |
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
| 24 | DESTROYED |
| 25 | DAMAGED |
| 26 | FULFILLED |
| 27 | Pickup Booked |
| 38 | REACHED AT DESTINATION HUB |
| 39 | MISROUTED |
| 40 | RTO_NDR |
| 41 | RTO_OFD |
| 42 | PICKED UP |
| 43 | SELF FULFILLED |
| 44 | DISPOSED OFF |
| 45 | CANCELLED_BEFORE_DISPATCHED |
| 46 | RTO IN INTRANSIT |
| 47 | QC FAILED |
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
| 68 | PROCESSED AT WAREHOUSE |
| 60 | FC Allocated |
| 61 | Picklist Generated |
| 62 | Ready To Pack |
| 63 | Packed |
| 67 | FC MANIFEST GENERATED |
| 71 | HANDOVER EXCEPTION |
| 72 | PACKED EXCEPTION |
| 75 | RTO_LOCK |
| 76 | UNTRACEABLE |
| 77 | ISSUE_RELATED_TO_THE_RECIPIENT |
| 78 | REACHED_BACK_AT_SELLER_CITY |
| 79 | RIDER ASSIGNED |
| 80 | RIDER UNASSIGNED |
| 81 | RIDER ASSIGNED |
| 82 | RIDER REACHED AT DROP |
| 83 | SEARCHING_FOR_RIDER |

### Get Tracking through AWB

`GET https://apiv2.shiprocket.in/v1/external/courier/track/awb/{awb_code}`

This entry in the source collection is byte-for-byte identical to **Tracking / Get Tracking through AWB**. See [Get Tracking through AWB](shipping.md#get-tracking-through-awb) for the full details (description, parameters, examples).

### Get Tracking Data for Multiple AWBS

`POST https://apiv2.shiprocket.in/v1/external/courier/track/awbs`

This entry in the source collection is byte-for-byte identical to **Tracking / Get Tracking Data for Multiple AWBS**. See [Get Tracking Data for Multiple AWBS](shipping.md#get-tracking-data-for-multiple-awbs) for the full details (description, parameters, examples).

### Get Tracking through Shipment ID

`GET https://apiv2.shiprocket.in/v1/external/courier/track/shipment/{shipment_id}`

This entry in the source collection is byte-for-byte identical to **Tracking / Get Tracking through Shipment ID**. See [Get Tracking through Shipment ID](shipping.md#get-tracking-through-shipment-id) for the full details (description, parameters, examples).

### Get Tracking Data through Order ID

`GET https://apiv2.shiprocket.in/v1/external/courier/track?order_id=123&channel_id=12345`

This entry in the source collection is byte-for-byte identical to **Tracking / Get Tracking Data through Order ID**. See [Get Tracking Data through Order ID](shipping.md#get-tracking-data-through-order-id) for the full details (description, parameters, examples).

## Hyperlocal / Pickup Addresses

Use these APIs to get a list of all available pickup locations in your account or add a new pickup location.

### Get All Pickup Locations

`GET https://apiv2.shiprocket.in/v1/external/settings/company/pickup`

This entry in the source collection is byte-for-byte identical to **Pickup Addresses / Get All Pickup Locations**. See [Get All Pickup Locations](shipping.md#get-all-pickup-locations) for the full details (description, parameters, examples).

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
| `email` | YES | *string* | The shipper's email address. | [deadpool@chimichanga.com](https://mailto:deadpool@chimichanga.com) |
| `phone` | YES | *integer* | Shipper's phone number. | 9777777779 |
| `address` | YES | *string* | Shipper's primary address. Max 80 characters. | Mutant Facility, Sector 3 |
| `address_2` | NO | *string* | Additional address details. | House number 34 |
| `city` | YES | *string* | Pickup location city name. | Pune |
| `state` | YES | *string* | Pickup location state name. | Maharashtra |
| `country` | YES | *string* | Pickup location country. | India |
| `pin_code` | YES | *integer* | Pickup location pincode. | 110022 |
| `lat` | YES | *float* | Pickup location Latitude. | 22.4064 |
| `long` | YES | *float* | Pickup location Longitude. | 69.0747 |
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

## International

These endpoints are used for operations related to international orders and shipments.

## International / Tracking

Use these APIs to get the tracking details of your shipments through the AWB code or the Shipment ID.

### Get Tracking through AWB

`GET https://apiv2.shiprocket.in/v1/external/courier/track/awb/{awb_code}`

This entry in the source collection is byte-for-byte identical to **Tracking / Get Tracking through AWB**. See [Get Tracking through AWB](shipping.md#get-tracking-through-awb) for the full details (description, parameters, examples).

### Get Tracking through Shipment ID

`GET https://apiv2.shiprocket.in/v1/external/courier/track/shipment/{shipment_id}`

This entry in the source collection is byte-for-byte identical to **Tracking / Get Tracking through Shipment ID**. See [Get Tracking through Shipment ID](shipping.md#get-tracking-through-shipment-id) for the full details (description, parameters, examples).

### Get Tracking Data through Order iD

`GET https://apiv2.shiprocket.in/v1/external/courier/track?order_id=123&channel_id=12345`

This entry in the source collection is byte-for-byte identical to **Tracking / Get Tracking Data through Order ID**. See [Get Tracking Data through Order ID](shipping.md#get-tracking-data-through-order-id) for the full details (description, parameters, examples).

### Tracking

`GET https://apiv2.shiprocket.in/v1/external/international/orders/track`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Description**

Get the tracking details of your shipments. No other body parameters are required to access this API.

**Example responses**

#### Tracking Successful — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/external/international/orders/track`

```json
{
    "data": [
        {
            "id": 153210141,
            "channel_id": 2252386,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_order_id": "7022318826",
            "customer_name": "Shubham OworldInt ",
            "customer_email": "shubham.tyagi+1@shiprocket.com",
            "customer_phone": "8923309680",
            "customer_address": "223, alaKam, Aramex",
            "customer_address_2": "",
            "customer_city": "UPTON",
            "customer_state": "New York",
            "customer_pincode": "11973",
            "customer_country": "United States",
            "customer_latitude": null,
            "customer_longitude": null,
            "customer_alternate_phone": "",
            "pickup_location": "SR International",
            "package_instructions": null,
            "order_type": 0,
            "payment_status": "",
            "total": "1000.00",
            "tax": "0.00",
            "sla": "2 days",
            "shipping_method": "SR",
            "expedited": 0,
            "status": "DELIVERED",
            "status_code": 7,
            "awd_etds": null,
            "master_status": "FULFILLED",
            "payment_method": "prepaid",
            "change_payment_mode": false,
            "is_international": 1,
            "purpose_of_shipment": 0,
            "channel_created_at": "16 Dec 2021, 09:16 AM",
            "created_at": "16 Dec 2021, 09:16 AM",
            "updated_at": "3 Jan 2022, 09:55 AM",
            "pickup_boy_name": "",
            "pickup_boy_contact_no": "",
            "products": [
                {
                    "id": 212990291,
                    "channel_order_product_id": "212990291",
                    "name": "Tshirt",
                    "channel_sku": "sku300",
                    "quantity": 1,
                    "product_id": 84244695,
                    "available": 1,
                    "status": "UNDEFINED",
                    "price": "1000.00",
                    "product_cost": "1000.00",
                    "status_code": 1,
                    "hsn": ""
                }
            ],
            "shipments": [
                {
                    "id": 152757127,
                    "isd_code": "",
                    "courier": "SR International",
                    "courier_id": 140,
                    "shipping_charges": "",
                    "weight": "0.45",
                    "dimensions": "10x10x10",
                    "shipped_date": "2021-11-27 23:52:19",
                    "pickup_scheduled_date": "27/11/2021",
                    "pickup_token_number": null,
                    "awb": "1527571275",
                    "return_awb": "",
                    "volumetric_weight": 0.2,
                    "pod": null,
                    "etd": "NA",
                    "saral_etd": "NA",
                    "rto_delivered_date": "0000-00-00 00:00:00",
                    "delivered_date": "2021-12-08 05:40:00",
                    "etd_escalation_btn": false,
                    "rto_initiated_date": "0000-00-00 00:00:00",
                    "package_images": "",
                    "weight_action": null,
                    "status": 7,
                    "pickup_id": "",
                    "delivery_executive_name": "",
                    "delivery_executive_number": ""
                }
            ],
            "cod": 0,
            "activities": [
                "ORDER_CREATED",
                "LABEL_GENERATED",
                "LABEL_GENERATED",
                "LABEL_GENERATED",
                "ORDER_DELIVERED"
            ],
            "allow_return": 0,
            "is_incomplete": 0,
            "errors": [],
            "pickup_exception_reason": "",
            "rto_prediction": "",
            "show_escalation_btn": 0,
            "escalation_status": "",
            "pii_removed": 0,
            "allow_multiship": true,
            "others": {
                "weight": 0.45,
                "isd_code": "+1",
                "quantity": 1,
                "self_ship": false,
                "buyer_psid": null,
                "dimensions": "10x10x10",
                "billing_city": "UPTON",
                "billing_name": "Shubham OworldInt  ",
                "company_name": "",
                "billing_email": "shubham.tyagi+1@shiprocket.com",
                "billing_phone": "",
                "billing_state": "New York",
                "currency_code": "INR",
                "package_count": 1,
                "reseller_name": "",
                "customer_gstin": "",
                "billing_address": "223, alaKam, Aramex",
                "billing_country": "United States",
                "billing_pincode": "11973",
                "billing_isd_code": "+1",
                "shipping_charges": "0",
                "billing_address_2": "",
                "is_order_verified": 0,
                "order_verified_at": "",
                "billing_alternate_phone": ""
            },
            "package_list": {
                "packages": [],
                "selected": false,
                "select_item": false,
                "all_packages": []
            },
            "pickup_address_detail": {
                "id": 1883284,
                "pickup_code": "SR International",
                "gstin": null,
                "invoice_prefix": null,
                "invoice_serial": null,
                "address": "123, tetsing line",
                "address_2": "",
                "address_type": null,
                "city": "East Delhi",
                "state": "Delhi",
                "country": "India",
                "pin_code": "110032",
                "email": "shubham.tyagi@shiprocket.com",
                "phone": "8923309680",
                "phone_verified": 1,
                "alternate_phone": "",
                "name": "Shubham",
                "company_id": 2,
                "status": 1,
                "rto_address_id": 1883284,
                "delhivery_clientware_id": null,
                "delhivery_surface_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_standard_clientware_id": "KARTROCKETSMALL SURFACE-East Delhi-110032-Shubham-2-1883284",
                "delhivery_surface_lite_clientware_id": "KARTROCKETMEDIUM SURFACE-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_10kg_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_20kg_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "sx_delhivery_clientware_id": null,
                "sx_delhivery_surface_clientware_id": null,
                "sx4_delhivery_surface_clientware_id": null,
                "sx3_delhivery_surface_clientware_id": null,
                "sx2_delhivery_surface_clientware_id": null,
                "sx1_delhivery_surface_clientware_id": null,
                "created_at": "2021-11-15 11:13:56",
                "updated_at": "2022-05-10 19:15:43",
                "gati_surface_customer_vendor_code": null,
                "gati_sg_customer_vendor_code": null,
                "delhivery_flash_air_clientware_id": null,
                "updated_on": "2022-05-10 13:45:43",
                "lat": "",
                "long": "",
                "delhivery_essential_5kg_clientware_id": null,
                "warehouse_code": null,
                "gati_surface_10kg_customer_vendor_code": null,
                "extra_info": "{\"source\": 1}",
                "delhivery_documents_clientware_id": null,
                "delhivery_documents_250gm_clientware_id": null
            },
            "rto_address": [],
            "re_escalate": 0,
            "seller_request": {
                "show_pod_actions": false,
                "can_request_pod": false,
                "pod_dispute": "",
                "pod_requested": false,
                "status": "",
                "lops_remark": "",
                "status_value": "",
                "can_raise_dispute": false
            },
            "engage": "NA",
            "total_dead_weight": null,
            "total_volumetric_weight": null,
            "total_order_value": null,
            "boxes": null,
            "order_tag": [],
            "is_return": 0,
            "is_b2b": false,
            "is_insurance_opt": false,
            "manifest_id": "",
            "manifest_generated": false,
            "awb_data": {
                "charges": {
                    "zone": "",
                    "cod_charges": "",
                    "applied_weight_amount": "1486.00",
                    "freight_charges": "1486.00",
                    "applied_weight": "0.450",
                    "charged_weight": "",
                    "charged_weight_amount": "",
                    "charged_weight_amount_rto": "0.00",
                    "applied_weight_amount_rto": "0.00",
                    "billing_amount": "",
                    "service_type_id": ""
                }
            },
            "escalation_history": []
        },
        {
            "id": 153210135,
            "channel_id": 2252386,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_order_id": "4327343108",
            "customer_name": "Shubham NARAmex ",
            "customer_email": "shubham.tyagi+1@shiprocket.com",
            "customer_phone": "8923309680",
            "customer_address": "223, alaKam, Aramex",
            "customer_address_2": "",
            "customer_city": "UPTON",
            "customer_state": "New York",
            "customer_pincode": "11973",
            "customer_country": "United States",
            "customer_latitude": null,
            "customer_longitude": null,
            "customer_alternate_phone": "",
            "pickup_location": "SR International",
            "package_instructions": null,
            "order_type": 0,
            "payment_status": "",
            "total": "1000.00",
            "tax": "0.00",
            "sla": "2 days",
            "shipping_method": "SR",
            "expedited": 0,
            "status": "DELIVERED",
            "status_code": 7,
            "awd_etds": null,
            "master_status": "FULFILLED",
            "payment_method": "prepaid",
            "change_payment_mode": false,
            "is_international": 1,
            "purpose_of_shipment": 0,
            "channel_created_at": "1 Dec 2021, 10:09 AM",
            "created_at": "1 Dec 2021, 10:10 AM",
            "updated_at": "3 Jan 2022, 09:48 AM",
            "pickup_boy_name": "",
            "pickup_boy_contact_no": "",
            "products": [
                {
                    "id": 212990285,
                    "channel_order_product_id": "212990285",
                    "name": "Tshirt",
                    "channel_sku": "sku300",
                    "quantity": 1,
                    "product_id": 84244695,
                    "available": 1,
                    "status": "UNDEFINED",
                    "price": "1000.00",
                    "product_cost": "1000.00",
                    "status_code": 1,
                    "hsn": ""
                }
            ],
            "shipments": [
                {
                    "id": 152757121,
                    "isd_code": "",
                    "courier": "SR International",
                    "courier_id": 140,
                    "shipping_charges": "",
                    "weight": "0.31",
                    "dimensions": "11x10x10",
                    "shipped_date": "2021-11-04 05:25:57",
                    "pickup_scheduled_date": "01/12/2021",
                    "pickup_token_number": "Pickup confirmation number: L01B2FA, GUID: 02462e3d-b203-46c8-a823-dfa320a43577, Reference: Shipment Pickup - 2021-12-01",
                    "awb": "1527571212",
                    "return_awb": "",
                    "volumetric_weight": 0.22,
                    "pod": null,
                    "etd": "NA",
                    "saral_etd": "NA",
                    "rto_delivered_date": "0000-00-00 00:00:00",
                    "delivered_date": "2021-11-29 08:07:00",
                    "etd_escalation_btn": false,
                    "rto_initiated_date": "0000-00-00 00:00:00",
                    "package_images": "",
                    "weight_action": null,
                    "status": 7,
                    "pickup_id": "",
                    "delivery_executive_name": "",
                    "delivery_executive_number": ""
                }
            ],
            "cod": 0,
            "activities": [
                "ORDER_CREATED",
                "INVOICED",
                "LABEL_GENERATED",
                "PICKUP_QUEUE",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "ORDER_DELIVERED"
            ],
            "allow_return": 0,
            "is_incomplete": 0,
            "errors": [],
            "pickup_exception_reason": "",
            "rto_prediction": "",
            "show_escalation_btn": 0,
            "escalation_status": "",
            "pii_removed": 0,
            "allow_multiship": true,
            "others": {
                "weight": 0.311,
                "isd_code": "+1",
                "quantity": 1,
                "self_ship": false,
                "buyer_psid": null,
                "dimensions": "11x10x10",
                "billing_city": "UPTON",
                "billing_name": "Shubham NARAmex ",
                "company_name": "",
                "billing_email": "shubham.tyagi+1@shiprocket.com",
                "billing_phone": "",
                "billing_state": "New York",
                "currency_code": "INR",
                "package_count": 1,
                "reseller_name": "",
                "customer_gstin": "",
                "billing_address": "223, alaKam, Aramex",
                "billing_country": "United States",
                "billing_pincode": "11973",
                "billing_isd_code": "+1",
                "shipping_charges": "0",
                "billing_address_2": "",
                "is_order_verified": 0,
                "order_verified_at": "",
                "billing_alternate_phone": ""
            },
            "package_list": {
                "packages": [],
                "selected": false,
                "select_item": false,
                "all_packages": []
            },
            "pickup_address_detail": {
                "id": 1883284,
                "pickup_code": "SR International",
                "gstin": null,
                "invoice_prefix": null,
                "invoice_serial": null,
                "address": "123, tetsing line",
                "address_2": "",
                "address_type": null,
                "city": "East Delhi",
                "state": "Delhi",
                "country": "India",
                "pin_code": "110032",
                "email": "shubham.tyagi@shiprocket.com",
                "phone": "8923309680",
                "phone_verified": 1,
                "alternate_phone": "",
                "name": "Shubham",
                "company_id": 2,
                "status": 1,
                "rto_address_id": 1883284,
                "delhivery_clientware_id": null,
                "delhivery_surface_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_standard_clientware_id": "KARTROCKETSMALL SURFACE-East Delhi-110032-Shubham-2-1883284",
                "delhivery_surface_lite_clientware_id": "KARTROCKETMEDIUM SURFACE-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_10kg_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_20kg_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "sx_delhivery_clientware_id": null,
                "sx_delhivery_surface_clientware_id": null,
                "sx4_delhivery_surface_clientware_id": null,
                "sx3_delhivery_surface_clientware_id": null,
                "sx2_delhivery_surface_clientware_id": null,
                "sx1_delhivery_surface_clientware_id": null,
                "created_at": "2021-11-15 11:13:56",
                "updated_at": "2022-05-10 19:15:43",
                "gati_surface_customer_vendor_code": null,
                "gati_sg_customer_vendor_code": null,
                "delhivery_flash_air_clientware_id": null,
                "updated_on": "2022-05-10 13:45:43",
                "lat": "",
                "long": "",
                "delhivery_essential_5kg_clientware_id": null,
                "warehouse_code": null,
                "gati_surface_10kg_customer_vendor_code": null,
                "extra_info": "{\"source\": 1}",
                "delhivery_documents_clientware_id": null,
                "delhivery_documents_250gm_clientware_id": null
            },
            "rto_address": [],
            "re_escalate": 0,
            "seller_request": {
                "show_pod_actions": false,
                "can_request_pod": false,
                "pod_dispute": "",
                "pod_requested": false,
                "status": "",
                "lops_remark": "",
                "status_value": "",
                "can_raise_dispute": false
            },
            "engage": "NA",
            "total_dead_weight": null,
            "total_volumetric_weight": null,
            "total_order_value": null,
            "boxes": null,
            "order_tag": [],
            "is_return": 0,
            "is_b2b": false,
            "is_insurance_opt": false,
            "manifest_id": "MANIFEST-0032",
            "manifest_generated": true,
            "awb_data": {
                "charges": {
                    "zone": "",
                    "cod_charges": "",
                    "applied_weight_amount": "1297.00",
                    "freight_charges": "1297.00",
                    "applied_weight": "0.310",
                    "charged_weight": "",
                    "charged_weight_amount": "",
                    "charged_weight_amount_rto": "0.00",
                    "applied_weight_amount_rto": "0.00",
                    "billing_amount": "",
                    "service_type_id": ""
                }
            },
            "escalation_history": []
        },
        {
            "id": 153210130,
            "channel_id": 2252386,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_order_id": "7850384358",
            "customer_name": "Shubham NAramex ",
            "customer_email": "shubham.tyagi+1@shiprocket.com",
            "customer_phone": "8923309680",
            "customer_address": "223, alaKam, Aramex",
            "customer_address_2": "",
            "customer_city": "UPTON",
            "customer_state": "New York",
            "customer_pincode": "11973",
            "customer_country": "United States",
            "customer_latitude": null,
            "customer_longitude": null,
            "customer_alternate_phone": "",
            "pickup_location": "SR International",
            "package_instructions": null,
            "order_type": 0,
            "payment_status": "",
            "total": "990.00",
            "tax": "0.00",
            "sla": "2 days",
            "shipping_method": "SR",
            "expedited": 0,
            "status": "DELIVERED",
            "status_code": 7,
            "awd_etds": null,
            "master_status": "FULFILLED",
            "payment_method": "prepaid",
            "change_payment_mode": false,
            "is_international": 1,
            "purpose_of_shipment": 0,
            "channel_created_at": "25 Nov 2021, 05:02 AM",
            "created_at": "25 Nov 2021, 05:03 AM",
            "updated_at": "3 Jan 2022, 09:48 AM",
            "pickup_boy_name": "",
            "pickup_boy_contact_no": "",
            "products": [
                {
                    "id": 212990280,
                    "channel_order_product_id": "212990280",
                    "name": "Tshirt",
                    "channel_sku": "sku300",
                    "quantity": 1,
                    "product_id": 84244695,
                    "available": 1,
                    "status": "UNDEFINED",
                    "price": "1000.00",
                    "product_cost": "1000.00",
                    "status_code": 1,
                    "hsn": ""
                }
            ],
            "shipments": [
                {
                    "id": 152757115,
                    "isd_code": "",
                    "courier": "SR International",
                    "courier_id": 140,
                    "shipping_charges": "",
                    "weight": "0.301",
                    "dimensions": "10x10x10",
                    "shipped_date": "2021-11-04 05:25:33",
                    "pickup_scheduled_date": "01/12/2021",
                    "pickup_token_number": "Pickup confirmation number: L018C2A, GUID: 88da068e-fc6e-43ee-86db-2b09f1b1a4a4, Reference: Shipment Pickup - 2021-12-01",
                    "awb": "1527571157",
                    "return_awb": "",
                    "volumetric_weight": 0.2,
                    "pod": null,
                    "etd": "NA",
                    "saral_etd": "NA",
                    "rto_delivered_date": "0000-00-00 00:00:00",
                    "delivered_date": "2021-11-11 05:16:00",
                    "etd_escalation_btn": false,
                    "rto_initiated_date": "0000-00-00 00:00:00",
                    "package_images": "",
                    "weight_action": null,
                    "status": 7,
                    "pickup_id": "",
                    "delivery_executive_name": "",
                    "delivery_executive_number": ""
                }
            ],
            "cod": 0,
            "activities": [
                "ORDER_CREATED",
                "DIMENSIONS_EDITED",
                "INVOICED",
                "LABEL_GENERATED",
                "PICKUP_QUEUE",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED",
                "MANIFEST_SCHEDULED"
            ],
            "allow_return": 0,
            "is_incomplete": 0,
            "errors": [],
            "pickup_exception_reason": "",
            "rto_prediction": "",
            "show_escalation_btn": 0,
            "escalation_status": "",
            "pii_removed": 0,
            "allow_multiship": true,
            "others": {
                "weight": 0.301,
                "isd_code": "+1",
                "quantity": 1,
                "self_ship": false,
                "buyer_psid": null,
                "dimensions": "10x10x10",
                "billing_city": "UPTON",
                "billing_name": "Shubham NAramex ",
                "company_name": "",
                "billing_email": "shubham.tyagi+1@shiprocket.com",
                "billing_phone": "",
                "billing_state": "New York",
                "currency_code": "USD",
                "package_count": 1,
                "reseller_name": "",
                "customer_gstin": "",
                "billing_address": "223, alaKam, Aramex",
                "billing_country": "United States",
                "billing_pincode": "11973",
                "billing_isd_code": "+1",
                "shipping_charges": "0",
                "billing_address_2": "",
                "is_order_verified": 0,
                "order_verified_at": "",
                "billing_alternate_phone": ""
            },
            "package_list": {
                "packages": [],
                "selected": false,
                "select_item": false,
                "all_packages": []
            },
            "pickup_address_detail": {
                "id": 1883284,
                "pickup_code": "SR International",
                "gstin": null,
                "invoice_prefix": null,
                "invoice_serial": null,
                "address": "123, tetsing line",
                "address_2": "",
                "address_type": null,
                "city": "East Delhi",
                "state": "Delhi",
                "country": "India",
                "pin_code": "110032",
                "email": "shubham.tyagi@shiprocket.com",
                "phone": "8923309680",
                "phone_verified": 1,
                "alternate_phone": "",
                "name": "Shubham",
                "company_id": 2,
                "status": 1,
                "rto_address_id": 1883284,
                "delhivery_clientware_id": null,
                "delhivery_surface_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_standard_clientware_id": "KARTROCKETSMALL SURFACE-East Delhi-110032-Shubham-2-1883284",
                "delhivery_surface_lite_clientware_id": "KARTROCKETMEDIUM SURFACE-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_10kg_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "delhivery_surface_20kg_clientware_id": "Kart Rocket Surface-East Delhi-110032-Shiprocket-2-1883284",
                "sx_delhivery_clientware_id": null,
                "sx_delhivery_surface_clientware_id": null,
                "sx4_delhivery_surface_clientware_id": null,
                "sx3_delhivery_surface_clientware_id": null,
                "sx2_delhivery_surface_clientware_id": null,
                "sx1_delhivery_surface_clientware_id": null,
                "created_at": "2021-11-15 11:13:56",
                "updated_at": "2022-05-10 19:15:43",
                "gati_surface_customer_vendor_code": null,
                "gati_sg_customer_vendor_code": null,
                "delhivery_flash_air_clientware_id": null,
                "updated_on": "2022-05-10 13:45:43",
                "lat": "",
                "long": "",
                "delhivery_essential_5kg_clientware_id": null,
                "warehouse_code": null,
                "gati_surface_10kg_customer_vendor_code": null,
                "extra_info": "{\"source\": 1}",
                "delhivery_documents_clientware_id": null,
                "delhivery_documents_250gm_clientware_id": null
            },
            "rto_address": [],
            "re_escalate": 0,
            "seller_request": {
                "show_pod_actions": false,
                "can_request_pod": false,
                "pod_dispute": "",
                "pod_requested": false,
                "status": "",
                "lops_remark": "",
                "status_value": "",
                "can_raise_dispute": false
            },
            "engage": "NA",
            "total_dead_weight": null,
            "total_volumetric_weight": null,
            "total_order_value": null,
            "boxes": null,
            "order_tag": [],
            "is_return": 0,
            "is_b2b": false,
            "is_insurance_opt": false,
            "manifest_id": "MANIFEST-0032",
            "manifest_generated": true,
            "awb_data": {
                "charges": {
                    "zone": "",
                    "cod_charges": "",
                    "applied_weight_amount": "1297.00",
                    "freight_charges": "1297.00",
                    "applied_weight": "0.301",
                    "charged_weight": "",
                    "charged_weight_amount": "",
                    "charged_weight_amount_rto": "0.00",
                    "applied_weight_amount_rto": "0.00",
                    "billing_amount": "",
                    "service_type_id": ""
                }
            },
            "escalation_history": []
        }
    ],
    "meta": {
        "counts": {
            "nofilter_pickup_count": 1,
            "awbs": 1,
            "nofilter_manifested_count": 1,
            "nofilter_readytoship_count": 1,
            "shipped": 1,
            "total": 1,
            "nofilter_processing_count": 1,
            "is_first": false
        },
        "pagination": {
            "total": 3,
            "count": 3,
            "per_page": 15,
            "current_page": 1,
            "total_pages": 1,
            "links": {}
        }
    }
}
```

### International KYC

`POST https://apiv2.shiprocket.in/v1/external/international/settings/international_kyc`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Description**

Use this API for your international KYC. This API return your KYC status.<br>Please check the documentation for the documents for KYC.

| PARAMS | REQUIRED DATA | DATA TYPE | DESCRIPTION | EXAMPLE |
|---|---|---|---|---|
| organization_type | Yes | string | Seller organization | company |
| ip_address | Yes | int unsigned | Seller's ip_address | 192.168.1.1.0 |
| documents | Yes | array | Documentaion for organization type (document size is not greater than 3 MB) | Please check request |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "organization_type": "Sole Proprietor",
    "ip_address": "35.207.230.249",
    "documents": [
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "iec_code",
            "document_value": "0514039752"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "ad_code",
            "document_value": "63919662900009"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "gstin_certificate",
            "document_value": "07AESPG5142H1ZE"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "pan_card",
            "document_value": "AESPG5142H"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "aadhar_card",
            "document_value": "613168861273"
        }
    ]
}
```

_5 long string value(s) (76,508, 76,508, 76,508, 76,508, 76,508 chars; base64 file data in the source) were truncated by the generator. The original values are in the source collection._

**Example responses**

#### Sucess Response — HTTP 200 OK

Example request body:

```json
{
    "organization_type": "Sole_Proprietor",
    "ip_address": "35.207.230.249",
    "documents": [
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "iec_code",
            "document_value": "0514039752"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "ad_code",
            "document_value": "63919662900009"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "gstin_certificate",
            "document_value": "07AESPG5142H1ZE"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "pan_card",
            "document_value": "AESPG5142H"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "aadhar_card",
            "document_value": "613168861273"
        }
    ]
}
```

_5 long string value(s) (76,508, 76,508, 76,508, 76,508, 76,508 chars; base64 file data in the source) were truncated by the generator. The original values are in the source collection._

```json
{
    "success" : true, 
    "message" : "Document is successfully updated!"
}
```

#### Failure Response — (no HTTP status recorded in source)

Example request body:

```json
{
    "organization_type": "Seller_type",
    "ip_address": "35.207.230.249",
    "documents": [
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "iec_code",
            "document_value": "0514039752"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "ad_code",
            "document_value": "63919662900009"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "gstin_certificate",
            "document_value": "07AESPG5142H1ZE"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "pan_card",
            "document_value": "AESPG5142H"
        },
        {
            "attachment": [
                {
                    "file": "/9j/4AAQSkZJRgABAQAAAQABAAD//gA8Q1JFQVRPUjogZ2Qt...[TRUNCATED: 76,508 more chars omitted by the doc generator]",
                    "extension": "pdf"
                }
            ],
            "document_name": "aadhar_card",
            "document_value": "613168861273"
        }
    ]
}
```

_5 long string value(s) (76,508, 76,508, 76,508, 76,508, 76,508 chars; base64 file data in the source) were truncated by the generator. The original values are in the source collection._

```json
{
  "status_code": 422,
  "message": "Kyc failed",
  "errors": "Please check the documentation and provide valid organization_type"
}
```

### Add Bank Details

`POST https://apiv2.shiprocket.in/v1/external/international/settings/add-bank-details`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

ThisdocumentationprovidesinformationabouttheBankDetailsAPI,whichallowsyouto<br>addbankdetailsforinternationalsettings.TheAPIendpoint, parameters, request sample<br>data, success response , and failed response are outlined below.

| PARAMS | REQUIRED | DATA TYPE | DESCRIPTION | EXAMPLE |
|---|---|---|---|---|
| bank_account_type | YES | string | ankaccounttype.Mustbe either"saving"or"current". | "saving" |
| beneficiary_name | YES | string | Beneficiaryname. Alphabetsandspacesonly | "JohnDoe" |
| bank_ifsc_code | YES | string | BankIFSCcode.Must followaspecificpattern | "ABCD0123456" |
| bank_account_number | YES | integer | Bankaccountnumber. Shouldbebetween9and 18digits. | 1234567890 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "bank_account_type": "saving",
    "beneficiary_name": "JohnDoe",
    "bank_ifsc_code": "ABCD0123456",
    "bank_account_number": "1234567890"
}
```

**Example responses**

#### Add Bank Details (Sucess Response) — (no HTTP status recorded in source)

```json
{
    "success":true,
    "status_code":200,
    "errors":[],
    "message":"BankDetailsissuccessfullyupdated!"
}
```

#### Add Bank Details Failed Response (Failure Response) — (no HTTP status recorded in source)

Example request body:

```json
{
    "bank_account_type": "saving",
    "beneficiary_name": "JohnDoe",
    "bank_ifsc_code": "ABCD0123456",
    "bank_account_number": "123456"
}
```

```json
{
    "message": "AddBankAccountFailed",
    "errors": {
        "bank_account_number": [
            "Thebankaccountnumbermustbebetween9and18digits."
        ]
    },
    "status_code": 422,
    "success": false
}
```

### Create order

`POST https://apiv2.shiprocket.in/v1/external/international/orders/create/adhoc`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Description**

Use this API to create a quick custom order. Quick orders are the ones where we do not store the product details in the master catalogue.

You have to pass all the required params at the minimum to create a quick custom order. You can add additional parameters as per your preference.

**Note:**

- In case the 'shipping_is_billing' field is false, further shipping detail fields are required.

If no channel id is passed, the order will be assigned to the default custom channel. If the channel id is not known, use the 'Get All Channels' API to get the list of all integrated channels in your Shiprocket account.

- order_id field cannot be equal to an already existing id. Doing so does not change or affect the existing order.
- New orders cannot be created with order id's same as that of cancelled orders. If error 422 shows up despite filling in the correct details, consider changing the order_id.
- Be sure to input the correct calculated sub_total amount. The total is not calculated automatically through the API.
- The 'order_id' returned in the response is the Shiprocket order_id. Please save this order ID as we will use this in future API calls.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *string* | The order id you want to specify to the order. Max char: 50. (Avoid passing character values as this contradicts some other API calls). | 224477 or 224-477 |
| `order_date` | YES | *string* | The date of order creation in yyyy-mm-dd format. Time is additional. | 2019-07-24 11:11 |
| `pickup_location` | YES | *string* | The name of the pickup location added in your Shiprocket account. This cannot be a new location. | Jammu |
| `channel_id` | NO | *integer* | Mention this in case you need to assign the order to a particular channel. Deafult is 'Custom'. | 27022 |
| `comment` | NO | *string* | Option to add 'From' field to the shipment. To do this, enter the name in the following format: 'Reseller: [name]'. | Reseller: Divine |
| `reseller_name` | NO | *string* | The 'from' name if you want to print. Use 'Reseller: [name]' | Reseller: Divine |
| `company_name` | NO | *string* | Name of the company. | Amazon |
| `billing_customer_name` | YES | *string* | First name of the billed customer. | John |
| `billing_last_name` | NO | *string* | Last name of the billed customer. | Doe |
| `billing_address` | NO | *string* | address details of the billed customer. | Civil line, House 20 |
| `billing_address_2` | NO | *string* | Further address details of the billed customer. | Near Hokage House |
| `billing_city` | YES | *string* | Billing address city. Max char: 30. | New Delhi |
| `billing_pincode` | YES | *integer* | Pincode of the billing address. | 110002 |
| `billing_state` | YES | *string* | Billing address state. | Delhi |
| `billing_country` | YES | *string* | Billing address country. | India |
| `billing_email` | YES | *string* | Email address of the billed customer. | [John@doe.com](mailto:John@doe.com) |
| `billing_phone` | YES | *integer* | The phone number of the billing customer. | 9856321472 |
| `billing_alternate_phone` | NO | *integer* | Alternate phone number of the billing customer. | 8604690454 |
| `shipping_is_billing` | YES | *boolean* | Whether the shipping address is the same as billing address. 1 or 'true' for yes and 0 or 'false' for no. | true |
| `shipping_customer_name` | CONDITIONAL YES | *string* | Name of the customer the order is shipped to. Required in case billing is not same as shipping. | Jane |
| `shipping_last_name` | NO | *string* | Last name of the shipping customer. | Doe |
| `shipping_address` | CONDITIONAL YES | *string* | Address of the Shipping customer. Required in case billing is not same as shipping. | Lane number 69 |
| `shipping_address_2` | NO | *string* | Further address details of shipping customer. | Andheri |
| `billing_isd_code` | NO | *string* | ISD code of the billing address. | +91 |
| `shipping_city` | CONDITIONAL YES | *string* | Shipping address city. | Mumbai |
| `shipping_pincode` | CONDITIONAL YES | *integer* | Shipping address pincode. | 200912 |
| `shipping_country` | CONDITIONAL YES | *string* | Shipping address country. | India |
| `shipping_state` | CONDITIONAL YES | *string* | Shipping address state. | Maharashtra |
| `shipping_email` | CONDITIONAL YES | *string* | Email of the shipping customer. | [Jane@doe.com](mailto:Jane@doe.com) |
| `shipping_phone` | CONDITIONAL YES | *integer* | Phone no. of the shipping customer. |  |
| `longitude` | NO | *float* | Destination (Shipping address) Longitude. | 69.0747 |
| `latitude` | NO | *float* | Destination (Shipping address) Latitude | 22.4064 |
| `order_items` | YES | / | List of items and their relevant fields in the form of Array. | / |
| `name` | YES | *string* | Name of the product. | Jeans |
| `sku` | YES | *string* | The sku id of the product. | cbs123 |
| `units` | YES | *integer* | No of units that are to be shipped. | 10 |
| `selling_price` | YES | *integer* | The selling price per unit in Rupee. Inclusive of GST. | 900 |
| `discount` | NO | *integer* | The discount amount in Rupee. Inclusive of tax. | 10 |
| `tax` | NO | *integer* | The tax percentage on the item. | 5 |
| `hsn` | NO | *integer* | Harmonised System Nomenclature code. Used to determine the category of taxation the goods fall under. | 44122 |
| `payment_method` | YES | *string* | The method of payment. Can be either COD (Cash on delivery) Or Prepaid. | COD |
| `shipping_charges` | NO | *integer* | Shipping charges if any in Rupee. | 5 |
| `giftwrap_charges` | NO | *integer* | Giftwrap charges if any in Rupee. | 5 |
| `transaction_charges` | NO | *integer* | Transaction charges if any in Rupee. | 5 |
| `total_discount` | NO | *integer* | The total discount amount in Rupee. | 15 |
| `sub_total` | YES | *integer* | Calculated sub total amount in Rupee after deductions. | 9010 |
| `length` | YES | *float* | The length of the item in cms. Must be more than 0.5. | 10 |
| `breadth` | YES | *float* | The breadth of the item in cms. Must be more than 0.5. | 10 |
| `height` | YES | *float* | The height of the item in cms. Must be more than 0.5. | 10 |
| `weight` | YES | *float* | The weight of the item in kgs. Must be more than 0. | 2.5 |
| `ewaybill_no` | NO | *string* | Details relating to the shipment of goods. . | K92373490 |
| `customer_gstin` | NO | *string* | Goods and Services Tax Identification Number. | 29ABCDE1234F2Z5 |
| `invoice_number` | NO | *string* |  |  |
| `order_type` | NO | *string* | Key to differentiate between Essentials or Non Essentials Shipments. Order type can only be ESSENTIALS or NON ESSENTIALS. Please note it is case sensitive and blank values are allowed. | ESSENTIALS |
| `checkout_shipping_method` | NO | *string* | Only for SRF users. | a. SR_RUSH: SDD, NDD b. SR_STANDARD: Surface Delivery c. SR_EXPRESS: Air Delivery d. SR_QUICK: 3 hrs delivery |
| `what3words_address` | NO | *string* | What3words is a proprietary geocode system designed to identify any location on the surface of Earth with a resolution of about 3 meters. The system encodes geographic coordinates into three permanently fixed dictionary words. | toddler.geologist.animated |
| purpose_of_shipment | NO | interger | The purpose of the shipment. values are 0 - gift, 1- sample, commercial - 2. | 1 |
| currency | YES | string | The currency of the order. Possible values are INR,USD,GBP, EUR, AUD, CAD, SAR, AED,SGD | USD |
| reasonOfExport | YES | integer | The reason for the export. Possible values are 0 - BONAFIDE_SAMPLE, 1 - SAMPLE, 2 - GIFT, 3 - COMMERCIAL | 2 |
| commodity | NO | boolean | Indicates if the order is a commodity or not | true |
| mies | NO |  |  | true |
| igstPaymentStatus | NO | char | possible values are 'A'- not applicable, 'B'- LUT or Export under Bond, 'C'- Export Against Payment of IGST | A |
| Terms_Of_Invoice | YES | string | FOB and CIF | FOB |
| ioss | YES | string |  |  |
| eori | YES | string |  |  |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
   "order_id":172647058,
   "isd_code":"+1",
   "billing_isd_code":"",
   "order_date":"2022-03-30T00:34:12.311Z",
   "channel_id":2252386,
   "billing_customer_name":"Elena  ",
   "billing_last_name":"",
   "billing_address":"Plot No. 348, Panchkula, 134113, India",
   "billing_address_2":"",
   "billing_city":"Panchkula",
   "billing_state":"Haryana",
   "billing_country":"India",
   "billing_pincode":"134113",
   "landmark":"",
   "shipping_is_billing":1,
   "shipping_customer_name":"Elena ",
   "shipping_last_name":"",
   "shipping_address":"Plot No. 348, Panchkula, 134113, India",
   "shipping_address_2":"",
   "shipping_city":"Dallas",
   "order_type":1,
   "shipping_country":"United States",
   "shipping_pincode":"134090",
   "shipping_state":"Texas",
   "shipping_email":"test.test@shiprocket.com",
   "product_category":"",
   "shipping_phone":9762343722,
   "billing_alternate_phone":"",
   "order_items":[
      {
         "name":"Combo of 4 Way Spanner and Hydraulic trolly Jack",
         "sku":"5-47606",
         "category_name":"Default Category",
         "tax":"",
         "hsn":"",
         "units":"1",
         "selling_price":"100",
         "discount":"",
         "category_id":"",
         "caetgroy_code":""
      }
   ],
   "payment_method":"Prepaid",
   "shipping_charges":0,
   "giftwrap_charges":0,
   "transaction_charges":0,
   "total_discount":0,
   "sub_total":100,
   "weight":0.41,
   "length":10,
   "breadth":10,
   "height":10,
   "pickup_location_id":255,
   "reseller_name":"",
   "company_name":"",
   "ewaybill_no":"",
   "customer_gstin":"",
   "is_order_revamp":1,
   "is_document":0,
   "delivery_challan":false,
   "order_tag":"",
   "purpose_of_shipment":0,
   "currency":"USD",
   "reasonOfExport":2,
   "commodity":"true",
   "mies":"true",
   "igstPaymentStatus":"C",
   "Terms_Of_Invoice":"FOB",
   "is_insurance_opt":false
}
```

**Example responses**

#### Create order (Sucess Response) — HTTP 200 OK

Example request body:

```json
{
   "order_id":1726470481,
   "isd_code":"+1",
   "billing_isd_code":"",
   "order_date":"2022-03-30T00:34:12.311Z",
   "channel_id":2252386,
   "billing_customer_name":"Elena  ",
   "billing_last_name":"",
   "billing_address":"Plot No. 348, 134113, India",
   "billing_address_2":"",
   "billing_city":"Panchkula",
   "billing_state":"Haryana",
   "billing_country":"India",
   "billing_pincode":"134113",
   "billing_email":"test.test@shiprocket.com",
   "billing_phone":9760858933,
   "landmark":"",
   "shipping_is_billing":1,
   "shipping_customer_name":"Elena ",
   "shipping_last_name":"",
   "shipping_address":"Plot 1, Panchkula, Haryana 134113, India",
   "shipping_address_2":"",
   "shipping_city":"Dallas",
   "order_type":1,
   "shipping_country":"United States",
   "shipping_pincode":"134090",
   "shipping_state":"Texas",
   "shipping_email":"test.test@shiprocket.com",
   "product_category":"",
   "shipping_phone":9760853722,
   "billing_alternate_phone":"",
   "order_items":[
      {
         "name":"Combo of 4 Way Spanner and Hydraulic trolly Jack",
         "sku":"5-47606",
         "category_name":"Default Category",
         "tax":"",
         "hsn":"",
         "units":"1",
         "selling_price":"100",
         "discount":"",
         "category_id":"",
         "category_code":""
      }
   ],
   "payment_method":"Prepaid",
   "shipping_charges":0,
   "giftwrap_charges":0,
   "transaction_charges":0,
   "total_discount":0,
   "sub_total":100,
   "weight":0.41,
   "length":10,
   "breadth":10,
   "height":10,
   "pickup_location_id":255,
   "reseller_name":"",
   "company_name":"",
   "ewaybill_no":"",
   "customer_gstin":"",
   "is_order_revamp":1,
   "is_document":0,
   "delivery_challan":false,
   "order_tag":"",
   "purpose_of_shipment":0,
   "currency":"USD",
   "reasonOfExport":2,
   "is_insurance_opt":false
}
```

```json
{
    "order_id": 153210169,
    "shipment_id": 152757155,
    "status": "NEW",
    "status_code": 1,
    "onboarding_completed_now": 0,
    "awb_code": "",
    "courier_company_id": "",
    "courier_name": ""
}
```

#### Create order (Failure Response) — HTTP 422 Unprocessable Entity

Example request body:

```json
{
   "order_id":172647058,
   "isd_code":"+1",
   "billing_isd_code":"",
   "order_date":"2022-03-30T00:34:12.311Z",
   "channel_id":2252386,
   "billing_customer_name":"Elena  ",
   "billing_last_name":"",
   "billing_address":"Plot No. 348,  Panchkula, 134113, India",
   "billing_address_2":"",
   "billing_city":"Panchkula",
   "billing_state":"Haryana",
   "billing_country":"India",
   "billing_pincode":"134113",

   "landmark":"",
   "shipping_is_billing":1,
   "shipping_customer_name":"Elena ",
   "shipping_last_name":"",
   "shipping_address":"Plot No. 348, Panchkula, 134113, India",
   "shipping_address_2":"",
   "shipping_city":"Dallas",
   "order_type":1,
   "shipping_country":"United States",
   "shipping_pincode":"134090",
   "shipping_state":"Texas",
   "shipping_email":"test.test@shiprocket.com",
   "product_category":"",
   "shipping_phone":9760889722,
   "billing_alternate_phone":"",
   "order_items":[
      {
         "name":"Combo of 4 Way Spanner and Hydraulic trolly Jack",
         "sku":"5-47606",
         "category_name":"Default Category",
         "tax":"",
         "hsn":"",
         "units":"1",
         "selling_price":"100",
         "discount":"",
         "category_id":"",
         "caetgroy_code":""
      }
   ],
   "shipping_charges":0,
   "giftwrap_charges":0,
   "transaction_charges":0,
   "total_discount":0,
   "sub_total":100,
   "weight":0.41,
   "length":10,
   "breadth":10,
   "height":10,
   "pickup_location_id":255,
   "reseller_name":"",
   "company_name":"",
   "ewaybill_no":"",
   "customer_gstin":"",
   "is_order_revamp":1,
   "is_document":0,
   "delivery_challan":false,
   "order_tag":"",
   "purpose_of_shipment":0,
   "currency":"USD",
   "reasonOfExport":2,
   "is_insurance_opt":false   
}
```

```json
{
    "message": "Oops! Invalid Data.",
    "errors": {
        "payment_method": [
            "The payment method field is required."
        ]
    },
    "status_code": 422
}
```

### Update order

`POST https://apiv2.shiprocket.in/v1/external/international/orders/update/adhoc`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Description**

Use this API to update your orders. You have to pass all the required params at the minimum to create a quick custom order. You can add additional parameters as per your preference.

You can update only the order_items details before assigning the AWB (before Ready to Ship status). You can only update these key-value pairs i.e increase/decrease the quantity, update tax/discount, add/remove product items. Some params specific to international order are.

| PARAMS | REQUIRED | DATA TYPE | DESCRIPTION | EXAMPLE |
|---|---|---|---|---|
| purpose_of_shipment | NO | interger | The purpose of the shipment. values are 0 - gift, 1- sample, commercial - 2. | 1 |
| currency | YES | string | The currency of the order. Possible values are INR,USD,GBP, EUR, AUD, CAD, SAR, AED,SGD | USD |
| reasonOfExport | No | integer | The reason for the export. Possible values are 0 - BONAFIDE_SAMPLE, 1 - SAMPLE, 2 - GIFT, 3 - COMMERCIAL | 2 |
| commodity | NO | boolean | Indicates if the order is a commodity or not | true |
| mies | No | boolean |  | true |
| igstPaymentStatus | YES | char | possible values are 'A'- not applicable, 'B'- LUT or Export under Bond, 'C'- Export Against Payment of IGST | A |
| Terms_Of_Invoice | no | string | The term of invoice either FOB and CIF | FOB |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
   "order_id":1726470481,
   "isd_code":"+1",
   "billing_isd_code":"",
   "order_date":"2022-03-30T00:34:12.311Z",
   "channel_id":2252386,
   "billing_customer_name":"Elena  ",
   "billing_last_name":"",
   "billing_address":"Plot No. 348, 134113, India",
   "billing_address_2":"",
   "billing_city":"Panchkula",
   "billing_state":"Haryana",
   "billing_country":"India",
   "billing_pincode":"134113",
   "billing_email":"test.test@shiprocket.com",
   "billing_phone":9760858933,
   "landmark":"",
   "shipping_is_billing":1,
   "shipping_customer_name":"Elena ",
   "shipping_last_name":"",
   "shipping_address":"Plot 1, Panchkula, Haryana 134113, India",
   "shipping_address_2":"",
   "shipping_city":"Dallas",
   "order_type":1,
   "shipping_country":"United States",
   "shipping_pincode":"134090",
   "shipping_state":"Texas",
   "shipping_email":"test.test@shiprocket.com",
   "product_category":"",
   "shipping_phone":9760853722,
   "billing_alternate_phone":"",
   "order_items":[
      {
         "name":"Combo of 4 Way Spanner and Hydraulic trolly Jack",
         "sku":"5-47606",
         "category_name":"Default Category",
         "tax":"",
         "hsn":"",
         "units":"1",
         "selling_price":"100",
         "discount":"",
         "category_id":"",
         "caetgroy_code":""
      }
   ],
   "payment_method":"Prepaid",
   "shipping_charges":0,
   "giftwrap_charges":0,
   "transaction_charges":0,
   "total_discount":0,
   "sub_total":100,
   "weight":0.51,
   "length":20,
   "breadth":20,
   "height":20,
   "pickup_location_id":255,
   "reseller_name":"",
   "company_name":"",
   "ewaybill_no":"",
   "customer_gstin":"",
   "is_order_revamp":1,
   "is_document":0,
   "delivery_challan":false,
   "order_tag":"",
   "purpose_of_shipment":0,
   "currency":"EUR",
   "reasonOfExport":2,
   "commodity":"true",
   "mies":"true",
   "igstPaymentStatus":"C",
   "Terms_Of_Invoice":"FOB",
   "is_insurance_opt":false
}
```

**Example responses**

#### Update order (Sucess Response) — HTTP 200 OK

Example request body:

```json
{
   "order_id":1726470481,
   "isd_code":"+1",
   "billing_isd_code":"",
   "order_date":"2022-03-30T00:34:12.311Z",
   "channel_id":2252386,
   "billing_customer_name":"Elena  ",
   "billing_last_name":"",
   "billing_address":"Plot No. 348, 134113, India",
   "billing_address_2":"",
   "billing_city":"Panchkula",
   "billing_state":"Haryana",
   "billing_country":"India",
   "billing_pincode":"134113",
   "billing_email":"test.test@shiprocket.com",
   "billing_phone":9761858933,
   "landmark":"",
   "shipping_is_billing":1,
   "shipping_customer_name":"Elena ",
   "shipping_last_name":"",
   "shipping_address":"Plot 1, Panchkula, Haryana 134113, India",
   "shipping_address_2":"",
   "shipping_city":"Dallas",
   "order_type":1,
   "shipping_country":"United States",
   "shipping_pincode":"134090",
   "shipping_state":"Texas",
   "shipping_email":"test.test@shiprocket.com",
   "product_category":"",
   "shipping_phone":9711153722,
   "billing_alternate_phone":"",
   "order_items":[
      {
         "name":"Combo of 4 Way Spanner and Hydraulic trolly Jack",
         "sku":"5-47606",
         "category_name":"Default Category",
         "tax":"",
         "hsn":"",
         "units":"1",
         "selling_price":"100",
         "discount":"",
         "category_id":"",
         "caetgroy_code":""
      }
   ],
   "payment_method":"Prepaid",
   "shipping_charges":0,
   "giftwrap_charges":0,
   "transaction_charges":0,
   "total_discount":0,
   "sub_total":100,
   "weight":0.51,
   "length":20,
   "breadth":20,
   "height":20,
   "pickup_location_id":255,
   "reseller_name":"",
   "company_name":"",
   "ewaybill_no":"",
   "customer_gstin":"",
   "is_order_revamp":1,
   "is_document":0,
   "delivery_challan":false,
   "order_tag":"",
   "purpose_of_shipment":0,
   "currency":"EUR",
   "reasonOfExport":2,
   "is_insurance_opt":false
}
```

```json
{
    "success": true,
    "partially_update": true,
    "not_updated_fields": "isd_code ,billing_isd_code ,order_date ,channel_...[TRUNCATED: 714 more chars omitted by the doc generator]",
    "order_id": 153210169,
    "shipment_id": 152757155,
    "new_order_status": "NEW",
    "old_order_status": 1,
    "awb_code": "",
    "courier_company_id": "",
    "courier_name": ""
}
```

_1 long string value(s) (714 chars; base64 file data in the source) were truncated by the generator. The original values are in the source collection._

#### Update order (Failure Response) — HTTP 400 Bad Request

Example request body:

```json
{
   "order_id":1726470381,
   "isd_code":"+1",
   "billing_isd_code":"",
   "order_date":"2022-03-30T00:34:12.311Z",
   "channel_id":846,
   "billing_customer_name":"Elena  ",
   "billing_last_name":"",
   "billing_address":"Plot No. 348, Panchkula, Haryana 134113, India",
   "billing_address_2":"",
   "billing_city":"Panchkula",
   "billing_state":"Haryana",
   "billing_country":"India",
   "billing_pincode":"134113",
   "billing_email":"test.test@shiprocket.com",
   "billing_phone":9760856722,
   "landmark":"",
   "shipping_is_billing":1,
   "shipping_customer_name":"Elena ",
   "shipping_last_name":"",
   "shipping_address":"Plot No. 348, Panchkula, 134113, India",
   "shipping_address_2":"",
   "shipping_city":"Dallas",
   "order_type":1,
   "shipping_country":"United States",
   "shipping_pincode":"134090",
   "shipping_state":"Texas",
   "shipping_email":"ruchi.parijat@shiprocket.com",
   "product_category":"",
   "shipping_phone":9760833322,
   "billing_alternate_phone":"",
   "order_items":[
      {
         "name":"Combo of 4 Way Spanner and Hydraulic trolly Jack",
         "sku":"5-47606",
         "category_name":"Default Category",
         "tax":"",
         "hsn":"",
         "units":"1",
         "selling_price":"100",
         "discount":"",
         "category_id":"",
         "caetgroy_code":""
      }
   ],
   "payment_method":"Prepaid",
   "shipping_charges":0,
   "giftwrap_charges":0,
   "transaction_charges":0,
   "total_discount":0,
   "sub_total":100,
   "weight":0.41,
   "length":10,
   "breadth":10,
   "height":10,
   "pickup_location_id":255,
   "reseller_name":"",
   "company_name":"",
   "ewaybill_no":"",
   "customer_gstin":"",
   "is_order_revamp":1,
   "is_document":0,
   "delivery_challan":false,
   "order_tag":"",
   "purpose_of_shipment":0,
   "currency":"USD",
   "reasonOfExport":2,
   "is_insurance_opt":false
}
```

```json
{
    "message": "Order Update for this channel is not allowed",
    "status_code": 400
}
```

### International Wrapper API

`POST https://apiv2.shiprocket.in/v1/external/international/shipments/create/forward-shipment`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to do multiple tasks in one go, namely creating a quick order, requesting its shipment, pickup generation generating the label and the manifest for the same order.

This API integrates several other APIs to perform all these tasks together.

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "order_id": "320988727",
    "order_date": "2022-05-08 12:23",
    "channel_id": "",
    "billing_customer_name": "Jax",
    "billing_last_name": "Tank",
    "billing_address": "Dust2",
    "billing_city": "New Delhi",
    "billing_pincode": "442001",
    "billing_state": "Delhi",
    "billing_country": "India",
    "billing_email": "jax@counterstike.com",
    "billing_phone": "9988998899",
    "shipping_is_billing": false,
    "shipping_customer_name": "Elena ",
    "shipping_last_name": "",
    "shipping_address": "Plot 1, Panchkula, Haryana 134113, India",
    "shipping_address_2": "",
    "shipping_city": "UPTON",
    "order_type": 1,
    "shipping_country": "United States",
    "shipping_pincode": "11973",
    "shipping_state": "New York",
    "shipping_email": "test.test@shiprocket.com",
    "product_category": "",
    "shipping_phone": 9760853722,
    "order_items": [
        {
            "name": "Delta",
            "sku": "delta123",
            "units": 10,
            "selling_price": "1000",
            "hsn": "24567870"
        }
    ],
    "payment_method": "PREPAID",
    "sub_total": 40,
    "length": 10,
    "breadth": 10,
    "height": 10,
    "weight": 0.7,
    "pickup_location": "rtryttest",
    "vendor_details": {
        "email": "mayur.p@iksula.com",
        "phone": 9879879879,
        "name": "Coco Cookie",
        "address": "F2004 Street  Street 1 Street ",
        "address_2": "",
        "city": "delhi",
        "state": "new delhi",
        "country": "india",
        "pin_code": "442001",
        "pickup_location": "rtryttest"
    },
    "purpose_of_shipment": 0,
    "currency": "INR",
    "igstPaymentStatus": "A",
    "Terms_Of_Invoice": "FOB",
    "igst_amount": 10,
    "ioss": "IM1234567890",
    "pickup_location_id": 647
}
```

**Example responses**

#### International Wrapper API — (no HTTP status recorded in source)

Example request body:

```json
{
    "order_id": "320988727",
    "order_date": "2022-05-08 12:23",
    "channel_id": "",
    "billing_customer_name": "Jax",
    "billing_last_name": "Tank",
    "billing_address": "Dust2",
    "billing_city": "New Delhi",
    "billing_pincode": "442001",
    "billing_state": "Delhi",
    "billing_country": "India",
    "billing_email": "jax@counterstike.com",
    "billing_phone": "9988998899",
    "shipping_is_billing": false,
    "shipping_customer_name": "Elena ",
    "shipping_last_name": "",
    "shipping_address": "Plot 1, Panchkula, Haryana 134113, India",
    "shipping_address_2": "",
    "shipping_city": "UPTON",
    "order_type": 1,
    "shipping_country": "United States",
    "shipping_pincode": "11973",
    "shipping_state": "New York",
    "shipping_email": "test.test@shiprocket.com",
    "product_category": "",
    "shipping_phone": 9760853722,
    "order_items": [
        {
            "name": "Delta",
            "sku": "delta123",
            "units": 10,
            "selling_price": "1000",
            "hsn": "24567870"
        }
    ],
    "payment_method": "PREPAID",
    "sub_total": 40,
    "length": 10,
    "breadth": 10,
    "height": 10,
    "weight": 0.7,
    "pickup_location": "rtryttest",
    "vendor_details": {
        "email": "mayur.p@iksula.com",
        "phone": 9879879879,
        "name": "Coco Cookie",
        "address": "F2004 Street  Street 1 Street ",
        "address_2": "",
        "city": "delhi",
        "state": "new delhi",
        "country": "india",
        "pin_code": "442001",
        "pickup_location": "rtryttest"
    },
    "purpose_of_shipment": 0,
    "currency": "INR",
    "igstPaymentStatus": "A",
    "Terms_Of_Invoice": "Paid",
    "igst_amount": 10,
    "ioss": "IM1234567890",
    "pickup_location_id": 647
}
```

```json
{
    "pickup_location_added": 0,
    "order_created": 1,
    "awb_generated": 1,
    "label_generated": 1,
    "pickup_generated": 1,
    "manifest_generated": 1,
    "pickup_scheduled_date": "2022-12-16 09:00:00",
    "pickup_booked_date": null,
    "order_id": 53861,
    "shipment_id": 52367,
    "awb_code": "8329468061579",
    "courier_company_id": 140,
    "courier_name": "Shiprocket Premium",
    "assigned_date_time": "2023-10-10T08:14:00.278336Z",
    "applied_weight": 1,
    "cod": 0,
    "label_url": "https://kr-multichannel.s3.ap-southeast-1.amazonaws.com/1049/labels/1671100520_shipping-label-52367-8329468061579.pdf",
    "manifest_url": "https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/1049/manifest/MANIFEST-0032.pdf",
    "routing_code": "",
    "rto_routing_code": "",
    "pickup_token_number": 102725472
}
```

### Serviceability

`GET https://apiv2.shiprocket.in/v1/external/international/courier/serviceability?order_id=247825513`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Query parameters in the request template**

| Name | Example value |
|---|---|
| `order_id` | `247825513` |

**Description**

This API checks courier serviceability for international orders and displays them as a list.

**Notes:**

- 'cod' field must be 0 as COD is not available for international orders.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `weight` | YES | *integer* | The weight of the shipment. | 10 |
| `cod` | YES | *integer* | Cash on delivery status. Must be 0. | 0 |
| `delivery_country` | YES | *string* | The destination country ISO Alpha 2 code. | US |
| `order_id` | NO | *integer* | The Shiprocket order_id of the shipment if available. | 1 |
| `pickup_postcode` | NO | *integer* | Use this field to select a different pickup postcode other than the primary pickup address. | 2 |

**Example responses**

#### success — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/international/courier/serviceability?weight=0.5&cod=0&delivery_country=United States&pickup_postcode=110030`

```json
{
    "status": 200,
    "data": {
        "is_recommendation_enabled": 1,
        "recommended_by": {
            "id": 6,
            "title": "Recommendation By Shiprocket"
        },
        "child_courier_id": null,
        "recommended_courier_company_id": 140,
        "shiprocket_recommended_courier_id": 140,
        "recommendation_advance_rule": null,
        "available_courier_companies": [
            {
                "courier_company_id": 140,
                "courier_name": "SRX Premium",
                "mode": 1,
                "description": "",
                "min_weight": 0.05,
                "charge_weight": 0.5,
                "realtime_tracking": "Real Time",
                "delivery_boy_contact": "Not Available",
                "pod_available": "On Request",
                "call_before_delivery": "Not Available",
                "is_international": 1,
                "pickup_performance": 4.7,
                "delivery_performance": 4.7,
                "rto_performance": 4.7,
                "weight_cases": 4.7,
                "rating": 4.7,
                "blocked": 0,
                "first_mile_courier_option": null,
                "service_type": 1,
                "pickup_availability": 0,
                "etd": "Feb 17, 2024 - Feb 22, 2024",
                "estimated_delivery_days": "10 - 15",
                "etd_hours": 360,
                "rate": {
                    "courier_id": 140,
                    "id": 9314109,
                    "rate": "108.01",
                    "zone_rates": {
                        "roc": "109.01",
                        "default": "108.01"
                    },
                    "extra_info": {
                        "edd": {
                            "to": 15,
                            "from": 10
                        }
                    },
                    "zone": "default"
                },
                "coverage_charges": 0,
                "insurance_applicable": 0,
                "courier_auto_secure": 0,
                "base_courier_id": null
            },
            {
                "courier_company_id": 326,
                "courier_name": "India Post EMS Merchandise",
                "mode": 1,
                "description": "",
                "min_weight": 0.5,
                "charge_weight": 0.5,
                "realtime_tracking": "MIS",
                "delivery_boy_contact": "Not Available",
                "pod_available": "On Request",
                "call_before_delivery": "Not Available",
                "is_international": 1,
                "pickup_performance": 3,
                "delivery_performance": 3,
                "rto_performance": 3,
                "weight_cases": 3,
                "rating": 3,
                "blocked": 0,
                "first_mile_courier_option": null,
                "service_type": null,
                "pickup_availability": 0,
                "etd": "Feb 11, 2024 - Feb 14, 2024",
                "estimated_delivery_days": "4 - 7",
                "etd_hours": 168,
                "rate": {
                    "courier_id": 326,
                    "id": 6741021,
                    "rate": 2330.6,
                    "zone_rates": {
                        "roc": "2324.6",
                        "default": "2324.6"
                    },
                    "extra_info": {
                        "edd": {
                            "to": 7,
                            "from": 4
                        }
                    },
                    "zone": "default",
                    "first_mile_charge": 6,
                    "last_mile_charge": "2324.6",
                    "total": 2330.6,
                    "first_mile_charge_uid": 5898
                },
                "coverage_charges": 0,
                "insurance_applicable": 0,
                "courier_auto_secure": 0,
                "base_courier_id": null
            },
            {
                "courier_company_id": 327,
                "courier_name": "India Post Air Parcel",
                "mode": 1,
                "description": "",
                "min_weight": 0.5,
                "charge_weight": 0.5,
                "realtime_tracking": "MIS",
                "delivery_boy_contact": "Not Available",
                "pod_available": "On Request",
                "call_before_delivery": "Not Available",
                "is_international": 1,
                "pickup_performance": 3,
                "delivery_performance": 3,
                "rto_performance": 3,
                "weight_cases": 3,
                "rating": 3,
                "blocked": 0,
                "first_mile_courier_option": null,
                "service_type": null,
                "pickup_availability": 0,
                "etd": "Feb 15, 2024 - Feb 22, 2024",
                "estimated_delivery_days": "8 - 15",
                "etd_hours": 360,
                "rate": {
                    "courier_id": 327,
                    "id": 7198630,
                    "rate": 1368.8000000000002,
                    "zone_rates": {
                        "roc": "1333.400",
                        "default": "1333.400"
                    },
                    "extra_info": null,
                    "zone": "default",
                    "first_mile_charge": 35.4,
                    "last_mile_charge": "1333.400",
                    "total": 1368.8000000000002,
                    "first_mile_charge_uid": 5341
                },
                "coverage_charges": 0,
                "insurance_applicable": 0,
                "courier_auto_secure": 0,
                "base_courier_id": null
            },
            {
                "courier_company_id": 328,
                "courier_name": "India Post Regd. Small Packet",
                "mode": 1,
                "description": "",
                "min_weight": 0.5,
                "charge_weight": 0.5,
                "realtime_tracking": "MIS",
                "delivery_boy_contact": "Not Available",
                "pod_available": "On Request",
                "call_before_delivery": "Not Available",
                "is_international": 1,
                "pickup_performance": 3,
                "delivery_performance": 3,
                "rto_performance": 3,
                "weight_cases": 3,
                "rating": 3,
                "blocked": 0,
                "first_mile_courier_option": null,
                "service_type": null,
                "pickup_availability": 0,
                "etd": "Feb 15, 2024 - Feb 22, 2024",
                "estimated_delivery_days": "8 - 15",
                "etd_hours": 360,
                "rate": {
                    "courier_id": 328,
                    "id": 7349730,
                    "rate": 1079.7,
                    "zone_rates": {
                        "roc": "1044.30",
                        "default": "1044.30"
                    },
                    "extra_info": null,
                    "zone": "default",
                    "first_mile_charge": 35.4,
                    "last_mile_charge": "1044.30",
                    "total": 1079.7,
                    "first_mile_charge_uid": 5340
                },
                "coverage_charges": 0,
                "insurance_applicable": 0,
                "courier_auto_secure": 0,
                "base_courier_id": null
            },
            {
                "courier_company_id": 329,
                "courier_name": "India Post Tracked Packet Service",
                "mode": 1,
                "description": "",
                "min_weight": 0.5,
                "charge_weight": 0.5,
                "realtime_tracking": "MIS",
                "delivery_boy_contact": "Not Available",
                "pod_available": "On Request",
                "call_before_delivery": "Not Available",
                "is_international": 1,
                "pickup_performance": 3,
                "delivery_performance": 3,
                "rto_performance": 3,
                "weight_cases": 3,
                "rating": 3,
                "blocked": 0,
                "first_mile_courier_option": null,
                "service_type": null,
                "pickup_availability": 0,
                "etd": "Feb 15, 2024 - Feb 22, 2024",
                "estimated_delivery_days": "8 - 15",
                "etd_hours": 360,
                "rate": {
                    "courier_id": 329,
                    "id": 6107371,
                    "rate": 1257.5,
                    "zone_rates": {
                        "roc": "1235.1",
                        "default": "1222.1"
                    },
                    "extra_info": null,
                    "zone": "default",
                    "first_mile_charge": 35.4,
                    "last_mile_charge": "1222.1",
                    "total": 1257.5,
                    "first_mile_charge_uid": 4837
                },
                "coverage_charges": 0,
                "insurance_applicable": 0,
                "courier_auto_secure": 0,
                "base_courier_id": null
            }
        ],
        "company_auto_shipment_insurance_setting": false,
        "eligible_for_insurance": false,
        "insurace_opted_at_order_creation": false,
        "user_insurance_manadatory": false
    },
    "currency": ""
}
```

#### Invalid data — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/international/courier/serviceability?order_id=52631`

```json
{
    "message": "Order does not exist!",
    "status_code": 404
}
```

### AWB Assignment

`POST https://apiv2.shiprocket.in/v1/external/international/courier/assign/awb`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `accept` | `application/json` |

**Description**

This API can be used to assign the AWB (Air Waybill Number) to your shipment. The AWB is a unique number that helps you track the shipment and get details about it.

| PARAMS | REQUIRED | DATA TYPE | DESCRIPTION | EXAMPLE |
|---|---|---|---|---|
| shipment_id | yes | integer | The shipment id of the order you want to create the AWB for. | 1603434 |
| courier_id | no | integer | The courier id of the courier service you want to select. The default courier is selected in case no id is specified. | 35 |
| status | no | string | Use this to change the courier of a shipment. Value: reassign. Note that this can be done only once in 24 hours. | reassign |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "shipment_id": 160169474,
    "courier_id": 332,
    "status": "reassign"
}
```

**Example responses**

#### AWB Assignment Successfully — (no HTTP status recorded in source)

Example request: `POST https://apiv2.shiprocket.in/external/international/courier/assign/awb`

Example request body:

```json
{
    "shipment_id": 160169474,
    "courier_id": 332,
    "status": ""
}
```

```json
{
    "awb_assign_status": 1,
    "response": {
        "data": {
            "courier_company_id": 10,
            "awb_code": "1091208940593",
            "cod": 0,
            "order_id": 181771297,
            "shipment_id": 160169474,
            "awb_code_status": 1,
            "assigned_date_time": {
                "date": "2022-05-10 11:18:37.397226",
                "timezone_type": 3,
                "timezone": "Asia/Kolkata"
            },
            "applied_weight": 1,
            "company_id": 25149,
            "courier_name": "Delhivery",
            "child_courier_name": null,
            "routing_code": "DEL/KIS",
            "rto_routing_code": "",
            "invoice_no": "test5769122383",
            "transporter_id": "06AAPCS9575E1ZR",
            "transporter_name": "Delhivery",
            "shipped_by": {
                "shipper_company_name": "New RtO",
                "shipper_address_1": "34- house",
                "shipper_address_2": "",
                "shipper_city": "South West Delhi",
                "shipper_state": "Delhi",
                "shipper_country": "India",
                "shipper_postcode": "110030",
                "shipper_first_mile_activated": 0,
                "shipper_phone": "7777777777",
                "lat": "28.517677",
                "long": "77.175261",
                "shipper_email": "new@rto.com",
                "rto_company_name": "New RtO",
                "rto_address_1": "34- house",
                "rto_address_2": "",
                "rto_city": "South West Delhi",
                "rto_state": "Delhi",
                "rto_country": "India",
                "rto_postcode": "110030",
                "rto_phone": "8888888888",
                "rto_email": "new@rto.com"
            }
        }
    }
}
```

#### AWB Assignment Failed — HTTP 200 OK

Example request: `POST https://apiv2.shiprocket.in/external/international/courier/assign/awb`

Example request body:

```json
{
    "shipment_id": [152757126],
    "courier_id": 140,
    "status": "reassign"
}
```

```json
{
    "awb_assign_status": 0,
    "response": {
        "data": {
            "awb_assign_error": "Awb Assignment Failed"
        }
    }
}
```

#### Invalid Courier Provided — HTTP 400 Bad Request

Example request: `POST https://apiv2.shiprocket.in/external/international/courier/assign/awb`

Example request body:

```json
{
    "shipment_id": [152757135],
    "courier_id": 32
}
```

```json
{
    "message": "Uh-oh! Invalid Courier for international shipment!"
}
```

### Manifest Generation

`POST https://apiv2.shiprocket.in/v1/external/international/manifests/generate`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Description**

Use this API to create a pickup request for your order shipment. The API returns the pickup status along with the estimated pickup time.<br>You will have to call the 'Generate Manifest' API after the successful response of this API.

**Note:**

- The AWB must be already generated for the shipment id to generate the pickup request.
- Only one shipment_id can be passed at a time.

| PARAMS | REQUIRED | DATA TYPE | DESCRIPTION | EXAMPLE |
|---|---|---|---|---|
| shipment_id | yes | integer | The shipment id of the shipment which is requested for pickup. | 1603434 |
| status | no | string | Use this field to retry if the pickup request fails. Value: retry | retry |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "shipment_id": [12345]
}
```

**Example responses**

#### Manifest Generated — HTTP 200 OK

Example request: `POST https://apiv2.shiprocket.in/external/international/manifests/generate`

Example request body:

```json
{
    "shipment_id": [152757135]
}
```

```json
{
    "status": 1,
    "manifest_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/25149/manifest/MANIFEST-3051.pdf"
}
```

#### Manisfest Not Generated — HTTP 200 OK

Example request: `POST https://apiv2.shiprocket.in/external/international/manifests/generate`

Example request body:

```json
{
    "shipment_id": [152757105]
}
```

```json
{
    "message": "Manifest not generated",
    "check_ids": [
        152757105
    ]
}
```

#### Manifest Already Generated — HTTP 400 Bad Request

Example request: `POST https://apiv2.shiprocket.in/external/international/manifests/generate`

Example request body:

```json
{
    "shipment_id": [152757121]
}
```

```json
{
    "message": "Manifest already generated.",
    "status_code": 400,
    "already_manifested_shipment_ids": [
        152757121
    ]
}
```

### Generate Pickup

`POST https://apiv2.shiprocket.in/v1/external/courier/generate/pickup`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `accept` | `application/json` |

**Description**

Use this API to create a pickup request for your order shipment. The API returns the pickup status along with the estimated pickup time.<br>You will have to call the 'Generate Manifest' API after the successful response of this API.

**Note:**

- The AWB must be already generated for the shipment id to generate the pickup request.
- Only one shipment_id can be passed at a time.

| PARAMS | REQUIRED | DATA TYPE | DESCRIPTION | EXAMPLE |
|---|---|---|---|---|
| shipment_id | yes | integer | The shipment id of the shipment which is requested for pickup. | 1603434 |
| status | no | string | Use this field to retry if the pickup request fails. Value: retry | retry |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "shipment_id": [12847483]
}
```

**Example responses**

#### Duplicate Request For Pickup Error — HTTP 500 Internal Server Error

Example request: `POST https://apiv2.shiprocket.in/external/international/courier/generate/pickup`

Example request body:

```json
{
    "shipment_id": 152757127
}
```

```json
{
    "message": "Duplicate request for pickup generation"
}
```

#### Pickup Generated Successfully — HTTP 500 Internal Server Error

Example request: `POST https://apiv2.shiprocket.in/external/international/courier/generate/pickup`

```json
{
    "pickup_token_number": "12345",
    "pickup_scheduled_date": "2022-05-17",
    "pickup_generated": 1,
    "shipment_id": 152757127,
    "manifest_generated": 1,
    "manifest_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/25149/manifest/MANIFEST-3051.pdf"
}
```

#### Pickup Not Generated Due To AWB Not Assigned — HTTP 400 Bad Request

Example request: `POST https://apiv2.shiprocket.in/external/international/courier/generate/pickup`

Example request body:

```json
{
    "shipment_id": 152757131
}
```

```json
{
    "message": "Awb not Assigned"
}
```

#### Already In PickUp Queue — HTTP 400 Bad Request

Example request: `POST https://apiv2.shiprocket.in/external/international/courier/generate/pickup`

Example request body:

```json
{
    "shipment_id": 152757133
}
```

```json
{
    "message": "Already in Pickup Queue."
}
```
