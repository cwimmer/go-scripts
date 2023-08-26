package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/cwimmer/awsutils/pkg/instances"
	"github.com/rodaine/table"
	"log"
	"os"
	"time"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	instanceIds := instances.Instances(cfg)
	tbl := table.New("Instance ID", "Name", "CreationDate", "Profile").WithHeaderFormatter(func(format string, vals ...interface{}) string {
		return ""
	})
	//	tbl := table.New("", "", "", "")
	for _, instanceId := range instanceIds {
		tbl.AddRow(instanceId,
			instances.InstanceName(cfg, instanceId),
			instances.InstanceDate(cfg, instanceId).Format(time.RFC3339),
			os.Getenv("AWS_VAULT"),
		)

	}
	tbl.Print()
}
