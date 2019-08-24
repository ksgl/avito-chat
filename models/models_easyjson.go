// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeAvitoChatModels(in *jlexer.Lexer, out *UserChat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user":
			out.UserID = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAvitoChatModels(out *jwriter.Writer, in UserChat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserChat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAvitoChatModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserChat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAvitoChatModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserChat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAvitoChatModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserChat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAvitoChatModels(l, v)
}
func easyjsonD2b7633eDecodeAvitoChatModels1(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "username":
			out.Username = string(in.String())
		case "user_id":
			out.UserID = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAvitoChatModels1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.Int32(int32(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAvitoChatModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAvitoChatModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAvitoChatModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAvitoChatModels1(l, v)
}
func easyjsonD2b7633eDecodeAvitoChatModels2(in *jlexer.Lexer, out *MessagesArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(MessagesArr, 0, 8)
			} else {
				*out = MessagesArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Message
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Message)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAvitoChatModels2(out *jwriter.Writer, in MessagesArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v MessagesArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAvitoChatModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagesArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAvitoChatModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagesArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAvitoChatModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagesArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAvitoChatModels2(l, v)
}
func easyjsonD2b7633eDecodeAvitoChatModels3(in *jlexer.Lexer, out *Message) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message_id":
			out.MessageID = int32(in.Int32())
		case "chat":
			out.Chat = int32(in.Int32())
		case "author":
			out.Author = int32(in.Int32())
		case "text":
			out.Text = string(in.String())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAvitoChatModels3(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message_id\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.MessageID))
	}
	{
		const prefix string = ",\"chat\":"
		out.RawString(prefix)
		out.Int32(int32(in.Chat))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.Int32(int32(in.Author))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAvitoChatModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAvitoChatModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAvitoChatModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAvitoChatModels3(l, v)
}
func easyjsonD2b7633eDecodeAvitoChatModels4(in *jlexer.Lexer, out *ChatsArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ChatsArr, 0, 8)
			} else {
				*out = ChatsArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 *Chat
			if in.IsNull() {
				in.Skip()
				v4 = nil
			} else {
				if v4 == nil {
					v4 = new(Chat)
				}
				(*v4).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAvitoChatModels4(out *jwriter.Writer, in ChatsArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ChatsArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAvitoChatModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatsArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAvitoChatModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatsArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAvitoChatModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatsArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAvitoChatModels4(l, v)
}
func easyjsonD2b7633eDecodeAvitoChatModels5(in *jlexer.Lexer, out *ChatMessages) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "chat":
			out.ChatID = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAvitoChatModels5(out *jwriter.Writer, in ChatMessages) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"chat\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.ChatID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatMessages) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAvitoChatModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatMessages) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAvitoChatModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatMessages) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAvitoChatModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatMessages) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAvitoChatModels5(l, v)
}
func easyjsonD2b7633eDecodeAvitoChatModels6(in *jlexer.Lexer, out *Chat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "chat_id":
			out.ChatID = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]int32, 0, 16)
					} else {
						out.Users = []int32{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v7 int32
					v7 = int32(in.Int32())
					out.Users = append(out.Users, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAvitoChatModels6(out *jwriter.Writer, in Chat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"chat_id\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.ChatID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if len(in.Users) != 0 {
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v8, v9 := range in.Users {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Chat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAvitoChatModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Chat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAvitoChatModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Chat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAvitoChatModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Chat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAvitoChatModels6(l, v)
}
