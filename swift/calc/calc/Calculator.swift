
import Foundation

// Фасад калькулятора, изолирущий логику чтения операндов и вычисления
// результатов от логики UI фреймворка iOS
class Calculator {
    let OPERATION_LOG = 0
    let OPERATION_LN = 1
    let OPERATION_SQRT = 2
    let OPERATION_PLUS = 3
    let OPERATION_MINUS = 4
    let OPERATION_MUL = 5
    let OPERATION_DIV = 6

    private var calculatorState = CalculatorState()
    var displayValue: String = ""
    
    init() {
        return
    }
    
    func getDisplayValue() -> String {
        return displayValue
    }
    
    func appendDot() -> Calculator {
        if calculatorState.appendDot() {
            displayValue = calculatorState.getDisplayValue()
        } else {
            displayValue = calculatorState.getLastError()
        }
        return self
    }
    
    func clearState() -> Calculator {
        calculatorState.setStateInitial()
        return self
    }
    
    func appendDigit(_ digit: Character) -> Calculator {
        if calculatorState.appendDigit(digit)  {
            displayValue = calculatorState.getDisplayValue()
        } else {
            displayValue = calculatorState.getLastError()
        }
        return self
    }
    
    func calculateUnary(_ operation: Int) -> Calculator {
        var op: UnaryOperation
        switch (operation) {
        case OPERATION_LOG:
            op = OperationLog()
            default:
                displayValue = "error \(operation)"
                return self
        }
        if calculatorState.calculateUnary(&op) {
            displayValue = calculatorState.getDisplayValue()
        } else {
            displayValue = calculatorState.getLastError()
        }
        return self
    }
    
    func binaryOperation(_ operation: Int) -> Calculator {
        var op: BinaryOperation
        switch (operation) {
        case OPERATION_PLUS:
            op = OperationPlus()
        default:
            displayValue = "error \(operation)"
            return self
        }
        
        if calculatorState.binaryOperation(op){
            displayValue = calculatorState.getDisplayValue()
        } else {
            displayValue = calculatorState.getLastError()
        }
        return self
    }
    
    func execute() -> Calculator {
        if calculatorState.calculateBinary() {
            displayValue = calculatorState.getDisplayValue()
        } else {
            displayValue = calculatorState.getLastError()
        }
        return self
    }
}
