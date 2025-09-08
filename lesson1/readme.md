### type 
# boolean
var x bool
mặc định false
# số 
// 	uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// float32     the set of all IEEE 754 32-bit floating-point numbers
// float64     the set of all IEEE 754 64-bit floating-point numbers

// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts

// byte        alias for uint8
// rune        alias for int32
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

uint là số nguyên dương
không có double, float chỉ có float32|float64 

# string
var String string;
dấu "" tương đương ``
thư viện strings 


docs :
 # các thư viện xem hàm viết thế nào https://pkg.go.dev/std

## PHÉP GÁN 

# 2 cú pháp 

* var i string   Explicit Declaration: khai báo tường minh
g := value       short Declaration: khai báo ngắn ngọn

giá trị mặc định sẽ được theo kiểu nếu theo khai báo thứ nhất
còn khai báo thứ 2 thì sẽ được tự ép kiểu


## if 
# cú pháp 
if initialization; condition {
    // khối lệnh nếu điều kiện đúng
}
if trong go cho phép sử dụng khai báo ngắn gọn và biến được khai báo chỉ tồn tại phạm vi trong if

## switch 
switch [init]; [expression] {
case value1, value2, ...:
    // khối lệnh khi biểu thức khớp với value1 hoặc value2
    [fallthrough] // nếu muốn tiếp tục thực thi case tiếp theo

case valueN:
    // khối lệnh khác

default:
    // khối lệnh mặc định nếu không khớp case nào
}

chỉ tồn tại khởi tạo hoặc biểu thức hoặc thiếu cả 2 và value thay bằng condition nếu thiếu cả 2 hoặc trường hợp khởi tạo
trong go break được mặc định cho tất cả dòng trừ trường hợp khai báo fallthrough nên khi sử dụng fallthrough nó chỉ thực thi thêm 1 case dưới nũa và trở lại chức năng ban đầu 
