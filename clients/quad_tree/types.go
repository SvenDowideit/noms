// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef = __mainPackageInFile_types_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __mainPackageInFile_types_Ref() ref.Ref {
	p := types.PackageDef{
		NamedTypes: types.MapOfStringToTypeRefDef{

			"Geoposition": types.MakeStructTypeRef("Geoposition",
				[]types.Field{
					types.Field{"Latitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
					types.Field{"Longitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
				},
				types.Choices{},
			),
			"Georectangle": types.MakeStructTypeRef("Georectangle",
				[]types.Field{
					types.Field{"TopLeft", types.MakeTypeRef("Geoposition", ref.Ref{}), false},
					types.Field{"BottomRight", types.MakeTypeRef("Geoposition", ref.Ref{}), false},
				},
				types.Choices{},
			),
			"Node": types.MakeStructTypeRef("Node",
				[]types.Field{
					types.Field{"Geoposition", types.MakeTypeRef("Geoposition", ref.Ref{}), false},
					types.Field{"Reference", types.MakeCompoundTypeRef("", types.RefKind, types.MakePrimitiveTypeRef(types.ValueKind)), false},
				},
				types.Choices{},
			),
			"QuadTree": types.MakeStructTypeRef("QuadTree",
				[]types.Field{
					types.Field{"Nodes", types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef("Node", ref.Ref{})), false},
					types.Field{"Tiles", types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef("QuadTree", ref.Ref{})), false},
					types.Field{"Depth", types.MakePrimitiveTypeRef(types.UInt8Kind), false},
					types.Field{"NumDescendents", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
					types.Field{"Path", types.MakePrimitiveTypeRef(types.StringKind), false},
					types.Field{"Georectangle", types.MakeTypeRef("Georectangle", ref.Ref{}), false},
				},
				types.Choices{},
			),
			"SQuadTree": types.MakeStructTypeRef("SQuadTree",
				[]types.Field{
					types.Field{"Nodes", types.MakeCompoundTypeRef("", types.ListKind, types.MakeCompoundTypeRef("", types.RefKind, types.MakePrimitiveTypeRef(types.ValueKind))), false},
					types.Field{"Tiles", types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef("", types.RefKind, types.MakeTypeRef("SQuadTree", ref.Ref{}))), false},
					types.Field{"Depth", types.MakePrimitiveTypeRef(types.UInt8Kind), false},
					types.Field{"NumDescendents", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
					types.Field{"Path", types.MakePrimitiveTypeRef(types.StringKind), false},
					types.Field{"Georectangle", types.MakeTypeRef("Georectangle", ref.Ref{}), false},
				},
				types.Choices{},
			),
		},
	}.New()
	return types.RegisterPackage(&p)
}

// Geoposition

type Geoposition struct {
	m types.Map
}

func NewGeoposition() Geoposition {
	return Geoposition{types.NewMap(
		types.NewString("$name"), types.NewString("Geoposition"),
		types.NewString("$type"), types.MakeTypeRef("Geoposition", __mainPackageInFile_types_CachedRef),
		types.NewString("Latitude"), types.Float32(0),
		types.NewString("Longitude"), types.Float32(0),
	)}
}

type GeopositionDef struct {
	Latitude  float32
	Longitude float32
}

func (def GeopositionDef) New() Geoposition {
	return Geoposition{
		types.NewMap(
			types.NewString("$name"), types.NewString("Geoposition"),
			types.NewString("$type"), types.MakeTypeRef("Geoposition", __mainPackageInFile_types_CachedRef),
			types.NewString("Latitude"), types.Float32(def.Latitude),
			types.NewString("Longitude"), types.Float32(def.Longitude),
		)}
}

func (s Geoposition) Def() (d GeopositionDef) {
	d.Latitude = float32(s.m.Get(types.NewString("Latitude")).(types.Float32))
	d.Longitude = float32(s.m.Get(types.NewString("Longitude")).(types.Float32))
	return
}

var __typeRefForGeoposition = types.MakeTypeRef("Geoposition", __mainPackageInFile_types_CachedRef)

func (m Geoposition) TypeRef() types.TypeRef {
	return __typeRefForGeoposition
}

func init() {
	types.RegisterFromValFunction(__typeRefForGeoposition, func(v types.Value) types.NomsValue {
		return GeopositionFromVal(v)
	})
}

func GeopositionFromVal(val types.Value) Geoposition {
	// TODO: Validate here
	return Geoposition{val.(types.Map)}
}

