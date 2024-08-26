package task_12

import (
	"reflect"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p := Product{
		ID:          1,
		Name:        "Плед",
		Description: "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5",
		Price:       3500,
	}

	product1 := NewProduct(p.Name, p.Description, p.Price)

	if !reflect.DeepEqual(p, product1) {
		t.Error("ошибка создания продукта")
		t.FailNow()
	}
}

func TestShop_AddProduct(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	p1 := NewProduct("Настольная лампа", "Материал: платик", 1342.31)
	p2 := NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)

	s.AddProduct(p1)
	s.AddProduct(p2)

	if len(s.products) != 2 {
		t.Error("Ошибка добавления продукта")
		t.FailNow()
	}
}

func TestShop_UpdateProduct(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	p1 := NewProduct("Настольная лампа", "Материал: платик", 1342.31)
	p2 := NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)

	s.AddProduct(p1)
	s.AddProduct(p2)

	p2.Price = 3500.52
	s.UpdateProduct(p2)

	for _, val := range s.products {
		if val.ID == p2.ID {
			if val.Price != p2.Price {
				t.Error("Ошибка обновления данных товара")
				t.FailNow()
			}
		}
	}
}

func TestShop_DeleteProduct(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	p1 := NewProduct("Настольная лампа", "Материал: платик", 1342.31)
	p2 := NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)

	s.AddProduct(p1)
	s.AddProduct(p2)

	s.DeleteProduct(p2)

	if len(s.products) != 1 {
		t.Error("Ошибка удаления товара")
		t.FailNow()
	}
}

func TestShop_FindByProductName(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	p1 := NewProduct("Настольная лампа", "Материал: платик", 1342.31)
	p2 := NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)
	p3 := NewProduct("Плед", "Состав: шерсть, Цвет: чёрный, Размер: 2.0x1.5", 3500)

	s.AddProduct(p1)
	s.AddProduct(p2)
	s.AddProduct(p3)

	f := s.FindByProductName(p2.Name)

	if !reflect.DeepEqual(f, []Product{p2, p3}) {
		t.Error("Ошибка поиска товара по имени")
		t.FailNow()
	}
}

func TestNewCustomer(t *testing.T) {
	c := Customer{
		ID:          1,
		Name:        "Нурдавлетова Алёна",
		PhoneNumber: 89871301212,
		Address:     "Россия, Уфа, Бехтерева, 16, 169",
		OrdersID:    make([]uint, 0),
	}

	customer1 := NewCustomer(c.Name, c.Address, c.PhoneNumber)

	if !reflect.DeepEqual(c, customer1) {
		t.Error("Ошибка создания клиента")
		t.FailNow()
	}
}

func TestShop_AddCustomer(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Бехтерева, 16, 169", 89871301212)
	c2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)

	s.AddCustomer(c1)
	s.AddCustomer(c2)

	if len(s.customers) != 2 {
		t.Error("Ошибка добавления клиента")
		t.FailNow()
	}
}

func TestShop_UpdateCustomer(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алена", "Россия, Уфа, Бехтерева, 16, 169", 89871301212)
	c2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)

	s.AddCustomer(c1)
	s.AddCustomer(c2)

	c1.Name = "Нурдавлетова Алёна"
	s.UpdateCustomer(c1)

	for _, val := range s.customers {
		if val.ID == c1.ID {
			if val.Name != c1.Name {
				t.Error("Ошибка обновления данных клиента")
				t.FailNow()
			}
		}
	}
}

func TestShop_DeleteCustomer(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Бехтерева, 16, 169", 89871301212)
	c2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)

	s.AddCustomer(c1)
	s.AddCustomer(c2)

	s.DeleteCustomer(c2)

	if len(s.customers) != 1 {
		t.Error("Ошибка удаления клиента")
		t.FailNow()
	}
}

