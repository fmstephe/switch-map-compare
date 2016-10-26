package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var size = flag.Int("s", 10, "Size of generated map/switch statement")

func main() {
	flag.Parse()
	in := make([]string, *size)
	out := make([]string, *size)
	for i := 0; i < *size; i++ {
		in[i] = strconv.Itoa(i)
		out[i] = strconv.Itoa(i * i)
	}
	fmt.Fprint(os.Stdout, "package main\n\n")
	fmt.Fprint(os.Stdout, "func main(){}\n\n")
	fmt.Fprint(os.Stdout, generateSlices(in, out))
	fmt.Fprint(os.Stdout, generateStringMap(in, out))
	fmt.Fprint(os.Stdout, generateStringSwitch(in, out))
}

func generateSlices(in, out []string) string {
	buf := &bytes.Buffer{}
	buf.WriteString("var in = []string{")
	for i := range in {
		buf.WriteString("\"")
		buf.WriteString(in[i])
		buf.WriteString("\",")
		if i%10 == 0 {
			buf.WriteString("\n")
		}
	}
	buf.WriteString("}\n")
	buf.WriteString("var out = []string{")
	for i := range in {
		buf.WriteString("\"")
		buf.WriteString(out[i])
		buf.WriteString("\",")
		if i%10 == 0 {
			buf.WriteString("\n")
		}
	}
	buf.WriteString("}\n\n")
	return buf.String()
}

func generateStringMap(in, out []string) string {
	// var stringMap = map[string]string{"foo":"bar",}
	buf := &bytes.Buffer{}
	buf.WriteString("var stringMap = map[string]string{")
	for i := 0; i < len(in); i++ {
		buf.WriteByte('"')
		buf.WriteString(in[i])
		buf.WriteByte('"')
		buf.WriteByte(':')
		buf.WriteByte('"')
		buf.WriteString(out[i])
		buf.WriteByte('"')
		buf.WriteByte(',')
		if i%10 == 0 {
			buf.WriteString("\n")
		}
	}
	buf.WriteString("}\n\n")
	buf.WriteString("func getFromMap(arg string) string {\n")
	buf.WriteString("return stringMap[arg]\n")
	buf.WriteString("}\n\n")
	return buf.String()
}

func generateStringSwitch(in, out []string) string {
	// switch arg {
	// case "foo":
	//	return "bar"
	// }
	buf := &bytes.Buffer{}
	buf.WriteString("func getFromSwitch(arg string) string {\n")
	buf.WriteString("switch arg {\n")
	for i := 0; i < len(in); i++ {
		buf.WriteString("case \"")
		buf.WriteString(in[i])
		buf.WriteString("\":\n")
		buf.WriteString("return \"")
		buf.WriteString(out[i])
		buf.WriteString("\"\n")
	}
	buf.WriteString("}\n")
	buf.WriteString("return \"\"\n")
	buf.WriteString("}\n")
	return buf.String()
}
