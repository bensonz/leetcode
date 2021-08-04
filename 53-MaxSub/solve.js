var maxSubArray = function(nums) {
    let sum = nums[0];
    let curSum = nums[0];
    for (let i = 1;i < nums.length; i++) {
        console.log(`sum ${sum}, curSum:${curSum}, nums[i]:${nums[i]}`);
        if (curSum >= 0) {
            curSum += nums[i];
        }else {
            curSum = nums[i];
        }
        if (sum < curSum){
            sum = curSum;
        }

    }
    return sum;
};

console.log(maxSubArray([-2,1,-3,4,-1,2,1,-5,4]));