func (s Geoposition) NomsValue() types.Value {
	return s.m
}

func (s Geoposition) Equals(other types.Value) bool {
	if other, ok := other.(Geoposition); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s Geoposition) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Geoposition) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s Geoposition) Latitude() float32 {
	return float32(s.m.Get(types.NewString("Latitude")).(types.Float32))
}

func (s Geoposition) SetLatitude(val float32) Geoposition {
	return Geoposition{s.m.Set(types.NewString("Latitude"), types.Float32(val))}
}

func (s Geoposition) Longitude() float32 {
	return float32(s.m.Get(types.NewString("Longitude")).(types.Float32))
}

func (s Geoposition) SetLongitude(val float32) Geoposition {
	return Geoposition{s.m.Set(types.NewString("Longitude"), types.Float32(val))}
}

// Georectangle

type Georectangle struct {
	m types.Map
}

func NewGeorectangle() Georectangle {
	return Georectangle{types.NewMap(
		types.NewString("$name"), types.NewString("Georectangle"),
		types.NewString("$type"), types.MakeTypeRef("Georectangle", __mainPackageInFile_types_CachedRef),
		types.NewString("TopLeft"), NewGeoposition().NomsValue(),
		types.NewString("BottomRight"), NewGeoposition().NomsValue(),
	)}
}

type GeorectangleDef struct {
	TopLeft     GeopositionDef
	BottomRight GeopositionDef
}

func (def GeorectangleDef) New() Georectangle {
	return Georectangle{
		types.NewMap(
			types.NewString("$name"), types.NewString("Georectangle"),
			types.NewString("$type"), types.MakeTypeRef("Georectangle", __mainPackageInFile_types_CachedRef),
			types.NewString("TopLeft"), def.TopLeft.New().NomsValue(),
			types.NewString("BottomRight"), def.BottomRight.New().NomsValue(),
		)}
}

func (s Georectangle) Def() (d GeorectangleDef) {
	d.TopLeft = GeopositionFromVal(s.m.Get(types.NewString("TopLeft"))).Def()
	d.BottomRight = GeopositionFromVal(s.m.Get(types.NewString("BottomRight"))).Def()
	return
}

var __typeRefForGeorectangle = types.MakeTypeRef("Georectangle", __mainPackageInFile_types_CachedRef)

func (m Georectangle) TypeRef() types.TypeRef {
	return __typeRefForGeorectangle
}

func init() {
	types.RegisterFromValFunction(__typeRefForGeorectangle, func(v types.Value) types.NomsValue {
		return GeorectangleFromVal(v)
	})
}

func GeorectangleFromVal(val types.Value) Georectangle {
	// TODO: Validate here
	return Georectangle{val.(types.Map)}
}

func (s Georectangle) NomsValue() types.Value {
	return s.m
}

func (s Georectangle) Equals(other types.Value) bool {
	if other, ok := other.(Georectangle); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s Georectangle) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Georectangle) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s Georectangle) TopLeft() Geoposition {
	return GeopositionFromVal(s.m.Get(types.NewString("TopLeft")))
}

func (s Georectangle) SetTopLeft(val Geoposition) Georectangle {
	return Georectangle{s.m.Set(types.NewString("TopLeft"), val.NomsValue())}
}

func (s Georectangle) BottomRight() Geoposition {
	return GeopositionFromVal(s.m.Get(types.NewString("BottomRight")))
}

func (s Georectangle) SetBottomRight(val Geoposition) Georectangle {
	return Georectangle{s.m.Set(types.NewString("BottomRight"), val.NomsValue())}
}

// Node

type Node struct {
	m types.Map
}

func NewNode() Node {
	return Node{types.NewMap(
		types.NewString("$name"), types.NewString("Node"),
		types.NewString("$type"), types.MakeTypeRef("Node", __mainPackageInFile_types_CachedRef),
		types.NewString("Geoposition"), NewGeoposition().NomsValue(),
		types.NewString("Reference"), types.Ref{R: ref.Ref{}},
	)}
}

type NodeDef struct {
	Geoposition GeopositionDef
	Reference   ref.Ref
}

func (def NodeDef) New() Node {
	return Node{
		types.NewMap(
			types.NewString("$name"), types.NewString("Node"),
			types.NewString("$type"), types.MakeTypeRef("Node", __mainPackageInFile_types_CachedRef),
			types.NewString("Geoposition"), def.Geoposition.New().NomsValue(),
			types.NewString("Reference"), types.Ref{R: def.Reference},
		)}
}

