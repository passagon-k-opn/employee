package model

type UrlStore struct {
	ShortUrl string `gorm:"primaryKey;column:short_url"`
	FullUrl  string `gorm:"column:full_url"`
}
type Employee struct {
	EmpId        int64  `gorm:"primaryKey;autoIncrement;column:em_pid"`
	RequestId    string `gorm:"column:request_id"`
	EmpName      string `gorm:"column:emp_name"`
	Age          int    `gorm:"column:age"`
	Address      string `gorm:"column:address"`
	Gender       string `gorm:"column:gender"`
	Department   string `gorm:"column:department"`
	MobileNumber string `gorm:"column:mobile_number"`
}
type Response struct {
	RequestId string `json:"requestId,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorDesc string `json:"errorDesc,omitempty"`
	EmpId     int64  `json:"empId,omitempty,omitempty"`
	EmpName   string `json:"empName,omitempty"`
}
type CreateRequest struct {
	RequestId    string `json:"requestId" validate:"required"`
	EmpName      string `json:"empName"  validate:"required"`
	Age          int    `json:"age"  validate:"required"`
	Address      string `json:"address"  validate:"required"`
	Gender       string `json:"gender"  validate:"required"`
	Department   string `json:"department"  validate:"required"`
	MobileNumber string `json:"mobileNumber"  validate:"required"`
}
type UpdateRequest struct {
	EmpId        int64  `json:"empId" validate:"required"`
	RequestId    string `json:"requestId" validate:"required"`
	EmpName      string `json:"empName,omitempty"`
	Age          int    `json:"age,omitempty"`
	Address      string `json:"address,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Department   string `json:"department,omitempty"`
	MobileNumber string `json:"mobileNumber,omitempty"`
}
type DeleteRequest struct {
	EmpId        int64  `json:"empId" validate:"required"`
	RequestId    string `json:"requestId" validate:"required"`
}