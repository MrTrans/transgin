package transutils

// import (
// 	"os"
// 	"path"

// 	"github.com/qiniu/x/log.v7"
// )

// // JoinPaths construct absolutePath
// func JoinPaths(absolutePath, relativePath string) string {
// 	if relativePath == "" {
// 		return absolutePath
// 	}

// 	finalPath := path.Join(absolutePath, relativePath)
// 	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
// 	if appendSlash {
// 		return finalPath + "/"
// 	}
// 	return finalPath
// }

// func lastChar(str string) uint8 {
// 	if str == "" {
// 		panic("The length of the string can't be 0")
// 	}
// 	return str[len(str)-1]
// }

// // ResolveAddress 解释地址
// func ResolveAddress(addr []string) string {
// 	switch len(addr) {
// 	case 0:
// 		if port := os.Getenv("PORT"); port != "" {
// 			log.Warnf("Environment variable PORT=\"%s\"", port)
// 			return ":" + port
// 		}
// 		log.Warnf("Environment variable PORT is undefined. Using port :8080 by default")
// 		return ":8086"
// 	case 1:
// 		return addr[0]
// 	default:
// 		panic("too many parameters")
// 	}
// }
