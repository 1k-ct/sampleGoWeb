package user

import (
	"fmt"
	"net/http"

	user "github.com/1k-ct/sampleGoWeb/service"
	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type Controller struct{}

// Index action: GET /users
func (pc Controller) Index(c *gin.Context) {
	var s user.Service
	p, err := s.GetAll()
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
