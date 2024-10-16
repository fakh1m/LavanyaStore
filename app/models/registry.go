package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: Address{}},
		{Model: Product{}},
		{Model: ProductImage{}},
		{Model: Section{}},
		{Model: Category{}},
		{Model: Order{}},
		{Model: OrderCustomer{}},
		{Model: OrderItem{}},
		{Model: Payment{}},
		{Model: Shipment{}},
	}
}
