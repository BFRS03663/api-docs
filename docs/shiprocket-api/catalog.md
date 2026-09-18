# Shiprocket API — Products, Listings, Channels & Inventory

> Generated from the published Postman collection behind [https://apidocs.shiprocket.in/](https://apidocs.shiprocket.in/) (collection "Shiprocket API", published id `SzYW1zB2`, version tag `latest`). Only content present in that collection is reproduced here; where the source omits something, the omission is stated.

## Products

Use these APIs to get information on your added products or update their details. You can also add a new product or import them in bulk.

### Get All Products

`GET https://apiv2.shiprocket.in/v1/external/products`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API can be used to display a detailed list of all the products that you have in your Shiprocket account.

There are no required parameters to access this API. However, the displayed result can be filtered or sorted using additional parameters.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `page` | NO | *integer* | The page number you want to display. | 5 |
| `per_page` | NO | *integer* | The number of products to get per page. | 2 |
| `sort` | NO | *string* | The order to sort by. Value: *ASC* or *DESC* | ASC |
| ` sort_by` | NO | *string* | Allows you to choose the value field by which the items will be sorted. Could be sorted by id, by sku, time created etc. | sku |
| `filter` | NO | *string* | The data to be matched for the filter value. | 11223344 |
| `filter_by` | NO | *string* | The filter value field . Can be id, sku, etc. | id |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": [
        {
            "id": 17484610,
            "sku": "chakra123",
            "hsn": "441122",
            "name": "Kunai",
            "description": "",
            "category_code": "default",
            "category_name": "Default Category",
            "category_tax_code": "",
            "image": "",
            "weight": "0 kg",
            "size": "",
            "cost_price": "0.00",
            "mrp": "0.00",
            "tax_code": "default",
            "low_stock": 0,
            "ean": "",
            "upc": "",
            "isbn": "",
            "created_at": "31 Jul 2019 12:37 PM",
            "updated_at": "31 Jul 2019 03:18 PM",
            "quantity": 41,
            "color": "",
            "brand": "",
            "dimensions": "10 x 10 x 10 cm",
            "status": "INACTIVE",
            "type": "Single"
        },
        {
            "id": 9741478,
            "sku": "LNO7K1670",
            "hsn": "",
            "name": "hehehprod",
            "description": "",
            "category_code": "default",
            "category_name": "Default Category",
            "category_tax_code": "",
            "image": "",
            "weight": "0.5 kg",
            "size": "",
            "cost_price": "0.00",
            "mrp": "0.00",
            "tax_code": "default",
            "low_stock": 0,
            "ean": "",
            "upc": "",
            "isbn": "",
            "created_at": "22 Apr 2019 12:48 PM",
            "updated_at": "22 Apr 2019 12:48 PM",
            "quantity": 0,
            "color": "",
            "brand": "",
            "dimensions": "10 x 10 x 10 cm",
            "status": "INACTIVE",
            "type": "Single"
        }
    ],
    "meta": {
        "pagination": {
            "total": 12080,
            "count": 15,
            "per_page": 15,
            "current_page": 1,
            "total_pages": 806,
            "links": {
                "next": "https://apiv2.shiprocket.in/v1/external/products?page=2"
            }
        }
    }
}
```

#### Invalid Data — HTTP 200 OK

```json
{
    "data": [],
    "meta": {
        "pagination": {
            "total": 0,
            "count": 0,
            "per_page": 15,
            "current_page": 3,
            "total_pages": 1,
            "links": {
                "previous": "https://apiv2.shiprocket.in/v1/external/products?page=2"
            }
        }
    }
}
```

### Get Specific Product Details

`GET https://apiv2.shiprocket.in/v1/external/products/show/{product_id}`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Path placeholders:** `product_id` (see the parameter table in the description for meaning where the source provides one).

**Description**

Use this API to get the details of a specific product. The product details will be displayed in JSON format.

