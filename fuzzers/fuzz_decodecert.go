package myfuzz
import librpki"github.com/cloudflare/cfrpki/validator/lib"

func Fuzz(data []byte) int {
	_, err := librpki.DecodeCertificate(data)
	if err != nil {
		return 1
	}
	return 0
}
