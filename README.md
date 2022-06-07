# SPapi Golang
## API для spworlds.ru 


## Features

- Создание транзакций
- Создание переводов
- Ник по Discord id

## Установка

```sh
go get -u github.com/purp1le/spgo
```
## Создание платежа

```
package main

import (
	"fmt"
	"github.com/purp1le/spgo"
)

func main() {
	sp := spworlds.New("your id", "your token")
	url, err := sp.NewPayment(sum(int), "success url", "notification url", "text")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)
}
```

## Создание транзакции
```
package main

import (
	"fmt"
	"github.com/purp1le/spgo"
)

func main() {
	sp := spworlds.New("id", "token")
	res, err := sp.NewTransaction("83934", "hi", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
```

## Узнать ник

```
package main

import (
	"fmt"
	"github.com/purp1le/spgo"
)

func main() {
	sp := spworlds.New("id", "token")
	res, err := sp.GetName("Discord id")
	fmt.Println(res)
}
```


