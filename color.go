package colorful

import (
	"fmt"
	"reflect"
	"time"
)

func Render(s interface{}) string {
	return pretty(reflect.ValueOf(s), 0)
}

func pretty(v reflect.Value, depth int) string {
	result := ""
	switch f := v; v.Kind() {
	case reflect.Float64:
		r := fmt.Sprintf("%v", f.Float())
		result += green(r)
	case reflect.Float32:
		r := fmt.Sprintf("%v", f.Float())
		result += green(r)
	case reflect.Slice:
		result += "[\n"
		for i := 0; i < v.Len(); i++ {
			result += pretty(v.Index(i), depth+1)
			if i != v.Len()-1 {
				result += ",\n"
			} else {
				result += "\n"
			}
		}
		result += "]"
	case reflect.String:
		result += green(f.String())
	case reflect.Int:
		r := fmt.Sprintf("%d", f.Int())
		result += green(r)
	case reflect.Struct:
		t := f.Type()
		str := ""
		for i := 0; i < depth; i++ {
			str += " "
		}
		result += "{\n"
		for i := 0; i < f.NumField(); i++ {
			value := getField(f, i)
			name := t.Field(i).Name
			result += fmt.Sprintf(str + red(name) + ":" + pretty(value, depth+1))
			if i != f.NumField()-1 {
				result += ",\n"
			} else {
				result += "\n"
			}
		}
		result += str + "}"
	case reflect.Ptr:
		switch a := f.Elem(); a.Kind() {
		case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
			return pretty(a, depth+1)
		}
	}

	return result
}

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta //洋红
	info          = "[INFO]"
	trac          = "[TRAC]"
	erro          = "[ERRO]"
	warn          = "[WARN]"
	succ          = "[SUCC]"
)

// see complete color rules in document in https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13
func Trace(format string, a ...interface{}) {
	prefix := yellow(trac)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}
func Info(format string, a ...interface{}) {
	prefix := blue(info)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}
func Success(format string, a ...interface{}) {
	prefix := green(succ)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}
func Warning(format string, a ...interface{}) {
	prefix := magenta(warn)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}
func Error(format string, a ...interface{}) {
	prefix := red(erro)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}
func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, s)
}
func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}
func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}
func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}
func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}
func formatLog(prefix string) string {
	return time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " "
}
