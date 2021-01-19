/**
  @author: xinyulu
  @date: 2021/1/19 11:05
  @note: find file
**/
package commons

import (
    "bytes"
    "errors"
    "os"
    "os/exec"
    "os/user"
    "runtime"
    "strings"
)

// 获取用户家目录
func HomeDir() (string, error) {
    userInfo, err := user.Current()
    if err == nil {
        return userInfo.HomeDir, nil
    }
    if runtime.GOOS == "windows" {
        return homeWindows()
    }
    return homeUnix()
}

// 获取Unix家目录
func homeUnix() (string ,error) {
    if home := os.Getenv("HOME"); home != "" {
        return home, nil
    }
    var stdout bytes.Buffer
    cmd := exec.Command("sh", "-c", "eval echo ~$USER")
    cmd.Stdout = &stdout
    if err := cmd.Run(); err != nil {
        return "", err
    }
    result := strings.TrimSpace(stdout.String())
    if result == "" {
        return "", errors.New("No home directory!!")
    }
    return result, nil
}

// 获取windows家目录
func homeWindows() (string, error) {
    drive := os.Getenv("HOMEDRIVE")
    path := os.Getenv("HOMEPATH")
    home := drive + path
    if drive == "" || path == "" {
        home = os.Getenv("USERPROFILE")
    }
    if home == "" {
        return "", errors.New("ENV: HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
    }
    return home, nil
}
