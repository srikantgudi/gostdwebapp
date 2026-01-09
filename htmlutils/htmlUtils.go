package htmlutils

import (
	"fmt"
	"net/http"
	"strings"
)

func HHeader() string {
	return `
	<!DOCTYPE>
	<html>
		<head>
			<script src="https://unpkg.com/htmx.org"></script>
			<style>
				* {
					font-family:helvetica;
					font-weight: 400;
				}
				body {
					margin: 1vh 5vw;
				}
				thead {
					position: sticky;
					top: 0;
					background-color: #303030;
					color: whitesmoke;
				}
				td {
					padding: 0.25rem 0.5rem;
				}
				.link {
					cursor:pointer;
				}
				.grid12 {
					display:grid;
					grid-template-columns:1fr 2fr;
					gap:1rem'
				}
				.grid13 {
					display:grid;
					grid-template-columns:1fr 3fr;
					gap:1rem'
				}
				nav {
					display: flex;
					align-items:center;
					gap: 0.5rem;
					margin-bottom: 1.5rem;
					justify-content: center;
				}
				.apptitle {
					font-size: 2rem;
					font-family: serif;
					font-weight: 600;
					margin-right: 2rem;
				}
			</style
		</head>
		<body>
		<nav>
			<div class='apptitle'>Template-less GoLang Demo App</div>
			<a href="/">Home</a>
			<a href="/categories">Categories</a>
			<a href="/products">Products</a>
			<a href="/customers">Customers</a>
		</nav>
		`
}

func HFooter() string {
	var b strings.Builder
	b.WriteString(`
		</body>
	</html>
	`)
	return b.String()
}

func SetHtml(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/html")
}

func Thead(cols ...string) string {
	var b strings.Builder
	b.WriteString("<thead><tr>")
	for _, col := range cols {
		b.WriteString(fmt.Sprintf(`%s`, col))
	}
	b.WriteString("</tr></thead>")
	return b.String()
}

func HTable(head, body string) string {
	return fmt.Sprintf("<table width='100%%'>%v %v</table>", head, body)
}
