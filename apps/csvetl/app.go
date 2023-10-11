package csvetl

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mymhimself/logger"
	"github.com/mymhimself/simple-csv-reader/apps"
	hCollections "github.com/mymhimself/simple-csv-reader/internal/gateway/handlers/collections"
	mWriter "github.com/mymhimself/simple-csv-reader/internal/models/writer"
	"github.com/mymhimself/simple-csv-reader/internal/services/processor"
	sReader "github.com/mymhimself/simple-csv-reader/internal/services/reader"
	sWriter "github.com/mymhimself/simple-csv-reader/internal/services/writer"
	"github.com/mymhimself/simple-csv-reader/internal/services/writer/publisher"
	"github.com/mymhimself/simple-csv-reader/pkg/config"
	"github.com/mymhimself/simple-csv-reader/pkg/constants"
	"github.com/mymhimself/simple-csv-reader/pkg/mongodb"
	"github.com/mymhimself/simple-csv-reader/pkg/xlogger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewApp(cmd *cobra.Command, args []string) (apps.Runnable, error) {

	{
		configPath, err := cmd.Flags().GetString(constants.ConfigPathFlag)
		if err != nil {
			return nil, err
		}
		config.InitConfig(configPath)
	}
	{
		err := xlogger.InitLogger()
		if err != nil {
			return nil, err
		}
	}

	echoApp, err := newEcho(cmd, args)
	if err != nil {
		return nil, err
	}

	logger.Info("port", viper.GetInt(constants.Port))
	return func() error {
		return echoApp.Start(fmt.Sprintf(":%d", viper.GetInt(constants.Port)))
	}, nil

}

// ─────────────────────────────────────────────────────────────────────────────
func newEcho(cmd *cobra.Command, args []string) (*echo.Echo, error) {
	logger.Debug(">>>>>>>>>>>>", args)
	e := echo.New()
	// Middleware
	//todo user our logger with caller Report
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover()) //TODO
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "*"},
		AllowMethods: []string{"*"},
	}))
	e.Use(middleware.CORS())
	// e.Use(instrument.DefaultHttpMiddleware)

	// e.HTTPErrorHandler = errorhandler.DefaultErrorHandler

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	const APPVersion = "1.0.0"
	const APIVersion = "1.0.0"

	e.OPTIONS("*", func(c echo.Context) error { return nil })
	e.GET("/version/api", func(c echo.Context) error {
		c.String(http.StatusOK, APPVersion)
		return nil
	})
	e.GET("/version/app", func(c echo.Context) error {
		c.String(http.StatusOK, APIVersion)
		return nil
	})

	type hcResponse struct {
		Status    bool   `json:"status"`
		Version   string `json:"version"`
		StartedAt int64  `json:"started_at"`
		Uptime    int64  `json:"uptime"`
	}

	{
		var resp hcResponse
		resp.Status = true
		resp.Version = APPVersion
		resp.StartedAt = time.Now().Unix()

		e.GET("/health-check", func(c echo.Context) error {
			resp.Uptime = time.Now().Unix() - resp.StartedAt
			return c.JSON(http.StatusOK, resp)
		})
	}

	os.Setenv(constants.ServiceName, strings.ToLower(cmd.Use))
	{
		// instantiate a mongo client
		mongoClient, err := mongodb.New(mongodb.InitOptionURI(viper.GetString(constants.MongoDBURI)))
		if err != nil {
			return nil, err
		}
		// instantiate a writerModel
		writerModel, err := mWriter.New(
			mWriter.InitOptionMongoClient(mongoClient),
			mWriter.InitOptionDatabaseName(viper.GetString(constants.MongoDBName)),
		)
		if err != nil {
			return nil, err
		}

		// instantiate a writer service
		writerService, err := sWriter.New(
			sWriter.InitModel(writerModel),
			sWriter.InitConfigsFromViper(viper.GetViper()),
		)
		if err != nil {
			return nil, err
		}

		etlHandler, err := hCollections.New(
			hCollections.InitOptionService(writerService),
		)

		err = etlHandler.Register(e)
		if err != nil {
			return nil, err
		}
	}

	// run etl to read form file and write it to the mongodb
	{
		err := runETL(context.Background(), &runETLParams{
			FileName:   args[0],
			Delimiter:  args[1],
			Collection: args[2],
		})
		if err != nil {
			return nil, err
		}
	}

	return e, nil
}

// ─────────────────────────────────────────────────────────────────────────────

func runETL(ctx context.Context, params *runETLParams) error {
	// instantiate a reader service
	readerService, err := sReader.New(
		sReader.InitOptionDelimiter(params.Delimiter),
		sReader.InitOptionFileName(params.FileName),
	)
	if err != nil {
		return err
	}

	m, err := readerService.ReadMetaData(ctx)
	if err != nil {
		return err
	}

	pub, err := publisher.New(publisher.PublisherOptionFromViper(viper.GetViper().Sub(constants.Publisher)))
	if err != nil {
		return err
	}

	processorService, err := processor.New(
		processor.InitOptionDelimiter(params.Delimiter),
		processor.InitOptionObject(m),
		processor.InitOptionPublisher(pub),
		processor.InitOptionCollectionName(params.Collection),
	)
	if err != nil {
		return err
	}

	linesChannel := make(chan string, 100)

	// read lines from the file concurrently and pass them to the processor
	go readerService.ReadLines(ctx, linesChannel)
	go processorService.ProcessLines(ctx, linesChannel)
	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
type runETLParams struct {
	FileName   string
	Delimiter  string
	Collection string
	Database   string
}
