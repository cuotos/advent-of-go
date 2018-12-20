

f = open("input")

total = 0

for l in f.readlines():
    total += int(l)

print(total)