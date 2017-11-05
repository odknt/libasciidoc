package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/bytesparadise/libasciidoc/types"
)

// *****************************************************************************************
// This file is generated after its sibling `asciidoc-grammar.peg` file. DO NOT MODIFY !
// *****************************************************************************************

var g = &grammar{
	rules: []*rule{
		{
			name: "Document",
			pos:  position{line: 16, col: 1, offset: 456},
			expr: &actionExpr{
				pos: position{line: 16, col: 13, offset: 468},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 16, col: 13, offset: 468},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 16, col: 13, offset: 468},
							label: "frontMatter",
							expr: &zeroOrOneExpr{
								pos: position{line: 16, col: 26, offset: 481},
								expr: &ruleRefExpr{
									pos:  position{line: 16, col: 26, offset: 481},
									name: "FrontMatter",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 16, col: 40, offset: 495},
							label: "documentHeader",
							expr: &zeroOrOneExpr{
								pos: position{line: 16, col: 56, offset: 511},
								expr: &ruleRefExpr{
									pos:  position{line: 16, col: 56, offset: 511},
									name: "DocumentHeader",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 16, col: 73, offset: 528},
							label: "blocks",
							expr: &ruleRefExpr{
								pos:  position{line: 16, col: 81, offset: 536},
								name: "DocumentBlocks",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 16, col: 97, offset: 552},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlocks",
			pos:  position{line: 20, col: 1, offset: 640},
			expr: &choiceExpr{
				pos: position{line: 20, col: 19, offset: 658},
				alternatives: []interface{}{
					&labeledExpr{
						pos:   position{line: 20, col: 19, offset: 658},
						label: "content",
						expr: &seqExpr{
							pos: position{line: 20, col: 28, offset: 667},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 20, col: 28, offset: 667},
									name: "Preamble",
								},
								&oneOrMoreExpr{
									pos: position{line: 20, col: 37, offset: 676},
									expr: &ruleRefExpr{
										pos:  position{line: 20, col: 37, offset: 676},
										name: "Section",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 20, col: 49, offset: 688},
						run: (*parser).callonDocumentBlocks7,
						expr: &labeledExpr{
							pos:   position{line: 20, col: 49, offset: 688},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 20, col: 58, offset: 697},
								expr: &ruleRefExpr{
									pos:  position{line: 20, col: 58, offset: 697},
									name: "StandaloneBlock",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "StandaloneBlock",
			pos:  position{line: 24, col: 1, offset: 744},
			expr: &choiceExpr{
				pos: position{line: 24, col: 20, offset: 763},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 24, col: 20, offset: 763},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 51, offset: 794},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 76, offset: 819},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 83, offset: 826},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 96, offset: 839},
						name: "LiteralBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 111, offset: 854},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 128, offset: 871},
						name: "Paragraph",
					},
					&seqExpr{
						pos: position{line: 24, col: 141, offset: 884},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 24, col: 141, offset: 884},
								name: "ElementAttribute",
							},
							&ruleRefExpr{
								pos:  position{line: 24, col: 158, offset: 901},
								name: "EOL",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 165, offset: 908},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "Preamble",
			pos:  position{line: 26, col: 1, offset: 963},
			expr: &actionExpr{
				pos: position{line: 26, col: 13, offset: 975},
				run: (*parser).callonPreamble1,
				expr: &labeledExpr{
					pos:   position{line: 26, col: 13, offset: 975},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 26, col: 23, offset: 985},
						expr: &ruleRefExpr{
							pos:  position{line: 26, col: 23, offset: 985},
							name: "StandaloneBlock",
						},
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 33, col: 1, offset: 1171},
			expr: &ruleRefExpr{
				pos:  position{line: 33, col: 16, offset: 1186},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 35, col: 1, offset: 1204},
			expr: &actionExpr{
				pos: position{line: 35, col: 16, offset: 1219},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 35, col: 16, offset: 1219},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 35, col: 16, offset: 1219},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 35, col: 37, offset: 1240},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 35, col: 45, offset: 1248},
								expr: &seqExpr{
									pos: position{line: 35, col: 46, offset: 1249},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 35, col: 46, offset: 1249},
											expr: &ruleRefExpr{
												pos:  position{line: 35, col: 47, offset: 1250},
												name: "YamlFrontMatterToken",
											},
										},
										&anyMatcher{
											line: 35, col: 68, offset: 1271,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 35, col: 72, offset: 1275},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 39, col: 1, offset: 1362},
			expr: &seqExpr{
				pos: position{line: 39, col: 26, offset: 1387},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 39, col: 26, offset: 1387},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 32, offset: 1393},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 45, col: 1, offset: 1582},
			expr: &actionExpr{
				pos: position{line: 45, col: 19, offset: 1600},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 45, col: 19, offset: 1600},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 45, col: 19, offset: 1600},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 45, col: 27, offset: 1608},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 45, col: 42, offset: 1623},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 45, col: 51, offset: 1632},
								expr: &ruleRefExpr{
									pos:  position{line: 45, col: 51, offset: 1632},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 45, col: 69, offset: 1650},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 45, col: 79, offset: 1660},
								expr: &ruleRefExpr{
									pos:  position{line: 45, col: 79, offset: 1660},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 45, col: 98, offset: 1679},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 45, col: 115, offset: 1696},
								expr: &ruleRefExpr{
									pos:  position{line: 45, col: 115, offset: 1696},
									name: "DocumentAttributeDeclaration",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentTitle",
			pos:  position{line: 49, col: 1, offset: 1827},
			expr: &actionExpr{
				pos: position{line: 49, col: 18, offset: 1844},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 49, col: 18, offset: 1844},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 49, col: 18, offset: 1844},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 49, col: 29, offset: 1855},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 30, offset: 1856},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 49, offset: 1875},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 49, col: 56, offset: 1882},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 49, col: 61, offset: 1887},
							expr: &ruleRefExpr{
								pos:  position{line: 49, col: 61, offset: 1887},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 65, offset: 1891},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 49, col: 73, offset: 1899},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 49, col: 87, offset: 1913},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 53, col: 1, offset: 2017},
			expr: &choiceExpr{
				pos: position{line: 53, col: 20, offset: 2036},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 53, col: 20, offset: 2036},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 48, offset: 2064},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 55, col: 1, offset: 2094},
			expr: &actionExpr{
				pos: position{line: 55, col: 30, offset: 2123},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 55, col: 30, offset: 2123},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 55, col: 30, offset: 2123},
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 30, offset: 2123},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 55, col: 34, offset: 2127},
							expr: &litMatcher{
								pos:        position{line: 55, col: 35, offset: 2128},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 55, col: 39, offset: 2132},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 55, col: 48, offset: 2141},
								expr: &ruleRefExpr{
									pos:  position{line: 55, col: 48, offset: 2141},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 65, offset: 2158},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 59, col: 1, offset: 2228},
			expr: &actionExpr{
				pos: position{line: 59, col: 33, offset: 2260},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 59, col: 33, offset: 2260},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 59, col: 33, offset: 2260},
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 33, offset: 2260},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 59, col: 37, offset: 2264},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 59, col: 48, offset: 2275},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 56, offset: 2283},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 63, col: 1, offset: 2376},
			expr: &actionExpr{
				pos: position{line: 63, col: 19, offset: 2394},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 63, col: 19, offset: 2394},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 63, col: 19, offset: 2394},
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 19, offset: 2394},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 63, col: 23, offset: 2398},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 34, offset: 2409},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 63, col: 58, offset: 2433},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 63, col: 68, offset: 2443},
								expr: &ruleRefExpr{
									pos:  position{line: 63, col: 69, offset: 2444},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 63, col: 94, offset: 2469},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 63, col: 104, offset: 2479},
								expr: &ruleRefExpr{
									pos:  position{line: 63, col: 105, offset: 2480},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 63, col: 130, offset: 2505},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 63, col: 136, offset: 2511},
								expr: &ruleRefExpr{
									pos:  position{line: 63, col: 137, offset: 2512},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 63, col: 159, offset: 2534},
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 159, offset: 2534},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 63, col: 163, offset: 2538},
							expr: &litMatcher{
								pos:        position{line: 63, col: 163, offset: 2538},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 63, col: 168, offset: 2543},
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 168, offset: 2543},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 68, col: 1, offset: 2708},
			expr: &seqExpr{
				pos: position{line: 68, col: 27, offset: 2734},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 68, col: 27, offset: 2734},
						expr: &litMatcher{
							pos:        position{line: 68, col: 28, offset: 2735},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 68, col: 32, offset: 2739},
						expr: &litMatcher{
							pos:        position{line: 68, col: 33, offset: 2740},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 37, offset: 2744},
						name: "Word",
					},
					&zeroOrMoreExpr{
						pos: position{line: 68, col: 42, offset: 2749},
						expr: &ruleRefExpr{
							pos:  position{line: 68, col: 42, offset: 2749},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 70, col: 1, offset: 2754},
			expr: &seqExpr{
				pos: position{line: 70, col: 24, offset: 2777},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 70, col: 24, offset: 2777},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 70, col: 28, offset: 2781},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 70, col: 34, offset: 2787},
							expr: &seqExpr{
								pos: position{line: 70, col: 35, offset: 2788},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 70, col: 35, offset: 2788},
										expr: &litMatcher{
											pos:        position{line: 70, col: 36, offset: 2789},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 70, col: 40, offset: 2793},
										expr: &ruleRefExpr{
											pos:  position{line: 70, col: 41, offset: 2794},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 70, col: 45, offset: 2798,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 70, col: 49, offset: 2802},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 74, col: 1, offset: 2938},
			expr: &actionExpr{
				pos: position{line: 74, col: 21, offset: 2958},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 74, col: 21, offset: 2958},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 74, col: 21, offset: 2958},
							expr: &ruleRefExpr{
								pos:  position{line: 74, col: 21, offset: 2958},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 74, col: 25, offset: 2962},
							expr: &litMatcher{
								pos:        position{line: 74, col: 26, offset: 2963},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 74, col: 30, offset: 2967},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 74, col: 40, offset: 2977},
								expr: &ruleRefExpr{
									pos:  position{line: 74, col: 41, offset: 2978},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 74, col: 66, offset: 3003},
							expr: &litMatcher{
								pos:        position{line: 74, col: 66, offset: 3003},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 74, col: 71, offset: 3008},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 74, col: 79, offset: 3016},
								expr: &ruleRefExpr{
									pos:  position{line: 74, col: 80, offset: 3017},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 74, col: 103, offset: 3040},
							expr: &litMatcher{
								pos:        position{line: 74, col: 103, offset: 3040},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 74, col: 108, offset: 3045},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 74, col: 118, offset: 3055},
								expr: &ruleRefExpr{
									pos:  position{line: 74, col: 119, offset: 3056},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 144, offset: 3081},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 79, col: 1, offset: 3254},
			expr: &choiceExpr{
				pos: position{line: 79, col: 27, offset: 3280},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 79, col: 27, offset: 3280},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 79, col: 27, offset: 3280},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 79, col: 32, offset: 3285},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 79, col: 39, offset: 3292},
								expr: &seqExpr{
									pos: position{line: 79, col: 40, offset: 3293},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 79, col: 40, offset: 3293},
											expr: &ruleRefExpr{
												pos:  position{line: 79, col: 41, offset: 3294},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 79, col: 45, offset: 3298},
											expr: &litMatcher{
												pos:        position{line: 79, col: 46, offset: 3299},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 79, col: 50, offset: 3303},
											expr: &litMatcher{
												pos:        position{line: 79, col: 51, offset: 3304},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 79, col: 55, offset: 3308,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 79, col: 61, offset: 3314},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 79, col: 61, offset: 3314},
								expr: &litMatcher{
									pos:        position{line: 79, col: 61, offset: 3314},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 79, col: 67, offset: 3320},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 79, col: 74, offset: 3327},
								expr: &seqExpr{
									pos: position{line: 79, col: 75, offset: 3328},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 79, col: 75, offset: 3328},
											expr: &ruleRefExpr{
												pos:  position{line: 79, col: 76, offset: 3329},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 79, col: 80, offset: 3333},
											expr: &litMatcher{
												pos:        position{line: 79, col: 81, offset: 3334},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 79, col: 85, offset: 3338},
											expr: &litMatcher{
												pos:        position{line: 79, col: 86, offset: 3339},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 79, col: 90, offset: 3343,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 79, col: 94, offset: 3347},
								expr: &ruleRefExpr{
									pos:  position{line: 79, col: 94, offset: 3347},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 79, col: 98, offset: 3351},
								expr: &litMatcher{
									pos:        position{line: 79, col: 99, offset: 3352},
									val:        ",",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionDate",
			pos:  position{line: 80, col: 1, offset: 3356},
			expr: &zeroOrMoreExpr{
				pos: position{line: 80, col: 25, offset: 3380},
				expr: &seqExpr{
					pos: position{line: 80, col: 26, offset: 3381},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 80, col: 26, offset: 3381},
							expr: &ruleRefExpr{
								pos:  position{line: 80, col: 27, offset: 3382},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 80, col: 31, offset: 3386},
							expr: &litMatcher{
								pos:        position{line: 80, col: 32, offset: 3387},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 80, col: 36, offset: 3391,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 81, col: 1, offset: 3396},
			expr: &zeroOrMoreExpr{
				pos: position{line: 81, col: 27, offset: 3422},
				expr: &seqExpr{
					pos: position{line: 81, col: 28, offset: 3423},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 81, col: 28, offset: 3423},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 29, offset: 3424},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 81, col: 33, offset: 3428,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 86, col: 1, offset: 3548},
			expr: &choiceExpr{
				pos: position{line: 86, col: 33, offset: 3580},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 86, col: 33, offset: 3580},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 86, col: 76, offset: 3623},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 88, col: 1, offset: 3670},
			expr: &actionExpr{
				pos: position{line: 88, col: 45, offset: 3714},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 88, col: 45, offset: 3714},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 88, col: 45, offset: 3714},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 88, col: 49, offset: 3718},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 55, offset: 3724},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 88, col: 70, offset: 3739},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 88, col: 74, offset: 3743},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 74, offset: 3743},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 78, offset: 3747},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 92, col: 1, offset: 3832},
			expr: &actionExpr{
				pos: position{line: 92, col: 49, offset: 3880},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 92, col: 49, offset: 3880},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 92, col: 49, offset: 3880},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 92, col: 53, offset: 3884},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 59, offset: 3890},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 92, col: 74, offset: 3905},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 92, col: 78, offset: 3909},
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 78, offset: 3909},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 92, col: 82, offset: 3913},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 92, col: 88, offset: 3919},
								expr: &seqExpr{
									pos: position{line: 92, col: 89, offset: 3920},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 92, col: 89, offset: 3920},
											expr: &ruleRefExpr{
												pos:  position{line: 92, col: 90, offset: 3921},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 92, col: 98, offset: 3929,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 92, col: 102, offset: 3933},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 96, col: 1, offset: 4036},
			expr: &choiceExpr{
				pos: position{line: 96, col: 27, offset: 4062},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 96, col: 27, offset: 4062},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 78, offset: 4113},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 98, col: 1, offset: 4159},
			expr: &actionExpr{
				pos: position{line: 98, col: 53, offset: 4211},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 98, col: 53, offset: 4211},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 98, col: 53, offset: 4211},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 98, col: 58, offset: 4216},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 64, offset: 4222},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 98, col: 79, offset: 4237},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 98, col: 83, offset: 4241},
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 83, offset: 4241},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 98, col: 87, offset: 4245},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 102, col: 1, offset: 4319},
			expr: &actionExpr{
				pos: position{line: 102, col: 49, offset: 4367},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 102, col: 49, offset: 4367},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 102, col: 49, offset: 4367},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 102, col: 53, offset: 4371},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 59, offset: 4377},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 102, col: 74, offset: 4392},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 102, col: 79, offset: 4397},
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 79, offset: 4397},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 83, offset: 4401},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 107, col: 1, offset: 4476},
			expr: &actionExpr{
				pos: position{line: 107, col: 34, offset: 4509},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 107, col: 34, offset: 4509},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 107, col: 34, offset: 4509},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 107, col: 38, offset: 4513},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 107, col: 44, offset: 4519},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 107, col: 59, offset: 4534},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 114, col: 1, offset: 4788},
			expr: &seqExpr{
				pos: position{line: 114, col: 18, offset: 4805},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 114, col: 19, offset: 4806},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 114, col: 19, offset: 4806},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 114, col: 27, offset: 4814},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 114, col: 35, offset: 4822},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 114, col: 43, offset: 4830},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 114, col: 48, offset: 4835},
						expr: &choiceExpr{
							pos: position{line: 114, col: 49, offset: 4836},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 114, col: 49, offset: 4836},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 114, col: 57, offset: 4844},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 114, col: 65, offset: 4852},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 114, col: 73, offset: 4860},
									val:        "-",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 119, col: 1, offset: 4971},
			expr: &choiceExpr{
				pos: position{line: 119, col: 12, offset: 4982},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 119, col: 12, offset: 4982},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 23, offset: 4993},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 34, offset: 5004},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 45, offset: 5015},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 56, offset: 5026},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 122, col: 1, offset: 5037},
			expr: &actionExpr{
				pos: position{line: 122, col: 13, offset: 5049},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 122, col: 13, offset: 5049},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 122, col: 13, offset: 5049},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 122, col: 21, offset: 5057},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 122, col: 36, offset: 5072},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 122, col: 46, offset: 5082},
								expr: &ruleRefExpr{
									pos:  position{line: 122, col: 46, offset: 5082},
									name: "Section1Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section1Block",
			pos:  position{line: 126, col: 1, offset: 5190},
			expr: &actionExpr{
				pos: position{line: 126, col: 18, offset: 5207},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 126, col: 18, offset: 5207},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 126, col: 18, offset: 5207},
							expr: &ruleRefExpr{
								pos:  position{line: 126, col: 19, offset: 5208},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 126, col: 28, offset: 5217},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 126, col: 37, offset: 5226},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 126, col: 37, offset: 5226},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 126, col: 48, offset: 5237},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 126, col: 59, offset: 5248},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 126, col: 70, offset: 5259},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 126, col: 81, offset: 5270},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2",
			pos:  position{line: 130, col: 1, offset: 5335},
			expr: &actionExpr{
				pos: position{line: 130, col: 13, offset: 5347},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 130, col: 13, offset: 5347},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 130, col: 13, offset: 5347},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 130, col: 21, offset: 5355},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 130, col: 36, offset: 5370},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 130, col: 46, offset: 5380},
								expr: &ruleRefExpr{
									pos:  position{line: 130, col: 46, offset: 5380},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 130, col: 62, offset: 5396},
							expr: &zeroOrMoreExpr{
								pos: position{line: 130, col: 63, offset: 5397},
								expr: &ruleRefExpr{
									pos:  position{line: 130, col: 64, offset: 5398},
									name: "Section2",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Block",
			pos:  position{line: 134, col: 1, offset: 5501},
			expr: &actionExpr{
				pos: position{line: 134, col: 18, offset: 5518},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 134, col: 18, offset: 5518},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 134, col: 18, offset: 5518},
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 19, offset: 5519},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 134, col: 28, offset: 5528},
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 29, offset: 5529},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 134, col: 38, offset: 5538},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 134, col: 47, offset: 5547},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 134, col: 47, offset: 5547},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 134, col: 58, offset: 5558},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 134, col: 69, offset: 5569},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 134, col: 80, offset: 5580},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3",
			pos:  position{line: 138, col: 1, offset: 5645},
			expr: &actionExpr{
				pos: position{line: 138, col: 13, offset: 5657},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 138, col: 13, offset: 5657},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 138, col: 13, offset: 5657},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 21, offset: 5665},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 138, col: 36, offset: 5680},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 138, col: 46, offset: 5690},
								expr: &ruleRefExpr{
									pos:  position{line: 138, col: 46, offset: 5690},
									name: "Section3Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Block",
			pos:  position{line: 142, col: 1, offset: 5798},
			expr: &actionExpr{
				pos: position{line: 142, col: 18, offset: 5815},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 142, col: 18, offset: 5815},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 142, col: 18, offset: 5815},
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 19, offset: 5816},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 142, col: 28, offset: 5825},
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 29, offset: 5826},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 142, col: 38, offset: 5835},
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 39, offset: 5836},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 142, col: 48, offset: 5845},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 142, col: 57, offset: 5854},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 142, col: 57, offset: 5854},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 142, col: 68, offset: 5865},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 142, col: 79, offset: 5876},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4",
			pos:  position{line: 146, col: 1, offset: 5941},
			expr: &actionExpr{
				pos: position{line: 146, col: 13, offset: 5953},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 146, col: 13, offset: 5953},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 146, col: 13, offset: 5953},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 21, offset: 5961},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 146, col: 36, offset: 5976},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 146, col: 46, offset: 5986},
								expr: &ruleRefExpr{
									pos:  position{line: 146, col: 46, offset: 5986},
									name: "Section4Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4Block",
			pos:  position{line: 150, col: 1, offset: 6094},
			expr: &actionExpr{
				pos: position{line: 150, col: 18, offset: 6111},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 150, col: 18, offset: 6111},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 150, col: 18, offset: 6111},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 19, offset: 6112},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 150, col: 28, offset: 6121},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 29, offset: 6122},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 150, col: 38, offset: 6131},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 39, offset: 6132},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 150, col: 48, offset: 6141},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 49, offset: 6142},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 150, col: 58, offset: 6151},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 150, col: 67, offset: 6160},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 150, col: 67, offset: 6160},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 150, col: 78, offset: 6171},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5",
			pos:  position{line: 154, col: 1, offset: 6236},
			expr: &actionExpr{
				pos: position{line: 154, col: 13, offset: 6248},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 154, col: 13, offset: 6248},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 154, col: 13, offset: 6248},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 21, offset: 6256},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 154, col: 36, offset: 6271},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 154, col: 46, offset: 6281},
								expr: &ruleRefExpr{
									pos:  position{line: 154, col: 46, offset: 6281},
									name: "Section5Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5Block",
			pos:  position{line: 158, col: 1, offset: 6389},
			expr: &actionExpr{
				pos: position{line: 158, col: 18, offset: 6406},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 158, col: 18, offset: 6406},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 158, col: 18, offset: 6406},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 19, offset: 6407},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 158, col: 28, offset: 6416},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 29, offset: 6417},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 158, col: 38, offset: 6426},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 39, offset: 6427},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 158, col: 48, offset: 6436},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 49, offset: 6437},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 158, col: 58, offset: 6446},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 59, offset: 6447},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 158, col: 68, offset: 6456},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 77, offset: 6465},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 166, col: 1, offset: 6641},
			expr: &choiceExpr{
				pos: position{line: 166, col: 17, offset: 6657},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 166, col: 17, offset: 6657},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 33, offset: 6673},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 49, offset: 6689},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 65, offset: 6705},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 81, offset: 6721},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 168, col: 1, offset: 6736},
			expr: &actionExpr{
				pos: position{line: 168, col: 18, offset: 6753},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 168, col: 18, offset: 6753},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 168, col: 18, offset: 6753},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 168, col: 29, offset: 6764},
								expr: &ruleRefExpr{
									pos:  position{line: 168, col: 30, offset: 6765},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 168, col: 49, offset: 6784},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 168, col: 56, offset: 6791},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 168, col: 62, offset: 6797},
							expr: &ruleRefExpr{
								pos:  position{line: 168, col: 62, offset: 6797},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 168, col: 66, offset: 6801},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 168, col: 74, offset: 6809},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 168, col: 88, offset: 6823},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 168, col: 93, offset: 6828},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 168, col: 93, offset: 6828},
									expr: &ruleRefExpr{
										pos:  position{line: 168, col: 93, offset: 6828},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 106, offset: 6841},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Title",
			pos:  position{line: 172, col: 1, offset: 6946},
			expr: &actionExpr{
				pos: position{line: 172, col: 18, offset: 6963},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 172, col: 18, offset: 6963},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 172, col: 18, offset: 6963},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 172, col: 29, offset: 6974},
								expr: &ruleRefExpr{
									pos:  position{line: 172, col: 30, offset: 6975},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 172, col: 49, offset: 6994},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 172, col: 56, offset: 7001},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 172, col: 63, offset: 7008},
							expr: &ruleRefExpr{
								pos:  position{line: 172, col: 63, offset: 7008},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 172, col: 67, offset: 7012},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 172, col: 75, offset: 7020},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 172, col: 89, offset: 7034},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 172, col: 94, offset: 7039},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 172, col: 94, offset: 7039},
									expr: &ruleRefExpr{
										pos:  position{line: 172, col: 94, offset: 7039},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 172, col: 107, offset: 7052},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Title",
			pos:  position{line: 176, col: 1, offset: 7156},
			expr: &actionExpr{
				pos: position{line: 176, col: 18, offset: 7173},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 176, col: 18, offset: 7173},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 176, col: 18, offset: 7173},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 176, col: 29, offset: 7184},
								expr: &ruleRefExpr{
									pos:  position{line: 176, col: 30, offset: 7185},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 49, offset: 7204},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 176, col: 56, offset: 7211},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 176, col: 64, offset: 7219},
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 64, offset: 7219},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 68, offset: 7223},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 76, offset: 7231},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 176, col: 90, offset: 7245},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 176, col: 95, offset: 7250},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 176, col: 95, offset: 7250},
									expr: &ruleRefExpr{
										pos:  position{line: 176, col: 95, offset: 7250},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 176, col: 108, offset: 7263},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4Title",
			pos:  position{line: 180, col: 1, offset: 7367},
			expr: &actionExpr{
				pos: position{line: 180, col: 18, offset: 7384},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 180, col: 18, offset: 7384},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 180, col: 18, offset: 7384},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 180, col: 29, offset: 7395},
								expr: &ruleRefExpr{
									pos:  position{line: 180, col: 30, offset: 7396},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 49, offset: 7415},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 180, col: 56, offset: 7422},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 180, col: 65, offset: 7431},
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 65, offset: 7431},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 69, offset: 7435},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 77, offset: 7443},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 180, col: 91, offset: 7457},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 180, col: 96, offset: 7462},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 180, col: 96, offset: 7462},
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 96, offset: 7462},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 109, offset: 7475},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5Title",
			pos:  position{line: 184, col: 1, offset: 7579},
			expr: &actionExpr{
				pos: position{line: 184, col: 18, offset: 7596},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 184, col: 18, offset: 7596},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 184, col: 18, offset: 7596},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 184, col: 29, offset: 7607},
								expr: &ruleRefExpr{
									pos:  position{line: 184, col: 30, offset: 7608},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 49, offset: 7627},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 184, col: 56, offset: 7634},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 184, col: 66, offset: 7644},
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 66, offset: 7644},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 70, offset: 7648},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 78, offset: 7656},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 184, col: 92, offset: 7670},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 184, col: 97, offset: 7675},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 184, col: 97, offset: 7675},
									expr: &ruleRefExpr{
										pos:  position{line: 184, col: 97, offset: 7675},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 184, col: 110, offset: 7688},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "List",
			pos:  position{line: 191, col: 1, offset: 7898},
			expr: &actionExpr{
				pos: position{line: 191, col: 9, offset: 7906},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 191, col: 9, offset: 7906},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 191, col: 9, offset: 7906},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 191, col: 20, offset: 7917},
								expr: &ruleRefExpr{
									pos:  position{line: 191, col: 21, offset: 7918},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 5, offset: 8010},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 193, col: 14, offset: 8019},
								expr: &seqExpr{
									pos: position{line: 193, col: 15, offset: 8020},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 193, col: 15, offset: 8020},
											name: "ListItem",
										},
										&zeroOrOneExpr{
											pos: position{line: 193, col: 24, offset: 8029},
											expr: &ruleRefExpr{
												pos:  position{line: 193, col: 24, offset: 8029},
												name: "BlankLine",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ListItem",
			pos:  position{line: 197, col: 1, offset: 8126},
			expr: &actionExpr{
				pos: position{line: 197, col: 13, offset: 8138},
				run: (*parser).callonListItem1,
				expr: &seqExpr{
					pos: position{line: 197, col: 13, offset: 8138},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 197, col: 13, offset: 8138},
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 13, offset: 8138},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 17, offset: 8142},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 197, col: 24, offset: 8149},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 197, col: 24, offset: 8149},
										expr: &litMatcher{
											pos:        position{line: 197, col: 24, offset: 8149},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 197, col: 31, offset: 8156},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 197, col: 36, offset: 8161},
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 36, offset: 8161},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 40, offset: 8165},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 49, offset: 8174},
								name: "ListItemContent",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContent",
			pos:  position{line: 201, col: 1, offset: 8271},
			expr: &actionExpr{
				pos: position{line: 201, col: 20, offset: 8290},
				run: (*parser).callonListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 201, col: 20, offset: 8290},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 201, col: 26, offset: 8296},
						expr: &seqExpr{
							pos: position{line: 201, col: 27, offset: 8297},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 201, col: 27, offset: 8297},
									expr: &seqExpr{
										pos: position{line: 201, col: 29, offset: 8299},
										exprs: []interface{}{
											&zeroOrMoreExpr{
												pos: position{line: 201, col: 29, offset: 8299},
												expr: &ruleRefExpr{
													pos:  position{line: 201, col: 29, offset: 8299},
													name: "WS",
												},
											},
											&choiceExpr{
												pos: position{line: 201, col: 34, offset: 8304},
												alternatives: []interface{}{
													&oneOrMoreExpr{
														pos: position{line: 201, col: 34, offset: 8304},
														expr: &litMatcher{
															pos:        position{line: 201, col: 34, offset: 8304},
															val:        "*",
															ignoreCase: false,
														},
													},
													&litMatcher{
														pos:        position{line: 201, col: 41, offset: 8311},
														val:        "-",
														ignoreCase: false,
													},
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 201, col: 46, offset: 8316},
												expr: &ruleRefExpr{
													pos:  position{line: 201, col: 46, offset: 8316},
													name: "WS",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 51, offset: 8321},
									name: "InlineContent",
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 65, offset: 8335},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 209, col: 1, offset: 8664},
			expr: &actionExpr{
				pos: position{line: 209, col: 14, offset: 8677},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 209, col: 14, offset: 8677},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 209, col: 14, offset: 8677},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 209, col: 25, offset: 8688},
								expr: &ruleRefExpr{
									pos:  position{line: 209, col: 26, offset: 8689},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 209, col: 45, offset: 8708},
							expr: &seqExpr{
								pos: position{line: 209, col: 47, offset: 8710},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 209, col: 47, offset: 8710},
										expr: &litMatcher{
											pos:        position{line: 209, col: 47, offset: 8710},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 209, col: 52, offset: 8715},
										expr: &ruleRefExpr{
											pos:  position{line: 209, col: 52, offset: 8715},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 209, col: 57, offset: 8720},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 209, col: 63, offset: 8726},
								expr: &seqExpr{
									pos: position{line: 209, col: 64, offset: 8727},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 209, col: 64, offset: 8727},
											name: "InlineContent",
										},
										&ruleRefExpr{
											pos:  position{line: 209, col: 78, offset: 8741},
											name: "EOL",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineContent",
			pos:  position{line: 215, col: 1, offset: 9031},
			expr: &actionExpr{
				pos: position{line: 215, col: 18, offset: 9048},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 215, col: 18, offset: 9048},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 215, col: 18, offset: 9048},
							expr: &ruleRefExpr{
								pos:  position{line: 215, col: 19, offset: 9049},
								name: "FencedBlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 215, col: 40, offset: 9070},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 215, col: 49, offset: 9079},
								expr: &seqExpr{
									pos: position{line: 215, col: 50, offset: 9080},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 215, col: 50, offset: 9080},
											expr: &ruleRefExpr{
												pos:  position{line: 215, col: 50, offset: 9080},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 215, col: 54, offset: 9084},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 215, col: 68, offset: 9098},
											expr: &ruleRefExpr{
												pos:  position{line: 215, col: 68, offset: 9098},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 215, col: 74, offset: 9104},
							expr: &ruleRefExpr{
								pos:  position{line: 215, col: 75, offset: 9105},
								name: "EOL",
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 219, col: 1, offset: 9223},
			expr: &choiceExpr{
				pos: position{line: 219, col: 18, offset: 9240},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 219, col: 18, offset: 9240},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 32, offset: 9254},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 45, offset: 9267},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 60, offset: 9282},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 92, offset: 9314},
						name: "Word",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 224, col: 1, offset: 9457},
			expr: &choiceExpr{
				pos: position{line: 224, col: 15, offset: 9471},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 224, col: 15, offset: 9471},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 224, col: 26, offset: 9482},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 224, col: 39, offset: 9495},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 225, col: 13, offset: 9523},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 225, col: 31, offset: 9541},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 225, col: 51, offset: 9561},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 227, col: 1, offset: 9583},
			expr: &choiceExpr{
				pos: position{line: 227, col: 13, offset: 9595},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 227, col: 13, offset: 9595},
						name: "BoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 227, col: 41, offset: 9623},
						name: "BoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 227, col: 73, offset: 9655},
						name: "BoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "BoldTextSimplePunctuation",
			pos:  position{line: 229, col: 1, offset: 9728},
			expr: &actionExpr{
				pos: position{line: 229, col: 30, offset: 9757},
				run: (*parser).callonBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 229, col: 30, offset: 9757},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 229, col: 30, offset: 9757},
							expr: &litMatcher{
								pos:        position{line: 229, col: 31, offset: 9758},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 229, col: 35, offset: 9762},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 229, col: 39, offset: 9766},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 48, offset: 9775},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 229, col: 67, offset: 9794},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextDoublePunctuation",
			pos:  position{line: 233, col: 1, offset: 9871},
			expr: &actionExpr{
				pos: position{line: 233, col: 30, offset: 9900},
				run: (*parser).callonBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 233, col: 30, offset: 9900},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 233, col: 30, offset: 9900},
							expr: &litMatcher{
								pos:        position{line: 233, col: 31, offset: 9901},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 233, col: 36, offset: 9906},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 233, col: 41, offset: 9911},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 233, col: 50, offset: 9920},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 233, col: 69, offset: 9939},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextUnbalancedPunctuation",
			pos:  position{line: 237, col: 1, offset: 10017},
			expr: &actionExpr{
				pos: position{line: 237, col: 34, offset: 10050},
				run: (*parser).callonBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 237, col: 34, offset: 10050},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 237, col: 34, offset: 10050},
							expr: &litMatcher{
								pos:        position{line: 237, col: 35, offset: 10051},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 237, col: 40, offset: 10056},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 237, col: 45, offset: 10061},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 54, offset: 10070},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 237, col: 73, offset: 10089},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 242, col: 1, offset: 10253},
			expr: &choiceExpr{
				pos: position{line: 242, col: 20, offset: 10272},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 242, col: 20, offset: 10272},
						name: "EscapedBoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 242, col: 55, offset: 10307},
						name: "EscapedBoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 242, col: 94, offset: 10346},
						name: "EscapedBoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedBoldTextSimplePunctuation",
			pos:  position{line: 244, col: 1, offset: 10426},
			expr: &actionExpr{
				pos: position{line: 244, col: 37, offset: 10462},
				run: (*parser).callonEscapedBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 244, col: 37, offset: 10462},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 244, col: 37, offset: 10462},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 244, col: 50, offset: 10475},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 244, col: 50, offset: 10475},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 244, col: 54, offset: 10479},
										expr: &litMatcher{
											pos:        position{line: 244, col: 54, offset: 10479},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 244, col: 60, offset: 10485},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 244, col: 64, offset: 10489},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 73, offset: 10498},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 244, col: 92, offset: 10517},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextDoublePunctuation",
			pos:  position{line: 248, col: 1, offset: 10623},
			expr: &actionExpr{
				pos: position{line: 248, col: 37, offset: 10659},
				run: (*parser).callonEscapedBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 248, col: 37, offset: 10659},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 248, col: 37, offset: 10659},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 248, col: 50, offset: 10672},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 248, col: 50, offset: 10672},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 248, col: 55, offset: 10677},
										expr: &litMatcher{
											pos:        position{line: 248, col: 55, offset: 10677},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 248, col: 61, offset: 10683},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 248, col: 66, offset: 10688},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 75, offset: 10697},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 248, col: 94, offset: 10716},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextUnbalancedPunctuation",
			pos:  position{line: 252, col: 1, offset: 10824},
			expr: &actionExpr{
				pos: position{line: 252, col: 42, offset: 10865},
				run: (*parser).callonEscapedBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 252, col: 42, offset: 10865},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 252, col: 42, offset: 10865},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 252, col: 55, offset: 10878},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 252, col: 55, offset: 10878},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 252, col: 59, offset: 10882},
										expr: &litMatcher{
											pos:        position{line: 252, col: 59, offset: 10882},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 252, col: 65, offset: 10888},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 252, col: 70, offset: 10893},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 79, offset: 10902},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 252, col: 98, offset: 10921},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 257, col: 1, offset: 11114},
			expr: &choiceExpr{
				pos: position{line: 257, col: 15, offset: 11128},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 257, col: 15, offset: 11128},
						name: "ItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 257, col: 45, offset: 11158},
						name: "ItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 257, col: 79, offset: 11192},
						name: "ItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "ItalicTextSimplePunctuation",
			pos:  position{line: 259, col: 1, offset: 11221},
			expr: &actionExpr{
				pos: position{line: 259, col: 32, offset: 11252},
				run: (*parser).callonItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 259, col: 32, offset: 11252},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 259, col: 32, offset: 11252},
							expr: &litMatcher{
								pos:        position{line: 259, col: 33, offset: 11253},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 259, col: 37, offset: 11257},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 259, col: 41, offset: 11261},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 259, col: 50, offset: 11270},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 259, col: 69, offset: 11289},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextDoublePunctuation",
			pos:  position{line: 263, col: 1, offset: 11368},
			expr: &actionExpr{
				pos: position{line: 263, col: 32, offset: 11399},
				run: (*parser).callonItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 263, col: 32, offset: 11399},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 263, col: 32, offset: 11399},
							expr: &litMatcher{
								pos:        position{line: 263, col: 33, offset: 11400},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 263, col: 38, offset: 11405},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 263, col: 43, offset: 11410},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 263, col: 52, offset: 11419},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 263, col: 71, offset: 11438},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextUnbalancedPunctuation",
			pos:  position{line: 267, col: 1, offset: 11518},
			expr: &actionExpr{
				pos: position{line: 267, col: 36, offset: 11553},
				run: (*parser).callonItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 267, col: 36, offset: 11553},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 267, col: 36, offset: 11553},
							expr: &litMatcher{
								pos:        position{line: 267, col: 37, offset: 11554},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 267, col: 42, offset: 11559},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 267, col: 47, offset: 11564},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 267, col: 56, offset: 11573},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 267, col: 75, offset: 11592},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 272, col: 1, offset: 11758},
			expr: &choiceExpr{
				pos: position{line: 272, col: 22, offset: 11779},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 272, col: 22, offset: 11779},
						name: "EscapedItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 272, col: 59, offset: 11816},
						name: "EscapedItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 272, col: 100, offset: 11857},
						name: "EscapedItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedItalicTextSimplePunctuation",
			pos:  position{line: 274, col: 1, offset: 11939},
			expr: &actionExpr{
				pos: position{line: 274, col: 39, offset: 11977},
				run: (*parser).callonEscapedItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 274, col: 39, offset: 11977},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 274, col: 39, offset: 11977},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 274, col: 52, offset: 11990},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 274, col: 52, offset: 11990},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 274, col: 56, offset: 11994},
										expr: &litMatcher{
											pos:        position{line: 274, col: 56, offset: 11994},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 274, col: 62, offset: 12000},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 274, col: 66, offset: 12004},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 274, col: 75, offset: 12013},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 274, col: 94, offset: 12032},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextDoublePunctuation",
			pos:  position{line: 278, col: 1, offset: 12138},
			expr: &actionExpr{
				pos: position{line: 278, col: 39, offset: 12176},
				run: (*parser).callonEscapedItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 278, col: 39, offset: 12176},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 278, col: 39, offset: 12176},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 278, col: 52, offset: 12189},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 278, col: 52, offset: 12189},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 278, col: 57, offset: 12194},
										expr: &litMatcher{
											pos:        position{line: 278, col: 57, offset: 12194},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 278, col: 63, offset: 12200},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 278, col: 68, offset: 12205},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 278, col: 77, offset: 12214},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 278, col: 96, offset: 12233},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextUnbalancedPunctuation",
			pos:  position{line: 282, col: 1, offset: 12341},
			expr: &actionExpr{
				pos: position{line: 282, col: 44, offset: 12384},
				run: (*parser).callonEscapedItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 282, col: 44, offset: 12384},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 282, col: 44, offset: 12384},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 282, col: 57, offset: 12397},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 282, col: 57, offset: 12397},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 282, col: 61, offset: 12401},
										expr: &litMatcher{
											pos:        position{line: 282, col: 61, offset: 12401},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 282, col: 67, offset: 12407},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 282, col: 72, offset: 12412},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 81, offset: 12421},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 282, col: 100, offset: 12440},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 287, col: 1, offset: 12633},
			expr: &choiceExpr{
				pos: position{line: 287, col: 18, offset: 12650},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 287, col: 18, offset: 12650},
						name: "MonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 287, col: 51, offset: 12683},
						name: "MonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 287, col: 88, offset: 12720},
						name: "MonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "MonospaceTextSimplePunctuation",
			pos:  position{line: 289, col: 1, offset: 12752},
			expr: &actionExpr{
				pos: position{line: 289, col: 35, offset: 12786},
				run: (*parser).callonMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 289, col: 35, offset: 12786},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 289, col: 35, offset: 12786},
							expr: &litMatcher{
								pos:        position{line: 289, col: 36, offset: 12787},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 289, col: 40, offset: 12791},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 289, col: 44, offset: 12795},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 53, offset: 12804},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 289, col: 72, offset: 12823},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextDoublePunctuation",
			pos:  position{line: 293, col: 1, offset: 12905},
			expr: &actionExpr{
				pos: position{line: 293, col: 35, offset: 12939},
				run: (*parser).callonMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 293, col: 35, offset: 12939},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 293, col: 35, offset: 12939},
							expr: &litMatcher{
								pos:        position{line: 293, col: 36, offset: 12940},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 293, col: 41, offset: 12945},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 293, col: 46, offset: 12950},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 55, offset: 12959},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 293, col: 74, offset: 12978},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextUnbalancedPunctuation",
			pos:  position{line: 297, col: 1, offset: 13061},
			expr: &actionExpr{
				pos: position{line: 297, col: 39, offset: 13099},
				run: (*parser).callonMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 297, col: 39, offset: 13099},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 297, col: 39, offset: 13099},
							expr: &litMatcher{
								pos:        position{line: 297, col: 40, offset: 13100},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 297, col: 45, offset: 13105},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 297, col: 50, offset: 13110},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 297, col: 59, offset: 13119},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 297, col: 78, offset: 13138},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 302, col: 1, offset: 13307},
			expr: &choiceExpr{
				pos: position{line: 302, col: 25, offset: 13331},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 302, col: 25, offset: 13331},
						name: "EscapedMonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 302, col: 65, offset: 13371},
						name: "EscapedMonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 302, col: 109, offset: 13415},
						name: "EscapedMonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextSimplePunctuation",
			pos:  position{line: 304, col: 1, offset: 13500},
			expr: &actionExpr{
				pos: position{line: 304, col: 42, offset: 13541},
				run: (*parser).callonEscapedMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 304, col: 42, offset: 13541},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 304, col: 42, offset: 13541},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 304, col: 55, offset: 13554},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 304, col: 55, offset: 13554},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 304, col: 59, offset: 13558},
										expr: &litMatcher{
											pos:        position{line: 304, col: 59, offset: 13558},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 304, col: 65, offset: 13564},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 304, col: 69, offset: 13568},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 304, col: 78, offset: 13577},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 304, col: 97, offset: 13596},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextDoublePunctuation",
			pos:  position{line: 308, col: 1, offset: 13702},
			expr: &actionExpr{
				pos: position{line: 308, col: 42, offset: 13743},
				run: (*parser).callonEscapedMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 308, col: 42, offset: 13743},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 308, col: 42, offset: 13743},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 308, col: 55, offset: 13756},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 308, col: 55, offset: 13756},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 308, col: 60, offset: 13761},
										expr: &litMatcher{
											pos:        position{line: 308, col: 60, offset: 13761},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 308, col: 66, offset: 13767},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 308, col: 71, offset: 13772},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 308, col: 80, offset: 13781},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 308, col: 99, offset: 13800},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextUnbalancedPunctuation",
			pos:  position{line: 312, col: 1, offset: 13908},
			expr: &actionExpr{
				pos: position{line: 312, col: 47, offset: 13954},
				run: (*parser).callonEscapedMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 312, col: 47, offset: 13954},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 312, col: 47, offset: 13954},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 312, col: 60, offset: 13967},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 312, col: 60, offset: 13967},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 312, col: 64, offset: 13971},
										expr: &litMatcher{
											pos:        position{line: 312, col: 64, offset: 13971},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 312, col: 70, offset: 13977},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 312, col: 75, offset: 13982},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 312, col: 84, offset: 13991},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 312, col: 103, offset: 14010},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 317, col: 1, offset: 14203},
			expr: &seqExpr{
				pos: position{line: 317, col: 22, offset: 14224},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 317, col: 22, offset: 14224},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 317, col: 47, offset: 14249},
						expr: &seqExpr{
							pos: position{line: 317, col: 48, offset: 14250},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 317, col: 48, offset: 14250},
									expr: &ruleRefExpr{
										pos:  position{line: 317, col: 48, offset: 14250},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 317, col: 52, offset: 14254},
									name: "QuotedTextContentElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContentElement",
			pos:  position{line: 319, col: 1, offset: 14282},
			expr: &choiceExpr{
				pos: position{line: 319, col: 29, offset: 14310},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 319, col: 29, offset: 14310},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 319, col: 42, offset: 14323},
						name: "QuotedTextWord",
					},
					&ruleRefExpr{
						pos:  position{line: 319, col: 59, offset: 14340},
						name: "WordWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextWord",
			pos:  position{line: 321, col: 1, offset: 14469},
			expr: &oneOrMoreExpr{
				pos: position{line: 321, col: 19, offset: 14487},
				expr: &seqExpr{
					pos: position{line: 321, col: 20, offset: 14488},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 321, col: 20, offset: 14488},
							expr: &ruleRefExpr{
								pos:  position{line: 321, col: 21, offset: 14489},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 321, col: 29, offset: 14497},
							expr: &ruleRefExpr{
								pos:  position{line: 321, col: 30, offset: 14498},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 321, col: 33, offset: 14501},
							expr: &litMatcher{
								pos:        position{line: 321, col: 34, offset: 14502},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 321, col: 38, offset: 14506},
							expr: &litMatcher{
								pos:        position{line: 321, col: 39, offset: 14507},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 321, col: 43, offset: 14511},
							expr: &litMatcher{
								pos:        position{line: 321, col: 44, offset: 14512},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 321, col: 48, offset: 14516,
						},
					},
				},
			},
		},
		{
			name: "WordWithQuotePunctuation",
			pos:  position{line: 322, col: 1, offset: 14558},
			expr: &actionExpr{
				pos: position{line: 322, col: 29, offset: 14586},
				run: (*parser).callonWordWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 322, col: 29, offset: 14586},
					expr: &seqExpr{
						pos: position{line: 322, col: 30, offset: 14587},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 322, col: 30, offset: 14587},
								expr: &ruleRefExpr{
									pos:  position{line: 322, col: 31, offset: 14588},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 322, col: 39, offset: 14596},
								expr: &ruleRefExpr{
									pos:  position{line: 322, col: 40, offset: 14597},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 322, col: 44, offset: 14601,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 327, col: 1, offset: 14846},
			expr: &choiceExpr{
				pos: position{line: 327, col: 31, offset: 14876},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 327, col: 31, offset: 14876},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 327, col: 37, offset: 14882},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 327, col: 43, offset: 14888},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 333, col: 1, offset: 14995},
			expr: &actionExpr{
				pos: position{line: 333, col: 17, offset: 15011},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 333, col: 17, offset: 15011},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 333, col: 17, offset: 15011},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 333, col: 22, offset: 15016},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 333, col: 22, offset: 15016},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 333, col: 33, offset: 15027},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 333, col: 38, offset: 15032},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 333, col: 43, offset: 15037},
								expr: &seqExpr{
									pos: position{line: 333, col: 44, offset: 15038},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 333, col: 44, offset: 15038},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 333, col: 48, offset: 15042},
											expr: &ruleRefExpr{
												pos:  position{line: 333, col: 49, offset: 15043},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 333, col: 60, offset: 15054},
											val:        "]",
											ignoreCase: false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BlockImage",
			pos:  position{line: 343, col: 1, offset: 15333},
			expr: &actionExpr{
				pos: position{line: 343, col: 15, offset: 15347},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 343, col: 15, offset: 15347},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 343, col: 15, offset: 15347},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 343, col: 26, offset: 15358},
								expr: &ruleRefExpr{
									pos:  position{line: 343, col: 27, offset: 15359},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 343, col: 46, offset: 15378},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 343, col: 52, offset: 15384},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 343, col: 69, offset: 15401},
							expr: &ruleRefExpr{
								pos:  position{line: 343, col: 69, offset: 15401},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 343, col: 73, offset: 15405},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 348, col: 1, offset: 15566},
			expr: &actionExpr{
				pos: position{line: 348, col: 20, offset: 15585},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 348, col: 20, offset: 15585},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 348, col: 20, offset: 15585},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 348, col: 30, offset: 15595},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 36, offset: 15601},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 348, col: 41, offset: 15606},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 348, col: 45, offset: 15610},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 348, col: 57, offset: 15622},
								expr: &ruleRefExpr{
									pos:  position{line: 348, col: 57, offset: 15622},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 348, col: 68, offset: 15633},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 352, col: 1, offset: 15700},
			expr: &actionExpr{
				pos: position{line: 352, col: 16, offset: 15715},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 352, col: 16, offset: 15715},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 352, col: 22, offset: 15721},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 357, col: 1, offset: 15868},
			expr: &actionExpr{
				pos: position{line: 357, col: 21, offset: 15888},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 357, col: 21, offset: 15888},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 357, col: 21, offset: 15888},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 357, col: 30, offset: 15897},
							expr: &litMatcher{
								pos:        position{line: 357, col: 31, offset: 15898},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 357, col: 35, offset: 15902},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 357, col: 41, offset: 15908},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 357, col: 46, offset: 15913},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 357, col: 50, offset: 15917},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 357, col: 62, offset: 15929},
								expr: &ruleRefExpr{
									pos:  position{line: 357, col: 62, offset: 15929},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 357, col: 73, offset: 15940},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 364, col: 1, offset: 16270},
			expr: &ruleRefExpr{
				pos:  position{line: 364, col: 19, offset: 16288},
				name: "FencedBlock",
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 366, col: 1, offset: 16302},
			expr: &actionExpr{
				pos: position{line: 366, col: 16, offset: 16317},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 366, col: 16, offset: 16317},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 366, col: 16, offset: 16317},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 366, col: 37, offset: 16338},
							expr: &ruleRefExpr{
								pos:  position{line: 366, col: 37, offset: 16338},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 366, col: 41, offset: 16342},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 366, col: 49, offset: 16350},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 366, col: 58, offset: 16359},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 366, col: 78, offset: 16379},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 366, col: 99, offset: 16400},
							expr: &ruleRefExpr{
								pos:  position{line: 366, col: 99, offset: 16400},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 366, col: 103, offset: 16404},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 370, col: 1, offset: 16492},
			expr: &litMatcher{
				pos:        position{line: 370, col: 25, offset: 16516},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 372, col: 1, offset: 16523},
			expr: &labeledExpr{
				pos:   position{line: 372, col: 23, offset: 16545},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 372, col: 31, offset: 16553},
					expr: &seqExpr{
						pos: position{line: 372, col: 32, offset: 16554},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 372, col: 32, offset: 16554},
								expr: &ruleRefExpr{
									pos:  position{line: 372, col: 33, offset: 16555},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 372, col: 54, offset: 16576,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 377, col: 1, offset: 16849},
			expr: &choiceExpr{
				pos: position{line: 377, col: 17, offset: 16865},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 377, col: 17, offset: 16865},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 377, col: 39, offset: 16887},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 377, col: 76, offset: 16924},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 380, col: 1, offset: 17019},
			expr: &actionExpr{
				pos: position{line: 380, col: 24, offset: 17042},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 380, col: 24, offset: 17042},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 380, col: 24, offset: 17042},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 380, col: 32, offset: 17050},
								expr: &ruleRefExpr{
									pos:  position{line: 380, col: 32, offset: 17050},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 380, col: 37, offset: 17055},
							expr: &ruleRefExpr{
								pos:  position{line: 380, col: 38, offset: 17056},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 380, col: 46, offset: 17064},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 380, col: 55, offset: 17073},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 380, col: 76, offset: 17094},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 385, col: 1, offset: 17275},
			expr: &actionExpr{
				pos: position{line: 385, col: 24, offset: 17298},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 385, col: 24, offset: 17298},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 385, col: 32, offset: 17306},
						expr: &seqExpr{
							pos: position{line: 385, col: 33, offset: 17307},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 385, col: 33, offset: 17307},
									expr: &seqExpr{
										pos: position{line: 385, col: 35, offset: 17309},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 385, col: 35, offset: 17309},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 385, col: 43, offset: 17317},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 385, col: 54, offset: 17328,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 390, col: 1, offset: 17413},
			expr: &choiceExpr{
				pos: position{line: 390, col: 22, offset: 17434},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 390, col: 22, offset: 17434},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 390, col: 22, offset: 17434},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 390, col: 30, offset: 17442},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 390, col: 42, offset: 17454},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 390, col: 52, offset: 17464},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 393, col: 1, offset: 17524},
			expr: &actionExpr{
				pos: position{line: 393, col: 39, offset: 17562},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 393, col: 39, offset: 17562},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 393, col: 39, offset: 17562},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 393, col: 61, offset: 17584},
							expr: &ruleRefExpr{
								pos:  position{line: 393, col: 61, offset: 17584},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 393, col: 65, offset: 17588},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 393, col: 73, offset: 17596},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 393, col: 81, offset: 17604},
								expr: &seqExpr{
									pos: position{line: 393, col: 82, offset: 17605},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 393, col: 82, offset: 17605},
											expr: &ruleRefExpr{
												pos:  position{line: 393, col: 83, offset: 17606},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 393, col: 105, offset: 17628,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 393, col: 109, offset: 17632},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 393, col: 131, offset: 17654},
							expr: &ruleRefExpr{
								pos:  position{line: 393, col: 131, offset: 17654},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 393, col: 135, offset: 17658},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 397, col: 1, offset: 17742},
			expr: &litMatcher{
				pos:        position{line: 397, col: 26, offset: 17767},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 400, col: 1, offset: 17829},
			expr: &actionExpr{
				pos: position{line: 400, col: 34, offset: 17862},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 400, col: 34, offset: 17862},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 400, col: 34, offset: 17862},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 400, col: 46, offset: 17874},
							expr: &ruleRefExpr{
								pos:  position{line: 400, col: 46, offset: 17874},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 400, col: 50, offset: 17878},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 400, col: 58, offset: 17886},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 400, col: 67, offset: 17895},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 400, col: 88, offset: 17916},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 407, col: 1, offset: 18128},
			expr: &labeledExpr{
				pos:   position{line: 407, col: 21, offset: 18148},
				label: "meta",
				expr: &choiceExpr{
					pos: position{line: 407, col: 27, offset: 18154},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 407, col: 27, offset: 18154},
							name: "ElementLink",
						},
						&ruleRefExpr{
							pos:  position{line: 407, col: 41, offset: 18168},
							name: "ElementID",
						},
						&ruleRefExpr{
							pos:  position{line: 407, col: 53, offset: 18180},
							name: "ElementTitle",
						},
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 410, col: 1, offset: 18251},
			expr: &actionExpr{
				pos: position{line: 410, col: 16, offset: 18266},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 410, col: 16, offset: 18266},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 410, col: 16, offset: 18266},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 410, col: 20, offset: 18270},
							expr: &ruleRefExpr{
								pos:  position{line: 410, col: 20, offset: 18270},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 410, col: 24, offset: 18274},
							val:        "link",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 410, col: 31, offset: 18281},
							expr: &ruleRefExpr{
								pos:  position{line: 410, col: 31, offset: 18281},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 410, col: 35, offset: 18285},
							val:        "=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 410, col: 39, offset: 18289},
							expr: &ruleRefExpr{
								pos:  position{line: 410, col: 39, offset: 18289},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 410, col: 43, offset: 18293},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 410, col: 48, offset: 18298},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 410, col: 52, offset: 18302},
							expr: &ruleRefExpr{
								pos:  position{line: 410, col: 52, offset: 18302},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 410, col: 56, offset: 18306},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 410, col: 60, offset: 18310},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 415, col: 1, offset: 18420},
			expr: &actionExpr{
				pos: position{line: 415, col: 14, offset: 18433},
				run: (*parser).callonElementID1,
				expr: &seqExpr{
					pos: position{line: 415, col: 14, offset: 18433},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 415, col: 14, offset: 18433},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 415, col: 18, offset: 18437},
							expr: &ruleRefExpr{
								pos:  position{line: 415, col: 18, offset: 18437},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 415, col: 22, offset: 18441},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 415, col: 26, offset: 18445},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 415, col: 30, offset: 18449},
								name: "ID",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 415, col: 34, offset: 18453},
							expr: &ruleRefExpr{
								pos:  position{line: 415, col: 34, offset: 18453},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 415, col: 38, offset: 18457},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 415, col: 42, offset: 18461},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 421, col: 1, offset: 18655},
			expr: &actionExpr{
				pos: position{line: 421, col: 17, offset: 18671},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 421, col: 17, offset: 18671},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 421, col: 17, offset: 18671},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 421, col: 21, offset: 18675},
							expr: &litMatcher{
								pos:        position{line: 421, col: 22, offset: 18676},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 421, col: 26, offset: 18680},
							expr: &ruleRefExpr{
								pos:  position{line: 421, col: 27, offset: 18681},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 421, col: 30, offset: 18684},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 421, col: 36, offset: 18690},
								expr: &seqExpr{
									pos: position{line: 421, col: 37, offset: 18691},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 421, col: 37, offset: 18691},
											expr: &ruleRefExpr{
												pos:  position{line: 421, col: 38, offset: 18692},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 421, col: 46, offset: 18700,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 421, col: 50, offset: 18704},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Word",
			pos:  position{line: 428, col: 1, offset: 18875},
			expr: &actionExpr{
				pos: position{line: 428, col: 9, offset: 18883},
				run: (*parser).callonWord1,
				expr: &oneOrMoreExpr{
					pos: position{line: 428, col: 9, offset: 18883},
					expr: &seqExpr{
						pos: position{line: 428, col: 10, offset: 18884},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 428, col: 10, offset: 18884},
								expr: &ruleRefExpr{
									pos:  position{line: 428, col: 11, offset: 18885},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 428, col: 19, offset: 18893},
								expr: &ruleRefExpr{
									pos:  position{line: 428, col: 20, offset: 18894},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 428, col: 23, offset: 18897,
							},
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 432, col: 1, offset: 18937},
			expr: &actionExpr{
				pos: position{line: 432, col: 14, offset: 18950},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 432, col: 14, offset: 18950},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 432, col: 14, offset: 18950},
							expr: &ruleRefExpr{
								pos:  position{line: 432, col: 15, offset: 18951},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 432, col: 19, offset: 18955},
							expr: &ruleRefExpr{
								pos:  position{line: 432, col: 19, offset: 18955},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 432, col: 23, offset: 18959},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 436, col: 1, offset: 19000},
			expr: &actionExpr{
				pos: position{line: 436, col: 8, offset: 19007},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 436, col: 8, offset: 19007},
					expr: &seqExpr{
						pos: position{line: 436, col: 9, offset: 19008},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 436, col: 9, offset: 19008},
								expr: &ruleRefExpr{
									pos:  position{line: 436, col: 10, offset: 19009},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 436, col: 18, offset: 19017},
								expr: &ruleRefExpr{
									pos:  position{line: 436, col: 19, offset: 19018},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 436, col: 22, offset: 19021},
								expr: &litMatcher{
									pos:        position{line: 436, col: 23, offset: 19022},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 436, col: 27, offset: 19026},
								expr: &litMatcher{
									pos:        position{line: 436, col: 28, offset: 19027},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 436, col: 32, offset: 19031,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 440, col: 1, offset: 19071},
			expr: &actionExpr{
				pos: position{line: 440, col: 7, offset: 19077},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 440, col: 7, offset: 19077},
					expr: &seqExpr{
						pos: position{line: 440, col: 8, offset: 19078},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 440, col: 8, offset: 19078},
								expr: &ruleRefExpr{
									pos:  position{line: 440, col: 9, offset: 19079},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 440, col: 17, offset: 19087},
								expr: &ruleRefExpr{
									pos:  position{line: 440, col: 18, offset: 19088},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 440, col: 21, offset: 19091},
								expr: &litMatcher{
									pos:        position{line: 440, col: 22, offset: 19092},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 440, col: 26, offset: 19096},
								expr: &litMatcher{
									pos:        position{line: 440, col: 27, offset: 19097},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 440, col: 31, offset: 19101,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 444, col: 1, offset: 19141},
			expr: &actionExpr{
				pos: position{line: 444, col: 13, offset: 19153},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 444, col: 13, offset: 19153},
					expr: &seqExpr{
						pos: position{line: 444, col: 14, offset: 19154},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 444, col: 14, offset: 19154},
								expr: &ruleRefExpr{
									pos:  position{line: 444, col: 15, offset: 19155},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 444, col: 23, offset: 19163},
								expr: &litMatcher{
									pos:        position{line: 444, col: 24, offset: 19164},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 444, col: 28, offset: 19168},
								expr: &litMatcher{
									pos:        position{line: 444, col: 29, offset: 19169},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 444, col: 33, offset: 19173,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 448, col: 1, offset: 19213},
			expr: &choiceExpr{
				pos: position{line: 448, col: 15, offset: 19227},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 448, col: 15, offset: 19227},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 448, col: 27, offset: 19239},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 448, col: 40, offset: 19252},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 448, col: 51, offset: 19263},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 448, col: 62, offset: 19274},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 450, col: 1, offset: 19285},
			expr: &charClassMatcher{
				pos:        position{line: 450, col: 13, offset: 19297},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 452, col: 1, offset: 19304},
			expr: &choiceExpr{
				pos: position{line: 452, col: 13, offset: 19316},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 452, col: 13, offset: 19316},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 452, col: 22, offset: 19325},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 452, col: 29, offset: 19332},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 454, col: 1, offset: 19338},
			expr: &choiceExpr{
				pos: position{line: 454, col: 13, offset: 19350},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 454, col: 13, offset: 19350},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 454, col: 19, offset: 19356},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 454, col: 19, offset: 19356},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 458, col: 1, offset: 19401},
			expr: &notExpr{
				pos: position{line: 458, col: 13, offset: 19413},
				expr: &anyMatcher{
					line: 458, col: 14, offset: 19414,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 460, col: 1, offset: 19417},
			expr: &choiceExpr{
				pos: position{line: 460, col: 13, offset: 19429},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 460, col: 13, offset: 19429},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 460, col: 23, offset: 19439},
						name: "EOF",
					},
				},
			},
		},
	},
}

