package User

import "NewProjectGo/Basket"

type user struct {
	name   string
	status bool

	basket Basket.Basket
}
