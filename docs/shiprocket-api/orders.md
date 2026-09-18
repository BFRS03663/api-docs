# Shiprocket API — Orders, Returns & Exchanges

> Generated from the published Postman collection behind [https://apidocs.shiprocket.in/](https://apidocs.shiprocket.in/) (collection "Shiprocket API", published id `SzYW1zB2`, version tag `latest`). Only content present in that collection is reproduced here; where the source omits something, the omission is stated.

## Create Or Update Order

Using these APIs, you can create new orders in your Shiprocket account or update existing orders.<br>You can also cancel an order, bulk import orders from a CSV file and update inventory for an ordered product.

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
| `channel_id` | NO | *integer* | Mention this in case you need to assign the order to a particular channel. Default is 'Custom'. | 27022 |
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
| `customer_gstin` | NO | *string* | Goods and Services Tax Identification Number. | 29ABCDE1234F2Z5 |
| `invoice_number` | NO | *string* |  |  |
| `order_type` | NO | *string* | Key to differentiate between Essentials or Non Essentials Shipments. Order type can only be ESSENTIALS or NON ESSENTIALS. Please note it is case sensitive and blank values are allowed. | ESSENTIALS |
| `checkout_shipping_method` | NO | *string* | Only for SRF users. | a. SR_RUSH: SDD, NDD b. SR_STANDARD: Surface Delivery c. SR_EXPRESS: Air Delivery d. SR_QUICK: 3 hrs delivery |
| `what3words_address` | NO | *string* | What3words is a proprietary geocode system designed to identify any location on the surface of Earth with a resolution of about 3 meters. The system encodes geographic coordinates into three permanently fixed dictionary words. | toddler.geologist.animated |
| `is_insurance_opt` | NO | *boolean* | To secure shipments above the order value of Rs 2500 | true |
| `is_document` | NO | *integer* | To create a document order | 1 or 0 |
| `order_tag` | NO | *string* | To add tags to your orders | abc, xyz |
| `reseller_name` | NO | *string* | To display the vendor name on the label | brandname |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "order_id": "",
    "order_date": "",
    "pickup_location": "",
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
    "order_type":""
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
  "comment": "Reseller: M/s Goku",
  "billing_customer_name": "Naruto",
  "billing_last_name": "Uzumaki",
  "billing_address": "House 221B, Leaf Village",
  "billing_address_2": "Near Hokage House",
  "billing_city": "New Delhi",
  "billing_pincode": 110002,
  "billing_state": "Delhi",
  "billing_country": "India",
  "billing_email": "naruto@uzumaki.com",
  "billing_phone": 9876543210,
  "shipping_is_billing": true,
  "shipping_customer_name": "",
  "shipping_last_name": "",
  "shipping_address": "",
  "shipping_address_2": "",
  "shipping_city": "",
  "shipping_pincode":"",
  "shipping_country": "",
  "shipping_state": "",
  "shipping_email": "",
  "shipping_phone": "",
  "order_items": [
    {
      "name": "Kunai",
      "sku": "chakra123",
      "units": 10,
      "selling_price": 900,
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

Example request body:

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

_(empty response body in source)_

### Create Channel Specific Order

`POST https://apiv2.shiprocket.in/v1/external/orders/create`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API can be used to create a custom order, the same as the Custom order API, except that you have to specify and select a custom channel to create the order.

The order created will be added under the specified channel. All the other parameters are the same.

**Note:**

- Channel_id field is required.
- Order_id cannot be the same as the already existing order id.
- Inventory Sync must be turned on to use this API. This can be done under the 'Channels' portion on the left-hand panel of your Shiprocket account.
- Inventory details of your Shiprocket account can be accessed using the 'Get Inventory Details' API.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *string* | The order id you want to specify to the order. Max char: 20. (Avoid passing character values as this contradicts some other API calls). | 224477 or 224-477 |
| `order_date` | YES | *string* | The date of order creation in yyyy-mm-dd format. Time is an additional option. | 2019-07-24 11:11 |
| `pickup_location` | NO | *string* | The name of the pickup location added in your Shiprocket account. This cannot be a new location. Default Pickup location is selected in case the parameter is not filled. | Jammu |
| `channel_id` | YES | *integer* | The id of the specific channel to be selected. | 27022 |
| `comment` | NO | *string* | Option to add 'From' field to the shipment. To do this, enter the name in the following format: 'Reseller: [name].' | Reseller: Divine |
| `billing_customer_name` | YES | *string* | First name of the customer who is billed. | John |
| `billing_last_name` | NO | *string* | Last name of the billed customer. | Doe |
| `billing_address` | YES | *string* | Primary address of the billed customer. Min char: 3. | House 221B, Leaf Village |
| `billing_address_2` | NO | *string* | Further address details of the billed customer. | Near Hokage House |
| `billing_city` | YES | *string* | Billing address city. Max char: 30. | New Delhi |
| `billing_pincode` | YES | *integer* | Pincode of the billing address. | 110002 |
| `billing_state` | YES | *string* | Billing address state. | Delhi |
| `billing_country` | YES | *string* | Billing address country. | India |
| `billing_email` | YES | *string* | Email address of the billed customer. | [John@doe.com](https://mailto:John@doe.com) |
| `billing_phone` | YES | *integer* | Phone number of the billed customer. | 9876543210 |
| `shipping_is_billing` | YES | *boolean* | Whether the shipping address is the same as billing address. 1 or 'true' for yes and 0 or 'false' for no. | true |
| `shipping_customer_name` | CONDITIONAL YES | *string* | Name of the customer the order is shipped to. Required in case billing is not same as shipping. | Jane |
| `shipping_last_name` | NO | *string* | Last name of the shipping customer. | Doe |
| `shipping_address` | CONDITIONAL YES | *string* | Address of the Shipping customer. Required in case billing is not same as shipping. | Lane 69 |
| `shipping_address_2` | NO | *string* | Further address details of shipping customer. | Andheri |
| `shipping_city` | CONDITIONAL YES | *string* | Shipping address city. | Mumbai |
| `shipping_pincode` | CONDITIONAL YES | *integer* | Shipping address pincode. | 200912 |
| `shipping_country` | CONDITIONAL YES | *string* | Shipping address country. | India |
| `shipping_state` | CONDITIONAL YES | *string* | Shipping address state. | Maharashtra |
| `shipping_email` | CONDITIONAL YES | *string* | Email of the shipping customer. | [jane@doe.com](https://mailto:jane@doe.com) |
| `shipping_phone` | CONDITIONAL YES | *integer* | Phone no. of the shipping customer | 9887655432 |
| `order_items` | YES | / | Array containing further fields. | / |
| `name` | YES | *string* | Name of the product. | Jeans |
| `sku` | YES | *string* | The sku id of the product. | cbs123 |
| `units` | YES | *integer* | No. of units that are to be shipped. | 10 |
| `selling_price` | YES | *integer* | The selling price per unit in Rupee. Inclusive of GST. | 900 |
| `discount` | NO | *integer* | The discount amount in Rupee. Inclusive of tax. | 10 |
| `tax` | NO | *integer* | The tax percentage on the item. | 5 |
| `hsn` | NO | *integer* | Harmonised System Nomenclature code. Used to determine the category of taxation the goods fall under. | 44122 |
| `payment_method` | YES | *string* | The method of payment. Can be either COD (Cash on delivery) Or Prepaid. | COD |
| `shipping_charges` | NO | *integer* | Shipping charges if any in Rupee. | 5 |
| `giftwrap_charges` | NO | *integer* | Giftwrap charges if any in Rupee. | 5 |
| `transaction_charges` | NO | *integer* | Transaction charges if any in Rupee. | 5 |
| `total_discount` | NO | *integer* | The total discount amount in Rupee. | 15 |
| `sub_total` | YES | *integer* | Calculated sub total amount in Rupee. | 9010 |
| `length` | YES | *integer* | The length of the item in cms. Must be more than 0.5. | 10 |
| `breadth` | YES | *integer* | The breadth of the item in cms. Must be more than 0.5. | 10 |
| `height` | YES | *integer* | The height of the item in cms. Must be more than 0.5. | 10 |
| `weight` | YES | *integer* | The weight of the item in kgs. Must be more than 0. | 2.5 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "order_id": "3167",
  "order_date": "2020-01-14 13:25",
  "pickup_location": "mrj",
  "channel_id": "443555",
  "comment": "fast and furious",
  "billing_customer_name": "rahul",
  "billing_last_name": "",
  "billing_address": "malviya nagar",
  "billing_address_2": "",
  "billing_city": "new delhi",
  "billing_pincode": "273303",
  "billing_state": "delhi",
  "billing_country": "india",
  "billing_email": "raushanra4@gmail.com",
  "billing_phone": "9721562372",
  "shipping_is_billing": 1,
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
      "name": "shoes",
      "sku": "shoes123",
      "units": "2",
      "selling_price": "1500",
      "discount": "100",
      "tax": "50",
      "hsn": ""
    }
  ],
  "payment_method": "COD",
  "shipping_charges": "",
  "giftwrap_charges": "",
  "transaction_charges": "",
  "total_discount": "",
  "sub_total": "2950",
  "length": "10",
  "breadth": "10",
  "height": "10",
  "weight": "1.5"
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
  "order_id": "224-4779",
  "order_date": "2019-07-24 11:11",
  "pickup_location": "Jammu",
  "channel_id": "76893",
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
    "order_id": 16161717,
    "shipment_id": 16000061,
    "status": "NEW",
    "status_code": 1
}
```

#### Invalid Data — HTTP 422 Unprocessable Entity

Example request body:

```json
{
  "order_id": "224-4779",
  "order_date": "2019-07-24 11:11",
  "pickup_location": "Jammu",
  "channel_id": "123321",
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
        "channel_id": [
            "The selected channel id is invalid."
        ]
    },
    "status_code": 422
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

Example request body:

```json
{
  "order_id": "224-4779",
  "order_date": "2019-07-24 11:11",
  "pickup_location": "Delhi",
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
        "channel_id": [
            "The channel id field is required."
        ]
    },
    "status_code": 422
}
```

#### Inventory Sync Error — HTTP 400 Bad Request

Example request body:

```json
{
  "order_id": "224-4779",
  "order_date": "2019-07-24 11:11",
  "pickup_location": "Delhi",
  "channel_id": "27202",
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
    "message": "Inventory sync is turned off. Please add a manual order!",
    "status_code": 400
}
```

### Change/Update Pickup Location of Created Orders

`PATCH https://apiv2.shiprocket.in/v1/external/orders/address/pickup`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Using this API, you can modify the pickup location of an already created order. Multiple order ids can be passed to update their pickup location together.

**Note:**

- Pickup location can only be changed/updated to an already existing pickup location in your account.
- The 'order_id' to be passed is the Shiprocket order_id received at the time of order creation.
- Multiple order ids can be passed as an array, separated by commas. eg: ["141414,142424,143434"]

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *integer* | The Shiprocket order_id specified to the order. | 16167171 |
| `pickup_location` | YES | *string* | The pickup location you want to change your current pickup location to. | Delhi |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "order_id": [],
  "pickup_location": ""
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
  "order_id": [16167171],
  "pickup_location": "Delhi"
}
```

```json
{
    "message": "Pickup location Updated"
}
```

#### Invalid Data — HTTP 400 Bad Request

Example request body:

```json
{
  "order_id": [16167171],
  "pickup_location": "Mumbai"
}
```

```json
{
    "message": "Pickup Code does not exist",
    "status_code": 400
}
```

#### Missing Fields — HTTP 400 Bad Request

```json
{
    "message": "Order Id does not exists",
    "status_code": 400
}
```

#### Wrong Format — HTTP 500 Internal Server Error

Example request body:

```json
{
  "order_id": "[16176659,16177223]",
  "pickup_location": "Delhi"
}
```

```json
{
    "message": "Invalid argument supplied for foreach()",
    "status_code": 500
}
```

### Update Customer Delivery Address

`POST https://apiv2.shiprocket.in/v1/external/orders/address/update`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

You can update the customer's name and delivery address through this API by passing the Shiprocket order id and the necessary customer details.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *integer* | The Shiprocket order_id specified to the order. | 16178831 |
| `shipping_customer_name` | YES | *string* | The name of the customer. | John Doe |
| `shipping_phone` | YES | *integer* | Phone number of the customer. | 9988998899 |
| `shipping_address` | YES | *string* | Primary address of the customer. | House no 123 |
| `shipping_address_2` | NO | *string* | Further address details of the customer. | Beside CM house |
| `shipping_city` | YES | *string* | Shipping city name. | Pune |
| `shipping_state` | YES | *string* | Shipping state name. | Maharashtra |
| `shipping_country` | YES | *string* | Shipping country name. | India |
| `shipping_pincode` | YES | *integer* | Shipping address pincode. | 120023 |
| `shipping_email` | NO | *string* | Customer's email address. | [john@doe.com](mailto:john@doe.com) |
| `billing_alternate_phone` | NO | *string* | The customer alternate phone. | 8604690454 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "order_id": "",
  "shipping_customer_name": "",
  "shipping_phone": "",
  "shipping_address": "",
  "shipping_address_2": "",
  "shipping_city": "",
  "shipping_state": "",
  "shipping_country": "",
  "shipping_pincode": "",
  "shipping_email": "",
  "billing_alternate_phone": ""
}
```

**Example responses**

#### Successful Call — HTTP 202 Accepted

Example request body:

```json
{
  "order_id": 16178831,
  "shipping_customer_name": "Majin Bu",
  "shipping_phone": "9988998899",
  "shipping_address": "Earth",
  "shipping_address_2": "",
  "shipping_city": "Pune",
  "shipping_state": "Maharashtra",
  "shipping_country": "India",
  "shipping_pincode": 110077
}
```

_(empty response body in source)_

#### Missing Fields — HTTP 422 Unprocessable Entity

Example request body:

```json
{
  "order_id": 16178831,
  "shipping_customer_name": "Majin Bu",
  "shipping_phone": "9988998899",
  "shipping_address": "Earth",
  "shipping_address_2": "",
  "shipping_city": "Pune",
  "shipping_state": "Maharashtra",
  "shipping_country": "",
  "shipping_pincode": 110077
}
```

```json
{
    "message": "Oops! Invalid Data.",
    "errors": {
        "shipping_country": [
            "The shipping country field is required."
        ]
    },
    "status_code": 422
}
```

#### Invalid Data — HTTP 422 Unprocessable Entity

Example request body:

```json
{
  "order_id": 123123,
  "shipping_customer_name": "Majin Bu",
  "shipping_phone": "9988998899",
  "shipping_address": "Earth",
  "shipping_address_2": "",
  "shipping_city": "Pune",
  "shipping_state": "Maharashtra",
  "shipping_country": "India",
  "shipping_pincode": 110077
}
```

```json
{
    "message": "Oops! Invalid Data.",
    "errors": {
        "order_id": [
            "The selected order id is invalid."
        ]
    },
    "status_code": 422
}
```

### Update Order

`POST https://apiv2.shiprocket.in/v1/external/orders/update/adhoc`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to update your orders. You have to pass all the required params at the minimum to create a quick custom order. You can add additional parameters as per your preference.

