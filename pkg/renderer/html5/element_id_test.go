package html5

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/stretchr/testify/assert"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("element ID generation", func() {

	It("should generate ID with default prefix", func() {
		// given
		ctx := &renderer.Context{
			Document: &types.Document{
				Attributes: types.DocumentAttributes{},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrID:       "foo",
			types.AttrCustomID: false,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		assert.Equal(GinkgoT(), "_foo", result)
	})

	It("should generate ID with custom prefix", func() {
		// given
		ctx := &renderer.Context{
			Document: &types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "id#",
				},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrID:       "foo",
			types.AttrCustomID: false,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		assert.Equal(GinkgoT(), "id#foo", result)
	})

	It("should generate custom ID", func() {
		// given
		ctx := &renderer.Context{
			Document: &types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "id#",
				},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrID:       "foo",
			types.AttrCustomID: true,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		assert.Equal(GinkgoT(), "foo", result)
	})

	It("should generate empty ID from empty value", func() {
		// given
		ctx := &renderer.Context{
			Document: &types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "id#",
				},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrID:       "",
			types.AttrCustomID: false,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		assert.Equal(GinkgoT(), "", result)
	})

	It("should generate empty ID from missing value", func() {
		// given
		ctx := &renderer.Context{
			Document: &types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "id#",
				},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrCustomID: false,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		assert.Equal(GinkgoT(), "", result)
	})
})
