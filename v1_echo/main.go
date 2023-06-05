package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/echo", getEcho)

    router.Run("localhost:8090")
}

// getAlbums responds with the list of all albums as JSON.
func getEcho(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, map[string]string{
        "Response": "Echo v1",
    })
}
