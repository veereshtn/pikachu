package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kutty-kumar/db_commons/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"pikachu/pkg/domain"
	"pikachu/pkg/pb"
	r "pikachu/pkg/repository"
	"pikachu/pkg/svc"
	"time"
)

func NewGRPCServer(logger *logrus.Logger, dbConnectionString string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    time.Duration(viper.GetInt("config.keepalive.time")) * time.Second,
				Timeout: time.Duration(viper.GetInt("config.keepalive.timeout")) * time.Second,
			},
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// logging middleware
				grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),

				// Request-Id interceptor
				requestid.UnaryServerInterceptor(),


				// Metrics middleware
				grpc_prometheus.UnaryServerInterceptor,

				// validation middleware
				grpc_validator.UnaryServerInterceptor(),

				// collection operators middleware
				gateway.UnaryServerInterceptor(),
			),
		),
	)

	// create new postgres database
	db, err := gorm.Open("mysql", dbConnectionString)
	db.LogMode(true)
	if err != nil {
		return nil, err
	}

	dropTables(db)
	createTables(db)

	domainFactory := db_commons.NewDomainFactory()
	domainFactory.RegisterMapping("user", func() db_commons.Base {
		return &domain.User{}
	})
	domainFactory.RegisterMapping("identity", func() db_commons.Base {
		return &domain.Identity{}
	})

	externalIdSetter := func(externalId string, base db_commons.Base) db_commons.Base {
		base.SetExternalId(externalId)
		return base
	}
	userBaseDao := db_commons.NewBaseGORMDao(db, domainFactory.GetMapping("user"), externalIdSetter)

	identityBaseDao := db_commons.NewBaseGORMDao(db, domainFactory.GetMapping("identity"), externalIdSetter)
	userAttributeBaseDao := db_commons.NewBaseGORMDao(db, domainFactory.GetMapping("user_attributes"), externalIdSetter)
	identityRepository := r.NewIdentityGormRepository(identityBaseDao)
	userAttributeRepository := r.NewUserAttributeGormRepository(userAttributeBaseDao)
	// register service implementation with the grpcServer
	userBaseSvc := db_commons.NewBaseSvc(userBaseDao)
	identityBaseSvc := db_commons.NewBaseSvc(identityBaseDao)
	userAttributeBaseSvc := db_commons.NewBaseSvc(userAttributeBaseDao)
	identityService := svc.NewIdentityService(identityBaseSvc, &identityRepository)
	userService := svc.NewUserService(userBaseSvc, identityService)
	userAttributeService := svc.UserAttributeService{userAttributeBaseSvc, &userAttributeRepository}

	pb.RegisterUserServiceServer(grpcServer, &userService)
	return grpcServer, nil
}

func createTables(db *gorm.DB) {
	db.CreateTable(domain.User{})
	db.CreateTable(domain.Identity{}).AddForeignKey("user_id", "users(external_id)", "CASCADE", "CASCADE")
	db.CreateTable(domain.Address{}).AddForeignKey("user_id", "users(external_id)", "CASCADE", "CASCADE")
}

func dropTables(db *gorm.DB) {
	db.DropTableIfExists(domain.Identity{}, domain.Address{}, domain.User{})
}
