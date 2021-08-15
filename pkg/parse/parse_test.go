package parse

import (
	"encoding/json"
	"fmt"
	"testing"
)

// TODO: table test

func TestParseDebugString1(t *testing.T) {
	input := `QueryStatement [0-19]
  Query [0-19]
    Select [0-19]
      SelectList [7-9]
        SelectColumn [7-9]
          PathExpression [7-9]
            Identifier(id) [7-9]
      FromClause [10-19]
        TablePathExpression [15-19]
          PathExpression [15-19]
            Identifier(user) [15-19]
`

	lexer := NewLexer(input)
	result := Parse(lexer)

	size := result.Len()
	for i := 0; i < size; i++ {
		if result.Len() == 1 {
			break
		}
		result.Pop()
	}

	top := result.Pop()
	bytes, err := json.MarshalIndent(top, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	Reset()
}

func TestParseDebugString2(t *testing.T) {
	input := `QueryStatement [0-521]
  Query [0-521]
    WithClause [0-467]
      WithClauseEntry [5-129]
        Identifier(A) [5-6]
        Query [14-127]
          Select [14-109]
            SelectList [25-26]
              SelectColumn [25-26]
                Star(*) [25-26]
            FromClause [29-56]
              TablePathExpression [34-56]
                PathExpression [34-56]
                  Identifier(hoge-11111.dataset.t) [34-56]
            WhereClause [59-109]
              OrExpr [65-109]
                AndExpr [65-89]
                  BinaryExpression(=) [65-71]
                    PathExpression [65-67]
                      Identifier(id) [65-67]
                    IntLiteral(1) [70-71]
                  BinaryExpression(=) [76-89]
                    PathExpression [76-80]
                      Identifier(name) [76-80]
                    StringLiteral("mike") [83-89]
                BinaryExpression(=) [93-109]
                  PathExpression [93-105]
                    Identifier(account_type) [93-105]
                  IntLiteral(3) [108-109]
          OrderBy [112-127]
            OrderingExpression(ASC) [121-127]
              PathExpression [121-123]
                Identifier(id) [121-123]
      WithClauseEntry [131-467]
        Identifier(B) [131-132]
        Query [139-465]
          SetOperation(UNION ALL) [139-465]
            Select [139-398]
              SelectList [150-216]
                SelectColumn [150-158]
                  PathExpression [150-152]
                    Identifier(id) [150-152]
                  Alias [153-158]
                    Identifier(id) [156-158]
                SelectColumn [164-187]
                  FunctionCall [164-173]
                    PathExpression [164-167]
                      Identifier(SUM) [164-167]
                    PathExpression [168-172]
                      Identifier(cost) [168-172]
                  Alias [174-187]
                    Identifier(total_cost) [177-187]
                SelectColumn [193-216]
                  FunctionCall [193-201]
                    PathExpression [193-198]
                      Identifier(COUNT) [193-198]
                    Star(*) [199-200]
                  Alias [202-216]
                    Identifier(total_count) [205-216]
              FromClause [219-384]
                TableSubquery [224-384]
                  Query [230-380]
                    Select [230-380]
                      SelectList [243-307]
                        SelectColumn [243-255]
                          PathExpression [243-247]
                            Identifier(cost) [243-247]
                          Alias [248-255]
                            Identifier(cost) [251-255]
                        SelectColumn [263-271]
                          PathExpression [263-265]
                            Identifier(id) [263-265]
                          Alias [266-271]
                            Identifier(id) [269-271]
                        SelectColumn [279-307]
                          FunctionCall [279-295]
                            PathExpression [279-284]
                              Identifier(SPLIT) [279-284]
                            PathExpression [285-289]
                              Identifier(hoge) [285-289]
                            StringLiteral(",") [291-294]
                          Alias [296-307]
                            Identifier(hoge_arr) [299-307]
                      FromClause [312-380]
                        Join(INNER) [331-380]
                          TablePathExpression [317-326]
                            PathExpression [317-324]
                              Identifier(payment) [317-324]
                            Alias [325-326]
                              Identifier(p) [325-326]
                          TablePathExpression [342-352]
                            PathExpression [342-350]
                              Identifier(customer) [342-350]
                            Alias [351-352]
                              Identifier(c) [351-352]
                          OnClause [357-380]
                            BinaryExpression(=) [360-380]
                              PathExpression [360-373]
                                Identifier(p) [360-361]
                                Identifier(customer_id) [362-373]
                              PathExpression [376-380]
                                Identifier(c) [376-377]
                                Identifier(id) [378-380]
              GroupBy [387-398]
                GroupingItem [396-398]
                  PathExpression [396-398]
                    Identifier(id) [396-398]
            Select [413-465]
              SelectList [424-465]
                SelectColumn [424-441]
                  IntLiteral(100) [424-427]
                  Alias [428-441]
                    Identifier(total_cost) [431-441]
                SelectColumn [447-465]
                  IntLiteral(300) [447-450]
                  Alias [451-465]
                    Identifier(total_count) [454-465]
    Select [468-521]
      SelectList [477-487]
        SelectColumn [477-480]
          DotStar [478-480]
            PathExpression [477-478]
              Identifier(A) [477-478]
        SelectColumn [484-487]
          DotStar [485-487]
            PathExpression [484-485]
              Identifier(B) [484-485]
      FromClause [488-521]
        Join(LEFT) [495-521]
          TablePathExpression [493-494]
            PathExpression [493-494]
              Identifier(A) [493-494]
          TablePathExpression [505-506]
            PathExpression [505-506]
              Identifier(B) [505-506]
          OnClause [507-521]
            BinaryExpression(=) [510-521]
              PathExpression [510-514]
                Identifier(A) [510-511]
                Identifier(id) [512-514]
              PathExpression [517-521]
                Identifier(B) [517-518]
                Identifier(id) [519-521]
`

	lexer := NewLexer(input)
	result := Parse(lexer)

	size := result.Len()
	for i := 0; i < size; i++ {
		if result.Len() == 1 {
			break
		}
		result.Pop()
	}

	top := result.Pop()
	bytes, err := json.MarshalIndent(top, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	Reset()
}