func TestShop_FindByCustomerName(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Бехтерева, 16, 169", 89871301212)
	c2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)
	c3 := NewCustomer("Иванов Иван", "Россия, Брянск, Брянского Фронта, 30, 1", 89098098899)
	c4 := NewCustomer("Иванов Иван", "Россия, Екатеринбург, Советская, 16, 34", 89995557665)

	s.AddCustomer(c1)
	s.AddCustomer(c2)
	s.AddCustomer(c3)
	s.AddCustomer(c4)

	f := s.FindByCustomerName(c3.Name)

	if !reflect.DeepEqual(f, []Customer{c3, c4}) {
		t.Error("Ошибка поиска клиента по имени")
		t.FailNow()
	}
}

func TestShop_FindByCustomerNumber(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Бехтерева, 16, 169", 89871301212)
	с2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)

	s.AddCustomer(c1)
	s.AddCustomer(с2)

	f := s.FindByCustomerNumber(c1.PhoneNumber)

	if !reflect.DeepEqual(f, c1) {
		t.Error("Ошибка поиска клиента по номеру телефона")
		t.FailNow()
	}
}

func TestShop_FindByCustomerAddress(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Комсомольская, 15, 94", 89871301212)
	c2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)
	c3 := NewCustomer("Петров Пётр", "Россия, Брянск, Брянского Фронта, 30, 1", 89098098899)
	c4 := NewCustomer("Иванов Иван", "Россия, Екатеринбург, Советская, 16, 34", 89995557665)

	s.AddCustomer(c1)
	s.AddCustomer(c2)
	s.AddCustomer(c3)
	s.AddCustomer(c4)

	f := s.FindByCustomerAddress(c1.Address)

	if !reflect.DeepEqual(f, []Customer{c1, c2}) {
		t.Error("Ошибка поиска клиента по адресу")
		t.FailNow()
	}
}

func TestShop_FindByCustomerOrderID(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Комсомольская, 15, 94", 89871301212)
	c2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)
	c3 := NewCustomer("Петров Пётр", "Россия, Брянск, Брянского Фронта, 30, 1", 89098098899)
	c4 := NewCustomer("Иванов Иван", "Россия, Екатеринбург, Советская, 16, 34", 89995557665)

	s.AddCustomer(c1)
	s.AddCustomer(c2)
	s.AddCustomer(c3)
	s.AddCustomer(c4)

	p1 := NewProduct("Настольная лампа", "Материал: платик", 1342.31)
	p2 := NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)

	s.AddProduct(p1)
	s.AddProduct(p2)

	newOr := NewOder(c3.ID, []Product{p1, p2})

	s.AddOrder(newOr)

	f := s.FindByCustomerOrderID(newOr.Number)
	if f.ID != c3.ID {
		t.Error("Ошибка поиска клиента по номеру заказа")
		t.FailNow()
	}
}

func TestNewOder(t *testing.T) {
	p1 := NewProduct("Настольная лампа", "Материал: платик", 1342.31)
	p2 := NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)

	or := Order{
		Number:     1,
		ProductsID: []uint{1, 2},
		Price:      p1.Price + p2.Price,
		CustomerID: 1,
	}

	order := NewOder(or.CustomerID, []Product{p1, p2})

	if !reflect.DeepEqual(order, or) {
		t.Error("Ошибка создания заказа")
		t.FailNow()
	}

	if !reflect.DeepEqual(order.ProductsID, []uint{p1.ID, p2.ID}) {
		t.Error("Ошибка создания заказа")
		t.FailNow()
	}
}

func TestShop_AddOrder(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Комсомольская, 15, 94", 89871301212)

	s.AddCustomer(c1)

	or1 := NewOder(1,
		[]Product{NewProduct("Настольная лампа", "Материал: платик", 1342.31),
			NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)})
	or2 := NewOder(1, []Product{NewProduct("Кухонный нож",
		"Материал лезвия: сталь; Материал рукояти: пластик; Цвет: чёрный", 2659.50)})

	s.AddOrder(or1)
	s.AddOrder(or2)

	if len(s.orders) != 2 {
		t.Error("Ошибка добавления заказа")
		t.FailNow()
	}

	c := s.FindByCustomerID(c1.ID)

	if !reflect.DeepEqual(c.OrdersID, []uint{1, 2}) {
		t.Error("Ошибка добавления заказа")
		t.FailNow()
	}
}

