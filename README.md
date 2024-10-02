# **Shelter Api**

![GitHub Repo stars](https://img.shields.io/github/stars/Thenemophilists/shelter_api?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/Thenemophilists/shelter_api?style=flat-square)
![GitHub License](https://img.shields.io/github/license/Thenemophilists/shelter_api?style=flat-square)

## **Introduction**

This api uses Golang and Gin framework to create a RESTful API that can return a sentence from the hitokoto database. 

## **Getting Started**

Visit the [GitHub repository](https://github.com/Evereen2023/shelter_api) to download the source code and follow the instructions below to run the API.

### **Prerequisites**

- Go 1.22.5
- Gin framework

### **Installing**

1. Clone the repository to your local machine:

```
git clone https://github.com/Evereen2023/shelter_api.git
```

2. Install the required packages:

```
go mod download
```

3. Run the API:

```
go run main.go
```

4. Open your browser and go to `http://localhost:8080/sentence` to see the sentences.

### **The Database**
Download the monogoDB version of hitokoto in https://github.com/Evereen2023/hikotoko_db.git

## **License**

This project is licensed under the AGPL-3.0 license.

## **Acknowledgments**

- [hitokoto](https://hitokoto.cn/)
