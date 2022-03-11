// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2/analizador/ast"
import "OLC2/analizador/ast/interfaces"
import "OLC2/analizador/ast/expresion"
import "OLC2/analizador/ast/expresion/Accesos"
import "OLC2/analizador/ast/funbasica"
import "OLC2/analizador/ast/definicion"
import "OLC2/analizador/ast/definicion/defobjetos"
import "OLC2/analizador/ast/definicion/defarreglos"
import "OLC2/analizador/ast/control"
import "OLC2/analizador/ast/expresion2"
import "OLC2/analizador/ast/transferenciaFlujo"
import "OLC2/analizador/entorno"
import "OLC2/analizador/entorno/Simbolos"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 50, 565,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3,
	100, 10, 3, 12, 3, 14, 3, 103, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 118, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 128, 10, 6, 12, 6, 14, 6,
	131, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 5, 7, 144, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	5, 8, 166, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 173, 10, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 184, 10,
	10, 12, 10, 14, 10, 187, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 198, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 6,
	13, 214, 10, 13, 13, 13, 14, 13, 215, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 251, 10, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7,
	17, 274, 10, 17, 12, 17, 14, 17, 277, 11, 17, 3, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 290, 10, 19, 12,
	19, 14, 19, 293, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 5, 21, 326, 10, 21, 3, 22, 6, 22, 329, 10, 22, 13,
	22, 14, 22, 330, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5,
	24, 351, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 5, 26, 373, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 7, 27, 384, 10, 27, 12, 27, 14, 27, 387, 11, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 405, 10, 30, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 415, 10, 31, 12, 31, 14,
	31, 418, 11, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 5, 32, 430, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33,
	447, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 7, 39, 481, 10, 39, 12, 39, 14, 39, 484, 11, 39, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 491, 10, 40, 3, 41, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 502, 10, 41, 12, 41, 14,
	41, 505, 11, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 520, 10, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 7, 42, 532, 10, 42,
	12, 42, 14, 42, 535, 11, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 549, 10, 43, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	5, 44, 563, 10, 44, 3, 44, 2, 12, 4, 10, 18, 32, 36, 52, 60, 76, 80, 82,
	45, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 78, 80, 82, 84, 86, 2, 5, 3, 2, 36, 39, 3, 2, 40, 41, 3, 2, 42,
	43, 2, 573, 2, 88, 3, 2, 2, 2, 4, 91, 3, 2, 2, 2, 6, 104, 3, 2, 2, 2, 8,
	117, 3, 2, 2, 2, 10, 119, 3, 2, 2, 2, 12, 143, 3, 2, 2, 2, 14, 165, 3,
	2, 2, 2, 16, 172, 3, 2, 2, 2, 18, 174, 3, 2, 2, 2, 20, 197, 3, 2, 2, 2,
	22, 199, 3, 2, 2, 2, 24, 213, 3, 2, 2, 2, 26, 250, 3, 2, 2, 2, 28, 252,
	3, 2, 2, 2, 30, 258, 3, 2, 2, 2, 32, 265, 3, 2, 2, 2, 34, 278, 3, 2, 2,
	2, 36, 281, 3, 2, 2, 2, 38, 294, 3, 2, 2, 2, 40, 325, 3, 2, 2, 2, 42, 328,
	3, 2, 2, 2, 44, 334, 3, 2, 2, 2, 46, 350, 3, 2, 2, 2, 48, 352, 3, 2, 2,
	2, 50, 372, 3, 2, 2, 2, 52, 374, 3, 2, 2, 2, 54, 388, 3, 2, 2, 2, 56, 394,
	3, 2, 2, 2, 58, 404, 3, 2, 2, 2, 60, 406, 3, 2, 2, 2, 62, 429, 3, 2, 2,
	2, 64, 446, 3, 2, 2, 2, 66, 448, 3, 2, 2, 2, 68, 453, 3, 2, 2, 2, 70, 458,
	3, 2, 2, 2, 72, 464, 3, 2, 2, 2, 74, 468, 3, 2, 2, 2, 76, 471, 3, 2, 2,
	2, 78, 490, 3, 2, 2, 2, 80, 492, 3, 2, 2, 2, 82, 519, 3, 2, 2, 2, 84, 548,
	3, 2, 2, 2, 86, 562, 3, 2, 2, 2, 88, 89, 5, 4, 3, 2, 89, 90, 8, 2, 1, 2,
	90, 3, 3, 2, 2, 2, 91, 92, 8, 3, 1, 2, 92, 93, 5, 6, 4, 2, 93, 94, 8, 3,
	1, 2, 94, 101, 3, 2, 2, 2, 95, 96, 12, 4, 2, 2, 96, 97, 5, 6, 4, 2, 97,
	98, 8, 3, 1, 2, 98, 100, 3, 2, 2, 2, 99, 95, 3, 2, 2, 2, 100, 103, 3, 2,
	2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 5, 3, 2, 2, 2, 103,
	101, 3, 2, 2, 2, 104, 105, 7, 14, 2, 2, 105, 106, 7, 49, 2, 2, 106, 107,
	5, 8, 5, 2, 107, 108, 8, 4, 1, 2, 108, 7, 3, 2, 2, 2, 109, 110, 7, 5, 2,
	2, 110, 111, 5, 10, 6, 2, 111, 112, 7, 6, 2, 2, 112, 113, 8, 5, 1, 2, 113,
	118, 3, 2, 2, 2, 114, 115, 7, 5, 2, 2, 115, 116, 7, 6, 2, 2, 116, 118,
	8, 5, 1, 2, 117, 109, 3, 2, 2, 2, 117, 114, 3, 2, 2, 2, 118, 9, 3, 2, 2,
	2, 119, 120, 8, 6, 1, 2, 120, 121, 5, 12, 7, 2, 121, 122, 8, 6, 1, 2, 122,
	129, 3, 2, 2, 2, 123, 124, 12, 4, 2, 2, 124, 125, 5, 12, 7, 2, 125, 126,
	8, 6, 1, 2, 126, 128, 3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 128, 131, 3, 2,
	2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 11, 3, 2, 2, 2,
	131, 129, 3, 2, 2, 2, 132, 133, 5, 14, 8, 2, 133, 134, 8, 7, 1, 2, 134,
	144, 3, 2, 2, 2, 135, 136, 5, 54, 28, 2, 136, 137, 7, 30, 2, 2, 137, 138,
	8, 7, 1, 2, 138, 144, 3, 2, 2, 2, 139, 140, 5, 56, 29, 2, 140, 141, 7,
	30, 2, 2, 141, 142, 8, 7, 1, 2, 142, 144, 3, 2, 2, 2, 143, 132, 3, 2, 2,
	2, 143, 135, 3, 2, 2, 2, 143, 139, 3, 2, 2, 2, 144, 13, 3, 2, 2, 2, 145,
	146, 5, 22, 12, 2, 146, 147, 8, 8, 1, 2, 147, 166, 3, 2, 2, 2, 148, 149,
	5, 16, 9, 2, 149, 150, 5, 62, 32, 2, 150, 151, 7, 49, 2, 2, 151, 152, 7,
	3, 2, 2, 152, 153, 7, 4, 2, 2, 153, 154, 5, 46, 24, 2, 154, 155, 8, 8,
	1, 2, 155, 166, 3, 2, 2, 2, 156, 157, 5, 16, 9, 2, 157, 158, 5, 62, 32,
	2, 158, 159, 7, 49, 2, 2, 159, 160, 7, 3, 2, 2, 160, 161, 5, 18, 10, 2,
	161, 162, 7, 4, 2, 2, 162, 163, 5, 46, 24, 2, 163, 164, 8, 8, 1, 2, 164,
	166, 3, 2, 2, 2, 165, 145, 3, 2, 2, 2, 165, 148, 3, 2, 2, 2, 165, 156,
	3, 2, 2, 2, 166, 15, 3, 2, 2, 2, 167, 168, 7, 18, 2, 2, 168, 173, 8, 9,
	1, 2, 169, 170, 7, 17, 2, 2, 170, 173, 8, 9, 1, 2, 171, 173, 8, 9, 1, 2,
	172, 167, 3, 2, 2, 2, 172, 169, 3, 2, 2, 2, 172, 171, 3, 2, 2, 2, 173,
	17, 3, 2, 2, 2, 174, 175, 8, 10, 1, 2, 175, 176, 5, 20, 11, 2, 176, 177,
	8, 10, 1, 2, 177, 185, 3, 2, 2, 2, 178, 179, 12, 4, 2, 2, 179, 180, 7,
	29, 2, 2, 180, 181, 5, 20, 11, 2, 181, 182, 8, 10, 1, 2, 182, 184, 3, 2,
	2, 2, 183, 178, 3, 2, 2, 2, 184, 187, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2,
	185, 186, 3, 2, 2, 2, 186, 19, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 188, 189,
	5, 62, 32, 2, 189, 190, 7, 49, 2, 2, 190, 191, 8, 11, 1, 2, 191, 198, 3,
	2, 2, 2, 192, 193, 7, 40, 2, 2, 193, 194, 5, 62, 32, 2, 194, 195, 7, 49,
	2, 2, 195, 196, 8, 11, 1, 2, 196, 198, 3, 2, 2, 2, 197, 188, 3, 2, 2, 2,
	197, 192, 3, 2, 2, 2, 198, 21, 3, 2, 2, 2, 199, 200, 7, 18, 2, 2, 200,
	201, 7, 19, 2, 2, 201, 202, 7, 25, 2, 2, 202, 203, 7, 16, 2, 2, 203, 204,
	7, 3, 2, 2, 204, 205, 7, 20, 2, 2, 205, 206, 7, 13, 2, 2, 206, 207, 7,
	7, 2, 2, 207, 208, 7, 8, 2, 2, 208, 209, 7, 4, 2, 2, 209, 210, 5, 46, 24,
	2, 210, 211, 8, 12, 1, 2, 211, 23, 3, 2, 2, 2, 212, 214, 5, 26, 14, 2,
	213, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215,
	216, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 8, 13, 1, 2, 218, 25,
	3, 2, 2, 2, 219, 220, 5, 40, 21, 2, 220, 221, 8, 14, 1, 2, 221, 251, 3,
	2, 2, 2, 222, 223, 5, 48, 25, 2, 223, 224, 7, 30, 2, 2, 224, 225, 8, 14,
	1, 2, 225, 251, 3, 2, 2, 2, 226, 227, 5, 54, 28, 2, 227, 228, 7, 30, 2,
	2, 228, 229, 8, 14, 1, 2, 229, 251, 3, 2, 2, 2, 230, 231, 5, 56, 29, 2,
	231, 232, 7, 30, 2, 2, 232, 233, 8, 14, 1, 2, 233, 251, 3, 2, 2, 2, 234,
	235, 5, 50, 26, 2, 235, 236, 7, 30, 2, 2, 236, 237, 8, 14, 1, 2, 237, 251,
	3, 2, 2, 2, 238, 239, 5, 58, 30, 2, 239, 240, 7, 30, 2, 2, 240, 241, 8,
	14, 1, 2, 241, 251, 3, 2, 2, 2, 242, 243, 5, 30, 16, 2, 243, 244, 7, 30,
	2, 2, 244, 245, 8, 14, 1, 2, 245, 251, 3, 2, 2, 2, 246, 247, 5, 28, 15,
	2, 247, 248, 7, 30, 2, 2, 248, 249, 8, 14, 1, 2, 249, 251, 3, 2, 2, 2,
	250, 219, 3, 2, 2, 2, 250, 222, 3, 2, 2, 2, 250, 226, 3, 2, 2, 2, 250,
	230, 3, 2, 2, 2, 250, 234, 3, 2, 2, 2, 250, 238, 3, 2, 2, 2, 250, 242,
	3, 2, 2, 2, 250, 246, 3, 2, 2, 2, 251, 27, 3, 2, 2, 2, 252, 253, 7, 49,
	2, 2, 253, 254, 5, 60, 31, 2, 254, 255, 7, 34, 2, 2, 255, 256, 5, 64, 33,
	2, 256, 257, 8, 15, 1, 2, 257, 29, 3, 2, 2, 2, 258, 259, 5, 62, 32, 2,
	259, 260, 5, 32, 17, 2, 260, 261, 7, 49, 2, 2, 261, 262, 7, 34, 2, 2, 262,
	263, 5, 64, 33, 2, 263, 264, 8, 16, 1, 2, 264, 31, 3, 2, 2, 2, 265, 266,
	8, 17, 1, 2, 266, 267, 5, 34, 18, 2, 267, 268, 8, 17, 1, 2, 268, 275, 3,
	2, 2, 2, 269, 270, 12, 4, 2, 2, 270, 271, 5, 34, 18, 2, 271, 272, 8, 17,
	1, 2, 272, 274, 3, 2, 2, 2, 273, 269, 3, 2, 2, 2, 274, 277, 3, 2, 2, 2,
	275, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 33, 3, 2, 2, 2, 277, 275,
	3, 2, 2, 2, 278, 279, 7, 7, 2, 2, 279, 280, 7, 8, 2, 2, 280, 35, 3, 2,
	2, 2, 281, 282, 8, 19, 1, 2, 282, 283, 5, 38, 20, 2, 283, 284, 8, 19, 1,
	2, 284, 291, 3, 2, 2, 2, 285, 286, 12, 4, 2, 2, 286, 287, 5, 38, 20, 2,
	287, 288, 8, 19, 1, 2, 288, 290, 3, 2, 2, 2, 289, 285, 3, 2, 2, 2, 290,
	293, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 37, 3,
	2, 2, 2, 293, 291, 3, 2, 2, 2, 294, 295, 7, 7, 2, 2, 295, 296, 5, 64, 33,
	2, 296, 297, 7, 8, 2, 2, 297, 298, 8, 20, 1, 2, 298, 39, 3, 2, 2, 2, 299,
	300, 7, 11, 2, 2, 300, 301, 7, 3, 2, 2, 301, 302, 5, 64, 33, 2, 302, 303,
	7, 4, 2, 2, 303, 304, 5, 46, 24, 2, 304, 305, 8, 21, 1, 2, 305, 326, 3,
	2, 2, 2, 306, 307, 7, 11, 2, 2, 307, 308, 7, 3, 2, 2, 308, 309, 5, 64,
	33, 2, 309, 310, 7, 4, 2, 2, 310, 311, 5, 46, 24, 2, 311, 312, 7, 12, 2,
	2, 312, 313, 5, 46, 24, 2, 313, 314, 8, 21, 1, 2, 314, 326, 3, 2, 2, 2,
	315, 316, 7, 11, 2, 2, 316, 317, 7, 3, 2, 2, 317, 318, 5, 64, 33, 2, 318,
	319, 7, 4, 2, 2, 319, 320, 5, 46, 24, 2, 320, 321, 5, 42, 22, 2, 321, 322,
	7, 12, 2, 2, 322, 323, 5, 46, 24, 2, 323, 324, 8, 21, 1, 2, 324, 326, 3,
	2, 2, 2, 325, 299, 3, 2, 2, 2, 325, 306, 3, 2, 2, 2, 325, 315, 3, 2, 2,
	2, 326, 41, 3, 2, 2, 2, 327, 329, 5, 44, 23, 2, 328, 327, 3, 2, 2, 2, 329,
	330, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 332,
	3, 2, 2, 2, 332, 333, 8, 22, 1, 2, 333, 43, 3, 2, 2, 2, 334, 335, 7, 12,
	2, 2, 335, 336, 7, 11, 2, 2, 336, 337, 7, 3, 2, 2, 337, 338, 5, 64, 33,
	2, 338, 339, 7, 4, 2, 2, 339, 340, 5, 46, 24, 2, 340, 341, 8, 23, 1, 2,
	341, 45, 3, 2, 2, 2, 342, 343, 7, 5, 2, 2, 343, 344, 5, 24, 13, 2, 344,
	345, 7, 6, 2, 2, 345, 346, 8, 24, 1, 2, 346, 351, 3, 2, 2, 2, 347, 348,
	7, 5, 2, 2, 348, 349, 7, 6, 2, 2, 349, 351, 8, 24, 1, 2, 350, 342, 3, 2,
	2, 2, 350, 347, 3, 2, 2, 2, 351, 47, 3, 2, 2, 2, 352, 353, 7, 26, 2, 2,
	353, 354, 7, 28, 2, 2, 354, 355, 7, 9, 2, 2, 355, 356, 7, 28, 2, 2, 356,
	357, 7, 10, 2, 2, 357, 358, 7, 3, 2, 2, 358, 359, 5, 64, 33, 2, 359, 360,
	7, 4, 2, 2, 360, 361, 8, 25, 1, 2, 361, 49, 3, 2, 2, 2, 362, 363, 7, 49,
	2, 2, 363, 364, 7, 3, 2, 2, 364, 365, 7, 4, 2, 2, 365, 373, 8, 26, 1, 2,
	366, 367, 7, 49, 2, 2, 367, 368, 7, 3, 2, 2, 368, 369, 5, 52, 27, 2, 369,
	370, 7, 4, 2, 2, 370, 371, 8, 26, 1, 2, 371, 373, 3, 2, 2, 2, 372, 362,
	3, 2, 2, 2, 372, 366, 3, 2, 2, 2, 373, 51, 3, 2, 2, 2, 374, 375, 8, 27,
	1, 2, 375, 376, 5, 64, 33, 2, 376, 377, 8, 27, 1, 2, 377, 385, 3, 2, 2,
	2, 378, 379, 12, 4, 2, 2, 379, 380, 7, 29, 2, 2, 380, 381, 5, 64, 33, 2,
	381, 382, 8, 27, 1, 2, 382, 384, 3, 2, 2, 2, 383, 378, 3, 2, 2, 2, 384,
	387, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 53, 3,
	2, 2, 2, 387, 385, 3, 2, 2, 2, 388, 389, 5, 62, 32, 2, 389, 390, 5, 60,
	31, 2, 390, 391, 7, 34, 2, 2, 391, 392, 5, 64, 33, 2, 392, 393, 8, 28,
	1, 2, 393, 55, 3, 2, 2, 2, 394, 395, 5, 62, 32, 2, 395, 396, 5, 60, 31,
	2, 396, 397, 8, 29, 1, 2, 397, 57, 3, 2, 2, 2, 398, 399, 7, 21, 2, 2, 399,
	405, 8, 30, 1, 2, 400, 401, 7, 21, 2, 2, 401, 402, 5, 64, 33, 2, 402, 403,
	8, 30, 1, 2, 403, 405, 3, 2, 2, 2, 404, 398, 3, 2, 2, 2, 404, 400, 3, 2,
	2, 2, 405, 59, 3, 2, 2, 2, 406, 407, 8, 31, 1, 2, 407, 408, 7, 49, 2, 2,
	408, 409, 8, 31, 1, 2, 409, 416, 3, 2, 2, 2, 410, 411, 12, 4, 2, 2, 411,
	412, 7, 29, 2, 2, 412, 413, 7, 49, 2, 2, 413, 415, 8, 31, 1, 2, 414, 410,
	3, 2, 2, 2, 415, 418, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 416, 417, 3, 2,
	2, 2, 417, 61, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 419, 420, 7, 22, 2, 2,
	420, 430, 8, 32, 1, 2, 421, 422, 7, 24, 2, 2, 422, 430, 8, 32, 1, 2, 423,
	424, 7, 23, 2, 2, 424, 430, 8, 32, 1, 2, 425, 426, 7, 27, 2, 2, 426, 430,
	8, 32, 1, 2, 427, 428, 7, 25, 2, 2, 428, 430, 8, 32, 1, 2, 429, 419, 3,
	2, 2, 2, 429, 421, 3, 2, 2, 2, 429, 423, 3, 2, 2, 2, 429, 425, 3, 2, 2,
	2, 429, 427, 3, 2, 2, 2, 430, 63, 3, 2, 2, 2, 431, 432, 5, 80, 41, 2, 432,
	433, 8, 33, 1, 2, 433, 447, 3, 2, 2, 2, 434, 435, 5, 82, 42, 2, 435, 436,
	8, 33, 1, 2, 436, 447, 3, 2, 2, 2, 437, 438, 5, 68, 35, 2, 438, 439, 8,
	33, 1, 2, 439, 447, 3, 2, 2, 2, 440, 441, 5, 66, 34, 2, 441, 442, 8, 33,
	1, 2, 442, 447, 3, 2, 2, 2, 443, 444, 5, 70, 36, 2, 444, 445, 8, 33, 1,
	2, 445, 447, 3, 2, 2, 2, 446, 431, 3, 2, 2, 2, 446, 434, 3, 2, 2, 2, 446,
	437, 3, 2, 2, 2, 446, 440, 3, 2, 2, 2, 446, 443, 3, 2, 2, 2, 447, 65, 3,
	2, 2, 2, 448, 449, 7, 5, 2, 2, 449, 450, 5, 52, 27, 2, 450, 451, 7, 6,
	2, 2, 451, 452, 8, 34, 1, 2, 452, 67, 3, 2, 2, 2, 453, 454, 7, 15, 2, 2,
	454, 455, 5, 62, 32, 2, 455, 456, 5, 36, 19, 2, 456, 457, 8, 35, 1, 2,
	457, 69, 3, 2, 2, 2, 458, 459, 7, 15, 2, 2, 459, 460, 7, 49, 2, 2, 460,
	461, 7, 3, 2, 2, 461, 462, 7, 4, 2, 2, 462, 463, 8, 36, 1, 2, 463, 71,
	3, 2, 2, 2, 464, 465, 7, 49, 2, 2, 465, 466, 5, 36, 19, 2, 466, 467, 8,
	37, 1, 2, 467, 73, 3, 2, 2, 2, 468, 469, 5, 76, 39, 2, 469, 470, 8, 38,
	1, 2, 470, 75, 3, 2, 2, 2, 471, 472, 8, 39, 1, 2, 472, 473, 5, 78, 40,
	2, 473, 474, 8, 39, 1, 2, 474, 482, 3, 2, 2, 2, 475, 476, 12, 4, 2, 2,
	476, 477, 7, 28, 2, 2, 477, 478, 5, 78, 40, 2, 478, 479, 8, 39, 1, 2, 479,
	481, 3, 2, 2, 2, 480, 475, 3, 2, 2, 2, 481, 484, 3, 2, 2, 2, 482, 480,
	3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 77, 3, 2, 2, 2, 484, 482, 3, 2,
	2, 2, 485, 486, 7, 49, 2, 2, 486, 491, 8, 40, 1, 2, 487, 488, 5, 72, 37,
	2, 488, 489, 8, 40, 1, 2, 489, 491, 3, 2, 2, 2, 490, 485, 3, 2, 2, 2, 490,
	487, 3, 2, 2, 2, 491, 79, 3, 2, 2, 2, 492, 493, 8, 41, 1, 2, 493, 494,
	5, 82, 42, 2, 494, 495, 8, 41, 1, 2, 495, 503, 3, 2, 2, 2, 496, 497, 12,
	4, 2, 2, 497, 498, 9, 2, 2, 2, 498, 499, 5, 80, 41, 5, 499, 500, 8, 41,
	1, 2, 500, 502, 3, 2, 2, 2, 501, 496, 3, 2, 2, 2, 502, 505, 3, 2, 2, 2,
	503, 501, 3, 2, 2, 2, 503, 504, 3, 2, 2, 2, 504, 81, 3, 2, 2, 2, 505, 503,
	3, 2, 2, 2, 506, 507, 8, 42, 1, 2, 507, 508, 7, 43, 2, 2, 508, 509, 5,
	64, 33, 2, 509, 510, 8, 42, 1, 2, 510, 520, 3, 2, 2, 2, 511, 512, 5, 84,
	43, 2, 512, 513, 8, 42, 1, 2, 513, 520, 3, 2, 2, 2, 514, 515, 7, 3, 2,
	2, 515, 516, 5, 64, 33, 2, 516, 517, 7, 4, 2, 2, 517, 518, 8, 42, 1, 2,
	518, 520, 3, 2, 2, 2, 519, 506, 3, 2, 2, 2, 519, 511, 3, 2, 2, 2, 519,
	514, 3, 2, 2, 2, 520, 533, 3, 2, 2, 2, 521, 522, 12, 6, 2, 2, 522, 523,
	9, 3, 2, 2, 523, 524, 5, 82, 42, 7, 524, 525, 8, 42, 1, 2, 525, 532, 3,
	2, 2, 2, 526, 527, 12, 5, 2, 2, 527, 528, 9, 4, 2, 2, 528, 529, 5, 82,
	42, 6, 529, 530, 8, 42, 1, 2, 530, 532, 3, 2, 2, 2, 531, 521, 3, 2, 2,
	2, 531, 526, 3, 2, 2, 2, 532, 535, 3, 2, 2, 2, 533, 531, 3, 2, 2, 2, 533,
	534, 3, 2, 2, 2, 534, 83, 3, 2, 2, 2, 535, 533, 3, 2, 2, 2, 536, 537, 5,
	86, 44, 2, 537, 538, 8, 43, 1, 2, 538, 549, 3, 2, 2, 2, 539, 540, 5, 50,
	26, 2, 540, 541, 8, 43, 1, 2, 541, 549, 3, 2, 2, 2, 542, 543, 5, 72, 37,
	2, 543, 544, 8, 43, 1, 2, 544, 549, 3, 2, 2, 2, 545, 546, 5, 74, 38, 2,
	546, 547, 8, 43, 1, 2, 547, 549, 3, 2, 2, 2, 548, 536, 3, 2, 2, 2, 548,
	539, 3, 2, 2, 2, 548, 542, 3, 2, 2, 2, 548, 545, 3, 2, 2, 2, 549, 85, 3,
	2, 2, 2, 550, 551, 7, 44, 2, 2, 551, 563, 8, 44, 1, 2, 552, 553, 7, 45,
	2, 2, 553, 563, 8, 44, 1, 2, 554, 555, 7, 46, 2, 2, 555, 563, 8, 44, 1,
	2, 556, 557, 7, 49, 2, 2, 557, 563, 8, 44, 1, 2, 558, 559, 7, 47, 2, 2,
	559, 563, 8, 44, 1, 2, 560, 561, 7, 48, 2, 2, 561, 563, 8, 44, 1, 2, 562,
	550, 3, 2, 2, 2, 562, 552, 3, 2, 2, 2, 562, 554, 3, 2, 2, 2, 562, 556,
	3, 2, 2, 2, 562, 558, 3, 2, 2, 2, 562, 560, 3, 2, 2, 2, 563, 87, 3, 2,
	2, 2, 31, 101, 117, 129, 143, 165, 172, 185, 197, 215, 250, 275, 291, 325,
	330, 350, 372, 385, 404, 416, 429, 446, 482, 490, 503, 519, 531, 533, 548,
	562,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'{'", "'}'", "'['", "']'", "'out'", "'println'", "'if'",
	"'else'", "'args'", "'class'", "'new'", "'main'", "'private'", "'public'",
	"'static'", "'String'", "'return'", "'int'", "'float'", "'string'", "'void'",
	"'System'", "'boolean'", "'.'", "','", "';'", "'&&'", "'||'", "'!'", "'='",
	"'!='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "", "",
	"", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "LP", "RP", "L_LLAVE", "R_LLAVE", "L_CORCH", "R_CORCH", "OUT", "PRINTLN",
	"IF_TOK", "ELSE", "ARGS", "CLASS", "NEW_", "MAIN", "PRIVATE", "PUBLIC",
	"STATIC", "STRINGARGS", "RETURN_P", "INTTYPE", "FLOATTYPE", "STRINGTYPE",
	"VOIDTYPE", "SYSTEM", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND", "OR",
	"NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR",
	"MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "TRUE", "FALSE",
	"ID", "WHITESPACE",
}

