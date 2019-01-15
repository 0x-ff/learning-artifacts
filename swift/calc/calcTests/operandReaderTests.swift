
import XCTest
@testable import calc

class operandReaderTests: XCTestCase {

    override func setUp() {
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }

    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }

    func testSetValue() {
        let reader = OperandReader()
        reader.setValue(1.357)
        XCTAssertEqual(1.357, reader.getValue())
        XCTAssertEqual("1.357", reader.getRawValue())
        reader.clearValue()
        XCTAssertEqual(0.0, reader.getValue())
        XCTAssertEqual("0", reader.getRawValue())
    }
    
    func testAppendDot() {
        let reader = OperandReader()
        
        XCTAssertTrue(reader.appendDot())
        XCTAssertEqual("0.", reader.getRawValue())
        
        XCTAssertTrue(reader.appendDot())
        XCTAssertEqual("0.", reader.getRawValue())
        XCTAssertTrue(reader.appendDigit("1"))
        XCTAssertTrue(reader.appendDot())
        XCTAssertEqual("0.1", reader.getRawValue())
    }
    
    func testAppendDigitFirstNotZero() {
        let reader = OperandReader()
        
        XCTAssertTrue(reader.appendDigit("1"))
        XCTAssertEqual("1", reader.getRawValue())
        XCTAssertTrue(reader.appendDigit("0"))
        XCTAssertEqual("10", reader.getRawValue())
        
        XCTAssertTrue(reader.appendDot())
        XCTAssertEqual("10.", reader.getRawValue())
        XCTAssertTrue(reader.appendDigit("0"))
        XCTAssertEqual("10.0", reader.getRawValue())
        XCTAssertTrue(reader.appendDigit("1"))
        XCTAssertEqual("10.01", reader.getRawValue())
    }
    
    func testAppendDigitFirstZero() {
        let reader = OperandReader()
        
        XCTAssertTrue(reader.appendDigit("0"))
        XCTAssertEqual("0", reader.getRawValue())
        XCTAssertTrue(reader.appendDigit("0"))
        XCTAssertEqual("0", reader.getRawValue())
        
        XCTAssertTrue(reader.appendDot())
        XCTAssertEqual("0.", reader.getRawValue())
        XCTAssertTrue(reader.appendDigit("0"))
        XCTAssertEqual("0.0", reader.getRawValue())
        XCTAssertTrue(reader.appendDigit("1"))
        XCTAssertEqual("0.01", reader.getRawValue())
    }
    
    func testParse() {
        let reader = OperandReader()
        
        XCTAssertEqual(0.0, reader.getValue())
        
        XCTAssertTrue(reader.appendDigit("5"))
        XCTAssertEqual("5", reader.getRawValue())
        XCTAssertTrue(reader.parse())
        XCTAssertEqual(5.0, reader.getValue())
        
        XCTAssertTrue(reader.appendDigit("0"))
        XCTAssertEqual("50", reader.getRawValue())
        XCTAssertTrue(reader.parse())
        XCTAssertEqual(50.0, reader.getValue())
    }
    
    func testParseError() {
        let reader = OperandReader()
        
        XCTAssertTrue(reader.appendDot())
        XCTAssertTrue(reader.appendDigit("a"))
        XCTAssertEqual("0.a", reader.getRawValue())
        
        XCTAssertFalse(reader.parse())
        XCTAssertEqual("error: 0.a", reader.getLastError())
    }
}