func (s Node) Def() (d NodeDef) {
	d.Geoposition = GeopositionFromVal(s.m.Get(types.NewString("Geoposition"))).Def()
	d.Reference = s.m.Get(types.NewString("Reference")).Ref()
	return
}

var __typeRefForNode = types.MakeTypeRef("Node", __mainPackageInFile_types_CachedRef)

func (m Node) TypeRef() types.TypeRef {
	return __typeRefForNode
}

func init() {
	types.RegisterFromValFunction(__typeRefForNode, func(v types.Value) types.NomsValue {
		return NodeFromVal(v)
	})
}

func NodeFromVal(val types.Value) Node {
	// TODO: Validate here
	return Node{val.(types.Map)}
}

func (s Node) NomsValue() types.Value {
	return s.m
}

func (s Node) Equals(other types.Value) bool {
	if other, ok := other.(Node); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s Node) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Node) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s Node) Geoposition() Geoposition {
	return GeopositionFromVal(s.m.Get(types.NewString("Geoposition")))
}

func (s Node) SetGeoposition(val Geoposition) Node {
	return Node{s.m.Set(types.NewString("Geoposition"), val.NomsValue())}
}

func (s Node) Reference() RefOfValue {
	return RefOfValueFromVal(s.m.Get(types.NewString("Reference")))
}

func (s Node) SetReference(val RefOfValue) Node {
	return Node{s.m.Set(types.NewString("Reference"), val.NomsValue())}
}

// RefOfValue

type RefOfValue struct {
	r ref.Ref
}

func NewRefOfValue(r ref.Ref) RefOfValue {
	return RefOfValue{r}
}

func (r RefOfValue) Ref() ref.Ref {
	return r.r
}

func (r RefOfValue) Equals(other types.Value) bool {
	if other, ok := other.(RefOfValue); ok {
		return r.r == other.r
	}
	return false
}

func (r RefOfValue) Chunks() []types.Future {
	return nil
}

func (r RefOfValue) NomsValue() types.Value {
	return types.Ref{R: r.r}
}

func RefOfValueFromVal(p types.Value) RefOfValue {
	return RefOfValue{p.(types.Ref).Ref()}
}

// A Noms Value that describes RefOfValue.
var __typeRefForRefOfValue = types.MakeCompoundTypeRef("", types.RefKind, types.MakePrimitiveTypeRef(types.ValueKind))

func (m RefOfValue) TypeRef() types.TypeRef {
	return __typeRefForRefOfValue
}

func init() {
	types.RegisterFromValFunction(__typeRefForRefOfValue, func(v types.Value) types.NomsValue {
		return RefOfValueFromVal(v)
	})
}

func (r RefOfValue) GetValue(cs chunks.ChunkSource) types.Value {
	return types.ReadValue(r.r, cs)
}

func (r RefOfValue) SetValue(val types.Value, cs chunks.ChunkSink) RefOfValue {
	ref := types.WriteValue(val, cs)
	return RefOfValue{ref}
}

// QuadTree

type QuadTree struct {
	m types.Map
}

func NewQuadTree() QuadTree {
	return QuadTree{types.NewMap(
		types.NewString("$name"), types.NewString("QuadTree"),
		types.NewString("$type"), types.MakeTypeRef("QuadTree", __mainPackageInFile_types_CachedRef),
		types.NewString("Nodes"), types.NewList(),
		types.NewString("Tiles"), types.NewMap(),
		types.NewString("Depth"), types.UInt8(0),
		types.NewString("NumDescendents"), types.UInt32(0),
		types.NewString("Path"), types.NewString(""),
		types.NewString("Georectangle"), NewGeorectangle().NomsValue(),
	)}
}

type QuadTreeDef struct {
	Nodes          ListOfNodeDef
	Tiles          MapOfStringToQuadTreeDef
	Depth          uint8
	NumDescendents uint32
	Path           string
	Georectangle   GeorectangleDef
}

func (def QuadTreeDef) New() QuadTree {
	return QuadTree{
		types.NewMap(
			types.NewString("$name"), types.NewString("QuadTree"),
			types.NewString("$type"), types.MakeTypeRef("QuadTree", __mainPackageInFile_types_CachedRef),
			types.NewString("Nodes"), def.Nodes.New().NomsValue(),
			types.NewString("Tiles"), def.Tiles.New().NomsValue(),
			types.NewString("Depth"), types.UInt8(def.Depth),
			types.NewString("NumDescendents"), types.UInt32(def.NumDescendents),
			types.NewString("Path"), types.NewString(def.Path),
			types.NewString("Georectangle"), def.Georectangle.New().NomsValue(),
		)}
}