var ruleNames = []string{
	"start", "listaClases", "clases", "cuerpoClase", "contenidoClase", "itemClase",
	"funcionItem", "modaccess", "parametros", "parametro", "funcmain", "instrucciones",
	"instruccion", "dec_objeto", "dec_arr", "dimensiones", "dimension", "listanchos",
	"ancho", "if_instr", "listaelseif", "else_if", "bloque", "consola", "llamada",
	"listaExpresiones", "declaracionIni", "declaracion", "retorno", "listides",
	"tiposvars", "expression", "arraydata", "instancia", "instanciaClase",
	"accesoarr", "accesoObjeto", "listaAccesos", "acceso", "expr_rel", "expr_arit",
	"expr_valor", "primitivo",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Calc struct {
	*antlr.BaseParser
}

func NewCalc(input antlr.TokenStream) *Calc {
	this := new(Calc)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calc.g4"

	return this
}

// Calc tokens.
const (
	CalcEOF        = antlr.TokenEOF
	CalcLP         = 1
	CalcRP         = 2
	CalcL_LLAVE    = 3
	CalcR_LLAVE    = 4
	CalcL_CORCH    = 5
	CalcR_CORCH    = 6
	CalcOUT        = 7
	CalcPRINTLN    = 8
	CalcIF_TOK     = 9
	CalcELSE       = 10
	CalcARGS       = 11
	CalcCLASS      = 12
	CalcNEW_       = 13
	CalcMAIN       = 14
	CalcPRIVATE    = 15
	CalcPUBLIC     = 16
	CalcSTATIC     = 17
	CalcSTRINGARGS = 18
	CalcRETURN_P   = 19
	CalcINTTYPE    = 20
	CalcFLOATTYPE  = 21
	CalcSTRINGTYPE = 22
	CalcVOIDTYPE   = 23
	CalcSYSTEM     = 24
	CalcBOOLTYPE   = 25
	CalcPUNTO      = 26
	CalcCOMA       = 27
	CalcPTCOMA     = 28
	CalcAND        = 29
	CalcOR         = 30
	CalcNOT        = 31
	CalcIGUAL      = 32
	CalcDIFERENTE  = 33
	CalcMAYORIGUAL = 34
	CalcMENORIGUAL = 35
	CalcMAYOR      = 36
	CalcMENOR      = 37
	CalcMUL        = 38
	CalcDIV        = 39
	CalcADD        = 40
	CalcSUB        = 41
	CalcNUMBER     = 42
	CalcFLOAT      = 43
	CalcSTRING     = 44
	CalcTRUE       = 45
	CalcFALSE      = 46
	CalcID         = 47
	CalcWHITESPACE = 48
)

// Calc rules.
const (
	CalcRULE_start            = 0
	CalcRULE_listaClases      = 1
	CalcRULE_clases           = 2
	CalcRULE_cuerpoClase      = 3
	CalcRULE_contenidoClase   = 4
	CalcRULE_itemClase        = 5
	CalcRULE_funcionItem      = 6
	CalcRULE_modaccess        = 7
	CalcRULE_parametros       = 8
	CalcRULE_parametro        = 9
	CalcRULE_funcmain         = 10
	CalcRULE_instrucciones    = 11
	CalcRULE_instruccion      = 12
	CalcRULE_dec_objeto       = 13
	CalcRULE_dec_arr          = 14
	CalcRULE_dimensiones      = 15
	CalcRULE_dimension        = 16
	CalcRULE_listanchos       = 17
	CalcRULE_ancho            = 18
	CalcRULE_if_instr         = 19
	CalcRULE_listaelseif      = 20
	CalcRULE_else_if          = 21
	CalcRULE_bloque           = 22
	CalcRULE_consola          = 23
	CalcRULE_llamada          = 24
	CalcRULE_listaExpresiones = 25
	CalcRULE_declaracionIni   = 26
	CalcRULE_declaracion      = 27
	CalcRULE_retorno          = 28
	CalcRULE_listides         = 29
	CalcRULE_tiposvars        = 30
	CalcRULE_expression       = 31
	CalcRULE_arraydata        = 32
	CalcRULE_instancia        = 33
	CalcRULE_instanciaClase   = 34
	CalcRULE_accesoarr        = 35
	CalcRULE_accesoObjeto     = 36
	CalcRULE_listaAccesos     = 37
	CalcRULE_acceso           = 38
	CalcRULE_expr_rel         = 39
	CalcRULE_expr_arit        = 40
	CalcRULE_expr_valor       = 41
	CalcRULE_primitivo        = 42
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listaClases returns the _listaClases rule contexts.
	Get_listaClases() IListaClasesContext

	// Set_listaClases sets the _listaClases rule contexts.
	Set_listaClases(IListaClasesContext)

	// GetAst returns the ast attribute.
	GetAst() ast.Ast

	// SetAst sets the ast attribute.
	SetAst(ast.Ast)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	ast          ast.Ast
	_listaClases IListaClasesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_listaClases() IListaClasesContext { return s._listaClases }

func (s *StartContext) Set_listaClases(v IListaClasesContext) { s._listaClases = v }

func (s *StartContext) GetAst() ast.Ast { return s.ast }

func (s *StartContext) SetAst(v ast.Ast) { s.ast = v }

func (s *StartContext) ListaClases() IListaClasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaClasesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaClasesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Calc) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)

		var _x = p.listaClases(0)

		localctx.(*StartContext)._listaClases = _x
	}
	localctx.(*StartContext).ast = ast.NewAst(localctx.(*StartContext).Get_listaClases().GetLista())

	return localctx
}

// IListaClasesContext is an interface to support dynamic dispatch.
type IListaClasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSUBLISTA returns the SUBLISTA rule contexts.
	GetSUBLISTA() IListaClasesContext

	// Get_clases returns the _clases rule contexts.
	Get_clases() IClasesContext

	// SetSUBLISTA sets the SUBLISTA rule contexts.
	SetSUBLISTA(IListaClasesContext)

	// Set_clases sets the _clases rule contexts.
	Set_clases(IClasesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListaClasesContext differentiates from other interfaces.
	IsListaClasesContext()
}

type ListaClasesContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	lista    *arrayList.List
	SUBLISTA IListaClasesContext
	_clases  IClasesContext
}

func NewEmptyListaClasesContext() *ListaClasesContext {
	var p = new(ListaClasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_listaClases
	return p
}

func (*ListaClasesContext) IsListaClasesContext() {}

func NewListaClasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaClasesContext {
	var p = new(ListaClasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_listaClases

	return p
}

func (s *ListaClasesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaClasesContext) GetSUBLISTA() IListaClasesContext { return s.SUBLISTA }

func (s *ListaClasesContext) Get_clases() IClasesContext { return s._clases }

func (s *ListaClasesContext) SetSUBLISTA(v IListaClasesContext) { s.SUBLISTA = v }

func (s *ListaClasesContext) Set_clases(v IClasesContext) { s._clases = v }

func (s *ListaClasesContext) GetLista() *arrayList.List { return s.lista }

func (s *ListaClasesContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListaClasesContext) Clases() IClasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClasesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClasesContext)
}

func (s *ListaClasesContext) ListaClases() IListaClasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaClasesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaClasesContext)
}

func (s *ListaClasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaClasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaClasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterListaClases(s)
	}
}

func (s *ListaClasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitListaClases(s)
	}
}

func (p *Calc) ListaClases() (localctx IListaClasesContext) {
	return p.listaClases(0)
}

func (p *Calc) listaClases(_p int) (localctx IListaClasesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaClasesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaClasesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalcRULE_listaClases, _p)

	localctx.(*ListaClasesContext).lista = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)

		var _x = p.Clases()

		localctx.(*ListaClasesContext)._clases = _x
	}
	localctx.(*ListaClasesContext).lista.Add(localctx.(*ListaClasesContext).Get_clases().GetInstr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaClasesContext(p, _parentctx, _parentState)
			localctx.(*ListaClasesContext).SUBLISTA = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_listaClases)
			p.SetState(93)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(94)

				var _x = p.Clases()

				localctx.(*ListaClasesContext)._clases = _x
			}

			localctx.(*ListaClasesContext).GetSUBLISTA().GetLista().Add(localctx.(*ListaClasesContext).Get_clases().GetInstr())
			localctx.(*ListaClasesContext).lista = localctx.(*ListaClasesContext).GetSUBLISTA().GetLista()

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IClasesContext is an interface to support dynamic dispatch.
type IClasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_cuerpoClase returns the _cuerpoClase rule contexts.
	Get_cuerpoClase() ICuerpoClaseContext

	// Set_cuerpoClase sets the _cuerpoClase rule contexts.
	Set_cuerpoClase(ICuerpoClaseContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsClasesContext differentiates from other interfaces.
	IsClasesContext()
}

type ClasesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruccion
	_ID          antlr.Token
	_cuerpoClase ICuerpoClaseContext
}

func NewEmptyClasesContext() *ClasesContext {
	var p = new(ClasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_clases
	return p
}

func (*ClasesContext) IsClasesContext() {}

func NewClasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClasesContext {
	var p = new(ClasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_clases

	return p
}

func (s *ClasesContext) GetParser() antlr.Parser { return s.parser }

func (s *ClasesContext) Get_ID() antlr.Token { return s._ID }

func (s *ClasesContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ClasesContext) Get_cuerpoClase() ICuerpoClaseContext { return s._cuerpoClase }

func (s *ClasesContext) Set_cuerpoClase(v ICuerpoClaseContext) { s._cuerpoClase = v }

func (s *ClasesContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *ClasesContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *ClasesContext) CLASS() antlr.TerminalNode {
	return s.GetToken(CalcCLASS, 0)
}

func (s *ClasesContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *ClasesContext) CuerpoClase() ICuerpoClaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICuerpoClaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICuerpoClaseContext)
}

