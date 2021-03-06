package parser_test

import (
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/require"
)

var _ = Describe("file location", func() {

	DescribeTable("'FileLocation' pattern",
		func(filename string, expected interface{}) {
			reader := strings.NewReader(filename)
			actual, err := parser.ParseReader(filename, reader, parser.Entrypoint("FileLocation"))
			require.NoError(GinkgoT(), err)
			GinkgoT().Log("actual result: %s", spew.Sdump(actual))
			GinkgoT().Log("expected result: %s", spew.Sdump(expected))
			assert.Equal(GinkgoT(), expected, actual)
		},
		Entry("'chapter'", "chapter", types.Location{
			&types.StringElement{
				Content: "chapter",
			},
		}),
		Entry("'chapter.adoc'", "chapter.adoc", types.Location{
			&types.StringElement{
				Content: "chapter.adoc",
			},
		}),
		Entry("'chapter-a.adoc'", "chapter-a.adoc", types.Location{
			&types.StringElement{
				Content: "chapter-a.adoc",
			},
		}),
		Entry("'chapter_a.adoc'", "chapter_a.adoc", types.Location{
			&types.StringElement{
				Content: "chapter_a.adoc",
			},
		}),
		Entry("'includes/chapter_a.adoc'", "includes/chapter_a.adoc", types.Location{
			&types.StringElement{
				Content: "includes/chapter_a.adoc",
			},
		}),
		Entry("'chapter-{foo}.adoc'", "chapter-{foo}.adoc", types.Location{
			&types.StringElement{
				Content: "chapter-",
			},
			&types.DocumentAttributeSubstitution{
				Name: "foo",
			},
			&types.StringElement{
				Content: ".adoc",
			},
		}),
		Entry("'{includedir}/chapter-{foo}.adoc'", "{includedir}/chapter-{foo}.adoc", types.Location{
			&types.DocumentAttributeSubstitution{
				Name: "includedir",
			},
			&types.StringElement{
				Content: "/chapter-",
			},
			&types.DocumentAttributeSubstitution{
				Name: "foo",
			},
			&types.StringElement{
				Content: ".adoc",
			},
		}),
	)
})

