package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/google/gops/agent"
	"github.com/smith-30/petit/logger"
	"github.com/smith-30/petit/server"
	"github.com/smith-30/petit/server/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configPath = ""
	conf       = config.Config{}
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve http server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// gops
		if err := agent.Listen(agent.Options{}); err != nil {
			panic(err)
		}

		//
		// create logger
		//
		zl, err := logger.NewLogger(logger.Config{
			Development:         conf.Logger.Development,
			Level:               conf.Logger.Level,
			Encoding:            conf.Logger.Encoding,
			OutputPaths:         conf.Logger.OutputPaths,
			AppErrorOutputPaths: conf.Logger.AppErrorOutputPaths,
			ErrorOutputPaths:    conf.Logger.ErrorOutputPaths,
			EncoderConfig: logger.EncoderConfig{
				MessageKey:    conf.Logger.MessageKey,
				LevelKey:      conf.Logger.LevelKey,
				TimeKey:       conf.Logger.TimeKey,
				NameKey:       conf.Logger.NameKey,
				CallerKey:     conf.Logger.CallerKey,
				StacktraceKey: conf.Logger.StacktraceKey,
				LevelEncoder:  conf.Logger.LevelEncoder,
				CallerEncoder: conf.Logger.CallerEncoder,
			},
		})
		if err != nil {
			panic(err)
		}

		// if err := rdb.InitRDB(); err != nil {
		// 	panic(err)
		// }

		zl.Info("start process", zap.String("version_info", verInfo()))

		s := server.NewServer(
			server.Address(conf.Server.Host, conf.Server.Port),
			server.Logger(zl),
			server.Routes(),
		)
		go func() {
			if err := s.Start(); err != nil {
				zl.Error("", zap.Error(err))
				panic(err)
			}
		}()

		// Setting up signal capturing
		stop := make(chan os.Signal)
		signal.Notify(stop, os.Interrupt)

		// Waiting for SIGINT (pkill -2)
		<-stop
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err := s.Shutdown(ctx); err != nil {
			zl.Error("", zap.Error(err))
		}
	},
}

func init() {
	// merge command
	rootCmd.AddCommand(serveCmd)

	//
	// command args
	//
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "example.toml", "config file name")

	// Here you will define your flags and configuration settings.
	viper.SetConfigFile(configPath)

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 設定ファイルの内容を構造体にコピーする
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
}
