/*
 Copyright 2023 adamswanglin

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package internal

import (
	"net/http"
	"time"
	"unicode/utf8"

	"github.com/blinkbean/dingtalk"
	"github.com/golang/glog"
)

type Body struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
}

// sendToDingTalkRobot 发送到钉钉机器人
func sendToDingTalkRobot(token, secret, title, text string) {

	//钉钉限制4000字符
	if utf8.RuneCountInString(text) > 3500 {
		runes := []rune(text)
		text = string(runes[:3500]) + "..."
	}

	theDingTalkBot := dingtalk.InitDingTalkWithSecret(token, secret)
	err := theDingTalkBot.SendTextMessage(title + "\n\n" + text)

	if err != nil {
		glog.Errorf("Dingtalk sent error to %s, %s, title %s, error: %w", token, secret, title, err)
	}

}