func (c *current) onDocument1(frontMatter, documentHeader, blocks interface{}) (interface{}, error) {
	return types.NewDocument(frontMatter, documentHeader, blocks.([]interface{}))
}

func (p *parser) callonDocument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument1(stack["frontMatter"], stack["documentHeader"], stack["blocks"])
}

func (c *current) onDocumentBlocks7(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonDocumentBlocks7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentBlocks7(stack["content"])
}

func (c *current) onPreamble1(elements interface{}) (interface{}, error) {
	return types.NewPreamble(elements.([]interface{}))
}

func (p *parser) callonPreamble1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPreamble1(stack["elements"])
}

func (c *current) onFrontMatter1(content interface{}) (interface{}, error) {
	return types.NewYamlFrontMatter(content.([]interface{}))
}

func (p *parser) callonFrontMatter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrontMatter1(stack["content"])
}

func (c *current) onDocumentHeader1(header, authors, revision, otherAttributes interface{}) (interface{}, error) {

	return types.NewDocumentHeader(header, authors, revision, otherAttributes.([]interface{}))
}

func (p *parser) callonDocumentHeader1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentHeader1(stack["header"], stack["authors"], stack["revision"], stack["otherAttributes"])
}

func (c *current) onDocumentTitle1(attributes, level, content interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonDocumentTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentTitle1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onDocumentAuthorsInlineForm1(authors interface{}) (interface{}, error) {
	return types.NewDocumentAuthors(authors.([]interface{}))
}

func (p *parser) callonDocumentAuthorsInlineForm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthorsInlineForm1(stack["authors"])
}

