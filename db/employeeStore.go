package db

import (
	"employee/model"
	"fmt"
	"gorm.io/gorm"
)

func InsertEmployee(engine *gorm.DB, employee *model.Employee) (*model.Employee, error) {
	if err := engine.Table("employee.employee_opn").Create(&employee).Error; err != nil {
		fmt.Printf("InsertEmployee error : %++v\n", err)
		return nil, err
	}
	return employee, nil
}
func UpdateEmployee(engine *gorm.DB, employee *model.Employee) (*model.Employee, error) {
	result := engine.Table("employee.employee_opn").Updates(&employee).Where("em_pid = ?", employee.EmpId)
	if result.Error != nil {
		fmt.Printf("UpdateEmployee error : %++v\n", result.Error)
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		err2 := fmt.Errorf("not found emp_id %++v", employee.EmpId)
		fmt.Printf("UpdateEmployee error : %++v", err2)
		return nil, err2
	}
	return employee, nil
}
func DeleteEmployee(engine *gorm.DB, employee *model.Employee) (*model.Employee, error) {
	result := engine.Table("employee.employee_opn").Delete(&employee).Where("em_pid = ?", employee.EmpId)
	if result.Error != nil {
		fmt.Printf("DeleteEmployee error : %++v\n", result.Error)
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		err2 := fmt.Errorf("not found emp_id %++v", employee.EmpId)
		fmt.Printf("DeleteEmployee error : %++v", err2)
		return nil, err2
	}
	return employee, nil
}
