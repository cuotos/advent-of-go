totals = {}
runningTotal = 0

while True:
    for l in ["+1000000", "-999999"]:
        runningTotal += int(l)
        if not totals.get(runningTotal, False):
            totals[runningTotal] = True
            print(runningTotal)
        else:
            print(runningTotal)
            exit(0)