func (c *current) onDocumentAuthorsAttributeForm1(author interface{}) (interface{}, error) {
	return []*types.DocumentAuthor{author.(*types.DocumentAuthor)}, nil
}

func (p *parser) callonDocumentAuthorsAttributeForm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthorsAttributeForm1(stack["author"])
}

func (c *current) onDocumentAuthor1(namePart1, namePart2, namePart3, email interface{}) (interface{}, error) {
	return types.NewDocumentAuthor(namePart1, namePart2, namePart3, email)
}

func (p *parser) callonDocumentAuthor1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthor1(stack["namePart1"], stack["namePart2"], stack["namePart3"], stack["email"])
}

func (c *current) onDocumentRevision1(revnumber, revdate, revremark interface{}) (interface{}, error) {
	return types.NewDocumentRevision(revnumber, revdate, revremark)
}

func (p *parser) callonDocumentRevision1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentRevision1(stack["revnumber"], stack["revdate"], stack["revremark"])
}

func (c *current) onDocumentAttributeDeclarationWithNameOnly1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeDeclaration(name.([]interface{}), nil)
}

func (p *parser) callonDocumentAttributeDeclarationWithNameOnly1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeDeclarationWithNameOnly1(stack["name"])
}

func (c *current) onDocumentAttributeDeclarationWithNameAndValue1(name, value interface{}) (interface{}, error) {
	return types.NewDocumentAttributeDeclaration(name.([]interface{}), value.([]interface{}))
}

