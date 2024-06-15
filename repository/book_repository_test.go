package repository

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"gin-go-testing/model/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type unitTestBookRepositorySuite struct {
	suite.Suite
	br   BookRepository
	mock sqlmock.Sqlmock
	db   *sql.DB
	ctx  *gin.Context
}

func TestUnitTestBookRepository(t *testing.T) {
	suite.Run(t, &unitTestBookRepositorySuite{})
}

func (u *unitTestBookRepositorySuite) SetupTest() {
	u.br = NewBookRepositoryImpl()

	u.ctx = &gin.Context{}
	db, mock, _ := sqlmock.New()

	u.mock = mock
	u.db = db
}

func (u *unitTestBookRepositorySuite) TearDownTest() {
	u.db.Close()
}

func (u *unitTestBookRepositorySuite) TestFindOneById_Success() {
	data := domain.Book{
		Id:     1,
		Title:  "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
		Author: "James Clear",
	}

	rows := sqlmock.NewRows([]string{"id", "title", "author"}).AddRow(data.Id, data.Title, data.Author)

	u.mock.ExpectQuery(`SELECT id, title, author FROM books WHERE id=\$1`).WithArgs(1).WillReturnRows(rows)

	result, err := u.br.FindOneById(u.ctx, u.db, 1)

	u.Nil(err)
	u.NotNil(result)
	u.Equal(data.Id, result.Id)
	u.Equal(data.Title, result.Title)
	u.Equal(data.Author, result.Author)

	if err := u.mock.ExpectationsWereMet(); err != nil {
		u.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (u *unitTestBookRepositorySuite) TestFindOneById_Failed() {
	u.mock.ExpectQuery(`SELECT id, title, author FROM books WHERE id=\$1`).WithArgs(2).WillReturnError(sql.ErrNoRows)

	result, err := u.br.FindOneById(u.ctx, u.db, 2)

	u.Nil(result)
	u.NotNil(err)

	if err := u.mock.ExpectationsWereMet(); err != nil {
		u.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (u *unitTestBookRepositorySuite) TestFindAll_Success() {
	data := []*domain.Book{{
		Id:     1,
		Title:  "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
		Author: "James Clear",
	}, {
		Id:     2,
		Title:  "The 7 Habits of Highly Effective People",
		Author: "Stephen R. Covey",
	}}

	// convert to domain.Book to driver.Value for mock purposes
	var values [][]driver.Value
	for _, e := range data {
		values = append(values, []driver.Value{e.Id, e.Title, e.Author})
	}

	rows := sqlmock.NewRows([]string{"id", "title", "author"}).AddRows(values...)

	u.mock.ExpectQuery(`SELECT id, title, author FROM books`).WithoutArgs().WillReturnRows(rows)

	result, err := u.br.FindAll(u.ctx, u.db)

	u.Nil(err)
	u.NotNil(result)
	u.Equal(result, data)

	if err := u.mock.ExpectationsWereMet(); err != nil {
		u.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (u *unitTestBookRepositorySuite) TestFindAll_Failed() {
	rows := sqlmock.NewRows([]string{"id", "title", "author"}).AddRows([][]driver.Value{}...)
	u.mock.ExpectQuery(`SELECT id, title, author FROM books`).WithoutArgs().WillReturnRows(rows)

	result, err := u.br.FindAll(u.ctx, u.db)
	u.Nil(result)
	u.NotNil(err)

	if err := u.mock.ExpectationsWereMet(); err != nil {
		u.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (u *unitTestBookRepositorySuite) TestCreate_Success() {
	data := &domain.Book{
		Id:     1,
		Title:  "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
		Author: "James Clear",
	}

	row := sqlmock.NewRows([]string{"id"}).AddRow(data.Id)
	u.mock.ExpectQuery(`INSERT INTO books\(title, author\) VALUES\(\$1,\$2\) RETURNING id`).WithArgs(data.Title, data.Author).WillReturnRows(row)

	result, err := u.br.Create(u.ctx, u.db, data)

	u.NotNil(result)
	u.Nil(err)
	u.Equal(data, result)

	if err := u.mock.ExpectationsWereMet(); err != nil {
		u.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (u *unitTestBookRepositorySuite) TestCreate_Failed() {
	data := &domain.Book{
		Id:     1,
		Title:  "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
		Author: "James Clear",
	}

	u.mock.ExpectQuery(`INSERT INTO books\(title, author\) VALUES\(\$1,\$2\) RETURNING id`).WithArgs(data.Title, data.Author).WillReturnError(errors.New("some error in db"))

	result, err := u.br.Create(u.ctx, u.db, data)

	u.Nil(result)
	u.NotNil(err)

	if err := u.mock.ExpectationsWereMet(); err != nil {
		u.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}
