# comment
# variable declaration
(var a true) # inline comment
(var b 2.5)
(var c "str")

# variable reference
(var x a)

# comparison
(= 3 5) # false
(>= 2 0) # true

# conditional
(if (< 2 3) (5) (6)) # returns 5

# basic arithmetic
(+ 5 4) # 9
(- 5 4) # 1

# function declaration
(fn fib n (if (<= n 1)
    (1)
    (+ (call fib (- n 1)) (call fib (- n 2)))))

# function call
(call print "printing fib(10):")
(call print (call fib 10))

# call a built-in function
(call str_concat "a" "b") # "ab"

# array-like structures can be implemented in the language directly
(fn car cons_cell (call cons_cell (fn _ a b (a)))) # get head of list
(fn cdr cons_cell (call cons_cell (fn _ a b (b)))) # get tail of list
(fn cons a b (fn _cons f (call f a b))) # used to define a list

(var lst (call cons 10 20))
(call print "printing car of lst:")
(call print (call car lst)) # 10
(call print "printing cdr of lst:")
(call print (call cdr lst)) # 20

(var longer_lst (call cons 30 (call cons 60 90)))
(call print "printing car of longer_lst:")
(call print (call car longer_lst)) # 30
(call print "printing car of cdr of longer_lst:")
(call print (call car (call cdr longer_lst))) # 60
(call print "printing cdr of cdr of longer_lst:")
(call print (call cdr (call cdr longer_lst))) # 90

# call a sequence of expressions
(seq [(call print "first expression in sequence")
    (call print "second expression in sequence")
    (call print "third expression in sequence")])

# using recursion
(fn pow base exp (if (= exp 0)
    (1)
    (* base (call pow base (- exp 1)))))
(call print "printing 2^3:")
(call print (call pow 2 3)) # 8

# uncomment to throw error
# (call throw_error "error!")

# dictionaries can be implemented using functions as well
(fn make_dict dict key val
    (fn _ k (if (= key k)
        (val)
        (if (= null dict)
        (null)
        (call dict k)))))

(var my_dict (call make_dict null "a" 200))
(var my_dict (call make_dict my_dict "my_key" true))

(call print "printing dictionary keys:")
(call print (call my_dict "a")) # 200
(call print (call my_dict "my_key")) # true
(call print (call my_dict "not_in_dict")) # null

# file manipulation using built-in functions
(var text (call read_file "./text.txt"))
(call write_file "./text2.txt" (call str_concat text text))

# expect function for testing
(fn expect val expected (if (= val expected)
    (true)
    (call throw_error (call str_concat "Expected "
        (call str_concat (call to_str val)
            (call str_concat " to equal " (call to_str expected)))))))
(call expect "a" "a")
# (call expect "a" "b") # would throw an error

# "and" functionality
(fn and cond1 cond2 (if (cond1)
    (cond2)
    (false)))

(call expect (call and true true) (true))
(call expect (call and true false) (false))
(call expect (call and false true) (false))
(call expect (call and false false) (false))

# "or" functionality
(fn or cond1 cond2 (if (cond1)
    (true)
    (cond2)))

(call expect (call or true true) (true))
(call expect (call or true false) (true))
(call expect (call or false true) (true))
(call expect (call or false false) (false))

# "not" functionality
(fn not cond (if (cond) (false) (true)))

(call expect (call not true) (false))
(call expect (call not false) (true))

# using null value and some recursion
# [1,2,3,4,5]
(var my_lst (call cons 1 (call cons 2 (call cons 3 (call cons 4 (call cons 5 null))))))
(fn map lst f (if (= (call cdr lst) null)
    (call cons (call f (call car lst)) null)
    (call cons (call f (call car lst)) (call map (call cdr lst) f))))

(fn double a (* a 2))
(var doubled_lst (call map my_lst double))
(call print "printing all items in a list:")
(call map doubled_lst print) # [2,4,6,8,10]

# filter function
(fn filter lst f (if (= (call cdr lst) null)
    (if (call f (call car lst))
        (call cons (call car lst) null)
        (null))
    (if (call f (call car lst))
        (call cons (call car lst) (call filter (call cdr lst) f))
        (call filter (call cdr lst) f))))

(var mixed_nums (call cons 14 (call cons 7 (call cons 8 (call cons 22 (call cons 13 null))))))
(fn is_even x (= 0 (% x 2)))
(fn is_odd x (call not (= 0 (% x 2))))
(var only_even (call filter mixed_nums is_even)) # [14,8,22]
(var only_odd (call filter mixed_nums is_odd)) # [7,13]

(fn sum_lst lst (if (= (call cdr lst) null)
    (call car lst)
    (+ (call car lst) (call sum_lst (call cdr lst)))))

(var another_lst (call cons 1 (call cons 2 (call cons 3 (call cons 4 (call cons 5 null))))))
(call print "printing sum of list:")
(call print (call sum_lst another_lst))

(fn size_lst lst (if (= (call car lst) null)
    (0)
    (if (= (call cdr lst) null)
        (1)
        (+ 1 (call size_lst (call cdr lst))))))

(call size_lst another_lst) # 5
(call size_lst (call cons null null)) # 0

(fn avg_lst lst (/ (call sum_lst lst) (call size_lst lst)))
(call print "printing average of list:")
(call print (call avg_lst another_lst))
