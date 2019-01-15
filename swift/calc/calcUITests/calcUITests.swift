
import XCTest

class calcUITests: XCTestCase {
        
    override func setUp() {
        super.setUp()
        
        // Put setup code here. This method is called before the invocation of each test method in the class.
        
        // In UI tests it is usually best to stop immediately when a failure occurs.
        continueAfterFailure = false
        // UI tests must launch the application that they test. Doing this in setup will make sure it happens for each test method.
        XCUIApplication().launch()

        // In UI tests it’s important to set the initial state - such as interface orientation - required for your tests before they run. The setUp method is a good place to do this.
    }
    
    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
        super.tearDown()
    }
    
    func testSmokePlus() {
        let app = XCUIApplication()
        app.buttons["1"].tap()
        app.buttons["6"].tap()
        app.buttons["+"].tap()
        app.buttons["3"].tap()
        app.buttons["="].tap()
        assertResultEqual("19.0")
    }
    
    func testSmokeMinus() {
        let app = XCUIApplication()
        app.buttons["1"].tap()
        app.buttons["6"].tap()
        app.buttons["-"].tap()
        app.buttons["3"].tap()
        app.buttons["="].tap()
        assertResultEqual("13.0")
    }
    
    func testSmokeMul() {
        let app = XCUIApplication()
        app.buttons["2"].tap()
        app.buttons["0"].tap()
        app.buttons["x"].tap()
        app.buttons["3"].tap()
        app.buttons["="].tap()
        assertResultEqual("60.0")
    }
    
    func testSmokeDiv() {
        let app = XCUIApplication()
        app.buttons["4"].tap()
        app.buttons["0"].tap()
        app.buttons["/"].tap()
        app.buttons["4"].tap()
        app.buttons["="].tap()
        assertResultEqual("10.0")
    }
    
    func testSmokeLog() {
        let app = XCUIApplication()
        app.buttons["1"].tap()
        app.buttons["0"].tap()
        app.buttons["0"].tap()
        app.buttons["0"].tap()
        app.buttons["Log"].tap()
        assertResultEqual("3.0")
    }
    
    func testSmokeLn() {
        let app = XCUIApplication()
        app.buttons["2"].tap()
        app.buttons["0"].tap()
        app.buttons["0"].tap()
        app.buttons["0"].tap()
        app.buttons["Ln"].tap()
        assertResultEqual("7.600902459542082")
    }
    
    func testSmokeSqrt() {
        let app = XCUIApplication()
        app.buttons["6"].tap()
        app.buttons["2"].tap()
        app.buttons["5"].tap()
        app.buttons["√"].tap()
        assertResultEqual("25.0")
    }
    
    func assertResultEqual(_ expected: String) {
        let app = XCUIApplication()
        let result: String
        let input = app.textFields.firstMatch
        if input.value != nil {
            result = String(input.value.debugDescription)
        } else {
            result = "unknown"
        }
        XCTAssertEqual("Optional(\(expected))", result)
    }
}