var _ = Describe("file inclusions - preflight", func() {

	It("should include adoc file without leveloffset", func() {
		source := "include::includes/chapter-a.adoc[]"
		expected := &types.PreflightDocument{
			Blocks: []interface{}{
				&types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID:       "chapter_a",
						types.AttrCustomID: false,
					},
					Level: 0,
					Title: types.InlineElements{
						&types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
				},
				&types.BlankLine{},
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "content",
							},
						},
					},
				},
			},
		}
		verifyPreflight(expected, source)
	})

	It("should include adoc file with leveloffset", func() {
		source := "include::includes/chapter-a.adoc[leveloffset=+1]"
		expected := &types.PreflightDocument{
			Blocks: []interface{}{
				&types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID:       "chapter_a",
						types.AttrCustomID: false,
					},
					Level: 1,
					Title: types.InlineElements{
						&types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
				},
				&types.BlankLine{},
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "content",
							},
						},
					},
				},
			},
		}
		verifyPreflight(expected, source)
	})

	Context("file inclusion in delimited blocks", func() {

		It("should include adoc file within fenced block", func() {
			source := "```\n" +
				"include::includes/chapter-a.adoc[]\n" +
				"```"
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Fenced,
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "= Chapter A",
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("should include adoc file within listing block", func() {
			source := `----
include::includes/chapter-a.adoc[]
----`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "= Chapter A",
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("should include adoc file within example block", func() {
			source := `====
include::includes/chapter-a.adoc[]
====`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Example,
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "= Chapter A",
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("should include adoc file within quote block", func() {
			source := `____
include::includes/chapter-a.adoc[]
____`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Quote,
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "= Chapter A",
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("should include adoc file within verse block", func() {
			source := `[verse]
____
include::includes/chapter-a.adoc[]
____`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Verse,
						},
						Kind: types.Verse,
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "= Chapter A",
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("should include adoc file within sidebar block", func() {
			source := `****
include::includes/chapter-a.adoc[]
****`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Sidebar,
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "= Chapter A",
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("should include adoc file within passthrough block", func() {
			Skip("missing support for passthrough blocks")
			source := `++++
include::includes/chapter-a.adoc[]
++++`
			expected := &types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				// Kind:       types.Passthrough,
				Elements: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "= Chapter A",
								},
							},
						},
					},
					&types.BlankLine{},
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "content",
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})
	})

	Context("file inclusion with line ranges", func() {

		Context("file inclusion with unquoted line ranges", func() {

			It("file inclusion with single unquoted line", func() {
				source := `include::includes/chapter-a.adoc[lines=1]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Level: 0,
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				verifyPreflight(expected, source)
			})

			It("file inclusion with multiple unquoted lines", func() {
				source := `include::includes/chapter-a.adoc[lines=1..2]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						&types.BlankLine{},
					},
				}
				verifyPreflight(expected, source)
			})

			It("file inclusion with multiple unquoted ranges", func() {
				source := `include::includes/chapter-a.adoc[lines=1;3..4;6..-1]` // paragraph becomes the author since the in-between blank line is stripped out
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "content",
									},
								},
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				verifyPreflight(expected, source)
			})

			It("file inclusion with invalid unquoted range - case 1", func() {
				source := `include::includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						&types.BlankLine{},
						&types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									&types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyPreflight(expected, source)
			})

			It("file inclusion with invalid unquoted range - case 2", func() {
				source := `include::includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				verifyPreflight(expected, source)
			})
		})

		Context("file inclusion with quoted line ranges", func() {

			It("file inclusion with single quoted line", func() {
				source := `include::includes/chapter-a.adoc[lines="1"]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				verifyPreflight(expected, source)
			})

			It("file inclusion with multiple quoted lines", func() {
				source := `include::includes/chapter-a.adoc[lines="1..2"]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						&types.BlankLine{},
					},
				}
				verifyPreflight(expected, source)
			})

			It("file inclusion with multiple quoted ranges", func() {
				// here, the `content` paragraph gets attached to the header and becomes the author
				source := `include::includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "content",
									},
								},
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				verifyPreflight(expected, source)
			})

			It("file inclusion with invalid quoted range - case 1", func() {
				source := `include::includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						&types.BlankLine{},
						&types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									&types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyPreflight(expected, source)
			})

			It("file inclusion with invalid quoted range - case 2", func() {
				source := `include::includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								&types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						&types.BlankLine{},
						&types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									&types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyPreflight(expected, source)
			})
		})
	})

	Context("missing file to include", func() {

		It("should replace with string element if directory does not exist in standalone block", func() {
			source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]",
								},
							},
						},
					},
				},
			}
			// TODO: also verify that an error was reported in the console.
			verifyPreflight(expected, source)
		})

		It("should replace with string element if file is missing in standalone block", func() {
			source := `include::includes/unknown.adoc[leveloffset=+1]`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "Unresolved directive in test.adoc - include::includes/unknown.adoc[leveloffset=+1]",
								},
							},
						},
					},
				},
			}
			// TODO: also verify that an error was reported in the console.
			verifyPreflight(expected, source)
		})

		It("should replace with string element if file is missing in delimited block", func() {
			source := `----
include::includes/unknown.adoc[leveloffset=+1]
----`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "Unresolved directive in test.adoc - include::includes/unknown.adoc[leveloffset=+1]",
										},
									},
								},
							},
						},
					},
				},
			}
			// TODO: also verify that an error was reported in the console.
			verifyPreflight(expected, source)
		})
	})

	Context("inclusion with attribute in path", func() {

		It("should resolve path with attribute in standalone block", func() {
			source := `:includedir: ./includes
			
include::{includedir}/grandchild-include.adoc[]`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DocumentAttributeDeclaration{
						Name:  "includedir",
						Value: "./includes",
					},
					&types.BlankLine{},
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "first line of grandchild",
								},
							},
						},
					},
					&types.BlankLine{},
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "last line of grandchild",
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("should resolve path with attribute in delimited block", func() {
			source := `:includedir: ./includes

----
include::{includedir}/grandchild-include.adoc[]
----`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DocumentAttributeDeclaration{
						Name:  "includedir",
						Value: "./includes",
					},
					&types.BlankLine{},
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "first line of grandchild",
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "last line of grandchild",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})
	})

	Context("inclusion of non-asciidoc file", func() {

		It("include go file without line range", func() {

			source := `----
include::includes/hello_world.go[] 
----`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Kind:       types.Listing,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: `package includes`,
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: `import "fmt"`,
										},
									},
								},
							},
							&types.BlankLine{},
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: `func helloworld() {`,
										},
									},
									{
										&types.StringElement{
											Content: `	fmt.Println("hello, world!")`,
										},
									},
									{
										&types.StringElement{
											Content: `}`,
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})
	})
})

var _ = Describe("file inclusions - ignored at preflight", func() {

	It("should include adoc file with leveloffset attribute", func() {
		source := "include::includes/chapter-a.adoc[leveloffset=+1]"
		expected := &types.PreflightDocument{
			Blocks: []interface{}{
				&types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLevelOffset: "+1",
					},
					Location: types.Location{
						&types.StringElement{
							Content: "includes/chapter-a.adoc",
						},
					},
					RawText: `include::includes/chapter-a.adoc[leveloffset=+1]`,
				},
			},
		}
		verifyPreflightWithoutPreprocessing(expected, source)
	})

	Context("file inclusion in delimited blocks", func() {

		It("should include adoc file within fenced block", func() {
			source := "```\n" +
				"include::includes/chapter-a.adoc[]\n" +
				"```"
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Fenced,
						Elements: []interface{}{
							&types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									&types.StringElement{
										Content: "includes/chapter-a.adoc",
									},
								},
								RawText: `include::includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			verifyPreflightWithoutPreprocessing(expected, source)
		})

		It("should include adoc file within listing block", func() {
			source := `----
include::includes/chapter-a.adoc[]
----`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							&types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									&types.StringElement{
										Content: "includes/chapter-a.adoc",
									},
								},
								RawText: `include::includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			verifyPreflightWithoutPreprocessing(expected, source)
		})

		It("should include adoc file within example block", func() {
			source := `====
include::includes/chapter-a.adoc[]
====`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Example,
						Elements: []interface{}{
							&types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									&types.StringElement{
										Content: "includes/chapter-a.adoc",
									},
								},
								RawText: `include::includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			verifyPreflightWithoutPreprocessing(expected, source)
		})

		It("should include adoc file within quote block", func() {
			source := `____
include::includes/chapter-a.adoc[]
____`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Quote,
						Elements: []interface{}{
							&types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									&types.StringElement{
										Content: "includes/chapter-a.adoc",
									},
								},
								RawText: `include::includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			verifyPreflightWithoutPreprocessing(expected, source)
		})

		It("should include adoc file within verse block", func() {
			source := `[verse]
____
include::includes/chapter-a.adoc[]
____`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Verse,
						},
						Kind: types.Verse,
						Elements: []interface{}{
							&types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									&types.StringElement{
										Content: "includes/chapter-a.adoc",
									},
								},
								RawText: `include::includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			verifyPreflightWithoutPreprocessing(expected, source)
		})

		It("should include adoc file within sidebar block", func() {
			source := `****
include::includes/chapter-a.adoc[]
****`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Sidebar,
						Elements: []interface{}{
							&types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									&types.StringElement{
										Content: "includes/chapter-a.adoc",
									},
								},
								RawText: `include::includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			verifyPreflightWithoutPreprocessing(expected, source)
		})

		It("should include adoc file within passthrough block", func() {
			Skip("missing support for passthrough blocks")
			source := `++++
include::includes/chapter-a.adoc[]
++++`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						// Kind:       types.Passthrough,
						Elements: []interface{}{
							&types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									&types.StringElement{
										Content: "includes/chapter-a.adoc",
									},
								},
								RawText: `include::includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			verifyPreflightWithoutPreprocessing(expected, source)
		})
	})

	Context("file inclusion with line ranges", func() {

		Context("file inclusion with unquoted line ranges", func() {

			It("file inclusion with single unquoted line", func() {
				source := `include::includes/chapter-a.adoc[lines=1]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{Start: 1, End: 1},
								},
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines=1]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with multiple unquoted lines", func() {
				source := `include::includes/chapter-a.adoc[lines=1..2]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{Start: 1, End: 2},
								},
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines=1..2]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with multiple unquoted ranges", func() {
				source := `include::includes/chapter-a.adoc[lines=1;3..4;6..-1]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{Start: 1, End: 1},
									{Start: 3, End: 4},
									{Start: 6, End: -1},
								},
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines=1;3..4;6..-1]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with invalid unquoted range - case 1", func() {
				source := `include::includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `1;3..4;6..foo`,
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines=1;3..4;6..foo]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with invalid unquoted range - case 2", func() {
				source := `include::includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{Start: 1, End: 1},
								},
								"3..4":  nil,
								"6..-1": nil,
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines=1,3..4,6..-1]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with invalid unquoted range - case 3", func() {
				source := `include::includes/chapter-a.adoc[lines=foo]` // using commas instead of semi-colons
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: "foo",
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines=foo]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})
		})

		Context("file inclusion with quoted line ranges", func() {

			It("file inclusion with single quoted line", func() {
				source := `include::includes/chapter-a.adoc[lines="1"]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{Start: 1, End: 1},
								},
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines="1"]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with multiple quoted lines", func() {
				source := `include::includes/chapter-a.adoc[lines="1..2"]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{Start: 1, End: 2},
								},
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines="1..2"]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with multiple quoted ranges", func() {
				source := `include::includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{Start: 1, End: 1},
									{Start: 3, End: 4},
									{Start: 6, End: -1},
								},
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines="1,3..4,6..-1"]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with invalid quoted range - case 1", func() {
				source := `include::includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `"1`, // viewed as a string
								"3..4":               nil,
								"6..foo":             nil,
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines="1,3..4,6..foo"]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})

			It("file inclusion with invalid quoted range - case 2", func() {
				source := `include::includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
				expected := &types.PreflightDocument{
					Blocks: []interface{}{
						&types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `"1;3..4;6..10"`,
							},
							Location: types.Location{
								&types.StringElement{
									Content: "includes/chapter-a.adoc",
								},
							},
							RawText: `include::includes/chapter-a.adoc[lines="1;3..4;6..10"]`,
						},
					},
				}
				verifyPreflightWithoutPreprocessing(expected, source)
			})
		})
	})

})

var _ = Describe("full document with file inclusions", func() {

	// TODO: verify that section in `chapter-a.adoc` in delimited blocks
	// becomes a paragraph in delimited blocks.
	// also, verify that quoted text are kept as-is or turn into string elements
	// (using their raw content)

})
