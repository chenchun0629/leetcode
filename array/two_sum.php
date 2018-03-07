<?php

// 给定一个整数数列，找出其中和为特定值的那两个数。
// 你可以假设每个输入都只会有一种答案，同样的元素不能被重用。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
//


# 双重循环
function twoSum($nums, $target) {
    $return = [];

    do {
        $first = array_shift($nums);
        // if ($first > $target) break;
        $other = $nums;

        do {
            $second = array_shift($other);
            // if ($second > $target) break;
            $sum = $second + $first;
            // if ($sum > $target) break;
            if ($sum == $target) $return[] = [$first, $second];
        } while (count($other));
    } while (count($nums));

    return $return;
}

# hashmap
function twoSumHashmap($nums, $target)
{
    $return = [];
    $nums = array_flip($nums);

    foreach ($nums as $key => $value) {
        if ($key > $target) {
            continue;
        }
        $sub = $target - $key;
        if (isset($nums[$sub])) {
            $return[] = [$key, $sub];
            unset($nums[$key], $nums[$sub]);
        }
    }

    return $return;
}

$nums   = [2, 15, 3, 11, 6, 7];
$target = 9;

print_r(twoSum($nums, $target));
print_r(twoSumHashmap($nums, $target));