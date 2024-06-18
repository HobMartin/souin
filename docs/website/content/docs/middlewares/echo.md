+++
weight = 505
title = "Echo"
icon = "extension"
description = "Use Souin directly in the Echo web server"
tags = ["Beginners", "Advanced"]
+++

## Usage
Here is the example about the Souin initialization.
```go
import (
	"net/http"

	souin_echo "github.com/HobMartin/souin/plugins/echo"
	"github.com/labstack/echo/v4"
)

func main(){

    // ...
	s := souin_echo.New(souin_echo.DevDefaultConfiguration)
	e.Use(s.Process)
    // ...

}
```
With that your application will be able to cache the responses if possible and returns at least the `Cache-Status` HTTP header with the different directives mentionned in the RFC specification.  
You have to pass a Souin `BaseConfiguration` structure into the `NewHTTPCache` method (you can use the `DefaultConfiguration` variable to have a built-in production ready configuration).  

Look at the configuration section to discover [all configurable keys here]({{% relref "/docs/configuration" %}}).

Other resources
---------------
You can find an example for a docker-compose stack inside the [`examples` folder on the Github repository](https://github.com/HobMartin/souin/tree/master/plugins/echo/examples).
Look at the [`BaseConfiguration` structure on pkg.go.dev documentation](https://pkg.go.dev/github.com/HobMartin/souin/pkg/middleware#BaseConfiguration).