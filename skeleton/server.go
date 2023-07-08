package main

import (
    "context"
    "log"
    "net"
    "os"
    "skeleton/pb/drivers"
    "skeleton/pb/generic"

    "google.golang.org/grpc"
)

func main() {

    log := log.New(os.Stdout, "grpc skeleton : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

    // listen tcp port
    lis, err := net.Listen("tcp", ":7070")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
        return
    }

    grpcServer := grpc.NewServer()

    // routing grpc services
    grpcRoute(grpcServer, log)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %s", err)
        return
    }
    log.Print("serve grpc on port: 7070")

}

func grpcRoute(grpcServer *grpc.Server, log *log.Logger) {
    driverServer := newDriverHandler(log)

    drivers.RegisterDriversServiceServer(grpcServer, driverServer)
}

type driverHandler struct {
    log *log.Logger
}

func newDriverHandler(log *log.Logger) *driverHandler {
    handler := new(driverHandler)
    handler.log = log
    return handler
}

func (u *driverHandler) List(ctx context.Context, in *drivers.DriverListInput) (*drivers.Drivers, error) {
    return &drivers.Drivers{}, nil
}

func (u *driverHandler) Create(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error) {
    return in, nil
}

func (u *driverHandler) Update(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error) {
    return in, nil
}

func (u *driverHandler) Delete(ctx context.Context, in *generic.Id) (*generic.BoolMessage, error) {
    return &generic.BoolMessage{}, nil
}