You can update only the order_items details before assigning the AWB (before Ready to Ship status). You can only update these key-value pairs i.e., increase/decrease the quantity, update tax/discount, add/remove product items. We've also enabled changing the nature of the order from a non-document to a document. You jus t need to pass the is_document key with the value of 1 in the payload.

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "order_id": "4TestOrderOct28",
  "order_date": "2024-10-28",
  "pickup_location": "23659_7026",
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
  "is_document":"0",
  "order_items": [
    {
      "name": "Agreement",
      "sku": "chakra123",
      "units": 1,
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

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
"success": true,
"partially_update": true,
"not_updated_fields": "order_date ,pickup_location ,channel_id ,comment...[TRUNCATED: 538 more chars omitted by the doc generator]",
"order_id": 79491,
"shipment_id": 77906,
"new_order_status": "NEW",
"old_order_status": 1,
"awb_code": "",
"courier_company_id": "",
"courier_name": ""
}
```

_1 long string value(s) (538 chars; base64 file data in the source) were truncated by the generator. The original values are in the source collection._

### Cancel an Order

`POST https://apiv2.shiprocket.in/v1/external/orders/cancel`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to cancel a created order. Multiple order_ids can be passed together as an array to cancel them simultaneously.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `ids` | YES | *integer* | The Shiprocket order id/ids of the orders that need to be canceled. | 16178831 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "ids": []
}
```

**Example responses**

#### Successful Call — HTTP 204 No Content

Example request body:

```json
{
  "ids": [16168898,16167171]
}
```

_(empty response body in source)_

#### Invalid Data — HTTP 500 Internal Server Error

Example request body:

```json
{
  "ids": [12312312]
}
```

```json
{
    "message": "Trying to get property of non-object",
    "status_code": 500
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

```json
{
    "message": "Required field missing",
    "errors": {
        "ids": [
            "The ids field is required."
        ]
    },
    "status_code": 422
}
```

### Add Inventory for Ordered Product

`PATCH https://apiv2.shiprocket.in/v1/external/orders/fulfill`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to add inventory for ordered products that are out of stock or low on quantity. You have to pass the order id and order product id. You can also specify the number of items.

**Notes:**

- Inventory sync of your account must be turned on to use this API.
- The order_id to be passed is the shiprocket order id. If you don't know the product id, Use the 'Get Product Details' API to get details about all the existing products.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *integer* | The Shiprocket order_id specified to the order. | 16167171 |
| `order_product_id` | YES | *integer* | The product id of the product to be added. | 17171717 |
| `quantity` | YES | *string* | The number of items you want to add. | 10 |
| `action` | YES | *string* | The action you want to carry out. Is 'add'. | add |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "data": [
    {
      "order_id": "",
      "order_product_id": "",
      "quantity": "",
      "action": ""
    }
  ]
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
  "data": [
    {
      "order_id": 14124005,
      "order_product_id": 43737767570843,
      "quantity": "1",
      "action": "add"
    }
  ]
}
```

```json
[
    {
        "data": {
            "order_id": 14124005,
            "order_product_id": 43737767570843,
            "quantity": "1",
            "action": "add"
        },
        "success": true,
        "message": "Inventory added successfully"
    }
]
```

#### Missing Fields — HTTP 200 OK

```json
[
    {
        "data": {
            "order_id": "",
            "order_product_id": "",
            "quantity": "",
            "action": ""
        },
        "success": false,
        "message": "Incorrect order_id or order status is no longer unfulfillable"
    }
]
```

#### Invalid Data — HTTP 200 OK

Example request body:

```json
{
  "data": [
    {
      "order_id": 10000001,
      "order_product_id": 17777771,
      "quantity": "1",
      "action": "add"
    }
  ]
}
```

```json
[
    {
        "data": {
            "order_id": 10000001,
            "order_product_id": 17777771,
            "quantity": "1",
            "action": "add"
        },
        "success": false,
        "message": "Incorrect order_id or order status is no longer unfulfillable"
    }
]
```

### Map Unmapped Products

`PATCH https://apiv2.shiprocket.in/v1/external/orders/mapping`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API maps your unmapped inventory products.

**Note:**

- Products must be unmapped to run this API successfully.
- Inventory sync must be turned on to use this API.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *integer* | The Shiprocket order_id specified to the order. | 16167171 |
| `order_product_id` | YES | *integer* | The product id of the product to be mapped. | 17171717 |
| `master_sku` | YES | *string* | The sku id of the product. In the case of a single integrated channel, master sku is the same as channel_sku; Otherwise, it can be found using the 'Get All Products' API. | chakra123 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "data": [
    {
      "order_id": "",
      "order_product_id": "",
      "master_sku": ""
    }
  ]
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
  "data": [
    {
      "order_id": 14303681,
      "order_product_id": 16487731,
      "master_sku": "delta123"
    }
  ]
}
```

```json
[
    {
        "data": {
            "order_id": 14303681,
            "order_product_id": 16487731,
            "master_sku": "delta123"
        },
        "status_code": 200,
        "success": true,
        "message": "Product mapped sucessfully."
    }
]
```

#### Missing Fields — HTTP 200 OK

```json
[
    {
        "data": {
            "order_id": "",
            "order_product_id": "",
            "master_sku": ""
        },
        "status_code": 0,
        "success": false,
        "message": "No product found matching this master sku"
    }
]
```

#### Invalid Data — HTTP 400 Bad Request

Example request body:

```json
{
  "data": [
    {
      "order_id": 11111111,
      "order_product_id": 17777771,
      "master_sku": "chakra123"
    }
  ]
}
```

```json
[
    {
        "data": {
            "order_id": 16178831,
            "order_product_id": 17484610,
            "master_sku": "chakra123"
        },
        "status_code": 0,
        "message": "Incorrect order_id or order status is  no longer unmapped",
        "success": false
    }
]
```

### Import Orders in Bulk

`POST https://apiv2.shiprocket.in/v1/external/orders/import`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to import orders in bulk to your Shiprocket account from an existing '.csv' file. The imported orders are automatically added to your panel.

**Request body** (`multipart/form-data`)

