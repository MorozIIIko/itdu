package repository

import "database/sql"

type Quiz struct {
	ID   int64
	Name string
}

func GetQuizzes(db *sql.DB) ([]Quiz, error) {
	rows, err := db.Query("SELECT id, name FROM quizzes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var quizzes []Quiz
	for rows.Next() {
		var quiz Quiz
		if err := rows.Scan(&quiz.ID, &quiz.Name); err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return quizzes, nil
}