func TestShop_UpdateOrder(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Комсомольская, 15, 94", 89871301212)
	c2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)

	s.AddCustomer(c1)
	s.AddCustomer(c2)

	or1 := NewOder(1,
		[]Product{NewProduct("Настольная лампа", "Материал: платик", 1342.31),
			NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)})
	or2 := NewOder(1, []Product{NewProduct("Кухонный нож",
		"Материал лезвия: сталь; Материал рукояти: пластик; Цвет: чёрный", 2659.50)})

	s.AddOrder(or1)
	s.AddOrder(or2)

	or1.CustomerID = 2
	s.UpdateOrder(or1)

	o := s.FindByOrderNumber(1)

	if !reflect.DeepEqual(o, or1) {
		t.Error("Ошбика обновления данных заказа")
		t.FailNow()
	}

	cust1 := s.FindByCustomerID(1)
	cust2 := s.FindByCustomerID(2)

	if !reflect.DeepEqual(cust1.OrdersID, []uint{2}) {
		t.Error("Ошибка обновления даных заказа")
		t.FailNow()
	}

	if !reflect.DeepEqual(cust2.OrdersID, []uint{1}) {
		t.Error("Ошибка обновления даных заказа")
		t.FailNow()
	}
}

func TestShop_DeleteOrder(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	c1 := NewCustomer("Нурдавлетова Алёна", "Россия, Уфа, Комсомольская, 15, 94", 89871301212)
	c2 := NewCustomer("Чеботарь Сергей", "Россия, Уфа, Комсомольская, 15, 94", 89050038097)

	s.AddCustomer(c1)
	s.AddCustomer(c2)

	or1 := NewOder(1,
		[]Product{NewProduct("Настольная лампа", "Материал: платик", 1342.31),
			NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)})
	or2 := NewOder(1, []Product{NewProduct("Кухонный нож",
		"Материал лезвия: сталь; Материал рукояти: пластик; Цвет: чёрный", 2659.50)})

	s.AddOrder(or1)
	s.AddOrder(or2)

	s.DeleteOrder(or2)

	if len(s.orders) != 1 {
		t.Error("Ошибка удаления заказа")
		t.FailNow()
	}

	cust1 := s.FindByCustomerID(c1.ID)

	if !reflect.DeepEqual(cust1.OrdersID, []uint{1}) {
		t.Error("Ошибка удаления заказа")
		t.FailNow()
	}
}
func TestShop_FindByOrderProductsID(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	or1 := NewOder(1,
		[]Product{NewProduct("Настольная лампа", "Материал: платик", 1342.31),
			NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)})
	or2 := NewOder(1, []Product{NewProduct("Кухонный нож",
		"Материал лезвия: сталь; Материал рукояти: пластик; Цвет: чёрный", 2659.50)})

	s.AddOrder(or1)
	s.AddOrder(or2)

	order := s.FindByOrderProductsID(3)

	if !reflect.DeepEqual(order, or2) {
		t.Error("Ошибка поиска заказа по ID продукта")
		t.FailNow()
	}
}

func TestShop_FindByOrderCustomerID(t *testing.T) {
	s := Shop{
		Name:       "Удобство и комфорт",
		DomainName: "udobstvo.com",
		products:   make([]Product, 0),
		customers:  make([]Customer, 0),
		orders:     make([]Order, 0),
	}

	or1 := NewOder(1,
		[]Product{NewProduct("Настольная лампа", "Материал: платик", 1342.31),
			NewProduct("Плед", "Состав: шерсть, Цвет: серый, Размер: 2.0x1.5", 3500)})
	or2 := NewOder(2, []Product{NewProduct("Кухонный нож",
		"Материал лезвия: сталь; Материал рукояти: пластик; Цвет: чёрный", 2659.50)})

	s.AddOrder(or1)
	s.AddOrder(or2)

	order := s.FindByOrderCustomerID(2)

	if !reflect.DeepEqual(order, or2) {
		t.Error("Ошибка поиска заказа по ID клиента")
		t.FailNow()
	}
}
