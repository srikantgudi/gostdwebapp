package db

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Category struct {
	Id    int    `db:"category_id"`
	Name  string `db:"category_name"`
	Descr string `db:"description"`
}

type Product struct {
	Id         int     `db:"product_id"`
	Name       string  `db:"product_name"`
	Qtyperunit string  `db:"quantity_per_unit"`
	Price      float64 `db:"unit_price"`
	Rorlevel   int     `db:"reorder_level"`
}

type Customer struct {
	Id      string         `db:"customer_id"`
	Name    string         `db:"company_name"`
	City    string         `db:"city"`
	Zipcode sql.NullString `db:"postal_code"`
	Country string         `db:"country"`
}

type Order struct {
	Orderid   int    `db:"order_id"`
	Orderdate string `db:"order_date"`
	Shipdate  string `db:"shipped_date"`
}

type Orderdetail struct {
	Orderid     int     `db:"od_order_id"`
	Productname string  `db:"product_name"`
	Quantity    int     `db:"quantity"`
	Unitprice   float64 `db:"od_unit_price"`
}

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("postgres", "host=localhost user=postgres password=toor port=5432 dbname=northwind")
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
}

func GetCategories() []Category {
	var data []Category
	db.Select(&data, "Select category_id, category_name, description From categories")
	return data
}

func GetProducts() []Product {
	var data []Product
	db.Select(&data, "Select product_id, product_name, quantity_per_unit, unit_price, reorder_level From products")
	return data
}

func GetCatProducts(catid string) []Product {
	var data []Product
	err := db.Select(&data, "Select product_id, product_name, quantity_per_unit, unit_price, reorder_level From products Where pr_category_id=$1", catid)
	if err != nil {
		log.Fatalln("Error getting category-products: ", err.Error())
	}
	return data
}
func GetCustomers() []Customer {
	var data []Customer
	db.Select(&data, "Select customer_id, company_name, city, postal_code, country From customers")
	return data
}

func GetOrders() []Order {
	var data []Order
	db.Select(&data, "Select order_id, to_char(order_date,'DD-MON-YYYY') order_date, to_char(shipped_date,'DD-MON-YYYY') shipped_date From orders")
	return data
}

func GetOrderdetails() []Orderdetail {
	var data []Orderdetail
	err := db.Select(&data, "Select od_order_id, product_name, quantity, od_unit_price From order_details Join products On product_id = od_product_id")
	if err != nil {
		log.Fatalln("error order details: ", err.Error())
	}
	return data
}

func GetOrdersByCustomer(custid string) []Order {
	var data []Order
	db.Select(&data, "Select order_id, to_char(order_date,'DD-MON-YYYY') order_date, to_char(shipped_date,'DD-MON-YYYY') shipped_date From orders Where ord_customer_id = $1", custid)
	return data
}

func GetOrderdetailsByOrder(oid string) []Orderdetail {
	var data []Orderdetail
	err := db.Select(&data, "Select od_order_id, product_name, quantity, od_unit_price From order_details Join products On product_id = od_product_id Where od_order_id = $1", oid)
	if err != nil {
		log.Fatalln("error: get order details by order: ", err.Error())
	}
	return data
}
