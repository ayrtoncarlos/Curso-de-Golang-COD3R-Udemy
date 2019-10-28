package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado retornado foi %v."

func TestMedia(t *testing.T) {
	t.Parallel()
	mediaEsperada := 7.28
	media := Media(7.2, 9.9, 6.1, 5.9)

	if media != mediaEsperada {
		t.Errorf(erroPadrao, mediaEsperada, media)
	}
}