func (s QuadTree) Def() (d QuadTreeDef) {
	d.Nodes = ListOfNodeFromVal(s.m.Get(types.NewString("Nodes"))).Def()
	d.Tiles = MapOfStringToQuadTreeFromVal(s.m.Get(types.NewString("Tiles"))).Def()
	d.Depth = uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
	d.NumDescendents = uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
	d.Path = s.m.Get(types.NewString("Path")).(types.String).String()
	d.Georectangle = GeorectangleFromVal(s.m.Get(types.NewString("Georectangle"))).Def()
	return
}

var __typeRefForQuadTree = types.MakeTypeRef("QuadTree", __mainPackageInFile_types_CachedRef)

func (m QuadTree) TypeRef() types.TypeRef {
	return __typeRefForQuadTree
}

func init() {
	types.RegisterFromValFunction(__typeRefForQuadTree, func(v types.Value) types.NomsValue {
		return QuadTreeFromVal(v)
	})
}

func QuadTreeFromVal(val types.Value) QuadTree {
	// TODO: Validate here
	return QuadTree{val.(types.Map)}
}

func (s QuadTree) NomsValue() types.Value {
	return s.m
}

func (s QuadTree) Equals(other types.Value) bool {
	if other, ok := other.(QuadTree); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s QuadTree) Ref() ref.Ref {
	return s.m.Ref()
}

func (s QuadTree) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s QuadTree) Nodes() ListOfNode {
	return ListOfNodeFromVal(s.m.Get(types.NewString("Nodes")))
}

func (s QuadTree) SetNodes(val ListOfNode) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Nodes"), val.NomsValue())}
}

func (s QuadTree) Tiles() MapOfStringToQuadTree {
	return MapOfStringToQuadTreeFromVal(s.m.Get(types.NewString("Tiles")))
}

func (s QuadTree) SetTiles(val MapOfStringToQuadTree) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Tiles"), val.NomsValue())}
}

func (s QuadTree) Depth() uint8 {
	return uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
}

func (s QuadTree) SetDepth(val uint8) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Depth"), types.UInt8(val))}
}

func (s QuadTree) NumDescendents() uint32 {
	return uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
}

func (s QuadTree) SetNumDescendents(val uint32) QuadTree {
	return QuadTree{s.m.Set(types.NewString("NumDescendents"), types.UInt32(val))}
}

func (s QuadTree) Path() string {
	return s.m.Get(types.NewString("Path")).(types.String).String()
}

func (s QuadTree) SetPath(val string) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Path"), types.NewString(val))}
}

func (s QuadTree) Georectangle() Georectangle {
	return GeorectangleFromVal(s.m.Get(types.NewString("Georectangle")))
}

func (s QuadTree) SetGeorectangle(val Georectangle) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Georectangle"), val.NomsValue())}
}

// ListOfNode

type ListOfNode struct {
	l types.List
}

func NewListOfNode() ListOfNode {
	return ListOfNode{types.NewList()}
}

type ListOfNodeDef []NodeDef

func (def ListOfNodeDef) New() ListOfNode {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New().NomsValue()
	}
	return ListOfNode{types.NewList(l...)}
}

func (l ListOfNode) Def() ListOfNodeDef {
	d := make([]NodeDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = NodeFromVal(l.l.Get(i)).Def()
	}
	return d
}

func ListOfNodeFromVal(val types.Value) ListOfNode {
	// TODO: Validate here
	return ListOfNode{val.(types.List)}
}

func (l ListOfNode) NomsValue() types.Value {
	return l.l
}

func (l ListOfNode) Equals(other types.Value) bool {
	if other, ok := other.(ListOfNode); ok {
		return l.l.Equals(other.l)
	}
	return false
}

func (l ListOfNode) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfNode) Chunks() []types.Future {
	return l.l.Chunks()
}

// A Noms Value that describes ListOfNode.
var __typeRefForListOfNode = types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef("Node", __mainPackageInFile_types_CachedRef))

func (m ListOfNode) TypeRef() types.TypeRef {
	return __typeRefForListOfNode
}

