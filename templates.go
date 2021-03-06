package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	//"regexp"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"text/template/parse"
)

// TODO: Turn this file into a library
var ctemplates []string
var tmplPtrMap = make(map[string]interface{})
var textOverlapList = make(map[string]int)

// nolint
type VarItem struct {
	Name        string
	Destination string
	Type        string
}
type VarItemReflect struct {
	Name        string
	Destination string
	Value       reflect.Value
}
type CTemplateSet struct {
	tlist          map[string]*parse.Tree
	dir            string
	funcMap        map[string]interface{}
	importMap      map[string]string
	Fragments      map[string]int
	FragmentCursor map[string]int
	FragOut        string
	varList        map[string]VarItem
	localVars      map[string]map[string]VarItemReflect
	stats          map[string]int
	pVarList       string
	pVarPosition   int
	previousNode   parse.NodeType
	currentNode    parse.NodeType
	nextNode       parse.NodeType
	//tempVars map[string]string
	doImports  bool
	expectsInt interface{}
}

func (c *CTemplateSet) compileTemplate(name string, dir string, expects string, expectsInt interface{}, varList map[string]VarItem) (out string, err error) {
	if dev.DebugMode {
		fmt.Println("Compiling template '" + name + "'")
	}

	c.dir = dir
	c.doImports = true
	c.funcMap = map[string]interface{}{
		"and":      "&&",
		"not":      "!",
		"or":       "||",
		"eq":       true,
		"ge":       true,
		"gt":       true,
		"le":       true,
		"lt":       true,
		"ne":       true,
		"add":      true,
		"subtract": true,
		"multiply": true,
		"divide":   true,
	}

	c.importMap = map[string]string{
		"net/http": "net/http",
	}
	c.varList = varList
	//c.pVarList = ""
	//c.pVarPosition = 0
	c.stats = make(map[string]int)
	c.expectsInt = expectsInt
	holdreflect := reflect.ValueOf(expectsInt)

	res, err := ioutil.ReadFile(dir + name)
	if err != nil {
		return "", err
	}

	content := string(res)
	if config.MinifyTemplates {
		content = minify(content)
	}

	tree := parse.New(name, c.funcMap)
	var treeSet = make(map[string]*parse.Tree)
	tree, err = tree.Parse(content, "{{", "}}", treeSet, c.funcMap)
	if err != nil {
		return "", err
	}
	if dev.TemplateDebug {
		fmt.Println(name)
	}

	out = ""
	fname := strings.TrimSuffix(name, filepath.Ext(name))
	c.tlist = make(map[string]*parse.Tree)
	c.tlist[fname] = tree
	varholder := "tmpl_" + fname + "_vars"

	if dev.TemplateDebug {
		fmt.Println(c.tlist)
	}
	c.localVars = make(map[string]map[string]VarItemReflect)
	c.localVars[fname] = make(map[string]VarItemReflect)
	c.localVars[fname]["."] = VarItemReflect{".", varholder, holdreflect}
	if c.Fragments == nil {
		c.Fragments = make(map[string]int)
	}
	c.FragmentCursor = make(map[string]int)
	c.FragmentCursor[fname] = 0

	subtree := c.tlist[fname]
	if dev.TemplateDebug {
		fmt.Println(subtree.Root)
	}

	treeLength := len(subtree.Root.Nodes)
	for index, node := range subtree.Root.Nodes {
		if dev.TemplateDebug {
			fmt.Println("Node: " + node.String())
		}

		c.previousNode = c.currentNode
		c.currentNode = node.Type()
		if treeLength != (index + 1) {
			c.nextNode = subtree.Root.Nodes[index+1].Type()
		}
		out += c.compileSwitch(varholder, holdreflect, fname, node)
	}

	var importList string
	if c.doImports {
		for _, item := range c.importMap {
			importList += "import \"" + item + "\"\n"
		}
	}

	var varString string
	for _, varItem := range c.varList {
		varString += "var " + varItem.Name + " " + varItem.Type + " = " + varItem.Destination + "\n"
	}

	fout := "// +build !no_templategen\n\n// Code generated by Gosora. More below:\n/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */\n"
	fout += "package main\n" + importList + c.pVarList + "\n"
	fout += "// nolint\nfunc init() {\n\ttemplate_" + fname + "_handle = template_" + fname + "\n\t//o_template_" + fname + "_handle = template_" + fname + "\n\tctemplates = append(ctemplates,\"" + fname + "\")\n\ttmplPtrMap[\"" + fname + "\"] = &template_" + fname + "_handle\n\ttmplPtrMap[\"o_" + fname + "\"] = template_" + fname + "\n}\n\n"
	fout += "// nolint\nfunc template_" + fname + "(tmpl_" + fname + "_vars " + expects + ", w http.ResponseWriter) {\n" + varString + out + "}\n"

	fout = strings.Replace(fout, `))
w.Write([]byte(`, " + ", -1)
	fout = strings.Replace(fout, "` + `", "", -1)
	//spstr := "`([:space:]*)`"
	//whitespace_writes := regexp.MustCompile(`(?s)w.Write\(\[\]byte\(`+spstr+`\)\)`)
	//fout = whitespace_writes.ReplaceAllString(fout,"")

	if dev.DebugMode {
		for index, count := range c.stats {
			fmt.Println(index + ": " + strconv.Itoa(count))
		}
		fmt.Println(" ")
	}

	if dev.TemplateDebug {
		fmt.Println("Output!")
		fmt.Println(fout)
	}
	return fout, nil
}

