package main

import (
	"context"
	"fmt"
	"github.com/lty/shortlink/admin/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := router.Register()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000), // TODO: 这里最终要从配置文件中读
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
