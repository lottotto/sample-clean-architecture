package infrastructure

import "github.com/gin-gonic/gin"

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: ":8080",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	//controller := controllers.NewUsersController(r.DB)
	r.Gin.LoadHTMLGlob("assets/templates/*.html")
	r.Gin.Static("/assets", "./assets")
	r.Gin.Static("/favicon.ico", "./assets/favicon.ico")
	r.Gin.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
}
func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