func (s *ClasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterClases(s)
	}
}

func (s *ClasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitClases(s)
	}
}

func (p *Calc) Clases() (localctx IClasesContext) {
	localctx = NewClasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcRULE_clases)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(CalcCLASS)
	}
	{
		p.SetState(103)

		var _m = p.Match(CalcID)

		localctx.(*ClasesContext)._ID = _m
	}
	{
		p.SetState(104)

		var _x = p.CuerpoClase()

		localctx.(*ClasesContext)._cuerpoClase = _x
	}
	localctx.(*ClasesContext).instr = definicion.NewDefClase((func() string {
		if localctx.(*ClasesContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ClasesContext).Get_ID().GetText()
		}
	}()), localctx.(*ClasesContext).Get_cuerpoClase().GetLista())

	return localctx
}

// ICuerpoClaseContext is an interface to support dynamic dispatch.
type ICuerpoClaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_contenidoClase returns the _contenidoClase rule contexts.
	Get_contenidoClase() IContenidoClaseContext

	// Set_contenidoClase sets the _contenidoClase rule contexts.
	Set_contenidoClase(IContenidoClaseContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsCuerpoClaseContext differentiates from other interfaces.
	IsCuerpoClaseContext()
}

type CuerpoClaseContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	lista           *arrayList.List
	_contenidoClase IContenidoClaseContext
}

func NewEmptyCuerpoClaseContext() *CuerpoClaseContext {
	var p = new(CuerpoClaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_cuerpoClase
	return p
}

func (*CuerpoClaseContext) IsCuerpoClaseContext() {}

func NewCuerpoClaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CuerpoClaseContext {
	var p = new(CuerpoClaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_cuerpoClase

	return p
}

func (s *CuerpoClaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CuerpoClaseContext) Get_contenidoClase() IContenidoClaseContext { return s._contenidoClase }

func (s *CuerpoClaseContext) Set_contenidoClase(v IContenidoClaseContext) { s._contenidoClase = v }

func (s *CuerpoClaseContext) GetLista() *arrayList.List { return s.lista }

func (s *CuerpoClaseContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *CuerpoClaseContext) L_LLAVE() antlr.TerminalNode {
	return s.GetToken(CalcL_LLAVE, 0)
}

func (s *CuerpoClaseContext) ContenidoClase() IContenidoClaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContenidoClaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContenidoClaseContext)
}

func (s *CuerpoClaseContext) R_LLAVE() antlr.TerminalNode {
	return s.GetToken(CalcR_LLAVE, 0)
}

func (s *CuerpoClaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CuerpoClaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CuerpoClaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterCuerpoClase(s)
	}
}

func (s *CuerpoClaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitCuerpoClase(s)
	}
}

func (p *Calc) CuerpoClase() (localctx ICuerpoClaseContext) {
	localctx = NewCuerpoClaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcRULE_cuerpoClase)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Match(CalcL_LLAVE)
		}
		{
			p.SetState(108)

			var _x = p.contenidoClase(0)

			localctx.(*CuerpoClaseContext)._contenidoClase = _x
		}
		{
			p.SetState(109)
			p.Match(CalcR_LLAVE)
		}
		localctx.(*CuerpoClaseContext).lista = localctx.(*CuerpoClaseContext).Get_contenidoClase().GetLista()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(CalcL_LLAVE)
		}
		{
			p.SetState(113)
			p.Match(CalcR_LLAVE)
		}
		localctx.(*CuerpoClaseContext).lista = arrayList.New()

	}

	return localctx
}

// IContenidoClaseContext is an interface to support dynamic dispatch.
type IContenidoClaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSUBLISTA returns the SUBLISTA rule contexts.
	GetSUBLISTA() IContenidoClaseContext

	// Get_itemClase returns the _itemClase rule contexts.
	Get_itemClase() IItemClaseContext

	// SetSUBLISTA sets the SUBLISTA rule contexts.
	SetSUBLISTA(IContenidoClaseContext)

	// Set_itemClase sets the _itemClase rule contexts.
	Set_itemClase(IItemClaseContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsContenidoClaseContext differentiates from other interfaces.
	IsContenidoClaseContext()
}

type ContenidoClaseContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	lista      *arrayList.List
	SUBLISTA   IContenidoClaseContext
	_itemClase IItemClaseContext
}

func NewEmptyContenidoClaseContext() *ContenidoClaseContext {
	var p = new(ContenidoClaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_contenidoClase
	return p
}

func (*ContenidoClaseContext) IsContenidoClaseContext() {}

func NewContenidoClaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContenidoClaseContext {
	var p = new(ContenidoClaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_contenidoClase

	return p
}

func (s *ContenidoClaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ContenidoClaseContext) GetSUBLISTA() IContenidoClaseContext { return s.SUBLISTA }

func (s *ContenidoClaseContext) Get_itemClase() IItemClaseContext { return s._itemClase }

func (s *ContenidoClaseContext) SetSUBLISTA(v IContenidoClaseContext) { s.SUBLISTA = v }

func (s *ContenidoClaseContext) Set_itemClase(v IItemClaseContext) { s._itemClase = v }

func (s *ContenidoClaseContext) GetLista() *arrayList.List { return s.lista }

func (s *ContenidoClaseContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ContenidoClaseContext) ItemClase() IItemClaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemClaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemClaseContext)
}

func (s *ContenidoClaseContext) ContenidoClase() IContenidoClaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContenidoClaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContenidoClaseContext)
}

func (s *ContenidoClaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContenidoClaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContenidoClaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterContenidoClase(s)
	}
}

func (s *ContenidoClaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitContenidoClase(s)
	}
}

func (p *Calc) ContenidoClase() (localctx IContenidoClaseContext) {
	return p.contenidoClase(0)
}

func (p *Calc) contenidoClase(_p int) (localctx IContenidoClaseContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewContenidoClaseContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IContenidoClaseContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, CalcRULE_contenidoClase, _p)

	localctx.(*ContenidoClaseContext).lista = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)

		var _x = p.ItemClase()

		localctx.(*ContenidoClaseContext)._itemClase = _x
	}

	localctx.(*ContenidoClaseContext).lista.Add(localctx.(*ContenidoClaseContext).Get_itemClase().GetInstr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewContenidoClaseContext(p, _parentctx, _parentState)
			localctx.(*ContenidoClaseContext).SUBLISTA = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_contenidoClase)
			p.SetState(121)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(122)

				var _x = p.ItemClase()

				localctx.(*ContenidoClaseContext)._itemClase = _x
			}

			localctx.(*ContenidoClaseContext).GetSUBLISTA().GetLista().Add(localctx.(*ContenidoClaseContext).Get_itemClase().GetInstr())
			localctx.(*ContenidoClaseContext).lista = localctx.(*ContenidoClaseContext).GetSUBLISTA().GetLista()

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IItemClaseContext is an interface to support dynamic dispatch.
type IItemClaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funcionItem returns the _funcionItem rule contexts.
	Get_funcionItem() IFuncionItemContext

	// Get_declaracionIni returns the _declaracionIni rule contexts.
	Get_declaracionIni() IDeclaracionIniContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Set_funcionItem sets the _funcionItem rule contexts.
	Set_funcionItem(IFuncionItemContext)

	// Set_declaracionIni sets the _declaracionIni rule contexts.
	Set_declaracionIni(IDeclaracionIniContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsItemClaseContext differentiates from other interfaces.
	IsItemClaseContext()
}

type ItemClaseContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	instr           interfaces.Instruccion
	_funcionItem    IFuncionItemContext
	_declaracionIni IDeclaracionIniContext
	_declaracion    IDeclaracionContext
}

func NewEmptyItemClaseContext() *ItemClaseContext {
	var p = new(ItemClaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_itemClase
	return p
}

func (*ItemClaseContext) IsItemClaseContext() {}

func NewItemClaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemClaseContext {
	var p = new(ItemClaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_itemClase

	return p
}

func (s *ItemClaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemClaseContext) Get_funcionItem() IFuncionItemContext { return s._funcionItem }

func (s *ItemClaseContext) Get_declaracionIni() IDeclaracionIniContext { return s._declaracionIni }

func (s *ItemClaseContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *ItemClaseContext) Set_funcionItem(v IFuncionItemContext) { s._funcionItem = v }

func (s *ItemClaseContext) Set_declaracionIni(v IDeclaracionIniContext) { s._declaracionIni = v }

func (s *ItemClaseContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *ItemClaseContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *ItemClaseContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *ItemClaseContext) FuncionItem() IFuncionItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncionItemContext)
}

func (s *ItemClaseContext) DeclaracionIni() IDeclaracionIniContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionIniContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionIniContext)
}

func (s *ItemClaseContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(CalcPTCOMA, 0)
}

func (s *ItemClaseContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *ItemClaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemClaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemClaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterItemClase(s)
	}
}

func (s *ItemClaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitItemClase(s)
	}
}

func (p *Calc) ItemClase() (localctx IItemClaseContext) {
	localctx = NewItemClaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalcRULE_itemClase)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)

			var _x = p.FuncionItem()

			localctx.(*ItemClaseContext)._funcionItem = _x
		}
		localctx.(*ItemClaseContext).instr = localctx.(*ItemClaseContext).Get_funcionItem().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)

			var _x = p.DeclaracionIni()

			localctx.(*ItemClaseContext)._declaracionIni = _x
		}
		{
			p.SetState(134)
			p.Match(CalcPTCOMA)
		}
		localctx.(*ItemClaseContext).instr = localctx.(*ItemClaseContext).Get_declaracionIni().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)

			var _x = p.Declaracion()

			localctx.(*ItemClaseContext)._declaracion = _x
		}
		{
			p.SetState(138)
			p.Match(CalcPTCOMA)
		}
		localctx.(*ItemClaseContext).instr = localctx.(*ItemClaseContext).Get_declaracion().GetInstr()

	}

	return localctx
}

// IFuncionItemContext is an interface to support dynamic dispatch.
type IFuncionItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_funcmain returns the _funcmain rule contexts.
	Get_funcmain() IFuncmainContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Get_tiposvars returns the _tiposvars rule contexts.
	Get_tiposvars() ITiposvarsContext

	// Get_parametros returns the _parametros rule contexts.
	Get_parametros() IParametrosContext

	// Set_funcmain sets the _funcmain rule contexts.
	Set_funcmain(IFuncmainContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// Set_tiposvars sets the _tiposvars rule contexts.
	Set_tiposvars(ITiposvarsContext)

	// Set_parametros sets the _parametros rule contexts.
	Set_parametros(IParametrosContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsFuncionItemContext differentiates from other interfaces.
	IsFuncionItemContext()
}

type FuncionItemContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruccion
	_funcmain   IFuncmainContext
	_ID         antlr.Token
	_bloque     IBloqueContext
	_tiposvars  ITiposvarsContext
	_parametros IParametrosContext
}

func NewEmptyFuncionItemContext() *FuncionItemContext {
	var p = new(FuncionItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_funcionItem
	return p
}

func (*FuncionItemContext) IsFuncionItemContext() {}

func NewFuncionItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionItemContext {
	var p = new(FuncionItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_funcionItem

	return p
}

func (s *FuncionItemContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionItemContext) Get_ID() antlr.Token { return s._ID }

func (s *FuncionItemContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FuncionItemContext) Get_funcmain() IFuncmainContext { return s._funcmain }

func (s *FuncionItemContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *FuncionItemContext) Get_tiposvars() ITiposvarsContext { return s._tiposvars }

func (s *FuncionItemContext) Get_parametros() IParametrosContext { return s._parametros }

func (s *FuncionItemContext) Set_funcmain(v IFuncmainContext) { s._funcmain = v }

func (s *FuncionItemContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *FuncionItemContext) Set_tiposvars(v ITiposvarsContext) { s._tiposvars = v }

func (s *FuncionItemContext) Set_parametros(v IParametrosContext) { s._parametros = v }

func (s *FuncionItemContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *FuncionItemContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *FuncionItemContext) Funcmain() IFuncmainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncmainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncmainContext)
}

func (s *FuncionItemContext) Modaccess() IModaccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModaccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModaccessContext)
}

func (s *FuncionItemContext) Tiposvars() ITiposvarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITiposvarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITiposvarsContext)
}

func (s *FuncionItemContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *FuncionItemContext) LP() antlr.TerminalNode {
	return s.GetToken(CalcLP, 0)
}

func (s *FuncionItemContext) RP() antlr.TerminalNode {
	return s.GetToken(CalcRP, 0)
}

func (s *FuncionItemContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncionItemContext) Parametros() IParametrosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *FuncionItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterFuncionItem(s)
	}
}

func (s *FuncionItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitFuncionItem(s)
	}
}

func (p *Calc) FuncionItem() (localctx IFuncionItemContext) {
	localctx = NewFuncionItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalcRULE_funcionItem)
	listaParams := arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)

			var _x = p.Funcmain()

			localctx.(*FuncionItemContext)._funcmain = _x
		}
		localctx.(*FuncionItemContext).instr = localctx.(*FuncionItemContext).Get_funcmain().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Modaccess()
		}
		{
			p.SetState(147)
			p.Tiposvars()
		}
		{
			p.SetState(148)

			var _m = p.Match(CalcID)

			localctx.(*FuncionItemContext)._ID = _m
		}
		{
			p.SetState(149)
			p.Match(CalcLP)
		}
		{
			p.SetState(150)
			p.Match(CalcRP)
		}
		{
			p.SetState(151)

			var _x = p.Bloque()

			localctx.(*FuncionItemContext)._bloque = _x
		}
		localctx.(*FuncionItemContext).instr = Simbolos.NewFuncion((func() string {
			if localctx.(*FuncionItemContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncionItemContext).Get_ID().GetText()
			}
		}()), listaParams, localctx.(*FuncionItemContext).Get_bloque().GetLista(), entorno.VOID)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Modaccess()
		}
		{
			p.SetState(155)

			var _x = p.Tiposvars()

			localctx.(*FuncionItemContext)._tiposvars = _x
		}
		{
			p.SetState(156)

			var _m = p.Match(CalcID)

			localctx.(*FuncionItemContext)._ID = _m
		}
		{
			p.SetState(157)
			p.Match(CalcLP)
		}
		{
			p.SetState(158)

			var _x = p.parametros(0)

			localctx.(*FuncionItemContext)._parametros = _x
		}
		{
			p.SetState(159)
			p.Match(CalcRP)
		}
		{
			p.SetState(160)

			var _x = p.Bloque()

			localctx.(*FuncionItemContext)._bloque = _x
		}
		localctx.(*FuncionItemContext).instr = Simbolos.NewFuncion((func() string {
			if localctx.(*FuncionItemContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncionItemContext).Get_ID().GetText()
			}
		}()), localctx.(*FuncionItemContext).Get_parametros().GetLista(), localctx.(*FuncionItemContext).Get_bloque().GetLista(), localctx.(*FuncionItemContext).Get_tiposvars().GetTipo())

	}

	return localctx
}

// IModaccessContext is an interface to support dynamic dispatch.
type IModaccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetModAccess returns the modAccess attribute.
	GetModAccess() entorno.TipoModAccess

	// SetModAccess sets the modAccess attribute.
	SetModAccess(entorno.TipoModAccess)

	// IsModaccessContext differentiates from other interfaces.
	IsModaccessContext()
}

type ModaccessContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	modAccess entorno.TipoModAccess
}

func NewEmptyModaccessContext() *ModaccessContext {
	var p = new(ModaccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_modaccess
	return p
}

func (*ModaccessContext) IsModaccessContext() {}

func NewModaccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModaccessContext {
	var p = new(ModaccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_modaccess

	return p
}

func (s *ModaccessContext) GetParser() antlr.Parser { return s.parser }

func (s *ModaccessContext) GetModAccess() entorno.TipoModAccess { return s.modAccess }

func (s *ModaccessContext) SetModAccess(v entorno.TipoModAccess) { s.modAccess = v }

func (s *ModaccessContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(CalcPUBLIC, 0)
}

func (s *ModaccessContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(CalcPRIVATE, 0)
}

func (s *ModaccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModaccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModaccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterModaccess(s)
	}
}

func (s *ModaccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitModaccess(s)
	}
}

func (p *Calc) Modaccess() (localctx IModaccessContext) {
	localctx = NewModaccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CalcRULE_modaccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcPUBLIC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(CalcPUBLIC)
		}
		localctx.(*ModaccessContext).modAccess = entorno.PUBLIC

	case CalcPRIVATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(CalcPRIVATE)
		}
		localctx.(*ModaccessContext).modAccess = entorno.PRIVATE

	case CalcINTTYPE, CalcFLOATTYPE, CalcSTRINGTYPE, CalcVOIDTYPE, CalcBOOLTYPE:
		p.EnterOuterAlt(localctx, 3)
		localctx.(*ModaccessContext).modAccess = entorno.PRIVATE

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSublista returns the sublista rule contexts.
	GetSublista() IParametrosContext

	// Get_parametro returns the _parametro rule contexts.
	Get_parametro() IParametroContext

	// SetSublista sets the sublista rule contexts.
	SetSublista(IParametrosContext)

	// Set_parametro sets the _parametro rule contexts.
	Set_parametro(IParametroContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	lista      *arrayList.List
	sublista   IParametrosContext
	_parametro IParametroContext
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_parametros
	return p
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) GetSublista() IParametrosContext { return s.sublista }

func (s *ParametrosContext) Get_parametro() IParametroContext { return s._parametro }

func (s *ParametrosContext) SetSublista(v IParametrosContext) { s.sublista = v }

func (s *ParametrosContext) Set_parametro(v IParametroContext) { s._parametro = v }

func (s *ParametrosContext) GetLista() *arrayList.List { return s.lista }

func (s *ParametrosContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ParametrosContext) Parametro() IParametroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametroContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *ParametrosContext) COMA() antlr.TerminalNode {
	return s.GetToken(CalcCOMA, 0)
}

