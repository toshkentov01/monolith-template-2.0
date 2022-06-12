package user

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/toshkentov01/template/config"
	entities "github.com/toshkentov01/template/entities/user"
	"github.com/toshkentov01/template/pkg/errs"
	"github.com/toshkentov01/template/pkg/logger"
	"github.com/toshkentov01/template/pkg/utils"
)

var (
	loggerTest  logger.Logger
	cfg         *config.Configuration
	userService *UserService
)

func TestMain(m *testing.M) {
	path := "/home/sardor/go/src/github.com/sardortoshkentov/monolith-template-2.0/.env"

	if info, err := os.Stat(path); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(path)
			if err != nil {
				fmt.Println("Err:", err)
			}
		}
	} else {
		fmt.Println("Not exists")
	}

	cfg = config.Config()
	loggerTest = logger.New(cfg.LogLevel, "user_service")
	userService = NewUserService(loggerTest)

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateUser(t *testing.T) {
	userID, _ := uuid.NewRandom()

	err := userService.CreateUser(context.Background(), &entities.CreateUserModel{
		ID:       userID.String(),
		Username: utils.UsernameGenerator(5),
		Email:    utils.EmailGenerator(5),
		Password: "sardor@11",
	})

	if err != nil {
		t.Errorf("Erorr while testing CreateUser method, error: %v", err.Error())
	}
}

func TestGetUser(t *testing.T) {
	type testCase struct {
		id        string
		expectErr error
	}

	testCases := []testCase{
		{
			id:        "d66c4934-690f-4ac0-8568-108bdd8397e7",
			expectErr: nil,
		},
		{
			id:        "eb19ef77-cfb7-427e-a83d-c2aff53fb100",
			expectErr: errs.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("TestGetUser_%s", tc.id), func(t *testing.T) {
			_, err := userService.GetUser(context.Background(), &entities.GetUserModel{
				ID: tc.id,
			})

			if err != tc.expectErr {
				t.Errorf("Error while testing GetUser method, error: %v", err.Error())
			}
		})
	}
}
