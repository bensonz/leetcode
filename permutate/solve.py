def permutate(inputList: list) -> list[list]:
    result = []
    for x in inputList:
        if len(result) == 0:
            result.append([x])
        else:
            newResult = []
            for y in result:
                for i in range(len(y) + 1):
                    newResult.append(y[:i] + [x] + y[i:])
            result = newResult
    return result


def heaps_algorithm(n, array):
    if n == 1:
        yield array.copy()
    else:
        for i in range(n):
            yield from heaps_algorithm(n - 1, array)
            if n % 2 == 0:
                array[i], array[n-1] = array[n-1], array[i]
            else:
                array[0], array[n-1] = array[n-1], array[0]

# Usage:


if __name__ == "__main__":
    # print(permutate([1, 2, 3]))
    # print(permutate([0, 1]))
    # print(permutate([1]))
    # print(permutate([]))
    list_of_permutations = list(heaps_algorithm(3, [1, 2, 3]))
    print(list_of_permutations)
