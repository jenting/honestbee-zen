package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq" // for sqlx.Connect usage
	"github.com/pkg/errors"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	sqlxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/jmoiron/sqlx"

	"github.com/honestbee/Zen/config"
)

type postgres struct {
	readTimeout           time.Duration
	writeTimeout          time.Duration
	transactionMaxTimeout time.Duration
	db                    *sqlx.DB
}

// NewPostgres returns a Database instance.
func NewPostgres(conf *config.Config) (Database, error) {
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName("helpcenter-zendesk-postgres"))

	connSchema := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		conf.Database.User,
		conf.Database.DBName,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(conf.Database.ConnectTimeoutSec)*time.Second,
	)
	defer cancel()

	db, err := sqlxtrace.Connect("postgres", connSchema)
	if err != nil {
		return nil, errors.Wrapf(err, "db: [NewPostgress] connect failed")
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Wrapf(err, "db: [NewPostgress] ping failed")
	}

	db.SetMaxIdleConns(conf.Database.MaxIdle)
	db.SetMaxOpenConns(conf.Database.MaxActive)

	return &postgres{
		db:                    db,
		readTimeout:           time.Duration(conf.Database.ReadTimeoutSec) * time.Second,
		writeTimeout:          time.Duration(conf.Database.WriteTimeoutSec) * time.Second,
		transactionMaxTimeout: time.Duration(conf.Database.TransactionMaxTimeoutSec) * time.Second,
	}, nil
}

// Close closes the database for preventing memory leaking.
func (p *postgres) Close() error {
	return errors.Wrapf(p.db.Close(), "db: [Close] close database failed")
}

// Select is the wrapper of sqlx SelectContext.
func (p *postgres) Select(ctx context.Context, dest interface{}, query string) error {
	ctx, cancel := context.WithTimeout(ctx, p.readTimeout)
	defer cancel()

	err := p.db.SelectContext(ctx, dest, query)
	if err == sql.ErrNoRows {
		return ErrNoRows
	}

	return errors.Wrapf(err, "db: [Select] failed on %q query", query)
}

// Get is the wrapper of sqlx GetContext.
func (p *postgres) Get(ctx context.Context, dest interface{}, query string) error {
	ctx, cancel := context.WithTimeout(ctx, p.readTimeout)
	defer cancel()

	err := p.db.GetContext(ctx, dest, query)
	if err == sql.ErrNoRows {
		return ErrNoRows
	}

	return errors.Wrapf(err, "db: [Get] failed on %q query", query)
}

// NameExec is the wrapper of sqlx NamedExecContext.
func (p *postgres) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(ctx, p.writeTimeout)
	defer cancel()

	result, err := p.db.NamedExecContext(ctx, query, arg)
	if err != nil {
		return nil, errors.Wrapf(err, "db: [NamedExec] failed query:%q, arg:%v", query, arg)
	}

	return result, nil
}

// postgresTransaction is the wrapper of sqlx transaction with ctx.
type postgresTransaction struct {
	tx     *sqlx.Tx
	ctx    context.Context
	cancel context.CancelFunc
	err    error
}

// Begin begins a transaction.
func (p *postgres) Begin() (DatabaseTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), p.transactionMaxTimeout)

	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		cancel()
		return nil, errors.Wrapf(err, "db: [Begin] BeginTxx failed")
	}

	return &postgresTransaction{
		tx:     tx,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// Select is the wrapper of sqlx tx SelectContext.
// It will doing rollback if there has an error occurred.
func (p *postgresTransaction) Select(dest interface{}, query string) {
	if p.err != nil {
		return
	}

	err := p.tx.SelectContext(p.ctx, dest, query)
	if err != nil {
		defer p.cancel()
		if er := p.tx.Rollback(); er != nil {
			err = errors.Wrapf(err, "db: [transaction Select] rollback failed:%v", er)
		}
	}

	p.err = errors.Wrapf(err, "db: [transaction Select] failed query:%q", query)
}

// Get is the wrapper of sqlx tx GetContext.
// It will doing rollback if there has an error occurred.
func (p *postgresTransaction) Get(dest interface{}, query string) {
	if p.err != nil {
		return
	}

	err := p.tx.GetContext(p.ctx, dest, query)
	if err != nil {
		defer p.cancel()
		if er := p.tx.Rollback(); er != nil {
			err = errors.Wrapf(err, "db: [transaction Get] rollback failed:%v", er)
		}
	}

	p.err = errors.Wrapf(err, "db: [transaction Get] failed query:%s", query)
}

// NamedExec is the wrapper of sqlx tx NamedExecContext.
// It will doing rollback if there has an error occurred.
func (p *postgresTransaction) NamedExec(query string, arg interface{}) sql.Result {
	if p.err != nil {
		return nil
	}

	result, err := p.tx.NamedExecContext(p.ctx, query, arg)
	if err != nil {
		defer p.cancel()
		if er := p.tx.Rollback(); er != nil {
			err = errors.Wrapf(err, "db: [transaction NamedExec] rollback failed:%v", er)
		}
	}

	p.err = errors.Wrapf(err, "db: [transaction NamedExec] failed query:%q, arg:%v", query, arg)
	return result
}

func (p *postgresTransaction) Err() error {
	return p.err
}

func (p *postgresTransaction) Commit() {
	if p.err != nil {
		return
	}

	defer p.cancel()
	p.err = errors.Wrapf(p.tx.Commit(), "db: [transaction Commit] commit failed")
}