func init() {
	types.RegisterFromValFunction(__typeRefForListOfNode, func(v types.Value) types.NomsValue {
		return ListOfNodeFromVal(v)
	})
}

func (l ListOfNode) Len() uint64 {
	return l.l.Len()
}

func (l ListOfNode) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfNode) Get(i uint64) Node {
	return NodeFromVal(l.l.Get(i))
}

func (l ListOfNode) Slice(idx uint64, end uint64) ListOfNode {
	return ListOfNode{l.l.Slice(idx, end)}
}

func (l ListOfNode) Set(i uint64, val Node) ListOfNode {
	return ListOfNode{l.l.Set(i, val.NomsValue())}
}

func (l ListOfNode) Append(v ...Node) ListOfNode {
	return ListOfNode{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfNode) Insert(idx uint64, v ...Node) ListOfNode {
	return ListOfNode{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfNode) Remove(idx uint64, end uint64) ListOfNode {
	return ListOfNode{l.l.Remove(idx, end)}
}

func (l ListOfNode) RemoveAt(idx uint64) ListOfNode {
	return ListOfNode{(l.l.RemoveAt(idx))}
}

func (l ListOfNode) fromElemSlice(p []Node) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfNodeIterCallback func(v Node, i uint64) (stop bool)

func (l ListOfNode) Iter(cb ListOfNodeIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(NodeFromVal(v), i)
	})
}

type ListOfNodeIterAllCallback func(v Node, i uint64)

func (l ListOfNode) IterAll(cb ListOfNodeIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(NodeFromVal(v), i)
	})
}

type ListOfNodeFilterCallback func(v Node, i uint64) (keep bool)

func (l ListOfNode) Filter(cb ListOfNodeFilterCallback) ListOfNode {
	nl := NewListOfNode()
	l.IterAll(func(v Node, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// MapOfStringToQuadTree

type MapOfStringToQuadTree struct {
	m types.Map
}

func NewMapOfStringToQuadTree() MapOfStringToQuadTree {
	return MapOfStringToQuadTree{types.NewMap()}
}

type MapOfStringToQuadTreeDef map[string]QuadTreeDef

func (def MapOfStringToQuadTreeDef) New() MapOfStringToQuadTree {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New().NomsValue())
	}
	return MapOfStringToQuadTree{types.NewMap(kv...)}
}

func (m MapOfStringToQuadTree) Def() MapOfStringToQuadTreeDef {
	def := make(map[string]QuadTreeDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = QuadTreeFromVal(v).Def()
		return false
	})
	return def
}

func MapOfStringToQuadTreeFromVal(p types.Value) MapOfStringToQuadTree {
	// TODO: Validate here
	return MapOfStringToQuadTree{p.(types.Map)}
}

func (m MapOfStringToQuadTree) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToQuadTree); ok {
		return m.m.Equals(other.m)
	}
	return false
}

func (m MapOfStringToQuadTree) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToQuadTree) Chunks() []types.Future {
	return m.m.Chunks()
}

// A Noms Value that describes MapOfStringToQuadTree.
var __typeRefForMapOfStringToQuadTree = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef("QuadTree", __mainPackageInFile_types_CachedRef))

func (m MapOfStringToQuadTree) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToQuadTree
}

func init() {
	types.RegisterFromValFunction(__typeRefForMapOfStringToQuadTree, func(v types.Value) types.NomsValue {
		return MapOfStringToQuadTreeFromVal(v)
	})
}

func (m MapOfStringToQuadTree) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToQuadTree) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToQuadTree) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToQuadTree) Get(p string) QuadTree {
	return QuadTreeFromVal(m.m.Get(types.NewString(p)))
}

func (m MapOfStringToQuadTree) Set(k string, v QuadTree) MapOfStringToQuadTree {
	return MapOfStringToQuadTree{m.m.Set(types.NewString(k), v.NomsValue())}
}

// TODO: Implement SetM?

func (m MapOfStringToQuadTree) Remove(p string) MapOfStringToQuadTree {
	return MapOfStringToQuadTree{m.m.Remove(types.NewString(p))}
}

type MapOfStringToQuadTreeIterCallback func(k string, v QuadTree) (stop bool)

func (m MapOfStringToQuadTree) Iter(cb MapOfStringToQuadTreeIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), QuadTreeFromVal(v))
	})
}

type MapOfStringToQuadTreeIterAllCallback func(k string, v QuadTree)

