
import Foundation

// Протокол числовых операций
protocol Operation {
    var result: Double { get }
    var lastError: String { get }
    // Возвращает true если операция успешно выполнена и
    // false в случае ошибки
    func calculate() -> Bool
}

// Унарные операции sqrt, ln, log ...
protocol UnaryOperation: Operation {
    var operand: Double { get set }
}

// Бинарные операции + - * /
protocol BinaryOperation: Operation {
    var leftOperand: Double { get set }
    var rightOperand: Double { get set }
}

// Десятичный логарифм
class OperationLog: UnaryOperation {
    var operand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        if operand > 0 {
            result = log(operand)
            return true
        } else {
            lastError = "error: log(\(operand))"
            return false
        }
    }
}

// Сложение
class OperationPlus: BinaryOperation {
    var leftOperand: Double = 0.0
    var rightOperand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        result = leftOperand + rightOperand
        return true
    }
}
