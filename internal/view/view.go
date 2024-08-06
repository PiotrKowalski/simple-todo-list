package view

import (
	"errors"
	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"log"
	"net/http"
	auth2 "simple-todo-list/pkg/auth"
	"time"
)

//	func createHandler(title string, body g.Node) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			return Page(title, c.Request().URL.Path, body).Render(c.Response())
//		}
//	}
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

		auth, err := c.Cookie("Authorization")
		if !errors.Is(err, http.ErrNoCookie) {
			log.Println("auth vaule", auth.Value)
			claims, _ := auth2.ExtractClaims(auth.Value)
			log.Println(claims)
			return Page(title, c.Request().URL.Path, body, user{username: claims.Username}).Render(c.Response())

			return err
		}

		return Page(title, c.Request().URL.Path, body, user{}).Render(c.Response())
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

type user struct {
	username string
	role     string
}

func Page(title, path string, body g.Node, user user) g.Node {
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
			navbar(user, path, []pageLink{
				{Path: "/", Name: "Home"},
				{Path: "/todo-lists", Name: "Todo lists"},
			}),
			Container(
				MainContent(body),
				PageFooter(),
			),
		},
	})
}

func Container(children ...g.Node) g.Node {
	return Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}

func MainContent(children ...g.Node) g.Node {
	//return Div(Class("prose"),
	//	g.Group(children))
	return g.Group(children)
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
