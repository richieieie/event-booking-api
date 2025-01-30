package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/richieieie/event-booking/internal/router"
)

func Run() error {
	err := godotenv.Load("/Users/trung/Workspace/go/api-with-gin/config/env/gin.env")
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	// gin.SetMode(setting.ServerSetting.RunMode)

	router := router.NewGinRouter()
	// readTimeout := setting.ServerSetting.ReadTimeout
	// writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%s", os.Getenv("PORT"))
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
		// ReadTimeout:    readTimeout,
		// WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	return server.ListenAndServe()
}