func (m MapOfStringToQuadTree) IterAll(cb MapOfStringToQuadTreeIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), QuadTreeFromVal(v))
	})
}

type MapOfStringToQuadTreeFilterCallback func(k string, v QuadTree) (keep bool)

func (m MapOfStringToQuadTree) Filter(cb MapOfStringToQuadTreeFilterCallback) MapOfStringToQuadTree {
	nm := NewMapOfStringToQuadTree()
	m.IterAll(func(k string, v QuadTree) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// SQuadTree

type SQuadTree struct {
	m types.Map
}

func NewSQuadTree() SQuadTree {
	return SQuadTree{types.NewMap(
		types.NewString("$name"), types.NewString("SQuadTree"),
		types.NewString("$type"), types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef),
		types.NewString("Nodes"), types.NewList(),
		types.NewString("Tiles"), types.NewMap(),
		types.NewString("Depth"), types.UInt8(0),
		types.NewString("NumDescendents"), types.UInt32(0),
		types.NewString("Path"), types.NewString(""),
		types.NewString("Georectangle"), NewGeorectangle().NomsValue(),
	)}
}

type SQuadTreeDef struct {
	Nodes          ListOfRefOfValueDef
	Tiles          MapOfStringToRefOfSQuadTreeDef
	Depth          uint8
	NumDescendents uint32
	Path           string
	Georectangle   GeorectangleDef
}

func (def SQuadTreeDef) New() SQuadTree {
	return SQuadTree{
		types.NewMap(
			types.NewString("$name"), types.NewString("SQuadTree"),
			types.NewString("$type"), types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef),
			types.NewString("Nodes"), def.Nodes.New().NomsValue(),
			types.NewString("Tiles"), def.Tiles.New().NomsValue(),
			types.NewString("Depth"), types.UInt8(def.Depth),
			types.NewString("NumDescendents"), types.UInt32(def.NumDescendents),
			types.NewString("Path"), types.NewString(def.Path),
			types.NewString("Georectangle"), def.Georectangle.New().NomsValue(),
		)}
}

func (s SQuadTree) Def() (d SQuadTreeDef) {
	d.Nodes = ListOfRefOfValueFromVal(s.m.Get(types.NewString("Nodes"))).Def()
	d.Tiles = MapOfStringToRefOfSQuadTreeFromVal(s.m.Get(types.NewString("Tiles"))).Def()
	d.Depth = uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
	d.NumDescendents = uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
	d.Path = s.m.Get(types.NewString("Path")).(types.String).String()
	d.Georectangle = GeorectangleFromVal(s.m.Get(types.NewString("Georectangle"))).Def()
	return
}

var __typeRefForSQuadTree = types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef)

func (m SQuadTree) TypeRef() types.TypeRef {
	return __typeRefForSQuadTree
}

func init() {
	types.RegisterFromValFunction(__typeRefForSQuadTree, func(v types.Value) types.NomsValue {
		return SQuadTreeFromVal(v)
	})
}

func SQuadTreeFromVal(val types.Value) SQuadTree {
	// TODO: Validate here
	return SQuadTree{val.(types.Map)}
}

func (s SQuadTree) NomsValue() types.Value {
	return s.m
}

func (s SQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(SQuadTree); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s SQuadTree) Ref() ref.Ref {
	return s.m.Ref()
}

func (s SQuadTree) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s SQuadTree) Nodes() ListOfRefOfValue {
	return ListOfRefOfValueFromVal(s.m.Get(types.NewString("Nodes")))
}

func (s SQuadTree) SetNodes(val ListOfRefOfValue) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Nodes"), val.NomsValue())}
}

func (s SQuadTree) Tiles() MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTreeFromVal(s.m.Get(types.NewString("Tiles")))
}

func (s SQuadTree) SetTiles(val MapOfStringToRefOfSQuadTree) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Tiles"), val.NomsValue())}
}

func (s SQuadTree) Depth() uint8 {
	return uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
}

func (s SQuadTree) SetDepth(val uint8) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Depth"), types.UInt8(val))}
}

func (s SQuadTree) NumDescendents() uint32 {
	return uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
}

func (s SQuadTree) SetNumDescendents(val uint32) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("NumDescendents"), types.UInt32(val))}
}

func (s SQuadTree) Path() string {
	return s.m.Get(types.NewString("Path")).(types.String).String()
}

func (s SQuadTree) SetPath(val string) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Path"), types.NewString(val))}
}

