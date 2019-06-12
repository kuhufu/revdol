package valid

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

//shortcut of map[string][string]
type M = map[string]string

func Params(rules M) gin.HandlerFunc {
	regexps := buildRegexps(rules)
	return func(c *gin.Context) {
		for param, r := range regexps {
			if !r.MatchString(c.Param(param)) {
				c.AbortWithStatusJSON(404, gin.H{
					"error": "wrong " + param,
				})
				return
			}
		}
		c.Next()
	}
}

func Param(param, expr string) gin.HandlerFunc {
	regexp := buildRegex(expr)
	return func(c *gin.Context) {
		if !regexp.MatchString(c.Param(param)) {
			c.AbortWithStatusJSON(404, gin.H{
				"error": "wrong " + param,
			})
			return
		}
		c.Next()
	}
}

func Querys(rules M) gin.HandlerFunc {
	regexps := buildRegexps(rules)
	return func(c *gin.Context) {
		for query, r := range regexps {
			if !r.MatchString(c.Query(query)) {
				c.AbortWithStatusJSON(404, gin.H{
					"error": "wrong " + query,
				})
				return
			}
		}
		c.Next()
	}
}

func Query(param, expr string) gin.HandlerFunc {
	regexp := buildRegex(expr)
	return func(c *gin.Context) {
		if !regexp.MatchString(c.Query(param)) {
			c.AbortWithStatusJSON(404, gin.H{
				"error": "wrong " + param,
			})
			return
		}
		c.Next()
	}
}

func Forms(rules M) gin.HandlerFunc {
	regexps := buildRegexps(rules)
	return func(c *gin.Context) {
		for field, r := range regexps {
			if !r.MatchString(c.PostForm(field)) {
				c.AbortWithStatusJSON(404, gin.H{
					"error": "wrong " + field,
				})
				return
			}
		}
		c.Next()
	}
}

func Form(field, expr string) gin.HandlerFunc {
	regexp := buildRegex(expr)
	return func(c *gin.Context) {
		if !regexp.MatchString(c.PostForm(field)) {
			c.AbortWithStatusJSON(404, gin.H{
				"error": "wrong " + field,
			})
			return
		}
		c.Next()
	}
}

func buildRegexps(rules M) map[string]*regexp.Regexp {
	regexps := map[string]*regexp.Regexp{}
	for param, expr := range rules {
		r, err := regexp.Compile(expr)
		if err != nil {
			panic("wrong regexp string, compile error")
		}
		regexps[param] = r
	}
	return regexps
}

func buildRegex(expr string) *regexp.Regexp {
	regexp, err := regexp.Compile(expr)
	if err != nil {
		panic("wrong regexp string, compile error")
	}
	return regexp
}
