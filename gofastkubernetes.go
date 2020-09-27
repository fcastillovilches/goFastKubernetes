package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

type comando struct {
	Name string
	Id   int
}

type namespace struct {
	Name string
	Id   int
}

type pod struct {
	Name string
	Id   int
}

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

func opciones() int {

	// Crear arreglo de Opciones

	comandos := []comando{
		{Name: "exec", Id: 1},
		{Name: "describe", Id: 2},
		{Name: "logs", Id: 3},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F527 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }} ",
		Selected: "\U0001F527 {{ .Name | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select a command option",
		Items:     comandos,
		Templates: templates,
		Size:      4,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 99 // Codigo de error
	}

	fmt.Printf("You choose number %d: %s\n", i+1, comandos[i].Name)

	return i

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

	// Var id de opcion de comando
	var i int

	// Var id de opcion de namespaces
	var n int

	title()

	i = opciones()

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

	// count := 1
	// for id_namespaces := range delete_empty(lines) {
	// 	fmt.Printf("Option %d - %s  \n", id_namespaces, lines[id_namespaces])
	// 	count++
	// }

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F30D {{ . | cyan }}",
		Inactive: "{{ . | cyan }}",
		Selected: "\U0001F30D {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select a namespace option",
		Items:     lines,
		Templates: templates,
		Size:      4,
	}

	n, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", n+1, lines[n])

	// fmt.Printf("You choose number %d: %s\n", i+1, lines[i])

	fmt.Print("\n")
	// fmt.Print("Select a namespace option (Option Number) \n\n")

	// // Read integer
	// fmt.Scanf("%d", &n)

	clear()
	// log.Print(n)

	// fmt.Print("Pods Options (Option Number) \n\n")

	cmd_pods := "kubectl  get   pod  -n  " + lines[n] + " | sed -n '1!p' | cut -d ' ' -f1 "
	stdout_pods, err_pods := exec.Command("bash", "-c", cmd_pods).Output()

	if err_pods != nil {
		fmt.Printf("Failed to execute command: \n")
	}

	lines_pod := strings.Split(string(stdout_pods), "\n")

	// for id_pod := range delete_empty(lines_pod) {
	// 	fmt.Printf("Option %d - %s  \n", id_pod, lines_pod[id_pod])
	// }

	templatesPod := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F588 {{ . | cyan }}",
		Inactive: "{{ . | cyan }}",
		Selected: "\U0001F588 {{ . | red | cyan }}",
	}

	promptPod := promptui.Select{
		Label:     "Select a pod option",
		Items:     lines_pod,
		Templates: templatesPod,
		Size:      4,
	}

	p, _, errPod := promptPod.Run()

	if errPod != nil {
		fmt.Printf("Prompt failed %v\n", errPod)
		return
	}

	fmt.Printf("You choose number %d: %s\n", p+1, lines_pod[p])

	// fmt.Print("\n")
	// fmt.Print("Select a pod option\n\n")

	// // Read integer
	// var p int
	// fmt.Scanf("%d", &p)

	clear()

	if i == 0 {

		cmd_command := "kubectl exec -it " + lines_pod[i] + " -n " + lines[n] + " bash"
		_, err_command := exec.Command("bash", "-c", cmd_command).Output()

		if err_command != nil {
			fmt.Printf("Failed to execute command: %s \n", cmd_command)
		} else {
			fmt.Printf("Copy, Paste and Run this command: \n")
			fmt.Printf(cmd_command)
			fmt.Print("\n")

		}

	} else if i == 1 {

		cmd_command := "kubectl describe pod " + lines_pod[i] + " -n " + lines[n]
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
