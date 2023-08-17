package ipchecker

import (
	"github.com/gin-gonic/gin"
	"github.com/tristan-club/kit/config"
	"github.com/tristan-club/kit/log"
	"net/http"
	"strings"
)

// GetRealIP retrieves the user's real IP address.
// This method is suitable for HTTP services configured with load balancing.
// Using the ClientIP() method from gin.context might fetch a potentially spoofed X-Forwarded-For data by users.
// By using this method, as long as there is a load balancer configured, the correct IP can always be retrieved.
func GetRealIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ipList := strings.Split(xForwardedFor, ",")
	// 取列表中的最后一个 IP
	realIP := strings.TrimSpace(ipList[len(ipList)-1])
	return realIP
}

func GetRealIPFromGin(c *gin.Context) string {
	realIP := GetRealIP(c.Request)
	if realIP == "" {
		if !config.EnvIsDev() {
			log.Error().Fields(map[string]interface{}{"action": "get real ip from x-forwarded-for failed", "header": c.Request.Header}).Send()
		}
		realIP = c.ClientIP()
	}
	return realIP
}
