package sample

import (
	"log"

	"livecode3/model"
	"livecode3/repo"
	"livecode3/usecase"
	"livecode3/utils"

	"gorm.io/gorm"
)

func Nomor01sampai07(db *gorm.DB) {
	/*
		1. Master Menu
	*/
	menuRepo := repo.NewMenuRepo(db)

	menu01 := model.Menu{
		MenuName: "Nasi Goreng",
	}
	menu02 := model.Menu{
		MenuName: "Mie Goreng",
	}

	// Create
	err := menuRepo.Create(&menu01)
	err = menuRepo.Create(&menu02)

	// Read
	menus, _ := menuRepo.FindAll()
	for _, m := range menus {
		log.Println(m.ToString())
	}

	menus, _ = menuRepo.FindAllBy(map[string]interface{}{
		"menu_name": "Nasi Goreng",
	})
	for _, m := range menus {
		log.Println(m.ToString())
	}

	// Update
	err = menuRepo.Update(&menu02, map[string]interface{}{
		"menu_name": "Mie Goreng Spesial",
	})

	// Delete
	err = menuRepo.Delete(&menu01)
	utils.IsError(err)

	/*
		2. Master Menu Price
	*/
	menuPriceRepo := repo.NewMenuPriceRepo(db)

	menuPrice01 := model.MenuPrice{
		Menu:  menu01,
		Price: 20000,
	}
	menuPrice02 := model.MenuPrice{
		Menu:  menu02,
		Price: 15000,
	}

	// Create
	err = menuPriceRepo.Create(&menuPrice01)
	err = menuPriceRepo.Create(&menuPrice02)

	// Read
	menuPrices, _ := menuPriceRepo.FindAll()
	for _, m := range menuPrices {
		log.Println(m.ToString())
	}

	menuPrices, _ = menuPriceRepo.FindAllBy(map[string]interface{}{
		"price": 20000,
	})
	for _, m := range menuPrices {
		log.Println(m.ToString())
	}

	// Update
	err = menuPriceRepo.Update(&menuPrice01, map[string]interface{}{
		"price": 18000,
	})

	// Delete
	err = menuPriceRepo.Delete(&menuPrice01)
	utils.IsError(err)

	/*
		3. Master Menu Table
	*/
	tableRepo := repo.NewTableRepo(db)

	table01 := model.Table{
		TableDescription: "hello",
		IsAvailable:      false,
	}
	table02 := model.Table{
		TableDescription: "hai",
		IsAvailable:      false,
	}

	// Create
	err = tableRepo.Create(&table01)
	err = tableRepo.Create(&table02)

	// Read
	tables, _ := tableRepo.FindAll()
	for _, m := range tables {
		log.Println(m.ToString())
	}

	tables, _ = tableRepo.FindAllBy(map[string]interface{}{
		"is_available": false,
	})
	for _, m := range tables {
		log.Println(m.ToString())
	}

	// Update
	err = tableRepo.Update(&table01, map[string]interface{}{
		"table_description": "uenak tenan",
	})

	// Delete
	err = tableRepo.Delete(&table01)
	utils.IsError(err)

	/*
		4. Master Trans Type
	*/
	transRepo := repo.NewTransRepo(db)

	transType01 := model.TransType{
		Id:          "DI",
		Description: "Dine In",
	}
	transType02 := model.TransType{
		Id:          "TA",
		Description: "Take Away",
	}
	transType03 := model.TransType{
		Id:          "DEL",
		Description: "Test untuk didelete",
	}

	err = transRepo.Create(&transType01)
	err = transRepo.Create(&transType02)
	err = transRepo.Create(&transType03)

	// Read
	transTypes, _ := transRepo.FindAll()
	for _, m := range transTypes {
		log.Println(m.ToString())
	}

	// Update
	err = transRepo.Update(&transType01, map[string]interface{}{
		"id":          "EI",
		"description": "Eat In",
	})

	// Delete
	err = transRepo.Delete(&transType03)
	utils.IsError(err)

	/*
		5. Master Discount
	*/
	discRepo := repo.NewDiscountRepo(db)
	customer01 := model.Customer{
		CustomerName:  "Peter",
		MobilePhoneNo: "097757859678",
		IsMember:      true,
	}

	disc01 := model.Discount{
		Description: "customer01",
		Pct:         0.1,
		Customers: []*model.Customer{
			&customer01,
		},
	}

	// Create
	err = discRepo.Create(&disc01)

	// Read
	discs, _ := discRepo.FindAll()
	for _, m := range discs {
		log.Println(m.ToString())
	}

	// Update
	err = discRepo.Update(&disc01, map[string]interface{}{
		"description": "customer01 naikin dics",
		"pct":         0.2,
	})

	// Delete
	err = discRepo.Delete(&disc01)
	utils.IsError(err)

	/*
		Customer Repo Test
	*/
	cstRepo := repo.NewCustomerRepo(db)
	customer02 := model.Customer{
		CustomerName:  "Bruce",
		MobilePhoneNo: "478923789",
		IsMember:      false,
	}
	err = cstRepo.Create(&customer02)

	err = cstRepo.Update(&customer02, map[string]interface{}{
		"mobile_phone_no": "911",
	})

	/*
		6. Melakukan customer registration
	*/
	custUseCase := usecase.NewCustomerUseCase(cstRepo)
	th, err := custUseCase.CustomerRegistration("Thomas", "3124124")

	/*
		7. Melakukan aktivasi member customer yang sudah terdaftar sekaligus memberikan privilege discount
	*/

	// (id dari m_customer), (persen discount), (repo discount)
	err = cstRepo.ActivateMember("3", 0.1, discRepo)
	utils.IsError(err)

	log.Println(th.ToString())
}

