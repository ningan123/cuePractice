
test1: "\(test.tt1)"
test2: test.tt1
test3: "test.tt1"

encode1: base64.Encode(null, "\(test.tt1)")
decode1: base64.Decode(null, base64.Encode(null, "\(test.tt1)"))

encode2: base64.Encode(null, test.tt1)
decode2: base64.Decode(null, base64.Encode(null, test.tt1))