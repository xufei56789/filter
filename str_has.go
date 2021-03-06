package filter

import (
	"strings"
	"unicode"
)

// 必须包含字母(不区分大小写)
func (self *Str) HasLetter(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if unicode.IsLetter(v) {
			return self
		}
	}
	self.err = wrapError(self.name, "", customError...)
	return self
}

// 必须包含小写字母
func (self *Str) HasLower(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}

	for _, v := range self.currentValue {
		if unicode.IsLower(v) {
			return self
		}
	}

	self.err = wrapError(self.name, "", customError...)
	return self
}

// 必须包含大写字母
func (self *Str) HasUpper(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}

	for _, v := range self.currentValue {
		if unicode.IsUpper(v) {
			return self
		}
	}

	self.err = wrapError(self.name, "", customError...)
	return self
}

// 必须包含数字
func (self *Str) HasDigit(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}

	for _, v := range self.currentValue {
		if unicode.IsDigit(v) {
			return self
		}
	}

	self.err = wrapError(self.name, "", customError...)
	return self
}

// 必须包含符号
func (self *Str) HasSymbol(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}

	for _, v := range self.currentValue {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) && !unicode.Is(unicode.Han, v) {
			return self
		}
	}

	self.err = wrapError(self.name, "", customError...)
	return self
}

// 必须包含指定的字符串
func (self *Str) HasString(sub string, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if strings.Contains(self.currentValue, sub) {
		return self
	}
	self.err = wrapError(self.name, "", customError...)
	return self
}

// 必须包含指定的前缀字符串
func (self *Str) HasPrefix(sub string, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if strings.HasPrefix(self.currentValue, sub) {
		return self
	}
	self.err = wrapError(self.name, "", customError...)
	return self
}

// 必须包含指定的后缀字符串
func (self *Str) HasSuffix(sub string, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if strings.HasSuffix(self.currentValue, sub) {
		return self
	}
	self.err = wrapError(self.name, "", customError...)
	return self
}
