package controlllerpage

import (
	modelpage "github.com/nguyenduclam1711/investment-management/app/model/page"
)

func renderIndexPage(renderPage renderPageFunc) {
	renderPage("", "index", modelpage.Index{
		Title: "Index",
	})
}
