ackermann :: Integer -> Integer -> Maybe Integer 
ackermann m n = if m < 0 || n < 0
    then Nothing
    else case m of
        0 -> Just (n+1)
        k -> case n of 
            0 -> ackermann (k-1) 1
            l -> let Just ack = ackermann k (l-1)
                 in ackermann (k-1) ack
main =
    print (ackermann 0 0) >>
    print (ackermann 1 0) >>
    print (ackermann 1 1) >>
    print (ackermann 1 2) >>
    print (ackermann 3 5) 