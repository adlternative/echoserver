package main

import (
    "io"
    "log"
    "net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
    // 设置响应头 Content-Type 与请求的相同，如果有的话
    if ct := r.Header.Get("Content-Type"); ct != "" {
        w.Header().Set("Content-Type", ct)
    }

    // 将请求的方法和URL打印到服务器日志（可选）
    log.Printf("Received %s request for %s", r.Method, r.URL.Path)

    // 将请求体复制到响应体
    if r.Body != nil {
        defer r.Body.Close()
        _, err := io.Copy(w, r.Body)
        if err != nil {
            http.Error(w, "Error reading request body", http.StatusInternalServerError)
            return
        }
    } else {
        // 如果没有请求体，返回空内容
        w.WriteHeader(http.StatusOK)
    }
}

func main() {
    // 设置路由，将所有路径都交给 echoHandler 处理
    http.HandleFunc("/", echoHandler)

    // 启动服务器，监听端口 8089
    addr := ":8089"
    log.Printf("Starting echo server on port %s", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
