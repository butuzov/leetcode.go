def containsNearbyAlmostDuplicate(nums, k, t):
    if t < 0: return False
    n = len(nums)
    d = {}
    w = t + 1
    for i in range(n):
        m = nums[i] / w
        if m in d:
            return True
        if m - 1 in d and abs(nums[i] - d[m - 1]) < w:
            return True
        if m + 1 in d and abs(nums[i] - d[m + 1]) < w:
            return True

        print(nums[i], d)

        d[m] = nums[i]



        if i >= k: del d[nums[i - k] / w]
    return False


print(containsNearbyAlmostDuplicate([2,1], 1, 1))
