def read_file():
    with open('../input/day1', 'r') as f:
            content = f.readlines()
    return content

def _sum_lines(lines):
    sums = []
    currMax = 0
    for line in lines:
        line = line.strip()
        if line:
            currMax += int(line)
        else:
            sums.append(currMax)
            currMax = 0
    return sums



def part1():
    content = read_file()
    return max(_sum_lines(content))

print(part1())

def part2():
    content = read_file()
    sort = sorted(_sum_lines(content))
    return sum(sort[-3:])

print(part2())

