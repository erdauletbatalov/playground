package main

func main() {

}

func CheckEmployee(db BaseDBClient, phone string) (err error, exists bool) {
	var employee interface{}
	err = db.Get(&employee, `SELECT name FROM employees WHERE phone = ?`, phone)
	if err != nil {
		return err, false
	}
	return nil, true
}

type BaseDBClient interface {
	Get(interface{}, string, ...interface{}) error
}
