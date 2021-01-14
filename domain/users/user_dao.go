package users

import (
	"github.com/bhanupratap1810/bookstore_users-api/constants"
	"github.com/bhanupratap1810/bookstore_users-api/datasources/mongo"
	"github.com/bhanupratap1810/bookstore_users-api/datasources/mysql"
	"github.com/bhanupratap1810/bookstore_users-api/utils/date_utils"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"github.com/bhanupratap1810/bookstore_users-api/utils/mysql_utils"
)

type UserDaoService interface {
	Get(*User) *errors.RestErr
	Save(*User) *errors.RestErr
	Update(*User) *errors.RestErr
	Delete(*User) *errors.RestErr
}

type userDaoMysql struct {
	DbService mysql.DbService
}

type UserDaoMongo struct {
	DbService mongo.DbService
}

func NewUserDaoMysqlService(dbService mysql.DbService) UserDaoService {

	return &userDaoMysql{
		DbService: dbService,
	}
}

func (u *userDaoMysql) Get(user *User) *errors.RestErr {
		stmt, err := mysql.Client.Prepare(constants.QueryGetUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		result := stmt.QueryRow(user.Id)
		if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
			return mysql_utils.ParseError(getErr)
		}
		return nil
}

func (u *userDaoMysql) Save(user *User) *errors.RestErr {
		stmt, err := mysql.Client.Prepare(constants.QueryInsertUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		user.DateCreated = date_utils.GetNowString()

		insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
		if saveErr != nil {
			return mysql_utils.ParseError(saveErr)
		}
		userId, err := insertResult.LastInsertId()
		if err != nil {
			return mysql_utils.ParseError(saveErr)
		}
		user.Id = userId
		return nil
}

func (u *userDaoMysql) Update(user *User) *errors.RestErr {
		stmt, err := mysql.Client.Prepare(constants.QueryUpdateUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
		if err != nil {
			return mysql_utils.ParseError(err)
		}
		return nil
}

func (u *userDaoMysql) Delete(user *User) *errors.RestErr {
		stmt, err := mysql.Client.Prepare(constants.QueryDeleteUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		if _, err = stmt.Exec(user.Id); err != nil {
			return mysql_utils.ParseError(err)
		}
		return nil
}

func (u *UserDaoMongo) Get(*User) *errors.RestErr {
	return nil
}

func (u *UserDaoMongo) Save(*User) *errors.RestErr {
	return nil
}

func (u *UserDaoMongo) Update(*User) *errors.RestErr {
	return nil
}

func (u *UserDaoMongo) Delete(*User) *errors.RestErr {
	return nil
}

//func (user *User) Get() *errors.RestErr {
//	stmt, err := users_db.Client.Prepare(queryGetUser)
//	if err != nil {
//		return errors.NewInternalServerError(err.Error())
//	}
//	defer stmt.Close()
//
//	result := stmt.QueryRow(user.Id)
//	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
//		return mysql_utils.ParseError(getErr)
//	}
//	return nil
//}
//
//func (user *User) Save() *errors.RestErr {
//	stmt, err := users_db.Client.Prepare(queryInsertUser)
//	if err != nil {
//		return errors.NewInternalServerError(err.Error())
//	}
//	defer stmt.Close()
//
//	user.DateCreated = date_utils.GetNowString()
//
//	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
//	if saveErr != nil {
//		return mysql_utils.ParseError(saveErr)
//	}
//	userId, err := insertResult.LastInsertId()
//	if err != nil {
//		return mysql_utils.ParseError(saveErr)
//	}
//	user.Id = userId
//	return nil
//}
//
//func (user *User) Update() *errors.RestErr {
//	stmt, err := users_db.Client.Prepare(queryUpdateUser)
//	if err != nil {
//		return errors.NewInternalServerError(err.Error())
//	}
//	defer stmt.Close()
//
//	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
//	if err != nil {
//		return mysql_utils.ParseError(err)
//	}
//	return nil
//}
//
//func (user *User) Delete() *errors.RestErr {
//	stmt, err := users_db.Client.Prepare(queryDeleteUser)
//	if err != nil {
//		return errors.NewInternalServerError(err.Error())
//	}
//	defer stmt.Close()
//
//	if _, err = stmt.Exec(user.Id); err != nil {
//		return mysql_utils.ParseError(err)
//	}
//	return nil
//}
