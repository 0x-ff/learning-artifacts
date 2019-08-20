fibonacci :: Integer -> Maybe Integer 
fibonacci n = if n < 0
    then Nothing
    else case n of
        0 -> Just 0
        1 -> Just 1
        n -> let Just f1 = fibonacci (n-1)
                 Just f2 = fibonacci (n-2)
             in Just (f1 + f2)

main =
    print (fibonacci (3))