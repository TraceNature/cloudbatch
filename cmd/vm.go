/**
  @author: xinyulu
  @date: 2021/1/18 11:38
  @note: 
**/
package cmd

import (
    "github.com/spf13/cobra"
)

func NewVmCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use: "vm <subcommand>",
        Short: "vm batch",
    }
    cmd.AddCommand(NewCreateVmInstancesCommand())
    return cmd
}

func NewCreateVmInstancesCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use: "create -f FILE -t FILETYPE",
        Short: "create instances by FILE",
        Run: createVmInstancesHandleFunc,
    }
    cmd.Flags().StringP("file", "f", "vm_instances.yaml", "file parameter of creating vm")
    return cmd
}

// 创建vm实例处理函数
func createVmInstancesHandleFunc(cmd *cobra.Command, args []string) {
}