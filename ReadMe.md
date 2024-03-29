# Merge Coding Assignment
Q. Build a Shopping cart API with following functionalities:
- Ability to create account with two roles (admin, user) and log in
- Admin should be able to
    - Add items
    - Suspend user
- User should be able to
    - List available items
    - Add items to a cart (if there are items in stock)
    - Remove items from their cart
- Restrict the access to APIs through RBAC mechanism
- Add unit test with extensive cases wherever possible (Bonus: E2E tests)
- Follow proper coding conventions and add documentation & code comments wherever necessary
- Create Readme for easy setup and testing of the API

Note: Candidates may use any coding language, though Golang would be preferable

> Merge Money Ltd │7 Bell Yard, London WC2A 2JR │ Tel +44.203.442.0175│Registration Number: 13463502
# Description

The project has app which exposes two http handlers:
1. APP API
    it contains the cart API (v1). It is divided into three groups for users, cart and items management. each group has it's own dependency system so that this can be broken down to microservices in the future.
2. Debug API
    it contains pprof, liveness, readiness probes as well Go's debug tooling
3. Scratch
    it contains bootstrap code like migration and seed scripts. It can be viewed in [schema](https://github.com/erdahuja/mergedup/blob/main/business/data/dbschema/sql/schema.sql), [seed](https://github.com/erdahuja/mergedup/blob/main/business/data/dbschema/sql/seed.sql)

As a rule of thumb, import graphs are as follows app imports business imports foundation.

    app: http handlers (can be swapped with socket, rpc easily)
    business: core business logic, data and RBAC system
    foundation: building blocks of a web server (swap router without changing anything!)


For RBAC mechanism, we are using jwt tokens with two roles (admin and user) saved in claims. By using a middleware for each handler we are doing authentication and authorization.
A cache layer is also added in users db to quick access user roles.

[DB Design](https://github.com/erdahuja/mergedup/blob/main/docs/db/dbb_design.pdf) | [API design](https://github.com/erdahuja/mergedup/blob/main/docs/api/mergedup.md)

## Getting started
For installing Go, please follow
[go official guide](https://go.dev/doc/install)

> A Makefile is also available to run basic commands. Please use the same. It is available in mac by default

The command line versions can now be installed straight from the command line itself;

    1. Open "Terminal" (it is located in Applications/Utilities)
    2. In the terminal window, run the command xcode-select --install
    3. In the windows that pops up, click Install, and agree to the Terms of Service.

## Commands
`make db`: db can be reset using 

`make run`: project can be run using 

`make status-debug`/`make status-api`: server status can be checked running or not

## First Step
Two users will be seeded on running migration (already done if you want to skip the step)

    username: admin@example.com
    password: admin

    username: user1@example.com
    password: user1

Please use [api](https://github.com/erdahuja/mergedup/blob/main/docs/api/mergedup.md#end-point-get-token) for getting bearer token. You have to use basic auth to be able to generate token. (authorization header with username/password)

Once you have the token you can use the same to try out different api. for users/admin apis will reject/accept as per role defined.

To use cart:

    Create cart using cart api
    post cart items into cart 
    remove item from cart

The quantites in the items table will be altered.

## API

Please download postman collection from [here](https://elements.getpostman.com/redirect?entityId=26793134-37605187-5b1a-4cdf-86b7-c82e7878094c&entityType=collection)

| API  | Policy |
| ------------- | ------------- |
| Get token  | public  |
| Create user | ruleAdmin  |
| Get all users  | ruleAdmin  |
| Get user by id | ruleAny  |
| Suspend user  | ruleAdmin |
| Create items  | ruleAdmin  |
| Get items  | ruleAny  |
| Create cart  |ruleUser  |
| Add item to cart  | ruleUser |
| View all items of cart  | ruleUser  |
| Delete item from cart  | ruleUser  |
| server status  | public  |

Documentation:

[API](https://github.com/erdahuja/mergedup/blob/main/docs/api/mergedup.md)

## DB Design
POSTGres db is used as the problem was of SQL type. however further scale we can add no sql/ cache server to specific problems (parts of app)
We are using hosted db for same for which "dev.env" file is provided already.
No db setup is required.

For pgAdmin exploration, credentials are available in a env file (though a secret manager would be ideal)

Sharing test credentials here
```DB_USERNAME=ocqbsygl
DB_PASSWORD=OTM2KccCSHPVll4u9rPZive247GuefZb
DB_NAME=ocqbsygl
DB_HOST=tiny.db.elephantsql.com
```

[Design](https://github.com/erdahuja/mergedup/blob/main/docs/db/dbb_design.pdf)

## E2E tests - WIP
We have used mockgen to mock database and tests can be found for business/core/item where quanities gets altered on add to cart and remove from cart.

I will add tests (unit or otherwise) on further discussions only. 