func (p *parser) callonDocumentAttributeDeclarationWithNameAndValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeDeclarationWithNameAndValue1(stack["name"], stack["value"])
}

func (c *current) onDocumentAttributeResetWithSectionTitleBangSymbol1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeReset(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeResetWithSectionTitleBangSymbol1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeResetWithSectionTitleBangSymbol1(stack["name"])
}

func (c *current) onDocumentAttributeResetWithTrailingBangSymbol1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeReset(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeResetWithTrailingBangSymbol1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeResetWithTrailingBangSymbol1(stack["name"])
}

func (c *current) onDocumentAttributeSubstitution1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeSubstitution(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeSubstitution1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeSubstitution1(stack["name"])
}

func (c *current) onSection11(header, elements interface{}) (interface{}, error) {
	return types.NewSection(1, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection11(stack["header"], stack["elements"])
}

func (c *current) onSection1Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection1Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Block1(stack["content"])
}

func (c *current) onSection21(header, elements interface{}) (interface{}, error) {
	return types.NewSection(2, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection21(stack["header"], stack["elements"])
}

func (c *current) onSection2Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection2Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Block1(stack["content"])
}

func (c *current) onSection31(header, elements interface{}) (interface{}, error) {
	return types.NewSection(3, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection31() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection31(stack["header"], stack["elements"])
}

func (c *current) onSection3Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection3Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Block1(stack["content"])
}

