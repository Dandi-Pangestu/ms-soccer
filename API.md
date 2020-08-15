**Add Team**
----
  Add a new team.

* **URL**

  http://localhost:8080/v1/teams

* **Method:**

  `POST`
  
*  **Headers**

   None
  
*  **URL Params**

   None

* **Data Params**

  **Required:**
   
  `{ "name" : "Chelsea" }`

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** `{ "message" : "Success" }`
 
  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error" }`

**Get List Team**
----
  Return all team.

* **URL**

  http://localhost:8080/v1/teams

* **Method:**

  `GET`
  
*  **Headers**

   None
  
*  **URL Params**

   None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `[{ "id": "dd6af1f5-bbaa-47b6-8a9a-4104e875e1ca", "name": "Chelsea" }, { "id": "0e0a4e1a-7a8f-4495-9fb8-25130f465b1d", "name": "Barcelona" }]`
 
  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error" }`

**Add Player**
----
  Add a new player.

* **URL**

  http://localhost:8080/v1/teams/:id/players

* **Method:**

  `POST`
  
*  **Headers**

   None
  
*  **URL Params**

   None

* **Data Params**

  **Required:**
   
  `{ "name": "John", "number": 10 }`

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** `{ "message" : "Success" }`
 
  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error" }`

**Get List Team**
----
  Return all team.

* **URL**

  http://localhost:8080/v1/teams/:id/players

* **Method:**

  `GET`
  
*  **Headers**

   None
  
*  **URL Params**

   None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `[{ "id": "f506406e-dd7e-435b-8aae-d5ed39a70733", "name": "Dandi", "number": 10 }]`
 
  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error" }`
