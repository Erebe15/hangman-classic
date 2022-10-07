package hangman

import (
	"math/rand"
	"time"
)

func RevealStartLettres(WordChosen string) []string { // select random lettres from the word for the start
	rand.Seed(time.Now().UnixNano())
	var ChosenLettres []string

	if len(WordChosen) >= 3 {
		for len(ChosenLettres) < (len(WordChosen)/2)-1 {
			RndLettre := string(WordChosen[rand.Intn(len(WordChosen))])
			if !DoesContain(ChosenLettres, RndLettre) {
				ChosenLettres = append(ChosenLettres, RndLettre)
			}
		}
	}
	return ChosenLettres
}

//func RevealWord(WordChosen string) string {
//	var RandomLetter []string // revealed letters list
//	rand.Seed(time.Now().UnixNano())
//	random := WordChosen[rand.Intn(len(WordChosen))] // une lettre du mot → random
//	if len(WordChosen) < 4 {                         // si le mot fais 3 lettres
//		RandomLetter = append(RandomLetter, string(random)) // ajoute random a la liste des revealed
//	} else {
//		for i := 0; i < (len(WordChosen)/2)-1; i++ { // sinon re-pioche une lettre du mot
//		RETURN:
//			random := WordChosen[rand.Intn(len(WordChosen)-1)]
//			if DoesContain(RandomLetter, string(random)) == true { // si on l'a deja on en pioche une autre
//				goto RETURN
//			} else {
//				RandomLetter = append(RandomLetter, string(random)) // ajoute la lettre a la liste
//			}
//		}
//	}
//	for i, l := range WordChosen {
//		println(i, "=", string(l))
//	}
//	fmt.Printf("liste de lettre revealed: %s\n", RandomLetter)
//	for i := 0; i < len(WordChosen); i++ {
//		if DoesContain(RandomLetter, string(WordChosen[i])) == true {
//			fmt.Print(string(WordChosen[i]))
//		} else {
//			fmt.Print("_")
//		}
//	}
//	return ""
//}