func (c *current) onSection41(header, elements interface{}) (interface{}, error) {
	return types.NewSection(4, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection41() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection41(stack["header"], stack["elements"])
}

func (c *current) onSection4Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection4Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Block1(stack["content"])
}

func (c *current) onSection51(header, elements interface{}) (interface{}, error) {
	return types.NewSection(5, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection51() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection51(stack["header"], stack["elements"])
}

func (c *current) onSection5Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection5Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Block1(stack["content"])
}

func (c *current) onSection1Title1(attributes, level, content interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection1Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onSection2Title1(attributes, level, content interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection2Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onSection3Title1(attributes, level, content interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection3Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onSection4Title1(attributes, level, content interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection4Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onSection5Title1(attributes, level, content interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection5Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onList1(attributes, elements interface{}) (interface{}, error) {
	return types.NewList(elements.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonList1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onList1(stack["attributes"], stack["elements"])
}

func (c *current) onListItem1(level, content interface{}) (interface{}, error) {
	return types.NewListItem(level, content.(*types.ListItemContent), nil)
}

func (p *parser) callonListItem1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListItem1(stack["level"], stack["content"])
}

func (c *current) onListItemContent1(lines interface{}) (interface{}, error) {

	return types.NewListItemContent(lines.([]interface{}))
}

func (p *parser) callonListItemContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListItemContent1(stack["lines"])
}

func (c *current) onParagraph1(attributes, lines interface{}) (interface{}, error) {
	return types.NewParagraph(lines.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonParagraph1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraph1(stack["attributes"], stack["lines"])
}

func (c *current) onInlineContent1(elements interface{}) (interface{}, error) {
	// needs an "EOL" but does not consume it here.
	return types.NewInlineContent(elements.([]interface{}))
}

func (p *parser) callonInlineContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineContent1(stack["elements"])
}

func (c *current) onBoldTextSimplePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldTextSimplePunctuation1(stack["content"])
}

func (c *current) onBoldTextDoublePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldTextDoublePunctuation1(stack["content"])
}

func (c *current) onBoldTextUnbalancedPunctuation1(content interface{}) (interface{}, error) {
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"*"}, content.([]interface{}))
	return types.NewQuotedText(types.Bold, result)
}