func (s *ParametrosContext) Parametros() IParametrosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (p *Calc) Parametros() (localctx IParametrosContext) {
	return p.parametros(0)
}

func (p *Calc) parametros(_p int) (localctx IParametrosContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParametrosContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, CalcRULE_parametros, _p)

	localctx.(*ParametrosContext).lista = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)

		var _x = p.Parametro()

		localctx.(*ParametrosContext)._parametro = _x
	}

	localctx.(*ParametrosContext).lista.Add(localctx.(*ParametrosContext).Get_parametro().GetInstr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParametrosContext(p, _parentctx, _parentState)
			localctx.(*ParametrosContext).sublista = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_parametros)
			p.SetState(176)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(177)
				p.Match(CalcCOMA)
			}
			{
				p.SetState(178)

				var _x = p.Parametro()

				localctx.(*ParametrosContext)._parametro = _x
			}

			localctx.(*ParametrosContext).GetSublista().GetLista().Add(localctx.(*ParametrosContext).Get_parametro().GetInstr())
			localctx.(*ParametrosContext).lista = localctx.(*ParametrosContext).GetSublista().GetLista()

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_tiposvars returns the _tiposvars rule contexts.
	Get_tiposvars() ITiposvarsContext

	// Set_tiposvars sets the _tiposvars rule contexts.
	Set_tiposvars(ITiposvarsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	instr      interfaces.Instruccion
	_tiposvars ITiposvarsContext
	_ID        antlr.Token
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_parametro
	return p
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) Get_ID() antlr.Token { return s._ID }

func (s *ParametroContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ParametroContext) Get_tiposvars() ITiposvarsContext { return s._tiposvars }

func (s *ParametroContext) Set_tiposvars(v ITiposvarsContext) { s._tiposvars = v }

func (s *ParametroContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *ParametroContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *ParametroContext) Tiposvars() ITiposvarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITiposvarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITiposvarsContext)
}

func (s *ParametroContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *ParametroContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalcMUL, 0)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (p *Calc) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CalcRULE_parametro)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcINTTYPE, CalcFLOATTYPE, CalcSTRINGTYPE, CalcVOIDTYPE, CalcBOOLTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)

			var _x = p.Tiposvars()

			localctx.(*ParametroContext)._tiposvars = _x
		}
		{
			p.SetState(187)

			var _m = p.Match(CalcID)

			localctx.(*ParametroContext)._ID = _m
		}

		listaIdes := arrayList.New()
		listaIdes.Add(expresion.NewIdentificador((func() string {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParametroContext).Get_ID().GetText()
			}
		}())))
		localctx.(*ParametroContext).instr = definicion.NewDeclaracionParametro(listaIdes, localctx.(*ParametroContext).Get_tiposvars().GetTipo(), false)

	case CalcMUL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(CalcMUL)
		}
		{
			p.SetState(191)

			var _x = p.Tiposvars()

			localctx.(*ParametroContext)._tiposvars = _x
		}
		{
			p.SetState(192)

			var _m = p.Match(CalcID)

			localctx.(*ParametroContext)._ID = _m
		}

		listaIdes := arrayList.New()
		listaIdes.Add(expresion.NewIdentificador((func() string {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParametroContext).Get_ID().GetText()
			}
		}())))
		localctx.(*ParametroContext).instr = definicion.NewDeclaracionParametro(listaIdes, localctx.(*ParametroContext).Get_tiposvars().GetTipo(), true)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncmainContext is an interface to support dynamic dispatch.
type IFuncmainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsFuncmainContext differentiates from other interfaces.
	IsFuncmainContext()
}

type FuncmainContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	instr   interfaces.Instruccion
	_bloque IBloqueContext
}

func NewEmptyFuncmainContext() *FuncmainContext {
	var p = new(FuncmainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_funcmain
	return p
}

func (*FuncmainContext) IsFuncmainContext() {}

func NewFuncmainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncmainContext {
	var p = new(FuncmainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_funcmain

	return p
}

func (s *FuncmainContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncmainContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *FuncmainContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *FuncmainContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *FuncmainContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *FuncmainContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(CalcPUBLIC, 0)
}

func (s *FuncmainContext) STATIC() antlr.TerminalNode {
	return s.GetToken(CalcSTATIC, 0)
}

func (s *FuncmainContext) VOIDTYPE() antlr.TerminalNode {
	return s.GetToken(CalcVOIDTYPE, 0)
}

func (s *FuncmainContext) MAIN() antlr.TerminalNode {
	return s.GetToken(CalcMAIN, 0)
}

func (s *FuncmainContext) LP() antlr.TerminalNode {
	return s.GetToken(CalcLP, 0)
}

func (s *FuncmainContext) STRINGARGS() antlr.TerminalNode {
	return s.GetToken(CalcSTRINGARGS, 0)
}

func (s *FuncmainContext) ARGS() antlr.TerminalNode {
	return s.GetToken(CalcARGS, 0)
}

func (s *FuncmainContext) L_CORCH() antlr.TerminalNode {
	return s.GetToken(CalcL_CORCH, 0)
}

func (s *FuncmainContext) R_CORCH() antlr.TerminalNode {
	return s.GetToken(CalcR_CORCH, 0)
}

func (s *FuncmainContext) RP() antlr.TerminalNode {
	return s.GetToken(CalcRP, 0)
}

func (s *FuncmainContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncmainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncmainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncmainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterFuncmain(s)
	}
}

func (s *FuncmainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitFuncmain(s)
	}
}

func (p *Calc) Funcmain() (localctx IFuncmainContext) {
	localctx = NewFuncmainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CalcRULE_funcmain)
	listaParams := arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(CalcPUBLIC)
	}
	{
		p.SetState(198)
		p.Match(CalcSTATIC)
	}
	{
		p.SetState(199)
		p.Match(CalcVOIDTYPE)
	}
	{
		p.SetState(200)
		p.Match(CalcMAIN)
	}
	{
		p.SetState(201)
		p.Match(CalcLP)
	}
	{
		p.SetState(202)
		p.Match(CalcSTRINGARGS)
	}
	{
		p.SetState(203)
		p.Match(CalcARGS)
	}
	{
		p.SetState(204)
		p.Match(CalcL_CORCH)
	}
	{
		p.SetState(205)
		p.Match(CalcR_CORCH)
	}
	{
		p.SetState(206)
		p.Match(CalcRP)
	}
	{
		p.SetState(207)

		var _x = p.Bloque()

		localctx.(*FuncmainContext)._bloque = _x
	}
	localctx.(*FuncmainContext).instr = Simbolos.NewFuncion("main", listaParams, localctx.(*FuncmainContext).Get_bloque().GetLista(), entorno.VOID)

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lista        *arrayList.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetLista() *arrayList.List { return s.lista }

func (s *InstruccionesContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Calc) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CalcRULE_instrucciones)
	localctx.(*InstruccionesContext).lista = arrayList.New()
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalcIF_TOK)|(1<<CalcRETURN_P)|(1<<CalcINTTYPE)|(1<<CalcFLOATTYPE)|(1<<CalcSTRINGTYPE)|(1<<CalcVOIDTYPE)|(1<<CalcSYSTEM)|(1<<CalcBOOLTYPE))) != 0) || _la == CalcID {
		{
			p.SetState(210)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).lista.Add(e.GetInstr())
	}
	fmt.Printf("tipo %T", localctx.(*InstruccionesContext).GetE())

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_if_instr returns the _if_instr rule contexts.
	Get_if_instr() IIf_instrContext

	// Get_consola returns the _consola rule contexts.
	Get_consola() IConsolaContext

	// Get_declaracionIni returns the _declaracionIni rule contexts.
	Get_declaracionIni() IDeclaracionIniContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_llamada returns the _llamada rule contexts.
	Get_llamada() ILlamadaContext

	// Get_retorno returns the _retorno rule contexts.
	Get_retorno() IRetornoContext

	// Get_dec_arr returns the _dec_arr rule contexts.
	Get_dec_arr() IDec_arrContext

	// Get_dec_objeto returns the _dec_objeto rule contexts.
	Get_dec_objeto() IDec_objetoContext

	// Set_if_instr sets the _if_instr rule contexts.
	Set_if_instr(IIf_instrContext)

	// Set_consola sets the _consola rule contexts.
	Set_consola(IConsolaContext)

	// Set_declaracionIni sets the _declaracionIni rule contexts.
	Set_declaracionIni(IDeclaracionIniContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_llamada sets the _llamada rule contexts.
	Set_llamada(ILlamadaContext)

	// Set_retorno sets the _retorno rule contexts.
	Set_retorno(IRetornoContext)

	// Set_dec_arr sets the _dec_arr rule contexts.
	Set_dec_arr(IDec_arrContext)

	// Set_dec_objeto sets the _dec_objeto rule contexts.
	Set_dec_objeto(IDec_objetoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	instr           interfaces.Instruccion
	_if_instr       IIf_instrContext
	_consola        IConsolaContext
	_declaracionIni IDeclaracionIniContext
	_declaracion    IDeclaracionContext
	_llamada        ILlamadaContext
	_retorno        IRetornoContext
	_dec_arr        IDec_arrContext
	_dec_objeto     IDec_objetoContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_if_instr() IIf_instrContext { return s._if_instr }

func (s *InstruccionContext) Get_consola() IConsolaContext { return s._consola }

func (s *InstruccionContext) Get_declaracionIni() IDeclaracionIniContext { return s._declaracionIni }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_llamada() ILlamadaContext { return s._llamada }

func (s *InstruccionContext) Get_retorno() IRetornoContext { return s._retorno }

func (s *InstruccionContext) Get_dec_arr() IDec_arrContext { return s._dec_arr }

func (s *InstruccionContext) Get_dec_objeto() IDec_objetoContext { return s._dec_objeto }

func (s *InstruccionContext) Set_if_instr(v IIf_instrContext) { s._if_instr = v }

func (s *InstruccionContext) Set_consola(v IConsolaContext) { s._consola = v }

func (s *InstruccionContext) Set_declaracionIni(v IDeclaracionIniContext) { s._declaracionIni = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_llamada(v ILlamadaContext) { s._llamada = v }

func (s *InstruccionContext) Set_retorno(v IRetornoContext) { s._retorno = v }

func (s *InstruccionContext) Set_dec_arr(v IDec_arrContext) { s._dec_arr = v }

func (s *InstruccionContext) Set_dec_objeto(v IDec_objetoContext) { s._dec_objeto = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *InstruccionContext) If_instr() IIf_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_instrContext)
}

func (s *InstruccionContext) Consola() IConsolaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConsolaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConsolaContext)
}

func (s *InstruccionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(CalcPTCOMA, 0)
}

func (s *InstruccionContext) DeclaracionIni() IDeclaracionIniContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionIniContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionIniContext)
}

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) Llamada() ILlamadaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILlamadaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILlamadaContext)
}

func (s *InstruccionContext) Retorno() IRetornoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetornoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetornoContext)
}

func (s *InstruccionContext) Dec_arr() IDec_arrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDec_arrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDec_arrContext)
}

func (s *InstruccionContext) Dec_objeto() IDec_objetoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDec_objetoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDec_objetoContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Calc) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CalcRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)

			var _x = p.If_instr()

			localctx.(*InstruccionContext)._if_instr = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_if_instr().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)

			var _x = p.Consola()

			localctx.(*InstruccionContext)._consola = _x
		}
		{
			p.SetState(221)
			p.Match(CalcPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_consola().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)

			var _x = p.DeclaracionIni()

			localctx.(*InstruccionContext)._declaracionIni = _x
		}
		{
			p.SetState(225)
			p.Match(CalcPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_declaracionIni().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(228)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(229)
			p.Match(CalcPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_declaracion().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(232)

			var _x = p.Llamada()

			localctx.(*InstruccionContext)._llamada = _x
		}
		{
			p.SetState(233)
			p.Match(CalcPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_llamada().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(236)

			var _x = p.Retorno()

			localctx.(*InstruccionContext)._retorno = _x
		}
		{
			p.SetState(237)
			p.Match(CalcPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_retorno().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(240)

			var _x = p.Dec_arr()

			localctx.(*InstruccionContext)._dec_arr = _x
		}
		{
			p.SetState(241)
			p.Match(CalcPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_dec_arr().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(244)

			var _x = p.Dec_objeto()

			localctx.(*InstruccionContext)._dec_objeto = _x
		}
		{
			p.SetState(245)
			p.Match(CalcPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_dec_objeto().GetInstr()

	}

	return localctx
}

// IDec_objetoContext is an interface to support dynamic dispatch.
type IDec_objetoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listides returns the _listides rule contexts.
	Get_listides() IListidesContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_listides sets the _listides rule contexts.
	Set_listides(IListidesContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsDec_objetoContext differentiates from other interfaces.
	IsDec_objetoContext()
}

type Dec_objetoContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruccion
	_ID         antlr.Token
	_listides   IListidesContext
	_expression IExpressionContext
}

func NewEmptyDec_objetoContext() *Dec_objetoContext {
	var p = new(Dec_objetoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_dec_objeto
	return p
}

func (*Dec_objetoContext) IsDec_objetoContext() {}

func NewDec_objetoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_objetoContext {
	var p = new(Dec_objetoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_dec_objeto

	return p
}

func (s *Dec_objetoContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_objetoContext) Get_ID() antlr.Token { return s._ID }

func (s *Dec_objetoContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Dec_objetoContext) Get_listides() IListidesContext { return s._listides }

func (s *Dec_objetoContext) Get_expression() IExpressionContext { return s._expression }

func (s *Dec_objetoContext) Set_listides(v IListidesContext) { s._listides = v }

func (s *Dec_objetoContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Dec_objetoContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *Dec_objetoContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *Dec_objetoContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *Dec_objetoContext) Listides() IListidesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListidesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListidesContext)
}

func (s *Dec_objetoContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(CalcIGUAL, 0)
}

func (s *Dec_objetoContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Dec_objetoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_objetoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dec_objetoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDec_objeto(s)
	}
}

func (s *Dec_objetoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDec_objeto(s)
	}
}

func (p *Calc) Dec_objeto() (localctx IDec_objetoContext) {
	localctx = NewDec_objetoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CalcRULE_dec_objeto)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)

		var _m = p.Match(CalcID)

		localctx.(*Dec_objetoContext)._ID = _m
	}
	{
		p.SetState(251)

		var _x = p.listides(0)

		localctx.(*Dec_objetoContext)._listides = _x
	}
	{
		p.SetState(252)
		p.Match(CalcIGUAL)
	}
	{
		p.SetState(253)

		var _x = p.Expression()

		localctx.(*Dec_objetoContext)._expression = _x
	}
	localctx.(*Dec_objetoContext).instr = defobjetos.NewDeclararObjeto((func() string {
		if localctx.(*Dec_objetoContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Dec_objetoContext).Get_ID().GetText()
		}
	}()), localctx.(*Dec_objetoContext).Get_listides().GetLista(), localctx.(*Dec_objetoContext).Get_expression().GetExpr())

	return localctx
}

// IDec_arrContext is an interface to support dynamic dispatch.
type IDec_arrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_tiposvars returns the _tiposvars rule contexts.
	Get_tiposvars() ITiposvarsContext

	// Get_dimensiones returns the _dimensiones rule contexts.
	Get_dimensiones() IDimensionesContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_tiposvars sets the _tiposvars rule contexts.
	Set_tiposvars(ITiposvarsContext)

	// Set_dimensiones sets the _dimensiones rule contexts.
	Set_dimensiones(IDimensionesContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsDec_arrContext differentiates from other interfaces.
	IsDec_arrContext()
}

type Dec_arrContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruccion
	_tiposvars   ITiposvarsContext
	_dimensiones IDimensionesContext
	_ID          antlr.Token
	_expression  IExpressionContext
}

func NewEmptyDec_arrContext() *Dec_arrContext {
	var p = new(Dec_arrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_dec_arr
	return p
}

func (*Dec_arrContext) IsDec_arrContext() {}

func NewDec_arrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_arrContext {
	var p = new(Dec_arrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_dec_arr

	return p
}

func (s *Dec_arrContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_arrContext) Get_ID() antlr.Token { return s._ID }

func (s *Dec_arrContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Dec_arrContext) Get_tiposvars() ITiposvarsContext { return s._tiposvars }

func (s *Dec_arrContext) Get_dimensiones() IDimensionesContext { return s._dimensiones }

func (s *Dec_arrContext) Get_expression() IExpressionContext { return s._expression }

func (s *Dec_arrContext) Set_tiposvars(v ITiposvarsContext) { s._tiposvars = v }

func (s *Dec_arrContext) Set_dimensiones(v IDimensionesContext) { s._dimensiones = v }

func (s *Dec_arrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Dec_arrContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *Dec_arrContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *Dec_arrContext) Tiposvars() ITiposvarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITiposvarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITiposvarsContext)
}

func (s *Dec_arrContext) Dimensiones() IDimensionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimensionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimensionesContext)
}

func (s *Dec_arrContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *Dec_arrContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(CalcIGUAL, 0)
}

func (s *Dec_arrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Dec_arrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_arrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dec_arrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDec_arr(s)
	}
}

func (s *Dec_arrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDec_arr(s)
	}
}

func (p *Calc) Dec_arr() (localctx IDec_arrContext) {
	localctx = NewDec_arrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CalcRULE_dec_arr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)

		var _x = p.Tiposvars()

		localctx.(*Dec_arrContext)._tiposvars = _x
	}
	{
		p.SetState(257)

		var _x = p.dimensiones(0)

		localctx.(*Dec_arrContext)._dimensiones = _x
	}
	{
		p.SetState(258)

		var _m = p.Match(CalcID)

		localctx.(*Dec_arrContext)._ID = _m
	}
	{
		p.SetState(259)
		p.Match(CalcIGUAL)
	}
	{
		p.SetState(260)

		var _x = p.Expression()

		localctx.(*Dec_arrContext)._expression = _x
	}
	localctx.(*Dec_arrContext).instr = defarreglos.NewDeclaracionArray(localctx.(*Dec_arrContext).Get_dimensiones().GetTam(), (func() string {
		if localctx.(*Dec_arrContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Dec_arrContext).Get_ID().GetText()
		}
	}()), localctx.(*Dec_arrContext).Get_expression().GetExpr(), localctx.(*Dec_arrContext).Get_tiposvars().GetTipo())

	return localctx
}

