
var checkRow = (board, pos) => {
    for (let x = 0; x < board.length; x++) {
        if (board[x][pos[1]] == "Q"){
            return false;
        }
    }
    return true;
}
var checkCol = (board, pos) => {
    for (let x = 0; x < board.length; x++) {
        if (board[pos[0]][x] == "Q"){
            return false;
        }
    }
    return true;
}
var checkDiag = (board, pos) => {
    let x = pos[0];
    let y = pos[1];
    while(x > 0 && y > 0) {
        if ((board[x-1][y-1]) == "Q"){
            return false;
        }
        x--;y--;
    }
    while(x<board.length-1 && y < board.length-1) {
        if ((board[x+1][y+1]) == "Q"){
            return false;
        }
        x++;y++;
    }
    return true;
}

function validateMove(board, pos) {
    return checkRow(board, pos) && checkCol(board, pos) && checkDiag(board, pos);
}

function initBoard(n) {
    let board = [];
    for (let i = 0; i < n; i++){
        let row = [];
        for (let j =0; j < n; j++){
            row.push(".");
        }
        board.push(row);
    }
    return board;
}

function solve(board,row,col){
    console.log("solving", board, row, col);
    board[row][col] = "Q";
    if (row == board.length && col == board.length) {
        return true;
    }
    for (let i = 0; i < board.length; i ++) {
        for(let j = 0; j < board.length; j++){
            if (validateMove(board, [i,j])) {
                return solve(board,i,j);
            }
        }
    }
    return false;
}

var solveNQueens = function(n) {
    let result = [];
    for (let i = 0; i < n; i++) {
        let b = initBoard(n);
        if (solve(b,i,0)){
            result.push(b);
        }else {
          console.log("No solution at ", i);
        }

    }
    return result;
};


console.log(solveNQueens(4));
