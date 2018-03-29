//Utils
package core

import (
    "io"
    "os"
    "log"
    "fmt"
    "bytes"
    "bufio"
    "os/exec"
)

func CountLine(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}

func ReadFromFile(filePath string){

    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalln(err.Error() + `: ` + filePath)
        return
    } else {
        defer file.Close()
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

}

func Readln(r *bufio.Reader) (string, error) {

    var (isPrefix bool = true
        err error = nil
        line, ln []byte
    )
    for isPrefix && err == nil {
        line, isPrefix, err = r.ReadLine()
        ln = append(ln, line...)
    }
    return string(ln),err

}

func Execute (pathExec string, args []string) (string, error) {

    path, err := exec.LookPath(pathExec)
    if err != nil {
        return "", err
    }
    cmd, err := exec.Command(path, args...).CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(cmd), nil

}