// IDimensionesContext is an interface to support dynamic dispatch.
type IDimensionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTamano returns the tamano rule contexts.
	GetTamano() IDimensionesContext

	// SetTamano sets the tamano rule contexts.
	SetTamano(IDimensionesContext)

	// GetTam returns the tam attribute.
	GetTam() int

	// SetTam sets the tam attribute.
	SetTam(int)

	// IsDimensionesContext differentiates from other interfaces.
	IsDimensionesContext()
}

type DimensionesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tam    int
	tamano IDimensionesContext
}

func NewEmptyDimensionesContext() *DimensionesContext {
	var p = new(DimensionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_dimensiones
	return p
}

func (*DimensionesContext) IsDimensionesContext() {}

func NewDimensionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimensionesContext {
	var p = new(DimensionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_dimensiones

	return p
}

func (s *DimensionesContext) GetParser() antlr.Parser { return s.parser }

func (s *DimensionesContext) GetTamano() IDimensionesContext { return s.tamano }

func (s *DimensionesContext) SetTamano(v IDimensionesContext) { s.tamano = v }

func (s *DimensionesContext) GetTam() int { return s.tam }

func (s *DimensionesContext) SetTam(v int) { s.tam = v }

func (s *DimensionesContext) Dimension() IDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimensionContext)
}

func (s *DimensionesContext) Dimensiones() IDimensionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimensionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimensionesContext)
}

func (s *DimensionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimensionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimensionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDimensiones(s)
	}
}

func (s *DimensionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDimensiones(s)
	}
}

func (p *Calc) Dimensiones() (localctx IDimensionesContext) {
	return p.dimensiones(0)
}

func (p *Calc) dimensiones(_p int) (localctx IDimensionesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDimensionesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDimensionesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, CalcRULE_dimensiones, _p)
	localctx.(*DimensionesContext).tam = 0

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Dimension()
	}
	localctx.(*DimensionesContext).tam = 1

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDimensionesContext(p, _parentctx, _parentState)
			localctx.(*DimensionesContext).tamano = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_dimensiones)
			p.SetState(267)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(268)
				p.Dimension()
			}

			localctx.(*DimensionesContext).tam = localctx.(*DimensionesContext).GetTamano().GetTam() + 1

		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IDimensionContext is an interface to support dynamic dispatch.
type IDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDimensionContext differentiates from other interfaces.
	IsDimensionContext()
}

type DimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimensionContext() *DimensionContext {
	var p = new(DimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_dimension
	return p
}

func (*DimensionContext) IsDimensionContext() {}

func NewDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimensionContext {
	var p = new(DimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_dimension

	return p
}

func (s *DimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *DimensionContext) L_CORCH() antlr.TerminalNode {
	return s.GetToken(CalcL_CORCH, 0)
}

func (s *DimensionContext) R_CORCH() antlr.TerminalNode {
	return s.GetToken(CalcR_CORCH, 0)
}

func (s *DimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDimension(s)
	}
}

func (s *DimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDimension(s)
	}
}

func (p *Calc) Dimension() (localctx IDimensionContext) {
	localctx = NewDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CalcRULE_dimension)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(CalcL_CORCH)
	}
	{
		p.SetState(277)
		p.Match(CalcR_CORCH)
	}

	return localctx
}

// IListanchosContext is an interface to support dynamic dispatch.
type IListanchosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSublist returns the sublist rule contexts.
	GetSublist() IListanchosContext

	// Get_ancho returns the _ancho rule contexts.
	Get_ancho() IAnchoContext

	// SetSublist sets the sublist rule contexts.
	SetSublist(IListanchosContext)

	// Set_ancho sets the _ancho rule contexts.
	Set_ancho(IAnchoContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListanchosContext differentiates from other interfaces.
	IsListanchosContext()
}

type ListanchosContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lista   *arrayList.List
	sublist IListanchosContext
	_ancho  IAnchoContext
}

func NewEmptyListanchosContext() *ListanchosContext {
	var p = new(ListanchosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_listanchos
	return p
}

func (*ListanchosContext) IsListanchosContext() {}

func NewListanchosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListanchosContext {
	var p = new(ListanchosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_listanchos

	return p
}

func (s *ListanchosContext) GetParser() antlr.Parser { return s.parser }

func (s *ListanchosContext) GetSublist() IListanchosContext { return s.sublist }

func (s *ListanchosContext) Get_ancho() IAnchoContext { return s._ancho }

func (s *ListanchosContext) SetSublist(v IListanchosContext) { s.sublist = v }

func (s *ListanchosContext) Set_ancho(v IAnchoContext) { s._ancho = v }

func (s *ListanchosContext) GetLista() *arrayList.List { return s.lista }

func (s *ListanchosContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListanchosContext) Ancho() IAnchoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnchoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnchoContext)
}

func (s *ListanchosContext) Listanchos() IListanchosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListanchosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListanchosContext)
}

func (s *ListanchosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListanchosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListanchosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterListanchos(s)
	}
}

func (s *ListanchosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitListanchos(s)
	}
}

func (p *Calc) Listanchos() (localctx IListanchosContext) {
	return p.listanchos(0)
}

func (p *Calc) listanchos(_p int) (localctx IListanchosContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListanchosContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListanchosContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, CalcRULE_listanchos, _p)

	localctx.(*ListanchosContext).lista = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)

		var _x = p.Ancho()

		localctx.(*ListanchosContext)._ancho = _x
	}
	localctx.(*ListanchosContext).lista.Add(localctx.(*ListanchosContext).Get_ancho().GetExpr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListanchosContext(p, _parentctx, _parentState)
			localctx.(*ListanchosContext).sublist = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_listanchos)
			p.SetState(283)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(284)

				var _x = p.Ancho()

				localctx.(*ListanchosContext)._ancho = _x
			}

			localctx.(*ListanchosContext).GetSublist().GetLista().Add(localctx.(*ListanchosContext).Get_ancho().GetExpr())
			localctx.(*ListanchosContext).lista = localctx.(*ListanchosContext).GetSublist().GetLista()

		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IAnchoContext is an interface to support dynamic dispatch.
type IAnchoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsAnchoContext differentiates from other interfaces.
	IsAnchoContext()
}

type AnchoContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	expr        interfaces.Expresion
	_expression IExpressionContext
}

func NewEmptyAnchoContext() *AnchoContext {
	var p = new(AnchoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_ancho
	return p
}

func (*AnchoContext) IsAnchoContext() {}

func NewAnchoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnchoContext {
	var p = new(AnchoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_ancho

	return p
}

func (s *AnchoContext) GetParser() antlr.Parser { return s.parser }

func (s *AnchoContext) Get_expression() IExpressionContext { return s._expression }

func (s *AnchoContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AnchoContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *AnchoContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *AnchoContext) L_CORCH() antlr.TerminalNode {
	return s.GetToken(CalcL_CORCH, 0)
}

func (s *AnchoContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AnchoContext) R_CORCH() antlr.TerminalNode {
	return s.GetToken(CalcR_CORCH, 0)
}

func (s *AnchoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnchoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnchoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAncho(s)
	}
}

func (s *AnchoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAncho(s)
	}
}

func (p *Calc) Ancho() (localctx IAnchoContext) {
	localctx = NewAnchoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CalcRULE_ancho)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(CalcL_CORCH)
	}
	{
		p.SetState(293)

		var _x = p.Expression()

		localctx.(*AnchoContext)._expression = _x
	}
	{
		p.SetState(294)
		p.Match(CalcR_CORCH)
	}
	localctx.(*AnchoContext).expr = localctx.(*AnchoContext).Get_expression().GetExpr()

	return localctx
}

// IIf_instrContext is an interface to support dynamic dispatch.
type IIf_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// GetBprincipal returns the bprincipal rule contexts.
	GetBprincipal() IBloqueContext

	// GetBelse returns the belse rule contexts.
	GetBelse() IBloqueContext

	// Get_listaelseif returns the _listaelseif rule contexts.
	Get_listaelseif() IListaelseifContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// SetBprincipal sets the bprincipal rule contexts.
	SetBprincipal(IBloqueContext)

	// SetBelse sets the belse rule contexts.
	SetBelse(IBloqueContext)

	// Set_listaelseif sets the _listaelseif rule contexts.
	Set_listaelseif(IListaelseifContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsIf_instrContext differentiates from other interfaces.
	IsIf_instrContext()
}

type If_instrContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruccion
	_expression  IExpressionContext
	_bloque      IBloqueContext
	bprincipal   IBloqueContext
	belse        IBloqueContext
	_listaelseif IListaelseifContext
}

func NewEmptyIf_instrContext() *If_instrContext {
	var p = new(If_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_if_instr
	return p
}

func (*If_instrContext) IsIf_instrContext() {}

func NewIf_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_instrContext {
	var p = new(If_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_if_instr

	return p
}

func (s *If_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *If_instrContext) Get_expression() IExpressionContext { return s._expression }

func (s *If_instrContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *If_instrContext) GetBprincipal() IBloqueContext { return s.bprincipal }

func (s *If_instrContext) GetBelse() IBloqueContext { return s.belse }

func (s *If_instrContext) Get_listaelseif() IListaelseifContext { return s._listaelseif }

func (s *If_instrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *If_instrContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *If_instrContext) SetBprincipal(v IBloqueContext) { s.bprincipal = v }

func (s *If_instrContext) SetBelse(v IBloqueContext) { s.belse = v }

func (s *If_instrContext) Set_listaelseif(v IListaelseifContext) { s._listaelseif = v }

func (s *If_instrContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *If_instrContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *If_instrContext) IF_TOK() antlr.TerminalNode {
	return s.GetToken(CalcIF_TOK, 0)
}

func (s *If_instrContext) LP() antlr.TerminalNode {
	return s.GetToken(CalcLP, 0)
}

func (s *If_instrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_instrContext) RP() antlr.TerminalNode {
	return s.GetToken(CalcRP, 0)
}

func (s *If_instrContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *If_instrContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *If_instrContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CalcELSE, 0)
}

func (s *If_instrContext) Listaelseif() IListaelseifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaelseifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaelseifContext)
}

func (s *If_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterIf_instr(s)
	}
}

func (s *If_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitIf_instr(s)
	}
}

func (p *Calc) If_instr() (localctx IIf_instrContext) {
	localctx = NewIf_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CalcRULE_if_instr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Match(CalcIF_TOK)
		}
		{
			p.SetState(298)
			p.Match(CalcLP)
		}
		{
			p.SetState(299)

			var _x = p.Expression()

			localctx.(*If_instrContext)._expression = _x
		}
		{
			p.SetState(300)
			p.Match(CalcRP)
		}
		{
			p.SetState(301)

			var _x = p.Bloque()

			localctx.(*If_instrContext)._bloque = _x
		}
		localctx.(*If_instrContext).instr = control.NewIfInstruccion(localctx.(*If_instrContext).Get_expression().GetExpr(), localctx.(*If_instrContext).Get_bloque().GetLista(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.Match(CalcIF_TOK)
		}
		{
			p.SetState(305)
			p.Match(CalcLP)
		}
		{
			p.SetState(306)

			var _x = p.Expression()

			localctx.(*If_instrContext)._expression = _x
		}
		{
			p.SetState(307)
			p.Match(CalcRP)
		}
		{
			p.SetState(308)

			var _x = p.Bloque()

			localctx.(*If_instrContext).bprincipal = _x
		}
		{
			p.SetState(309)
			p.Match(CalcELSE)
		}
		{
			p.SetState(310)

			var _x = p.Bloque()

			localctx.(*If_instrContext).belse = _x
		}
		localctx.(*If_instrContext).instr = control.NewIfInstruccion(localctx.(*If_instrContext).Get_expression().GetExpr(), localctx.(*If_instrContext).GetBprincipal().GetLista(), nil, localctx.(*If_instrContext).GetBelse().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(313)
			p.Match(CalcIF_TOK)
		}
		{
			p.SetState(314)
			p.Match(CalcLP)
		}
		{
			p.SetState(315)

			var _x = p.Expression()

			localctx.(*If_instrContext)._expression = _x
		}
		{
			p.SetState(316)
			p.Match(CalcRP)
		}
		{
			p.SetState(317)

			var _x = p.Bloque()

			localctx.(*If_instrContext).bprincipal = _x
		}
		{
			p.SetState(318)

			var _x = p.Listaelseif()

			localctx.(*If_instrContext)._listaelseif = _x
		}
		{
			p.SetState(319)
			p.Match(CalcELSE)
		}
		{
			p.SetState(320)

			var _x = p.Bloque()

			localctx.(*If_instrContext).belse = _x
		}

		localctx.(*If_instrContext).instr = control.NewIfInstruccion(localctx.(*If_instrContext).Get_expression().GetExpr(), localctx.(*If_instrContext).GetBprincipal().GetLista(), localctx.(*If_instrContext).Get_listaelseif().GetLista(), localctx.(*If_instrContext).GetBelse().GetLista())

	}

	return localctx
}

// IListaelseifContext is an interface to support dynamic dispatch.
type IListaelseifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_else_if returns the _else_if rule contexts.
	Get_else_if() IElse_ifContext

	// Set_else_if sets the _else_if rule contexts.
	Set_else_if(IElse_ifContext)

	// GetList returns the list rule context list.
	GetList() []IElse_ifContext

	// SetList sets the list rule context list.
	SetList([]IElse_ifContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListaelseifContext differentiates from other interfaces.
	IsListaelseifContext()
}

type ListaelseifContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	lista    *arrayList.List
	_else_if IElse_ifContext
	list     []IElse_ifContext
}

func NewEmptyListaelseifContext() *ListaelseifContext {
	var p = new(ListaelseifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_listaelseif
	return p
}

func (*ListaelseifContext) IsListaelseifContext() {}

func NewListaelseifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaelseifContext {
	var p = new(ListaelseifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_listaelseif

	return p
}

func (s *ListaelseifContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaelseifContext) Get_else_if() IElse_ifContext { return s._else_if }

func (s *ListaelseifContext) Set_else_if(v IElse_ifContext) { s._else_if = v }

func (s *ListaelseifContext) GetList() []IElse_ifContext { return s.list }

func (s *ListaelseifContext) SetList(v []IElse_ifContext) { s.list = v }

func (s *ListaelseifContext) GetLista() *arrayList.List { return s.lista }

func (s *ListaelseifContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListaelseifContext) AllElse_if() []IElse_ifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElse_ifContext)(nil)).Elem())
	var tst = make([]IElse_ifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElse_ifContext)
		}
	}

	return tst
}

func (s *ListaelseifContext) Else_if(i int) IElse_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_ifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElse_ifContext)
}

func (s *ListaelseifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaelseifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaelseifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterListaelseif(s)
	}
}

func (s *ListaelseifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitListaelseif(s)
	}
}

func (p *Calc) Listaelseif() (localctx IListaelseifContext) {
	localctx = NewListaelseifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CalcRULE_listaelseif)
	localctx.(*ListaelseifContext).lista = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(325)

				var _x = p.Else_if()

				localctx.(*ListaelseifContext)._else_if = _x
			}
			localctx.(*ListaelseifContext).list = append(localctx.(*ListaelseifContext).list, localctx.(*ListaelseifContext)._else_if)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	listInt := localctx.(*ListaelseifContext).GetList()
	for _, e := range listInt {
		localctx.(*ListaelseifContext).lista.Add(e.GetInstr())
	}

	return localctx
}

// IElse_ifContext is an interface to support dynamic dispatch.
type IElse_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsElse_ifContext differentiates from other interfaces.
	IsElse_ifContext()
}

type Else_ifContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruccion
	_expression IExpressionContext
	_bloque     IBloqueContext
}

func NewEmptyElse_ifContext() *Else_ifContext {
	var p = new(Else_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_else_if
	return p
}

func (*Else_ifContext) IsElse_ifContext() {}

func NewElse_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_ifContext {
	var p = new(Else_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_else_if

	return p
}

func (s *Else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_ifContext) Get_expression() IExpressionContext { return s._expression }

func (s *Else_ifContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *Else_ifContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Else_ifContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *Else_ifContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *Else_ifContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *Else_ifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CalcELSE, 0)
}

func (s *Else_ifContext) IF_TOK() antlr.TerminalNode {
	return s.GetToken(CalcIF_TOK, 0)
}

func (s *Else_ifContext) LP() antlr.TerminalNode {
	return s.GetToken(CalcLP, 0)
}

func (s *Else_ifContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Else_ifContext) RP() antlr.TerminalNode {
	return s.GetToken(CalcRP, 0)
}

func (s *Else_ifContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterElse_if(s)
	}
}

func (s *Else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitElse_if(s)
	}
}

func (p *Calc) Else_if() (localctx IElse_ifContext) {
	localctx = NewElse_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CalcRULE_else_if)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(CalcELSE)
	}
	{
		p.SetState(333)
		p.Match(CalcIF_TOK)
	}
	{
		p.SetState(334)
		p.Match(CalcLP)
	}
	{
		p.SetState(335)

		var _x = p.Expression()

		localctx.(*Else_ifContext)._expression = _x
	}
	{
		p.SetState(336)
		p.Match(CalcRP)
	}
	{
		p.SetState(337)

		var _x = p.Bloque()

		localctx.(*Else_ifContext)._bloque = _x
	}
	localctx.(*Else_ifContext).instr = control.NewIfInstruccion(localctx.(*Else_ifContext).Get_expression().GetExpr(), localctx.(*Else_ifContext).Get_bloque().GetLista(), nil, nil)

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *BloqueContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *BloqueContext) GetLista() *arrayList.List { return s.lista }

func (s *BloqueContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *BloqueContext) L_LLAVE() antlr.TerminalNode {
	return s.GetToken(CalcL_LLAVE, 0)
}

func (s *BloqueContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueContext) R_LLAVE() antlr.TerminalNode {
	return s.GetToken(CalcR_LLAVE, 0)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *Calc) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CalcRULE_bloque)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.Match(CalcL_LLAVE)
		}
		{
			p.SetState(341)

			var _x = p.Instrucciones()

			localctx.(*BloqueContext)._instrucciones = _x
		}
		{
			p.SetState(342)
			p.Match(CalcR_LLAVE)
		}
		localctx.(*BloqueContext).lista = localctx.(*BloqueContext).Get_instrucciones().GetLista()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.Match(CalcL_LLAVE)
		}
		{
			p.SetState(346)
			p.Match(CalcR_LLAVE)
		}
		localctx.(*BloqueContext).lista = arrayList.New()

	}

	return localctx
}