func (c *CTemplateSet) compileSwitch(varholder string, holdreflect reflect.Value, templateName string, node interface{}) (out string) {
	if dev.TemplateDebug {
		fmt.Println("in compileSwitch")
	}
	switch node := node.(type) {
	case *parse.ActionNode:
		if dev.TemplateDebug {
			fmt.Println("Action Node")
		}
		if node.Pipe == nil {
			break
		}
		for _, cmd := range node.Pipe.Cmds {
			out += c.compileSubswitch(varholder, holdreflect, templateName, cmd)
		}
		return out
	case *parse.IfNode:
		if dev.TemplateDebug {
			fmt.Println("If Node:")
			fmt.Println("node.Pipe", node.Pipe)
		}

		var expr string
		for _, cmd := range node.Pipe.Cmds {
			if dev.TemplateDebug {
				fmt.Println("If Node Bit:", cmd)
				fmt.Println("If Node Bit Type:", reflect.ValueOf(cmd).Type().Name())
			}
			expr += c.compileVarswitch(varholder, holdreflect, templateName, cmd)
			if dev.TemplateDebug {
				fmt.Println("If Node Expression Step:", c.compileVarswitch(varholder, holdreflect, templateName, cmd))
			}
		}

		if dev.TemplateDebug {
			fmt.Println("If Node Expression:", expr)
		}

		c.previousNode = c.currentNode
		c.currentNode = parse.NodeList
		c.nextNode = -1
		if node.ElseList == nil {
			if dev.TemplateDebug {
				fmt.Println("Selected Branch 1")
			}
			return "if " + expr + " {\n" + c.compileSwitch(varholder, holdreflect, templateName, node.List) + "}\n"
		}

		if dev.TemplateDebug {
			fmt.Println("Selected Branch 2")
		}
		return "if " + expr + " {\n" + c.compileSwitch(varholder, holdreflect, templateName, node.List) + "} else {\n" + c.compileSwitch(varholder, holdreflect, templateName, node.ElseList) + "}\n"
	case *parse.ListNode:
		if dev.TemplateDebug {
			fmt.Println("List Node")
		}
		for _, subnode := range node.Nodes {
			out += c.compileSwitch(varholder, holdreflect, templateName, subnode)
		}
		return out
	case *parse.RangeNode:
		if dev.TemplateDebug {
			fmt.Println("Range Node!")
			fmt.Println(node.Pipe)
		}

		var outVal reflect.Value
		for _, cmd := range node.Pipe.Cmds {
			if dev.TemplateDebug {
				fmt.Println("Range Bit:", cmd)
			}
			out, outVal = c.compileReflectswitch(varholder, holdreflect, templateName, cmd)
		}

		if dev.TemplateDebug {
			fmt.Println("Returned:", out)
			fmt.Println("Range Kind Switch!")
		}

		switch outVal.Kind() {
		case reflect.Map:
			var item reflect.Value
			for _, key := range outVal.MapKeys() {
				item = outVal.MapIndex(key)
			}
			if dev.DebugMode {
				fmt.Println("Range item:", item)
			}
			if !item.IsValid() {
				panic("item" + "^\n" + "Invalid map. Maybe, it doesn't have any entries for the template engine to analyse?")
			}

			if node.ElseList != nil {
				out = "if len(" + out + ") != 0 {\nfor _, item := range " + out + " {\n" + c.compileSwitch("item", item, templateName, node.List) + "}\n} else {\n" + c.compileSwitch("item", item, templateName, node.ElseList) + "}\n"
			} else {
				out = "if len(" + out + ") != 0 {\nfor _, item := range " + out + " {\n" + c.compileSwitch("item", item, templateName, node.List) + "}\n}"
			}
		case reflect.Slice:
			if outVal.Len() == 0 {
				panic("The sample data needs at-least one or more elements for the slices. We're looking into removing this requirement at some point!")
			}
			item := outVal.Index(0)
			out = "if len(" + out + ") != 0 {\nfor _, item := range " + out + " {\n" + c.compileSwitch("item", item, templateName, node.List) + "}\n}"
		case reflect.Invalid:
			return ""
		}

		if node.ElseList != nil {
			out += " else {\n" + c.compileSwitch(varholder, holdreflect, templateName, node.ElseList) + "}\n"
		} else {
			out += "\n"
		}
		return out
	case *parse.TemplateNode:
		return c.compileSubtemplate(varholder, holdreflect, node)
	case *parse.TextNode:
		c.previousNode = c.currentNode
		c.currentNode = node.Type()
		c.nextNode = 0
		tmpText := bytes.TrimSpace(node.Text)
		if len(tmpText) == 0 {
			return ""
		}

		//return "w.Write([]byte(`" + string(node.Text) + "`))\n"
		fragmentName := templateName + "_" + strconv.Itoa(c.FragmentCursor[templateName])
		_, ok := c.Fragments[fragmentName]
		if !ok {
			c.Fragments[fragmentName] = len(node.Text)
			c.FragOut += "var " + fragmentName + " = []byte(`" + string(node.Text) + "`)\n"
		}
		c.FragmentCursor[templateName] = c.FragmentCursor[templateName] + 1
		return "w.Write(" + fragmentName + ")\n"
	default:
		panic("Unknown Node in main switch")
	}
	return ""
}

