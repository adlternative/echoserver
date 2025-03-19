package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
    "strings"
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
    http.HandleFunc("/api/echo", echoHandler)
    
    // 添加静态文件服务
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // 添加游戏API路由
    http.HandleFunc("/api/game/", gameHandler)
    
    // 添加主页路由
    http.HandleFunc("/", indexHandler)

    // 启动服务器，监听端口 8089
    addr := ":8089"
    log.Printf("Starting chess game server on port %s", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    http.ServeFile(w, r, "./static/index.html")
}

// 游戏状态
type GameState struct {
    Board      [10][9]string `json:"board"`
    CurrentTurn string       `json:"currentTurn"`
    GameOver   bool          `json:"gameOver"`
    Winner     string        `json:"winner"`
}

var (
    // 当前游戏状态
    currentGame GameState
    // 游戏是否已初始化
    gameInitialized bool
)

func gameHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // 根据路径处理不同的游戏操作
    path := strings.TrimPrefix(r.URL.Path, "/api/game/")
    
    switch path {
    case "state":
        // 获取游戏状态
        if !gameInitialized {
            initializeGame()
        }
        json.NewEncoder(w).Encode(currentGame)
        
    case "move":
        // 处理棋子移动
        var move struct {
            FromX int `json:"fromX"`
            FromY int `json:"fromY"`
            ToX   int `json:"toX"`
            ToY   int `json:"toY"`
        }
        
        if err := json.NewDecoder(r.Body).Decode(&move); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        
        // 验证并执行移动
        if isValidMove(move.FromX, move.FromY, move.ToX, move.ToY) {
            executeMove(move.FromX, move.FromY, move.ToX, move.ToY)
            json.NewEncoder(w).Encode(currentGame)
        } else {
            http.Error(w, "Invalid move", http.StatusBadRequest)
        }
        
    case "reset":
        // 重置游戏
        initializeGame()
        json.NewEncoder(w).Encode(currentGame)
        
    default:
        http.NotFound(w, r)
    }
}

// 初始化游戏
func initializeGame() {
    // 清空棋盘
    for i := range currentGame.Board {
        for j := range currentGame.Board[i] {
            currentGame.Board[i][j] = ""
        }
    }
    
    // 设置中国象棋部分（红方）
    currentGame.Board[0][0] = "r-車"
    currentGame.Board[0][1] = "r-馬"
    currentGame.Board[0][2] = "r-象"
    currentGame.Board[0][3] = "r-士"
    currentGame.Board[0][4] = "r-帥"
    currentGame.Board[0][5] = "r-士"
    currentGame.Board[0][6] = "r-象"
    currentGame.Board[0][7] = "r-馬"
    currentGame.Board[0][8] = "r-車"
    currentGame.Board[2][1] = "r-炮"
    currentGame.Board[2][7] = "r-炮"
    currentGame.Board[3][0] = "r-兵"
    currentGame.Board[3][2] = "r-兵"
    currentGame.Board[3][4] = "r-兵"
    currentGame.Board[3][6] = "r-兵"
    currentGame.Board[3][8] = "r-兵"
    
    // 设置国际象棋部分（黑方）
    currentGame.Board[9][0] = "b-♜"
    currentGame.Board[9][1] = "b-♞"
    currentGame.Board[9][2] = "b-♝"
    currentGame.Board[9][3] = "b-♛"
    currentGame.Board[9][4] = "b-♚"
    currentGame.Board[9][5] = "b-♝"
    currentGame.Board[9][6] = "b-♞"
    currentGame.Board[9][7] = "b-♜"
    currentGame.Board[8][0] = "b-♟"
    currentGame.Board[8][1] = "b-♟"
    currentGame.Board[8][2] = "b-♟"
    currentGame.Board[8][3] = "b-♟"
    currentGame.Board[8][4] = "b-♟"
    currentGame.Board[8][5] = "b-♟"
    currentGame.Board[8][6] = "b-♟"
    currentGame.Board[8][7] = "b-♟"
    currentGame.Board[8][8] = "b-♟"
    
    // 设置游戏状态
    currentGame.CurrentTurn = "r" // 红方先行
    currentGame.GameOver = false
    currentGame.Winner = ""
    
    gameInitialized = true
}

// 验证移动是否合法
func isValidMove(fromX, fromY, toX, toY int) bool {
    // 检查坐标是否在棋盘范围内
    if fromX < 0 || fromX >= 10 || fromY < 0 || fromY >= 9 ||
       toX < 0 || toX >= 10 || toY < 0 || toY >= 9 {
        return false
    }
    
    // 检查起始位置是否有棋子
    piece := currentGame.Board[fromX][fromY]
    if piece == "" {
        return false
    }
    
    // 检查是否是当前玩家的棋子
    pieceColor := string(piece[0])
    if pieceColor != currentGame.CurrentTurn {
        return false
    }
    
    // 检查目标位置是否是自己的棋子
    targetPiece := currentGame.Board[toX][toY]
    if targetPiece != "" && string(targetPiece[0]) == currentGame.CurrentTurn {
        return false
    }
    
    // 这里简化了移动规则，实际实现中需要根据不同棋子类型检查移动规则
    // 在完整实现中，需要为每种棋子类型实现特定的移动规则
    
    return true
}

// 执行移动
func executeMove(fromX, fromY, toX, toY int) {
    // 获取棋子
    piece := currentGame.Board[fromX][fromY]
    
    // 检查是否吃掉对方的王/帅，结束游戏
    targetPiece := currentGame.Board[toX][toY]
    if targetPiece == "r-帥" || targetPiece == "b-K" {
        currentGame.GameOver = true
        currentGame.Winner = string(piece[0])
    }
    
    // 移动棋子
    currentGame.Board[toX][toY] = piece
    currentGame.Board[fromX][fromY] = ""
    
    // 切换回合
    if currentGame.CurrentTurn == "r" {
        currentGame.CurrentTurn = "b"
    } else {
        currentGame.CurrentTurn = "r"
    }
}
