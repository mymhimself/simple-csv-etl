package csvetl

import (
	"net"
	"os"
	"strings"

	mWriter "github.com/mymhimself/simple-csv-reader/internal/models/writer"
	sReader "github.com/mymhimself/simple-csv-reader/internal/services/reader"
	sWriter "github.com/mymhimself/simple-csv-reader/internal/services/writer"
	"github.com/mymhimself/simple-csv-reader/pkg/constants"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

func NewApp(cmd *cobra.Command) error {
	var (
		listener   net.Listener
		grpcServer *grpc.Server
		err        error
	)

	os.Setenv(constants.ServiceName, strings.ToLower(cmd.Use))
	// instantiate a mongo client
	// instantiate a writer writerModel
	writerModel, err := mWriter.New()
	if err != nil {
		return err
	}

	// instantiate a writer service
	writerService, err := sWriter.New(
		sWriter.InitOptionModel(writerModel),
	)
	if err != nil {
		return err
	}
	// instantiate a reader service
	service, err := sReader.New()
	if err != nil {
		return nil, err
	}

	return grpcServer, func() error { return grpcServer.Serve(listener) }
}