func Nomor08(db *gorm.DB) {
	/*
		8. Melakukan transaksi penjualan dengan validasi apabila meja sudah dipakai tidak bisa dibuat bill
	*/

	// buat trans type
	transRepo := repo.NewTransRepo(db)

	transTypeDI := model.TransType{
		Id:          "DI",
		Description: "Dine In",
	}
	transTypeTA := model.TransType{
		Id:          "TA",
		Description: "Take Away",
	}

	err := transRepo.Create(&transTypeDI)
	utils.IsError(err)
	err = transRepo.Create(&transTypeTA)
	utils.IsError(err)

	custRepo := repo.NewCustomerRepo(db)
	tableRepo := repo.NewTableRepo(db)

	billRepo := repo.NewBillRepo(db)

	// register customer baru
	custUseCase := usecase.NewCustomerUseCase(custRepo)
	thomas, err := custUseCase.CustomerRegistration("Thomas Wayne", "911")
	utils.IsError(err)

	// buat meja
	table01 := model.Table{
		TableDescription: "Thomas' table",
		IsAvailable:      true, // kalau dibuat false maka bill tidak dapat dibuat pada t_bill
	}

	// memasukan meja pada repo
	err = tableRepo.Create(&table01)
	utils.IsError(err)

	// buat bill

	thomasBill := model.Bill{
		CustomerId:  thomas.Model.ID,
		Customer:    thomas,
		TableId:     table01.Model.ID,
		Table:       table01, // memasukan meja thomas ke bruce yang sudah terisi
		TransTypeId: transTypeDI.Id,
		TransType:   transTypeDI,
	}

	err = billRepo.Create(&thomasBill) // bill thomas tidak terbuat karena nempatin meja thomas
	utils.IsError(err)

}

