package gracefully

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunGinGracefully(r *gin.Engine, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutdown Server ...")

	var timeout time.Duration = 1

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		logger.Printf("timeout of %d second(s).", timeout)
	}
	logger.Println("Server exiting")
}
