**Fetch Server**
----
Fetch server that handles various data fetch and verify token. Running on port `9080`

* **URL**
  /fetch/aggregate
* **Method:**
  ``
*  **URL Params**
    None
* **Data Params**
    None
* **Success Response:**
  * **Code:** 200 <br />
* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
    **Content:** `{ error : "Internal Server Error" }`
----
* **URL**
  /fetch/with-usd
* **Method:**
  `POST`
*  **URL Params**
    None
* **Data Params**
    None
* **Success Response:**
  * **Code:** 200 <br />
* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
    **Content:** `{ error : "Internal Server Error" }`
  OR
  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "unauthorize token" }`
   ----
* **URL**
  /fetch/verify-token
* **Method:**
  `GET`
*  **URL Params**`
  None
* **Data Params**
  None
* **Success Response:**
  * **Code:** 200 <br />
* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
    **Content:** `{ error : "Internal Server Error" }`
  OR
  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "unauthorize token" }`
  
