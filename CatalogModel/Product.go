package CatalogModel

type Context struct {
	setDiscount SetDiscount
}

func (Context) setStrategy(setDiscount SetDiscount) *Context {
	return &Context{setDiscount}
}

type SetDiscount interface {
	Discount()
}

//Concrete strategy with realization
type DiscountMeat struct {
}

func (DiscountMeat) Discount() {

}
