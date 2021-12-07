package controllers

import (
	"employee/handler"
	"employee/model"
	"github.com/gin-gonic/gin"
	"net/http"

)

func CreateEmployee(c *gin.Context)  {

	service := handler.Handler{}
	 input := &model.CreateRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response,err :=service.SaveEmployee(input)
	if err != nil{
		res := &model.Response{
			ErrorCode: string(http.StatusInternalServerError),
			ErrorDesc: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else{
	c.JSON(http.StatusOK, response)
	}
}
func UpdateEmployee(c *gin.Context)  {
	service := handler.Handler{}
	input := &model.UpdateRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response,err :=service.UpdateEmployee(input)
	if err != nil{
		res := &model.Response{
			ErrorCode: "500",
			ErrorDesc: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else{
		c.JSON(http.StatusOK, response)
	}
}
func DeleteEmployee(c *gin.Context)  {
	service := handler.Handler{}
	input := &model.DeleteRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response,err :=service.DeleteEmployee(input)
	if err != nil{
		res := &model.Response{
			ErrorCode: "500",
			ErrorDesc: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else{
		c.JSON(http.StatusOK, response)
	}
}