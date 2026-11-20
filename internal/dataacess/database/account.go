package database

import (
	"context"
	"database/sql"
	"log"
	"github.com/doug-martin/goqu/v9"
)

var TabNameAccounts = goqu.T("accounts")

const (
	ColNameAccountsID          = "id"
	ColNameAccountsAccountName = "account_name"
)

type Account struct{
	ID uint64 `db:"id"`
	AccountName string `db:"account_name"`
}

type AccountDataAccessor interface{
	CreateAccount(ctx context.Context,account Account) (uint64, error) 
	GetAccountByID (ctx context.Context, id uint64) (Account, error)
	GetAccountByAccountName (ctx context.Context, accountName string) (Account, error)
	WithDatabase(db Database) AccountDataAccessor
}

type accountDataAccessor struct{
	database Database 
}

func NewAccountDataAccessor(
	database *goqu.Database,
) AccountDataAccessor{
	return &accountDataAccessor{
		database: database,
	}
}

func (a *accountDataAccessor) WithDatabase(database Database) AccountDataAccessor{
	return &accountDataAccessor{
		database: database,
	}
}


func (a *accountDataAccessor) CreateAccount(ctx context.Context, account Account) (uint64 ,error) {
	
	result, err := a.database.
		Insert(TabNameAccounts).
		Rows(goqu.Record{
			ColNameAccountsAccountName: account.AccountName,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		log.Printf("failed to create account %v",err)
		return 0,err
	}
	accountID, err := result.LastInsertId()
	if err != nil{
		return 0, err
	}
	
	return uint64(accountID),nil 
}

func (a *accountDataAccessor) GetAccountByID (ctx context.Context, id uint64) (Account, error){
	account := Account{}
	found, err:= a.database.
		From(TabNameAccounts).
		Where(goqu.C(ColNameAccountsID).Eq(id)).
		ScanStructContext(ctx, &account)
	
	if err != nil{
		return Account{}, err
	}
	if !found{
		return Account{}, sql.ErrNoRows
	}
	return account, nil 
}

func (a *accountDataAccessor) GetAccountByAccountName (ctx context.Context, accountName string) (Account, error){
	account := Account{}
	found, err:= a.database.
		From(TabNameAccounts).
		Where(goqu.C(ColNameAccountsAccountName).Eq(accountName)).
		ScanStructContext(ctx, &account)
		
	if err != nil{
		return Account{}, err
	}
	if !found{
		return Account{}, sql.ErrNoRows
	}
	return account, nil 
}