// IConsolaContext is an interface to support dynamic dispatch.
type IConsolaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsConsolaContext differentiates from other interfaces.
	IsConsolaContext()
}

type ConsolaContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruccion
	_expression IExpressionContext
}

func NewEmptyConsolaContext() *ConsolaContext {
	var p = new(ConsolaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_consola
	return p
}

func (*ConsolaContext) IsConsolaContext() {}

func NewConsolaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsolaContext {
	var p = new(ConsolaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_consola

	return p
}

func (s *ConsolaContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsolaContext) Get_expression() IExpressionContext { return s._expression }

func (s *ConsolaContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ConsolaContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *ConsolaContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *ConsolaContext) SYSTEM() antlr.TerminalNode {
	return s.GetToken(CalcSYSTEM, 0)
}

func (s *ConsolaContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(CalcPUNTO)
}

func (s *ConsolaContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(CalcPUNTO, i)
}

func (s *ConsolaContext) OUT() antlr.TerminalNode {
	return s.GetToken(CalcOUT, 0)
}

func (s *ConsolaContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(CalcPRINTLN, 0)
}

func (s *ConsolaContext) LP() antlr.TerminalNode {
	return s.GetToken(CalcLP, 0)
}

func (s *ConsolaContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConsolaContext) RP() antlr.TerminalNode {
	return s.GetToken(CalcRP, 0)
}

func (s *ConsolaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsolaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConsolaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterConsola(s)
	}
}

func (s *ConsolaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitConsola(s)
	}
}

func (p *Calc) Consola() (localctx IConsolaContext) {
	localctx = NewConsolaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CalcRULE_consola)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(CalcSYSTEM)
	}
	{
		p.SetState(351)
		p.Match(CalcPUNTO)
	}
	{
		p.SetState(352)
		p.Match(CalcOUT)
	}
	{
		p.SetState(353)
		p.Match(CalcPUNTO)
	}
	{
		p.SetState(354)
		p.Match(CalcPRINTLN)
	}
	{
		p.SetState(355)
		p.Match(CalcLP)
	}
	{
		p.SetState(356)

		var _x = p.Expression()

		localctx.(*ConsolaContext)._expression = _x
	}
	{
		p.SetState(357)
		p.Match(CalcRP)
	}
	localctx.(*ConsolaContext).instr = funbasica.NewImprimir(localctx.(*ConsolaContext).Get_expression().GetExpr())

	return localctx
}

// ILlamadaContext is an interface to support dynamic dispatch.
type ILlamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listaExpresiones returns the _listaExpresiones rule contexts.
	Get_listaExpresiones() IListaExpresionesContext

	// Set_listaExpresiones sets the _listaExpresiones rule contexts.
	Set_listaExpresiones(IListaExpresionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsLlamadaContext differentiates from other interfaces.
	IsLlamadaContext()
}

type LlamadaContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	instr             interfaces.Instruccion
	expr              interfaces.Expresion
	_ID               antlr.Token
	_listaExpresiones IListaExpresionesContext
}

func NewEmptyLlamadaContext() *LlamadaContext {
	var p = new(LlamadaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_llamada
	return p
}

func (*LlamadaContext) IsLlamadaContext() {}

func NewLlamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaContext {
	var p = new(LlamadaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_llamada

	return p
}

func (s *LlamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaContext) Get_ID() antlr.Token { return s._ID }

func (s *LlamadaContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *LlamadaContext) Get_listaExpresiones() IListaExpresionesContext { return s._listaExpresiones }

func (s *LlamadaContext) Set_listaExpresiones(v IListaExpresionesContext) { s._listaExpresiones = v }

func (s *LlamadaContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *LlamadaContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *LlamadaContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *LlamadaContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *LlamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *LlamadaContext) LP() antlr.TerminalNode {
	return s.GetToken(CalcLP, 0)
}

func (s *LlamadaContext) RP() antlr.TerminalNode {
	return s.GetToken(CalcRP, 0)
}

func (s *LlamadaContext) ListaExpresiones() IListaExpresionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaExpresionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaExpresionesContext)
}

func (s *LlamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterLlamada(s)
	}
}

func (s *LlamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitLlamada(s)
	}
}

func (p *Calc) Llamada() (localctx ILlamadaContext) {
	localctx = NewLlamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CalcRULE_llamada)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)

			var _m = p.Match(CalcID)

			localctx.(*LlamadaContext)._ID = _m
		}
		{
			p.SetState(361)
			p.Match(CalcLP)
		}
		{
			p.SetState(362)
			p.Match(CalcRP)
		}

		localctx.(*LlamadaContext).instr = expresion2.NewLlamada((func() string {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetText()
			}
		}()), arrayList.New())
		localctx.(*LlamadaContext).expr = expresion2.NewLlamada((func() string {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetText()
			}
		}()), arrayList.New())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(364)

			var _m = p.Match(CalcID)

			localctx.(*LlamadaContext)._ID = _m
		}
		{
			p.SetState(365)
			p.Match(CalcLP)
		}
		{
			p.SetState(366)

			var _x = p.listaExpresiones(0)

			localctx.(*LlamadaContext)._listaExpresiones = _x
		}
		{
			p.SetState(367)
			p.Match(CalcRP)
		}

		localctx.(*LlamadaContext).instr = expresion2.NewLlamada((func() string {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetText()
			}
		}()), localctx.(*LlamadaContext).Get_listaExpresiones().GetLista())
		localctx.(*LlamadaContext).expr = expresion2.NewLlamada((func() string {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetText()
			}
		}()), localctx.(*LlamadaContext).Get_listaExpresiones().GetLista())

	}

	return localctx
}

// IListaExpresionesContext is an interface to support dynamic dispatch.
type IListaExpresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLISTA returns the LISTA rule contexts.
	GetLISTA() IListaExpresionesContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetLISTA sets the LISTA rule contexts.
	SetLISTA(IListaExpresionesContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListaExpresionesContext differentiates from other interfaces.
	IsListaExpresionesContext()
}

type ListaExpresionesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	lista       *arrayList.List
	LISTA       IListaExpresionesContext
	_expression IExpressionContext
}

func NewEmptyListaExpresionesContext() *ListaExpresionesContext {
	var p = new(ListaExpresionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_listaExpresiones
	return p
}

func (*ListaExpresionesContext) IsListaExpresionesContext() {}

func NewListaExpresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaExpresionesContext {
	var p = new(ListaExpresionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_listaExpresiones

	return p
}

func (s *ListaExpresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaExpresionesContext) GetLISTA() IListaExpresionesContext { return s.LISTA }

func (s *ListaExpresionesContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListaExpresionesContext) SetLISTA(v IListaExpresionesContext) { s.LISTA = v }

func (s *ListaExpresionesContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListaExpresionesContext) GetLista() *arrayList.List { return s.lista }

func (s *ListaExpresionesContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListaExpresionesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListaExpresionesContext) COMA() antlr.TerminalNode {
	return s.GetToken(CalcCOMA, 0)
}

func (s *ListaExpresionesContext) ListaExpresiones() IListaExpresionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaExpresionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaExpresionesContext)
}

func (s *ListaExpresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaExpresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaExpresionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterListaExpresiones(s)
	}
}

func (s *ListaExpresionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitListaExpresiones(s)
	}
}

func (p *Calc) ListaExpresiones() (localctx IListaExpresionesContext) {
	return p.listaExpresiones(0)
}

func (p *Calc) listaExpresiones(_p int) (localctx IListaExpresionesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaExpresionesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaExpresionesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, CalcRULE_listaExpresiones, _p)

	localctx.(*ListaExpresionesContext).lista = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)

		var _x = p.Expression()

		localctx.(*ListaExpresionesContext)._expression = _x
	}

	localctx.(*ListaExpresionesContext).lista.Add(localctx.(*ListaExpresionesContext).Get_expression().GetExpr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaExpresionesContext(p, _parentctx, _parentState)
			localctx.(*ListaExpresionesContext).LISTA = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_listaExpresiones)
			p.SetState(376)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(377)
				p.Match(CalcCOMA)
			}
			{
				p.SetState(378)

				var _x = p.Expression()

				localctx.(*ListaExpresionesContext)._expression = _x
			}

			localctx.(*ListaExpresionesContext).GetLISTA().GetLista().Add(localctx.(*ListaExpresionesContext).Get_expression().GetExpr())
			localctx.(*ListaExpresionesContext).lista = localctx.(*ListaExpresionesContext).GetLISTA().GetLista()

		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclaracionIniContext is an interface to support dynamic dispatch.
type IDeclaracionIniContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_tiposvars returns the _tiposvars rule contexts.
	Get_tiposvars() ITiposvarsContext

	// Get_listides returns the _listides rule contexts.
	Get_listides() IListidesContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_tiposvars sets the _tiposvars rule contexts.
	Set_tiposvars(ITiposvarsContext)

	// Set_listides sets the _listides rule contexts.
	Set_listides(IListidesContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsDeclaracionIniContext differentiates from other interfaces.
	IsDeclaracionIniContext()
}

type DeclaracionIniContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruccion
	_tiposvars  ITiposvarsContext
	_listides   IListidesContext
	_expression IExpressionContext
}

func NewEmptyDeclaracionIniContext() *DeclaracionIniContext {
	var p = new(DeclaracionIniContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_declaracionIni
	return p
}

func (*DeclaracionIniContext) IsDeclaracionIniContext() {}

func NewDeclaracionIniContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionIniContext {
	var p = new(DeclaracionIniContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_declaracionIni

	return p
}

func (s *DeclaracionIniContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionIniContext) Get_tiposvars() ITiposvarsContext { return s._tiposvars }

func (s *DeclaracionIniContext) Get_listides() IListidesContext { return s._listides }

func (s *DeclaracionIniContext) Get_expression() IExpressionContext { return s._expression }

func (s *DeclaracionIniContext) Set_tiposvars(v ITiposvarsContext) { s._tiposvars = v }

func (s *DeclaracionIniContext) Set_listides(v IListidesContext) { s._listides = v }

func (s *DeclaracionIniContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *DeclaracionIniContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *DeclaracionIniContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *DeclaracionIniContext) Tiposvars() ITiposvarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITiposvarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITiposvarsContext)
}

func (s *DeclaracionIniContext) Listides() IListidesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListidesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListidesContext)
}

func (s *DeclaracionIniContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(CalcIGUAL, 0)
}

func (s *DeclaracionIniContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclaracionIniContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionIniContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionIniContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDeclaracionIni(s)
	}
}

func (s *DeclaracionIniContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDeclaracionIni(s)
	}
}

func (p *Calc) DeclaracionIni() (localctx IDeclaracionIniContext) {
	localctx = NewDeclaracionIniContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CalcRULE_declaracionIni)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)

		var _x = p.Tiposvars()

		localctx.(*DeclaracionIniContext)._tiposvars = _x
	}
	{
		p.SetState(387)

		var _x = p.listides(0)

		localctx.(*DeclaracionIniContext)._listides = _x
	}
	{
		p.SetState(388)
		p.Match(CalcIGUAL)
	}
	{
		p.SetState(389)

		var _x = p.Expression()

		localctx.(*DeclaracionIniContext)._expression = _x
	}

	localctx.(*DeclaracionIniContext).instr = definicion.NewDeclaracionInicializacion(localctx.(*DeclaracionIniContext).Get_listides().GetLista(), localctx.(*DeclaracionIniContext).Get_tiposvars().GetTipo(), localctx.(*DeclaracionIniContext).Get_expression().GetExpr())

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_tiposvars returns the _tiposvars rule contexts.
	Get_tiposvars() ITiposvarsContext

	// Get_listides returns the _listides rule contexts.
	Get_listides() IListidesContext

	// Set_tiposvars sets the _tiposvars rule contexts.
	Set_tiposvars(ITiposvarsContext)

	// Set_listides sets the _listides rule contexts.
	Set_listides(IListidesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	instr      interfaces.Instruccion
	_tiposvars ITiposvarsContext
	_listides  IListidesContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_tiposvars() ITiposvarsContext { return s._tiposvars }

func (s *DeclaracionContext) Get_listides() IListidesContext { return s._listides }

func (s *DeclaracionContext) Set_tiposvars(v ITiposvarsContext) { s._tiposvars = v }

func (s *DeclaracionContext) Set_listides(v IListidesContext) { s._listides = v }

func (s *DeclaracionContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *DeclaracionContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *DeclaracionContext) Tiposvars() ITiposvarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITiposvarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITiposvarsContext)
}

func (s *DeclaracionContext) Listides() IListidesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListidesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListidesContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *Calc) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CalcRULE_declaracion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)

		var _x = p.Tiposvars()

		localctx.(*DeclaracionContext)._tiposvars = _x
	}
	{
		p.SetState(393)

		var _x = p.listides(0)

		localctx.(*DeclaracionContext)._listides = _x
	}

	localctx.(*DeclaracionContext).instr = definicion.NewDeclaracion(localctx.(*DeclaracionContext).Get_listides().GetLista(), localctx.(*DeclaracionContext).Get_tiposvars().GetTipo())

	return localctx
}

// IRetornoContext is an interface to support dynamic dispatch.
type IRetornoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsRetornoContext differentiates from other interfaces.
	IsRetornoContext()
}

type RetornoContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruccion
	_expression IExpressionContext
}

func NewEmptyRetornoContext() *RetornoContext {
	var p = new(RetornoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_retorno
	return p
}

func (*RetornoContext) IsRetornoContext() {}

func NewRetornoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetornoContext {
	var p = new(RetornoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_retorno

	return p
}

func (s *RetornoContext) GetParser() antlr.Parser { return s.parser }

func (s *RetornoContext) Get_expression() IExpressionContext { return s._expression }

func (s *RetornoContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *RetornoContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *RetornoContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *RetornoContext) RETURN_P() antlr.TerminalNode {
	return s.GetToken(CalcRETURN_P, 0)
}

func (s *RetornoContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetornoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterRetorno(s)
	}
}

func (s *RetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitRetorno(s)
	}
}

func (p *Calc) Retorno() (localctx IRetornoContext) {
	localctx = NewRetornoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CalcRULE_retorno)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)
			p.Match(CalcRETURN_P)
		}
		localctx.(*RetornoContext).instr = transferenciaFlujo.NewReturn(entorno.VOID, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(398)
			p.Match(CalcRETURN_P)
		}
		{
			p.SetState(399)

			var _x = p.Expression()

			localctx.(*RetornoContext)._expression = _x
		}
		localctx.(*RetornoContext).instr = transferenciaFlujo.NewReturn(entorno.NULL, localctx.(*RetornoContext).Get_expression().GetExpr())

	}

	return localctx
}

// IListidesContext is an interface to support dynamic dispatch.
type IListidesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetSub returns the sub rule contexts.
	GetSub() IListidesContext

	// SetSub sets the sub rule contexts.
	SetSub(IListidesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListidesContext differentiates from other interfaces.
	IsListidesContext()
}

type ListidesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  *arrayList.List
	sub    IListidesContext
	_ID    antlr.Token
}

func NewEmptyListidesContext() *ListidesContext {
	var p = new(ListidesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_listides
	return p
}

func (*ListidesContext) IsListidesContext() {}

func NewListidesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListidesContext {
	var p = new(ListidesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_listides

	return p
}

func (s *ListidesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListidesContext) Get_ID() antlr.Token { return s._ID }

func (s *ListidesContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListidesContext) GetSub() IListidesContext { return s.sub }

func (s *ListidesContext) SetSub(v IListidesContext) { s.sub = v }

func (s *ListidesContext) GetLista() *arrayList.List { return s.lista }

func (s *ListidesContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListidesContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *ListidesContext) COMA() antlr.TerminalNode {
	return s.GetToken(CalcCOMA, 0)
}

func (s *ListidesContext) Listides() IListidesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListidesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListidesContext)
}

func (s *ListidesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListidesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListidesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterListides(s)
	}
}

func (s *ListidesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitListides(s)
	}
}

func (p *Calc) Listides() (localctx IListidesContext) {
	return p.listides(0)
}

func (p *Calc) listides(_p int) (localctx IListidesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListidesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListidesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, CalcRULE_listides, _p)
	localctx.(*ListidesContext).lista = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)

		var _m = p.Match(CalcID)

		localctx.(*ListidesContext)._ID = _m
	}
	localctx.(*ListidesContext).lista.Add(expresion.NewIdentificador((func() string {
		if localctx.(*ListidesContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListidesContext).Get_ID().GetText()
		}
	}())))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListidesContext(p, _parentctx, _parentState)
			localctx.(*ListidesContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_listides)
			p.SetState(408)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(409)
				p.Match(CalcCOMA)
			}
			{
				p.SetState(410)

				var _m = p.Match(CalcID)

				localctx.(*ListidesContext)._ID = _m
			}

			localctx.(*ListidesContext).GetSub().GetLista().Add(expresion.NewIdentificador((func() string {
				if localctx.(*ListidesContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListidesContext).Get_ID().GetText()
				}
			}())))
			localctx.(*ListidesContext).lista = localctx.(*ListidesContext).GetSub().GetLista()

		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// ITiposvarsContext is an interface to support dynamic dispatch.
type ITiposvarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo returns the tipo attribute.
	GetTipo() entorno.TipoDato

	// SetTipo sets the tipo attribute.
	SetTipo(entorno.TipoDato)

	// IsTiposvarsContext differentiates from other interfaces.
	IsTiposvarsContext()
}

type TiposvarsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tipo   entorno.TipoDato
}

func NewEmptyTiposvarsContext() *TiposvarsContext {
	var p = new(TiposvarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_tiposvars
	return p
}

func (*TiposvarsContext) IsTiposvarsContext() {}

func NewTiposvarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TiposvarsContext {
	var p = new(TiposvarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_tiposvars

	return p
}

func (s *TiposvarsContext) GetParser() antlr.Parser { return s.parser }

func (s *TiposvarsContext) GetTipo() entorno.TipoDato { return s.tipo }

func (s *TiposvarsContext) SetTipo(v entorno.TipoDato) { s.tipo = v }

func (s *TiposvarsContext) INTTYPE() antlr.TerminalNode {
	return s.GetToken(CalcINTTYPE, 0)
}

func (s *TiposvarsContext) STRINGTYPE() antlr.TerminalNode {
	return s.GetToken(CalcSTRINGTYPE, 0)
}

func (s *TiposvarsContext) FLOATTYPE() antlr.TerminalNode {
	return s.GetToken(CalcFLOATTYPE, 0)
}

func (s *TiposvarsContext) BOOLTYPE() antlr.TerminalNode {
	return s.GetToken(CalcBOOLTYPE, 0)
}

func (s *TiposvarsContext) VOIDTYPE() antlr.TerminalNode {
	return s.GetToken(CalcVOIDTYPE, 0)
}

func (s *TiposvarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TiposvarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TiposvarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterTiposvars(s)
	}
}

func (s *TiposvarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitTiposvars(s)
	}
}

