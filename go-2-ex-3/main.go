package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Grundlagen der Programmierung",
		117: "Datenbanken entwerfen und nutzen",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	// Ausgabe der ursprünglichen Module
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)

	modules[205] = "Web-Anwendungen entwickeln"

	modules[346] = "Cloud-Architekturen planen"

	fmt.Println(modules)
}
