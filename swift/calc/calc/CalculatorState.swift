
import Foundation

// Состояние конечного автомата
class CalculatorState {
    
    // Начальное состояние - когда не введены ни левый ни правый операнды
    let STATE_INITIAL = 0
    
    // Состояние чтения левого операнда
    let STATE_COLLECT_LEFT = 1
    
    // Состояние когда завершили вводить левый операнд
    let STATE_COMPLETE_LEFT = 2
    
    // Состояние чтения правого операнда
    let STATE_COLLECT_RIGHT = 3
    
    // Текущее состояние автомата
    private var state: Int
    
    private var left: OperandReader = OperandReader()
    private var right: OperandReader = OperandReader()
    
    private var binaryOperation: BinaryOperation? = nil
    
    private var displayValue: String = ""
    private var lastError: String = ""
    
    init() {
        state = STATE_INITIAL
        return
    }
    
    // Отладочный вывод в строку всего текущего состояния объекта
    func describe() -> String {
        var description: String
        switch (state) {
            case STATE_INITIAL:
                description = "initial\n"
            case STATE_COLLECT_LEFT:
                description = "collecting left\n"
            case STATE_COMPLETE_LEFT:
                description = "complete left\n"
            case STATE_COLLECT_RIGHT:
                description = "colecting right\n"
            default:
                description = "<unknown state>"
        }
        description += "left op is '" + left.getRawValue() + "'\n"
        description += "right op is '" + right.getRawValue() + "'\n"
        var opName: String = "<nil>"
        if binaryOperation != nil {
            opName = binaryOperation!.name
        }
        description += "operation is " + opName + "\n"
        return description
    }
    
    func getState() -> Int {
        return state
    }
    
    func setStateInitial() {
        left.clearValue()
        right.clearValue()
        displayValue = left.getRawValue()
        state = STATE_INITIAL
    }
    
    func setStateCollectLeft() {
        state = STATE_COLLECT_LEFT
    }
    
    func setStateCompleteLeft() {
        state = STATE_COMPLETE_LEFT
    }
    
    func setStateCollectRight() {
        state = STATE_COLLECT_RIGHT
    }
    
    func setLastError(err: String) {
        lastError = err
    }
    
    func getLastError() -> String {
        return lastError
    }
    
    func getDisplayValue() -> String {
        return displayValue
    }
    
    func appendDot() -> Bool {
        switch state {
        case STATE_INITIAL:
            setStateCollectLeft()
            return appendDotLeft()
            
        case STATE_COLLECT_LEFT:
            return appendDotLeft()
            
        case STATE_COMPLETE_LEFT:
            setStateCollectRight()
            return appendDotRight()
            
        case STATE_COLLECT_RIGHT:
            return appendDotRight()
            
        default:
            lastError = "calculatorState error: \(state)"
            return false
        }
    }
    
    private func appendDotLeft() -> Bool {
        if left.appendDot() {
            displayValue = left.getRawValue()
            return true
        } else {
            lastError = left.getLastError()
            return false
        }
    }
    
    private func appendDotRight() -> Bool {
        if right.appendDot() {
            displayValue = right.getRawValue()
            return true
        } else {
            lastError = right.getLastError()
            return false
        }
    }
    
    func appendDigit(_ digit: Character) -> Bool {
        switch state {
        case STATE_INITIAL:
            setStateCollectLeft()
            return appendDigitLeft(digit)
            
        case STATE_COLLECT_LEFT:
            return appendDigitLeft(digit)
            
        case STATE_COMPLETE_LEFT:
            setStateCollectRight()
            return appendDigitRight(digit)
            
        case STATE_COLLECT_RIGHT:
            return appendDigitRight(digit)
            
        default:
            lastError = "calculatorState error: \(state)"
            return false
        }
    }
    
    private func appendDigitLeft(_ digit: Character) -> Bool {
        if left.appendDigit(digit) {
            displayValue = left.getRawValue()
            return true
        } else {
            lastError = left.getLastError()
            return false
        }
    }
    
    private func appendDigitRight(_ digit: Character) -> Bool {
        if right.appendDigit(digit) {
            displayValue = right.getRawValue()
            return true
        } else {
            lastError = right.getLastError()
            return false
        }
    }
    
    func calculateUnary(_ operation: inout UnaryOperation) -> Bool {
        binaryOperation = nil
        switch state {
        case STATE_INITIAL, STATE_COLLECT_LEFT, STATE_COMPLETE_LEFT:
            // Считаем, что производим операцию с левым операндом
            // В этом состоянии у него должно быть значение 0
            if left.parse() {
                operation.operand = left.getValue()
                if operation.calculate() {
                    left.setValue(operation.result)
                    displayValue = left.getRawValue()
                    return true
                } else {
                    setStateInitial()
                    lastError = operation.lastError
                    return false
                }
            } else {
                lastError = left.getLastError()
                return false
            }
        case STATE_COLLECT_RIGHT:
            // Считаем, что производим операцию с правым операндом
            // Но при этом перезатираем вычисленным значением левый
            // А правый очищаем и переходим в состояние ввода левого
            if right.parse() {
                operation.operand = right.getValue()
                if operation.calculate() {
                    left.setValue(operation.result)
                    setStateCollectLeft()
                    displayValue = left.getRawValue()
                    right.clearValue()
                    return true
                } else {
                    setStateInitial()
                    lastError = operation.lastError
                    return false
                }
            } else {
                lastError = right.getLastError()
                return false
            }
        default:
            lastError = "calculatorState error: \(state)"
            return false
        }
    }
    
    func inputBinary(_ op: BinaryOperation) -> Bool {
        switch (state) {
            case STATE_INITIAL, STATE_COLLECT_LEFT:
                setStateCompleteLeft()
                binaryOperation = op
                return true
            case STATE_COMPLETE_LEFT:
                binaryOperation = op
                return true
            case STATE_COLLECT_RIGHT:
                if binaryOperation == nil {
                    binaryOperation = op
                }
                let result: Bool = calculateBinary()
                binaryOperation = op
                return result
            default:
                lastError = "calculatorState error: \(state)"
                return false
        }
    }
    
    func calculateBinary() -> Bool {
        switch state {
        case STATE_INITIAL, STATE_COLLECT_LEFT, STATE_COMPLETE_LEFT:
            // Считаем, что производим пустую операцию, которая не делает ничего
            return true
            
        case STATE_COLLECT_RIGHT:
            if binaryOperation == nil {
                // Считаем, что производим пустую операцию, которая не делает ничего
                return false
            }
            if left.parse() {
                binaryOperation!.leftOperand = left.getValue()
            } else {
                lastError = left.getLastError()
                return false
            }
            if right.parse() {
                binaryOperation?.rightOperand = right.getValue()
            } else {
                lastError = right.getLastError()
                return false
            }
            if binaryOperation!.calculate() {
                left.setValue(binaryOperation!.result)
                binaryOperation = nil
                displayValue = left.getRawValue()
                right.clearValue()
                setStateCompleteLeft()
                return true
            } else {
                lastError = binaryOperation!.lastError
                setStateInitial()
                return false
            }
        default:
            lastError = "calculatorState error: \(state)"
            return false
        }
    }
}
