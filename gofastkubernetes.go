package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func title() {
	// fmt.Print("GOFASTKUBERNETES\n\n")
	// basic stuff

	asciiArt :=
		`
	####  ###    ###   #   ### #####   # ## #  # #### #### ###  #  # #### ##### #### ### 
	##    ## ##   #    ###  #     #     # #  #  # #  # #    # #  ## # #      #   #    #   
	#     #   #   #    # #  ##    #     ##   #  # ###  ###  # #  ## # ###    #   ###  ##  
	#   # #   #   ###  ###   ##   #     ##   #  # # ## #    ##   # ## #      #   #     ## 
	##  # ## ##   #   #   #   #   #     # #  #  # #  # #    # #  # ## #      #   #      # 
	 ####  ###    #   #   # ###   #     #  # #### ###  #### #  # #  # ####   #   #### ### 
	`
	fmt.Println(asciiArt)
}

func opciones() {

	// Crear arreglo de Opciones
	opciones := [3]string{"exec", "describe", "logs"}

	//fmt.Println(opciones)
	//fmt.Println(len(opciones))
	fmt.Print("Command Options\n\n")

	for id_opcion, gl_opcion := range opciones {
		fmt.Printf("Option %d - %s\n", id_opcion, gl_opcion)
	}

	fmt.Print("\n")
	fmt.Print("Select a command option (Option Number) \n\n")
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {

	clear()

	var i int

	title()

	for {
		opciones()

		// Read integer
		fmt.Scanf("%d", &i)
		// fmt.Print(i)

		if i < 3 {
			break
		} else {

			clear()

			title()

			fmt.Print("Invalid Option\n\n")
		}
	}

	clear()

	fmt.Print("Namespaces Options\n\n")

	cmd_namespaces := "kubectl get namespace | grep Active | cut -d ' ' -f1"
	stdout_namespaces, err_namespaces := exec.Command("bash", "-c", cmd_namespaces).Output()

	// fmt.Print(string(stdout_namespaces))
	// fmt.Print("\n")

	if err_namespaces != nil {
		fmt.Printf("Failed to execute command: %s", cmd_namespaces)
	}

	lines := strings.Split(string(stdout_namespaces), "\n")

	for id_namespaces := range delete_empty(lines) {
		fmt.Printf("Option %d - %s  \n", id_namespaces, lines[id_namespaces])
	}

	fmt.Print("\n")
	fmt.Print("Select a namespace option (Option Number) \n\n")

	// Read integer
	var n int
	fmt.Scanf("%d", &n)

	clear()
	// log.Print(n)

	fmt.Print("Pods Options (Option Number) \n\n")

	cmd_pods := "kubectl  get   pod  -n  " + lines[n] + " | sed -n '1!p' | cut -d ' ' -f1 "
	stdout_pods, err_pods := exec.Command("bash", "-c", cmd_pods).Output()

	if err_pods != nil {
		fmt.Printf("Failed to execute command: \n")
	}

	lines_pod := strings.Split(string(stdout_pods), "\n")

	for id_pod := range delete_empty(lines_pod) {
		fmt.Printf("Option %d - %s  \n", id_pod, lines_pod[id_pod])
	}

	fmt.Print("\n")
	fmt.Print("Select a pod option\n\n")

	// Read integer
	var p int
	fmt.Scanf("%d", &p)

	clear()

	if i == 0 {

		cmd_command := "kubectl exec -it " + lines_pod[p] + " -n " + lines[n] + " bash"
		_, err_command := exec.Command("bash", "-c", cmd_command).Output()

		if err_command != nil {
			fmt.Printf("Failed to execute command: %s \n", cmd_command)
		} else {
			fmt.Printf("Copy, Paste and Run this command: \n")
			fmt.Printf(cmd_command)
			fmt.Print("\n")

		}

	} else if i == 1 {

		cmd_command := "kubectl describe pod " + lines_pod[p] + " -n " + lines[n]
		stdout_command, err_command := exec.Command("bash", "-c", cmd_command).Output()

		if err_command != nil {
			fmt.Printf("Failed to execute command: %s \n", cmd_command)
		} else {
			fmt.Print(string(stdout_command))
			fmt.Print("\n")
		}

	} else {

		cmd_command := "kubectl logs  " + lines_pod[p] + " -n " + lines[n]
		stdout_command, err_command := exec.Command("bash", "-c", cmd_command).Output()

		if err_command != nil {
			fmt.Printf("Failed to execute command: %s \n", cmd_command)
		} else {
			fmt.Print(string(stdout_command))
		}

	}

}
