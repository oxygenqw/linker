package postgres_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository/postgres"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestGetByID(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	logger := logger.GetLogger()
	repo := postgres.NewStudentRepository(db, &logger)

	testID := uuid.New().String()
	expectedStudent := models.Student{
		ID:         uuid.MustParse(testID),
		TelegramID: 12345,
		FirstName:  "John",
		LastName:   "Doe",
	}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "telegram_id", "first_name", "middle_name", "last_name",
			"github", "job", "idea", "about", "university", "faculty", "course", "education",
		}).AddRow(
			expectedStudent.ID, expectedStudent.TelegramID, expectedStudent.FirstName, "", expectedStudent.LastName,
			"", "", "", "", "", "", "", "",
		)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about, 
              university, faculty, course, education FROM students WHERE id = $1`)).
			WithArgs(testID).
			WillReturnRows(rows)

		student, err := repo.GetByID(testID)
		assert.NoError(t, err)
		assert.Equal(t, expectedStudent.ID, student.ID)
		assert.Equal(t, expectedStudent.FirstName, student.FirstName)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about, 
              university, faculty, course, education FROM students WHERE id = $1`)).
			WithArgs(testID).
			WillReturnError(sql.ErrNoRows)

		_, err := repo.GetByID(testID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "student not found")
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("invalid uuid", func(t *testing.T) {
		_, err := repo.GetByID("invalid-uuid")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid UUID format")
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about, 
              university, faculty, course, education FROM students WHERE id = $1`)).
			WithArgs(testID).
			WillReturnError(errors.New("db error"))

		_, err := repo.GetByID(testID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get student")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetByTelegramID(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	logger := logger.GetLogger()
	repo := postgres.NewStudentRepository(db, &logger)

	telegramID := int64(12345)
	expectedStudent := models.Student{
		ID:         uuid.New(),
		TelegramID: telegramID,
		FirstName:  "Jane",
		LastName:   "Smith",
	}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "telegram_id", "first_name", "middle_name", "last_name",
			"github", "job", "idea", "about", "university", "faculty", "course", "education",
		}).AddRow(
			expectedStudent.ID, expectedStudent.TelegramID, expectedStudent.FirstName, "", expectedStudent.LastName,
			"", "", "", "", "", "", "", "",
		)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about,
              university, faculty, course, education FROM students WHERE telegram_id = $1`)).
			WithArgs(telegramID).
			WillReturnRows(rows)

		student, err := repo.GetByTelegramID(telegramID)
		assert.NoError(t, err)
		assert.Equal(t, expectedStudent.TelegramID, student.TelegramID)
		assert.Equal(t, expectedStudent.FirstName, student.FirstName)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about,
              university, faculty, course, education FROM students WHERE telegram_id = $1`)).
			WithArgs(telegramID).
			WillReturnError(sql.ErrNoRows)

		_, err := repo.GetByTelegramID(telegramID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "student not found")
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about,
              university, faculty, course, education FROM students WHERE telegram_id = $1`)).
			WithArgs(telegramID).
			WillReturnError(errors.New("db error"))

		_, err := repo.GetByTelegramID(telegramID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to retrieve student")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAll(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	logger := logger.GetLogger()
	repo := postgres.NewStudentRepository(db, &logger)

	students := []models.Student{
		{
			ID:        uuid.New(),
			FirstName: "Alice",
			LastName:  "Johnson",
		},
		{
			ID:        uuid.New(),
			FirstName: "Bob",
			LastName:  "Brown",
		},
	}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "telegram_id", "first_name", "middle_name", "last_name",
			"github", "job", "idea", "about", "university", "faculty", "course", "education",
		})
		for _, s := range students {
			rows.AddRow(
				s.ID, 0, s.FirstName, "", s.LastName,
				"", "", "", "", "", "", "", "",
			)
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about,
              university, faculty, course, education FROM students`)).
			WillReturnRows(rows)

		result, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, students[0].FirstName, result[0].FirstName)
		assert.Equal(t, students[1].FirstName, result[1].FirstName)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("empty result", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "telegram_id", "first_name", "middle_name", "last_name",
			"github", "job", "idea", "about", "university", "faculty", "course", "education",
		})

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about,
              university, faculty, course, education FROM students`)).
			WillReturnRows(rows)

		result, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Empty(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, telegram_id, first_name, middle_name, last_name, github, job, idea, about,
              university, faculty, course, education FROM students`)).
			WillReturnError(errors.New("db error"))

		_, err := repo.GetAll()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to retrieve students")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCreate(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	logger := logger.GetLogger()
	repo := postgres.NewStudentRepository(db, &logger)

	student := models.Student{
		TelegramID: 54321,
		FirstName:  "New",
		LastName:   "Student",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO students 
              (id, telegram_id, first_name, middle_name, last_name, github, job, idea, about,
               university, faculty, course, education) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`)).
			WithArgs(
				sqlmock.AnyArg(),
				student.TelegramID,
				student.FirstName,
				"",
				student.LastName,
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := repo.Create(student)
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO students 
              (id, telegram_id, first_name, middle_name, last_name, github, job, idea, about,
               university, faculty, course, education) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`)).
			WithArgs(
				sqlmock.AnyArg(),
				student.TelegramID,
				student.FirstName,
				"",
				student.LastName,
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
			).
			WillReturnError(errors.New("db error"))

		_, err := repo.Create(student)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to insert student")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	logger := logger.GetLogger()
	repo := postgres.NewStudentRepository(db, &logger)

	student := models.Student{
		ID:        uuid.New(),
		FirstName: "Updated",
		LastName:  "Name",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`
	UPDATE students 
	SET 
		first_name = $1,
		middle_name = $2,
		last_name = $3,
		github = $4,
		job = $5,
		idea = $6,
		about = $7,
		university = $8,
		faculty = $9,
		course = $10,
		education = $11
	WHERE id = $12
	`)).
			WithArgs(
				student.FirstName,
				"",
				student.LastName,
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				student.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(student)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`
	UPDATE students 
	SET 
		first_name = $1,
		middle_name = $2,
		last_name = $3,
		github = $4,
		job = $5,
		idea = $6,
		about = $7,
		university = $8,
		faculty = $9,
		course = $10,
		education = $11
	WHERE id = $12
	`)).
			WithArgs(
				student.FirstName,
				"",
				student.LastName,
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				student.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Update(student)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`
	UPDATE students 
	SET 
		first_name = $1,
		middle_name = $2,
		last_name = $3,
		github = $4,
		job = $5,
		idea = $6,
		about = $7,
		university = $8,
		faculty = $9,
		course = $10,
		education = $11
	WHERE id = $12
	`)).
			WithArgs(
				student.FirstName,
				"",
				student.LastName,
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				student.ID,
			).
			WillReturnError(errors.New("db error"))

		err := repo.Update(student)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to update student")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
