package main

import (
	asciiart ".."
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	normal            = "\033[0m"
	grave             = "\033[1m"
	soulignerGrave    = "\033[21m"
	fondBlanc         = "\033[7m"
	colorBlack        = "\033[30m"
	colorRedSombre    = "\033[31m"
	colorGreenSombre  = "\033[32m"
	colorYellowSombre = "\033[33m"
	colorBlueSombre   = "\033[34m"
	colorPurpleSombre = "\033[35m"
	colorCyanSombre   = "\033[36m"
	colorRed          = "\033[91m"
	colorGreen        = "\033[92m"
	colorYellow       = "\033[93m"
	colorBlue         = "\033[94m"
	colorRose         = "\033[95m"
	colorCyan         = "\033[96m"
	colorWhite        = "\033[97m"
	fondNoir          = "\033[40m"
	fondRouge         = "\033[41m"
	fondVert          = "\033[42m"
	fondJaune         = "\033[43m"
	fondBleu          = "\033[44m"
	fondViolet        = "\033[45m"
	fondCyan          = "\033[46m"
	fondGris          = "\033[100m"
	fondRoseRouge     = "\033[101m"
	fondRose          = "\033[105m"
	fondBleuCiel      = "\033[106m"
	couleur           = ""
	typeOfWriting     = ""
	text              = ""
	result            = ""
	fond              = ""
)

