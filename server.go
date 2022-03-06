/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/microsoft/vscode-remote-try-go/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: randomhellos")
	json.NewEncoder(w).Encode(hello.RandomHellos())
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/randomhellos", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
