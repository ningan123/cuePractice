import "encoding/base64"


t1: base64.Encode(null, "foo")
t2: base64.Decode(null, base64.Encode(null, "foo"))
