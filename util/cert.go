package util

import (
	"io/ioutil"
	"path/filepath"
)

var (
	CertPath = filepath.Join(LianDiDir, "cert.pem")
	KeyPath  = filepath.Join(LianDiDir, "key.pem")
)

func InitCert() {
	ioutil.WriteFile(CertPath, cert, 0644)
	ioutil.WriteFile(KeyPath, key, 0644)
}

// go run d:\\go\\src\\crypto\\tls\\generate_cert.go --host localhost

var (
	cert = []byte(`-----BEGIN CERTIFICATE-----
MIIC+TCCAeGgAwIBAgIQMGveXlCWypp1V+/TVlKahDANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMB4XDTIwMDEyMzA2MTcxMloXDTIxMDEyMjA2MTcx
MlowEjEQMA4GA1UEChMHQWNtZSBDbzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAKO3hdfT/tRmgFuMi3RNhJM3dZ14PwR5/gbVlwVbSJ6b7IgNVG2lrqWV
pfYZob4jEK0RKrrRYIEYnl6sWTkorv3uqbYFwOUa89xnr3BtD5HZnzDa4PdmGrCP
PJnw/v2nI33mypJ27M7Jy3peAHIS1WEtgOCKpMFOtkbCfNePHyYniqyhzFWQ3ztx
rexFhjllcImDaSFd2k5UomzPPgIgg6nttAMx9ZOwdrzHt474q6E4BHtuPOcyraPx
sUiv3jB3yfwL+IT2qJPxVA7Z43STyv8lQPJc4XduB7uqgphwlmZCNWtLjpWDPETG
unzCkvkpnXSdplHHeBn/njLWjeA2vkcCAwEAAaNLMEkwDgYDVR0PAQH/BAQDAgWg
MBMGA1UdJQQMMAoGCCsGAQUFBwMBMAwGA1UdEwEB/wQCMAAwFAYDVR0RBA0wC4IJ
bG9jYWxob3N0MA0GCSqGSIb3DQEBCwUAA4IBAQBactVxCfV6EsQWBdcL+zTL1xWv
B1NV+ubH18Wu57BV8Tm4yyeN7aw8t51HtGyhNnMGKwLnVT2XsxQgpzUVPSGcIlci
CGgFprHGinkS5nmVQp1qvH5YndDTOXGsepjtRfITYG3vdQW8/iRV9cp4hWK6UPdS
8GXQMrBSlZBa/i8WOtcqlbG38xZfnkTx62l8eKW4ZUoH0sO/lcENDoTqBMOX3Ihw
1mjUxqjl4sLzlxy140ker6/3YbGS0fgff7jUMEjBnyGsHF8axJGl6YnAr/TPTaqk
TN2GC5026DJOAkD/dD/DbHBdn42T++XZ5spU0tQh392A1WeER79RRMnRTx21
-----END CERTIFICATE-----
`)

	key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAo7eF19P+1GaAW4yLdE2Ekzd1nXg/BHn+BtWXBVtInpvsiA1U
baWupZWl9hmhviMQrREqutFggRieXqxZOSiu/e6ptgXA5Rrz3GevcG0PkdmfMNrg
92YasI88mfD+/acjfebKknbszsnLel4AchLVYS2A4IqkwU62RsJ8148fJieKrKHM
VZDfO3Gt7EWGOWVwiYNpIV3aTlSibM8+AiCDqe20AzH1k7B2vMe3jviroTgEe248
5zKto/GxSK/eMHfJ/Av4hPaok/FUDtnjdJPK/yVA8lzhd24Hu6qCmHCWZkI1a0uO
lYM8RMa6fMKS+SmddJ2mUcd4Gf+eMtaN4Da+RwIDAQABAoIBAQCfWkzOrSxHNZYO
hhhfFezLFppGsfo0s0FNCa07q4RdzctHCiJ5r2D1D1pA1uYWGi9rJWCAArh6L6pG
4F7gQfdCd4F9CNCer8NZbxrWxXD4N+i5NvLuG9YysNj9xhOMdZkYPQJD2bnZ9ZU3
ADRS1H4ZXp8o8Ao5lhHZggjeEBLzoM63e+UryCKHxaOss3c0Uu6RBZns5LcnbCyr
9CUeN8lup2k3MstPVwfoE0prgu6bkVPDnSMi6nE+E24z32fMz5k8b9D5mzZwM4+I
eKcPSeUha1aazRri2guBbX1cTBmCVUQ1zOQ8Is8duaxgSUYC29g1Eoadfg7ZQTkg
fMJLq4NxAoGBAMgTg3koKrmYwTH8G0rfph3ZgZfFak+uH6uG42AlZEnaphQqMhGX
IwHlYwhlrejjkIWnG/zEBYRVnobVPyaTkw+F8McEh+CeTc1oeDf9lQzLtV3V0H/g
fbrTv2eju8TpXExcYxUCpF74+ssl6m0Ldo3v0q1sukxOxUPjj4kIph+LAoGBANF6
T5iQ+f2qK1RVt7ZjGwMlj2YJfHZsR3kJYODYufrXX/4twfUS9fI/mjnWTWSS6AF6
cmB/15eJnujEJY1KBBQ8XEwejQISkt9VMcMw34ju0WpCJyVyzOUSQub4XEfTkkYW
U/rA467o5JdnIZR1atfomO/0DMre3jukZ7sqpXO1AoGBAJJz1jUFdeDNM/3GsKOp
+gq3f+j8lMkmiAAiiKWEOVjmN9Ni36uImhN1OXyYESj8CnoKzK3FPtSTZ6sCxWsT
cv5V6N2FL4D337OP2RADCuuG4YeCT1CLnvz2qpDOhU+qEenDZrOmb3MlhU7WvjZ7
wZg8CFbx05yvvC3pkNOaBR/pAoGBAJY3xN6ekHdE0b+vnIexeEyAYd4FTXjK21JD
1HJJWM02J8Np4t4xU6f9zkZNlMBbzP5KZb0n0F3+NbOr/VXdzyEHzBRWED59PU/k
k0PrR1G2GPy/Jb1oHuDigJGLZvmnEg0qs9xs+JmUr2CXd3Az54OfbBUfCRFuUyWo
Fy8UPQPFAoGAJouvqHhxS+5zRzasMEkmk/I27eUTw3/haa4efVzck1PN2Z1KCda8
4Ruyf8RusAsaKAWm46149ZXsVUpUTP6dQwwKFEhutF4rrwhIHQLYUAu0zKnEtvIE
nsTX0/0JA9FiPYGG+jL6h69eLImhYpHilxpTZ7DnXEd/ytZjkywDNxM=
-----END RSA PRIVATE KEY-----
`)
)
