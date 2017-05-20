package utilityData

import (
	"fmt"

	"github.com/go-errors/errors"
	_ "github.com/go-sql-driver/mysql"
)

func Authenticate(userName, userPassword string) bool {

	var isExists bool
	rows, err := DBtravel.Query("select 1 from user_authentication where is_active = 1 and user_name =? and user_password = ?", userName, userPassword)

	if err != nil {
		InsertErrorLogs(1, fmt.Sprint(errors.Errorf(err.Error()).ErrorStack()))
	}

	if rows.Next() {
		isExists = true
	}

	return isExists
}
