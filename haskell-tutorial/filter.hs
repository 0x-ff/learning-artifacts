filterNumber :: Integer -> [Integer] -> [Integer]
filterNumber num list = filter (\x -> num == x) list

filterNotNumber :: Integer -> [Integer] -> [Integer]
filterNotNumber num list = filter (\x -> num /= x) list

main =
    print (filterNumber 1 [1,2,1,4]) >> 
    print (filterNotNumber 1 [1,2,1,4])