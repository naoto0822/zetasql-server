#include <memory>
#include <iostream>
#include <string.h>
#include "zetasql/public/sql_formatter.h"
#include "zetasql/public/parse_helpers.h"
#include "zetasql/public/analyzer.h"
#include "zetasql/public/simple_catalog.h"
#include "zetasql/public/options.pb.h"
#include "zetasql/public/analyzer_options.h"
#include "zetasql/parser/parser.h"
#include "absl/status/status.h"
#include "absl/memory/memory.h"
#include "absl/strings/cord.h"
#include "absl/strings/str_cat.h"
#include "absl/strings/strip.h"
#include "zetasql/common/status_payload_utils.h"
#include "parse_query.h"

char *isValidStatement(char* sql) {
  absl::Status status = zetasql::IsValidStatementSyntax(sql, zetasql::ERROR_MESSAGE_WITH_PAYLOAD);
  std::string return_status = zetasql::internal::StatusToString(status);
  return strdup(return_status.c_str());
}

char *parseStatement(char *sql) {
  std::unique_ptr<zetasql::ParserOutput> parser_output;
  const absl::Status parse_status = zetasql::ParseStatement(sql, zetasql::ParserOptions(), &parser_output);
  if (!parse_status.ok()) {
    return strdup(parse_status.c_str());
  }

  std::string ast_str = parser_output->statement()->DebugString();
  return strdup(ast_str.c_str());
}

//char *analyzeStatement(char* sql) {
//  std::unique_ptr<const zetasql::AnalyzerOutput> analyzer_output;
//
//  zetasql::AnalyzerOptions options = zetasql::AnalyzerOptions();
//  options.set_enabled_rewrites({});
//
//  zetasql::TypeFactory type_factory_;
//  zetasql::SimpleCatalog catalog{"zetasql"};
//
//  absl::Status status = zetasql::AnalyzeStatement(sql, options, &catalog, &type_factory_, &analyzer_output);
//  std::string return_status = zetasql::internal::StatusToString(status);
//
//  std::string ret = analyzer_output->resolved_statement()->DebugString();
//
//  return strdup(ret.c_str());
//}
