import codecs

bits = 486604799
p = ''

# calculando a string do alvo para checarmos
exp = bits >> 24
print(exp)
mant = bits & 0xffffff
print(mant)
target = mant * (1 << (8 * (exp - 3)))
print(target)
target_hexstr = '%064x' % target
print(target_hexstr)
target_str = codecs.decode(target_hexstr, 'hex')
nonce = 100000000
