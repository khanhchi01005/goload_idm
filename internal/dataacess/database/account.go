package database

import (
	"context"
	"log"
)

var TabNameAccounts = goqu.T("accounts")

const (
	ColNameAccountsID          = "id"
	ColNameAccountsAccountName = "account_name"
)

type Account struct{
	ID uint64 `sql:"user_id"`
	AccountName string `sql:"username"`
}

type AccountDataAccessor interface{
	CreateAccount(ctx context.Context,account Account) error 
	GetAccountByID (ctx context.Context, id uint64) (Account, error)
	GetAccountByAccountName (ctx context.Context, accountName string) (Account, error)
}

type accountDataAccessor struct{
	database *goqu.Database
}

func NewAccountDataAccessor(
	database *goqu.Database,
) AccountDataAccessor{
	return &accountDataAccessor{
		database: database,
	}
}


func (a *accountDataAccessor) CreateAccount(ctx context.Context, account Account) error {
	if _, err:= a.database.
	Insert(TabNameAccounts).
	Rows(goqu.Record{
		ColNameAccountsAccountName: account.AccountName,
	}).
	Executor().
	ExecContext(ctx); err != nil{
		log.Printf("failed to create account %v",err)
		return err
	}
	return nil 
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
		return Account{}, err
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
		return Account{}, err
	}
	return account, nil 
}