You need to pass the product id in the endpoint URL for the successful call of the API. No other body parameters are required.

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/products/show/17484610](https://apiv2.shiprocket.in/v1/external/products/show/17484610) |

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/products/show/17484610`

```json
{
    "data": {
        "id": 17484610,
        "sku": "chakra123",
        "name": "Kunai",
        "description": "",
        "category_code": "",
        "category_name": "",
        "category_tax_code": "",
        "image": "",
        "weight": "0.000",
        "size": "",
        "cost_price": "0.00",
        "mrp": "0.00",
        "tax_code": "",
        "low_stock": 0,
        "ean": "",
        "upc": "",
        "isbn": "",
        "created_at": "31 Jul 2019 12:37 PM",
        "updated_at": "31 Jul 2019 03:18 PM",
        "quantity": 41,
        "color": "",
        "brand": "",
        "dimensions": "0.00 x 0.00 x 0.00",
        "status": "INACTIVE",
        "is_combo": 0
    }
}
```

#### Invalid Data — HTTP 400 Bad Request

Example request: `GET https://apiv2.shiprocket.in/v1/external/products/show/11111111`

```json
{
    "message": "This product is either inactive or does not exist",
    "status_code": 400
}
```

#### Missing Fields — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/products/show/`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Add New Products

`POST https://apiv2.shiprocket.in/v1/external/products`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to add a new product to your Shiprocket account.<br>Provide the required product details and any additional info to successfully add a new product to your product list.

**Notes:**

- 'sku' Id has to be unique. It cannot be the same as an existing sku.
- In case no category code is added, the code should be default.
- 'type' field should be either 'single' or 'multiple.'

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `sku` | YES | *string* | Stock Keeping Unit or the identification unit of an individual product (generally alphanumeric). | bat123 |
| `HSN` | NO | *string* | Harmonised System Nomenclature. A code number is used to classify goods for taxation purposes. Done to determine which category of taxes do the goods come under. | 4412 |
| `name` | YES | *string* | Name of the product. | Batman Toy |
| `tax_code` | NO | *string* | The percentage of tax that is to be imposed. | 10 |
| `type` | YES | *string* | If there is only one product or multiple types of products. **Single** or **Multiple** | Single. |
| `qty` | YES | *integer* | Total Quantity of the products to be shipped. | 5 |
| `low_stock` | NO | *string* | Specifies when the low stock notification should come on | / |
| `category_code` | YES | *string* | You can add a category code to your ShipRocket account from “add category” | default |
| `description` | NO | *string* | Gives a description of the product. | Batman plastic toy. |
| `brand` | NO | *string* | The product brand name. | Bat |
| `size` | NO | *integer* | The size of the product. | 25 |
| `weight` | NO | *integer* | The weight of the product in kgs. | 0.5 |
| `length` | NO | *integer* | The length of the product in cms. | 10 |
| `width` | NO | *integer* | The width of the product in cms. | 5 |
| `height` | NO | *integer* | The height of the product in cms. | 15 |
| `ean` | NO | *string* | European Article Number - A barcode for product identification (which helps manufacturers identify how many products have been sold once a sale is made). It is 13 digits long and required for international selling. | / |
| `upc` | NO | *string* | Universal Product Code – Barcode for product identification which is used across the world. It is 12 digits long. | / |
| `isbn` | NO | *string* | International Standard Book Number – Identification barcode for books, magazines, e-books and other published media. It is 10 digits long. | / |
| `color` | NO | *string* | The colour of the product. | Black |
| `imei_serialnumber` | NO | *string* | The International Mobile Equipment Identity Number, which is used by a network to identify valid devices. E.g. if two iPhones have to be shipped, they will have 2 IMEI numbers | / |
| `cost_price` | NO | *integer* | The manufacture cost price of the product. | 500 |
| `mrp` | NO | *string* | Maximum Retail Price. How much is the maximum price which the product can be sold at. | 1000 |
| `status` | NO | *boolean* | In Boolean, if the product details have been successfully or unsuccessfully added. | 1 |
| `image_url` | NO | *string* | Shows the URL of the product images which have been uploaded. | / |
| `qc_details` | NO | / | List of items and their relevant fields in the form of Array. | / |
| `product_image` | CONDITIONAL YES | *string* | Pickup agent will cross check theshared product with the actual product received from the buyer. Mandatory for all QC Products | [https://abc/xyz.jpg](https://abc/xyz.jpg) |
| `brand` | CONDITIONAL YES | *string* | The pickup agent will cross check the provided brand name visible on the item(s) or its packaging. | shiprocket |
| `brand_tag` | CONDITIONAL YES | *string* | The pickup agent will cross-check the provided brand name, which should match the brand tag affixed to the item(s) upon delivery. | shiprocket |
| `color` | CONDITIONAL YES | *string* | Pickup agent will cross check shared product color with the actual product received from the buyer. | green |
| `size` | CONDITIONAL YES | *string* | Pickup agent will cross check shared product size with the size on the label/tag. | L |
| `product_imei` | CONDITIONAL YES | *string* | It is a 15 digit unique number. It is displayed on the screen, on the box, or at the back of the appliance. This can be used to check production & garauntee of the appliance. | 0123456781234 |
| `serial_no` | CONDITIONAL YES | *string* | A serial number (SN) is a unique alphanumeric value assigned to each individual product. | 123456 |

**Request body template** (as published; empty strings are placeholders to fill in)

```text
{
	 "name": "Batman451",
    "category_code": "default",
    "type": "Single",
    "qty": "10",
    "sku": "b118771212",
    "qc_details": {
        "product_image": "https://kr-multichannel-stage.s3.ap-south-1.amazonaws.com/1310/qc_product_img/538c2939-24e7-4b60-98cd-b13cee264c1e.jpg",
        "brand": "redlabel",
        "color": "white",
        "size": "L",
        "product_imei": "",
        "serial_no": "790878",
        "ean_barcode": "",
        "check_damaged_product": true
}
```

**Example responses**

#### Successful Call — HTTP 201 Created

Example request body:

```json
{
	"name": "Batman",
	"category_code": "default",
	"type": "Single",
	"qty": "10",
	"sku": "bat1234"
}
```

_(empty response body in source)_

#### Missing Fields — HTTP 422 Unprocessable Entity

Example request body:

```json
{
	"name": "Batman",
	"category_code": "default",
	"type": "Single",
	"qty": "10"

}
```

```json
{
    "message": "There were errors in adding new product!",
    "errors": {
        "sku": [
            "The sku field is required."
        ]
    },
    "status_code": 422
}
```

### Convert to QC Product

`POST https://apiv2.shiprocket.in/v1/external/products/qc-product-update/{productID}`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Path placeholders:** `productID` (see the parameter table in the description for meaning where the source provides one).

**Description**

Use this API to convert an existing product to a QC product. The {productID} will be the "id" from the "[Get all products](https://apidocs.shiprocket.in/#0b8d1f26-3abd-4f4e-9cd8-3928bcfcf30b)" or "[Get specific product details](https://apidocs.shiprocket.in/#134f7710-660c-464f-b579-6da46ba9402f)" API.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `sku` | YES | *string* | Stock Keeping Unit or the identification unit of an individual product (generally alphanumeric). | bat123 |
| `product_image` | YES | *string* | Pickup agent will cross check shared product color with the actual product received from the buyer. Mandatory for all QC Products | [https://abc/xyz.jpg](https://abc/xyz.jpg) |
| `serial_no` | NO | *string* | A serial number (SN) is a unique alphanumeric value assigned to each individual product. | 12345 |
| `size` | NO | *string* | Pickup agent will cross check shared product size with the size on the label/tag. | L |
| `color` | NO | *string* | Pickup agent will cross check shared product color with the actual product received from the buyer. | Green |
| `brand` | NO | *integer* | The pickup agent will cross-check the provided brand name visible on the item(s) or its packaging. | shiprocket |
| `brand_box` | NO | *string* | The pickup agent will cross-check the provided brand name, which should match the brand tag affixed to the item(s) upon delivery | shiprocket |
| `product_imei` | NO | *string* | "It is a 15-digit unique number. It is displayed on the screen, on the box, or at the back of the appliance. This can be used to check production & guarantee of the appliance. | 1234567890 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
  "sku": "PROD12345",
  "product_image": "https://kr-multichannel.s3.ap-southeast-1.amazonaws.com/2/products/images/1640350781_08d63a46-a0f1-4d6b-9eea-9e476fce7c4e1575377853407-HIGHLANDER-by-Rohit-Sharma-Men-White--Blue-Slim-Fit-Checked--1.j",
  "brand_box": "12323chacha",
  "brand": "blue",
  "color": "Red",
  "size": "M",
  "serial_no": "",
  "check_damaged_product": 0
}
```

**Example responses**

#### Convert to QC Product — (no HTTP status recorded in source)

Example request: `POST https://apiv2.shiprocket.in/v1/external/products/qc-product-update/76216`

```json
{
    "status": 200,
    "message": "Product details updated successfully!"
}
```

### Bulk Import Products

`POST https://apiv2.shiprocket.in/v1/external/products/import`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to import your products in bulk from a .csv file.<br>No parameters are required. Choose the target file as required. You will receive an import id upon successful import operation.

**Request body** (`multipart/form-data`)

| Field | Type | Value |
|---|---|---|
| `file` | file | (file upload) |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "id": 20290943
}
```

#### Missing Fields — HTTP 422 Unprocessable Entity (WebDAV) (RFC 4918)

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

### Get Sample .csv Format

`GET https://apiv2.shiprocket.in/v1/external/products/sample`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API will provide a sample format for a CSV file that can be used for importing orders. You can use this format to create your CSV file, which you want to use to import products into your Shiprocket account.

No additional parameters are required.

**Example responses**

#### Successful Call — HTTP 200 OK

```text
"Category Name","*Master Sku Code","*Product Name","Low Stock Warning At","Description","Length (cm)","Width (cm)","Height (cm)","Weight (kgs)","ean","upc","isbn","Color","Brand","Size","Tax Code","Image Url","Custom Detail Fields (IMEI/SerialNumber)","MRP","Cost Price","Active(True/False)","Type(Single)","HSN code"
"Below given is the sample data. Please delete the same and create/upload data as per you requirements. ","","","","","","","","","","","","","","","","","","","","","",""
"Default","Tshirt-Blue-42","Blue tshirt 42 size","2","blue polo tshirt with round collor large size","50","42","","500","1235467891111","","","blue","nike","42","5","http://www.onlinebachat.com/photos/MjAxNi0wMy0yMCAwNzowMzo1Nw==_tshirt%20yellow.jpg","","2500","1500","","",""
"Electronics","Iphone5s","Iphone5s","1","iphone 5s with 1 year warranty...","","","","","","","","Black","Iphone","","","","IMEI","20000","","","",""
"Electronics","Iphone6","Iphone6","1","iphone 6 with 1 year manufacturer warranty..","","","","","","","","Silver","Iphone","","","","IMEI","50000","","True","",""
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/products/sampl`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

## Listings

_No folder-level description in the source collection._

### Get All Listings

`GET https://apiv2.shiprocket.in/v1/external/listings`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API can be used to get a JSON representation of all the product listings in your Shiprocket account, i.e., all the products associated with a specific channel.

No parameters are required to access this API. However, the displayed data can be filtered and sorted using further parameters. If no sort parameter is used, the data is displayed in the default format.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `page` | NO | *integer* | The page number you want to display. | 5 |
| `per_page` | NO | *integer* | The number of listings to get per page. | 2 |
| `sort` | NO | *string* | The order to sort by. *Value*: **ASC** or **DESC** | ASC |
| `sort_by` | NO | *string* | Allows you to choose the value field by which the listings will be sorted. Could be sorted by id, by sku, time created etc. | sku |
| `filter` | NO | *string* | The data to be matched for the filter value. | 11223344 |
| `filter_by` | NO | *string* | The filter value field . Can be id, sku, etc. | id |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": [
        {
            "id": 15897064,
            "title": "Kunai",
            "image": null,
            "price": "900.00",
            "quantity": 0,
            "sku": "chakra123",
            "channel_sku": "chakra123",
            "channel_id": 76893,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_product_id": "",
            "inventory": 41,
            "synced_on": "Never Synced",
            "product": {
                "dimensions": {
                    "length": "0.000",
                    "width": "0.000",
                    "height": "0.000"
                },
                "weight": "0.000"
            },
            "category_name": ""
        },
        {
            "id": 14312566,
            "title": "Ferrari Pitstop Black iPhone 7 Case",
            "image": null,
            "price": "100.00",
            "quantity": 0,
            "sku": "SW-546",
            "channel_sku": "SW-546",
            "channel_id": 76893,
            "channel_name": "CUSTOM",
            "base_channel_code": "CS",
            "channel_product_id": "",
            "inventory": 0,
            "synced_on": "Never Synced",
            "product": {
                "dimensions": {
                    "length": "0.000",
                    "width": "0.000",
                    "height": "0.000"
                },
                "weight": "0.000"
            },
            "category_name": ""
        }
    ],
    "meta": {
        "pagination": {
            "total": 12345,
            "count": 2,
            "per_page": 2,
            "current_page": 1,
            "total_pages": 6173,
            "links": {
                "next": "https://apiv2.shiprocket.in/v1/external/listings?page=2"
            }
        }
    }
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/listing`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Map Channel Product

`POST https://apiv2.shiprocket.in/v1/external/listings/link`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to map a product present in the channel catalogue to a product present in the master catalogue.

