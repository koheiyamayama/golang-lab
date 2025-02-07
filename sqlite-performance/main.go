package main

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"math/rand"
	"os"
	"sync"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"golang.org/x/sync/semaphore"
)

// コマンドライン引数
// --sec int = 30 テストを行う秒数
// --write-con int = 1 テストを行う間の書き込み同時接続数
// --read-con int = 1 テストを行う間の読み込み同時接続数

var (
	readSem  *semaphore.Weighted
	writeSem *semaphore.Weighted
)

func main() {
	ctx := context.Background()
	readSem = semaphore.NewWeighted(1)                      // TODO: コマンドライン引数 --read-con で指定された数にする
	writeSem = semaphore.NewWeighted(1)                     // TODO: コマンドライン引数 --write-con で指定された数にする
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second) // TODO: コマンドライン引数 --sec で指定された秒数にする
	defer cancel()

	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		slog.ErrorContext(ctx, "Error opening database", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()

	s := &sqlite{db: db}

	vv, err := s.Version(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error getting version", slog.Any("error", err))
		os.Exit(1)
	}
	slog.DebugContext(ctx, "SQLite version", slog.Any("version", vv))

	result, err := CreateUser(ctx, s)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			slog.InfoContext(ctx, "done")
			os.Exit(0)
		}

		slog.ErrorContext(ctx, "Error creating user", slog.Any("error", err))
		os.Exit(1)
	}
	slog.InfoContext(ctx, "Result", slog.Any("result", result))
}

func CreateUser(ctx context.Context, s *sqlite) (*Result, error) {
	slog.DebugContext(ctx, "calling CreateUser")

	createUser := func(wg *sync.WaitGroup, sem *semaphore.Weighted) {
		defer wg.Done()
		defer sem.Release(1)

		u := NewRandomUser()
		err := s.CreateUser(ctx, u)
		if err != nil {
			return
		}
		slog.DebugContext(ctx, "Created user", slog.Any("user", u))
	}

	var wg sync.WaitGroup
LOOP:
	for {
		select {
		case <-ctx.Done():
			slog.DebugContext(ctx, "CreateUser: Context done")
			break LOOP
		default:
			wg.Add(1)
			if err := writeSem.Acquire(ctx, 1); err != nil {
				return nil, err
			}

			go createUser(&wg, writeSem)
		}
	}
	wg.Wait()

	slog.DebugContext(ctx, "CreateUser done")

	return nil, nil
}

func init() {
	level := os.Getenv("LOG_LEVEL")

	switch level {
	case "DEBUG":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "INFO":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case "WARN":
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case "ERROR":
		slog.SetLogLoggerLevel(slog.LevelError)
	default:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}
}

func randomString() string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var result []byte
	for i := 0; i < 10; i++ {
		result = append(result, alphabet[rand.Intn(len(alphabet))])
	}
	return string(result)
}
