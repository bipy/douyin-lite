package queries

const (
	getHashByUsernameQuery = `SELECT USER_ID, USER_PASSWORD FROM USERS WHERE USER_NAME = ?`
	createUserQuery        = `INSERT INTO USERS(USER_NAME, USER_PASSWORD) VALUES (?,?)`
)

func (db *DouyinQuery) GetHashByUsername(username string) (id int, hash string, err error) {
	row := db.QueryRowx(getHashByUsernameQuery, username)
	err = row.Scan(&id, &hash)
	return
}

func (db *DouyinQuery) CreateUser(username, hash string) (int, error) {
	r, err := db.Exec(createUserQuery, username, hash)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