| Field | Type | Value |
|---|---|---|
| `file` | file | (file upload) |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "id": 19739203
}
```

#### Invalid Data — HTTP 400 Bad Request

```json
{
    "message": "Sorry, text/x-c file type is not allowed",
    "status_code": 400
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity

```json
{
    "message": "Oops! Something went wrong.",
    "errors": {
        "file": [
            "The file field is required."
        ]
    },
    "status_code": 422
}
```

## Orders

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
| `fbs_all_orders` | NO | *integer* | Use this filter if you want to view both SRF and Last mile (core) orders | 0 or 1 |

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

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value `Bearer {token}`.

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |
| `Authorization` | `Bearer {token}` |

**Description**

Get the order and shipment details of a particular order through this API by passing the Shiprocket order_id in the endpoint URL itself — type in your order_id in place of {id}.

No other body parameters are required.

**Note:**

For SRF orders, you'll receive an extra parameter viz., fulfillment_status. This key will have four values:

- Ready to Pack,
- Packed
- Added to Picklist
- Picked up

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/orders/show/16167171](https://apiv2.shiprocket.in/v1/external/orders/show/16167171) |

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/orders/show/16167171`

```json
{
  "data": {
    "id": 259492257,
    "channel_id": 38026,
    "channel_name": "MANUAL1",
    "base_channel_code": "CS",
    "is_international": 0,
    "is_document": 0,
    "channel_order_id": "1873081902",
    "customer_name": "DemoHome ",
    "customer_email": "abc@gmail.com",
    "customer_phone": "9876543236",
    "customer_address": "408, Gautami, Kondapur",
    "customer_address_2": null,
    "customer_city": "North West Delhi",
    "customer_state": "Delhi",
    "customer_pincode": "110088",
    "customer_country": "India",
    "pickup_code": "",
    "pickup_location": "",
    "pickup_location_id": "",
    "pickup_id": "",
    "ship_type": "",
    "courier_mode": "",
    "currency": "INR",
    "country_code": 99,
    "exchange_rate_usd": 0,
    "exchange_rate_inr": 0,
    "state_code": 1483,
    "payment_status": "",
    "delivery_code": "110088",
    "total": 345,
    "total_inr": 0,
    "total_usd": 0,
    "net_total": "345.00",
    "other_charges": "0.00",
    "other_discounts": "0.00",
    "giftwrap_charges": "0.00",
    "expedited": 0,
    "sla": "2 days",
    "cod": 0,
    "tax": 0,
    "total_kerala_cess": "",
    "discount": 0,
    "status": "RETURN PENDING",
    "sub_status": null,
    "status_code": 21,
    "master_status": "",
    "payment_method": "prepaid",
    "purpose_of_shipment": 0,
    "channel_created_at": "21 Sep 2022 05:25 PM",
    "created_at": "21 Sep 2022 05:28 PM",
    "order_date": "21 Sep 2022",
    "updated_at": "21 Sep 2022 05:28 PM",
    "products": [
      {
        "id": 365076966,
        "order_id": 259492257,
        "product_id": 1620533,
        "name": "watch",
        "sku": "Tshirt-Blue-41",
        "description": "WHEAT AND MESLIN DURUM WHEAT : OF SEED QUALITY",
        "channel_order_product_id": "365076966",
        "channel_sku": "Tshirt-Blue-41",
        "hsn": "",
        "model": null,
        "manufacturer": null,
        "brand": "",
        "color": "",
        "size": null,
        "custom_field": "",
        "custom_field_value": "",
        "custom_field_value_string": "",
        "weight": 0,
        "dimensions": "0x0x0",
        "price": 345,
        "cost": 345,
        "mrp": 400,
        "quantity": 1,
        "returnable_quantity": 0,
        "tax": 0,
        "status": 1,
        "net_total": 345,
        "discount": 0,
        "product_options": [],
        "selling_price": 345,
        "tax_percentage": 0,
        "discount_including_tax": 0,
        "channel_category": "Default Category",
        "packaging_material": "",
        "additional_material": "",
        "is_free_product": ""
      }
    ],
    "invoice_no": "",
    "shipments": {
      "id": 258878960,
      "order_id": 259492257,
      "order_product_id": null,
      "channel_id": 38026,
      "code": "",
      "cost": "0.00",
      "tax": "0.00",
      "awb": null,
      "rto_awb": "",
      "awb_assign_date": null,
      "etd": "",
      "delivered_date": "",
      "quantity": 1,
      "cod_charges": "0.00",
      "number": null,
      "name": null,
      "order_item_id": null,
      "weight": 1,
      "volumetric_weight": 0.266,
      "dimensions": "11.000x11.000x11.000",
      "comment": "",
      "courier": "",
      "courier_id": "",
      "manifest_id": "",
      "manifest_escalate": false,
      "status": "PENDING",
      "isd_code": "+91",
      "created_at": "21st Sep 2022 05:28 PM",
      "updated_at": "21st Sep 2022 05:28 PM",
      "pod": null,
      "eway_bill_number": "-",
      "eway_bill_date": null,
      "length": 11,
      "breadth": 11,
      "height": 11,
      "rto_initiated_date": "",
      "rto_delivered_date": "",
      "shipped_date": "",
      "package_images": "",
      "is_rto": false,
      "eway_required": false,
      "invoice_link": "",
      "is_darkstore_courier": 0,
      "courier_custom_rule": "",
      "is_single_shipment": true
    },
    "awb_data": {
      "awb": "",
      "applied_weight": "",
      "charged_weight": "",
      "billed_weight": "",
      "routing_code": "",
      "rto_routing_code": "",
      "charges": {
        "zone": "",
        "cod_charges": "",
        "applied_weight_amount": "",
        "freight_charges": "",
        "applied_weight": "",
        "charged_weight": "",
        "charged_weight_amount": "",
        "charged_weight_amount_rto": "",
        "applied_weight_amount_rto": "",
        "service_type_id": ""
      }
    },
    "order_insurance": {
      "insurance_status": "No",
      "policy_no": "N/A",
      "claim_enable": false
    },
    "return_pickup_data": {
      "id": 2143757,
      "name": "ashwin ashwin",
      "email": "ashwingunadeep@gmail.com",
      "address": "shiprocket",
      "address_2": "shiprocket",
      "city": "South West Delhi",
      "state": "Delhi",
      "country": "India",
      "pin_code": "110030",
      "phone": "9562817406",
      "lat": null,
      "long": null,
      "order_id": 259492257,
      "created_at": "2022-09-21 17:28:40",
      "updated_at": "2022-09-21 17:28:40"
    },
    "company_logo": null,
    "allow_return": 0,
    "is_return": 1,
    "is_incomplete": 0,
    "errors": null,
    "payment_code": null,
    "coupon_is_visible": false,
    "coupons": "",
    "billing_city": "",
    "billing_name": "",
    "billing_email": "",
    "billing_phone": "",
    "billing_alternate_phone": "",
    "billing_state_name": "",
    "billing_address": "",
    "billing_country_name": "",
    "billing_pincode": "",
    "billing_address_2": "",
    "billing_mobile_country_code": "+91",
    "isd_code": "",
    "billing_state_id": "",
    "billing_country_id": "",
    "freight_description": "Forward charges",
    "reseller_name": "",
    "shipping_is_billing": 0,
    "company_name": "shiprocket",
    "shipping_title": "",
    "allow_channel_order_sync": false,
    "uib-tooltip-text": "Re-fetch orders with updated details",
    "api_order_id": "",
    "allow_multiship": 0,
    "other_sub_orders": [],
    "others": {
      "weight": "1",
      "quantity": 1,
      "buyer_psid": null,
      "dimensions": "11x11x11",
      "api_order_id": "",
      "company_name": "shiprocket",
      "currency_code": "INR",
      "package_count": "1",
      "shipping_city": "North West Delhi",
      "shipping_name": "DemoHome ",
      "shipping_email": "abc@gmail.com",
      "shipping_phone": "9876543236",
      "shipping_state": "Delhi",
      "custom_order_id": null,
      "billing_isd_code": "+91",
      "forward_order_id": null,
      "shipping_address": "408, Gautami, Kondapur",
      "shipping_charges": "0",
      "shipping_country": "India",
      "shipping_pincode": "110088",
      "shipping_address_2": ""
    },
    "is_order_verified": 0,
    "extra_info": {
      "qc_check": 1,
      "qc_params": "Product Name,Size,Color,Brand,Product Image",
      "order_type": 1,
      "amazon_dg_status": false,
      "forward_order_id": "",
      "bluedart_dg_status": false,
      "other_courier_dg_status": false,
      "insurace_opted_at_order_creation": false
    },
    "dup": 0,
    "is_blackbox_seller": false,
    "shipping_method": "SR",
    "refund_detail": {
      "refund_mode": "Store Credits",
      "account_holder_name": "",
      "account_number": "",
      "bank_ifsc": "",
      "bank_name": ""
    },
    "pickup_address": [],
    "eway_bill_number": "",
    "eway_bill_url": "",
    "eway_required": false,
    "irn_no": "",
    "engage": null,
    "seller_can_edit": false,
    "seller_can_cancell": false,
    "is_post_ship_status": false,
    "order_tag": "",
    "qc_status": "",
    "qc_reason": "",
    "qc_image": "",
    "product_qc": [
      {
        "product_id": 365076966,
        "qc_values": {
          "qc_product_name": {
            "value": "watch",
            "name": "Product Name"
          },
          "qc_size": {
            "value": "asdasd",
            "name": "Size"
          },
          "qc_color": {
            "value": "asdasd",
            "name": "Color"
          },
          "qc_brand": {
            "value": "asdas",
            "name": "Brand"
          },
          "qc_product_image": {
            "value": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/1663238198WTVf4.jpeg",
            "name": "Product Image"
          }
        }
      }
    ],
    "seller_request": null,
    "change_payment_mode": true,
    "etd_date": null,
    "out_for_delivery_date": null,
    "delivered_date": null,
    "remittance_date": "",
    "remittance_utr": "",
    "remittance_status": "",
    "insurance_excluded": true,
    "can_edit_dimension": true
  }
}
```

#### Incorrect Endpoint Path — HTTP 404 Not Found

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

#### Invalid Data — HTTP 400 Bad Request

Example request: `GET https://apiv2.shiprocket.in/v1/external/orders/show/00000001`

```json
{
    "message": "Order ID not found",
    "status_code": 400
}
```

### Export your orders

`POST https://apiv2.shiprocket.in/v1/external/orders/export`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API downloads and creates a CSV file of all the orders in your Shiprocket account and sends the download URL to the linked email account of the API user.

The CSV file containing is accessible via this URL. No parameters are required to access this API.

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "status": 200,
    "is_background_downloading": 1
}
```

#### Invalid Endpoint — HTTP 502 Bad Gateway

Example request: `POST https://apiv2.shiprocket.in/v1/external/orders/expor`

```html
<html>
    <head>
        <title>502 Bad Gateway</title>
    </head>
    <body bgcolor="white">
        <center>
            <h1>502 Bad Gateway</h1>
        </center>
        <hr>
        <center>nginx</center>
    </body>
</html>
```

## Return & Exchange Orders

_No folder-level description in the source collection._

### Create a Return Order

