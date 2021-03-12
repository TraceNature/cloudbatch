package commons

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
)

//AppendLineToFile 向文件追加行
func AppendLineToFile(line *bytes.Buffer, filename string) {
    lock.Lock()
    defer lock.Unlock()
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }

    defer f.Close()
    w := bufio.NewWriter(f)
    fmt.Fprintln(w, line.String())
    w.Flush()
}

func WriteFile(content []byte) error {
    err := ioutil.WriteFile("output.txt", content, 0666)
    return err
}

// Exists 用于判断所给路径文件或文件夹是否存在
func FileExists(path string) bool {
    _, err := os.Stat(path) //os.Stat获取文件信息
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
    s, err := os.Stat(path)
    if err != nil {
        return false
    }
    return s.IsDir()
}

//IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
    return !IsDir(path)
}

//复制文件
func CopyFile(src string, dst string, buffersize int) error {
    buf := make([]byte, buffersize)
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()
    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()
    for {
        n, err := source.Read(buf)
        if err != nil && err != io.EOF {
            return err
        }
        if n == 0 {
            break
        }

        if _, err := destination.Write(buf[:n]); err != nil {
            return err
        }
    }
    return nil
}

// 判断文件是否存在，不存在则创建
func CreateFile(filePath string) (*os.File,error) {
    dir := filepath.Dir(filePath)
    if err := MkdirAll(dir, 0755); err != nil {
        return nil, err
    }
    file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    return file, nil
}

// 判断目录是否存在，如果不存在则创建
func MkdirAll(dir string, mode os.FileMode) error {
    if FileExists(dir) {
        return nil
    }
    err := os.MkdirAll(dir, mode)
    if err != nil {
        return err
    }
    return nil
}