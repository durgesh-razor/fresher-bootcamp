module studentRecord

go 1.16

//replace api/Config => ./Config

//replace api/Routes => ./Routes

//replace api/Models => ./Models

//replace api/Controllers => ./Controllers

require (
	github.com/gin-gonic/gin v1.7.2
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)