`POST https://apiv2.shiprocket.in/v1/external/orders/create/return`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to create a new return order in your Shiprocket panel. Return orders are created in case the buyer refuses/rejects/returns a specific order.<br>The parameter specifications are the same as the custom order API, with a few exceptions.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *string* | The order id you want to specify to the order. Max char: 50. (Avoid passing character values as this contradicts some other API calls) | 99711997 |
| `order_date` | YES | *string* | The date of order creation in yyyy-mm-dd format. Time is an additional option. | 2019-08-05 |
| `channel_id` | NO | *integer* | Id of the desired channel where the order is to be placed. 'Custom' channel id is selected in case parameter is not filled. | 768903 |
| `pickup_customer_name` | YES | *string* | The customer’s first name. | John |
| `pickup_last_name` | NO | *string* | The customer’s last name. | Doe |
| `pickup_address` | YES | *string* | The customer's primary address. | 416, Udyog Vihar III, Sector 20 |
| `pickup_address_2` | NO | *string* | Additional customer address details. | DDA |
| `pickup_city` | YES | *string* | The customer's city name. | Delhi |
| `pickup_state` | YES | *string* | The customer's state. | New Delhi |
| `pickup_country` | YES | *string* | Customer's country name. | India |
| `pickup_pincode` | YES | *integer* | Pincode of the customer address. | 110002 |
| `pickup_email` | YES | *string* | Customer's email address. | [john@doe.com](https://mailto:john@doe.com) |
| `pickup_phone` | YES | *string* | Customer's phone number. | 9999999999 |
| `pickup_isd_code` | NO | *string* | ISD code. | 91 |
| `shipping_customer_name` | YES | *string* | The name of the seller the package is shipped back to. | Jane |
| `shipping_last_name` | NO | *string* | The last name of the seller. | Doe |
| `shipping_address` | YES | *string* | The address the package is shipped to. | Castle |
| `shipping_address_2` | NO | *string* | Further shipping address details. | Bridge |
| `shipping_city` | YES | *string* | The shipping address city. | Mumbai |
| `shipping_country` | YES | *string* | The shipping address country. | India |
| `shipping_pincode` | YES | *integer* | The shipping pincode. | 220022 |
| `shipping_state` | YES | *string* | Shipping address state. | Maharashtra |
| `shipping_email` | NO | *string* | The email of the seller the package is shipped to. | [jane@doe.com](https://mailto:jane@doe.com) |
| `shipping_isd_code` | NO | *string* | The shipping isd code. | 91 |
| `shipping_phone` | YES | *integer* | Phone no. of the shipping customer | 8888888888 |
| `order_items` | YES | / | Array containing further fields. | / |
| `name` | YES | *string* | Name of the product. | ball123 |
| `sku` | YES | *string* | The sku id of the product. | Tennis Ball |
| `units` | YES | *integer* | No of units that are to be shipped. | 1 |
| `selling_price` | YES | *integer* | The selling price per unit in Rupee. Inclusive of GST. | 10 |
| `discount` | NO | *integer* | The discount amount in Rupee. Inclusive of tax. | 0 |
| `hsn` | NO | *string* | Harmonised System Nomenclature code. Used to determine the category of taxation the goods fall under. | 4412 |
| `return_reason` | NO | *string* | Bought by Mistake, Both product and shipping box damaged | Please refer to the "Possible return_reason values" section |
| `qc_enable` | CONDITIONAL YES | *string* | If True, QC will be performed for that product and QC will be performed only for a single SKU per order | true/false |
| `qc_color` | NO | *varchar(180)* | The color of the product can be passed in this parameter | Red |
| `qc_brand` | NO | *varchar(255)* | The brand of the product can be passed in this parameter | 768903 |
| `qc_serial_no` | NO | *varchar(255)* | The serial number of the product can be passed in this parameter | T13123124 |
| `qc_ean_barcode` | NO | *varchar(255)* | EAN/Barcode of the product can be passed in this parameter | QWRE123 |
| `qc_size` | NO | *varchar(180)* | The size of the product can be passed in this parameter | 8 |
| `qc_product_name` | CONDITIONAL YES | *varchar(255)* | If qc_enable set True, then Product name should be passed in this parameter | Shoes |
| `qc_product_image` | CONDITIONAL YES | *varchar(255)* | If qc_enable set True, then Product image should be passed in this parameter (only png/jpg format supported) | [https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/1636713733zxja.png](https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/1636713733zxja.png) |
| `qc_product_imei` | NO | *varchar(255)* | IMEI of the device | 398612387501872509 |
| `qc_brand_tag` | NO | *boolean* | Eligible Categories : Selective like Footwear, Apparels The pickup agent will cross-check the provided brand name, which should match the brand tag affixed to the item(s) upon delivery. Can be either 0 or 1 | 1 |
| `qc_used_check` | NO | *boolean* | Eligible Product Categories : All The pickup agent will check the product being handed over for clear signs of usage. Can be either 0 or 1 | 0 |
| `qc_sealtag_check` | NO | *boolean* | Eligible Product Categories : All The pickup agent will check if the seal tag is intact in the products received from the buyer. Can be either 0 or 1 | 1 |
| `qc_check_damaged_product` | NO | *string* | Eligible Product Categories : All The pickup agent will check the product for any signs of damages. Can be either yes or no | yes |
| `payment_method` | YES | *string* | The method of payment. This should always be prepaid. | Prepaid |
| `total_discount` | NO | *string* | The total discount amount in Rupee. | 0 |
| `sub_total` | YES | *integer* | Calculated sub total amount in Rupee after deductions. | 10 |
| `length` | YES | *float* | The length of the shipment in cms. | 10 |
| `breadth` | YES | *float* | The breadth of the shipment in cms. | 15 |
| `height` | YES | *float* | The height of the shipment in cms. | 20 |
| `weight` | YES | *float* | The shipment weight in kgs. | 1 |

#### **Possible return_reason values**:

```
1. Bought by Mistake
2. Better price available
3. Performance or quality not adequate
4. Incompatible or not useful
5. Product damaged, but shipping box OK
6. Item arrived too late
7. Missing parts or accessories
8. Both product and shipping box damaged
9. Wrong item was sent
10. Item defective or doesn't work
11. No longer needed
12. Didn't approve purchase
13. Inaccurate website description
14. Return against replacement
15. Delay Refund
16. Delivered Late
17. Product does not Match Description on Website
18. Both Product & Outer Box Damaged
19. Defective or does not work
20. Product damaged, but outer Box OK
21. Missing Parts or Accessories
22. Incorrect Item Delivered
23. Product Defective or Doesn't Work
24. Product performance/quality is not up to my expectations
25. Other
26. Changed my mind
27. Does not fit
28. Size not as expected
29. Item is damaged
30. Received wrong item
31. Parcel damaged on arrival
32. Quality not as expected
33. Missing Item or accessories
34. Performance not adequate
35. Not as described
36. Arrived too late
37. Order Not Received
38. Empty Package
39. Wrong item or Wrong colour was sent
40. Item defective, expired, spoilt or does not work
41. Items or parts missing
42. Size or Quantity issues
43. Status as delivered but order not received
44. N/A
```

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "order_id": "r121579B09ap3o",
  "order_date": "2021-12-30",
  "channel_id": "27202",
  "pickup_customer_name": "iron man",
  "pickup_last_name": "",
  "company_name":"iorn pvt ltd",
  "pickup_address": "b 123",
  "pickup_address_2": "",
  "pickup_city": "Delhi",
  "pickup_state": "New Delhi",
  "pickup_country": "India",
  "pickup_pincode": 110030,
  "pickup_email": "deadpool@red.com",
  "pickup_phone": "9810363552",
  "pickup_isd_code": "91",
  "shipping_customer_name": "Jax",
  "shipping_last_name": "Doe",
  "shipping_address": "Castle",
  "shipping_address_2": "Bridge",
  "shipping_city": "ghaziabad",
  "shipping_country": "India",
  "shipping_pincode": 201005,
  "shipping_state": "Uttarpardesh",
  "shipping_email": "kumar.abhishek@shiprocket.com",
  "shipping_isd_code": "91",
  "shipping_phone": 8888888888,
  "order_items": [
    {
      "sku": "WSH234",
      "name": "shoes",
      "units": 2,
      "selling_price": 100,
      "discount": 0,
      "qc_enable":true,
      "hsn": "123",
      "brand":"",
      "qc_size":"43"
       }
    ],
  "payment_method": "PREPAID",
  "total_discount": "0",
  "sub_total": 400,
  "length": 11,
  "breadth": 11,
  "height": 11,
  "weight": 0.5
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
  "order_id": "r121579B09ap3o",
  "order_date": "2021-12-30",
  "channel_id": "27202",
  "pickup_customer_name": "iron man",
  "pickup_last_name": "",
  "company_name":"iorn pvt ltd",
  "pickup_address": "b 123",
  "pickup_address_2": "",
  "pickup_city": "Delhi",
  "pickup_state": "New Delhi",
  "pickup_country": "India",
  "pickup_pincode": 110030,
  "pickup_email": "deadpool@red.com",
  "pickup_phone": "9810363552",
  "pickup_isd_code": "91",
  "shipping_customer_name": "Jax",
  "shipping_last_name": "Doe",
  "shipping_address": "Castle",
  "shipping_address_2": "Bridge",
  "shipping_city": "ghaziabad",
  "shipping_country": "India",
  "shipping_pincode": 201005,
  "shipping_state": "Uttarpardesh",
  "shipping_email": "kumar.abhishek@shiprocket.com",
  "shipping_isd_code": "91",
  "shipping_phone": 8888888888,
  "order_items": [
    {
      "name": "shoes",
      "qc_enable":true,
      "qc_product_name": "shoes",
      "sku": "WSH234",
      "units": 1,
      "selling_price": 100,
      "discount": 0,
      "qc_brand":"Levi",
      "qc_product_image":"https://assets.vogue.in/photos/5d7224d50ce95e0008696c55/2:3/w_2240,c_limit/Joker.jpg"
       }
    ],
  "payment_method": "PREPAID",
  "total_discount": "0",
  "sub_total": 400,
  "length": 11,
  "breadth": 11,
  "height": 11,
  "weight": 0.5
}
```

```json
{
    "order_id": 170872392,
    "shipment_id": 170411259,
    "status": "RETURN PENDING",
    "status_code": 21,
    "company_name": "shiprocket"
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity

Example request body:

```json
{
  "order_id": "997119978",
  "order_date": "2019-08-05",
  "channel_id": "76893",
  "pickup_customer_name": "Deadpool",
  "pickup_last_name": "",
  "pickup_address": "Home",
  "pickup_address_2": "",
  "pickup_city": "Delhi",
  "pickup_state": "New Delhi",
  "pickup_pincode": "110002",
  "pickup_email": "deadpool@red.com",
  "pickup_phone": "9999999999",
  "pickup_isd_code": "",
  "pickup_location_id": "41514",
  "shipping_customer_name": "Jax",
  "shipping_last_name": "",
  "shipping_address": "Castle",
  "shipping_address_2": "",
  "shipping_city": "Mumbai",
  "shipping_country": "India",
  "shipping_pincode": "220022",
  "shipping_state": "Maharashtra",
  "shipping_email": "jax@tank.com",
  "shipping_isd_code": "",
  "shipping_phone": "8888888888",
  "order_items": [
    {
      "sku": "ball123",
      "name": "Tennis Ball",
      "units": 1,
      "selling_price": 10,
      "discount": "",
      "hsn": ""
    }
  ],
  "payment_method": "Prepaid",
  "total_discount": "",
  "sub_total": 10,
  "length": 10,
  "breadth": 15,
  "height": 20,
  "weight": 1
}
```

```json
{
    "message": "Oops! Invalid Data.",
    "errors": {
        "pickup_country": [
            "The pickup country field is required"
        ],
        "order_id": [
            "The order id has already been taken."
        ]
    },
    "status_code": 422
}
```

#### Wrong Format — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

Example request body:

```text
{
  "order_id": "997119978",
  "order_date": "2019-08-05",
  "channel_id": "76893",
  "pickup_customer_name": "Deadpool",
  "pickup_last_name": "",
  "pickup_address": "Home",
  "pickup_address_2": "",
  "pickup_city": "Delhi",
  "pickup_state": "New Delhi",
  "pickup_country": "India"
  "pickup_pincode": "110002",
  "pickup_email": "deadpool@red.com",
  "pickup_phone": "9999999999",
  "pickup_isd_code": "",
  "pickup_location_id": "41514",
  "shipping_customer_name": "Jax",
  "shipping_last_name": "",
  "shipping_address": "Castle",
  "shipping_address_2": "",
  "shipping_city": "Mumbai",
  "shipping_country": "India",
  "shipping_pincode": "220022",
  "shipping_state": "Maharashtra",
  "shipping_email": "jax@tank.com",
  "shipping_isd_code": "",
  "shipping_phone": "8888888888",
  "order_items": [
    {
      "sku": "ball123",
      "name": "Tennis Ball",
      "units": 1,
      "selling_price": 10,
      "discount": "",
      "hsn": ""
    }
  ],
  "payment_method": "Prepaid",
  "total_discount": "",
  "sub_total": 10,
  "length": 10,
  "breadth": 15,
  "height": 20,
  "weight": 1
}
```

```json
{
    "message": "Oops! Invalid Data.",
    "errors": {
        "order_date": [
            "The order date field is required."
        ],
        "payment_method": [
            "The payment method field is required."
        ],
        "pickup_customer_name": [
            "The pickup customer name field is required"
        ],
        "pickup_email": [
            "The pickup email field is required."
        ],
        "pickup_address": [
            "The pickup address field is required"
        ],
        "pickup_city": [
            "The pickup city field is required"
        ],
        "pickup_state": [
            "The pickup state field is required"
        ],
        "pickup_country": [
            "The pickup country field is required"
        ],
        "pickup_pincode": [
            "The pickup pincode field is required"
        ],
        "pickup_phone": [
            "The pickup phone field is required."
        ],
        "pickup_location_id": [
            "The pickup location id field is required."
        ],
        "shipping_customer_name": [
            "The shipping customer name field is required"
        ],
        "shipping_email": [
            "The shipping email field is required."
        ],
        "shipping_address": [
            "The shipping address field is required"
        ],
        "shipping_city": [
            "The shipping city field is required"
        ],
        "shipping_state": [
            "The shipping state field is required"
        ],
        "shipping_country": [
            "The shipping country field is required"
        ],
        "shipping_pincode": [
            "The shipping pincode field is required"
        ],
        "shipping_phone": [
            "The shipping phone field is required."
        ],
        "order_items": [
            "The order items field is required."
        ],
        "sub_total": [
            "The sub total field is required."
        ],
        "order_id": [
            "The order id field is required."
        ]
    },
    "status_code": 422
}
```

#### Invalid Data — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

Example request body:

```json
{
  "order_id": "997119978",
  "order_date": "2019-08-05",
  "channel_id": "76893",
  "pickup_customer_name": "Deadpool",
  "pickup_last_name": "",
  "pickup_address": "Home",
  "pickup_address_2": "",
  "pickup_city": "Delhi",
  "pickup_state": "New Delhi",
  "pickup_country": "India",
  "pickup_pincode": "110002",
  "pickup_email": "deadpool@red.com",
  "pickup_phone": "9999999999",
  "pickup_isd_code": "",
  "pickup_location_id": "41514",
  "shipping_customer_name": "Jax",
  "shipping_last_name": "",
  "shipping_address": "Castle",
  "shipping_address_2": "",
  "shipping_city": "Mumbai",
  "shipping_country": "India",
  "shipping_pincode": "220022",
  "shipping_state": "Maharashtra",
  "shipping_email": "jax@tank.com",
  "shipping_isd_code": "",
  "shipping_phone": "8888888888",
  "order_items": [
    {
      "sku": "ball123",
      "name": "Tennis Ball",
      "units": 1,
      "selling_price": 10,
      "discount": "",
      "hsn": ""
    }
  ],
  "payment_method": "Prepaid",
  "total_discount": "",
  "sub_total": 10,
  "length": 10,
  "breadth": 15,
  "height": 20,
  "weight": 1
}
```

```json
{
    "message": "Oops! Invalid Data.",
    "errors": {
        "order_id": [
            "The order id has already been taken."
        ]
    },
    "status_code": 422
}
```

### Create Exchange Order

`POST https://apiv2.shiprocket.in/v1/external/orders/create/exchange`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value `Bearer {{token}}`.

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |
| `Authorization` | `Bearer {{token}}` |

**Description**

Use this API to create a new exchange order in your Shiprocket panel. Exchange orders are created in case the buyer needs to replace an item due to reasons such as size mismatch, wrong product received, or defective product. This API helps sellers seamlessly process exchange requests while ensuring the correct quality checks and logistics handling

#### Parameters:

| PARAMS | REQUIRED | DATA TYPE | DESCRIPTION | EXAMPLE |
|---|---|---|---|---|
| `exchange_order_id` | YES | *String* | Unique ID for the exchange order | EX_TEST002 |
| `seller_pickup_location_id` | YES | *String* | Pickup location ID of the seller | 5723898 |
| `seller_shipping_location_id` | YES | *String* | Shipping location ID of the seller | 5723898 |
| `return_order_id` | YES | *String* | Unique ID for the return order | R_TEST002 |
| `order_date` | YES | *Date* | Date of order placement (YYYY-MM-DD) | 2024-12-10 |
| `payment_method` | YES | *String* | Payment method used for the order | prepaid |
| `buyer_shipping_first_name` | YES | *String* | First name of the shipping buyer | Test |
| `buyer_shipping_last_name` | NO | *String* | Last name of the shipping buyer | Test |
| `buyer_shipping_email` | NO | *String* | Email of the shipping buyer (valid email format) | [test@gmail.com](https://null) |
| `buyer_shipping_address` | YES | *String* | Full address of the shipping buyer | dkalsd |
| `buyer_shipping_address_2` | NO | *String* | Additional address details |  |
| `buyer_shipping_city` | YES | *String* | City of the shipping buyer | South West Delhi |
| `buyer_shipping_state` | YES | *String* | State of the shipping buyer | Delhi |
| `buyer_shipping_country` | YES | *String* | Country of the shipping buyer | India |
| `buyer_shipping_pincode` | YES | *String* | Pincode of the shipping address | 110045 |
| `buyer_shipping_phone` | YES | *String* | Contact number of the shipping buyer (10 digits) | 9716414139 |
| `buyer_pickup_first_name` | YES | *String* | First name of the pickup buyer | Test |
| `buyer_pickup_last_name` | NO | *String* | Last name of the pickup buyer | Test |
| `buyer_pickup_email` | NO | *String* | Email of the pickup buyer | [test@gmail.com](https://null) |
| `buyer_pickup_address` | YES | *String* | Full address of the pickup buyer | Test |
| `buyer_pickup_address_2` | NO | *String* | Additional pickup address details |  |
| `buyer_pickup_city` | YES | *String* | City of the pickup buyer | South West Delhi |
| `buyer_pickup_state` | YES | *String* | State of the pickup buyer | Delhi |
| `buyer_pickup_country` | YES | *String* | Country of the pickup buyer | India |
| `buyer_pickup_pincode` | YES | *String* | Pincode of the pickup address | 110045 |
| `buyer_pickup_phone` | YES | *String* | Contact number of the pickup buyer (10 digits) | 9716414139 |
| `order_items` | YES | *Array* | List of items in the order |  |
| `order_items[].name` | YES | *String* | Name of the product | Black tshirt XL |
| `order_items[].selling_price` | YES | *Float* | Price of the product | 500.00 |
| `order_items[].units` | YES | *Integer* | Quantity of the product | 1 |
| `order_items[].hsn` | YES | *String* | HSN code of the product | 1733808730720 |
| `order_items[].sku` | YES | *String* | SKU of the product | mackbook |
| `order_items[].tax` | NO | *Float* | Tax amount |  |
| `order_items[].discount` | NO | *Float* | Discount on the product |  |
| `order_items[].exchange_item_id` | NO | *String* | Exchange item ID | 193658024 |
| `order_items[].exchange_item_name` | YES | *String* | Exchange item name | Black tshirt XL |
| `order_items[].exchange_item_sku` | YES | *String* | Exchange item SKU | mackbook |
| `sub_total` | YES | *Float* | Subtotal amount | 500.00 |
| `shipping_charges` | NO | *Float* | Shipping charges |  |
| `giftwrap_charges` | NO | *Float* | Gift wrapping charges |  |
| `total_discount` | NO | *Float* | Total discount on the order | 0 |
| `transaction_charges` | NO | *Float* | Transaction charges |  |
| `return_length` | YES | *Float* | Return package length (cm) | 10.00 |
| `return_breadth` | YES | *Float* | Return package breadth (cm) | 10.00 |
| `return_height` | YES | *Float* | Return package height (cm) | 10.00 |
| `return_weight` | YES | *Float* | Return package weight (kg) | 0.500 |
| `exchange_length` | YES | *Float* | Exchange package length (cm) | 11.00 |
| `exchange_breadth` | YES | *Float* | Exchange package breadth (cm) | 11.00 |
| `exchange_height` | YES | *Float* | Exchange package height (cm) | 11.00 |
| `exchange_weight` | YES | *Float* | Exchange package weight (kg) | 11.00 |
| `return_reason` | YES | *String* | Reason for return | 29 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "order_items": [
    {
      "name": "Black tshirt XL",
      "selling_price": "500.00",
      "units": "1",
      "hsn": "1733808730720",
      "sku": "mackbook",
      "tax": "",
      "discount": "",
      "brand": "",
      "color": "",
      "exchange_item_id": "193658024",
      "exchange_item_name": "Black tshirt XL",
      "exchange_item_sku": "mackbook",
      "qc_enable": true,
      "qc_product_name": "Black tshirt XL",
      "qc_product_image": "https://sr-multichannel-stage.s3.ap-south-1.amazonaws.com/1310/qc_product_img/547950c2-9c2f-4908-98d5-276f9ad5b63a.png",
      "qc_brand": "changedname1",
      "qc_color": "changecolr",
      "qc_size": "changesize112",
      "accessories": "",
      "qc_used_check": "1",
      "qc_sealtag_check": "1",
      "qc_brand_box": "1",
      "qc_check_damaged_product": "yes"
    }
  ],
  "buyer_pickup_first_name": "Test",
  "buyer_pickup_last_name": "Test",
  "buyer_pickup_email": "test@gmail.com",
  "buyer_pickup_address": "Test",
  "buyer_pickup_address_2": "",
  "buyer_pickup_city": "South West Delhi",
  "buyer_pickup_state": "Delhi",
  "buyer_pickup_country": "India",
  "buyer_pickup_phone": "9716414139",
  "buyer_pickup_pincode": "110045",
  "buyer_shipping_first_name": "Test",
  "buyer_shipping_last_name": "Test",
  "buyer_shipping_email": "test@gmail.com",
  "buyer_shipping_address": "dkalsd",
  "buyer_shipping_address_2": "",
  "buyer_shipping_city": "South West Delhi",
  "buyer_shipping_state": "Delhi",
  "buyer_shipping_country": "India",
  "buyer_shipping_phone": "9716414139",
  "buyer_shipping_pincode": "110045",
  "seller_pickup_location_id": "5723898",
  "seller_shipping_location_id": "5723898",
  "exchange_order_id": "EX_TEST002",
  "return_order_id": "R_TEST002",
  "payment_method": "prepaid",
  "order_date": "2024-12-10",
  "channel_id": "1960878",
  "existing_order_id": "",
  "return_reason": "29",
  "sub_total": "500.00",
  "shipping_charges": "",
  "giftwrap_charges": "",
  "total_discount": "0",
  "transaction_charges": "",
  "exchange_length": "11",
  "exchange_breadth": "11",
  "exchange_height": "11",
  "exchange_weight": "11",
  "return_length": "10.00",
  "return_breadth": "10.00",
  "return_height": "10.00",
  "return_weight": "0.500",
  "qc_check": "true"
}
```

**Example responses**

#### Sucessful Call — (no HTTP status recorded in source)

```json
{
    "success": true,
    "data": {
        "forward_orders": {
            "order_id": 76175,
            "channel_order_id": "EX_TEST_101",
            "shipment_id": 659559333,
            "status": "NEW",
            "status_code": 1,
            "awb_code": "",
            "courier_company_id": "",
            "courier_name": ""
        },
        "return_orders": {
            "order_id": 76176,
            "channel_order_id": "R_TEST_101",
            "shipment_id": 659559334,
            "status": "RETURN PENDING",
            "status_code": 21,
            "awb_code": "",
            "courier_company_id": "",
            "courier_name": ""
        }
    }
}
```

### Update Return Order

`POST https://apiv2.shiprocket.in/v1/external/orders/edit`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to update your return orders. Please specify the parameters based on the "action" key.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *string* | Your return order ID | R_1231234 |
| `action` | YES | *string* | Pass array of action. Allowed actions: `1. product_details: Allows you to edit the weight and dimensions` `2. warehouse_address: Allows you to change the return address` | "action": ["product_details"] |
| `length` | CONDITIONAL YES | *float* | The length of the item in cms. Must be more than 0.5 | 12 |
| `breadth` | CONDITIONAL YES | *float* | The breadth of the item in cms. Must be more than 0.5 | 23 |
| `height` | CONDITIONAL YES | *float* | The height of the item in cms. Must be more than 0.5 | 30 |
| `weight` | CONDITIONAL YES | *float* | The weight of the item in kgs. Must be more than 0. | 10 |
| `return_warehouse_id` | CONDITIONAL YES | *integer* | Id of pickup location | 213443 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "order_id": "79596",
  "action": ["product_details"],
  "length":"11",
  "breadth":"10",
  "height":"10",
  "return_warehouse_id":1072,
  "weight":1.5
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "product_details": {
        "success": true,
        "msg": "Product Details is updated successfully"
    },
    "return_warehouse_address": {
        "success": true,
        "msg": "Shipping Address is updated successfully"
    }
}
```

### Get All Return Orders

`GET https://apiv2.shiprocket.in/v1/external/orders/processing/return`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Using this API, you can get a list of all created return orders in your Shiprocket account, along with their details.

No parameters are required to use the API. However, further parameters can be defined to sort and filter the data.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `page` | NO | *integer* | The page number to display. | 1 |
| `per_page` | NO | *integer* | The number of orders per page. | 2 |
| `to` | NO | *string* | Ending date of search. | 2019-08-04 |
| `from` | NO | *string* | Starting date of search. | 2019-08-05 |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": [
        {
            "id": 16525924,
            "channel_id": 76893,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_order_id": "997119978",
            "customer_name": "Jax Doe",
            "customer_email": "jax@tank.com",
            "customer_phone": "8888888888",
            "customer_pincode": "220022",
            "pickup_code": "110002",
            "pickup_location": "Home ,Home ,Delhi ,New Delhi ,India",
            "payment_status": "",
            "total": "10.00",
            "expedited": 0,
            "sla": "2 days",
            "shipping_method": "SR",
            "status": "RETURN PENDING",
            "status_code": 21,
            "payment_method": "prepaid",
            "is_international": 0,
            "purpose_of_shipment": 0,
            "channel_created_at": "5 Aug 2019, 12:00 AM",
            "created_at": "5 Aug 2019, 03:53 PM",
            "products": [
                {
                    "id": 19192381,
                    "name": "Tennis Ball",
                    "channel_sku": "ball123",
                    "channel_order_product_id": "19192381",
                    "quantity": 1,
                    "product_id": 17949825,
                    "sku": "ball123",
                    "custom_field": "",
                    "custom_field_value": "",
                    "status": "UNDEFINED",
                    "hsn": "4412"
                }
            ],
            "delivery_code": "220022",
            "cod": 0,
            "shipment_id": 16370752,
            "in_queue": 0,
            "shipments": [
                {
                    "isd_code": "+91",
                    "courier": "",
                    "sr_courier_id": "",
                    "weight": "1",
                    "length": "10",
                    "breadth": "15",
                    "height": "20",
                    "volumetric_weight": 0.6,
                    "awb": ""
                }
            ]
        }
    ],
    "meta": {
        "pagination": {
            "total": 16,
            "count": 15,
            "per_page": 1,
            "current_page": 1,
            "total_pages": 2,
            "links": {
                "next": "https://apiv2.shiprocket.in/v1/external/orders/processing/return?page=2"
            }
        }
    }
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/orders/processing/retur`

```json
{
    "message": "404 Not Found",
    "status_code": 404
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
| `pickup_postcode` | YES | *integer* | The shipment id of the order you want to create the AWB of. | 16016920 |
| `delivery_postcode` | YES | *integer* | The courier id of the courier service you want to select. The default courier is selected in case no id is specified. | 10 |
| `order_id` | CONDITIONAL YES | *integer* | Use this to change the courier of a shipment. Value: reassign Note that this can be done only once in 24 hours. | reassign |
| `cod` | CONDITIONAL YES | *boolean* | 1 for Cash on Delivery and 0 for Prepaid orders. | 1 |
| `weight` | CONDITIONAL YES | *string* | The weight of shipment in kgs. | 2 |
| `length` | NO | *integer* | The length of the shipment in cms. | 15 |
| `breadth` | NO | *integer* | The breadth of the shipment in cms. | 10 |
| `height` | NO | *integer* | The height of the shipment in cms. | 5 |
| `declared_value` | NO | *integer* | The price of the order shipment rupee. | 50 |
| `mode` | NO | *string* | The mode of travel. Either: Surface or Air | Air |
| `is_return` | YES | *integer* | Whether the order is return order or not. 1 in case of Yes and 0 for No | 1 |
| `qc_check` | NO | *integer* | Filter out QC couriers from the serviceability list | 1 |

**Example responses:** none published in the source collection for this request.

### Generate AWB for Return Shipment

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
| `shipment_id` | YES | *integer* | The shipment id of the order you want to create the AWB of. | 16016920 |
| `courier_id` | NO | *integer* | The courier id of the courier service you want to select. The default courier is selected in case no id is specified. | 10 |
| `status` | NO | *string* | Use this to change the courier of a shipment. Value: reassign Note that this can be done only once in 24 hours. | reassign |
| `is_return` | YES | *integer* | Whether the order is return order or not. 1 in case of Yes and 0 for No. | 1 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "shipment_id": "",
  "courier_id": "",
  "status": "",
  "is_return": ""
}
```

**Example responses:** none published in the source collection for this request.

## Wrapper API

This is an all-in-one API to create an order, ship the order, add a new pickup location and generate a label along with the manifest for the same.

### Forward

`POST https://apiv2.shiprocket.in/v1/external/shipments/create/forward-shipment`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to do multiple tasks in one go, namely creating a quick order, requesting its shipment, and finally generating the label and the manifest for the same order.

