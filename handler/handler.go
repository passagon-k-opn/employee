package handler

import (
	"employee/db"
	"employee/model"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (t *Handler) initDB() {
	if nil ==t.DB{
		t.DB = db.EngineDB
	}
}

func (t *Handler) SaveEmployee(request *model.CreateRequest) (*model.Response,error) {
	t.initDB()
	engineDb := t.DB
	employee := &model.Employee{
		RequestId: request.RequestId,
		EmpName: request.EmpName,
		Age: request.Age,
		Address: request.Address,
		Gender: request.Gender,
		Department: request.Department,
		MobileNumber: request.MobileNumber,
	}
	res,err := db.InsertEmployee(engineDb, employee)
	if err != nil {
		return nil,err
	}
	response :=&model.Response{
		RequestId: res.RequestId,
		EmpId: res.EmpId,
		EmpName: res.EmpName,
	}
	return response,nil
}
func (t *Handler) UpdateEmployee(request *model.UpdateRequest) (*model.Response,error) {
	t.initDB()
	engineDb := t.DB
	employee := &model.Employee{
		EmpId: request.EmpId,
		RequestId: request.RequestId,
		EmpName: request.EmpName,
		Age: request.Age,
		Address: request.Address,
		Gender: request.Gender,
		Department: request.Department,
		MobileNumber: request.MobileNumber,
	}
	res,err := db.UpdateEmployee(engineDb, employee)
	if err != nil {
		return nil,err
	}
	response :=&model.Response{
		RequestId: res.RequestId,
		EmpId: res.EmpId,
	}
	return response,nil
}
func (t *Handler) DeleteEmployee(request *model.DeleteRequest) (*model.Response,error) {
	t.initDB()
	engineDb := t.DB
	employee := &model.Employee{
		EmpId: request.EmpId,
		RequestId: request.RequestId,
	}
	res,err := db.DeleteEmployee(engineDb, employee)
	if err != nil {
		return nil,err
	}
	response :=&model.Response{
		RequestId: res.RequestId,
		EmpId: res.EmpId,
	}
	return response,nil
}