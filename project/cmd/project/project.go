package main

import (
	"github.com/SleepNFire/go_bootstrap/project/internal/app"
)

func main() {

	app.Init().Run()

	// appInstance := app.Init()

	// if err := appInstance.Start(context.Background()); err != nil {
	// 	fmt.Printf("Erreur lors du démarrage de l'application : %v\n", err)
	// }

	// // Attente du signal d'arrêt (Ctrl+C)
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)

	// select {
	// case <-quit:
	// 	fmt.Println("Arrêt de l'application par signal d'interruption (Ctrl+C)")
	// case <-appInstance.Done():
	// 	fmt.Println("L'application s'est arrêtée correctement.")
	// }

	// // Arrêt propre de l'application
	// if err := appInstance.Stop(context.Background()); err != nil {
	// 	fmt.Printf("Erreur lors de l'arrêt de l'application : %v\n", err)
	// }
}
