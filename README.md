# urlshortener

DockerHub Image : https://hub.docker.com/repository/docker/kudligi97/urlshortener

## Endpoints

 1. **POST** /shorten
    > **Request Body Schema**

    > long_url  string, required, url

    > **Response Body Schema**

    > long_url  string, url

    > short_url string, url

    Sample request body :
    ```
    {
    	"long_url" : "https://github.com/Laeeqdev"
    }
    ```
    Sample response body :
    ```
    {

    "long_url":  "https://github.com/Laeeqdev",

    "short_url": "http://localhost:9090/96XD4spw"

    }
    ```
if we generate short url for the same given link again and again it will give same shorten url for life time

  2. **GET** http://localhost:9090/96XD4spw
    > Redirects to the registered long url (actual url)

 3. **POST** /metrics
    it will list out top 3 domains
    > **Response Body Schema**
  ```
  [
    {
        "Domain": "github.com",
        "Count": 4
    },
    {
        "Domain": "localhost.com",
        "Count": 2
    },
    {
        "Domain": "localhost",
        "Count": 1
    }
]

```
 4. **POST** /lengthen
    > **Request Body Schema**

    > short_url  string, required, url

    > **Response Body Schema**

    > long_url  string, url

    > short_url string, url

    Sample request body :
    ```
    {
    	"short_url" : "http://localhost:9090/96XD4spw"
    }
    ```
    Sample response body :
    ```
    {

    "long_url":  "https://github.com/Laeeqdev",

    "short_url": "http://localhost:9090/96XD4spw"

    }
    ```

## Design

 - AppLayer
	 - Config loaded from environement variables
   - App
   - Routers
   - Handlers
   - Services
   - Repository
   - Utils
   - Constants
   - In Memory Data Store:
	  - Map + RwMutex  
## Unit Testing
Minimal tests for:

 - Router
 - DataService
 - RandomUrl Utils
# Running the application
## To run this application locally 
> $ git clone https://github.com/Laeeqdev/urlShortner.git

> $ cd urlShortner

> $ go run .

Server will be listening on http://localhost:9090/

## to start application using docker image

> $ docker pull 

> $ docker run -p 9090:9090 laeeqa222/url_shortener:latest

Server will be listening on http://localhost:9090/


 if face any issue then please drop a mail on Laeeqa222@gmail.com or open the issue thank you
