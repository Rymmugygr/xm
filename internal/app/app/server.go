package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/Rymmugygr/xm/config"
	"github.com/Rymmugygr/xm/internal/app/controller/restapi"
	repository "github.com/Rymmugygr/xm/internal/app/infra/db"
	"github.com/Rymmugygr/xm/internal/app/usecase"
	"github.com/Rymmugygr/xm/internal/pkg/httpserver"
	"github.com/Rymmugygr/xm/internal/pkg/postgres"
	"github.com/sirupsen/logrus"
)

func Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := logrus.New()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Errorf("error loading config: %v\n", err)
		return
	}

	db, err := postgres.New(logger, cfg.Db)
	if err != nil {
		logger.Errorf("error db: %v\n", err)
		return
	}
	defer db.Close()

	useCases := initUseCases(db)

	server := httpserver.New(logger, httpserver.Addr(cfg.Http.Addr))
	restapi.NewRouter(logger, server.App, cfg, useCases)

	logger.Infof("Starting restapi server on %s...\n", cfg.Http.Addr)
	go func() {
		err := server.App.Listen(cfg.Http.Addr)
		logger.Error(err, "Error restapi server error")
	}()
	logger.Infof("Started")

	<-ctx.Done()
	stop()
}

func initUseCases(db *postgres.Postgres) usecase.UseCase {
	companyRepo := repository.NewCompanyRepository(db.Gorm)

	return usecase.UseCase{
		FindOneCompanyUseCase: usecase.NewFindOneCompanyUseCaseImpl(companyRepo),
		CreateCompanyUseCase:  usecase.NewCreateCompanyUseCaseImpl(companyRepo),
		ModifyCompanyUseCase:  usecase.NewModifyCompanyUseCaseImpl(companyRepo),
		DeleteCompanyUseCase:  usecase.NewDeleteCompanyUseCaseImpl(companyRepo),
	}
}
