package valueobject

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func containMutibyte(value string) bool {
	return len(value) != utf8.RuneCountInString(value)
}

func newRequiredError(field string) error {
	msg := fmt.Sprintf("%sは必須です。", field)
	return errors.New(msg)
}

func newFormatError(field string) error {
	msg := fmt.Sprintf("%sは正しい形式ではありません。", field)
	return errors.New(msg)
}

func newContainMutibyteError(field string) error {
	msg := fmt.Sprintf("%sに利用できない文字が含まれています。", field)
	return errors.New(msg)
}

func newContainItemsError(field string) error {
	msg := fmt.Sprintf("%sに利用できない選択肢が含まれています。", field)
	return errors.New(msg)
}