Pass the product and listing id for the successful call of the API

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `product_id` | YES | *integer* | The id of item in the master catalog. | 17908342 |
| `listing_id` | YES | *integer* | The id of the product in the channel catalog. | 15897064 |
| `ID` | NO | *integer* | The id placed in the respective 'GET' codes. | 15897064 |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	"product_id": "",
	"listing_id": "",
	"ID": ""

	
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request body:

```json
{
	"product_id": 17908342,
	"listing_id": 15897064

	
}
```

```json
{
    "message": "Linking success",
    "status_code": 200
}
```

#### MIssing Fields — HTTP 400 Bad Request

Example request body:

```json
{
	"product_id": 17908342

	
}
```

```json
{
    "message": "Listing ID not found",
    "status_code": 400
}
```

#### Already Mapped Product — HTTP 400 Bad Request

```json
{
    "message": "Listing already mapped",
    "status_code": 400
}
```

#### Invalid Data — HTTP 400 Bad Request

Example request body:

```json
{
    "product_id": 17777771,
    "listing_id": 15897064
}
```

```json
{
    "message": "Product ID not found",
    "status_code": 400
}
```

### Import Catalog Mappings

`POST https://apiv2.shiprocket.in/v1/external/listings/import`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Description**

Use this API to import a CSV file containing channel catalogue to the master catalogue mappings.

