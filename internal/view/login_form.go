package view

import (
	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	. "github.com/maragudk/gomponents/html"
	"net/http"
	"simple-todo-list/internal/dtos"
)

func createLoginPageHandler(app app) echo.HandlerFunc {
	return func(c echo.Context) error {

		title, body := loginPage()
		return Page(title, c.Request().URL.Path, body, user{}).Render(c.Response())
	}
}

func createLoginPageActionHandler(app app) echo.HandlerFunc {
	return func(c echo.Context) error {
		params, err := c.FormParams()
		if err != nil {
			return err
		}
		login, err := app.Login(c.Request().Context(), dtos.LoginInput{
			Username: params.Get("username"),
			Password: params.Get("password"),
		})
		if err != nil {
			return loginError(err.Error()).Render(c.Response())
		}

		c.SetCookie(&http.Cookie{Name: "Authorization", Value: login.JWT})
		hxhttp.SetRedirect(c.Response().Header(), "/")
		//err = c.Redirect(http.StatusSeeOther, "/")
		//if err != nil {
		//	return err
		//}

		return nil
	}
}

func loginPage() (string, g.Node) {
	return "Login Page",
		Div(Class("flex min-h-full flex-col justify-center px-6 py-12 lg:px-8"),
			Div(Class("sm:mx-auto sm:w-full sm:max-w-sm"),
				Img(Class("x-auto h-10 w-auto"), Src("https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"), Alt("Your Company")),
				H2(Class("mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900"), g.Text("Sign in to your account")),
			),
			Div(Class("mt-10 sm:mx-auto sm:w-full sm:max-w-sm"),
				FormEl(Class("space-y-6"), Action("/login"), Method("POST"), hx.Boost("true"), hx.Target("#error"), hx.Swap("innerHTML"),
					Div(
						Label(For("username"), Class("block text-sm font-medium leading-6 text-gray-900"), g.Text("Username")),
						Div(Class("mt-2"), Input(ID("username"), Name("username"), Type("text"), AutoComplete("username"), Class("block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"), Required())),
					),
					Div(
						Div(Class("flex items-center justify-between"),
							Label(For("password"), Class("block text-sm font-medium leading-6 text-gray-900"), g.Text("Password")),
							Div(Class("text-sm"),
								A(Href("#"), Class("font-semibold text-indigo-600 hover:text-indigo-500"), g.Text("Forgot password?"))),
						),
						Div(Class("mt-2"),
							Input(ID("password"), Name("password"), Type("password"), AutoComplete("current-password"), Required(), Class("block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6")),
						),
					),
					Div(
						Button(Type("submit"), g.Text("Sign in"), Class("flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600")),
					),
				),

				Div(Class(""), ID("error"), hx.Trigger("login-failed")),

				P(Class("mt-10 text-center text-sm text-gray-500"),
					g.Text("Not a member?  "),
					A(Href("#"), Class("font-semibold leading-6 text-indigo-600 hover:text-indigo-500"), g.Text("Sign up")),
				),
			),
		)

	//Div(
	//
	//	H1(g.Text("Welcome to this example page")),
	//	P(g.Text("I hope it will make you happy. 😄 It's using TailwindCSS for styling.")),
	//	partial(now),
	//	FormEl(Method("post"), Action("/"), hx.Boost("true"), hx.Target("#partial"), hx.Swap("outerHTML"),
	//		Button(Type("submit"), g.Text(`Update time`),
	//			Class("rounded-md border border-transparent bg-orange-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:ring-offset-2"),
	//		),
	//	),
	//)
}

func loginError(text string) g.Node {
	return Div(Class("bg-orange-100 border-l-4 border-orange-500 text-orange-700 p-4"), Role("alert"),
		P(Class("font-bold"), g.Text("Login Error")),
		P(g.Text(text)))

}

//<div class="bg-orange-100 border-l-4 border-orange-500 text-orange-700 p-4" role="alert">
//<p class="font-bold">Be Warned</p>
//<p>Something not ideal might be happening.</p>
//</div>

//<div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
//<form class="space-y-6" action="#" method="POST">
//<div>
//<label for="email" class="block text-sm font-medium leading-6 text-gray-900">Email address</label>
//<div class="mt-2">
//<input id="email" name="email" type="email" autocomplete="email" required class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6">
//</div>
//</div>
//
//<div>
//<div class="flex items-center justify-between">
//<label for="password" class="block text-sm font-medium leading-6 text-gray-900">Password</label>
//<div class="text-sm">
//<a href="#" class="font-semibold text-indigo-600 hover:text-indigo-500">Forgot password?</a>
//</div>
//</div>
//<div class="mt-2">
//<input id="password" name="password" type="password" autocomplete="current-password" required class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6">
//</div>
//</div>
//
//<div>
//<button type="submit" class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Sign in</button>
//</div>
//</form>
//
//<p class="mt-10 text-center text-sm text-gray-500">
//Not a member?
//<a href="#" class="font-semibold leading-6 text-indigo-600 hover:text-indigo-500">Start a 14 day free trial</a>
//</p>
//</div>
