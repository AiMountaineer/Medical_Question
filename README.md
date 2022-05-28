# iris_web

* go version go1.17.3

* links
  * <https://www.iris-go.com/docs/#/>
  * <https://github.com/kataras/iris/>

* Golang ORM
  * <https://gorm.io/ja_JP/docs/models.html>

* 環境設定
  *
  * ```zsh:環境設定
    mkdir myapp
    go mod init myapp
    go get github.com/kataras/iris/v12@master
    ```

* swagger: <https://github.com/iris-contrib/swagger>

* 開発環境：http://0.0.0.0:8085/swagger/index.html

```zsh:path
vi .zshrc
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

* Start using it:
  1. Add comments to your API source code, See Declarative Comments Format.
  2. Download Swag for Go by using:

   ```zsh
   go get -u github.com/swaggo/swag/cmd/swag
   ```

  3. Run the Swag in your Go project root folder which contains main.go file, Swag will parse comments and generate required files(docs folder and docs/doc.go).

  ```zsh
  swag init
  ```

  4. Download swagger for Iris by using:

  ```zsh
  go get github.com/iris-contrib/swagger/v12@v12.2.0-alpha
  ```
  Reference information
https://trello.com/c/LiYNA6tU/55-swagger%E9%85%8D%E7%BD%AE

* database:
  * endpoint: medicalquestionnaireservice.cz7nrjgqkgdb.ap-northeast-1.rds.amazonaws.com
  * port: 3306

  * username: admin
  * password: Okulo202127!


lastly:

```:command
go run main.go
```

live-reloading：

```:command
air
```