func (p *parser) callonBoldTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldTextUnbalancedPunctuation1(stack["content"])
}

func (c *current) onEscapedBoldTextSimplePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "*", content.([]interface{}))
}

func (p *parser) callonEscapedBoldTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldTextSimplePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedBoldTextDoublePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "**", content.([]interface{}))
}

func (p *parser) callonEscapedBoldTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldTextDoublePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedBoldTextUnbalancedPunctuation1(backslashes, content interface{}) (interface{}, error) {
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"*"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "*", result)
}

func (p *parser) callonEscapedBoldTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldTextUnbalancedPunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onItalicTextSimplePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicTextSimplePunctuation1(stack["content"])
}

func (c *current) onItalicTextDoublePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicTextDoublePunctuation1(stack["content"])
}

func (c *current) onItalicTextUnbalancedPunctuation1(content interface{}) (interface{}, error) {
	// unbalanced `__` vs `_` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewQuotedText(types.Italic, result)
}

func (p *parser) callonItalicTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicTextUnbalancedPunctuation1(stack["content"])
}

func (c *current) onEscapedItalicTextSimplePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "_", content.([]interface{}))
}

func (p *parser) callonEscapedItalicTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicTextSimplePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedItalicTextDoublePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "__", content.([]interface{}))
}