No other body parameters are required.

**Request body** (`multipart/form-data`)

| Field | Type | Value |
|---|---|---|
| `file` | file | (file upload) |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "id": 20294650
}
```

#### Missing Fields — HTTP 400 Bad Request

```json
{
    "message": "File required",
    "status_code": 400
}
```

### Export Mapped Products

`GET https://apiv2.shiprocket.in/v1/external/listings/export/mapped`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API downloads the list of mapped items in your channel catalogue sheet. After mapping the items, you can see the number of products in one channel present in the Master Catalogue.

The downloaded CSV file URL is shared as the response. No other body parameters are required.

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "download_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/imports/c_67216/export-2019-08-051564990885.csv"
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/listings/export/mappe`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Export Unmapped Products

`GET https://apiv2.shiprocket.in/v1/external/listings/export/unmapped`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Get a list of all the unmapped products in your channel catalogue using this API.

The list is downloaded into a CSV file, and the download URL is displayed as the response. No other body parameters are required.

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "download_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/imports/c_67216/export-2019-08-051564992035.csv"
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/listings/export/unmappe`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Export Catalog Sample

`GET https://apiv2.shiprocket.in/v1/external/listings/sample`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API displays the download link of a sample catalogue sheet for reference purposes.

No additional parameters are required.

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "download_url": "https://s3-ap-southeast-1.amazonaws.com/kr-shipmultichannel/imports/c_67216/export-2019-08-051564991776.csv"
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/listings/sampl`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

## Channels

Use this API to get details about all the integrated channels in your Shiprocket account, including channel id. This channel id can be used to select or specify a custom channel at the time of order creation.

### Get Integrated Channel Details

`GET https://apiv2.shiprocket.in/v1/external/channels`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API shows a list of all the channels that have already been integrated with your Shiprocket account.