func (s SQuadTree) Georectangle() Georectangle {
	return GeorectangleFromVal(s.m.Get(types.NewString("Georectangle")))
}

func (s SQuadTree) SetGeorectangle(val Georectangle) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Georectangle"), val.NomsValue())}
}

// ListOfRefOfValue

type ListOfRefOfValue struct {
	l types.List
}

func NewListOfRefOfValue() ListOfRefOfValue {
	return ListOfRefOfValue{types.NewList()}
}

type ListOfRefOfValueDef []ref.Ref

func (def ListOfRefOfValueDef) New() ListOfRefOfValue {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = types.Ref{R: d}
	}
	return ListOfRefOfValue{types.NewList(l...)}
}

func (l ListOfRefOfValue) Def() ListOfRefOfValueDef {
	d := make([]ref.Ref, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).Ref()
	}
	return d
}

func ListOfRefOfValueFromVal(val types.Value) ListOfRefOfValue {
	// TODO: Validate here
	return ListOfRefOfValue{val.(types.List)}
}

func (l ListOfRefOfValue) NomsValue() types.Value {
	return l.l
}

func (l ListOfRefOfValue) Equals(other types.Value) bool {
	if other, ok := other.(ListOfRefOfValue); ok {
		return l.l.Equals(other.l)
	}
	return false
}

func (l ListOfRefOfValue) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfRefOfValue) Chunks() []types.Future {
	return l.l.Chunks()
}

// A Noms Value that describes ListOfRefOfValue.
var __typeRefForListOfRefOfValue = types.MakeCompoundTypeRef("", types.ListKind, types.MakeCompoundTypeRef("", types.RefKind, types.MakePrimitiveTypeRef(types.ValueKind)))

func (m ListOfRefOfValue) TypeRef() types.TypeRef {
	return __typeRefForListOfRefOfValue
}

func init() {
	types.RegisterFromValFunction(__typeRefForListOfRefOfValue, func(v types.Value) types.NomsValue {
		return ListOfRefOfValueFromVal(v)
	})
}

func (l ListOfRefOfValue) Len() uint64 {
	return l.l.Len()
}

func (l ListOfRefOfValue) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfRefOfValue) Get(i uint64) RefOfValue {
	return RefOfValueFromVal(l.l.Get(i))
}

func (l ListOfRefOfValue) Slice(idx uint64, end uint64) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Slice(idx, end)}
}

func (l ListOfRefOfValue) Set(i uint64, val RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Set(i, val.NomsValue())}
}

func (l ListOfRefOfValue) Append(v ...RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfRefOfValue) Insert(idx uint64, v ...RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfRefOfValue) Remove(idx uint64, end uint64) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Remove(idx, end)}
}

func (l ListOfRefOfValue) RemoveAt(idx uint64) ListOfRefOfValue {
	return ListOfRefOfValue{(l.l.RemoveAt(idx))}
}

func (l ListOfRefOfValue) fromElemSlice(p []RefOfValue) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfRefOfValueIterCallback func(v RefOfValue, i uint64) (stop bool)

func (l ListOfRefOfValue) Iter(cb ListOfRefOfValueIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(RefOfValueFromVal(v), i)
	})
}

type ListOfRefOfValueIterAllCallback func(v RefOfValue, i uint64)

func (l ListOfRefOfValue) IterAll(cb ListOfRefOfValueIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(RefOfValueFromVal(v), i)
	})
}

type ListOfRefOfValueFilterCallback func(v RefOfValue, i uint64) (keep bool)

func (l ListOfRefOfValue) Filter(cb ListOfRefOfValueFilterCallback) ListOfRefOfValue {
	nl := NewListOfRefOfValue()
	l.IterAll(func(v RefOfValue, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// MapOfStringToRefOfSQuadTree

type MapOfStringToRefOfSQuadTree struct {
	m types.Map
}

func NewMapOfStringToRefOfSQuadTree() MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{types.NewMap()}
}

type MapOfStringToRefOfSQuadTreeDef map[string]ref.Ref

func (def MapOfStringToRefOfSQuadTreeDef) New() MapOfStringToRefOfSQuadTree {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), types.Ref{R: v})
	}
	return MapOfStringToRefOfSQuadTree{types.NewMap(kv...)}
}

func (m MapOfStringToRefOfSQuadTree) Def() MapOfStringToRefOfSQuadTreeDef {
	def := make(map[string]ref.Ref)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.Ref()
		return false
	})
	return def
}

