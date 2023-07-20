/*
Copyright © 2021 Upvest GmbH <support@upvest.co>

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

package material

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func GetRequestBody(req *http.Request) ([]byte, error) {
	reader := io.NopCloser(bytes.NewBufferString(""))
	if req.GetBody != nil {
		var err error
		reader, err = req.GetBody()
		if err != nil {
			return nil, errors.Wrap(err, "AddSignatureHeaders: Get Body error")
		}
		defer func() {
			_ = reader.Close()
		}()
	} else if req.Body != nil {
		reader = req.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "AddSignatureHeaders: ReadAll error")
	}
	return body, nil
}

func Format(k, v string) string {
	return fmt.Sprintf("\"%s\": %s", k, v)
}
