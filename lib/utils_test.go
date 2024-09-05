package lib

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	u "github.com/sunshine69/golang-tools/utils"
	"gopkg.in/yaml.v3"
)

var project_dir string

func init() {
	project_dir, _ = os.Getwd()
}

func TestAddhoc(t *testing.T) {
	// varAnsible := ParseVarAnsibleNext(project_dir+"/../work/azure-devops/vars-ansible.yaml", project_dir)
	// HelmChartValidation("/home/stevek/src/helm_playground-1_v1/", []string{
	// 	"//home/stevek/src/helm_playground-1_v1/values.yaml",
	// })
	// tmplStr := `{{ "test".upper() }}`
	// o := TemplateString(tmplStr, map[string]interface{}{})
	// fmt.Printf("%s\n", o)
	// str := "Let freedom ring from the mighty mountains of New York. Let freedom ring from the heightening Alleghenies of Pennsylvania. Let freedom ring from the snow-capped Rockies of Colorado. Let freedom ring from the curvaceous slopes of California."
	// counter := 1
	// repl := func(match string) string {
	// 	old := counter
	// 	counter++
	// 	if old != 1 {
	// 		return fmt.Sprintf("[%d] %s%d", old, match, old)
	// 	}
	// 	return fmt.Sprintf("[%d] %s", old, match)
	// }
	// re := regexp.MustCompile("Let freedom")
	// str2 := re.ReplaceAllStringFunc(str, repl)
	// fmt.Println(str2)

	// u.CheckErr(err, "")
	// words, _ := loadDictionary("/tmp/words", 0)

	// words := strings.FieldsFunc(strings.ToLower("#A<V_$AvQ{Yj!Y"), func(r rune) bool {
	// 	return !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'))
	// })
	// fmt.Printf("%v\n", words)

	fmt.Printf("Likely a password: %v\n", IsLikelyPasswordOrToken("VHuGgaJvV", "letter+word", "/tmp/words.txt", 0, 1))
	// fmt.Println(changed)
	ptn := regexp.MustCompile(`(?i)['"]?(password|passwd|token|api_key|secret)['"]?[=:\s][\s]*?['"]?([^'"\s]+)['"]?`)
	matches := ptn.FindAllStringSubmatch(`PublicKeyToken=null</TypeInfo>`, -1)
	fmt.Printf("%q\n", matches)
	fmt.Println("Done test")
}

func TestLineinfile(t *testing.T) {
	err, changed := LineInFile("../tmp/tests", NewLineInfileOpt(&LineInfileOpt{
		// Regexp:     `v1.0.1(.*)`,
		Search_string: "This is new line",
		Line:          "This is new line to be reaplced at line 4",
		// ReplaceAll: true,
	}))
	u.CheckErr(err, "Error")
	fmt.Println(changed)
}

func TestExtractBlock(t *testing.T) {
	o, _, _ := ExtractTextBlock("../tmp/tests", []string{
		"[\\s]+- name: \"Run setup-project-go validatehelm command\"",
	},
		[]string{"[\\s]+- name: \"Install chart using helm local package\""})
	// o = `MYVAR:\n` + o
	fmt.Println(o)
	o1 := []map[string]interface{}{}
	u.CheckErr(yaml.Unmarshal([]byte(o), &o1), "ERR")
	fmt.Printf("%s\n", u.JsonDump(o1, "  "))
}

func TestExtractBlockContains(t *testing.T) {
	o, _, _ := ExtractTextBlockContains("../tmp/tests", []string{`- name: [^\s]+`}, []string{`- name: [^\s]+`}, 70)
	o1 := []map[string]interface{}{}
	u.CheckErr(yaml.Unmarshal([]byte(o), &o1), "ERR")
	fmt.Printf("USE INT %s\n", u.JsonDump(o1, "  "))

	o, _, _ = ExtractTextBlockContains("../tmp/tests", []string{`- name: [^\s]+`}, []string{`- name: [^\s]+`}, []string{`tar --strip-components=1 -xf "{{ helm_chart_resource_fact }}" -C`})
	o1 = []map[string]interface{}{}
	u.CheckErr(yaml.Unmarshal([]byte(o), &o1), "ERR")
	fmt.Printf("USE INT %s\n", u.JsonDump(o1, "  "))
}
