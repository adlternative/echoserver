document.addEventListener('DOMContentLoaded', () => {
    const chessBoard = document.getElementById('chess-board');
    const turnValue = document.getElementById('turn-value');
    const resetButton = document.getElementById('reset-button');
    
    let gameState = null;
    let selectedCell = null;
    
    // 初始化棋盘
    function initializeBoard() {
        chessBoard.innerHTML = '';
        
        // 创建棋盘格子
        for (let i = 0; i < 10; i++) {
            for (let j = 0; j < 9; j++) {
                const cell = document.createElement('div');
                cell.className = 'cell';
                cell.dataset.x = i;
                cell.dataset.y = j;
                cell.addEventListener('click', handleCellClick);
                chessBoard.appendChild(cell);
            }
        }
        
        // 获取游戏状态
        fetchGameState();
    }
    
    // 获取游戏状态
    function fetchGameState() {
        fetch('/api/game/state')
            .then(response => response.json())
            .then(data => {
                gameState = data;
                updateBoard();
                updateTurnIndicator();
            })
            .catch(error => console.error('Error fetching game state:', error));
    }
    
    // 更新棋盘显示
    function updateBoard() {
        const cells = document.querySelectorAll('.cell');
        
        cells.forEach(cell => {
            const x = parseInt(cell.dataset.x);
            const y = parseInt(cell.dataset.y);
            
            // 清除之前的棋子
            cell.innerHTML = '';
            
            // 如果有棋子，则显示
            const piece = gameState.board[x][y];
            if (piece) {
                const pieceDiv = document.createElement('div');
                const [color, type] = piece.split('-');
                
                pieceDiv.className = `piece ${color === 'r' ? 'red' : 'black'}`;
                pieceDiv.textContent = type;
                
                cell.appendChild(pieceDiv);
            }
        });
        
        // 检查游戏是否结束
        if (gameState.gameOver) {
            const winner = gameState.winner === 'r' ? '红方' : '黑方';
            setTimeout(() => {
                alert(`游戏结束！${winner}获胜！`);
            }, 100);
        }
    }
    
    // 更新回合指示器
    function updateTurnIndicator() {
        turnValue.textContent = gameState.currentTurn === 'r' ? '红方' : '黑方';
    }
    
    // 处理格子点击事件
    function handleCellClick(event) {
        const cell = event.currentTarget;
        const x = parseInt(cell.dataset.x);
        const y = parseInt(cell.dataset.y);
        
        // 如果游戏已结束，不处理点击
        if (gameState.gameOver) {
            return;
        }
        
        // 如果已经选中了一个格子
        if (selectedCell) {
            const fromX = parseInt(selectedCell.dataset.x);
            const fromY = parseInt(selectedCell.dataset.y);
            
            // 如果点击的是同一个格子，取消选择
            if (fromX === x && fromY === y) {
                selectedCell.classList.remove('selected');
                selectedCell = null;
                return;
            }
            
            // 尝试移动棋子
            movePiece(fromX, fromY, x, y);
            
            // 清除选中状态
            selectedCell.classList.remove('selected');
            selectedCell = null;
            
            // 清除所有有效移动标记
            document.querySelectorAll('.valid-move').forEach(cell => {
                cell.classList.remove('valid-move');
            });
        } else {
            // 检查是否选中了当前回合的棋子
            const piece = gameState.board[x][y];
            if (piece && piece.startsWith(gameState.currentTurn)) {
                selectedCell = cell;
                cell.classList.add('selected');
                
                // 这里可以添加显示有效移动的逻辑
                // highlightValidMoves(x, y, piece);
            }
        }
    }
    
    // 移动棋子
    function movePiece(fromX, fromY, toX, toY) {
        const moveData = {
            fromX: fromX,
            fromY: fromY,
            toX: toX,
            toY: toY
        };
        
        fetch('/api/game/move', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(moveData)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Invalid move');
            }
            return response.json();
        })
        .then(data => {
            gameState = data;
            updateBoard();
            updateTurnIndicator();
        })
        .catch(error => {
            console.error('Error moving piece:', error);
            // 可以添加一些视觉反馈，表明移动无效
        });
    }
    
    // 重置游戏
    function resetGame() {
        fetch('/api/game/reset')
            .then(response => response.json())
            .then(data => {
                gameState = data;
                updateBoard();
                updateTurnIndicator();
                
                // 清除选中状态
                if (selectedCell) {
                    selectedCell.classList.remove('selected');
                    selectedCell = null;
                }
            })
            .catch(error => console.error('Error resetting game:', error));
    }
    
    // 添加重置按钮事件监听器
    resetButton.addEventListener('click', resetGame);
    
    // 初始化棋盘
    initializeBoard();
});
