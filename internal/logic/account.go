package logic

import (
	"context"
	"database/sql"
	"errors"
	"go_load/internal/dataacess/database"
	go_load "go_load/internal/generated"
)


type CreateAccountParams struct{
	AccountName string 
	Password string
}

type CreateSessionParams struct{
	AccountName string 
	Password string
}

type CreateAccountOuput struct{
	ID uint64
	AccountName string
}

type CreateSessionOutput struct{
	Account *go_load.Account
	Token string 
}

type Account interface{
	CreateAccount(ctx context.Context, param CreateAccountParams) (CreateAccountOuput, error)
	CreateSession(ctx context.Context, param CreateSessionParams ) (CreateSessionOutput, error)
}

type account struct{
	accountDataAccessor database.AccountDataAccessor
	accountPasswordDataAccessor database.AccountPasswordDataAccessor

}

func NewAccount(
	accountDataAccessor database.AccountDataAccessor,
	accountPasswordDataAccessor database.AccountPasswordDataAccessor,
) Account{
	return &account{
		accountDataAccessor: accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
	}
}

func (a *account)isAccountAccountNameTaken(ctx context.Context, accountName string) (bool, error){
	if_, err := a.accountDataAccessor.GetAccountByAccountName(ctx,accountName); err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return false, nil
		}
		return false, err
	}
	return true, nil 
}

func (a *account)CreateAccount(ctx context.Context, param CreateAccountParams) (CreateAccountOuput, error){
	accountNameTake, err := a.isAccountAccountNameTaken(ctx, param.AccountName) 
	

}

func  (a *account)CreateSession(ctx context.Context, param CreateSessionParams ) (CreateSessionOutput, error){
	
}
