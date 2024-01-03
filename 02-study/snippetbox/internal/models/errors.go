/**
 * file: internal/models/errors.go
 * description: file responsible for the errors of the application.
 * data: 01/02/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")
