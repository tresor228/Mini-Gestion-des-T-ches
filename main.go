package main

import (
	"fmt"
)

// ajout des tâches
func AjoutTasks(Titre string, Description string) {
	ID := len(tasks) + 1
	Tache := Task{
		ID:          ID,
		Titre:       Titre,
		Description: Description,
		Etat:        false,
	}
	tasks = append(tasks, Tache)
	fmt.Println("Tâche ajouté", Tache)

}

// Affichage des tâches
func AffichageDesTache() {
	if len(tasks) == 0 {
		fmt.Println("Aucune tâche disponible")
		return
	}
	for _, task := range tasks {
		Status := "Non Terminé"
		if task.Etat {
			Status = "Terminé"
		}
		fmt.Printf("L'ID %d , Titre %s, Status %s", task.ID, task.Titre, Status)
	}
}

// Marquer qu'une tâche est complète
func Tachecomplet(ID int) {
	for i, task := range tasks {
		if task.ID == ID {
			tasks[i].Etat = true
			fmt.Println("Tâche terminée", task.Titre)
			return
		}
	}
	fmt.Println(" Tache non disponible")
}

// Supprimer une tâche
func SuppTache(ID int) {
	for i, task := range tasks {
		if task.ID == ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
		fmt.Println("Tâches supprimé!", task.ID)
		return
	}
	fmt.Println("Tâche introuvable")
}
func main() {
	for {
		fmt.Println("====== Gestion des tâches====== \n")
		fmt.Println("1. Ajoutez une tâche \n")
		fmt.Println("2. Affichez une tâche \n")
		fmt.Println("3. Marquez une tâche comme terminé \n")
		fmt.Println("4. Supprimez une tâche \n")
		fmt.Println("5. Sortir du programme \n")
		fmt.Println("Choisisez une option \n")

		var choix int
		fmt.Scan(&choix)

		//Examination du choix
		switch choix {
		case 1:
			var Titre, Description string
			fmt.Println("Veuillez saisir le Titre: ")
			fmt.Scanln(&Titre)
			fmt.Println("Veuillez saisir la description")
			fmt.Scanln(&Description)
			AjoutTasks(Titre, Description)
		case 2:
			AffichageDesTache()
		case 3:
			var ID int
			fmt.Println("Veuillez saisir l'ID")
			fmt.Scan(&ID)
			Tachecomplet(ID)
		case 4:
			var ID int
			fmt.Println("Veuillez saisir l'ID")
			fmt.Scan(&ID)
			SuppTache(ID)
		case 5:
			fmt.Println("Au revoir")
			return
		default:
			fmt.Println("option invalide")
		}

	}

}
