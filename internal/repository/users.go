package repository

import "database/sql"

/* if exists вместо ON CONFLICT (telegram_id) DO NOTHING
проверка на if exists внутри транкзации
валидация username на спец символы и правила тг ника
*/
func CreateUser(db *sql.DB, telegramID int64, username string) error {
	_, err := db.Exec("INSERT INTO users (telegram_id, username) VALUES ($1, $2) ON CONFLICT (telegram_id) DO NOTHING", telegramID, username)
	if err != nil {
		return err
	}
	return nil
}
