db.createUser(
    {
        user: "admin",
        pwd: "pass",
        roles: [
            {
                role: "readWrite",
                db: "coffee_shop"
            }
        ]
    }
);