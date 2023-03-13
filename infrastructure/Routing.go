package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"

	bc "github.com/set2002satoshi/my-site-api_v2/interfaces/controllers/blog"
	cc "github.com/set2002satoshi/my-site-api_v2/interfaces/controllers/category"
	uc "github.com/set2002satoshi/my-site-api_v2/interfaces/controllers/user"
)

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

	usersController := uc.NewUserController(r.DB)
	blogsController := bc.NewBlogController(r.DB)
	categoriesController := cc.NewCategoryController(r.DB)

	userNotLoggedIn := r.Gin.Group("/api")
	{
		userNotLoggedIn.POST("/users/get", func(c *gin.Context) { usersController.Find(c) })
		userNotLoggedIn.POST("/users/get/all", func(c *gin.Context) { usersController.FindAll(c) })
		userNotLoggedIn.POST("/users/create", func(c *gin.Context) { usersController.Create(c) })
		userNotLoggedIn.POST("/users/update", func(c *gin.Context) { usersController.Update(c) })
		userNotLoggedIn.POST("/users/delete", func(c *gin.Context) { usersController.DeleteById(c) })
	}

	blogNotLoggedIn := r.Gin.Group("/api")
	{
		blogNotLoggedIn.POST("/blogs/get", func(c *gin.Context) { blogsController.Find(c) })
		blogNotLoggedIn.POST("/blogs/get/all", func(c *gin.Context) { blogsController.FindAll(c) })
		blogNotLoggedIn.POST("/blogs/create", func(c *gin.Context) { blogsController.Create(c) })
		blogNotLoggedIn.POST("/blogs/update", func(c *gin.Context) { blogsController.Update(c) })
		blogNotLoggedIn.POST("/blogs/delete", func(c *gin.Context) { blogsController.DeleteById(c) })
	}

	categoryNotLoggedIn := r.Gin.Group("/api")
	{
		categoryNotLoggedIn.POST("/category/get", func(c *gin.Context) { categoriesController.Find(c) })
		categoryNotLoggedIn.POST("/category/get/all", func(c *gin.Context) { categoriesController.FindAll(c) })
		categoryNotLoggedIn.POST("/category/create", func(c *gin.Context) { categoriesController.Create(c) })
	}

	r.Gin.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})
}

func (r *Routing) Run() {
	err := r.Gin.Run(r.Port)
	if err != nil {
		panic(err)
	}
}
