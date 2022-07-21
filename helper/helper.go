package helper

import (
	"time"
	"todo/database/migration"
	"todo/models"
)

func SignUp(name, password string) (*string, error) {
	// language=SQL
	SQL := `INSERT INTO users(name, password)
            VALUES ($1, $2) RETURNING id;`

	var userDetails string
	err := migration.To.Get(&userDetails, SQL, name, password)
	if err != nil {
		return nil, err

	}
	return &userDetails, nil

}
func RetrieveCredentials(id string) (*models.User, error) {
	//language=SQL
	SQL := `SELECT password, id
            FROM users
            WHERE id = $1;`
	var credentials models.User
	err := migration.To.Get(&credentials, SQL, id)
	if err != nil {
		return nil, err

	}
	return &credentials, nil

}

func CreateSession(id string, expiryTime time.Time) (*string, error) {
	//language=SQL
	SQL := `INSERT INTO sessions(expiry_time, user_id)
            VALUES ($1, $2)
            RETURNING id;`
	var SessionId string
	err := migration.To.Get(&SessionId, SQL, expiryTime, id)
	if err != nil {
		return nil, err
	}
	return &SessionId, nil

}
func FetchSession(id string) (*models.Session, error) {
	//language=SQL
	SQL := `SELECT id, expiry_time, user_id 
            FROM sessions
            WHERE id = $1 and archived_at IS NULL`
	var SessionModel models.Session
	err := migration.To.Get(&SessionModel, SQL, id)
	if err != nil {
		return nil, err
	}
	return &SessionModel, nil

}
func DeleteSession(id string) error {
	//language=SQL
	SQL := `UPDATE sessions
            SET archived_at = $1
            WHERE id = $2`
	_, err := migration.To.Exec(SQL, time.Now(), id)
	if err != nil {
		return err
	}
	return nil

}

func UpdateUser(name, password, id string) error {
	//language=SQL
	SQL := `UPDATE users
            SET name= (CASE WHEN $1 = '' THEN name
                    ELSE $1 end),
             password= (CASE WHEN $2 = '' THEN password
                     ELSE $2 END)
            WHERE id = $3;`
	_, err := migration.To.Exec(SQL, name, password, id)
	if err != nil {
		return err
	}
	return nil

}

func DeleteUser(id string) error {
	//language=SQL
	SQL := `UPDATE users
            SET archived_at = $1
            WHERE id = $2;`
	_, err := migration.To.Exec(SQL, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func CreateTask(name, userId string) (*models.Task, error) {
	//language=SQL
	SQL := `INSERT INTO tasks(name, user_id)
            VALUES ($1, $2) RETURNING id ;`

	var taskDetails models.Task
	err := migration.To.Get(&taskDetails, SQL, name, userId)
	if err != nil {
		return nil, err

	}
	return &taskDetails, nil
}

func UpdateTask(id string) error {
	//language=SQL

	SQL := `UPDATE tasks
            SET is_completed = TRUE
            WHERE id = $1`
	_, err := migration.To.Exec(SQL, id)
	if err != nil {
		return err
	}
	return nil
}

func FetchTask(id string) (*models.Task, error) {
	//language=SQL
	var taskDetails models.Task
	SQL := `SELECT id, name, created_at, is_completed, user_id 
            FROM tasks
            WHERE id = $1 and archived_at IS NULL`
	err := migration.To.Get(&taskDetails, SQL, id)
	if err != nil {
		return nil, err

	}
	return &taskDetails, nil

}

func DeleteTask(id, name string) error {
	//language=SQL
	SQL := `UPDATE tasks
            SET archived_at = $1
            WHERE id = $2 and name = $3`
	_, err := migration.To.Exec(SQL, time.Now(), id, name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllTasks(userId string) error {
	//language=SQL
	SQL := `UPDATE tasks
            SET archived_at = $1
            WHERE user_id = $2`
	_, err := migration.To.Exec(SQL, time.Now(), userId)
	if err != nil {
		return err
	}
	return nil
}
