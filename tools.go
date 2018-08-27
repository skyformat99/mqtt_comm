package mqtt_comm

import (
	// "fmt"
	"strings"
)

func GetResponseTopic(serverVersion string, serverName string, action string, id string) string {
	return strings.Join([]string{"/", serverVersion, "/", serverName, "/", action, "/", id}, "")
}

func GetResponseUri(serverVersion string, serverName string) string {
	return strings.Join([]string{"/", serverVersion, "/", serverName, "/+/+"}, "")
}

func GetSubscribeUri(action string, topic string) string {
	length := len(topic)
	end := []byte(topic)[length-1]
	if string(end) != "/" {
		topic += "/"
	}
	return strings.Join([]string{"/+/+/", action, "/+/", topic}, "")
}

func GetFullUri(serverVersion string, serverName string, action string, topic string, id string) string {
	length := len(topic)
	end := []byte(topic)[length-1]
	if string(end) != "/" {
		topic += "/"
	}
	return strings.Join([]string{"/", serverVersion, "/", serverName, "/", action, "/", id, "/", topic}, "")
}

func SpliteFullUri(uri string) (serverVersion string, serverName string, action string, id string, topic string) {
	strs := strings.Split(uri, "/")
	var top string
	length := len(strs)
	for i := 5; i < length; i++ {
		top += strs[i]
		if i < length-1 {
			top += "/"
		}
	}
	return strs[1], strs[2], strs[3], strs[4], top
}

func SpliteResponseUri(uri string) (action string, id string) {
	strs := strings.Split(uri, "/")
	return strs[3], strs[4]
}
