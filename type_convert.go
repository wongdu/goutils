package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type APIBuilder struct {
	api int
}

type Router struct {
	router string
}

func (r *APIBuilder) Use() {
	fmt.Println("APIBuilder Use...")
}

func (r *Router) Use() {
	fmt.Println("Route Use...")
}

type Application struct {
	*APIBuilder
	*Router
}

func main() {
	app := &Application{
		APIBuilder: &APIBuilder{
			api: 123,
		},
		Router: &Router{
			router: "abc",
		},
	}
	//app.Use()
	_ = app
	fmt.Println("in main func")
	fmt.Println(getCaller())
	var AllMethods = []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPatch,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodConnect,
		http.MethodTrace,
	}
	for v := range AllMethods {
		fmt.Println(v, AllMethods[v])
	}
	time.Sleep(1)

	mm := int(32.0)
	fmt.Println(mm)
	f := 3.2
	fmt.Println(int32(f)) //ok
	//fmt.Println(int(3.2)) //error
	tt("a", "b", "c")
}

func tt(p ...string) {
	fmt.Println([]string(p)) //ok
	//fmt.Println([]string("a", "b", "c")) //error
}
func getCaller() (string, int) {
	var pcs [32]uintptr
	n := runtime.Callers(1, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	wd, _ := os.Getwd()

	var (
		frame runtime.Frame
		more  = true
	)

	for {
		if !more {
			break
		}

		frame, more = frames.Next()
		file := filepath.ToSlash(frame.File)
		// fmt.Printf("%s:%d | %s\n", file, frame.Line, frame.Function)

		if strings.Contains(file, "go/src/runtime/") {
			continue
		}

		if !strings.Contains(file, "_test.go") {
			if strings.Contains(file, "/kataras/iris") &&
				!strings.Contains(file, "kataras/iris/_examples") &&
				!strings.Contains(file, "kataras/iris/middleware") &&
				!strings.Contains(file, "iris-contrib/examples") {
				continue
			}
		}

		if relFile, err := filepath.Rel(wd, file); err == nil {
			if !strings.HasPrefix(relFile, "..") {
				// Only if it's relative to this path, not parent.
				file = "./" + relFile
			}
		}

		return file, frame.Line
	}

	return "???", 0
}

func splitPath(pathMany string) (paths []string) {
	pathMany = strings.Trim(pathMany, " ")
	pathsWithoutSlashFromFirstAndSoOn := strings.Split(pathMany, "/")
	fmt.Println("pathsWithoutSlashFromFirstAndSoOn is:", pathsWithoutSlashFromFirstAndSoOn)
	for _, path := range pathsWithoutSlashFromFirstAndSoOn {
		if path == "" {
			continue
		}
		if path[0] != '/' {
			path = "/" + path
		}
		paths = append(paths, path)
	}
	return
}

func ConvertSecond(spec string) string {
	if spec != "" {
		_spec := strings.Split(spec, " ")
		hour, minute, second := time.Now().Clock()

		// second处理
		_second := strconv.FormatInt(int64(second), 10) + "/60"
		spec = strings.Replace(spec, _spec[0], _second, 1)

		// minute处理(类似: */5)
		if strings.HasPrefix(_spec[1], "*/") {
			_minute := strconv.FormatInt(int64(minute), 10) + "/" + strings.Trim(_spec[1], "*/")
			spec = strings.Replace(spec, _spec[1], _minute, 1)
		}

		// hour处理(类似: */5)
		if strings.HasPrefix(_spec[2], "*/") {
			_hour := strconv.FormatInt(int64(hour), 10) + "/" + strings.Trim(_spec[2], "*/")
			spec = strings.Replace(spec, _spec[2], _hour, 1)
		}

		return spec
	}

	return ""
}