func (c *CTemplateSet) compileSubswitch(varholder string, holdreflect reflect.Value, templateName string, node *parse.CommandNode) (out string) {
	if dev.TemplateDebug {
		fmt.Println("in compileSubswitch")
	}
	firstWord := node.Args[0]
	switch n := firstWord.(type) {
	case *parse.FieldNode:
		if dev.TemplateDebug {
			fmt.Println("Field Node:", n.Ident)
		}

		/* Use reflect to determine if the field is for a method, otherwise assume it's a variable. Variable declarations are coming soon! */
		cur := holdreflect

		var varbit string
		if cur.Kind() == reflect.Interface {
			cur = cur.Elem()
			varbit += ".(" + cur.Type().Name() + ")"
		}

		for _, id := range n.Ident {
			if dev.TemplateDebug {
				fmt.Println("Data Kind:", cur.Kind().String())
				fmt.Println("Field Bit:", id)
			}

			if cur.Kind() == reflect.Ptr {
				if dev.TemplateDebug {
					fmt.Println("Looping over pointer")
				}
				for cur.Kind() == reflect.Ptr {
					cur = cur.Elem()
				}

				if dev.TemplateDebug {
					fmt.Println("Data Kind:", cur.Kind().String())
					fmt.Println("Field Bit:", id)
				}
			}

			if !cur.IsValid() {
				if dev.DebugMode {
					fmt.Println("Debug Data:")
					fmt.Println("Holdreflect:", holdreflect)
					fmt.Println("Holdreflect.Kind()", holdreflect.Kind())
					if !dev.TemplateDebug {
						fmt.Println("cur.Kind():", cur.Kind().String())
					}
					fmt.Println("")
				}

				panic(varholder + varbit + "^\n" + "Invalid value. Maybe, it doesn't exist?")
			}

			cur = cur.FieldByName(id)
			if cur.Kind() == reflect.Interface {
				cur = cur.Elem()
				// TODO: Surely, there's a better way of detecting this?
				/*if cur.Kind() == reflect.String && cur.Type().Name() != "string" {
				varbit = "string(" + varbit + "." + id + ")"*/
				//if cur.Kind() == reflect.String && cur.Type().Name() != "string" {
				if cur.Type().PkgPath() != "main" && cur.Type().PkgPath() != "" {
					c.importMap["html/template"] = "html/template"
					varbit += "." + id + ".(" + strings.TrimPrefix(cur.Type().PkgPath(), "html/") + "." + cur.Type().Name() + ")"
				} else {
					varbit += "." + id + ".(" + cur.Type().Name() + ")"
				}
			} else {
				varbit += "." + id
			}
			if dev.TemplateDebug {
				fmt.Println("End Cycle")
			}
		}
		out = c.compileVarsub(varholder+varbit, cur)

		for _, varItem := range c.varList {
			if strings.HasPrefix(out, varItem.Destination) {
				out = strings.Replace(out, varItem.Destination, varItem.Name, 1)
			}
		}
		return out
	case *parse.DotNode:
		if dev.TemplateDebug {
			fmt.Println("Dot Node:", node.String())
		}
		return c.compileVarsub(varholder, holdreflect)
	case *parse.NilNode:
		panic("Nil is not a command x.x")
	case *parse.VariableNode:
		if dev.TemplateDebug {
			fmt.Println("Variable Node:", n.String())
			fmt.Println(n.Ident)
		}
		varname, reflectVal := c.compileIfVarsub(n.String(), varholder, templateName, holdreflect)
		return c.compileVarsub(varname, reflectVal)
	case *parse.StringNode:
		return n.Quoted
	case *parse.IdentifierNode:
		if dev.TemplateDebug {
			fmt.Println("Identifier Node:", node)
			fmt.Println("Identifier Node Args:", node.Args)
		}
		return c.compileVarsub(c.compileIdentswitch(varholder, holdreflect, templateName, node))
	default:
		fmt.Println("Unknown Kind:", reflect.ValueOf(firstWord).Elem().Kind())
		fmt.Println("Unknown Type:", reflect.ValueOf(firstWord).Elem().Type().Name())
		panic("I don't know what node this is")
	}
}

