package strategy_pattern

import (
	"errors"
	"fmt"
	"testing"
)

type SendMessage interface {
	Send(from, to, content string)
}

type SendImgMsg struct{}

func (s SendImgMsg) Send(from, to, content string) {
	fmt.Printf("Send image from [%v] to [%v]: content [%v]\n", from, to, content)
}

type SendVideoMsg struct{}

func (s SendVideoMsg) Send(from, to, content string) {
	fmt.Printf("Send video from [%v] to [%v]: content [%v]\n", from, to, content)
}

type SendTxtMsg struct{}

func (s SendTxtMsg) Send(from, to, content string) {
	fmt.Printf("Send text from [%v] to [%v]: content [%v]\n", from, to, content)
}

type MessageParams struct {
	Type    string
	Content string
	From    string
	To      string
}

var MsgTemplate = map[string]SendMessage{
	"image": new(SendImgMsg),
	"video": new(SendVideoMsg),
	"text":  new(SendTxtMsg),
}

func SendStrategy(params MessageParams) error {
	if _, ok := MsgTemplate[params.Type]; !ok {
		return errors.New("type invalid")
	}
	MsgTemplate[params.Type].Send(params.From, params.To, params.Content)
	return nil
}

func TestStrategy(t *testing.T) {
	imgCtx := MessageParams{
		Type:    "image",
		Content: "发送图片",
		From:    "1",
		To:      "2",
	}
	_ = SendStrategy(imgCtx)
	videoCtx := MessageParams{
		Type:    "video",
		Content: "发送视频",
		From:    "1",
		To:      "2",
	}
	_ = SendStrategy(videoCtx)
	txtCtx := MessageParams{
		Type:    "text",
		Content: "发送文本",
		From:    "1",
		To:      "2",
	}
	_ = SendStrategy(txtCtx)
}
