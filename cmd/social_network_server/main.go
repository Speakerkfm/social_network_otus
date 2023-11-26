package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Speakerkfm/social_network_otus/internal/app/admin"
	"github.com/Speakerkfm/social_network_otus/internal/app/rest"
	"github.com/Speakerkfm/social_network_otus/internal/bootstrap"
	"github.com/Speakerkfm/social_network_otus/internal/service/auth"
	auth_pg_adapter "github.com/Speakerkfm/social_network_otus/internal/service/auth/adapter/pg"
	auth_repo "github.com/Speakerkfm/social_network_otus/internal/service/auth/repository"
	friend_service "github.com/Speakerkfm/social_network_otus/internal/service/friend"
	friend_pg_adapter "github.com/Speakerkfm/social_network_otus/internal/service/friend/adapter/pg"
	friend_repository "github.com/Speakerkfm/social_network_otus/internal/service/friend/repository"
	post_service "github.com/Speakerkfm/social_network_otus/internal/service/post"
	post_pg_adapter "github.com/Speakerkfm/social_network_otus/internal/service/post/adapter/pg"
	post_repo "github.com/Speakerkfm/social_network_otus/internal/service/post/repository"
	user_service "github.com/Speakerkfm/social_network_otus/internal/service/user"
	user_pg_adapter "github.com/Speakerkfm/social_network_otus/internal/service/user/adapter/pg"
	user_repo "github.com/Speakerkfm/social_network_otus/internal/service/user/repository"
)

func main() {
	db, err := bootstrap.CreatePostgres()
	if err != nil {
		log.Fatalf("fail to create pg conn: %s", err.Error())
	}
	defer db.Close()

	cache, err := bootstrap.NewCache("test")
	if err != nil {
		log.Fatalf("fail to create cache: %s", err.Error())
	}

	authRepo := auth_pg_adapter.New(auth_repo.New(db))
	authService := auth.New(authRepo)

	userRepo := user_pg_adapter.New(user_repo.New(db))
	userService := user_service.New(userRepo, authService)

	friendRepo := friend_pg_adapter.New(friend_repository.New(db))
	friendService := friend_service.New(friendRepo)

	postRepo := post_pg_adapter.New(post_repo.New(db))
	postService := post_service.New(cache, postRepo, friendService)

	httpServer, err := rest.New(userService, authService, postService)
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
