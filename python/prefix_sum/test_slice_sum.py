import unittest

from range_sum_query import NumArray, NumArray2
from slice_sum import slice_sum

from python.prefix_sum.contigious_array import find_max_length


class TestPrefixSum(unittest.TestCase):
    def setUp(self) -> None:
        self.num_array = NumArray([-2, 0, 3, -5, 2, -1])
        self.num_array2 = NumArray2([-2, 0, 3, -5, 2, -1])

    def test_slice_sum(self) -> None:
        args_list = [
            ([1, 2, 3, 4, 5, 6], 1, 3),
            ([1, 2, 3, 4, 5, 6], 0, 3),
        ]
        expected_list = [9, 10]
        for args, expected in zip(args_list, expected_list):
            assert slice_sum(*args) == expected

    def test_sum_range(self) -> None:
        args_list = [(0, 2), (2, 5), (0, 5)]
        expected_list = [1, -1, -3]
        for args, expected in zip(args_list, expected_list):
            assert self.num_array.sum_range(*args) == expected

    def test_sum_range2(self) -> None:
        args_list = [(0, 2), (2, 5), (0, 5)]
        expected_list = [1, -1, -3]
        for args, expected in zip(args_list, expected_list):
            assert self.num_array2.sum_range(*args) == expected

    def test_contigious_array(self) -> None:
        args_list = [
            [0, 1],
            [0, 1, 0],
            [1, 0, 1, 0, 1, 1, 1, 0],
            [0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0],
        ]
        expected_list = [2, 2, 4, 20]
        for args, expected in zip(args_list, expected_list):
            assert find_max_length(args) == expected
