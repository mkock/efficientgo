package efficientgo

import "strings"

var snippets = []string{
	"Hello",
	"If invoked on a pair of input files, benchstat adds to the output a column showing the statistics from the second file and a column showing the percent change in mean from the first to the second file. Next to the percent change, benchstat shows the p-value and sample sizes from a test of the two distributions of benchmark times. Small p-values indicate that the two distributions are significantly different.",
	"The order of the lines in the file does not matter, except that the output lists benchmarks in order of appearance.",
	"Martin",
	"Use built-in Amazon Alexa to get quick news, info and weather, set bedtime reminders and alarms, control your smart home devices and more-just by speaking to your smartwatch.",
	"Disc Brake System - Provides long lasting and stronger fasting stopping power than the traditional V-Brake Systems.",
	"Durability: Duralock technology keeps the full energy of your unused Duracell batteries for up to 10 years in their packaging (ambient storage)",
	"Simba idolizes his father, King Mufasa, and takes to heart his own royal destiny. But not everyone in the kingdom celebrates the new cub’s arrival. Scar, Mufasa’s brother—and former heir to the throne—has plans of his own. The battle for Pride Rock is ravaged with betrayal, tragedy and drama, ultimately resulting in Simba’s exile. With help from a curious pair of newfound friends, Simba will have to figure out how to grow up and take back what is rightfully his.",
	"S.Pellegrino embodies the finest side of the Italian way of living - it's all about beauty, passion, food and togetherness",
	"Naturally carbonated, S.Pellegrino has unique mid-size bubbles that reveal the finest tastes of food",
}

var firstResult, secondResult string

func concatString() {
	str := ""

	for i := 0; i < 5000; i++ {
		for j := 0; j < len(snippets); j++ {
			str += snippets[j]
		}
	}

	firstResult = str
}

func concatStringBuilder() {
	var str strings.Builder

	for i := 0; i < 5000; i++ {
		for j := 0; j < len(snippets); j++ {
			_, err := str.WriteString(snippets[j])
			if err != nil {
				panic(err)
			}
		}
	}

	secondResult = str.String()
}
