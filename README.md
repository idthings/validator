# validator

[![ci](https://github.com/idthings/validator/actions/workflows/ci.yaml/badge.svg)](https://github.com/idthings/validator/actions/workflows/ci.yaml)

A validator module library.

#### Install
```
go get github.com/idthings/validator
```

#### Use
```
package somepkg

import (
	"github.com/idthings/validator"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

    // function truncated for clarity

    if !validator.IsValidEmail(r.FormValue("email")) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

}
```