No parameters are required to access this API.

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": [
        {
            "id": 76893,
            "name": "CUSTOM",
            "status": "Active",
            "connection_response": null,
            "channel_updated_at": "2018-05-24 11:55:28",
            "status_code": 1,
            "settings": {
                "dimensions": "0x0x0",
                "weight": 0,
                "order_status": ""
            },
            "auth": [],
            "connection": 1,
            "orders_sync": 0,
            "inventory_sync": 0,
            "catalog_sync": 0,
            "orders_synced_on": "Not Available",
            "inventory_synced_on": "Not Available",
            "base_channel_code": "CS",
            "base_channel": {
                "id": 4,
                "name": "MANUAL",
                "code": "CS",
                "type": "Carts",
                "logo": "custom.png",
                "settings_sample": {
                    "name": "Channels Settings",
                    "help": "",
                    "settings": {
                        "brand_name": {
                            "code": "brand_name",
                            "name": "Brand Name",
                            "placeholder": "Your brand name",
                            "type": "text"
                        },
                        "brand_logo": {
                            "code": "brand_logo",
                            "name": "Brand Logo",
                            "placeholder": "Your brand logo",
                            "type": "file"
                        }
                    }
                },
                "auth_sample": [],
                "description": "Manual channel"
            },
            "catalog_synced_on": "31 Jul 12:49 PM",
            "order_status_mapper": "",
            "payment_status_mapper": "",
            "brand_name": "",
            "brand_logo": ""
        }
]
        
}
```

#### Wrong Endpoint Name — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/channel`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Create Custom Channel

