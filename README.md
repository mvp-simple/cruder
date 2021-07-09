# CRUDER



[![golang](https://user-images.githubusercontent.com/26686461/125044308-a8c73380-e0bd-11eb-98ca-06795b3021c3.png)](https://golang.org/) [![gorm](https://user-images.githubusercontent.com/26686461/125044305-a82e9d00-e0bd-11eb-8f89-50e72053af47.png)](http://gorm.io/) [![crud](https://user-images.githubusercontent.com/26686461/125044299-a6fd7000-e0bd-11eb-99d9-2a7d97f3c066.png)](https://ru.wikipedia.org/wiki/CRUD)

[![Test](https://github.com/mvp-simple/cruder/actions/workflows/tests.yml/badge.svg)](https://github.com/mvp-simple/cruder/actions/workflows/tests.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Пример.

```
package main  
  
import (  
	"net/http"  
	"time"  
	// пакеты экосистемы gorm  
	"gorm.io/driver/postgres"  
	"gorm.io/gorm"  
	// пакеты экосистемы cruder  
	"github.com/mvp-simple/cruder"  
	"github.com/mvp-simple/cruder-gormdb")  
  
// ShowStructure пример структуры
type ShowStructure struct {  
	ID          uint `gorm:"primarykey"`  
	CreatedAt   time.Time  
	UpdatedAt   time.Time  
	DeletedAt   gorm.DeletedAt `gorm:"index"`  
	Name        string  
	Description string  
}  
  
// SecondStructure пример структуры
type SecondStructure struct {  
	ID          uint `gorm:"primarykey"`  
	Name        string  
	Description string  
}  
  
func main() {  
	// инициализируем соединение с БД  
	db, err := gorm.Open(
		postgres.New(postgres.Config{DSN: "user=test password=test dbname=test port=5432 sslmode=disable TimeZone=Asia/Almaty", PreferSimpleProtocol: true}),
		&gorm.Config{},
	)  
	if err != nil {  
		panic(err)  
	}  
  
	// создаем экземпляр cruder.IDB  
	gormDB := cruder_gormdb.New(db)  
  
	// создаем cruder для структуры ShowStructure  
	cruderShowStructure := cruder.New(ShowStructure{}, gormDB)  
	cruderShowStructure.  
		// C	создаем метод Create		  
		Create().	
		// R	создаем метод Get		 
		Get().		
		// R	создаем метод Get Many	  
		List().	
		// U	создаем метод Update		  	
		Update().
		// U	создаем метод Restore	  	
		Restore().	
		// D	создаем метод Delete		  
		Delete()	
	// создаем cruder для структуры SecondStructure
	cruderSecondStructure := cruder.New(SecondStructure{}, gormDB)  
	cruderSecondStructure.  
		// C	создаем метод Create	
		Create().
		// R	создаем метод Get	
		Get().		
		// R	создаем метод Get Many	 
		List()		
  
	// Создаем роутер  
	router := cruder.NewRouter()  
  
	router.
		Add(cruderShowStructure).     // Подключаем сформированные доя этого крудеры к роутеру
		Add(cruderSecondStructure)    // Подключаем сформированные доя этого крудеры к роутеру  

    // запускаем сервер
	http.ListenAndServe(":2000", router)       
}
```



## Авторизация.

Для добавления авторизации необходимо вызвать метод SetIdentity у роутера и в нем указать функцию которая будет проверять имеет ли право пользователь проходить внутрь по маршруту.

```
func auth(ctx cruder.IContext) (out bool) {  
	if rand.Int63() < math.MaxInt64 / 2 {  
		return false  
	}  
	return true  
}

func main() {
	//...
	router.SetIdentity(auth)
	//...
}
```

Хорошим тоном является в функции отвечающей за авторизацию добавлять идентификатор пользователя через методы:

```
// IContext.SetUserID(in int64) (out IContext)  
// IContext.SetUserUUID(in string) (out IContext)

func auth(ctx cruder.IContext) (out bool) {  
	if rand.Int63() < math.MaxInt64 / 2 {  
		return false  
	}
	ctx.SetUserID(rand.Int63())				// ID 	если вы используете идентификатором тип	int
	ctx.SetUserUUID("0000-EXAMPLE-UUID")	// UUID	если Вы используете идентификатором тип	string
	return true  
}
```



## ICruderOption.

Параметры по умолчанию для функций CRUD:

**tableName** - таблица где хранится сущность  и где все точки заменены на **__**

> strings.ReplaceAll(tableName, ".", "__")

|функция|endpoint|http метод|  
|--|--|--|  
|Create|/api/**tableName**/create|POST|  
|Delete|/api/**tableName**/delete|DELETE|  
|Get|/api/**tableName**/get_by_id|GET|  
|List|/api/**tableName**/list|GET|  
|Restore|/api/**tableName**/restore|PATCH|  
|Update|/api/**tableName**/change|PUT|


Для гибкого изменения добавляемых функций CRUD в пакет добавлены ICruderOption.

```
	cruder.Uri(uriIn string)
	cruder.OneModelCreator(model interface{})
	cruder.Method(methodIn string)
	cruder.ManyModelCreator(modelSlice interface{})
```

* **функция** одна из функций CRUD.


### cruder.Uri(uriIn string)
Изменяет url по которому будет доступна функция


### cruder.OneModelCreator(model interface{})
Изменяет структуру с которой необходимо работать в функции.


### cruder.Method(methodIn string)
Изменяет метод по которому будет доступна функция.

В пакете предопределены методы для более упрощения использования cruder.Method
```
	// Get завернутый в Method http.MethodGet  
	cruder.Get = Method(http.MethodGet)  
	// Head завернутый в Method http.MethodHead  
	cruder.Head = Method(http.MethodHead)  
	// Post завернутый в Method http.MethodPost  
	cruder.Post = Method(http.MethodPost)  
	// Put завернутый в Method http.MethodPut  
	cruder.Put = Method(http.MethodPut)  
	// Patch завернутый в Method http.MethodPatch  
	cruder.Patch = Method(http.MethodPatch)  
	// Delete завернутый в Method http.MethodDelete  
	cruder.Delete = Method(http.MethodDelete)  
	// Connect завернутый в Method http.MethodConnect  
	cruder.Connect = Method(http.MethodConnect)  
	// Options завернутый в Method http.MethodOptions  
	cruder.Options = Method(http.MethodOptions)  
	// Trace завернутый в Method http.MethodTrace  
	cruder.Trace = Method(http.MethodTrace)
```


### cruder.ManyModelCreator(modelSlice interface{})
Изменяет слайс структур с которыми необходимо работать в функции.


### Примеры использования ICruderOption

```
// ShowStructure пример структуры
type ShowStructure struct {  
	ID          uint `gorm:"primarykey"`  
	CreatedAt   time.Time  
	UpdatedAt   time.Time  
	DeletedAt   gorm.DeletedAt `gorm:"index"`  
	Name        string  
	Description string  
}  
  
// SecondStructure пример структуры
type SecondStructure struct {  
	ID          uint `gorm:"primarykey"`  
	Name        string  
	Description string  
}  

func main() {
	// создаем cruder для структуры ShowStructure
	cruderShowStructure := cruder.New(ShowStructure{}, gormDB)
	cruderShowStructure.  
		// для функции Create будет использоваться PATCH метод запроса
		Create(cruder.Method(http.MethodPatch)).							
		// для функции Getбудет использоваться PUT метод запроса
		Get(cruder.Put).	
		// для функции List будет использваться структура SecondStructure 										
		List(cruder.OneModelCreator(SecondStructure{})).
		// функция Update с параметрами по умолчанию			
		Update().
		// функция Restore с параметрами по умолчанию													
		Restore().	
		// функция Delete доступна по url "/api_special/delete/model") и с методом POST 												
		Delete(cruder.Post,cruder.Uri("/api_special/delete/model"))	
}
```
