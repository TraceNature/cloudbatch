/**
  @author: xinyulu
  @date: 2021/1/20 17:58
  @note: 
**/
package commons

import (
    "fmt"
    "testing"
)

func TestCreateFile(t *testing.T) {
    file, err := CreateFile("E:/a/b/c/")
    defer file.Close()
    if err != nil {
        fmt.Println(err)
    }
}
