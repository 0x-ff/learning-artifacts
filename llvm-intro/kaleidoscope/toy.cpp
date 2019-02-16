#include "llvm/ADT/STLExtras.h"
#include <algorithm>
#include <cctype>
#include <cstdio>
#include <cstdlib>
#include <map>
#include <memory>
#include <string>
#include <vector>

#include "llvm/IR/LegacyPassManager.h"

using namespace llvm;
#include "Parser.cpp"

static std::unique_ptr<legacy::FunctionPassManager> TheFPM;

void InitializeModuleAndPassManager(void) {
  TheModule = llvm::make_unique<Module>("JIT", TheContext);

  TheFPM = llvm::make_unique<FunctionPassManager>(TheModule.get());
  TheFPM->add(createInstructionCombiningPass());
  TheFPM->add(createReassociatePass());
  TheFPM->add(createGVNPass());
  TheFPM->add(createCFGSimplificationPass());
  TheFPM->doInitialization();
}

static void HandleDefinition() {
  if (auto FnAST = ParseDefinition()) {
    if (auto *FnIR = FnAST->codegen()) {
      if (FnIR == nullptr) {
        fprintf(stderr, "Can't read function definition:");
      } else {
        fprintf(stderr, "Read function definition:");
        FnIR->print(errs());
        fprintf(stderr, "\n");
      }
    }
  } else {
    getNextToken();
  }
}

static void HandleExtern() {
  if (auto ProtoAST = ParseExtern()) {
    if (auto *FnIR = ProtoAST->codegen()) {
      if (FnIR == nullptr) {
        fprintf(stderr, "Can't read extern definition:");
      } else {
        fprintf(stderr, "Read extern: ");
        FnIR->print(errs());
        fprintf(stderr, "\n");
      }
    }
  } else {
    getNextToken();
  }
}

static void HandleTopLevelExpression() {
  if (auto FnAST = ParseTopLevelExpr()) {
    fprintf(stderr, "Parsed top-level expression:");
    if (auto *FnIR = FnAST->codegen()) {
      fprintf(stderr, "Read top-level expression:");
      FnIR->print(errs());
      fprintf(stderr, "\n");
    }
  } else {
    getNextToken();
  }
}

static void MainLoop() {
  while (true) {
    fprintf(stderr, "ready> ");
    switch (CurTok) {
    case tok_eof:
      return;
    case ';':
      getNextToken();
      break;
    case tok_def:
      HandleDefinition();
      break;
    case tok_extern:
      HandleExtern();
      break;
    default:
      HandleTopLevelExpression();
      break;
    }
  }
}

int main() {
  BinopPrecedence['<'] = 10;
  BinopPrecedence['+'] = 20;
  BinopPrecedence['-'] = 20;
  BinopPrecedence['*'] = 40;

  fprintf(stderr, "ready> ");
  getNextToken();

  TheModule = llvm::make_unique<Module>("my cool jit", TheContext);

  MainLoop();

  TheModule->print(errs(), nullptr);
  return 0;
}