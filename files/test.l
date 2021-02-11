# variable declaration
#(var x 1)
#(var a true) # inline comment
#(var b 2.5)
#(var c "str")

# variable reference
#(var x a)

# comparison
#(= 3 5) # false
#(>= 2 0) # true

# conditional
#(if (< 2 3) (5) (6)) # returns 5

# basic arithmetic
#(+ 5 4) # 9
#(- 5 4) # 1

# function declaration
#(fn fib n (if (<= n 1)
#    (1)
#    (+ (call fib (- n 1)) (call fib (- n 2)))))

(fn myfn n (* 2 n))

# function call
#(call print "printing fib(10):")
#(call print (call fib 10))