func (p *Calc) Tiposvars() (localctx ITiposvarsContext) {
	localctx = NewTiposvarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CalcRULE_tiposvars)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(427)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcINTTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(417)
			p.Match(CalcINTTYPE)
		}
		localctx.(*TiposvarsContext).tipo = entorno.INTEGER

	case CalcSTRINGTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(419)
			p.Match(CalcSTRINGTYPE)
		}
		localctx.(*TiposvarsContext).tipo = entorno.STRING

	case CalcFLOATTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(421)
			p.Match(CalcFLOATTYPE)
		}
		localctx.(*TiposvarsContext).tipo = entorno.FLOAT

	case CalcBOOLTYPE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(423)
			p.Match(CalcBOOLTYPE)
		}
		localctx.(*TiposvarsContext).tipo = entorno.BOOLEAN

	case CalcVOIDTYPE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(425)
			p.Match(CalcVOIDTYPE)
		}
		localctx.(*TiposvarsContext).tipo = entorno.VOID

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr_rel returns the _expr_rel rule contexts.
	Get_expr_rel() IExpr_relContext

	// Get_expr_arit returns the _expr_arit rule contexts.
	Get_expr_arit() IExpr_aritContext

	// Get_instancia returns the _instancia rule contexts.
	Get_instancia() IInstanciaContext

	// Get_arraydata returns the _arraydata rule contexts.
	Get_arraydata() IArraydataContext

	// Get_instanciaClase returns the _instanciaClase rule contexts.
	Get_instanciaClase() IInstanciaClaseContext

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// Set_instancia sets the _instancia rule contexts.
	Set_instancia(IInstanciaContext)

	// Set_arraydata sets the _arraydata rule contexts.
	Set_arraydata(IArraydataContext)

	// Set_instanciaClase sets the _instanciaClase rule contexts.
	Set_instanciaClase(IInstanciaClaseContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	expr            interfaces.Expresion
	_expr_rel       IExpr_relContext
	_expr_arit      IExpr_aritContext
	_instancia      IInstanciaContext
	_arraydata      IArraydataContext
	_instanciaClase IInstanciaClaseContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Get_expr_rel() IExpr_relContext { return s._expr_rel }

func (s *ExpressionContext) Get_expr_arit() IExpr_aritContext { return s._expr_arit }

func (s *ExpressionContext) Get_instancia() IInstanciaContext { return s._instancia }

func (s *ExpressionContext) Get_arraydata() IArraydataContext { return s._arraydata }

func (s *ExpressionContext) Get_instanciaClase() IInstanciaClaseContext { return s._instanciaClase }

func (s *ExpressionContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *ExpressionContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *ExpressionContext) Set_instancia(v IInstanciaContext) { s._instancia = v }

func (s *ExpressionContext) Set_arraydata(v IArraydataContext) { s._arraydata = v }

func (s *ExpressionContext) Set_instanciaClase(v IInstanciaClaseContext) { s._instanciaClase = v }

func (s *ExpressionContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *ExpressionContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *ExpressionContext) Expr_rel() IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *ExpressionContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *ExpressionContext) Instancia() IInstanciaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstanciaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstanciaContext)
}

func (s *ExpressionContext) Arraydata() IArraydataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArraydataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArraydataContext)
}

func (s *ExpressionContext) InstanciaClase() IInstanciaClaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstanciaClaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstanciaClaseContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Calc) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CalcRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(429)

			var _x = p.expr_rel(0)

			localctx.(*ExpressionContext)._expr_rel = _x
		}
		localctx.(*ExpressionContext).expr = localctx.(*ExpressionContext).Get_expr_rel().GetExpr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(432)

			var _x = p.expr_arit(0)

			localctx.(*ExpressionContext)._expr_arit = _x
		}
		localctx.(*ExpressionContext).expr = localctx.(*ExpressionContext).Get_expr_arit().GetExpr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(435)

			var _x = p.Instancia()

			localctx.(*ExpressionContext)._instancia = _x
		}
		localctx.(*ExpressionContext).expr = localctx.(*ExpressionContext).Get_instancia().GetExpr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(438)

			var _x = p.Arraydata()

			localctx.(*ExpressionContext)._arraydata = _x
		}
		localctx.(*ExpressionContext).expr = localctx.(*ExpressionContext).Get_arraydata().GetExpr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(441)

			var _x = p.InstanciaClase()

			localctx.(*ExpressionContext)._instanciaClase = _x
		}
		localctx.(*ExpressionContext).expr = localctx.(*ExpressionContext).Get_instanciaClase().GetExpr()

	}

	return localctx
}

// IArraydataContext is an interface to support dynamic dispatch.
type IArraydataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listaExpresiones returns the _listaExpresiones rule contexts.
	Get_listaExpresiones() IListaExpresionesContext

	// Set_listaExpresiones sets the _listaExpresiones rule contexts.
	Set_listaExpresiones(IListaExpresionesContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsArraydataContext differentiates from other interfaces.
	IsArraydataContext()
}

type ArraydataContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	expr              interfaces.Expresion
	_listaExpresiones IListaExpresionesContext
}

func NewEmptyArraydataContext() *ArraydataContext {
	var p = new(ArraydataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_arraydata
	return p
}

func (*ArraydataContext) IsArraydataContext() {}

func NewArraydataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraydataContext {
	var p = new(ArraydataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_arraydata

	return p
}

func (s *ArraydataContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraydataContext) Get_listaExpresiones() IListaExpresionesContext {
	return s._listaExpresiones
}

func (s *ArraydataContext) Set_listaExpresiones(v IListaExpresionesContext) { s._listaExpresiones = v }

func (s *ArraydataContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *ArraydataContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *ArraydataContext) L_LLAVE() antlr.TerminalNode {
	return s.GetToken(CalcL_LLAVE, 0)
}

func (s *ArraydataContext) ListaExpresiones() IListaExpresionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaExpresionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaExpresionesContext)
}

func (s *ArraydataContext) R_LLAVE() antlr.TerminalNode {
	return s.GetToken(CalcR_LLAVE, 0)
}

func (s *ArraydataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraydataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArraydataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterArraydata(s)
	}
}

func (s *ArraydataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitArraydata(s)
	}
}

func (p *Calc) Arraydata() (localctx IArraydataContext) {
	localctx = NewArraydataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CalcRULE_arraydata)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(CalcL_LLAVE)
	}
	{
		p.SetState(447)

		var _x = p.listaExpresiones(0)

		localctx.(*ArraydataContext)._listaExpresiones = _x
	}
	{
		p.SetState(448)
		p.Match(CalcR_LLAVE)
	}
	localctx.(*ArraydataContext).expr = expresion2.NewValorArreglo(localctx.(*ArraydataContext).Get_listaExpresiones().GetLista())

	return localctx
}

// IInstanciaContext is an interface to support dynamic dispatch.
type IInstanciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_tiposvars returns the _tiposvars rule contexts.
	Get_tiposvars() ITiposvarsContext

	// Get_listanchos returns the _listanchos rule contexts.
	Get_listanchos() IListanchosContext

	// Set_tiposvars sets the _tiposvars rule contexts.
	Set_tiposvars(ITiposvarsContext)

	// Set_listanchos sets the _listanchos rule contexts.
	Set_listanchos(IListanchosContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsInstanciaContext differentiates from other interfaces.
	IsInstanciaContext()
}

type InstanciaContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	expr        interfaces.Expresion
	_tiposvars  ITiposvarsContext
	_listanchos IListanchosContext
}

func NewEmptyInstanciaContext() *InstanciaContext {
	var p = new(InstanciaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_instancia
	return p
}

func (*InstanciaContext) IsInstanciaContext() {}

func NewInstanciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanciaContext {
	var p = new(InstanciaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_instancia

	return p
}

func (s *InstanciaContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanciaContext) Get_tiposvars() ITiposvarsContext { return s._tiposvars }

func (s *InstanciaContext) Get_listanchos() IListanchosContext { return s._listanchos }

func (s *InstanciaContext) Set_tiposvars(v ITiposvarsContext) { s._tiposvars = v }

func (s *InstanciaContext) Set_listanchos(v IListanchosContext) { s._listanchos = v }

func (s *InstanciaContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *InstanciaContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *InstanciaContext) NEW_() antlr.TerminalNode {
	return s.GetToken(CalcNEW_, 0)
}

func (s *InstanciaContext) Tiposvars() ITiposvarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITiposvarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITiposvarsContext)
}

func (s *InstanciaContext) Listanchos() IListanchosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListanchosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListanchosContext)
}

func (s *InstanciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterInstancia(s)
	}
}

func (s *InstanciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitInstancia(s)
	}
}

func (p *Calc) Instancia() (localctx IInstanciaContext) {
	localctx = NewInstanciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CalcRULE_instancia)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(CalcNEW_)
	}
	{
		p.SetState(452)

		var _x = p.Tiposvars()

		localctx.(*InstanciaContext)._tiposvars = _x
	}
	{
		p.SetState(453)

		var _x = p.listanchos(0)

		localctx.(*InstanciaContext)._listanchos = _x
	}
	localctx.(*InstanciaContext).expr = expresion2.NewInstanciaArreglo(localctx.(*InstanciaContext).Get_tiposvars().GetTipo(), localctx.(*InstanciaContext).Get_listanchos().GetLista())

	return localctx
}

// IInstanciaClaseContext is an interface to support dynamic dispatch.
type IInstanciaClaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsInstanciaClaseContext differentiates from other interfaces.
	IsInstanciaClaseContext()
}

type InstanciaClaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   interfaces.Expresion
	_ID    antlr.Token
}

func NewEmptyInstanciaClaseContext() *InstanciaClaseContext {
	var p = new(InstanciaClaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_instanciaClase
	return p
}

func (*InstanciaClaseContext) IsInstanciaClaseContext() {}

func NewInstanciaClaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanciaClaseContext {
	var p = new(InstanciaClaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_instanciaClase

	return p
}

func (s *InstanciaClaseContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanciaClaseContext) Get_ID() antlr.Token { return s._ID }

func (s *InstanciaClaseContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *InstanciaClaseContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *InstanciaClaseContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *InstanciaClaseContext) NEW_() antlr.TerminalNode {
	return s.GetToken(CalcNEW_, 0)
}

func (s *InstanciaClaseContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *InstanciaClaseContext) LP() antlr.TerminalNode {
	return s.GetToken(CalcLP, 0)
}

func (s *InstanciaClaseContext) RP() antlr.TerminalNode {
	return s.GetToken(CalcRP, 0)
}

func (s *InstanciaClaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanciaClaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanciaClaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterInstanciaClase(s)
	}
}

func (s *InstanciaClaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitInstanciaClase(s)
	}
}

func (p *Calc) InstanciaClase() (localctx IInstanciaClaseContext) {
	localctx = NewInstanciaClaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CalcRULE_instanciaClase)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.Match(CalcNEW_)
	}
	{
		p.SetState(457)

		var _m = p.Match(CalcID)

		localctx.(*InstanciaClaseContext)._ID = _m
	}
	{
		p.SetState(458)
		p.Match(CalcLP)
	}
	{
		p.SetState(459)
		p.Match(CalcRP)
	}
	localctx.(*InstanciaClaseContext).expr = expresion2.NewInstanciaObjeto((func() string {
		if localctx.(*InstanciaClaseContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*InstanciaClaseContext).Get_ID().GetText()
		}
	}()))

	return localctx
}

// IAccesoarrContext is an interface to support dynamic dispatch.
type IAccesoarrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listanchos returns the _listanchos rule contexts.
	Get_listanchos() IListanchosContext

	// Set_listanchos sets the _listanchos rule contexts.
	Set_listanchos(IListanchosContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsAccesoarrContext differentiates from other interfaces.
	IsAccesoarrContext()
}

type AccesoarrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	expr        interfaces.Expresion
	_ID         antlr.Token
	_listanchos IListanchosContext
}

func NewEmptyAccesoarrContext() *AccesoarrContext {
	var p = new(AccesoarrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_accesoarr
	return p
}

func (*AccesoarrContext) IsAccesoarrContext() {}

func NewAccesoarrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesoarrContext {
	var p = new(AccesoarrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_accesoarr

	return p
}

func (s *AccesoarrContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesoarrContext) Get_ID() antlr.Token { return s._ID }

func (s *AccesoarrContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AccesoarrContext) Get_listanchos() IListanchosContext { return s._listanchos }

func (s *AccesoarrContext) Set_listanchos(v IListanchosContext) { s._listanchos = v }

func (s *AccesoarrContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *AccesoarrContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *AccesoarrContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *AccesoarrContext) Listanchos() IListanchosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListanchosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListanchosContext)
}

func (s *AccesoarrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoarrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesoarrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAccesoarr(s)
	}
}

func (s *AccesoarrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAccesoarr(s)
	}
}

func (p *Calc) Accesoarr() (localctx IAccesoarrContext) {
	localctx = NewAccesoarrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CalcRULE_accesoarr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)

		var _m = p.Match(CalcID)

		localctx.(*AccesoarrContext)._ID = _m
	}
	{
		p.SetState(463)

		var _x = p.listanchos(0)

		localctx.(*AccesoarrContext)._listanchos = _x
	}
	localctx.(*AccesoarrContext).expr = Accesos.NewAccessoArr((func() string {
		if localctx.(*AccesoarrContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AccesoarrContext).Get_ID().GetText()
		}
	}()), localctx.(*AccesoarrContext).Get_listanchos().GetLista())

	return localctx
}

// IAccesoObjetoContext is an interface to support dynamic dispatch.
type IAccesoObjetoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listaAccesos returns the _listaAccesos rule contexts.
	Get_listaAccesos() IListaAccesosContext

	// Set_listaAccesos sets the _listaAccesos rule contexts.
	Set_listaAccesos(IListaAccesosContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsAccesoObjetoContext differentiates from other interfaces.
	IsAccesoObjetoContext()
}

type AccesoObjetoContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	expr          interfaces.Expresion
	_listaAccesos IListaAccesosContext
}

func NewEmptyAccesoObjetoContext() *AccesoObjetoContext {
	var p = new(AccesoObjetoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_accesoObjeto
	return p
}

func (*AccesoObjetoContext) IsAccesoObjetoContext() {}

func NewAccesoObjetoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesoObjetoContext {
	var p = new(AccesoObjetoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_accesoObjeto

	return p
}

func (s *AccesoObjetoContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesoObjetoContext) Get_listaAccesos() IListaAccesosContext { return s._listaAccesos }

func (s *AccesoObjetoContext) Set_listaAccesos(v IListaAccesosContext) { s._listaAccesos = v }

func (s *AccesoObjetoContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *AccesoObjetoContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *AccesoObjetoContext) ListaAccesos() IListaAccesosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaAccesosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaAccesosContext)
}

func (s *AccesoObjetoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoObjetoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesoObjetoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAccesoObjeto(s)
	}
}

func (s *AccesoObjetoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAccesoObjeto(s)
	}
}

func (p *Calc) AccesoObjeto() (localctx IAccesoObjetoContext) {
	localctx = NewAccesoObjetoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CalcRULE_accesoObjeto)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)

		var _x = p.listaAccesos(0)

		localctx.(*AccesoObjetoContext)._listaAccesos = _x
	}
	localctx.(*AccesoObjetoContext).expr = Accesos.NewAccesoObjeto(localctx.(*AccesoObjetoContext).Get_listaAccesos().GetLista())

	return localctx
}

// IListaAccesosContext is an interface to support dynamic dispatch.
type IListaAccesosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSUBLISTA returns the SUBLISTA rule contexts.
	GetSUBLISTA() IListaAccesosContext

	// Get_acceso returns the _acceso rule contexts.
	Get_acceso() IAccesoContext

	// SetSUBLISTA sets the SUBLISTA rule contexts.
	SetSUBLISTA(IListaAccesosContext)

	// Set_acceso sets the _acceso rule contexts.
	Set_acceso(IAccesoContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListaAccesosContext differentiates from other interfaces.
	IsListaAccesosContext()
}

type ListaAccesosContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	lista    *arrayList.List
	SUBLISTA IListaAccesosContext
	_acceso  IAccesoContext
}

func NewEmptyListaAccesosContext() *ListaAccesosContext {
	var p = new(ListaAccesosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_listaAccesos
	return p
}

func (*ListaAccesosContext) IsListaAccesosContext() {}

func NewListaAccesosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaAccesosContext {
	var p = new(ListaAccesosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_listaAccesos

	return p
}

func (s *ListaAccesosContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaAccesosContext) GetSUBLISTA() IListaAccesosContext { return s.SUBLISTA }

func (s *ListaAccesosContext) Get_acceso() IAccesoContext { return s._acceso }

func (s *ListaAccesosContext) SetSUBLISTA(v IListaAccesosContext) { s.SUBLISTA = v }

func (s *ListaAccesosContext) Set_acceso(v IAccesoContext) { s._acceso = v }

func (s *ListaAccesosContext) GetLista() *arrayList.List { return s.lista }

func (s *ListaAccesosContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListaAccesosContext) Acceso() IAccesoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccesoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccesoContext)
}

func (s *ListaAccesosContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(CalcPUNTO, 0)
}

func (s *ListaAccesosContext) ListaAccesos() IListaAccesosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaAccesosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaAccesosContext)
}

func (s *ListaAccesosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaAccesosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaAccesosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterListaAccesos(s)
	}
}

func (s *ListaAccesosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitListaAccesos(s)
	}
}

func (p *Calc) ListaAccesos() (localctx IListaAccesosContext) {
	return p.listaAccesos(0)
}

