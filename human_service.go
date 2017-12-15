package human

type Human struct {
	Name    string
	Surname string
	Age     int
	Gender  string
}

var humans []Human

func init() {
	humans = append(humans, Human{"Batuhan", "Apaydın", 23, "Male"})
	humans = append(humans, Human{"Asena", "Apaydın", 19, "Female"})
	humans = append(humans, Human{"Mehmet", "Apaydın", 50, "Male"})
}

func HumanByName(name string) interface{} {
	for _, human := range humans {
		if human.Name == name {
			return human
		}
	}
	return nil
}
