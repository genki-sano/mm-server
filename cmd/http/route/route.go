package route

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/genki-sano/mm-server/cmd/http/di"
)

// Route ルーティングの設定
func Route() *gin.Engine {
	e := createEngine()

	e.NoRoute(func(c *gin.Context) {
		msgs := []string{"no route to host"}
		c.JSON(http.StatusNotFound, gin.H{"errors": msgs})
	})

	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"healthCheck": "ok"})
	})

	api := e.Group("/api")

	auth := api.Group("/auth")
	{
		auth.GET("/verify", di.InitializeAuthVerify().Handler)
	}

	user := api.Group("/user")
	{
		user.GET("/list", di.InitializeUserList().Handler)
	}

	return e
}

func createEngine() *gin.Engine {
	gin.DisableConsoleColor()
	setMode()

	r := gin.Default()

	// COR設定の追加
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"https://mm-client-0506.netlify.app",
	}
	r.Use(cors.New(config))

	return r
}

func setMode() {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
}