This API integrates several other APIs to perform all these tasks together.

**Notes:**

- Use the 'vendor_details' array to add a new pickup location to your account and assign it to your order.
- The 'pickup_location' field must contain a new pickup location name for adding a new pickup location to your Shiprocket account.
- In case of multiple items per order, please pass the final weight (sum total weight of items) of the shipment.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `mode` | NO | *string* | The mode of shipment, either surface or air. *Value*: **Surface** or **Air** | Air |
| `request_pickup` | NO | *boolean* | Use false if you dont want to request pickup. Default value is true. | true |
| `print_label` | NO | *boolean* | Use false if you dont want to print label. Default value is true. | true |
| `generate_manifest` | NO | *boolean* | Use false if you dont want to generate manifest. Default value is true. | true |
| `courier_id` | NO | *integer* | The courier id of the courier you want to assign. Refer to the servicability API to get id. | 10 |
| `reseller_name` | NO | *string* | The 'from' name if you want to print. Use 'Reseller: [name]' | Reseller: Divine |
| `order_id` | YES | *string* | The custom reference id you want to assign to the order. | 2477 |
| `isd_code` | NO | *string* | The isd code | 91 |
| `billing_isd_code` | NO | *string* | The billing isd code. | 91 |
| `order_date` | YES | *string* | The date of the creation of order. | 2019-08-04 |
| `channel_id` | NO | *string* | The channel id of the specific channel. Use the Channels API to get id. | 72505 |
| `company_name` | NO | *string* | Name of the company. | Amazon |
| `billing_customer_name` | YES | *string* | The first name of customer to be billed. | John |
| `billing_last_name` | NO | *string* | The last name of the billing customer. | Doe |
| `billing_address` | YES | *string* | The primary billing address. | House no 21 |
| `billing_address_2` | NO | *string* | Additional billing address details. | Street 2, Dwarka |
| `billing_city` | YES | *string* | The billing city. | Delhi |
| `billing_state` | YES | *string* | The billing address state. | New Delhi |
| `billing_country` | YES | *string* | The billing address country. | India |
| `billing_pincode` | YES | *integer* | The pincode of the billing address. | 110002 |
| `billing_email` | YES | *string* | The billing customer email. | [john@doe.com](https://mailto:john@doe.com) |
| `billing_phone` | YES | *integer* | The billing customer phone. | 9999998899 |
| `billing_alternate_phone` | NO | *integer* | The billing customer alternate phone. | 8404690454 |
| `shipping_is_billing` | YES | *boolean* | Whether shipping details are the same as billing details. **true** for yes **false** for no. | true |
| `shipping_customer_name` | CONDITIONAL YES | *string* | Shipping customer's first name. | Jane |
| `shipping_last_name` | NO | *string* | Shipping customer's last name. | Doe |
| `shipping_address` | CONDITIONAL YES | *string* | The shipping address. | House no X |
| `shipping_address_2` | NO | *string* | Additional shipping address details. | Street X |
| `shipping_city` | CONDITIONAL YES | *string* | The shipping city. | Mumbai |
| `shipping_state` | CONDITIONAL YES | *string* | The state of the shipping address. | Maharashtra |
| `shipping_country` | CONDITIONAL YES | *string* | The shipping address country. | India |
| `shipping_pincode` | CONDITIONAL YES | *integer* | Shipping pincode. | 230023 |
| `shipping_email` | CONDITIONAL YES | *string* | The email of the shipping customer. | [Jane@Doe.com](https://mailto:Jane@Doe.com) |
| `shipping_phone` | CONDITIONAL YES | *integer* | The phone number of the shipping customer. | 8877997799 |
| `order_items` | YES | / | Array containing further parameters. | / |
| `name` | YES | *string* | The name of the product. | Jeans |
| `sku` | YES | *string* | sku Code of the product. | Bat |
| `units` | YES | *integer* | Number of units. | 10 |
| `hsn` | NO | *integer* | HSN code if available. | 4412 |
| `selling_price` | YES | *integer* | The selling price of each unit inclusive of GST. | 200 |
| `tax` | NO | *integer* | The tax applied in percent. | 20 |
| `discount` | NO | *integer* | The discount amount inclusive of tax. | 20 |
| `payment_method` | YES | *string* | If the payment method is Cash on delivery (**COD**) or **Prepaid**. | COD |
| `shipping_charges` | NO | *integer* | The shipping charges if any in rupees. | 5 |
| `giftwrap_charges` | NO | *integer* | The gift-wrap charges if any in rupees. | 5 |
| `transaction_charges` | NO | *integer* | The transaction charges if any in rupees. | 10 |
| `total_discount` | NO | *integer* | The discount amount in rupees. | 15 |
| `sub_total` | YES | *integer* | The sub total amount in rupees. | 1800 |
| `weight` | YES | *integer* | The weight of the shipment in kgs. | 2 |
| `length` | YES | *integer* | The length of the shipment in cms. Must be more than 0.5 | 10 |
| `breadth` | YES | *integer* | The breadth of the shipment in cms. Must be more than 0.5 | 15 |
| `height` | YES | *integer* | The height of the shipment in cms. Must be more than 0.5 | 20 |
| `pickup_location` | YES | *string* | The pickup location name. Equal to an existing pickup location. If you use 'vendor details' to add a new location, it must be equal to the new pickup location name. | Office |
| `customer_gstin` | NO | *string* | Goods and Services Tax Identification Number. | 29ABCDE1234F2Z5 |
| `vendor_details` | NO | / | Array containing further parameters. Use to assign/add a new pickup location to your account. | / |
| `email` | CONDITIONAL YES | *string* | The shipper's email address. | [john@doe.com](https://mailto:john@doe.com) |
| `phone` | CONDITIONAL YES | *integer* | The shipper's phone number. | 8888999888 |
| `name` | CONDITIONAL YES | *string* | The shipper's name. | John Doe |
| `address` | CONDITIONAL YES | *string* | The pickup location address. Min 10 characters. | Office Building |
| `address_2` | NO | *string* | Additional address details. Min 10 characters. | Street 2 house 4 |
| `city` | CONDITIONAL YES | *string* | The pickup location city. | Pune |
| `state` | CONDITIONAL YES | *string* | The pickup state. | Maharashtra |
| `country` | CONDITIONAL YES | *string* | The pickup location country. | India |
| `pin_code` | CONDITIONAL YES | *integer* | The pickup pincode. | 200099 |
| `pickup_location` | CONDITIONAL YES | *string* | New pickup location name. Max: 36 char. Alphanumeric only. | New Office |
| `order_type` | NO | *string* | Key to differentiate between Essentials or Non Essentials Shipments. Order type can only be ESSENTIALS or NON ESSENTIALS. Please note it is case sensitive and blank values are allowed | ESSENTIALS |
| `longitude` | NO | *float* | Destination (Shipping address) Longitude. | 69.0747 |
| `latitude` | NO | *float* | Destination (Shipping address) Latitude | 22.4064 |
| `what3words_address` | NO | *string* | What3words is a proprietary geocode system designed to identify any location on the surface of Earth with a resolution of about 3 meters. The system encodes geographic coordinates into three permanently fixed dictionary words. | toddler.geologist.animated |
| `is_document` | NO | *integer* | To create a document order | 1 or 0 |
| `reseller_name` | NO | *string* | To display the vendor name on the label | brandname |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "mode": "",
  "request_pickup": "",
  "print_label": "",
  "generate_manifest": "",
  "ewaybill_no": "",
  "courier_id": "",
  "reseller_name": "",
  "order_id": "",
  "isd_code": "",
  "billing_isd_code": "",
  "order_date": "",
  "channel_id": "",
  "company_name": "",
  "billing_customer_name": "",
  "billing_last_name": "",
  "billing_address": "",
  "billing_address_2": "",
  "billing_city": "",
  "billing_state": "",
  "billing_country": "",
  "billing_pincode": "",
  "billing_email": "",
  "billing_phone": "",
  "billing_alternate_phone": "",
  "shipping_is_billing": "1",
  "shipping_customer_name": "",
  "shipping_last_name": "",
  "shipping_address": "",
  "shipping_address_2": "",
  "shipping_city": "",
  "shipping_state": "",
  "shipping_country": "",
  "shipping_pincode": "",
  "shipping_email": "",
  "shipping_phone": "",
  "order_items": [
    {
      "name": "",
      "sku": "",
      "units": "",
      "hsn": "",
      "selling_price": "",
      "tax": "",
      "discount": ""
    }
  ],
  "payment_method": "",
  "shipping_charges": "",
  "giftwrap_charges": "",
  "transaction_charges": "",
  "total_discount": "",
  "sub_total": "",
  "weight": "",
  "length": "",
  "breadth": "",
  "height": "",
  "pickup_location": "",
  "customer_gstin": "",
  "vendor_details": {
    "email": "",
    "phone": "",
    "name": "",
    "address": "",
    "address_2": "",
    "city": "",
    "state": "",
    "country": "",
    "pin_code": "",
    "pickup_location": ""
  }
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
 
  "order_id": "22114477",
  "order_date": "2018-05-08 12:23",
  "channel_id": "27202",
  "billing_customer_name": "Jax",
  "billing_last_name": "Tank",
  "billing_address": "Dust2",
  "billing_city": "New Delhi",
  "billing_pincode": 110002,
  "billing_state": "Delhi",
  "billing_country": "India",
  "billing_email": "jax@counterstike.com",
  "billing_phone": 9988998899,
  "shipping_is_billing": true,
   "order_items": [
    {
      "name": "T-shirt Round Neck",
      "sku": "t-shirt-round1474",
      "units": 10,
      "selling_price": 400

    }
  ],

  "payment_method": "COD",
  "sub_total": 4000,
  "length": 100,
  "breadth": 50,
  "height": 10,
  "weight": 0.50,
  "pickup_location": "HomeNew",
  "vendor_details": {
    "email": "abcdd@abcdd.com",
    "phone": 9879879879,
    "name": "Coco Cookie",
    "address": "Street 1",
    "address_2": "",
    "city": "delhi",
    "state": "new delhi",
    "country": "india",
    "pin_code": "110077",
    "pickup_location": "HomeNew"
  }
}
```

```json
{
    "status": 1,
    "payload": {
        "pickup_location_added": 0,
        "order_created": 1,
        "awb_generated": 1,
        "label_generated": 1,
        "pickup_generated": 1,
        "manifest_generated": 1,
        "pickup_scheduled_date": "2022-06-04 09:00:00",
        "pickup_booked_date": null,
        "order_id": 222521420,
        "shipment_id": 222002884,
        "awb_code": "14326421307048",
        "courier_company_id": 25,
        "courier_name": "Xpressbees 5kg",
        "assigned_date_time": {
            "date": "2022-06-03 13:52:05.051557",
            "timezone_type": 3,
            "timezone": "Asia/Kolkata"
        },
        "applied_weight": 10,
        "cod": 1,
        "label_url": "https://kr-shipmultichannel.s3.ap-southeast-1.amazonaws.com/25149/labels/shipping-label-222002884-14326421307048.pdf",
        "manifest_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/25149/manifest/MANIFEST-4757.pdf",
        "routing_code": "N/S-01/12B/2",
        "rto_routing_code": "",
        "pickup_token_number": "Reference No: 194_BIGFOOT 2540335_04062022"
    }
}
```

#### Invalid Data — HTTP 200 OK

Example request body:

```json
{
 
  "order_id": "22907867",
  "order_date": "2018-05-08 12:23",
  "channel_id": "27202",
  "billing_customer_name": "Jax",
  "billing_last_name": "Tank",
  "billing_address": "Dust2",
  "billing_city": "New Delhi",
  "billing_pincode": "110002",
  "billing_state": "Delhi",
  "billing_country": "India",
  "billing_email": "jax@counterstike.com",
  "billing_phone": "9988998899",
  "shipping_is_billing": true,
   "order_items": [
    {
      "name": "Delta",
      "sku": "delta123",
      "units": 10,
      "selling_price": "1000"
    }
  ],

  "payment_method": "COD",
  "sub_total": 4000,
  "length": 100,
  "breadth": 50,
  "height": 10,
  "weight": 0.50,
  "pickup_location": "HomeNew",
  "vendor_details": {
    "email": "abcdd@abcdd.com",
    "phone": 9879879879,
    "name": "CustName",
    "address": "Street1",
    "address_2": "",
    "city": "delhi",
    "state": "new delhi",
    "country": "india",
    "pin_code": "110077",
    "pickup_location": "HomeNew"
  }
}
```

```json
{
    "status": 0,
    "payload": {
        "action": "Adding pickup location",
        "error_message": "Oops! Invalid Data."
    }
}
```

#### Missing Data — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

Example request body:

```json
{
 
  "order_id": "",
  "order_date": "2018-05-08 12:23",
  "channel_id": "27202",
  "billing_customer_name": "Jax",
  "billing_last_name": "Tank",
  "billing_address": "Dust2",
  "billing_city": "New Delhi",
  "billing_pincode": "110002",
  "billing_state": "Delhi",
  "billing_country": "India",
  "billing_email": "jax@counterstike.com",
  "billing_phone": "9988998899",
  "shipping_is_billing": true,
   "order_items": [
    {
      "name": "Delta",
      "sku": "delta123",
      "units": 10,
      "selling_price": "1000"
    }
  ],

  "payment_method": "COD",
  "sub_total": 4000,
  "length": 100,
  "breadth": 50,
  "height": 10,
  "weight": 0.50,
  "vendor_details": {
    "email": "abcdd@abcdd.com",
    "phone": 9879879879,
    "name": "",
    "address": "",
    "address_2": "",
    "city": "delhi",
    "state": "new delhi",
    "country": "india",
    "pin_code": "110077",
    "pickup_location": "Home"
  }
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

#### Forward — (no HTTP status recorded in source)

_(empty response body in source)_

### Return

`POST https://apiv2.shiprocket.in/v1/external/shipments/create/return-shipment`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to perform multiple tasks like Create, AWB generation & scheduling reverse pickups for your Returns.<br>The specifications are the same as the custom return order API, with a few exceptions.

**Notes:**

- pickup_location field is not required.
- Label and Manifest are not required in case of returns.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `order_id` | YES | *string* | The order id you want to specify to the order. Max char: 50. (Avoid passing character values as this contradicts some other API calls) | 99711997 |
| `order_date` | YES | *string* | The date of order creation in yyyy-mm-dd format. Time is an additional option. | 2019-08-05 |
| `channel_id` | NO | *integer* | Id of the desired channel where the order is to be placed. 'Custom' channel id is selected in case parameter is not filled. | 768903 |
| `pickup_customer_name` | YES | *string* | The customer’s first name. | John |
| `pickup_last_name` | NO | *string* | The customer’s last name. | Doe |
| `company_name` | NO | *string* | Name of the company | Amazon |
| `pickup_address` | YES | *string* | The customer's primary address. | Home |
| `pickup_address_2` | NO | *string* | Additional customer address details. | DDA |
| `pickup_city` | YES | *string* | The customer's city name. | Delhi |
| `pickup_state` | YES | *string* | The customer's state. | New Delhi |
| `pickup_country` | YES | *string* | Customer's country name. | India |
| `pickup_pincode` | YES | *integer* | Pincode of the customer address. | 110002 |
| `pickup_email` | YES | *string* | Customer's email address. | [john@doe.com](https://mailto:john@doe.com) |
| `pickup_phone` | YES | *string* | Customer's phone number. | 9999999999 |
| `pickup_isd_code` | NO | *string* | ISD code. | 91 |
| `shipping_customer_name` | YES | *string* | The name of the seller the package is shipped back to. | Jane |
| `shipping_last_name` | NO | *string* | The last name of the seller. | Doe |
| `shipping_address` | YES | *string* | The address the package is shipped to. | Castle |
| `shipping_address_2` | NO | *string* | Further shipping address details. | Bridge |
| `shipping_city` | YES | *string* | The shipping address city. | Mumbai |
| `shipping_country` | YES | *string* | The shipping address country. | India |
| `shipping_pincode` | YES | *integer* | The shipping pincode. | 220022 |
| `shipping_state` | YES | *string* | Shipping address state. | Maharashtra |
| `shipping_email` | YES | *string* | The email of the seller the package is shipped to. | [jane@doe.com](https://mailto:jane@doe.com) |
| `shipping_isd_code` | NO | *string* | The shipping isd code. | 91 |
| `shipping_phone` | YES | *integer* | Phone no. of the shipping customer | 8888888888 |
| `order_items` | YES | / | Array containing further fields. | / |
| `name` | YES | *string* | Name of the product. | ball123 |
| `sku` | YES | *string* | The sku id of the product. | Tennis Ball |
| `units` | YES | *integer* | No of units that are to be shipped. | 1 |
| `selling_price` | YES | *integer* | The selling price per unit in Rupee. Inclusive of GST. | 10 |
| `discount` | NO | *integer* | The discount amount in Rupee. Inclusive of tax. | 0 |
| `hsn` | NO | *string* | Harmonised System Nomenclature code. Used to determine the category of taxation the goods fall under. | 4412 |
| `qc_enable` | CONDITIONAL YES | *string* | If True, QC will be performed for that product and QC will be performed only for a single SKU per order | TRUE/FALSE |
| `qc_color` | NO | *varchar(180)* | The color of the product can be passed in this parameter | Red |
| `qc_brand` | NO | *varchar(255)* | The brand of the product can be passed in this parameter | 768903 |
| `qc_serial_no` | NO | *varchar(255)* | The serial number of the product can be passed in this parameter | T13123124 |
| `qc_ean_barcode` | NO | *varchar(255)* | EAN/Barcode of the product can be passed in this parameter | QWRE123 |
| `qc_size` | NO | *varchar(180)* | The size of the product can be passed in this parameter | 8 |
| `qc_product_name` | CONDITIONAL YES | *varchar(255)* | If qc_enable set True, then Product name should be passed in this parameter | Shoes |
| `qc_product_image` | CONDITIONAL YES | *varchar(255)* | If qc_enable set True, then Product image should be passed in this parameter (only png/jpg format supported) | [https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/1636713733zxja.png](https://s3-ap-southeast-1.amazonaws.com/kr-multichannel/1636713733zxja.png) |
| `qc_product_imei` | NO | *varchar(255)* | IMEI of the device | 86532976457823 |
| `payment_method` | YES | *string* | The method of payment. Can be either COD (Cash on delivery) Or Prepaid. | Prepaid |
| `total_discount` | NO | *string* | The total discount amount in Rupee. | 0 |
| `sub_total` | YES | *integer* | Calculated sub total amount in Rupee after deductions. | 10 |
| `length` | YES | *integer* | The length of the shipment in cms. | 10 |
| `breadth` | YES | *integer* | The breadth of the shipment in cms. | 15 |
| `height` | YES | *integer* | The height of the shipment in cms. | 20 |
| `weight` | YES | *integer* | The shipment weight in kgs. | 1 |

**Request body template** (as published; empty strings are placeholders to fill in)

```text
{
"order_id": "",
"order_date": "",
"channel_id": "",
"pickup_customer_name": "",
"pickup_last_name": "",
"company_name":"",
"pickup_address": "",
"pickup_address_2": "",
"pickup_city": "",
"pickup_state": "",
"pickup_country": "",
"pickup_pincode": ,
"pickup_email": "",
"pickup_phone": "",
"pickup_isd_code": "",
"shipping_customer_name": "",
"shipping_last_name": "",
"shipping_address": "",
"shipping_address_2": "",
"shipping_city": "",
"shipping_country": "",
"shipping_pincode": ,
"shipping_state": "",
"shipping_email": "",
"shipping_isd_code": "",
"shipping_phone": ,
"order_items": [
{
"sku": "",
"name": "",
"units":"",
"selling_price":"",
"discount":""
}
],
"payment_method": "",
"total_discount": "",
"sub_total":"",
"length":"",
"breadth":"",
"height":"",
"weight":"",
"request_pickup":""
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
"order_id": "r121579B09ap492",
"order_date": "2022-02-16",
"channel_id": "2113680",
"pickup_customer_name": "iron man",
"pickup_last_name": "",
"company_name":"iorn pvt ltd",
"pickup_address": "b-123",
"pickup_address_2": "",
"pickup_city": "Delhi",
"pickup_state": "New Delhi",
"pickup_country": "India",
"pickup_pincode": 110030,
"pickup_email": "deadpool@red.com",
"pickup_phone": "9810363552",
"pickup_isd_code": "91",
"shipping_customer_name": "Jax",
"shipping_last_name": "Doe",
"shipping_address": "Castle",
"shipping_address_2": "Bridge",
"shipping_city": "Delhi",
"shipping_country": "India",
"shipping_pincode": 110015,
"shipping_state": "New Delhi",
"shipping_email": "kumar.abhishek@shiprocket.com",
"shipping_isd_code": "91",
"shipping_phone": 8888888888,
"order_items": [
{
"name": "shoes",
      "qc_enable":true,
      "qc_product_name": "shoes",
      "sku": "WSH234",
      "units": 1,
      "selling_price": 100,
      "discount": 0,
      "qc_brand":"Levi",
      "qc_product_image":"https://assets.vogue.in/photos/5d7224d50ce95e0008696c55/2:3/w_2240,c_limit/Joker.jpg"
}
],
"payment_method": "PREPAID",
"total_discount": "0",
"sub_total": 400,
"length": 11,
"breadth": 11,
"height": 11,
"weight": 0.5,
"request_pickup": true
}
```

```text
{
"status": 1,
"payload": {
"order_created": 1,
"awb_generated": 1,
"pickup_generated": 1,
"pickup_scheduled_date": "2022-02-17 09:00:00",
"order_id": 186255214,
"shipment_id": 185783600,
"awb_code": "24112321126686",
"courier_company_id": 125,
"courier_name": "Xpressbees Reverse",
"assigned_date_time": {
"date": "2022-02-16 16:34:06.963419",
"timezone_type": 3,
"timezone": "Asia/Kolkata"
},
"applied_weight": 0.5,
"cod": 0,
"is_return": 1,
"routing_code": "N/S-01/13B/015",
"rto_routing_code": "",
"pickup_token_number": "",
}
}
```

## File Imports

Check the status response of files that have been imported using the API.

### Get File import Results

`GET https://apiv2.shiprocket.in/v1/external/errors/{import_id}/check`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Path placeholders:** `import_id` (see the parameter table in the description for meaning where the source provides one).

**Description**

Check the import response of various file imports for any error, including bulk order imports, products and product listings.

You will have to pass the file import id as a path parameter in the endpoint URL. This id is provided at the time of importing the file. No other body parameters are required.

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/errors/20212061/check](https://apiv2.shiprocket.in/v1/external/errors/20212061/check) |

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/errors/20212061/check`

```json
{
    "data": {
        "status": "3",
        "message": "Error in reading file data! Please check that the column names are same as that in the sample file provided."
    }
}
```

#### Invalid Data — HTTP 500 Internal Server Error

Example request: `GET https://apiv2.shiprocket.in/v1/external/errors/22222222/check`

```json
{
    "message": "Trying to get property of non-object",
    "status_code": 500
}
```

#### Missing Fields — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/errors//check`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```
