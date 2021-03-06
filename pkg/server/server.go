package server

import (
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/services/grpc"
	"github.com/gigamono/gigamono/pkg/services/proto/generated"

	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"

	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
)

// APIServer represents an new REST-based server instance.
type APIServer struct {
	inits.App
	*gin.Engine
	Validate   validator.Validate
	AuthClient generated.AuthClient
}

// NewAPIServer creates a new server instance.
func NewAPIServer(app inits.App) (APIServer, error) {
	validate := *validator.New()
	var authClient generated.AuthClient

	// TODO: Sec: Insecure connection.
	// TODO: Remove this. It is not needed anymore. user.CreateIfNotExist already addresses this issue.
	client, err := grpc.GetInsecureClient("localhost", app.Config.Services.Auth.Ports.Private, app.Config)
	if err != nil {
		logs.FmtPrintln("initialising api server: unable to connect to Auth Service:", err)
	} else {
		authClient = client.(generated.AuthClient)
	}

	return APIServer{
		App:        app,
		Engine:     gin.Default(),
		Validate:   validate,
		AuthClient: authClient,
	}, nil
}

// Listen makes the server listen on specified port.
func (server *APIServer) Listen() error {
	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe() })
	grp.Go(func() error { return server.httpServe() })
	return grp.Wait()
}
