unzipCustom :: [Integer] -> [Integer] -> [[Maybe Integer]]
unzipCustom first second = 
    if null first 
        then [] 
        else 
            [
                Just (head first), 
                if null (second) 
                    then Nothing 
                    else Just (head second)
            ] 
            : unzipCustom (tail first) (tail second)

main =
    print (unzipCustom [] []) >>
    print (unzipCustom [] [3]) >>
    print (unzipCustom [1] [3]) >>
    print (unzipCustom [1] [4, 5]) >> 
    print (unzipCustom [1, 2] [3]) >>
    print (unzipCustom [1, 2] [3, 4]) >> 
    print (unzipCustom [1, 2, 3] [4, 5, 6])