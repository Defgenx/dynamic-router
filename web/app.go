package web

import (
	"github.com/defgenx/dynamic-router/web/config"
	"github.com/defgenx/dynamic-router/web/core"
	"github.com/howeyc/fsnotify"
	"log"
	"net/http"
)


type App struct {
	Router *core.Router
	Config     *config.Config
}

// Implement a singleton Pattern
func newApp() *App {
	app := &App{
		Router: core.NewRouter(),
		Config:     config.NewConfig()}

	go func() {
		app.watchConfigFile()
	}()

	return app
}

func (app *App) HandleRoute() {
	app.Router.SetRoutes(app.Config.Routes)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(app.Config)
	app.Router.ServeHTTP(w, r)
}

func (app *App) watchConfigFile() {
	// Create watcher instance
	configFileWatcher, _ := fsnotify.NewWatcher()
	defer configFileWatcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case ev := <-configFileWatcher.Event:
				if (ev.IsModify()) {
					log.Println("Route file has been modified, reload. ", ev)
					// Get new routes array hash
					app.Config.Routes = app.Config.GetRoutes()
					// Recreate controller structure containing mux router instance
					// Put instance in tmp var => we need to add route before replacing
					// That way it is invisible for the end user
					tmpNewRouterInstance := core.NewRouter()
					// Reload routes with new array hash
					tmpNewRouterInstance.SetRoutes(app.Config.Routes)
					// Set the new instance to the app
					app.Router = tmpNewRouterInstance
				}
			}
		}
	}()

	errWatcher := configFileWatcher.Watch(config.CONFIG_DIR)
	if errWatcher != nil {
		log.Println(errWatcher)
	}

	<-done
}

// Initialize my app object
var MyApp = newApp()
