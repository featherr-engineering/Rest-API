# Featherr Rest API


**CircleCI Build** [![CircleCI](https://circleci.com/gh/featherr-engineering/rest-api/tree/develop.svg?style=svg)](https://circleci.com/gh/featherr-engineering/rest-api/tree/develop) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/914c33d772094350ab7535716a5970eb)](https://www.codacy.com/app/featherr-engineering/rest-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=featherr-engineering/rest-api&amp;utm_campaign=Badge_Grade)

&nbsp;

Golang rest api for Featherr mobile app. Repo contains src, and build. Using git flow for version control.


## Codebase Architecture
- **Packages**
  - [Gorrila/Mux](https://github.com/gorilla/mux) A powerful URL router and dispatcher for golang.
  - [Gorm](https://github.com/jinzhu/gorm) The fantastic ORM library for Golang, aims to be developer friendly.
  - [GoDotEnv](https://github.com/joho/godotenv) A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) 
  - [JWT-Go](https://github.com/dgrijalva/jwt-go)  Golang implementation of JSON Web Tokens (JWT).
  - [UUID](https://github.com/satori/go.uuid)  UUID package for Go.
  
- ****
  - [Go](https://golang.org/doc/devel/release.html#go1.11) statically typed, compiled programming language 
  - [Mysql](https://www.mysql.com/)  open-source relational database management system.
  - [Go Deps](https://github.com/golang/dep) Go dependency management tool 
  
&nbsp;
  
  
## Getting started :raised_hands:

### Install Golang
- Install react-native-cli. And then install dependencies. Run the project using the cli. And you're all set!
    - https://golang.org/doc/install
- make sure your ~/.*shrc have those varible:
    - ```
      ➜  echo $GOPATH
      /Users/zitwang/test/
      ➜  echo $GOROOT
      /usr/local/go/
      ➜  echo $PATH
      ...:/usr/local/go/bin:/Users/zitwang/test//bin:/usr/local/go//bin
      ```
- Install package 
- Run ```dep ensure```
- Create .env file

&nbsp;
&nbsp;
&nbsp;

### Project Workflow :zap:


> Use git flow for working on changes

```shell
# Create a new feature branch
$ git flow feature start feature_branch_name
# Finishing a feature branch
$ git flow feature finish feature_branch
```

- Working on features
  - Each new feature should reside in its own branch
  - feature branches use `develop` as their parent branch

### Testing
```
➜  go test -v ./tests
```


### Deploying / Publishing :rocket:

- Review circle ci config
- Release branches
    - ```shell
      # Create a release branch
      $ git flow release start x.x.x
      
      # Finish a release branch
      $ git flow release finish 'x.x.x'
      ```
- Approve build in CircleCI UI
    