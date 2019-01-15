
import XCTest
@testable import calc

class calcTests: XCTestCase {
    
    override func setUp() {
        super.setUp()
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }
    
    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
        super.tearDown()
    }
    
    // Минимальная проверка фасада для конечного автомата, если она не работает - все очень и очень плохо!
    // Последовательные проверки тривиальных случаев выполнения каждой операции
    func testSmoke() {
        let calc = Calculator()
        var result = calc
            .appendDigit("1")
            .binaryOperation(calc.OPERATION_PLUS)
            .appendDigit("1")
            .execute()
            .getDisplayValue()
        XCTAssertEqual("2.0", result)
        
        result = calc
            .clearState()
            .appendDigit("2")
            .binaryOperation(calc.OPERATION_MINUS)
            .appendDigit("3")
            .execute()
            .getDisplayValue()
        XCTAssertEqual("-1.0", result)
        
        result = calc
            .clearState()
            .appendDigit("7")
            .binaryOperation(calc.OPERATION_MUL)
            .appendDigit("5")
            .execute()
            .getDisplayValue()
        XCTAssertEqual("35.0", result)
        
        result = calc
            .clearState()
            .appendDigit("6")
            .binaryOperation(calc.OPERATION_DIV)
            .appendDigit("3")
            .execute()
            .getDisplayValue()
        XCTAssertEqual("2.0", result)
        
        result = calc
            .clearState()
            .appendDigit("2")
            .appendDigit("5")
            .executeUnary(calc.OPERATION_SQRT)
            .execute()
            .getDisplayValue()
        XCTAssertEqual("5.0", result)
        
        result = calc
            .clearState()
            .appendDigit("1")
            .appendDigit("0")
            .appendDigit("0")
            .executeUnary(calc.OPERATION_LOG)
            .getDisplayValue()
        XCTAssertEqual("2.0", result)
        
        result = calc
            .clearState()
            .appendDigit("1")
            .appendDigit("0")
            .appendDigit("0")
            .executeUnary(calc.OPERATION_LN)
            .getDisplayValue()
        XCTAssertEqual("4.60517", result.prefix(7))
    }
}
