# **Shelter Api**

>[!NOTE]
>By @FloatStudio2024

## **Introduction**
This api uses Golang and Gin framework to create a RESTful API that can return a sentence from the hitokoto database. 

## **Getting Started**
Visit the [GitHub repository](https://github.com/FloatStudio2024/shelter_api) to download the source code and follow the instructions below to run the API.
### **Prerequisites**
- Go 1.14 or later
- Gin framework

### **Installing**
1. Clone the repository to your local machine:
```
git clone https://github.com/FloatStudio2024/shelter_api.git
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
>[!NOTE]
>You can edit the `config.json` file to change the port.

## **License**
This project is licensed under the AGPL-3.0 license.

## **Acknowledgments**
- [hitokoto](https://hitokoto.cn/)