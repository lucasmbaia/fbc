import hashlib
header_hex = ("01000000" +
     "81cd02ab7e569e8bcd9317e2fe99f2de44d49ab2b8851ba4a308000000000000" +
      "e320b6c2fffc8d750423db8b1eb942ae710e951ed797f7affc8892b0f1fc122b" +
       "c7f5d74d" +
        "f2b9441a" +
	 "42a14695")
header_bin = header_hex.decode('hex')
print(header_hex)
print(hashlib.sha256(hashlib.sha256(header_hex).digest()).digest())
hash = hashlib.sha256(hashlib.sha256(header_bin).digest()).digest()
hash.encode('hex_codec')
print(hash.encode('hex_codec'))
#'1dbd981fe6985776b644b173a4d0385ddc1aa2a829688d1e0000000000000000'
print(hash[::-1].encode('hex_codec'))

#'00000000000000001e8d6829a8a21adc5d38d0a473b144b6765798e61f98bd1d'
