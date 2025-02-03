package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Amos-Do/astudio/server/config"
	"github.com/gin-gonic/gin"
)

func TestCORS(t *testing.T) {
	// set config
	conf := config.Config{
		Server: config.ServerConf{
			Run: "debug",
		},
	}
	// the allowed origin that want to check
	origin := "*"

	// create gin router with middleware config
	g := gin.Default()
	g.Use(CORS(&conf))
	g.GET("/", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	server := httptest.NewServer(g)
	defer server.Close()

	// set client to request the server with the test origin
	client := &http.Client{}
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://%s", server.Listener.Addr().String()),
		nil,
	)
	req.Header.Add("Origin", origin)

	get, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// check the allow origin
	o := get.Header.Get("Access-Control-Allow-Origin")
	if o != origin {
		t.Errorf("Got '%s' ; expecting origin '%s'", o, origin)
	}
}
