package sintatico

import (
	core "github.com/maiconio/portugo/core"
	lex "github.com/maiconio/portugo/lex"
	util "github.com/maiconio/portugo/util"
	"testing"
)

func TestCarregaTokens(t *testing.T) {
	r1 :=
		`P[]
  INICIO[início]
  V[V]
    V1[V1]
      _[_]
  AC[AC]
    AC1[AC1]
      A[A]
        ESCREVA[escreva]
        ([(]
        ESCV1[ESCV1]
          STRING[1.a) Ele tem ]
        ESCV2[ESCV2]
          ,[,]
          ESCV1[ESCV1]
            L[L]
              L1[L1]
                L3[L3]
                  L5[L5]
                    R1[R1]
                      M1[M1]
                        M3[M3]
                          M5[M5]
                            M7[M7]
                              M9[M9]
                                INTEIRO[15]
                              M10[M10]
                                _[_]
                            M8[M8]
                              _[_]
                          M6[M6]
                            _[_]
                        M4[M4]
                          _[_]
                      M2[M2]
                        _[_]
                    R2[R2]
                      _[_]
                  L6[L6]
                    _[_]
                L4[L4]
                  _[_]
              L2[L2]
                _[_]
          ESCV2[ESCV2]
            ,[,]
            ESCV1[ESCV1]
              STRING[ irmãos.]
            ESCV2[ESCV2]
              _[_]
        )[)]
      ;[;]
      AC1[AC1]
        A[A]
          ESCREVA[escreva]
          ([(]
          ESCV1[ESCV1]
            STRING[1.b) A temperatura desta noite será de ]
          ESCV2[ESCV2]
            ,[,]
            ESCV1[ESCV1]
              L[L]
                L1[L1]
                  L3[L3]
                    L5[L5]
                      R1[R1]
                        M1[M1]
                          M3[M3]
                            M5[M5]
                              M7[M7]
                                M9[M9]
                                  +-[+]
                                  M1[M1]
                                    M3[M3]
                                      M5[M5]
                                        M7[M7]
                                          M9[M9]
                                            +-[-]
                                            M1[M1]
                                              M3[M3]
                                                M5[M5]
                                                  M7[M7]
                                                    M9[M9]
                                                      INTEIRO[2]
                                                    M10[M10]
                                                      _[_]
                                                  M8[M8]
                                                    _[_]
                                                M6[M6]
                                                  _[_]
                                              M4[M4]
                                                _[_]
                                          M10[M10]
                                            _[_]
                                        M8[M8]
                                          _[_]
                                      M6[M6]
                                        _[_]
                                    M4[M4]
                                      _[_]
                                M10[M10]
                                  _[_]
                              M8[M8]
                                _[_]
                            M6[M6]
                              _[_]
                          M4[M4]
                            _[_]
                        M2[M2]
                          _[_]
                      R2[R2]
                        _[_]
                    L6[L6]
                      _[_]
                  L4[L4]
                    _[_]
                L2[L2]
                  _[_]
            ESCV2[ESCV2]
              ,[,]
              ESCV1[ESCV1]
                STRING[ graus.]
              ESCV2[ESCV2]
                _[_]
          )[)]
        ;[;]
        AC1[AC1]
          A[A]
            ESCREVA[escreva]
            ([(]
            ESCV1[ESCV1]
              STRING[]
            ESCV2[ESCV2]
              _[_]
            )[)]
          ;[;]
          AC1[AC1]
            A[A]
              ESCREVA[escreva]
              ([(]
              ESCV1[ESCV1]
                STRING[2.a) Ela tem ]
              ESCV2[ESCV2]
                ,[,]
                ESCV1[ESCV1]
                  L[L]
                    L1[L1]
                      L3[L3]
                        L5[L5]
                          R1[R1]
                            M1[M1]
                              M3[M3]
                                M5[M5]
                                  M7[M7]
                                    M9[M9]
                                      REAL[1.73]
                                    M10[M10]
                                      _[_]
                                  M8[M8]
                                    _[_]
                                M6[M6]
                                  _[_]
                              M4[M4]
                                _[_]
                            M2[M2]
                              _[_]
                          R2[R2]
                            _[_]
                        L6[L6]
                          _[_]
                      L4[L4]
                        _[_]
                    L2[L2]
                      _[_]
                ESCV2[ESCV2]
                  ,[,]
                  ESCV1[ESCV1]
                    STRING[ de altura]
                  ESCV2[ESCV2]
                    _[_]
              )[)]
            ;[;]
            AC1[AC1]
              A[A]
                ESCREVA[escreva]
                ([(]
                ESCV1[ESCV1]
                  STRING[2.b) Meu saldo bancário é de -R$]
                ESCV2[ESCV2]
                  ,[,]
                  ESCV1[ESCV1]
                    L[L]
                      L1[L1]
                        L3[L3]
                          L5[L5]
                            R1[R1]
                              M1[M1]
                                M3[M3]
                                  M5[M5]
                                    M7[M7]
                                      M9[M9]
                                        REAL[121.07]
                                      M10[M10]
                                        _[_]
                                    M8[M8]
                                      _[_]
                                  M6[M6]
                                    _[_]
                                M4[M4]
                                  _[_]
                              M2[M2]
                                _[_]
                            R2[R2]
                              _[_]
                          L6[L6]
                            _[_]
                        L4[L4]
                          _[_]
                      L2[L2]
                        _[_]
                  ESCV2[ESCV2]
                    ,[,]
                    ESCV1[ESCV1]
                      STRING[ .]
                    ESCV2[ESCV2]
                      _[_]
                )[)]
              ;[;]
              AC1[AC1]
                A[A]
                  ESCREVA[escreva]
                  ([(]
                  ESCV1[ESCV1]
                    STRING[]
                  ESCV2[ESCV2]
                    _[_]
                  )[)]
                ;[;]
                AC1[AC1]
                  A[A]
                    ESCREVA[escreva]
                    ([(]
                    ESCV1[ESCV1]
                      STRING[3.a) Constava na prova: ]
                    ESCV2[ESCV2]
                      ,[,]
                      ESCV1[ESCV1]
                        STRING['Use somente caneta!']
                      ESCV2[ESCV2]
                        _[_]
                    )[)]
                  ;[;]
                  AC1[AC1]
                    A[A]
                      ESCREVA[escreva]
                      ([(]
                      ESCV1[ESCV1]
                        STRING[3.b) O parque municipal estava repleto de placas: ]
                      ESCV2[ESCV2]
                        ,[,]
                        ESCV1[ESCV1]
                          STRING['Não pise na grama.']
                        ESCV2[ESCV2]
                          _[_]
                      )[)]
                    ;[;]
                    AC1[AC1]
                      A[A]
                        ESCREVA[escreva]
                        ([(]
                        ESCV1[ESCV1]
                          STRING[]
                        ESCV2[ESCV2]
                          _[_]
                        )[)]
                      ;[;]
                      AC1[AC1]
                        A[A]
                          ESCREVA[escreva]
                          ([(]
                          ESCV1[ESCV1]
                            STRING[4.a) A porta pode estar aberta ou fechada. ]
                          ESCV2[ESCV2]
                            ,[,]
                            ESCV1[ESCV1]
                              L[L]
                                L1[L1]
                                  L3[L3]
                                    L5[L5]
                                      LOGICO[verdadeiro]
                                    L6[L6]
                                      _[_]
                                  L4[L4]
                                    _[_]
                                L2[L2]
                                  _[_]
                            ESCV2[ESCV2]
                              _[_]
                          )[)]
                        ;[;]
                        AC1[AC1]
                          A[A]
                            ESCREVA[escreva]
                            ([(]
                            ESCV1[ESCV1]
                              STRING[4.b) A lâmpada pode estar acesa ou apagada. ]
                            ESCV2[ESCV2]
                              ,[,]
                              ESCV1[ESCV1]
                                L[L]
                                  L1[L1]
                                    L3[L3]
                                      L5[L5]
                                        LOGICO[falso]
                                      L6[L6]
                                        _[_]
                                    L4[L4]
                                      _[_]
                                  L2[L2]
                                    _[_]
                              ESCV2[ESCV2]
                                _[_]
                            )[)]
                          ;[;]
                          AC1[AC1]
                            _[_]
  FIM[fim]
  PONTO[.]
`

	listaTokens := lex.CarregaTokens("../testes_src/cap02/2.1_tipos_primitivos.ptg")

	parseTree := core.Node{nil, nil, "P", 0, 0, core.Token{"", ""}}
	MontaParsingTree(&parseTree, listaTokens)
	if r1 != util.MostraTree(&parseTree) {
		t.Error(len(util.MostraTree(&parseTree)), "---------", len(r1))
	}
}
