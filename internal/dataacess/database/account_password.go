package database

import(
	"context"
)

var (
	TabNameAccountPassword =  goqu.T("account_passwords")
)

const(
	ColNameAccountPasswordsOfAccountID  = "of_account_id"
	ColNameAccountPasswordHash			= "hash"
)
type AccountPassword struct{
	OfAccountID  uint64 `sql:"of_account_id "`
	Hash string `sql:"hash"`
}

type AccountPasswordDataAccessor interface{
	CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error
	GetAccountPassword(ctx context.Context, ofAccountID uint64) (AccountPassword, error)
	UpdateAccountPassword(ctx context.Context, accountPassword AccountPassword) error
	
}

type accountPasswordDataAccessor struct{
	database *goqu.Database;
}

func NewAccounPasswordDataAccessor(
	database *goqu.Database,
) AccountPasswordDataAccessor{
	return &accountPasswordDataAccessor{
		database: database
	}
}

func (a *accountPasswordDataAccessor) CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	if _, err:= a.database.
		Insert("password").
		Rows(goqu.Record{
			"of_user_id": accountPassword.OfAccountID,
			"password": accountPassword.Hash,
		}).
		Executor().
		ExecContext(ctx); err != nil{
			log.Printf("failed to create account %v",err)
			return err
		}
	return nil 
}

func (a *accountPasswordDataAccessor) GetAccountPassword(ctx context.Context, ofAccountID uint64) (AccountPassword, error) {
	accountPassword := AccountPassword{}
	found, err := a.database.
		From(TabNameAccountPassword).
		Where(goqu.C(ColNameAccountPasswordsOfAccountID).Eq(ofAccountID)).
		ScanStructContext(ctx, &accountPassword)
	
	if err != nil{
		return AccountPassword{}, err
	}

	if !found{
		return AccountPassword{}, err
	}

	return accountPassword, nil 

}

func (a *accountPasswordDataAccessor) UpdateAccountPassword(ctx context.Context, password string) error{
	accountPassword := AccountPassword{}
	_, err := a.database.
		Update(TabNameAccountPassword).
		Set(goqu.Record{
			ColNameAccountPasswordHash: password
		}).
		Where(goqu.Ex{ColNameAccountPasswordsOfAccountID: accountPassword.OfAccountID}).
		Executor.
		ExecContext(ctx);
	
	if err != nil {
		return err
	}
	return nil 
}