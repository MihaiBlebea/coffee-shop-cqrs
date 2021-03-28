## Plan

Create a coffee shop CQRS microservice

1. Create an api order coffee
    - api:
        - POST /api/v1/user - authenticates an user, returns a user id token that must be sent when placing an order
        - POST /api/v1/order - places an order for a coffee
    - database:
        - mariadb:
            - user
            - transaction
            - coffee

2. Create an api to get the coffee orders
    - api:
        - GET /api/v1/coffee - returns a list of all the available coffee types
        - GET /api/v1/coffee/{id} - returns the ingredients of a coffee by id
        - GET /api/v1/user/{id}/orders - returns a list of orders for a user
        - GET /api/v1/user - returns a list of users
    - database:
        - mongodb
            - coffee_view
            - user_view

3. Add rabbit mq or sqs for event sourcing
    - https://hub.docker.com/r/roribio16/alpine-sqs/
    - https://tutorialedge.net/golang/go-rabbitmq-tutorial/
    - https://www.youtube.com/watch?v=J_CGqiQx04c