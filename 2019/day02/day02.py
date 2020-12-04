def read_input():
    with open("input.txt", "r") as f:
        for line in f:
            nums = [int(x) for x in line.split(",")]
            del nums[-1]
            return nums

def part_one(nums, noun, verb):
    nums[1] = noun
    nums[2] = verb
    i = 0
    while i < len(nums):
        if nums[i] == 1:
            nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
        elif nums[i] == 2:
            nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
        else:
            break
        i += 4
    return nums[0]

def part_two(nums):
    for noun in range(1000):
        for verb in range(1000):
            res = part_one(nums[:], noun, verb)
            if res == 19690720:
                print(f"Part 2 Intcode: {100 * noun + verb}")
                return 100 * noun + verb

def main():
    nums = read_input()
    part_one(nums, 12, 2)
    part_two(nums)


if __name__ == "__main__":
    main()
