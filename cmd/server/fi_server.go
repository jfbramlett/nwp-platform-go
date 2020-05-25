package main

import (
	"context"
	"fmt"

	"github.com/jfbramlett/go-aop/pkg/aop"
	"github.com/jfbramlett/go-aop/pkg/logging"
	"github.com/jfbramlett/go-aop/pkg/metrics"
	"github.com/jfbramlett/go-aop/pkg/tracing"
	"github.com/jfbramlett/nwp-platform-go/pkg/fi"
	"github.com/spf13/cobra"
)

// The serveCmd will execute the generate command
var fiServerCmd = &cobra.Command{
	Use:   "fi",
	Short: "starts a financial institution server",
	Run:   serveFi,
}

func init() {
	// Here we create the command line flags for our app, and bind them to our package-local
	// config variable.
	flags := fiServerCmd.Flags()
	flags.String("service_name", "fibank", "name of the fi")
	flags.Bool("metrics_enabled", true, "if metrics should be produced")
	flags.Float64("trace_sample_rate", 1.0, "the rate to sample trace at (percentage as decimal between 0 and 1)")
	flags.String("web_root", "./", "root directory for serving our content")

	// Add the "serve" sub-command to the root command.
	rootCmd.AddCommand(fiServerCmd)
}

func serveFi(cmd *cobra.Command, args []string) {
	logging.InitLogging()
	logger := logging.GetLogger(context.Background())

	service, err := cmd.Flags().GetString("service_name")
	if err != nil {
		logger.Error(err, "failed to resolve service name")
		return
	}

	webRoot, err := cmd.Flags().GetString("web_root")
	if err != nil {
		logger.Error(err, "failed to resolve web root")
		return
	}

	aop.InitAOP(service)
	//aop.RegisterJoinPoint(aop.NewRegexPointcut(".*"), aop.NewSpanFuncAdvice())
	aop.RegisterJoinPoint(aop.NewRegexPointcut(".*"), aop.NewTimedFuncAdvice(service, fmt.Sprintf("%s metrics", service)))
	aop.RegisterJoinPoint(aop.NewRegexPointcut(".*"), aop.NewLoggingFuncAdvice())

	metrics.InitMetrics(metrics.MetricsConfig{Port: 1005, URL: "/metrics"})
	tracing.InitTracing(tracing.TracingConfig{Service: service, ReporterPort: 8090, ReporterUrl: "http://localhost:9411/api/v2/spans"})

	runner := fi.NewFIWebRunner(webRoot)
	runner.Run()
}
