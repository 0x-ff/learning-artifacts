
import Foundation

// Интерфейсный фасад калькулятора, изолирущий
// 1. Cтруктуру классов чтения операндов/состояния конечного автомата
// 2. Логику вычисления результатов
// 3. Обработку исключений
// от UI инструментария ViewController iOS
class Calculator {
    let OPERATION_LOG = 0
    let OPERATION_LN = 1
    let OPERATION_SQRT = 2
    let OPERATION_PLUS = 3
    let OPERATION_MINUS = 4
    let OPERATION_MUL = 5
    let OPERATION_DIV = 6

    private var state = CalculatorState()
    private var displayValue: String = ""
    
    init() {
        return
    }
    
    func getDisplayValue() -> String {
        return displayValue
    }
    
    func getStateDescription() -> String {
        return state.describe()
    }
    
    func appendDot() -> Calculator {
        if state.appendDot() {
            displayValue = state.getDisplayValue()
        } else {
            displayValue = state.getLastError()
        }
        return self
    }
    
    func clearState() -> Calculator {
        state.setStateInitial()
        displayValue = state.getDisplayValue()
        return self
    }
    
    func appendDigit(_ digit: Character) -> Calculator {
        if state.appendDigit(digit)  {
            displayValue = state.getDisplayValue()
        } else {
            displayValue = state.getLastError()
        }
        return self
    }
    
    func executeUnary(_ operation: Int) -> Calculator {
        var op: UnaryOperation
        switch (operation) {
            case OPERATION_LOG:
                op = OperationLog()
            case OPERATION_LN:
                op = OperationLn()
            case OPERATION_SQRT:
                op = OperationSqrt()
            default:
                displayValue = "error unknown operation\(operation)"
                return self
        }
        if state.calculateUnary(&op) {
            displayValue = state.getDisplayValue()
        } else {
            displayValue = state.getLastError()
        }
        return self
    }
    
    func binaryOperation(_ operation: Int) -> Calculator {
        var op: BinaryOperation
        switch (operation) {
            case OPERATION_PLUS:
                op = OperationPlus()
            case OPERATION_MINUS:
                op = OperationMinus()
            case OPERATION_MUL:
                op = OperationMul()
            case OPERATION_DIV:
                op = OperationDiv()
            default:
                displayValue = "error \(operation)"
                return self
        }
        
        if state.inputBinary(op){
            displayValue = state.getDisplayValue()
        } else {
            displayValue = state.getLastError()
        }
        return self
    }
    
    func execute() -> Calculator {
        if state.calculateBinary() {
            displayValue = state.getDisplayValue()
        } else {
            displayValue = state.getLastError()
        }
        return self
    }
}
