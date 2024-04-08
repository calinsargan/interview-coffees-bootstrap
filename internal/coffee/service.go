package coffee

type Service struct {
}

// todo: #4 link repository to service and use it in the service methods
func NewService() Service {
	return Service{}
}

func (self Service) Create(coffee Coffee) (Coffee, error) {
	return Coffee{}, nil
}

func (self Service) Get(coffeeID int) (Coffee, error) {
	return Coffee{}, nil
}

func (self Service) Update(coffee Coffee) error {
	// todo: #9 update only non-zero values from payload
	return nil
}

func (self Service) Delete(coffeeID int) error {
	// todo: #7 delete only if exists (and return error if not found)
	// todo: #8 test Service.Delete()
	return nil
}