func Nomor09(db *gorm.DB) {
	// buat trans type
	transRepo := repo.NewTransRepo(db)

	transTypeDI := model.TransType{
		Id:          "DI",
		Description: "Dine In",
	}
	transTypeTA := model.TransType{
		Id:          "TA",
		Description: "Take Away",
	}
	err := transRepo.Create(&transTypeDI)
	utils.IsError(err)
	err = transRepo.Create(&transTypeTA)
	utils.IsError(err)

	// discRepo := repo.DiscountRepo(db)

	// billRepo := repo.NewBillRepo(db)

	/*
		9. Mencetak bill berdasarkan bill_id, sekaligus melakukan update meja menjadi available
	*/

	billDetailRepo := repo.NewBillDetailRepo(db)
	billRepo := repo.NewBillRepo(db)
	custRepo := repo.NewCustomerRepo(db)
	discRepo := repo.NewDiscountRepo(db)
	menuPriceRepo := repo.NewMenuPriceRepo(db)
	menuRepo := repo.NewMenuRepo(db)
	tableRepo := repo.NewTableRepo(db)

	//buat meja untuk bruce
	table02 := model.Table{
		TableDescription: "Bruce's table",
		IsAvailable:      true,
	}
	err = tableRepo.Create(&table02)
	utils.IsError(err)

	// Registrasi customer
	custUseCase := usecase.NewCustomerUseCase(custRepo)
	bruce, err := custUseCase.CustomerRegistration("Bruce Wayne", "119")
	utils.IsError(err)

	// buat bikin bill bruce
	bruceBill := model.Bill{
		CustomerId:  bruce.Model.ID,
		Customer:    bruce,
		TableId:     table02.Model.ID,
		Table:       table02,
		TransTypeId: transTypeDI.Id,
		TransType:   transTypeDI,
	}
	err = billRepo.Create(&bruceBill) // bill bruce akan terbuat pada table t_bill
	utils.IsError(err)

	// loads, err := billRepo.FindAllWithPreload("Customer")

	// for _, load := range loads {
	// 	log.Println(load.ToString())
	// }

	// buat menu
	menu01 := model.Menu{
		MenuName: "Nasi Goreng Spesial",
	}
	menu02 := model.Menu{
		MenuName: "Mie Goreng Sedikit Spesial",
	}
	err = menuRepo.Create(&menu01)
	err = menuRepo.Create(&menu02)

	// buat harga menu
	menuPrice01 := model.MenuPrice{
		Menu:  menu01,
		Price: 20000,
	}
	menuPrice02 := model.MenuPrice{
		Menu:  menu02,
		Price: 15000,
	}
	err = menuPriceRepo.Create(&menuPrice01)
	err = menuPriceRepo.Create(&menuPrice02)

	// bikin discount
	disc01 := model.Discount{
		Description: "bruce discount",
		Pct:         0.1,
		Customers: []*model.Customer{
			&bruce,
		},
	}
	custRepo.Update(&bruce, map[string]interface{}{
		"is_member": true, // update jadi member
	})
	err = discRepo.Create(&disc01)

	// bikin bill detail
	bruceBillDetail01 := model.BillDetail{
		Bill:      bruceBill,
		MenuPrice: menuPrice01,
		Quantity:  2,
	}

	bruceBillDetail02 := model.BillDetail{
		Bill:      bruceBill,
		MenuPrice: menuPrice02,
		Quantity:  3,
	}
	err = billDetailRepo.Create(&bruceBillDetail01)
	err = billDetailRepo.Create(&bruceBillDetail02)

	billDetails, _ := billDetailRepo.FindAll()
	log.Println(billDetails)

}

func Nomor10(db *gorm.DB) {
	var result int
	// stmt := "SELECT CASE WHEN mc.is_member = true THEN SUM(tbd.quantity*mmp.price)*(1-md.pct) ELSE SUM(tbd.quantity*mmp.price) END FROM t_bill tb JOIN t_bill_detail tbd ON tbd.bill_id = tb.id JOIN m_menu_price mmp ON mmp.id = tbd.menu_price_id JOIN m_customer mc ON tb.customer_id = mc.id JOIN m_customer_discounts mcd ON mcd.customer_id = mc.id JOIN m_discount md ON md.id = mcd.discount_id GROUP BY tb.customer_id, mc.is_member, md.pct"
	stmt := "SELECT SUM(tbd.quantity*mmp.price) FROM t_bill tb JOIN t_bill_detail tbd ON tbd.bill_id = tb.id JOIN m_menu_price mmp ON mmp.id = tbd.menu_price_id GROUP BY tb.trans_date"
	db.Raw(stmt).Scan(&result)
	log.Println("Total penjualan harian: ", result)
}
