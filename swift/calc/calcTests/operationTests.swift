
import XCTest
@testable import calc

class operationTests: XCTestCase {

    func testOperationLog_Zero_Error() {
        let operation = OperationLog()
        operation.operand = 0
        XCTAssertFalse(operation.calculate())
        XCTAssertEqual(0.0, operation.result)
        XCTAssertEqual("error: negative or zero value", operation.lastError)
    }
    
    func testOperationLog_Ok() {
        let operation = OperationLog()
        operation.operand = 100
        XCTAssertTrue(operation.calculate())
        XCTAssertEqual(2.0, operation.result)
        XCTAssertEqual("", operation.lastError)
    }
    
    func testOperationLog_NegativeValue() {
        let operation = OperationLog()
        operation.operand = -25
        XCTAssertFalse(operation.calculate())
        XCTAssertEqual(0.0, operation.result)
        XCTAssertEqual("error: negative or zero value", operation.lastError)
    }
    
    func testOperationLn_Zero_Error() {
        let operation = OperationLn()
        operation.operand = 0
        XCTAssertFalse(operation.calculate())
        XCTAssertEqual(0.0, operation.result)
        XCTAssertEqual("error: negative or zero value", operation.lastError)
    }
    
    func testOperationLn_Ok() {
        let operation = OperationLn()
        operation.operand = M_E
        XCTAssertTrue(operation.calculate())
        XCTAssertEqual(1.0, operation.result)
        XCTAssertEqual("", operation.lastError)
    }
    
    func testOperationLn_NegativeValue() {
        let operation = OperationLn()
        operation.operand = -25
        XCTAssertFalse(operation.calculate())
        XCTAssertEqual(0.0, operation.result)
        XCTAssertEqual("error: negative or zero value", operation.lastError)
    }
    
    func testOperationSqrt_Zero_Ok() {
        let operation = OperationSqrt()
        operation.operand = 0
        XCTAssertTrue(operation.calculate())
        XCTAssertEqual(0.0, operation.result)
        XCTAssertEqual("", operation.lastError)
    }
    
    func testOperationSqrt_Ok() {
        let operation = OperationSqrt()
        operation.operand = 25
        XCTAssertTrue(operation.calculate())
        XCTAssertEqual(5, operation.result)
        XCTAssertEqual("", operation.lastError)
    }
    
    func testOperationSqrt_NegativeValue() {
        let operation = OperationSqrt()
        operation.operand = -25
        XCTAssertFalse(operation.calculate())
        XCTAssertEqual(0.0, operation.result)
        XCTAssertEqual("error: negative value", operation.lastError)
    }
    
    func testOperationPlus_Ok() {
        let operation = OperationPlus()
        operation.leftOperand = 77
        operation.rightOperand = 33
        XCTAssertTrue(operation.calculate())
        XCTAssertEqual(110.0, operation.result)
        XCTAssertEqual("", operation.lastError)
    }
    
    func testOperationMinus_Ok() {
        let operation = OperationMinus()
        operation.leftOperand = 2
        operation.rightOperand = 3
        XCTAssertTrue(operation.calculate())
        XCTAssertEqual(-1.0, operation.result)
        XCTAssertEqual("", operation.lastError)
    }
    
    func testOperationMul_Ok() {
        let operation = OperationMul()
        operation.leftOperand = 7
        operation.rightOperand = 8
        XCTAssertTrue(operation.calculate())
        XCTAssertEqual(56.0, operation.result)
        XCTAssertEqual("", operation.lastError)
    }
    
    func testOperationDiv_Ok() {
        let operation = OperationDiv()
        operation.leftOperand = 22
        operation.rightOperand = 11
        XCTAssertTrue(operation.calculate())
        XCTAssertEqual(2.0, operation.result)
        XCTAssertEqual("", operation.lastError)
    }
    
    func testOperationDiv_DivisionByZero() {
        let operation = OperationDiv()
        operation.leftOperand = 22
        operation.rightOperand = 0
        XCTAssertFalse(operation.calculate())
        XCTAssertEqual(0.0, operation.result)
        XCTAssertEqual("error: division by zero!", operation.lastError)
    }
}
