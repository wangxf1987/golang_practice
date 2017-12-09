n = 12
k = 1

num_seq = ''.join([str(i) for i in range(n+1)])
print num_seq
k_seq = ''.join([j for j in num_seq if j == str(k)])
print k_seq
k_len = len(k_seq)
print k_len
