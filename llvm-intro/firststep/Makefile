LLVM_CONFIG?=/usr/local/opt/llvm/bin/llvm-config
SRC_DIR=?$(PWD)
LDFLAGS+=$(shell $(LLVM_CONFIG) --ldflags)
COMMON_FLAGS=-Wall -Wextra
CXXFLAGS+=$(COMMON_FLAGS) $(shell $(LLVM_CONFIG) --cxxflags)
CPPFLAGS+=$(shell $(LLVM_CONFIG) --cppflags) -I$(SRC_DIR)

ifndef VERBOSE
QUIET:=@
endif

HELLO=helloworld
HELLO_OBJECTS=hello.o

hello: $(HELLO)

%.o: $(SRC_DIR)/%.cpp
	@echo Compiling $*.cpp
	$(QUIET)$(CXX) -c $(CPPFLAGS) $(CXXFLAGS) $@ $^

$(HELLO): $(HELLO_OBJECTS)
	@echo Linking $@
	$(QUIET)$(CXX) -o $@ $(CXXFLAGS) $(LDFLAGS) $^ `$(LLVM_CONFIG) --libs bitreader core support`