`POST https://apiv2.shiprocket.in/v1/external/channels`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token). The request template lists the header explicitly with value `Bearer {{jwt_token}}`.

**Request headers**

| Header | Value |
|---|---|
| `Authorization` | `Bearer {{jwt_token}}` |
| `Content-Type` | `application/json` |

**Description**

Creates a new manual (custom) channel in the seller's Shiprocket account.

**Pre-condition:** This endpoint is gated. Your account must be allow-listed for "Channel Creation by API" by the Shiprocket Tech Support / KAM team before it will work.

#### Request body

| Field | Type | Required | Description |
|---|---|---|---|
| name | string | Yes | Channel display name. Max 12 characters. Unique within your account. |
| brand_name | string | No | Brand label on invoices and shipping labels. Max 50 characters. |

#### Error responses

- 401 Unauthorized

  — Missing or invalid bearer token.
- 403 Forbidden

  — "Channel Creation by API" flag is not enabled for your account, or channel limit reached.
- 422 Unprocessable Entity

  — Validation failure (name missing, name > 12 chars, name already exists, or brand_name > 50 chars).

#### Notes

- Default order-status mapping is applied automatically (matches the canonical mapping used when you create a manual channel from the Shiprocket dashboard).

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
    "name": "MANUAL-25149",
    "brand_name": "SIORA-12"
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "channel_id": 123456,
    "name": "MANUAL-25149",
    "brand_name": "SIORA-12",
    "base_channel_code": "CUSTOM",
    "status": 1,
    "company_id": 25149,
    "created_at": "2026-04-29 12:34:56"
}
```

## Inventory

Use this API to update and get Inventory details.

### Get Inventory Details

`GET https://apiv2.shiprocket.in/v1/external/inventory`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API can be used to check the inventory details of a product sku. There are no required parameters for this API, but additional parameters can be used to sort and filter the data.

