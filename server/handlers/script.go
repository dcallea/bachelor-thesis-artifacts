package handlers

import (
	"net/http"
	"os/exec"
	"fmt"
)

func runScript() (string, error) {
	// path to shell script
	scriptPath := "/home/boss/ftp/date.sh"
	// execute the script
	cmd := exec.Command("/bin/bash", scriptPath)
	// capture output
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func ScriptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		output, err := runScript()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing script: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, output) // send script output
	}
}
