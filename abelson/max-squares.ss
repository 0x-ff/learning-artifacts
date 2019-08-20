(define (square x) (* x x))

(define (square-sum x y) (+ (square x) (square y)))

(define (max-square-sum x y z)
    (if (> x y)
        (if (> x z) 
            (if (> y z) (square-sum x y) (square-sum x z))
            (square-sum x z))
        (if (> x z) (square-sum x y) (square-sum y z))
    )
)

(max-square-sum 1 2 3)
(max-square-sum 1 3 2)
(max-square-sum 2 1 3)
(max-square-sum 3 1 2)
(max-square-sum 2 3 1)
(max-square-sum 3 2 1)