func (p *parser) callonEscapedItalicTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicTextDoublePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedItalicTextUnbalancedPunctuation1(backslashes, content interface{}) (interface{}, error) {
	// unbalanced `__` vs `_` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "_", result)
}

func (p *parser) callonEscapedItalicTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicTextUnbalancedPunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onMonospaceTextSimplePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceTextSimplePunctuation1(stack["content"])
}

func (c *current) onMonospaceTextDoublePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceTextDoublePunctuation1(stack["content"])
}

func (c *current) onMonospaceTextUnbalancedPunctuation1(content interface{}) (interface{}, error) {
	// unbalanced "``" vs "`" punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewQuotedText(types.Monospace, result)
}

func (p *parser) callonMonospaceTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceTextUnbalancedPunctuation1(stack["content"])
}

func (c *current) onEscapedMonospaceTextSimplePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "`", content.([]interface{}))
}

func (p *parser) callonEscapedMonospaceTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceTextSimplePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedMonospaceTextDoublePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "``", content.([]interface{}))
}

func (p *parser) callonEscapedMonospaceTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceTextDoublePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedMonospaceTextUnbalancedPunctuation1(backslashes, content interface{}) (interface{}, error) {
	// unbalanced "``" vs "`" punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "`", result)
}

func (p *parser) callonEscapedMonospaceTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceTextUnbalancedPunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onWordWithQuotePunctuation1() (interface{}, error) {
	// can have "*", "_" or "`" within, maybe because the user inserted another quote, or made an error (extra or missing space, for example)
	return c.text, nil
}

func (p *parser) callonWordWithQuotePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWordWithQuotePunctuation1()
}

func (c *current) onExternalLink1(url, text interface{}) (interface{}, error) {
	if text != nil {
		return types.NewExternalLink(url.([]interface{}), text.([]interface{}))
	}
	return types.NewExternalLink(url.([]interface{}), nil)
}

func (p *parser) callonExternalLink1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternalLink1(stack["url"], stack["text"])
}

func (c *current) onBlockImage1(attributes, image interface{}) (interface{}, error) {
	// here we can ignore the blank line in the returned element
	return types.NewBlockImage(*image.(*types.ImageMacro), attributes.([]interface{}))
}

func (p *parser) callonBlockImage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlockImage1(stack["attributes"], stack["image"])
}

func (c *current) onBlockImageMacro1(path, attributes interface{}) (interface{}, error) {
	return types.NewImageMacro(path.(string), attributes)
}

func (p *parser) callonBlockImageMacro1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlockImageMacro1(stack["path"], stack["attributes"])
}

func (c *current) onInlineImage1(image interface{}) (interface{}, error) {
	// here we can ignore the blank line in the returned element
	return types.NewInlineImage(*image.(*types.ImageMacro))
}

func (p *parser) callonInlineImage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineImage1(stack["image"])
}

func (c *current) onInlineImageMacro1(path, attributes interface{}) (interface{}, error) {
	return types.NewImageMacro(path.(string), attributes)
}

func (p *parser) callonInlineImageMacro1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineImageMacro1(stack["path"], stack["attributes"])
}

func (c *current) onFencedBlock1(content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.FencedBlock, content.([]interface{}))
}

func (p *parser) callonFencedBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFencedBlock1(stack["content"])
}

func (c *current) onParagraphWithSpaces1(spaces, content interface{}) (interface{}, error) {
	return types.NewLiteralBlock(spaces.([]interface{}), content.([]interface{}))
}

func (p *parser) callonParagraphWithSpaces1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithSpaces1(stack["spaces"], stack["content"])
}

func (c *current) onLiteralBlockContent1(content interface{}) (interface{}, error) {

	return content, nil
}

func (p *parser) callonLiteralBlockContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLiteralBlockContent1(stack["content"])
}

func (c *current) onParagraphWithLiteralBlockDelimiter1(content interface{}) (interface{}, error) {
	return types.NewLiteralBlock([]interface{}{}, content.([]interface{}))
}

func (p *parser) callonParagraphWithLiteralBlockDelimiter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithLiteralBlockDelimiter1(stack["content"])
}

func (c *current) onParagraphWithLiteralAttribute1(content interface{}) (interface{}, error) {
	return types.NewLiteralBlock([]interface{}{}, content.([]interface{}))
}

func (p *parser) callonParagraphWithLiteralAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithLiteralAttribute1(stack["content"])
}

func (c *current) onElementLink1(path interface{}) (interface{}, error) {
	return types.NewElementLink(path.(string))
}

func (p *parser) callonElementLink1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementLink1(stack["path"])
}

func (c *current) onElementID1(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonElementID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementID1(stack["id"])
}

func (c *current) onElementTitle1(title interface{}) (interface{}, error) {
	return types.NewElementTitle(title.([]interface{}))
}

func (p *parser) callonElementTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementTitle1(stack["title"])
}

func (c *current) onWord1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonWord1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWord1()
}

func (c *current) onBlankLine1() (interface{}, error) {
	return types.NewBlankLine()
}

func (p *parser) callonBlankLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlankLine1()
}

func (c *current) onURL1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonURL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onURL1()
}

func (c *current) onID1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onID1()
}

func (c *current) onURL_TEXT1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonURL_TEXT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onURL_TEXT1()
}

func (c *current) onWS3() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonWS3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWS3()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEntrypoint is returned when the specified entrypoint rule
	// does not exit.
	errInvalidEntrypoint = errors.New("invalid entrypoint")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errMaxExprCnt is used to signal that the maximum number of
	// expressions have been parsed.
	errMaxExprCnt = errors.New("max number of expresssions parsed")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// MaxExpressions creates an Option to stop parsing after the provided
// number of expressions have been parsed, if the value is 0 then the parser will
// parse for as many steps as needed (possibly an infinite number).
//
// The default for maxExprCnt is 0.
func MaxExpressions(maxExprCnt uint64) Option {
	return func(p *parser) Option {
		oldMaxExprCnt := p.maxExprCnt
		p.maxExprCnt = maxExprCnt
		return MaxExpressions(oldMaxExprCnt)
	}
}

