
(define (average x y) (/ (+ x y) 2))

(define (cubic x y) (/ (+ (* 2 y) (/ x (* y y))) 3))

(define (improve fn guess x)
    (fn guess (/ x guess))
)

(define (good-enough? guess x)
    (< 
        (abs (- (square guess) x))
        0.001
    )
)

(define (sqrt-iter fn guess x)
    (if (good-enough? guess x)
        guess
        (sqrt-iter fn (improve fn guess x) x)
    )
)

(define (new-if predicate then-clause else-clause)
    (cond (predicate then-clause)
        (else else-clause))
)

(define (new-sqrt-iter guess x)
    (new-if (good-enough? guess x)
        guess
        (new-sqrt-iter (improve average guess x) x)
    )
)

(define (sqrt x)
    (sqrt-iter average 1.0 x)
)

(define (new-sqrt x)
    (new-sqrt-iter 1.0 x)
)

(average 1 4)
(improve average 1 5)
(good-enough? 0.000001 0.0001)
(sqrt-iter average 0.000001 0.0001)

(sqrt 9)
