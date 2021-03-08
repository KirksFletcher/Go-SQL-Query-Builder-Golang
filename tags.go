// Copyright 2020 The GoQueryBuilder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package querybuilder

import (
	"reflect"
	"strings"
)

// TODO: Temporary placeholder for struct tags.

const (
	tagName    = "sqlb"
	skipAll    = "skipAll"
	createTime = "autoCreateTime"
	updateTime = "autoUpdateTime"
)

func getTags(field reflect.StructField) (string, bool) {
	return field.Tag.Lookup(tagName)
}

func splitTags(str string) []string {
	return strings.Split(str, ",")
}

func isSkipped(tags []string) bool {
	for _, v := range tags {
		if v == skipAll {
			return true
		}
	}
	return false
}

func isAutoCreateTime(value string) bool {
	return strings.Contains(value, createTime)
}

func isAutoUpdateTime(value string) bool {
	return strings.Contains(value, updateTime)
}