In case no filter conditions are passed, the details are displayed in the default JSON format.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `page` | NO | *integer* | The page number to be displayed. | 4 |
| `per_page` | NO | *integer* | The total products you want to diplay in each page. | 2 |
| `sort` | NO | *string* | Sort conditions if any. Value: ASC or DESC | ASC |
| `sort_by` | NO | *integer* | Allows you to choose the field by which the data will be sorted. Could be sorted by id, by sku, time created etc. | sku |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "data": [
        {
            "id": 3448631,
            "sku": "123",
            "category_name": "Default Category",
            "is_combo": 0,
            "name": "gsdgw",
            "type": "Single",
            "color": "",
            "brand": "",
            "total_quantity": 6,
            "available_quantity": 6,
            "blocked_quantity": 0,
            "updated_on": "15 Oct 2018 04:14 PM"
        },
        {
            "id": 3448637,
            "sku": "BlackTshirt",
            "category_name": "Default Category",
            "is_combo": 0,
            "name": "Black tshirt XL",
            "type": "Single",
            "color": "",
            "brand": "",
            "total_quantity": 12,
            "available_quantity": 12,
            "blocked_quantity": 0,
            "updated_on": "23 Jan 2019 12:03 PM"
        }
    ],
    "meta": {
        "pagination": {
            "total": 12080,
            "count": 2,
            "per_page": 2,
            "current_page": 1,
            "total_pages": 6040,
            "links": {
                "next": "https://apiv2.shiprocket.in/v1/external/inventory?page=2"
            }
        }
    }
}
```

#### Wrong Endpoint — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/inventor`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Update Your Inventory

`PUT https://apiv2.shiprocket.in/v1/external/inventory/{product_id}/update`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Path placeholders:** `product_id` (see the parameter table in the description for meaning where the source provides one).

**Description**

This API is used to update your product inventory details.

First, you need to pass the product_id of the product in the endpoint URL. You can then set the quantity and action you want to perform on the existing inventory of the specified product.

**Note:**

