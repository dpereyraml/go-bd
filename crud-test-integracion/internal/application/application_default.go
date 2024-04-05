package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// ConfigDefault is a struct that represents the default application configuration
type ConfigDefault struct {
	// Database is the database configuration
	Database mysql.Config
	// Address is the address of the application
	Address string
}

// NewApplicationDefault creates a new default application.
// func NewApplicationDefault(addr, filePathStore string) (a *ApplicationDefault) {
func NewApplicationDefault(cfg *ConfigDefault) (a *ApplicationDefault) {
	// default config
	/* defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}

	a = &ApplicationDefault{
		rt:            defaultRouter,
		addr:          defaultAddr,
		filePathStore: filePathStore,
	}
	return */
	cfgDefault := &ConfigDefault{
		Address: ":8080",
	}
	if cfg != nil {
		cfgDefault.Database = cfg.Database
		if cfg.Address != "" {
			cfgDefault.Address = cfg.Address
		}
	}

	return &ApplicationDefault{
		cfgDb: cfgDefault.Database,
		addr:  cfgDefault.Address,
	}

}

// ApplicationDefault is the default application.
type ApplicationDefault struct {
	// rt is the router.
	// rt *chi.Mux
	// addr is the address to listen.
	addr string
	// filePathStore is the file path to store.
	// filePathStore string

	// Database is the database configuration
	cfgDb mysql.Config
}

// TearDown tears down the application.
func (a *ApplicationDefault) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) Run() (err error) {
	// func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - store
	// st := store.NewStoreProductJSON(a.filePathStore)
	// dependencies
	// - database: connection
	db, err := sql.Open("mysql", a.cfgDb.FormatDSN())
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return
	}
	// - repository
	// rp := repository.NewRepositoryProductStore(st)
	rp := repository.NewProductMySQL(db)
	// - handler
	hd := handler.NewHandlerProduct(rp)

	// router

	// - router: chi
	rt := chi.NewRouter()
	// - middlewares
	// a.rt.Use(middleware.Logger)
	rt.Use(middleware.Logger)
	// a.rt.Use(middleware.Recoverer)
	rt.Use(middleware.Recoverer)
	// - endpoints
	//a.rt.Route("/products", func(r chi.Router) {
	rt.Route("/products", func(r chi.Router) {
		// GET /products/{id}
		r.Get("/{id}", hd.GetById())
		// POST /products
		r.Post("/", hd.Create())
		// PUT /products/{id}
		r.Put("/{id}", hd.UpdateOrCreate())
		// PATCH /products/{id}
		r.Patch("/{id}", hd.Update())
		// DELETE /products/{id}
		r.Delete("/{id}", hd.Delete())
	})

	// run
	err = http.ListenAndServe(a.addr, rt)
	if err != nil {
		return
	}
	return
}

/* // Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}
*/
