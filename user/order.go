package user

func orderCoffee(
	store *store,
	coffeeService CoffeeService,
	transactionService TransactionService,
	userID,
	coffeeID string) (string, error) {

	user, err := store.GetByID(userID)
	if err != nil {
		return "", err
	}

	coffee, err := coffeeService.GetCoffeeByID(coffeeID)
	if err != nil {
		return "", err
	}

	transaction, err := transactionService.NewTransaction(user.ID, coffee.ID)
	if err != nil {
		return "", err
	}

	return transaction.ID, nil
}
