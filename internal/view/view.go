package view

import (
	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"time"
)

func createHandler(title string, body g.Node) echo.HandlerFunc {
	return func(c echo.Context) error {
		return Page(title, c.Request().URL.Path, body).Render(c.Response())
	}
}
func createUpdateTimeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		if hxhttp.IsBoosted(c.Request().Header) {
			now := time.Now()
			return partial(now).Render(c.Response())
		}
		return nil
	}
}

func createIndexPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		now := time.Now()
		title, body := indexPage(now)
		return Page(title, c.Request().URL.Path, body).Render(c.Response())
	}
}

func indexPage(now time.Time) (string, g.Node) {
	return "Welcome!", Div(
		H1(g.Text("Welcome to this example page")),
		P(g.Text("I hope it will make you happy. 😄 It's using TailwindCSS for styling.")),
		partial(now),
		FormEl(Method("post"), Action("/"), hx.Boost("true"), hx.Target("#partial"), hx.Swap("outerHTML"),
			Button(Type("submit"), g.Text(`Update time`),
				Class("rounded-md border border-transparent bg-orange-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:ring-offset-2"),
			),
		),
	)
}

func Page(title, path string, body g.Node) g.Node {
	// HTML5 boilerplate document
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{

			Script(Src("https://cdn.tailwindcss.com?plugins=typography,forms")),
			Script(Src("https://unpkg.com/htmx.org")),
		},
		Body: []g.Node{
			Class("h-full"),
			Navbar(path, []PageLink{
				{Path: "/contact", Name: "Contact"},
				{Path: "/about", Name: "About"},
			}),
			Container(
				MainConte(body),
				PageFooter(),
			),
		},
	})
}

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) g.Node {
	return Nav(Class("bg-gray-700 mb-4"),
		Container(
			Div(Class("flex items-center space-x-4 h-16"),
				NavbarLink("/", "Home", currentPath == "/"),

				// We can Map custom slices to Nodes
				g.Group(g.Map(links, func(l PageLink) g.Node {
					return NavbarLink(l.Path, l.Name, currentPath == l.Path)
				})),
			),
		),
	)
}

// NavbarLink is a link in the Navbar.
func NavbarLink(path, text string, active bool) g.Node {
	return A(Href(path), g.Text(text),
		// Apply CSS classes conditionally
		c.Classes{
			"px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700": true,
			"text-white bg-gray-900":                           active,
			"text-gray-300 hover:text-white hover:bg-gray-700": !active,
		},
	)
}

func Container(children ...g.Node) g.Node {
	return Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}

func MainConte(children ...g.Node) g.Node {
	return Div(Class("prose"),
		g.Group(children))
}

func PageFooter() g.Node {
	return Footer(Class("prose prose-sm prose-indigo"),
		P(
			// We can use string interpolation directly, like fmt.Sprintf.
			g.Textf("Rendered %v. ", time.Now().Format(time.RFC3339)),

			// Conditional inclusion
			g.If(time.Now().Second()%2 == 0, g.Text("It's an even second.")),
			g.If(time.Now().Second()%2 == 1, g.Text("It's an odd second.")),
		),

		P(A(Href("https://www.gomponents.com"), g.Text("gomponents"))),
	)
}

const timeFormat = "15:04:05"

func partial(now time.Time) g.Node {
	return P(ID("partial"), g.Textf(`Time was last updated at %v.`, now.Format(timeFormat)))
}
