
import Foundation

func filterNoDividedBy(_ n: Int, _ list: [Int]) -> [Int] {
    if n <= 0 {
        return list
    }
    var filtered:[Int] = []
    for number in list {
        if number % n != 0 {
            filtered.append(number)
        }
    }
    return filtered;
}

func filterEratosfen(_ list: [Int]) -> [Int] {
    if list.count < 2 {
        return list
    }
    var tail = list
    let head = tail.removeFirst()
    
    var filtered:[Int] = []
    filtered = filterEratosfen(filterNoDividedBy(head, tail))
    filtered.insert(head, at:0)
    return filtered
}

func primes(_ max: Int) -> [Int] {
    if (max <= 1) {
        return []
    }
    var intList:[Int] = []
    for i in 2...max {
        intList.append(i)
    }
    return filterEratosfen(intList)
}

print("Hello, World!")

print(primes(1000))
