<?php

// 给定一个包含 n 个整数的数组 S，是否存在属于 S 的三个元素 a，b，c 使得 a + b + c = 0 ？
// 找出所有不重复的三个元素组合使三个数的和为零。
// 注意：结果不能包括重复的三个数的组合。
// 例如, 给定数组 S = [-1, 0, 1, 2, -1, -4]，
// 一个结果集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

function threeSum($nums) {
    // 排序
    sort($nums);

    $return = [];

    $count = count($nums);
    for ($left = 0; $left < $count; $left++) {
        $mid = $left + 1; $right = $count - 1;
        $sub = 0 - $nums[$left];

        // 跳过 left 相同
        if ($left > 0 && $nums[$left] == $nums[$left-1]) {
            continue;
        }

        while ($mid < $right) {
            echo $mid, " ", $right, "\n";
            if ($nums[$mid] + $nums[$right] == $sub) {
                $tmp_mid = $nums[$mid];
                $tmp_right = $nums[$right];
                $return[] = [$nums[$left], $tmp_mid, $tmp_right];
                while ($mid < $right && $nums[++$mid] == $tmp_mid) {}   // 跳过mid 相同
                while ($mid < $right && $nums[--$right] == $tmp_right) {}   // 跳过 right 相同
            } else if ($nums[$mid] + $nums[$right] > $sub) {
                $right--;
                continue;
            } else {
                $mid++;
            }
        }
    }
    return $return;
}

$arr = [
    -2, -2, -1, -1, 0, 0, 1, 1, 2, 2,
    // -1, 0, 1, 2, -1, -4,
];

print_r(threeSum($arr));