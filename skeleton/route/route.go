package route

import (
	"database/sql"
	"log"
	"skeleton/domain/ddrivers"
	"skeleton/pb/drivers"

	"google.golang.org/grpc"
)

func GrpcRoute(grpcServer *grpc.Server, log *log.Logger, db *sql.DB) {
	driverServer := ddrivers.NewDriverHandler(log, db)

	drivers.RegisterDriversServiceServer(grpcServer, driverServer)
}
