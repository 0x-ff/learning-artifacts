sumCustom :: [Integer] -> Integer
sumCustom list = foldr (\x y -> y + x) 0 list

filterNotDividedBy :: Integer -> [Integer] -> [Integer]
filterNotDividedBy n list = 
    filter (\x -> (rem x n) > 0) list

filterEratosfen :: [Integer] -> [Integer]
filterEratosfen list = 
    if null (tail list)
    then [head list]
    else head list : filterEratosfen (filterNotDividedBy (head list) (tail list))

primes :: Integer -> [Integer]
primes m = 
    let list = [2..m] 
    in filterEratosfen list

main =
    -- сумма простых чисел меньших заданного
    print (sumCustom (primes 10000))