func (p *Calc) listaAccesos(_p int) (localctx IListaAccesosContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaAccesosContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaAccesosContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, CalcRULE_listaAccesos, _p)

	localctx.(*ListaAccesosContext).lista = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)

		var _x = p.Acceso()

		localctx.(*ListaAccesosContext)._acceso = _x
	}
	localctx.(*ListaAccesosContext).lista.Add(localctx.(*ListaAccesosContext).Get_acceso().GetExpr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaAccesosContext(p, _parentctx, _parentState)
			localctx.(*ListaAccesosContext).SUBLISTA = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_listaAccesos)
			p.SetState(473)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(474)
				p.Match(CalcPUNTO)
			}
			{
				p.SetState(475)

				var _x = p.Acceso()

				localctx.(*ListaAccesosContext)._acceso = _x
			}

			localctx.(*ListaAccesosContext).GetSUBLISTA().GetLista().Add(localctx.(*ListaAccesosContext).Get_acceso().GetExpr())
			localctx.(*ListaAccesosContext).lista = localctx.(*ListaAccesosContext).GetSUBLISTA().GetLista()

		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IAccesoContext is an interface to support dynamic dispatch.
type IAccesoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_accesoarr returns the _accesoarr rule contexts.
	Get_accesoarr() IAccesoarrContext

	// Set_accesoarr sets the _accesoarr rule contexts.
	Set_accesoarr(IAccesoarrContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsAccesoContext differentiates from other interfaces.
	IsAccesoContext()
}

type AccesoContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	expr       interfaces.Expresion
	_ID        antlr.Token
	_accesoarr IAccesoarrContext
}

func NewEmptyAccesoContext() *AccesoContext {
	var p = new(AccesoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_acceso
	return p
}

func (*AccesoContext) IsAccesoContext() {}

func NewAccesoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesoContext {
	var p = new(AccesoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_acceso

	return p
}

func (s *AccesoContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesoContext) Get_ID() antlr.Token { return s._ID }

func (s *AccesoContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AccesoContext) Get_accesoarr() IAccesoarrContext { return s._accesoarr }

func (s *AccesoContext) Set_accesoarr(v IAccesoarrContext) { s._accesoarr = v }

func (s *AccesoContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *AccesoContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *AccesoContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *AccesoContext) Accesoarr() IAccesoarrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccesoarrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccesoarrContext)
}

func (s *AccesoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAcceso(s)
	}
}

func (s *AccesoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAcceso(s)
	}
}

func (p *Calc) Acceso() (localctx IAccesoContext) {
	localctx = NewAccesoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CalcRULE_acceso)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)

			var _m = p.Match(CalcID)

			localctx.(*AccesoContext)._ID = _m
		}
		localctx.(*AccesoContext).expr = expresion.NewIdentificador((func() string {
			if localctx.(*AccesoContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AccesoContext).Get_ID().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(485)

			var _x = p.Accesoarr()

			localctx.(*AccesoContext)._accesoarr = _x
		}
		localctx.(*AccesoContext).expr = localctx.(*AccesoContext).Get_accesoarr().GetExpr()

	}

	return localctx
}

// IExpr_relContext is an interface to support dynamic dispatch.
type IExpr_relContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpr_relContext

	// Get_expr_arit returns the _expr_arit rule contexts.
	Get_expr_arit() IExpr_aritContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpr_relContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpr_relContext)

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpr_relContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsExpr_relContext differentiates from other interfaces.
	IsExpr_relContext()
}

type Expr_relContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	expr       interfaces.Expresion
	opIz       IExpr_relContext
	_expr_arit IExpr_aritContext
	op         antlr.Token
	opDe       IExpr_relContext
}

func NewEmptyExpr_relContext() *Expr_relContext {
	var p = new(Expr_relContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_expr_rel
	return p
}

func (*Expr_relContext) IsExpr_relContext() {}

func NewExpr_relContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_relContext {
	var p = new(Expr_relContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_expr_rel

	return p
}

func (s *Expr_relContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_relContext) GetOp() antlr.Token { return s.op }

func (s *Expr_relContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_relContext) GetOpIz() IExpr_relContext { return s.opIz }

func (s *Expr_relContext) Get_expr_arit() IExpr_aritContext { return s._expr_arit }

func (s *Expr_relContext) GetOpDe() IExpr_relContext { return s.opDe }

func (s *Expr_relContext) SetOpIz(v IExpr_relContext) { s.opIz = v }

func (s *Expr_relContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *Expr_relContext) SetOpDe(v IExpr_relContext) { s.opDe = v }

func (s *Expr_relContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *Expr_relContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *Expr_relContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *Expr_relContext) AllExpr_rel() []IExpr_relContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_relContext)(nil)).Elem())
	var tst = make([]IExpr_relContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_relContext)
		}
	}

	return tst
}

func (s *Expr_relContext) Expr_rel(i int) IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *Expr_relContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(CalcMAYORIGUAL, 0)
}

func (s *Expr_relContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(CalcMENORIGUAL, 0)
}

func (s *Expr_relContext) MENOR() antlr.TerminalNode {
	return s.GetToken(CalcMENOR, 0)
}

func (s *Expr_relContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(CalcMAYOR, 0)
}

func (s *Expr_relContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_relContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_relContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExpr_rel(s)
	}
}

func (s *Expr_relContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExpr_rel(s)
	}
}

func (p *Calc) Expr_rel() (localctx IExpr_relContext) {
	return p.expr_rel(0)
}

func (p *Calc) expr_rel(_p int) (localctx IExpr_relContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_relContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_relContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, CalcRULE_expr_rel, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)

		var _x = p.expr_arit(0)

		localctx.(*Expr_relContext)._expr_arit = _x
	}
	localctx.(*Expr_relContext).expr = localctx.(*Expr_relContext).Get_expr_arit().GetExpr()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_relContext(p, _parentctx, _parentState)
			localctx.(*Expr_relContext).opIz = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CalcRULE_expr_rel)
			p.SetState(494)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(495)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expr_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CalcMAYORIGUAL-34))|(1<<(CalcMENORIGUAL-34))|(1<<(CalcMAYOR-34))|(1<<(CalcMENOR-34)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expr_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(496)

				var _x = p.expr_rel(3)

				localctx.(*Expr_relContext).opDe = _x
			}
			localctx.(*Expr_relContext).expr = expresion.NewOperacion(localctx.(*Expr_relContext).GetOpIz().GetExpr(), (func() string {
				if localctx.(*Expr_relContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expr_relContext).GetOp().GetText()
				}
			}()), localctx.(*Expr_relContext).GetOpDe().GetExpr(), false)

		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr_aritContext is an interface to support dynamic dispatch.
type IExpr_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpr_aritContext

	// GetOpU returns the opU rule contexts.
	GetOpU() IExpressionContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpr_aritContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpr_aritContext)

	// SetOpU sets the opU rule contexts.
	SetOpU(IExpressionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpr_aritContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsExpr_aritContext differentiates from other interfaces.
	IsExpr_aritContext()
}

type Expr_aritContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	expr        interfaces.Expresion
	opIz        IExpr_aritContext
	opU         IExpressionContext
	_expression IExpressionContext
	_expr_valor IExpr_valorContext
	op          antlr.Token
	opDe        IExpr_aritContext
}

func NewEmptyExpr_aritContext() *Expr_aritContext {
	var p = new(Expr_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_expr_arit
	return p
}

func (*Expr_aritContext) IsExpr_aritContext() {}

func NewExpr_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_aritContext {
	var p = new(Expr_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_expr_arit

	return p
}

func (s *Expr_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expr_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_aritContext) GetOpIz() IExpr_aritContext { return s.opIz }

func (s *Expr_aritContext) GetOpU() IExpressionContext { return s.opU }

func (s *Expr_aritContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_aritContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *Expr_aritContext) GetOpDe() IExpr_aritContext { return s.opDe }

func (s *Expr_aritContext) SetOpIz(v IExpr_aritContext) { s.opIz = v }

func (s *Expr_aritContext) SetOpU(v IExpressionContext) { s.opU = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *Expr_aritContext) SetOpDe(v IExpr_aritContext) { s.opDe = v }

func (s *Expr_aritContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *Expr_aritContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *Expr_aritContext) SUB() antlr.TerminalNode {
	return s.GetToken(CalcSUB, 0)
}

func (s *Expr_aritContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expr_aritContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *Expr_aritContext) LP() antlr.TerminalNode {
	return s.GetToken(CalcLP, 0)
}

func (s *Expr_aritContext) RP() antlr.TerminalNode {
	return s.GetToken(CalcRP, 0)
}

func (s *Expr_aritContext) AllExpr_arit() []IExpr_aritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem())
	var tst = make([]IExpr_aritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_aritContext)
		}
	}

	return tst
}

func (s *Expr_aritContext) Expr_arit(i int) IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *Expr_aritContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalcMUL, 0)
}

func (s *Expr_aritContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalcDIV, 0)
}

func (s *Expr_aritContext) ADD() antlr.TerminalNode {
	return s.GetToken(CalcADD, 0)
}

func (s *Expr_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExpr_arit(s)
	}
}

func (s *Expr_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExpr_arit(s)
	}
}

func (p *Calc) Expr_arit() (localctx IExpr_aritContext) {
	return p.expr_arit(0)
}

func (p *Calc) expr_arit(_p int) (localctx IExpr_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 80
	p.EnterRecursionRule(localctx, 80, CalcRULE_expr_arit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(517)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcSUB:
		{
			p.SetState(505)
			p.Match(CalcSUB)
		}
		{
			p.SetState(506)

			var _x = p.Expression()

			localctx.(*Expr_aritContext).opU = _x
			localctx.(*Expr_aritContext)._expression = _x
		}
		localctx.(*Expr_aritContext).expr = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpU().GetExpr(), "-", nil, true)

	case CalcNUMBER, CalcFLOAT, CalcSTRING, CalcTRUE, CalcFALSE, CalcID:
		{
			p.SetState(509)

			var _x = p.Expr_valor()

			localctx.(*Expr_aritContext)._expr_valor = _x
		}
		localctx.(*Expr_aritContext).expr = localctx.(*Expr_aritContext).Get_expr_valor().GetExpr()

	case CalcLP:
		{
			p.SetState(512)
			p.Match(CalcLP)
		}
		{
			p.SetState(513)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(514)
			p.Match(CalcRP)
		}
		localctx.(*Expr_aritContext).expr = localctx.(*Expr_aritContext).Get_expression().GetExpr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(529)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalcRULE_expr_arit)
				p.SetState(519)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(520)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcMUL || _la == CalcDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(521)

					var _x = p.expr_arit(5)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).expr = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetExpr(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetExpr(), false)

			case 2:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalcRULE_expr_arit)
				p.SetState(524)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(525)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcADD || _la == CalcSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(526)

					var _x = p.expr_arit(4)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).expr = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetExpr(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetExpr(), false)

			}

		}
		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr_valorContext is an interface to support dynamic dispatch.
type IExpr_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_llamada returns the _llamada rule contexts.
	Get_llamada() ILlamadaContext

	// Get_accesoarr returns the _accesoarr rule contexts.
	Get_accesoarr() IAccesoarrContext

	// Get_accesoObjeto returns the _accesoObjeto rule contexts.
	Get_accesoObjeto() IAccesoObjetoContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_llamada sets the _llamada rule contexts.
	Set_llamada(ILlamadaContext)

	// Set_accesoarr sets the _accesoarr rule contexts.
	Set_accesoarr(IAccesoarrContext)

	// Set_accesoObjeto sets the _accesoObjeto rule contexts.
	Set_accesoObjeto(IAccesoObjetoContext)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsExpr_valorContext differentiates from other interfaces.
	IsExpr_valorContext()
}

type Expr_valorContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	expr          interfaces.Expresion
	_primitivo    IPrimitivoContext
	_llamada      ILlamadaContext
	_accesoarr    IAccesoarrContext
	_accesoObjeto IAccesoObjetoContext
}

func NewEmptyExpr_valorContext() *Expr_valorContext {
	var p = new(Expr_valorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_expr_valor
	return p
}

func (*Expr_valorContext) IsExpr_valorContext() {}

func NewExpr_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_valorContext {
	var p = new(Expr_valorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_expr_valor

	return p
}

func (s *Expr_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_valorContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expr_valorContext) Get_llamada() ILlamadaContext { return s._llamada }

func (s *Expr_valorContext) Get_accesoarr() IAccesoarrContext { return s._accesoarr }

func (s *Expr_valorContext) Get_accesoObjeto() IAccesoObjetoContext { return s._accesoObjeto }

func (s *Expr_valorContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expr_valorContext) Set_llamada(v ILlamadaContext) { s._llamada = v }

func (s *Expr_valorContext) Set_accesoarr(v IAccesoarrContext) { s._accesoarr = v }

func (s *Expr_valorContext) Set_accesoObjeto(v IAccesoObjetoContext) { s._accesoObjeto = v }

func (s *Expr_valorContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *Expr_valorContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *Expr_valorContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Expr_valorContext) Llamada() ILlamadaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILlamadaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILlamadaContext)
}

func (s *Expr_valorContext) Accesoarr() IAccesoarrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccesoarrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccesoarrContext)
}

func (s *Expr_valorContext) AccesoObjeto() IAccesoObjetoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccesoObjetoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccesoObjetoContext)
}

func (s *Expr_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_valorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExpr_valor(s)
	}
}

func (s *Expr_valorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExpr_valor(s)
	}
}

func (p *Calc) Expr_valor() (localctx IExpr_valorContext) {
	localctx = NewExpr_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CalcRULE_expr_valor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(534)

			var _x = p.Primitivo()

			localctx.(*Expr_valorContext)._primitivo = _x
		}
		localctx.(*Expr_valorContext).expr = localctx.(*Expr_valorContext).Get_primitivo().GetExpr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(537)

			var _x = p.Llamada()

			localctx.(*Expr_valorContext)._llamada = _x
		}
		localctx.(*Expr_valorContext).expr = localctx.(*Expr_valorContext).Get_llamada().GetExpr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(540)

			var _x = p.Accesoarr()

			localctx.(*Expr_valorContext)._accesoarr = _x
		}
		localctx.(*Expr_valorContext).expr = localctx.(*Expr_valorContext).Get_accesoarr().GetExpr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(543)

			var _x = p.AccesoObjeto()

			localctx.(*Expr_valorContext)._accesoObjeto = _x
		}
		localctx.(*Expr_valorContext).expr = localctx.(*Expr_valorContext).Get_accesoObjeto().GetExpr()

	}

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetExpr returns the expr attribute.
	GetExpr() interfaces.Expresion

	// SetExpr sets the expr attribute.
	SetExpr(interfaces.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	expr    interfaces.Expresion
	_NUMBER antlr.Token
	_FLOAT  antlr.Token
	_STRING antlr.Token
	_ID     antlr.Token
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitivoContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *PrimitivoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivoContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitivoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivoContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *PrimitivoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivoContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitivoContext) GetExpr() interfaces.Expresion { return s.expr }

func (s *PrimitivoContext) SetExpr(v interfaces.Expresion) { s.expr = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcNUMBER, 0)
}

func (s *PrimitivoContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CalcFLOAT, 0)
}

func (s *PrimitivoContext) STRING() antlr.TerminalNode {
	return s.GetToken(CalcSTRING, 0)
}

func (s *PrimitivoContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcID, 0)
}

func (s *PrimitivoContext) TRUE() antlr.TerminalNode {
	return s.GetToken(CalcTRUE, 0)
}

func (s *PrimitivoContext) FALSE() antlr.TerminalNode {
	return s.GetToken(CalcFALSE, 0)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Calc) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, CalcRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(560)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(548)

			var _m = p.Match(CalcNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).expr = expresion.NewPrimitivo(num, entorno.INTEGER)

	case CalcFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(550)

			var _m = p.Match(CalcFLOAT)

			localctx.(*PrimitivoContext)._FLOAT = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_FLOAT() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_FLOAT().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).expr = expresion.NewPrimitivo(num, entorno.FLOAT)

	case CalcSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(552)

			var _m = p.Match(CalcSTRING)

			localctx.(*PrimitivoContext)._STRING = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).expr = expresion.NewPrimitivo(str, entorno.STRING)

	case CalcID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(554)

			var _m = p.Match(CalcID)

			localctx.(*PrimitivoContext)._ID = _m
		}
		localctx.(*PrimitivoContext).expr = expresion.NewIdentificador((func() string {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetText()
			}
		}()))

	case CalcTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(556)
			p.Match(CalcTRUE)
		}
		localctx.(*PrimitivoContext).expr = expresion.NewPrimitivo(true, entorno.BOOLEAN)

	case CalcFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(558)
			p.Match(CalcFALSE)
		}
		localctx.(*PrimitivoContext).expr = expresion.NewPrimitivo(false, entorno.BOOLEAN)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Calc) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ListaClasesContext = nil
		if localctx != nil {
			t = localctx.(*ListaClasesContext)
		}
		return p.ListaClases_Sempred(t, predIndex)

	case 4:
		var t *ContenidoClaseContext = nil
		if localctx != nil {
			t = localctx.(*ContenidoClaseContext)
		}
		return p.ContenidoClase_Sempred(t, predIndex)

	case 8:
		var t *ParametrosContext = nil
		if localctx != nil {
			t = localctx.(*ParametrosContext)
		}
		return p.Parametros_Sempred(t, predIndex)

	case 15:
		var t *DimensionesContext = nil
		if localctx != nil {
			t = localctx.(*DimensionesContext)
		}
		return p.Dimensiones_Sempred(t, predIndex)

	case 17:
		var t *ListanchosContext = nil
		if localctx != nil {
			t = localctx.(*ListanchosContext)
		}
		return p.Listanchos_Sempred(t, predIndex)

	case 25:
		var t *ListaExpresionesContext = nil
		if localctx != nil {
			t = localctx.(*ListaExpresionesContext)
		}
		return p.ListaExpresiones_Sempred(t, predIndex)

	case 29:
		var t *ListidesContext = nil
		if localctx != nil {
			t = localctx.(*ListidesContext)
		}
		return p.Listides_Sempred(t, predIndex)

	case 37:
		var t *ListaAccesosContext = nil
		if localctx != nil {
			t = localctx.(*ListaAccesosContext)
		}
		return p.ListaAccesos_Sempred(t, predIndex)

	case 39:
		var t *Expr_relContext = nil
		if localctx != nil {
			t = localctx.(*Expr_relContext)
		}
		return p.Expr_rel_Sempred(t, predIndex)

	case 40:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Calc) ListaClases_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) ContenidoClase_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) Parametros_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) Dimensiones_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) Listanchos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) ListaExpresiones_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) Listides_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) ListaAccesos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) Expr_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Calc) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