- The product_id can be found using the 'Get Inventory Details' API.
- The id is to be passed in the endpoint URL itself.

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/inventory/17454637/update](https://apiv2.shiprocket.in/v1/external/inventory/17454637/update) |

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `quantity` | YES | *integer* | The quantity of the product you want. | 2 |
| `action` | YES | *integer* | The action you want to perform. *Value*: - **add** : Adds the specific quantity to the product inventory. - **replace** : Replaces the existing quantity with the specified number. - **remove** : Removes the specific number from the product inventory. | add |

**Request body template** (as published; empty strings are placeholders to fill in)

```json
{
	"quantity": "",
	"action": ""
}
```

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `PUT https://apiv2.shiprocket.in/v1/external/inventory/17454637/update`

Example request body:

```json
{
	"quantity": 2,
	"action": "remove"
}
```

```json
{
    "data": {
        "available_quantity": 51,
        "blocked_quantity": 0,
        "total_quantity": 51
    }
}
```

#### Missing Fields — HTTP 400 Bad Request

Example request: `PUT https://apiv2.shiprocket.in/v1/external/inventory/17454637/update`

```json
{
    "message": "Quantity must be numeric, and upto 8 digits are allowed!",
    "status_code": 400
}
```

#### Invalid Data — HTTP 400 Bad Request

Example request: `PUT https://apiv2.shiprocket.in/v1/external/inventory/12345678/update`

Example request body:

```json
{
	"quantity": "2",
	"action": "add"
}
```

```json
{
    "message": "Cannot update inventory for a combo product",
    "status_code": 400
}
```

## Countries

The following APIs provide details about the various country codes and their respective zone details along with the locality details.

### Get Country Codes

`GET https://apiv2.shiprocket.in/v1/external/countries`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

This API lists all the countries present in the Shiprocket database and the respective country ids, ISO 2 and ISO 3 codes.

There are a total of 44 available countries in the database. You can use these codes to check the serviceability and use them in your dropdown menu.

No parameters are required to access this API.

**Example responses**

#### Successful Call — HTTP 200 OK

```text
{
    "status": 200,
    "data": [
        {
            "id": 1,
            "name": "Afghanistan",
            "iso_code_2": "AF",
            "iso_code_3": "AFG",
            "isd_code": "+93",
            "address_format": "",
            "postcode_required": 1,
            "status": 1
        },
        {
            "id": 2,
            "name": "Albania",
            "iso_code_2": "AL",
            "iso_code_3": "ALB",
            "isd_code": "+355",
            "address_format": "",
            "postcode_required": 1,
            "status": 1
        },
        {
            "id": 3,
            "name": "Algeria",
            "iso_code_2": "DZ",
            "iso_code_3": "DZA",
            "isd_code": "+213",
            "address_format": "",
            "postcode_required": 1,
            "status": 1
        },
        {
            "id": 4,
            "name": "American Samoa",
            "iso_code_2": "AS",
            "iso_code_3": "ASM",
            "isd_code": "+1",
            "address_format": "",
            "postcode_required": 1,
            "status": 1
        },
        {
            "id": 5,
            "name": "Andorra",
            "iso_code_2": "AD",
            "iso_code_3": "AND",
            "isd_code": "+376",
            "address_format": "",
            "postcode_required": 1,
            "status": 1
        }
        	.
        	.
        	.
        	.
    ]
}
```

#### Wrong Endpoint — HTTP 404 Not Found

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

### Get All Zones

`GET https://apiv2.shiprocket.in/v1/external/countries/show/{country_id}`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Path placeholders:** `country_id` (see the parameter table in the description for meaning where the source provides one).

**Description**

Use this API to get a further list of all the available zones within a country, along with their ids and details.

The country ID must be passed as a path parameter to access this API. No other body parameters are required.

#### Path:

| **EXAMPLE** |
|---|
| [https://apiv2.shiprocket.in/v1/external/countries/show/4](https://apiv2.shiprocket.in/v1/external/countries/show/4) |

**Example responses**

#### Successful Call — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/countries/show/4`

```json
{
    "status": 200,
    "data": [
        {
            "id": 117,
            "country_id": 4,
            "code": "E",
            "state_code": null,
            "name": "Eastern",
            "status": 1
        },
        {
            "id": 118,
            "country_id": 4,
            "code": "M",
            "state_code": null,
            "name": "Manu'a",
            "status": 1
        },
        {
            "id": 119,
            "country_id": 4,
            "code": "R",
            "state_code": null,
            "name": "Rose Island",
            "status": 1
        },
        {
            "id": 120,
            "country_id": 4,
            "code": "S",
            "state_code": null,
            "name": "Swains Island",
            "status": 1
        },
        {
            "id": 121,
            "country_id": 4,
            "code": "W",
            "state_code": null,
            "name": "Western",
            "status": 1
        }
    ]
}
```

#### Missing Fields — HTTP 404 Not Found

Example request: `GET https://apiv2.shiprocket.in/v1/external/countries/show/`

```json
{
    "message": "404 Not Found",
    "status_code": 404
}
```

#### Invalid Data — HTTP 200 OK

Example request: `GET https://apiv2.shiprocket.in/v1/external/countries/show/900`

```json
{
    "status": 200,
    "data": [
        {
            "name": "None"
        }
    ]
}
```

### Get Locality Details

`GET https://apiv2.shiprocket.in/v1/external/open/postcode/details`

**Authentication:** Bearer token in the `Authorization: Bearer <token>` header (collection-wide guideline; token comes from Generate Token).

**Request headers**

| Header | Value |
|---|---|
| `Content-Type` | `application/json` |

**Description**

Use this API to get further locality details of any given postcode. Just pass the valid locality Pincode, and the details will be displayed in JSON format.

#### Parameters:

| **PARAMS** | **REQUIRED** | **DATA TYPE** | **DESCRIPTION** | **EXAMPLE** |
|---|---|---|---|---|
| `postcode` | YES | *integer* | The Pincode you want to get the locality details. | 110077 |

**Example responses**

#### Successful Call — HTTP 200 OK

```json
{
    "success": true,
    "postcode_details": {
        "postcode": "110077",
        "city": "South West Delhi",
        "locality": [
            "Bagdola",
            "Barthal",
            "Palam Extn (Harijan Basti)",
            "Dhulsiras",
            "Raj Nagar - II"
        ],
        "state": "Delhi",
        "state_code": "DL",
        "longitude": "77.399",
        "latitude": "28.2636"
    }
}
```

#### Invalid Data — HTTP 403 Forbidden

```json
{
    "message": "Mapped City-State details not found for 10000893",
    "status_code": 403
}
```

#### Missing Fields — HTTP 403 Forbidden

```json
{
    "message": "Mapped City-State details not found for ",
    "status_code": 403
}
```
