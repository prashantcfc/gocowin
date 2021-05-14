# gocowin
Golang based API bindings for CoWIN project. 

This project also includes a compiled binary that can be used on Windows,
for checking slot availability of COVID-19 vaccines in India. 

In addition, you can clone the source code & build your own binary on Windows or Linux platform using:

``go build``

## CLI Examples

In order to fetch list of state codes, use:

```./gocowin.exe --states```

In order to fetch district codes of a state, use:

```./gocowin.exe --stateId "12"```

In order to check COVID-19 vaccine availability by districtId and date, use:

```./gocowin.exe --districtId "188" --date "15-05-2021"```

In order to check COVID-19 vaccine availability by pin and date, use:

```./gocowin.exe --pin "110005" --date "21-05-2021"```




## API Examples 

```
import (
	"fmt"
	"github.com/prashantcfc/gocowin/api"
)

func main() {
	res, err := api.GetSessionsbyDistrictIdAndDate("141", "15-05-2021")
	if err != nil {
		return
	}
	fmt.Println(res)
}
```

## Contributions

Contributions to this project are welcome. Let's build a stronger community together!!

## Support

Please raise github issues incase of any concerns.
This is a 'side' project & developer might take time to revert back.
