package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Speakerkfm/social_network_otus/internal/app/admin"
	"github.com/Speakerkfm/social_network_otus/internal/app/rest"
	"github.com/Speakerkfm/social_network_otus/internal/bootstrap"
	"github.com/Speakerkfm/social_network_otus/internal/service/user"
	pg_adapter "github.com/Speakerkfm/social_network_otus/internal/service/user/adapter/pg"
	"github.com/Speakerkfm/social_network_otus/internal/service/user/repository"
)

func main() {
	db, err := bootstrap.CreatePostgres()
	if err != nil {
		log.Fatalf("fail to create pg conn: %s", err.Error())
	}
	defer db.Close()

	userRepo := pg_adapter.New(repository.New(db))
	userService := user.New(userRepo)
	httpServer, err := rest.New(userService)
	if err != nil {
		log.Fatalf("fail to create http server: %s", err.Error())
	}
	adminServer, err := admin.New()
	if err != nil {
		log.Fatalf("fail to create admin server: %s", err.Error())
	}

	done := make(chan struct{})
	go func() {
		if err = httpServer.ListenAndServe(); err != nil {
			log.Fatalf("fail to start http server: %s", err.Error())
		}
		close(done)
	}()
	go func() {
		if err = adminServer.ListenAndServe(); err != nil {
			log.Fatalf("fail to start admin server: %s", err.Error())
		}
		close(done)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		<-c
		close(done)
	}()

	log.Println("app started")
	<-done
	log.Println("app stopped")
}
