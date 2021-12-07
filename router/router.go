package router

import (
	"employee/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
   router := gin.Default()
   router.POST("/create/employee",controllers.CreateEmployee)
   router.POST("/update/employee",controllers.UpdateEmployee)
   router.DELETE("/delete/employee",controllers.DeleteEmployee)
   router.Run()
}