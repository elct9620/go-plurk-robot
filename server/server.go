package server

import (
	"github.com/elct9620/go-plurk-robot/db"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"html/template"
	"io"
	"path/filepath"
	"runtime"
)

var (
	cookie  *sessions.CookieStore
	appRoot string
)

type Renderer struct {
	templates *template.Template
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getDatabase() (mdb *mgo.Database, err error) {
	session, err := db.OpenSession("") // Load mongodb session from env var
	mdb = session.DB("")               // Fetch default database

	return
}

func Serve(Port string, AppRoot string) {

	// Fallback port to 5000 for local test
	if len(Port) <= 0 {
		Port = "5000"
	}

	appRoot = AppRoot

	cookie = NewCookieStore()

	server := echo.New()

	// Setup Middleware
	server.Use(mw.Logger())
	server.Use(mw.Recover())
	server.Use(mw.Gzip())

	server.Use(AuthMiddleware())

	// Setup Route
	setupRoute(server)

	// Setup Template
	server.SetRenderer(getRenderer())

	// Host static files
	setupStatic(server)

	// Start server
	gracehttp.Serve(server.Server(":" + Port))
}

func getRenderer() *Renderer {
	var tmplPath string
	if len(appRoot) <= 0 {
		// Get correctly path to template file
		_, filename, _, _ := runtime.Caller(1)
		tmplPath = filepath.Dir(filename)
	} else {
		tmplPath = appRoot + "/server"
	}
	templates := template.Must(template.ParseGlob(tmplPath + "/template/*.tmpl"))
	templates.ParseGlob(tmplPath + "/template/*/*.tmpl")
	return &Renderer{
		templates: templates,
	}
}

func setupStatic(server *echo.Echo) {
	var packagePath string
	if len(appRoot) <= 0 {
		_, filename, _, _ := runtime.Caller(1)
		packagePath = filepath.Dir(filename)
	} else {
		packagePath = appRoot + "/server"
	}
	server.Static("/js", packagePath+"/static/js")
	server.Static("/css", packagePath+"/static/css")
	server.Static("/vendor", packagePath+"/static/vendor")
	server.Static("/img", packagePath+"/static/img")
}

func setupRoute(s *echo.Echo) {
	s.Get("/", index)
	s.Get("/login", login)
	s.Post("/login", verifyLogin)

	s.Get("/jobs", jobs)
	s.Get("/jobs/new", newJob)
	s.Post("/jobs", createJob)

	s.Get("/job/:id", getJob)
	s.Put("/job/:id", updateJob)
	s.Delete("/job/:id", deleteJob)
}

func index(c *echo.Context) error {
	return c.Render(200, "index", nil)
}