func MapOfStringToRefOfSQuadTreeFromVal(p types.Value) MapOfStringToRefOfSQuadTree {
	// TODO: Validate here
	return MapOfStringToRefOfSQuadTree{p.(types.Map)}
}

func (m MapOfStringToRefOfSQuadTree) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToRefOfSQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToRefOfSQuadTree); ok {
		return m.m.Equals(other.m)
	}
	return false
}

func (m MapOfStringToRefOfSQuadTree) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToRefOfSQuadTree) Chunks() []types.Future {
	return m.m.Chunks()
}

// A Noms Value that describes MapOfStringToRefOfSQuadTree.
var __typeRefForMapOfStringToRefOfSQuadTree = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef("", types.RefKind, types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef)))

func (m MapOfStringToRefOfSQuadTree) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToRefOfSQuadTree
}

func init() {
	types.RegisterFromValFunction(__typeRefForMapOfStringToRefOfSQuadTree, func(v types.Value) types.NomsValue {
		return MapOfStringToRefOfSQuadTreeFromVal(v)
	})
}

func (m MapOfStringToRefOfSQuadTree) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToRefOfSQuadTree) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToRefOfSQuadTree) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToRefOfSQuadTree) Get(p string) RefOfSQuadTree {
	return RefOfSQuadTreeFromVal(m.m.Get(types.NewString(p)))
}

func (m MapOfStringToRefOfSQuadTree) Set(k string, v RefOfSQuadTree) MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{m.m.Set(types.NewString(k), v.NomsValue())}
}

// TODO: Implement SetM?

func (m MapOfStringToRefOfSQuadTree) Remove(p string) MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{m.m.Remove(types.NewString(p))}
}

type MapOfStringToRefOfSQuadTreeIterCallback func(k string, v RefOfSQuadTree) (stop bool)

func (m MapOfStringToRefOfSQuadTree) Iter(cb MapOfStringToRefOfSQuadTreeIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), RefOfSQuadTreeFromVal(v))
	})
}

type MapOfStringToRefOfSQuadTreeIterAllCallback func(k string, v RefOfSQuadTree)

func (m MapOfStringToRefOfSQuadTree) IterAll(cb MapOfStringToRefOfSQuadTreeIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), RefOfSQuadTreeFromVal(v))
	})
}

type MapOfStringToRefOfSQuadTreeFilterCallback func(k string, v RefOfSQuadTree) (keep bool)

func (m MapOfStringToRefOfSQuadTree) Filter(cb MapOfStringToRefOfSQuadTreeFilterCallback) MapOfStringToRefOfSQuadTree {
	nm := NewMapOfStringToRefOfSQuadTree()
	m.IterAll(func(k string, v RefOfSQuadTree) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// RefOfSQuadTree

type RefOfSQuadTree struct {
	r ref.Ref
}

func NewRefOfSQuadTree(r ref.Ref) RefOfSQuadTree {
	return RefOfSQuadTree{r}
}

func (r RefOfSQuadTree) Ref() ref.Ref {
	return r.r
}

func (r RefOfSQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(RefOfSQuadTree); ok {
		return r.r == other.r
	}
	return false
}

func (r RefOfSQuadTree) Chunks() []types.Future {
	return nil
}

func (r RefOfSQuadTree) NomsValue() types.Value {
	return types.Ref{R: r.r}
}

func RefOfSQuadTreeFromVal(p types.Value) RefOfSQuadTree {
	return RefOfSQuadTree{p.(types.Ref).Ref()}
}

// A Noms Value that describes RefOfSQuadTree.
var __typeRefForRefOfSQuadTree = types.MakeCompoundTypeRef("", types.RefKind, types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef))

func (m RefOfSQuadTree) TypeRef() types.TypeRef {
	return __typeRefForRefOfSQuadTree
}

func init() {
	types.RegisterFromValFunction(__typeRefForRefOfSQuadTree, func(v types.Value) types.NomsValue {
		return RefOfSQuadTreeFromVal(v)
	})
}

func (r RefOfSQuadTree) GetValue(cs chunks.ChunkSource) SQuadTree {
	return SQuadTreeFromVal(types.ReadValue(r.r, cs))
}

func (r RefOfSQuadTree) SetValue(val SQuadTree, cs chunks.ChunkSink) RefOfSQuadTree {
	ref := types.WriteValue(val.NomsValue(), cs)
	return RefOfSQuadTree{ref}
}
