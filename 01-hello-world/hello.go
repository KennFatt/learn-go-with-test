package main

const bahasaIndonesia = "Bahasa Indonesia"
const french = "French"
const englishHelloPrefix = "Hello,"
const bahasaIndonesiaPrefix = "Halo,"
const frenchPrefix = "Bonjour,"

func main() {
}

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "there"
	}

	return getLocaleGreeting(language) + " " + name + "!"
}

func getLocaleGreeting(locale string) (prefix string) {
	switch locale {
	case bahasaIndonesia:
		prefix = bahasaIndonesiaPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
