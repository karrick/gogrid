package gogrid

import (
	"fmt"
	"strings"
)

func align(alignment Alignment, width int, text string) string {
	needed := width - len(text)
	if needed < 0 {
		return text[:width] // trim text to fit
	}
	if needed == 0 {
		return text // no padding required
	}

	switch alignment {
	case Left:
		return text + strings.Repeat(" ", needed)
	case Center:
		eachSize := needed >> 1
		double := eachSize << 1
		spaces := strings.Repeat(" ", eachSize)
		if double == needed {
			return spaces + text + spaces
		}
		// Need extra whitespace on one of the sides:
		return spaces + text + spaces + " "
	case Right:
		return strings.Repeat(" ", needed) + text
	default:
		panic(fmt.Errorf("cannot recognize alignment: %v", alignment))
	}
}

func withColors(colors, text string) string {
	pre, post := ansiFromLabel(colors)
	return pre + text + post
}

func ansiFromLabel(colors string) (string, string) {
	chunks := strings.Split(colors, ";")
	pres := make([]string, 0, len(chunks))
	posts := make([]string, 0, len(chunks))

	for _, chunk := range chunks {
		switch strings.ToLower(chunk) {
		case "bold":
			pres = append(pres, "1")
			posts = append(posts, "22") // NOTE: 22 is neither bold nor dim
		case "dim":
			pres = append(pres, "2")
			posts = append(posts, "22") // NOTE: 22 is neither bold nor dim
		case "italic":
			pres = append(pres, "3")
			posts = append(posts, "23")
		case "underline", "underscore":
			pres = append(pres, "4")
			posts = append(posts, "24")
		case "blink", "slow-blink":
			pres = append(pres, "5")
			posts = append(posts, "25")
		case "fast-blink", "rapid-blink":
			pres = append(pres, "6")
			posts = append(posts, "25")
		case "invert", "reverse":
			pres = append(pres, "7")
			posts = append(posts, "27")
		case "conceal", "hidden":
			pres = append(pres, "8")
			posts = append(posts, "28")
		case "strike", "strikethrough":
			pres = append(pres, "9")
			posts = append(posts, "29")

		case "black", "black-foreground":
			pres = append(pres, "30")
			posts = append(posts, "39")
		case "red", "red-foreground":
			pres = append(pres, "31")
			posts = append(posts, "39")
		case "green", "green-foreground":
			pres = append(pres, "32")
			posts = append(posts, "39")
		case "yellow", "yellow-foreground":
			pres = append(pres, "33")
			posts = append(posts, "39")
		case "blue", "blue-foreground":
			pres = append(pres, "34")
			posts = append(posts, "39")
		case "magenta", "purple", "magenta-foreground", "purple-foreground":
			pres = append(pres, "35")
			posts = append(posts, "39")
		case "cyan", "teal", "cyan-foreground", "teal-foreground":
			pres = append(pres, "36")
			posts = append(posts, "39")
		case "white", "white-foreground":
			pres = append(pres, "37")
			posts = append(posts, "39")

		case "black-background":
			pres = append(pres, "40")
			posts = append(posts, "49")
		case "red-background":
			pres = append(pres, "41")
			posts = append(posts, "49")
		case "green-background":
			pres = append(pres, "42")
			posts = append(posts, "49")
		case "yellow-background":
			pres = append(pres, "43")
			posts = append(posts, "49")
		case "blue-background":
			pres = append(pres, "44")
			posts = append(posts, "49")
		case "magenta-background", "purple-background":
			pres = append(pres, "45")
			posts = append(posts, "49")
		case "cyan-background":
			pres = append(pres, "46")
			posts = append(posts, "49")
		case "white-background":
			pres = append(pres, "47")
			posts = append(posts, "49")

		default:
			// pres = append(pres, chunk)
			// posts = append(posts, "0")
		}
	}

	if len(pres) == 0 {
		return "", ""
	}

	pre := "\033[" + strings.Join(pres, ";") + "m"
	post := "\033[" + strings.Join(posts, ";") + "m"
	return pre, post
}
