package trend

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func findEmoji(session *discordgo.Session, emojiString string, guildId string, isAnimated bool) string {
	allEmojis, _ := session.GuildEmojis(guildId)
	for _, emoji := range allEmojis {
		if emoji.Name == emojiString {
			if isAnimated {
				return fmt.Sprintf("<a:%s:%s>", emojiString, emoji.ID)
			} else {
				return fmt.Sprintf("<:%s:%s", emojiString, emoji.ID)
			}
		}
	}
	return ""
}

var numberEmojiStrings = [10]string{"", "", "", "four", "five", "six", "seven", "eight", "nine", "keycap_ten"}
var dateRangeStarGainedMap = map[string]string{"daily": "Day", "weekly": "Week", "monthly": "Month"}
var spokenLangMap = map[string]string{"en": "English", "zh": "Chinese", "ru": "Russian"}

func (e *Trend) List(message *model.DiscordMessage, repos []*model.Repo) error {
	var (
		badge1       = findEmoji(e.ses, "badge1", message.GuildId, true)
		badge2       = findEmoji(e.ses, "badge2", message.GuildId, true)
		badge3       = findEmoji(e.ses, "badge3", message.GuildId, true)
		starAnimated = findEmoji(e.ses, "star_animated", message.GuildId, true)
	)
	numberEmojiStrings[0] = badge1
	numberEmojiStrings[1] = badge2
	numberEmojiStrings[2] = badge3
	title := "### :chart_with_upwards_trend: Github Trending Repositories\n"
	var content string
	if len(repos) != 0 {
		content += fmt.Sprintf("Trending repo in %s, %s, %s \n\n", cases.Title(language.AmericanEnglish).String(repos[0].ProgrammingLanguage), spokenLangMap[repos[0].SpokenLanguage], dateRangeStarGainedMap[repos[0].DateRange])
		// Set star gained text(today/last week/last month)
		for i := range repos {
			repo := repos[i]
			// Top 3 repos will have bigger title
			if i < 3 {
				content += fmt.Sprintf("%s [**%s**](%s)\n", numberEmojiStrings[i], repo.Name, repo.URL)

			} else {
				content += fmt.Sprintf("%s [**%s**](%s)\n", fmt.Sprintf(":%s:", numberEmojiStrings[i]), repo.Name, repo.URL)

			}
			truncatedDescription, _ := Truncate(repo.Description, 60)
			content += fmt.Sprintf("*%s*\n%s`%s: %s  |  Total: %s`\n",
				truncatedDescription,
				starAnimated,
				dateRangeStarGainedMap[repos[0].DateRange],
				rightPadding(fmt.Sprint(repo.StarGained), 3),
				fmt.Sprint(repo.StarCount),
			)
		}
	}
	msg := &discordgo.MessageEmbed{
		Title:       "",
		Description: title + content,
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Trend) ListDateRange(message *model.DiscordMessage) error {
	title := "### :chart_with_upwards_trend: Github Trending Repositories\n"
	content := ""
	for k := range dateRangeStarGainedMap {
		content += k + "\n"
	}
	msg := &discordgo.MessageEmbed{
		Description: title + "Available value for <date_range> parameter\n" + content,
	}
	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Trend) ListProgramLang(message *model.DiscordMessage) error {
	title := "### :chart_with_upwards_trend: Github Trending Repositories\n"
	content := ""
	for _, v := range commonProgrammingLanguage {
		content += v + ","
	}
	msg := &discordgo.MessageEmbed{
		Description: title + "Available value for <lang> parameter\n" + content,
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Trend) ListSpokenLang(message *model.DiscordMessage) error {
	title := "### :chart_with_upwards_trend: Github Trending Repositories\n"
	content := ""
	for k := range spokenLangMap {
		content += k + "\n"
	}
	msg := &discordgo.MessageEmbed{
		Description: title + "Available value for <spoken_lang> parameter\n" + content,
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func rightPadding(value interface{}, n int) string {
	// Convert the value to a string
	valueStr := fmt.Sprint(value)

	// Calculate the padding count
	paddingCount := n - len(valueStr)
	if paddingCount < 0 {
		paddingCount = 0
	}

	// Pad the value and return
	paddedValue := valueStr + strings.Repeat(" ", paddingCount)
	return paddedValue
}

func (e *Trend) GetAvailableProgrammingLang() []string {
	return programmingLanguages
}
func (e *Trend) GetAvailableSpokenLangMap() map[string]string {
	return spokenLangMap
}
func (e *Trend) GetAvaiableDateRangeMap() map[string]string {
	return dateRangeStarGainedMap
}
func (e *Trend) NotFound(message *model.DiscordMessage) error {
	msg := &discordgo.MessageEmbed{
		Description: "No trending repo found for your query",
	}
	return base.SendEmbededMessage(e.ses, message, msg)
}
func Truncate(text string, width int) (string, error) {
	if width < 0 {
		return "", fmt.Errorf("invalid width size")
	}
	if width >= len(text) {
		return text, nil
	}
	r := []rune(text)

	trunc := r[:width]
	return string(trunc) + "...", nil
}

var programmingLanguages = []string{"unknown", "1c-enterprise", "2-dimensional-array", "4d", "abap", "abap-cds", "abnf", "actionscript", "ada", "adblock-filter-list", "adobe-font-metrics", "agda", "ags-script", "aidl", "al", "alloy", "alpine-abuild", "altium-designer", "ampl", "angelscript", "ant-build-system", "antlers", "antlr", "apacheconf", "apex", "api-blueprint", "apl", "apollo-guidance-computer", "applescript", "arc", "asciidoc", "asl", "asn.1", "classic-asp", "asp.net", "aspectj", "assembly", "astro", "asymptote", "ats", "augeas", "autohotkey", "autoit", "avro-idl", "awk", "ballerina", "basic", "batchfile", "beef", "befunge", "berry", "bibtex", "bicep", "bikeshed", "bison", "bitbake", "blade", "blitzbasic", "blitzmax", "bluespec", "bluespec-bh", "boo", "boogie", "brainfuck", "brighterscript", "brightscript", "zeek", "browserslist", "c", "c%23", "c++", "c-objdump", "c2hs-haskell", "cabal-config", "cadence", "cairo", "cameligo", "cap-cds", "cap'n-proto", "cartocss", "ceylon", "chapel", "charity", "checksums", "chuck", "cil", "circom", "cirru", "clarion", "clarity", "classic-asp", "clean", "click", "clips", "clojure", "closure-templates", "cloud-firestore-security-rules", "cmake", "cobol", "codeowners", "codeql", "coffeescript", "coldfusion", "coldfusion-cfc", "collada", "common-lisp", "common-workflow-language", "component-pascal", "conll-u", "cool", "coq", "cpp-objdump", "creole", "crystal", "cson", "csound", "csound-document", "csound-score", "css", "csv", "cuda", "cue", "cue-sheet", "curl-config", "curry", "cweb", "cycript", "cypher", "cython", "d", "d-objdump", "d2", "dafny", "darcs-patch", "dart", "dataweave", "debian-package-control-file", "denizenscript", "desktop", "dhall", "diff", "digital-command-language", "dircolors", "directx-3d-file", "dm", "dns-zone", "dockerfile", "dogescript", "dotenv", "dtrace", "dylan", "e", "e-mail", "eagle", "earthly", "easybuild", "ebnf", "ec", "ecere-projects", "ecl", "eclipse", "ecmarkup", "editorconfig", "edje-data-collection", "edn", "eiffel", "ejs", "elixir", "elm", "elvish", "elvish-transcript", "emacs-lisp", "emberscript", "e-mail", "eq", "erlang", "euphoria", "f%23", "f*", "factor", "fancy", "fantom", "faust", "fennel", "figlet-font", "filebench-wml", "filterscript", "fish", "fluent", "flux", "formatted", "forth", "fortran", "fortran-free-form", "freebasic", "freemarker", "frege", "futhark", "g-code", "game-maker-language", "gaml", "gams", "gap", "gcc-machine-description", "gdb", "gdscript", "gedcom", "gemfile.lock", "gemini", "genero", "genero-forms", "genie", "genshi", "gentoo-ebuild", "gentoo-eclass", "gerber-image", "gettext-catalog", "gherkin", "git-attributes", "git-config", "git-revision-list", "gleam", "glsl", "glyph", "glyph-bitmap-distribution-format", "gn", "gnuplot", "/trending", "go-checksums", "go-module", "go-workspace", "godot-resource", "golo", "gosu", "grace", "gradle", "gradle-kotlin-dsl", "grammatical-framework", "graph-modeling-language", "graphql", "graphviz-(dot)", "groovy", "groovy-server-pages", "gsc", "hack", "haml", "handlebars", "haproxy", "harbour", "haskell", "haxe", "hcl", "hiveql", "hlsl", "hocon", "holyc", "hoon", "hosts-file", "html", "jinja", "html+ecr", "html+eex", "html+erb", "html+php", "html+razor", "http", "hxml", "hy", "hyphy", "idl", "idris", "ignore-list", "igor-pro", "imagej-macro", "imba", "inform-7", "ini", "ink", "inno-setup", "io", "ioke", "irc-log", "isabelle", "isabelle-root", "j", "janet", "jar-manifest", "jasmin", "java", "java-properties", "java-server-pages", "javascript", "javascript+erb", "jcl", "jest-snapshot", "jetbrains-mps", "jflex", "jinja", "jison", "jison-lex", "jolie", "jq", "json", "json-with-comments", "json5", "jsoniq", "jsonld", "jsonnet", "julia", "jupyter-notebook", "just", "kaitai-struct", "kakounescript", "kerboscript", "kicad-layout", "kicad-legacy-layout", "kicad-schematic", "kickstart", "kit", "kotlin", "krl", "kusto", "kvlang", "labview", "lark", "lasso", "latte", "lean", "less", "lex", "lfe", "ligolang", "lilypond", "limbo", "linker-script", "linux-kernel-module", "liquid", "literate-agda", "literate-coffeescript", "literate-haskell", "livescript", "llvm", "logos", "logtalk", "lolcode", "lookml", "loomscript", "lsl", "ltspice-symbol", "lua", "m", "m4", "m4sugar", "macaulay2", "makefile", "mako", "markdown", "marko", "mask", "mathematica", "matlab", "maven-pom", "max", "maxscript", "mcfunction", "mdx", "wikitext", "mercury", "mermaid", "meson", "metal", "microsoft-developer-studio-project", "microsoft-visual-studio-solution", "minid", "miniyaml", "mint", "mirah", "mirc-script", "mlir", "modelica", "modula-2", "modula-3", "module-management-system", "monkey", "monkey-c", "moocode", "moonscript", "motoko", "motorola-68k-assembly", "move", "mql4", "mql5", "mtml", "muf", "mupad", "muse", "mustache", "myghty", "nanorc", "nasal", "nasl", "ncl", "nearley", "nemerle", "neon", "nesc", "netlinx", "netlinx+erb", "netlogo", "newlisp", "nextflow", "nginx", "nim", "ninja", "nit", "nix", "nl", "npm-config", "nsis", "nu", "numpy", "nunjucks", "nushell", "nwscript", "oasv2-json", "oasv2-yaml", "oasv3-json", "oasv3-yaml", "objdump", "object-data-instance-notation", "objective-c", "objective-c++", "objective-j", "objectscript", "ocaml", "odin", "omgrofl", "ooc", "opa", "opal", "open-policy-agent", "openapi-specification-v2", "openapi-specification-v3", "opencl", "openedge-abl", "openqasm", "openrc-runscript", "openscad", "openstep-property-list", "opentype-feature-file", "option-list", "org", "ox", "oxygene", "oz", "p4", "pact", "pan", "papyrus", "parrot", "parrot-assembly", "parrot-internal-representation", "pascal", "pawn", "pddl", "peg.js", "pep8", "perl", "php", "pic", "pickle", "picolisp", "piglatin", "pike", "plantuml", "plpgsql", "plsql", "pod", "pod-6", "pogoscript", "polar", "pony", "portugol", "postcss", "postscript", "pov-ray-sdl", "powerbuilder", "powershell", "prisma", "processing", "procfile", "proguard", "prolog", "promela", "propeller-spin", "protocol-buffer", "protocol-buffer-text-format", "public-key", "pug", "puppet", "pure-data", "purebasic", "purescript", "pyret", "python", "python-console", "python-traceback", "q", "q%23", "qmake", "qml", "qt-script", "quake", "r", "racket", "ragel", "raku", "raml", "rascal", "raw-token-data", "rbs", "rdoc", "readline-config", "realbasic", "reason", "reasonligo", "rebol", "record-jar", "red", "redcode", "redirect-rules", "regular-expression", "ren'py", "renderscript", "rescript", "restructuredtext", "rexx", "rez", "rich-text-format", "ring", "riot", "rmarkdown", "robotframework", "robots.txt", "roff", "roff-manpage", "rouge", "routeros-script", "rpc", "rpgle", "rpm-spec", "ruby", "runoff", "rust", "sage", "saltstack", "sas", "sass", "scala", "scaml", "scenic", "scheme", "scilab", "scss", "sed", "self", "selinux-policy", "shaderlab", "shell", "shellcheck-config", "shellsession", "shen", "sieve", "simple-file-verification", "singularity", "slash", "slice", "slim", "smali", "smalltalk", "smarty", "smithy", "smpl", "smt", "snakemake", "solidity", "soong", "sourcepawn", "sparql", "spline-font-database", "sqf", "sql", "sqlpl", "squirrel", "srecode-template", "ssh-config", "stan", "standard-ml", "star", "starlark", "stata", "stl", "ston", "stringtemplate", "stylus", "subrip-text", "sugarss", "supercollider", "svelte", "svg", "sway", "sweave", "swift", "swig", "systemverilog", "talon", "tcl", "tcsh", "tea", "terra", "tex", "texinfo", "text", "textile", "textmate-properties", "thrift", "ti-program", "tl-verilog", "tla", "toml", "tsql", "tsv", "tsx", "turing", "turtle", "twig", "txl", "type-language", "typescript", "typst", "unified-parallel-c", "unity3d-asset", "unix-assembly", "uno", "unrealscript", "urweb", "v", "vala", "valve-data-format", "vba", "vbscript", "vcl", "velocity-template-language", "verilog", "vhdl", "vim-help-file", "vim-script", "vim-snippet", "visual-basic-.net", "visual-basic-.net", "visual-basic-6.0", "volt", "vue", "vyper", "wavefront-material", "wavefront-object", "wdl", "web-ontology-language", "webassembly", "webassembly-interface-type", "webidl", "webvtt", "wget-config", "wgsl", "whiley", "wikitext", "win32-message-file", "windows-registry-entries", "wisp", "witcher-script", "wollok", "world-of-warcraft-addon-data", "wren", "x-bitmap", "x-font-directory-index", "x-pixmap", "x10", "xbase", "xc", "xcompose", "xml", "xml-property-list", "xojo", "xonsh", "xpages", "xproc", "xquery", "xs", "xslt", "xtend", "yacc", "yaml", "yang", "yara", "yasnippet", "yul", "zap", "zeek", "zenscript", "zephir", "zig", "zil", "zimpl"}

// most popular 50 languages, value based on the variable above
var commonProgrammingLanguage = []string{"javascript", "python", "java", "ruby", "php", "c++", "c", "go", "c#", "shell", "objective-c", "r", "swift", "scala", "rust", "kotlin", "typescript", "dart", "elixir", "clojure", "haskell", "lua", "julia", "perl", "crystal", "nim", "vala", "ocaml", "groovy", "powershell", "erlang", "coffeescript", "fortran", "assembly", "matlab", "abap", "pascal", "vba", "elm", "hack", "racket", "d", "julia", "common-lisp", "scheme", "awk", "forth", "smalltalk", "prolog", "haxe", "f#", "apex", "puppet", "rebol", "tcl", "autohotkey", "dylan", "eiffel", "forth", "f"}
