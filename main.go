package main

import (
	"fmt"
	"gostdwebapp/db"
	hu "gostdwebapp/htmlutils"
	"log"
	"net/http"
	"strings"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	hu.SetHtml(w)
	body := `
	<div style="text-align:center;font-size:1.2rem;">
	<h2>@Home</h2>
	<p>This concept of template-less rendering was thought of to avoid the hassles of installing and configuring external templating engine</p>
	<p><strong>No templates. Pure performance.</strong> Direct HTML streaming skips parse/render overhead.</p>
	<p>Created utility package to eliminate repetition; generate perfect tables from Go structs instantly.</p>
	</div>
	`
	fmt.Fprintf(w, "%v <div>%v</div> %v", hu.HHeader(), body, hu.HFooter())
}

func categoriesPage(w http.ResponseWriter, r *http.Request) {
	categories := db.GetCategories()
	var data strings.Builder
	for _, c := range categories {
		data.WriteString(fmt.Sprintf("<tr class='link' hx-get='category/%d/products' hx-target='#products'><td>%v</td><td>%v</td></tr>", c.Id, c.Name, c.Descr))
	}
	htable := hu.HTable(hu.Thead("<td>Product name</td>", "<td>Description</td>"), data.String())
	hu.SetHtml(w)
	fmt.Fprintf(w, "%v <div class='grid12'><div><h2>Categories</h2><div>%v</div></div><div id='products'>-- products --</div></div>%v", hu.HHeader(), htable, hu.HFooter())
}

func productsTable(products []db.Product) string {
	var data strings.Builder
	for _, p := range products {
		data.WriteString(fmt.Sprintf("<tr><td>%v</td><td>%v</td><td align='right'>%.2f</td><td align='right'>%d</td></tr>", p.Name, p.Qtyperunit, p.Price, p.Rorlevel))
	}
	return hu.HTable(hu.Thead("<td width='40%'>Product name</td>", "<td width='30%'>Quantity per unit</td>", "<td align='right'>Price</td>", "<td align='right'>Reorder level</td>"), data.String())
}

func productsPage(w http.ResponseWriter, r *http.Request) {
	products := db.GetProducts()
	htable := productsTable(products)
	hu.SetHtml(w)
	fmt.Fprintf(w, "%v <h2>Products</h2><div style='height:70vh;overflow-y:auto'>%v</div>%v", hu.HHeader(), htable, hu.HFooter())
}

func categoryProductsPage(w http.ResponseWriter, r *http.Request) {
	catid := r.PathValue("catid")
	products := db.GetCatProducts(catid)
	htable := productsTable(products)
	fmt.Fprintf(w, "<div><h2>Products | <small>%d fetched</small></h2><div style='height:70vh;overflow-y:auto'>%v</div></div>", len(products), htable)
}

func customersPage(w http.ResponseWriter, r *http.Request) {
	customers := db.GetCustomers()

	var data strings.Builder
	for _, c := range customers {
		data.WriteString(fmt.Sprintf("<tr class='link' hx-get='/customer/%v/orders' hx-target='#orders'><td title='%v'>%v</td></tr>", c.Id, c.City+" "+c.Zipcode.String+" "+c.Country, c.Name))
	}

	htable := hu.HTable(hu.Thead("<td>Customer name</td>"), data.String())
	hu.SetHtml(w)
	fmt.Fprintf(w, `%v <div class='grid13'>
		<div>
			<h2>Customers</h2>
			<div style='height:70vh;overflow-y:auto'>%v</div>
		</div>
		<div id='orders'></div>
		</div>
		%v`, hu.HHeader(), htable, hu.HFooter())
}

func customerOrdersPage(w http.ResponseWriter, r *http.Request) {
	custid := r.PathValue("custid")
	orders := db.GetOrdersByCustomer(custid)
	var data strings.Builder
	for _, o := range orders {
		data.WriteString(fmt.Sprintf("<tr class='link' hx-get='/order/%v/details' hx-target='#orderdetails'><td>%v</td><td>%v</td></tr>", o.Orderid, o.Orderid, o.Orderdate))
	}
	htable := hu.HTable(hu.Thead("<td>Order-Id</td>", "<td>Order date</td>"), data.String())
	fmt.Fprintf(w, `<div class='grid13'>
		<div>
			<h2>Orders</h2>
			<div style='height:70vh;overflow-y:auto'>%v</div>
		</div>
		<div id='orderdetails'></div>
	</div>
	`, htable)
}

func orderDetailsPage(w http.ResponseWriter, r *http.Request) {
	orderid := r.PathValue("orderid")
	orderdetails := db.GetOrderdetailsByOrder(orderid)
	var data strings.Builder
	for _, od := range orderdetails {
		data.WriteString(fmt.Sprintf("<tr><td>%v</td><td align='right'>%v</td><td align='right'>%.2f</td><td align='right'>%.2f</td></tr>", od.Productname, od.Quantity, od.Unitprice, (float64(od.Quantity) * od.Unitprice)))
	}
	htable := hu.HTable(hu.Thead("<td width='40%%'>Product name</td>", "<td align='right'>Quantity</td>", "<td align='right'>Unit Price</td>", "<td align='right'>Line total</td>"), data.String())
	fmt.Fprintf(w, `<div>
			<h2>Order Details</h2>
			<div style='height:70vh;overflow-y:auto'>%v</div>
		</div>
	`, htable)
}

func main() {
	app := http.NewServeMux()

	app.HandleFunc("GET /", rootPage)
	app.HandleFunc("GET /categories", categoriesPage)
	app.HandleFunc("GET /products", productsPage)
	app.HandleFunc("GET /category/{catid}/products", categoryProductsPage)
	app.HandleFunc("GET /customers", customersPage)
	app.HandleFunc("GET /customer/{custid}/orders", customerOrdersPage)
	app.HandleFunc("GET /order/{orderid}/details", orderDetailsPage)

	log.Println("http://localhost:8090")
	http.ListenAndServe(":8090", app)
}
