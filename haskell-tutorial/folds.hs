sumCustom :: [Integer] -> Integer
sumCustom list = foldr (\x y -> y + x) 0 list

countCustom :: [Integer] -> Integer
countCustom list = foldr (\x y -> y + 1) 0 list

main =
    print (sumCustom ([1])) >>
    print (sumCustom ([1,2])) >>
    print (sumCustom ([1,2,3])) >>
    print (sumCustom ([1,2,3,4,5,6,7,8,9,10])) >>
    print (countCustom ([1])) >>
    print (countCustom ([1,2])) >>
    print (countCustom ([1,2,3])) >>
    print (countCustom ([1,2,3,4,5,6,7,8,9,10]))