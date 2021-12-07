package handler

import (
	"employee/model"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func getSession() (*gorm.DB,sqlmock.Sqlmock)  {
	db,mock,err:= sqlmock.New()
    convey.So(err,convey.ShouldBeNil)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	convey.So(err,convey.ShouldBeNil)
	return gormDB,mock

}
func TestCreateEmployee(t *testing.T) {
	convey.Convey("test..",t, func() {
		db,mock := getSession()
		service := Handler{DB: db}
		convey.Convey("TestCreateEmployee Success", func() {
			mock.ExpectBegin()
			rows := sqlmock.NewRows([]string{"em_pid"}).
				AddRow(1)
			mock.ExpectQuery("INSERT INTO \"employee\".\"employee_opn\"").WillReturnRows(rows)
			mock.ExpectCommit()
			request := &model.CreateRequest{
				RequestId: "11234",
				EmpName: "Passagon",
				Age: 25,
				Address: "132/22",
				Gender: "male",
				Department: "opn",
				MobileNumber: "0969643641",
			}
           response,err := service.SaveEmployee(request)
			convey.So(err,convey.ShouldBeNil)
			convey.So(response,convey.ShouldNotBeNil)
		})
		convey.Convey("TestCreateEmployee Fail", func() {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO \"url\".\"url_store\"").WillReturnError(errors.New("cannot database"))
			mock.ExpectCommit()
			request := &model.CreateRequest{
				RequestId: "11234",
				EmpName: "Passagon",
				Age: 25,
				Address: "132/22",
				Gender: "male",
				Department: "opn",
				MobileNumber: "0969643641",
			}
			response,err := service.SaveEmployee(request)
			convey.So(err,convey.ShouldNotBeNil)
			convey.So(response,convey.ShouldBeNil)

		})
	})
}
func TestUpdateEmployee(t *testing.T) {
	convey.Convey("test..",t, func() {
		db,mock := getSession()
		service := Handler{DB: db}
		convey.Convey("TestUpdateEmployee Success", func() {
			mock.ExpectBegin()
			mock.ExpectExec("UPDATE \"employee\".\"employee_opn\"").WillReturnResult(sqlmock.NewResult(1,1))
			mock.ExpectCommit()
			request := &model.UpdateRequest{
				EmpId: 1,
				RequestId: "11234",
				EmpName: "Passagon",
				Age: 25,
				Address: "132/22",
				Gender: "male",
				Department: "opn",
				MobileNumber: "0969643641",
			}
			response,err := service.UpdateEmployee(request)
			convey.So(err,convey.ShouldBeNil)
			convey.So(response,convey.ShouldNotBeNil)
		})
		convey.Convey("TestUpdateEmployee Fail", func() {
			mock.ExpectBegin()
			mock.ExpectExec("UPDATE \"employee\".\"employee_opn\"").WillReturnError(errors.New("cannot database"))
			mock.ExpectCommit()
			request := &model.UpdateRequest{
				EmpId: 1,
				RequestId: "11234",
				EmpName: "Passagon",
				Age: 25,
				Address: "132/22",
				Gender: "male",
				Department: "opn",
				MobileNumber: "0969643641",
			}
			response,err := service.UpdateEmployee(request)
			convey.So(err,convey.ShouldNotBeNil)
			convey.So(response,convey.ShouldBeNil)

		})
		convey.Convey("TestUpdateEmployee Notfound id", func() {
			mock.ExpectBegin()
			mock.ExpectExec("UPDATE \"employee\".\"employee_opn\"").WillReturnResult(sqlmock.NewResult(1,0))
			mock.ExpectCommit()
			request := &model.UpdateRequest{
				EmpId: 1,
				RequestId: "11234",
				EmpName: "Passagon",
				Age: 25,
				Address: "132/22",
				Gender: "male",
				Department: "opn",
				MobileNumber: "0969643641",
			}
			response,err := service.UpdateEmployee(request)
			convey.So(err,convey.ShouldNotBeNil)
			convey.So(response,convey.ShouldBeNil)

		})
	})
}
func TestDeleteEmployee(t *testing.T) {
	convey.Convey("test..",t, func() {
		db,mock := getSession()
		service := Handler{DB: db}
		convey.Convey("TestDeleteEmployee Success", func() {
			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM \"employee\".\"employee_opn\"").WillReturnResult(sqlmock.NewResult(1,1))
			mock.ExpectCommit()
			request := &model.DeleteRequest{
				EmpId: 1,
				RequestId: "11234",
			}
			response,err := service.DeleteEmployee(request)
			convey.So(err,convey.ShouldBeNil)
			convey.So(response,convey.ShouldNotBeNil)
		})
		convey.Convey("TestDeleteEmployee Fail", func() {
			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM \"employee\".\"employee_opn\"").WillReturnError(errors.New("cannot database"))
			mock.ExpectCommit()
			request := &model.DeleteRequest{
				EmpId: 1,
				RequestId: "11234",
			}
			response,err := service.DeleteEmployee(request)
			convey.So(err,convey.ShouldNotBeNil)
			convey.So(response,convey.ShouldBeNil)

		})
		convey.Convey("TestDeleteEmployee Notfound id", func() {
			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM \"employee\".\"employee_opn\"").WillReturnResult(sqlmock.NewResult(1,0))
			mock.ExpectCommit()
			request := &model.DeleteRequest{
				EmpId: 1,
				RequestId: "11234",
			}
			response,err := service.DeleteEmployee(request)
			convey.So(err,convey.ShouldNotBeNil)
			convey.So(response,convey.ShouldBeNil)

		})
	})
}
