**Auth Server**
----
Auth server that handles register, login, and verify token. Running on port `9081`

* **URL**
  /register
* **Method:**
  `POST`
*  **URL Params**
    None
* **Data Params**
   **Required:**
   `name=[string]`
   `role=[string]`
   `phone=[string]`
* **Success Response:**
  * **Code:** 200 <br />
    **Content:** `{"message":{"id":2,"name":"zulkoda","password":"q5iD","role":"admin","phone":"081322766362","updatedAt":"2021-02-01T03:17:11.297Z","createdAt":"2021-02-01T03:17:11.297Z"}}`
* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
    **Content:** `{ error : "Internal Server Error" }`
----
* **URL**
  /login
* **Method:**
  `POST`
*  **URL Params**
    none
* **Data Params**
   **Required:**
   `phone=[string]`
   `password=[string]`
* **Success Response:**
  * **Code:** 200 <br />
    **Content:** `{ token : "syTbgH_egD_YwJofpuv_7I7tTvGu0hKAqO9wbsqvtHFVUEavUUmLrF1byZIOVb8sceMFA" }`
* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
    **Content:** `{ error : "Internal Server Error" }`
  OR
  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "unauthorize token" }`
   ----
* **URL**
  /verify
* **Method:**
  `GET`
*  **URL Params**`
  None
* **Data Params**
  None
* **Success Response:**
  * **Code:** 200 <br />
    **Content:** `{"name":"zuronga","role":"not admin","phone":"081322766365","timestamp":1612149376969,"iat":1612149376}`
* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
    **Content:** `{ error : "Internal Server Error" }`
  OR
  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "unauthorize token" }`
  
