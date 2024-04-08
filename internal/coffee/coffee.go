package coffee

type Repository interface {
	Create(coffee Coffee) (Coffee, error)
	Get(coffeeID int) (Coffee, error)
	Update(coffee Coffee) error
	Delete(coffeeID int) error
}

/*
{
	"name": "Coffee",
	"flavours": ["citrus", "roibos", "caramel"],
	"price": {
		"amount": 11.99,
		"currency": "USD"
	},
}
*/

type Coffee struct {
	ID int
	// todo: #3 define Coffee struct
	// hint: for price, you can use common.Money
	// todo: #10 Add `roasted_on` DateOnly
}
