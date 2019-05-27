package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func main() {
  client, err := ecs.NewClientWithAccessKey("cn-hangzhou", "LTAIK61t4t7qwRuv", "YA6chzF0tiwPAxtx47MkOBlckJtxEz")

  request := ecs.CreateDescribeAccessPointsRequest()

  response, err := client.DescribeAccessPoints(request)
  if err != nil {
    fmt.Print(err.Error())
  }
  fmt.Printf("response is %#v\n", response)
}