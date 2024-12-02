package repository

import (
	"database/sql"
	"log"

	"cv_service/internal/models"    // Замените на путь к вашим моделям
	_ "github.com/mattn/go-sqlite3" // Импортируем драйвер SQLite
)

type Storage struct {
	db *sql.DB
}

// New создает новое соединение с базой данных и инициализирует таблицы
func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	// Создание таблицы resumes
	resumeTableQuery := `
	CREATE TABLE IF NOT EXISTS resumes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		candidate_name TEXT NOT NULL,
		email TEXT NOT NULL,
		phone TEXT NOT NULL,
		experience TEXT,
		education TEXT
	);`

	// Создание таблицы vacancies
	vacancyTableQuery := `
	CREATE TABLE IF NOT EXISTS vacancies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		company TEXT NOT NULL,
		location TEXT NOT NULL,
		description TEXT
	);`

	// Выполнение запросов для создания таблиц
	if _, err := db.Exec(resumeTableQuery); err != nil {
		return nil, err
	}

	if _, err := db.Exec(vacancyTableQuery); err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully.")
	return &Storage{db: db}, nil
}

// Resume CRUD Methods

// CreateResume создает новое резюме
func (s *Storage) CreateResume(resume *models.Resume) (int, error) {
	result, err := s.db.Exec(`INSERT INTO resumes (candidate_name, email, phone, experience, education) VALUES (?, ?, ?, ?, ?)`,
		resume.CandidateName, resume.Email, resume.Phone, resume.Experience, resume.Education)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

// GetResume получает резюме по ID
func (s *Storage) GetResume(id int) (*models.Resume, error) {
	var resume models.Resume
	err := s.db.QueryRow(`SELECT id, candidate_name, email, phone, experience, education FROM resumes WHERE id = ?`, id).Scan(
		&resume.Id, &resume.CandidateName, &resume.Email, &resume.Phone, &resume.Experience, &resume.Education)
	if err != nil {
		return nil, err
	}
	return &resume, nil
}

// UpdateResume обновляет существующее резюме
func (s *Storage) UpdateResume(resume *models.Resume) error {
	_, err := s.db.Exec(`UPDATE resumes SET candidate_name = ?, email = ?, phone = ?, experience = ?, education = ? WHERE id = ?`,
		resume.CandidateName, resume.Email, resume.Phone, resume.Experience, resume.Education, resume.Id)
	return err
}

// DeleteResume удаляет резюме по ID
func (s *Storage) DeleteResume(id int) error {
	_, err := s.db.Exec(`DELETE FROM resumes WHERE id = ?`, id)
	return err
}

// Vacancy CRUD Methods

// CreateVacancy создает новую вакансию
func (s *Storage) CreateVacancy(vacancy *models.Vacancy) (int, error) {
	result, err := s.db.Exec(`INSERT INTO vacancies (title, company, location, description) VALUES (?, ?, ?, ?)`,
		vacancy.Title, vacancy.Company, vacancy.Location, vacancy.Description)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

// GetVacancy получает вакансию по ID
func (s *Storage) GetVacancy(id int) (*models.Vacancy, error) {
	var vacancy models.Vacancy
	err := s.db.QueryRow(`SELECT id, title, company, location, description FROM vacancies WHERE id = ?`, id).Scan(
		&vacancy.Id, &vacancy.Title, &vacancy.Company, &vacancy.Location, &vacancy.Description)
	if err != nil {
		return nil, err
	}
	return &vacancy, nil
}

// UpdateVacancy обновляет существующую вакансию
func (s *Storage) UpdateVacancy(vacancy *models.Vacancy) error {
	_, err := s.db.Exec(`UPDATE vacancies SET title = ?, company = ?, location = ?, description = ? WHERE id = ?`, vacancy.Title, vacancy.Company, vacancy.Location, vacancy.Description, vacancy.Id)
	return err
}

// DeleteVacancy удаляет вакансию по ID
func (s *Storage) DeleteVacancy(id int) error {
	_, err := s.db.Exec(`DELETE FROM vacancies WHERE id = ?`, id)
	return err
}

// GetAllResumes получает все резюме
func (s *Storage) GetAllResumes() ([]*models.Resume, error) {
	rows, err := s.db.Query(`SELECT id, candidate_name, email, phone, experience, education FROM resumes`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []*models.Resume
	for rows.Next() {
		var resume models.Resume
		if err := rows.Scan(&resume.Id, &resume.CandidateName, &resume.Email, &resume.Phone, &resume.Experience, &resume.Education); err != nil {
			return nil, err
		}
		resumes = append(resumes, &resume)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resumes, nil
}

// GetAllVacancies получает все вакансии
func (s *Storage) GetAllVacancies() ([]*models.Vacancy, error) {
	rows, err := s.db.Query(`SELECT id, title, company, location, description FROM vacancies`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []*models.Vacancy
	for rows.Next() {
		var vacancy models.Vacancy
		if err := rows.Scan(&vacancy.Id, &vacancy.Title, &vacancy.Company, &vacancy.Location, &vacancy.Description); err != nil {
			return nil, err
		}
		vacancies = append(vacancies, &vacancy)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vacancies, nil
}
