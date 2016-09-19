#!/usr/bin/env python3
# https://www.coursera.org/learn/algorithm-design-analysis/home/week/1
# Programming assignment

import sys


def merge_sort(lst):
    """
    The implementation of the merge sort algorithm
    """

    def merge(lst1, lst2):
        """Merge two lists"""
        i, j = 0, 0
        len1 = len(lst1)
        len2 = len(lst2)
        merged = [None] * (len1 + len2)  # allocate once
        inversions = 0
        while i < len1 or j < len2:
            x1 = lst1[i] if i < len1 else None
            x2 = lst2[j] if j < len2 else None
            if x1 is None or (x2 is not None and x1 > x2):
                merged[i + j] = x2
                j += 1
                if x1 is not None:
                    inversions += len1 - i
            elif x2 is None or x1 <= x2:
                merged[i + j] = x1
                i += 1
        return merged, inversions

    l = len(lst)
    if l <= 1:
        return lst, 0
    elif l == 2:
        inv = 0
        if lst[0] > lst[1]:
            lst.sort()  # don't act derp
            inv = 1
        return lst, inv
    else:
        div = l // 2
        lst1, inv1 = merge_sort(lst[:div])
        lst2, inv2 = merge_sort(lst[div:])
        merged, inv3 = merge(lst1, lst2)
        inversions = inv1 + inv2 + inv3
        return merged, inversions


def test():
    cases = [{
        'name': 'Case 1',
        'nums': [1, 3, 5, 2, 4, 6],
        'expected': 3
    }, {
        'name': 'Case 2',
        'nums': [1, 5, 3, 2, 4],
        'expected': 4,
    }, {
        'name': 'Case 3',
        'nums': [5, 4, 3, 2, 1],
        'expected': 10
    }, {
        'name': 'Case 4',
        'nums': [1, 6, 3, 2, 4, 5],
        'expected': 5
    }, {
        'name': 'Case 5',
        'nums': [9, 12, 3, 1, 6, 8, 2, 5, 14, 13, 11, 7, 10, 4, 0],
        'expected': 56
    }, {
        'name': 'Case 6',
        'nums':
        [37, 7, 2, 14, 35, 47, 10, 24, 44, 17, 34, 11, 16, 48, 1, 39, 6, 33,
         43, 26, 40, 4, 28, 5, 38, 41, 42, 12, 13, 21, 29, 18, 3, 19, 0, 32,
         46, 27, 31, 25, 15, 36, 20, 8, 9, 49, 22, 23, 30, 45],
        'expected': 590
    }, {
        'name': 'Case 7',
        'nums':
        [4, 80, 70, 23, 9, 60, 68, 27, 66, 78, 12, 40, 52, 53, 44, 8, 49, 28,
         18, 46, 21, 39, 51, 7, 87, 99, 69, 62, 84, 6, 79, 67, 14, 98, 83, 0,
         96, 5, 82, 10, 26, 48, 3, 2, 15, 92, 11, 55, 63, 97, 43, 45, 81, 42,
         95, 20, 25, 74, 24, 72, 91, 35, 86, 19, 75, 58, 71, 47, 76, 59, 64,
         93, 17, 50, 56, 94, 90, 89, 32, 37, 34, 65, 1, 73, 41, 36, 57, 77, 30,
         22, 13, 29, 38, 16, 88, 61, 31, 85, 33, 54],
        'expected': 2372
    }]

    for case in cases:
        _, inversions = merge_sort(case['nums'])
        assert inversions == case['expected'], \
                '%s failed: expected %d inversions, got %d' % (
                    case['name'], case['expected'], inversions)


def list_of_nums(path):
    try:
        with open(path) as f:
            return [int(line) for line in f.readlines()]
    except:
        print('Error: Cannot read data, check the file')
        sys.exit(1)


if __name__ == '__main__':
    # test()

    if len(sys.argv) > 1:
        path = sys.argv[1]
        nums = list_of_nums(path)
        srt, inversions = merge_sort(nums)
        print('Found %d inversions in %s' % (inversions, path))
