// Code generated from .\SwiftLan.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SwiftLanLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SwiftLanLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftlanlexerLexerInit() {
	staticData := &SwiftLanLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "'('", "')'", "'?'", "':'", "'['", "']'", "'var'", "'print'",
		"'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "", "", "",
		"", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'^'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'>'", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "Declarar", "Print", "IntDecla", "FloatDecla",
		"BoolDecla", "StringDecla", "CharDecla", "BoolVal", "Number", "Id",
		"String", "Or", "And", "Equals", "NotEquals", "MayQEquals", "MinQEquals",
		"Power", "Add", "Minus", "Mult", "Div", "Mod", "MayorQue", "MenorQue",
		"Space", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "Declarar",
		"Print", "IntDecla", "FloatDecla", "BoolDecla", "StringDecla", "CharDecla",
		"BoolVal", "Number", "Id", "String", "Int", "Digit", "Or", "And", "Equals",
		"NotEquals", "MayQEquals", "MinQEquals", "Power", "Add", "Minus", "Mult",
		"Div", "Mod", "MayorQue", "MenorQue", "Space", "COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 256, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3,
		14, 141, 8, 14, 1, 15, 1, 15, 1, 15, 5, 15, 146, 8, 15, 10, 15, 12, 15,
		149, 9, 15, 3, 15, 151, 8, 15, 1, 16, 1, 16, 5, 16, 155, 8, 16, 10, 16,
		12, 16, 158, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 164, 8, 17, 10,
		17, 12, 17, 167, 9, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 174,
		8, 17, 10, 17, 12, 17, 177, 9, 17, 1, 17, 3, 17, 180, 8, 17, 1, 18, 1,
		18, 5, 18, 184, 8, 18, 10, 18, 12, 18, 187, 9, 18, 1, 18, 3, 18, 190, 8,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 35, 1, 35, 5, 35, 236, 8, 35, 10, 35, 12, 35, 239, 9, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 250, 8,
		36, 10, 36, 12, 36, 253, 9, 36, 1, 36, 1, 36, 1, 237, 0, 37, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 0, 39, 0, 41, 19, 43, 20,
		45, 21, 47, 22, 49, 23, 51, 24, 53, 25, 55, 26, 57, 27, 59, 28, 61, 29,
		63, 30, 65, 31, 67, 32, 69, 33, 71, 34, 73, 35, 1, 0, 10, 2, 0, 65, 90,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 34, 34, 4, 0, 10,
		10, 13, 13, 34, 34, 92, 92, 2, 0, 10, 10, 13, 13, 1, 0, 39, 39, 4, 0, 10,
		10, 13, 13, 39, 39, 92, 92, 1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 9, 10, 12,
		13, 32, 32, 266, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 1, 75, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 79, 1,
		0, 0, 0, 7, 81, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 85, 1, 0, 0, 0, 13,
		87, 1, 0, 0, 0, 15, 89, 1, 0, 0, 0, 17, 93, 1, 0, 0, 0, 19, 99, 1, 0, 0,
		0, 21, 103, 1, 0, 0, 0, 23, 109, 1, 0, 0, 0, 25, 114, 1, 0, 0, 0, 27, 121,
		1, 0, 0, 0, 29, 140, 1, 0, 0, 0, 31, 142, 1, 0, 0, 0, 33, 152, 1, 0, 0,
		0, 35, 179, 1, 0, 0, 0, 37, 189, 1, 0, 0, 0, 39, 191, 1, 0, 0, 0, 41, 193,
		1, 0, 0, 0, 43, 196, 1, 0, 0, 0, 45, 199, 1, 0, 0, 0, 47, 202, 1, 0, 0,
		0, 49, 205, 1, 0, 0, 0, 51, 208, 1, 0, 0, 0, 53, 211, 1, 0, 0, 0, 55, 213,
		1, 0, 0, 0, 57, 215, 1, 0, 0, 0, 59, 217, 1, 0, 0, 0, 61, 219, 1, 0, 0,
		0, 63, 221, 1, 0, 0, 0, 65, 223, 1, 0, 0, 0, 67, 225, 1, 0, 0, 0, 69, 227,
		1, 0, 0, 0, 71, 231, 1, 0, 0, 0, 73, 245, 1, 0, 0, 0, 75, 76, 5, 61, 0,
		0, 76, 2, 1, 0, 0, 0, 77, 78, 5, 40, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5,
		41, 0, 0, 80, 6, 1, 0, 0, 0, 81, 82, 5, 63, 0, 0, 82, 8, 1, 0, 0, 0, 83,
		84, 5, 58, 0, 0, 84, 10, 1, 0, 0, 0, 85, 86, 5, 91, 0, 0, 86, 12, 1, 0,
		0, 0, 87, 88, 5, 93, 0, 0, 88, 14, 1, 0, 0, 0, 89, 90, 5, 118, 0, 0, 90,
		91, 5, 97, 0, 0, 91, 92, 5, 114, 0, 0, 92, 16, 1, 0, 0, 0, 93, 94, 5, 112,
		0, 0, 94, 95, 5, 114, 0, 0, 95, 96, 5, 105, 0, 0, 96, 97, 5, 110, 0, 0,
		97, 98, 5, 116, 0, 0, 98, 18, 1, 0, 0, 0, 99, 100, 5, 73, 0, 0, 100, 101,
		5, 110, 0, 0, 101, 102, 5, 116, 0, 0, 102, 20, 1, 0, 0, 0, 103, 104, 5,
		70, 0, 0, 104, 105, 5, 108, 0, 0, 105, 106, 5, 111, 0, 0, 106, 107, 5,
		97, 0, 0, 107, 108, 5, 116, 0, 0, 108, 22, 1, 0, 0, 0, 109, 110, 5, 66,
		0, 0, 110, 111, 5, 111, 0, 0, 111, 112, 5, 111, 0, 0, 112, 113, 5, 108,
		0, 0, 113, 24, 1, 0, 0, 0, 114, 115, 5, 83, 0, 0, 115, 116, 5, 116, 0,
		0, 116, 117, 5, 114, 0, 0, 117, 118, 5, 105, 0, 0, 118, 119, 5, 110, 0,
		0, 119, 120, 5, 103, 0, 0, 120, 26, 1, 0, 0, 0, 121, 122, 5, 67, 0, 0,
		122, 123, 5, 104, 0, 0, 123, 124, 5, 97, 0, 0, 124, 125, 5, 114, 0, 0,
		125, 126, 5, 97, 0, 0, 126, 127, 5, 99, 0, 0, 127, 128, 5, 116, 0, 0, 128,
		129, 5, 101, 0, 0, 129, 130, 5, 114, 0, 0, 130, 28, 1, 0, 0, 0, 131, 132,
		5, 116, 0, 0, 132, 133, 5, 114, 0, 0, 133, 134, 5, 117, 0, 0, 134, 141,
		5, 101, 0, 0, 135, 136, 5, 102, 0, 0, 136, 137, 5, 97, 0, 0, 137, 138,
		5, 108, 0, 0, 138, 139, 5, 115, 0, 0, 139, 141, 5, 101, 0, 0, 140, 131,
		1, 0, 0, 0, 140, 135, 1, 0, 0, 0, 141, 30, 1, 0, 0, 0, 142, 150, 3, 37,
		18, 0, 143, 147, 5, 46, 0, 0, 144, 146, 3, 39, 19, 0, 145, 144, 1, 0, 0,
		0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 143, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 32, 1, 0, 0, 0, 152, 156, 7, 0, 0, 0, 153, 155, 7, 1,
		0, 0, 154, 153, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0,
		156, 157, 1, 0, 0, 0, 157, 34, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 165,
		7, 2, 0, 0, 160, 164, 8, 3, 0, 0, 161, 162, 5, 92, 0, 0, 162, 164, 8, 4,
		0, 0, 163, 160, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0,
		165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 168, 180, 7, 2, 0, 0, 169, 175, 7, 5, 0, 0, 170, 174,
		8, 6, 0, 0, 171, 172, 5, 92, 0, 0, 172, 174, 8, 4, 0, 0, 173, 170, 1, 0,
		0, 0, 173, 171, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0,
		175, 176, 1, 0, 0, 0, 176, 178, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178,
		180, 7, 5, 0, 0, 179, 159, 1, 0, 0, 0, 179, 169, 1, 0, 0, 0, 180, 36, 1,
		0, 0, 0, 181, 185, 7, 7, 0, 0, 182, 184, 3, 39, 19, 0, 183, 182, 1, 0,
		0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0,
		186, 190, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 190, 5, 48, 0, 0, 189,
		181, 1, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 38, 1, 0, 0, 0, 191, 192, 7,
		8, 0, 0, 192, 40, 1, 0, 0, 0, 193, 194, 5, 124, 0, 0, 194, 195, 5, 124,
		0, 0, 195, 42, 1, 0, 0, 0, 196, 197, 5, 38, 0, 0, 197, 198, 5, 38, 0, 0,
		198, 44, 1, 0, 0, 0, 199, 200, 5, 61, 0, 0, 200, 201, 5, 61, 0, 0, 201,
		46, 1, 0, 0, 0, 202, 203, 5, 33, 0, 0, 203, 204, 5, 61, 0, 0, 204, 48,
		1, 0, 0, 0, 205, 206, 5, 62, 0, 0, 206, 207, 5, 61, 0, 0, 207, 50, 1, 0,
		0, 0, 208, 209, 5, 60, 0, 0, 209, 210, 5, 61, 0, 0, 210, 52, 1, 0, 0, 0,
		211, 212, 5, 94, 0, 0, 212, 54, 1, 0, 0, 0, 213, 214, 5, 43, 0, 0, 214,
		56, 1, 0, 0, 0, 215, 216, 5, 45, 0, 0, 216, 58, 1, 0, 0, 0, 217, 218, 5,
		42, 0, 0, 218, 60, 1, 0, 0, 0, 219, 220, 5, 47, 0, 0, 220, 62, 1, 0, 0,
		0, 221, 222, 5, 37, 0, 0, 222, 64, 1, 0, 0, 0, 223, 224, 5, 62, 0, 0, 224,
		66, 1, 0, 0, 0, 225, 226, 5, 60, 0, 0, 226, 68, 1, 0, 0, 0, 227, 228, 7,
		9, 0, 0, 228, 229, 1, 0, 0, 0, 229, 230, 6, 34, 0, 0, 230, 70, 1, 0, 0,
		0, 231, 232, 5, 47, 0, 0, 232, 233, 5, 42, 0, 0, 233, 237, 1, 0, 0, 0,
		234, 236, 9, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237,
		238, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 240, 241, 5, 42, 0, 0, 241, 242, 5, 47, 0, 0, 242, 243, 1,
		0, 0, 0, 243, 244, 6, 35, 0, 0, 244, 72, 1, 0, 0, 0, 245, 246, 5, 47, 0,
		0, 246, 247, 5, 47, 0, 0, 247, 251, 1, 0, 0, 0, 248, 250, 8, 4, 0, 0, 249,
		248, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252,
		1, 0, 0, 0, 252, 254, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 255, 6, 36,
		0, 0, 255, 74, 1, 0, 0, 0, 14, 0, 140, 147, 150, 156, 163, 165, 173, 175,
		179, 185, 189, 237, 251, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftLanLexerInit initializes any static state used to implement SwiftLanLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSwiftLanLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftLanLexerInit() {
	staticData := &SwiftLanLexerLexerStaticData
	staticData.once.Do(swiftlanlexerLexerInit)
}

// NewSwiftLanLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSwiftLanLexer(input antlr.CharStream) *SwiftLanLexer {
	SwiftLanLexerInit()
	l := new(SwiftLanLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SwiftLanLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SwiftLan.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftLanLexer tokens.
const (
	SwiftLanLexerT__0         = 1
	SwiftLanLexerT__1         = 2
	SwiftLanLexerT__2         = 3
	SwiftLanLexerT__3         = 4
	SwiftLanLexerT__4         = 5
	SwiftLanLexerT__5         = 6
	SwiftLanLexerT__6         = 7
	SwiftLanLexerDeclarar     = 8
	SwiftLanLexerPrint        = 9
	SwiftLanLexerIntDecla     = 10
	SwiftLanLexerFloatDecla   = 11
	SwiftLanLexerBoolDecla    = 12
	SwiftLanLexerStringDecla  = 13
	SwiftLanLexerCharDecla    = 14
	SwiftLanLexerBoolVal      = 15
	SwiftLanLexerNumber       = 16
	SwiftLanLexerId           = 17
	SwiftLanLexerString_      = 18
	SwiftLanLexerOr           = 19
	SwiftLanLexerAnd          = 20
	SwiftLanLexerEquals       = 21
	SwiftLanLexerNotEquals    = 22
	SwiftLanLexerMayQEquals   = 23
	SwiftLanLexerMinQEquals   = 24
	SwiftLanLexerPower        = 25
	SwiftLanLexerAdd          = 26
	SwiftLanLexerMinus        = 27
	SwiftLanLexerMult         = 28
	SwiftLanLexerDiv          = 29
	SwiftLanLexerMod          = 30
	SwiftLanLexerMayorQue     = 31
	SwiftLanLexerMenorQue     = 32
	SwiftLanLexerSpace        = 33
	SwiftLanLexerCOMMENT      = 34
	SwiftLanLexerLINE_COMMENT = 35
)