func (c *CTemplateSet) compileVarswitch(varholder string, holdreflect reflect.Value, templateName string, node *parse.CommandNode) (out string) {
	if dev.TemplateDebug {
		fmt.Println("in compile_varswitch")
	}
	firstWord := node.Args[0]
	switch n := firstWord.(type) {
	case *parse.FieldNode:
		if dev.TemplateDebug {
			fmt.Println("Field Node:", n.Ident)
			for _, id := range n.Ident {
				fmt.Println("Field Bit:", id)
			}
		}

		/* Use reflect to determine if the field is for a method, otherwise assume it's a variable. Coming Soon. */
		return c.compileBoolsub(n.String(), varholder, templateName, holdreflect)
	case *parse.ChainNode:
		if dev.TemplateDebug {
			fmt.Println("Chain Node:", n.Node)
			fmt.Println("Chain Node Args:", node.Args)
		}
		break
	case *parse.IdentifierNode:
		if dev.TemplateDebug {
			fmt.Println("Identifier Node:", node)
			fmt.Println("Identifier Node Args:", node.Args)
		}
		return c.compileIdentswitchN(varholder, holdreflect, templateName, node)
	case *parse.DotNode:
		return varholder
	case *parse.VariableNode:
		if dev.TemplateDebug {
			fmt.Println("Variable Node:", n.String())
			fmt.Println("Variable Node Identifier:", n.Ident)
		}
		out, _ = c.compileIfVarsub(n.String(), varholder, templateName, holdreflect)
		return out
	case *parse.NilNode:
		panic("Nil is not a command x.x")
	case *parse.PipeNode:
		if dev.TemplateDebug {
			fmt.Println("Pipe Node!")
			fmt.Println(n)
			fmt.Println("Args:", node.Args)
		}
		out += c.compileIdentswitchN(varholder, holdreflect, templateName, node)

		if dev.TemplateDebug {
			fmt.Println("Out:", out)
		}
		return out
	default:
		fmt.Println("Unknown Kind:", reflect.ValueOf(firstWord).Elem().Kind())
		fmt.Println("Unknown Type:", reflect.ValueOf(firstWord).Elem().Type().Name())
		panic("I don't know what node this is! Grr...")
	}
	return ""
}

func (c *CTemplateSet) compileIdentswitchN(varholder string, holdreflect reflect.Value, templateName string, node *parse.CommandNode) (out string) {
	if dev.TemplateDebug {
		fmt.Println("in compile_identswitch_n")
	}
	out, _ = c.compileIdentswitch(varholder, holdreflect, templateName, node)
	return out
}

