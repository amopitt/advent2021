
# read a file
def read_file(filename):
    with open(filename, 'r') as f:
        return f.read()


input = read_file('input.txt')

# loop through each character and check if it is ( or )
# if it is (, add 1 to the counter
# if it is ), subtract 1 from the counter  
loop_counter = 0
index = 1
for elem in input:
    if elem == '(': 
        loop_counter += 1
    elif elem == ')':  
        loop_counter -= 1
    
    if loop_counter < 0:
        print(index)
        break
    index += 1

print(index)
print(loop_counter)
 

