package controlllerpage

import (
	modelpage "github.com/nguyenduclam1711/investment-management/app/model/page"
)

func Index(renderPage RenderPage) {
	renderPage("", "index", modelpage.Index{
		Title: "heheehhehe",
	})
}
