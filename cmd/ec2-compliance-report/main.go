package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/cwimmer/awsutils/pkg/instances"
	"github.com/rodaine/table"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	instanceIds := instances.Instances(cfg)
	tbl := table.New(
		"Instance ID",
		"Name",
		"CreationDate",
		"Profile",
		"Owner",
		"Compliance Ticket",
	).WithHeaderFormatter(func(format string, vals ...interface{}) string {
		return ""
	})
	var instanceAge map[string][]string = make(map[string][]string)
	for _, instanceId := range instanceIds {
		var date = instances.InstanceDate(cfg, instanceId).String()
		if instanceAge[date] == nil {
			a := []string{instanceId}
			instanceAge[date] = a
		} else {
			instanceAge[date] = append(instanceAge[date], instanceId)
		}
	}
	keys := make([]string, 0, len(instanceAge))
	for k := range instanceAge {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		for _, instanceId := range instanceAge[k] {
			var owner string
			var compliance_ticket string
			for _, tag := range instances.GetInstance(cfg, instanceId).Tags {
				if *tag.Key == "owner" {
					owner = *tag.Value
				} else if *tag.Key == "compliance_ticket" {
					compliance_ticket = *tag.Value
				}
			}
			tbl.AddRow(instanceId,
				instances.InstanceName(cfg, instanceId),
				instances.InstanceDate(cfg, instanceId).Format(time.RFC3339),
				os.Getenv("AWS_VAULT"),
				owner,
				compliance_ticket,
			)
		}
	}
	tbl.Print()
}
