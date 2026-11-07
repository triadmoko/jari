package main

import (
	_ "app/docs"
	"app/internal/app"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//	@title			Jari API
//	@version		1.0
//	@description	REST API for Jari, an open-source work and project management platform (Jira clone).
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	https://www.apache.org/licenses/LICENSE-2.0

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

func main() {
	// Subcommand: `main migrate <up|down [N|all]|version|force V>`
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := runMigrate(os.Args[2:]); err != nil {
			log.Fatal(err)
		}
		return
	}

	application, err := app.New()
	if err != nil {
		log.Fatal("Failed to initialize application:", err)
	}
	defer application.Close()

	addr := application.Config.Server.Host + ":" + application.Config.Server.Port
	server := &http.Server{
		Addr:    addr,
		Handler: application.Engine,
	}

	go func() {
		log.Printf("Server starting on %s (ENV=%s)", addr, application.Config.Env)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
