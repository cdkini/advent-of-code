def read_input_to_list():
    output = []
    with open("input.txt", "r") as f:
        for line in f:
            output.append(int(line.strip()))
    return output

def part_one(nums):
    res = 0
    for num in nums:
        res += num // 3 - 2
    print(f"Total mass is {res}")
    return res

def part_two(nums):
    res = 0
    for num in nums:
        fuel = num
        while fuel > 0:
            temp = fuel // 3 - 2
            if temp <= 0:
                break
            res += temp
            fuel = temp
    print(f"Total mass is {res}")
    return res

def a():
    print("a")

def main():
    nums = read_input_to_list()
    part_one(nums)
    part_two(nums)

if __name__ == "__main__":
    main()
