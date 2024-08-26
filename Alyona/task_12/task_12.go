package task_12

var prod_ID uint
var cust_ID uint
var ord_Number uint

type Shop struct {
	Name       string
	DomainName string
	products   []Product
	customers  []Customer
	orders     []Order
}

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       float64
}

type Customer struct {
	ID          uint
	Name        string
	PhoneNumber uint
	Address     string
	OrdersID    []uint
}

type Order struct {
	Number     uint
	ProductsID []uint
	Price      float64
	CustomerID uint
}

func NewProduct(name, description string, price float64) Product {
	prod_ID += 1
	p := Product{prod_ID, name, description, price}
	return p
}

func (s *Shop) AddProduct(p Product) {
	s.products = append(s.products, p)
}

func (s *Shop) UpdateProduct(p Product) {
	for i, val := range s.products {
		if val.ID == p.ID {
			s.products[i] = p
		}
	}
}

func (s *Shop) DeleteProduct(p Product) {
	for i, val := range s.products {
		if val.ID == p.ID && (i+1) < len(s.products) {
			s.products = append(s.products[:i], s.products[i+1])
		} else if val.ID == p.ID && (i+1) == len(s.products) {
			s.products = append(s.products[:i])
		}
	}
}

func (s *Shop) FindByProductName(name string) []Product {
	var pn []Product
	for _, val := range s.products {
		if val.Name == name {
			pn = append(pn, val)
		}
	}
	return pn
}

func NewCustomer(name, address string, number uint) Customer {
	cust_ID += 1
	oID := make([]uint, 0)
	c := Customer{cust_ID, name, number, address, oID}
	return c
}

func (s *Shop) AddCustomer(c Customer) {
	s.customers = append(s.customers, c)
}

func (s *Shop) UpdateCustomer(c Customer) {
	for i, val := range s.customers {
		if val.ID == c.ID {
			s.customers[i] = c
		}
	}
}

func (s *Shop) DeleteCustomer(c Customer) {
	for i, val := range s.customers {
		if val.ID == c.ID && (i+1) < len(s.customers) {
			s.customers = append(s.customers[:i], s.customers[i+1])
		} else if val.ID == c.ID && (i+1) == len(s.customers) {
			s.customers = append(s.customers[:i])
		}
	}
}

func (s *Shop) FindByCustomerName(name string) []Customer {
	var c_name []Customer
	for _, val := range s.customers {
		if val.Name == name {
			c_name = append(c_name, val)
		}
	}
	return c_name
}

func (s *Shop) FindByCustomerNumber(number uint) Customer {
	var c_num Customer
	for _, val := range s.customers {
		if val.PhoneNumber == number {
			c_num = val
		}
	}
	return c_num
}

func (s *Shop) FindByCustomerID(id uint) Customer {
	var c_id Customer
	for _, val := range s.customers {
		if val.ID == id {
			c_id = val
		}
	}
	return c_id
}

func (s *Shop) FindByCustomerAddress(address string) []Customer {
	var ca []Customer
	for _, val := range s.customers {
		if val.Address == address {
			ca = append(ca, val)
		}
	}
	return ca
}

func (s *Shop) FindByCustomerOrderID(id uint) Customer {
	var co Customer
	for _, val := range s.customers {
		for _, v := range val.OrdersID {
			if v == id {
				co = val
			}
		}
	}
	return co
}

func NewOder(customerID uint, products []Product) Order {
	ord_Number += 1
	pID := make([]uint, 0)
	var price float64
	for _, val := range products {
		pID = append(pID, val.ID)
		price += val.Price
	}
	o := Order{ord_Number, pID, price, customerID}
	return o
}

func (s *Shop) AddOrder(o Order) {
	s.orders = append(s.orders, o)
	for i, val := range s.customers {
		if val.ID == o.CustomerID {
			val.OrdersID = append(val.OrdersID, o.Number)
			s.customers[i] = val
		}
	}
}

func (s *Shop) UpdateOrder(o Order) {
	var (
		orderIndex int
		orderFound bool
	)

	for i, order := range s.orders {
		if order.Number == o.Number {
			orderIndex = i
			orderFound = true
		}
	}

	// если заказ не найден, мы не можем его обновить и просто выходим из функции
	if !orderFound {
		return
	}

	oldOrder := s.orders[orderIndex]

	if oldOrder.CustomerID != o.CustomerID {

		for i, val := range s.customers {
			if val.ID == oldOrder.CustomerID {
				// удалить айди ордера из ордерсайди
				for index, number := range val.OrdersID {
					if number == o.Number && (index+1) < len(val.OrdersID) {
						s.customers[i].OrdersID = append(val.OrdersID[:index], val.OrdersID[index+1])
					}
					if number == o.Number && (index+1) == len(val.OrdersID) {
						s.customers[i].OrdersID = append(val.OrdersID[:index])
					}
				}
			}
			//
			if s.customers[i].ID == o.CustomerID {
				s.customers[i].OrdersID = append(s.customers[i].OrdersID, o.Number)
			}

		}

		s.orders[orderIndex] = o

	}
}

func (s *Shop) DeleteOrder(o Order) {
	for i, val := range s.orders {
		if val.Number == o.Number && (i+1) < len(s.orders) {
			s.orders = append(s.orders[:i], s.orders[i+1])
		} else if val.Number == o.Number && (i+1) == len(s.orders) {
			s.orders = append(s.orders[:i])
		}
	}

	for i, val := range s.customers {
		if val.ID == o.CustomerID {
			// удалить айди ордера из ордерсайди
			for index, number := range val.OrdersID {
				if number == o.Number && (index+1) < len(val.OrdersID) {
					s.customers[i].OrdersID = append(val.OrdersID[:index], val.OrdersID[index+1])
				}
				if number == o.Number && (index+1) == len(val.OrdersID) {
					s.customers[i].OrdersID = append(val.OrdersID[:index])
				}
			}
		}
	}
}

func (s *Shop) FindByOrderNumber(number uint) Order {
	var o Order
	for _, val := range s.orders {
		if val.Number == number {
			o = val
		}
	}
	return o
}

func (s *Shop) FindByOrderProductsID(id uint) Order {
	var op Order
	for _, val := range s.orders {
		for _, v := range val.ProductsID {
			if v == id {
				op = val
			}
		}
	}
	return op
}

func (s *Shop) FindByOrderCustomerID(id uint) Order {
	var oc Order
	for _, val := range s.orders {
		if val.CustomerID == id {
			oc = val
		}
	}
	return oc
}
