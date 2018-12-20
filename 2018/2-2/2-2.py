f = open("input")

for line in f.readlines():
    containsTwo = False
    containsThree = False

    letterCount = {}

    for letter in line.rstrip():
        if not letterCount.get(letter, ""):
            letterCount[letter] = 1
        else:
            letterCount[letter] += 1

    for l, c in letterCount.items():
        if c == 2:
            containsTwo = True
        if c == 3:
            containsThree = True

    break