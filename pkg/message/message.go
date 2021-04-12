/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

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
package message

import (
	"github.com/romberli/go-util/config"
)

const DefaultMessageHeader = "DAS"

var Messages = map[int]*config.ErrMessage{}

func init() {
	initInfoMessage()
	initErrorMessage()
}

// NewMessage returns *config.ErrMessage with specified values
func NewMessage(code int, values ...interface{}) *config.ErrMessage {
	return Messages[code].Renew(values...)
}
