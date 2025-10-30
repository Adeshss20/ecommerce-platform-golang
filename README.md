# ecommerce-platform-golang

Functional Requirements :
1. Product should have add to cart and buy now  options
2. At home screen all products should be visible
3. Should be able to see and purchase the products
4. Should be able to see the products
5. At end after purchase screen now can have page for user details to send the products

Models :

Product
Properties
* id
* name
* desc
* quantity
* Price

Apis
* Get
* Post
* Delete — not needed as of now
* Update — not needed as of now

Order Properties
* Id
* Product-id
* Quantity
* Bill-id

Apis
* Get
* Post
* Delete —not needed as of now

Bill
Properties
* Id
* Order-id
* Total cost

Apis
* Get
* Post