func (c *CTemplateSet) compileIdentswitch(varholder string, holdreflect reflect.Value, templateName string, node *parse.CommandNode) (out string, val reflect.Value) {
	if dev.TemplateDebug {
		fmt.Println("in compileIdentswitch")
	}

	//var outbuf map[int]string
ArgLoop:
	for pos := 0; pos < len(node.Args); pos++ {
		id := node.Args[pos]
		if dev.TemplateDebug {
			fmt.Println("pos:", pos)
			fmt.Println("ID:", id)
		}
		switch id.String() {
		case "not":
			out += "!"
		case "or":
			if dev.TemplateDebug {
				fmt.Println("Building or function")
			}
			if pos == 0 {
				fmt.Println("pos:", pos)
				panic("or is missing a left operand")
			}
			if len(node.Args) <= pos {
				fmt.Println("post pos:", pos)
				fmt.Println("len(node.Args):", len(node.Args))
				panic("or is missing a right operand")
			}

			left := c.compileBoolsub(node.Args[pos-1].String(), varholder, templateName, holdreflect)
			_, funcExists := c.funcMap[node.Args[pos+1].String()]

			var right string
			if !funcExists {
				right = c.compileBoolsub(node.Args[pos+1].String(), varholder, templateName, holdreflect)
			}

			out += left + " || " + right

			if dev.TemplateDebug {
				fmt.Println("Left operand:", node.Args[pos-1])
				fmt.Println("Right operand:", node.Args[pos+1])
			}

			if !funcExists {
				pos++
			}

			if dev.TemplateDebug {
				fmt.Println("pos:", pos)
				fmt.Println("len(node.Args):", len(node.Args))
			}
		case "and":
			if dev.TemplateDebug {
				fmt.Println("Building and function")
			}
			if pos == 0 {
				fmt.Println("pos:", pos)
				panic("and is missing a left operand")
			}
			if len(node.Args) <= pos {
				fmt.Println("post pos:", pos)
				fmt.Println("len(node.Args):", len(node.Args))
				panic("and is missing a right operand")
			}

			left := c.compileBoolsub(node.Args[pos-1].String(), varholder, templateName, holdreflect)
			_, funcExists := c.funcMap[node.Args[pos+1].String()]

			var right string
			if !funcExists {
				right = c.compileBoolsub(node.Args[pos+1].String(), varholder, templateName, holdreflect)
			}

			out += left + " && " + right

			if dev.TemplateDebug {
				fmt.Println("Left operand:", node.Args[pos-1])
				fmt.Println("Right operand:", node.Args[pos+1])
			}

			if !funcExists {
				pos++
			}

			if dev.TemplateDebug {
				fmt.Println("pos:", pos)
				fmt.Println("len(node.Args):", len(node.Args))
			}
		case "le":
			out += c.compileIfVarsubN(node.Args[pos+1].String(), varholder, templateName, holdreflect) + " <= " + c.compileIfVarsubN(node.Args[pos+2].String(), varholder, templateName, holdreflect)
			if dev.TemplateDebug {
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "lt":
			out += c.compileIfVarsubN(node.Args[pos+1].String(), varholder, templateName, holdreflect) + " < " + c.compileIfVarsubN(node.Args[pos+2].String(), varholder, templateName, holdreflect)
			if dev.TemplateDebug {
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "gt":
			out += c.compileIfVarsubN(node.Args[pos+1].String(), varholder, templateName, holdreflect) + " > " + c.compileIfVarsubN(node.Args[pos+2].String(), varholder, templateName, holdreflect)
			if dev.TemplateDebug {
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "ge":
			out += c.compileIfVarsubN(node.Args[pos+1].String(), varholder, templateName, holdreflect) + " >= " + c.compileIfVarsubN(node.Args[pos+2].String(), varholder, templateName, holdreflect)
			if dev.TemplateDebug {
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "eq":
			out += c.compileIfVarsubN(node.Args[pos+1].String(), varholder, templateName, holdreflect) + " == " + c.compileIfVarsubN(node.Args[pos+2].String(), varholder, templateName, holdreflect)
			if dev.TemplateDebug {
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "ne":
			out += c.compileIfVarsubN(node.Args[pos+1].String(), varholder, templateName, holdreflect) + " != " + c.compileIfVarsubN(node.Args[pos+2].String(), varholder, templateName, holdreflect)
			if dev.TemplateDebug {
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "add":
			param1, val2 := c.compileIfVarsub(node.Args[pos+1].String(), varholder, templateName, holdreflect)
			param2, val3 := c.compileIfVarsub(node.Args[pos+2].String(), varholder, templateName, holdreflect)

			if val2.IsValid() {
				val = val2
			} else if val3.IsValid() {
				val = val3
			} else {
				numSample := 1
				val = reflect.ValueOf(numSample)
			}

			out += param1 + " + " + param2
			if dev.TemplateDebug {
				fmt.Println("add")
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "subtract":
			param1, val2 := c.compileIfVarsub(node.Args[pos+1].String(), varholder, templateName, holdreflect)
			param2, val3 := c.compileIfVarsub(node.Args[pos+2].String(), varholder, templateName, holdreflect)

			if val2.IsValid() {
				val = val2
			} else if val3.IsValid() {
				val = val3
			} else {
				numSample := 1
				val = reflect.ValueOf(numSample)
			}

			out += param1 + " - " + param2
			if dev.TemplateDebug {
				fmt.Println("subtract")
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "divide":
			param1, val2 := c.compileIfVarsub(node.Args[pos+1].String(), varholder, templateName, holdreflect)
			param2, val3 := c.compileIfVarsub(node.Args[pos+2].String(), varholder, templateName, holdreflect)

			if val2.IsValid() {
				val = val2
			} else if val3.IsValid() {
				val = val3
			} else {
				numSample := 1
				val = reflect.ValueOf(numSample)
			}

			out += param1 + " / " + param2
			if dev.TemplateDebug {
				fmt.Println("divide")
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		case "multiply":
			param1, val2 := c.compileIfVarsub(node.Args[pos+1].String(), varholder, templateName, holdreflect)
			param2, val3 := c.compileIfVarsub(node.Args[pos+2].String(), varholder, templateName, holdreflect)

			if val2.IsValid() {
				val = val2
			} else if val3.IsValid() {
				val = val3
			} else {
				numSample := 1
				val = reflect.ValueOf(numSample)
			}

			out += param1 + " * " + param2
			if dev.TemplateDebug {
				fmt.Println("multiply")
				fmt.Println("node.Args[pos + 1]", node.Args[pos+1])
				fmt.Println("node.Args[pos + 2]", node.Args[pos+2])
			}
			break ArgLoop
		default:
			if dev.TemplateDebug {
				fmt.Println("Variable!")
			}
			if len(node.Args) > (pos + 1) {
				nextNode := node.Args[pos+1].String()
				if nextNode == "or" || nextNode == "and" {
					continue
				}
			}
			out += c.compileIfVarsubN(id.String(), varholder, templateName, holdreflect)
		}
	}

	//for _, outval := range outbuf {
	//	out += outval
	//}
	return out, val
}

func (c *CTemplateSet) compileReflectswitch(varholder string, holdreflect reflect.Value, templateName string, node *parse.CommandNode) (out string, outVal reflect.Value) {
	if dev.TemplateDebug {
		fmt.Println("in compileReflectswitch")
	}
	firstWord := node.Args[0]
	switch n := firstWord.(type) {
	case *parse.FieldNode:
		if dev.TemplateDebug {
			fmt.Println("Field Node:", n.Ident)
			for _, id := range n.Ident {
				fmt.Println("Field Bit:", id)
			}
		}
		/* Use reflect to determine if the field is for a method, otherwise assume it's a variable. Coming Soon. */
		return c.compileIfVarsub(n.String(), varholder, templateName, holdreflect)
	case *parse.ChainNode:
		if dev.TemplateDebug {
			fmt.Println("Chain Node:", n.Node)
			fmt.Println("node.Args", node.Args)
		}
		return "", outVal
	case *parse.DotNode:
		return varholder, holdreflect
	case *parse.NilNode:
		panic("Nil is not a command x.x")
	default:
		//panic("I don't know what node this is")
	}
	return "", outVal
}

func (c *CTemplateSet) compileIfVarsubN(varname string, varholder string, templateName string, cur reflect.Value) (out string) {
	if dev.TemplateDebug {
		fmt.Println("in compileIfVarsubN")
	}
	out, _ = c.compileIfVarsub(varname, varholder, templateName, cur)
	return out
}

func (c *CTemplateSet) compileIfVarsub(varname string, varholder string, templateName string, cur reflect.Value) (out string, val reflect.Value) {
	if dev.TemplateDebug {
		fmt.Println("in compileIfVarsub")
	}
	if varname[0] != '.' && varname[0] != '$' {
		return varname, cur
	}

	bits := strings.Split(varname, ".")
	if varname[0] == '$' {
		var res VarItemReflect
		if varname[1] == '.' {
			res = c.localVars[templateName]["."]
		} else {
			res = c.localVars[templateName][strings.TrimPrefix(bits[0], "$")]
		}
		out += res.Destination
		cur = res.Value

		if cur.Kind() == reflect.Interface {
			cur = cur.Elem()
		}
	} else {
		if cur.Kind() == reflect.Interface {
			cur = cur.Elem()
			out += varholder + ".(" + cur.Type().Name() + ")"
		} else {
			out += varholder
		}
	}
	bits[0] = strings.TrimPrefix(bits[0], "$")

	if dev.TemplateDebug {
		fmt.Println("Cur Kind:", cur.Kind())
		fmt.Println("Cur Type:", cur.Type().Name())
	}

	for _, bit := range bits {
		if dev.TemplateDebug {
			fmt.Println("Variable Field:", bit)
		}
		if bit == "" {
			continue
		}

		// TODO: Fix this up so that it works for regular pointers and not just struct pointers. Ditto for the other cur.Kind() == reflect.Ptr we have in this file
		if cur.Kind() == reflect.Ptr {
			if dev.TemplateDebug {
				fmt.Println("Looping over pointer")
			}
			for cur.Kind() == reflect.Ptr {
				cur = cur.Elem()
			}

			if dev.TemplateDebug {
				fmt.Println("Data Kind:", cur.Kind().String())
				fmt.Println("Field Bit:", bit)
			}
		}

		cur = cur.FieldByName(bit)
		if cur.Kind() == reflect.Interface {
			cur = cur.Elem()
			out += "." + bit + ".(" + cur.Type().Name() + ")"
		} else {
			out += "." + bit
		}

		if !cur.IsValid() {
			panic(out + "^\n" + "Invalid value. Maybe, it doesn't exist?")
		}

		if dev.TemplateDebug {
			fmt.Println("Data Kind:", cur.Kind())
			fmt.Println("Data Type:", cur.Type().Name())
		}
	}

	if dev.TemplateDebug {
		fmt.Println("Out Value:", out)
		fmt.Println("Out Kind:", cur.Kind())
		fmt.Println("Out Type:", cur.Type().Name())
	}

	for _, varItem := range c.varList {
		if strings.HasPrefix(out, varItem.Destination) {
			out = strings.Replace(out, varItem.Destination, varItem.Name, 1)
		}
	}

	if dev.TemplateDebug {
		fmt.Println("Out Value:", out)
		fmt.Println("Out Kind:", cur.Kind())
		fmt.Println("Out Type:", cur.Type().Name())
	}

	_, ok := c.stats[out]
	if ok {
		c.stats[out]++
	} else {
		c.stats[out] = 1
	}

	return out, cur
}

func (c *CTemplateSet) compileBoolsub(varname string, varholder string, templateName string, val reflect.Value) string {
	if dev.TemplateDebug {
		fmt.Println("in compileBoolsub")
	}
	out, val := c.compileIfVarsub(varname, varholder, templateName, val)
	// TODO: What if it's a pointer or an interface? I *think* we've got pointers handled somewhere, but not interfaces which we don't know the types of at compile time
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		out += " > 0"
	case reflect.Bool: // Do nothing
	case reflect.String:
		out += " != \"\""
	case reflect.Slice, reflect.Map:
		out = "len(" + out + ") != 0"
	default:
		fmt.Println("Variable Name:", varname)
		fmt.Println("Variable Holder:", varholder)
		fmt.Println("Variable Kind:", val.Kind())
		panic("I don't know what this variable's type is o.o\n")
	}
	return out
}

func (c *CTemplateSet) compileVarsub(varname string, val reflect.Value) string {
	if dev.TemplateDebug {
		fmt.Println("in compileVarsub")
	}
	for _, varItem := range c.varList {
		if strings.HasPrefix(varname, varItem.Destination) {
			varname = strings.Replace(varname, varItem.Destination, varItem.Name, 1)
		}
	}

	_, ok := c.stats[varname]
	if ok {
		c.stats[varname]++
	} else {
		c.stats[varname] = 1
	}

	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Int:
		c.importMap["strconv"] = "strconv"
		return "w.Write([]byte(strconv.Itoa(" + varname + ")))\n"
	case reflect.Bool:
		return "if " + varname + " {\nw.Write([]byte(\"true\"))} else {\nw.Write([]byte(\"false\"))\n}\n"
	case reflect.String:
		if val.Type().Name() != "string" && !strings.HasPrefix(varname, "string(") {
			return "w.Write([]byte(string(" + varname + ")))\n"
		}
		return "w.Write([]byte(" + varname + "))\n"
	case reflect.Int64:
		c.importMap["strconv"] = "strconv"
		return "w.Write([]byte(strconv.FormatInt(" + varname + ", 10)))"
	default:
		if !val.IsValid() {
			panic(varname + "^\n" + "Invalid value. Maybe, it doesn't exist?")
		}
		fmt.Println("Unknown Variable Name:", varname)
		fmt.Println("Unknown Kind:", val.Kind())
		fmt.Println("Unknown Type:", val.Type().Name())
		panic("// I don't know what this variable's type is o.o\n")
	}
}

func (c *CTemplateSet) compileSubtemplate(pvarholder string, pholdreflect reflect.Value, node *parse.TemplateNode) (out string) {
	if dev.TemplateDebug {
		fmt.Println("in compileSubtemplate")
		fmt.Println("Template Node:", node.Name)
	}

	fname := strings.TrimSuffix(node.Name, filepath.Ext(node.Name))
	varholder := "tmpl_" + fname + "_vars"
	var holdreflect reflect.Value
	if node.Pipe != nil {
		for _, cmd := range node.Pipe.Cmds {
			firstWord := cmd.Args[0]
			switch firstWord.(type) {
			case *parse.DotNode:
				varholder = pvarholder
				holdreflect = pholdreflect
				break
			case *parse.NilNode:
				panic("Nil is not a command x.x")
			default:
				out = "var " + varholder + " := false\n"
				out += c.compileCommand(cmd)
			}
		}
	}

	// TODO: Cascade errors back up the tree to the caller?
	res, err := ioutil.ReadFile(c.dir + node.Name)
	if err != nil {
		log.Fatal(err)
	}

	content := string(res)
	if config.MinifyTemplates {
		content = minify(content)
	}

	tree := parse.New(node.Name, c.funcMap)
	var treeSet = make(map[string]*parse.Tree)
	tree, err = tree.Parse(content, "{{", "}}", treeSet, c.funcMap)
	if err != nil {
		log.Fatal(err)
	}

	c.tlist[fname] = tree
	subtree := c.tlist[fname]
	if dev.TemplateDebug {
		fmt.Println("subtree.Root", subtree.Root)
	}

	c.localVars[fname] = make(map[string]VarItemReflect)
	c.localVars[fname]["."] = VarItemReflect{".", varholder, holdreflect}
	c.FragmentCursor[fname] = 0

	treeLength := len(subtree.Root.Nodes)
	for index, node := range subtree.Root.Nodes {
		if dev.TemplateDebug {
			fmt.Println("Node:", node.String())
		}

		c.previousNode = c.currentNode
		c.currentNode = node.Type()
		if treeLength != (index + 1) {
			c.nextNode = subtree.Root.Nodes[index+1].Type()
		}
		out += c.compileSwitch(varholder, holdreflect, fname, node)
	}
	return out
}

func (c *CTemplateSet) compileCommand(*parse.CommandNode) (out string) {
	panic("Uh oh! Something went wrong!")
}

// TODO: Write unit tests for this
func minify(data string) string {
	data = strings.Replace(data, "\t", "", -1)
	data = strings.Replace(data, "\v", "", -1)
	data = strings.Replace(data, "\n", "", -1)
	data = strings.Replace(data, "\r", "", -1)
	data = strings.Replace(data, "  ", " ", -1)
	return data
}

// TODO: Strip comments
// TODO: Handle CSS nested in <style> tags?
// TODO: Write unit tests for this
func minifyHTML(data string) string {
	return minify(data)
}

// TODO: Have static files use this
// TODO: Strip comments
// TODO: Convert the rgb()s to hex codes?
// TODO: Write unit tests for this
func minifyCSS(data string) string {
	return minify(data)
}

// TODO: Convert this to three character hex strings whenever possible?
// TODO: Write unit tests for this
// nolint
func rgbToHexstr(red int, green int, blue int) string {
	return strconv.FormatInt(int64(red), 16) + strconv.FormatInt(int64(green), 16) + strconv.FormatInt(int64(blue), 16)
}

/*
// TODO: Write unit tests for this
func hexstrToRgb(hexstr string) (red int, blue int, green int, err error) {
	// Strip the # at the start
	if hexstr[0] == '#' {
		hexstr = strings.TrimPrefix(hexstr,"#")
	}
	if len(hexstr) != 3 && len(hexstr) != 6 {
		return 0, 0, 0, errors.New("Hex colour codes may only be three or six characters long")
	}

	if len(hexstr) == 3 {
		hexstr = hexstr[0] + hexstr[0] + hexstr[1] + hexstr[1] + hexstr[2] + hexstr[2]
	}
}*/
