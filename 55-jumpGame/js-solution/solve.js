/**
 * @param {number[]} nums
 * @return {boolean}
 */
function max(a,b) {
    if (a >= b){
        return a;
    }
    return b;
}

var canJump = function(nums) {
    return jump(nums, 0);
};

function jump(nums) {
    let farest = 0;
    for (let i = 0; i < nums.length-1; i ++){
        farest = max(farest, i + nums[i]);

        if (i == farest && nums[i] == 0){
            return false;
        }
    }
    return true;
}

let a = [2,0,0];
console.log(canJump(a));
let b = [3,2,1,1,4];
console.log(canJump(b));
