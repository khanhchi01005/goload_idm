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
	hashLogic Hash  

}

func NewAccount(
	accountDataAccessor database.AccountDataAccessor,
	accountPasswordDataAccessor database.AccountPasswordDataAccessor,
	hashLogic Hash,
) Account{
	return &account{
		accountDataAccessor: accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
		hashLogic : hashLogic,
	}
}

func (a *account)isAccountAccountNameTaken(ctx context.Context, accountName string) (bool, error){
	if _, err := a.accountDataAccessor.GetAccountByAccountName(ctx, accountName); err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return false, nil
		}
		return false, err
	}
	return true, nil 
}

func (a *account)CreateAccount(ctx context.Context, param CreateAccountParams) (CreateAccountOuput, error){
	accountNameTaken, err := a.isAccountAccountNameTaken(ctx, param.AccountName)
	if err != nil{
		return CreateAccountOuput{}, err
	}

	if accountNameTaken{
		return CreateAccountOuput{}, errors.New("Account name has been taken")
	}

	accountID, err := a.accountDataAccessor.CreateAccount(ctx, database.Account{
		AccountName: param.AccountName,
	})
	if err != nil {
		return CreateAccountOuput{}, err
	}

	hashedPassword, err := a.hashLogic.Hash(ctx, param.Password)
	if err != nil {
		return err
	}

	err := a.accountPasswordDataAccessor.CreateAccountPassword(ctx, database.AccountPassword{
		OfAccountID: accountID,
		Hash: hashedPassword
	})
	if err != nil{
		return CreateAccountOuput{}, err
	}

	return CreateAccountOuput{
		ID: accountID,
		AccountName: param.AccountName,
	}, nil 

}

func  (a *account)CreateSession(ctx context.Context, param CreateSessionParams ) (CreateSessionOutput, error){
	return CreateSessionOutput{}, errors.New("not implemented")
}
