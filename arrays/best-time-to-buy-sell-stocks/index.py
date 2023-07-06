stocks = [1, 1, 1, 3, 1]


def calc_best_stock(stocks):
    min_price = float("inf")
    max_price = 0
    for i in range(len(stocks)):
        if stocks[i] < min_price:
            min_price = stocks[i]
            continue
        if stocks[i] - min_price > max_price:
            max_price = stocks[i] - min_price
    return max_price


print(calc_best_stock(stocks))
