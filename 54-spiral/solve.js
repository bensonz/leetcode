/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
// right, down, up, left
const dirs = [[0,1],[1,0],[0,-1],[-1,0]];

let buildGraph = (width, height) => {
    let rg = [];
    for (let i = 0; i < height; i++) {
        let row = [];
        for (let j = 0; j < width; j++){
           row.push(0);
        }
        rg.push(row);
    }
    return rg;
};

let isValidMove= (mat,r,c,dirI)=>{
    let aa = r + dirs[dirI][0];
    let bb = c + dirs[dirI][1];
    let h = mat.length;
    let w = mat[0].length;
    return aa >= 0 && aa < h && bb >= 0 && bb < w && mat[aa][bb] != 1;
};

function move(g, a, b, di) {
    if (isValidMove(g,a,b,di)) {
      console.log("continue cur dir");
        return {
            row: a + dirs[di][0],
            col: b + dirs[di][1],
            dirI: di,
            ok: true
        }
    }
    let ndi = (di + 1) % 4;
    if (isValidMove(g,a,b,ndi)) {
      console.log("Next dir:", ndi);
        return {
            row: a + dirs[ndi][0],
            col: b + dirs[ndi][1],
            dirI: ndi,
            ok: true
        }
    }
    return {
        row: a,
        col: b,
        dirI: di,
        ok: false
    }
}
var spiralOrder = function(matrix) {
    let heigth = matrix.length;
    let width = matrix[0].length;
    let cleanGraph = buildGraph(width,heigth);
    let cur = {
        row: 0,
        col:0
    };
    let dirI = 0;
    let result = [];
    while(true) {
        let {row, col} = cur
        console.log("Moving", row, col, dirI);
        let r = move(cleanGraph,row,col,dirI)
        if (r.ok){
            cur.row = r.row;
            cur.col = r.col;
            dirI = r.dirI;
            cleanGraph[row][col] = 1;
            result.push(matrix[row][col]);
        }else{
            break;
        }
    }
    result.push(matrix[cur.row][cur.col]);
    return result;
};

let a = [[1,2,3],
[4,5,6],
[7,8,9]]
let b = [[1,2,3,4],
[5,6,7,8],
[9,10,11,12]]
console.log(spiralOrder(b));
