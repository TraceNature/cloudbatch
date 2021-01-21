/**
  @author: xinyulu
  @date: 2021/1/21 10:41
  @note: 
**/
package cmd

import (
    "cloudbatch/client"
    "cloudbatch/log"
    "fmt"
    "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/apis"
    lbClient "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/client"
    "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/models"
    "github.com/spf13/cobra"
)

func NewLbCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use: "lb <subcommand>",
        Short: "lb batch",
    }
    cmd.AddCommand(NewCheckUnhealthyBackend())
    return cmd
}

func NewCheckUnhealthyBackend() *cobra.Command {
    cmd := &cobra.Command{
        Use: "checkhealthy -r REGION",
        Short: "check unhealthy backends of all lbs",
        Run: getUnhealthyBackend,
    }
    cmd.Flags().StringP("region", "r", "cn-north-1", "")
    return cmd
}

// 获取所有Backends
func getBackends(regionId string, lbclient *lbClient.LbClient) ([]models.Backend, error) {
    pageSize := 100
    backends := make([]models.Backend, 0)
    req := apis.NewDescribeBackendsRequest(regionId)
    req.SetPageNumber(1)
    req.SetPageSize(10)
    resp, err := lbclient.DescribeBackends(req)
    if err != nil{
        return nil, err
    }
    // 保存Backends的总数
    var counts int
    if resp.Result.TotalCount % pageSize == 0 {
        counts =resp.Result.TotalCount / pageSize
    } else {
        counts = resp.Result.TotalCount / pageSize + 1
    }

    for i := 1; i <= counts; i++ {
        req := apis.NewDescribeBackendsRequest(regionId)
        req.SetPageNumber(i)
        req.SetPageSize(pageSize)
        resp, err := lbclient.DescribeBackends(req)
        if err != nil {
            return nil, err
        }
        backends = append(backends, resp.Result.Backends...)
    }
    return backends, nil
}


func getUnhealthyBackend(cmd *cobra.Command, args []string) {
    lbclient := client.GetLbClient()
    regionId, err := cmd.Flags().GetString("region")
    if err != nil {
        log.Error(err.Error())
        return
    }
    backends, err := getBackends(regionId, lbclient)
    if err != nil {
        log.Error(err.Error())
        return
    }
    for _, backend := range backends {
        req := apis.NewDescribeTargetHealthRequest(regionId, backend.BackendId)
        resp, err := lbclient.DescribeTargetHealth(req)
        if err != nil {
            log.Error(err.Error())
            return
        }
        for _, targetHealthy := range resp.Result.TargetHealths {
            if targetHealthy.Status == "unhealthy" {
                fmt.Println(backend.LoadBalancerId, backend.BackendName, targetHealthy.InstanceId, targetHealthy.Status)
            }
        }
    }
}
