package cmd

import (
	"grpc-kit-service/provider/config"
	"grpc-kit-service/provider/db"
	"grpc-kit-service/service"
	"grpc-kit-service/worker/consumer"
	"grpc-kit-service/worker/cron"
	"log"
	"net"

	"github.com/espitman/protos-kit/kit"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Gin",
	Long:  `Start Gin`,
	Run: func(cmd *cobra.Command, args []string) {
		PORT := config.GetString("port")

		db.Connect()
		consumer.Start()
		cron.Run()

		lis, err := net.Listen("tcp", ":"+PORT)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		s := service.Server{}
		grpcServer := grpc.NewServer()

		kit.RegisterKitServiceServer(grpcServer, &s)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
