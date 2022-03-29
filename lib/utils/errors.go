package utils

import "github.com/gin-gonic/gin"

func ErrStatus(m gin.H, err error) gin.H {
	m["error"] = err.Error()
	return m
}
