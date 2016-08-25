// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package dockerinspect

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var _ = json.RawMessage{} // suppress unused package warning

func easyjson_f525986a_decode_github_com_gooops_micadvisor_open_dockerinspect_Inspect(in *jlexer.Lexer, out *Inspect) {
	if in.IsNull() {
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
		case "Id":
			out.Id = string(in.String())
		case "Created":
			out.Created = string(in.String())
		case "Config":
			if in.IsNull() {
				in.Skip()
				out.Config = nil
			} else {
				out.Config = new(InspectConfig)
				(*out.Config).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_f525986a_encode_github_com_gooops_micadvisor_open_dockerinspect_Inspect(out *jwriter.Writer, in Inspect) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Created\":")
	out.String(string(in.Created))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Config\":")
	if in.Config == nil {
		out.RawString("null")
	} else {
		(*in.Config).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func (v Inspect) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_f525986a_encode_github_com_gooops_micadvisor_open_dockerinspect_Inspect(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v Inspect) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_f525986a_encode_github_com_gooops_micadvisor_open_dockerinspect_Inspect(w, v)
}
func (v *Inspect) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_f525986a_decode_github_com_gooops_micadvisor_open_dockerinspect_Inspect(&r, v)
	return r.Error()
}
func (v *Inspect) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_f525986a_decode_github_com_gooops_micadvisor_open_dockerinspect_Inspect(l, v)
}
func easyjson_f525986a_decode_github_com_gooops_micadvisor_open_dockerinspect_InspectConfig(in *jlexer.Lexer, out *InspectConfig) {
	if in.IsNull() {
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
		case "Hostname":
			out.Hostname = string(in.String())
		case "Labels":
			out.Labels = in.Interface()
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_f525986a_encode_github_com_gooops_micadvisor_open_dockerinspect_InspectConfig(out *jwriter.Writer, in InspectConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Hostname\":")
	out.String(string(in.Hostname))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Labels\":")
	out.Raw(json.Marshal(in.Labels))
	out.RawByte('}')
}
func (v InspectConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_f525986a_encode_github_com_gooops_micadvisor_open_dockerinspect_InspectConfig(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v InspectConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_f525986a_encode_github_com_gooops_micadvisor_open_dockerinspect_InspectConfig(w, v)
}
func (v *InspectConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_f525986a_decode_github_com_gooops_micadvisor_open_dockerinspect_InspectConfig(&r, v)
	return r.Error()
}
func (v *InspectConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_f525986a_decode_github_com_gooops_micadvisor_open_dockerinspect_InspectConfig(l, v)
}
