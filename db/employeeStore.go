package db

import (
	"employee/log"
	"employee/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InsertEmployee(engine *gorm.DB, employee *model.Employee) (*model.Employee, error) {
	if err := engine.Table("employee.employee_opn").Create(&employee).Error; err != nil {
		//fmt.Printf("InsertEmployee error : %++v\n", err)
		log.Log.WithFields(logrus.Fields{
			"Error":err,
			"RequestId" : employee.RequestId,
		}).Error("InsertEmployee error")
		return nil, err
	}
	return employee, nil
}
func UpdateEmployee(engine *gorm.DB, employee *model.Employee) (*model.Employee, error) {
	result := engine.Table("employee.employee_opn").Updates(&employee).Where("em_pid = ?", employee.EmpId)
	if result.Error != nil {
		log.Log.WithFields(logrus.Fields{
			"Error":result.Error,
			"RequestId" : employee.RequestId,
		}).Error("UpdateEmployee error")
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		err2 := fmt.Errorf("not found emp_id %++v", employee.EmpId)
		//fmt.Printf("UpdateEmployee error : %++v", err2)
		log.Log.WithFields(logrus.Fields{
			"Error":"not found emp_id",
			"RequestId" : employee.RequestId,
			"empId" : employee.EmpId,
		}).Error("UpdateEmployee error : not found emp_id")
		return nil, err2
	}
	return employee, nil
}
func DeleteEmployee(engine *gorm.DB, employee *model.Employee) (*model.Employee, error) {
	result := engine.Table("employee.employee_opn").Delete(&employee).Where("em_pid = ?", employee.EmpId)
	if result.Error != nil {
		log.Log.WithFields(logrus.Fields{
			"Error":result.Error,
			"RequestId" : employee.RequestId,
		}).Error("DeleteEmployee error")
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		err2 := fmt.Errorf("not found emp_id %++v", employee.EmpId)
		log.Log.WithFields(logrus.Fields{
			"Error":"not found emp_id",
			"RequestId" : employee.RequestId,
			"empId" : employee.EmpId,
		}).Error("DeleteEmployee error : not found emp_id")
		return nil, err2
	}
	return employee, nil
}
