package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/redis.v5"
	log "github.com/golang/glog"
	"net/http"
)

const (
	ERR_ALREADY_BEEN_LIKED = "object already been liked"
)

var (
	redisCli *redis.Client
)

func like(c *gin.Context) {
	oid := c.Param("oid")
	uid := c.Param("uid")
	cmd := redisCli.SAdd(uid, oid)
	code := http.StatusOK
	message := http.StatusText(code)
	if n, err := cmd.Result(); err != nil {
		log.Error(err)
		code = http.StatusInternalServerError
		message = http.StatusText(code)
	} else if n == 0 {
		code = http.StatusConflict
		message = ERR_ALREADY_BEEN_LIKED
	}
	c.JSON(http.StatusOK, gin.H{
		"error_code": code,
		"error_message": message,
		"oid": oid,
		"uid": uid,
	})
}
func isLike(c *gin.Context) {
	oid := c.Param("oid")
	uid := c.Param("uid")
	cmd := redisCli.SIsMember(uid, oid)
	if yes, err := cmd.Result(); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"error_code": http.StatusInternalServerError,
			"error_message": http.StatusText(http.StatusInternalServerError),
			"oid": oid,
			"uid": uid,
		})
	} else {
		var isLike int
		if yes {
			isLike = 1
		}
		c.JSON(http.StatusOK, gin.H{
			"oid": oid,
			"uid": uid,
			"is_like": isLike,
		})
	}
}
func count(c *gin.Context) {
	oid := c.Param("oid")
	cmd := redisCli.Get(oid)
	if val, err := cmd.Result(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_code": http.StatusInternalServerError,
			"error_message": http.StatusText(http.StatusInternalServerError),
			"oid": oid,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"oid": oid,
			"count": val,
		})
	}
}
func list(c *gin.Context) {

}