func main() {
	fmt.Print(grave)
	fmt.Println("\n\n        				" + colorBlueSombre + "Welcome In Ascii-Art : " + normal + "                                  ")
	fmt.Println(colorGreen + grave + "\n-" + "Choose your language ")
	fmt.Println(colorRed + "\nFrench    ->  1\n" + colorPurpleSombre + "English   ->  2\n" + colorCyanSombre + "Spanish   ->  3" + normal)
	fmt.Print(colorYellowSombre + "\n" + "Your language : ")
	fmt.Print(normal)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	language, _ := strconv.Atoi(scanner.Text())
	switch language {
	case 1:
		fmt.Println(colorBlue + "\nVous avez" + normal + " choisis" + colorRed + " le français")
		fmt.Print(colorGreen+"\n-"+"Choisissez votre texte à modifier en ascii-art : ", normal)
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text = scanner.Text()
		fmt.Println(normal + "\nVotre text est : " + text)
		fmt.Println(colorGreen + "\n-" + "Choisissez la couleur de votre text ")
		fmt.Println(colorRed + "\nSi vous voulez plus d'information sur les couleurs disponibles écrivez : " + colorWhite + "Aide" + colorGreen + normal)
		fmt.Print(colorYellowSombre+"\n"+"Quelle couleur choisissez-vous : ", normal)
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		couleur = scanner.Text()
		if couleur == "Aide" || couleur == "aide" {
			fmt.Println(colorCyan + "\nLe choix des couleurs se partagent en plusieurs catégories : ")
			fmt.Println("- Les Couleurs            ->  1")
			fmt.Println("- Les Fonds de Couleurs spécifique (ces fonds suppriment les couleurs choisis)   ->  2")
			fmt.Print("Sur quelle catègorie souhaitez vous aller : " + normal)
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			categorie, _ := strconv.Atoi(scanner.Text())
			switch categorie {
			case 1:
				fmt.Println(colorYellowSombre + "\nVoici les couleur disponible : ")
				fmt.Println("Normal, Grave, Blanc, Rouge, Vert, Jaune, Bleu, Rose, Cyan, Noir\nRouge Sombre, Jaune Sombre, Vert Sombre, Bleu Sombre, Violet Sombre, Cyan Sombre ")
			case 2:
				fmt.Println(colorYellowSombre + "\nVoici les fonds disponible : ")
				fmt.Println("Aucun, Fond Gris, Fond Blanc, Fond Rouge, Fond Vert, Fond Jaune, Fond Bleu, Fond Violet, Fond Cyan, Fond Noir, Fond Rose-Rouge, Fond Bleu-Ciel, Fond Rose")
			default:
				fmt.Println(soulignerGrave + colorRedSombre + "\nChoisissez 1 ou 2" + normal)
				return

			}
			fmt.Print(grave + colorGreen + "\nQuelle couleur choisissez-vous : " + normal)
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			couleur = scanner.Text()

		}

		switch {

		case couleur == "noir" || couleur == "Noir" || couleur == "NOIR":
			couleur = colorBlack
		case couleur == "grave" || couleur == "Grave" || couleur == "GRAVE":
			couleur = grave
		case couleur == "blanc" || couleur == "Blanc" || couleur == "BLANC":
			couleur = colorWhite
		case couleur == "rouge" || couleur == "Rouge" || couleur == "ROUGE":
			couleur = colorRed
		case couleur == "vert" || couleur == "Vert" || couleur == "VERT":
			couleur = colorGreen
		case couleur == "jaune" || couleur == "Jaune" || couleur == "JAUNE":
			couleur = colorYellow
		case couleur == "bleu" || couleur == "Bleu" || couleur == "BLEU":
			couleur = colorBlue
		case couleur == "rose" || couleur == "Rose" || couleur == "ROSE":
			couleur = colorRose
		case couleur == "cyan" || couleur == "Cyan" || couleur == "CYAN":
			couleur = colorCyan
		case couleur == "normal" || couleur == "Normal" || couleur == "NORMAL":
			couleur = normal
		case couleur == "Rouge Sombre" || couleur == "Rouge sombre" || couleur == "rouge sombre" || couleur == "rouge Sombre" || couleur == "Rougesombre" || couleur == "RougeSombre" || couleur == "rougesombre" || couleur == "rougeSombre":
			couleur = colorRedSombre
		case couleur == "Jaune Sombre" || couleur == "Jaune sombre" || couleur == "jaune sombre" || couleur == "jaune Sombre" || couleur == "Jaunesombre" || couleur == "JauneSombre" || couleur == "jaunesombre" || couleur == "jauneSombre":
			couleur = colorYellowSombre
		case couleur == "Vert Sombre" || couleur == "Vert sombre" || couleur == "vert sombre" || couleur == "vert Sombre" || couleur == "Vertsombre" || couleur == "VertSombre" || couleur == "vertsombre" || couleur == "vertSombre":
			couleur = colorGreenSombre
		case couleur == "Bleu Sombre" || couleur == "Bleu sombre" || couleur == "bleu sombre" || couleur == "bleu Sombre" || couleur == "Bleusombre" || couleur == "BleuSombre" || couleur == "bleusombre" || couleur == "bleuSombre":
			couleur = colorBlueSombre
		case couleur == "Violet Sombre" || couleur == "Violet sombre" || couleur == "violet sombre" || couleur == "violet Sombre" || couleur == "Violetsombre" || couleur == "VioletSombre" || couleur == "violetsombre" || couleur == "violetSombre":
			couleur = colorPurpleSombre
		case couleur == "Cyan Sombre" || couleur == "Cyan sombre" || couleur == "cyan sombre" || couleur == "cyan Sombre" || couleur == "Cyansombre" || couleur == "CyanSombre" || couleur == "cyansombre" || couleur == "cyanSombre":
			couleur = colorCyanSombre
		case couleur == "Fond Noir" || couleur == "Fond noir" || couleur == "fond noir" || couleur == "fond Noir" || couleur == "Fondnoir" || couleur == "FondNoir" || couleur == "fondnoir" || couleur == "fondNoir":
			couleur = fondNoir
		default:
			fmt.Println(soulignerGrave + colorRedSombre + "\nVeuillez écrire correctement la couleur" + normal)
			return

		}
		fmt.Println(colorGreen + "\n-" + "Choisissez le fond de votre text ")
		fmt.Println(colorRed + "\nSi vous voulez plus d'information sur les couleurs disponibles écrivez : " + colorWhite + "Aide" + colorGreen + normal)
		fmt.Println(colorCyan + "\nAttention ! Certains fonds ne sont pas compatible avec des couleurs : \nSi vous choisissez un fond et une couleurs , il se peut que le résultat ne soit pas le résultat voulu"+ normal)
		fmt.Print(colorYellowSombre+"\n"+"Quelle fond choisissez-vous : ", normal)

		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fond = scanner.Text()
		if fond == "Aide" || fond == "aide" {
			fmt.Println(colorYellowSombre + "\nVoici les fonds disponible : ")
			fmt.Println("Aucun, Fond Gris, Fond Blanc, Fond Rouge, Fond Vert, Fond Jaune, Fond Bleu, Fond Violet, Fond Cyan, Fond Noir, Fond Rose-Rouge, Fond Bleu-Ciel, Fond Rose")
			fmt.Print(grave + colorGreen + "\nQuelle fond choisissez-vous : " + normal)
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			fond = scanner.Text()
		}
		switch {
		case fond == "aucun" || fond == "Aucun" || fond == "AUCUN" || fond == "":
			fond = ""
		case fond == "Fond Blanc" || fond == "Fond blanc" || fond == "fond blanc" || fond == "fond Blanc" || fond == "Fondblanc" || fond == "FondBlanc" || fond == "fondblanc" || fond == "fondBlanc":
			fond = fondBlanc
		case fond == "Fond Rouge" || fond == "Fond rouge" || fond == "fond rouge" || fond == "fond Rouge" || fond == "Fondrouge" || fond == "FondRouge" || fond == "fondrouge" || fond == "fondRouge":
			fond = fondRouge
		case fond == "Fond Vert" || fond == "Fond vert" || fond == "fond vert" || fond == "fond Vert" || fond == "Fondvert" || fond == "FondVert" || fond == "fondvert" || fond == "fondVert":
			fond = fondVert
		case fond == "Fond Cyan" || fond == "Fond cyan" || fond == "fond cyan" || fond == "fond Cyan" || fond == "Fondcyan" || fond == "FondCyan" || fond == "fondcyan" || fond == "fondCyan":
			fond = fondCyan
		case fond == "Fond Jaune" || fond == "Fond jaune" || fond == "fond jaune" || fond == "fond Jaune" || fond == "Fondjaune" || fond == "FondJaune" || fond == "fondjaune" || fond == "fondJaune":
			fond = fondJaune
		case fond == "Fond Bleu" || fond == "Fond bleu" || fond == "fond bleu" || fond == "fond Bleu" || fond == "Fondbleu" || fond == "FondBleu" || fond == "fondbleu" || fond == "fondBleu":
			fond = fondBleu
		case fond == "Fond Violet" || fond == "Fond violet" || fond == "fond violet" || fond == "fond Violet" || fond == "Fondviolet" || fond == "FondViolet" || fond == "fondviolet" || fond == "fondViolet":
			fond = fondViolet
		case fond == "Fond Gris" || fond == "Fond gris" || fond == "fond gris" || fond == "fond Gris" || fond == "Fondgris" || fond == "FondGris" || fond == "fondgris" || fond == "fondGris":
			fond = fondGris
		case fond == "Fond Rose-Rouge" || fond == "Fond rose-rouge" || fond == "fond Rose-rouge" || fond == "fond rose-Rouge" || fond == "Fondrose-rouge" || fond == "FondRose-Rouge" || fond == "fondrose-rouge" || fond == "fondRose-Rouge":
			fond = fondRoseRouge
		case fond == "Fond Rose" || fond == "Fond rose" || fond == "fond Rose" || fond == "fond rose" || fond == "Fondrose" || fond == "FondRose" || fond == "fondrose" || fond == "fondRose":
			fond = fondRose
		case fond == "Fond Bleu-Ciel" || fond == "Fond bleu-ciel" || fond == "fond Bleu-ciel" || fond == "fond bleu-Ciel" || fond == "Fondbleu-ciel" || fond == "FondBleu-Ciel" || fond == "fondbleu-ciel" || fond == "fondBleu-Ciel":
			fond = fondBleuCiel
		case fond == "Fond Noir" || fond == "Fond noir" || fond == "fond noir" || fond == "fond Noir" || fond == "Fondnoir" || fond == "FondNoir" || fond == "fondnoir" || fond == "fondNoir":
			fond = fondNoir
		default:
			fmt.Println(soulignerGrave + colorRedSombre + "\nVeuillez écrire correctement le fond" + normal)
			return
		}

		fmt.Println(grave + normal + colorGreen + "\n-" + "Choisissez le style d'écriture de votre text : " + normal)
		fmt.Println(colorRed + "\nSi vous voulez plus d'information sur les styles d'écritures disponibles écrivez : " + colorWhite + "Aide" + normal)
		fmt.Print(normal + colorYellowSombre + "\n-" + "Quel type d'écriture choisissez-vous : " + normal)
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		writingType := scanner.Text()
		if writingType == "Aide" || writingType == "aide" {
			fmt.Println(colorCyan + "Il existe 3 types d'écritures: ")
			fmt.Println(colorYellowSombre+"Standard\nShadow\nThinkertoy ")
			fmt.Print("Quel type d'écriture choisissez-vous : " + normal)
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			writingType = scanner.Text()
		}
		switch {
		case writingType == "standard" || writingType == "standard.txt" || writingType == "Standard":
			typeOfWriting = "standard.txt"
		case writingType == "shadow" || writingType == "shadow.txt" || writingType == "Shadow":
			typeOfWriting = "shadow.txt"
		case typeOfWriting == "thinkertoy" || writingType == "thinkertoy.txt" || writingType == "Thinkertoy":
			typeOfWriting = "thinkertoy.txt"
		default:
			fmt.Println(soulignerGrave + colorRedSombre + "Veuillez écrire correctement le style d'écriture" + normal)
			return

		}
	case 2:
		fmt.Println(colorRed+"Cette langue n'est pas encore répertorié , veuillez réessayer plus tard")
		return
	case 3:
		fmt.Println(colorRed+"Cette langue n'est pas encore répertorié , veuillez réessayer plus tard")
		return
	default:
		fmt.Println(soulignerGrave + colorRedSombre + "Veuillez choisir 1 - 2 - 3"+ normal)



	}
	fmt.Println(normal)

	prev := '0'
	newLine := false
	for _, v := range text {
		if v == 'n' && prev == '\\' {
			newLine = true
		}
		prev = v
	}

	if newLine {

		args := strings.Split(text, "\\n")
		for _, v := range args {

			for i := 0; i < 8; i++ {

				for _, w := range v {

					result += asciiart.ScanLine(typeOfWriting, 1+int(w-32)*9+i)

				}
				fmt.Println(fond+couleur, result)
				result = ""
			}

		}

	} else {

		for i := 0; i < 8; i++ {

			for _, v := range text {

				result += asciiart.ScanLine(typeOfWriting, 1+int(v-32)*9+i)

			}
			fmt.Println(fond+couleur, result)
			result = ""
		}

	}

	fmt.Println(normal)

}