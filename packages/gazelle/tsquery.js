const {readFileSync} = require('fs');
const {tsquery} = require('@phenomnomnominal/tsquery');
const builtins = require('builtins')();

const tsFilePath = process.argv[2];

const IMPORTS_QUERY = `ImportDeclaration:has(StringLiteral)`;
const EXPORTS_QUERY = `ExportDeclaration:has(StringLiteral)`;

const tsFile = readFileSync(tsFilePath, {encoding: 'utf8'});

const ast = tsquery.ast(tsFile, tsFilePath);

const importSpecifiers = tsquery(ast, IMPORTS_QUERY)
                             .map(node => node.moduleSpecifier)
                             .map(moduleSpecifier => moduleSpecifier.getText().split(`'`)[1])
                             .filter(imp => !(imp.startsWith('./') && imp.split('/').length === 2))
                             .map(imp => builtins.indexOf(imp) > -1 ? `@types/node` : imp)
                             .join(',');

process.stdout.write(importSpecifiers);
