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
	FindByRole(*User) ([]User, *errors.RestErr)
	FindByEmailAndPassword(*User) *errors.RestErr
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
		stmt, err := u.DbService.Client.Prepare(constants.QueryGetUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		result := stmt.QueryRow(user.Id)
		if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Role, &user.State); getErr != nil {
			return mysql_utils.ParseError(getErr)
		}
		return nil
}

func (u *userDaoMysql) Save(user *User) *errors.RestErr {
		stmt, err := u.DbService.Client.Prepare(constants.QueryInsertUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		user.DateCreated = date_utils.GetNowString()

		insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Role, user.State)
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
		stmt, err := u.DbService.Client.Prepare(constants.QueryUpdateUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.Id)
		if err != nil {
			return mysql_utils.ParseError(err)
		}
		return nil
}

func (u *userDaoMysql) Delete(user *User) *errors.RestErr {
		stmt, err := u.DbService.Client.Prepare(constants.QueryDeleteUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		user.State="inactive"
		if _, err = stmt.Exec(user.State,user.Id); err != nil {
			return mysql_utils.ParseError(err)
		}
		return nil
}

func (u *userDaoMysql) FindByRole(user *User) ([]User, *errors.RestErr) {
	stmt, err := u.DbService.Client.Prepare(constants.QueryFindByRole)
	if err != nil {
		//logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Role, "active")
	if err != nil {
		//logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.State); err != nil {
			//logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError("no users matching role")
	}
	return results, nil
}

func (u *userDaoMysql) FindByEmailAndPassword(user *User) *errors.RestErr {
	stmt, err := u.DbService.Client.Prepare(constants.QueryFindByEmailAndPassword)
	if err != nil {
		//logger.Error("error when trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, "active")
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Role, &user.State); getErr != nil {
		//if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
		//	return rest_errors.NewNotFoundError("invalid user credentials")
		//}
		//logger.Error("error when trying to get user by email and password", getErr)
		return errors.NewInternalServerError(getErr.Error())
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
