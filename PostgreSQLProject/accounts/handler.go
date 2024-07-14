package accounts

import (
	"PostgreSQLProject/accounts/dto"
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"net/http"
)

func New(ctx *context.Context) (*Handler, error) {
	connectionString := "host=0.0.0.0 port=5432 user=postgres password=12345 dbname=postgres"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	return &Handler{
		accounts: db,
		ctx:      ctx,
	}, nil
}

type Handler struct {
	accounts *sql.DB
	ctx      *context.Context
}

// Close Закрывает handler
func (h *Handler) Close() error {
	return h.accounts.Close()
}

// CreateAccount Создаёт аккаунт с балансом
func (h *Handler) CreateAccount(c echo.Context) (err error) {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	row, err := h.accounts.QueryContext(*h.ctx, "SELECT name, balance FROM accounts WHERE name = $1", request.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "database lookup error")
	}
	defer func() {
		tempErr := row.Close()
		if err == nil && tempErr != nil {
			err = c.String(http.StatusInternalServerError, "database close error")
		}
	}()

	if ok := row.Next(); ok {
		err = c.String(http.StatusForbidden, "account already exists")
		return err
	}

	_, err = h.accounts.ExecContext(*h.ctx, "INSERT INTO accounts(name, balance) VALUES($1, $2)", request.Name, request.Amount)

	if err != nil {
		err = c.String(http.StatusInternalServerError, "database insertion error")
		return err
	}

	err = c.NoContent(http.StatusOK)
	return err
}

// GetAccount Получает баланс аккаунта
func (h *Handler) GetAccount(c echo.Context) (err error) {
	var request dto.GetAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	row, err := h.accounts.QueryContext(*h.ctx, "SELECT name, balance FROM accounts WHERE name = $1", request.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "database lookup error")
	}
	defer func() {
		tempErr := row.Close()
		if err == nil && tempErr != nil {
			err = c.String(http.StatusInternalServerError, "database close error")
		}
	}()

	if ok := row.Next(); !ok {
		err = c.String(http.StatusForbidden, "account does not exist")
		return err
	}

	var response dto.GetAccountResponse

	err = row.Scan(&response.Name, &response.Amount)
	if err != nil {
		err = c.String(http.StatusInternalServerError, "database scan error")
		return err
	}

	err = c.JSON(http.StatusOK, response)
	return err
}

// DeleteAccount Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) (err error) {

	var request dto.DeleteAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	row, err := h.accounts.QueryContext(*h.ctx, "SELECT name, balance FROM accounts WHERE name = $1", request.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "database lookup error")
	}
	defer func() {
		tempErr := row.Close()
		if err == nil && tempErr != nil {
			err = c.String(http.StatusInternalServerError, "database close error")
		}
	}()

	if ok := row.Next(); !ok {
		err = c.String(http.StatusForbidden, "account does not exist")
		return err
	}

	_, err = h.accounts.ExecContext(*h.ctx, "DELETE FROM accounts WHERE name = $1", request.Name)
	if err != nil {
		err = c.String(http.StatusInternalServerError, "database delete error")
		return err
	}

	err = c.NoContent(http.StatusOK)
	return err
}

// PatchAccount Меняет баланс
func (h *Handler) PatchAccount(c echo.Context) (err error) {

	var request dto.PatchAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	row, err := h.accounts.QueryContext(*h.ctx, "SELECT name, balance FROM accounts WHERE name = $1", request.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "database lookup error")
	}
	defer func() {
		tempErr := row.Close()
		if err == nil && tempErr != nil {
			err = c.String(http.StatusInternalServerError, "database close error")
		}
	}()

	if ok := row.Next(); !ok {
		err = c.String(http.StatusForbidden, "account does not exist")
		return err
	}

	_, err = h.accounts.ExecContext(*h.ctx, "UPDATE accounts SET balance=$1 WHERE name=$2", request.Amount, request.Name)
	if err != nil {
		err = c.String(http.StatusInternalServerError, "database update error")
		return err
	}

	err = c.NoContent(http.StatusOK)
	return err
}

// ChangeAccount Меняет имя
func (h *Handler) ChangeAccount(c echo.Context) (err error) {
	var request dto.ChangeAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	row, err := h.accounts.QueryContext(*h.ctx, "SELECT name, balance FROM accounts WHERE name = $1", request.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "database lookup error")
	}
	defer func() {
		tempErr := row.Close()
		if err == nil && tempErr != nil {
			err = c.String(http.StatusInternalServerError, "database close error")
		}
	}()

	if ok := row.Next(); !ok {
		err = c.String(http.StatusForbidden, "account does not exist")
		return err
	}

	row, err = h.accounts.QueryContext(*h.ctx, "SELECT name, balance FROM accounts WHERE name = $1", request.NameNew)
	if err != nil {
		return c.String(http.StatusInternalServerError, "database lookup error")
	}
	defer func() {
		tempErr := row.Close()
		if err == nil && tempErr != nil {
			err = c.String(http.StatusInternalServerError, "database close error")
		}
	}()

	if ok := row.Next(); ok {
		err = c.String(http.StatusForbidden, "name already taken")
		return err
	}

	_, err = h.accounts.ExecContext(*h.ctx, "UPDATE accounts SET name=$1 WHERE name=$2", request.NameNew, request.Name)
	if err != nil {
		err = c.String(http.StatusInternalServerError, "database update error")
		return err
	}

	err = c.NoContent(http.StatusOK)
	return err
}