// Entrypoint creates an Option to set the rule name to use as entrypoint.
// The rule name must have been specified in the -alternate-entrypoints
// if generating the parser with the -optimize-grammar flag, otherwise
// it may have been optimized out. Passing an empty string sets the
// entrypoint to the first rule in the grammar.
//
// The default is to start parsing at the first rule in the grammar.
func Entrypoint(ruleName string) Option {
	return func(p *parser) Option {
		oldEntrypoint := p.entrypoint
		p.entrypoint = ruleName
		if ruleName == "" {
			p.entrypoint = g.rules[0].name
		}
		return Entrypoint(oldEntrypoint)
	}
}

// Statistics adds a user provided Stats struct to the parser to allow
// the user to process the results after the parsing has finished.
// Also the key for the "no match" counter is set.
//
// Example usage:
//
//     input := "input"
//     stats := Stats{}
//     _, err := Parse("input-file", []byte(input), Statistics(&stats, "no match"))
//     if err != nil {
//         log.Panicln(err)
//     }
//     b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
//     if err != nil {
//         log.Panicln(err)
//     }
//     fmt.Println(string(b))
//
func Statistics(stats *Stats, choiceNoMatch string) Option {
	return func(p *parser) Option {
		oldStats := p.Stats
		p.Stats = stats
		oldChoiceNoMatch := p.choiceNoMatch
		p.choiceNoMatch = choiceNoMatch
		if p.Stats.ChoiceAltCnt == nil {
			p.Stats.ChoiceAltCnt = make(map[string]map[string]int)
		}
		return Statistics(oldStats, oldChoiceNoMatch)
	}
}

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// GlobalStore creates an Option to set a key to a certain value in
// the globalStore.
func GlobalStore(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.globalStore[key]
		p.cur.globalStore[key] = value
		return GlobalStore(key, old)
	}
}

// InitState creates an Option to set a key to a certain value in
// the global "state" store.
func InitState(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.state[key]
		p.cur.state[key] = value
		return InitState(key, old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (i interface{}, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match

	// state is a store for arbitrary key,value pairs that the user wants to be
	// tied to the backtracking of the parser.
	// This is always rolled back if a parsing rule fails.
	state storeDict

	// globalStore is a general store for the user to store arbitrary key-value
	// pairs that they need to manage and that they do not want tied to the
	// backtracking of the parser. This is only modified by the user and never
	// rolled back by the parser. It is always up to the user to keep this in a
	// consistent state.
	globalStore storeDict
}

type storeDict map[string]interface{}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type recoveryExpr struct {
	pos          position
	expr         interface{}
	recoverExpr  interface{}
	failureLabel []string
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type throwExpr struct {
	pos   position
	label string
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type stateCodeExpr struct {
	pos position
	run func(*parser) error
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos             position
	val             string
	basicLatinChars [128]bool
	chars           []rune
	ranges          []rune
	classes         []*unicode.RangeTable
	ignoreCase      bool
	inverted        bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner    error
	pos      position
	prefix   string
	expected []string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	stats := Stats{
		ChoiceAltCnt: make(map[string]map[string]int),
	}

	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
		cur: current{
			state:       make(storeDict),
			globalStore: make(storeDict),
		},
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make([]string, 0, 20),
		Stats:           &stats,
		// start rule is rule [0] unless an alternate entrypoint is specified
		entrypoint: g.rules[0].name,
		emptyState: make(storeDict),
	}
	p.setOptions(opts)

	if p.maxExprCnt == 0 {
		p.maxExprCnt = math.MaxUint64
	}

	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

const choiceNoMatch = -1

// Stats stores some statistics, gathered during parsing
type Stats struct {
	// ExprCnt counts the number of expressions processed during parsing
	// This value is compared to the maximum number of expressions allowed
	// (set by the MaxExpressions option).
	ExprCnt uint64

	// ChoiceAltCnt is used to count for each ordered choice expression,
	// which alternative is used how may times.
	// These numbers allow to optimize the order of the ordered choice expression
	// to increase the performance of the parser
	//
	// The outer key of ChoiceAltCnt is composed of the name of the rule as well
	// as the line and the column of the ordered choice.
	// The inner key of ChoiceAltCnt is the number (one-based) of the matching alternative.
	// For each alternative the number of matches are counted. If an ordered choice does not
	// match, a special counter is incremented. The name of this counter is set with
	// the parser option Statistics.
	// For an alternative to be included in ChoiceAltCnt, it has to match at least once.
	ChoiceAltCnt map[string]map[string]int
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	depth   int
	recover bool
	debug   bool

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// parse fail
	maxFailPos            position
	maxFailExpected       []string
	maxFailInvertExpected bool

	// max number of expressions to be parsed
	maxExprCnt uint64
	// entrypoint for the parser
	entrypoint string

	*Stats

	choiceNoMatch string
	// recovery expression stack, keeps track of the currently available recovery expression, these are traversed in reverse
	recoveryStack []map[string]interface{}

	// emptyState contains an empty storeDict, which is used to optimize cloneState if global "state" store is not used.
	emptyState storeDict
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

// push a recovery expression with its labels to the recoveryStack
func (p *parser) pushRecovery(labels []string, expr interface{}) {
	if cap(p.recoveryStack) == len(p.recoveryStack) {
		// create new empty slot in the stack
		p.recoveryStack = append(p.recoveryStack, nil)
	} else {
		// slice to 1 more
		p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)+1]
	}

	m := make(map[string]interface{}, len(labels))
	for _, fl := range labels {
		m[fl] = expr
	}
	p.recoveryStack[len(p.recoveryStack)-1] = m
}

// pop a recovery expression from the recoveryStack
func (p *parser) popRecovery() {
	// GC that map
	p.recoveryStack[len(p.recoveryStack)-1] = nil

	p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position, []string{})
}

func (p *parser) addErrAt(err error, pos position, expected []string) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String(), expected: expected}
	p.errs.add(pe)
}

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if fail == p.maxFailInvertExpected {
		if pos.offset < p.maxFailPos.offset {
			return
		}

		if pos.offset > p.maxFailPos.offset {
			p.maxFailPos = pos
			p.maxFailExpected = p.maxFailExpected[:0]
		}

		if p.maxFailInvertExpected {
			want = "!" + want
		}
		p.maxFailExpected = append(p.maxFailExpected, want)
	}
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError && n == 1 { // see utf8.DecodeRune
		p.addErr(errInvalidEncoding)
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// Cloner is implemented by any value that has a Clone method, which returns a
// copy of the value. This is mainly used for types which are not passed by
// value (e.g map, slice, chan) or structs that contain such types.
//
// This is used in conjunction with the global state feature to create proper
// copies of the state to allow the parser to properly restore the state in
// the case of backtracking.
type Cloner interface {
	Clone() interface{}
}

// clone and return parser current state.
func (p *parser) cloneState() storeDict {
	if p.debug {
		defer p.out(p.in("cloneState"))
	}

	if len(p.cur.state) == 0 {
		return p.emptyState
	}

	state := make(storeDict, len(p.cur.state))
	for k, v := range p.cur.state {
		if c, ok := v.(Cloner); ok {
			state[k] = c.Clone()
		} else {
			state[k] = v
		}
	}
	return state
}

// restore parser current state to the state storeDict.
// every restoreState should applied only one time for every cloned state
func (p *parser) restoreState(state storeDict) {
	if p.debug {
		defer p.out(p.in("restoreState"))
	}
	p.cur.state = state
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	startRule, ok := p.rules[p.entrypoint]
	if !ok {
		p.addErr(errInvalidEntrypoint)
		return nil, p.errs.err()
	}

	p.read() // advance to first rune
	val, ok = p.parseRule(startRule)
	if !ok {
		if len(*p.errs) == 0 {
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			maxFailExpectedMap := make(map[string]struct{}, len(p.maxFailExpected))
			for _, v := range p.maxFailExpected {
				maxFailExpectedMap[v] = struct{}{}
			}
			expected := make([]string, 0, len(maxFailExpectedMap))
			eof := false
			if _, ok := maxFailExpectedMap["!."]; ok {
				delete(maxFailExpectedMap, "!.")
				eof = true
			}
			for k := range maxFailExpectedMap {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}

		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func listJoin(list []string, sep string, lastSep string) string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return fmt.Sprintf("%s %s %s", strings.Join(list[:len(list)-1], sep), lastSep, list[len(list)-1])
	}
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.ExprCnt++
	if p.ExprCnt > p.maxExprCnt {
		panic(errMaxExprCnt)
	}

	var val interface{}
	var ok bool
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *recoveryExpr:
		val, ok = p.parseRecoveryExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *stateCodeExpr:
		val, ok = p.parseStateCodeExpr(expr)
	case *throwExpr:
		val, ok = p.parseThrowExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		state := p.cloneState()
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position, []string{})
		}
		p.restoreState(state)

		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	state := p.cloneState()

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError || p.pt.w > 1 { // see utf8.DecodeRune
		start := p.pt
		p.read()
		p.failAt(true, start.position, ".")
		return p.sliceFrom(start), true
	}
	p.failAt(false, p.pt.position, ".")
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt

	// can't match EOF
	if cur == utf8.RuneError && p.pt.w == 0 { // see utf8.DecodeRune
		p.failAt(false, start.position, chr.val)
		return nil, false
	}

	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		p.failAt(true, start.position, chr.val)
		return p.sliceFrom(start), true
	}
	p.failAt(false, start.position, chr.val)
	return nil, false
}

func (p *parser) incChoiceAltCnt(ch *choiceExpr, altI int) {
	choiceIdent := fmt.Sprintf("%s %d:%d", p.rstack[len(p.rstack)-1].name, ch.pos.line, ch.pos.col)
	m := p.ChoiceAltCnt[choiceIdent]
	if m == nil {
		m = make(map[string]int)
		p.ChoiceAltCnt[choiceIdent] = m
	}
	// We increment altI by 1, so the keys do not start at 0
	alt := strconv.Itoa(altI + 1)
	if altI == choiceNoMatch {
		alt = p.choiceNoMatch
	}
	m[alt]++
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for altI, alt := range ch.alternatives {
		// dummy assignment to prevent compile error if optimized
		_ = altI

		state := p.cloneState()
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			p.incChoiceAltCnt(ch, altI)
			return val, ok
		}
		p.restoreState(state)
	}
	p.incChoiceAltCnt(ch, choiceNoMatch)
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	ignoreCase := ""
	if lit.ignoreCase {
		ignoreCase = "i"
	}
	val := fmt.Sprintf("%q%s", lit.val, ignoreCase)
	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.failAt(false, start.position, val)
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	p.failAt(true, start.position, val)
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	state := p.cloneState()

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRecoveryExpr(recover *recoveryExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRecoveryExpr (" + strings.Join(recover.failureLabel, ",") + ")"))
	}

	p.pushRecovery(recover.failureLabel, recover.recoverExpr)
	val, ok := p.parseExpr(recover.expr)
	p.popRecovery()

	return val, ok
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	vals := make([]interface{}, 0, len(seq.exprs))

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseStateCodeExpr(state *stateCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseStateCodeExpr"))
	}

	err := state.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, true
}

func (p *parser) parseThrowExpr(expr *throwExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseThrowExpr"))
	}

	for i := len(p.recoveryStack) - 1; i >= 0; i-- {
		if recoverExpr, ok := p.recoveryStack[i][expr.label]; ok {
			if val, ok := p.parseExpr(recoverExpr); ok {
				return val, ok
			}
		}
	}

	return nil, false
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}